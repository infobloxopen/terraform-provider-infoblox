package dns

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	uddidns "github.com/infobloxopen/universal-ddi-go-client/dnsconfig"
)

// AuthZoneInheritanceModel is the Terraform model for AuthZoneInheritance
type AuthZoneInheritanceModel struct {
	GssTsigEnabled           types.Object `tfsdk:"gss_tsig_enabled"`
	Notify                   types.Object `tfsdk:"notify"`
	QueryAcl                 types.Object `tfsdk:"query_acl"`
	TransferAcl              types.Object `tfsdk:"transfer_acl"`
	UpdateAcl                types.Object `tfsdk:"update_acl"`
	UseForwardersForSubzones types.Object `tfsdk:"use_forwarders_for_subzones"`
	ZoneAuthority            types.Object `tfsdk:"zone_authority"`
}

// AuthZoneInheritanceAttrTypes contains the attribute types for AuthZoneInheritanceModel
var AuthZoneInheritanceAttrTypes = map[string]attr.Type{
	"gss_tsig_enabled":            types.ObjectType{AttrTypes: Inheritance2InheritedBoolAttrTypes},
	"notify":                      types.ObjectType{AttrTypes: Inheritance2InheritedBoolAttrTypes},
	"query_acl":                   types.ObjectType{AttrTypes: InheritedACLItemsAttrTypes},
	"transfer_acl":                types.ObjectType{AttrTypes: InheritedACLItemsAttrTypes},
	"update_acl":                  types.ObjectType{AttrTypes: InheritedACLItemsAttrTypes},
	"use_forwarders_for_subzones": types.ObjectType{AttrTypes: Inheritance2InheritedBoolAttrTypes},
	"zone_authority":              types.ObjectType{AttrTypes: InheritedZoneAuthorityAttrTypes},
}

// AuthZoneInheritanceResourceSchemaAttributes contains the schema attributes for AuthZoneInheritanceModel
var AuthZoneInheritanceResourceSchemaAttributes = map[string]schema.Attribute{
	"gss_tsig_enabled": schema.SingleNestedAttribute{
		Attributes:          Inheritance2InheritedBoolResourceSchemaAttributes,
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "Optional. Field config for _gss_tsig_enabled_ field from _AuthZone_ object.",
	},
	"notify": schema.SingleNestedAttribute{
		Attributes:          Inheritance2InheritedBoolResourceSchemaAttributes,
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "Field config for _notify_ field from _AuthZone_ object.",
	},
	"query_acl": schema.SingleNestedAttribute{
		Attributes:          InheritedACLItemsResourceSchemaAttributes,
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "Optional. Field config for _query_acl_ field from _AuthZone_ object.",
	},
	"transfer_acl": schema.SingleNestedAttribute{
		Attributes:          InheritedACLItemsResourceSchemaAttributes,
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "Optional. Field config for _transfer_acl_ field from _AuthZone_ object.",
	},
	"update_acl": schema.SingleNestedAttribute{
		Attributes:          InheritedACLItemsResourceSchemaAttributes,
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "Optional. Field config for _update_acl_ field from _AuthZone_ object.",
	},
	"use_forwarders_for_subzones": schema.SingleNestedAttribute{
		Attributes:          Inheritance2InheritedBoolResourceSchemaAttributes,
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "Optional. Field config for _use_forwarders_for_subzones_ field from _AuthZone_ object.",
	},
	"zone_authority": schema.SingleNestedAttribute{
		Attributes:          InheritedZoneAuthorityResourceSchemaAttributes,
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "Optional. Field config for _zone_authority_ field from _AuthZone_ object.",
	},
}

// ExpandAuthZoneInheritance converts a Terraform Object to SDK type
func ExpandAuthZoneInheritance(ctx context.Context, o types.Object, diags *diag.Diagnostics) *uddidns.AuthZoneInheritance {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m AuthZoneInheritanceModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

// Expand converts the Terraform model to SDK type
func (m *AuthZoneInheritanceModel) Expand(ctx context.Context, diags *diag.Diagnostics) *uddidns.AuthZoneInheritance {
	if m == nil {
		return nil
	}
	to := &uddidns.AuthZoneInheritance{
		GssTsigEnabled:           ExpandInheritance2InheritedBool(ctx, m.GssTsigEnabled, diags),
		Notify:                   ExpandInheritance2InheritedBool(ctx, m.Notify, diags),
		QueryAcl:                 ExpandInheritedACLItems(ctx, m.QueryAcl, diags),
		TransferAcl:              ExpandInheritedACLItems(ctx, m.TransferAcl, diags),
		UpdateAcl:                ExpandInheritedACLItems(ctx, m.UpdateAcl, diags),
		UseForwardersForSubzones: ExpandInheritance2InheritedBool(ctx, m.UseForwardersForSubzones, diags),
		ZoneAuthority:            ExpandInheritedZoneAuthority(ctx, m.ZoneAuthority, diags),
	}
	return to
}

// FlattenAuthZoneInheritance converts an SDK type to Terraform Object
func FlattenAuthZoneInheritance(ctx context.Context, from *uddidns.AuthZoneInheritance, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(AuthZoneInheritanceAttrTypes)
	}
	m := &AuthZoneInheritanceModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, AuthZoneInheritanceAttrTypes, m)
	diags.Append(d...)
	return t
}

// Flatten populates the Terraform model from SDK type
func (m *AuthZoneInheritanceModel) Flatten(ctx context.Context, from *uddidns.AuthZoneInheritance, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	m.GssTsigEnabled = FlattenInheritance2InheritedBool(ctx, from.GssTsigEnabled, diags)
	m.Notify = FlattenInheritance2InheritedBool(ctx, from.Notify, diags)
	m.QueryAcl = FlattenInheritedACLItems(ctx, from.QueryAcl, diags)
	m.TransferAcl = FlattenInheritedACLItems(ctx, from.TransferAcl, diags)
	m.UpdateAcl = FlattenInheritedACLItems(ctx, from.UpdateAcl, diags)
	m.UseForwardersForSubzones = FlattenInheritance2InheritedBool(ctx, from.UseForwardersForSubzones, diags)
	m.ZoneAuthority = FlattenInheritedZoneAuthority(ctx, from.ZoneAuthority, diags)
}
