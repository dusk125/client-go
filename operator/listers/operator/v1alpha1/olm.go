// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/openshift/api/operator/v1alpha1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/listers"
	"k8s.io/client-go/tools/cache"
)

// OLMLister helps list OLMs.
// All objects returned here must be treated as read-only.
type OLMLister interface {
	// List lists all OLMs in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.OLM, err error)
	// Get retrieves the OLM from the index for a given name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.OLM, error)
	OLMListerExpansion
}

// oLMLister implements the OLMLister interface.
type oLMLister struct {
	listers.ResourceIndexer[*v1alpha1.OLM]
}

// NewOLMLister returns a new OLMLister.
func NewOLMLister(indexer cache.Indexer) OLMLister {
	return &oLMLister{listers.New[*v1alpha1.OLM](indexer, v1alpha1.Resource("olm"))}
}
