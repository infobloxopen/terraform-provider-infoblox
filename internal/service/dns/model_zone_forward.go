package dns

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
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
	stringplanmodifier "github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"

	coremodel "github.com/infobloxopen/terraform-provider-infoblox/internal/core/model/dns"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
	immutable "github.com/infobloxopen/terraform-provider-infoblox/internal/planmodifiers/immutable"
	importmod "github.com/infobloxopen/terraform-provider-infoblox/internal/planmodifiers/import"
	internaltypes "github.com/infobloxopen/terraform-provider-infoblox/internal/types"
	customvalidator "github.com/infobloxopen/terraform-provider-infoblox/internal/validator"
)

type ZoneForwardModel struct {
	Id   types.String `tfsdk:"id"`
	NIOS types.Object `tfsdk:"nios"`
	UDDI types.Object `tfsdk:"uddi"`
}

var ZoneForwardAttrTypes = map[string]attr.Type{
	"id":   types.StringType,
	"nios": types.ObjectType{AttrTypes: NIOSZoneForwardAttrTypes},
	"uddi": types.ObjectType{AttrTypes: UDDIZoneForwardAttrTypes},
}

type NIOSZoneForwardModel struct {
	Comment             types.String                        `tfsdk:"comment"`
	Disable             types.Bool                          `tfsdk:"disable"`
	DisableNsGeneration types.Bool                          `tfsdk:"disable_ns_generation"`
	ExtAttrs            types.Map                           `tfsdk:"ext_attrs"`
	ExtAttrsAll         types.Map                           `tfsdk:"ext_attrs_all"`
	ExternalNsGroup     types.String                        `tfsdk:"external_ns_group"`
	ForwardTo           types.List                          `tfsdk:"forward_to"`
	ForwardersOnly      types.Bool                          `tfsdk:"forwarders_only"`
	ForwardingServers   types.List                          `tfsdk:"forwarding_servers"`
	Fqdn                types.String                        `tfsdk:"fqdn"`
	Locked              types.Bool                          `tfsdk:"locked"`
	MsAdIntegrated      types.Bool                          `tfsdk:"ms_ad_integrated"`
	MsDdnsMode          types.String                        `tfsdk:"ms_ddns_mode"`
	NsGroup             types.String                        `tfsdk:"ns_group"`
	Prefix              internaltypes.CaseInsensitiveString `tfsdk:"prefix"`
	View                types.String                        `tfsdk:"view"`
	ZoneFormat          types.String                        `tfsdk:"zone_format"`
}

var NIOSZoneForwardAttrTypes = map[string]attr.Type{
	"comment":               types.StringType,
	"disable":               types.BoolType,
	"disable_ns_generation": types.BoolType,
	"ext_attrs":             types.MapType{ElemType: types.StringType},
	"ext_attrs_all":         types.MapType{ElemType: types.StringType},
	"external_ns_group":     types.StringType,
	"forward_to":            types.ListType{ElemType: types.ObjectType{AttrTypes: ZoneForwardForwardToAttrTypes}},
	"forwarders_only":       types.BoolType,
	"forwarding_servers":    types.ListType{ElemType: types.ObjectType{AttrTypes: ZoneForwardForwardingServersAttrTypes}},
	"fqdn":                  types.StringType,
	"locked":                types.BoolType,
	"ms_ad_integrated":      types.BoolType,
	"ms_ddns_mode":          types.StringType,
	"ns_group":              types.StringType,
	"prefix":                internaltypes.CaseInsensitiveStringType{},
	"view":                  types.StringType,
	"zone_format":           types.StringType,
}

type UDDIZoneForwardModel struct {
	Comment            types.String `tfsdk:"comment"`
	CompartmentId      types.String `tfsdk:"compartment_id"`
	Disabled           types.Bool   `tfsdk:"disabled"`
	ExternalForwarders types.List   `tfsdk:"external_forwarders"`
	ForwardOnly        types.Bool   `tfsdk:"forward_only"`
	Fqdn               types.String `tfsdk:"fqdn"`
	Hosts              types.List   `tfsdk:"hosts"`
	InternalForwarders types.List   `tfsdk:"internal_forwarders"`
	Nsgs               types.List   `tfsdk:"nsgs"`
	Parent             types.String `tfsdk:"parent"`
	Tags               types.Map    `tfsdk:"tags"`
	TagsAll            types.Map    `tfsdk:"tags_all"`
	View               types.String `tfsdk:"view"`
}

var UDDIZoneForwardAttrTypes = map[string]attr.Type{
	"comment":             types.StringType,
	"compartment_id":      types.StringType,
	"disabled":            types.BoolType,
	"external_forwarders": types.ListType{ElemType: types.ObjectType{AttrTypes: ForwarderAttrTypes}},
	"forward_only":        types.BoolType,
	"fqdn":                types.StringType,
	"hosts":               types.ListType{ElemType: types.StringType},
	"internal_forwarders": types.ListType{ElemType: types.StringType},
	"nsgs":                types.ListType{ElemType: types.StringType},
	"parent":              types.StringType,
	"tags":                types.MapType{ElemType: types.StringType},
	"tags_all":            types.MapType{ElemType: types.StringType},
	"view":                types.StringType,
}

const (
	ZoneForwardReturnFields = "address,comment,disable,disable_ns_generation,display_domain,dns_fqdn,extattrs,external_ns_group,forward_to,forwarders_only,forwarding_servers,fqdn,locked,locked_by,mask_prefix,ms_ad_integrated,ms_ddns_mode,ms_managed,ms_read_only,ms_sync_master_name,ns_group,parent,prefix,using_srg_associations,view,zone_format"
)

var ZoneForwardResourceSchemaAttributes = map[string]schema.Attribute{
	"id": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The reference to the object.",
	},
	"nios": schema.SingleNestedAttribute{
		Optional:            true,
		MarkdownDescription: "NIOS backend-specific fields.",
		Attributes:          ZoneForwardResourceNiosSchemaAttributes,
	},
	"uddi": schema.SingleNestedAttribute{
		Optional:            true,
		MarkdownDescription: "UDDI backend-specific fields.",
		Attributes:          ZoneForwardResourceUddiSchemaAttributes,
	},
}

var ZoneForwardResourceNiosSchemaAttributes = map[string]schema.Attribute{
	"comment": schema.StringAttribute{
		Optional: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
			stringvalidator.LengthBetween(0, 256),
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
	"disable_ns_generation": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Determines whether a auto-generation of NS records in parent zone is disabled or not. When this is set to False, the auto-generation is enabled.",
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
			stringvalidator.ConflictsWith(path.MatchRelative().AtParent().AtName("forward_to")),
			customvalidator.ValidateTrimmedString(),
		},
		MarkdownDescription: "A forward stub server name server group.",
	},
	"forward_to": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: ZoneForwardForwardToResourceSchemaAttributes,
		},
		Optional: true,
		Validators: []validator.List{
			customvalidator.ListNotEmpty(),
			listvalidator.ConflictsWith(path.MatchRelative().AtParent().AtName("external_ns_group")),
		},
		MarkdownDescription: "The information for the remote name servers to which you want the Infoblox appliance to forward queries for a specified domain name.",
	},
	"forwarders_only": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Determines if the appliance sends queries to forwarders only, and not to other internal or Internet root servers.",
	},
	"forwarding_servers": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: ZoneForwardForwardingServersResourceSchemaAttributes,
		},
		Optional: true,
		Validators: []validator.List{
			customvalidator.ListNotEmpty(),
		},
		MarkdownDescription: "The information for the Grid members to which you want the Infoblox appliance to forward queries for a specified domain name.",
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
		MarkdownDescription: "A forwarding member name server group.",
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

var ZoneForwardResourceUddiSchemaAttributes = map[string]schema.Attribute{
	"comment": schema.StringAttribute{
		Default:             stringdefault.StaticString(""),
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "Optional. Comment for zone configuration.",
	},
	"compartment_id": schema.StringAttribute{
		Default:             stringdefault.StaticString(""),
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The access view associated with the object. If no access view is associated with the object, the value defaults to empty.",
	},
	"disabled": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Optional. _true_ to disable object. A disabled object is effectively non-existent when generating configuration.",
	},
	"external_forwarders": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: ForwarderResourceSchemaAttributes(false),
		},
		Optional: true,
		Validators: []validator.List{
			customvalidator.ListNotEmpty(),
		},
		MarkdownDescription: "Optional. External DNS servers to forward to. Order is not significant.",
	},
	"forward_only": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Optional. _true_ to only forward.",
	},
	"fqdn": schema.StringAttribute{
		Required: true,
		PlanModifiers: []planmodifier.String{
			stringplanmodifier.RequiresReplaceIfConfigured(),
		},
		Validators: []validator.String{
			customvalidator.IsValidUDDIDomainName(),
		},
		MarkdownDescription: "Zone FQDN. The FQDN supplied at creation will be converted to canonical form.  Read-only after creation.",
	},
	"hosts": schema.ListAttribute{
		ElementType: types.StringType,
		Optional:    true,
		Validators: []validator.List{
			customvalidator.ListNotEmpty(),
		},
		MarkdownDescription: "The resource identifier.",
	},
	"internal_forwarders": schema.ListAttribute{
		ElementType: types.StringType,
		Optional:    true,
		Validators: []validator.List{
			customvalidator.ListNotEmpty(),
		},
		MarkdownDescription: "The resource identifier.",
	},
	"nsgs": schema.ListAttribute{
		ElementType: types.StringType,
		Optional:    true,
		Validators: []validator.List{
			customvalidator.ListNotEmpty(),
		},
		MarkdownDescription: "The resource identifier.",
	},
	"parent": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The resource identifier.",
	},
	"tags": schema.MapAttribute{
		Optional:    true,
		Computed:    true,
		ElementType: types.StringType,
		Default:     mapdefault.StaticValue(types.MapNull(types.StringType)),
		Validators: []validator.Map{
			mapvalidator.SizeAtLeast(1),
		},
		MarkdownDescription: "Tagging specifics.",
	},
	"tags_all": schema.MapAttribute{
		Computed:            true,
		ElementType:         types.StringType,
		MarkdownDescription: "All tags including inherited values.",
	},
	"view": schema.StringAttribute{
		Optional: true,
		Computed: true,
		PlanModifiers: []planmodifier.String{
			stringplanmodifier.RequiresReplaceIfConfigured(),
		},
		MarkdownDescription: "The resource identifier.",
	},
}

// Expand converts the TF model to the infoblox core model
func (m *ZoneForwardModel) Expand(ctx context.Context, diags *diag.Diagnostics, isCreate bool) *coremodel.ZoneForward {
	if m == nil {
		return nil
	}

	obj := &coremodel.ZoneForward{}

	// Expand NIOS nested attribute (returns nil if not present)
	niosModel := flex.ExpandNestedObject[NIOSZoneForwardModel](ctx, m.NIOS, diags)
	if niosModel != nil {
		obj.NIOS = niosModel.Expand(ctx, diags, isCreate)
	}

	// Expand UDDI nested attribute (returns nil if not present)
	uddiModel := flex.ExpandNestedObject[UDDIZoneForwardModel](ctx, m.UDDI, diags)
	if uddiModel != nil {
		obj.UDDI = uddiModel.Expand(ctx, diags, isCreate)
	}

	return obj
}

// Expand converts the NIOS TF model to the core model.
func (m *NIOSZoneForwardModel) Expand(ctx context.Context, diags *diag.Diagnostics, isCreate bool) *coremodel.NIOSZoneForwardExt {
	ext := &coremodel.NIOSZoneForwardExt{
		Comment:             flex.ExpandStringPointerNullAsEmpty(m.Comment),
		Disable:             flex.ExpandBoolPointer(m.Disable),
		DisableNsGeneration: flex.ExpandBoolPointer(m.DisableNsGeneration),
		ExtAttrs:            flex.ExpandMapStringAny(ctx, m.ExtAttrs, diags),
		ExternalNsGroup:     flex.ExpandStringPointer(m.ExternalNsGroup),
		ForwardTo:           flex.ExpandFrameworkListNestedBlock(ctx, m.ForwardTo, diags, ExpandZoneForwardForwardTo),
		ForwardersOnly:      flex.ExpandBoolPointer(m.ForwardersOnly),
		ForwardingServers:   flex.ExpandFrameworkListNestedBlock(ctx, m.ForwardingServers, diags, ExpandZoneForwardForwardingServers),
		Locked:              flex.ExpandBoolPointer(m.Locked),
		MsAdIntegrated:      flex.ExpandBoolPointer(m.MsAdIntegrated),
		MsDdnsMode:          flex.ExpandStringPointerNullAsEmpty(m.MsDdnsMode),
		NsGroup:             flex.ExpandStringPointer(m.NsGroup),
		Prefix:              flex.ExpandStringPointer(m.Prefix.StringValue),
	}
	if isCreate {
		ext.Fqdn = flex.ExpandStringPointerNullAsEmpty(m.Fqdn)
		ext.View = flex.ExpandStringPointerNullAsEmpty(m.View)
		ext.ZoneFormat = flex.ExpandStringPointerNullAsEmpty(m.ZoneFormat)
	}
	return ext
}

// Expand converts the UDDI TF model to the core model.
func (m *UDDIZoneForwardModel) Expand(ctx context.Context, diags *diag.Diagnostics, isCreate bool) *coremodel.UDDIZoneForwardExt {
	ext := &coremodel.UDDIZoneForwardExt{
		Comment:            flex.ExpandStringPointer(m.Comment),
		CompartmentId:      flex.ExpandStringPointer(m.CompartmentId),
		Disabled:           flex.ExpandBoolPointer(m.Disabled),
		ExternalForwarders: flex.ExpandFrameworkListNestedBlock(ctx, m.ExternalForwarders, diags, ExpandForwarder),
		ForwardOnly:        flex.ExpandBoolPointer(m.ForwardOnly),
		Hosts:              flex.ExpandFrameworkListString(ctx, m.Hosts, diags),
		InternalForwarders: flex.ExpandFrameworkListString(ctx, m.InternalForwarders, diags),
		Nsgs:               flex.ExpandFrameworkListString(ctx, m.Nsgs, diags),
		Parent:             flex.ExpandStringPointer(m.Parent),
		Tags:               flex.ExpandMapStringAny(ctx, m.Tags, diags),
	}
	if isCreate {
		ext.Fqdn = flex.ExpandStringPointer(m.Fqdn)
		ext.View = flex.ExpandStringPointer(m.View)
	}
	return ext
}

// Flatten populates the TF model from a core response.
func (m *ZoneForwardModel) Flatten(ctx context.Context, resp *coremodel.ZoneForward, diags *diag.Diagnostics) {
	if resp == nil {
		return
	}

	m.Id = flex.FlattenStringPointer(resp.Id)

	// Extract existing NIOS model, flatten API response onto it, convert back
	niosModel := flex.ExpandNestedObject[NIOSZoneForwardModel](ctx, m.NIOS, diags)
	if niosModel == nil {
		niosModel = &NIOSZoneForwardModel{}
	}
	plannedNIOS := flex.ExpandNestedObject[NIOSZoneForwardModel](ctx, m.NIOS, diags)
	niosModel.Flatten(ctx, resp.NIOS, diags)
	if resp.NIOS != nil {
		PostFlattenZoneForwardNIOS(ctx, plannedNIOS, niosModel, diags)
		m.NIOS = flex.FlattenNestedObject(ctx, niosModel, NIOSZoneForwardAttrTypes, diags)
	} else {
		m.NIOS = types.ObjectNull(NIOSZoneForwardAttrTypes)
	}

	// Extract existing UDDI model, flatten API response onto it, convert back
	uddiModel := flex.ExpandNestedObject[UDDIZoneForwardModel](ctx, m.UDDI, diags)
	if uddiModel == nil {
		uddiModel = &UDDIZoneForwardModel{}
	}
	uddiModel.Flatten(ctx, resp.UDDI, diags)
	if resp.UDDI != nil {
		m.UDDI = flex.FlattenNestedObject(ctx, uddiModel, UDDIZoneForwardAttrTypes, diags)
	} else {
		m.UDDI = types.ObjectNull(UDDIZoneForwardAttrTypes)
	}
}

// Flatten merges API response onto existing NIOS model.
func (m *NIOSZoneForwardModel) Flatten(ctx context.Context, from *coremodel.NIOSZoneForwardExt, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	planExtAttrs := m.ExtAttrs
	if planExtAttrs.IsUnknown() {
		planExtAttrs = types.MapNull(types.StringType)
	}
	m.Comment = flex.FlattenStringPointerEmptyAsNull(from.Comment)
	m.Disable = flex.FlattenBoolPointer(from.Disable)
	m.DisableNsGeneration = flex.FlattenBoolPointer(from.DisableNsGeneration)
	m.ExtAttrs, m.ExtAttrsAll = flex.FlattenEAs(planExtAttrs, from.ExtAttrs)
	m.ExternalNsGroup = flex.FlattenStringPointerEmptyAsNull(from.ExternalNsGroup)
	m.ForwardTo = flex.FlattenFrameworkListNestedBlock(ctx, from.ForwardTo, ZoneForwardForwardToAttrTypes, diags, FlattenZoneForwardForwardTo)
	m.ForwardersOnly = flex.FlattenBoolPointer(from.ForwardersOnly)
	m.ForwardingServers = flex.FlattenFrameworkListNestedBlock(ctx, from.ForwardingServers, ZoneForwardForwardingServersAttrTypes, diags, FlattenZoneForwardForwardingServers)
	m.Fqdn = flex.FlattenStringPointerEmptyAsNull(from.Fqdn)
	m.Locked = flex.FlattenBoolPointer(from.Locked)
	m.MsAdIntegrated = flex.FlattenBoolPointer(from.MsAdIntegrated)
	m.MsDdnsMode = flex.FlattenStringPointerEmptyAsNull(from.MsDdnsMode)
	m.NsGroup = flex.FlattenStringPointerEmptyAsNull(from.NsGroup)
	m.Prefix.StringValue = flex.FlattenStringPointer(from.Prefix)
	m.View = flex.FlattenStringPointerEmptyAsNull(from.View)
	m.ZoneFormat = flex.FlattenStringPointerEmptyAsNull(from.ZoneFormat)
}

// Flatten merges API response onto existing UDDI model.
func (m *UDDIZoneForwardModel) Flatten(ctx context.Context, from *coremodel.UDDIZoneForwardExt, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	m.Comment = flex.FlattenStringPointer(from.Comment)
	m.CompartmentId = flex.FlattenStringPointer(from.CompartmentId)
	m.Disabled = flex.FlattenBoolPointer(from.Disabled)
	m.ExternalForwarders = flex.FlattenFrameworkListNestedBlock(ctx, from.ExternalForwarders, ForwarderAttrTypes, diags, FlattenForwarder)
	m.ForwardOnly = flex.FlattenBoolPointer(from.ForwardOnly)
	m.Fqdn = flex.FlattenStringPointer(from.Fqdn)
	m.Hosts = flex.FlattenFrameworkListString(ctx, from.Hosts, diags)
	m.InternalForwarders = flex.FlattenFrameworkListString(ctx, from.InternalForwarders, diags)
	m.Nsgs = flex.FlattenFrameworkListString(ctx, from.Nsgs, diags)
	m.Parent = flex.FlattenStringPointer(from.Parent)
	tagsAll := flex.FlattenMapStringAny(ctx, from.Tags, diags)
	if m.Tags.IsNull() || m.Tags.IsUnknown() {
		m.Tags = tagsAll
	}
	m.TagsAll = tagsAll
	m.View = flex.FlattenStringPointer(from.View)
}
