package main

import (
	 "github.com/hashicorp/terraform-plugin-sdk"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: Provider,
	})
}
