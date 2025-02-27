/*
 * Copyright Kungfusoftware.es. All Rights Reserved.
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package v1alpha1

// FabricCAAffiliationApplyConfiguration represents an declarative configuration of the FabricCAAffiliation type for use
// with apply.
type FabricCAAffiliationApplyConfiguration struct {
	Name        *string  `json:"name,omitempty"`
	Departments []string `json:"departments,omitempty"`
}

// FabricCAAffiliationApplyConfiguration constructs an declarative configuration of the FabricCAAffiliation type for use with
// apply.
func FabricCAAffiliation() *FabricCAAffiliationApplyConfiguration {
	return &FabricCAAffiliationApplyConfiguration{}
}

// WithName sets the Name field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Name field is set to the value of the last call.
func (b *FabricCAAffiliationApplyConfiguration) WithName(value string) *FabricCAAffiliationApplyConfiguration {
	b.Name = &value
	return b
}

// WithDepartments adds the given value to the Departments field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Departments field.
func (b *FabricCAAffiliationApplyConfiguration) WithDepartments(values ...string) *FabricCAAffiliationApplyConfiguration {
	for i := range values {
		b.Departments = append(b.Departments, values[i])
	}
	return b
}
