package dhcp

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"

	coremodel "github.com/infobloxopen/terraform-provider-infoblox/internal/core/model/dhcp"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
	customvalidator "github.com/infobloxopen/terraform-provider-infoblox/internal/validator"
)

type Ipv6rangetemplateModel struct {
	Id   types.String `tfsdk:"id"`
	NIOS types.Object `tfsdk:"nios"`
}

var Ipv6rangetemplateAttrTypes = map[string]attr.Type{
	"id":   types.StringType,
	"nios": types.ObjectType{AttrTypes: NIOSIpv6rangetemplateAttrTypes},
}

type NIOSIpv6rangetemplateModel struct {
	CloudApiCompatible    types.Bool   `tfsdk:"cloud_api_compatible"`
	Comment               types.String `tfsdk:"comment"`
	DelegatedMember       types.Object `tfsdk:"delegated_member"`
	Exclude               types.List   `tfsdk:"exclude"`
	LogicFilterRules      types.List   `tfsdk:"logic_filter_rules"`
	Member                types.Object `tfsdk:"member"`
	Name                  types.String `tfsdk:"name"`
	NumberOfAddresses     types.Int64  `tfsdk:"number_of_addresses"`
	Offset                types.Int64  `tfsdk:"offset"`
	OptionFilterRules     types.List   `tfsdk:"option_filter_rules"`
	RecycleLeases         types.Bool   `tfsdk:"recycle_leases"`
	ServerAssociationType types.String `tfsdk:"server_association_type"`
}

var NIOSIpv6rangetemplateAttrTypes = map[string]attr.Type{
	"cloud_api_compatible":    types.BoolType,
	"comment":                 types.StringType,
	"delegated_member":        types.ObjectType{AttrTypes: Ipv6rangetemplateDelegatedMemberAttrTypes},
	"exclude":                 types.ListType{ElemType: types.ObjectType{AttrTypes: Ipv6rangetemplateExcludeAttrTypes}},
	"logic_filter_rules":      types.ListType{ElemType: types.ObjectType{AttrTypes: Ipv6rangetemplateLogicFilterRulesAttrTypes}},
	"member":                  types.ObjectType{AttrTypes: Ipv6rangetemplateMemberAttrTypes},
	"name":                    types.StringType,
	"number_of_addresses":     types.Int64Type,
	"offset":                  types.Int64Type,
	"option_filter_rules":     types.ListType{ElemType: types.ObjectType{AttrTypes: Ipv6rangetemplateOptionFilterRulesAttrTypes}},
	"recycle_leases":          types.BoolType,
	"server_association_type": types.StringType,
}

const (
	Ipv6rangetemplateReturnFields = "cloud_api_compatible,comment,delegated_member,exclude,logic_filter_rules,member,name,number_of_addresses,offset,option_filter_rules,recycle_leases,server_association_type,use_logic_filter_rules,use_recycle_leases"
)

var Ipv6rangetemplateResourceSchemaAttributes = map[string]schema.Attribute{
	"id": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The reference to the object.",
	},
	"nios": schema.SingleNestedAttribute{
		Optional:            true,
		MarkdownDescription: "NIOS backend-specific fields.",
		Attributes:          Ipv6rangetemplateResourceNiosSchemaAttributes,
	},
}

var Ipv6rangetemplateResourceNiosSchemaAttributes = map[string]schema.Attribute{
	"cloud_api_compatible": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(true),
		MarkdownDescription: "Determines whether the IPv6 DHCP range template can be used to create network objects in a cloud-computing deployment.",
	},
	"comment": schema.StringAttribute{
		Optional: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
			customvalidator.ValidateTrimmedString(),
		},
		MarkdownDescription: "The IPv6 DHCP range template descriptive comment.",
	},
	"delegated_member": schema.SingleNestedAttribute{
		Attributes:          Ipv6rangetemplateDelegatedMemberResourceSchemaAttributes,
		Optional:            true,
		MarkdownDescription: "",
	},
	"exclude": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: Ipv6rangetemplateExcludeResourceSchemaAttributes,
		},
		Optional: true,
		Validators: []validator.List{
			customvalidator.ListNotEmpty(),
		},
		MarkdownDescription: "These are ranges of IPv6 addresses that the appliance does not use to assign to clients. You can use these excluded addresses as static IPv6 addresses. They contain the start and end addresses of the excluded range, and optionally, information about this excluded range.",
	},
	"logic_filter_rules": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: Ipv6rangetemplateLogicFilterRulesResourceSchemaAttributes,
		},
		Optional: true,
		Validators: []validator.List{
			customvalidator.ListNotEmpty(),
		},
		MarkdownDescription: "This field contains the logic filters to be applied on this IPv6 range. This list corresponds to the match rules that are written to the DHCPv6 configuration file.",
	},
	"member": schema.SingleNestedAttribute{
		Attributes:          Ipv6rangetemplateMemberResourceSchemaAttributes,
		Optional:            true,
		MarkdownDescription: "",
	},
	"name": schema.StringAttribute{
		Required: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
			customvalidator.ValidateTrimmedString(),
		},
		MarkdownDescription: "Name of the IPv6 DHCP range template.",
	},
	"number_of_addresses": schema.Int64Attribute{
		Required:            true,
		MarkdownDescription: "The number of addresses for the IPv6 DHCP range.",
	},
	"offset": schema.Int64Attribute{
		Required:            true,
		MarkdownDescription: "The start address offset for the IPv6 DHCP range.",
	},
	"option_filter_rules": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: Ipv6rangetemplateOptionFilterRulesResourceSchemaAttributes,
		},
		Optional: true,
		Validators: []validator.List{
			customvalidator.ListNotEmpty(),
		},
		MarkdownDescription: "This field contains the Option filters to be applied to this IPv6 range. The appliance uses the matching rules of these filters to select the address range from which it assigns a lease.",
	},
	"recycle_leases": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(true),
		MarkdownDescription: "Determines whether the leases are kept in Recycle Bin until one week after expiry. If this is set to False, the leases are permanently deleted.",
	},
	"server_association_type": schema.StringAttribute{
		Default: stringdefault.StaticString("NONE"),
		Validators: []validator.String{
			stringvalidator.OneOf("NONE", "MEMBER", "FAILOVER", "MS_SERVER", "MS_FAILOVER"),
		},
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The type of server that is going to serve the IPv6 DHCP range.",
	},
}

// Expand converts the TF model to the infoblox core model
func (m *Ipv6rangetemplateModel) Expand(ctx context.Context, diags *diag.Diagnostics, isCreate bool) *coremodel.Ipv6rangetemplate {
	if m == nil {
		return nil
	}

	obj := &coremodel.Ipv6rangetemplate{}

	// Expand NIOS nested attribute (returns nil if not present)
	niosModel := flex.ExpandNestedObject[NIOSIpv6rangetemplateModel](ctx, m.NIOS, diags)
	if niosModel != nil {
		obj.NIOS = niosModel.Expand(ctx, diags)
	}

	return obj
}

// Expand converts the NIOS TF model to the core model.
func (m *NIOSIpv6rangetemplateModel) Expand(ctx context.Context, diags *diag.Diagnostics) *coremodel.NIOSIpv6rangetemplateExt {
	return &coremodel.NIOSIpv6rangetemplateExt{
		CloudApiCompatible:    flex.ExpandBoolPointer(m.CloudApiCompatible),
		Comment:               flex.ExpandStringPointerNullAsEmpty(m.Comment),
		DelegatedMember:       ExpandIpv6rangetemplateDelegatedMember(ctx, m.DelegatedMember, diags),
		Exclude:               flex.ExpandFrameworkListNestedBlock(ctx, m.Exclude, diags, ExpandIpv6rangetemplateExclude),
		LogicFilterRules:      flex.ExpandFrameworkListNestedBlock(ctx, m.LogicFilterRules, diags, ExpandIpv6rangetemplateLogicFilterRules),
		Member:                ExpandIpv6rangetemplateMember(ctx, m.Member, diags),
		Name:                  flex.ExpandStringPointerNullAsEmpty(m.Name),
		NumberOfAddresses:     flex.ExpandInt64Pointer(m.NumberOfAddresses),
		Offset:                flex.ExpandInt64Pointer(m.Offset),
		OptionFilterRules:     flex.ExpandFrameworkListNestedBlock(ctx, m.OptionFilterRules, diags, ExpandIpv6rangetemplateOptionFilterRules),
		RecycleLeases:         flex.ExpandBoolPointer(m.RecycleLeases),
		ServerAssociationType: flex.ExpandStringPointerNullAsEmpty(m.ServerAssociationType),
	}
}

// ApplyIpv6rangetemplateNIOSUseFlags derives NIOS use flags from the raw config
// value(s) and writes them onto the core model. A flag is true when the user
// set any of its governed value fields in config.
func ApplyIpv6rangetemplateNIOSUseFlags(ctx context.Context, config tfsdk.Config, obj *coremodel.Ipv6rangetemplate, diags *diag.Diagnostics) {
	if obj == nil || obj.NIOS == nil {
		return
	}
	obj.NIOS.UseLogicFilterRules = flex.DeriveUseFlag(ctx, config, diags, path.Root("nios").AtName("logic_filter_rules"))
	obj.NIOS.UseRecycleLeases = flex.DeriveUseFlag(ctx, config, diags, path.Root("nios").AtName("recycle_leases"))
}

// Flatten populates the TF model from a core response.
func (m *Ipv6rangetemplateModel) Flatten(ctx context.Context, resp *coremodel.Ipv6rangetemplate, diags *diag.Diagnostics) {
	if resp == nil {
		return
	}

	m.Id = flex.FlattenStringPointer(resp.Id)

	// Extract existing NIOS model, flatten API response onto it, convert back
	niosModel := flex.ExpandNestedObject[NIOSIpv6rangetemplateModel](ctx, m.NIOS, diags)
	if niosModel == nil {
		niosModel = &NIOSIpv6rangetemplateModel{}
	}
	niosModel.Flatten(ctx, resp.NIOS, diags)
	if resp.NIOS != nil {
		m.NIOS = flex.FlattenNestedObject(ctx, niosModel, NIOSIpv6rangetemplateAttrTypes, diags)
	} else {
		m.NIOS = types.ObjectNull(NIOSIpv6rangetemplateAttrTypes)
	}

}

// Flatten merges API response onto existing NIOS model.
func (m *NIOSIpv6rangetemplateModel) Flatten(ctx context.Context, from *coremodel.NIOSIpv6rangetemplateExt, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	m.CloudApiCompatible = flex.FlattenBoolPointer(from.CloudApiCompatible)
	m.Comment = flex.FlattenStringPointerEmptyAsNull(from.Comment)
	m.DelegatedMember = FlattenIpv6rangetemplateDelegatedMember(ctx, from.DelegatedMember, diags)
	m.Exclude = flex.FlattenFrameworkListNestedBlock(ctx, from.Exclude, Ipv6rangetemplateExcludeAttrTypes, diags, FlattenIpv6rangetemplateExclude)
	m.LogicFilterRules = flex.FlattenFrameworkListNestedBlock(ctx, from.LogicFilterRules, Ipv6rangetemplateLogicFilterRulesAttrTypes, diags, FlattenIpv6rangetemplateLogicFilterRules)
	m.Member = FlattenIpv6rangetemplateMember(ctx, from.Member, diags)
	m.Name = flex.FlattenStringPointerEmptyAsNull(from.Name)
	m.NumberOfAddresses = flex.FlattenInt64Pointer(from.NumberOfAddresses)
	m.Offset = flex.FlattenInt64Pointer(from.Offset)
	m.OptionFilterRules = flex.FlattenFrameworkListNestedBlock(ctx, from.OptionFilterRules, Ipv6rangetemplateOptionFilterRulesAttrTypes, diags, FlattenIpv6rangetemplateOptionFilterRules)
	m.RecycleLeases = flex.FlattenBoolPointer(from.RecycleLeases)
	m.ServerAssociationType = flex.FlattenStringPointerEmptyAsNull(from.ServerAssociationType)
}
