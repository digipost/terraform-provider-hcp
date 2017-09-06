package hcp

import (
	"github.com/digipost/hcp"
	"github.com/hashicorp/terraform/helper/schema"
)

const defaultLocalAuthentication = true
const defaultEnabled = true
const defaultForcePasswordChange = false
const defaultDescription = "User is managed by Terraform"
const defaultAllowNamespaceManagement = false

func resourceUserAccount() *schema.Resource {
	return &schema.Resource{
		Create: resourceUserAccountCreate,
		Read:   resourceUserAccountRead,
		Update: resourceUserAccountUpdate,
		Delete: resourceUserAccountDelete,
		Exists: resourceUserAccountExists,
		Schema: map[string]*schema.Schema{
			"username": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"password": &schema.Schema{
				Type:      schema.TypeString,
				Required:  true,
				Sensitive: true,
			},

			"full_name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceUserAccountCreate(d *schema.ResourceData, m interface{}) error {

	username := d.Get("username").(string)
	password := d.Get("password").(string)
	fullName := d.Get("full_name").(string)

	uA := &hcp.UserAccount{
		Username:                 username,
		FullName:                 fullName,
		Description:              defaultDescription,
		LocalAuthentication:      defaultLocalAuthentication,
		ForcePasswordChange:      defaultForcePasswordChange,
		Enabled:                  defaultEnabled,
		AllowNamespaceManagement: defaultAllowNamespaceManagement,
	}

	if err := hcpClient(m).CreateUserAccount(uA, password); err == nil {
		d.SetId(username)
		return nil
	} else {
		return err
	}

}

func resourceUserAccountUpdate(d *schema.ResourceData, m interface{}) error {

	username := d.Get("username").(string)
	password := d.Get("password").(string)
	fullName := d.Get("full_name").(string)

	uA := &hcp.UserAccount{
		Username:                 username,
		FullName:                 fullName,
		Description:              defaultDescription,
		ForcePasswordChange:      defaultForcePasswordChange,
		Enabled:                  defaultEnabled,
		AllowNamespaceManagement: defaultAllowNamespaceManagement,
	}

	if err := hcpClient(m).UpdateUserAccount(uA, password); err == nil {
		d.SetId(username)
		return nil
	} else {
		return err
	}
}

func resourceUserAccountRead(d *schema.ResourceData, m interface{}) error {

	username := d.Get("username").(string)
	if userAccount, err := hcpClient(m).UserAccount(username); err == nil {
		d.Set("full_name", userAccount.FullName)
		return nil
	} else {
		return err
	}

}

func resourceUserAccountDelete(d *schema.ResourceData, m interface{}) error {
	username := d.Get("username").(string)
	return hcpClient(m).DeleteUserAccount(username)
}

func resourceUserAccountExists(d *schema.ResourceData, m interface{}) (bool, error) {
	username := d.Get("username").(string)
	return hcpClient(m).UserAccountExists(username)
}
