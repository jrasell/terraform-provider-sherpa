package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/plugin"
	"github.com/jrasell/terraform-provider-sherpa/sherpa"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: sherpa.Provider})
}
