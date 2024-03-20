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

// FakeBGPFilters implements BGPFilterInterface
type FakeBGPFilters struct {
	Fake *FakeProjectcalicoV3
}

var bgpfiltersResource = schema.GroupVersionResource{Group: "projectcalico.org", Version: "v3", Resource: "bgpfilters"}

var bgpfiltersKind = schema.GroupVersionKind{Group: "projectcalico.org", Version: "v3", Kind: "BGPFilter"}

// Get takes name of the bGPFilter, and returns the corresponding bGPFilter object, and an error if there is any.
func (c *FakeBGPFilters) Get(ctx context.Context, name string, options v1.GetOptions) (result *v3.BGPFilter, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(bgpfiltersResource, name), &v3.BGPFilter{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v3.BGPFilter), err
}

// List takes label and field selectors, and returns the list of BGPFilters that match those selectors.
func (c *FakeBGPFilters) List(ctx context.Context, opts v1.ListOptions) (result *v3.BGPFilterList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(bgpfiltersResource, bgpfiltersKind, opts), &v3.BGPFilterList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v3.BGPFilterList{ListMeta: obj.(*v3.BGPFilterList).ListMeta}
	for _, item := range obj.(*v3.BGPFilterList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested bGPFilters.
func (c *FakeBGPFilters) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(bgpfiltersResource, opts))
}

// Create takes the representation of a bGPFilter and creates it.  Returns the server's representation of the bGPFilter, and an error, if there is any.
func (c *FakeBGPFilters) Create(ctx context.Context, bGPFilter *v3.BGPFilter, opts v1.CreateOptions) (result *v3.BGPFilter, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(bgpfiltersResource, bGPFilter), &v3.BGPFilter{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v3.BGPFilter), err
}

// Update takes the representation of a bGPFilter and updates it. Returns the server's representation of the bGPFilter, and an error, if there is any.
func (c *FakeBGPFilters) Update(ctx context.Context, bGPFilter *v3.BGPFilter, opts v1.UpdateOptions) (result *v3.BGPFilter, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(bgpfiltersResource, bGPFilter), &v3.BGPFilter{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v3.BGPFilter), err
}

// Delete takes name of the bGPFilter and deletes it. Returns an error if one occurs.
func (c *FakeBGPFilters) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteActionWithOptions(bgpfiltersResource, name, opts), &v3.BGPFilter{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeBGPFilters) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(bgpfiltersResource, listOpts)

	_, err := c.Fake.Invokes(action, &v3.BGPFilterList{})
	return err
}

// Patch applies the patch and returns the patched bGPFilter.
func (c *FakeBGPFilters) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v3.BGPFilter, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(bgpfiltersResource, name, pt, data, subresources...), &v3.BGPFilter{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v3.BGPFilter), err
}
