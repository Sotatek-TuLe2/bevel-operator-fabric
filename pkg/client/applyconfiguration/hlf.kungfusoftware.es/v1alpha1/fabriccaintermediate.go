/*
 * Copyright Kungfusoftware.es. All Rights Reserved.
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package v1alpha1

// FabricCAIntermediateApplyConfiguration represents an declarative configuration of the FabricCAIntermediate type for use
// with apply.
type FabricCAIntermediateApplyConfiguration struct {
	ParentServer *FabricCAIntermediateParentServerApplyConfiguration `json:"parentServer,omitempty"`
}

// FabricCAIntermediateApplyConfiguration constructs an declarative configuration of the FabricCAIntermediate type for use with
// apply.
func FabricCAIntermediate() *FabricCAIntermediateApplyConfiguration {
	return &FabricCAIntermediateApplyConfiguration{}
}

// WithParentServer sets the ParentServer field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ParentServer field is set to the value of the last call.
func (b *FabricCAIntermediateApplyConfiguration) WithParentServer(value *FabricCAIntermediateParentServerApplyConfiguration) *FabricCAIntermediateApplyConfiguration {
	b.ParentServer = value
	return b
}
