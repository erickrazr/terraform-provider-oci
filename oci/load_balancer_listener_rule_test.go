// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	listenerRuleDataSourceRepresentation = map[string]interface{}{
		"listener_name":    Representation{repType: Required, create: `${oci_load_balancer_listener.test_listener.name}`},
		"load_balancer_id": Representation{repType: Required, create: `${oci_load_balancer_load_balancer.test_load_balancer.id}`},
	}

	ListenerRuleResourceConfig = generateResourceFromRepresentationMap("oci_load_balancer_backend_set", "test_backend_set", Required, Create, backendSetRepresentation) +
		generateResourceFromRepresentationMap("oci_load_balancer_certificate", "test_certificate", Required, Create, certificateRepresentation) +
		generateResourceFromRepresentationMap("oci_load_balancer_listener", "test_listener", Required, Create, representationCopyWithNewProperties(listenerRepresentation, map[string]interface{}{
			"rule_set_names": Representation{repType: Required, create: []string{`${oci_load_balancer_rule_set.test_rule_set.name}`}},
		})) +
		generateResourceFromRepresentationMap("oci_load_balancer_load_balancer", "test_load_balancer", Required, Create, loadBalancerRepresentation) +
		generateResourceFromRepresentationMap("oci_load_balancer_rule_set", "test_rule_set", Required, Create, ruleSetRepresentation) +
		LoadBalancerSubnetDependencies
)

// issue-routing-tag: load_balancer/default
func TestLoadBalancerListenerRuleResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestLoadBalancerListenerRuleResource_basic")
	defer httpreplay.SaveScenario()

	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_load_balancer_listener_rules.test_listener_rules"

	saveConfigContent("", "", "", t)

	ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				generateDataSourceFromRepresentationMap("oci_load_balancer_listener_rules", "test_listener_rules", Required, Create, listenerRuleDataSourceRepresentation) +
				compartmentIdVariableStr + ListenerRuleResourceConfig,
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "listener_name"),
				resource.TestCheckResourceAttrSet(datasourceName, "load_balancer_id"),

				resource.TestCheckResourceAttrSet(datasourceName, "listener_rules.#"),
				resource.TestCheckResourceAttrSet(datasourceName, "listener_rules.0.name"),
				resource.TestCheckResourceAttr(datasourceName, "listener_rules.0.rule.#", "1"),
			),
		},
	})
}
