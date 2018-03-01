package main

import (
	"github.com/hashicorp/terraform/plugin"
	"github.com/infobloxopen/terraform-provider-infoblox/infoblox"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: infoblox.Provider})
}
