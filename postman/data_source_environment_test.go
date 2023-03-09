package postman

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccEnvironmentDataSource__basic(t *testing.T) {
	dataSourceName := "data.postman_environment.default"
	resourceName := "postman_environment.default"
	environmentName := acctest.RandomWithPrefix("tf-test")
	workspaceName := acctest.RandomWithPrefix("tf-test")
	workspaceType := "personal"
	context := map[string]interface{}{
		"name":          environmentName,
		"workspaceName": workspaceName,
		"workspaceType": workspaceType,
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 testAccPreCheck(t),
		ErrorCheck:               testAccErrorCheck(t),
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Read testing
			{
				Config: testAccEnvironmentDataSource__basic(context),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrPair(dataSourceName, "id", resourceName, "id"),
					resource.TestCheckResourceAttrPair(dataSourceName, "name", resourceName, "name"),
					resource.TestCheckResourceAttrPair(dataSourceName, "workspace", resourceName, "workspace"),
					resource.TestCheckResourceAttrPair(dataSourceName, "is_public", resourceName, "is_public"),
					resource.TestCheckResourceAttrPair(dataSourceName, "owner", resourceName, "owner"),
					resource.TestCheckResourceAttrPair(dataSourceName, "created_at", resourceName, "created_at"),
					resource.TestCheckResourceAttrPair(dataSourceName, "updated_at", resourceName, "updated_at"),
				),
			},
		},
	})
}

func testAccEnvironmentDataSource__basic(context map[string]interface{}) string {
	return Nprintf(providerConfig+`
resource "postman_workspace" "default" {
	name = "%{workspaceName}"
	type = "%{workspaceType}"
}

resource "postman_environment" "default" {
	name = "%{name}"
	workspace = postman_workspace.default.id
}

data "postman_environment" "default" {
	id = postman_environment.default.id
	workspace = postman_environment.default.workspace
}
`, context)
}
