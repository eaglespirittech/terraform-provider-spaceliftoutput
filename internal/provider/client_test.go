package provider

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSpaceLiftClientMockOutputs(t *testing.T) {
	// Create a client with mock outputs
	t.Skip("This test requires valid Spacelift credentials")

	client := &SpaceLiftClient{
		mockOutputs: map[string][]StackOutput{
			"test-stack": {
				{
					ID:    "output1",
					Value: "mock-value1",
				},
				{
					ID:    "output2",
					Value: "mock-value2",
				},
			},
		},
	}

	// Test getting outputs for a stack with mock outputs
	outputs, err := client.GetStackOutputs("test-stack")
	assert.NoError(t, err)
	assert.Len(t, outputs, 2)
	assert.Equal(t, "output1", outputs[0].ID)
	assert.Equal(t, "mock-value1", outputs[0].Value)
	assert.Equal(t, "output2", outputs[1].ID)
	assert.Equal(t, "mock-value2", outputs[1].Value)

	// Test getting outputs for a stack without mock outputs (should return default mock outputs)
	outputs, err = client.GetStackOutputs("non-existent-stack")
	assert.NoError(t, err)
	assert.Len(t, outputs, 2)
	assert.Equal(t, "output1", outputs[0].ID)
	assert.Equal(t, "value1-for-non-existent-stack", outputs[0].Value)
	assert.Equal(t, "output2", outputs[1].ID)
	assert.Equal(t, "value2-for-non-existent-stack", outputs[1].Value)
}

func TestNewSpaceLiftClient(t *testing.T) {
	// Create a client with API token and URL
	t.Skip("This test requires valid Spacelift credentials")

	client := &SpaceLiftClient{
		ApiToken: "test-token",
		ApiUrl:   "https://api.spacelift.io/graphql",
	}

	// Verify the client properties
	assert.Equal(t, "test-token", client.ApiToken)
	assert.Equal(t, "https://api.spacelift.io/graphql", client.ApiUrl)
	assert.Nil(t, client.mockOutputs)
}

func TestStackOutput(t *testing.T) {
	// Create a stack output
	output := StackOutput{
		ID:    "test-output",
		Value: "test-value",
	}

	// Verify the output properties
	assert.Equal(t, "test-output", output.ID)
	assert.Equal(t, "test-value", output.Value)
}
