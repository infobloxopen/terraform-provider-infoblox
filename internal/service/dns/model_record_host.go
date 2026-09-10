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
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/mapdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"

	coremodel "github.com/infobloxopen/terraform-provider-infoblox/internal/core/model/dns"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
	immutable "github.com/infobloxopen/terraform-provider-infoblox/internal/planmodifiers/immutable"
	importmod "github.com/infobloxopen/terraform-provider-infoblox/internal/planmodifiers/import"
	internaltypes "github.com/infobloxopen/terraform-provider-infoblox/internal/types"
	customvalidator "github.com/infobloxopen/terraform-provider-infoblox/internal/validator"
)

type RecordHostModel struct {
	Id   types.String `tfsdk:"id"`
	NIOS types.Object `tfsdk:"nios"`
}

var RecordHostAttrTypes = map[string]attr.Type{
	"id":   types.StringType,
	"nios": types.ObjectType{AttrTypes: NIOSRecordHostAttrTypes},
}

type NIOSRecordHostModel struct {
	Aliases                  internaltypes.UnorderedListValue `tfsdk:"aliases"`
	CliCredentials           types.List                       `tfsdk:"cli_credentials"`
	Comment                  types.String                     `tfsdk:"comment"`
	ConfigureForDns          types.Bool                       `tfsdk:"configure_for_dns"`
	DdnsProtected            types.Bool                       `tfsdk:"ddns_protected"`
	DeviceDescription        types.String                     `tfsdk:"device_description"`
	DeviceLocation           types.String                     `tfsdk:"device_location"`
	DeviceType               types.String                     `tfsdk:"device_type"`
	DeviceVendor             types.String                     `tfsdk:"device_vendor"`
	Disable                  types.Bool                       `tfsdk:"disable"`
	DisableDiscovery         types.Bool                       `tfsdk:"disable_discovery"`
	EnableImmediateDiscovery types.Bool                       `tfsdk:"enable_immediate_discovery"`
	ExtAttrs                 types.Map                        `tfsdk:"ext_attrs"`
	ExtAttrsAll              types.Map                        `tfsdk:"ext_attrs_all"`
	Ipv4addrs                types.List                       `tfsdk:"ipv4addrs"`
	Ipv6addrs                types.List                       `tfsdk:"ipv6addrs"`
	Name                     types.String                     `tfsdk:"name"`
	NetworkView              types.String                     `tfsdk:"network_view"`
	RestartIfNeeded          types.Bool                       `tfsdk:"restart_if_needed"`
	RrsetOrder               types.String                     `tfsdk:"rrset_order"`
	Snmp3Credential          types.Object                     `tfsdk:"snmp3_credential"`
	SnmpCredential           types.Object                     `tfsdk:"snmp_credential"`
	Ttl                      types.Int64                      `tfsdk:"ttl"`
	UseDnsEaInheritance      types.Bool                       `tfsdk:"use_dns_ea_inheritance"`
	View                     types.String                     `tfsdk:"view"`
}

var NIOSRecordHostAttrTypes = map[string]attr.Type{
	"aliases":                    internaltypes.UnorderedListOfStringType,
	"cli_credentials":            types.ListType{ElemType: types.ObjectType{AttrTypes: RecordHostCliCredentialsAttrTypes}},
	"comment":                    types.StringType,
	"configure_for_dns":          types.BoolType,
	"ddns_protected":             types.BoolType,
	"device_description":         types.StringType,
	"device_location":            types.StringType,
	"device_type":                types.StringType,
	"device_vendor":              types.StringType,
	"disable":                    types.BoolType,
	"disable_discovery":          types.BoolType,
	"enable_immediate_discovery": types.BoolType,
	"ext_attrs":                  types.MapType{ElemType: types.StringType},
	"ext_attrs_all":              types.MapType{ElemType: types.StringType},
	"ipv4addrs":                  types.ListType{ElemType: types.ObjectType{AttrTypes: RecordHostIpv4addrAttrTypes}},
	"ipv6addrs":                  types.ListType{ElemType: types.ObjectType{AttrTypes: RecordHostIpv6addrAttrTypes}},
	"name":                       types.StringType,
	"network_view":               types.StringType,
	"restart_if_needed":          types.BoolType,
	"rrset_order":                types.StringType,
	"snmp3_credential":           types.ObjectType{AttrTypes: RecordHostSnmp3CredentialAttrTypes},
	"snmp_credential":            types.ObjectType{AttrTypes: RecordHostSnmpCredentialAttrTypes},
	"ttl":                        types.Int64Type,
	"use_dns_ea_inheritance":     types.BoolType,
	"view":                       types.StringType,
}

const (
	RecordHostReturnFields = "aliases,allow_telnet,cli_credentials,cloud_info,comment,configure_for_dns,creation_time,ddns_protected,device_description,device_location,device_type,device_vendor,disable,disable_discovery,dns_aliases,dns_name,extattrs,ipv4addrs,ipv6addrs,last_queried,ms_ad_user_data,name,network_view,rrset_order,snmp3_credential,snmp_credential,ttl,use_cli_credentials,use_dns_ea_inheritance,use_snmp3_credential,use_snmp_credential,use_ttl,view,zone"
)

var RecordHostResourceSchemaAttributes = map[string]schema.Attribute{
	"id": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The reference to the object.",
	},
	"nios": schema.SingleNestedAttribute{
		Optional:            true,
		MarkdownDescription: "NIOS backend-specific fields.",
		Attributes:          RecordHostResourceNiosSchemaAttributes,
	},
}

var RecordHostResourceNiosSchemaAttributes = map[string]schema.Attribute{
	"aliases": schema.ListAttribute{
		ElementType: types.StringType,
		Optional:    true,
		Computed:    true,
		CustomType:  internaltypes.UnorderedListOfStringType,
		Default:     listdefault.StaticValue(types.ListNull(types.StringType)),
		Validators: []validator.List{
			customvalidator.ListNotEmpty(),
		},
		MarkdownDescription: "This is a list of aliases for the host. The aliases must be in FQDN format. This value can be in unicode format.",
	},
	"cli_credentials": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: RecordHostCliCredentialsResourceSchemaAttributes,
		},
		Optional: true,
		Validators: []validator.List{
			customvalidator.ListNotEmpty(),
		},
		MarkdownDescription: "The CLI credentials for the host record.",
	},
	"comment": schema.StringAttribute{
		Optional: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
			customvalidator.ValidateTrimmedString(),
		},
		MarkdownDescription: "Comment for the record; maximum 256 characters.",
	},
	"configure_for_dns": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(true),
		MarkdownDescription: "When configure_for_dns is false, the host does not have parent zone information.",
	},
	"ddns_protected": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Determines if the DDNS updates for this record are allowed or not.",
	},
	"device_description": schema.StringAttribute{
		Optional: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
			customvalidator.ValidateTrimmedString(),
		},
		MarkdownDescription: "The description of the device.",
	},
	"device_location": schema.StringAttribute{
		Optional: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
			customvalidator.ValidateTrimmedString(),
		},
		MarkdownDescription: "The location of the device.",
	},
	"device_type": schema.StringAttribute{
		Optional: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
			customvalidator.ValidateTrimmedString(),
		},
		MarkdownDescription: "The type of the device.",
	},
	"device_vendor": schema.StringAttribute{
		Optional: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
			customvalidator.ValidateTrimmedString(),
		},
		MarkdownDescription: "The vendor of the device.",
	},
	"disable": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Determines if the record is disabled or not. False means that the record is enabled.",
	},
	"disable_discovery": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Determines if the discovery for the record is disabled or not. False means that the discovery is enabled.",
	},
	"enable_immediate_discovery": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Determines if the discovery for the record should be immediately enabled.",
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
	"ipv4addrs": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: RecordHostIpv4addrResourceSchemaAttributes,
		},
		Optional: true,
		Validators: []validator.List{
			customvalidator.ListNotEmpty(),
		},
		MarkdownDescription: "This is a list of IPv4 Addresses for the host.",
	},
	"ipv6addrs": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: RecordHostIpv6addrResourceSchemaAttributes,
		},
		Optional: true,
		Validators: []validator.List{
			customvalidator.ListNotEmpty(),
		},
		MarkdownDescription: "This is a list of IPv6 Addresses for the host.",
	},
	"name": schema.StringAttribute{
		Required: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
			customvalidator.IsValidNIOSDomainName(),
		},
		MarkdownDescription: "The host name in FQDN format This value can be in unicode format. Regular expression search is not supported for unicode values.",
	},
	"network_view": schema.StringAttribute{
		Default:  stringdefault.StaticString("default"),
		Optional: true,
		Computed: true,
		PlanModifiers: []planmodifier.String{
			immutable.ImmutableString(),
		},
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
		},
		MarkdownDescription: "The name of the network view in which the host record resides.",
	},
	"restart_if_needed": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Restarts the member service.",
	},
	"rrset_order": schema.StringAttribute{
		Default: stringdefault.StaticString("cyclic"),
		Validators: []validator.String{
			stringvalidator.OneOf("cyclic", "random", "fixed"),
		},
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The value of this field specifies the order in which resource record sets are returned. The possible values are \"cyclic\", \"random\" and \"fixed\".",
	},
	"snmp3_credential": schema.SingleNestedAttribute{
		Attributes:          RecordHostSnmp3CredentialResourceSchemaAttributes,
		Optional:            true,
		MarkdownDescription: "The SNMPv3 credential for this host record.",
	},
	"snmp_credential": schema.SingleNestedAttribute{
		Attributes:          RecordHostSnmpCredentialResourceSchemaAttributes,
		Optional:            true,
		MarkdownDescription: "The SNMP credential for this host record. If set to true, the SNMP credential will override member-level settings.",
	},
	"ttl": schema.Int64Attribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The Time To Live (TTL) value for record. A 32-bit unsigned integer that represents the duration, in seconds, for which the record is valid (cached). Zero indicates that the record should not be cached.",
	},
	"use_dns_ea_inheritance": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "When use_dns_ea_inheritance is True, the EA is inherited from associated zone.",
	},
	"view": schema.StringAttribute{
		Default:  stringdefault.StaticString("default"),
		Optional: true,
		Computed: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
		},
		MarkdownDescription: "The name of the DNS view in which the record resides. Example: \"external\".",
	},
}

// Expand converts the TF model to the infoblox core model
func (m *RecordHostModel) Expand(ctx context.Context, diags *diag.Diagnostics, isCreate bool) *coremodel.RecordHost {
	if m == nil {
		return nil
	}

	obj := &coremodel.RecordHost{}

	// Expand NIOS nested attribute (returns nil if not present)
	niosModel := flex.ExpandNestedObject[NIOSRecordHostModel](ctx, m.NIOS, diags)
	if niosModel != nil {
		obj.NIOS = niosModel.Expand(ctx, diags, isCreate)
		obj.NIOS = PostExpandRecordHostNIOS(ctx, obj.NIOS, diags)
	}

	return obj
}

// Expand converts the NIOS TF model to the core model.
func (m *NIOSRecordHostModel) Expand(ctx context.Context, diags *diag.Diagnostics, isCreate bool) *coremodel.NIOSRecordHostExt {
	ext := &coremodel.NIOSRecordHostExt{
		Aliases:                  flex.ExpandFrameworkListString(ctx, m.Aliases, diags),
		CliCredentials:           flex.ExpandFrameworkListNestedBlock(ctx, m.CliCredentials, diags, ExpandRecordHostCliCredentials),
		Comment:                  flex.ExpandStringPointerNullAsEmpty(m.Comment),
		ConfigureForDns:          flex.ExpandBoolPointer(m.ConfigureForDns),
		DdnsProtected:            flex.ExpandBoolPointer(m.DdnsProtected),
		DeviceDescription:        flex.ExpandStringPointerNullAsEmpty(m.DeviceDescription),
		DeviceLocation:           flex.ExpandStringPointerNullAsEmpty(m.DeviceLocation),
		DeviceType:               flex.ExpandStringPointerNullAsEmpty(m.DeviceType),
		DeviceVendor:             flex.ExpandStringPointerNullAsEmpty(m.DeviceVendor),
		Disable:                  flex.ExpandBoolPointer(m.Disable),
		DisableDiscovery:         flex.ExpandBoolPointer(m.DisableDiscovery),
		EnableImmediateDiscovery: flex.ExpandBoolPointer(m.EnableImmediateDiscovery),
		ExtAttrs:                 flex.ExpandMapStringAny(ctx, m.ExtAttrs, diags),
		Ipv4addrs:                flex.ExpandFrameworkListNestedBlock(ctx, m.Ipv4addrs, diags, ExpandRecordHostIpv4addr),
		Ipv6addrs:                flex.ExpandFrameworkListNestedBlock(ctx, m.Ipv6addrs, diags, ExpandRecordHostIpv6addr),
		Name:                     flex.ExpandStringPointerNullAsEmpty(m.Name),
		RestartIfNeeded:          flex.ExpandBoolPointer(m.RestartIfNeeded),
		RrsetOrder:               flex.ExpandStringPointerNullAsEmpty(m.RrsetOrder),
		Snmp3Credential:          ExpandRecordHostSnmp3Credential(ctx, m.Snmp3Credential, diags),
		SnmpCredential:           ExpandRecordHostSnmpCredential(ctx, m.SnmpCredential, diags),
		Ttl:                      flex.ExpandInt64Pointer(m.Ttl),
		UseDnsEaInheritance:      flex.ExpandBoolPointer(m.UseDnsEaInheritance),
		View:                     flex.ExpandStringPointerNullAsEmpty(m.View),
	}
	if isCreate {
		ext.NetworkView = flex.ExpandStringPointerNullAsEmpty(m.NetworkView)
	}
	return ext
}

// ApplyRecordHostNIOSUseFlags derives NIOS use flags from the raw config
// value(s) and writes them onto the core model. A flag is true when the user
// set any of its governed value fields in config.
func ApplyRecordHostNIOSUseFlags(ctx context.Context, config tfsdk.Config, obj *coremodel.RecordHost, diags *diag.Diagnostics) {
	if obj == nil || obj.NIOS == nil {
		return
	}
	obj.NIOS.UseCliCredentials = flex.DeriveUseFlag(ctx, config, diags, path.Root("nios").AtName("cli_credentials"))
	obj.NIOS.UseSnmp3Credential = flex.DeriveUseFlag(ctx, config, diags, path.Root("nios").AtName("snmp3_credential"))
	obj.NIOS.UseSnmpCredential = flex.DeriveUseFlag(ctx, config, diags, path.Root("nios").AtName("snmp_credential"))
	obj.NIOS.UseTtl = flex.DeriveUseFlag(ctx, config, diags, path.Root("nios").AtName("ttl"))
}

// Flatten populates the TF model from a core response.
func (m *RecordHostModel) Flatten(ctx context.Context, resp *coremodel.RecordHost, diags *diag.Diagnostics) {
	if resp == nil {
		return
	}

	m.Id = flex.FlattenStringPointer(resp.Id)

	// Extract existing NIOS model, flatten API response onto it, convert back
	niosModel := flex.ExpandNestedObject[NIOSRecordHostModel](ctx, m.NIOS, diags)
	if niosModel == nil {
		niosModel = &NIOSRecordHostModel{}
	}
	plannedNIOS := flex.ExpandNestedObject[NIOSRecordHostModel](ctx, m.NIOS, diags)
	niosModel.Flatten(ctx, resp.NIOS, diags)
	if resp.NIOS != nil {
		PostFlattenRecordHostNIOS(ctx, plannedNIOS, niosModel, diags)
		m.NIOS = flex.FlattenNestedObject(ctx, niosModel, NIOSRecordHostAttrTypes, diags)
	} else {
		m.NIOS = types.ObjectNull(NIOSRecordHostAttrTypes)
	}

}

// Flatten merges API response onto existing NIOS model.
func (m *NIOSRecordHostModel) Flatten(ctx context.Context, from *coremodel.NIOSRecordHostExt, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	planExtAttrs := m.ExtAttrs
	if planExtAttrs.IsUnknown() {
		planExtAttrs = types.MapNull(types.StringType)
	}
	m.Aliases = flex.FlattenFrameworkUnorderedListString(ctx, from.Aliases, diags)
	m.CliCredentials = flex.FlattenFrameworkListNestedBlock(ctx, from.CliCredentials, RecordHostCliCredentialsAttrTypes, diags, FlattenRecordHostCliCredentials)
	m.Comment = flex.FlattenStringPointerEmptyAsNull(from.Comment)
	m.ConfigureForDns = flex.FlattenBoolPointer(from.ConfigureForDns)
	m.DdnsProtected = flex.FlattenBoolPointer(from.DdnsProtected)
	m.DeviceDescription = flex.FlattenStringPointerEmptyAsNull(from.DeviceDescription)
	m.DeviceLocation = flex.FlattenStringPointerEmptyAsNull(from.DeviceLocation)
	m.DeviceType = flex.FlattenStringPointerEmptyAsNull(from.DeviceType)
	m.DeviceVendor = flex.FlattenStringPointerEmptyAsNull(from.DeviceVendor)
	m.Disable = flex.FlattenBoolPointer(from.Disable)
	m.DisableDiscovery = flex.FlattenBoolPointer(from.DisableDiscovery)
	m.ExtAttrs, m.ExtAttrsAll = flex.FlattenEAs(planExtAttrs, from.ExtAttrs)
	m.Ipv4addrs = flex.FlattenFrameworkListNestedBlock(ctx, from.Ipv4addrs, RecordHostIpv4addrAttrTypes, diags, FlattenRecordHostIpv4addr)
	m.Ipv6addrs = flex.FlattenFrameworkListNestedBlock(ctx, from.Ipv6addrs, RecordHostIpv6addrAttrTypes, diags, FlattenRecordHostIpv6addr)
	m.Name = flex.FlattenStringPointerEmptyAsNull(from.Name)
	m.NetworkView = flex.FlattenStringPointerEmptyAsNull(from.NetworkView)
	m.RrsetOrder = flex.FlattenStringPointerEmptyAsNull(from.RrsetOrder)
	m.Snmp3Credential = FlattenRecordHostSnmp3Credential(ctx, from.Snmp3Credential, diags)
	m.SnmpCredential = FlattenRecordHostSnmpCredential(ctx, from.SnmpCredential, diags)
	m.Ttl = flex.FlattenInt64Pointer(from.Ttl)
	m.UseDnsEaInheritance = flex.FlattenBoolPointer(from.UseDnsEaInheritance)
	m.View = flex.FlattenStringPointerEmptyAsNull(from.View)
}
