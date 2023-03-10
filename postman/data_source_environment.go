package postman

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/jonnydgreen/terraform-provider-postman/client/postman"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasource.DataSource              = &environmentDataSource{}
	_ datasource.DataSourceWithConfigure = &environmentDataSource{}
)

// NewEnvironmentDataSource is a helper function to simplify the provider implementation.
func NewEnvironmentDataSource() datasource.DataSource {
	return &environmentDataSource{}
}

// environmentDataSource is the data source implementation.
type environmentDataSource struct {
	client *postman.APIClient
}

// Metadata returns the data source type name.
func (d *environmentDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_environment"
}

// Schema defines the schema for the data source.
func (d *environmentDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "The data source postman_environment fetches a Postman Environment.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description: "The environment's ID.",
				Required:    true,
			},
			"workspace": schema.StringAttribute{
				Description: "The environment's workspace ID. If not specified, the default workspace is used.",
				Optional:    true,
			},
			"name": schema.StringAttribute{
				Description: "The environment's name.",
				Computed:    true,
			},
			"created_at": schema.StringAttribute{
				Description: "The date and time at which the environment was created.",
				Computed:    true,
			},
			"updated_at": schema.StringAttribute{
				Description: "The date and time at which the environment was last updated.",
				Computed:    true,
			},
			"owner": schema.StringAttribute{
				Description: "The environment owner's ID.",
				Computed:    true,
				Optional:    true,
			},
			"is_public": schema.BoolAttribute{
				Description: "If true, the environment is public.",
				Computed:    true,
				Optional:    true,
			},
			"values": schema.ListNestedAttribute{
				Description: "The environment's values. If defined, existing values will be overridden. This can be bypassed through the use of lifecycle.ignore_changes.",
				Optional:    true,
				Computed:    true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"key": schema.StringAttribute{
							Description: "The environment value's key.",
							Computed:    true,
						},
						"value": schema.StringAttribute{
							Description: "The environment value's value.",
							Computed:    true,
							Sensitive:   true,
						},
						"type": schema.StringAttribute{
							Description: "The environment value's key. Valid values: default|secret|any. Default: `default`",
							Computed:    true,
							Optional:    true,
						},
						"enabled": schema.BoolAttribute{
							Description: "If true, the value is enabled. Default: `true`",
							Computed:    true,
							Optional:    true,
						},
					},
				},
			},
		},
	}
}

// environmentDataSourceModel maps the data source schema data.
type environmentDataSourceModel struct {
	ID        types.String                    `tfsdk:"id"`
	Name      types.String                    `tfsdk:"name"`
	Values    []environmentValueResourceModel `tfsdk:"values"`
	Workspace types.String                    `tfsdk:"workspace"`
	IsPublic  types.Bool                      `tfsdk:"is_public"`
	Owner     types.String                    `tfsdk:"owner"`
	CreatedAt types.String                    `tfsdk:"created_at"`
	UpdatedAt types.String                    `tfsdk:"updated_at"`
}

// Configure adds the provider configured client to the data source.
func (r *environmentDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = req.ProviderData.(*postman.APIClient)
}

// Read refreshes the Terraform state with the latest data.
func (d *environmentDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	// Get current state
	var state environmentDataSourceModel
	diags := resp.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Get refreshed environment value from Postman
	environmentID, err := expandEnvironmentID(state.ID)
	if err != nil {
		resp.Diagnostics.AddError("Error parsing environment ID", err.Error())
		return
	}
	response, _, err := d.client.EnvironmentsApi.SingleEnvironment(ctx, environmentID).Execute()
	if err != nil {
		resp.Diagnostics.AddError("Error reading environment", "Could not read environment, unexpected error: "+err.Error())
		return
	}

	// Overwrite with refreshed state
	state.Name = flattenEnvironmentName(response.Environment.Name)
	state.Values = flattenEnvironmentValues(response.Environment.Values)
	state.CreatedAt = flattenEnvironmentCreatedAt(response.Environment.CreatedAt)
	state.UpdatedAt = flattenEnvironmentUpdatedAt(response.Environment.UpdatedAt)
	state.IsPublic = flattenEnvironmentIsPublic(response.Environment.IsPublic)
	state.Owner = flattenEnvironmentOwner(response.Environment.Owner)

	// Set refreshed state
	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}
