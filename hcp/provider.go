package hcp

import (
	"github.com/digipost/hcp"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
)

func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"username": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("HCP_USERNAME", nil),
				Description: "The username to use for HCP MAPI operations.",
			},
			"password": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("HCP_PASSWORD", nil),
				Description: "The password to use for HCP MAPI operations.",
			},
			"mapi_url": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("HCP_MAPI_URL", nil),
				Description: "The url of the HCP MAPI",
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"hcp_user_account": resourceHCPUserAccount(),
		},
		ConfigureFunc: providerConfigure,
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {

	return &hcp.HCP{
		Username: d.Get("username").(string),
		Password: d.Get("password").(string),
		URL:      d.Get("mapi_url").(string),
		Insecure: true,
	}, nil

}

func hcpClient(meta interface{}) *hcp.HCP {
	return meta.(*hcp.HCP)
}
