package postman

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccWorkspace__basic(t *testing.T) {
	resourceName := "postman_workspace.default"
	name := acctest.RandomWithPrefix("tf-test")
	context := map[string]interface{}{
		"name": name,
		"type": "personal",
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          testAccPreCheck(t),
		ErrorCheck:        testAccErrorCheck(t),
		ProviderFactories: providerFactories,
		CheckDestroy:      testAccCheckWorkspaceDoesNotExist(t, resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccWorkspace__basic(context),
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

func TestAccWorkspace__name(t *testing.T) {
	nameBefore := acctest.RandomWithPrefix("tf-test-before")
	nameAfter := acctest.RandomWithPrefix("tf-test-before")
	resourceName := "postman_workspace.default"
	contextBefore := map[string]interface{}{
		"name": nameBefore,
		"type": "personal",
	}
	contextAfter := map[string]interface{}{
		"name": nameAfter,
		"type": "personal",
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          testAccPreCheck(t),
		ErrorCheck:        testAccErrorCheck(t),
		ProviderFactories: providerFactories,
		CheckDestroy:      testAccCheckWorkspaceDoesNotExist(t, resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccWorkspace__basic(contextBefore),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckWorkspaceExists(t, resourceName),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccWorkspace__basic(contextAfter),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckWorkspaceExists(t, resourceName),
				),
			},
		},
	})
}

func TestAccWorkspace__description(t *testing.T) {
	name := acctest.RandomWithPrefix("tf-test")
	resourceName := "postman_workspace.description"
	contextBefore := map[string]interface{}{
		"name":        name,
		"type":        "personal",
		"description": "Description before",
	}
	contextAfter := map[string]interface{}{
		"name":        name,
		"type":        "personal",
		"description": "Description after",
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          testAccPreCheck(t),
		ErrorCheck:        testAccErrorCheck(t),
		ProviderFactories: providerFactories,
		CheckDestroy:      testAccCheckWorkspaceDoesNotExist(t, resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccWorkspace__description(contextBefore),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckWorkspaceExists(t, resourceName),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccWorkspace__description(contextAfter),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckWorkspaceExists(t, resourceName),
				),
			},
		},
	})
}

func TestAccWorkspace__disappears_basic(t *testing.T) {
	name := acctest.RandomWithPrefix("tf-test")
	resourceName := "postman_workspace.default"
	context := map[string]interface{}{
		"name": name,
		"type": "personal",
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          testAccPreCheck(t),
		ErrorCheck:        testAccErrorCheck(t),
		ProviderFactories: providerFactories,
		CheckDestroy:      testAccCheckWorkspaceDoesNotExist(t, resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccWorkspace__basic(context),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckWorkspaceExists(t, resourceName),
					testAccCheckWorkspaceDisappears(t, resourceName),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func testAccWorkspace__basic(context map[string]interface{}) string {
	return Nprintf(`
resource "postman_workspace" "default" {
  name = "%{name}"
  type = "%{type}"
}
`, context)
}

func testAccWorkspace__description(context map[string]interface{}) string {
	return Nprintf(`
resource "postman_workspace" "description" {
  name        = "%{name}"
  type        = "%{type}"
	description = "%{description}"
}
`, context)
}

func testAccCheckWorkspaceExists(t *testing.T, resourceName string) resource.TestCheckFunc {
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
