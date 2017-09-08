package hcp

import (
	"github.com/digipost/hcp"
	"github.com/hashicorp/terraform/helper/schema"
)

const (
	defaultHashScheme                    = hcp.SHA_512
	defaultEnterpriseMode                = true
	defaultOptimizedFor                  = hcp.CLOUD
	defaultSearchEnabled                 = false
	defaultIndexingEnabled               = false
	defaultCustomMetadataIndexingEnabled = false
	deafultServiceRemoteSystemRequests   = false
)

func resourceNamespace() *schema.Resource {
	return &schema.Resource{
		Create: resourceNamespaceCreate,
		Read:   resourceNamespaceRead,
		Update: resourceNamespaceUpdate,
		Delete: resourceNamespaceDelete,
		Exists: resourceNamespaceExists,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: false,
			},
			"hard_quota": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validateHardQuota,
			},
			"soft_quota": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"replication_enabled": {
				Type:     schema.TypeBool,
				Required: true,
			},
			"read_from_replica": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"versioning_settings": {
				Type:     schema.TypeList,
				Required: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"enabled": {
							Type:     schema.TypeBool,
							Required: true,
						},
						"prune": {
							Type:     schema.TypeBool,
							Optional: true,
							Default:  false,
						},
						"prune_days": {
							Type:     schema.TypeInt,
							Optional: true,
							Default:  10,
						},
					},
				},
			},
		},
	}
}

func resourceNamespaceCreate(d *schema.ResourceData, m interface{}) error {

	name := d.Get("name").(string)
	hardQuota := d.Get("hard_quota").(string)
	softQuota := d.Get("soft_quota").(int)
	replicationEnabled := d.Get("replication_enabled").(bool)
	readFromReplica := d.Get("read_from_replica").(bool)

	v := d.Get("versioning_settings").([]interface{})
	versioningSettings := expandVersioningSettings(v)

	namespace := &hcp.Namespace{
		Name:                          name,
		HardQuota:                     hardQuota,
		SoftQuota:                     softQuota,
		ReplicationEnabled:            replicationEnabled,
		ReadFromReplica:               readFromReplica,
		HashScheme:                    defaultHashScheme,
		EnterpriseMode:                defaultEnterpriseMode,
		OptimizedFor:                  defaultOptimizedFor,
		SearchEnabled:                 defaultSearchEnabled,
		IndexingEnabled:               defaultIndexingEnabled,
		CustomMetadataIndexingEnabled: defaultCustomMetadataIndexingEnabled,
		ServiceRemoteSystemRequests:   deafultServiceRemoteSystemRequests,
		VersioningSettings:            versioningSettings,
	}

	if err := hcpClient(m).CreateNamespace(namespace); err == nil {
		d.SetId(name)
		return nil
	} else {
		return err
	}

}

func resourceNamespaceUpdate(d *schema.ResourceData, m interface{}) error {

	name := d.Get("name").(string)
	hardQuota := d.Get("hard_quota").(string)
	softQuota := d.Get("soft_quota").(int)
	replicationEnabled := d.Get("replication_enabled").(bool)
	readFromReplica := d.Get("read_from_replica").(bool)

	namespace := &hcp.Namespace{
		Name:                          name,
		HardQuota:                     hardQuota,
		SoftQuota:                     softQuota,
		ReplicationEnabled:            replicationEnabled,
		ReadFromReplica:               readFromReplica,
		EnterpriseMode:                defaultEnterpriseMode,
		OptimizedFor:                  defaultOptimizedFor,
		SearchEnabled:                 defaultSearchEnabled,
		IndexingEnabled:               defaultIndexingEnabled,
		CustomMetadataIndexingEnabled: defaultCustomMetadataIndexingEnabled,
		ServiceRemoteSystemRequests:   deafultServiceRemoteSystemRequests,
	}

	if err := hcpClient(m).UpdateNamespace(namespace); err == nil {
		d.SetId(name)
		return nil
	} else {
		return err
	}
}

func resourceNamespaceRead(d *schema.ResourceData, m interface{}) error {

	name := d.Get("name").(string)
	if namespace, err := hcpClient(m).Namespace(name); err == nil {
		// TODO?
		d.Set("soft_quota", namespace.SoftQuota)
		return nil
	} else {
		return err
	}

}

func resourceNamespaceDelete(d *schema.ResourceData, m interface{}) error {
	name := d.Get("name").(string)
	return hcpClient(m).DeleteNamespace(name)
}

func resourceNamespaceExists(d *schema.ResourceData, m interface{}) (bool, error) {
	name := d.Get("name").(string)
	return hcpClient(m).NamespaceExists(name)
}

func expandVersioningSettings(config []interface{}) hcp.VersioningSettings {
	versioningSettingsConfig := config[0].(map[string]interface{})

	return hcp.VersioningSettings{
		Enabled:   versioningSettingsConfig["enabled"].(bool),
		Prune:     versioningSettingsConfig["prune"].(bool),
		PruneDays: versioningSettingsConfig["prune_days"].(int),
	}
}
