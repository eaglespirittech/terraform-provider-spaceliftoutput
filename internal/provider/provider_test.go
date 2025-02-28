package provider

import (
	"context"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
)

// testAccProtoV6ProviderFactories are used to instantiate a provider during
// acceptance testing. The factory function will be invoked for every Terraform
// CLI command executed to create a provider server to which the CLI can
// reattach.
var testAccProtoV6ProviderFactories = map[string]func() (tfprotov6.ProviderServer, error){
	"spaceliftoutput": providerserver.NewProtocol6WithError(New("test", WithMockClient)()),
}

func testAccPreCheck(t *testing.T) {
	// You can add code here to run prior to any test case execution, for example assertions
	// about the appropriate environment variables being set are common to see in a pre-check
	// function.

	// Set environment variables for testing
	os.Setenv("SPACELIFT_API_TOKEN", "test-token")
	os.Setenv("SPACELIFT_API_URL", "https://example.com/api")
}

// WithMockClient is a provider option that configures the provider to use a mock client.
func WithMockClient(p *SpaceLiftOutputProvider) {
	p.CreateClient = func(ctx context.Context, apiToken, apiUrl string) (*SpaceLiftClient, error) {
		// Create a mock client that returns predefined outputs
		client := &SpaceLiftClient{
			ApiToken: apiToken,
			ApiUrl:   apiUrl,
			mockOutputs: map[string][]StackOutput{
				"test-stack-id": {
					{
						ID:    "output1",
						Value: "value1-for-test-stack-id",
					},
					{
						ID:    "output2",
						Value: "value2-for-test-stack-id",
					},
				},
				"updated-stack-id": {
					{
						ID:    "output1",
						Value: "value1-for-updated-stack-id",
					},
					{
						ID:    "output2",
						Value: "value2-for-updated-stack-id",
					},
				},
			},
		}

		// Return the mock client
		return client, nil
	}
}
