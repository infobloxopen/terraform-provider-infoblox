package ipam

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/mapvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/mapdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectdefault"
	objectplanmodifier "github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"

	coremodel "github.com/infobloxopen/terraform-provider-infoblox/internal/core/model/ipam"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
	importmod "github.com/infobloxopen/terraform-provider-infoblox/internal/planmodifiers/import"
	customvalidator "github.com/infobloxopen/terraform-provider-infoblox/internal/validator"
)

type NetworkviewModel struct {
	Id   types.String `tfsdk:"id"`
	NIOS types.Object `tfsdk:"nios"`
	UDDI types.Object `tfsdk:"uddi"`
}

var NetworkviewAttrTypes = map[string]attr.Type{
	"id":   types.StringType,
	"nios": types.ObjectType{AttrTypes: NIOSNetworkviewAttrTypes},
	"uddi": types.ObjectType{AttrTypes: UDDINetworkviewAttrTypes},
}

type NIOSNetworkviewModel struct {
	CloudInfo            types.Object `tfsdk:"cloud_info"`
	Comment              types.String `tfsdk:"comment"`
	DdnsDnsView          types.String `tfsdk:"ddns_dns_view"`
	DdnsZonePrimaries    types.List   `tfsdk:"ddns_zone_primaries"`
	ExtAttrs             types.Map    `tfsdk:"ext_attrs"`
	ExtAttrsAll          types.Map    `tfsdk:"ext_attrs_all"`
	FederatedRealms      types.List   `tfsdk:"federated_realms"`
	InternalForwardZones types.List   `tfsdk:"internal_forward_zones"`
	MgmPrivate           types.Bool   `tfsdk:"mgm_private"`
	Name                 types.String `tfsdk:"name"`
	RemoteForwardZones   types.List   `tfsdk:"remote_forward_zones"`
	RemoteReverseZones   types.List   `tfsdk:"remote_reverse_zones"`
}

var NIOSNetworkviewAttrTypes = map[string]attr.Type{
	"cloud_info":             types.ObjectType{AttrTypes: NetworkviewCloudInfoAttrTypes},
	"comment":                types.StringType,
	"ddns_dns_view":          types.StringType,
	"ddns_zone_primaries":    types.ListType{ElemType: types.ObjectType{AttrTypes: NetworkviewDdnsZonePrimariesAttrTypes}},
	"ext_attrs":              types.MapType{ElemType: types.StringType},
	"ext_attrs_all":          types.MapType{ElemType: types.StringType},
	"federated_realms":       types.ListType{ElemType: types.ObjectType{AttrTypes: NetworkviewFederatedRealmsAttrTypes}},
	"internal_forward_zones": types.ListType{ElemType: types.StringType},
	"mgm_private":            types.BoolType,
	"name":                   types.StringType,
	"remote_forward_zones":   types.ListType{ElemType: types.ObjectType{AttrTypes: NetworkviewRemoteForwardZonesAttrTypes}},
	"remote_reverse_zones":   types.ListType{ElemType: types.ObjectType{AttrTypes: NetworkviewRemoteReverseZonesAttrTypes}},
}

type UDDINetworkviewModel struct {
	AsmConfig                       types.Object  `tfsdk:"asm_config"`
	Comment                         types.String  `tfsdk:"comment"`
	CompartmentId                   types.String  `tfsdk:"compartment_id"`
	DdnsClientUpdate                types.String  `tfsdk:"ddns_client_update"`
	DdnsConflictResolutionMode      types.String  `tfsdk:"ddns_conflict_resolution_mode"`
	DdnsDomain                      types.String  `tfsdk:"ddns_domain"`
	DdnsGenerateName                types.Bool    `tfsdk:"ddns_generate_name"`
	DdnsGeneratedPrefix             types.String  `tfsdk:"ddns_generated_prefix"`
	DdnsSendUpdates                 types.Bool    `tfsdk:"ddns_send_updates"`
	DdnsTtlPercent                  types.Float64 `tfsdk:"ddns_ttl_percent"`
	DdnsUpdateOnRenew               types.Bool    `tfsdk:"ddns_update_on_renew"`
	DdnsUseConflictResolution       types.Bool    `tfsdk:"ddns_use_conflict_resolution"`
	DefaultRealms                   types.List    `tfsdk:"default_realms"`
	DhcpConfig                      types.Object  `tfsdk:"dhcp_config"`
	DhcpOptions                     types.List    `tfsdk:"dhcp_options"`
	DhcpOptionsV6                   types.List    `tfsdk:"dhcp_options_v6"`
	HeaderOptionFilename            types.String  `tfsdk:"header_option_filename"`
	HeaderOptionServerAddress       types.String  `tfsdk:"header_option_server_address"`
	HeaderOptionServerName          types.String  `tfsdk:"header_option_server_name"`
	HostnameRewriteChar             types.String  `tfsdk:"hostname_rewrite_char"`
	HostnameRewriteEnabled          types.Bool    `tfsdk:"hostname_rewrite_enabled"`
	HostnameRewriteRegex            types.String  `tfsdk:"hostname_rewrite_regex"`
	InheritanceSources              types.Object  `tfsdk:"inheritance_sources"`
	Name                            types.String  `tfsdk:"name"`
	Tags                            types.Map     `tfsdk:"tags"`
	TagsAll                         types.Map     `tfsdk:"tags_all"`
	VendorSpecificOptionOptionSpace types.String  `tfsdk:"vendor_specific_option_option_space"`
}

var UDDINetworkviewAttrTypes = map[string]attr.Type{
	"asm_config":                          types.ObjectType{AttrTypes: ASMConfigAttrTypes},
	"comment":                             types.StringType,
	"compartment_id":                      types.StringType,
	"ddns_client_update":                  types.StringType,
	"ddns_conflict_resolution_mode":       types.StringType,
	"ddns_domain":                         types.StringType,
	"ddns_generate_name":                  types.BoolType,
	"ddns_generated_prefix":               types.StringType,
	"ddns_send_updates":                   types.BoolType,
	"ddns_ttl_percent":                    types.Float64Type,
	"ddns_update_on_renew":                types.BoolType,
	"ddns_use_conflict_resolution":        types.BoolType,
	"default_realms":                      types.ListType{ElemType: types.StringType},
	"dhcp_config":                         types.ObjectType{AttrTypes: DHCPConfigAttrTypes},
	"dhcp_options":                        types.ListType{ElemType: types.ObjectType{AttrTypes: OptionItemAttrTypes}},
	"dhcp_options_v6":                     types.ListType{ElemType: types.ObjectType{AttrTypes: OptionItemAttrTypes}},
	"header_option_filename":              types.StringType,
	"header_option_server_address":        types.StringType,
	"header_option_server_name":           types.StringType,
	"hostname_rewrite_char":               types.StringType,
	"hostname_rewrite_enabled":            types.BoolType,
	"hostname_rewrite_regex":              types.StringType,
	"inheritance_sources":                 types.ObjectType{AttrTypes: IPSpaceInheritanceAttrTypes},
	"name":                                types.StringType,
	"tags":                                types.MapType{ElemType: types.StringType},
	"tags_all":                            types.MapType{ElemType: types.StringType},
	"vendor_specific_option_option_space": types.StringType,
}

const (
	NetworkviewInheritanceType = "full"
	NetworkviewReturnFields    = "associated_dns_views,associated_members,cloud_info,comment,ddns_dns_view,ddns_zone_primaries,extattrs,federated_realms,internal_forward_zones,is_default,mgm_private,ms_ad_user_data,name,remote_forward_zones,remote_reverse_zones"
)

var NetworkviewResourceSchemaAttributes = map[string]schema.Attribute{
	"id": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The reference to the object.",
	},
	"nios": schema.SingleNestedAttribute{
		Optional:            true,
		MarkdownDescription: "NIOS backend-specific fields.",
		Attributes:          NetworkviewResourceNiosSchemaAttributes,
	},
	"uddi": schema.SingleNestedAttribute{
		Optional:            true,
		MarkdownDescription: "UDDI backend-specific fields.",
		Attributes:          NetworkviewResourceUddiSchemaAttributes,
	},
}

var NetworkviewResourceNiosSchemaAttributes = map[string]schema.Attribute{
	"cloud_info": schema.SingleNestedAttribute{
		Attributes:          NetworkviewCloudInfoResourceSchemaAttributes,
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "",
	},
	"comment": schema.StringAttribute{
		Optional: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
			customvalidator.ValidateTrimmedString(),
		},
		MarkdownDescription: "Comment for the network view; maximum 256 characters.",
	},
	"ddns_dns_view": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
		},
		MarkdownDescription: "DNS views that will receive the updates if you enable the appliance to send updates to Grid members.",
	},
	"ddns_zone_primaries": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: NetworkviewDdnsZonePrimariesResourceSchemaAttributes,
		},
		Optional: true,
		Computed: true,
		Validators: []validator.List{
			customvalidator.ListNotEmpty(),
		},
		MarkdownDescription: "An array of Ddns Zone Primary dhcpddns structs that lists the information of primary zone to wich DDNS updates should be sent.",
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
	"federated_realms": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: NetworkviewFederatedRealmsResourceSchemaAttributes,
		},
		Optional: true,
		Computed: true,
		Validators: []validator.List{
			customvalidator.ListNotEmpty(),
		},
		MarkdownDescription: "This field contains the federated realms associated to this network view",
	},
	"internal_forward_zones": schema.ListAttribute{
		ElementType: types.StringType,
		Optional:    true,
		Computed:    true,
		Default:     listdefault.StaticValue(types.ListNull(types.StringType)),
		Validators: []validator.List{
			customvalidator.ListNotEmpty(),
		},
		MarkdownDescription: "The list of linked authoritative DNS zones.",
	},
	"mgm_private": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "This field controls whether this object is synchronized with the Multi-Grid Master. If this field is set to True, objects are not synchronized.",
	},
	"name": schema.StringAttribute{
		Required: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
			customvalidator.ValidateTrimmedString(),
		},
		MarkdownDescription: "Name of the network view.",
	},
	"remote_forward_zones": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: NetworkviewRemoteForwardZonesResourceSchemaAttributes,
		},
		Optional: true,
		Computed: true,
		Validators: []validator.List{
			customvalidator.ListNotEmpty(),
		},
		MarkdownDescription: "The list of forward-mapping zones to which the DHCP server sends the updates.",
	},
	"remote_reverse_zones": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: NetworkviewRemoteReverseZonesResourceSchemaAttributes,
		},
		Optional: true,
		Computed: true,
		Validators: []validator.List{
			customvalidator.ListNotEmpty(),
		},
		MarkdownDescription: "The list of reverse-mapping zones to which the DHCP server sends the updates.",
	},
}

var NetworkviewResourceUddiSchemaAttributes = map[string]schema.Attribute{
	"asm_config": schema.SingleNestedAttribute{
		Attributes: ASMConfigResourceSchemaAttributes,
		Optional:   true,
		Computed:   true,
		Default: objectdefault.StaticValue(types.ObjectValueMust(ASMConfigAttrTypes, map[string]attr.Value{
			"asm_threshold":       types.Int64Value(90),
			"enable":              types.BoolValue(true),
			"enable_notification": types.BoolValue(true),
			"forecast_period":     types.Int64Value(14),
			"growth_factor":       types.Int64Value(20),
			"growth_type":         types.StringValue("percent"),
			"history":             types.Int64Value(30),
			"min_total":           types.Int64Value(10),
			"min_unused":          types.Int64Value(10),
			"reenable_date":       timetypes.NewRFC3339ValueMust("1970-01-01T00:00:00Z"),
		})),
		MarkdownDescription: "The __ASMConfig__ object represents Automated Scope Management configuration.",
	},
	"comment": schema.StringAttribute{
		Default:             stringdefault.StaticString(""),
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The description for the IP space. May contain 0 to 1024 characters. Can include UTF-8.",
	},
	"compartment_id": schema.StringAttribute{
		Default:             stringdefault.StaticString(""),
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The compartment associated with the object. If no compartment is associated with the object, the value defaults to empty.",
	},
	"ddns_client_update": schema.StringAttribute{
		Default: stringdefault.StaticString("client"),
		Validators: []validator.String{
			stringvalidator.OneOf("client", "server", "ignore", "over_client_update", "over_no_update"),
		},
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "Controls who does the DDNS updates.  Valid values are: * _client_: DHCP server updates DNS if requested by client. * _server_: DHCP server always updates DNS, overriding an update request from the client, unless the client requests no updates. * _ignore_: DHCP server always updates DNS, even if the client says not to. * _over_client_update_: Same as _server_. DHCP server always updates DNS, overriding an update request from the client, unless the client requests no updates. * _over_no_update_: DHCP server updates DNS even if the client requests that no updates be done. If the client requests to do the update, DHCP server allows it.  Defaults to _client_.",
	},
	"ddns_conflict_resolution_mode": schema.StringAttribute{
		Default: stringdefault.StaticString("check_with_dhcid"),
		Validators: []validator.String{
			stringvalidator.OneOf("check_with_dhcid", "no_check_with_dhcid", "check_exists_with_dhcid", "no_check_without_dhcid"),
		},
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The mode used for resolving conflicts while performing DDNS updates.  Valid values are: * _check_with_dhcid_: It includes adding a DHCID record and checking that record via conflict detection as per RFC 4703. * _no_check_with_dhcid_: This will ignore conflict detection but add a DHCID record when creating/updating an entry. * _check_exists_with_dhcid_: This will check if there is an existing DHCID record but does not verify the value of the record matches the update. This will also update the DHCID record for the entry. * _no_check_without_dhcid_: This ignores conflict detection and will not add a DHCID record when creating/updating a DDNS entry.  Defaults to _check_with_dhcid_.",
	},
	"ddns_domain": schema.StringAttribute{
		Default:             stringdefault.StaticString(""),
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The domain suffix for DDNS updates. FQDN, may be empty.  Defaults to empty.",
	},
	"ddns_generate_name": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Indicates if DDNS needs to generate a hostname when not supplied by the client.  Defaults to _false_.",
	},
	"ddns_generated_prefix": schema.StringAttribute{
		Default:             stringdefault.StaticString("myhost"),
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The prefix used in the generation of an FQDN.  When generating a name, DHCP server will construct the name in the format: [ddns-generated-prefix]-[address-text].[ddns-qualifying-suffix]. where address-text is simply the lease IP address converted to a hyphenated string.  Defaults to \"myhost\".",
	},
	"ddns_send_updates": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(true),
		MarkdownDescription: "Determines if DDNS updates are enabled at the IP space level. Defaults to _true_.",
	},
	"ddns_ttl_percent": schema.Float64Attribute{
		Optional:            true,
		MarkdownDescription: "DDNS TTL value - to be calculated as a simple percentage of the lease's lifetime, using the parameter's value as the percentage. It is specified as a percentage (e.g. 25, 75). Defaults to unspecified.",
	},
	"ddns_update_on_renew": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Instructs the DHCP server to always update the DNS information when a lease is renewed even if its DNS information has not changed.  Defaults to _false_.",
	},
	"ddns_use_conflict_resolution": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(true),
		MarkdownDescription: "When true, DHCP server will apply conflict resolution, as described in RFC 4703, when attempting to fulfill the update request.  When false, DHCP server will simply attempt to update the DNS entries per the request, regardless of whether or not they conflict with existing entries owned by other DHCP4 clients.  Defaults to _true_.",
	},
	"default_realms": schema.ListAttribute{
		ElementType: types.StringType,
		Optional:    true,
		Computed:    true,
		Validators: []validator.List{
			customvalidator.ListNotEmpty(),
		},
		MarkdownDescription: "Reserved for future use.",
	},
	"dhcp_config": schema.SingleNestedAttribute{
		Attributes: DHCPConfigResourceSchemaAttributes,
		Optional:   true,
		Computed:   true,
		Default: objectdefault.StaticValue(types.ObjectValueMust(DHCPConfigAttrTypes, map[string]attr.Value{
			"abandoned_reclaim_time":    types.Int64Value(3600),
			"abandoned_reclaim_time_v6": types.Int64Value(3600),
			"allow_unknown":             types.BoolValue(true),
			"allow_unknown_v6":          types.BoolValue(true),
			"echo_client_id":            types.BoolValue(true),
			"filters":                   types.ListNull(types.StringType),
			"filters_v6":                types.ListNull(types.StringType),
			"filters_large_selection":   types.ListNull(types.StringType),
			"ignore_client_uid":         types.BoolValue(false),
			"ignore_list":               types.ListNull(types.ObjectType{AttrTypes: IgnoreItemAttrTypes}),
			"lease_time":                types.Int64Value(3600),
			"lease_time_v6":             types.Int64Value(3600),
		})),
		MarkdownDescription: "A DHCP Config object (_dhcp/dhcp_config_) represents a shared DHCP configuration that controls how leases are issued.",
	},
	"dhcp_options": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: OptionItemResourceSchemaAttributes,
		},
		Optional: true,
		Validators: []validator.List{
			customvalidator.ListNotEmpty(),
		},
		MarkdownDescription: "The list of IPv4 DHCP options for IP space. May be either a specific option or a group of options.",
	},
	"dhcp_options_v6": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: OptionItemResourceSchemaAttributes,
		},
		Optional: true,
		Validators: []validator.List{
			customvalidator.ListNotEmpty(),
		},
		MarkdownDescription: "The list of IPv6 DHCP options for IP space. May be either a specific option or a group of options.",
	},
	"header_option_filename": schema.StringAttribute{
		Default:             stringdefault.StaticString(""),
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The configuration for header option filename field.",
	},
	"header_option_server_address": schema.StringAttribute{
		Default:             stringdefault.StaticString(""),
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The configuration for header option server address field.",
	},
	"header_option_server_name": schema.StringAttribute{
		Default:             stringdefault.StaticString(""),
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The configuration for header option server name field.",
	},
	"hostname_rewrite_char": schema.StringAttribute{
		Default:  stringdefault.StaticString("-"),
		Optional: true,
		Computed: true,
		Validators: []validator.String{
			stringvalidator.LengthAtMost(1),
		},
		MarkdownDescription: "The character to replace non-matching characters with, when hostname rewrite is enabled.  Any single ASCII character or no character if the invalid characters should be removed without replacement.  Defaults to \"-\".",
	},
	"hostname_rewrite_enabled": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Indicates if client supplied hostnames will be rewritten prior to DDNS update by replacing every character that does not match _hostname_rewrite_regex_ by _hostname_rewrite_char_.  Defaults to _false_.",
	},
	"hostname_rewrite_regex": schema.StringAttribute{
		Default:             stringdefault.StaticString("[^a-zA-Z0-9_.]"),
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The regex bracket expression to match valid characters.  Must begin with \"[\" and end with \"]\" and be a compilable POSIX regex.  Defaults to \"[^a-zA-Z0-9_.]\".",
	},
	"inheritance_sources": schema.SingleNestedAttribute{
		Attributes: IPSpaceInheritanceResourceSchemaAttributes,
		Optional:   true,
		Computed:   true,
		PlanModifiers: []planmodifier.Object{
			objectplanmodifier.UseStateForUnknown(),
		},
		MarkdownDescription: "The __IPSpaceInheritance__ object specifies how and which fields _IPSpace_ object inherits from the parent.",
	},
	"name": schema.StringAttribute{
		Required:            true,
		MarkdownDescription: "The name of the IP space. Must contain 1 to 256 characters. Can include UTF-8.",
	},
	"tags": schema.MapAttribute{
		Optional:    true,
		Computed:    true,
		ElementType: types.StringType,
		Default:     mapdefault.StaticValue(types.MapNull(types.StringType)),
		Validators: []validator.Map{
			mapvalidator.SizeAtLeast(1),
		},
		MarkdownDescription: "The tags for the IP space in JSON format.",
	},
	"tags_all": schema.MapAttribute{
		Computed:            true,
		ElementType:         types.StringType,
		MarkdownDescription: "All tags including inherited values.",
	},
	"vendor_specific_option_option_space": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The resource identifier.",
	},
}

// Expand converts the TF model to the infoblox core model
func (m *NetworkviewModel) Expand(ctx context.Context, diags *diag.Diagnostics, isCreate bool) *coremodel.Networkview {
	if m == nil {
		return nil
	}

	obj := &coremodel.Networkview{}

	// Expand NIOS nested attribute (returns nil if not present)
	niosModel := flex.ExpandNestedObject[NIOSNetworkviewModel](ctx, m.NIOS, diags)
	if niosModel != nil {
		obj.NIOS = niosModel.Expand(ctx, diags)
	}

	// Expand UDDI nested attribute (returns nil if not present)
	uddiModel := flex.ExpandNestedObject[UDDINetworkviewModel](ctx, m.UDDI, diags)
	if uddiModel != nil {
		obj.UDDI = uddiModel.Expand(ctx, diags)
	}

	return obj
}

// Expand converts the NIOS TF model to the core model.
func (m *NIOSNetworkviewModel) Expand(ctx context.Context, diags *diag.Diagnostics) *coremodel.NIOSNetworkviewExt {
	return &coremodel.NIOSNetworkviewExt{
		CloudInfo:            ExpandNetworkviewCloudInfo(ctx, m.CloudInfo, diags),
		Comment:              flex.ExpandStringPointerNullAsEmpty(m.Comment),
		DdnsDnsView:          flex.ExpandStringPointer(m.DdnsDnsView),
		DdnsZonePrimaries:    flex.ExpandFrameworkListNestedBlock(ctx, m.DdnsZonePrimaries, diags, ExpandNetworkviewDdnsZonePrimaries),
		ExtAttrs:             flex.ExpandMapStringAny(ctx, m.ExtAttrs, diags),
		FederatedRealms:      flex.ExpandFrameworkListNestedBlock(ctx, m.FederatedRealms, diags, ExpandNetworkviewFederatedRealms),
		InternalForwardZones: flex.ExpandFrameworkListString(ctx, m.InternalForwardZones, diags),
		MgmPrivate:           flex.ExpandBoolPointer(m.MgmPrivate),
		Name:                 flex.ExpandStringPointerNullAsEmpty(m.Name),
		RemoteForwardZones:   flex.ExpandFrameworkListNestedBlock(ctx, m.RemoteForwardZones, diags, ExpandNetworkviewRemoteForwardZones),
		RemoteReverseZones:   flex.ExpandFrameworkListNestedBlock(ctx, m.RemoteReverseZones, diags, ExpandNetworkviewRemoteReverseZones),
	}
}

// Expand converts the UDDI TF model to the core model.
func (m *UDDINetworkviewModel) Expand(ctx context.Context, diags *diag.Diagnostics) *coremodel.UDDINetworkviewExt {
	return &coremodel.UDDINetworkviewExt{
		AsmConfig:                       ExpandASMConfig(ctx, m.AsmConfig, diags),
		Comment:                         flex.ExpandStringPointer(m.Comment),
		CompartmentId:                   flex.ExpandStringPointer(m.CompartmentId),
		DdnsClientUpdate:                flex.ExpandStringPointer(m.DdnsClientUpdate),
		DdnsConflictResolutionMode:      flex.ExpandStringPointer(m.DdnsConflictResolutionMode),
		DdnsDomain:                      flex.ExpandStringPointer(m.DdnsDomain),
		DdnsGenerateName:                flex.ExpandBoolPointer(m.DdnsGenerateName),
		DdnsGeneratedPrefix:             flex.ExpandStringPointer(m.DdnsGeneratedPrefix),
		DdnsSendUpdates:                 flex.ExpandBoolPointer(m.DdnsSendUpdates),
		DdnsTtlPercent:                  flex.ExpandFloat32Pointer(m.DdnsTtlPercent),
		DdnsUpdateOnRenew:               flex.ExpandBoolPointer(m.DdnsUpdateOnRenew),
		DdnsUseConflictResolution:       flex.ExpandBoolPointer(m.DdnsUseConflictResolution),
		DefaultRealms:                   flex.ExpandFrameworkListString(ctx, m.DefaultRealms, diags),
		DhcpConfig:                      ExpandDHCPConfig(ctx, m.DhcpConfig, diags),
		DhcpOptions:                     flex.ExpandFrameworkListNestedBlock(ctx, m.DhcpOptions, diags, ExpandOptionItem),
		DhcpOptionsV6:                   flex.ExpandFrameworkListNestedBlock(ctx, m.DhcpOptionsV6, diags, ExpandOptionItem),
		HeaderOptionFilename:            flex.ExpandStringPointer(m.HeaderOptionFilename),
		HeaderOptionServerAddress:       flex.ExpandStringPointer(m.HeaderOptionServerAddress),
		HeaderOptionServerName:          flex.ExpandStringPointer(m.HeaderOptionServerName),
		HostnameRewriteChar:             flex.ExpandStringPointer(m.HostnameRewriteChar),
		HostnameRewriteEnabled:          flex.ExpandBoolPointer(m.HostnameRewriteEnabled),
		HostnameRewriteRegex:            flex.ExpandStringPointer(m.HostnameRewriteRegex),
		InheritanceSources:              ExpandIPSpaceInheritance(ctx, m.InheritanceSources, diags),
		Name:                            flex.ExpandString(m.Name),
		Tags:                            flex.ExpandMapStringAny(ctx, m.Tags, diags),
		VendorSpecificOptionOptionSpace: flex.ExpandStringPointer(m.VendorSpecificOptionOptionSpace),
	}
}

// Flatten populates the TF model from a core response.
func (m *NetworkviewModel) Flatten(ctx context.Context, resp *coremodel.Networkview, diags *diag.Diagnostics) {
	if resp == nil {
		return
	}

	m.Id = flex.FlattenStringPointer(resp.Id)

	// Extract existing NIOS model, flatten API response onto it, convert back
	niosModel := flex.ExpandNestedObject[NIOSNetworkviewModel](ctx, m.NIOS, diags)
	if niosModel == nil {
		niosModel = &NIOSNetworkviewModel{}
	}
	niosModel.Flatten(ctx, resp.NIOS, diags)
	if resp.NIOS != nil {
		m.NIOS = flex.FlattenNestedObject(ctx, niosModel, NIOSNetworkviewAttrTypes, diags)
	} else {
		m.NIOS = types.ObjectNull(NIOSNetworkviewAttrTypes)
	}

	// Extract existing UDDI model, flatten API response onto it, convert back
	uddiModel := flex.ExpandNestedObject[UDDINetworkviewModel](ctx, m.UDDI, diags)
	if uddiModel == nil {
		uddiModel = &UDDINetworkviewModel{}
	}
	uddiModel.Flatten(ctx, resp.UDDI, diags)
	if resp.UDDI != nil {
		m.UDDI = flex.FlattenNestedObject(ctx, uddiModel, UDDINetworkviewAttrTypes, diags)
	} else {
		m.UDDI = types.ObjectNull(UDDINetworkviewAttrTypes)
	}
}

// Flatten merges API response onto existing NIOS model.
func (m *NIOSNetworkviewModel) Flatten(ctx context.Context, from *coremodel.NIOSNetworkviewExt, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	planExtAttrs := m.ExtAttrs
	if planExtAttrs.IsUnknown() {
		planExtAttrs = types.MapNull(types.StringType)
	}
	m.CloudInfo = FlattenNetworkviewCloudInfo(ctx, from.CloudInfo, diags)
	m.Comment = flex.FlattenStringPointerEmptyAsNull(from.Comment)
	m.DdnsDnsView = flex.FlattenStringPointerEmptyAsNull(from.DdnsDnsView)
	m.DdnsZonePrimaries = flex.FlattenFrameworkListNestedBlock(ctx, from.DdnsZonePrimaries, NetworkviewDdnsZonePrimariesAttrTypes, diags, FlattenNetworkviewDdnsZonePrimaries)
	m.ExtAttrs, m.ExtAttrsAll = flex.FlattenEAs(planExtAttrs, from.ExtAttrs)
	m.FederatedRealms = flex.FlattenFrameworkListNestedBlock(ctx, from.FederatedRealms, NetworkviewFederatedRealmsAttrTypes, diags, FlattenNetworkviewFederatedRealms)
	m.InternalForwardZones = flex.FlattenFrameworkListString(ctx, from.InternalForwardZones, diags)
	m.MgmPrivate = flex.FlattenBoolPointer(from.MgmPrivate)
	m.Name = flex.FlattenStringPointerEmptyAsNull(from.Name)
	m.RemoteForwardZones = flex.FlattenFrameworkListNestedBlock(ctx, from.RemoteForwardZones, NetworkviewRemoteForwardZonesAttrTypes, diags, FlattenNetworkviewRemoteForwardZones)
	m.RemoteReverseZones = flex.FlattenFrameworkListNestedBlock(ctx, from.RemoteReverseZones, NetworkviewRemoteReverseZonesAttrTypes, diags, FlattenNetworkviewRemoteReverseZones)
}

// Flatten merges API response onto existing UDDI model.
func (m *UDDINetworkviewModel) Flatten(ctx context.Context, from *coremodel.UDDINetworkviewExt, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	m.AsmConfig = FlattenASMConfig(ctx, from.AsmConfig, diags)
	m.Comment = flex.FlattenStringPointer(from.Comment)
	m.CompartmentId = flex.FlattenStringPointer(from.CompartmentId)
	m.DdnsClientUpdate = flex.FlattenStringPointer(from.DdnsClientUpdate)
	m.DdnsConflictResolutionMode = flex.FlattenStringPointer(from.DdnsConflictResolutionMode)
	m.DdnsDomain = flex.FlattenStringPointer(from.DdnsDomain)
	m.DdnsGenerateName = flex.FlattenBoolPointer(from.DdnsGenerateName)
	m.DdnsGeneratedPrefix = flex.FlattenStringPointer(from.DdnsGeneratedPrefix)
	m.DdnsSendUpdates = flex.FlattenBoolPointer(from.DdnsSendUpdates)
	m.DdnsTtlPercent = flex.FlattenFloat32PointerZeroAsNull(from.DdnsTtlPercent)
	m.DdnsUpdateOnRenew = flex.FlattenBoolPointer(from.DdnsUpdateOnRenew)
	m.DdnsUseConflictResolution = flex.FlattenBoolPointer(from.DdnsUseConflictResolution)
	m.DefaultRealms = flex.FlattenFrameworkListString(ctx, from.DefaultRealms, diags)
	m.DhcpConfig = FlattenDHCPConfig(ctx, from.DhcpConfig, diags)
	m.DhcpOptions = flex.FlattenFrameworkListNestedBlock(ctx, from.DhcpOptions, OptionItemAttrTypes, diags, FlattenOptionItem)
	m.DhcpOptionsV6 = flex.FlattenFrameworkListNestedBlock(ctx, from.DhcpOptionsV6, OptionItemAttrTypes, diags, FlattenOptionItem)
	m.HeaderOptionFilename = flex.FlattenStringPointer(from.HeaderOptionFilename)
	m.HeaderOptionServerAddress = flex.FlattenStringPointer(from.HeaderOptionServerAddress)
	m.HeaderOptionServerName = flex.FlattenStringPointer(from.HeaderOptionServerName)
	m.HostnameRewriteChar = flex.FlattenStringPointer(from.HostnameRewriteChar)
	m.HostnameRewriteEnabled = flex.FlattenBoolPointer(from.HostnameRewriteEnabled)
	m.HostnameRewriteRegex = flex.FlattenStringPointer(from.HostnameRewriteRegex)
	m.InheritanceSources = FlattenIPSpaceInheritance(ctx, from.InheritanceSources, diags)
	m.Name = flex.FlattenString(from.Name)
	tagsAll := flex.FlattenMapStringAny(ctx, from.Tags, diags)
	if m.Tags.IsNull() || m.Tags.IsUnknown() {
		m.Tags = tagsAll
	}
	m.TagsAll = tagsAll
	m.VendorSpecificOptionOptionSpace = flex.FlattenStringPointer(from.VendorSpecificOptionOptionSpace)
}
