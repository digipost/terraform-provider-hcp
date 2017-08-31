package hcp

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/digipost/hcp"
)

/**
    <name>Accounts-Receivable</name>
    <hashScheme>SHA-256</hashScheme>
    <enterpriseMode>true</enterpriseMode>
    <hardQuota>50 GB</hardQuota>
    <softQuota>75</softQuota>
    <optimizedFor>ALL</optimizedFor>
    <versioningSettings>
        <enabled>true</enabled>
        <prune>true</prune>
        <pruneDays>10</pruneDays>
    </versioningSettings>
    <searchEnabled>true</searchEnabled>
    <indexingEnabled>true</indexingEnabled>
    <customMetadataIndexingEnabled>true</customMetadataIndexingEnabled>
    <replicationEnabled>true</replicationEnabled>
    <readFromReplica>true</readFromReplica>
    <serviceRemoteSystemRequests>true</serviceRemoteSystemRequests>
 */
func resourceNamespace() *schema.Resource {
	return &schema.Resource{
		Create: resourceNamespaceCreate,
		Read:   resourceNamespaceRead,
		Update: resourceNamespaceUpdate,
		Delete: resourceNamespaceDelete,
		Exists: resourceNamespaceExists,
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: false, //
			},
			"hash_scheme": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"enterprise_mode": &schema.Schema{
				Type:     schema.TypeBool,
				Required: true,
			},
			"hard_quota": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"soft_quota": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
			},
			"replication_enabled": &schema.Schema{
				Type:     schema.TypeBool,
				Required: true,
			},
		},
	}
}

func resourceNamespaceCreate(d *schema.ResourceData, m interface{}) error {
	hcpClient := hcpClient(m)

	name := d.Get("name").(string)
	hashScheme := d.Get("hash_scheme").(string)
	enterpriseMode := d.Get("enterprise_mode").(bool)
	hardQuota := d.Get("hard_quota").(string)
	softQuota := d.Get("soft_quota").(int)
	replicationEnabled := d.Get("replication_enabled").(bool)

	namespace := &hcp.Namespace{
		Name:               name,
		HashScheme:         hashScheme,
		EnterpriseMode:     enterpriseMode,
		HardQuota:          hardQuota,
		SoftQuota:          softQuota,
		ReplicationEnabled: replicationEnabled,
	}

	if err := hcpClient.CreateNamespace(namespace); err == nil {
		d.SetId(name)
		return nil
	} else {
		return err
	}

}

func resourceNamespaceRead(d *schema.ResourceData, m interface{}) error {
	return nil

}

func resourceNamespaceUpdate(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceNamespaceDelete(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceNamespaceExists(d *schema.ResourceData, meta interface{}) (bool, error) {
	return false, nil
}
