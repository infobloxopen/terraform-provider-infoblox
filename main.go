package main

import (
	"context"
	"flag"
	"log"

	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/provider"
)

// Run "go generate" to format example terraform files and generate the docs for the registry/website

// If you do not have terraform installed, you can remove the formatting command, but its suggested to
// ensure the documentation is formatted properly.
//go:generate terraform fmt -recursive ./examples/

// Run the docs generation tool, check its repository for more information on how it works and how docs
// can be customized. The provider type name is "infoblox" (it differs from the
// module name terraform-provider-infoblox), so it must be passed explicitly.
//go:generate go run github.com/hashicorp/terraform-plugin-docs/cmd/tfplugindocs generate

// Post-process the rendered docs: inject the subcategory (group) and per-backend
// Example Usage sections, driven by the internal/service/<group>/ package layout
// and the examples/ folder layout. Must run AFTER tfplugindocs. This avoids
// committing a per-object template file for every resource/data-source.
//go:generate go run ./tools/gendocs

var (
	// these will be set by the goreleaser configuration
	// to appropriate values for the compiled binary.
	version string = "dev"
	commit  string = "none"

	// goreleaser can pass other information to the main package, such as the specific commit
	// https://goreleaser.com/cookbooks/using-main.version/
)

func main() {
	var debug bool

	flag.BoolVar(&debug, "debug", false, "set to true to run the provider with support for debuggers like delve")
	flag.Parse()

	opts := providerserver.ServeOpts{
		Address: "registry.terraform.io/infobloxopen/infoblox",
		Debug:   debug,
	}

	err := providerserver.Serve(context.Background(), provider.New(version, commit), opts)

	if err != nil {
		log.Fatal(err.Error())
	}
}
