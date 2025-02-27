/*
 * Copyright Kungfusoftware.es. All Rights Reserved.
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package v1alpha1

// ChannelConfigApplyConfiguration represents an declarative configuration of the ChannelConfig type for use
// with apply.
type ChannelConfigApplyConfiguration struct {
	BatchTimeout            *string                                    `json:"batchTimeout,omitempty"`
	MaxMessageCount         *int                                       `json:"maxMessageCount,omitempty"`
	AbsoluteMaxBytes        *int                                       `json:"absoluteMaxBytes,omitempty"`
	PreferredMaxBytes       *int                                       `json:"preferredMaxBytes,omitempty"`
	OrdererCapabilities     *OrdererCapabilitiesApplyConfiguration     `json:"ordererCapabilities,omitempty"`
	ApplicationCapabilities *ApplicationCapabilitiesApplyConfiguration `json:"applicationCapabilities,omitempty"`
	ChannelCapabilities     *ChannelCapabilitiesApplyConfiguration     `json:"channelCapabilities,omitempty"`
	SnapshotIntervalSize    *int                                       `json:"snapshotIntervalSize,omitempty"`
	TickInterval            *string                                    `json:"tickInterval,omitempty"`
	ElectionTick            *int                                       `json:"electionTick,omitempty"`
	HeartbeatTick           *int                                       `json:"heartbeatTick,omitempty"`
	MaxInflightBlocks       *int                                       `json:"maxInflightBlocks,omitempty"`
}

// ChannelConfigApplyConfiguration constructs an declarative configuration of the ChannelConfig type for use with
// apply.
func ChannelConfig() *ChannelConfigApplyConfiguration {
	return &ChannelConfigApplyConfiguration{}
}

// WithBatchTimeout sets the BatchTimeout field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the BatchTimeout field is set to the value of the last call.
func (b *ChannelConfigApplyConfiguration) WithBatchTimeout(value string) *ChannelConfigApplyConfiguration {
	b.BatchTimeout = &value
	return b
}

// WithMaxMessageCount sets the MaxMessageCount field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the MaxMessageCount field is set to the value of the last call.
func (b *ChannelConfigApplyConfiguration) WithMaxMessageCount(value int) *ChannelConfigApplyConfiguration {
	b.MaxMessageCount = &value
	return b
}

// WithAbsoluteMaxBytes sets the AbsoluteMaxBytes field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the AbsoluteMaxBytes field is set to the value of the last call.
func (b *ChannelConfigApplyConfiguration) WithAbsoluteMaxBytes(value int) *ChannelConfigApplyConfiguration {
	b.AbsoluteMaxBytes = &value
	return b
}

// WithPreferredMaxBytes sets the PreferredMaxBytes field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the PreferredMaxBytes field is set to the value of the last call.
func (b *ChannelConfigApplyConfiguration) WithPreferredMaxBytes(value int) *ChannelConfigApplyConfiguration {
	b.PreferredMaxBytes = &value
	return b
}

// WithOrdererCapabilities sets the OrdererCapabilities field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the OrdererCapabilities field is set to the value of the last call.
func (b *ChannelConfigApplyConfiguration) WithOrdererCapabilities(value *OrdererCapabilitiesApplyConfiguration) *ChannelConfigApplyConfiguration {
	b.OrdererCapabilities = value
	return b
}

// WithApplicationCapabilities sets the ApplicationCapabilities field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ApplicationCapabilities field is set to the value of the last call.
func (b *ChannelConfigApplyConfiguration) WithApplicationCapabilities(value *ApplicationCapabilitiesApplyConfiguration) *ChannelConfigApplyConfiguration {
	b.ApplicationCapabilities = value
	return b
}

// WithChannelCapabilities sets the ChannelCapabilities field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ChannelCapabilities field is set to the value of the last call.
func (b *ChannelConfigApplyConfiguration) WithChannelCapabilities(value *ChannelCapabilitiesApplyConfiguration) *ChannelConfigApplyConfiguration {
	b.ChannelCapabilities = value
	return b
}

// WithSnapshotIntervalSize sets the SnapshotIntervalSize field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the SnapshotIntervalSize field is set to the value of the last call.
func (b *ChannelConfigApplyConfiguration) WithSnapshotIntervalSize(value int) *ChannelConfigApplyConfiguration {
	b.SnapshotIntervalSize = &value
	return b
}

// WithTickInterval sets the TickInterval field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the TickInterval field is set to the value of the last call.
func (b *ChannelConfigApplyConfiguration) WithTickInterval(value string) *ChannelConfigApplyConfiguration {
	b.TickInterval = &value
	return b
}

// WithElectionTick sets the ElectionTick field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ElectionTick field is set to the value of the last call.
func (b *ChannelConfigApplyConfiguration) WithElectionTick(value int) *ChannelConfigApplyConfiguration {
	b.ElectionTick = &value
	return b
}

// WithHeartbeatTick sets the HeartbeatTick field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the HeartbeatTick field is set to the value of the last call.
func (b *ChannelConfigApplyConfiguration) WithHeartbeatTick(value int) *ChannelConfigApplyConfiguration {
	b.HeartbeatTick = &value
	return b
}

// WithMaxInflightBlocks sets the MaxInflightBlocks field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the MaxInflightBlocks field is set to the value of the last call.
func (b *ChannelConfigApplyConfiguration) WithMaxInflightBlocks(value int) *ChannelConfigApplyConfiguration {
	b.MaxInflightBlocks = &value
	return b
}
