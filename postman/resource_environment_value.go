package postman

import (
	"context"
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/jonnydgreen/terraform-provider-postman/client/postman"
	"golang.org/x/exp/slices"
)

func resourceEnvironmentValue() *schema.Resource {
	return &schema.Resource{
		// CreateContext: resourceEnvironmentValueCreate,
		ReadContext: resourceEnvironmentValueRead,
		// UpdateContext: resourceEnvironmentValueUpdate,
		// DeleteContext: resourceEnvironmentValueDelete,
		Schema: map[string]*schema.Schema{
			"environment": {
				Type:        schema.TypeString,
				Description: "The value's environment.",
				Required:    true,
			},
			"key": {
				Type:        schema.TypeString,
				Description: "The value's name.",
				Required:    true,
			},
			"type": {
				Type:        schema.TypeString,
				Description: "The value's type. One of: default|secret|any",
				Optional:    true,
				Default:     "default",
			},
			"value": {
				Type:        schema.TypeString,
				Description: "The value's value.",
				Optional:    true,
				Sensitive:   true,
			},
			"enabled": {
				Type:        schema.TypeBool,
				Description: "If true, the value is enabled.",
				Optional:    true,
				Default:     true,
			},
		},
	}
}

func resourceEnvironmentValueRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	environmentID := d.Get("environment").(string)
	environmentValueKey := d.Get("key").(string)

	c := m.(*postman.APIClient)

	response, raw, err := c.EnvironmentsApi.SingleEnvironment(ctx, environmentID).Execute()
	if err != nil {
		if raw.StatusCode == 404 {
			log.Printf("[DEBUG] %s for: %s, removing from state file", err, d.Id())
			d.SetId("")
			return diags
		}
		return diag.FromErr(err)
	}
	idx := slices.IndexFunc(response.Environment.Values, func(c postman.CreateEnvironmentRequestEnvironmentValuesInner) bool {
		return *c.Key == environmentValueKey
	})
	if idx == -1 {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Unable to find environment value",
			Detail:   fmt.Sprintf("No environment value with key %s found in Postman API response.", environmentValueKey),
		})
		return diags
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

// func resourceEnvironmentValueCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
// 	// Warning or errors can be collected in a slice type
// 	var diags diag.Diagnostics

// 	c := m.(*postman.APIClient)

// 	environmentName := d.Get("name").(string)
// 	input := postman.CreateEnvironmentRequest{
// 		Environment: &postman.CreateEnvironmentRequestEnvironment{
// 			Name:   environmentName,
// 			Values: mapToEnvironmentValueItemsResponse(d.Get("values")),
// 		},
// 	}

// 	createEnvironment := c.EnvironmentsApi.CreateEnvironmentValue(ctx)
// 	workspaceID := d.Get("workspace")
// 	if workspaceID != nil {
// 		diags = append(diags, diag.Diagnostic{
// 			Severity: diag.Warning,
// 			Summary:  "Create for workspace",
// 			Detail:   fmt.Sprintf("Create for workspace %s", workspaceID),
// 		})
// 		createEnvironment = createEnvironment.Workspace(workspaceID.(string))
// 	}
// 	response, _, err := createEnvironment.CreateEnvironmentRequest(input).Execute()
// 	if err != nil {
// 		return diag.FromErr(err)
// 	}

// 	environmentID := *response.Environment.Id
// 	d.SetId(environmentID)

// 	return diags
// }

// func resourceEnvironmentValueUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
// 	// // Warning or errors can be collected in a slice type
// 	// var diags diag.Diagnostics

// 	// c := m.(*postman.APIClient)

// 	// environmentID := d.Id()
// 	// environmentName := d.Get("name").(string)

// 	// updateWorkspaceRequest := postman.UpdateWorkspaceRequest{
// 	// 	Workspace: &postman.UpdateWorkspaceRequestWorkspace{
// 	// 		Name:        &environmentName,
// 	// 		Type:        &workspaceType,
// 	// 		Description: &workspaceDescription,
// 	// 	},
// 	// }
// 	// response, _, err := c.WorkspacesApi.UpdateWorkspace(ctx, environmentID).UpdateWorkspaceRequest(updateWorkspaceRequest).Execute()
// 	// if err != nil {
// 	// 	return diag.FromErr(err)
// 	// }

// 	// responseWorkspace, isWorkspaceDefined := response.GetWorkspaceOk()
// 	// if responseWorkspace == nil || isWorkspaceDefined != true {
// 	// 	diags = append(diags, diag.Diagnostic{
// 	// 		Severity: diag.Error,
// 	// 		Summary:  "Unable to find workspace",
// 	// 		Detail:   fmt.Sprintf("No workspace with ID %s found in Postman API response.", environmentID),
// 	// 	})
// 	// 	return diags
// 	// }

// 	return resourceEnvironmentValueRead(ctx, d, m)
// }

// func resourceEnvironmentValueDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
// 	// Warning or errors can be collected in a slice type
// 	var diags diag.Diagnostics

// 	return diags
// }

// func setEnvironmentResourceData(d *schema.ResourceData, responseEnvironment *postman.SingleEnvironment200ResponseEnvironment) error {
// 	if err := d.Set("name", responseEnvironment.Name); err != nil {
// 		return fmt.Errorf("Error setting name: %s", err)
// 	}
// 	if err := d.Set("values", mapToProviderEnvironmentValueItems(&responseEnvironment.Values)); err != nil {
// 		return fmt.Errorf("Error setting name: %s", err)
// 	}
// 	return nil
// }

// func mapToProviderEnvironmentValueItems(valueItems *[]postman.CreateEnvironmentRequestEnvironmentValuesInner) []interface{} {
// 	if valueItems != nil {
// 		vis := make([]interface{}, len(*valueItems), len(*valueItems))
// 		for i, valueItem := range *valueItems {
// 			vi := make(map[string]interface{})
// 			vi["key"] = valueItem.Key
// 			vi["value"] = valueItem.Value
// 			vi["type"] = valueItem.Type
// 			vi["enabled"] = valueItem.Enabled
// 			vis[i] = vi
// 		}
// 		return vis
// 	}
// 	return make([]interface{}, 0)
// }

// func mapToEnvironmentValueItemsResponse(rawValueItems interface{}) []postman.CreateEnvironmentRequestEnvironmentValuesInner {
// 	if rawValueItems != nil {
// 		valueItems := rawValueItems.([]interface{})
// 		vis := make([]postman.CreateEnvironmentRequestEnvironmentValuesInner, len(valueItems), len(valueItems))
// 		for idx, valueItem := range valueItems {
// 			i := valueItem.(map[string]interface{})
// 			key := i["key"].(string)
// 			valueType := i["type"].(string)
// 			value := i["value"].(string)
// 			enabled := i["enabled"].(bool)
// 			vi := postman.CreateEnvironmentRequestEnvironmentValuesInner{
// 				Key:     &key,
// 				Type:    &valueType,
// 				Value:   &value,
// 				Enabled: &enabled,
// 			}
// 			vis[idx] = vi
// 		}
// 		return vis
// 	}
// 	return make([]postman.CreateEnvironmentRequestEnvironmentValuesInner, 0)
// }
