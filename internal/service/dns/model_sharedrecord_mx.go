package dns

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
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
	customvalidator "github.com/infobloxopen/terraform-provider-infoblox/internal/validator"
)

type SharedrecordMxModel struct {
	Id   types.String `tfsdk:"id"`
	NIOS types.Object `tfsdk:"nios"`
}

var SharedrecordMxAttrTypes = map[string]attr.Type{
	"id":   types.StringType,
	"nios": types.ObjectType{AttrTypes: NIOSSharedrecordMxAttrTypes},
}

type NIOSSharedrecordMxModel struct {
	Comment           types.String `tfsdk:"comment"`
	Disable           types.Bool   `tfsdk:"disable"`
	ExtAttrs          types.Map    `tfsdk:"ext_attrs"`
	ExtAttrsAll       types.Map    `tfsdk:"ext_attrs_all"`
	MailExchanger     types.String `tfsdk:"mail_exchanger"`
	Name              types.String `tfsdk:"name"`
	Preference        types.Int64  `tfsdk:"preference"`
	SharedRecordGroup types.String `tfsdk:"shared_record_group"`
	Ttl               types.Int64  `tfsdk:"ttl"`
}

var NIOSSharedrecordMxAttrTypes = map[string]attr.Type{
	"comment":             types.StringType,
	"disable":             types.BoolType,
	"ext_attrs":           types.MapType{ElemType: types.StringType},
	"ext_attrs_all":       types.MapType{ElemType: types.StringType},
	"mail_exchanger":      types.StringType,
	"name":                types.StringType,
	"preference":          types.Int64Type,
	"shared_record_group": types.StringType,
	"ttl":                 types.Int64Type,
}

const (
	SharedrecordMxReturnFields = "comment,disable,dns_mail_exchanger,dns_name,extattrs,mail_exchanger,name,preference,shared_record_group,ttl,use_ttl"
)

var SharedrecordMxResourceSchemaAttributes = map[string]schema.Attribute{
	"id": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The reference to the object.",
	},
	"nios": schema.SingleNestedAttribute{
		Optional:            true,
		MarkdownDescription: "NIOS backend-specific fields.",
		Attributes:          SharedrecordMxResourceNiosSchemaAttributes,
	},
}

var SharedrecordMxResourceNiosSchemaAttributes = map[string]schema.Attribute{
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
		MarkdownDescription: "All ext_attrs including inherited values.",
	},
	"mail_exchanger": schema.StringAttribute{
		Required: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
			customvalidator.IsValidNIOSDomainName(),
		},
		MarkdownDescription: "The name of the mail exchanger in FQDN format. This value can be in unicode format.",
	},
	"name": schema.StringAttribute{
		Required: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
			customvalidator.ValidateTrimmedString(),
		},
		MarkdownDescription: "Name for this shared record. This value can be in unicode format.",
	},
	"preference": schema.Int64Attribute{
		Required: true,
		Validators: []validator.Int64{
			int64validator.Between(0, 65535),
		},
		MarkdownDescription: "The preference value. Valid values are from 0 to 65535 (inclusive), in 32-bit unsigned integer format.",
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
func (m *SharedrecordMxModel) Expand(ctx context.Context, diags *diag.Diagnostics, isCreate bool) *coremodel.SharedrecordMx {
	if m == nil {
		return nil
	}

	obj := &coremodel.SharedrecordMx{}

	// Expand NIOS nested attribute (returns nil if not present)
	niosModel := flex.ExpandNestedObject[NIOSSharedrecordMxModel](ctx, m.NIOS, diags)
	if niosModel != nil {
		obj.NIOS = niosModel.Expand(ctx, diags)
	}

	return obj
}

// Expand converts the NIOS TF model to the core model.
func (m *NIOSSharedrecordMxModel) Expand(ctx context.Context, diags *diag.Diagnostics) *coremodel.NIOSSharedrecordMxExt {
	return &coremodel.NIOSSharedrecordMxExt{
		Comment:           flex.ExpandStringPointerNullAsEmpty(m.Comment),
		Disable:           flex.ExpandBoolPointer(m.Disable),
		ExtAttrs:          flex.ExpandMapStringAny(ctx, m.ExtAttrs, diags),
		MailExchanger:     flex.ExpandStringPointerNullAsEmpty(m.MailExchanger),
		Name:              flex.ExpandStringPointerNullAsEmpty(m.Name),
		Preference:        flex.ExpandInt64Pointer(m.Preference),
		SharedRecordGroup: flex.ExpandStringPointerNullAsEmpty(m.SharedRecordGroup),
		Ttl:               flex.ExpandInt64Pointer(m.Ttl),
	}
}

// ApplySharedrecordMxNIOSUseFlags derives NIOS use flags from the raw config
// value(s) and writes them onto the core model. A flag is true when the user
// set any of its governed value fields in config.
func ApplySharedrecordMxNIOSUseFlags(ctx context.Context, config tfsdk.Config, obj *coremodel.SharedrecordMx, diags *diag.Diagnostics) {
	if obj == nil || obj.NIOS == nil {
		return
	}
	obj.NIOS.UseTtl = flex.DeriveUseFlag(ctx, config, diags, path.Root("nios").AtName("ttl"))
}

// Flatten populates the TF model from a core response.
func (m *SharedrecordMxModel) Flatten(ctx context.Context, resp *coremodel.SharedrecordMx, diags *diag.Diagnostics) {
	if resp == nil {
		return
	}

	m.Id = flex.FlattenStringPointer(resp.Id)

	// Extract existing NIOS model, flatten API response onto it, convert back
	niosModel := flex.ExpandNestedObject[NIOSSharedrecordMxModel](ctx, m.NIOS, diags)
	if niosModel == nil {
		niosModel = &NIOSSharedrecordMxModel{}
	}
	niosModel.Flatten(ctx, resp.NIOS, diags)
	if resp.NIOS != nil {
		m.NIOS = flex.FlattenNestedObject(ctx, niosModel, NIOSSharedrecordMxAttrTypes, diags)
	} else {
		m.NIOS = types.ObjectNull(NIOSSharedrecordMxAttrTypes)
	}

}

// Flatten merges API response onto existing NIOS model.
func (m *NIOSSharedrecordMxModel) Flatten(ctx context.Context, from *coremodel.NIOSSharedrecordMxExt, diags *diag.Diagnostics) {
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
	m.MailExchanger = flex.FlattenStringPointerEmptyAsNull(from.MailExchanger)
	m.Name = flex.FlattenStringPointerEmptyAsNull(from.Name)
	m.Preference = flex.FlattenInt64Pointer(from.Preference)
	m.SharedRecordGroup = flex.FlattenStringPointerEmptyAsNull(from.SharedRecordGroup)
	m.Ttl = flex.FlattenInt64Pointer(from.Ttl)
}
