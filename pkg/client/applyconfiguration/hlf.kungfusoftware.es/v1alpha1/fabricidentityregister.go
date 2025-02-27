/*
 * Copyright Kungfusoftware.es. All Rights Reserved.
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package v1alpha1

// FabricIdentityRegisterApplyConfiguration represents an declarative configuration of the FabricIdentityRegister type for use
// with apply.
type FabricIdentityRegisterApplyConfiguration struct {
	Enrollid       *string  `json:"enrollid,omitempty"`
	Enrollsecret   *string  `json:"enrollsecret,omitempty"`
	Type           *string  `json:"type,omitempty"`
	Affiliation    *string  `json:"affiliation,omitempty"`
	MaxEnrollments *int     `json:"maxenrollments,omitempty"`
	Attrs          []string `json:"attrs,omitempty"`
}

// FabricIdentityRegisterApplyConfiguration constructs an declarative configuration of the FabricIdentityRegister type for use with
// apply.
func FabricIdentityRegister() *FabricIdentityRegisterApplyConfiguration {
	return &FabricIdentityRegisterApplyConfiguration{}
}

// WithEnrollid sets the Enrollid field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Enrollid field is set to the value of the last call.
func (b *FabricIdentityRegisterApplyConfiguration) WithEnrollid(value string) *FabricIdentityRegisterApplyConfiguration {
	b.Enrollid = &value
	return b
}

// WithEnrollsecret sets the Enrollsecret field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Enrollsecret field is set to the value of the last call.
func (b *FabricIdentityRegisterApplyConfiguration) WithEnrollsecret(value string) *FabricIdentityRegisterApplyConfiguration {
	b.Enrollsecret = &value
	return b
}

// WithType sets the Type field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Type field is set to the value of the last call.
func (b *FabricIdentityRegisterApplyConfiguration) WithType(value string) *FabricIdentityRegisterApplyConfiguration {
	b.Type = &value
	return b
}

// WithAffiliation sets the Affiliation field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Affiliation field is set to the value of the last call.
func (b *FabricIdentityRegisterApplyConfiguration) WithAffiliation(value string) *FabricIdentityRegisterApplyConfiguration {
	b.Affiliation = &value
	return b
}

// WithMaxEnrollments sets the MaxEnrollments field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the MaxEnrollments field is set to the value of the last call.
func (b *FabricIdentityRegisterApplyConfiguration) WithMaxEnrollments(value int) *FabricIdentityRegisterApplyConfiguration {
	b.MaxEnrollments = &value
	return b
}

// WithAttrs adds the given value to the Attrs field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Attrs field.
func (b *FabricIdentityRegisterApplyConfiguration) WithAttrs(values ...string) *FabricIdentityRegisterApplyConfiguration {
	for i := range values {
		b.Attrs = append(b.Attrs, values[i])
	}
	return b
}
