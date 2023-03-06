package postman

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccWorkspaceResource__basic(t *testing.T) {
	resourceName := "postman_workspace.default"
	workspaceName := acctest.RandomWithPrefix("tf-test")
	workspaceType := "personal"
	context := map[string]interface{}{
		"name": workspaceName,
		"type": workspaceType,
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: testAccPreCheck(t),
		// TODO
		// ErrorCheck:        testAccErrorCheck(t),
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckWorkspaceDoesNotExist(t, resourceName),
		Steps: []resource.TestStep{
			// Create and Read testing
			{
				Config: testAccWorkspaceResource__basic(context),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", workspaceName),
					resource.TestCheckResourceAttr(resourceName, "type", workspaceType),
					// TODO: description
					// resource.TestCheckResourceAttr(resourceName, "description", ""),
					// TODO: other computed props
					testAccCheckWorkspaceExists(t, resourceName),
				),
			},
			// ImportState testing
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
			// Update and Read testing
			{
				Config: testAccWorkspaceResource__basic(map[string]interface{}{
					"name": workspaceName + "-2",
					"type": workspaceType,
				}),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", workspaceName+"-2"),
					resource.TestCheckResourceAttr(resourceName, "type", workspaceType),
					// TODO
					// https://github.com/hashicorp/terraform-plugin-framework/issues/175
					// https://github.com/hashicorp/terraform-provider-aws/issues/28638
					// https://github.com/hashicorp/terraform-plugin-framework/pull/176
					// TODO: description
					// resource.TestCheckResourceAttr(resourceName, "description", ""),
					// TODO: other computed props
					testAccCheckWorkspaceExists(t, resourceName),
				),
			},
		},
	})
}

// func TestAccWorkspaceResource__name(t *testing.T) {
// 	nameBefore := acctest.RandomWithPrefix("tf-test-before")
// 	nameAfter := acctest.RandomWithPrefix("tf-test-before")
// 	resourceName := "postman_workspace.default"
// 	contextBefore := map[string]interface{}{
// 		"name": nameBefore,
// 		"type": "personal",
// 	}
// 	contextAfter := map[string]interface{}{
// 		"name": nameAfter,
// 		"type": "personal",
// 	}

// 	resource.ParallelTest(t, resource.TestCase{
// 		PreCheck:          testAccPreCheck(t),
// 		ErrorCheck:        testAccErrorCheck(t),
// 		ProviderFactories: providerFactories,
// 		CheckDestroy:      testAccCheckWorkspaceDoesNotExist(t, resourceName),
// 		Steps: []resource.TestStep{
// 			{
// 				Config: testAccWorkspaceResource__basic(contextBefore),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckWorkspaceExists(t, resourceName),
// 				),
// 			},
// 			{
// 				ResourceName:      resourceName,
// 				ImportState:       true,
// 				ImportStateVerify: true,
// 			},
// 			{
// 				Config: testAccWorkspaceResource__basic(contextAfter),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckWorkspaceExists(t, resourceName),
// 				),
// 			},
// 		},
// 	})
// }

// func TestAccWorkspaceResource__description(t *testing.T) {
// 	name := acctest.RandomWithPrefix("tf-test")
// 	resourceName := "postman_workspace.description"
// 	contextBefore := map[string]interface{}{
// 		"name":        name,
// 		"type":        "personal",
// 		"description": "Description before",
// 	}
// 	contextAfter := map[string]interface{}{
// 		"name":        name,
// 		"type":        "personal",
// 		"description": "Description after",
// 	}

// 	resource.ParallelTest(t, resource.TestCase{
// 		PreCheck:          testAccPreCheck(t),
// 		ErrorCheck:        testAccErrorCheck(t),
// 		ProviderFactories: providerFactories,
// 		CheckDestroy:      testAccCheckWorkspaceDoesNotExist(t, resourceName),
// 		Steps: []resource.TestStep{
// 			{
// 				Config: testAccWorkspaceResource__description(contextBefore),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckWorkspaceExists(t, resourceName),
// 				),
// 			},
// 			{
// 				ResourceName:      resourceName,
// 				ImportState:       true,
// 				ImportStateVerify: true,
// 			},
// 			{
// 				Config: testAccWorkspaceResource__description(contextAfter),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckWorkspaceExists(t, resourceName),
// 				),
// 			},
// 		},
// 	})
// }

// func TestAccWorkspaceResource__disappears_basic(t *testing.T) {
// 	name := acctest.RandomWithPrefix("tf-test")
// 	resourceName := "postman_workspace.default"
// 	context := map[string]interface{}{
// 		"name": name,
// 		"type": "personal",
// 	}

// 	resource.ParallelTest(t, resource.TestCase{
// 		PreCheck:          testAccPreCheck(t),
// 		ErrorCheck:        testAccErrorCheck(t),
// 		ProviderFactories: providerFactories,
// 		CheckDestroy:      testAccCheckWorkspaceDoesNotExist(t, resourceName),
// 		Steps: []resource.TestStep{
// 			{
// 				Config: testAccWorkspaceResource__basic(context),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckWorkspaceExists(t, resourceName),
// 					testAccCheckWorkspaceDisappears(t, resourceName),
// 				),
// 				ExpectNonEmptyPlan: true,
// 			},
// 		},
// 	})
// }

func testAccWorkspaceResource__basic(context map[string]interface{}) string {
	return Nprintf(providerConfig+`
resource "postman_workspace" "default" {
  name = "%{name}"
  type = "%{type}"
}
`, context)
}

func testAccWorkspaceResource__description(context map[string]interface{}) string {
	return Nprintf(providerConfig+`
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
