/*
Copyright 2023 The ImagineKube Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by client-gen. DO NOT EDIT.

package v1beta1

import (
	"context"
	"time"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
	v1beta1 "imaginekube.com/api/types/v1beta1"
	scheme "imaginekube.com/imaginekube/pkg/client/clientset/versioned/scheme"
)

// FederatedIngressesGetter has a method to return a FederatedIngressInterface.
// A group's client should implement this interface.
type FederatedIngressesGetter interface {
	FederatedIngresses(namespace string) FederatedIngressInterface
}

// FederatedIngressInterface has methods to work with FederatedIngress resources.
type FederatedIngressInterface interface {
	Create(ctx context.Context, federatedIngress *v1beta1.FederatedIngress, opts v1.CreateOptions) (*v1beta1.FederatedIngress, error)
	Update(ctx context.Context, federatedIngress *v1beta1.FederatedIngress, opts v1.UpdateOptions) (*v1beta1.FederatedIngress, error)
	UpdateStatus(ctx context.Context, federatedIngress *v1beta1.FederatedIngress, opts v1.UpdateOptions) (*v1beta1.FederatedIngress, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1beta1.FederatedIngress, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1beta1.FederatedIngressList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.FederatedIngress, err error)
	FederatedIngressExpansion
}

// federatedIngresses implements FederatedIngressInterface
type federatedIngresses struct {
	client rest.Interface
	ns     string
}

// newFederatedIngresses returns a FederatedIngresses
func newFederatedIngresses(c *TypesV1beta1Client, namespace string) *federatedIngresses {
	return &federatedIngresses{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the federatedIngress, and returns the corresponding federatedIngress object, and an error if there is any.
func (c *federatedIngresses) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta1.FederatedIngress, err error) {
	result = &v1beta1.FederatedIngress{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("federatedingresses").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of FederatedIngresses that match those selectors.
func (c *federatedIngresses) List(ctx context.Context, opts v1.ListOptions) (result *v1beta1.FederatedIngressList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1beta1.FederatedIngressList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("federatedingresses").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested federatedIngresses.
func (c *federatedIngresses) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("federatedingresses").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a federatedIngress and creates it.  Returns the server's representation of the federatedIngress, and an error, if there is any.
func (c *federatedIngresses) Create(ctx context.Context, federatedIngress *v1beta1.FederatedIngress, opts v1.CreateOptions) (result *v1beta1.FederatedIngress, err error) {
	result = &v1beta1.FederatedIngress{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("federatedingresses").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(federatedIngress).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a federatedIngress and updates it. Returns the server's representation of the federatedIngress, and an error, if there is any.
func (c *federatedIngresses) Update(ctx context.Context, federatedIngress *v1beta1.FederatedIngress, opts v1.UpdateOptions) (result *v1beta1.FederatedIngress, err error) {
	result = &v1beta1.FederatedIngress{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("federatedingresses").
		Name(federatedIngress.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(federatedIngress).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *federatedIngresses) UpdateStatus(ctx context.Context, federatedIngress *v1beta1.FederatedIngress, opts v1.UpdateOptions) (result *v1beta1.FederatedIngress, err error) {
	result = &v1beta1.FederatedIngress{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("federatedingresses").
		Name(federatedIngress.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(federatedIngress).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the federatedIngress and deletes it. Returns an error if one occurs.
func (c *federatedIngresses) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("federatedingresses").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *federatedIngresses) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("federatedingresses").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched federatedIngress.
func (c *federatedIngresses) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.FederatedIngress, err error) {
	result = &v1beta1.FederatedIngress{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("federatedingresses").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
