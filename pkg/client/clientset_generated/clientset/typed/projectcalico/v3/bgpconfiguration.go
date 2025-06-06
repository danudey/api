// Copyright (c) 2025 Tigera, Inc. All rights reserved.

// Code generated by client-gen. DO NOT EDIT.

package v3

import (
	context "context"

	projectcalicov3 "github.com/tigera/api/pkg/apis/projectcalico/v3"
	scheme "github.com/tigera/api/pkg/client/clientset_generated/clientset/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	gentype "k8s.io/client-go/gentype"
)

// BGPConfigurationsGetter has a method to return a BGPConfigurationInterface.
// A group's client should implement this interface.
type BGPConfigurationsGetter interface {
	BGPConfigurations() BGPConfigurationInterface
}

// BGPConfigurationInterface has methods to work with BGPConfiguration resources.
type BGPConfigurationInterface interface {
	Create(ctx context.Context, bGPConfiguration *projectcalicov3.BGPConfiguration, opts v1.CreateOptions) (*projectcalicov3.BGPConfiguration, error)
	Update(ctx context.Context, bGPConfiguration *projectcalicov3.BGPConfiguration, opts v1.UpdateOptions) (*projectcalicov3.BGPConfiguration, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*projectcalicov3.BGPConfiguration, error)
	List(ctx context.Context, opts v1.ListOptions) (*projectcalicov3.BGPConfigurationList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *projectcalicov3.BGPConfiguration, err error)
	BGPConfigurationExpansion
}

// bGPConfigurations implements BGPConfigurationInterface
type bGPConfigurations struct {
	*gentype.ClientWithList[*projectcalicov3.BGPConfiguration, *projectcalicov3.BGPConfigurationList]
}

// newBGPConfigurations returns a BGPConfigurations
func newBGPConfigurations(c *ProjectcalicoV3Client) *bGPConfigurations {
	return &bGPConfigurations{
		gentype.NewClientWithList[*projectcalicov3.BGPConfiguration, *projectcalicov3.BGPConfigurationList](
			"bgpconfigurations",
			c.RESTClient(),
			scheme.ParameterCodec,
			"",
			func() *projectcalicov3.BGPConfiguration { return &projectcalicov3.BGPConfiguration{} },
			func() *projectcalicov3.BGPConfigurationList { return &projectcalicov3.BGPConfigurationList{} },
		),
	}
}
