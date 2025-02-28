package provider

import (
	"context"
	"os"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ provider.Provider = &SpaceLiftOutputProvider{}
)

// SpaceLiftOutputProvider is the provider implementation.
type SpaceLiftOutputProvider struct {
	// version is set to the provider version on release.
	version string
	// CreateClient is a function that creates a SpaceLiftClient.
	// This can be overridden for testing.
	CreateClient func(ctx context.Context, apiToken, apiUrl string) (*SpaceLiftClient, error)
}

// SpaceLiftOutputProviderModel describes the provider data model.
type SpaceLiftOutputProviderModel struct {
	ApiToken   types.String `tfsdk:"api_token"`
	ApiUrl     types.String `tfsdk:"api_url"`
	AccountName types.String `tfsdk:"account_name"`
}

// ProviderOption is a function that configures a provider.
type ProviderOption func(*SpaceLiftOutputProvider)

// New creates a new provider instance.
func New(version string, opts ...ProviderOption) func() provider.Provider {
	return func() provider.Provider {
		p := &SpaceLiftOutputProvider{
			version: version,
			CreateClient: func(ctx context.Context, apiToken, apiUrl string) (*SpaceLiftClient, error) {
				return &SpaceLiftClient{
					ApiToken: apiToken,
					ApiUrl:   apiUrl,
				}, nil
			},
		}
		
		// Apply options
		for _, opt := range opts {
			opt(p)
		}
		
		return p
	}
}

// Metadata returns the provider type name.
func (p *SpaceLiftOutputProvider) Metadata(_ context.Context, _ provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "spaceliftoutput"
	resp.Version = p.version
}

// Schema defines the provider-level schema for configuration data.
func (p *SpaceLiftOutputProvider) Schema(_ context.Context, _ provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Interact with SpaceLift Stack Outputs.",
		Attributes: map[string]schema.Attribute{
			"api_token": schema.StringAttribute{
				Description: "The SpaceLift API token. Can also be set with the SPACELIFT_API_TOKEN environment variable.",
				Optional:    true,
				Sensitive:   true,
			},
			"api_url": schema.StringAttribute{
				Description: "The SpaceLift API URL. If not specified, it will be constructed using the account_name.",
				Optional:    true,
			},
			"account_name": schema.StringAttribute{
				Description: "Your account name in Spacelift. Used to construct the API URL if api_url is not specified. Can also be set with the spacelift_account_name environment variable.",
				Optional:    true,
			},
		},
	}
}

// Configure prepares a SpaceLift API client for data sources and resources.
func (p *SpaceLiftOutputProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
	// Retrieve provider data from configuration
	var config SpaceLiftOutputProviderModel
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// If practitioner provided a configuration value for any of the
	// attributes, it must be a known value.
	if config.ApiToken.IsUnknown() {
		resp.Diagnostics.AddAttributeError(
			path.Root("api_token"),
			"Unknown SpaceLift API Token",
			"The provider cannot create the SpaceLift API client as there is an unknown configuration value for the SpaceLift API token. "+
				"Either target apply the source of the value first, set the value statically in the configuration, or use the SPACELIFT_API_TOKEN environment variable.",
		)
	}

	if config.ApiUrl.IsUnknown() {
		resp.Diagnostics.AddAttributeError(
			path.Root("api_url"),
			"Unknown SpaceLift API URL",
			"The provider cannot create the SpaceLift API client as there is an unknown configuration value for the SpaceLift API URL. "+
				"Either target apply the source of the value first, set the value statically in the configuration, or use the default value.",
		)
	}

	if config.AccountName.IsUnknown() {
		resp.Diagnostics.AddAttributeError(
			path.Root("account_name"),
			"Unknown Account Name",
			"The provider cannot create the SpaceLift API client as there is an unknown configuration value for the account name. "+
				"Either target apply the source of the value first, set the value statically in the configuration, or use the default value.",
		)
	}

	if resp.Diagnostics.HasError() {
		return
	}

	// Default values to environment variables, but override
	// with Terraform configuration value if set.
	apiToken := os.Getenv("SPACELIFT_API_TOKEN")
	var apiUrl string
	accountName := os.Getenv("spacelift_account_name")
	
	// If account name is not set in environment variable, use default
	if accountName == "" {
		accountName = "eaglespirittech"
	}

	if !config.ApiToken.IsNull() {
		apiToken = config.ApiToken.ValueString()
	}

	if !config.AccountName.IsNull() {
		accountName = config.AccountName.ValueString()
	}

	if !config.ApiUrl.IsNull() {
		apiUrl = config.ApiUrl.ValueString()
	} else {
		// Construct the API URL using the account name
		apiUrl = "https://" + accountName + ".app.spacelift.io/graphql"
	}

	// If any of the expected configurations are missing, return
	// errors with provider-specific guidance.
	if apiToken == "" {
		resp.Diagnostics.AddAttributeError(
			path.Root("api_token"),
			"Missing SpaceLift API Token",
			"The provider cannot create the SpaceLift API client as there is a missing or empty value for the SpaceLift API token. "+
				"Set the api_token value in the configuration or use the SPACELIFT_API_TOKEN environment variable. "+
				"If either is already set, ensure the value is not empty.",
		)
	}

	if resp.Diagnostics.HasError() {
		return
	}

	// Create a new SpaceLift client using the configuration values
	client, err := p.CreateClient(ctx, apiToken, apiUrl)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Create SpaceLift API Client",
			"An unexpected error occurred when creating the SpaceLift API client. "+
				"If the error is not clear, please contact the provider developers.\n\n"+
				"SpaceLift Client Error: "+err.Error(),
		)
		return
	}

	// Make the SpaceLift client available during DataSource and Resource
	// type Configure methods.
	resp.DataSourceData = client
	resp.ResourceData = client
}

// DataSources defines the data sources implemented in the provider.
func (p *SpaceLiftOutputProvider) DataSources(_ context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{
		NewStackOutputsDataSource,
		NewStackOutputDataSource,
	}
}

// Resources defines the resources implemented in the provider.
func (p *SpaceLiftOutputProvider) Resources(_ context.Context) []func() resource.Resource {
	return []func() resource.Resource{}
} 