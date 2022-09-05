// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1beta1

import (
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// HelmChartRepositoryStatusApplyConfiguration represents an declarative configuration of the HelmChartRepositoryStatus type for use
// with apply.
type HelmChartRepositoryStatusApplyConfiguration struct {
	Conditions []v1.Condition `json:"conditions,omitempty"`
}

// HelmChartRepositoryStatusApplyConfiguration constructs an declarative configuration of the HelmChartRepositoryStatus type for use with
// apply.
func HelmChartRepositoryStatus() *HelmChartRepositoryStatusApplyConfiguration {
	return &HelmChartRepositoryStatusApplyConfiguration{}
}

// WithConditions adds the given value to the Conditions field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Conditions field.
func (b *HelmChartRepositoryStatusApplyConfiguration) WithConditions(values ...v1.Condition) *HelmChartRepositoryStatusApplyConfiguration {
	for i := range values {
		b.Conditions = append(b.Conditions, values[i])
	}
	return b
}