package postman

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	postmanSDK "github.com/jonnydgreen/terraform-provider-postman/client/postman"
)

func resourceEnvironment() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceEnvironmentCreate,
		ReadContext:   resourceEnvironmentRead,
		UpdateContext: resourceEnvironmentUpdate,
		DeleteContext: resourceEnvironmentDelete,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Description: "The environment's name.",
				Required:    true,
				ForceNew:    true,
			},
			"workspace": {
				Type:        schema.TypeString,
				Description: "The workspace's ID.",
				Optional:    true,
				ForceNew:    true,
			},
			"values": {
				Type:        schema.TypeList,
				Description: "Information about the environment's variables",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"key": {
							Type:        schema.TypeString,
							Description: "The variable's name.",
							Required:    true,
						},
						"type": {
							Type:        schema.TypeString,
							Description: "The variable type.",
							Optional:    true,
							Default:     "default",
						},
						"value": {
							Type:        schema.TypeString,
							Description: "The variable's value.",
							Optional:    true,
						},
						"enabled": {
							Type:        schema.TypeBool,
							Description: "If true, the variable is enabled.",
							Optional:    true,
							Default:     true,
						},
					},
				},
			},
		},
	}
}

func resourceEnvironmentCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	c := m.(*postmanSDK.APIClient)

	name := d.Get("name").(string)
	input := postmanSDK.CreateEnvironmentRequest{
		Environment: &postmanSDK.CreateEnvironmentRequestEnvironment{
			Name: name,
		},
	}

	createEnvironment := c.EnvironmentsApi.CreateEnvironment(ctx)
	workspaceID := d.Get("workspace")
	if workspaceID != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Warning,
			Summary:  "Create for workspace",
			Detail:   fmt.Sprintf("Create for workspace %s", workspaceID),
		})
		createEnvironment = createEnvironment.Workspace(workspaceID.(string))
	}
	response, _, err := createEnvironment.CreateEnvironmentRequest(input).Execute()
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(*response.Environment.Id)

	return diags
}

func resourceEnvironmentRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	c := m.(*postmanSDK.APIClient)

	environmentID := d.Id()

	response, _, err := c.EnvironmentsApi.SingleEnvironment(ctx, environmentID).Execute()
	if err != nil {
		return diag.FromErr(err)
	}

	responseEnvironment, isEnvironmentDefined := response.GetEnvironmentOk()
	if responseEnvironment == nil || isEnvironmentDefined != true {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Unable to find environment",
			Detail:   fmt.Sprintf("No environment with ID %s found in Postman API response.", environmentID),
		})
		return diags
	}

	err = setEnvironmentResourceData(d, responseEnvironment)
	if err != nil {
		return diag.FromErr(err)
	}
	return diags
}

func resourceEnvironmentUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	return resourceEnvironmentRead(ctx, d, m)
}

func resourceEnvironmentDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	return diags
}

func setEnvironmentResourceData(d *schema.ResourceData, responseEnvironment *postmanSDK.SingleEnvironment200ResponseEnvironment) error {
	if err := d.Set("id", responseEnvironment.Id); err != nil {
		return fmt.Errorf("Error setting id: %s", err)
	}
	if err := d.Set("name", responseEnvironment.Name); err != nil {
		return fmt.Errorf("Error setting name: %s", err)
	}
	return nil
}
