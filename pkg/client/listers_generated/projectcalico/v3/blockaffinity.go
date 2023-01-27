// Copyright (c) 2023 Tigera, Inc. All rights reserved.

// Code generated by lister-gen. DO NOT EDIT.

package v3

import (
	v3 "github.com/tigera/api/pkg/apis/projectcalico/v3"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// BlockAffinityLister helps list BlockAffinities.
// All objects returned here must be treated as read-only.
type BlockAffinityLister interface {
	// List lists all BlockAffinities in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v3.BlockAffinity, err error)
	// Get retrieves the BlockAffinity from the index for a given name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v3.BlockAffinity, error)
	BlockAffinityListerExpansion
}

// blockAffinityLister implements the BlockAffinityLister interface.
type blockAffinityLister struct {
	indexer cache.Indexer
}

// NewBlockAffinityLister returns a new BlockAffinityLister.
func NewBlockAffinityLister(indexer cache.Indexer) BlockAffinityLister {
	return &blockAffinityLister{indexer: indexer}
}

// List lists all BlockAffinities in the indexer.
func (s *blockAffinityLister) List(selector labels.Selector) (ret []*v3.BlockAffinity, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v3.BlockAffinity))
	})
	return ret, err
}

// Get retrieves the BlockAffinity from the index for a given name.
func (s *blockAffinityLister) Get(name string) (*v3.BlockAffinity, error) {
	obj, exists, err := s.indexer.GetByKey(name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v3.Resource("blockaffinity"), name)
	}
	return obj.(*v3.BlockAffinity), nil
}
