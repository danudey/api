// Copyright (c) 2025 Tigera, Inc. All rights reserved.

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v3 "github.com/tigera/api/pkg/apis/projectcalico/v3"
	projectcalicov3 "github.com/tigera/api/pkg/client/clientset_generated/clientset/typed/projectcalico/v3"
	gentype "k8s.io/client-go/gentype"
)

// fakeExternalNetworks implements ExternalNetworkInterface
type fakeExternalNetworks struct {
	*gentype.FakeClientWithList[*v3.ExternalNetwork, *v3.ExternalNetworkList]
	Fake *FakeProjectcalicoV3
}

func newFakeExternalNetworks(fake *FakeProjectcalicoV3) projectcalicov3.ExternalNetworkInterface {
	return &fakeExternalNetworks{
		gentype.NewFakeClientWithList[*v3.ExternalNetwork, *v3.ExternalNetworkList](
			fake.Fake,
			"",
			v3.SchemeGroupVersion.WithResource("externalnetworks"),
			v3.SchemeGroupVersion.WithKind("ExternalNetwork"),
			func() *v3.ExternalNetwork { return &v3.ExternalNetwork{} },
			func() *v3.ExternalNetworkList { return &v3.ExternalNetworkList{} },
			func(dst, src *v3.ExternalNetworkList) { dst.ListMeta = src.ListMeta },
			func(list *v3.ExternalNetworkList) []*v3.ExternalNetwork { return gentype.ToPointerSlice(list.Items) },
			func(list *v3.ExternalNetworkList, items []*v3.ExternalNetwork) {
				list.Items = gentype.FromPointerSlice(items)
			},
		),
		fake,
	}
}
