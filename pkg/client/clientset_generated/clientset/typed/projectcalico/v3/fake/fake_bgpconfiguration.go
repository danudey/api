// Copyright (c) 2024 Tigera, Inc. All rights reserved.

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	v3 "github.com/tigera/api/pkg/apis/projectcalico/v3"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeBGPConfigurations implements BGPConfigurationInterface
type FakeBGPConfigurations struct {
	Fake *FakeProjectcalicoV3
}

var bgpconfigurationsResource = schema.GroupVersionResource{Group: "projectcalico.org", Version: "v3", Resource: "bgpconfigurations"}

var bgpconfigurationsKind = schema.GroupVersionKind{Group: "projectcalico.org", Version: "v3", Kind: "BGPConfiguration"}

// Get takes name of the bGPConfiguration, and returns the corresponding bGPConfiguration object, and an error if there is any.
func (c *FakeBGPConfigurations) Get(ctx context.Context, name string, options v1.GetOptions) (result *v3.BGPConfiguration, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(bgpconfigurationsResource, name), &v3.BGPConfiguration{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v3.BGPConfiguration), err
}

// List takes label and field selectors, and returns the list of BGPConfigurations that match those selectors.
func (c *FakeBGPConfigurations) List(ctx context.Context, opts v1.ListOptions) (result *v3.BGPConfigurationList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(bgpconfigurationsResource, bgpconfigurationsKind, opts), &v3.BGPConfigurationList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v3.BGPConfigurationList{ListMeta: obj.(*v3.BGPConfigurationList).ListMeta}
	for _, item := range obj.(*v3.BGPConfigurationList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested bGPConfigurations.
func (c *FakeBGPConfigurations) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(bgpconfigurationsResource, opts))
}

// Create takes the representation of a bGPConfiguration and creates it.  Returns the server's representation of the bGPConfiguration, and an error, if there is any.
func (c *FakeBGPConfigurations) Create(ctx context.Context, bGPConfiguration *v3.BGPConfiguration, opts v1.CreateOptions) (result *v3.BGPConfiguration, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(bgpconfigurationsResource, bGPConfiguration), &v3.BGPConfiguration{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v3.BGPConfiguration), err
}

// Update takes the representation of a bGPConfiguration and updates it. Returns the server's representation of the bGPConfiguration, and an error, if there is any.
func (c *FakeBGPConfigurations) Update(ctx context.Context, bGPConfiguration *v3.BGPConfiguration, opts v1.UpdateOptions) (result *v3.BGPConfiguration, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(bgpconfigurationsResource, bGPConfiguration), &v3.BGPConfiguration{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v3.BGPConfiguration), err
}

// Delete takes name of the bGPConfiguration and deletes it. Returns an error if one occurs.
func (c *FakeBGPConfigurations) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteActionWithOptions(bgpconfigurationsResource, name, opts), &v3.BGPConfiguration{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeBGPConfigurations) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(bgpconfigurationsResource, listOpts)

	_, err := c.Fake.Invokes(action, &v3.BGPConfigurationList{})
	return err
}

// Patch applies the patch and returns the patched bGPConfiguration.
func (c *FakeBGPConfigurations) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v3.BGPConfiguration, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(bgpconfigurationsResource, name, pt, data, subresources...), &v3.BGPConfiguration{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v3.BGPConfiguration), err
}
