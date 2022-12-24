package postman

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/jonnydgreen/terraform-provider-postman/client/postman"
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
							Sensitive:   true,
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

	c := m.(*postman.APIClient)

	name := d.Get("name").(string)
	input := postman.CreateEnvironmentRequest{
		Environment: &postman.CreateEnvironmentRequestEnvironment{
			Name:   name,
			Values: mapToEnvironmentValueItemsResponse(d.Get("values")),
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

	c := m.(*postman.APIClient)

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

func setEnvironmentResourceData(d *schema.ResourceData, responseEnvironment *postman.SingleEnvironment200ResponseEnvironment) error {
	if err := d.Set("name", responseEnvironment.Name); err != nil {
		return fmt.Errorf("Error setting name: %s", err)
	}
	if err := d.Set("values", mapToProviderEnvironmentValueItems(&responseEnvironment.Values)); err != nil {
		return fmt.Errorf("Error setting name: %s", err)
	}
	return nil
}

func mapToProviderEnvironmentValueItems(valueItems *[]postman.CreateEnvironmentRequestEnvironmentValuesInner) []interface{} {
	if valueItems != nil {
		vis := make([]interface{}, len(*valueItems), len(*valueItems))
		for i, valueItem := range *valueItems {
			vi := make(map[string]interface{})
			vi["key"] = valueItem.Key
			vi["value"] = valueItem.Value
			vi["type"] = valueItem.Type
			vi["enabled"] = valueItem.Enabled
			vis[i] = vi
		}
		return vis
	}
	return make([]interface{}, 0)
}

func mapToEnvironmentValueItemsResponse(rawValueItems interface{}) []postman.CreateEnvironmentRequestEnvironmentValuesInner {
	if rawValueItems != nil {
		valueItems := rawValueItems.([]interface{})
		vis := make([]postman.CreateEnvironmentRequestEnvironmentValuesInner, len(valueItems), len(valueItems))
		for idx, valueItem := range valueItems {
			i := valueItem.(map[string]interface{})
			key := i["key"].(string)
			valueType := i["type"].(string)
			value := i["value"].(string)
			enabled := i["enabled"].(bool)
			vi := postman.CreateEnvironmentRequestEnvironmentValuesInner{
				Key:     &key,
				Type:    &valueType,
				Value:   &value,
				Enabled: &enabled,
			}
			vis[idx] = vi
		}
		return vis
	}
	return make([]postman.CreateEnvironmentRequestEnvironmentValuesInner, 0)
}
