// Code generated by informer-gen. DO NOT EDIT.

package machine

import (
	internalinterfaces "github.com/gardener/machine-controller-manager/pkg/client/informers/internalversion/internalinterfaces"
	internalversion "github.com/gardener/machine-controller-manager/pkg/client/informers/internalversion/machine/internalversion"
)

// Interface provides access to each of this group's versions.
type Interface interface {
	// InternalVersion provides access to shared informers for resources in InternalVersion.
	InternalVersion() internalversion.Interface
}

type group struct {
	factory          internalinterfaces.SharedInformerFactory
	namespace        string
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// New returns a new Interface.
func New(f internalinterfaces.SharedInformerFactory, namespace string, tweakListOptions internalinterfaces.TweakListOptionsFunc) Interface {
	return &group{factory: f, namespace: namespace, tweakListOptions: tweakListOptions}
}

// InternalVersion returns a new internalversion.Interface.
func (g *group) InternalVersion() internalversion.Interface {
	return internalversion.New(g.factory, g.namespace, g.tweakListOptions)
}
