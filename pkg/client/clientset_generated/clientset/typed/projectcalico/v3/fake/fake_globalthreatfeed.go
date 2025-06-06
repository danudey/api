// Copyright (c) 2025 Tigera, Inc. All rights reserved.

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v3 "github.com/tigera/api/pkg/apis/projectcalico/v3"
	projectcalicov3 "github.com/tigera/api/pkg/client/clientset_generated/clientset/typed/projectcalico/v3"
	gentype "k8s.io/client-go/gentype"
)

// fakeGlobalThreatFeeds implements GlobalThreatFeedInterface
type fakeGlobalThreatFeeds struct {
	*gentype.FakeClientWithList[*v3.GlobalThreatFeed, *v3.GlobalThreatFeedList]
	Fake *FakeProjectcalicoV3
}

func newFakeGlobalThreatFeeds(fake *FakeProjectcalicoV3) projectcalicov3.GlobalThreatFeedInterface {
	return &fakeGlobalThreatFeeds{
		gentype.NewFakeClientWithList[*v3.GlobalThreatFeed, *v3.GlobalThreatFeedList](
			fake.Fake,
			"",
			v3.SchemeGroupVersion.WithResource("globalthreatfeeds"),
			v3.SchemeGroupVersion.WithKind("GlobalThreatFeed"),
			func() *v3.GlobalThreatFeed { return &v3.GlobalThreatFeed{} },
			func() *v3.GlobalThreatFeedList { return &v3.GlobalThreatFeedList{} },
			func(dst, src *v3.GlobalThreatFeedList) { dst.ListMeta = src.ListMeta },
			func(list *v3.GlobalThreatFeedList) []*v3.GlobalThreatFeed { return gentype.ToPointerSlice(list.Items) },
			func(list *v3.GlobalThreatFeedList, items []*v3.GlobalThreatFeed) {
				list.Items = gentype.FromPointerSlice(items)
			},
		),
		fake,
	}
}
