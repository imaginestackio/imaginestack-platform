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

package fake

import (
	"context"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
	v1beta1 "imaginekube.com/api/types/v1beta1"
)

// FakeFederatedJobs implements FederatedJobInterface
type FakeFederatedJobs struct {
	Fake *FakeTypesV1beta1
	ns   string
}

var federatedjobsResource = schema.GroupVersionResource{Group: "types.kubefed.io", Version: "v1beta1", Resource: "federatedjobs"}

var federatedjobsKind = schema.GroupVersionKind{Group: "types.kubefed.io", Version: "v1beta1", Kind: "FederatedJob"}

// Get takes name of the federatedJob, and returns the corresponding federatedJob object, and an error if there is any.
func (c *FakeFederatedJobs) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta1.FederatedJob, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(federatedjobsResource, c.ns, name), &v1beta1.FederatedJob{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.FederatedJob), err
}

// List takes label and field selectors, and returns the list of FederatedJobs that match those selectors.
func (c *FakeFederatedJobs) List(ctx context.Context, opts v1.ListOptions) (result *v1beta1.FederatedJobList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(federatedjobsResource, federatedjobsKind, c.ns, opts), &v1beta1.FederatedJobList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1beta1.FederatedJobList{ListMeta: obj.(*v1beta1.FederatedJobList).ListMeta}
	for _, item := range obj.(*v1beta1.FederatedJobList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested federatedJobs.
func (c *FakeFederatedJobs) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(federatedjobsResource, c.ns, opts))

}

// Create takes the representation of a federatedJob and creates it.  Returns the server's representation of the federatedJob, and an error, if there is any.
func (c *FakeFederatedJobs) Create(ctx context.Context, federatedJob *v1beta1.FederatedJob, opts v1.CreateOptions) (result *v1beta1.FederatedJob, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(federatedjobsResource, c.ns, federatedJob), &v1beta1.FederatedJob{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.FederatedJob), err
}

// Update takes the representation of a federatedJob and updates it. Returns the server's representation of the federatedJob, and an error, if there is any.
func (c *FakeFederatedJobs) Update(ctx context.Context, federatedJob *v1beta1.FederatedJob, opts v1.UpdateOptions) (result *v1beta1.FederatedJob, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(federatedjobsResource, c.ns, federatedJob), &v1beta1.FederatedJob{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.FederatedJob), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeFederatedJobs) UpdateStatus(ctx context.Context, federatedJob *v1beta1.FederatedJob, opts v1.UpdateOptions) (*v1beta1.FederatedJob, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(federatedjobsResource, "status", c.ns, federatedJob), &v1beta1.FederatedJob{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.FederatedJob), err
}

// Delete takes name of the federatedJob and deletes it. Returns an error if one occurs.
func (c *FakeFederatedJobs) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(federatedjobsResource, c.ns, name), &v1beta1.FederatedJob{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeFederatedJobs) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(federatedjobsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1beta1.FederatedJobList{})
	return err
}

// Patch applies the patch and returns the patched federatedJob.
func (c *FakeFederatedJobs) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.FederatedJob, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(federatedjobsResource, c.ns, name, pt, data, subresources...), &v1beta1.FederatedJob{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.FederatedJob), err
}
