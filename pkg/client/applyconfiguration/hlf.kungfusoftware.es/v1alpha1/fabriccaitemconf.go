/*
 * Copyright Kungfusoftware.es. All Rights Reserved.
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package v1alpha1

// FabricCAItemConfApplyConfiguration represents an declarative configuration of the FabricCAItemConf type for use
// with apply.
type FabricCAItemConfApplyConfiguration struct {
	Name         *string                                 `json:"name,omitempty"`
	CFG          *FabricCACFGApplyConfiguration          `json:"cfg,omitempty"`
	Subject      *FabricCASubjectApplyConfiguration      `json:"subject,omitempty"`
	CSR          *FabricCACSRApplyConfiguration          `json:"csr,omitempty"`
	Signing      *FabricCASigningApplyConfiguration      `json:"signing,omitempty"`
	CRL          *FabricCACRLApplyConfiguration          `json:"crl,omitempty"`
	Registry     *FabricCARegistryApplyConfiguration     `json:"registry,omitempty"`
	Intermediate *FabricCAIntermediateApplyConfiguration `json:"intermediate,omitempty"`
	BCCSP        *FabricCABCCSPApplyConfiguration        `json:"bccsp,omitempty"`
	Affiliations []FabricCAAffiliationApplyConfiguration `json:"affiliations,omitempty"`
	CA           *FabricCACryptoApplyConfiguration       `json:"ca,omitempty"`
	TlsCA        *FabricTLSCACryptoApplyConfiguration    `json:"tlsCa,omitempty"`
}

// FabricCAItemConfApplyConfiguration constructs an declarative configuration of the FabricCAItemConf type for use with
// apply.
func FabricCAItemConf() *FabricCAItemConfApplyConfiguration {
	return &FabricCAItemConfApplyConfiguration{}
}

// WithName sets the Name field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Name field is set to the value of the last call.
func (b *FabricCAItemConfApplyConfiguration) WithName(value string) *FabricCAItemConfApplyConfiguration {
	b.Name = &value
	return b
}

// WithCFG sets the CFG field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the CFG field is set to the value of the last call.
func (b *FabricCAItemConfApplyConfiguration) WithCFG(value *FabricCACFGApplyConfiguration) *FabricCAItemConfApplyConfiguration {
	b.CFG = value
	return b
}

// WithSubject sets the Subject field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Subject field is set to the value of the last call.
func (b *FabricCAItemConfApplyConfiguration) WithSubject(value *FabricCASubjectApplyConfiguration) *FabricCAItemConfApplyConfiguration {
	b.Subject = value
	return b
}

// WithCSR sets the CSR field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the CSR field is set to the value of the last call.
func (b *FabricCAItemConfApplyConfiguration) WithCSR(value *FabricCACSRApplyConfiguration) *FabricCAItemConfApplyConfiguration {
	b.CSR = value
	return b
}

// WithSigning sets the Signing field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Signing field is set to the value of the last call.
func (b *FabricCAItemConfApplyConfiguration) WithSigning(value *FabricCASigningApplyConfiguration) *FabricCAItemConfApplyConfiguration {
	b.Signing = value
	return b
}

// WithCRL sets the CRL field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the CRL field is set to the value of the last call.
func (b *FabricCAItemConfApplyConfiguration) WithCRL(value *FabricCACRLApplyConfiguration) *FabricCAItemConfApplyConfiguration {
	b.CRL = value
	return b
}

// WithRegistry sets the Registry field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Registry field is set to the value of the last call.
func (b *FabricCAItemConfApplyConfiguration) WithRegistry(value *FabricCARegistryApplyConfiguration) *FabricCAItemConfApplyConfiguration {
	b.Registry = value
	return b
}

// WithIntermediate sets the Intermediate field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Intermediate field is set to the value of the last call.
func (b *FabricCAItemConfApplyConfiguration) WithIntermediate(value *FabricCAIntermediateApplyConfiguration) *FabricCAItemConfApplyConfiguration {
	b.Intermediate = value
	return b
}

// WithBCCSP sets the BCCSP field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the BCCSP field is set to the value of the last call.
func (b *FabricCAItemConfApplyConfiguration) WithBCCSP(value *FabricCABCCSPApplyConfiguration) *FabricCAItemConfApplyConfiguration {
	b.BCCSP = value
	return b
}

// WithAffiliations adds the given value to the Affiliations field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Affiliations field.
func (b *FabricCAItemConfApplyConfiguration) WithAffiliations(values ...*FabricCAAffiliationApplyConfiguration) *FabricCAItemConfApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithAffiliations")
		}
		b.Affiliations = append(b.Affiliations, *values[i])
	}
	return b
}

// WithCA sets the CA field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the CA field is set to the value of the last call.
func (b *FabricCAItemConfApplyConfiguration) WithCA(value *FabricCACryptoApplyConfiguration) *FabricCAItemConfApplyConfiguration {
	b.CA = value
	return b
}

// WithTlsCA sets the TlsCA field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the TlsCA field is set to the value of the last call.
func (b *FabricCAItemConfApplyConfiguration) WithTlsCA(value *FabricTLSCACryptoApplyConfiguration) *FabricCAItemConfApplyConfiguration {
	b.TlsCA = value
	return b
}
