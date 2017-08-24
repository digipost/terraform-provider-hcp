package main

import (
	"github.com/digipost/terraform-provider-hcp/hcp"
	"github.com/hashicorp/terraform/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: hcp.Provider})
}
