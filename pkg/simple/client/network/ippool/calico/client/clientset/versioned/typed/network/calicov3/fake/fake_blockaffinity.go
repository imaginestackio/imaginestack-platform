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

	calicov3 "imaginekube.com/api/network/calicov3"
)

// FakeBlockAffinities implements BlockAffinityInterface
type FakeBlockAffinities struct {
	Fake *FakeCrdCalicov3
}

var blockaffinitiesResource = schema.GroupVersionResource{Group: "crd.projectcalico.org", Version: "calicov3", Resource: "blockaffinities"}

var blockaffinitiesKind = schema.GroupVersionKind{Group: "crd.projectcalico.org", Version: "calicov3", Kind: "BlockAffinity"}

// Get takes name of the blockAffinity, and returns the corresponding blockAffinity object, and an error if there is any.
func (c *FakeBlockAffinities) Get(ctx context.Context, name string, options v1.GetOptions) (result *calicov3.BlockAffinity, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(blockaffinitiesResource, name), &calicov3.BlockAffinity{})
	if obj == nil {
		return nil, err
	}
	return obj.(*calicov3.BlockAffinity), err
}

// List takes label and field selectors, and returns the list of BlockAffinities that match those selectors.
func (c *FakeBlockAffinities) List(ctx context.Context, opts v1.ListOptions) (result *calicov3.BlockAffinityList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(blockaffinitiesResource, blockaffinitiesKind, opts), &calicov3.BlockAffinityList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &calicov3.BlockAffinityList{ListMeta: obj.(*calicov3.BlockAffinityList).ListMeta}
	for _, item := range obj.(*calicov3.BlockAffinityList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested blockAffinities.
func (c *FakeBlockAffinities) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(blockaffinitiesResource, opts))
}

// Create takes the representation of a blockAffinity and creates it.  Returns the server's representation of the blockAffinity, and an error, if there is any.
func (c *FakeBlockAffinities) Create(ctx context.Context, blockAffinity *calicov3.BlockAffinity, opts v1.CreateOptions) (result *calicov3.BlockAffinity, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(blockaffinitiesResource, blockAffinity), &calicov3.BlockAffinity{})
	if obj == nil {
		return nil, err
	}
	return obj.(*calicov3.BlockAffinity), err
}

// Update takes the representation of a blockAffinity and updates it. Returns the server's representation of the blockAffinity, and an error, if there is any.
func (c *FakeBlockAffinities) Update(ctx context.Context, blockAffinity *calicov3.BlockAffinity, opts v1.UpdateOptions) (result *calicov3.BlockAffinity, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(blockaffinitiesResource, blockAffinity), &calicov3.BlockAffinity{})
	if obj == nil {
		return nil, err
	}
	return obj.(*calicov3.BlockAffinity), err
}

// Delete takes name of the blockAffinity and deletes it. Returns an error if one occurs.
func (c *FakeBlockAffinities) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(blockaffinitiesResource, name), &calicov3.BlockAffinity{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeBlockAffinities) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(blockaffinitiesResource, listOpts)

	_, err := c.Fake.Invokes(action, &calicov3.BlockAffinityList{})
	return err
}

// Patch applies the patch and returns the patched blockAffinity.
func (c *FakeBlockAffinities) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *calicov3.BlockAffinity, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(blockaffinitiesResource, name, pt, data, subresources...), &calicov3.BlockAffinity{})
	if obj == nil {
		return nil, err
	}
	return obj.(*calicov3.BlockAffinity), err
}
