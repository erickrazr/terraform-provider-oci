---
subcategory: "Database Migration"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_migration_connection"
sidebar_current: "docs-oci-datasource-database_migration-connection"
description: |-
  Provides details about a specific Connection in Oracle Cloud Infrastructure Database Migration service
---

# Data Source: oci_database_migration_connection
This data source provides details about a specific Connection resource in Oracle Cloud Infrastructure Database Migration service.

Note: Deprecated. This resource will be deprecated. A new resource model will be provided soon.
Display Database Connection details.


## Example Usage

```hcl
data "oci_database_migration_connection" "test_connection" {
	#Required
	connection_id = oci_database_migration_connection.test_connection.id
}
```

## Argument Reference

The following arguments are supported:

* `connection_id` - (Required) The OCID of the database connection 


## Attributes Reference

The following attributes are exported:

* `admin_credentials` - Note: Deprecated. Use the new resource model APIs instead. Database Administrator Credentials details. 
	* `username` - Administrator username 
* `certificate_tdn` - This name is the distinguished name used while creating the certificate on target database. 
* `compartment_id` - OCID of the compartment 
* `connect_descriptor` - Note: Deprecated. Use the new resource model APIs instead. Connect Descriptor details. 
	* `connect_string` - Connect string. 
	* `database_service_name` - Database service name. 
	* `host` - Host of the connect descriptor. 
	* `port` - Port of the connect descriptor. 
* `credentials_secret_id` - OCID of the Secret in the Oracle Cloud Infrastructure vault containing the Database Connection credentials. 
* `database_id` - The OCID of the cloud database. 
* `database_type` - Database connection type. 
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `display_name` - Database Connection display name identifier. 
* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `id` - The OCID of the resource 
* `lifecycle_details` - A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in Failed state. 
* `private_endpoint` - Note: Deprecated. Use the new resource model APIs instead. Oracle Cloud Infrastructure Private Endpoint configuration details. 
	* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment to contain the private endpoint. 
	* `id` - [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of a previously created Private Endpoint. 
	* `subnet_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the customer's subnet where the private endpoint VNIC will reside. 
	* `vcn_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the VCN where the Private Endpoint will be bound to. 
* `ssh_details` - Note: Deprecated. Use the new resource model APIs instead. Details of the SSH key that will be used. 
	* `host` - Name of the host the SSH key is valid for. 
	* `sudo_location` - Sudo location 
	* `user` - SSH user 
* `state` - The current state of the Connection resource. 
* `system_tags` - Usage of system tag keys. These predefined keys are scoped to namespaces. Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_created` - The time the Connection resource was created. An RFC3339 formatted datetime string. 
* `time_updated` - The time of the last Connection resource details update. An RFC3339 formatted datetime string. 
* `vault_details` - Note: Deprecated. Use the new resource model APIs instead. Oracle Cloud Infrastructure Vault details to store migration and connection credentials secrets 
	* `compartment_id` - OCID of the compartment where the secret containing the credentials will be created. 
	* `key_id` - OCID of the vault encryption key 
	* `vault_id` - OCID of the vault 

