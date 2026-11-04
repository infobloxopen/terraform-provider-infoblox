package dhcp

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	niosdhcp "github.com/infobloxopen/infoblox-nios-go-client/dhcp"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
	customvalidator "github.com/infobloxopen/terraform-provider-infoblox/internal/validator"
)

// Ipv6rangesubscribesettingsMappedEaAttributesModel is the Terraform model for Ipv6rangesubscribesettingsMappedEaAttributes
type Ipv6rangesubscribesettingsMappedEaAttributesModel struct {
	Name     types.String `tfsdk:"name"`
	MappedEa types.String `tfsdk:"mapped_ea"`
}

// Ipv6rangesubscribesettingsMappedEaAttributesAttrTypes contains the attribute types for Ipv6rangesubscribesettingsMappedEaAttributesModel
var Ipv6rangesubscribesettingsMappedEaAttributesAttrTypes = map[string]attr.Type{
	"name":      types.StringType,
	"mapped_ea": types.StringType,
}

// Ipv6rangesubscribesettingsMappedEaAttributesResourceSchemaAttributes contains the schema attributes for Ipv6rangesubscribesettingsMappedEaAttributesModel
var Ipv6rangesubscribesettingsMappedEaAttributesResourceSchemaAttributes = map[string]schema.Attribute{
	"name": schema.StringAttribute{
		Validators: []validator.String{
			stringvalidator.OneOf("ACCOUNT_SESSION_ID", "AUDIT_SESSION_ID", "EPS_STATUS", "IP_ADDRESS", "MAC", "NAS_IP_ADDRESS", "NAS_PORT_ID", "POSTURE_STATUS", "POSTURE_TIMESTAMP"),
		},
		Optional:            true,
		MarkdownDescription: "The Cisco ISE attribute name that is enabled for publishsing from a Cisco ISE endpoint.",
	},
	"mapped_ea": schema.StringAttribute{
		Optional: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
		},
		MarkdownDescription: "The name of the extensible attribute definition object the Cisco ISE attribute that is enabled for subscription is mapped on.",
	},
}

// ExpandIpv6rangesubscribesettingsMappedEaAttributes converts a Terraform Object to SDK type
func ExpandIpv6rangesubscribesettingsMappedEaAttributes(ctx context.Context, o types.Object, diags *diag.Diagnostics) *niosdhcp.Ipv6rangesubscribesettingsMappedEaAttributes {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m Ipv6rangesubscribesettingsMappedEaAttributesModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

// Expand converts the Terraform model to SDK type
func (m *Ipv6rangesubscribesettingsMappedEaAttributesModel) Expand(ctx context.Context, diags *diag.Diagnostics) *niosdhcp.Ipv6rangesubscribesettingsMappedEaAttributes {
	if m == nil {
		return nil
	}
	to := &niosdhcp.Ipv6rangesubscribesettingsMappedEaAttributes{
		Name:     flex.ExpandStringPointerNullAsEmpty(m.Name),
		MappedEa: flex.ExpandStringPointerNullAsEmpty(m.MappedEa),
	}
	return to
}

// FlattenIpv6rangesubscribesettingsMappedEaAttributes converts an SDK type to Terraform Object
func FlattenIpv6rangesubscribesettingsMappedEaAttributes(ctx context.Context, from *niosdhcp.Ipv6rangesubscribesettingsMappedEaAttributes, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(Ipv6rangesubscribesettingsMappedEaAttributesAttrTypes)
	}
	m := &Ipv6rangesubscribesettingsMappedEaAttributesModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, Ipv6rangesubscribesettingsMappedEaAttributesAttrTypes, m)
	diags.Append(d...)
	return t
}

// Flatten populates the Terraform model from SDK type
func (m *Ipv6rangesubscribesettingsMappedEaAttributesModel) Flatten(ctx context.Context, from *niosdhcp.Ipv6rangesubscribesettingsMappedEaAttributes, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	m.Name = flex.FlattenStringPointerEmptyAsNull(from.Name)
	m.MappedEa = flex.FlattenStringPointerEmptyAsNull(from.MappedEa)
}
