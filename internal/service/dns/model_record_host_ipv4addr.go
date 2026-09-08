package dns

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/hashicorp/terraform-plugin-framework-nettypes/iptypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/boolvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	niosdns "github.com/infobloxopen/infoblox-nios-go-client/dns"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/dynamicallocation"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
	customvalidator "github.com/infobloxopen/terraform-provider-infoblox/internal/validator"
)

// RecordHostIpv4addrModel is the Terraform model for RecordHostIpv4addr
type RecordHostIpv4addrModel struct {
	Bootfile                        types.String        `tfsdk:"bootfile"`
	Bootserver                      types.String        `tfsdk:"bootserver"`
	ConfigureForDhcp                types.Bool          `tfsdk:"configure_for_dhcp"`
	DenyBootp                       types.Bool          `tfsdk:"deny_bootp"`
	EnablePxeLeaseTime              types.Bool          `tfsdk:"enable_pxe_lease_time"`
	IgnoreClientRequestedOptions    types.Bool          `tfsdk:"ignore_client_requested_options"`
	Ipv4addr                        iptypes.IPv4Address `tfsdk:"ipv4addr"`
	LogicFilterRules                types.List          `tfsdk:"logic_filter_rules"`
	Mac                             types.String        `tfsdk:"mac"`
	MatchClient                     types.String        `tfsdk:"match_client"`
	Nextserver                      types.String        `tfsdk:"nextserver"`
	Options                         types.List          `tfsdk:"options"`
	PxeLeaseTime                    types.Int64         `tfsdk:"pxe_lease_time"`
	ReservedInterface               types.String        `tfsdk:"reserved_interface"`
	UseBootfile                     types.Bool          `tfsdk:"use_bootfile"`
	UseBootserver                   types.Bool          `tfsdk:"use_bootserver"`
	UseDenyBootp                    types.Bool          `tfsdk:"use_deny_bootp"`
	UseForEaInheritance             types.Bool          `tfsdk:"use_for_ea_inheritance"`
	UseIgnoreClientRequestedOptions types.Bool          `tfsdk:"use_ignore_client_requested_options"`
	UseLogicFilterRules             types.Bool          `tfsdk:"use_logic_filter_rules"`
	UseNextserver                   types.Bool          `tfsdk:"use_nextserver"`
	UseOptions                      types.Bool          `tfsdk:"use_options"`
	UsePxeLeaseTime                 types.Bool          `tfsdk:"use_pxe_lease_time"`
	DynamicAllocation               types.Object        `tfsdk:"dynamic_allocation"`
}

// RecordHostIpv4addrAttrTypes contains the attribute types for RecordHostIpv4addrModel
var RecordHostIpv4addrAttrTypes = map[string]attr.Type{
	"bootfile":                            types.StringType,
	"bootserver":                          types.StringType,
	"configure_for_dhcp":                  types.BoolType,
	"deny_bootp":                          types.BoolType,
	"enable_pxe_lease_time":               types.BoolType,
	"ignore_client_requested_options":     types.BoolType,
	"ipv4addr":                            iptypes.IPv4AddressType{},
	"logic_filter_rules":                  types.ListType{ElemType: types.ObjectType{AttrTypes: RecordHostIpv4addrLogicFilterRulesAttrTypes}},
	"mac":                                 types.StringType,
	"match_client":                        types.StringType,
	"nextserver":                          types.StringType,
	"options":                             types.ListType{ElemType: types.ObjectType{AttrTypes: RecordHostIpv4addrOptionsAttrTypes}},
	"pxe_lease_time":                      types.Int64Type,
	"reserved_interface":                  types.StringType,
	"use_bootfile":                        types.BoolType,
	"use_bootserver":                      types.BoolType,
	"use_deny_bootp":                      types.BoolType,
	"use_for_ea_inheritance":              types.BoolType,
	"use_ignore_client_requested_options": types.BoolType,
	"use_logic_filter_rules":              types.BoolType,
	"use_nextserver":                      types.BoolType,
	"use_options":                         types.BoolType,
	"use_pxe_lease_time":                  types.BoolType,
	"dynamic_allocation":                  types.ObjectType{AttrTypes: dynamicallocation.NextAvailableIpAttrTypes},
}

// RecordHostIpv4addrResourceSchemaAttributes contains the schema attributes for RecordHostIpv4addrModel
var RecordHostIpv4addrResourceSchemaAttributes = map[string]schema.Attribute{
	"bootfile": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
			stringvalidator.AlsoRequires(path.MatchRelative().AtParent().AtName("use_bootfile")),
		},
		MarkdownDescription: "The name of the boot file the client must download.",
	},
	"bootserver": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
			stringvalidator.AlsoRequires(path.MatchRelative().AtParent().AtName("use_bootserver")),
			customvalidator.IsValidIPv4OrFQDN(),
		},
		MarkdownDescription: "The IP address or hostname of the boot file server where the boot file is stored.",
	},
	"configure_for_dhcp": schema.BoolAttribute{
		Computed:            true,
		MarkdownDescription: "Set this to True to enable the DHCP configuration for this host address.",
	},
	"deny_bootp": schema.BoolAttribute{
		Optional: true,
		Computed: true,
		Validators: []validator.Bool{
			boolvalidator.AlsoRequires(path.MatchRelative().AtParent().AtName("use_deny_bootp")),
		},
		MarkdownDescription: "Set this to True to disable the BOOTP settings and deny BOOTP boot requests.",
	},
	"enable_pxe_lease_time": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Set this to True if you want the DHCP server to use a different lease time for PXE clients. You can specify the duration of time it takes a host to connect to a boot server, such as a TFTP server, and download the file it needs to boot. For example, set a longer lease time if the client downloads an OS (operating system) or configuration file, or set a shorter lease time if the client downloads only configuration changes. Enter the lease time for the preboot execution environment for hosts to boot remotely from a server.",
	},
	"ignore_client_requested_options": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "If this field is set to false, the appliance returns all DHCP options the client is eligible to receive, rather than only the list of options the client has requested.",
	},
	"ipv4addr": schema.StringAttribute{
		Optional:   true,
		Computed:   true,
		CustomType: iptypes.IPv4AddressType{},
		Validators: []validator.String{
			stringvalidator.ExactlyOneOf(
				path.MatchRelative().AtParent().AtName("dynamic_allocation"),
			),
			customvalidator.StringNotEmpty(),
		},
		MarkdownDescription: "The IPv4 Address of the record.",
	},
	"logic_filter_rules": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: RecordHostIpv4addrLogicFilterRulesResourceSchemaAttributes,
		},
		Optional: true,
		Validators: []validator.List{
			customvalidator.ListNotEmpty(),
			listvalidator.AlsoRequires(path.MatchRelative().AtParent().AtName("use_logic_filter_rules")),
		},
		MarkdownDescription: "This field contains the logic filters to be applied on the this host address. This list corresponds to the match rules that are written to the dhcpd configuration file.",
	},
	"mac": schema.StringAttribute{
		Computed: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
		},
		MarkdownDescription: "The MAC address for this host address.",
	},
	"match_client": schema.StringAttribute{
		Computed: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
		},
		MarkdownDescription: "Set this to 'MAC_ADDRESS' to assign the IP address to the selected host, provided that the MAC address of the requesting host matches the MAC address that you specify in the field. Set this to 'RESERVED' to reserve this particular IP address for future use, or if the IP address is statically configured on a system (the Infoblox server does not assign the address from a DHCP request).",
	},
	"nextserver": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
			stringvalidator.AlsoRequires(path.MatchRelative().AtParent().AtName("use_nextserver")),
			customvalidator.IsValidIPv4OrFQDN(),
		},
		MarkdownDescription: "The name in FQDN format and/or IPv4 Address of the next server that the host needs to boot.",
	},
	"options": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: RecordHostIpv4addrOptionsResourceSchemaAttributes,
		},
		Optional: true,
		Validators: []validator.List{
			customvalidator.ListNotEmpty(),
		},
		MarkdownDescription: "An array of DHCP option dhcpoption structs that lists the DHCP options associated with the object.",
	},
	"pxe_lease_time": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "The lease time for PXE clients, see *enable_pxe_lease_time* for more information.",
	},
	"reserved_interface": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
		},
		MarkdownDescription: "The reference to the reserved interface to which the device belongs.",
	},
	"use_bootfile": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Use flag for: bootfile",
	},
	"use_bootserver": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Use flag for: bootserver",
	},
	"use_deny_bootp": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Use flag for: deny_bootp",
	},
	"use_for_ea_inheritance": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Set this to True when using this host address for EA inheritance.",
	},
	"use_ignore_client_requested_options": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Use flag for: ignore_client_requested_options",
	},
	"use_logic_filter_rules": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Use flag for: logic_filter_rules",
	},
	"use_nextserver": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Use flag for: nextserver",
	},
	"use_options": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Use flag for: options",
	},
	"use_pxe_lease_time": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Use flag for: pxe_lease_time",
	},
	"dynamic_allocation": schema.SingleNestedAttribute{
		Attributes:          dynamicallocation.NextAvailableIpResourceSchemaAttributes,
		Optional:            true,
		MarkdownDescription: "Dynamically allocate the ip using the NIOS next_available_ip function call. Mutually exclusive with the static value field.",
	},
}

// ExpandRecordHostIpv4addr converts a Terraform Object to SDK type
func ExpandRecordHostIpv4addr(ctx context.Context, o types.Object, diags *diag.Diagnostics) *niosdns.RecordHostIpv4addr {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m RecordHostIpv4addrModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

// Expand converts the Terraform model to SDK type
func (m *RecordHostIpv4addrModel) Expand(ctx context.Context, diags *diag.Diagnostics) *niosdns.RecordHostIpv4addr {
	if m == nil {
		return nil
	}
	to := &niosdns.RecordHostIpv4addr{
		Bootfile:                        flex.ExpandStringPointerNullAsEmpty(m.Bootfile),
		Bootserver:                      flex.ExpandStringPointerNullAsEmpty(m.Bootserver),
		ConfigureForDhcp:                flex.ExpandBoolPointer(m.ConfigureForDhcp),
		DenyBootp:                       flex.ExpandBoolPointer(m.DenyBootp),
		EnablePxeLeaseTime:              flex.ExpandBoolPointer(m.EnablePxeLeaseTime),
		IgnoreClientRequestedOptions:    flex.ExpandBoolPointer(m.IgnoreClientRequestedOptions),
		Ipv4addr:                        ExpandRecordHostIpv4addrIpv4addr(m.Ipv4addr),
		LogicFilterRules:                flex.ExpandFrameworkListNestedBlock(ctx, m.LogicFilterRules, diags, ExpandRecordHostIpv4addrLogicFilterRules),
		Mac:                             flex.ExpandStringPointer(m.Mac),
		MatchClient:                     flex.ExpandStringPointer(m.MatchClient),
		Nextserver:                      flex.ExpandStringPointerNullAsEmpty(m.Nextserver),
		Options:                         flex.ExpandFrameworkListNestedBlock(ctx, m.Options, diags, ExpandRecordHostIpv4addrOptions),
		PxeLeaseTime:                    flex.ExpandInt64Pointer(m.PxeLeaseTime),
		ReservedInterface:               flex.ExpandStringPointer(m.ReservedInterface),
		UseBootfile:                     flex.ExpandBoolPointer(m.UseBootfile),
		UseBootserver:                   flex.ExpandBoolPointer(m.UseBootserver),
		UseDenyBootp:                    flex.ExpandBoolPointer(m.UseDenyBootp),
		UseForEaInheritance:             flex.ExpandBoolPointer(m.UseForEaInheritance),
		UseIgnoreClientRequestedOptions: flex.ExpandBoolPointer(m.UseIgnoreClientRequestedOptions),
		UseLogicFilterRules:             flex.ExpandBoolPointer(m.UseLogicFilterRules),
		UseNextserver:                   flex.ExpandBoolPointer(m.UseNextserver),
		UseOptions:                      flex.ExpandBoolPointer(m.UseOptions),
		UsePxeLeaseTime:                 flex.ExpandBoolPointer(m.UsePxeLeaseTime),
		FuncCall:                        BuildRecordHostIpv4addrFuncCall(ctx, m.DynamicAllocation, diags),
	}
	return to
}

// FlattenRecordHostIpv4addr converts an SDK type to Terraform Object
func FlattenRecordHostIpv4addr(ctx context.Context, from *niosdns.RecordHostIpv4addr, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(RecordHostIpv4addrAttrTypes)
	}
	m := &RecordHostIpv4addrModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, RecordHostIpv4addrAttrTypes, m)
	diags.Append(d...)
	return t
}

// Flatten populates the Terraform model from SDK type
func (m *RecordHostIpv4addrModel) Flatten(ctx context.Context, from *niosdns.RecordHostIpv4addr, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	m.Bootfile = flex.FlattenStringPointerEmptyAsNull(from.Bootfile)
	m.Bootserver = flex.FlattenStringPointerEmptyAsNull(from.Bootserver)
	m.ConfigureForDhcp = flex.FlattenBoolPointer(from.ConfigureForDhcp)
	m.DenyBootp = flex.FlattenBoolPointer(from.DenyBootp)
	m.EnablePxeLeaseTime = flex.FlattenBoolPointer(from.EnablePxeLeaseTime)
	m.IgnoreClientRequestedOptions = flex.FlattenBoolPointer(from.IgnoreClientRequestedOptions)
	m.Ipv4addr = FlattenRecordHostIpv4addrIpv4addr(from.Ipv4addr)
	m.LogicFilterRules = flex.FlattenFrameworkListNestedBlock(ctx, from.LogicFilterRules, RecordHostIpv4addrLogicFilterRulesAttrTypes, diags, FlattenRecordHostIpv4addrLogicFilterRules)
	m.Mac = flex.FlattenStringPointerEmptyAsNull(from.Mac)
	m.MatchClient = flex.FlattenStringPointerEmptyAsNull(from.MatchClient)
	m.Nextserver = flex.FlattenStringPointerEmptyAsNull(from.Nextserver)
	m.Options = flex.FlattenFrameworkListNestedBlock(ctx, from.Options, RecordHostIpv4addrOptionsAttrTypes, diags, FlattenRecordHostIpv4addrOptions)
	m.PxeLeaseTime = flex.FlattenInt64Pointer(from.PxeLeaseTime)
	m.ReservedInterface = flex.FlattenStringPointerEmptyAsNull(from.ReservedInterface)
	m.UseBootfile = flex.FlattenBoolPointer(from.UseBootfile)
	m.UseBootserver = flex.FlattenBoolPointer(from.UseBootserver)
	m.UseDenyBootp = flex.FlattenBoolPointer(from.UseDenyBootp)
	m.UseForEaInheritance = flex.FlattenBoolPointer(from.UseForEaInheritance)
	m.UseIgnoreClientRequestedOptions = flex.FlattenBoolPointer(from.UseIgnoreClientRequestedOptions)
	m.UseLogicFilterRules = flex.FlattenBoolPointer(from.UseLogicFilterRules)
	m.UseNextserver = flex.FlattenBoolPointer(from.UseNextserver)
	m.UseOptions = flex.FlattenBoolPointer(from.UseOptions)
	m.UsePxeLeaseTime = flex.FlattenBoolPointer(from.UsePxeLeaseTime)
	if len(m.DynamicAllocation.AttributeTypes(ctx)) == 0 {
		m.DynamicAllocation = types.ObjectNull(dynamicallocation.NextAvailableIpAttrTypes)
	}
}

func ExpandRecordHostIpv4addrIpv4addr(v iptypes.IPv4Address) *niosdns.RecordHostIpv4addrIpv4addr {
	return &niosdns.RecordHostIpv4addrIpv4addr{String: flex.ExpandIPv4Address(v)}
}

func FlattenRecordHostIpv4addrIpv4addr(v *niosdns.RecordHostIpv4addrIpv4addr) iptypes.IPv4Address {
	var value *string
	if v != nil {
		value = v.String
	}
	return flex.FlattenIPv4Address(value)
}
