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

package v1alpha2

import (
	"context"
	"time"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
	v1alpha2 "imaginekube.com/api/iam/v1alpha2"
	scheme "imaginekube.com/imaginekube/pkg/client/clientset/versioned/scheme"
)

// WorkspaceRolesGetter has a method to return a WorkspaceRoleInterface.
// A group's client should implement this interface.
type WorkspaceRolesGetter interface {
	WorkspaceRoles() WorkspaceRoleInterface
}

// WorkspaceRoleInterface has methods to work with WorkspaceRole resources.
type WorkspaceRoleInterface interface {
	Create(ctx context.Context, workspaceRole *v1alpha2.WorkspaceRole, opts v1.CreateOptions) (*v1alpha2.WorkspaceRole, error)
	Update(ctx context.Context, workspaceRole *v1alpha2.WorkspaceRole, opts v1.UpdateOptions) (*v1alpha2.WorkspaceRole, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha2.WorkspaceRole, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha2.WorkspaceRoleList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha2.WorkspaceRole, err error)
	WorkspaceRoleExpansion
}

// workspaceRoles implements WorkspaceRoleInterface
type workspaceRoles struct {
	client rest.Interface
}

// newWorkspaceRoles returns a WorkspaceRoles
func newWorkspaceRoles(c *IamV1alpha2Client) *workspaceRoles {
	return &workspaceRoles{
		client: c.RESTClient(),
	}
}

// Get takes name of the workspaceRole, and returns the corresponding workspaceRole object, and an error if there is any.
func (c *workspaceRoles) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha2.WorkspaceRole, err error) {
	result = &v1alpha2.WorkspaceRole{}
	err = c.client.Get().
		Resource("workspaceroles").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of WorkspaceRoles that match those selectors.
func (c *workspaceRoles) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha2.WorkspaceRoleList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha2.WorkspaceRoleList{}
	err = c.client.Get().
		Resource("workspaceroles").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested workspaceRoles.
func (c *workspaceRoles) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("workspaceroles").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a workspaceRole and creates it.  Returns the server's representation of the workspaceRole, and an error, if there is any.
func (c *workspaceRoles) Create(ctx context.Context, workspaceRole *v1alpha2.WorkspaceRole, opts v1.CreateOptions) (result *v1alpha2.WorkspaceRole, err error) {
	result = &v1alpha2.WorkspaceRole{}
	err = c.client.Post().
		Resource("workspaceroles").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(workspaceRole).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a workspaceRole and updates it. Returns the server's representation of the workspaceRole, and an error, if there is any.
func (c *workspaceRoles) Update(ctx context.Context, workspaceRole *v1alpha2.WorkspaceRole, opts v1.UpdateOptions) (result *v1alpha2.WorkspaceRole, err error) {
	result = &v1alpha2.WorkspaceRole{}
	err = c.client.Put().
		Resource("workspaceroles").
		Name(workspaceRole.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(workspaceRole).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the workspaceRole and deletes it. Returns an error if one occurs.
func (c *workspaceRoles) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("workspaceroles").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *workspaceRoles) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("workspaceroles").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched workspaceRole.
func (c *workspaceRoles) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha2.WorkspaceRole, err error) {
	result = &v1alpha2.WorkspaceRole{}
	err = c.client.Patch(pt).
		Resource("workspaceroles").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
