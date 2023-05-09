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
	v1alpha2 "imaginekube.com/api/servicemesh/v1alpha2"
)

// FakeStrategies implements StrategyInterface
type FakeStrategies struct {
	Fake *FakeServicemeshV1alpha2
	ns   string
}

var strategiesResource = schema.GroupVersionResource{Group: "servicemesh.imaginekube.com", Version: "v1alpha2", Resource: "strategies"}

var strategiesKind = schema.GroupVersionKind{Group: "servicemesh.imaginekube.com", Version: "v1alpha2", Kind: "Strategy"}

// Get takes name of the strategy, and returns the corresponding strategy object, and an error if there is any.
func (c *FakeStrategies) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha2.Strategy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(strategiesResource, c.ns, name), &v1alpha2.Strategy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha2.Strategy), err
}

// List takes label and field selectors, and returns the list of Strategies that match those selectors.
func (c *FakeStrategies) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha2.StrategyList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(strategiesResource, strategiesKind, c.ns, opts), &v1alpha2.StrategyList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha2.StrategyList{ListMeta: obj.(*v1alpha2.StrategyList).ListMeta}
	for _, item := range obj.(*v1alpha2.StrategyList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested strategies.
func (c *FakeStrategies) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(strategiesResource, c.ns, opts))

}

// Create takes the representation of a strategy and creates it.  Returns the server's representation of the strategy, and an error, if there is any.
func (c *FakeStrategies) Create(ctx context.Context, strategy *v1alpha2.Strategy, opts v1.CreateOptions) (result *v1alpha2.Strategy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(strategiesResource, c.ns, strategy), &v1alpha2.Strategy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha2.Strategy), err
}

// Update takes the representation of a strategy and updates it. Returns the server's representation of the strategy, and an error, if there is any.
func (c *FakeStrategies) Update(ctx context.Context, strategy *v1alpha2.Strategy, opts v1.UpdateOptions) (result *v1alpha2.Strategy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(strategiesResource, c.ns, strategy), &v1alpha2.Strategy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha2.Strategy), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeStrategies) UpdateStatus(ctx context.Context, strategy *v1alpha2.Strategy, opts v1.UpdateOptions) (*v1alpha2.Strategy, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(strategiesResource, "status", c.ns, strategy), &v1alpha2.Strategy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha2.Strategy), err
}

// Delete takes name of the strategy and deletes it. Returns an error if one occurs.
func (c *FakeStrategies) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(strategiesResource, c.ns, name), &v1alpha2.Strategy{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeStrategies) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(strategiesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha2.StrategyList{})
	return err
}

// Patch applies the patch and returns the patched strategy.
func (c *FakeStrategies) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha2.Strategy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(strategiesResource, c.ns, name, pt, data, subresources...), &v1alpha2.Strategy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha2.Strategy), err
}
