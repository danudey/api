// Copyright (c) 2025 Tigera, Inc. All rights reserved.

// Code generated by lister-gen. DO NOT EDIT.

package v3

import (
	projectcalicov3 "github.com/tigera/api/pkg/apis/projectcalico/v3"
	labels "k8s.io/apimachinery/pkg/labels"
	listers "k8s.io/client-go/listers"
	cache "k8s.io/client-go/tools/cache"
)

// NetworkSetLister helps list NetworkSets.
// All objects returned here must be treated as read-only.
type NetworkSetLister interface {
	// List lists all NetworkSets in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*projectcalicov3.NetworkSet, err error)
	// NetworkSets returns an object that can list and get NetworkSets.
	NetworkSets(namespace string) NetworkSetNamespaceLister
	NetworkSetListerExpansion
}

// networkSetLister implements the NetworkSetLister interface.
type networkSetLister struct {
	listers.ResourceIndexer[*projectcalicov3.NetworkSet]
}

// NewNetworkSetLister returns a new NetworkSetLister.
func NewNetworkSetLister(indexer cache.Indexer) NetworkSetLister {
	return &networkSetLister{listers.New[*projectcalicov3.NetworkSet](indexer, projectcalicov3.Resource("networkset"))}
}

// NetworkSets returns an object that can list and get NetworkSets.
func (s *networkSetLister) NetworkSets(namespace string) NetworkSetNamespaceLister {
	return networkSetNamespaceLister{listers.NewNamespaced[*projectcalicov3.NetworkSet](s.ResourceIndexer, namespace)}
}

// NetworkSetNamespaceLister helps list and get NetworkSets.
// All objects returned here must be treated as read-only.
type NetworkSetNamespaceLister interface {
	// List lists all NetworkSets in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*projectcalicov3.NetworkSet, err error)
	// Get retrieves the NetworkSet from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*projectcalicov3.NetworkSet, error)
	NetworkSetNamespaceListerExpansion
}

// networkSetNamespaceLister implements the NetworkSetNamespaceLister
// interface.
type networkSetNamespaceLister struct {
	listers.ResourceIndexer[*projectcalicov3.NetworkSet]
}
