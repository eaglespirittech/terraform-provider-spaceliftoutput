package provider

import (
	"context"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/stretchr/testify/assert"
)

// TestStackOutputDataSourceMetadata tests the data source metadata.
func TestStackOutputDataSourceMetadata(t *testing.T) {
	ctx := context.Background()

	// Create an instance of the data source
	ds := &stackOutputDataSource{}

	// Create a metadata request and response
	req := datasource.MetadataRequest{
		ProviderTypeName: "spaceliftoutput",
	}
	resp := &datasource.MetadataResponse{}

	// Call the Metadata method
	ds.Metadata(ctx, req, resp)

	// Assert that the type name is correct
	assert.Equal(t, "spaceliftoutput_stack_output", resp.TypeName)
}

// TestStackOutputDataSourceSchema tests the data source schema.
func TestStackOutputDataSourceSchema(t *testing.T) {
	ctx := context.Background()

	// Create an instance of the data source
	ds := &stackOutputDataSource{}

	// Create a schema request and response
	req := datasource.SchemaRequest{}
	resp := &datasource.SchemaResponse{}

	// Call the Schema method
	ds.Schema(ctx, req, resp)

	// Assert that the schema has the expected attributes
	assert.NotNil(t, resp.Schema.Attributes["stack_id"])
	assert.NotNil(t, resp.Schema.Attributes["output_name"])
	assert.NotNil(t, resp.Schema.Attributes["value"])
	assert.NotNil(t, resp.Schema.Attributes["last_check"])
}

// TestStackOutputsDataSourceMetadata tests the data source metadata.
func TestStackOutputsDataSourceMetadata(t *testing.T) {
	ctx := context.Background()

	// Create an instance of the data source
	ds := &stackOutputsDataSource{}

	// Create a metadata request and response
	req := datasource.MetadataRequest{
		ProviderTypeName: "spaceliftoutput",
	}
	resp := &datasource.MetadataResponse{}

	// Call the Metadata method
	ds.Metadata(ctx, req, resp)

	// Assert that the type name is correct
	assert.Equal(t, "spaceliftoutput_stack_outputs", resp.TypeName)
}

// TestStackOutputsDataSourceSchema tests the data source schema.
func TestStackOutputsDataSourceSchema(t *testing.T) {
	ctx := context.Background()

	// Create an instance of the data source
	ds := &stackOutputsDataSource{}

	// Create a schema request and response
	req := datasource.SchemaRequest{}
	resp := &datasource.SchemaResponse{}

	// Call the Schema method
	ds.Schema(ctx, req, resp)

	// Assert that the schema has the expected attributes
	assert.NotNil(t, resp.Schema.Attributes["stack_id"])
	assert.NotNil(t, resp.Schema.Attributes["outputs"])
	assert.NotNil(t, resp.Schema.Attributes["last_check"])
}
