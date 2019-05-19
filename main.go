package main

import (
	"github.com/hashicorp/terraform/plugin"
	"github.com/ianmiell/terraform-provider-ltthw/provider"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: provider.Provider,
	})
}
