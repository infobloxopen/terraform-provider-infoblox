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

// InheritedZoneAuthorityModel is the Terraform model for InheritedZoneAuthority
type InheritedZoneAuthorityModel struct {
	DefaultTtl    types.Object `tfsdk:"default_ttl"`
	Expire        types.Object `tfsdk:"expire"`
	MnameBlock    types.Object `tfsdk:"mname_block"`
	NegativeTtl   types.Object `tfsdk:"negative_ttl"`
	ProtocolRname types.Object `tfsdk:"protocol_rname"`
	Refresh       types.Object `tfsdk:"refresh"`
	Retry         types.Object `tfsdk:"retry"`
	Rname         types.Object `tfsdk:"rname"`
}

// InheritedZoneAuthorityAttrTypes contains the attribute types for InheritedZoneAuthorityModel
var InheritedZoneAuthorityAttrTypes = map[string]attr.Type{
	"default_ttl":    types.ObjectType{AttrTypes: Inheritance2InheritedUInt32DnsconfigAttrTypes},
	"expire":         types.ObjectType{AttrTypes: Inheritance2InheritedUInt32DnsconfigAttrTypes},
	"mname_block":    types.ObjectType{AttrTypes: InheritedZoneAuthorityMNameBlockAttrTypes},
	"negative_ttl":   types.ObjectType{AttrTypes: Inheritance2InheritedUInt32DnsconfigAttrTypes},
	"protocol_rname": types.ObjectType{AttrTypes: Inheritance2InheritedStringAttrTypes},
	"refresh":        types.ObjectType{AttrTypes: Inheritance2InheritedUInt32DnsconfigAttrTypes},
	"retry":          types.ObjectType{AttrTypes: Inheritance2InheritedUInt32DnsconfigAttrTypes},
	"rname":          types.ObjectType{AttrTypes: Inheritance2InheritedStringAttrTypes},
}

// InheritedZoneAuthorityResourceSchemaAttributes contains the schema attributes for InheritedZoneAuthorityModel
var InheritedZoneAuthorityResourceSchemaAttributes = map[string]schema.Attribute{
	"default_ttl": schema.SingleNestedAttribute{
		Attributes:          Inheritance2InheritedUInt32DnsconfigResourceSchemaAttributes,
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "Optional. Field config for _default_ttl_ field from _ZoneAuthority_ object.",
	},
	"expire": schema.SingleNestedAttribute{
		Attributes:          Inheritance2InheritedUInt32DnsconfigResourceSchemaAttributes,
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "Optional. Field config for _expire_ field from _ZoneAuthority_ object.",
	},
	"mname_block": schema.SingleNestedAttribute{
		Attributes:          InheritedZoneAuthorityMNameBlockResourceSchemaAttributes,
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "Optional. Field config for _mname_ block from _ZoneAuthority_ object.",
	},
	"negative_ttl": schema.SingleNestedAttribute{
		Attributes:          Inheritance2InheritedUInt32DnsconfigResourceSchemaAttributes,
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "Optional. Field config for _negative_ttl_ field from _ZoneAuthority_ object.",
	},
	"protocol_rname": schema.SingleNestedAttribute{
		Attributes:          Inheritance2InheritedStringResourceSchemaAttributes,
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "Optional. Field config for _protocol_rname_ field from _ZoneAuthority_ object.",
	},
	"refresh": schema.SingleNestedAttribute{
		Attributes:          Inheritance2InheritedUInt32DnsconfigResourceSchemaAttributes,
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "Optional. Field config for _refresh_ field from _ZoneAuthority_ object.",
	},
	"retry": schema.SingleNestedAttribute{
		Attributes:          Inheritance2InheritedUInt32DnsconfigResourceSchemaAttributes,
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "Optional. Field config for _retry_ field from _ZoneAuthority_ object.",
	},
	"rname": schema.SingleNestedAttribute{
		Attributes:          Inheritance2InheritedStringResourceSchemaAttributes,
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "Optional. Field config for _rname_ field from _ZoneAuthority_ object.",
	},
}

// ExpandInheritedZoneAuthority converts a Terraform Object to SDK type
func ExpandInheritedZoneAuthority(ctx context.Context, o types.Object, diags *diag.Diagnostics) *uddidns.InheritedZoneAuthority {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m InheritedZoneAuthorityModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

// Expand converts the Terraform model to SDK type
func (m *InheritedZoneAuthorityModel) Expand(ctx context.Context, diags *diag.Diagnostics) *uddidns.InheritedZoneAuthority {
	if m == nil {
		return nil
	}
	to := &uddidns.InheritedZoneAuthority{
		DefaultTtl:    ExpandInheritance2InheritedUInt32Dnsconfig(ctx, m.DefaultTtl, diags),
		Expire:        ExpandInheritance2InheritedUInt32Dnsconfig(ctx, m.Expire, diags),
		MnameBlock:    ExpandInheritedZoneAuthorityMNameBlock(ctx, m.MnameBlock, diags),
		NegativeTtl:   ExpandInheritance2InheritedUInt32Dnsconfig(ctx, m.NegativeTtl, diags),
		ProtocolRname: ExpandInheritance2InheritedString(ctx, m.ProtocolRname, diags),
		Refresh:       ExpandInheritance2InheritedUInt32Dnsconfig(ctx, m.Refresh, diags),
		Retry:         ExpandInheritance2InheritedUInt32Dnsconfig(ctx, m.Retry, diags),
		Rname:         ExpandInheritance2InheritedString(ctx, m.Rname, diags),
	}
	return to
}

// FlattenInheritedZoneAuthority converts an SDK type to Terraform Object
func FlattenInheritedZoneAuthority(ctx context.Context, from *uddidns.InheritedZoneAuthority, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(InheritedZoneAuthorityAttrTypes)
	}
	m := &InheritedZoneAuthorityModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, InheritedZoneAuthorityAttrTypes, m)
	diags.Append(d...)
	return t
}

// Flatten populates the Terraform model from SDK type
func (m *InheritedZoneAuthorityModel) Flatten(ctx context.Context, from *uddidns.InheritedZoneAuthority, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	m.DefaultTtl = FlattenInheritance2InheritedUInt32Dnsconfig(ctx, from.DefaultTtl, diags)
	m.Expire = FlattenInheritance2InheritedUInt32Dnsconfig(ctx, from.Expire, diags)
	m.MnameBlock = FlattenInheritedZoneAuthorityMNameBlock(ctx, from.MnameBlock, diags)
	m.NegativeTtl = FlattenInheritance2InheritedUInt32Dnsconfig(ctx, from.NegativeTtl, diags)
	m.ProtocolRname = FlattenInheritance2InheritedString(ctx, from.ProtocolRname, diags)
	m.Refresh = FlattenInheritance2InheritedUInt32Dnsconfig(ctx, from.Refresh, diags)
	m.Retry = FlattenInheritance2InheritedUInt32Dnsconfig(ctx, from.Retry, diags)
	m.Rname = FlattenInheritance2InheritedString(ctx, from.Rname, diags)
}
