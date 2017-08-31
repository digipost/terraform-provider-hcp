package hcp

import (
	"github.com/digipost/hcp"
	"github.com/hashicorp/terraform/helper/schema"
)

/*
resource "hcp_user_account" "foo" {
    // internal User ID: abdd0999-92db-4c20-b578-f6ff353a4d01

    // Required (can be renamed)
    username = "Username"

    // Required
    full_name = "Full Name"

     // Optional
    roles = ["COMPLIANCE", "SECURITY"] // default empty

    --------

    // Optional
    description = "Description" // default empty

    // Optional
    localAuthentication = "true" // default

    // Optional
    forcePasswordChange = "false" // default

    // Optional
    enabled = "true" // default

    // Optional
    allowNamespaceManagement = "false" // default
}
*/

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
				Type:     schema.TypeString,
				Required: true,
			},

			"full_name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},

			"roles": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Set:      schema.HashString,
			},
		},
	}
}

func resourceUserAccountCreate(d *schema.ResourceData, m interface{}) error {
	return resourceUserAccountCreateOrUpdate(true, d, m)
}

func resourceUserAccountUpdate(d *schema.ResourceData, m interface{}) error {
	return resourceUserAccountCreateOrUpdate(false, d, m)
}

func resourceUserAccountCreateOrUpdate(create bool, d *schema.ResourceData, m interface{}) error {
	hcpClient := hcpClient(m)

	username := d.Get("username").(string)
	password := d.Get("password").(string)
	fullName := d.Get("full_name").(string)

	uA := &hcp.UserAccount{Username: username, FullName: fullName}

	var err error
	if create {
		err = hcpClient.CreateUserAccount(uA, password)
	} else {
		err = hcpClient.UpdateUserAccount(uA, password)
	}

	if err == nil {
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
