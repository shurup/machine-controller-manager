package internalversion

import (
	machine "github.com/gardener/node-controller-manager/pkg/apis/machine"
	scheme "github.com/gardener/node-controller-manager/pkg/client/clientset/internalversion/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// AWSMachineClassesGetter has a method to return a AWSMachineClassInterface.
// A group's client should implement this interface.
type AWSMachineClassesGetter interface {
	AWSMachineClasses(namespace string) AWSMachineClassInterface
}

// AWSMachineClassInterface has methods to work with AWSMachineClass resources.
type AWSMachineClassInterface interface {
	Create(*machine.AWSMachineClass) (*machine.AWSMachineClass, error)
	Update(*machine.AWSMachineClass) (*machine.AWSMachineClass, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*machine.AWSMachineClass, error)
	List(opts v1.ListOptions) (*machine.AWSMachineClassList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *machine.AWSMachineClass, err error)
	AWSMachineClassExpansion
}

// aWSMachineClasses implements AWSMachineClassInterface
type aWSMachineClasses struct {
	client rest.Interface
	ns     string
}

// newAWSMachineClasses returns a AWSMachineClasses
func newAWSMachineClasses(c *MachineClient, namespace string) *aWSMachineClasses {
	return &aWSMachineClasses{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the aWSMachineClass, and returns the corresponding aWSMachineClass object, and an error if there is any.
func (c *aWSMachineClasses) Get(name string, options v1.GetOptions) (result *machine.AWSMachineClass, err error) {
	result = &machine.AWSMachineClass{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("awsmachineclasses").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AWSMachineClasses that match those selectors.
func (c *aWSMachineClasses) List(opts v1.ListOptions) (result *machine.AWSMachineClassList, err error) {
	result = &machine.AWSMachineClassList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("awsmachineclasses").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested aWSMachineClasses.
func (c *aWSMachineClasses) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("awsmachineclasses").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a aWSMachineClass and creates it.  Returns the server's representation of the aWSMachineClass, and an error, if there is any.
func (c *aWSMachineClasses) Create(aWSMachineClass *machine.AWSMachineClass) (result *machine.AWSMachineClass, err error) {
	result = &machine.AWSMachineClass{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("awsmachineclasses").
		Body(aWSMachineClass).
		Do().
		Into(result)
	return
}

// Update takes the representation of a aWSMachineClass and updates it. Returns the server's representation of the aWSMachineClass, and an error, if there is any.
func (c *aWSMachineClasses) Update(aWSMachineClass *machine.AWSMachineClass) (result *machine.AWSMachineClass, err error) {
	result = &machine.AWSMachineClass{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("awsmachineclasses").
		Name(aWSMachineClass.Name).
		Body(aWSMachineClass).
		Do().
		Into(result)
	return
}

// Delete takes name of the aWSMachineClass and deletes it. Returns an error if one occurs.
func (c *aWSMachineClasses) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("awsmachineclasses").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *aWSMachineClasses) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("awsmachineclasses").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched aWSMachineClass.
func (c *aWSMachineClasses) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *machine.AWSMachineClass, err error) {
	result = &machine.AWSMachineClass{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("awsmachineclasses").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
