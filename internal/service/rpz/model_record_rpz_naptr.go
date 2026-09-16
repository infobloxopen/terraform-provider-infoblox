package rpz

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
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"

	coremodel "github.com/infobloxopen/terraform-provider-infoblox/internal/core/model/rpz"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
	immutable "github.com/infobloxopen/terraform-provider-infoblox/internal/planmodifiers/immutable"
	importmod "github.com/infobloxopen/terraform-provider-infoblox/internal/planmodifiers/import"
	internaltypes "github.com/infobloxopen/terraform-provider-infoblox/internal/types"
	customvalidator "github.com/infobloxopen/terraform-provider-infoblox/internal/validator"
)

type RecordRpzNaptrModel struct {
	Id   types.String `tfsdk:"id"`
	NIOS types.Object `tfsdk:"nios"`
}

var RecordRpzNaptrAttrTypes = map[string]attr.Type{
	"id":   types.StringType,
	"nios": types.ObjectType{AttrTypes: NIOSRecordRpzNaptrAttrTypes},
}

type NIOSRecordRpzNaptrModel struct {
	Comment     types.String                        `tfsdk:"comment"`
	Disable     types.Bool                          `tfsdk:"disable"`
	ExtAttrs    types.Map                           `tfsdk:"ext_attrs"`
	ExtAttrsAll types.Map                           `tfsdk:"ext_attrs_all"`
	Flags       types.String                        `tfsdk:"flags"`
	Name        internaltypes.CaseInsensitiveString `tfsdk:"name"`
	Order       types.Int64                         `tfsdk:"order"`
	Preference  types.Int64                         `tfsdk:"preference"`
	Regexp      types.String                        `tfsdk:"regexp"`
	Replacement types.String                        `tfsdk:"replacement"`
	RpZone      types.String                        `tfsdk:"rp_zone"`
	Services    types.String                        `tfsdk:"services"`
	Ttl         types.Int64                         `tfsdk:"ttl"`
	View        types.String                        `tfsdk:"view"`
}

var NIOSRecordRpzNaptrAttrTypes = map[string]attr.Type{
	"comment":       types.StringType,
	"disable":       types.BoolType,
	"ext_attrs":     types.MapType{ElemType: types.StringType},
	"ext_attrs_all": types.MapType{ElemType: types.StringType},
	"flags":         types.StringType,
	"name":          internaltypes.CaseInsensitiveStringType{},
	"order":         types.Int64Type,
	"preference":    types.Int64Type,
	"regexp":        types.StringType,
	"replacement":   types.StringType,
	"rp_zone":       types.StringType,
	"services":      types.StringType,
	"ttl":           types.Int64Type,
	"view":          types.StringType,
}

const (
	RecordRpzNaptrReturnFields = "comment,disable,extattrs,flags,last_queried,name,order,preference,regexp,replacement,rp_zone,services,ttl,use_ttl,view,zone"
)

var RecordRpzNaptrResourceSchemaAttributes = map[string]schema.Attribute{
	"id": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The reference to the object.",
	},
	"nios": schema.SingleNestedAttribute{
		Optional:            true,
		MarkdownDescription: "NIOS backend-specific fields.",
		Attributes:          RecordRpzNaptrResourceNiosSchemaAttributes,
	},
}

var RecordRpzNaptrResourceNiosSchemaAttributes = map[string]schema.Attribute{
	"comment": schema.StringAttribute{
		Optional: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
			stringvalidator.LengthBetween(0, 256),
			customvalidator.ValidateTrimmedString(),
		},
		MarkdownDescription: "The comment for the record; maximum 256 characters.",
	},
	"disable": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Determines if the record is disabled or not. False means that the record is enabled.",
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
	"flags": schema.StringAttribute{
		Default: stringdefault.StaticString(""),
		Validators: []validator.String{
			customvalidator.ValidateTrimmedString(),
			stringvalidator.OneOf("U", "S", "P", "A", ""),
		},
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The flags used to control the interpretation of the fields for a Substitute (NAPTR Record) Rule object. Supported values for the flags field are \"U\", \"S\", \"P\" and \"A\".",
	},
	"name": schema.StringAttribute{
		Required:   true,
		CustomType: internaltypes.CaseInsensitiveStringType{},
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
			customvalidator.IsValidNIOSDomainName(),
		},
		MarkdownDescription: "The name for a record in FQDN format. This value cannot be in unicode format.",
	},
	"order": schema.Int64Attribute{
		Required: true,
		Validators: []validator.Int64{
			int64validator.Between(0, 65535),
		},
		MarkdownDescription: "The order parameter of the Substitute (NAPTR Record) Rule records. This parameter specifies the order in which the NAPTR rules are applied when multiple rules are present. Valid values are from 0 to 65535 (inclusive), in 32-bit unsigned integer format.",
	},
	"preference": schema.Int64Attribute{
		Required: true,
		Validators: []validator.Int64{
			int64validator.Between(0, 65535),
		},
		MarkdownDescription: "The preference of the Substitute (NAPTR Record) Rule record. The preference field determines the order NAPTR records are processed when multiple records with the same order parameter are present. Valid values are from 0 to 65535 (inclusive), in 32-bit unsigned integer format.",
	},
	"regexp": schema.StringAttribute{
		Default:  stringdefault.StaticString(""),
		Optional: true,
		Computed: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
			customvalidator.ValidateTrimmedString(),
		},
		MarkdownDescription: "The regular expression-based rewriting rule of the Substitute (NAPTR Record) Rule record. This should be a POSIX compliant regular expression, including the substitution rule and flags. Refer to RFC 2915 for the field syntax details.",
	},
	"replacement": schema.StringAttribute{
		Required: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
			customvalidator.IsValidNIOSDomainName(),
		},
		MarkdownDescription: "The replacement field of the Substitute (NAPTR Record) Rule object. For nonterminal NAPTR records, this field specifies the next domain name to look up. This value can be in unicode format.",
	},
	"rp_zone": schema.StringAttribute{
		Required: true,
		PlanModifiers: []planmodifier.String{
			immutable.ImmutableString(),
		},
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
		},
		MarkdownDescription: "The name of a response policy zone in which the record resides.",
	},
	"services": schema.StringAttribute{
		Default:  stringdefault.StaticString(""),
		Optional: true,
		Computed: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
			customvalidator.ValidateTrimmedString(),
			stringvalidator.LengthBetween(0, 128),
		},
		MarkdownDescription: "The services field of the Substitute (NAPTR Record) Rule object; maximum 128 characters. The services field contains protocol and service identifiers, such as \"http+E2U\" or \"SIPS+D2T\".",
	},
	"ttl": schema.Int64Attribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The Time To Live (TTL) value for record. A 32-bit unsigned integer that represents the duration, in seconds, for which the record is valid (cached). Zero indicates that the record should not be cached.",
	},
	"view": schema.StringAttribute{
		Default:  stringdefault.StaticString("default"),
		Optional: true,
		Computed: true,
		PlanModifiers: []planmodifier.String{
			immutable.ImmutableString(),
		},
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
			customvalidator.ValidateTrimmedString(),
		},
		MarkdownDescription: "The name of the DNS View in which the record resides. Example: \"external\".",
	},
}

// Expand converts the TF model to the infoblox core model
func (m *RecordRpzNaptrModel) Expand(ctx context.Context, diags *diag.Diagnostics, isCreate bool) *coremodel.RecordRpzNaptr {
	if m == nil {
		return nil
	}

	obj := &coremodel.RecordRpzNaptr{}

	// Expand NIOS nested attribute (returns nil if not present)
	niosModel := flex.ExpandNestedObject[NIOSRecordRpzNaptrModel](ctx, m.NIOS, diags)
	if niosModel != nil {
		obj.NIOS = niosModel.Expand(ctx, diags)
	}

	return obj
}

// Expand converts the NIOS TF model to the core model.
func (m *NIOSRecordRpzNaptrModel) Expand(ctx context.Context, diags *diag.Diagnostics) *coremodel.NIOSRecordRpzNaptrExt {
	return &coremodel.NIOSRecordRpzNaptrExt{
		Comment:     flex.ExpandStringPointerNullAsEmpty(m.Comment),
		Disable:     flex.ExpandBoolPointer(m.Disable),
		ExtAttrs:    flex.ExpandMapStringAny(ctx, m.ExtAttrs, diags),
		Flags:       flex.ExpandStringPointerNullAsEmpty(m.Flags),
		Name:        flex.ExpandStringPointer(m.Name.StringValue),
		Order:       flex.ExpandInt64Pointer(m.Order),
		Preference:  flex.ExpandInt64Pointer(m.Preference),
		Regexp:      flex.ExpandStringPointerNullAsEmpty(m.Regexp),
		Replacement: flex.ExpandStringPointerNullAsEmpty(m.Replacement),
		RpZone:      flex.ExpandStringPointerNullAsEmpty(m.RpZone),
		Services:    flex.ExpandStringPointerNullAsEmpty(m.Services),
		Ttl:         flex.ExpandInt64Pointer(m.Ttl),
		View:        flex.ExpandStringPointerNullAsEmpty(m.View),
	}
}

// ApplyRecordRpzNaptrNIOSUseFlags derives NIOS use flags from the raw config
// value(s) and writes them onto the core model. A flag is true when the user
// set any of its governed value fields in config.
func ApplyRecordRpzNaptrNIOSUseFlags(ctx context.Context, config tfsdk.Config, obj *coremodel.RecordRpzNaptr, diags *diag.Diagnostics) {
	if obj == nil || obj.NIOS == nil {
		return
	}
	obj.NIOS.UseTtl = flex.DeriveUseFlag(ctx, config, diags, path.Root("nios").AtName("ttl"))
}

// Flatten populates the TF model from a core response.
func (m *RecordRpzNaptrModel) Flatten(ctx context.Context, resp *coremodel.RecordRpzNaptr, diags *diag.Diagnostics) {
	if resp == nil {
		return
	}

	m.Id = flex.FlattenStringPointer(resp.Id)

	// Extract existing NIOS model, flatten API response onto it, convert back
	niosModel := flex.ExpandNestedObject[NIOSRecordRpzNaptrModel](ctx, m.NIOS, diags)
	if niosModel == nil {
		niosModel = &NIOSRecordRpzNaptrModel{}
	}
	niosModel.Flatten(ctx, resp.NIOS, diags)
	if resp.NIOS != nil {
		m.NIOS = flex.FlattenNestedObject(ctx, niosModel, NIOSRecordRpzNaptrAttrTypes, diags)
	} else {
		m.NIOS = types.ObjectNull(NIOSRecordRpzNaptrAttrTypes)
	}

}

// Flatten merges API response onto existing NIOS model.
func (m *NIOSRecordRpzNaptrModel) Flatten(ctx context.Context, from *coremodel.NIOSRecordRpzNaptrExt, diags *diag.Diagnostics) {
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
	m.Flags = flex.FlattenStringPointer(from.Flags)
	m.Name.StringValue = flex.FlattenStringPointer(from.Name)
	m.Order = flex.FlattenInt64Pointer(from.Order)
	m.Preference = flex.FlattenInt64Pointer(from.Preference)
	m.Regexp = flex.FlattenStringPointer(from.Regexp)
	m.Replacement = flex.FlattenStringPointerEmptyAsNull(from.Replacement)
	m.RpZone = flex.FlattenStringPointerEmptyAsNull(from.RpZone)
	m.Services = flex.FlattenStringPointer(from.Services)
	m.Ttl = flex.FlattenInt64Pointer(from.Ttl)
	m.View = flex.FlattenStringPointerEmptyAsNull(from.View)
}
