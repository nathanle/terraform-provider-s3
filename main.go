package main

import (
	"github.com/hashicorp/terraform/internal/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: Provider,
	})
}
