/*
 * Copyright Kungfusoftware.es. All Rights Reserved.
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package v1alpha1

// FabricCASigningDefaultApplyConfiguration represents an declarative configuration of the FabricCASigningDefault type for use
// with apply.
type FabricCASigningDefaultApplyConfiguration struct {
	Expiry *string  `json:"expiry,omitempty"`
	Usage  []string `json:"usage,omitempty"`
}

// FabricCASigningDefaultApplyConfiguration constructs an declarative configuration of the FabricCASigningDefault type for use with
// apply.
func FabricCASigningDefault() *FabricCASigningDefaultApplyConfiguration {
	return &FabricCASigningDefaultApplyConfiguration{}
}

// WithExpiry sets the Expiry field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Expiry field is set to the value of the last call.
func (b *FabricCASigningDefaultApplyConfiguration) WithExpiry(value string) *FabricCASigningDefaultApplyConfiguration {
	b.Expiry = &value
	return b
}

// WithUsage adds the given value to the Usage field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Usage field.
func (b *FabricCASigningDefaultApplyConfiguration) WithUsage(values ...string) *FabricCASigningDefaultApplyConfiguration {
	for i := range values {
		b.Usage = append(b.Usage, values[i])
	}
	return b
}
