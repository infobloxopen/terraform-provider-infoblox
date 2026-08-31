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

type RecordRpzTxtModel struct {
	Id   types.String `tfsdk:"id"`
	NIOS types.Object `tfsdk:"nios"`
}

var RecordRpzTxtAttrTypes = map[string]attr.Type{
	"id":   types.StringType,
	"nios": types.ObjectType{AttrTypes: NIOSRecordRpzTxtAttrTypes},
}

type NIOSRecordRpzTxtModel struct {
	Comment     types.String                        `tfsdk:"comment"`
	Disable     types.Bool                          `tfsdk:"disable"`
	ExtAttrs    types.Map                           `tfsdk:"ext_attrs"`
	ExtAttrsAll types.Map                           `tfsdk:"ext_attrs_all"`
	Name        internaltypes.CaseInsensitiveString `tfsdk:"name"`
	RpZone      types.String                        `tfsdk:"rp_zone"`
	Text        types.String                        `tfsdk:"text"`
	Ttl         types.Int64                         `tfsdk:"ttl"`
	View        types.String                        `tfsdk:"view"`
}

var NIOSRecordRpzTxtAttrTypes = map[string]attr.Type{
	"comment":       types.StringType,
	"disable":       types.BoolType,
	"ext_attrs":     types.MapType{ElemType: types.StringType},
	"ext_attrs_all": types.MapType{ElemType: types.StringType},
	"name":          internaltypes.CaseInsensitiveStringType{},
	"rp_zone":       types.StringType,
	"text":          types.StringType,
	"ttl":           types.Int64Type,
	"view":          types.StringType,
}

const (
	RecordRpzTxtReturnFields = "comment,disable,extattrs,name,rp_zone,text,ttl,use_ttl,view,zone"
)

var RecordRpzTxtResourceSchemaAttributes = map[string]schema.Attribute{
	"id": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The reference to the object.",
	},
	"nios": schema.SingleNestedAttribute{
		Optional:            true,
		MarkdownDescription: "NIOS backend-specific fields.",
		Attributes:          RecordRpzTxtResourceNiosSchemaAttributes,
	},
}

var RecordRpzTxtResourceNiosSchemaAttributes = map[string]schema.Attribute{
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
		CustomType: internaltypes.CaseInsensitiveStringType{},
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
	"text": schema.StringAttribute{
		Required: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
			customvalidator.ValidateTXT(true),
			customvalidator.ValidateTrimmedString(),
		},
		MarkdownDescription: "Text associated with the record. It can contain up to 255 bytes per substring, up to a total of 512 bytes. To enter leading, trailing, or embedded spaces in the text, add quotes around the text to preserve the spaces.",
	},
	"ttl": schema.Int64Attribute{
		Optional: true,
		Computed: true,
		Validators: []validator.Int64{
			int64validator.Between(0, 4294967295),
		},
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
func (m *RecordRpzTxtModel) Expand(ctx context.Context, diags *diag.Diagnostics, isCreate bool) *coremodel.RecordRpzTxt {
	if m == nil {
		return nil
	}

	obj := &coremodel.RecordRpzTxt{}

	// Expand NIOS nested attribute (returns nil if not present)
	niosModel := flex.ExpandNestedObject[NIOSRecordRpzTxtModel](ctx, m.NIOS, diags)
	if niosModel != nil {
		obj.NIOS = niosModel.Expand(ctx, diags)
	}

	return obj
}

// Expand converts the NIOS TF model to the core model.
func (m *NIOSRecordRpzTxtModel) Expand(ctx context.Context, diags *diag.Diagnostics) *coremodel.NIOSRecordRpzTxtExt {
	return &coremodel.NIOSRecordRpzTxtExt{
		Comment:  flex.ExpandStringPointerNullAsEmpty(m.Comment),
		Disable:  flex.ExpandBoolPointer(m.Disable),
		ExtAttrs: flex.ExpandMapStringAny(ctx, m.ExtAttrs, diags),
		Name:     flex.ExpandStringPointer(m.Name.StringValue),
		RpZone:   flex.ExpandStringPointerNullAsEmpty(m.RpZone),
		Text:     flex.ExpandStringPointerNullAsEmpty(m.Text),
		Ttl:      flex.ExpandInt64Pointer(m.Ttl),
		View:     flex.ExpandStringPointerNullAsEmpty(m.View),
	}
}

// ApplyRecordRpzTxtNIOSUseFlags derives NIOS use flags from the raw config
// value(s) and writes them onto the core model. A flag is true when the user
// set any of its governed value fields in config.
func ApplyRecordRpzTxtNIOSUseFlags(ctx context.Context, config tfsdk.Config, obj *coremodel.RecordRpzTxt, diags *diag.Diagnostics) {
	if obj == nil || obj.NIOS == nil {
		return
	}
	obj.NIOS.UseTtl = flex.DeriveUseFlag(ctx, config, diags, path.Root("nios").AtName("ttl"))
}

// Flatten populates the TF model from a core response.
func (m *RecordRpzTxtModel) Flatten(ctx context.Context, resp *coremodel.RecordRpzTxt, diags *diag.Diagnostics) {
	if resp == nil {
		return
	}

	m.Id = flex.FlattenStringPointer(resp.Id)

	// Extract existing NIOS model, flatten API response onto it, convert back
	niosModel := flex.ExpandNestedObject[NIOSRecordRpzTxtModel](ctx, m.NIOS, diags)
	if niosModel == nil {
		niosModel = &NIOSRecordRpzTxtModel{}
	}
	niosModel.Flatten(ctx, resp.NIOS, diags)
	if resp.NIOS != nil {
		m.NIOS = flex.FlattenNestedObject(ctx, niosModel, NIOSRecordRpzTxtAttrTypes, diags)
	} else {
		m.NIOS = types.ObjectNull(NIOSRecordRpzTxtAttrTypes)
	}

}

// Flatten merges API response onto existing NIOS model.
func (m *NIOSRecordRpzTxtModel) Flatten(ctx context.Context, from *coremodel.NIOSRecordRpzTxtExt, diags *diag.Diagnostics) {
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
	m.Name.StringValue = flex.FlattenStringPointer(from.Name)
	m.RpZone = flex.FlattenStringPointerEmptyAsNull(from.RpZone)
	m.Text = flex.FlattenStringPointerEmptyAsNull(from.Text)
	m.Ttl = flex.FlattenInt64Pointer(from.Ttl)
	m.View = flex.FlattenStringPointerEmptyAsNull(from.View)
}
