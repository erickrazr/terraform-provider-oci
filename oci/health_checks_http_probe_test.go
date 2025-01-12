// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	HttpProbeRequiredOnlyResource = HttpProbeResourceDependencies +
		generateResourceFromRepresentationMap("oci_health_checks_http_probe", "test_http_probe", Required, Create, httpProbeRepresentation)

	httpProbeRepresentation = map[string]interface{}{
		"compartment_id":      Representation{repType: Required, create: `${var.compartment_id}`},
		"protocol":            Representation{repType: Required, create: `HTTP`},
		"targets":             Representation{repType: Required, create: []string{`www.oracle.com`}},
		"headers":             Representation{repType: Optional, create: map[string]string{"headers": "headers"}, update: map[string]string{"headers2": "headers2"}},
		"method":              Representation{repType: Optional, create: `GET`},
		"path":                Representation{repType: Optional, create: `/`},
		"port":                Representation{repType: Optional, create: `80`},
		"timeout_in_seconds":  Representation{repType: Optional, create: `10`},
		"vantage_point_names": Representation{repType: Optional, create: []string{`goo-chs`}},
	}

	HttpProbeResourceDependencies = ""
)

// issue-routing-tag: health_checks/default
func TestHealthChecksHttpProbeResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestHealthChecksHttpProbeResource_basic")
	defer httpreplay.SaveScenario()

	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_health_checks_http_probe.test_http_probe"

	var resId string
	// Save TF content to create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	saveConfigContent(config+compartmentIdVariableStr+HttpProbeResourceDependencies+
		generateResourceFromRepresentationMap("oci_health_checks_http_probe", "test_http_probe", Optional, Create, httpProbeRepresentation), "healthchecks", "httpProbe", t)

	ResourceTest(t, nil, []resource.TestStep{
		// verify create
		{
			Config: config + compartmentIdVariableStr + HttpProbeResourceDependencies +
				generateResourceFromRepresentationMap("oci_health_checks_http_probe", "test_http_probe", Required, Create, httpProbeRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "protocol", "HTTP"),
				resource.TestCheckResourceAttr(resourceName, "targets.#", "1"),
			),
		},

		// delete before next create
		{
			Config: config + compartmentIdVariableStr + HttpProbeResourceDependencies,
		},
		// verify create with optionals
		{
			Config: config + compartmentIdVariableStr + HttpProbeResourceDependencies +
				generateResourceFromRepresentationMap("oci_health_checks_http_probe", "test_http_probe", Optional, Create, httpProbeRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "headers.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "method", "GET"),
				resource.TestCheckResourceAttr(resourceName, "path", "/"),
				resource.TestCheckResourceAttr(resourceName, "port", "80"),
				resource.TestCheckResourceAttr(resourceName, "protocol", "HTTP"),
				resource.TestCheckResourceAttr(resourceName, "targets.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "timeout_in_seconds", "10"),
				resource.TestCheckResourceAttr(resourceName, "vantage_point_names.#", "1"),

				func(s *terraform.State) (err error) {
					resId, err = fromInstanceState(s, resourceName, "id")
					if isEnableExportCompartment, _ := strconv.ParseBool(getEnvSettingWithDefault("enable_export_compartment", "true")); isEnableExportCompartment {
						if errExport := testExportCompartmentWithResourceName(&resId, &compartmentId, resourceName); errExport != nil {
							return errExport
						}
					}
					return err
				},
			),
		},
	})
}
