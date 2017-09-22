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
				Type:             schema.TypeString,
				Required:         true,
				ValidateFunc:     validateHardQuota,
				DiffSuppressFunc: suppressHardQuotaDiffs,
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
			"http_protocol": {
				Type:     schema.TypeList,
				Required: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"hs3_enabled": {
							Type:     schema.TypeBool,
							Required: true,
						},
						"hs3_requires_authentication": {
							Type:     schema.TypeBool,
							Optional: true,
							Default:  true,
						},
						"http_enabled": {
							Type:     schema.TypeBool,
							Optional: true,
							Default:  false,
						},
						"https_enabled": {
							Type:     schema.TypeBool,
							Optional: true,
							Default:  true,
						},
						"rest_enabled": {
							Type:     schema.TypeBool,
							Optional: true,
							Default:  false,
						},
						"rest_requires_authentication": {
							Type:     schema.TypeBool,
							Optional: true,
							Default:  false,
						},
						"allow_addresses": {
							Type:     schema.TypeList,
							Required: true,
							Elem:     &schema.Schema{Type: schema.TypeString},
						},
						"deny_addresses": {
							Type:     schema.TypeList,
							Required: true,
							Elem:     &schema.Schema{Type: schema.TypeString},
						},
						"allow_if_in_both_lists": {
							Type:     schema.TypeBool,
							Optional: true,
							Default:  false,
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

	versioningSettings := expandVersioningSettings(d.Get("versioning_settings").([]interface{}))

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

	d.Partial(true)

	name := d.Get("name").(string)

	if d.HasChange("namespace") {

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

		if errNamespace := hcpClient(m).UpdateNamespace(namespace); errNamespace == nil {
			d.SetId(name)
			d.SetPartial("namespace")
		} else {
			return errNamespace
		}
	}

	if d.HasChange("http_protocol") {

		httpProtocol := expandHttpProtocol(d.Get("http_protocol").([]interface{}))

		if errProtocolHttp := hcpClient(m).UpdateNamespaceProtocolHttp(name, httpProtocol); errProtocolHttp == nil {
			d.SetPartial("http_protocol")
		} else {
			return errProtocolHttp
		}

	}

	return nil

}

func resourceNamespaceRead(d *schema.ResourceData, m interface{}) error {

	name := d.Get("name").(string)
	if namespace, err := hcpClient(m).ReadNamespace(name); err == nil {
		d.Set("soft_quota", namespace.SoftQuota)
		d.Set("hard_quota", namespace.HardQuota)
		d.Set("replication_enabled", namespace.ReplicationEnabled)
		d.Set("read_from_replica", namespace.ReadFromReplica)
		d.Set("enterprise_mode", namespace.EnterpriseMode)
		d.Set("optimized_for", namespace.OptimizedFor)
		d.Set("search_enabled", namespace.SearchEnabled)
		d.Set("indexing_enabled", namespace.IndexingEnabled)
		d.Set("custom_metadata_indexing_enabled", namespace.CustomMetadataIndexingEnabled)
		d.Set("service_remote_system_requests", namespace.ServiceRemoteSystemRequests)
		d.Set("acls_usage", namespace.AclsUsage)
		d.Set("owner", namespace.Owner)

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

func expandHttpProtocol(config []interface{}) *hcp.HttpProtocol {

	httpProtocolConfig := config[0].(map[string]interface{})

	hs3Enabled := httpProtocolConfig["hs3_enabled"].(bool)
	hs3RequiresAuthentication := httpProtocolConfig["hs3_requires_authentication"].(bool)
	httpEnabled := httpProtocolConfig["http_enabled"].(bool)
	httpsEnabled := httpProtocolConfig["https_enabled"].(bool)
	restEnabled := httpProtocolConfig["rest_enabled"].(bool)
	restRequiresAuthentication := httpProtocolConfig["rest_requires_authentication"].(bool)
	//allowAddresses := httpProtocolConfig["allow_addresses"].([]interface{})
	//denyAddresses := httpProtocolConfig["deny_addresses"].([]interface{})
	allowIfInBothLists := httpProtocolConfig["allow_if_in_both_lists"].(bool)

	httpProtocol := &hcp.HttpProtocol{

		Hs3Enabled:                 hs3Enabled,
		Hs3RequiresAuthentication:  hs3RequiresAuthentication,
		HttpEnabled:                httpEnabled,
		HttpsEnabled:               httpsEnabled,
		RestEnabled:                restEnabled,
		RestRequiresAuthentication: restRequiresAuthentication,
		IpSettings: &hcp.IpSettings{
			//AllowAddresses:     []string{""},
			//DenyAddresses:      []string{""},
			AllowIfInBothLists: allowIfInBothLists,
		},
	}

	return httpProtocol

}

func expandVersioningSettings(config []interface{}) *hcp.VersioningSettings {
	versioningSettingsConfig := config[0].(map[string]interface{})

	return &hcp.VersioningSettings{
		Enabled:   versioningSettingsConfig["enabled"].(bool),
		Prune:     versioningSettingsConfig["prune"].(bool),
		PruneDays: versioningSettingsConfig["prune_days"].(int),
	}
}
