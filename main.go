package main

import (
	 
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
	"github.com/nathanle/terraform-provider-s3"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: Provider,
	})
}
