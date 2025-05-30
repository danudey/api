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

// PacketCapturesGetter has a method to return a PacketCaptureInterface.
// A group's client should implement this interface.
type PacketCapturesGetter interface {
	PacketCaptures(namespace string) PacketCaptureInterface
}

// PacketCaptureInterface has methods to work with PacketCapture resources.
type PacketCaptureInterface interface {
	Create(ctx context.Context, packetCapture *projectcalicov3.PacketCapture, opts v1.CreateOptions) (*projectcalicov3.PacketCapture, error)
	Update(ctx context.Context, packetCapture *projectcalicov3.PacketCapture, opts v1.UpdateOptions) (*projectcalicov3.PacketCapture, error)
	// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
	UpdateStatus(ctx context.Context, packetCapture *projectcalicov3.PacketCapture, opts v1.UpdateOptions) (*projectcalicov3.PacketCapture, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*projectcalicov3.PacketCapture, error)
	List(ctx context.Context, opts v1.ListOptions) (*projectcalicov3.PacketCaptureList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *projectcalicov3.PacketCapture, err error)
	PacketCaptureExpansion
}

// packetCaptures implements PacketCaptureInterface
type packetCaptures struct {
	*gentype.ClientWithList[*projectcalicov3.PacketCapture, *projectcalicov3.PacketCaptureList]
}

// newPacketCaptures returns a PacketCaptures
func newPacketCaptures(c *ProjectcalicoV3Client, namespace string) *packetCaptures {
	return &packetCaptures{
		gentype.NewClientWithList[*projectcalicov3.PacketCapture, *projectcalicov3.PacketCaptureList](
			"packetcaptures",
			c.RESTClient(),
			scheme.ParameterCodec,
			namespace,
			func() *projectcalicov3.PacketCapture { return &projectcalicov3.PacketCapture{} },
			func() *projectcalicov3.PacketCaptureList { return &projectcalicov3.PacketCaptureList{} },
		),
	}
}
