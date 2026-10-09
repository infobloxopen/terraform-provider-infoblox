package dtc

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
	uddidtc "github.com/infobloxopen/universal-ddi-go-client/dtc"
)

// DesignatorServiceModel is the Terraform model for DesignatorService
type DesignatorServiceModel struct {
	DnsServiceId   types.String `tfsdk:"dns_service_id"`
	DnsServiceName types.String `tfsdk:"dns_service_name"`
}

// DesignatorServiceAttrTypes contains the attribute types for DesignatorServiceModel
var DesignatorServiceAttrTypes = map[string]attr.Type{
	"dns_service_id":   types.StringType,
	"dns_service_name": types.StringType,
}

// DesignatorServiceResourceSchemaAttributes contains the schema attributes for DesignatorServiceModel
var DesignatorServiceResourceSchemaAttributes = map[string]schema.Attribute{
	"dns_service_id": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The resource identifier.",
	},
	"dns_service_name": schema.StringAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "Display name of the __DNS Service__. Response-only; ignored on request.",
	},
}

// ExpandDesignatorService converts a Terraform Object to SDK type
func ExpandDesignatorService(ctx context.Context, o types.Object, diags *diag.Diagnostics) *uddidtc.DesignatorService {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m DesignatorServiceModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

// Expand converts the Terraform model to SDK type
func (m *DesignatorServiceModel) Expand(ctx context.Context, diags *diag.Diagnostics) *uddidtc.DesignatorService {
	if m == nil {
		return nil
	}
	to := &uddidtc.DesignatorService{
		DnsServiceId:   flex.ExpandString(m.DnsServiceId),
		DnsServiceName: flex.ExpandStringPointer(m.DnsServiceName),
	}
	return to
}

// FlattenDesignatorService converts an SDK type to Terraform Object
func FlattenDesignatorService(ctx context.Context, from *uddidtc.DesignatorService, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(DesignatorServiceAttrTypes)
	}
	m := &DesignatorServiceModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, DesignatorServiceAttrTypes, m)
	diags.Append(d...)
	return t
}

// Flatten populates the Terraform model from SDK type
func (m *DesignatorServiceModel) Flatten(ctx context.Context, from *uddidtc.DesignatorService, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	m.DnsServiceId = flex.FlattenString(from.DnsServiceId)
	m.DnsServiceName = flex.FlattenStringPointer(from.DnsServiceName)
}
