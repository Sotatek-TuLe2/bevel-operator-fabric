/*
 * Copyright Kungfusoftware.es. All Rights Reserved.
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package v1alpha1

// FabricCAMetricsApplyConfiguration represents an declarative configuration of the FabricCAMetrics type for use
// with apply.
type FabricCAMetricsApplyConfiguration struct {
	Provider *string                                  `json:"provider,omitempty"`
	Statsd   *FabricCAMetricsStatsdApplyConfiguration `json:"statsd,omitempty"`
}

// FabricCAMetricsApplyConfiguration constructs an declarative configuration of the FabricCAMetrics type for use with
// apply.
func FabricCAMetrics() *FabricCAMetricsApplyConfiguration {
	return &FabricCAMetricsApplyConfiguration{}
}

// WithProvider sets the Provider field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Provider field is set to the value of the last call.
func (b *FabricCAMetricsApplyConfiguration) WithProvider(value string) *FabricCAMetricsApplyConfiguration {
	b.Provider = &value
	return b
}

// WithStatsd sets the Statsd field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Statsd field is set to the value of the last call.
func (b *FabricCAMetricsApplyConfiguration) WithStatsd(value *FabricCAMetricsStatsdApplyConfiguration) *FabricCAMetricsApplyConfiguration {
	b.Statsd = value
	return b
}
