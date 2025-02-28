package provider

import (
	"context"
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasource.DataSource              = &stackOutputsDataSource{}
	_ datasource.DataSourceWithConfigure = &stackOutputsDataSource{}
)

// NewStackOutputsDataSource is a helper function to simplify the provider implementation.
func NewStackOutputsDataSource() datasource.DataSource {
	return &stackOutputsDataSource{}
}

// stackOutputsDataSource is the data source implementation.
type stackOutputsDataSource struct {
	client *SpaceLiftClient
}

// stackOutputsDataSourceModel maps the data source schema data.
type stackOutputsDataSourceModel struct {
	ID        types.String            `tfsdk:"id"`
	StackID   types.String            `tfsdk:"stack_id"`
	Outputs   types.Map               `tfsdk:"outputs"`
	LastCheck types.String            `tfsdk:"last_check"`
}

// Configure adds the provider configured client to the data source.
func (d *stackOutputsDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*SpaceLiftClient)
	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Data Source Configure Type",
			fmt.Sprintf("Expected *SpaceLiftClient, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)
		return
	}

	d.client = client
}

// Metadata returns the data source type name.
func (d *stackOutputsDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_stack_outputs"
}

// Schema defines the schema for the data source.
func (d *stackOutputsDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Retrieves all outputs from a SpaceLift stack.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description: "Identifier of the data source.",
				Computed:    true,
			},
			"stack_id": schema.StringAttribute{
				Description: "The ID of the SpaceLift stack.",
				Required:    true,
			},
			"outputs": schema.MapAttribute{
				Description: "The outputs of the SpaceLift stack.",
				Computed:    true,
				ElementType: types.StringType,
			},
			"last_check": schema.StringAttribute{
				Description: "The timestamp of the last check.",
				Computed:    true,
			},
		},
	}
}

// Read refreshes the Terraform state with the latest data.
func (d *stackOutputsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var state stackOutputsDataSourceModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Get stack outputs from SpaceLift
	stackID := state.StackID.ValueString()
	outputs, err := d.client.GetStackOutputs(stackID)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Reading SpaceLift Stack Outputs",
			"Could not read stack outputs: "+err.Error(),
		)
		return
	}

	// Create a map of string values for the outputs
	outputMap := make(map[string]attr.Value)
	for _, output := range outputs {
		outputMap[output.ID] = types.StringValue(output.Value)
	}
	
	// Create a Map value from the map of string values
	outputsValue, diags := types.MapValue(types.StringType, outputMap)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	
	state.ID = types.StringValue(stackID)
	state.Outputs = outputsValue
	state.LastCheck = types.StringValue(time.Now().Format(time.RFC3339))

	// Set state
	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
} 