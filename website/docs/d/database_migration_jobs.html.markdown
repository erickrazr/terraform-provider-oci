---
subcategory: "Database Migration"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_migration_jobs"
sidebar_current: "docs-oci-datasource-database_migration-jobs"
description: |-
  Provides the list of Jobs in Oracle Cloud Infrastructure Database Migration service
---

# Data Source: oci_database_migration_jobs
This data source provides the list of Jobs in Oracle Cloud Infrastructure Database Migration service.

Note: Deprecated. This resource will be deprecated. A new resource model will be provided soon.
List all the names of the Migration jobs associated to the specified
migration site.


## Example Usage

```hcl
data "oci_database_migration_jobs" "test_jobs" {
	#Required
	migration_id = oci_database_migration_migration.test_migration.id

	#Optional
	display_name = var.job_display_name
	state = var.job_state
}
```

## Argument Reference

The following arguments are supported:

* `display_name` - (Optional) A filter to return only resources that match the entire display name given. 
* `migration_id` - (Required) The ID of the migration in which to list resources. 
* `state` - (Optional) The lifecycle state of the Migration Job. 


## Attributes Reference

The following attributes are exported:

* `job_collection` - The list of job_collection.

### Job Reference

The following attributes are exported:

* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `display_name` - Name of the job. 
* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `id` - The OCID of the Migration Job. 
* `lifecycle_details` - A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in Failed state. 
* `migration_id` - The OCID of the Migration that this job belongs to. 
* `progress` - Note: Deprecated. Use the new resource model APIs instead. Progress details of a Migration Job. 
	* `current_phase` - Current phase of the job. 
	* `current_status` - Current status of the job. 
	* `phases` - List of phase status for the job. 
		* `duration_in_ms` - Duration of the phase in milliseconds 
		* `name` - Phase name 
		* `progress` - Percent progress of job phase. 
		* `status` - Phase status 
* `state` - The current state of the migration job. 
* `system_tags` - Usage of system tag keys. These predefined keys are scoped to namespaces. Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_created` - The time the Migration Job was created. An RFC3339 formatted datetime string 
* `time_updated` - The time the Migration Job was last updated. An RFC3339 formatted datetime string 
* `type` - The job type. 
* `unsupported_objects` - Database objects not supported. 
	* `object` - Name of the object (regular expression is allowed) 
	* `owner` - Owner of the object (regular expression is allowed) 
	* `type` - Type of unsupported object 

