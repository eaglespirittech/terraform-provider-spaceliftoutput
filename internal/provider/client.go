package provider

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"context"

	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// SpaceLiftClient is the client used to communicate with the SpaceLift API.
type SpaceLiftClient struct {
	ApiToken string
	ApiUrl   string
	// For testing purposes
	mockOutputs map[string][]StackOutput
	ctx         context.Context
}

// GraphQLRequest represents a GraphQL request.
type GraphQLRequest struct {
	Query     string                 `json:"query"`
	Variables map[string]interface{} `json:"variables,omitempty"`
}

// GraphQLResponse represents a GraphQL response.
type GraphQLResponse struct {
	Data   map[string]interface{} `json:"data,omitempty"`
	Errors []GraphQLError         `json:"errors,omitempty"`
}

// GraphQLError represents a GraphQL error.
type GraphQLError struct {
	Message string `json:"message"`
}

// StackOutput represents a stack output.
type StackOutput struct {
	ID    string `json:"id"`
	Value string `json:"value"`
}

// GetStackOutputs retrieves the outputs for a stack.
func (c *SpaceLiftClient) GetStackOutputs(stackID string) ([]StackOutput, error) {
	tflog.Debug(c.ctx, "Getting stack outputs", map[string]interface{}{
		"stack_id": stackID,
	})

	// For testing purposes
	if c.mockOutputs != nil {
		tflog.Debug(c.ctx, "Using mock outputs", map[string]interface{}{
			"stack_id": stackID,
		})
		if outputs, ok := c.mockOutputs[stackID]; ok {
			return outputs, nil
		}
		// If the stack ID is not found in the mock outputs, return default mock outputs
		return []StackOutput{
			{
				ID:    "output1",
				Value: fmt.Sprintf("value1-for-%s", stackID),
			},
			{
				ID:    "output2",
				Value: fmt.Sprintf("value2-for-%s", stackID),
			},
		}, nil
	}

	query := `
		query getStackOutputs($id: ID!) {
			stack(id: $id) {
				outputs {
					id
					value
				}
			}
		}
	`

	variables := map[string]interface{}{
		"id": stackID,
	}

	request := GraphQLRequest{
		Query:     query,
		Variables: variables,
	}

	requestBody, err := json.Marshal(request)
	if err != nil {
		tflog.Error(c.ctx, "Failed to marshal request", map[string]interface{}{
			"error": err.Error(),
		})
		return nil, fmt.Errorf("error marshalling request: %w", err)
	}

	req, err := http.NewRequest("POST", c.ApiUrl, bytes.NewBuffer(requestBody))
	if err != nil {
		tflog.Error(c.ctx, "Failed to create request", map[string]interface{}{
			"error": err.Error(),
			"url":   c.ApiUrl,
		})
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+c.ApiToken)

	tflog.Debug(c.ctx, "Sending request to SpaceLift API", map[string]interface{}{
		"url": c.ApiUrl,
	})

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		tflog.Error(c.ctx, "Failed to make request", map[string]interface{}{
			"error": err.Error(),
		})
		return nil, fmt.Errorf("error making request: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		tflog.Error(c.ctx, "Failed to read response body", map[string]interface{}{
			"error": err.Error(),
		})
		return nil, fmt.Errorf("error reading response body: %w", err)
	}

	var graphQLResponse GraphQLResponse
	err = json.Unmarshal(body, &graphQLResponse)
	if err != nil {
		tflog.Error(c.ctx, "Failed to unmarshal response", map[string]interface{}{
			"error": err.Error(),
			"body":  string(body),
		})
		return nil, fmt.Errorf("error unmarshalling response: %w", err)
	}

	if len(graphQLResponse.Errors) > 0 {
		tflog.Error(c.ctx, "GraphQL error in response", map[string]interface{}{
			"error": graphQLResponse.Errors[0].Message,
		})
		return nil, fmt.Errorf("GraphQL error: %s", graphQLResponse.Errors[0].Message)
	}

	// Extract the stack outputs from the response
	stackData, ok := graphQLResponse.Data["stack"].(map[string]interface{})
	if !ok {
		tflog.Error(c.ctx, "Invalid response format", map[string]interface{}{
			"error": "stack data not found",
			"data":  graphQLResponse.Data,
		})
		return nil, fmt.Errorf("invalid response format: stack data not found")
	}

	outputsData, ok := stackData["outputs"].([]interface{})
	if !ok {
		tflog.Error(c.ctx, "Invalid response format", map[string]interface{}{
			"error":     "outputs data not found",
			"stackData": stackData,
		})
		return nil, fmt.Errorf("invalid response format: outputs data not found")
	}

	var outputs []StackOutput
	for _, outputData := range outputsData {
		outputMap, ok := outputData.(map[string]interface{})
		if !ok {
			tflog.Error(c.ctx, "Invalid output format", nil)
			return nil, fmt.Errorf("invalid output format")
		}

		id, ok := outputMap["id"].(string)
		if !ok {
			tflog.Error(c.ctx, "Invalid output id format", map[string]interface{}{
				"output": outputMap,
			})
			return nil, fmt.Errorf("invalid output id format")
		}

		value, ok := outputMap["value"].(string)
		if !ok {
			tflog.Error(c.ctx, "Invalid output value format", map[string]interface{}{
				"output": outputMap,
			})
			return nil, fmt.Errorf("invalid output value format")
		}

		outputs = append(outputs, StackOutput{
			ID:    id,
			Value: value,
		})
	}

	tflog.Debug(c.ctx, "Successfully retrieved stack outputs", map[string]interface{}{
		"stack_id":     stackID,
		"output_count": len(outputs),
	})

	return outputs, nil
}
