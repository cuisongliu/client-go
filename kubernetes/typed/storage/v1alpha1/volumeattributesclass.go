/*
Copyright The Kubernetes Authors.

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
	json "encoding/json"
	"fmt"
	"time"

	v1alpha1 "k8s.io/api/storage/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	storagev1alpha1 "k8s.io/client-go/applyconfigurations/storage/v1alpha1"
	scheme "k8s.io/client-go/kubernetes/scheme"
	rest "k8s.io/client-go/rest"
	consistencydetector "k8s.io/client-go/util/consistencydetector"
)

// VolumeAttributesClassesGetter has a method to return a VolumeAttributesClassInterface.
// A group's client should implement this interface.
type VolumeAttributesClassesGetter interface {
	VolumeAttributesClasses() VolumeAttributesClassInterface
}

// VolumeAttributesClassInterface has methods to work with VolumeAttributesClass resources.
type VolumeAttributesClassInterface interface {
	Create(ctx context.Context, volumeAttributesClass *v1alpha1.VolumeAttributesClass, opts v1.CreateOptions) (*v1alpha1.VolumeAttributesClass, error)
	Update(ctx context.Context, volumeAttributesClass *v1alpha1.VolumeAttributesClass, opts v1.UpdateOptions) (*v1alpha1.VolumeAttributesClass, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.VolumeAttributesClass, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.VolumeAttributesClassList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.VolumeAttributesClass, err error)
	Apply(ctx context.Context, volumeAttributesClass *storagev1alpha1.VolumeAttributesClassApplyConfiguration, opts v1.ApplyOptions) (result *v1alpha1.VolumeAttributesClass, err error)
	VolumeAttributesClassExpansion
}

// volumeAttributesClasses implements VolumeAttributesClassInterface
type volumeAttributesClasses struct {
	client rest.Interface
}

// newVolumeAttributesClasses returns a VolumeAttributesClasses
func newVolumeAttributesClasses(c *StorageV1alpha1Client) *volumeAttributesClasses {
	return &volumeAttributesClasses{
		client: c.RESTClient(),
	}
}

// Get takes name of the volumeAttributesClass, and returns the corresponding volumeAttributesClass object, and an error if there is any.
func (c *volumeAttributesClasses) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.VolumeAttributesClass, err error) {
	result = &v1alpha1.VolumeAttributesClass{}
	err = c.client.Get().
		Resource("volumeattributesclasses").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of VolumeAttributesClasses that match those selectors.
func (c *volumeAttributesClasses) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.VolumeAttributesClassList, err error) {
	defer func() {
		if err == nil {
			consistencydetector.CheckListFromCacheDataConsistencyIfRequested(ctx, "list request for volumeattributesclasses", c.list, opts, result)
		}
	}()
	return c.list(ctx, opts)
}

// list takes label and field selectors, and returns the list of VolumeAttributesClasses that match those selectors.
func (c *volumeAttributesClasses) list(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.VolumeAttributesClassList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.VolumeAttributesClassList{}
	err = c.client.Get().
		Resource("volumeattributesclasses").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested volumeAttributesClasses.
func (c *volumeAttributesClasses) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("volumeattributesclasses").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a volumeAttributesClass and creates it.  Returns the server's representation of the volumeAttributesClass, and an error, if there is any.
func (c *volumeAttributesClasses) Create(ctx context.Context, volumeAttributesClass *v1alpha1.VolumeAttributesClass, opts v1.CreateOptions) (result *v1alpha1.VolumeAttributesClass, err error) {
	result = &v1alpha1.VolumeAttributesClass{}
	err = c.client.Post().
		Resource("volumeattributesclasses").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(volumeAttributesClass).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a volumeAttributesClass and updates it. Returns the server's representation of the volumeAttributesClass, and an error, if there is any.
func (c *volumeAttributesClasses) Update(ctx context.Context, volumeAttributesClass *v1alpha1.VolumeAttributesClass, opts v1.UpdateOptions) (result *v1alpha1.VolumeAttributesClass, err error) {
	result = &v1alpha1.VolumeAttributesClass{}
	err = c.client.Put().
		Resource("volumeattributesclasses").
		Name(volumeAttributesClass.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(volumeAttributesClass).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the volumeAttributesClass and deletes it. Returns an error if one occurs.
func (c *volumeAttributesClasses) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("volumeattributesclasses").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *volumeAttributesClasses) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("volumeattributesclasses").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched volumeAttributesClass.
func (c *volumeAttributesClasses) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.VolumeAttributesClass, err error) {
	result = &v1alpha1.VolumeAttributesClass{}
	err = c.client.Patch(pt).
		Resource("volumeattributesclasses").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}

// Apply takes the given apply declarative configuration, applies it and returns the applied volumeAttributesClass.
func (c *volumeAttributesClasses) Apply(ctx context.Context, volumeAttributesClass *storagev1alpha1.VolumeAttributesClassApplyConfiguration, opts v1.ApplyOptions) (result *v1alpha1.VolumeAttributesClass, err error) {
	if volumeAttributesClass == nil {
		return nil, fmt.Errorf("volumeAttributesClass provided to Apply must not be nil")
	}
	patchOpts := opts.ToPatchOptions()
	data, err := json.Marshal(volumeAttributesClass)
	if err != nil {
		return nil, err
	}
	name := volumeAttributesClass.Name
	if name == nil {
		return nil, fmt.Errorf("volumeAttributesClass.Name must be provided to Apply")
	}
	result = &v1alpha1.VolumeAttributesClass{}
	err = c.client.Patch(types.ApplyPatchType).
		Resource("volumeattributesclasses").
		Name(*name).
		VersionedParams(&patchOpts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
