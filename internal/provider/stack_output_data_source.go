package provider

import (
	"context"
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasource.DataSource              = &stackOutputDataSource{}
	_ datasource.DataSourceWithConfigure = &stackOutputDataSource{}
)

// NewStackOutputDataSource is a helper function to simplify the provider implementation.
func NewStackOutputDataSource() datasource.DataSource {
	return &stackOutputDataSource{}
}

// stackOutputDataSource is the data source implementation.
type stackOutputDataSource struct {
	client *SpaceLiftClient
}

// stackOutputDataSourceModel maps the data source schema data.
type stackOutputDataSourceModel struct {
	ID         types.String `tfsdk:"id"`
	StackID    types.String `tfsdk:"stack_id"`
	OutputName types.String `tfsdk:"output_name"`
	Value      types.String `tfsdk:"value"`
	LastCheck  types.String `tfsdk:"last_check"`
}

// Configure adds the provider configured client to the data source.
func (d *stackOutputDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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
func (d *stackOutputDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_stack_output"
}

// Schema defines the schema for the data source.
func (d *stackOutputDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Retrieves a single output from a SpaceLift stack.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description: "Identifier of the data source.",
				Computed:    true,
			},
			"stack_id": schema.StringAttribute{
				Description: "The ID of the SpaceLift stack.",
				Required:    true,
			},
			"output_name": schema.StringAttribute{
				Description: "The name of the output to retrieve.",
				Required:    true,
			},
			"value": schema.StringAttribute{
				Description: "The value of the specified output.",
				Computed:    true,
			},
			"last_check": schema.StringAttribute{
				Description: "The timestamp of the last check.",
				Computed:    true,
			},
		},
	}
}

// Read refreshes the Terraform state with the latest data.
func (d *stackOutputDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var state stackOutputDataSourceModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Get stack outputs from SpaceLift
	stackID := state.StackID.ValueString()
	outputName := state.OutputName.ValueString()
	outputs, err := d.client.GetStackOutputs(stackID)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Reading SpaceLift Stack Outputs",
			"Could not read stack outputs: "+err.Error(),
		)
		return
	}

	// Find the specific output
	var outputValue string
	found := false
	for _, output := range outputs {
		if output.ID == outputName {
			outputValue = output.Value
			found = true
			break
		}
	}

	if !found {
		resp.Diagnostics.AddError(
			"Output Not Found",
			fmt.Sprintf("Output with name '%s' not found in stack '%s'", outputName, stackID),
		)
		return
	}

	// Update state with the data
	state.ID = types.StringValue(stackID + ":" + outputName)
	state.Value = types.StringValue(outputValue)
	state.LastCheck = types.StringValue(time.Now().Format(time.RFC3339))

	// Set state
	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
} 