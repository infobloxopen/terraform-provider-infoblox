package dhcp

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/boolvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/mapvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/mapdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"

	coremodel "github.com/infobloxopen/terraform-provider-infoblox/internal/core/model/dhcp"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
	immutable "github.com/infobloxopen/terraform-provider-infoblox/internal/planmodifiers/immutable"
	importmod "github.com/infobloxopen/terraform-provider-infoblox/internal/planmodifiers/import"
	internaltypes "github.com/infobloxopen/terraform-provider-infoblox/internal/types"
	customvalidator "github.com/infobloxopen/terraform-provider-infoblox/internal/validator"
)

type SharednetworkModel struct {
	Id   types.String `tfsdk:"id"`
	NIOS types.Object `tfsdk:"nios"`
}

var SharednetworkAttrTypes = map[string]attr.Type{
	"id":   types.StringType,
	"nios": types.ObjectType{AttrTypes: NIOSSharednetworkAttrTypes},
}

type NIOSSharednetworkModel struct {
	Authority                   types.Bool                       `tfsdk:"authority"`
	Bootfile                    types.String                     `tfsdk:"bootfile"`
	Bootserver                  types.String                     `tfsdk:"bootserver"`
	Comment                     types.String                     `tfsdk:"comment"`
	DdnsGenerateHostname        types.Bool                       `tfsdk:"ddns_generate_hostname"`
	DdnsServerAlwaysUpdates     types.Bool                       `tfsdk:"ddns_server_always_updates"`
	DdnsTtl                     types.Int64                      `tfsdk:"ddns_ttl"`
	DdnsUpdateFixedAddresses    types.Bool                       `tfsdk:"ddns_update_fixed_addresses"`
	DdnsUseOption81             types.Bool                       `tfsdk:"ddns_use_option81"`
	DenyBootp                   types.Bool                       `tfsdk:"deny_bootp"`
	Disable                     types.Bool                       `tfsdk:"disable"`
	EnableDdns                  types.Bool                       `tfsdk:"enable_ddns"`
	EnablePxeLeaseTime          types.Bool                       `tfsdk:"enable_pxe_lease_time"`
	ExtAttrs                    types.Map                        `tfsdk:"ext_attrs"`
	ExtAttrsAll                 types.Map                        `tfsdk:"ext_attrs_all"`
	IgnoreClientIdentifier      types.Bool                       `tfsdk:"ignore_client_identifier"`
	IgnoreDhcpOptionListRequest types.Bool                       `tfsdk:"ignore_dhcp_option_list_request"`
	IgnoreId                    types.String                     `tfsdk:"ignore_id"`
	IgnoreMacAddresses          internaltypes.UnorderedListValue `tfsdk:"ignore_mac_addresses"`
	LeaseScavengeTime           types.Int64                      `tfsdk:"lease_scavenge_time"`
	LogicFilterRules            types.List                       `tfsdk:"logic_filter_rules"`
	Name                        types.String                     `tfsdk:"name"`
	NetworkView                 types.String                     `tfsdk:"network_view"`
	Networks                    types.List                       `tfsdk:"networks"`
	Nextserver                  types.String                     `tfsdk:"nextserver"`
	Options                     types.List                       `tfsdk:"options"`
	PxeLeaseTime                types.Int64                      `tfsdk:"pxe_lease_time"`
	UpdateDnsOnLeaseRenewal     types.Bool                       `tfsdk:"update_dns_on_lease_renewal"`
}

var NIOSSharednetworkAttrTypes = map[string]attr.Type{
	"authority":                       types.BoolType,
	"bootfile":                        types.StringType,
	"bootserver":                      types.StringType,
	"comment":                         types.StringType,
	"ddns_generate_hostname":          types.BoolType,
	"ddns_server_always_updates":      types.BoolType,
	"ddns_ttl":                        types.Int64Type,
	"ddns_update_fixed_addresses":     types.BoolType,
	"ddns_use_option81":               types.BoolType,
	"deny_bootp":                      types.BoolType,
	"disable":                         types.BoolType,
	"enable_ddns":                     types.BoolType,
	"enable_pxe_lease_time":           types.BoolType,
	"ext_attrs":                       types.MapType{ElemType: types.StringType},
	"ext_attrs_all":                   types.MapType{ElemType: types.StringType},
	"ignore_client_identifier":        types.BoolType,
	"ignore_dhcp_option_list_request": types.BoolType,
	"ignore_id":                       types.StringType,
	"ignore_mac_addresses":            internaltypes.UnorderedListOfStringType,
	"lease_scavenge_time":             types.Int64Type,
	"logic_filter_rules":              types.ListType{ElemType: types.ObjectType{AttrTypes: SharednetworkLogicFilterRulesAttrTypes}},
	"name":                            types.StringType,
	"network_view":                    types.StringType,
	"networks":                        types.ListType{ElemType: types.ObjectType{AttrTypes: SharednetworkNetworksAttrTypes}},
	"nextserver":                      types.StringType,
	"options":                         types.ListType{ElemType: types.ObjectType{AttrTypes: SharednetworkOptionsAttrTypes}},
	"pxe_lease_time":                  types.Int64Type,
	"update_dns_on_lease_renewal":     types.BoolType,
}

const (
	SharednetworkReturnFields = "authority,bootfile,bootserver,comment,ddns_generate_hostname,ddns_server_always_updates,ddns_ttl,ddns_update_fixed_addresses,ddns_use_option81,deny_bootp,dhcp_utilization,dhcp_utilization_status,disable,dynamic_hosts,enable_ddns,enable_pxe_lease_time,extattrs,ignore_client_identifier,ignore_dhcp_option_list_request,ignore_id,ignore_mac_addresses,lease_scavenge_time,logic_filter_rules,ms_ad_user_data,name,network_view,networks,nextserver,options,pxe_lease_time,static_hosts,total_hosts,update_dns_on_lease_renewal,use_authority,use_bootfile,use_bootserver,use_ddns_generate_hostname,use_ddns_ttl,use_ddns_update_fixed_addresses,use_ddns_use_option81,use_deny_bootp,use_enable_ddns,use_ignore_client_identifier,use_ignore_dhcp_option_list_request,use_ignore_id,use_lease_scavenge_time,use_logic_filter_rules,use_nextserver,use_options,use_pxe_lease_time,use_update_dns_on_lease_renewal"
)

var SharednetworkResourceSchemaAttributes = map[string]schema.Attribute{
	"id": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The reference to the object.",
	},
	"nios": schema.SingleNestedAttribute{
		Optional:            true,
		MarkdownDescription: "NIOS backend-specific fields.",
		Attributes:          SharednetworkResourceNiosSchemaAttributes,
	},
}

var SharednetworkResourceNiosSchemaAttributes = map[string]schema.Attribute{
	"authority": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Authority for the shared network.",
	},
	"bootfile": schema.StringAttribute{
		Optional: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
		},
		MarkdownDescription: "The bootfile name for the shared network. You can configure the DHCP server to support clients that use the boot file name option in their DHCPREQUEST messages.",
	},
	"bootserver": schema.StringAttribute{
		Optional: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
			customvalidator.IsValidIPv4OrFQDN(),
		},
		MarkdownDescription: "The bootserver address for the shared network. You can specify the name and/or IP address of the boot server that the host needs to boot. The boot server IPv4 Address or name in FQDN format.",
	},
	"comment": schema.StringAttribute{
		Optional: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
			customvalidator.ValidateTrimmedString(),
		},
		MarkdownDescription: "Comment for the shared network, maximum 256 characters.",
	},
	"ddns_generate_hostname": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "If this field is set to True, the DHCP server generates a hostname and updates DNS with it when the DHCP client request does not contain a hostname.",
	},
	"ddns_server_always_updates": schema.BoolAttribute{
		Optional: true,
		Computed: true,
		Default:  booldefault.StaticBool(true),
		Validators: []validator.Bool{
			boolvalidator.AlsoRequires(path.MatchRelative().AtParent().AtName("ddns_use_option81")),
		},
		MarkdownDescription: "This field controls whether only the DHCP server is allowed to update DNS, regardless of the DHCP clients requests. Note that changes for this field take effect only if ddns_use_option81 is True.",
	},
	"ddns_ttl": schema.Int64Attribute{
		Optional:            true,
		Computed:            true,
		Default:             int64default.StaticInt64(0),
		MarkdownDescription: "The DNS update Time to Live (TTL) value of a shared network object. The TTL is a 32-bit unsigned integer that represents the duration, in seconds, for which the update is cached. Zero indicates that the update is not cached.",
	},
	"ddns_update_fixed_addresses": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "By default, the DHCP server does not update DNS when it allocates a fixed address to a client. You can configure the DHCP server to update the A and PTR records of a client with a fixed address. When this feature is enabled and the DHCP server adds A and PTR records for a fixed address, the DHCP server never discards the records.",
	},
	"ddns_use_option81": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "The support for DHCP Option 81 at the shared network level.",
	},
	"deny_bootp": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "If set to true, BOOTP settings are disabled and BOOTP requests will be denied.",
	},
	"disable": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Determines whether a shared network is disabled or not. When this is set to False, the shared network is enabled.",
	},
	"enable_ddns": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "The dynamic DNS updates flag of a shared network object. If set to True, the DHCP server sends DDNS updates to DNS servers in the same Grid, and to external DNS servers.",
	},
	"enable_pxe_lease_time": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Set this to True if you want the DHCP server to use a different lease time for PXE clients.",
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
	"ignore_client_identifier": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "If set to true, the client identifier will be ignored.",
	},
	"ignore_dhcp_option_list_request": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "If this field is set to False, the appliance returns all DHCP options the client is eligible to receive, rather than only the list of options the client has requested.",
	},
	"ignore_id": schema.StringAttribute{
		Default: stringdefault.StaticString("NONE"),
		Validators: []validator.String{
			stringvalidator.OneOf("NONE", "MAC_ADDRESS", "CLIENT_ID"),
		},
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "Indicates whether the appliance will ignore DHCP client IDs or MAC addresses. Valid values are \"NONE\", \"CLIENT\", or \"MACADDR\". The default is \"NONE\".",
	},
	"ignore_mac_addresses": schema.ListAttribute{
		ElementType: types.StringType,
		Optional:    true,
		CustomType:  internaltypes.UnorderedListOfStringType,
		Validators: []validator.List{
			customvalidator.ListNotEmpty(),
		},
		MarkdownDescription: "A list of MAC addresses the appliance will ignore.",
	},
	"lease_scavenge_time": schema.Int64Attribute{
		Optional:            true,
		Computed:            true,
		Default:             int64default.StaticInt64(-1),
		MarkdownDescription: "An integer that specifies the period of time (in seconds) that frees and backs up leases remained in the database before they are automatically deleted. To disable lease scavenging, set the parameter to -1. The minimum positive value must be greater than 86400 seconds (1 day).",
	},
	"logic_filter_rules": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: SharednetworkLogicFilterRulesResourceSchemaAttributes,
		},
		Optional: true,
		Validators: []validator.List{
			customvalidator.ListNotEmpty(),
		},
		MarkdownDescription: "This field contains the logic filters to be applied on the this shared network. This list corresponds to the match rules that are written to the dhcpd configuration file.",
	},
	"name": schema.StringAttribute{
		Required: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
			customvalidator.ValidateTrimmedString(),
		},
		MarkdownDescription: "The name of the IPv6 Shared Network.",
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
			customvalidator.ValidateTrimmedString(),
		},
		MarkdownDescription: "The name of the network view in which this shared network resides.",
	},
	"networks": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: SharednetworkNetworksResourceSchemaAttributes,
		},
		Required: true,
		Validators: []validator.List{
			customvalidator.ListNotEmpty(),
		},
		MarkdownDescription: "A list of networks belonging to the shared network Each individual list item must be specified as an object containing a '_ref' parameter to a network reference, for example:: [{ \"_ref\": \"network/ZG5zLm5ldHdvcmskMTAuMwLvMTYvMA\", }] if the reference of the wanted network is not known, it is possible to specify search parameters for the network instead in the following way:: [{ \"_ref\": { 'network': '10.0.0.0/8', } }] note that in this case the search must match exactly one network for the assignment to be successful.",
	},
	"nextserver": schema.StringAttribute{
		Optional: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
			customvalidator.IsValidIPv4OrFQDN(),
		},
		MarkdownDescription: "The name in FQDN and/or IPv4 Address of the next server that the host needs to boot.",
	},
	"options": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: SharednetworkOptionsResourceSchemaAttributes,
		},
		Optional: true,
		Computed: true,
		Default:  listdefault.StaticValue(types.ListValueMust(types.ObjectType{AttrTypes: SharednetworkOptionsAttrTypes}, []attr.Value{})),
		Validators: []validator.List{
			customvalidator.ListNotEmpty(),
		},
		MarkdownDescription: "An array of DHCP option dhcpoption structs that lists the DHCP options associated with the object.",
	},
	"pxe_lease_time": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "The PXE lease time value of a shared network object. Some hosts use PXE (Preboot Execution Environment) to boot remotely from a server. To better manage your IP resources, set a different lease time for PXE boot requests. You can configure the DHCP server to allocate an IP address with a shorter lease time to hosts that send PXE boot requests, so IP addresses are not leased longer than necessary. A 32-bit unsigned integer that represents the duration, in seconds, for which the update is cached. Zero indicates that the update is not cached.",
	},
	"update_dns_on_lease_renewal": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "This field controls whether the DHCP server updates DNS when a DHCP lease is renewed.",
	},
}

// Expand converts the TF model to the infoblox core model
func (m *SharednetworkModel) Expand(ctx context.Context, diags *diag.Diagnostics, isCreate bool) *coremodel.Sharednetwork {
	if m == nil {
		return nil
	}

	obj := &coremodel.Sharednetwork{}

	// Expand NIOS nested attribute (returns nil if not present)
	niosModel := flex.ExpandNestedObject[NIOSSharednetworkModel](ctx, m.NIOS, diags)
	if niosModel != nil {
		obj.NIOS = niosModel.Expand(ctx, diags)
	}

	return obj
}

// Expand converts the NIOS TF model to the core model.
func (m *NIOSSharednetworkModel) Expand(ctx context.Context, diags *diag.Diagnostics) *coremodel.NIOSSharednetworkExt {
	return &coremodel.NIOSSharednetworkExt{
		Authority:                   flex.ExpandBoolPointer(m.Authority),
		Bootfile:                    flex.ExpandStringPointerNullAsEmpty(m.Bootfile),
		Bootserver:                  flex.ExpandStringPointerNullAsEmpty(m.Bootserver),
		Comment:                     flex.ExpandStringPointerNullAsEmpty(m.Comment),
		DdnsGenerateHostname:        flex.ExpandBoolPointer(m.DdnsGenerateHostname),
		DdnsServerAlwaysUpdates:     flex.ExpandBoolPointer(m.DdnsServerAlwaysUpdates),
		DdnsTtl:                     flex.ExpandInt64Pointer(m.DdnsTtl),
		DdnsUpdateFixedAddresses:    flex.ExpandBoolPointer(m.DdnsUpdateFixedAddresses),
		DdnsUseOption81:             flex.ExpandBoolPointer(m.DdnsUseOption81),
		DenyBootp:                   flex.ExpandBoolPointer(m.DenyBootp),
		Disable:                     flex.ExpandBoolPointer(m.Disable),
		EnableDdns:                  flex.ExpandBoolPointer(m.EnableDdns),
		EnablePxeLeaseTime:          flex.ExpandBoolPointer(m.EnablePxeLeaseTime),
		ExtAttrs:                    flex.ExpandMapStringAny(ctx, m.ExtAttrs, diags),
		IgnoreClientIdentifier:      flex.ExpandBoolPointer(m.IgnoreClientIdentifier),
		IgnoreDhcpOptionListRequest: flex.ExpandBoolPointer(m.IgnoreDhcpOptionListRequest),
		IgnoreId:                    flex.ExpandStringPointerNullAsEmpty(m.IgnoreId),
		IgnoreMacAddresses:          flex.ExpandFrameworkListString(ctx, m.IgnoreMacAddresses, diags),
		LeaseScavengeTime:           flex.ExpandInt64Pointer(m.LeaseScavengeTime),
		LogicFilterRules:            flex.ExpandFrameworkListNestedBlock(ctx, m.LogicFilterRules, diags, ExpandSharednetworkLogicFilterRules),
		Name:                        flex.ExpandStringPointerNullAsEmpty(m.Name),
		NetworkView:                 flex.ExpandStringPointerNullAsEmpty(m.NetworkView),
		Networks:                    flex.ExpandFrameworkListNestedBlock(ctx, m.Networks, diags, ExpandSharednetworkNetworks),
		Nextserver:                  flex.ExpandStringPointerNullAsEmpty(m.Nextserver),
		Options:                     flex.ExpandFrameworkListNestedBlock(ctx, m.Options, diags, ExpandSharednetworkOptions),
		PxeLeaseTime:                flex.ExpandInt64Pointer(m.PxeLeaseTime),
		UpdateDnsOnLeaseRenewal:     flex.ExpandBoolPointer(m.UpdateDnsOnLeaseRenewal),
	}
}

// ApplySharednetworkNIOSUseFlags derives NIOS use flags from the raw config
// value(s) and writes them onto the core model. A flag is true when the user
// set any of its governed value fields in config.
func ApplySharednetworkNIOSUseFlags(ctx context.Context, config tfsdk.Config, obj *coremodel.Sharednetwork, diags *diag.Diagnostics) {
	if obj == nil || obj.NIOS == nil {
		return
	}
	obj.NIOS.UseAuthority = flex.DeriveUseFlag(ctx, config, diags, path.Root("nios").AtName("authority"))
	obj.NIOS.UseBootfile = flex.DeriveUseFlag(ctx, config, diags, path.Root("nios").AtName("bootfile"))
	obj.NIOS.UseBootserver = flex.DeriveUseFlag(ctx, config, diags, path.Root("nios").AtName("bootserver"))
	obj.NIOS.UseDdnsGenerateHostname = flex.DeriveUseFlag(ctx, config, diags, path.Root("nios").AtName("ddns_generate_hostname"))
	obj.NIOS.UseDdnsTtl = flex.DeriveUseFlag(ctx, config, diags, path.Root("nios").AtName("ddns_ttl"))
	obj.NIOS.UseDdnsUpdateFixedAddresses = flex.DeriveUseFlag(ctx, config, diags, path.Root("nios").AtName("ddns_update_fixed_addresses"))
	obj.NIOS.UseDdnsUseOption81 = flex.DeriveUseFlag(ctx, config, diags, path.Root("nios").AtName("ddns_use_option81"))
	obj.NIOS.UseDenyBootp = flex.DeriveUseFlag(ctx, config, diags, path.Root("nios").AtName("deny_bootp"))
	obj.NIOS.UseEnableDdns = flex.DeriveUseFlag(ctx, config, diags, path.Root("nios").AtName("enable_ddns"))
	obj.NIOS.UseIgnoreClientIdentifier = flex.DeriveUseFlag(ctx, config, diags, path.Root("nios").AtName("ignore_client_identifier"))
	obj.NIOS.UseIgnoreDhcpOptionListRequest = flex.DeriveUseFlag(ctx, config, diags, path.Root("nios").AtName("ignore_dhcp_option_list_request"))
	obj.NIOS.UseIgnoreId = flex.DeriveUseFlag(ctx, config, diags, path.Root("nios").AtName("ignore_id"))
	obj.NIOS.UseLeaseScavengeTime = flex.DeriveUseFlag(ctx, config, diags, path.Root("nios").AtName("lease_scavenge_time"))
	obj.NIOS.UseLogicFilterRules = flex.DeriveUseFlag(ctx, config, diags, path.Root("nios").AtName("logic_filter_rules"))
	obj.NIOS.UseNextserver = flex.DeriveUseFlag(ctx, config, diags, path.Root("nios").AtName("nextserver"))
	obj.NIOS.UseOptions = flex.DeriveUseFlag(ctx, config, diags, path.Root("nios").AtName("options"))
	obj.NIOS.UsePxeLeaseTime = flex.DeriveUseFlag(ctx, config, diags, path.Root("nios").AtName("pxe_lease_time"))
	obj.NIOS.UseUpdateDnsOnLeaseRenewal = flex.DeriveUseFlag(ctx, config, diags, path.Root("nios").AtName("update_dns_on_lease_renewal"))
}

// Flatten populates the TF model from a core response.
func (m *SharednetworkModel) Flatten(ctx context.Context, resp *coremodel.Sharednetwork, diags *diag.Diagnostics) {
	if resp == nil {
		return
	}

	m.Id = flex.FlattenStringPointer(resp.Id)

	// Extract existing NIOS model, flatten API response onto it, convert back
	niosModel := flex.ExpandNestedObject[NIOSSharednetworkModel](ctx, m.NIOS, diags)
	if niosModel == nil {
		niosModel = &NIOSSharednetworkModel{}
	}
	niosModel.Flatten(ctx, resp.NIOS, diags)
	if resp.NIOS != nil {
		m.NIOS = flex.FlattenNestedObject(ctx, niosModel, NIOSSharednetworkAttrTypes, diags)
	} else {
		m.NIOS = types.ObjectNull(NIOSSharednetworkAttrTypes)
	}

}

// Flatten merges API response onto existing NIOS model.
func (m *NIOSSharednetworkModel) Flatten(ctx context.Context, from *coremodel.NIOSSharednetworkExt, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	planExtAttrs := m.ExtAttrs
	if planExtAttrs.IsUnknown() {
		planExtAttrs = types.MapNull(types.StringType)
	}
	m.Authority = flex.FlattenBoolPointer(from.Authority)
	m.Bootfile = flex.FlattenStringPointerEmptyAsNull(from.Bootfile)
	m.Bootserver = flex.FlattenStringPointerEmptyAsNull(from.Bootserver)
	m.Comment = flex.FlattenStringPointerEmptyAsNull(from.Comment)
	m.DdnsGenerateHostname = flex.FlattenBoolPointer(from.DdnsGenerateHostname)
	m.DdnsServerAlwaysUpdates = flex.FlattenBoolPointer(from.DdnsServerAlwaysUpdates)
	m.DdnsTtl = flex.FlattenInt64Pointer(from.DdnsTtl)
	m.DdnsUpdateFixedAddresses = flex.FlattenBoolPointer(from.DdnsUpdateFixedAddresses)
	m.DdnsUseOption81 = flex.FlattenBoolPointer(from.DdnsUseOption81)
	m.DenyBootp = flex.FlattenBoolPointer(from.DenyBootp)
	m.Disable = flex.FlattenBoolPointer(from.Disable)
	m.EnableDdns = flex.FlattenBoolPointer(from.EnableDdns)
	m.EnablePxeLeaseTime = flex.FlattenBoolPointer(from.EnablePxeLeaseTime)
	m.ExtAttrs, m.ExtAttrsAll = flex.FlattenEAs(planExtAttrs, from.ExtAttrs)
	m.IgnoreClientIdentifier = flex.FlattenBoolPointer(from.IgnoreClientIdentifier)
	m.IgnoreDhcpOptionListRequest = flex.FlattenBoolPointer(from.IgnoreDhcpOptionListRequest)
	m.IgnoreId = flex.FlattenStringPointerEmptyAsNull(from.IgnoreId)
	m.IgnoreMacAddresses = flex.FlattenFrameworkUnorderedListString(ctx, from.IgnoreMacAddresses, diags)
	m.LeaseScavengeTime = flex.FlattenInt64Pointer(from.LeaseScavengeTime)
	m.LogicFilterRules = flex.FlattenFrameworkListNestedBlock(ctx, from.LogicFilterRules, SharednetworkLogicFilterRulesAttrTypes, diags, FlattenSharednetworkLogicFilterRules)
	m.Name = flex.FlattenStringPointerEmptyAsNull(from.Name)
	m.NetworkView = flex.FlattenStringPointerEmptyAsNull(from.NetworkView)
	m.Networks = flex.FlattenFrameworkListNestedBlock(ctx, from.Networks, SharednetworkNetworksAttrTypes, diags, FlattenSharednetworkNetworks)
	m.Nextserver = flex.FlattenStringPointerEmptyAsNull(from.Nextserver)
	m.Options = flex.FlattenFrameworkListNestedBlock(ctx, from.Options, SharednetworkOptionsAttrTypes, diags, FlattenSharednetworkOptions)
	m.PxeLeaseTime = flex.FlattenInt64Pointer(from.PxeLeaseTime)
	m.UpdateDnsOnLeaseRenewal = flex.FlattenBoolPointer(from.UpdateDnsOnLeaseRenewal)
}
