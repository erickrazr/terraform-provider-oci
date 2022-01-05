// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}
variable "compartment_ocid" {}

variable "fleet_description" {
  default = "Example Fleet created by Terraform"
}

variable "fleet_display_name" {
  default = "Example Fleet"
}

variable "fleet_freeform_tags" {
  default = { "bar-key" = "value" }
}

variable "fleet_defined_tags" {
  default = { "example-tag-namespace-all.example-tag" = "value" }
}

variable "fleet_id" {
  default = "id"
}

variable "fleet_state" {
  default = "ACTIVE"
}

provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

resource "oci_jms_fleet" "example_fleet" {
  #Required
  compartment_id = var.compartment_ocid
  display_name   = var.fleet_display_name

  #Optional
  description = var.fleet_description
  freeform_tags = merge(var.fleet_freeform_tags, {
    yor_trace = "53a74db1-308b-489c-a2a6-553c59486f8f"
  })

  # Create the Tag namespace in OCI before enabling
  # See user guide: https://docs.oracle.com/en-us/iaas/Content/Tagging/Tasks/managingtagsandtagnamespaces.htm
  # defined_tags  = var.fleet_defined_tags
}

data "oci_jms_fleets" "example_fleets" {

  #Optional
  compartment_id = var.compartment_ocid
  display_name   = var.fleet_display_name
  id             = var.fleet_id
  state          = var.fleet_state
}