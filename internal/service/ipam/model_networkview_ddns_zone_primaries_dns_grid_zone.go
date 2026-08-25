package ipam

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	niosipam "github.com/infobloxopen/infoblox-nios-go-client/ipam"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
	customvalidator "github.com/infobloxopen/terraform-provider-infoblox/internal/validator"
)

// NetworkviewDdnsZonePrimariesDnsGridZoneModel is the Terraform model for NetworkviewDdnsZonePrimariesDnsGridZone
type NetworkviewDdnsZonePrimariesDnsGridZoneModel struct {
	Ref types.String `tfsdk:"ref"`
}

// NetworkviewDdnsZonePrimariesDnsGridZoneAttrTypes contains the attribute types for NetworkviewDdnsZonePrimariesDnsGridZoneModel
var NetworkviewDdnsZonePrimariesDnsGridZoneAttrTypes = map[string]attr.Type{
	"ref": types.StringType,
}

// NetworkviewDdnsZonePrimariesDnsGridZoneResourceSchemaAttributes contains the schema attributes for NetworkviewDdnsZonePrimariesDnsGridZoneModel
var NetworkviewDdnsZonePrimariesDnsGridZoneResourceSchemaAttributes = map[string]schema.Attribute{
	"ref": schema.StringAttribute{
		Required: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
		},
		MarkdownDescription: "The reference to the DNS zone object.",
	},
}

// ExpandNetworkviewDdnsZonePrimariesDnsGridZone converts a Terraform Object to SDK type
func ExpandNetworkviewDdnsZonePrimariesDnsGridZone(ctx context.Context, o types.Object, diags *diag.Diagnostics) *niosipam.NetworkviewDdnsZonePrimariesDnsGridZone {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m NetworkviewDdnsZonePrimariesDnsGridZoneModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

// Expand converts the Terraform model to SDK type
func (m *NetworkviewDdnsZonePrimariesDnsGridZoneModel) Expand(ctx context.Context, diags *diag.Diagnostics) *niosipam.NetworkviewDdnsZonePrimariesDnsGridZone {
	if m == nil {
		return nil
	}
	to := &niosipam.NetworkviewDdnsZonePrimariesDnsGridZone{
		Ref: flex.ExpandStringPointerNullAsEmpty(m.Ref),
	}
	return to
}

// FlattenNetworkviewDdnsZonePrimariesDnsGridZone converts an SDK type to Terraform Object
func FlattenNetworkviewDdnsZonePrimariesDnsGridZone(ctx context.Context, from *niosipam.NetworkviewDdnsZonePrimariesDnsGridZone, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(NetworkviewDdnsZonePrimariesDnsGridZoneAttrTypes)
	}
	m := &NetworkviewDdnsZonePrimariesDnsGridZoneModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, NetworkviewDdnsZonePrimariesDnsGridZoneAttrTypes, m)
	diags.Append(d...)
	return t
}

// Flatten populates the Terraform model from SDK type
func (m *NetworkviewDdnsZonePrimariesDnsGridZoneModel) Flatten(ctx context.Context, from *niosipam.NetworkviewDdnsZonePrimariesDnsGridZone, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	m.Ref = flex.FlattenStringPointerEmptyAsNull(from.Ref)
}
