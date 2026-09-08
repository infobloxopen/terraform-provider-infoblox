package rpz

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

type RecordRpzAaaaIpaddressModel struct {
	Id   types.String `tfsdk:"id"`
	NIOS types.Object `tfsdk:"nios"`
}

var RecordRpzAaaaIpaddressAttrTypes = map[string]attr.Type{
	"id":   types.StringType,
	"nios": types.ObjectType{AttrTypes: NIOSRecordRpzAaaaIpaddressAttrTypes},
}

type NIOSRecordRpzAaaaIpaddressModel struct {
	Comment     types.String           `tfsdk:"comment"`
	Disable     types.Bool             `tfsdk:"disable"`
	ExtAttrs    types.Map              `tfsdk:"ext_attrs"`
	ExtAttrsAll types.Map              `tfsdk:"ext_attrs_all"`
	Ipv6addr    iptypes.IPv6Address    `tfsdk:"ipv6addr"`
	Name        internaltypes.IPv6Name `tfsdk:"name"`
	RpZone      types.String           `tfsdk:"rp_zone"`
	Ttl         types.Int64            `tfsdk:"ttl"`
	View        types.String           `tfsdk:"view"`
}

var NIOSRecordRpzAaaaIpaddressAttrTypes = map[string]attr.Type{
	"comment":       types.StringType,
	"disable":       types.BoolType,
	"ext_attrs":     types.MapType{ElemType: types.StringType},
	"ext_attrs_all": types.MapType{ElemType: types.StringType},
	"ipv6addr":      iptypes.IPv6AddressType{},
	"name":          internaltypes.IPv6NameType{},
	"rp_zone":       types.StringType,
	"ttl":           types.Int64Type,
	"view":          types.StringType,
}

const (
	RecordRpzAaaaIpaddressReturnFields = "comment,disable,extattrs,ipv6addr,name,rp_zone,ttl,use_ttl,view,zone"
)

var RecordRpzAaaaIpaddressResourceSchemaAttributes = map[string]schema.Attribute{
	"id": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The reference to the object.",
	},
	"nios": schema.SingleNestedAttribute{
		Optional:            true,
		MarkdownDescription: "NIOS backend-specific fields.",
		Attributes:          RecordRpzAaaaIpaddressResourceNiosSchemaAttributes,
	},
}

var RecordRpzAaaaIpaddressResourceNiosSchemaAttributes = map[string]schema.Attribute{
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
	"ipv6addr": schema.StringAttribute{
		Required:   true,
		CustomType: iptypes.IPv6AddressType{},
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
		},
		MarkdownDescription: "The IPv6 Address of the substitute rule.",
	},
	"name": schema.StringAttribute{
		Required:   true,
		CustomType: internaltypes.IPv6NameType{},
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
			customvalidator.IsValidNIOSDomainName(),
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
func (m *RecordRpzAaaaIpaddressModel) Expand(ctx context.Context, diags *diag.Diagnostics, isCreate bool) *coremodel.RecordRpzAaaaIpaddress {
	if m == nil {
		return nil
	}

	obj := &coremodel.RecordRpzAaaaIpaddress{}

	// Expand NIOS nested attribute (returns nil if not present)
	niosModel := flex.ExpandNestedObject[NIOSRecordRpzAaaaIpaddressModel](ctx, m.NIOS, diags)
	if niosModel != nil {
		obj.NIOS = niosModel.Expand(ctx, diags)
	}

	return obj
}

// Expand converts the NIOS TF model to the core model.
func (m *NIOSRecordRpzAaaaIpaddressModel) Expand(ctx context.Context, diags *diag.Diagnostics) *coremodel.NIOSRecordRpzAaaaIpaddressExt {
	return &coremodel.NIOSRecordRpzAaaaIpaddressExt{
		Comment:  flex.ExpandStringPointerNullAsEmpty(m.Comment),
		Disable:  flex.ExpandBoolPointer(m.Disable),
		ExtAttrs: flex.ExpandMapStringAny(ctx, m.ExtAttrs, diags),
		Ipv6addr: flex.ExpandIPv6Address(m.Ipv6addr),
		Name:     flex.ExpandStringPointer(m.Name.StringValue),
		RpZone:   flex.ExpandStringPointerNullAsEmpty(m.RpZone),
		Ttl:      flex.ExpandInt64Pointer(m.Ttl),
		View:     flex.ExpandStringPointerNullAsEmpty(m.View),
	}
}

// ApplyRecordRpzAaaaIpaddressNIOSUseFlags derives NIOS use flags from the raw config
// value(s) and writes them onto the core model. A flag is true when the user
// set any of its governed value fields in config.
func ApplyRecordRpzAaaaIpaddressNIOSUseFlags(ctx context.Context, config tfsdk.Config, obj *coremodel.RecordRpzAaaaIpaddress, diags *diag.Diagnostics) {
	if obj == nil || obj.NIOS == nil {
		return
	}
	obj.NIOS.UseTtl = flex.DeriveUseFlag(ctx, config, diags, path.Root("nios").AtName("ttl"))
}

// Flatten populates the TF model from a core response.
func (m *RecordRpzAaaaIpaddressModel) Flatten(ctx context.Context, resp *coremodel.RecordRpzAaaaIpaddress, diags *diag.Diagnostics) {
	if resp == nil {
		return
	}

	m.Id = flex.FlattenStringPointer(resp.Id)

	// Extract existing NIOS model, flatten API response onto it, convert back
	niosModel := flex.ExpandNestedObject[NIOSRecordRpzAaaaIpaddressModel](ctx, m.NIOS, diags)
	if niosModel == nil {
		niosModel = &NIOSRecordRpzAaaaIpaddressModel{}
	}
	niosModel.Flatten(ctx, resp.NIOS, diags)
	if resp.NIOS != nil {
		m.NIOS = flex.FlattenNestedObject(ctx, niosModel, NIOSRecordRpzAaaaIpaddressAttrTypes, diags)
	} else {
		m.NIOS = types.ObjectNull(NIOSRecordRpzAaaaIpaddressAttrTypes)
	}

}

// Flatten merges API response onto existing NIOS model.
func (m *NIOSRecordRpzAaaaIpaddressModel) Flatten(ctx context.Context, from *coremodel.NIOSRecordRpzAaaaIpaddressExt, diags *diag.Diagnostics) {
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
	m.Name.StringValue = flex.FlattenStringPointer(from.Name)
	m.RpZone = flex.FlattenStringPointerEmptyAsNull(from.RpZone)
	m.Ttl = flex.FlattenInt64Pointer(from.Ttl)
	m.View = flex.FlattenStringPointerEmptyAsNull(from.View)
}
