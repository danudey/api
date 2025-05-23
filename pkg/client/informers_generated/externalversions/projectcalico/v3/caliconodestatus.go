// Copyright (c) 2025 Tigera, Inc. All rights reserved.

// Code generated by informer-gen. DO NOT EDIT.

package v3

import (
	context "context"
	time "time"

	apisprojectcalicov3 "github.com/tigera/api/pkg/apis/projectcalico/v3"
	clientset "github.com/tigera/api/pkg/client/clientset_generated/clientset"
	internalinterfaces "github.com/tigera/api/pkg/client/informers_generated/externalversions/internalinterfaces"
	projectcalicov3 "github.com/tigera/api/pkg/client/listers_generated/projectcalico/v3"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// CalicoNodeStatusInformer provides access to a shared informer and lister for
// CalicoNodeStatuses.
type CalicoNodeStatusInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() projectcalicov3.CalicoNodeStatusLister
}

type calicoNodeStatusInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// NewCalicoNodeStatusInformer constructs a new informer for CalicoNodeStatus type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewCalicoNodeStatusInformer(client clientset.Interface, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredCalicoNodeStatusInformer(client, resyncPeriod, indexers, nil)
}

// NewFilteredCalicoNodeStatusInformer constructs a new informer for CalicoNodeStatus type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredCalicoNodeStatusInformer(client clientset.Interface, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.ProjectcalicoV3().CalicoNodeStatuses().List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.ProjectcalicoV3().CalicoNodeStatuses().Watch(context.TODO(), options)
			},
		},
		&apisprojectcalicov3.CalicoNodeStatus{},
		resyncPeriod,
		indexers,
	)
}

func (f *calicoNodeStatusInformer) defaultInformer(client clientset.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredCalicoNodeStatusInformer(client, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *calicoNodeStatusInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&apisprojectcalicov3.CalicoNodeStatus{}, f.defaultInformer)
}

func (f *calicoNodeStatusInformer) Lister() projectcalicov3.CalicoNodeStatusLister {
	return projectcalicov3.NewCalicoNodeStatusLister(f.Informer().GetIndexer())
}
