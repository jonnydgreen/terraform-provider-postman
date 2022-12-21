package postman

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
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
	return resourceWorkspaceRead(ctx, d, m)
}
