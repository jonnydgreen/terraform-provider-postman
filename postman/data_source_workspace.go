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
	_ datasource.DataSource              = &workspaceDataSource{}
	_ datasource.DataSourceWithConfigure = &workspaceDataSource{}
)

// NewWorkspaceDataSource is a helper function to simplify the provider implementation.
func NewWorkspaceDataSource() datasource.DataSource {
	return &workspaceDataSource{}
}

// workspaceDataSource is the data source implementation.
type workspaceDataSource struct {
	client *postman.APIClient
}

// Metadata returns the data source type name.
func (d *workspaceDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_workspace"
}

// Schema defines the schema for the data source.
func (d *workspaceDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "The data source postman_workspace fetches a Postman Workspace.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description: "The workspace's ID.",
				Required:    true,
			},
			"name": schema.StringAttribute{
				Description: "The workspace's name.",
				Computed:    true,
			},
			"type": schema.StringAttribute{
				Description: "The type of workspace. One of: personal|team",
				Computed:    true,
			},
			"description": schema.StringAttribute{
				Description: "The workspace's description.",
				Computed:    true,
				Optional:    true,
			},
			"visibility": schema.StringAttribute{
				Description: "The workspace's visibility. [Visibility](https://learning.postman.com/docs/collaborating-in-postman/using-workspaces/managing-workspaces/#changing-workspace-visibility) determines who can access the workspace.",
				Computed:    true,
			},
			"created_by": schema.StringAttribute{
				Description: "The user ID of the user who created the workspace.",
				Computed:    true,
			},
			"updated_by": schema.StringAttribute{
				Description: "The user ID of the user who last updated the workspace.",
				Computed:    true,
			},
			"created_at": schema.StringAttribute{
				Description: "The date and time at which the workspace was created.",
				Computed:    true,
			},
			"updated_at": schema.StringAttribute{
				Description: "The date and time at which the workspace was last updated.",
				Computed:    true,
			},
			"collections": schema.ListNestedAttribute{
				Description: "The workspace's Collections.",
				Computed:    true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							Description: "The ID of the workspace Collection.",
							Computed:    true,
						},
						"name": schema.StringAttribute{
							Description: "The name of the workspace Collection.",
							Computed:    true,
						},
						"uid": schema.StringAttribute{
							Description: "The UID of the workspace Collection.",
							Computed:    true,
						},
					},
				},
			},
			"environments": schema.ListNestedAttribute{
				Description: "The Workspace's Environments.",
				Computed:    true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							Description: "The ID of the Workspace Environment.",
							Computed:    true,
						},
						"name": schema.StringAttribute{
							Description: "The name of the Workspace Environment.",
							Computed:    true,
						},
						"uid": schema.StringAttribute{
							Description: "The UID of the Workspace Environment.",
							Computed:    true,
						},
					},
				},
			},
			"mocks": schema.ListNestedAttribute{
				Description: "The Workspace's Mocks.",
				Computed:    true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							Description: "The ID of the Workspace Mock.",
							Computed:    true,
						},
						"name": schema.StringAttribute{
							Description: "The name of the Workspace Mock.",
							Computed:    true,
						},
						"uid": schema.StringAttribute{
							Description: "The UID of the Workspace Mock.",
							Computed:    true,
						},
					},
				},
			},
			"monitors": schema.ListNestedAttribute{
				Description: "The Workspace's Monitors.",
				Computed:    true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							Description: "The ID of the Workspace Monitor.",
							Computed:    true,
						},
						"name": schema.StringAttribute{
							Description: "The name of the Workspace Monitor.",
							Computed:    true,
						},
						"uid": schema.StringAttribute{
							Description: "The UID of the Workspace Monitor.",
							Computed:    true,
						},
					},
				},
			},
			"apis": schema.ListNestedAttribute{
				Description: "The Workspace's APIs.",
				Computed:    true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							Description: "The ID of the Workspace API.",
							Computed:    true,
						},
						"name": schema.StringAttribute{
							Description: "The name of the Workspace API.",
							Computed:    true,
						},
						"uid": schema.StringAttribute{
							Description: "The UID of the Workspace API.",
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

// workspaceDataSourceModel maps the data source schema data.
type workspaceDataSourceModel struct {
	ID           types.String `tfsdk:"id"`
	Name         types.String `tfsdk:"name"`
	Type         types.String `tfsdk:"type"`
	Description  types.String `tfsdk:"description"`
	Visibility   types.String `tfsdk:"visibility"`
	CreatedBy    types.String `tfsdk:"created_by"`
	UpdatedBy    types.String `tfsdk:"updated_by"`
	CreatedAt    types.String `tfsdk:"created_at"`
	UpdatedAt    types.String `tfsdk:"updated_at"`
	Collections  types.List   `tfsdk:"collections"`
	Environments types.List   `tfsdk:"environments"`
	Mocks        types.List   `tfsdk:"mocks"`
	Monitors     types.List   `tfsdk:"monitors"`
	Apis         types.List   `tfsdk:"apis"`
}

// Configure adds the provider configured client to the data source.
func (r *workspaceDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = req.ProviderData.(*postman.APIClient)
}

// Read refreshes the Terraform state with the latest data.
func (d *workspaceDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	// Get current state
	var state workspaceDataSourceModel
	diags := resp.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Get refreshed workspace value from Postman
	workspaceID, err := expandWorkspaceID(state.ID)
	if err != nil {
		resp.Diagnostics.AddError("Error parsing workspace ID", err.Error())
		return
	}
	response, _, err := d.client.WorkspacesApi.SingleWorkspace(ctx, workspaceID).Execute()
	if err != nil {
		resp.Diagnostics.AddError("Error reading workspace", "Could not read workspace, unexpected error: "+err.Error())
		return
	}

	// Overwrite with refreshed state
	state.Name = flattenWorkspaceName(response.Workspace.Name)
	state.Type = flattenWorkspaceType(response.Workspace.Type)
	state.Description = flattenWorkspaceDescription(response.Workspace.Description)
	state.Collections, diags = flattenWorkspaceCollections(ctx, response.Workspace.Collections)
	resp.Diagnostics.Append(diags...)
	state.Environments, diags = flattenWorkspaceEnvironments(ctx, response.Workspace.Environments)
	resp.Diagnostics.Append(diags...)
	state.Mocks, diags = flattenWorkspaceMocks(ctx, response.Workspace.Mocks)
	resp.Diagnostics.Append(diags...)
	state.Monitors, diags = flattenWorkspaceMonitors(ctx, response.Workspace.Monitors)
	resp.Diagnostics.Append(diags...)
	state.Apis, diags = flattenWorkspaceApis(ctx, response.Workspace.Apis)
	resp.Diagnostics.Append(diags...)
	state.CreatedAt = flattenWorkspaceCreatedAt(response.Workspace.CreatedAt)
	state.CreatedBy = flattenWorkspaceCreatedBy(response.Workspace.CreatedBy)
	state.UpdatedAt = flattenWorkspaceUpdatedAt(response.Workspace.UpdatedAt)
	state.UpdatedBy = flattenWorkspaceUpdatedBy(response.Workspace.UpdatedBy)
	state.Visibility = flattenWorkspaceVisibility(response.Workspace.Visibility)

	// Set refreshed state
	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}
