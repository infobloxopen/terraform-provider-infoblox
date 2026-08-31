package dns

import (
	"context"

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

	coremodel "github.com/infobloxopen/terraform-provider-infoblox/internal/core/model/dns"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
	importmod "github.com/infobloxopen/terraform-provider-infoblox/internal/planmodifiers/import"
	customvalidator "github.com/infobloxopen/terraform-provider-infoblox/internal/validator"
)

type RecordAliasModel struct {
	Id   types.String `tfsdk:"id"`
	NIOS types.Object `tfsdk:"nios"`
}

var RecordAliasAttrTypes = map[string]attr.Type{
	"id":   types.StringType,
	"nios": types.ObjectType{AttrTypes: NIOSRecordAliasAttrTypes},
}

type NIOSRecordAliasModel struct {
	Comment     types.String `tfsdk:"comment"`
	Creator     types.String `tfsdk:"creator"`
	Disable     types.Bool   `tfsdk:"disable"`
	ExtAttrs    types.Map    `tfsdk:"ext_attrs"`
	ExtAttrsAll types.Map    `tfsdk:"ext_attrs_all"`
	Name        types.String `tfsdk:"name"`
	TargetName  types.String `tfsdk:"target_name"`
	TargetType  types.String `tfsdk:"target_type"`
	Ttl         types.Int64  `tfsdk:"ttl"`
	View        types.String `tfsdk:"view"`
}

var NIOSRecordAliasAttrTypes = map[string]attr.Type{
	"comment":       types.StringType,
	"creator":       types.StringType,
	"disable":       types.BoolType,
	"ext_attrs":     types.MapType{ElemType: types.StringType},
	"ext_attrs_all": types.MapType{ElemType: types.StringType},
	"name":          types.StringType,
	"target_name":   types.StringType,
	"target_type":   types.StringType,
	"ttl":           types.Int64Type,
	"view":          types.StringType,
}

const (
	RecordAliasReturnFields = "aws_rte53_record_info,cloud_info,comment,creator,disable,dns_name,dns_target_name,extattrs,last_queried,name,target_name,target_type,ttl,use_ttl,view,zone"
)

var RecordAliasResourceSchemaAttributes = map[string]schema.Attribute{
	"id": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The reference to the object.",
	},
	"nios": schema.SingleNestedAttribute{
		Optional:            true,
		MarkdownDescription: "NIOS backend-specific fields.",
		Attributes:          RecordAliasResourceNiosSchemaAttributes,
	},
}

var RecordAliasResourceNiosSchemaAttributes = map[string]schema.Attribute{
	"comment": schema.StringAttribute{
		Optional: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
			customvalidator.ValidateTrimmedString(),
			stringvalidator.LengthBetween(0, 256),
		},
		MarkdownDescription: "Comment for the record; maximum 256 characters.",
	},
	"creator": schema.StringAttribute{
		Default: stringdefault.StaticString("STATIC"),
		Validators: []validator.String{
			stringvalidator.OneOf("STATIC"),
		},
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The record creator.",
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
	"name": schema.StringAttribute{
		Required: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
			customvalidator.IsValidNIOSDomainName(),
		},
		MarkdownDescription: "The name for an Alias record in FQDN format. This value can be in unicode format. Regular expression search is not supported for unicode values.",
	},
	"target_name": schema.StringAttribute{
		Required: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
			customvalidator.NotEqualsField(path.MatchRoot("nios").AtName("name")),
		},
		MarkdownDescription: "Target name in FQDN format. This value can be in unicode format.",
	},
	"target_type": schema.StringAttribute{
		Validators: []validator.String{
			stringvalidator.OneOf("A", "AAAA", "MX", "NAPTR", "PTR", "SPF", "SRV", "TXT"),
		},
		Required:            true,
		MarkdownDescription: "Target type.",
	},
	"ttl": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "The Time To Live (TTL) value for record. A 32-bit unsigned integer that represents the duration, in seconds, for which the record is valid (cached). Zero indicates that the record should not be cached.",
	},
	"view": schema.StringAttribute{
		Default:  stringdefault.StaticString("default"),
		Optional: true,
		Computed: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
			customvalidator.ValidateTrimmedString(),
		},
		MarkdownDescription: "The name of the DNS View in which the record resides. Example: \"external\".",
	},
}

// Expand converts the TF model to the infoblox core model
func (m *RecordAliasModel) Expand(ctx context.Context, diags *diag.Diagnostics, isCreate bool) *coremodel.RecordAlias {
	if m == nil {
		return nil
	}

	obj := &coremodel.RecordAlias{}

	// Expand NIOS nested attribute (returns nil if not present)
	niosModel := flex.ExpandNestedObject[NIOSRecordAliasModel](ctx, m.NIOS, diags)
	if niosModel != nil {
		obj.NIOS = niosModel.Expand(ctx, diags)
	}

	return obj
}

// Expand converts the NIOS TF model to the core model.
func (m *NIOSRecordAliasModel) Expand(ctx context.Context, diags *diag.Diagnostics) *coremodel.NIOSRecordAliasExt {
	return &coremodel.NIOSRecordAliasExt{
		Comment:    flex.ExpandStringPointerNullAsEmpty(m.Comment),
		Creator:    flex.ExpandStringPointerNullAsEmpty(m.Creator),
		Disable:    flex.ExpandBoolPointer(m.Disable),
		ExtAttrs:   flex.ExpandMapStringAny(ctx, m.ExtAttrs, diags),
		Name:       flex.ExpandStringPointerNullAsEmpty(m.Name),
		TargetName: flex.ExpandStringPointerNullAsEmpty(m.TargetName),
		TargetType: flex.ExpandStringPointerNullAsEmpty(m.TargetType),
		Ttl:        flex.ExpandInt64Pointer(m.Ttl),
		View:       flex.ExpandStringPointerNullAsEmpty(m.View),
	}
}

// ApplyRecordAliasNIOSUseFlags derives NIOS use flags from the raw config
// value(s) and writes them onto the core model. A flag is true when the user
// set any of its governed value fields in config.
func ApplyRecordAliasNIOSUseFlags(ctx context.Context, config tfsdk.Config, obj *coremodel.RecordAlias, diags *diag.Diagnostics) {
	if obj == nil || obj.NIOS == nil {
		return
	}
	obj.NIOS.UseTtl = flex.DeriveUseFlag(ctx, config, diags, path.Root("nios").AtName("ttl"))
}

// Flatten populates the TF model from a core response.
func (m *RecordAliasModel) Flatten(ctx context.Context, resp *coremodel.RecordAlias, diags *diag.Diagnostics) {
	if resp == nil {
		return
	}

	m.Id = flex.FlattenStringPointer(resp.Id)

	// Extract existing NIOS model, flatten API response onto it, convert back
	niosModel := flex.ExpandNestedObject[NIOSRecordAliasModel](ctx, m.NIOS, diags)
	if niosModel == nil {
		niosModel = &NIOSRecordAliasModel{}
	}
	niosModel.Flatten(ctx, resp.NIOS, diags)
	if resp.NIOS != nil {
		m.NIOS = flex.FlattenNestedObject(ctx, niosModel, NIOSRecordAliasAttrTypes, diags)
	} else {
		m.NIOS = types.ObjectNull(NIOSRecordAliasAttrTypes)
	}

}

// Flatten merges API response onto existing NIOS model.
func (m *NIOSRecordAliasModel) Flatten(ctx context.Context, from *coremodel.NIOSRecordAliasExt, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	planExtAttrs := m.ExtAttrs
	if planExtAttrs.IsUnknown() {
		planExtAttrs = types.MapNull(types.StringType)
	}
	m.Comment = flex.FlattenStringPointerEmptyAsNull(from.Comment)
	m.Creator = flex.FlattenStringPointerEmptyAsNull(from.Creator)
	m.Disable = flex.FlattenBoolPointer(from.Disable)
	m.ExtAttrs, m.ExtAttrsAll = flex.FlattenEAs(planExtAttrs, from.ExtAttrs)
	m.Name = flex.FlattenStringPointerEmptyAsNull(from.Name)
	m.TargetName = flex.FlattenStringPointerEmptyAsNull(from.TargetName)
	m.TargetType = flex.FlattenStringPointerEmptyAsNull(from.TargetType)
	m.Ttl = flex.FlattenInt64Pointer(from.Ttl)
	m.View = flex.FlattenStringPointerEmptyAsNull(from.View)
}
