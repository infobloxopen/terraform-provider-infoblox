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
	objectplanmodifier "github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	stringplanmodifier "github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"

	coremodel "github.com/infobloxopen/terraform-provider-infoblox/internal/core/model/dns"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
	immutable "github.com/infobloxopen/terraform-provider-infoblox/internal/planmodifiers/immutable"
	importmod "github.com/infobloxopen/terraform-provider-infoblox/internal/planmodifiers/import"
	customvalidator "github.com/infobloxopen/terraform-provider-infoblox/internal/validator"
)

type RecordMxModel struct {
	Id   types.String `tfsdk:"id"`
	NIOS types.Object `tfsdk:"nios"`
	UDDI types.Object `tfsdk:"uddi"`
}

var RecordMxAttrTypes = map[string]attr.Type{
	"id":   types.StringType,
	"nios": types.ObjectType{AttrTypes: NIOSRecordMxAttrTypes},
	"uddi": types.ObjectType{AttrTypes: UDDIRecordMxAttrTypes},
}

type NIOSRecordMxModel struct {
	Comment           types.String `tfsdk:"comment"`
	Creator           types.String `tfsdk:"creator"`
	DdnsPrincipal     types.String `tfsdk:"ddns_principal"`
	DdnsProtected     types.Bool   `tfsdk:"ddns_protected"`
	Disable           types.Bool   `tfsdk:"disable"`
	ExtAttrs          types.Map    `tfsdk:"ext_attrs"`
	ExtAttrsAll       types.Map    `tfsdk:"ext_attrs_all"`
	ForbidReclamation types.Bool   `tfsdk:"forbid_reclamation"`
	MailExchanger     types.String `tfsdk:"mail_exchanger"`
	Name              types.String `tfsdk:"name"`
	Preference        types.Int64  `tfsdk:"preference"`
	Ttl               types.Int64  `tfsdk:"ttl"`
	View              types.String `tfsdk:"view"`
}

var NIOSRecordMxAttrTypes = map[string]attr.Type{
	"comment":            types.StringType,
	"creator":            types.StringType,
	"ddns_principal":     types.StringType,
	"ddns_protected":     types.BoolType,
	"disable":            types.BoolType,
	"ext_attrs":          types.MapType{ElemType: types.StringType},
	"ext_attrs_all":      types.MapType{ElemType: types.StringType},
	"forbid_reclamation": types.BoolType,
	"mail_exchanger":     types.StringType,
	"name":               types.StringType,
	"preference":         types.Int64Type,
	"ttl":                types.Int64Type,
	"view":               types.StringType,
}

type UDDIRecordMxModel struct {
	AbsoluteNameSpec   types.String `tfsdk:"absolute_name_spec"`
	Comment            types.String `tfsdk:"comment"`
	Disabled           types.Bool   `tfsdk:"disabled"`
	InheritanceSources types.Object `tfsdk:"inheritance_sources"`
	NameInZone         types.String `tfsdk:"name_in_zone"`
	Rdata              types.Object `tfsdk:"rdata"`
	Tags               types.Map    `tfsdk:"tags"`
	TagsAll            types.Map    `tfsdk:"tags_all"`
	Ttl                types.Int64  `tfsdk:"ttl"`
	Type               types.String `tfsdk:"type"`
	View               types.String `tfsdk:"view"`
	Zone               types.String `tfsdk:"zone"`
}

var UDDIRecordMxAttrTypes = map[string]attr.Type{
	"absolute_name_spec":  types.StringType,
	"comment":             types.StringType,
	"disabled":            types.BoolType,
	"inheritance_sources": types.ObjectType{AttrTypes: RecordInheritanceAttrTypes},
	"name_in_zone":        types.StringType,
	"rdata":               types.ObjectType{AttrTypes: UDDIRecordMxRdataAttrTypes},
	"tags":                types.MapType{ElemType: types.StringType},
	"tags_all":            types.MapType{ElemType: types.StringType},
	"ttl":                 types.Int64Type,
	"type":                types.StringType,
	"view":                types.StringType,
	"zone":                types.StringType,
}

const (
	RecordMxType            = "MX"
	RecordMxInheritanceType = "full"
	RecordMxReturnFields    = "aws_rte53_record_info,cloud_info,comment,creation_time,creator,ddns_principal,ddns_protected,disable,dns_mail_exchanger,dns_name,extattrs,forbid_reclamation,last_queried,mail_exchanger,name,preference,reclaimable,shared_record_group,ttl,use_ttl,view,zone"
)

var RecordMxResourceSchemaAttributes = map[string]schema.Attribute{
	"id": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The reference to the object.",
	},
	"nios": schema.SingleNestedAttribute{
		Optional:            true,
		MarkdownDescription: "NIOS backend-specific fields.",
		Attributes:          RecordMxResourceNiosSchemaAttributes,
	},
	"uddi": schema.SingleNestedAttribute{
		Optional:            true,
		MarkdownDescription: "UDDI backend-specific fields.",
		Attributes:          RecordMxResourceUddiSchemaAttributes,
	},
}

var RecordMxResourceNiosSchemaAttributes = map[string]schema.Attribute{
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
			stringvalidator.OneOf("STATIC", "DYNAMIC"),
		},
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The record creator. Note that changing creator from or to 'SYSTEM' value is not allowed.",
	},
	"ddns_principal": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
		},
		MarkdownDescription: "The GSS-TSIG principal that owns this record.",
	},
	"ddns_protected": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Determines if the DDNS updates for this record are allowed or not.",
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
	"forbid_reclamation": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Determines if the reclamation is allowed for the record or not.",
	},
	"mail_exchanger": schema.StringAttribute{
		Required: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
			customvalidator.IsValidNIOSDomainName(),
		},
		MarkdownDescription: "Mail exchanger name in FQDN format. This value can be in unicode format.",
	},
	"name": schema.StringAttribute{
		Required: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
			customvalidator.IsValidNIOSDomainName(),
		},
		MarkdownDescription: "Name for the MX record in FQDN format. This value can be in unicode format.",
	},
	"preference": schema.Int64Attribute{
		Required: true,
		Validators: []validator.Int64{
			int64validator.Between(0, 65535),
		},
		MarkdownDescription: "Preference value, 0 to 65535 (inclusive) in 32-bit unsigned integer format.",
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
		},
		MarkdownDescription: "The name of the DNS view in which the record resides. Example: \"external\".",
	},
}

var RecordMxResourceUddiSchemaAttributes = map[string]schema.Attribute{
	"absolute_name_spec": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Validators: []validator.String{
			stringvalidator.AlsoRequires(path.MatchRelative().AtParent().AtName("view")),
			stringvalidator.ConflictsWith(
				path.MatchRelative().AtParent().AtName("zone"),
				path.MatchRelative().AtParent().AtName("name_in_zone"),
			),
			customvalidator.IsValidUDDIDomainName(),
		},
		MarkdownDescription: "Synthetic field, used to determine _zone_ and/or _name_in_zone_ field for records.",
	},
	"comment": schema.StringAttribute{
		Default:             stringdefault.StaticString(""),
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The description for the DNS resource record. May contain 0 to 1024 characters. Can include UTF-8.",
	},
	"disabled": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Indicates if the DNS resource record is disabled. A disabled object is effectively non-existent when generating configuration.  Defaults to _false_.",
	},
	"inheritance_sources": schema.SingleNestedAttribute{
		Attributes: RecordInheritanceResourceSchemaAttributes,
		Optional:   true,
		Computed:   true,
		PlanModifiers: []planmodifier.Object{
			objectplanmodifier.UseStateForUnknown(),
		},
		MarkdownDescription: "The inheritance configuration specifies how the _Record_ object inherits the _ttl_ field.",
	},
	"name_in_zone": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Validators: []validator.String{
			stringvalidator.AlsoRequires(path.MatchRelative().AtParent().AtName("zone")),
			stringvalidator.ConflictsWith(
				path.MatchRelative().AtParent().AtName("absolute_name_spec"),
				path.MatchRelative().AtParent().AtName("view"),
			),
		},
		MarkdownDescription: "The relative owner name to the zone origin. Must be specified for creating the DNS resource record and is read only for other operations.",
	},
	"rdata": schema.SingleNestedAttribute{
		Attributes:          UDDIRecordMxRdataResourceSchemaAttributes,
		Required:            true,
		MarkdownDescription: "The DNS resource record data in JSON format. Certain DNS resource record-specific subfields are required for creating the DNS resource record.",
	},
	"tags": schema.MapAttribute{
		Optional:    true,
		Computed:    true,
		ElementType: types.StringType,
		Default:     mapdefault.StaticValue(types.MapNull(types.StringType)),
		Validators: []validator.Map{
			mapvalidator.SizeAtLeast(1),
		},
		MarkdownDescription: "The tags for the DNS resource record in JSON format.",
	},
	"tags_all": schema.MapAttribute{
		Computed:            true,
		ElementType:         types.StringType,
		MarkdownDescription: "All tags including inherited values.",
	},
	"ttl": schema.Int64Attribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The record time to live value in seconds. The range of this value is 0 to 2147483647.  Defaults to TTL value from the SOA record of the zone.",
	},
	"type": schema.StringAttribute{
		Default:             stringdefault.StaticString("MX"),
		Computed:            true,
		MarkdownDescription: "The DNS resource record type. Always _MX_ for this resource (numeric type 15, Mail Exchanger record).",
	},
	"view": schema.StringAttribute{
		Optional: true,
		Computed: true,
		PlanModifiers: []planmodifier.String{
			stringplanmodifier.RequiresReplaceIfConfigured(),
		},
		Validators: []validator.String{
			stringvalidator.AlsoRequires(path.MatchRelative().AtParent().AtName("absolute_name_spec")),
			stringvalidator.ConflictsWith(
				path.MatchRelative().AtParent().AtName("zone"),
				path.MatchRelative().AtParent().AtName("name_in_zone"),
			),
		},
		MarkdownDescription: "The resource identifier.",
	},
	"zone": schema.StringAttribute{
		Optional: true,
		Computed: true,
		PlanModifiers: []planmodifier.String{
			stringplanmodifier.RequiresReplaceIfConfigured(),
		},
		Validators: []validator.String{
			stringvalidator.ConflictsWith(
				path.MatchRelative().AtParent().AtName("absolute_name_spec"),
				path.MatchRelative().AtParent().AtName("view"),
			),
		},
		MarkdownDescription: "The resource identifier.",
	},
}

// Expand converts the TF model to the infoblox core model
func (m *RecordMxModel) Expand(ctx context.Context, diags *diag.Diagnostics, isCreate bool) *coremodel.RecordMx {
	if m == nil {
		return nil
	}

	obj := &coremodel.RecordMx{}

	// Expand NIOS nested attribute (returns nil if not present)
	niosModel := flex.ExpandNestedObject[NIOSRecordMxModel](ctx, m.NIOS, diags)
	if niosModel != nil {
		obj.NIOS = niosModel.Expand(ctx, diags)
	}

	// Expand UDDI nested attribute (returns nil if not present)
	uddiModel := flex.ExpandNestedObject[UDDIRecordMxModel](ctx, m.UDDI, diags)
	if uddiModel != nil {
		obj.UDDI = uddiModel.Expand(ctx, diags, isCreate)
	}

	return obj
}

// Expand converts the NIOS TF model to the core model.
func (m *NIOSRecordMxModel) Expand(ctx context.Context, diags *diag.Diagnostics) *coremodel.NIOSRecordMxExt {
	return &coremodel.NIOSRecordMxExt{
		Comment:           flex.ExpandStringPointerNullAsEmpty(m.Comment),
		Creator:           flex.ExpandStringPointerNullAsEmpty(m.Creator),
		DdnsPrincipal:     flex.ExpandStringPointerNullAsEmpty(m.DdnsPrincipal),
		DdnsProtected:     flex.ExpandBoolPointer(m.DdnsProtected),
		Disable:           flex.ExpandBoolPointer(m.Disable),
		ExtAttrs:          flex.ExpandMapStringAny(ctx, m.ExtAttrs, diags),
		ForbidReclamation: flex.ExpandBoolPointer(m.ForbidReclamation),
		MailExchanger:     flex.ExpandStringPointerNullAsEmpty(m.MailExchanger),
		Name:              flex.ExpandStringPointerNullAsEmpty(m.Name),
		Preference:        flex.ExpandInt64Pointer(m.Preference),
		Ttl:               flex.ExpandInt64Pointer(m.Ttl),
		View:              flex.ExpandStringPointerNullAsEmpty(m.View),
	}
}

// ApplyRecordMxNIOSUseFlags derives NIOS use flags from the raw config
// value(s) and writes them onto the core model. A flag is true when the user
// set any of its governed value fields in config.
func ApplyRecordMxNIOSUseFlags(ctx context.Context, config tfsdk.Config, obj *coremodel.RecordMx, diags *diag.Diagnostics) {
	if obj == nil || obj.NIOS == nil {
		return
	}
	obj.NIOS.UseTtl = flex.DeriveUseFlag(ctx, config, diags, path.Root("nios").AtName("ttl"))
}

// Expand converts the UDDI TF model to the core model.
func (m *UDDIRecordMxModel) Expand(ctx context.Context, diags *diag.Diagnostics, isCreate bool) *coremodel.UDDIRecordMxExt {
	ext := &coremodel.UDDIRecordMxExt{
		AbsoluteNameSpec:   flex.ExpandStringPointer(m.AbsoluteNameSpec),
		Comment:            flex.ExpandStringPointer(m.Comment),
		Disabled:           flex.ExpandBoolPointer(m.Disabled),
		InheritanceSources: ExpandRecordInheritance(ctx, m.InheritanceSources, diags),
		NameInZone:         flex.ExpandStringPointer(m.NameInZone),
		Rdata:              ExpandUDDIRecordMxRdata(ctx, m.Rdata, diags),
		Tags:               flex.ExpandMapStringAny(ctx, m.Tags, diags),
		Ttl:                flex.ExpandInt64Pointer(m.Ttl),
	}
	if isCreate {
		ext.Type = flex.ExpandStringPointer(m.Type)
		ext.View = flex.ExpandStringPointer(m.View)
		ext.Zone = flex.ExpandStringPointer(m.Zone)
	}
	return ext
}

// Flatten populates the TF model from a core response.
func (m *RecordMxModel) Flatten(ctx context.Context, resp *coremodel.RecordMx, diags *diag.Diagnostics) {
	if resp == nil {
		return
	}

	m.Id = flex.FlattenStringPointer(resp.Id)

	// Extract existing NIOS model, flatten API response onto it, convert back
	niosModel := flex.ExpandNestedObject[NIOSRecordMxModel](ctx, m.NIOS, diags)
	if niosModel == nil {
		niosModel = &NIOSRecordMxModel{}
	}
	niosModel.Flatten(ctx, resp.NIOS, diags)
	if resp.NIOS != nil {
		m.NIOS = flex.FlattenNestedObject(ctx, niosModel, NIOSRecordMxAttrTypes, diags)
	} else {
		m.NIOS = types.ObjectNull(NIOSRecordMxAttrTypes)
	}

	// Extract existing UDDI model, flatten API response onto it, convert back
	uddiModel := flex.ExpandNestedObject[UDDIRecordMxModel](ctx, m.UDDI, diags)
	if uddiModel == nil {
		uddiModel = &UDDIRecordMxModel{}
	}
	uddiModel.Flatten(ctx, resp.UDDI, diags)
	if resp.UDDI != nil {
		m.UDDI = flex.FlattenNestedObject(ctx, uddiModel, UDDIRecordMxAttrTypes, diags)
	} else {
		m.UDDI = types.ObjectNull(UDDIRecordMxAttrTypes)
	}
}

// Flatten merges API response onto existing NIOS model.
func (m *NIOSRecordMxModel) Flatten(ctx context.Context, from *coremodel.NIOSRecordMxExt, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	planExtAttrs := m.ExtAttrs
	if planExtAttrs.IsUnknown() {
		planExtAttrs = types.MapNull(types.StringType)
	}
	m.Comment = flex.FlattenStringPointerEmptyAsNull(from.Comment)
	m.Creator = flex.FlattenStringPointerEmptyAsNull(from.Creator)
	m.DdnsPrincipal = flex.FlattenStringPointerEmptyAsNull(from.DdnsPrincipal)
	m.DdnsProtected = flex.FlattenBoolPointer(from.DdnsProtected)
	m.Disable = flex.FlattenBoolPointer(from.Disable)
	m.ExtAttrs, m.ExtAttrsAll = flex.FlattenEAs(planExtAttrs, from.ExtAttrs)
	m.ForbidReclamation = flex.FlattenBoolPointer(from.ForbidReclamation)
	m.MailExchanger = flex.FlattenStringPointerEmptyAsNull(from.MailExchanger)
	m.Name = flex.FlattenStringPointerEmptyAsNull(from.Name)
	m.Preference = flex.FlattenInt64Pointer(from.Preference)
	m.Ttl = flex.FlattenInt64Pointer(from.Ttl)
	m.View = flex.FlattenStringPointerEmptyAsNull(from.View)
}

// Flatten merges API response onto existing UDDI model.
func (m *UDDIRecordMxModel) Flatten(ctx context.Context, from *coremodel.UDDIRecordMxExt, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	m.AbsoluteNameSpec = flex.FlattenStringPointer(from.AbsoluteNameSpec)
	m.Comment = flex.FlattenStringPointer(from.Comment)
	m.Disabled = flex.FlattenBoolPointer(from.Disabled)
	m.InheritanceSources = FlattenRecordInheritance(ctx, from.InheritanceSources, diags)
	m.NameInZone = flex.FlattenStringPointer(from.NameInZone)
	m.Rdata = FlattenUDDIRecordMxRdata(ctx, from.Rdata, diags)
	tagsAll := flex.FlattenMapStringAny(ctx, from.Tags, diags)
	if m.Tags.IsNull() || m.Tags.IsUnknown() {
		m.Tags = tagsAll
	}
	m.TagsAll = tagsAll
	m.Ttl = flex.FlattenInt64Pointer(from.Ttl)
	m.Type = flex.FlattenStringPointer(from.Type)
	m.View = flex.FlattenStringPointer(from.View)
	m.Zone = flex.FlattenStringPointer(from.Zone)
}
