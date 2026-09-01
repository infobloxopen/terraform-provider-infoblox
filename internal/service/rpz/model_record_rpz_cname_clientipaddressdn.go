package rpz

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

	coremodel "github.com/infobloxopen/terraform-provider-infoblox/internal/core/model/rpz"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
	immutable "github.com/infobloxopen/terraform-provider-infoblox/internal/planmodifiers/immutable"
	importmod "github.com/infobloxopen/terraform-provider-infoblox/internal/planmodifiers/import"
	internaltypes "github.com/infobloxopen/terraform-provider-infoblox/internal/types"
	customvalidator "github.com/infobloxopen/terraform-provider-infoblox/internal/validator"
)

type RecordRpzCnameClientipaddressdnModel struct {
	Id   types.String `tfsdk:"id"`
	NIOS types.Object `tfsdk:"nios"`
}

var RecordRpzCnameClientipaddressdnAttrTypes = map[string]attr.Type{
	"id":   types.StringType,
	"nios": types.ObjectType{AttrTypes: NIOSRecordRpzCnameClientipaddressdnAttrTypes},
}

type NIOSRecordRpzCnameClientipaddressdnModel struct {
	Canonical   types.String         `tfsdk:"canonical"`
	Comment     types.String         `tfsdk:"comment"`
	Disable     types.Bool           `tfsdk:"disable"`
	ExtAttrs    types.Map            `tfsdk:"ext_attrs"`
	ExtAttrsAll types.Map            `tfsdk:"ext_attrs_all"`
	Name        internaltypes.IPName `tfsdk:"name"`
	RpZone      types.String         `tfsdk:"rp_zone"`
	Ttl         types.Int64          `tfsdk:"ttl"`
	View        types.String         `tfsdk:"view"`
}

var NIOSRecordRpzCnameClientipaddressdnAttrTypes = map[string]attr.Type{
	"canonical":     types.StringType,
	"comment":       types.StringType,
	"disable":       types.BoolType,
	"ext_attrs":     types.MapType{ElemType: types.StringType},
	"ext_attrs_all": types.MapType{ElemType: types.StringType},
	"name":          internaltypes.IPNameType{},
	"rp_zone":       types.StringType,
	"ttl":           types.Int64Type,
	"view":          types.StringType,
}

const (
	RecordRpzCnameClientipaddressdnReturnFields = "canonical,comment,disable,extattrs,is_ipv4,name,rp_zone,ttl,use_ttl,view,zone"
)

var RecordRpzCnameClientipaddressdnResourceSchemaAttributes = map[string]schema.Attribute{
	"id": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The reference to the object.",
	},
	"nios": schema.SingleNestedAttribute{
		Optional:            true,
		MarkdownDescription: "NIOS backend-specific fields.",
		Attributes:          RecordRpzCnameClientipaddressdnResourceNiosSchemaAttributes,
	},
}

var RecordRpzCnameClientipaddressdnResourceNiosSchemaAttributes = map[string]schema.Attribute{
	"canonical": schema.StringAttribute{
		Required: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
			customvalidator.IsValidNIOSDomainName(),
		},
		MarkdownDescription: "The canonical name in FQDN format. This value can be in unicode format.",
	},
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
	"name": schema.StringAttribute{
		Required:   true,
		CustomType: internaltypes.IPNameType{},
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
			customvalidator.IsValidNIOSDomainName(),
			customvalidator.ValidateTrimmedString(),
		},
		MarkdownDescription: "The name for a record in FQDN format. This value cannot be in unicode format.",
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
func (m *RecordRpzCnameClientipaddressdnModel) Expand(ctx context.Context, diags *diag.Diagnostics, isCreate bool) *coremodel.RecordRpzCnameClientipaddressdn {
	if m == nil {
		return nil
	}

	obj := &coremodel.RecordRpzCnameClientipaddressdn{}

	// Expand NIOS nested attribute (returns nil if not present)
	niosModel := flex.ExpandNestedObject[NIOSRecordRpzCnameClientipaddressdnModel](ctx, m.NIOS, diags)
	if niosModel != nil {
		obj.NIOS = niosModel.Expand(ctx, diags)
	}

	return obj
}

// Expand converts the NIOS TF model to the core model.
func (m *NIOSRecordRpzCnameClientipaddressdnModel) Expand(ctx context.Context, diags *diag.Diagnostics) *coremodel.NIOSRecordRpzCnameClientipaddressdnExt {
	return &coremodel.NIOSRecordRpzCnameClientipaddressdnExt{
		Canonical: flex.ExpandStringPointerNullAsEmpty(m.Canonical),
		Comment:   flex.ExpandStringPointerNullAsEmpty(m.Comment),
		Disable:   flex.ExpandBoolPointer(m.Disable),
		ExtAttrs:  flex.ExpandMapStringAny(ctx, m.ExtAttrs, diags),
		Name:      flex.ExpandStringPointer(m.Name.StringValue),
		RpZone:    flex.ExpandStringPointerNullAsEmpty(m.RpZone),
		Ttl:       flex.ExpandInt64Pointer(m.Ttl),
		View:      flex.ExpandStringPointerNullAsEmpty(m.View),
	}
}

// ApplyRecordRpzCnameClientipaddressdnNIOSUseFlags derives NIOS use flags from the raw config
// value(s) and writes them onto the core model. A flag is true when the user
// set any of its governed value fields in config.
func ApplyRecordRpzCnameClientipaddressdnNIOSUseFlags(ctx context.Context, config tfsdk.Config, obj *coremodel.RecordRpzCnameClientipaddressdn, diags *diag.Diagnostics) {
	if obj == nil || obj.NIOS == nil {
		return
	}
	obj.NIOS.UseTtl = flex.DeriveUseFlag(ctx, config, diags, path.Root("nios").AtName("ttl"))
}

// Flatten populates the TF model from a core response.
func (m *RecordRpzCnameClientipaddressdnModel) Flatten(ctx context.Context, resp *coremodel.RecordRpzCnameClientipaddressdn, diags *diag.Diagnostics) {
	if resp == nil {
		return
	}

	m.Id = flex.FlattenStringPointer(resp.Id)

	// Extract existing NIOS model, flatten API response onto it, convert back
	niosModel := flex.ExpandNestedObject[NIOSRecordRpzCnameClientipaddressdnModel](ctx, m.NIOS, diags)
	if niosModel == nil {
		niosModel = &NIOSRecordRpzCnameClientipaddressdnModel{}
	}
	niosModel.Flatten(ctx, resp.NIOS, diags)
	if resp.NIOS != nil {
		m.NIOS = flex.FlattenNestedObject(ctx, niosModel, NIOSRecordRpzCnameClientipaddressdnAttrTypes, diags)
	} else {
		m.NIOS = types.ObjectNull(NIOSRecordRpzCnameClientipaddressdnAttrTypes)
	}

}

// Flatten merges API response onto existing NIOS model.
func (m *NIOSRecordRpzCnameClientipaddressdnModel) Flatten(ctx context.Context, from *coremodel.NIOSRecordRpzCnameClientipaddressdnExt, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	planExtAttrs := m.ExtAttrs
	if planExtAttrs.IsUnknown() {
		planExtAttrs = types.MapNull(types.StringType)
	}
	m.Canonical = flex.FlattenStringPointerEmptyAsNull(from.Canonical)
	m.Comment = flex.FlattenStringPointerEmptyAsNull(from.Comment)
	m.Disable = flex.FlattenBoolPointer(from.Disable)
	m.ExtAttrs, m.ExtAttrsAll = flex.FlattenEAs(planExtAttrs, from.ExtAttrs)
	m.Name.StringValue = flex.FlattenStringPointer(from.Name)
	m.RpZone = flex.FlattenStringPointerEmptyAsNull(from.RpZone)
	m.Ttl = flex.FlattenInt64Pointer(from.Ttl)
	m.View = flex.FlattenStringPointerEmptyAsNull(from.View)
}
