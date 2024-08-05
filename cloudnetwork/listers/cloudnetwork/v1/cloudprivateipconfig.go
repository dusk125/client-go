// Code generated by lister-gen. DO NOT EDIT.

package v1

import (
	v1 "github.com/openshift/api/cloudnetwork/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/listers"
	"k8s.io/client-go/tools/cache"
)

// CloudPrivateIPConfigLister helps list CloudPrivateIPConfigs.
// All objects returned here must be treated as read-only.
type CloudPrivateIPConfigLister interface {
	// List lists all CloudPrivateIPConfigs in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.CloudPrivateIPConfig, err error)
	// Get retrieves the CloudPrivateIPConfig from the index for a given name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1.CloudPrivateIPConfig, error)
	CloudPrivateIPConfigListerExpansion
}

// cloudPrivateIPConfigLister implements the CloudPrivateIPConfigLister interface.
type cloudPrivateIPConfigLister struct {
	listers.ResourceIndexer[*v1.CloudPrivateIPConfig]
}

// NewCloudPrivateIPConfigLister returns a new CloudPrivateIPConfigLister.
func NewCloudPrivateIPConfigLister(indexer cache.Indexer) CloudPrivateIPConfigLister {
	return &cloudPrivateIPConfigLister{listers.New[*v1.CloudPrivateIPConfig](indexer, v1.Resource("cloudprivateipconfig"))}
}
