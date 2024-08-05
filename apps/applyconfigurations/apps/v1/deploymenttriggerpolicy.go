// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

import (
	v1 "github.com/openshift/api/apps/v1"
)

// DeploymentTriggerPolicyApplyConfiguration represents a declarative configuration of the DeploymentTriggerPolicy type for use
// with apply.
type DeploymentTriggerPolicyApplyConfiguration struct {
	Type              *v1.DeploymentTriggerType                             `json:"type,omitempty"`
	ImageChangeParams *DeploymentTriggerImageChangeParamsApplyConfiguration `json:"imageChangeParams,omitempty"`
}

// DeploymentTriggerPolicyApplyConfiguration constructs a declarative configuration of the DeploymentTriggerPolicy type for use with
// apply.
func DeploymentTriggerPolicy() *DeploymentTriggerPolicyApplyConfiguration {
	return &DeploymentTriggerPolicyApplyConfiguration{}
}

// WithType sets the Type field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Type field is set to the value of the last call.
func (b *DeploymentTriggerPolicyApplyConfiguration) WithType(value v1.DeploymentTriggerType) *DeploymentTriggerPolicyApplyConfiguration {
	b.Type = &value
	return b
}

// WithImageChangeParams sets the ImageChangeParams field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ImageChangeParams field is set to the value of the last call.
func (b *DeploymentTriggerPolicyApplyConfiguration) WithImageChangeParams(value *DeploymentTriggerImageChangeParamsApplyConfiguration) *DeploymentTriggerPolicyApplyConfiguration {
	b.ImageChangeParams = value
	return b
}
