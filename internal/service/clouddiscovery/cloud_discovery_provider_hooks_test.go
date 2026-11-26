package clouddiscovery

import (
	"context"
	"errors"
	"net/http"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/infobloxopen/terraform-provider-infoblox/internal/core"
	coremodel "github.com/infobloxopen/terraform-provider-infoblox/internal/core/model/clouddiscovery"
	coresvc "github.com/infobloxopen/terraform-provider-infoblox/internal/core/service/clouddiscovery"
)

// fakeProviderService returns a fixed response from Update.
type fakeProviderService struct {
	coresvc.CloudDiscoveryProviderService
	status      int
	err         error
	updateCalls int
}

func (f *fakeProviderService) Update(_ context.Context, _ string, obj *coremodel.CloudDiscoveryProvider, _ *core.Options) (*coremodel.CloudDiscoveryProvider, *http.Response, error) {
	f.updateCalls++
	var resp *http.Response
	if f.status != 0 {
		resp = &http.Response{StatusCode: f.status}
	}
	if f.err != nil {
		return nil, resp, f.err
	}
	return obj, resp, nil
}

func runPreDelete(t *testing.T, svc *fakeProviderService) diag.Diagnostics {
	t.Helper()
	r := &CloudDiscoveryProviderResource{service: svc}
	data := &CloudDiscoveryProviderModel{
		Id:            types.StringValue("cloud_discovery/provider/1"),
		UpdateTrigger: types.StringNull(),
		UDDI:          types.ObjectNull(UDDICloudDiscoveryProviderAttrTypes),
	}
	var diags diag.Diagnostics
	r.preDeleteCloudDiscoveryProvider(context.Background(), data, &diags)
	if svc.updateCalls != 1 {
		t.Fatalf("expected one disable call, got %d", svc.updateCalls)
	}
	return diags
}

func TestPreDeleteCloudDiscoveryProvider_NotFoundIsSuccess(t *testing.T) {
	diags := runPreDelete(t, &fakeProviderService{status: http.StatusNotFound, err: errors.New("404 Not Found")})
	if diags.HasError() {
		t.Fatalf("404 on disable should not fail delete: %v", diags)
	}
}

func TestPreDeleteCloudDiscoveryProvider_OtherErrorFails(t *testing.T) {
	diags := runPreDelete(t, &fakeProviderService{status: http.StatusInternalServerError, err: errors.New("500")})
	if !diags.HasError() {
		t.Fatal("expected an error for a 500 on disable")
	}
}

func TestPreDeleteCloudDiscoveryProvider_NoResponseFails(t *testing.T) {
	diags := runPreDelete(t, &fakeProviderService{err: errors.New("connection refused")})
	if !diags.HasError() {
		t.Fatal("expected an error when the disable call fails without a response")
	}
}
