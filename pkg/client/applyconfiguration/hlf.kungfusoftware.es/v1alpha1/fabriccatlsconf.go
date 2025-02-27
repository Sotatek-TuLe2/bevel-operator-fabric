/*
 * Copyright Kungfusoftware.es. All Rights Reserved.
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package v1alpha1

// FabricCATLSConfApplyConfiguration represents an declarative configuration of the FabricCATLSConf type for use
// with apply.
type FabricCATLSConfApplyConfiguration struct {
	Subject *FabricCASubjectApplyConfiguration `json:"subject,omitempty"`
}

// FabricCATLSConfApplyConfiguration constructs an declarative configuration of the FabricCATLSConf type for use with
// apply.
func FabricCATLSConf() *FabricCATLSConfApplyConfiguration {
	return &FabricCATLSConfApplyConfiguration{}
}

// WithSubject sets the Subject field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Subject field is set to the value of the last call.
func (b *FabricCATLSConfApplyConfiguration) WithSubject(value *FabricCASubjectApplyConfiguration) *FabricCATLSConfApplyConfiguration {
	b.Subject = value
	return b
}
