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

package v1

import (
	"context"
	json "encoding/json"
	"fmt"
	"time"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	corev1 "k8s.io/client-go/applyconfigurations/core/v1"
	scheme "k8s.io/client-go/kubernetes/scheme"
	rest "k8s.io/client-go/rest"
	consistencydetector "k8s.io/client-go/util/consistencydetector"
)

// LimitRangesGetter has a method to return a LimitRangeInterface.
// A group's client should implement this interface.
type LimitRangesGetter interface {
	LimitRanges(namespace string) LimitRangeInterface
}

// LimitRangeInterface has methods to work with LimitRange resources.
type LimitRangeInterface interface {
	Create(ctx context.Context, limitRange *v1.LimitRange, opts metav1.CreateOptions) (*v1.LimitRange, error)
	Update(ctx context.Context, limitRange *v1.LimitRange, opts metav1.UpdateOptions) (*v1.LimitRange, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*v1.LimitRange, error)
	List(ctx context.Context, opts metav1.ListOptions) (*v1.LimitRangeList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.LimitRange, err error)
	Apply(ctx context.Context, limitRange *corev1.LimitRangeApplyConfiguration, opts metav1.ApplyOptions) (result *v1.LimitRange, err error)
	LimitRangeExpansion
}

// limitRanges implements LimitRangeInterface
type limitRanges struct {
	client rest.Interface
	ns     string
}

// newLimitRanges returns a LimitRanges
func newLimitRanges(c *CoreV1Client, namespace string) *limitRanges {
	return &limitRanges{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the limitRange, and returns the corresponding limitRange object, and an error if there is any.
func (c *limitRanges) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.LimitRange, err error) {
	result = &v1.LimitRange{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("limitranges").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of LimitRanges that match those selectors.
func (c *limitRanges) List(ctx context.Context, opts metav1.ListOptions) (result *v1.LimitRangeList, err error) {
	defer func() {
		if err == nil {
			consistencydetector.CheckListFromCacheDataConsistencyIfRequested(ctx, "list request for limitranges", c.list, opts, result)
		}
	}()
	return c.list(ctx, opts)
}

// list takes label and field selectors, and returns the list of LimitRanges that match those selectors.
func (c *limitRanges) list(ctx context.Context, opts metav1.ListOptions) (result *v1.LimitRangeList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.LimitRangeList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("limitranges").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested limitRanges.
func (c *limitRanges) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("limitranges").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a limitRange and creates it.  Returns the server's representation of the limitRange, and an error, if there is any.
func (c *limitRanges) Create(ctx context.Context, limitRange *v1.LimitRange, opts metav1.CreateOptions) (result *v1.LimitRange, err error) {
	result = &v1.LimitRange{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("limitranges").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(limitRange).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a limitRange and updates it. Returns the server's representation of the limitRange, and an error, if there is any.
func (c *limitRanges) Update(ctx context.Context, limitRange *v1.LimitRange, opts metav1.UpdateOptions) (result *v1.LimitRange, err error) {
	result = &v1.LimitRange{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("limitranges").
		Name(limitRange.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(limitRange).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the limitRange and deletes it. Returns an error if one occurs.
func (c *limitRanges) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("limitranges").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *limitRanges) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("limitranges").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched limitRange.
func (c *limitRanges) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.LimitRange, err error) {
	result = &v1.LimitRange{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("limitranges").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}

// Apply takes the given apply declarative configuration, applies it and returns the applied limitRange.
func (c *limitRanges) Apply(ctx context.Context, limitRange *corev1.LimitRangeApplyConfiguration, opts metav1.ApplyOptions) (result *v1.LimitRange, err error) {
	if limitRange == nil {
		return nil, fmt.Errorf("limitRange provided to Apply must not be nil")
	}
	patchOpts := opts.ToPatchOptions()
	data, err := json.Marshal(limitRange)
	if err != nil {
		return nil, err
	}
	name := limitRange.Name
	if name == nil {
		return nil, fmt.Errorf("limitRange.Name must be provided to Apply")
	}
	result = &v1.LimitRange{}
	err = c.client.Patch(types.ApplyPatchType).
		Namespace(c.ns).
		Resource("limitranges").
		Name(*name).
		VersionedParams(&patchOpts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
