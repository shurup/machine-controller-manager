// This file was automatically generated by lister-gen

package internalversion

import (
	machine "github.com/gardener/node-controller-manager/pkg/apis/machine"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// MachineLister helps list Machines.
type MachineLister interface {
	// List lists all Machines in the indexer.
	List(selector labels.Selector) (ret []*machine.Machine, err error)
	// Machines returns an object that can list and get Machines.
	Machines(namespace string) MachineNamespaceLister
	MachineListerExpansion
}

// machineLister implements the MachineLister interface.
type machineLister struct {
	indexer cache.Indexer
}

// NewMachineLister returns a new MachineLister.
func NewMachineLister(indexer cache.Indexer) MachineLister {
	return &machineLister{indexer: indexer}
}

// List lists all Machines in the indexer.
func (s *machineLister) List(selector labels.Selector) (ret []*machine.Machine, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*machine.Machine))
	})
	return ret, err
}

// Machines returns an object that can list and get Machines.
func (s *machineLister) Machines(namespace string) MachineNamespaceLister {
	return machineNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// MachineNamespaceLister helps list and get Machines.
type MachineNamespaceLister interface {
	// List lists all Machines in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*machine.Machine, err error)
	// Get retrieves the Machine from the indexer for a given namespace and name.
	Get(name string) (*machine.Machine, error)
	MachineNamespaceListerExpansion
}

// machineNamespaceLister implements the MachineNamespaceLister
// interface.
type machineNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all Machines in the indexer for a given namespace.
func (s machineNamespaceLister) List(selector labels.Selector) (ret []*machine.Machine, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*machine.Machine))
	})
	return ret, err
}

// Get retrieves the Machine from the indexer for a given namespace and name.
func (s machineNamespaceLister) Get(name string) (*machine.Machine, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(machine.Resource("machine"), name)
	}
	return obj.(*machine.Machine), nil
}
