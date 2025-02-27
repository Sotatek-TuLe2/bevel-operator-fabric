/*
 * Copyright Kungfusoftware.es. All Rights Reserved.
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package v1alpha1

import (
	v1 "k8s.io/api/core/v1"
)

// FabricExplorerSpecApplyConfiguration represents an declarative configuration of the FabricExplorerSpec type for use
// with apply.
type FabricExplorerSpecApplyConfiguration struct {
	Resources *v1.ResourceRequirements `json:"resources,omitempty"`
}

// FabricExplorerSpecApplyConfiguration constructs an declarative configuration of the FabricExplorerSpec type for use with
// apply.
func FabricExplorerSpec() *FabricExplorerSpecApplyConfiguration {
	return &FabricExplorerSpecApplyConfiguration{}
}

// WithResources sets the Resources field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Resources field is set to the value of the last call.
func (b *FabricExplorerSpecApplyConfiguration) WithResources(value v1.ResourceRequirements) *FabricExplorerSpecApplyConfiguration {
	b.Resources = &value
	return b
}
