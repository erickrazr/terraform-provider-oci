// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Cloud Advisor API
//
// APIs for managing Cloud Advisor. Cloud Advisor provides recommendations that help you maximize cost savings and improve the security posture of your tenancy.
//

package optimizer

import (
	"github.com/oracle/oci-go-sdk/v48/common"
)

// Strategy The metadata associated with the strategy. The strategy is the method used to apply the recommendation.
type Strategy struct {

	// The name of the strategy.
	StrategyName *string `mandatory:"true" json:"strategyName"`

	// Whether this is the default recommendation strategy.
	IsDefault *bool `mandatory:"true" json:"isDefault"`

	// The list of strategies for the parameters.
	ParametersDefinition []StrategyParameter `mandatory:"false" json:"parametersDefinition"`
}

func (m Strategy) String() string {
	return common.PointerString(m)
}