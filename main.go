package main

import (
	"github.com/arkiaconsulting/terraform-provider-akc/akc"
	"github.com/hashicorp/terraform-plugin-sdk/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: akc.Provider,
	})
}
