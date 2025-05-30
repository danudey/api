// Copyright (c) 2025 Tigera, Inc. All rights reserved.

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v3 "github.com/tigera/api/pkg/apis/projectcalico/v3"
	projectcalicov3 "github.com/tigera/api/pkg/client/clientset_generated/clientset/typed/projectcalico/v3"
	gentype "k8s.io/client-go/gentype"
)

// fakeGlobalAlerts implements GlobalAlertInterface
type fakeGlobalAlerts struct {
	*gentype.FakeClientWithList[*v3.GlobalAlert, *v3.GlobalAlertList]
	Fake *FakeProjectcalicoV3
}

func newFakeGlobalAlerts(fake *FakeProjectcalicoV3) projectcalicov3.GlobalAlertInterface {
	return &fakeGlobalAlerts{
		gentype.NewFakeClientWithList[*v3.GlobalAlert, *v3.GlobalAlertList](
			fake.Fake,
			"",
			v3.SchemeGroupVersion.WithResource("globalalerts"),
			v3.SchemeGroupVersion.WithKind("GlobalAlert"),
			func() *v3.GlobalAlert { return &v3.GlobalAlert{} },
			func() *v3.GlobalAlertList { return &v3.GlobalAlertList{} },
			func(dst, src *v3.GlobalAlertList) { dst.ListMeta = src.ListMeta },
			func(list *v3.GlobalAlertList) []*v3.GlobalAlert { return gentype.ToPointerSlice(list.Items) },
			func(list *v3.GlobalAlertList, items []*v3.GlobalAlert) { list.Items = gentype.FromPointerSlice(items) },
		),
		fake,
	}
}
