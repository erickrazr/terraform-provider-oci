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
	adbWalletMgmtDbName = randomString(1, charsetWithoutDigits) + randomString(13, charset)

	AutonomousDatabaseInstanceWalletManagementResourceConfig = AutonomousDatabaseInstanceWalletManagementResourceDependencies +
		generateResourceFromRepresentationMap("oci_database_autonomous_database_instance_wallet_management", "test_autonomous_database_instance_wallet_management", Optional, Update, autonomousDatabaseInstanceWalletManagementRepresentation)

	autonomousDatabaseInstanceWalletManagementSingularDataSourceRepresentation = map[string]interface{}{
		"autonomous_database_id": Representation{repType: Required, create: `${oci_database_autonomous_database.test_autonomous_database.id}`},
	}

	autonomousDatabaseInstanceWalletManagementRepresentation = map[string]interface{}{
		"autonomous_database_id": Representation{repType: Required, create: `${oci_database_autonomous_database.test_autonomous_database.id}`},
		"should_rotate":          Representation{repType: Optional, create: `false`, update: `true`},
	}

	AutonomousDatabaseInstanceWalletManagementResourceDependencies = generateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database", Required, Create,
		getUpdatedRepresentationCopy("db_name", Representation{repType: Required, create: adbWalletMgmtDbName}, autonomousDatabaseRepresentation))
)

// issue-routing-tag: database/dbaas-adb
func TestDatabaseAutonomousDatabaseInstanceWalletManagementResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatabaseAutonomousDatabaseInstanceWalletManagementResource_basic")
	defer httpreplay.SaveScenario()

	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_database_autonomous_database_instance_wallet_management.test_autonomous_database_instance_wallet_management"

	singularDatasourceName := "data.oci_database_autonomous_database_instance_wallet_management.test_autonomous_database_instance_wallet_management"

	var resId, resId2 string
	// Save TF content to create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	saveConfigContent(config+compartmentIdVariableStr+AutonomousDatabaseInstanceWalletManagementResourceDependencies+
		generateResourceFromRepresentationMap("oci_database_autonomous_database_instance_wallet_management", "test_autonomous_database_instance_wallet_management", Optional, Create, autonomousDatabaseInstanceWalletManagementRepresentation), "database", "autonomousDatabaseInstanceWalletManagement", t)

	ResourceTest(t, nil, []resource.TestStep{
		// verify create
		{
			Config: config + compartmentIdVariableStr + AutonomousDatabaseInstanceWalletManagementResourceDependencies +
				generateResourceFromRepresentationMap("oci_database_autonomous_database_instance_wallet_management", "test_autonomous_database_instance_wallet_management", Required, Create, autonomousDatabaseInstanceWalletManagementRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "autonomous_database_id"),
				resource.TestCheckResourceAttr(resourceName, "state", "ACTIVE"),

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

		// verify updates to updatable parameters
		{
			Config: config + compartmentIdVariableStr + AutonomousDatabaseInstanceWalletManagementResourceDependencies +
				generateResourceFromRepresentationMap("oci_database_autonomous_database_instance_wallet_management", "test_autonomous_database_instance_wallet_management", Optional, Update, autonomousDatabaseInstanceWalletManagementRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "autonomous_database_id"),
				resource.TestCheckResourceAttr(resourceName, "state", "ACTIVE"),
				resource.TestCheckResourceAttrSet(resourceName, "time_rotated"),

				func(s *terraform.State) (err error) {
					resId2, err = fromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("Resource recreated when it was supposed to be updated.")
					}
					return err
				},
			),
		},
		// verify singular datasource
		{
			Config: config +
				generateDataSourceFromRepresentationMap("oci_database_autonomous_database_instance_wallet_management", "test_autonomous_database_instance_wallet_management", Required, Create, autonomousDatabaseInstanceWalletManagementSingularDataSourceRepresentation) +
				compartmentIdVariableStr + AutonomousDatabaseInstanceWalletManagementResourceConfig,
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "autonomous_database_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "state", "ACTIVE"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_rotated"),
			),
		},
	})
}
