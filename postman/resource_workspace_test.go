package postman

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccPostmanWorkspace__basic(t *testing.T) {
	// TODO
	// rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	resourceName := "postman_workspace.default"
	context := map[string]interface{}{
		"random_suffix": randString(t, 10),
		"type":          "personal",
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          testAccPreCheck(t),
		ErrorCheck:        testAccErrorCheck(t),
		ProviderFactories: providerFactories,
		CheckDestroy:      testAccCheckWorkspaceDoesNotExist(t, resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccPostmanWorkspace__basic(context),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckWorkspaceExists(t, resourceName),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccPostmanWorkspace_Disappears_basic(t *testing.T) {
	resourceName := "postman_workspace.default"

	context := map[string]interface{}{
		"random_suffix": randString(t, 10),
		"type":          "personal",
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          testAccPreCheck(t),
		ErrorCheck:        testAccErrorCheck(t),
		ProviderFactories: providerFactories,
		CheckDestroy:      testAccCheckWorkspaceDoesNotExist(t, resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccPostmanWorkspace__basic(context),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckWorkspaceExists(t, resourceName),
					testAccCheckWorkspaceDisappears(t, resourceName),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func testAccPostmanWorkspace__basic(context map[string]interface{}) string {
	return Nprintf(`
resource "postman_workspace" "default" {
  name = "tf-test-%{random_suffix}"
  type = "%{type}"
}
`, context)
}

func testAccCheckWorkspaceExists(t *testing.T, resourceName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		// TODO: util func
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("Resource not found: %s", resourceName)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No resource ID is set")
		}

		workspaceID := rs.Primary.ID
		c := getProviderClient(t)
		response, _, err := c.WorkspacesApi.SingleWorkspace(context.Background(), workspaceID).Execute()
		if err != nil {
			return err
		}
		if *response.Workspace.Id == workspaceID {
			return nil
		}
		return fmt.Errorf("Postman Workspace with ID %s does not exist", workspaceID)
	}
}

func testAccCheckWorkspaceDoesNotExist(t *testing.T, resourceName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("Resource not found: %s", resourceName)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No resource ID is set")
		}

		workspaceID := rs.Primary.ID
		c := getProviderClient(t)
		_, raw, err := c.WorkspacesApi.SingleWorkspace(context.Background(), workspaceID).Execute()
		if err != nil {
			if raw.StatusCode == 404 {
				return nil
			}
			return err
		}
		return fmt.Errorf("Postman Workspace with ID %s exists", workspaceID)
	}
}

func testAccCheckWorkspaceDisappears(t *testing.T, resourceName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		c := getProviderClient(t)
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("Resource not found: %s", resourceName)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No resource ID is set")
		}

		workspaceID := rs.Primary.ID
		_, _, err := c.WorkspacesApi.DeleteWorkspace(context.Background(), workspaceID).Execute()
		if err != nil {
			return fmt.Errorf("Error deleting Postman Workspace Resource: %s", err)
		}
		return nil
	}
}

// TODO
//  - Per Attribute
//  - rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
