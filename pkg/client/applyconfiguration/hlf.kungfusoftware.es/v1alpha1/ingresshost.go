/*
 * Copyright Kungfusoftware.es. All Rights Reserved.
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package v1alpha1

// IngressHostApplyConfiguration represents an declarative configuration of the IngressHost type for use
// with apply.
type IngressHostApplyConfiguration struct {
	Host  *string                         `json:"host,omitempty"`
	Paths []IngressPathApplyConfiguration `json:"paths,omitempty"`
}

// IngressHostApplyConfiguration constructs an declarative configuration of the IngressHost type for use with
// apply.
func IngressHost() *IngressHostApplyConfiguration {
	return &IngressHostApplyConfiguration{}
}

// WithHost sets the Host field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Host field is set to the value of the last call.
func (b *IngressHostApplyConfiguration) WithHost(value string) *IngressHostApplyConfiguration {
	b.Host = &value
	return b
}

// WithPaths adds the given value to the Paths field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Paths field.
func (b *IngressHostApplyConfiguration) WithPaths(values ...*IngressPathApplyConfiguration) *IngressHostApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithPaths")
		}
		b.Paths = append(b.Paths, *values[i])
	}
	return b
}
