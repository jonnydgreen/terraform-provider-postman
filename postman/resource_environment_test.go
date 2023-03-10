package postman

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccEnvironmentResource__basic(t *testing.T) {
	resourceName := "postman_environment.default"
	workspaceResourceName := "postman_workspace.default"
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
		CheckDestroy:             testAccCheckEnvironmentDoesNotExist(t, resourceName),
		Steps: []resource.TestStep{
			// Create and Read testing
			{
				Config: testAccEnvironmentResource__basic(context),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", environmentName),
					resource.TestCheckResourceAttrPair(resourceName, "workspace", workspaceResourceName, "id"),
					// Verify dynamic values have any value set in the state.
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttrSet(resourceName, "is_public"),
					resource.TestCheckResourceAttrSet(resourceName, "owner"),
					resource.TestCheckResourceAttrSet(resourceName, "created_at"),
					resource.TestCheckResourceAttrSet(resourceName, "updated_at"),
					testAccCheckEnvironmentExists(t, resourceName),
				),
			},
			// ImportState testing
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateIdFunc: func(state *terraform.State) (string, error) {
					rs, ok := state.RootModule().Resources[resourceName]
					if !ok {
						return "", fmt.Errorf("Not found: %s", resourceName)
					}
					id := rs.Primary.ID
					workspace := rs.Primary.Attributes["workspace"]
					return fmt.Sprintf("%s,%s", workspace, id), nil
				},
			},
			// Update and Read testing
			{
				Config: testAccEnvironmentResource__basic(map[string]interface{}{
					"name":          environmentName + "-2",
					"workspaceName": workspaceName,
					"workspaceType": workspaceType,
				}),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", environmentName+"-2"),
					resource.TestCheckResourceAttrPair(resourceName, "workspace", workspaceResourceName, "id"),
					// Verify dynamic values have any value set in the state.
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttrSet(resourceName, "is_public"),
					resource.TestCheckResourceAttrSet(resourceName, "owner"),
					resource.TestCheckResourceAttrSet(resourceName, "created_at"),
					resource.TestCheckResourceAttrSet(resourceName, "updated_at"),

					testAccCheckEnvironmentExists(t, resourceName),
				),
			},
		},
	})
}

func TestAccEnvironmentResource__basicValues(t *testing.T) {
	resourceName := "postman_environment.default"
	workspaceResourceName := "postman_workspace.default"
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
		CheckDestroy:             testAccCheckEnvironmentDoesNotExist(t, resourceName),
		Steps: []resource.TestStep{
			// Create and Read testing
			{
				Config: testAccEnvironmentResource__basicValues(context),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", environmentName),
					resource.TestCheckResourceAttrPair(resourceName, "workspace", workspaceResourceName, "id"),
					// Verify number of values
					resource.TestCheckResourceAttr(resourceName, "values.#", "3"),
					// Verify 1st value
					resource.TestCheckResourceAttr(resourceName, "values.0.key", "key1"),
					resource.TestCheckResourceAttr(resourceName, "values.0.value", "value1"),
					// Verify 2nd value
					resource.TestCheckResourceAttr(resourceName, "values.1.key", "key2"),
					resource.TestCheckResourceAttr(resourceName, "values.1.value", "value2"),
					resource.TestCheckResourceAttr(resourceName, "values.1.type", "secret"),
					// Verify 3rd value
					resource.TestCheckResourceAttr(resourceName, "values.2.key", "key3"),
					resource.TestCheckResourceAttr(resourceName, "values.2.value", "value3"),
					resource.TestCheckResourceAttr(resourceName, "values.2.enabled", "false"),
					// Verify dynamic values have any value set in the state.
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttrSet(resourceName, "is_public"),
					resource.TestCheckResourceAttrSet(resourceName, "owner"),
					resource.TestCheckResourceAttrSet(resourceName, "created_at"),
					resource.TestCheckResourceAttrSet(resourceName, "updated_at"),
					testAccCheckEnvironmentExists(t, resourceName),
				),
			},
			// ImportState testing
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateIdFunc: func(state *terraform.State) (string, error) {
					rs, ok := state.RootModule().Resources[resourceName]
					if !ok {
						return "", fmt.Errorf("Not found: %s", resourceName)
					}
					id := rs.Primary.ID
					workspace := rs.Primary.Attributes["workspace"]
					return fmt.Sprintf("%s,%s", workspace, id), nil
				},
			},
			// Update and Read testing
			{
				Config: testAccEnvironmentResource__basicValues(map[string]interface{}{
					"name":          environmentName + "-2",
					"workspaceName": workspaceName,
					"workspaceType": workspaceType,
				}),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", environmentName+"-2"),
					resource.TestCheckResourceAttrPair(resourceName, "workspace", workspaceResourceName, "id"),
					// Verify number of values
					resource.TestCheckResourceAttr(resourceName, "values.#", "3"),
					// Verify 1st value
					resource.TestCheckResourceAttr(resourceName, "values.0.key", "key1"),
					resource.TestCheckResourceAttr(resourceName, "values.0.value", "value1"),
					// Verify 2nd value
					resource.TestCheckResourceAttr(resourceName, "values.1.key", "key2"),
					resource.TestCheckResourceAttr(resourceName, "values.1.value", "value2"),
					resource.TestCheckResourceAttr(resourceName, "values.1.type", "secret"),
					// Verify 3rd value
					resource.TestCheckResourceAttr(resourceName, "values.2.key", "key3"),
					resource.TestCheckResourceAttr(resourceName, "values.2.value", "value3"),
					resource.TestCheckResourceAttr(resourceName, "values.2.enabled", "false"),
					// Verify dynamic values have any value set in the state.
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttrSet(resourceName, "is_public"),
					resource.TestCheckResourceAttrSet(resourceName, "owner"),
					resource.TestCheckResourceAttrSet(resourceName, "created_at"),
					resource.TestCheckResourceAttrSet(resourceName, "updated_at"),

					testAccCheckEnvironmentExists(t, resourceName),
				),
			},
		},
	})
}

func testAccEnvironmentResource__basic(context map[string]interface{}) string {
	return Nprintf(providerConfig+`
resource "postman_workspace" "default" {
	name = "%{workspaceName}"
	type = "%{workspaceType}"
}

resource "postman_environment" "default" {
  name = "%{name}"
  workspace = postman_workspace.default.id
}
`, context)
}

func testAccEnvironmentResource__basicValues(context map[string]interface{}) string {
	return Nprintf(providerConfig+`
resource "postman_workspace" "default" {
	name = "%{workspaceName}"
	type = "%{workspaceType}"
}

resource "postman_environment" "default" {
  name = "%{name}"
  workspace = postman_workspace.default.id
	values = [
		{
			key   = "key1"
			value = "value1"
		},
		{
			key   = "key2"
			value = "value2"
			type = "secret"
		},
		{
			key   = "key3"
			value = "value3"
			enabled = false
		}
	]
}
`, context)
}

func testAccEnvironmentResource__basicDefaultWorkspace(context map[string]interface{}) string {
	return Nprintf(providerConfig+`
resource "postman_environment" "description" {
  name        = "%{name}"
}
`, context)
}

func testAccCheckEnvironmentExists(t *testing.T, resourceName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("Resource not found: %s", resourceName)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No resource ID is set")
		}

		environmentID := rs.Primary.ID
		c := getProviderClient(t)
		response, _, err := c.EnvironmentsApi.SingleEnvironment(context.Background(), environmentID).Execute()
		if err != nil {
			return err
		}
		if *response.Environment.Id == environmentID {
			return nil
		}
		return fmt.Errorf("Postman Environment with ID %s does not exist", environmentID)
	}
}

func testAccCheckEnvironmentDoesNotExist(t *testing.T, resourceName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("Resource not found: %s", resourceName)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No resource ID is set")
		}

		environmentID := rs.Primary.ID
		c := getProviderClient(t)
		_, raw, err := c.EnvironmentsApi.SingleEnvironment(context.Background(), environmentID).Execute()
		if err != nil {
			if raw.StatusCode == 404 {
				return nil
			}
			return err
		}
		return fmt.Errorf("Postman Environment with ID %s exists", environmentID)
	}
}

func testAccCheckEnvironmentDisappears(t *testing.T, resourceName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		c := getProviderClient(t)
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("Resource not found: %s", resourceName)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No resource ID is set")
		}

		environmentID := rs.Primary.ID
		_, _, err := c.EnvironmentsApi.DeleteEnvironment(context.Background(), environmentID).Execute()
		if err != nil {
			return fmt.Errorf("Error deleting Postman Environment Resource: %s", err)
		}
		return nil
	}
}
