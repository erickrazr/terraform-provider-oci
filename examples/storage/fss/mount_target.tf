// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

resource "oci_file_storage_mount_target" "my_mount_target_1" {
  #Required
  availability_domain = data.oci_identity_availability_domain.ad.name
  compartment_id      = var.compartment_ocid
  subnet_id           = oci_core_subnet.my_subnet.id

  #Optional
  display_name = var.mount_target_1_display_name
  # defined_tags = {
  #   "example-tag-namespace-all.example-tag" = "value"
  # }

  freeform_tags = {
    "Department" = "Finance"
    yor_trace    = "99954674-e499-4591-9191-50fdf84a7e75"
  }

  nsg_ids = [oci_core_network_security_group.test_network_security_group.id]
}

resource "oci_file_storage_mount_target" "my_mount_target_2" {
  #Required
  availability_domain = data.oci_identity_availability_domain.ad.name
  compartment_id      = var.compartment_ocid
  subnet_id           = oci_core_subnet.my_subnet.id

  #Optional
  display_name = var.mount_target_2_display_name
  # defined_tags = {
  #   "example-tag-namespace-all.example-tag" = "value"
  # }

  freeform_tags = {
    "Department" = "Accounting"
    yor_trace    = "3f4493c6-bbb6-4218-a76c-5c28688f8680"
  }

  nsg_ids = [oci_core_network_security_group.test_network_security_group.id]
}

# Use export_set.tf config to update the size for a mount target
