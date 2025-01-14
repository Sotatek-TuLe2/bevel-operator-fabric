/*
 * Copyright Kungfusoftware.es. All Rights Reserved.
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package v1alpha1

// CorsApplyConfiguration represents an declarative configuration of the Cors type for use
// with apply.
type CorsApplyConfiguration struct {
	Enabled *bool    `json:"enabled,omitempty"`
	Origins []string `json:"origins,omitempty"`
}

// CorsApplyConfiguration constructs an declarative configuration of the Cors type for use with
// apply.
func Cors() *CorsApplyConfiguration {
	return &CorsApplyConfiguration{}
}

// WithEnabled sets the Enabled field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Enabled field is set to the value of the last call.
func (b *CorsApplyConfiguration) WithEnabled(value bool) *CorsApplyConfiguration {
	b.Enabled = &value
	return b
}

// WithOrigins adds the given value to the Origins field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Origins field.
func (b *CorsApplyConfiguration) WithOrigins(values ...string) *CorsApplyConfiguration {
	for i := range values {
		b.Origins = append(b.Origins, values[i])
	}
	return b
}
