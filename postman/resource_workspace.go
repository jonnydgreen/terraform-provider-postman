package postman

import (
	"context"
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/jonnydgreen/terraform-provider-postman/client/postman"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ resource.Resource                = &workspaceResource{}
	_ resource.ResourceWithConfigure   = &workspaceResource{}
	_ resource.ResourceWithImportState = &workspaceResource{}
)

// NewWorkspaceResource is a helper function to simplify the provider implementation.
func NewWorkspaceResource() resource.Resource {
	return &workspaceResource{}
}

// workspaceResource is the resource implementation.
type workspaceResource struct {
	client *postman.APIClient
}

// Metadata returns the resource type name.
func (r *workspaceResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_workspace"
}

func workspaceSchema() schema.Schema {
	return schema.Schema{
		Description: "The resource `postman_workspace` creates a Postman Workspace.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description: "The workspace's ID.",
				Computed:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"name": schema.StringAttribute{
				Description: "The workspace's name.",
				Required:    true,
			},
			"type": schema.StringAttribute{
				Description: "The type of workspace. One of: personal|team",
				Required:    true,
			},
			"description": schema.StringAttribute{
				Description: "The workspace's description.",
				Optional:    true,
			},
			"visibility": schema.StringAttribute{
				Description: "The workspace's visibility. [Visibility](https://learning.postman.com/docs/collaborating-in-postman/using-workspaces/managing-workspaces/#changing-workspace-visibility) determines who can access the workspace.",
				Computed:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"created_by": schema.StringAttribute{
				Description: "The user ID of the user who created the workspace.",
				Computed:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"updated_by": schema.StringAttribute{
				Description: "The user ID of the user who last updated the workspace.",
				Computed:    true,
			},
			"created_at": schema.StringAttribute{
				Description: "The date and time at which the workspace was created.",
				Computed:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"updated_at": schema.StringAttribute{
				Description: "The date and time at which the workspace was last updated.",
				Computed:    true,
			},
			"collections": schema.ListNestedAttribute{
				Description: "The workspace's Collections.",
				Computed:    true,
				PlanModifiers: []planmodifier.List{
					listplanmodifier.UseStateForUnknown(),
				},
				NestedObject: schema.NestedAttributeObject{
					PlanModifiers: []planmodifier.Object{
						objectplanmodifier.UseStateForUnknown(),
					},
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							Description: "The ID of the workspace Collection.",
							Computed:    true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
						},
						"name": schema.StringAttribute{
							Description: "The name of the workspace Collection.",
							Computed:    true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
						},
						"uid": schema.StringAttribute{
							Description: "The UID of the workspace Collection.",
							Computed:    true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
						},
					},
				},
			},
			"environments": schema.ListNestedAttribute{
				Description: "The Workspace's Environments.",
				Computed:    true,
				PlanModifiers: []planmodifier.List{
					listplanmodifier.UseStateForUnknown(),
				},
				NestedObject: schema.NestedAttributeObject{
					PlanModifiers: []planmodifier.Object{
						objectplanmodifier.UseStateForUnknown(),
					},
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							Description: "The ID of the Workspace Environment.",
							Computed:    true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
						},
						"name": schema.StringAttribute{
							Description: "The name of the Workspace Environment.",
							Computed:    true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
						},
						"uid": schema.StringAttribute{
							Description: "The UID of the Workspace Environment.",
							Computed:    true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
						},
					},
				},
			},
			"mocks": schema.ListNestedAttribute{
				Description: "The Workspace's Mocks.",
				Computed:    true,
				PlanModifiers: []planmodifier.List{
					listplanmodifier.UseStateForUnknown(),
				},
				NestedObject: schema.NestedAttributeObject{
					PlanModifiers: []planmodifier.Object{
						objectplanmodifier.UseStateForUnknown(),
					},
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							Description: "The ID of the Workspace Mock.",
							Computed:    true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
						},
						"name": schema.StringAttribute{
							Description: "The name of the Workspace Mock.",
							Computed:    true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
						},
						"uid": schema.StringAttribute{
							Description: "The UID of the Workspace Mock.",
							Computed:    true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
						},
					},
				},
			},
			"monitors": schema.ListNestedAttribute{
				Description: "The Workspace's Monitors.",
				Computed:    true,
				PlanModifiers: []planmodifier.List{
					listplanmodifier.UseStateForUnknown(),
				},
				NestedObject: schema.NestedAttributeObject{
					PlanModifiers: []planmodifier.Object{
						objectplanmodifier.UseStateForUnknown(),
					},
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							Description: "The ID of the Workspace Monitor.",
							Computed:    true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
						},
						"name": schema.StringAttribute{
							Description: "The name of the Workspace Monitor.",
							Computed:    true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
						},
						"uid": schema.StringAttribute{
							Description: "The UID of the Workspace Monitor.",
							Computed:    true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
						},
					},
				},
			},
			"apis": schema.ListNestedAttribute{
				Description: "The Workspace's APIs.",
				Computed:    true,
				PlanModifiers: []planmodifier.List{
					listplanmodifier.UseStateForUnknown(),
				},
				NestedObject: schema.NestedAttributeObject{
					PlanModifiers: []planmodifier.Object{
						objectplanmodifier.UseStateForUnknown(),
					},
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							Description: "The ID of the Workspace API.",
							Computed:    true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
						},
						"name": schema.StringAttribute{
							Description: "The name of the Workspace API.",
							Computed:    true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
						},
						"uid": schema.StringAttribute{
							Description: "The UID of the Workspace API.",
							Computed:    true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
						},
					},
				},
			},
		},
	}
}

// Schema defines the schema for the resource.
func (r *workspaceResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = workspaceSchema()
}

// workspaceResourceModel maps the resource schema data.
type workspaceResourceModel struct {
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

// workspaceCollectionModel maps Workspace Collection data.
type workspaceCollectionModel struct {
	ID   types.String `tfsdk:"id"`
	Name types.String `tfsdk:"name"`
	UID  types.String `tfsdk:"uid"`
}

// workspaceEnvironmentModel maps Workspace Environment data.
type workspaceEnvironmentModel struct {
	ID   types.String `tfsdk:"id"`
	Name types.String `tfsdk:"name"`
	UID  types.String `tfsdk:"uid"`
}

// workspaceMockModel maps Workspace Mock data.
type workspaceMockModel struct {
	ID   types.String `tfsdk:"id"`
	Name types.String `tfsdk:"name"`
	UID  types.String `tfsdk:"uid"`
}

// workspaceMonitorModel maps Workspace Monitor data.
type workspaceMonitorModel struct {
	ID   types.String `tfsdk:"id"`
	Name types.String `tfsdk:"name"`
	UID  types.String `tfsdk:"uid"`
}

// workspaceApiModel maps Workspace Api data.
type workspaceApiModel struct {
	ID   types.String `tfsdk:"id"`
	Name types.String `tfsdk:"name"`
	UID  types.String `tfsdk:"uid"`
}

// Configure adds the provider configured client to the resource.
func (r *workspaceResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = req.ProviderData.(*postman.APIClient)
}

// Create creates the resource and sets the initial Terraform state.
func (r *workspaceResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Retrieve values from plan
	var plan workspaceResourceModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Generate API request body from plan
	workspaceName, err := expandWorkspaceName(plan.Name)
	if err != nil {
		resp.Diagnostics.AddError("Error parsing workspace name", err.Error())
		return
	}
	workspaceType, err := expandWorkspaceType(plan.Type)
	if err != nil {
		resp.Diagnostics.AddError("Error parsing workspace type", err.Error())
		return
	}
	workspaceDescription, err := expandWorkspaceDescription(plan.Description)
	if err != nil {
		resp.Diagnostics.AddError("Error parsing workspace description", err.Error())
		return
	}
	workspace := postman.NewCreateWorkspaceRequestWorkspace(workspaceName, workspaceType)
	if workspaceDescription != nil {
		workspace.SetDescription(*workspaceDescription)
	}

	// Create new workspace
	input := postman.NewCreateWorkspaceRequest()
	input.SetWorkspace(*workspace)
	response, _, err := r.client.WorkspacesApi.CreateWorkspace(ctx).CreateWorkspaceRequest(*input).Execute()
	if err != nil {
		resp.Diagnostics.AddError("Error creating workspace", "Could not create workspace, unexpected error: "+err.Error())
		return
	}

	// Populate Computed attribute values
	workspaceID := *response.Workspace.Id
	plan.ID = flattenWorkspaceID(workspaceID)
	singleWorkspaceResponse, _, err := r.client.WorkspacesApi.SingleWorkspace(ctx, workspaceID).Execute()
	if err != nil {
		resp.Diagnostics.AddError("Error creating workspace", "Error finding created workspace, unexpected error: "+err.Error())
		return
	}
	responseWorkspace, isWorkspaceDefined := singleWorkspaceResponse.GetWorkspaceOk()
	if responseWorkspace == nil || isWorkspaceDefined != true {
		resp.Diagnostics.AddError("Error creating workspace", "Created workspace does not exist")
		return
	}

	// Map response body to schema and populate Computed attribute values
	plan.Description = flattenWorkspaceDescription(responseWorkspace.Description)
	plan.Collections, diags = flattenWorkspaceCollections(ctx, responseWorkspace.Collections)
	resp.Diagnostics.Append(diags...)
	plan.Environments, diags = flattenWorkspaceEnvironments(ctx, responseWorkspace.Environments)
	resp.Diagnostics.Append(diags...)
	plan.Mocks, diags = flattenWorkspaceMocks(ctx, responseWorkspace.Mocks)
	resp.Diagnostics.Append(diags...)
	plan.Monitors, diags = flattenWorkspaceMonitors(ctx, responseWorkspace.Monitors)
	resp.Diagnostics.Append(diags...)
	plan.Apis, diags = flattenWorkspaceApis(ctx, responseWorkspace.Apis)
	resp.Diagnostics.Append(diags...)
	plan.CreatedAt = flattenWorkspaceCreatedAt(responseWorkspace.CreatedAt)
	plan.CreatedBy = flattenWorkspaceCreatedBy(responseWorkspace.CreatedBy)
	plan.UpdatedAt = flattenWorkspaceUpdatedAt(responseWorkspace.UpdatedAt)
	plan.UpdatedBy = flattenWorkspaceUpdatedBy(responseWorkspace.UpdatedBy)
	plan.Visibility = flattenWorkspaceVisibility(responseWorkspace.Visibility)

	// Set state to fully populated data
	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

// Read refreshes the Terraform state with the latest data.
func (r *workspaceResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	// Get current state
	var state workspaceResourceModel
	diags := req.State.Get(ctx, &state)
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
	response, raw, err := r.client.WorkspacesApi.SingleWorkspace(ctx, workspaceID).Execute()
	if err != nil {
		if raw.StatusCode == 404 {
			tflog.Debug(ctx, fmt.Sprintf("[DEBUG] %s for: %s, removing from state file", err, workspaceID))
			state.ID = flattenWorkspaceID("")
			return
		}
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

// Update updates the resource and sets the updated Terraform state on success.
func (r *workspaceResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	// Retrieve values from plan
	var plan workspaceResourceModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Generate API request body from plan
	workspaceID, err := expandWorkspaceID(plan.ID)
	if err != nil {
		resp.Diagnostics.AddError("Error parsing workspace ID", err.Error())
		return
	}
	workspaceName, err := expandWorkspaceName(plan.Name)
	if err != nil {
		resp.Diagnostics.AddError("Error parsing workspace name", err.Error())
		return
	}
	workspaceType, err := expandWorkspaceType(plan.Type)
	if err != nil {
		resp.Diagnostics.AddError("Error parsing workspace type", err.Error())
		return
	}
	workspaceDescription, err := expandWorkspaceDescription(plan.Description)
	if err != nil {
		resp.Diagnostics.AddError("Error parsing workspace description", err.Error())
		return
	}
	workspace := postman.NewUpdateWorkspaceRequestWorkspace()
	workspace.SetName(workspaceName)
	workspace.SetType(workspaceType)
	if workspaceDescription != nil {
		workspace.SetDescription(*workspaceDescription)
	} else {
		workspace.SetDescription("")
	}
	updateWorkspaceRequest := postman.NewUpdateWorkspaceRequest()
	updateWorkspaceRequest.SetWorkspace(*workspace)
	_, _, err = r.client.WorkspacesApi.UpdateWorkspace(ctx, workspaceID).UpdateWorkspaceRequest(*updateWorkspaceRequest).Execute()
	if err != nil {
		resp.Diagnostics.AddError("Error updating workspace", "Could not update workspace, unexpected error: "+err.Error())
		return
	}

	singleWorkspaceResponse, _, err := r.client.WorkspacesApi.SingleWorkspace(ctx, workspaceID).Execute()
	if err != nil {
		resp.Diagnostics.AddError("Error updating workspace", "Error finding updated workspace, unexpected error: "+err.Error())
		return
	}
	responseWorkspace, isWorkspaceDefined := singleWorkspaceResponse.GetWorkspaceOk()
	if responseWorkspace == nil || isWorkspaceDefined != true {
		resp.Diagnostics.AddError("Error updating workspace", "Updated workspace does not exist")
		return
	}

	// Map response body to schema and populate Computed attribute values
	plan.Description = flattenWorkspaceDescription(responseWorkspace.Description)
	plan.Collections, diags = flattenWorkspaceCollections(ctx, responseWorkspace.Collections)
	resp.Diagnostics.Append(diags...)
	plan.Environments, diags = flattenWorkspaceEnvironments(ctx, responseWorkspace.Environments)
	resp.Diagnostics.Append(diags...)
	plan.Mocks, diags = flattenWorkspaceMocks(ctx, responseWorkspace.Mocks)
	resp.Diagnostics.Append(diags...)
	plan.Monitors, diags = flattenWorkspaceMonitors(ctx, responseWorkspace.Monitors)
	resp.Diagnostics.Append(diags...)
	plan.Apis, diags = flattenWorkspaceApis(ctx, responseWorkspace.Apis)
	resp.Diagnostics.Append(diags...)
	plan.CreatedAt = flattenWorkspaceCreatedAt(responseWorkspace.CreatedAt)
	plan.CreatedBy = flattenWorkspaceCreatedBy(responseWorkspace.CreatedBy)
	plan.UpdatedAt = flattenWorkspaceUpdatedAt(responseWorkspace.UpdatedAt)
	plan.UpdatedBy = flattenWorkspaceUpdatedBy(responseWorkspace.UpdatedBy)
	plan.Visibility = flattenWorkspaceVisibility(responseWorkspace.Visibility)

	// Set refreshed state
	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

// Delete deletes the resource and removes the Terraform state on success.
func (r *workspaceResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	// Retrieve values from state
	var state workspaceResourceModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Get workspace data from state
	workspaceID, err := expandWorkspaceID(state.ID)
	if err != nil {
		resp.Diagnostics.AddError("Error parsing workspace ID", err.Error())
		return
	}

	// If the resource doesn't exist, leave as is and delegate to Terraform
	_, response, err := r.client.WorkspacesApi.SingleWorkspace(context.Background(), workspaceID).Execute()
	if response.StatusCode == 404 && err != nil {
		tflog.Debug(ctx, fmt.Sprintf("[DEBUG] %s for: %s, workspace already exists, removing from state file", err, workspaceID))
		return
	}

	// Delete existing workspace
	_, _, err = r.client.WorkspacesApi.DeleteWorkspace(ctx, workspaceID).Execute()
	if err != nil {
		resp.Diagnostics.AddError("Error deleting workspace", "Could not delete workspace, unexpected error: "+err.Error())
		return
	}
}

func (r *workspaceResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	// Retrieve import ID and save to id attribute
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

// Maps
func flattenWorkspaceCollections(ctx context.Context, collectionItems []postman.SingleWorkspace200ResponseWorkspaceCollectionsInner) (types.List, diag.Diagnostics) {
	attrTypes := workspaceSchema().Attributes["collections"].GetType().(types.ListType).ElemType.(types.ObjectType).AttrTypes
	if collectionItems != nil {
		cis := make([]workspaceCollectionModel, len(collectionItems), len(collectionItems))
		for i, collectionItem := range collectionItems {
			ci := workspaceCollectionModel{
				ID:   types.StringValue(*collectionItem.Id),
				Name: types.StringValue(*collectionItem.Name),
				UID:  types.StringValue(*collectionItem.Uid),
			}
			cis[i] = ci
		}
		return types.ListValueFrom(ctx, types.ObjectType{AttrTypes: attrTypes}, cis)
	}
	return types.ListValueFrom(ctx, types.ObjectType{
		AttrTypes: attrTypes,
	}, make([]workspaceCollectionModel, 0))
}

func flattenWorkspaceEnvironments(ctx context.Context, environmentItems []postman.SingleWorkspace200ResponseWorkspaceEnvironmentsInner) (types.List, diag.Diagnostics) {
	attrTypes := workspaceSchema().Attributes["environments"].GetType().(types.ListType).ElemType.(types.ObjectType).AttrTypes
	if environmentItems != nil {
		eis := make([]workspaceEnvironmentModel, len(environmentItems), len(environmentItems))
		for i, environmentItem := range environmentItems {
			ei := workspaceEnvironmentModel{
				ID:   types.StringValue(*environmentItem.Id),
				Name: types.StringValue(*environmentItem.Name),
				UID:  types.StringValue(*environmentItem.Uid),
			}
			eis[i] = ei
		}
		return types.ListValueFrom(ctx, types.ObjectType{AttrTypes: attrTypes}, eis)
	}
	return types.ListValueFrom(ctx, types.ObjectType{
		AttrTypes: attrTypes,
	}, make([]workspaceEnvironmentModel, 0))
}

func flattenWorkspaceMocks(ctx context.Context, mockItems []postman.SingleWorkspace200ResponseWorkspaceMocksInner) (types.List, diag.Diagnostics) {
	attrTypes := workspaceSchema().Attributes["mocks"].GetType().(types.ListType).ElemType.(types.ObjectType).AttrTypes
	if mockItems != nil {
		mis := make([]workspaceMockModel, len(mockItems), len(mockItems))
		for i, mockItem := range mockItems {
			mi := workspaceMockModel{
				ID:   types.StringValue(*mockItem.Id),
				Name: types.StringValue(*mockItem.Name),
				UID:  types.StringValue(*mockItem.Uid),
			}
			mis[i] = mi
		}
		return types.ListValueFrom(ctx, types.ObjectType{AttrTypes: attrTypes}, mis)
	}
	return types.ListValueFrom(ctx, types.ObjectType{
		AttrTypes: attrTypes,
	}, make([]workspaceMockModel, 0))
}

func flattenWorkspaceMonitors(ctx context.Context, monitorItems []postman.SingleWorkspace200ResponseWorkspaceMonitorsInner) (types.List, diag.Diagnostics) {
	attrTypes := workspaceSchema().Attributes["monitors"].GetType().(types.ListType).ElemType.(types.ObjectType).AttrTypes
	if monitorItems != nil {
		mis := make([]workspaceMonitorModel, len(monitorItems), len(monitorItems))
		for i, monitorItem := range monitorItems {
			mi := workspaceMonitorModel{
				ID:   types.StringValue(*monitorItem.Id),
				Name: types.StringValue(*monitorItem.Name),
				UID:  types.StringValue(*monitorItem.Uid),
			}
			mis[i] = mi
		}
		return types.ListValueFrom(ctx, types.ObjectType{AttrTypes: attrTypes}, mis)
	}
	return types.ListValueFrom(ctx, types.ObjectType{
		AttrTypes: attrTypes,
	}, make([]workspaceMonitorModel, 0))
}

func flattenWorkspaceApis(ctx context.Context, apiItems []postman.SingleWorkspace200ResponseWorkspaceApisInner) (types.List, diag.Diagnostics) {
	attrTypes := workspaceSchema().Attributes["apis"].GetType().(types.ListType).ElemType.(types.ObjectType).AttrTypes
	if apiItems != nil {
		mis := make([]workspaceApiModel, len(apiItems), len(apiItems))
		for i, apiItem := range apiItems {
			mi := workspaceApiModel{
				ID:   types.StringValue(*apiItem.Id),
				Name: types.StringValue(*apiItem.Name),
				UID:  types.StringValue(*apiItem.Uid),
			}
			mis[i] = mi
		}
		return types.ListValueFrom(ctx, types.ObjectType{AttrTypes: attrTypes}, mis)
	}
	return types.ListValueFrom(ctx, types.ObjectType{
		AttrTypes: attrTypes,
	}, make([]workspaceApiModel, 0))
}

func expandWorkspaceID(v basetypes.StringValue) (string, error) {
	return v.ValueString(), nil
}

func flattenWorkspaceID(v string) basetypes.StringValue {
	return types.StringValue(v)
}

func expandWorkspaceName(v basetypes.StringValue) (string, error) {
	return v.ValueString(), nil
}

func flattenWorkspaceName(v *string) basetypes.StringValue {
	return types.StringValue(*v)
}

func expandWorkspaceType(v basetypes.StringValue) (string, error) {
	return v.ValueString(), nil
}

func flattenWorkspaceType(v *string) basetypes.StringValue {
	return types.StringValue(*v)
}

func expandWorkspaceDescription(v basetypes.StringValue) (*string, error) {
	if v.IsNull() {
		return nil, nil
	}
	workspaceDescription := v.ValueString()
	return &workspaceDescription, nil
}

func flattenWorkspaceDescription(v *string) basetypes.StringValue {
	if v == nil {
		return types.StringNull()
	}
	stringValue := *v
	if stringValue == "" {
		return types.StringNull()
	}
	return types.StringValue(stringValue)
}

func flattenWorkspaceVisibility(v *string) basetypes.StringValue {
	return types.StringValue(*v)
}

func flattenWorkspaceCreatedBy(v *string) basetypes.StringValue {
	return types.StringValue(*v)
}

func flattenWorkspaceUpdatedBy(v *string) basetypes.StringValue {
	return types.StringValue(*v)
}

func flattenWorkspaceCreatedAt(v *time.Time) basetypes.StringValue {
	return types.StringValue((*v).Format(time.RFC3339))
}

func flattenWorkspaceUpdatedAt(v *time.Time) basetypes.StringValue {
	return types.StringValue((*v).Format(time.RFC3339))
}
