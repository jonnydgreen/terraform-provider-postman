package main

import (
	"context"
	"flag"

	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/jonnydgreen/terraform-provider-postman/postman"
)

// Provider documentation generation.
//go:generate go run github.com/hashicorp/terraform-plugin-docs/cmd/tfplugindocs generate --provider-name postman

var (
	// these will be set by the goreleaser configuration
	// to appropriate values for the compiled binary
	version string = "0.2"

	// goreleaser can also pass the specific commit if you want
	// commit  string = ""
)

func main() {
	var debugMode bool

	flag.BoolVar(&debugMode, "debug", false, "set to true to run the provider with support for debuggers like delve")
	flag.Parse()

	providerserver.Serve(context.Background(), postman.New(version), providerserver.ServeOpts{
		Address: "registry.terraform.io/jonnydgreen/postman",
		Debug:   debugMode,
	})
}
