package postman

import (
	"os"
	"sync"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/jonnydgreen/terraform-provider-postman/client/postman"
)

var (
	// providerConfig is a shared configuration to combine with the actual
	// test configuration so the Postman client is properly configured.
	// The POSTMAN_API_KEY environment must be set.
	providerConfig = `
provider "postman" {}
`
)

var (
	// testAccProtoV6ProviderFactories are used to instantiate a provider during
	// acceptance testing. The factory function will be invoked for every Terraform
	// CLI command executed to create a provider server to which the CLI can
	// reattach.
	testAccProtoV6ProviderFactories = map[string]func() (tfprotov6.ProviderServer, error){
		"postman": providerserver.NewProtocol6WithError(New("dev")()),
	}
)

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
		failIfEmpty(t, "POSTMAN_API_KEY", "static credentials value when using POSTMAN_API_KEY")
	}
}

func randString(t *testing.T, length int) string {
	return acctest.RandString(length)
}

// TODO
// Make better
func getProviderClient(t *testing.T) *postman.APIClient {
	failIfEmpty(t, "POSTMAN_API_KEY", "static credentials value when using POSTMAN_API_KEY")
	apiKey := os.Getenv("POSTMAN_API_KEY")
	configuration := postman.NewConfiguration()
	configuration.AddDefaultHeader("x-api-key", apiKey)
	return postman.NewAPIClient(configuration)
}

// TODO
// func testAccErrorCheck(t *testing.T) resource.ErrorCheckFunc {
// 	return func(err error) error {
// 		// Placeholder function for now to allow for common issues to be skipped
// 		return err
// 	}
// }
