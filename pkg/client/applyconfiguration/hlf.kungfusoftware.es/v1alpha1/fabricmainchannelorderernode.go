/*
 * Copyright Kungfusoftware.es. All Rights Reserved.
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package v1alpha1

// FabricMainChannelOrdererNodeApplyConfiguration represents an declarative configuration of the FabricMainChannelOrdererNode type for use
// with apply.
type FabricMainChannelOrdererNodeApplyConfiguration struct {
	Name      *string `json:"name,omitempty"`
	Namespace *string `json:"namespace,omitempty"`
}

// FabricMainChannelOrdererNodeApplyConfiguration constructs an declarative configuration of the FabricMainChannelOrdererNode type for use with
// apply.
func FabricMainChannelOrdererNode() *FabricMainChannelOrdererNodeApplyConfiguration {
	return &FabricMainChannelOrdererNodeApplyConfiguration{}
}

// WithName sets the Name field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Name field is set to the value of the last call.
func (b *FabricMainChannelOrdererNodeApplyConfiguration) WithName(value string) *FabricMainChannelOrdererNodeApplyConfiguration {
	b.Name = &value
	return b
}

// WithNamespace sets the Namespace field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Namespace field is set to the value of the last call.
func (b *FabricMainChannelOrdererNodeApplyConfiguration) WithNamespace(value string) *FabricMainChannelOrdererNodeApplyConfiguration {
	b.Namespace = &value
	return b
}
