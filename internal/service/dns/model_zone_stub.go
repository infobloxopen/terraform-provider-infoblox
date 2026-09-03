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
	"github.com/hashicorp/terraform-plugin-framework/types"

	coremodel "github.com/infobloxopen/terraform-provider-infoblox/internal/core/model/dns"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
	immutable "github.com/infobloxopen/terraform-provider-infoblox/internal/planmodifiers/immutable"
	importmod "github.com/infobloxopen/terraform-provider-infoblox/internal/planmodifiers/import"
	internaltypes "github.com/infobloxopen/terraform-provider-infoblox/internal/types"
	customvalidator "github.com/infobloxopen/terraform-provider-infoblox/internal/validator"
)

type ZoneStubModel struct {
	Id   types.String `tfsdk:"id"`
	NIOS types.Object `tfsdk:"nios"`
}

var ZoneStubAttrTypes = map[string]attr.Type{
	"id":   types.StringType,
	"nios": types.ObjectType{AttrTypes: NIOSZoneStubAttrTypes},
}

type NIOSZoneStubModel struct {
	Comment           types.String                        `tfsdk:"comment"`
	Disable           types.Bool                          `tfsdk:"disable"`
	DisableForwarding types.Bool                          `tfsdk:"disable_forwarding"`
	ExtAttrs          types.Map                           `tfsdk:"ext_attrs"`
	ExtAttrsAll       types.Map                           `tfsdk:"ext_attrs_all"`
	ExternalNsGroup   types.String                        `tfsdk:"external_ns_group"`
	Fqdn              types.String                        `tfsdk:"fqdn"`
	Locked            types.Bool                          `tfsdk:"locked"`
	MsAdIntegrated    types.Bool                          `tfsdk:"ms_ad_integrated"`
	MsDdnsMode        types.String                        `tfsdk:"ms_ddns_mode"`
	NsGroup           types.String                        `tfsdk:"ns_group"`
	Prefix            internaltypes.CaseInsensitiveString `tfsdk:"prefix"`
	StubFrom          types.List                          `tfsdk:"stub_from"`
	StubMembers       types.List                          `tfsdk:"stub_members"`
	StubMsservers     types.List                          `tfsdk:"stub_msservers"`
	View              types.String                        `tfsdk:"view"`
	ZoneFormat        types.String                        `tfsdk:"zone_format"`
}

var NIOSZoneStubAttrTypes = map[string]attr.Type{
	"comment":            types.StringType,
	"disable":            types.BoolType,
	"disable_forwarding": types.BoolType,
	"ext_attrs":          types.MapType{ElemType: types.StringType},
	"ext_attrs_all":      types.MapType{ElemType: types.StringType},
	"external_ns_group":  types.StringType,
	"fqdn":               types.StringType,
	"locked":             types.BoolType,
	"ms_ad_integrated":   types.BoolType,
	"ms_ddns_mode":       types.StringType,
	"ns_group":           types.StringType,
	"prefix":             internaltypes.CaseInsensitiveStringType{},
	"stub_from":          types.ListType{ElemType: types.ObjectType{AttrTypes: ZoneStubStubFromAttrTypes}},
	"stub_members":       types.ListType{ElemType: types.ObjectType{AttrTypes: ZoneStubStubMembersAttrTypes}},
	"stub_msservers":     types.ListType{ElemType: types.ObjectType{AttrTypes: ZoneStubStubMsserversAttrTypes}},
	"view":               types.StringType,
	"zone_format":        types.StringType,
}

const (
	ZoneStubReturnFields = "address,comment,disable,disable_forwarding,display_domain,dns_fqdn,extattrs,external_ns_group,fqdn,locked,locked_by,mask_prefix,ms_ad_integrated,ms_ddns_mode,ms_managed,ms_read_only,ms_sync_master_name,ns_group,parent,prefix,soa_email,soa_expire,soa_mname,soa_negative_ttl,soa_refresh,soa_retry,soa_serial_number,stub_from,stub_members,stub_msservers,using_srg_associations,view,zone_format"
)

var ZoneStubResourceSchemaAttributes = map[string]schema.Attribute{
	"id": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The reference to the object.",
	},
	"nios": schema.SingleNestedAttribute{
		Optional:            true,
		MarkdownDescription: "NIOS backend-specific fields.",
		Attributes:          ZoneStubResourceNiosSchemaAttributes,
	},
}

var ZoneStubResourceNiosSchemaAttributes = map[string]schema.Attribute{
	"comment": schema.StringAttribute{
		Optional: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
			customvalidator.ValidateTrimmedString(),
		},
		MarkdownDescription: "Comment for the zone; maximum 256 characters.",
	},
	"disable": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Determines whether a zone is disabled or not. When this is set to False, the zone is enabled.",
	},
	"disable_forwarding": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Determines if the name servers that host the zone should not forward queries that end with the domain name of the zone to any configured forwarders.",
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
	"external_ns_group": schema.StringAttribute{
		Optional: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
			customvalidator.ValidateTrimmedString(),
		},
		MarkdownDescription: "A forward stub server name server group.",
	},
	"fqdn": schema.StringAttribute{
		Required: true,
		PlanModifiers: []planmodifier.String{
			immutable.ImmutableString(),
		},
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
			customvalidator.IsValidNIOSDomainName(),
			customvalidator.IsNotArpa(),
			stringvalidator.AlsoRequires(path.MatchRelative().AtParent().AtName("stub_from")),
		},
		MarkdownDescription: "The name of this DNS zone. For a reverse zone, this is in \"address/cidr\" format. For other zones, this is in FQDN format. This value can be in unicode format. Note that for a reverse zone, the corresponding zone_format value should be set.",
	},
	"locked": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "If you enable this flag, other administrators cannot make conflicting changes. This is for administration purposes only. The zone will continue to serve DNS data even when it is locked.",
	},
	"ms_ad_integrated": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "The flag that determines whether Active Directory is integrated or not. This field is valid only when ms_managed is \"STUB\", \"AUTH_PRIMARY\", or \"AUTH_BOTH\".",
	},
	"ms_ddns_mode": schema.StringAttribute{
		Default: stringdefault.StaticString("NONE"),
		Validators: []validator.String{
			stringvalidator.OneOf("ANY", "NONE", "SECURE"),
		},
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "Determines whether an Active Directory-integrated zone with a Microsoft DNS server as primary allows dynamic updates. Valid values are: \"SECURE\" if the zone allows secure updates only. \"NONE\" if the zone forbids dynamic updates. \"ANY\" if the zone accepts both secure and nonsecure updates. This field is valid only if ms_managed is either \"AUTH_PRIMARY\" or \"AUTH_BOTH\". If the flag ms_ad_integrated is false, the value \"SECURE\" is not allowed.",
	},
	"ns_group": schema.StringAttribute{
		Optional: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
			customvalidator.ValidateTrimmedString(),
		},
		MarkdownDescription: "A stub member name server group.",
	},
	"prefix": schema.StringAttribute{
		Optional:   true,
		CustomType: internaltypes.CaseInsensitiveStringType{},
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
			customvalidator.ValidateTrimmedString(),
		},
		MarkdownDescription: "The RFC2317 prefix value of this DNS zone. Use this field only when the netmask is greater than 24 bits; that is, for a mask between 25 and 31 bits. Enter a prefix, such as the name of the allocated address block. The prefix can be alphanumeric characters, such as 128/26 , 128-189 , or sub-B.",
	},
	"stub_from": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: ZoneStubStubFromResourceSchemaAttributes,
		},
		Optional: true,
		Validators: []validator.List{
			customvalidator.ListNotEmpty(),
		},
		MarkdownDescription: "The primary servers (masters) of this stub zone.",
	},
	"stub_members": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: ZoneStubStubMembersResourceSchemaAttributes,
		},
		Optional: true,
		Validators: []validator.List{
			customvalidator.ListNotEmpty(),
		},
		MarkdownDescription: "The Grid member servers of this stub zone. Note that the lead/stealth/grid_replicate/ preferred_primaries/override_preferred_primaries fields of the struct will be ignored when set in this field.",
	},
	"stub_msservers": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: ZoneStubStubMsserversResourceSchemaAttributes,
		},
		Optional: true,
		Validators: []validator.List{
			customvalidator.ListNotEmpty(),
		},
		MarkdownDescription: "The Microsoft DNS servers of this stub zone. Note that the stealth field of the struct will be ignored when set in this field.",
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
		MarkdownDescription: "The name of the DNS view in which the zone resides. Example \"external\".",
	},
	"zone_format": schema.StringAttribute{
		Default: stringdefault.StaticString("FORWARD"),
		Validators: []validator.String{
			stringvalidator.OneOf("FORWARD", "IPV4", "IPV6"),
		},
		Optional: true,
		Computed: true,
		PlanModifiers: []planmodifier.String{
			immutable.ImmutableString(),
		},
		MarkdownDescription: "Determines the format of this zone.",
	},
}

// Expand converts the TF model to the infoblox core model
func (m *ZoneStubModel) Expand(ctx context.Context, diags *diag.Diagnostics, isCreate bool) *coremodel.ZoneStub {
	if m == nil {
		return nil
	}

	obj := &coremodel.ZoneStub{}

	// Expand NIOS nested attribute (returns nil if not present)
	niosModel := flex.ExpandNestedObject[NIOSZoneStubModel](ctx, m.NIOS, diags)
	if niosModel != nil {
		obj.NIOS = niosModel.Expand(ctx, diags)
	}

	return obj
}

// Expand converts the NIOS TF model to the core model.
func (m *NIOSZoneStubModel) Expand(ctx context.Context, diags *diag.Diagnostics) *coremodel.NIOSZoneStubExt {
	return &coremodel.NIOSZoneStubExt{
		Comment:           flex.ExpandStringPointerNullAsEmpty(m.Comment),
		Disable:           flex.ExpandBoolPointer(m.Disable),
		DisableForwarding: flex.ExpandBoolPointer(m.DisableForwarding),
		ExtAttrs:          flex.ExpandMapStringAny(ctx, m.ExtAttrs, diags),
		ExternalNsGroup:   flex.ExpandStringPointerNullAsEmpty(m.ExternalNsGroup),
		Fqdn:              flex.ExpandStringPointerNullAsEmpty(m.Fqdn),
		Locked:            flex.ExpandBoolPointer(m.Locked),
		MsAdIntegrated:    flex.ExpandBoolPointer(m.MsAdIntegrated),
		MsDdnsMode:        flex.ExpandStringPointerNullAsEmpty(m.MsDdnsMode),
		NsGroup:           flex.ExpandStringPointerNullAsEmpty(m.NsGroup),
		Prefix:            flex.ExpandStringPointer(m.Prefix.StringValue),
		StubFrom:          flex.ExpandFrameworkListNestedBlock(ctx, m.StubFrom, diags, ExpandZoneStubStubFrom),
		StubMembers:       flex.ExpandFrameworkListNestedBlock(ctx, m.StubMembers, diags, ExpandZoneStubStubMembers),
		StubMsservers:     flex.ExpandFrameworkListNestedBlock(ctx, m.StubMsservers, diags, ExpandZoneStubStubMsservers),
		View:              flex.ExpandStringPointerNullAsEmpty(m.View),
		ZoneFormat:        flex.ExpandStringPointerNullAsEmpty(m.ZoneFormat),
	}
}

// Flatten populates the TF model from a core response.
func (m *ZoneStubModel) Flatten(ctx context.Context, resp *coremodel.ZoneStub, diags *diag.Diagnostics) {
	if resp == nil {
		return
	}

	m.Id = flex.FlattenStringPointer(resp.Id)

	// Extract existing NIOS model, flatten API response onto it, convert back
	niosModel := flex.ExpandNestedObject[NIOSZoneStubModel](ctx, m.NIOS, diags)
	if niosModel == nil {
		niosModel = &NIOSZoneStubModel{}
	}
	niosModel.Flatten(ctx, resp.NIOS, diags)
	if resp.NIOS != nil {
		m.NIOS = flex.FlattenNestedObject(ctx, niosModel, NIOSZoneStubAttrTypes, diags)
	} else {
		m.NIOS = types.ObjectNull(NIOSZoneStubAttrTypes)
	}

}

// Flatten merges API response onto existing NIOS model.
func (m *NIOSZoneStubModel) Flatten(ctx context.Context, from *coremodel.NIOSZoneStubExt, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	planExtAttrs := m.ExtAttrs
	if planExtAttrs.IsUnknown() {
		planExtAttrs = types.MapNull(types.StringType)
	}
	m.Comment = flex.FlattenStringPointerEmptyAsNull(from.Comment)
	m.Disable = flex.FlattenBoolPointer(from.Disable)
	m.DisableForwarding = flex.FlattenBoolPointer(from.DisableForwarding)
	m.ExtAttrs, m.ExtAttrsAll = flex.FlattenEAs(planExtAttrs, from.ExtAttrs)
	m.ExternalNsGroup = flex.FlattenStringPointerEmptyAsNull(from.ExternalNsGroup)
	m.Fqdn = flex.FlattenStringPointerEmptyAsNull(from.Fqdn)
	m.Locked = flex.FlattenBoolPointer(from.Locked)
	m.MsAdIntegrated = flex.FlattenBoolPointer(from.MsAdIntegrated)
	m.MsDdnsMode = flex.FlattenStringPointerEmptyAsNull(from.MsDdnsMode)
	m.NsGroup = flex.FlattenStringPointerEmptyAsNull(from.NsGroup)
	m.Prefix.StringValue = flex.FlattenStringPointer(from.Prefix)
	m.StubFrom = flex.FlattenFrameworkListNestedBlock(ctx, from.StubFrom, ZoneStubStubFromAttrTypes, diags, FlattenZoneStubStubFrom)
	m.StubMembers = flex.FlattenFrameworkListNestedBlock(ctx, from.StubMembers, ZoneStubStubMembersAttrTypes, diags, FlattenZoneStubStubMembers)
	m.StubMsservers = flex.FlattenFrameworkListNestedBlock(ctx, from.StubMsservers, ZoneStubStubMsserversAttrTypes, diags, FlattenZoneStubStubMsservers)
	m.View = flex.FlattenStringPointerEmptyAsNull(from.View)
	m.ZoneFormat = flex.FlattenStringPointerEmptyAsNull(from.ZoneFormat)
}
