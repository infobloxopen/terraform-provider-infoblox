package dns

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-nettypes/iptypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/mapvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/mapdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"

	coremodel "github.com/infobloxopen/terraform-provider-infoblox/internal/core/model/dns"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
	immutable "github.com/infobloxopen/terraform-provider-infoblox/internal/planmodifiers/immutable"
	importmod "github.com/infobloxopen/terraform-provider-infoblox/internal/planmodifiers/import"
	customvalidator "github.com/infobloxopen/terraform-provider-infoblox/internal/validator"
)

type SharedrecordAaaaModel struct {
	Id   types.String `tfsdk:"id"`
	NIOS types.Object `tfsdk:"nios"`
}

var SharedrecordAaaaAttrTypes = map[string]attr.Type{
	"id":   types.StringType,
	"nios": types.ObjectType{AttrTypes: NIOSSharedrecordAaaaAttrTypes},
}

type NIOSSharedrecordAaaaModel struct {
	Comment           types.String        `tfsdk:"comment"`
	Disable           types.Bool          `tfsdk:"disable"`
	ExtAttrs          types.Map           `tfsdk:"ext_attrs"`
	ExtAttrsAll       types.Map           `tfsdk:"ext_attrs_all"`
	Ipv6addr          iptypes.IPv6Address `tfsdk:"ipv6addr"`
	Name              types.String        `tfsdk:"name"`
	SharedRecordGroup types.String        `tfsdk:"shared_record_group"`
	Ttl               types.Int64         `tfsdk:"ttl"`
}

var NIOSSharedrecordAaaaAttrTypes = map[string]attr.Type{
	"comment":             types.StringType,
	"disable":             types.BoolType,
	"ext_attrs":           types.MapType{ElemType: types.StringType},
	"ext_attrs_all":       types.MapType{ElemType: types.StringType},
	"ipv6addr":            iptypes.IPv6AddressType{},
	"name":                types.StringType,
	"shared_record_group": types.StringType,
	"ttl":                 types.Int64Type,
}

const (
	SharedrecordAaaaReturnFields = "comment,disable,dns_name,extattrs,ipv6addr,name,shared_record_group,ttl,use_ttl"
)

var SharedrecordAaaaResourceSchemaAttributes = map[string]schema.Attribute{
	"id": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The reference to the object.",
	},
	"nios": schema.SingleNestedAttribute{
		Optional:            true,
		MarkdownDescription: "NIOS backend-specific fields.",
		Attributes:          SharedrecordAaaaResourceNiosSchemaAttributes,
	},
}

var SharedrecordAaaaResourceNiosSchemaAttributes = map[string]schema.Attribute{
	"comment": schema.StringAttribute{
		Optional: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
			customvalidator.ValidateTrimmedString(),
			stringvalidator.LengthBetween(0, 256),
		},
		MarkdownDescription: "Comment for this shared record; maximum 256 characters.",
	},
	"disable": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Determines if this shared record is disabled or not. False means that the record is enabled.",
	},
	"ext_attrs": schema.MapAttribute{
		Optional:    true,
		Computed:    true,
		ElementType: types.StringType,
		Default:     mapdefault.StaticValue(types.MapNull(types.StringType)),
		Validators: []validator.Map{
			mapvalidator.SizeAtLeast(1),
		},
		MarkdownDescription: "Extensible attributes associated with the object. For valid values for extensible attributes, see {extattrs:values}.",
	},
	"ext_attrs_all": schema.MapAttribute{
		Computed:            true,
		ElementType:         types.StringType,
		MarkdownDescription: "All ext_attrs including Terraform Internal ID and inherited attributes.",
		PlanModifiers: []planmodifier.Map{
			importmod.AssociateInternalId(),
		},
	},
	"ipv6addr": schema.StringAttribute{
		Required:   true,
		CustomType: iptypes.IPv6AddressType{},
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
		},
		MarkdownDescription: "The IPv6 Address of the shared record.",
	},
	"name": schema.StringAttribute{
		Required: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
			customvalidator.ValidateTrimmedString(),
		},
		MarkdownDescription: "Name for this shared record. This value can be in unicode format.",
	},
	"shared_record_group": schema.StringAttribute{
		Required: true,
		PlanModifiers: []planmodifier.String{
			immutable.ImmutableString(),
		},
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
		},
		MarkdownDescription: "The name of the shared record group in which the record resides.",
	},
	"ttl": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "The Time To Live (TTL) value for this shared record. A 32-bit unsigned integer that represents the duration, in seconds, for which the shared record is valid (cached). Zero indicates that the shared record should not be cached.",
	},
}

// Expand converts the TF model to the infoblox core model
func (m *SharedrecordAaaaModel) Expand(ctx context.Context, diags *diag.Diagnostics, isCreate bool) *coremodel.SharedrecordAaaa {
	if m == nil {
		return nil
	}

	obj := &coremodel.SharedrecordAaaa{}

	// Expand NIOS nested attribute (returns nil if not present)
	niosModel := flex.ExpandNestedObject[NIOSSharedrecordAaaaModel](ctx, m.NIOS, diags)
	if niosModel != nil {
		obj.NIOS = niosModel.Expand(ctx, diags)
	}

	return obj
}

// Expand converts the NIOS TF model to the core model.
func (m *NIOSSharedrecordAaaaModel) Expand(ctx context.Context, diags *diag.Diagnostics) *coremodel.NIOSSharedrecordAaaaExt {
	return &coremodel.NIOSSharedrecordAaaaExt{
		Comment:           flex.ExpandStringPointerNullAsEmpty(m.Comment),
		Disable:           flex.ExpandBoolPointer(m.Disable),
		ExtAttrs:          flex.ExpandMapStringAny(ctx, m.ExtAttrs, diags),
		Ipv6addr:          flex.ExpandIPv6Address(m.Ipv6addr),
		Name:              flex.ExpandStringPointerNullAsEmpty(m.Name),
		SharedRecordGroup: flex.ExpandStringPointerNullAsEmpty(m.SharedRecordGroup),
		Ttl:               flex.ExpandInt64Pointer(m.Ttl),
	}
}

// ApplySharedrecordAaaaNIOSUseFlags derives NIOS use flags from the raw config
// value(s) and writes them onto the core model. A flag is true when the user
// set any of its governed value fields in config.
func ApplySharedrecordAaaaNIOSUseFlags(ctx context.Context, config tfsdk.Config, obj *coremodel.SharedrecordAaaa, diags *diag.Diagnostics) {
	if obj == nil || obj.NIOS == nil {
		return
	}
	obj.NIOS.UseTtl = flex.DeriveUseFlag(ctx, config, diags, path.Root("nios").AtName("ttl"))
}

// Flatten populates the TF model from a core response.
func (m *SharedrecordAaaaModel) Flatten(ctx context.Context, resp *coremodel.SharedrecordAaaa, diags *diag.Diagnostics) {
	if resp == nil {
		return
	}

	m.Id = flex.FlattenStringPointer(resp.Id)

	// Extract existing NIOS model, flatten API response onto it, convert back
	niosModel := flex.ExpandNestedObject[NIOSSharedrecordAaaaModel](ctx, m.NIOS, diags)
	if niosModel == nil {
		niosModel = &NIOSSharedrecordAaaaModel{}
	}
	niosModel.Flatten(ctx, resp.NIOS, diags)
	if resp.NIOS != nil {
		m.NIOS = flex.FlattenNestedObject(ctx, niosModel, NIOSSharedrecordAaaaAttrTypes, diags)
	} else {
		m.NIOS = types.ObjectNull(NIOSSharedrecordAaaaAttrTypes)
	}

}

// Flatten merges API response onto existing NIOS model.
func (m *NIOSSharedrecordAaaaModel) Flatten(ctx context.Context, from *coremodel.NIOSSharedrecordAaaaExt, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	planExtAttrs := m.ExtAttrs
	if planExtAttrs.IsUnknown() {
		planExtAttrs = types.MapNull(types.StringType)
	}
	m.Comment = flex.FlattenStringPointerEmptyAsNull(from.Comment)
	m.Disable = flex.FlattenBoolPointer(from.Disable)
	m.ExtAttrs, m.ExtAttrsAll = flex.FlattenEAs(planExtAttrs, from.ExtAttrs)
	m.Ipv6addr = flex.FlattenIPv6Address(from.Ipv6addr)
	m.Name = flex.FlattenStringPointerEmptyAsNull(from.Name)
	m.SharedRecordGroup = flex.FlattenStringPointerEmptyAsNull(from.SharedRecordGroup)
	m.Ttl = flex.FlattenInt64Pointer(from.Ttl)
}
