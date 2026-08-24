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
	objectplanmodifier "github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	stringplanmodifier "github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"

	coremodel "github.com/infobloxopen/terraform-provider-infoblox/internal/core/model/dns"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/dynamicallocation"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
	immutable "github.com/infobloxopen/terraform-provider-infoblox/internal/planmodifiers/immutable"
	importmod "github.com/infobloxopen/terraform-provider-infoblox/internal/planmodifiers/import"
	customvalidator "github.com/infobloxopen/terraform-provider-infoblox/internal/validator"
)

type RecordAaaaModel struct {
	Id   types.String `tfsdk:"id"`
	NIOS types.Object `tfsdk:"nios"`
	UDDI types.Object `tfsdk:"uddi"`
}

var RecordAaaaAttrTypes = map[string]attr.Type{
	"id":   types.StringType,
	"nios": types.ObjectType{AttrTypes: NIOSRecordAaaaAttrTypes},
	"uddi": types.ObjectType{AttrTypes: UDDIRecordAaaaAttrTypes},
}

type NIOSRecordAaaaModel struct {
	Comment           types.String        `tfsdk:"comment"`
	Creator           types.String        `tfsdk:"creator"`
	DdnsPrincipal     types.String        `tfsdk:"ddns_principal"`
	DdnsProtected     types.Bool          `tfsdk:"ddns_protected"`
	Disable           types.Bool          `tfsdk:"disable"`
	ExtAttrs          types.Map           `tfsdk:"ext_attrs"`
	ExtAttrsAll       types.Map           `tfsdk:"ext_attrs_all"`
	ForbidReclamation types.Bool          `tfsdk:"forbid_reclamation"`
	Ipv6addr          iptypes.IPv6Address `tfsdk:"ipv6addr"`
	Name              types.String        `tfsdk:"name"`
	Ttl               types.Int64         `tfsdk:"ttl"`
	View              types.String        `tfsdk:"view"`
	DynamicAllocation types.Object        `tfsdk:"dynamic_allocation"`
}

var NIOSRecordAaaaAttrTypes = map[string]attr.Type{
	"comment":            types.StringType,
	"creator":            types.StringType,
	"ddns_principal":     types.StringType,
	"ddns_protected":     types.BoolType,
	"disable":            types.BoolType,
	"ext_attrs":          types.MapType{ElemType: types.StringType},
	"ext_attrs_all":      types.MapType{ElemType: types.StringType},
	"forbid_reclamation": types.BoolType,
	"ipv6addr":           iptypes.IPv6AddressType{},
	"name":               types.StringType,
	"ttl":                types.Int64Type,
	"view":               types.StringType,
	"dynamic_allocation": types.ObjectType{AttrTypes: dynamicallocation.NextAvailableIpAttrTypes},
}

type UDDIRecordAaaaModel struct {
	AbsoluteNameSpec   types.String `tfsdk:"absolute_name_spec"`
	Comment            types.String `tfsdk:"comment"`
	Disabled           types.Bool   `tfsdk:"disabled"`
	InheritanceSources types.Object `tfsdk:"inheritance_sources"`
	NameInZone         types.String `tfsdk:"name_in_zone"`
	Options            types.Object `tfsdk:"options"`
	Rdata              types.Object `tfsdk:"rdata"`
	Tags               types.Map    `tfsdk:"tags"`
	TagsAll            types.Map    `tfsdk:"tags_all"`
	Ttl                types.Int64  `tfsdk:"ttl"`
	Type               types.String `tfsdk:"type"`
	View               types.String `tfsdk:"view"`
	Zone               types.String `tfsdk:"zone"`
}

var UDDIRecordAaaaAttrTypes = map[string]attr.Type{
	"absolute_name_spec":  types.StringType,
	"comment":             types.StringType,
	"disabled":            types.BoolType,
	"inheritance_sources": types.ObjectType{AttrTypes: RecordInheritanceAttrTypes},
	"name_in_zone":        types.StringType,
	"options":             types.ObjectType{AttrTypes: UDDIRecordAaaaOptionsAttrTypes},
	"rdata":               types.ObjectType{AttrTypes: UDDIRecordAaaaRdataAttrTypes},
	"tags":                types.MapType{ElemType: types.StringType},
	"tags_all":            types.MapType{ElemType: types.StringType},
	"ttl":                 types.Int64Type,
	"type":                types.StringType,
	"view":                types.StringType,
	"zone":                types.StringType,
}

const (
	RecordAaaaType            = "AAAA"
	RecordAaaaInheritanceType = "full"
	RecordAaaaReturnFields    = "aws_rte53_record_info,cloud_info,comment,creation_time,creator,ddns_principal,ddns_protected,disable,discovered_data,dns_name,extattrs,forbid_reclamation,ipv6addr,last_queried,ms_ad_user_data,name,reclaimable,shared_record_group,ttl,use_ttl,view,zone"
)

var RecordAaaaResourceSchemaAttributes = map[string]schema.Attribute{
	"id": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The reference to the object.",
	},
	"nios": schema.SingleNestedAttribute{
		Optional:            true,
		MarkdownDescription: "NIOS backend-specific fields.",
		Attributes:          RecordAaaaResourceNiosSchemaAttributes,
	},
	"uddi": schema.SingleNestedAttribute{
		Optional:            true,
		MarkdownDescription: "UDDI backend-specific fields.",
		Attributes:          RecordAaaaResourceUddiSchemaAttributes,
	},
}

var RecordAaaaResourceNiosSchemaAttributes = map[string]schema.Attribute{
	"comment": schema.StringAttribute{
		Optional: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
			customvalidator.ValidateTrimmedString(),
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
	"ipv6addr": schema.StringAttribute{
		Optional:   true,
		Computed:   true,
		CustomType: iptypes.IPv6AddressType{},
		Validators: []validator.String{
			stringvalidator.ExactlyOneOf(
				path.MatchRelative().AtParent().AtName("dynamic_allocation"),
			),
			customvalidator.StringNotEmpty(),
		},
		MarkdownDescription: "The IPv6 Address of the record.",
	},
	"name": schema.StringAttribute{
		Required: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
			customvalidator.IsValidNIOSDomainName(),
		},
		MarkdownDescription: "Name for the AAAA record in FQDN format. This value can be in unicode format.",
	},
	"ttl": schema.Int64Attribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The Time To Live (TTL) value for the record. A 32-bit unsigned integer that represents the duration, in seconds, for which the record is valid (cached). Zero indicates that the record should not be cached.",
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
	"dynamic_allocation": schema.SingleNestedAttribute{
		Attributes:          dynamicallocation.NextAvailableIpResourceSchemaAttributes,
		Optional:            true,
		MarkdownDescription: "Dynamically allocate the ip using the NIOS next_available_ip function call. Mutually exclusive with the static value field.",
	},
}

var RecordAaaaResourceUddiSchemaAttributes = map[string]schema.Attribute{
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
	"options": schema.SingleNestedAttribute{
		Attributes:          UDDIRecordAaaaOptionsResourceSchemaAttributes,
		Optional:            true,
		MarkdownDescription: "The DNS resource record type-specific non-protocol options.",
	},
	"rdata": schema.SingleNestedAttribute{
		Attributes:          UDDIRecordAaaaRdataResourceSchemaAttributes,
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
		Default:             stringdefault.StaticString("AAAA"),
		Computed:            true,
		MarkdownDescription: "The DNS resource record type. Always _AAAA_ for this resource (numeric type 28, IPv6 Address record).",
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
func (m *RecordAaaaModel) Expand(ctx context.Context, diags *diag.Diagnostics, isCreate bool) *coremodel.RecordAaaa {
	if m == nil {
		return nil
	}

	obj := &coremodel.RecordAaaa{}

	// Expand NIOS nested attribute (returns nil if not present)
	niosModel := flex.ExpandNestedObject[NIOSRecordAaaaModel](ctx, m.NIOS, diags)
	if niosModel != nil {
		obj.NIOS = niosModel.Expand(ctx, diags, isCreate)
	}

	// Expand UDDI nested attribute (returns nil if not present)
	uddiModel := flex.ExpandNestedObject[UDDIRecordAaaaModel](ctx, m.UDDI, diags)
	if uddiModel != nil {
		obj.UDDI = uddiModel.Expand(ctx, diags, isCreate)
	}

	return obj
}

// Expand converts the NIOS TF model to the core model.
func (m *NIOSRecordAaaaModel) Expand(ctx context.Context, diags *diag.Diagnostics, isCreate bool) *coremodel.NIOSRecordAaaaExt {
	ext := &coremodel.NIOSRecordAaaaExt{
		Comment:           flex.ExpandStringPointerNullAsEmpty(m.Comment),
		Creator:           flex.ExpandStringPointerNullAsEmpty(m.Creator),
		DdnsPrincipal:     flex.ExpandStringPointerNullAsEmpty(m.DdnsPrincipal),
		DdnsProtected:     flex.ExpandBoolPointer(m.DdnsProtected),
		Disable:           flex.ExpandBoolPointer(m.Disable),
		ExtAttrs:          flex.ExpandMapStringAny(ctx, m.ExtAttrs, diags),
		ForbidReclamation: flex.ExpandBoolPointer(m.ForbidReclamation),
		Ipv6addr:          flex.ExpandIPv6Address(m.Ipv6addr),
		Name:              flex.ExpandStringPointerNullAsEmpty(m.Name),
		Ttl:               flex.ExpandInt64Pointer(m.Ttl),
	}
	if isCreate {
		ext.View = flex.ExpandStringPointerNullAsEmpty(m.View)
		ext.FuncCall = BuildRecordAaaaFuncCall(ctx, m.DynamicAllocation, diags)
	}
	return ext
}

// ApplyRecordAaaaNIOSUseFlags derives NIOS use flags from the raw config
// value(s) and writes them onto the core model. A flag is true when the user
// set any of its governed value fields in config.
func ApplyRecordAaaaNIOSUseFlags(ctx context.Context, config tfsdk.Config, obj *coremodel.RecordAaaa, diags *diag.Diagnostics) {
	if obj == nil || obj.NIOS == nil {
		return
	}
	obj.NIOS.UseTtl = flex.DeriveUseFlag(ctx, config, diags, path.Root("nios").AtName("ttl"))
}

// Expand converts the UDDI TF model to the core model.
func (m *UDDIRecordAaaaModel) Expand(ctx context.Context, diags *diag.Diagnostics, isCreate bool) *coremodel.UDDIRecordAaaaExt {
	ext := &coremodel.UDDIRecordAaaaExt{
		AbsoluteNameSpec:   flex.ExpandStringPointer(m.AbsoluteNameSpec),
		Comment:            flex.ExpandStringPointer(m.Comment),
		Disabled:           flex.ExpandBoolPointer(m.Disabled),
		InheritanceSources: ExpandRecordInheritance(ctx, m.InheritanceSources, diags),
		NameInZone:         flex.ExpandStringPointer(m.NameInZone),
		Options:            ExpandUDDIRecordAaaaOptions(ctx, m.Options, diags),
		Rdata:              ExpandUDDIRecordAaaaRdata(ctx, m.Rdata, diags),
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
func (m *RecordAaaaModel) Flatten(ctx context.Context, resp *coremodel.RecordAaaa, diags *diag.Diagnostics) {
	if resp == nil {
		return
	}

	m.Id = flex.FlattenStringPointer(resp.Id)

	// Extract existing NIOS model, flatten API response onto it, convert back
	niosModel := flex.ExpandNestedObject[NIOSRecordAaaaModel](ctx, m.NIOS, diags)
	if niosModel == nil {
		niosModel = &NIOSRecordAaaaModel{}
	}
	niosModel.Flatten(ctx, resp.NIOS, diags)
	if resp.NIOS != nil {
		m.NIOS = flex.FlattenNestedObject(ctx, niosModel, NIOSRecordAaaaAttrTypes, diags)
	} else {
		m.NIOS = types.ObjectNull(NIOSRecordAaaaAttrTypes)
	}

	// Extract existing UDDI model, flatten API response onto it, convert back
	uddiModel := flex.ExpandNestedObject[UDDIRecordAaaaModel](ctx, m.UDDI, diags)
	if uddiModel == nil {
		uddiModel = &UDDIRecordAaaaModel{}
	}
	plannedUDDI := flex.ExpandNestedObject[UDDIRecordAaaaModel](ctx, m.UDDI, diags)
	uddiModel.Flatten(ctx, resp.UDDI, diags)
	if resp.UDDI != nil {
		PostFlattenRecordAaaaUDDI(ctx, plannedUDDI, uddiModel, diags)
		m.UDDI = flex.FlattenNestedObject(ctx, uddiModel, UDDIRecordAaaaAttrTypes, diags)
	} else {
		m.UDDI = types.ObjectNull(UDDIRecordAaaaAttrTypes)
	}
}

// Flatten merges API response onto existing NIOS model.
func (m *NIOSRecordAaaaModel) Flatten(ctx context.Context, from *coremodel.NIOSRecordAaaaExt, diags *diag.Diagnostics) {
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
	m.Ipv6addr = flex.FlattenIPv6Address(from.Ipv6addr)
	m.Name = flex.FlattenStringPointerEmptyAsNull(from.Name)
	m.Ttl = flex.FlattenInt64Pointer(from.Ttl)
	m.View = flex.FlattenStringPointerEmptyAsNull(from.View)
	if len(m.DynamicAllocation.AttributeTypes(ctx)) == 0 {
		m.DynamicAllocation = types.ObjectNull(dynamicallocation.NextAvailableIpAttrTypes)
	}
}

// Flatten merges API response onto existing UDDI model.
func (m *UDDIRecordAaaaModel) Flatten(ctx context.Context, from *coremodel.UDDIRecordAaaaExt, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	m.AbsoluteNameSpec = flex.FlattenStringPointer(from.AbsoluteNameSpec)
	m.Comment = flex.FlattenStringPointer(from.Comment)
	m.Disabled = flex.FlattenBoolPointer(from.Disabled)
	m.InheritanceSources = FlattenRecordInheritance(ctx, from.InheritanceSources, diags)
	m.NameInZone = flex.FlattenStringPointer(from.NameInZone)
	m.Options = FlattenUDDIRecordAaaaOptions(ctx, from.Options, diags)
	m.Rdata = FlattenUDDIRecordAaaaRdata(ctx, from.Rdata, diags)
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
