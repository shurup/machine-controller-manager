// This file was automatically generated by lister-gen

package internalversion

import (
	machine "github.com/gardener/node-controller-manager/pkg/apis/machine"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// GCPMachineClassLister helps list GCPMachineClasses.
type GCPMachineClassLister interface {
	// List lists all GCPMachineClasses in the indexer.
	List(selector labels.Selector) (ret []*machine.GCPMachineClass, err error)
	// GCPMachineClasses returns an object that can list and get GCPMachineClasses.
	GCPMachineClasses(namespace string) GCPMachineClassNamespaceLister
	GCPMachineClassListerExpansion
}

// gCPMachineClassLister implements the GCPMachineClassLister interface.
type gCPMachineClassLister struct {
	indexer cache.Indexer
}

// NewGCPMachineClassLister returns a new GCPMachineClassLister.
func NewGCPMachineClassLister(indexer cache.Indexer) GCPMachineClassLister {
	return &gCPMachineClassLister{indexer: indexer}
}

// List lists all GCPMachineClasses in the indexer.
func (s *gCPMachineClassLister) List(selector labels.Selector) (ret []*machine.GCPMachineClass, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*machine.GCPMachineClass))
	})
	return ret, err
}

// GCPMachineClasses returns an object that can list and get GCPMachineClasses.
func (s *gCPMachineClassLister) GCPMachineClasses(namespace string) GCPMachineClassNamespaceLister {
	return gCPMachineClassNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// GCPMachineClassNamespaceLister helps list and get GCPMachineClasses.
type GCPMachineClassNamespaceLister interface {
	// List lists all GCPMachineClasses in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*machine.GCPMachineClass, err error)
	// Get retrieves the GCPMachineClass from the indexer for a given namespace and name.
	Get(name string) (*machine.GCPMachineClass, error)
	GCPMachineClassNamespaceListerExpansion
}

// gCPMachineClassNamespaceLister implements the GCPMachineClassNamespaceLister
// interface.
type gCPMachineClassNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all GCPMachineClasses in the indexer for a given namespace.
func (s gCPMachineClassNamespaceLister) List(selector labels.Selector) (ret []*machine.GCPMachineClass, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*machine.GCPMachineClass))
	})
	return ret, err
}

// Get retrieves the GCPMachineClass from the indexer for a given namespace and name.
func (s gCPMachineClassNamespaceLister) Get(name string) (*machine.GCPMachineClass, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(machine.Resource("gcpmachineclass"), name)
	}
	return obj.(*machine.GCPMachineClass), nil
}
