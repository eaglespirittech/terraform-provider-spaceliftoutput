package provider

import (
	"context"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/provider"
)

// TestProviderMetadata tests the provider metadata.
func TestProviderMetadata(t *testing.T) {
	ctx := context.Background()
	p := &SpaceLiftOutputProvider{
		version: "test",
	}

	metadataResp := &provider.MetadataResponse{}
	p.Metadata(ctx, provider.MetadataRequest{}, metadataResp)

	if metadataResp.TypeName != "spaceliftoutput" {
		t.Errorf("Expected provider type name to be 'spaceliftoutput', got '%s'", metadataResp.TypeName)
	}

	if metadataResp.Version != "test" {
		t.Errorf("Expected provider version to be 'test', got '%s'", metadataResp.Version)
	}
}

// TestProviderSchema tests the provider schema.
func TestProviderSchema(t *testing.T) {
	ctx := context.Background()
	p := &SpaceLiftOutputProvider{
		version: "test",
	}

	schemaResp := &provider.SchemaResponse{}
	p.Schema(ctx, provider.SchemaRequest{}, schemaResp)

	if schemaResp.Schema.Description != "Interact with SpaceLift Stack Outputs." {
		t.Errorf("Expected provider description to be 'Interact with SpaceLift Stack Outputs.', got '%s'", schemaResp.Schema.Description)
	}

	if _, ok := schemaResp.Schema.Attributes["api_token"]; !ok {
		t.Errorf("Expected provider schema to have 'api_token' attribute")
	}

	if _, ok := schemaResp.Schema.Attributes["api_url"]; !ok {
		t.Errorf("Expected provider schema to have 'api_url' attribute")
	}

	if _, ok := schemaResp.Schema.Attributes["account_name"]; !ok {
		t.Errorf("Expected provider schema to have 'account_name' attribute")
	}
}

// TestProviderDataSources tests the provider data sources.
func TestProviderDataSources(t *testing.T) {
	ctx := context.Background()
	p := &SpaceLiftOutputProvider{
		version: "test",
	}

	dataSources := p.DataSources(ctx)

	if len(dataSources) != 2 {
		t.Errorf("Expected provider to have 2 data sources, got %d", len(dataSources))
	}
}

// TestProviderResources tests the provider resources.
func TestProviderResources(t *testing.T) {
	ctx := context.Background()
	p := &SpaceLiftOutputProvider{
		version: "test",
	}

	resources := p.Resources(ctx)

	if len(resources) != 0 {
		t.Errorf("Expected provider to have 0 resources, got %d", len(resources))
	}
}

// TestProviderFactory tests the provider factory.
func TestProviderFactory(t *testing.T) {
	factory := New("test")
	p := factory()

	if p == nil {
		t.Errorf("Expected provider factory to return a non-nil provider")
	}
}
