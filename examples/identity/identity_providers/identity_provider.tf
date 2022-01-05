// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

resource "oci_identity_identity_provider" "test_identity_provider" {
  #Required
  compartment_id = var.tenancy_ocid
  description    = var.identity_provider_description
  metadata       = var.identity_provider_metadata != "" ? var.identity_provider_metadata : file(var.identity_provider_metadata_file)
  metadata_url   = var.identity_provider_metadata_url
  name           = var.identity_provider_name
  product_type   = var.identity_provider_product_type
  protocol       = var.identity_provider_protocol

  #Optional
  freeform_tags = merge(var.identity_provider_freeform_tags, {
    yor_trace = "b6cf86e7-4fce-41de-aad5-d3bb7c802350"
  })
}

data "oci_identity_identity_providers" "test_identity_providers" {
  #Required
  compartment_id = var.tenancy_ocid
  protocol       = var.identity_provider_protocol
}

output "identity_providers" {
  value = data.oci_identity_identity_providers.test_identity_providers.identity_providers
}

