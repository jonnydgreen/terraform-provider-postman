package postman

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	postmanSDK "github.com/jonnydgreen/terraform-provider-postman/client/postman"
)

func dataSourceWorkspace() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceWorkspaceRead,
		Schema: map[string]*schema.Schema{
			"id": {
				Type:        schema.TypeString,
				Description: "The workspace's ID.",
				Required:    true,
			},
			"name": {
				Type:        schema.TypeString,
				Description: "The workspace's name.",
				Computed:    true,
			},
			"type": {
				Type:        schema.TypeString,
				Description: "The type of workspace. One of: personal|team",
				Computed:    true,
			},
			"description": {
				Type:        schema.TypeString,
				Description: "The workspace's description.",
				Computed:    true,
			},
			"visibility": {
				Type:        schema.TypeString,
				Description: "The workspace's visibility. [Visibility](https://learning.postman.com/docs/collaborating-in-postman/using-workspaces/managing-workspaces/#changing-workspace-visibility) determines who can access the workspace.",
				Computed:    true,
			},
			"created_by": {
				Type:        schema.TypeString,
				Description: "The user ID of the user who created the workspace.",
				Computed:    true,
			},
			"updated_by": {
				Type:        schema.TypeString,
				Description: "The user ID of the user who last updated the workspace.",
				Computed:    true,
			},
			"created_at": {
				Type:        schema.TypeString,
				Description: "The date and time at which the workspace was created.",
				Computed:    true,
			},
			"updated_at": {
				Type:        schema.TypeString,
				Description: "The date and time at which the workspace was last updated.",
				Computed:    true,
			},
			"collections": {
				Type:        schema.TypeList,
				Description: "The workspace's collections.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:        schema.TypeString,
							Description: "",
							Computed:    true,
						},
						"name": {
							Type:        schema.TypeString,
							Description: "",
							Computed:    true,
						},
						"uid": {
							Type:        schema.TypeString,
							Description: "",
							Computed:    true,
						},
					},
				},
			},
			"environments": {
				Type:        schema.TypeList,
				Description: "The workspace's collections.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:        schema.TypeString,
							Description: "",
							Computed:    true,
						},
						"name": {
							Type:        schema.TypeString,
							Description: "",
							Computed:    true,
						},
						"uid": {
							Type:        schema.TypeString,
							Description: "",
							Computed:    true,
						},
					},
				},
			},
			"mocks": {
				Type:        schema.TypeList,
				Description: "The workspace's collections.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:        schema.TypeString,
							Description: "",
							Computed:    true,
						},
						"name": {
							Type:        schema.TypeString,
							Description: "",
							Computed:    true,
						},
						"uid": {
							Type:        schema.TypeString,
							Description: "",
							Computed:    true,
						},
					},
				},
			},
			"monitors": {
				Type:        schema.TypeList,
				Description: "The workspace's collections.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:        schema.TypeString,
							Description: "",
							Computed:    true,
						},
						"name": {
							Type:        schema.TypeString,
							Description: "",
							Computed:    true,
						},
						"uid": {
							Type:        schema.TypeString,
							Description: "",
							Computed:    true,
						},
					},
				},
			},
			"apis": {
				Type:        schema.TypeList,
				Description: "The workspace's collections.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:        schema.TypeString,
							Description: "",
							Computed:    true,
						},
						"name": {
							Type:        schema.TypeString,
							Description: "",
							Computed:    true,
						},
						"uid": {
							Type:        schema.TypeString,
							Description: "",
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func dataSourceWorkspaceRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*postmanSDK.APIClient)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	workspaceID := d.Get("id").(string)
	d.SetId(workspaceID)

	response, _, err := c.WorkspacesApi.SingleWorkspace(ctx, workspaceID).Execute()
	if err != nil {
		return diag.FromErr(err)
	}

	responseWorkspace, isWorkspaceDefined := response.GetWorkspaceOk()
	if responseWorkspace == nil || isWorkspaceDefined != true {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Unable to find workspace",
			Detail:   fmt.Sprintf("No workspace with ID %s found in Postman API response.", workspaceID),
		})
		return diags
	}

	diags = append(diags, diag.Diagnostic{
		Severity: diag.Warning,
		Summary:  "INFO",
		Detail:   fmt.Sprintf("Workspace ID %s", *responseWorkspace.Id),
	})
	err = setWorkspaceResourceData(d, responseWorkspace)
	if err != nil {
		return diag.FromErr(err)
	}
	return diags
}

func setWorkspaceResourceData(d *schema.ResourceData, responseWorkspace *postmanSDK.SingleWorkspace200ResponseWorkspace) error {
	if err := d.Set("id", responseWorkspace.Id); err != nil {
		return fmt.Errorf("Error setting id: %s", err)
	}
	if err := d.Set("name", responseWorkspace.Name); err != nil {
		return fmt.Errorf("Error setting name: %s", err)
	}
	if err := d.Set("type", responseWorkspace.Type); err != nil {
		return fmt.Errorf("Error setting type: %s", err)
	}
	if err := d.Set("description", responseWorkspace.Description); err != nil {
		return fmt.Errorf("Error setting description: %s", err)
	}
	if err := d.Set("visibility", responseWorkspace.Visibility); err != nil {
		return fmt.Errorf("Error setting visibility: %s", err)
	}
	if err := d.Set("created_by", responseWorkspace.CreatedBy); err != nil {
		return fmt.Errorf("Error setting created_by: %s", err)
	}
	if err := d.Set("updated_by", responseWorkspace.UpdatedBy); err != nil {
		return fmt.Errorf("Error setting updated_by: %s", err)
	}
	if err := d.Set("created_at", responseWorkspace.CreatedAt.String()); err != nil {
		return fmt.Errorf("Error setting created_at: %s", err)
	}
	if err := d.Set("updated_at", responseWorkspace.UpdatedAt.String()); err != nil {
		return fmt.Errorf("Error setting updated_at: %s", err)
	}
	if err := d.Set("collections", flattenCollectionItemsData(&responseWorkspace.Collections)); err != nil {
		return fmt.Errorf("Error setting collections: %s", err)
	}
	if err := d.Set("environments", flattenEnvironmentItemsData(&responseWorkspace.Environments)); err != nil {
		return fmt.Errorf("Error setting environments: %s", err)
	}
	if err := d.Set("mocks", flattenMockItemsData(&responseWorkspace.Mocks)); err != nil {
		return fmt.Errorf("Error setting mocks: %s", err)
	}
	if err := d.Set("monitors", flattenMonitorItemsData(&responseWorkspace.Monitors)); err != nil {
		return fmt.Errorf("Error setting monitors: %s", err)
	}
	if err := d.Set("apis", flattenApiItemsData(&responseWorkspace.Apis)); err != nil {
		return fmt.Errorf("Error setting apis: %s", err)
	}
	return nil
}

func flattenCollectionItemsData(collectionItems *[]postmanSDK.SingleWorkspace200ResponseWorkspaceCollectionsInner) []interface{} {
	if collectionItems != nil {
		cis := make([]interface{}, len(*collectionItems), len(*collectionItems))
		for i, collectionItem := range *collectionItems {
			ci := make(map[string]interface{})
			ci["id"] = collectionItem.Id
			ci["name"] = collectionItem.Name
			ci["uid"] = collectionItem.Uid
			cis[i] = ci
		}
		return cis
	}
	return make([]interface{}, 0)
}

func flattenEnvironmentItemsData(environmentItems *[]postmanSDK.SingleWorkspace200ResponseWorkspaceEnvironmentsInner) []interface{} {
	if environmentItems != nil {
		eis := make([]interface{}, len(*environmentItems), len(*environmentItems))
		for i, environmentItem := range *environmentItems {
			ei := make(map[string]interface{})
			ei["id"] = environmentItem.Id
			ei["name"] = environmentItem.Name
			ei["uid"] = environmentItem.Uid
			eis[i] = ei
		}
		return eis
	}
	return make([]interface{}, 0)
}

func flattenMockItemsData(mockItems *[]postmanSDK.SingleWorkspace200ResponseWorkspaceMocksInner) []interface{} {
	if mockItems != nil {
		mis := make([]interface{}, len(*mockItems), len(*mockItems))
		for i, mockItem := range *mockItems {
			mi := make(map[string]interface{})
			mi["id"] = mockItem.Id
			mi["name"] = mockItem.Name
			mi["uid"] = mockItem.Uid
			mis[i] = mi
		}
		return mis
	}
	return make([]interface{}, 0)
}

func flattenMonitorItemsData(monitorItems *[]postmanSDK.SingleWorkspace200ResponseWorkspaceMonitorsInner) []interface{} {
	if monitorItems != nil {
		mis := make([]interface{}, len(*monitorItems), len(*monitorItems))
		for i, monitorItem := range *monitorItems {
			mi := make(map[string]interface{})
			mi["id"] = monitorItem.Id
			mi["name"] = monitorItem.Name
			mi["uid"] = monitorItem.Uid
			mis[i] = mi
		}
		return mis
	}
	return make([]interface{}, 0)
}

func flattenApiItemsData(apiItems *[]postmanSDK.SingleWorkspace200ResponseWorkspaceApisInner) []interface{} {
	if apiItems != nil {
		ais := make([]interface{}, len(*apiItems), len(*apiItems))
		for i, apiItem := range *apiItems {
			ai := make(map[string]interface{})
			ai["id"] = apiItem.Id
			ai["name"] = apiItem.Name
			ai["uid"] = apiItem.Uid
			ais[i] = ai
		}
		return ais
	}
	return make([]interface{}, 0)
}
