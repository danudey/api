// Copyright (c) 2024 Tigera, Inc. All rights reserved.

// Code generated by lister-gen. DO NOT EDIT.

package v3

import (
	v3 "github.com/tigera/api/pkg/apis/projectcalico/v3"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// BGPPeerLister helps list BGPPeers.
// All objects returned here must be treated as read-only.
type BGPPeerLister interface {
	// List lists all BGPPeers in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v3.BGPPeer, err error)
	// Get retrieves the BGPPeer from the index for a given name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v3.BGPPeer, error)
	BGPPeerListerExpansion
}

// bGPPeerLister implements the BGPPeerLister interface.
type bGPPeerLister struct {
	indexer cache.Indexer
}

// NewBGPPeerLister returns a new BGPPeerLister.
func NewBGPPeerLister(indexer cache.Indexer) BGPPeerLister {
	return &bGPPeerLister{indexer: indexer}
}

// List lists all BGPPeers in the indexer.
func (s *bGPPeerLister) List(selector labels.Selector) (ret []*v3.BGPPeer, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v3.BGPPeer))
	})
	return ret, err
}

// Get retrieves the BGPPeer from the index for a given name.
func (s *bGPPeerLister) Get(name string) (*v3.BGPPeer, error) {
	obj, exists, err := s.indexer.GetByKey(name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v3.Resource("bgppeer"), name)
	}
	return obj.(*v3.BGPPeer), nil
}
