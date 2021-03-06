/*
 * Tencent is pleased to support the open source community by making TKEStack
 * available.
 *
 * Copyright (C) 2012-2020 Tencent. All Rights Reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License"); you may not use
 * this file except in compliance with the License. You may obtain a copy of the
 * License at
 *
 * https://opensource.org/licenses/Apache-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
 * WARRANTIES OF ANY KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations under the License.
 */

// Code generated by client-gen. DO NOT EDIT.

package internalversion

import (
	"context"
	"time"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
	scheme "tkestack.io/tke/api/client/clientset/internalversion/scheme"
	platform "tkestack.io/tke/api/platform"
)

// TappControllersGetter has a method to return a TappControllerInterface.
// A group's client should implement this interface.
type TappControllersGetter interface {
	TappControllers() TappControllerInterface
}

// TappControllerInterface has methods to work with TappController resources.
type TappControllerInterface interface {
	Create(ctx context.Context, tappController *platform.TappController, opts v1.CreateOptions) (*platform.TappController, error)
	Update(ctx context.Context, tappController *platform.TappController, opts v1.UpdateOptions) (*platform.TappController, error)
	UpdateStatus(ctx context.Context, tappController *platform.TappController, opts v1.UpdateOptions) (*platform.TappController, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*platform.TappController, error)
	List(ctx context.Context, opts v1.ListOptions) (*platform.TappControllerList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *platform.TappController, err error)
	TappControllerExpansion
}

// tappControllers implements TappControllerInterface
type tappControllers struct {
	client rest.Interface
}

// newTappControllers returns a TappControllers
func newTappControllers(c *PlatformClient) *tappControllers {
	return &tappControllers{
		client: c.RESTClient(),
	}
}

// Get takes name of the tappController, and returns the corresponding tappController object, and an error if there is any.
func (c *tappControllers) Get(ctx context.Context, name string, options v1.GetOptions) (result *platform.TappController, err error) {
	result = &platform.TappController{}
	err = c.client.Get().
		Resource("tappcontrollers").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of TappControllers that match those selectors.
func (c *tappControllers) List(ctx context.Context, opts v1.ListOptions) (result *platform.TappControllerList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &platform.TappControllerList{}
	err = c.client.Get().
		Resource("tappcontrollers").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested tappControllers.
func (c *tappControllers) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("tappcontrollers").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a tappController and creates it.  Returns the server's representation of the tappController, and an error, if there is any.
func (c *tappControllers) Create(ctx context.Context, tappController *platform.TappController, opts v1.CreateOptions) (result *platform.TappController, err error) {
	result = &platform.TappController{}
	err = c.client.Post().
		Resource("tappcontrollers").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(tappController).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a tappController and updates it. Returns the server's representation of the tappController, and an error, if there is any.
func (c *tappControllers) Update(ctx context.Context, tappController *platform.TappController, opts v1.UpdateOptions) (result *platform.TappController, err error) {
	result = &platform.TappController{}
	err = c.client.Put().
		Resource("tappcontrollers").
		Name(tappController.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(tappController).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *tappControllers) UpdateStatus(ctx context.Context, tappController *platform.TappController, opts v1.UpdateOptions) (result *platform.TappController, err error) {
	result = &platform.TappController{}
	err = c.client.Put().
		Resource("tappcontrollers").
		Name(tappController.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(tappController).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the tappController and deletes it. Returns an error if one occurs.
func (c *tappControllers) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("tappcontrollers").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched tappController.
func (c *tappControllers) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *platform.TappController, err error) {
	result = &platform.TappController{}
	err = c.client.Patch(pt).
		Resource("tappcontrollers").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
