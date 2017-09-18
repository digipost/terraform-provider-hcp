package hcp

import (
	"github.com/digipost/hcp"
	"github.com/hashicorp/terraform/helper/schema"
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
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     false,
				ValidateFunc: validateNamespaceName,
			},
			"hard_quota": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validateHardQuota,
			},
			"soft_quota": {
				Type:         schema.TypeInt,
				Optional:     true,
				Default:      85,
				ValidateFunc: validateSoftQuota,
			},
			"replication_enabled": {
				Type:     schema.TypeBool,
				Required: true,
			},
			"read_from_replica": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"enterprise_mode": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"optimized_for": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      hcp.CLOUD,
				ValidateFunc: validateOptimizedFor,
			},
			"search_enabled": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"indexing_enabled": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"custom_metadata_indexing_enabled": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"service_remote_system_requests": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"hash_scheme": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      hcp.SHA_512,
				ForceNew:     true,
				ValidateFunc: validateHashScheme,
			},
			"acls_usage": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      hcp.ENFORCED,
				ValidateFunc: validateAclsUsage,
			},
			"owner": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"owner_type": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validateOwnerType,
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
	enterpriseMode := d.Get("enterprise_mode").(bool)
	optimizedFor := d.Get("optimized_for").(string)
	searchEnabled := d.Get("search_enabled").(bool)
	indexingEnabled := d.Get("indexing_enabled").(bool)
	customMetadataIndexingEnabled := d.Get("custom_metadata_indexing_enabled").(bool)
	serviceRemoteSystemRequests := d.Get("service_remote_system_requests").(bool)
	aclsUsage := d.Get("acls_usage").(string)
	owner := d.Get("owner").(string)
	ownerType := d.Get("owner_type").(string)

	hashScheme := d.Get("hash_scheme").(string)

	v := d.Get("versioning_settings").([]interface{})
	versioningSettings := expandVersioningSettings(v)

	namespace := &hcp.Namespace{
		Name:                          name,
		HardQuota:                     hardQuota,
		SoftQuota:                     softQuota,
		ReplicationEnabled:            replicationEnabled,
		ReadFromReplica:               readFromReplica,
		HashScheme:                    hashScheme,
		EnterpriseMode:                enterpriseMode,
		OptimizedFor:                  optimizedFor,
		SearchEnabled:                 searchEnabled,
		IndexingEnabled:               indexingEnabled,
		CustomMetadataIndexingEnabled: customMetadataIndexingEnabled,
		ServiceRemoteSystemRequests:   serviceRemoteSystemRequests,
		AclsUsage:                     aclsUsage,
		OwnerType:                     ownerType,
		VersioningSettings:            versioningSettings,
		Owner:                         owner,
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
	enterpriseMode := d.Get("enterprise_mode").(bool)
	optimizedFor := d.Get("optimized_for").(string)
	searchEnabled := d.Get("search_enabled").(bool)
	indexingEnabled := d.Get("indexing_enabled").(bool)
	customMetadataIndexingEnabled := d.Get("custom_metadata_indexing_enabled").(bool)
	serviceRemoteSystemRequests := d.Get("service_remote_system_requests").(bool)
	aclsUsage := d.Get("acls_usage").(string)
	ownerType := d.Get("owner_type").(string)
	owner := d.Get("owner").(string)

	namespace := &hcp.Namespace{
		Name:                          name,
		HardQuota:                     hardQuota,
		SoftQuota:                     softQuota,
		ReplicationEnabled:            replicationEnabled,
		ReadFromReplica:               readFromReplica,
		EnterpriseMode:                enterpriseMode,
		OptimizedFor:                  optimizedFor,
		SearchEnabled:                 searchEnabled,
		IndexingEnabled:               indexingEnabled,
		CustomMetadataIndexingEnabled: customMetadataIndexingEnabled,
		ServiceRemoteSystemRequests:   serviceRemoteSystemRequests,
		OwnerType:                     ownerType,
		Owner:                         owner,
		AclsUsage:                     aclsUsage,
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
	if namespace, err := hcpClient(m).ReadNamespace(name); err == nil {
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

func expandVersioningSettings(config []interface{}) *hcp.VersioningSettings {
	versioningSettingsConfig := config[0].(map[string]interface{})

	return &hcp.VersioningSettings{
		Enabled:   versioningSettingsConfig["enabled"].(bool),
		Prune:     versioningSettingsConfig["prune"].(bool),
		PruneDays: versioningSettingsConfig["prune_days"].(int),
	}
}
