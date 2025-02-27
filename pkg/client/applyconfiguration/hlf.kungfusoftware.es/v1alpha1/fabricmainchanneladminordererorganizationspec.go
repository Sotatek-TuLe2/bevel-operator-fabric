/*
 * Copyright Kungfusoftware.es. All Rights Reserved.
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package v1alpha1

// FabricMainChannelAdminOrdererOrganizationSpecApplyConfiguration represents an declarative configuration of the FabricMainChannelAdminOrdererOrganizationSpec type for use
// with apply.
type FabricMainChannelAdminOrdererOrganizationSpecApplyConfiguration struct {
	MSPID *string `json:"mspID,omitempty"`
}

// FabricMainChannelAdminOrdererOrganizationSpecApplyConfiguration constructs an declarative configuration of the FabricMainChannelAdminOrdererOrganizationSpec type for use with
// apply.
func FabricMainChannelAdminOrdererOrganizationSpec() *FabricMainChannelAdminOrdererOrganizationSpecApplyConfiguration {
	return &FabricMainChannelAdminOrdererOrganizationSpecApplyConfiguration{}
}

// WithMSPID sets the MSPID field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the MSPID field is set to the value of the last call.
func (b *FabricMainChannelAdminOrdererOrganizationSpecApplyConfiguration) WithMSPID(value string) *FabricMainChannelAdminOrdererOrganizationSpecApplyConfiguration {
	b.MSPID = &value
	return b
}
