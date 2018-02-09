// This file was automatically generated by lister-gen

package internalversion

import (
	machine "github.com/gardener/node-controller-manager/pkg/apis/machine"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// AWSMachineClassLister helps list AWSMachineClasses.
type AWSMachineClassLister interface {
	// List lists all AWSMachineClasses in the indexer.
	List(selector labels.Selector) (ret []*machine.AWSMachineClass, err error)
	// AWSMachineClasses returns an object that can list and get AWSMachineClasses.
	AWSMachineClasses(namespace string) AWSMachineClassNamespaceLister
	AWSMachineClassListerExpansion
}

// aWSMachineClassLister implements the AWSMachineClassLister interface.
type aWSMachineClassLister struct {
	indexer cache.Indexer
}

// NewAWSMachineClassLister returns a new AWSMachineClassLister.
func NewAWSMachineClassLister(indexer cache.Indexer) AWSMachineClassLister {
	return &aWSMachineClassLister{indexer: indexer}
}

// List lists all AWSMachineClasses in the indexer.
func (s *aWSMachineClassLister) List(selector labels.Selector) (ret []*machine.AWSMachineClass, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*machine.AWSMachineClass))
	})
	return ret, err
}

// AWSMachineClasses returns an object that can list and get AWSMachineClasses.
func (s *aWSMachineClassLister) AWSMachineClasses(namespace string) AWSMachineClassNamespaceLister {
	return aWSMachineClassNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// AWSMachineClassNamespaceLister helps list and get AWSMachineClasses.
type AWSMachineClassNamespaceLister interface {
	// List lists all AWSMachineClasses in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*machine.AWSMachineClass, err error)
	// Get retrieves the AWSMachineClass from the indexer for a given namespace and name.
	Get(name string) (*machine.AWSMachineClass, error)
	AWSMachineClassNamespaceListerExpansion
}

// aWSMachineClassNamespaceLister implements the AWSMachineClassNamespaceLister
// interface.
type aWSMachineClassNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all AWSMachineClasses in the indexer for a given namespace.
func (s aWSMachineClassNamespaceLister) List(selector labels.Selector) (ret []*machine.AWSMachineClass, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*machine.AWSMachineClass))
	})
	return ret, err
}

// Get retrieves the AWSMachineClass from the indexer for a given namespace and name.
func (s aWSMachineClassNamespaceLister) Get(name string) (*machine.AWSMachineClass, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(machine.Resource("awsmachineclass"), name)
	}
	return obj.(*machine.AWSMachineClass), nil
}
