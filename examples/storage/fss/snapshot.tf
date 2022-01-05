// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

resource "oci_file_storage_snapshot" "my_snapshot" {
  #Required
  file_system_id = oci_file_storage_file_system.my_fs_1.id
  name           = var.snapshot_name
  # defined_tags = {
  #   "example-tag-namespace-all.example-tag" = "value"
  # }

  freeform_tags = {
    "Department" = "Finance"
    yor_trace    = "331dcbe8-3918-4a8b-9668-604c415f9f42"
  }
}

