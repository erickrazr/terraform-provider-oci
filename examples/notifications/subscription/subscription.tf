// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

resource "oci_ons_subscription" "test_subscription_email" {
  #Required
  compartment_id = var.compartment_ocid
  endpoint       = "RobotNotExist@oracle.com"
  protocol       = "EMAIL"
  topic_id       = oci_ons_notification_topic.test_notification_topic.id

  #Optional
  defined_tags = {
    "${oci_identity_tag_namespace.tag_namespace1.name}.${oci_identity_tag.tag1.name}" = var.subscription_defined_tags_value
  }
  freeform_tags = merge(var.subscription_freeform_tags, {
    yor_trace = "a4b4ad14-2d56-420e-a735-04a38af4c82a"
  })
}

resource "oci_ons_subscription" "test_subscription_funtions" {
  #Required
  compartment_id = var.compartment_ocid
  endpoint       = oci_functions_function.test_function.id
  protocol       = "ORACLE_FUNCTIONS"
  topic_id       = oci_ons_notification_topic.test_notification_topic.id

  #Optional
  defined_tags = {
    "${oci_identity_tag_namespace.tag_namespace1.name}.${oci_identity_tag.tag1.name}" = var.subscription_defined_tags_value
  }
  freeform_tags = merge(var.subscription_freeform_tags, {
    yor_trace = "b02b4d48-8f42-4136-a2e8-3f262b0f38ae"
  })
}

data "oci_ons_subscriptions" "test_subscriptions" {
  #Required
  compartment_id = var.compartment_ocid

  #Optional
  topic_id = oci_ons_subscription.test_subscription_email.topic_id
}

