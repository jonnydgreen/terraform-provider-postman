package postman

import (
	"context"
	"log"
	"os"
	"sync"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/jonnydgreen/terraform-provider-postman/client/postman"
)

var testAccProvider *schema.Provider

// providerFactories are used to instantiate a provider during acceptance testing.
// The factory function will be invoked for every Terraform CLI command executed
// to create a provider server to which the CLI can reattach.
var providerFactories = map[string]func() (*schema.Provider, error){
	"postman": func() (*schema.Provider, error) {
		return testAccProvider, nil
	},
}

func init() {
	testAccProvider = Provider("dev")()
	err := testAccProvider.Configure(context.Background(), terraform.NewResourceConfigRaw(nil))
	if err != nil {
		log.Fatal(err)
	}
}

func TestProvider(t *testing.T) {
	if err := Provider("dev")().InternalValidate(); err != nil {
		t.Fatalf("err: %s", err)
	}
}

var testAccProviderConfigure sync.Once

func failIfEmpty(t *testing.T, name string, usageMessage string) string {
	t.Helper()

	value := os.Getenv(name)

	if value == "" {
		t.Fatalf("environment variable %s must be set. Usage: %s", name, usageMessage)
	}

	return value
}

func testAccPreCheck(t *testing.T) func() {
	return func() {
		// Since we are outside the scope of the Terraform configuration we must
		// call Configure() to properly initialize the provider configuration.

		failIfEmpty(t, "POSTMAN_API_KEY", "static credentials value when using POSTMAN_API_KEY")
		err := testAccProvider.Configure(context.Background(), terraform.NewResourceConfigRaw(nil))
		if err != nil {
			t.Fatal(err)
		}
	}
}

func randString(t *testing.T, length int) string {
	return acctest.RandString(length)
}

func getProviderClient(t *testing.T) *postman.APIClient {
	return testAccProvider.Meta().(*postman.APIClient)
}

func testAccErrorCheck(t *testing.T) resource.ErrorCheckFunc {
	return func(err error) error {
		// Placeholder function for now to allow for common issues to be skipped
		return err
	}
}
