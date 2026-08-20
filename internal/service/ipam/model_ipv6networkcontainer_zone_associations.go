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

// Ipv6networkcontainerZoneAssociationsModel is the Terraform model for Ipv6networkcontainerZoneAssociations
type Ipv6networkcontainerZoneAssociationsModel struct {
	Fqdn      types.String `tfsdk:"fqdn"`
	IsDefault types.Bool   `tfsdk:"is_default"`
	View      types.String `tfsdk:"view"`
}

// Ipv6networkcontainerZoneAssociationsAttrTypes contains the attribute types for Ipv6networkcontainerZoneAssociationsModel
var Ipv6networkcontainerZoneAssociationsAttrTypes = map[string]attr.Type{
	"fqdn":       types.StringType,
	"is_default": types.BoolType,
	"view":       types.StringType,
}

// Ipv6networkcontainerZoneAssociationsResourceSchemaAttributes contains the schema attributes for Ipv6networkcontainerZoneAssociationsModel
var Ipv6networkcontainerZoneAssociationsResourceSchemaAttributes = map[string]schema.Attribute{
	"fqdn": schema.StringAttribute{
		Optional: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
			customvalidator.IsValidNIOSDomainName(),
		},
		MarkdownDescription: "The FQDN of the authoritative forward zone.",
	},
	"is_default": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "True if this is the default zone.",
	},
	"view": schema.StringAttribute{
		Optional: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
			customvalidator.ValidateTrimmedString(),
		},
		MarkdownDescription: "The view to which the zone belongs. If a view is not specified, the default view is used.",
	},
}

// ExpandIpv6networkcontainerZoneAssociations converts a Terraform Object to SDK type
func ExpandIpv6networkcontainerZoneAssociations(ctx context.Context, o types.Object, diags *diag.Diagnostics) *niosipam.Ipv6networkcontainerZoneAssociations {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m Ipv6networkcontainerZoneAssociationsModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

// Expand converts the Terraform model to SDK type
func (m *Ipv6networkcontainerZoneAssociationsModel) Expand(ctx context.Context, diags *diag.Diagnostics) *niosipam.Ipv6networkcontainerZoneAssociations {
	if m == nil {
		return nil
	}
	to := &niosipam.Ipv6networkcontainerZoneAssociations{
		Fqdn:      flex.ExpandStringPointerNullAsEmpty(m.Fqdn),
		IsDefault: flex.ExpandBoolPointer(m.IsDefault),
		View:      flex.ExpandStringPointerNullAsEmpty(m.View),
	}
	return to
}

// FlattenIpv6networkcontainerZoneAssociations converts an SDK type to Terraform Object
func FlattenIpv6networkcontainerZoneAssociations(ctx context.Context, from *niosipam.Ipv6networkcontainerZoneAssociations, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(Ipv6networkcontainerZoneAssociationsAttrTypes)
	}
	m := &Ipv6networkcontainerZoneAssociationsModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, Ipv6networkcontainerZoneAssociationsAttrTypes, m)
	diags.Append(d...)
	return t
}

// Flatten populates the Terraform model from SDK type
func (m *Ipv6networkcontainerZoneAssociationsModel) Flatten(ctx context.Context, from *niosipam.Ipv6networkcontainerZoneAssociations, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	m.Fqdn = flex.FlattenStringPointerEmptyAsNull(from.Fqdn)
	m.IsDefault = flex.FlattenBoolPointer(from.IsDefault)
	m.View = flex.FlattenStringPointerEmptyAsNull(from.View)
}
