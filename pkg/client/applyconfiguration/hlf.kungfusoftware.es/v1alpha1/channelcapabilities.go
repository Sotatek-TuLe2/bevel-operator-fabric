/*
 * Copyright Kungfusoftware.es. All Rights Reserved.
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package v1alpha1

// ChannelCapabilitiesApplyConfiguration represents an declarative configuration of the ChannelCapabilities type for use
// with apply.
type ChannelCapabilitiesApplyConfiguration struct {
	V2_0 *bool `json:"V2_0,omitempty"`
}

// ChannelCapabilitiesApplyConfiguration constructs an declarative configuration of the ChannelCapabilities type for use with
// apply.
func ChannelCapabilities() *ChannelCapabilitiesApplyConfiguration {
	return &ChannelCapabilitiesApplyConfiguration{}
}

// WithV2_0 sets the V2_0 field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the V2_0 field is set to the value of the last call.
func (b *ChannelCapabilitiesApplyConfiguration) WithV2_0(value bool) *ChannelCapabilitiesApplyConfiguration {
	b.V2_0 = &value
	return b
}
