// Copyright (c) 2023 Tigera, Inc. All rights reserved.

// Code generated by client-gen. DO NOT EDIT.

package v3

import (
	"context"
	"time"

	v3 "github.com/tigera/api/pkg/apis/projectcalico/v3"
	scheme "github.com/tigera/api/pkg/client/clientset_generated/clientset/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// ExternalNetworksGetter has a method to return a ExternalNetworkInterface.
// A group's client should implement this interface.
type ExternalNetworksGetter interface {
	ExternalNetworks() ExternalNetworkInterface
}

// ExternalNetworkInterface has methods to work with ExternalNetwork resources.
type ExternalNetworkInterface interface {
	Create(ctx context.Context, externalNetwork *v3.ExternalNetwork, opts v1.CreateOptions) (*v3.ExternalNetwork, error)
	Update(ctx context.Context, externalNetwork *v3.ExternalNetwork, opts v1.UpdateOptions) (*v3.ExternalNetwork, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v3.ExternalNetwork, error)
	List(ctx context.Context, opts v1.ListOptions) (*v3.ExternalNetworkList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v3.ExternalNetwork, err error)
	ExternalNetworkExpansion
}

// externalNetworks implements ExternalNetworkInterface
type externalNetworks struct {
	client rest.Interface
}

// newExternalNetworks returns a ExternalNetworks
func newExternalNetworks(c *ProjectcalicoV3Client) *externalNetworks {
	return &externalNetworks{
		client: c.RESTClient(),
	}
}

// Get takes name of the externalNetwork, and returns the corresponding externalNetwork object, and an error if there is any.
func (c *externalNetworks) Get(ctx context.Context, name string, options v1.GetOptions) (result *v3.ExternalNetwork, err error) {
	result = &v3.ExternalNetwork{}
	err = c.client.Get().
		Resource("externalnetworks").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of ExternalNetworks that match those selectors.
func (c *externalNetworks) List(ctx context.Context, opts v1.ListOptions) (result *v3.ExternalNetworkList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v3.ExternalNetworkList{}
	err = c.client.Get().
		Resource("externalnetworks").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested externalNetworks.
func (c *externalNetworks) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("externalnetworks").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a externalNetwork and creates it.  Returns the server's representation of the externalNetwork, and an error, if there is any.
func (c *externalNetworks) Create(ctx context.Context, externalNetwork *v3.ExternalNetwork, opts v1.CreateOptions) (result *v3.ExternalNetwork, err error) {
	result = &v3.ExternalNetwork{}
	err = c.client.Post().
		Resource("externalnetworks").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(externalNetwork).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a externalNetwork and updates it. Returns the server's representation of the externalNetwork, and an error, if there is any.
func (c *externalNetworks) Update(ctx context.Context, externalNetwork *v3.ExternalNetwork, opts v1.UpdateOptions) (result *v3.ExternalNetwork, err error) {
	result = &v3.ExternalNetwork{}
	err = c.client.Put().
		Resource("externalnetworks").
		Name(externalNetwork.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(externalNetwork).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the externalNetwork and deletes it. Returns an error if one occurs.
func (c *externalNetworks) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("externalnetworks").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *externalNetworks) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("externalnetworks").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched externalNetwork.
func (c *externalNetworks) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v3.ExternalNetwork, err error) {
	result = &v3.ExternalNetwork{}
	err = c.client.Patch(pt).
		Resource("externalnetworks").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
