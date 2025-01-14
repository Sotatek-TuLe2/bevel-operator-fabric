/*
 * Copyright Kungfusoftware.es. All Rights Reserved.
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package v1alpha1

import (
	v1 "k8s.io/api/core/v1"
)

// FabricFSServerApplyConfiguration represents an declarative configuration of the FabricFSServer type for use
// with apply.
type FabricFSServerApplyConfiguration struct {
	Image      *string        `json:"image,omitempty"`
	Tag        *string        `json:"tag,omitempty"`
	PullPolicy *v1.PullPolicy `json:"pullPolicy,omitempty"`
}

// FabricFSServerApplyConfiguration constructs an declarative configuration of the FabricFSServer type for use with
// apply.
func FabricFSServer() *FabricFSServerApplyConfiguration {
	return &FabricFSServerApplyConfiguration{}
}

// WithImage sets the Image field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Image field is set to the value of the last call.
func (b *FabricFSServerApplyConfiguration) WithImage(value string) *FabricFSServerApplyConfiguration {
	b.Image = &value
	return b
}

// WithTag sets the Tag field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Tag field is set to the value of the last call.
func (b *FabricFSServerApplyConfiguration) WithTag(value string) *FabricFSServerApplyConfiguration {
	b.Tag = &value
	return b
}

// WithPullPolicy sets the PullPolicy field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the PullPolicy field is set to the value of the last call.
func (b *FabricFSServerApplyConfiguration) WithPullPolicy(value v1.PullPolicy) *FabricFSServerApplyConfiguration {
	b.PullPolicy = &value
	return b
}
