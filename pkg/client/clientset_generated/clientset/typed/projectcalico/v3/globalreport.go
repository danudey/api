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

// GlobalReportsGetter has a method to return a GlobalReportInterface.
// A group's client should implement this interface.
type GlobalReportsGetter interface {
	GlobalReports() GlobalReportInterface
}

// GlobalReportInterface has methods to work with GlobalReport resources.
type GlobalReportInterface interface {
	Create(ctx context.Context, globalReport *projectcalicov3.GlobalReport, opts v1.CreateOptions) (*projectcalicov3.GlobalReport, error)
	Update(ctx context.Context, globalReport *projectcalicov3.GlobalReport, opts v1.UpdateOptions) (*projectcalicov3.GlobalReport, error)
	// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
	UpdateStatus(ctx context.Context, globalReport *projectcalicov3.GlobalReport, opts v1.UpdateOptions) (*projectcalicov3.GlobalReport, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*projectcalicov3.GlobalReport, error)
	List(ctx context.Context, opts v1.ListOptions) (*projectcalicov3.GlobalReportList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *projectcalicov3.GlobalReport, err error)
	GlobalReportExpansion
}

// globalReports implements GlobalReportInterface
type globalReports struct {
	*gentype.ClientWithList[*projectcalicov3.GlobalReport, *projectcalicov3.GlobalReportList]
}

// newGlobalReports returns a GlobalReports
func newGlobalReports(c *ProjectcalicoV3Client) *globalReports {
	return &globalReports{
		gentype.NewClientWithList[*projectcalicov3.GlobalReport, *projectcalicov3.GlobalReportList](
			"globalreports",
			c.RESTClient(),
			scheme.ParameterCodec,
			"",
			func() *projectcalicov3.GlobalReport { return &projectcalicov3.GlobalReport{} },
			func() *projectcalicov3.GlobalReportList { return &projectcalicov3.GlobalReportList{} },
		),
	}
}
