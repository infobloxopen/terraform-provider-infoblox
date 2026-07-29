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
	"github.com/hashicorp/terraform-plugin-framework/types"

	coremodel "github.com/infobloxopen/terraform-provider-infoblox/internal/core/model/dns"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
	immutable "github.com/infobloxopen/terraform-provider-infoblox/internal/planmodifiers/immutable"
	customvalidator "github.com/infobloxopen/terraform-provider-infoblox/internal/validator"
)

type RecordNsModel struct {
	Id   types.String `tfsdk:"id"`
	NIOS types.Object `tfsdk:"nios"`
	UDDI types.Object `tfsdk:"uddi"`
}

var RecordNsAttrTypes = map[string]attr.Type{
	"id":   types.StringType,
	"nios": types.ObjectType{AttrTypes: NIOSRecordNsAttrTypes},
	"uddi": types.ObjectType{AttrTypes: UDDIRecordNsAttrTypes},
}

type NIOSRecordNsModel struct {
	Addresses        types.List   `tfsdk:"addresses"`
	MsDelegationName types.String `tfsdk:"ms_delegation_name"`
	Name             types.String `tfsdk:"name"`
	Nameserver       types.String `tfsdk:"nameserver"`
	View             types.String `tfsdk:"view"`
}

var NIOSRecordNsAttrTypes = map[string]attr.Type{
	"addresses":          types.ListType{ElemType: types.ObjectType{AttrTypes: RecordNsAddressesAttrTypes}},
	"ms_delegation_name": types.StringType,
	"name":               types.StringType,
	"nameserver":         types.StringType,
	"view":               types.StringType,
}

type UDDIRecordNsModel struct {
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

var UDDIRecordNsAttrTypes = map[string]attr.Type{
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
	RecordNsType            = "Ns"
	RecordNsInheritanceType = "full"
	RecordNsReturnFields    = "addresses,cloud_info,creator,dns_name,last_queried,ms_delegation_name,name,nameserver,policy,view,zone"
)

var RecordNsResourceSchemaAttributes = map[string]schema.Attribute{
	"id": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The reference to the object.",
	},
	"nios": schema.SingleNestedAttribute{
		Optional:            true,
		MarkdownDescription: "NIOS backend-specific fields.",
		Attributes:          RecordNsResourceNiosSchemaAttributes,
	},
	"uddi": schema.SingleNestedAttribute{
		Optional:            true,
		MarkdownDescription: "UDDI backend-specific fields.",
		Attributes:          RecordNsResourceUddiSchemaAttributes,
	},
}

var RecordNsResourceNiosSchemaAttributes = map[string]schema.Attribute{
	"addresses": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: RecordNsAddressesResourceSchemaAttributes,
		},
		Required: true,
		Validators: []validator.List{
			customvalidator.ListNotEmpty(),
		},
		MarkdownDescription: "The list of zone name servers.",
	},
	"ms_delegation_name": schema.StringAttribute{
		Optional: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
		},
		MarkdownDescription: "The MS delegation point name.",
	},
	"name": schema.StringAttribute{
		Required: true,
		PlanModifiers: []planmodifier.String{
			immutable.ImmutableString(),
		},
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
			customvalidator.IsValidDomainName(),
		},
		MarkdownDescription: "The name of the NS record in FQDN format. This value can be in unicode format.",
	},
	"nameserver": schema.StringAttribute{
		Required: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
			customvalidator.ValidateTrimmedString(),
		},
		MarkdownDescription: "The domain name of an authoritative server for the redirected zone.",
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

var RecordNsResourceUddiSchemaAttributes = map[string]schema.Attribute{
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
		MarkdownDescription: "The DNS resource record data in JSON format. Certain DNS resource record-specific subfields are required for creating the DNS resource record.  Subfields for _NS_ (Name Server) record:  Subfield | Description                         | Required ---------|-------------------------------------|--------- dname    | A domain-name which specifies a host which should be authoritative for the specified class and domain. Can be absolute or relative domain name and include UTF-8. <br><br> | Yes",
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
func (m *RecordNsModel) Expand(ctx context.Context, diags *diag.Diagnostics, isCreate bool) *coremodel.RecordNs {
	if m == nil {
		return nil
	}

	obj := &coremodel.RecordNs{}

	// Expand NIOS nested attribute (returns nil if not present)
	niosModel := flex.ExpandNestedObject[NIOSRecordNsModel](ctx, m.NIOS, diags)
	if niosModel != nil {
		obj.NIOS = niosModel.Expand(ctx, diags, isCreate)
	}

	// Expand UDDI nested attribute (returns nil if not present)
	uddiModel := flex.ExpandNestedObject[UDDIRecordNsModel](ctx, m.UDDI, diags)
	if uddiModel != nil {
		obj.UDDI = uddiModel.Expand(ctx, diags)
	}

	return obj
}

// Expand converts the NIOS TF model to the core model.
func (m *NIOSRecordNsModel) Expand(ctx context.Context, diags *diag.Diagnostics, isCreate bool) *coremodel.NIOSRecordNsExt {
	ext := &coremodel.NIOSRecordNsExt{
		Addresses:        flex.ExpandFrameworkListNestedBlock(ctx, m.Addresses, diags, ExpandRecordNsAddresses),
		MsDelegationName: flex.ExpandStringPointerNullAsEmpty(m.MsDelegationName),
		Nameserver:       flex.ExpandStringPointerNullAsEmpty(m.Nameserver),
	}
	if isCreate {
		ext.Name = flex.ExpandStringPointerNullAsEmpty(m.Name)
		ext.View = flex.ExpandStringPointerNullAsEmpty(m.View)
	}
	return ext
}

// Expand converts the UDDI TF model to the core model.
func (m *UDDIRecordNsModel) Expand(ctx context.Context, diags *diag.Diagnostics) *coremodel.UDDIRecordNsExt {
	return &coremodel.UDDIRecordNsExt{
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
func (m *RecordNsModel) Flatten(ctx context.Context, resp *coremodel.RecordNs, diags *diag.Diagnostics) {
	if resp == nil {
		return
	}

	m.Id = flex.FlattenStringPointer(resp.Id)

	// Extract existing NIOS model, flatten API response onto it, convert back
	niosModel := flex.ExpandNestedObject[NIOSRecordNsModel](ctx, m.NIOS, diags)
	if niosModel == nil {
		niosModel = &NIOSRecordNsModel{}
	}
	niosModel.Flatten(ctx, resp.NIOS, diags)
	if resp.NIOS != nil {
		m.NIOS = flex.FlattenNestedObject(ctx, niosModel, NIOSRecordNsAttrTypes, diags)
	} else {
		m.NIOS = types.ObjectNull(NIOSRecordNsAttrTypes)
	}

	// Extract existing UDDI model, flatten API response onto it, convert back
	uddiModel := flex.ExpandNestedObject[UDDIRecordNsModel](ctx, m.UDDI, diags)
	if uddiModel == nil {
		uddiModel = &UDDIRecordNsModel{}
	}
	uddiModel.Flatten(ctx, resp.UDDI, diags)
	if resp.UDDI != nil {
		m.UDDI = flex.FlattenNestedObject(ctx, uddiModel, UDDIRecordNsAttrTypes, diags)
	} else {
		m.UDDI = types.ObjectNull(UDDIRecordNsAttrTypes)
	}
}

// Flatten merges API response onto existing NIOS model.
func (m *NIOSRecordNsModel) Flatten(ctx context.Context, from *coremodel.NIOSRecordNsExt, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	m.Addresses = flex.FlattenFrameworkListNestedBlock(ctx, from.Addresses, RecordNsAddressesAttrTypes, diags, FlattenRecordNsAddresses)
	m.MsDelegationName = flex.FlattenStringPointerEmptyAsNull(from.MsDelegationName)
	m.Name = flex.FlattenStringPointerEmptyAsNull(from.Name)
	m.Nameserver = flex.FlattenStringPointerEmptyAsNull(from.Nameserver)
	m.View = flex.FlattenStringPointerEmptyAsNull(from.View)
}

// Flatten merges API response onto existing UDDI model.
func (m *UDDIRecordNsModel) Flatten(ctx context.Context, from *coremodel.UDDIRecordNsExt, diags *diag.Diagnostics) {
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
