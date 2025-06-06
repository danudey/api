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

// LicenseKeysGetter has a method to return a LicenseKeyInterface.
// A group's client should implement this interface.
type LicenseKeysGetter interface {
	LicenseKeys() LicenseKeyInterface
}

// LicenseKeyInterface has methods to work with LicenseKey resources.
type LicenseKeyInterface interface {
	Create(ctx context.Context, licenseKey *projectcalicov3.LicenseKey, opts v1.CreateOptions) (*projectcalicov3.LicenseKey, error)
	Update(ctx context.Context, licenseKey *projectcalicov3.LicenseKey, opts v1.UpdateOptions) (*projectcalicov3.LicenseKey, error)
	// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
	UpdateStatus(ctx context.Context, licenseKey *projectcalicov3.LicenseKey, opts v1.UpdateOptions) (*projectcalicov3.LicenseKey, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*projectcalicov3.LicenseKey, error)
	List(ctx context.Context, opts v1.ListOptions) (*projectcalicov3.LicenseKeyList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *projectcalicov3.LicenseKey, err error)
	LicenseKeyExpansion
}

// licenseKeys implements LicenseKeyInterface
type licenseKeys struct {
	*gentype.ClientWithList[*projectcalicov3.LicenseKey, *projectcalicov3.LicenseKeyList]
}

// newLicenseKeys returns a LicenseKeys
func newLicenseKeys(c *ProjectcalicoV3Client) *licenseKeys {
	return &licenseKeys{
		gentype.NewClientWithList[*projectcalicov3.LicenseKey, *projectcalicov3.LicenseKeyList](
			"licensekeys",
			c.RESTClient(),
			scheme.ParameterCodec,
			"",
			func() *projectcalicov3.LicenseKey { return &projectcalicov3.LicenseKey{} },
			func() *projectcalicov3.LicenseKeyList { return &projectcalicov3.LicenseKeyList{} },
		),
	}
}
