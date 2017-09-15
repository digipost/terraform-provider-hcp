package hcp

import (
	"github.com/digipost/hcp"
	"github.com/hashicorp/terraform/helper/schema"
	"log"
)

const (
	defaultDescription = "User is managed by Terraform"
)

var defaultRoles = []string{hcp.MONITOR}

func resourceUserAccount() *schema.Resource {
	return &schema.Resource{
		Create: resourceUserAccountCreate,
		Read:   resourceUserAccountRead,
		Update: resourceUserAccountUpdate,
		Delete: resourceUserAccountDelete,
		Exists: resourceUserAccountExists,
		Schema: map[string]*schema.Schema{
			"username": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"password": {
				Type:             schema.TypeString,
				Required:         true,
				Sensitive:        true,
				DiffSuppressFunc: suppressPasswordDiffs,
			},
			"full_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"local_authentication": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"enabled": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"force_password_change": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"allow_namespace_management": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
		},
	}
}

func resourceUserAccountCreate(d *schema.ResourceData, m interface{}) error {

	username := d.Get("username").(string)
	password := d.Get("password").(string)
	fullName := d.Get("full_name").(string)
	enabled := d.Get("enabled").(bool)
	forcePasswordChange := d.Get("force_password_change").(bool)
	allowNamespaceManagement := d.Get("allow_namespace_management").(bool)

	localAuthentication := d.Get("local_authentication").(bool)

	uA := &hcp.UserAccount{
		Username:                 username,
		FullName:                 fullName,
		Description:              defaultDescription,
		LocalAuthentication:      localAuthentication,
		ForcePasswordChange:      forcePasswordChange,
		Enabled:                  enabled,
		AllowNamespaceManagement: allowNamespaceManagement,
		Roles: defaultRoles,
	}

	if err := hcpClient(m).CreateUserAccount(uA, password); err == nil {
		d.SetId(username)
		d.Set("password", sha512sum(password))
		return nil
	} else {
		return err
	}

}

func resourceUserAccountUpdate(d *schema.ResourceData, m interface{}) error {

	username := d.Get("username").(string)
	password := d.Get("password").(string)
	fullName := d.Get("full_name").(string)
	enabled := d.Get("enabled").(bool)
	forcePasswordChange := d.Get("force_password_change").(bool)
	allowNamespaceManagement := d.Get("allow_namespace_management").(bool)

	uA := &hcp.UserAccount{
		Username:                 username,
		FullName:                 fullName,
		Description:              defaultDescription,
		ForcePasswordChange:      forcePasswordChange,
		Enabled:                  enabled,
		AllowNamespaceManagement: allowNamespaceManagement,
		Roles: defaultRoles,
	}

	hasPasswordChange := d.HasChange("password")
	log.Printf("[DEBUG] resourceUserAccountUpdate - hasPasswordChange = %t", hasPasswordChange)
	if hasPasswordChange {
		if err := hcpClient(m).UpdateUserAccount(uA, password); err == nil {
			d.SetId(username)
			d.Set("password", sha512sum(password))
			return nil
		} else {
			return err
		}

	} else {

		if err := hcpClient(m).UpdateUserAccountExceptPassword(uA); err == nil {
			d.SetId(username)
			return nil
		} else {
			return err
		}
	}
}

func resourceUserAccountRead(d *schema.ResourceData, m interface{}) error {

	username := d.Get("username").(string)
	if userAccount, err := hcpClient(m).ReadUserAccount(username); err == nil {
		// TODO?
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
