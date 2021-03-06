package main

import (
	"github.com/hashicorp/terraform/plugin"
	"github.com/wearespindle/terraform-provider-phpipam/plugin/providers/phpipam"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: phpipam.Provider,
	})
}
