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

package v1alpha1

import (
	"context"
	"time"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
	v1alpha1 "imaginekube.com/api/auditing/v1alpha1"
	scheme "imaginekube.com/imaginekube/pkg/client/clientset/versioned/scheme"
)

// WebhooksGetter has a method to return a WebhookInterface.
// A group's client should implement this interface.
type WebhooksGetter interface {
	Webhooks() WebhookInterface
}

// WebhookInterface has methods to work with Webhook resources.
type WebhookInterface interface {
	Create(ctx context.Context, webhook *v1alpha1.Webhook, opts v1.CreateOptions) (*v1alpha1.Webhook, error)
	Update(ctx context.Context, webhook *v1alpha1.Webhook, opts v1.UpdateOptions) (*v1alpha1.Webhook, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.Webhook, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.WebhookList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.Webhook, err error)
	WebhookExpansion
}

// webhooks implements WebhookInterface
type webhooks struct {
	client rest.Interface
}

// newWebhooks returns a Webhooks
func newWebhooks(c *AuditingV1alpha1Client) *webhooks {
	return &webhooks{
		client: c.RESTClient(),
	}
}

// Get takes name of the webhook, and returns the corresponding webhook object, and an error if there is any.
func (c *webhooks) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.Webhook, err error) {
	result = &v1alpha1.Webhook{}
	err = c.client.Get().
		Resource("webhooks").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of Webhooks that match those selectors.
func (c *webhooks) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.WebhookList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.WebhookList{}
	err = c.client.Get().
		Resource("webhooks").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested webhooks.
func (c *webhooks) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("webhooks").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a webhook and creates it.  Returns the server's representation of the webhook, and an error, if there is any.
func (c *webhooks) Create(ctx context.Context, webhook *v1alpha1.Webhook, opts v1.CreateOptions) (result *v1alpha1.Webhook, err error) {
	result = &v1alpha1.Webhook{}
	err = c.client.Post().
		Resource("webhooks").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(webhook).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a webhook and updates it. Returns the server's representation of the webhook, and an error, if there is any.
func (c *webhooks) Update(ctx context.Context, webhook *v1alpha1.Webhook, opts v1.UpdateOptions) (result *v1alpha1.Webhook, err error) {
	result = &v1alpha1.Webhook{}
	err = c.client.Put().
		Resource("webhooks").
		Name(webhook.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(webhook).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the webhook and deletes it. Returns an error if one occurs.
func (c *webhooks) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("webhooks").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *webhooks) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("webhooks").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched webhook.
func (c *webhooks) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.Webhook, err error) {
	result = &v1alpha1.Webhook{}
	err = c.client.Patch(pt).
		Resource("webhooks").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
