// Code generated by applyconfiguration-gen. DO NOT EDIT.

package applyconfigurations

import (
	v1 "github.com/openshift/api/securityinternal/v1"
	internal "github.com/openshift/client-go/securityinternal/applyconfigurations/internal"
	securityinternalv1 "github.com/openshift/client-go/securityinternal/applyconfigurations/securityinternal/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	testing "k8s.io/client-go/testing"
)

// ForKind returns an apply configuration type for the given GroupVersionKind, or nil if no
// apply configuration type exists for the given GroupVersionKind.
func ForKind(kind schema.GroupVersionKind) interface{} {
	switch kind {
	// Group=security.internal.openshift.io, Version=v1
	case v1.SchemeGroupVersion.WithKind("RangeAllocation"):
		return &securityinternalv1.RangeAllocationApplyConfiguration{}

	}
	return nil
}

func NewTypeConverter(scheme *runtime.Scheme) *testing.TypeConverter {
	return &testing.TypeConverter{Scheme: scheme, TypeResolver: internal.Parser()}
}
