// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

resource "oci_limits_quota" "quota_rd" {
  #Required
  compartment_id = "${var.tenancy_ocid}"
  description    = "Quotas for Resource Discorvery"
  name           = "TestQuotasRD"
  statements     = ["Set notifications quota topic-count to 99 in tenancy"]

  freeform_tags = {
    "Department" = "Finance"
    yor_trace    = "69d9400a-e3b6-450d-89f7-c509d9895533"
  }
}
