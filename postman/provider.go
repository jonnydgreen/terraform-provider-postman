package postman

import (
	"context"
	"os"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/jonnydgreen/terraform-provider-postman/client/postman"
)

// Ensure the implementation satisfies the expected interfaces
var (
	_ provider.Provider = &postmanProvider{}
)

// New is a helper function to simplify provider server and testing implementation.
func New(version string) func() provider.Provider {
	return func() provider.Provider {
		return &postmanProvider{}
	}
}

// postmanProvider is the provider implementation.
type postmanProvider struct{}

// Metadata returns the provider type name.
func (p *postmanProvider) Metadata(_ context.Context, _ provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "postman"
}

// Schema defines the provider-level schema for configuration data.
func (p *postmanProvider) Schema(_ context.Context, _ provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"api_key": schema.StringAttribute{
				Description: "API Key for the Postman API. May also be provided via POSTMAN_API_KEY environment variable.",
				Optional:    true,
				Sensitive:   true,
			},
		},
	}
}

// Configure prepares a Postman API client for data sources and resources.
func (p *postmanProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
	// Retrieve provider data from configuration
	var config postmanProviderModel
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// If practitioner provided a configuration value for any of the
	// attributes, it must be a known value.
	if config.ApiKey.IsUnknown() {
		resp.Diagnostics.AddAttributeError(
			path.Root("api_key"),
			"Unknown Postman API Key",
			"The provider cannot create the Postman API client as there is an unknown configuration value for the Postman API Key. "+
				"Either target apply the source of the value first, set the value statically in the configuration, or use the POSTMAN_API_KEY environment variable.",
		)
	}

	if resp.Diagnostics.HasError() {
		return
	}

	// Default values to environment variables, but override
	// with Terraform configuration value if set.
	apiKey := os.Getenv("POSTMAN_API_KEY")

	if !config.ApiKey.IsNull() {
		apiKey = config.ApiKey.ValueString()
	}

	// If any of the expected configurations are missing, return
	// errors with provider-specific guidance.
	if apiKey == "" {
		resp.Diagnostics.AddAttributeError(
			path.Root("api_key"),
			"Missing Postman API Key",
			"The provider cannot create the Postman API client as there is a missing or empty value for the Postman API password. "+
				"Set the api_key value in the configuration or use the POSTMAN_API_KEY environment variable. "+
				"If either is already set, ensure the value is not empty.",
		)
	}

	if resp.Diagnostics.HasError() {
		return
	}

	// Create a new Postman client using the configuration values
	configuration := postman.NewConfiguration()
	configuration.AddDefaultHeader("x-api-key", apiKey)
	client := postman.NewAPIClient(configuration)

	// Make the Postman client available during DataSource and Resource
	// type Configure methods.
	resp.DataSourceData = client
	resp.ResourceData = client
}

// DataSources defines the data sources implemented in the provider.
func (p *postmanProvider) DataSources(_ context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{
		NewWorkspaceDataSource,
		NewEnvironmentDataSource,
	}
}

// Resources defines the resources implemented in the provider.
func (p *postmanProvider) Resources(_ context.Context) []func() resource.Resource {
	return []func() resource.Resource{
		NewWorkspaceResource,
		NewEnvironmentResource,
	}
}

// postmanProviderModel maps provider schema data to a Go type.
type postmanProviderModel struct {
	ApiKey types.String `tfsdk:"api_key"`
}
