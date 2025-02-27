/*
 * Copyright Kungfusoftware.es. All Rights Reserved.
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package v1alpha1

// FabricPeerDiscoveryApplyConfiguration represents an declarative configuration of the FabricPeerDiscovery type for use
// with apply.
type FabricPeerDiscoveryApplyConfiguration struct {
	Period      *string `json:"period,omitempty"`
	TouchPeriod *string `json:"touchPeriod,omitempty"`
}

// FabricPeerDiscoveryApplyConfiguration constructs an declarative configuration of the FabricPeerDiscovery type for use with
// apply.
func FabricPeerDiscovery() *FabricPeerDiscoveryApplyConfiguration {
	return &FabricPeerDiscoveryApplyConfiguration{}
}

// WithPeriod sets the Period field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Period field is set to the value of the last call.
func (b *FabricPeerDiscoveryApplyConfiguration) WithPeriod(value string) *FabricPeerDiscoveryApplyConfiguration {
	b.Period = &value
	return b
}

// WithTouchPeriod sets the TouchPeriod field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the TouchPeriod field is set to the value of the last call.
func (b *FabricPeerDiscoveryApplyConfiguration) WithTouchPeriod(value string) *FabricPeerDiscoveryApplyConfiguration {
	b.TouchPeriod = &value
	return b
}
