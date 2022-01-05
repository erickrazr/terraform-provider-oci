// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0


variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}
variable "managed_agent_id" {}
variable "compartment_ocid" {}

provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

resource "oci_identity_tag_namespace" "tag-namespace1" {
  compartment_id = var.tenancy_ocid
  description    = "example tag namespace"
  name           = "examples-tag-namespace-all"
  is_retired     = false
}

resource "oci_identity_tag" "tag1" {
  description      = "example tag"
  name             = "example-tag"
  tag_namespace_id = oci_identity_tag_namespace.tag-namespace1.id
  is_cost_tracking = false
  is_retired       = false
}

variable "host_insight_defined_tags_value" {
  default = "value"
}

variable "host_insight_entity_source" {
  default = "MACS_MANAGED_EXTERNAL_HOST"
}

variable "host_insight_freeform_tags" {
  default = { "bar-key" = "value" }
}

variable "resource_status" {
  default = "ENABLED"
}

resource "oci_management_agent_management_agent" "test_management_agent" {
  managed_agent_id = var.managed_agent_id
}

// To Create a Host insight
resource "oci_opsi_host_insight" "test_host_insight" {
  compartment_id      = var.compartment_ocid
  entity_source       = var.host_insight_entity_source
  management_agent_id = oci_management_agent_management_agent.test_management_agent.id
  defined_tags        = "${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "${var.host_insight_defined_tags_value}")}"
  freeform_tags = merge(var.host_insight_freeform_tags, {
    yor_trace = "6c36b436-5422-4462-a9d2-b8ef57776f6e"
  })
  status = var.resource_status
}


variable "host_insight_host_type" {
  default = ["EXTERNAL-HOST"]
}

variable "host_insight_state" {
  default = ["ACTIVE"]
}

variable "host_insight_status" {
  default = ["ENABLED"]
}

// List host insight
data "oci_opsi_host_insights" "test_host_insights" {
  compartment_id = var.compartment_ocid
  host_type      = var.host_insight_host_type
  state          = var.host_insight_state
  status         = var.host_insight_status
}

// Get an host insight
data "oci_opsi_host_insight" "test_host_insight" {
  host_insight_id = oci_opsi_host_insight.test_host_insight.id
  freeform_tags = {
    yor_trace = "6c36b436-5422-4462-a9d2-b8ef57776f6e"
  }
}

