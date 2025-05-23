// Copyright (c) 2025 Tigera, Inc. All rights reserved.

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v3 "github.com/tigera/api/pkg/apis/projectcalico/v3"
	projectcalicov3 "github.com/tigera/api/pkg/client/clientset_generated/clientset/typed/projectcalico/v3"
	gentype "k8s.io/client-go/gentype"
)

// fakeRemoteClusterConfigurations implements RemoteClusterConfigurationInterface
type fakeRemoteClusterConfigurations struct {
	*gentype.FakeClientWithList[*v3.RemoteClusterConfiguration, *v3.RemoteClusterConfigurationList]
	Fake *FakeProjectcalicoV3
}

func newFakeRemoteClusterConfigurations(fake *FakeProjectcalicoV3) projectcalicov3.RemoteClusterConfigurationInterface {
	return &fakeRemoteClusterConfigurations{
		gentype.NewFakeClientWithList[*v3.RemoteClusterConfiguration, *v3.RemoteClusterConfigurationList](
			fake.Fake,
			"",
			v3.SchemeGroupVersion.WithResource("remoteclusterconfigurations"),
			v3.SchemeGroupVersion.WithKind("RemoteClusterConfiguration"),
			func() *v3.RemoteClusterConfiguration { return &v3.RemoteClusterConfiguration{} },
			func() *v3.RemoteClusterConfigurationList { return &v3.RemoteClusterConfigurationList{} },
			func(dst, src *v3.RemoteClusterConfigurationList) { dst.ListMeta = src.ListMeta },
			func(list *v3.RemoteClusterConfigurationList) []*v3.RemoteClusterConfiguration {
				return gentype.ToPointerSlice(list.Items)
			},
			func(list *v3.RemoteClusterConfigurationList, items []*v3.RemoteClusterConfiguration) {
				list.Items = gentype.FromPointerSlice(items)
			},
		),
		fake,
	}
}
