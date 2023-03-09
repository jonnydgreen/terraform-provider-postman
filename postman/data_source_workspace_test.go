package postman

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccWorkspaceDataSource__basic(t *testing.T) {
	resourceName := "postman_workspace.default"
	dataSourceName := "data.postman_workspace.default"
	workspaceName := acctest.RandomWithPrefix("tf-test")
	workspaceType := "personal"
	context := map[string]interface{}{
		"name": workspaceName,
		"type": workspaceType,
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 testAccPreCheck(t),
		ErrorCheck:               testAccErrorCheck(t),
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Read testing
			{
				Config: testAccWorkspaceDataSource__basic(context),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrPair(dataSourceName, "id", resourceName, "id"),
					resource.TestCheckResourceAttrPair(dataSourceName, "name", resourceName, "name"),
					resource.TestCheckResourceAttrPair(dataSourceName, "type", resourceName, "type"),
					// TODO: description
					// resource.TestCheckResourceAttr(dataSourceName, "description", ""),
					resource.TestCheckResourceAttrPair(dataSourceName, "apis", resourceName, "apis"),
					resource.TestCheckResourceAttrPair(dataSourceName, "collections", resourceName, "collections"),
					resource.TestCheckResourceAttrPair(dataSourceName, "environments", resourceName, "environments"),
					resource.TestCheckResourceAttrPair(dataSourceName, "mocks", resourceName, "mocks"),
					resource.TestCheckResourceAttrPair(dataSourceName, "monitors", resourceName, "monitors"),
					resource.TestCheckResourceAttrPair(dataSourceName, "visibility", resourceName, "visibility"),
					resource.TestCheckResourceAttrPair(dataSourceName, "created_at", resourceName, "created_at"),
					resource.TestCheckResourceAttrPair(dataSourceName, "created_by", resourceName, "created_by"),
					resource.TestCheckResourceAttrPair(dataSourceName, "updated_at", resourceName, "updated_at"),
					resource.TestCheckResourceAttrPair(dataSourceName, "updated_by", resourceName, "updated_by"),
				),
			},
		},
	})
}

func testAccWorkspaceDataSource__basic(context map[string]interface{}) string {
	return Nprintf(providerConfig+`
resource "postman_workspace" "default" {
  name = "%{name}"
  type = "%{type}"
}

data "postman_workspace" "default" {
	id = postman_workspace.default.id
}
`, context)
}
