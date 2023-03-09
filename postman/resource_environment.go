package postman

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/jonnydgreen/terraform-provider-postman/client/postman"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ resource.Resource                = &environmentResource{}
	_ resource.ResourceWithConfigure   = &environmentResource{}
	_ resource.ResourceWithImportState = &environmentResource{}
)

// NewEnvironmentResource is a helper function to simplify the provider implementation.
func NewEnvironmentResource() resource.Resource {
	return &environmentResource{}
}

// environmentResource is the resource implementation.
type environmentResource struct {
	client *postman.APIClient
}

// Metadata returns the resource type name.
func (r *environmentResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_environment"
}

func environmentSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description: "The environment's ID.",
				Computed:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"name": schema.StringAttribute{
				Description: "The environment's name.",
				Required:    true,
			},
			"workspace": schema.StringAttribute{
				Description: "The environment's workspace ID. If not specified, the default workspace is used.",
				Optional:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"created_at": schema.StringAttribute{
				Description: "The date and time at which the environment was created.",
				Computed:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"updated_at": schema.StringAttribute{
				Description: "The date and time at which the environment was last updated.",
				Computed:    true,
			},
			"owner": schema.StringAttribute{
				Description: "The environment owner's ID.",
				Computed:    true,
				Optional:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"is_public": schema.BoolAttribute{
				Description: "If true, the environment is public.",
				Computed:    true,
				Optional:    true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.UseStateForUnknown(),
				},
			},
			// TODO: move to a separate resource
			// "values": {
			// 	Type:        schema.TypeList,
			// 	Description: "Information about the environment's variables",
			// 	Optional:    true,
			// 	Elem: &schema.Resource{
			// 		Schema: map[string]*schema.Schema{
			// 			"key": {
			// 				Type:        schema.TypeString,
			// 				Description: "The variable's name.",
			// 				Required:    true,
			// 			},
			// 			"type": {
			// 				Type:        schema.TypeString,
			// 				Description: "The variable type.",
			// 				Optional:    true,
			// 				Default:     "default",
			// 			},
			// 			"value": {
			// 				Type:        schema.TypeString,
			// 				Description: "The variable's value.",
			// 				Optional:    true,
			// 				Sensitive:   true,
			// 			},
			// 			"enabled": {
			// 				Type:        schema.TypeBool,
			// 				Description: "If true, the variable is enabled.",
			// 				Optional:    true,
			// 				Default:     true,
			// 			},
			// 		},
			// 	},
			// },
		},
	}
}

// Schema defines the schema for the resource.
func (r *environmentResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = environmentSchema()
}

// environmentResourceModel maps the resource schema data.
type environmentResourceModel struct {
	ID        types.String `tfsdk:"id"`
	Name      types.String `tfsdk:"name"`
	Workspace types.String `tfsdk:"workspace"`
	IsPublic  types.Bool   `tfsdk:"is_public"`
	Owner     types.String `tfsdk:"owner"`
	CreatedAt types.String `tfsdk:"created_at"`
	UpdatedAt types.String `tfsdk:"updated_at"`
}

// Configure adds the provider configured client to the resource.
func (r *environmentResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = req.ProviderData.(*postman.APIClient)
}

// Create creates the resource and sets the initial Terraform state.
func (r *environmentResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Retrieve values from plan
	var plan environmentResourceModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Generate API request body from plan
	environmentName, err := expandEnvironmentName(plan.Name)
	if err != nil {
		resp.Diagnostics.AddError("Error parsing environment name", err.Error())
		return
	}
	environmentWorkspace, err := expandEnvironmentWorkspace(plan.Workspace)
	if err != nil {
		resp.Diagnostics.AddError("Error parsing environment workspace", err.Error())
		return
	}
	environment := postman.NewCreateEnvironmentRequestEnvironment(environmentName)

	// Create new environment
	input := postman.NewCreateEnvironmentRequest()
	input.SetEnvironment(*environment)
	createEnvironmentRequest := r.client.EnvironmentsApi.CreateEnvironment(ctx).CreateEnvironmentRequest(*input)
	if environmentWorkspace != nil {
		createEnvironmentRequest.Workspace(*environmentWorkspace)
	}
	response, _, err := createEnvironmentRequest.Execute()
	if err != nil {
		resp.Diagnostics.AddError("Error creating environment", "Could not create environment, unexpected error: "+err.Error())
		return
	}

	environmentID := *response.Environment.Id
	plan.ID = flattenEnvironmentID(environmentID)
	singleEnvironmentResponse, _, err := r.client.EnvironmentsApi.SingleEnvironment(ctx, environmentID).Execute()
	if err != nil {
		resp.Diagnostics.AddError("Error creating environment", "Error finding created environment, unexpected error: "+err.Error())
		return
	}
	responseEnvironment, isEnvironmentDefined := singleEnvironmentResponse.GetEnvironmentOk()
	if responseEnvironment == nil || isEnvironmentDefined != true {
		resp.Diagnostics.AddError("Error creating environment", "Created environment does not exist")
		return
	}

	// Map response body to schema and populate Computed attribute values
	plan.CreatedAt = flattenEnvironmentCreatedAt(responseEnvironment.CreatedAt)
	plan.UpdatedAt = flattenEnvironmentUpdatedAt(responseEnvironment.UpdatedAt)
	plan.IsPublic = flattenEnvironmentIsPublic(responseEnvironment.IsPublic)
	plan.Owner = flattenEnvironmentOwner(responseEnvironment.Owner)

	// Set state to fully populated data
	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

// Read refreshes the Terraform state with the latest data.
func (r *environmentResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	// Get current state
	var state environmentResourceModel
	diags := req.State.Get(ctx, &state)
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
	response, raw, err := r.client.EnvironmentsApi.SingleEnvironment(ctx, environmentID).Execute()
	if err != nil {
		if raw.StatusCode == 404 {
			tflog.Debug(ctx, fmt.Sprintf("[DEBUG] %s for: %s, removing from state file", err, environmentID))
			state.ID = flattenWorkspaceID("")
			return
		}
		resp.Diagnostics.AddError("Error reading environment", "Could not read environment, unexpected error: "+err.Error())
		return
	}

	// TODO: ensure that the environment belongs to this workspace

	// Overwrite with refreshed state
	state.Name = flattenEnvironmentName(response.Environment.Name)
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

// Update updates the resource and sets the updated Terraform state on success.
func (r *environmentResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	// Retrieve values from plan
	var plan environmentResourceModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Generate API request body from plan
	environmentID, err := expandEnvironmentID(plan.ID)
	if err != nil {
		resp.Diagnostics.AddError("Error parsing environment ID", err.Error())
		return
	}
	environmentName, err := expandEnvironmentName(plan.Name)
	if err != nil {
		resp.Diagnostics.AddError("Error parsing environment name", err.Error())
		return
	}
	environment := postman.NewUpdateEnvironmentRequestEnvironment(environmentName)

	// Update environment
	input := postman.NewUpdateEnvironmentRequest()
	input.SetEnvironment(*environment)
	_, _, err = r.client.EnvironmentsApi.UpdateEnvironment(ctx, environmentID).UpdateEnvironmentRequest(*input).Execute()
	if err != nil {
		resp.Diagnostics.AddError("Error updating environment", "Could not update environment, unexpected error: "+err.Error())
		return
	}

	// Get new computed values
	singleEnvironmentResponse, _, err := r.client.EnvironmentsApi.SingleEnvironment(ctx, environmentID).Execute()
	if err != nil {
		resp.Diagnostics.AddError("Error updating environment", "Error finding updated environment, unexpected error: "+err.Error())
		return
	}
	responseEnvironment, isEnvironmentDefined := singleEnvironmentResponse.GetEnvironmentOk()
	if responseEnvironment == nil || isEnvironmentDefined != true {
		resp.Diagnostics.AddError("Error updating environment", "Updated environment does not exist")
		return
	}

	// Map response body to schema and populate Computed attribute values
	plan.CreatedAt = flattenEnvironmentCreatedAt(responseEnvironment.CreatedAt)
	plan.UpdatedAt = flattenEnvironmentUpdatedAt(responseEnvironment.UpdatedAt)
	plan.IsPublic = flattenEnvironmentIsPublic(responseEnvironment.IsPublic)
	plan.Owner = flattenEnvironmentOwner(responseEnvironment.Owner)

	// Set refreshed state
	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

// Delete deletes the resource and removes the Terraform state on success.
func (r *environmentResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	// Retrieve values from state
	var state environmentResourceModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Get environment data from state
	environmentID, err := expandEnvironmentID(state.ID)
	if err != nil {
		resp.Diagnostics.AddError("Error parsing environment ID", err.Error())
		return
	}

	// If the resource doesn't exist, leave as is and delegate to Terraform
	_, response, err := r.client.EnvironmentsApi.SingleEnvironment(context.Background(), environmentID).Execute()
	if response.StatusCode == 404 && err != nil {
		tflog.Debug(ctx, fmt.Sprintf("[DEBUG] %s for: %s, environment already exists, removing from state file", err, environmentID))
		return
	}

	// Delete existing environment
	_, _, err = r.client.EnvironmentsApi.DeleteEnvironment(ctx, environmentID).Execute()
	if err != nil {
		resp.Diagnostics.AddError("Error deleting environment", "Could not delete environment, unexpected error: "+err.Error())
		return
	}
}

func (r *environmentResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	idParts := strings.Split(req.ID, ",")

	if len(idParts) == 1 && idParts[0] != "" {
		resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
		return
	}

	if len(idParts) != 2 || idParts[0] == "" || idParts[1] == "" {
		resp.Diagnostics.AddError(
			"Unexpected Import Identifier",
			fmt.Sprintf("Expected import identifier with format: [workspace,]environment. Got: %q", req.ID),
		)
		return
	}

	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("workspace"), idParts[0])...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), idParts[1])...)
}

func flattenEnvironmentID(v string) basetypes.StringValue {
	return types.StringValue(v)
}

func expandEnvironmentID(v basetypes.StringValue) (string, error) {
	return v.ValueString(), nil
}

func flattenEnvironmentName(v *string) basetypes.StringValue {
	return types.StringValue(*v)
}

func expandEnvironmentName(v basetypes.StringValue) (string, error) {
	return v.ValueString(), nil
}

func expandEnvironmentWorkspace(v basetypes.StringValue) (*string, error) {
	if v.IsNull() {
		return nil, nil
	}
	environmentWorkspace := v.ValueString()
	return &environmentWorkspace, nil
}

func flattenEnvironmentIsPublic(v *bool) basetypes.BoolValue {
	return types.BoolValue(*v)
}

func flattenEnvironmentOwner(v *string) basetypes.StringValue {
	return types.StringValue(*v)
}

func flattenEnvironmentCreatedAt(v *time.Time) basetypes.StringValue {
	return types.StringValue((*v).Format(time.RFC3339))
}

func flattenEnvironmentUpdatedAt(v *time.Time) basetypes.StringValue {
	return types.StringValue((*v).Format(time.RFC3339))
}
