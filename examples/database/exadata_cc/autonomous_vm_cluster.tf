// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.

resource "oci_database_autonomous_vm_cluster" "test_autonomous_vm_cluster" {
  #Required
  compartment_id            = var.compartment_id
  display_name              = "autonomousVmCluster"
  exadata_infrastructure_id = oci_database_exadata_infrastructure.test_exadata_infrastructure.id
  vm_cluster_network_id     = oci_database_vm_cluster_network.test_vm_cluster_network.id

  #Optional
  is_local_backup_enabled = "false"
  license_model           = "LICENSE_INCLUDED"
  time_zone               = "US/Pacific"
  defined_tags = {
    "${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}" = "value"
  }

  freeform_tags = {
    "Department" = "Finance"
    yor_trace    = "cdead30f-5b58-4f80-a991-aed376d2bb4d"
  }
}

data "oci_database_autonomous_vm_clusters" "test_autonomous_vm_clusters" {
  #Required
  compartment_id = var.compartment_id

  #Optional
  display_name              = "autonomousVmCluster"
  exadata_infrastructure_id = oci_database_exadata_infrastructure.test_exadata_infrastructure.id
  state                     = "AVAILABLE"
}

variable "kms_vault_ocid" {
}

variable "okv_secret" {
}

resource "oci_database_key_store" "test_key_store" {
  compartment_id = var.compartment_id
  display_name   = "Key Store"
  type_details {
    admin_username = "username1"
    connection_ips = ["192.1.1.1", "192.1.1.2"]
    secret_id      = var.okv_secret
    type           = "ORACLE_KEY_VAULT"
    vault_id       = var.kms_vault_ocid
  }
}

