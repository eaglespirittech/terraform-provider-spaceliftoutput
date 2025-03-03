package provider

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSpaceLiftClientLive(t *testing.T) {
	// Skip this test when running in CI or automated test environments
	t.Skip("This test requires valid Spacelift credentials")
	client := &SpaceLiftClient{
		ApiToken: "test-token", // Replace with your actual API token
		ApiUrl:   "https://account.app.spacelift.io/graphql",
	}

	// Test getting outputs for a specific stack
	stackID := "test-stack"
	outputs, err := client.GetStackOutputs(stackID)

	// Assert no error occurred
	assert.NoError(t, err)

	// Assert we got some outputs
	assert.NotNil(t, outputs)
	assert.Greater(t, len(outputs), 0)

	// Look for specific output
	var foundOutput *StackOutput
	for _, output := range outputs {
		if output.ID == "test-output" {
			foundOutput = &output
			break
		}
	}

	// Assert we found the specific output we were looking for
	assert.NotNil(t, foundOutput, "Output 'test-output' not found")
	if foundOutput != nil {
		t.Logf("Found output %s with value: %s", foundOutput.ID, foundOutput.Value)
	}
}
