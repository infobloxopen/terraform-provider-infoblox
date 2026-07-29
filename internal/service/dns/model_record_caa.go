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
	objectplanmodifier "github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	stringplanmodifier "github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"

	coremodel "github.com/infobloxopen/terraform-provider-infoblox/internal/core/model/dns"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
	importmod "github.com/infobloxopen/terraform-provider-infoblox/internal/planmodifiers/import"
	customvalidator "github.com/infobloxopen/terraform-provider-infoblox/internal/validator"
)

type RecordCaaModel struct {
	Id   types.String `tfsdk:"id"`
	NIOS types.Object `tfsdk:"nios"`
	UDDI types.Object `tfsdk:"uddi"`
}

var RecordCaaAttrTypes = map[string]attr.Type{
	"id":   types.StringType,
	"nios": types.ObjectType{AttrTypes: NIOSRecordCaaAttrTypes},
	"uddi": types.ObjectType{AttrTypes: UDDIRecordCaaAttrTypes},
}

type NIOSRecordCaaModel struct {
	CaFlag            types.Int64  `tfsdk:"ca_flag"`
	CaTag             types.String `tfsdk:"ca_tag"`
	CaValue           types.String `tfsdk:"ca_value"`
	Comment           types.String `tfsdk:"comment"`
	Creator           types.String `tfsdk:"creator"`
	DdnsPrincipal     types.String `tfsdk:"ddns_principal"`
	DdnsProtected     types.Bool   `tfsdk:"ddns_protected"`
	Disable           types.Bool   `tfsdk:"disable"`
	ExtAttrs          types.Map    `tfsdk:"ext_attrs"`
	ExtAttrsAll       types.Map    `tfsdk:"ext_attrs_all"`
	ForbidReclamation types.Bool   `tfsdk:"forbid_reclamation"`
	Name              types.String `tfsdk:"name"`
	Ttl               types.Int64  `tfsdk:"ttl"`
	View              types.String `tfsdk:"view"`
}

var NIOSRecordCaaAttrTypes = map[string]attr.Type{
	"ca_flag":            types.Int64Type,
	"ca_tag":             types.StringType,
	"ca_value":           types.StringType,
	"comment":            types.StringType,
	"creator":            types.StringType,
	"ddns_principal":     types.StringType,
	"ddns_protected":     types.BoolType,
	"disable":            types.BoolType,
	"ext_attrs":          types.MapType{ElemType: types.StringType},
	"ext_attrs_all":      types.MapType{ElemType: types.StringType},
	"forbid_reclamation": types.BoolType,
	"name":               types.StringType,
	"ttl":                types.Int64Type,
	"view":               types.StringType,
}

type UDDIRecordCaaModel struct {
	AbsoluteNameSpec   types.String `tfsdk:"absolute_name_spec"`
	Comment            types.String `tfsdk:"comment"`
	Disabled           types.Bool   `tfsdk:"disabled"`
	InheritanceSources types.Object `tfsdk:"inheritance_sources"`
	NameInZone         types.String `tfsdk:"name_in_zone"`
	Options            types.Map    `tfsdk:"options"`
	Rdata              types.Map    `tfsdk:"rdata"`
	Tags               types.Map    `tfsdk:"tags"`
	TagsAll            types.Map    `tfsdk:"tags_all"`
	Ttl                types.Int64  `tfsdk:"ttl"`
	Type               types.String `tfsdk:"type"`
	View               types.String `tfsdk:"view"`
	Zone               types.String `tfsdk:"zone"`
}

var UDDIRecordCaaAttrTypes = map[string]attr.Type{
	"absolute_name_spec":  types.StringType,
	"comment":             types.StringType,
	"disabled":            types.BoolType,
	"inheritance_sources": types.ObjectType{AttrTypes: RecordInheritanceAttrTypes},
	"name_in_zone":        types.StringType,
	"options":             types.MapType{ElemType: types.StringType},
	"rdata":               types.MapType{ElemType: types.StringType},
	"tags":                types.MapType{ElemType: types.StringType},
	"tags_all":            types.MapType{ElemType: types.StringType},
	"ttl":                 types.Int64Type,
	"type":                types.StringType,
	"view":                types.StringType,
	"zone":                types.StringType,
}

const (
	RecordCaaType            = "Caa"
	RecordCaaInheritanceType = "full"
	RecordCaaReturnFields    = "ca_flag,ca_tag,ca_value,cloud_info,comment,creation_time,creator,ddns_principal,ddns_protected,disable,dns_name,extattrs,forbid_reclamation,last_queried,name,reclaimable,ttl,use_ttl,view,zone"
)

var RecordCaaResourceSchemaAttributes = map[string]schema.Attribute{
	"id": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The reference to the object.",
	},
	"nios": schema.SingleNestedAttribute{
		Optional:            true,
		MarkdownDescription: "NIOS backend-specific fields.",
		Attributes:          RecordCaaResourceNiosSchemaAttributes,
	},
	"uddi": schema.SingleNestedAttribute{
		Optional:            true,
		MarkdownDescription: "UDDI backend-specific fields.",
		Attributes:          RecordCaaResourceUddiSchemaAttributes,
	},
}

var RecordCaaResourceNiosSchemaAttributes = map[string]schema.Attribute{
	"ca_flag": schema.Int64Attribute{
		Required:            true,
		MarkdownDescription: "Flag of CAA record.",
	},
	"ca_tag": schema.StringAttribute{
		Required: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
			stringvalidator.LengthBetween(1, 255),
			customvalidator.ValidateTrimmedString(),
		},
		MarkdownDescription: "Tag of CAA record.",
	},
	"ca_value": schema.StringAttribute{
		Required: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
		},
		MarkdownDescription: "Value of CAA record",
	},
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
			stringvalidator.OneOf("STATIC", "DYNAMIC", "SYSTEM"),
		},
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The record creator. Note that changing creator from or to 'SYSTEM' value is not allowed.",
	},
	"ddns_principal": schema.StringAttribute{
		Optional: true,
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
	"name": schema.StringAttribute{
		Required: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
			customvalidator.IsValidDomainName(),
		},
		MarkdownDescription: "The CAA record name in FQDN format. This value can be in unicode format.",
	},
	"ttl": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "The Time to Live (TTL) value for the record. A 32-bit unsigned integer that represents the duration, in seconds, for which the record is valid (cached). Zero indicates that the record should not be cached.",
	},
	"view": schema.StringAttribute{
		Default:  stringdefault.StaticString("default"),
		Optional: true,
		Computed: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
			customvalidator.ValidateTrimmedString(),
		},
		MarkdownDescription: "The name of the DNS view in which the record resides. Example: \"external\".",
	},
}

var RecordCaaResourceUddiSchemaAttributes = map[string]schema.Attribute{
	"absolute_name_spec": schema.StringAttribute{
		Optional: true,
		Validators: []validator.String{
			stringvalidator.AlsoRequires(path.MatchRelative().AtParent().AtName("view")),
			stringvalidator.ConflictsWith(
				path.MatchRelative().AtParent().AtName("zone"),
				path.MatchRelative().AtParent().AtName("name_in_zone"),
			),
		},
		MarkdownDescription: "Synthetic field, used to determine _zone_ and/or _name_in_zone_ field for records.",
	},
	"comment": schema.StringAttribute{
		Optional:            true,
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
		PlanModifiers: []planmodifier.Object{
			objectplanmodifier.UseStateForUnknown(),
		},
		MarkdownDescription: "The inheritance configuration specifies how the _Record_ object inherits the _ttl_ field.",
	},
	"name_in_zone": schema.StringAttribute{
		Optional: true,
		Validators: []validator.String{
			stringvalidator.AlsoRequires(path.MatchRelative().AtParent().AtName("zone")),
			stringvalidator.ConflictsWith(
				path.MatchRelative().AtParent().AtName("absolute_name_spec"),
				path.MatchRelative().AtParent().AtName("view"),
			),
		},
		MarkdownDescription: "The relative owner name to the zone origin. Must be specified for creating the DNS resource record and is read only for other operations.",
	},
	"options": schema.MapAttribute{
		ElementType:         types.StringType,
		Optional:            true,
		MarkdownDescription: "The DNS resource record type-specific non-protocol options.  Valid value for _A_ (Address) and _AAAA_ (IPv6 Address) records:  Option     | Description -----------|----------------------------------------- create_ptr | A boolean flag which can be set to _true_ for POST operation to automatically create the corresponding PTR record. check_rmz  | A boolean flag which can be set to _true_ for POST operation to check the existence of reverse zone for creating the corresponding PTR record. Only applicable if the _create_ptr_ option is set to _true_.   Valid value for _PTR_ (Pointer) records:  Option     | Description -----------|---------------------------------------- address    | For GET operation it contains the IPv4 or IPv6 address represented by the PTR record.<br><br>For POST and PATCH operations it can be used to create/update a PTR record based on the IP address it represents. In this case, in addition to the _address_ in the options field, need to specify the _view_ field. |",
	},
	"rdata": schema.MapAttribute{
		ElementType:         types.StringType,
		Required:            true,
		MarkdownDescription: "The DNS resource record data in JSON format. Certain DNS resource record-specific subfields are required for creating the DNS resource record.  Subfields for _CAA_ (Certification Authority Authorization) record:  Subfield | Description                           | Required ---------|---------------------------------------|--------- flags    | An unsigned 8-bit integer which specifies the CAA record flags. RFC 6844 defines one (highest) bit in flag octet, remaining bits are deferred for future use. This bit is referenced as _Critical_. When the bit is set (flag value == 128), issuers must not issue certificates in case CAA records contain unknown property tags.<br><br>Defaults to 0.<br><br> | No tag      | The CAA record property tag string which indicates the type of CAA record. The following property tags are defined by RFC 6844:<ul><li>_issue_: Used to explicitly authorize CA to issue certificates for the domain in which the property is published.</li><li>_issuewild_: Used to explicitly authorize a single CA to issue wildcard certificates for the domain in which the property is published.</li><li>_iodef_: Used to specify an email address or URL to report invalid certificate requests or issuers’ certificate policy violations.</li></ul>Note: _issuewild_ type takes precedence over _issue_.<br><br> | Yes value    | A string which contains the CAA record property value.<br><br>Specifies the CA who is authorized to issue a certificate for the domain if the CAA record property tag is _issue_ or _issuewild_.<br><br> Specifies the URL/email address to report CAA policy violation for the domain if the CAA record property tag is _iodef_.<br><br> | Yes",
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
		MarkdownDescription: "The record time to live value in seconds. The range of this value is 0 to 2147483647.  Defaults to TTL value from the SOA record of the zone.",
	},
	"type": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The DNS resource record type specified in the textual mnemonic format or in the \"TYPEnnn\" format where \"nnn\" indicates the numeric type value.  Value  | Numeric Type | Description -------|--------------|--------------------------------------------- A      | 1            | Address record AAAA   | 28           | IPv6 Address record CAA    | 257          | Certification Authority Authorization record CNAME  | 5            | Canonical Name record DNAME  | 39           | Delegation Name record DHCID  | 49           | DHCP Identifier record MX     | 15           | Mail Exchanger record NAPTR  | 35           | Naming Authority Pointer record NS     | 2            | Name Server record PTR    | 12           | Pointer record SOA    | 6            | Start of Authority record SRV    | 33           | Service record TXT    | 16           | Text record IBMETA | 65536        | Infoblox meta records, not valid for DNS protocol (read-only)",
	},
	"view": schema.StringAttribute{
		Optional: true,
		PlanModifiers: []planmodifier.String{
			stringplanmodifier.RequiresReplaceIfConfigured(),
		},
		Validators: []validator.String{
			stringvalidator.Any(stringvalidator.AlsoRequires(path.MatchRelative().AtParent().AtName("absolute_name_spec")), stringvalidator.AlsoRequires(path.MatchRelative().AtParent().AtName("options").AtName("address"))),
			stringvalidator.ConflictsWith(
				path.MatchRelative().AtParent().AtName("zone"),
				path.MatchRelative().AtParent().AtName("name_in_zone"),
			),
		},
		MarkdownDescription: "The resource identifier.",
	},
	"zone": schema.StringAttribute{
		Optional: true,
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
func (m *RecordCaaModel) Expand(ctx context.Context, diags *diag.Diagnostics, isCreate bool) *coremodel.RecordCaa {
	if m == nil {
		return nil
	}

	obj := &coremodel.RecordCaa{}

	// Expand NIOS nested attribute (returns nil if not present)
	niosModel := flex.ExpandNestedObject[NIOSRecordCaaModel](ctx, m.NIOS, diags)
	if niosModel != nil {
		obj.NIOS = niosModel.Expand(ctx, diags)
	}

	// Expand UDDI nested attribute (returns nil if not present)
	uddiModel := flex.ExpandNestedObject[UDDIRecordCaaModel](ctx, m.UDDI, diags)
	if uddiModel != nil {
		obj.UDDI = uddiModel.Expand(ctx, diags)
	}

	return obj
}

// Expand converts the NIOS TF model to the core model.
func (m *NIOSRecordCaaModel) Expand(ctx context.Context, diags *diag.Diagnostics) *coremodel.NIOSRecordCaaExt {
	return &coremodel.NIOSRecordCaaExt{
		CaFlag:            flex.ExpandInt64Pointer(m.CaFlag),
		CaTag:             flex.ExpandStringPointerNullAsEmpty(m.CaTag),
		CaValue:           flex.ExpandStringPointerNullAsEmpty(m.CaValue),
		Comment:           flex.ExpandStringPointerNullAsEmpty(m.Comment),
		Creator:           flex.ExpandStringPointerNullAsEmpty(m.Creator),
		DdnsPrincipal:     flex.ExpandStringPointerNullAsEmpty(m.DdnsPrincipal),
		DdnsProtected:     flex.ExpandBoolPointer(m.DdnsProtected),
		Disable:           flex.ExpandBoolPointer(m.Disable),
		ExtAttrs:          flex.ExpandMapStringAny(ctx, m.ExtAttrs, diags),
		ForbidReclamation: flex.ExpandBoolPointer(m.ForbidReclamation),
		Name:              flex.ExpandStringPointerNullAsEmpty(m.Name),
		Ttl:               flex.ExpandInt64Pointer(m.Ttl),
		View:              flex.ExpandStringPointerNullAsEmpty(m.View),
	}
}

// ApplyRecordCaaNIOSUseFlags derives NIOS use flags from the raw config
// value(s) and writes them onto the core model. A flag is true when the user
// set any of its governed value fields in config.
func ApplyRecordCaaNIOSUseFlags(ctx context.Context, config tfsdk.Config, obj *coremodel.RecordCaa, diags *diag.Diagnostics) {
	if obj == nil || obj.NIOS == nil {
		return
	}
	obj.NIOS.UseTtl = flex.DeriveUseFlag(ctx, config, diags, path.Root("nios").AtName("ttl"))
}

// Expand converts the UDDI TF model to the core model.
func (m *UDDIRecordCaaModel) Expand(ctx context.Context, diags *diag.Diagnostics) *coremodel.UDDIRecordCaaExt {
	return &coremodel.UDDIRecordCaaExt{
		AbsoluteNameSpec:   flex.ExpandStringPointer(m.AbsoluteNameSpec),
		Comment:            flex.ExpandStringPointer(m.Comment),
		Disabled:           flex.ExpandBoolPointer(m.Disabled),
		InheritanceSources: ExpandRecordInheritance(ctx, m.InheritanceSources, diags),
		NameInZone:         flex.ExpandStringPointer(m.NameInZone),
		Options:            flex.ExpandMapStringAny(ctx, m.Options, diags),
		Rdata:              flex.ExpandMapStringAny(ctx, m.Rdata, diags),
		Tags:               flex.ExpandMapStringAny(ctx, m.Tags, diags),
		Ttl:                flex.ExpandInt64Pointer(m.Ttl),
		Type:               flex.ExpandStringPointer(m.Type),
		View:               flex.ExpandStringPointer(m.View),
		Zone:               flex.ExpandStringPointer(m.Zone),
	}
}

// Flatten populates the TF model from a core response.
func (m *RecordCaaModel) Flatten(ctx context.Context, resp *coremodel.RecordCaa, diags *diag.Diagnostics) {
	if resp == nil {
		return
	}

	m.Id = flex.FlattenStringPointer(resp.Id)

	// Extract existing NIOS model, flatten API response onto it, convert back
	niosModel := flex.ExpandNestedObject[NIOSRecordCaaModel](ctx, m.NIOS, diags)
	if niosModel == nil {
		niosModel = &NIOSRecordCaaModel{}
	}
	niosModel.Flatten(ctx, resp.NIOS, diags)
	if resp.NIOS != nil {
		m.NIOS = flex.FlattenNestedObject(ctx, niosModel, NIOSRecordCaaAttrTypes, diags)
	} else {
		m.NIOS = types.ObjectNull(NIOSRecordCaaAttrTypes)
	}

	// Extract existing UDDI model, flatten API response onto it, convert back
	uddiModel := flex.ExpandNestedObject[UDDIRecordCaaModel](ctx, m.UDDI, diags)
	if uddiModel == nil {
		uddiModel = &UDDIRecordCaaModel{}
	}
	uddiModel.Flatten(ctx, resp.UDDI, diags)
	if resp.UDDI != nil {
		m.UDDI = flex.FlattenNestedObject(ctx, uddiModel, UDDIRecordCaaAttrTypes, diags)
	} else {
		m.UDDI = types.ObjectNull(UDDIRecordCaaAttrTypes)
	}
}

// Flatten merges API response onto existing NIOS model.
func (m *NIOSRecordCaaModel) Flatten(ctx context.Context, from *coremodel.NIOSRecordCaaExt, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	planExtAttrs := m.ExtAttrs
	if planExtAttrs.IsUnknown() {
		planExtAttrs = types.MapNull(types.StringType)
	}
	m.CaFlag = flex.FlattenInt64Pointer(from.CaFlag)
	m.CaTag = flex.FlattenStringPointerEmptyAsNull(from.CaTag)
	m.CaValue = flex.FlattenStringPointerEmptyAsNull(from.CaValue)
	m.Comment = flex.FlattenStringPointerEmptyAsNull(from.Comment)
	m.Creator = flex.FlattenStringPointerEmptyAsNull(from.Creator)
	m.DdnsPrincipal = flex.FlattenStringPointerEmptyAsNull(from.DdnsPrincipal)
	m.DdnsProtected = flex.FlattenBoolPointer(from.DdnsProtected)
	m.Disable = flex.FlattenBoolPointer(from.Disable)
	m.ExtAttrs, m.ExtAttrsAll = flex.FlattenEAs(planExtAttrs, from.ExtAttrs)
	m.ForbidReclamation = flex.FlattenBoolPointer(from.ForbidReclamation)
	m.Name = flex.FlattenStringPointerEmptyAsNull(from.Name)
	m.Ttl = flex.FlattenInt64Pointer(from.Ttl)
	m.View = flex.FlattenStringPointerEmptyAsNull(from.View)
}

// Flatten merges API response onto existing UDDI model.
func (m *UDDIRecordCaaModel) Flatten(ctx context.Context, from *coremodel.UDDIRecordCaaExt, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	m.AbsoluteNameSpec = flex.FlattenStringPointer(from.AbsoluteNameSpec)
	m.Comment = flex.FlattenStringPointer(from.Comment)
	m.Disabled = flex.FlattenBoolPointer(from.Disabled)
	m.InheritanceSources = FlattenRecordInheritance(ctx, from.InheritanceSources, diags)
	m.NameInZone = flex.FlattenStringPointer(from.NameInZone)
	m.Options = flex.FlattenMapStringAny(ctx, from.Options, diags)
	m.Rdata = flex.FlattenMapStringAny(ctx, from.Rdata, diags)
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
