/*
Copyright The Kubernetes Authors.

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "../pkg/apis/nodehealth/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// NodeHealthLister helps list NodeHealths.
// All objects returned here must be treated as read-only.
type NodeHealthLister interface {
	// List lists all NodeHealths in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.NodeHealth, err error)
	// NodeHealths returns an object that can list and get NodeHealths.
	NodeHealths(namespace string) NodeHealthNamespaceLister
	NodeHealthListerExpansion
}

// nodeHealthLister implements the NodeHealthLister interface.
type nodeHealthLister struct {
	indexer cache.Indexer
}

// NewNodeHealthLister returns a new NodeHealthLister.
func NewNodeHealthLister(indexer cache.Indexer) NodeHealthLister {
	return &nodeHealthLister{indexer: indexer}
}

// List lists all NodeHealths in the indexer.
func (s *nodeHealthLister) List(selector labels.Selector) (ret []*v1alpha1.NodeHealth, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.NodeHealth))
	})
	return ret, err
}

// NodeHealths returns an object that can list and get NodeHealths.
func (s *nodeHealthLister) NodeHealths(namespace string) NodeHealthNamespaceLister {
	return nodeHealthNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// NodeHealthNamespaceLister helps list and get NodeHealths.
// All objects returned here must be treated as read-only.
type NodeHealthNamespaceLister interface {
	// List lists all NodeHealths in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.NodeHealth, err error)
	// Get retrieves the NodeHealth from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.NodeHealth, error)
	NodeHealthNamespaceListerExpansion
}

// nodeHealthNamespaceLister implements the NodeHealthNamespaceLister
// interface.
type nodeHealthNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all NodeHealths in the indexer for a given namespace.
func (s nodeHealthNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.NodeHealth, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.NodeHealth))
	})
	return ret, err
}

// Get retrieves the NodeHealth from the indexer for a given namespace and name.
func (s nodeHealthNamespaceLister) Get(name string) (*v1alpha1.NodeHealth, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("nodehealth"), name)
	}
	return obj.(*v1alpha1.NodeHealth), nil
}