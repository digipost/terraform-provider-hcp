package hcp

import (
	"github.com/digipost/hcp"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
)

func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"username": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("HCP_USERNAME", nil),
				Description: "The username to use for HCP MAPI operations.",
			},
			"password": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("HCP_PASSWORD", nil),
				Description: "The password to use for HCP MAPI operations.",
			},
			"mapi_url": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("HCP_MAPI_URL", nil),
				Description: "The url of the HCP MAPI",
			},
			"insecure": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: "Ignore TLS validation errors",
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"hcp_user_account": resourceUserAccount(),
			"hcp_namespace":    resourceNamespace(),
		},
		ConfigureFunc: providerConfigure,
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {

	return &hcp.HCP{
		Username: d.Get("username").(string),
		Password: d.Get("password").(string),
		URL:      d.Get("mapi_url").(string),
		Insecure: d.Get("insecure").(bool),
	}, nil

}

func hcpClient(meta interface{}) *hcp.HCP {
	return meta.(*hcp.HCP)
}
