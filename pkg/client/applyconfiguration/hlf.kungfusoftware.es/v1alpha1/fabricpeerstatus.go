/*
 * Copyright Kungfusoftware.es. All Rights Reserved.
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package v1alpha1

import (
	v1alpha1 "github.com/kfsoftware/hlf-operator/api/hlf.kungfusoftware.es/v1alpha1"
	status "github.com/kfsoftware/hlf-operator/pkg/status"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// FabricPeerStatusApplyConfiguration represents an declarative configuration of the FabricPeerStatus type for use
// with apply.
type FabricPeerStatusApplyConfiguration struct {
	Conditions            *status.Conditions         `json:"conditions,omitempty"`
	Message               *string                    `json:"message,omitempty"`
	Status                *v1alpha1.DeploymentStatus `json:"status,omitempty"`
	LastCertificateUpdate *v1.Time                   `json:"lastCertificateUpdate,omitempty"`
	SignCert              *string                    `json:"signCert,omitempty"`
	TlsCert               *string                    `json:"tlsCert,omitempty"`
	TlsCACert             *string                    `json:"tlsCaCert,omitempty"`
	SignCACert            *string                    `json:"signCaCert,omitempty"`
	NodePort              *int                       `json:"port,omitempty"`
}

// FabricPeerStatusApplyConfiguration constructs an declarative configuration of the FabricPeerStatus type for use with
// apply.
func FabricPeerStatus() *FabricPeerStatusApplyConfiguration {
	return &FabricPeerStatusApplyConfiguration{}
}

// WithConditions sets the Conditions field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Conditions field is set to the value of the last call.
func (b *FabricPeerStatusApplyConfiguration) WithConditions(value status.Conditions) *FabricPeerStatusApplyConfiguration {
	b.Conditions = &value
	return b
}

// WithMessage sets the Message field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Message field is set to the value of the last call.
func (b *FabricPeerStatusApplyConfiguration) WithMessage(value string) *FabricPeerStatusApplyConfiguration {
	b.Message = &value
	return b
}

// WithStatus sets the Status field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Status field is set to the value of the last call.
func (b *FabricPeerStatusApplyConfiguration) WithStatus(value v1alpha1.DeploymentStatus) *FabricPeerStatusApplyConfiguration {
	b.Status = &value
	return b
}

// WithLastCertificateUpdate sets the LastCertificateUpdate field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the LastCertificateUpdate field is set to the value of the last call.
func (b *FabricPeerStatusApplyConfiguration) WithLastCertificateUpdate(value v1.Time) *FabricPeerStatusApplyConfiguration {
	b.LastCertificateUpdate = &value
	return b
}

// WithSignCert sets the SignCert field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the SignCert field is set to the value of the last call.
func (b *FabricPeerStatusApplyConfiguration) WithSignCert(value string) *FabricPeerStatusApplyConfiguration {
	b.SignCert = &value
	return b
}

// WithTlsCert sets the TlsCert field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the TlsCert field is set to the value of the last call.
func (b *FabricPeerStatusApplyConfiguration) WithTlsCert(value string) *FabricPeerStatusApplyConfiguration {
	b.TlsCert = &value
	return b
}

// WithTlsCACert sets the TlsCACert field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the TlsCACert field is set to the value of the last call.
func (b *FabricPeerStatusApplyConfiguration) WithTlsCACert(value string) *FabricPeerStatusApplyConfiguration {
	b.TlsCACert = &value
	return b
}

// WithSignCACert sets the SignCACert field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the SignCACert field is set to the value of the last call.
func (b *FabricPeerStatusApplyConfiguration) WithSignCACert(value string) *FabricPeerStatusApplyConfiguration {
	b.SignCACert = &value
	return b
}

// WithNodePort sets the NodePort field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the NodePort field is set to the value of the last call.
func (b *FabricPeerStatusApplyConfiguration) WithNodePort(value int) *FabricPeerStatusApplyConfiguration {
	b.NodePort = &value
	return b
}
