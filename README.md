Terraform HCP Provider
======================

Provider for the [Hitachi Content Platform](https://www.hds.com/en-us/products-solutions/storage/content-platform.html).

Uses the [HCP Management API](https://community.hds.com/docs/DOC-1000121) to create users and namespaces. 

Using the provider
----------------------

### Provider configuration


    provider hcp {
      mapi_url = "https://finance.hcp.example.com:9090/mapi/tenants/finance"
      username = "admin"
      password = "password"
    }

You can also configure the provider using the following environment variables instead:
	
- HCP_MAPI_URL
- HCP_USERNAME
- HCP_PASSWORD

#### hcp_user_account resource

    resource "hcp_user_account" "sftp" {
      username = "username" 
      full_name = "full username"
      password = "password" // sha512 of this will be stored in state file
    }


Requirements
------------

-	[Terraform](https://www.terraform.io/downloads.html) 0.10.2
-	[Go](https://golang.org/doc/install) 1.8 (to build the provider plugin)

Building The Provider
---------------------

```
$ make build
```

Developing the Provider
---------------------------

See [GNUmakefile](GNUmakefile)
