package dhcp

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/boolvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/mapvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/mapdefault"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"

	coremodel "github.com/infobloxopen/terraform-provider-infoblox/internal/core/model/dhcp"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
	customvalidator "github.com/infobloxopen/terraform-provider-infoblox/internal/validator"
)

type FixedaddresstemplateModel struct {
	Id            types.String `tfsdk:"id"`
	UpdateTrigger types.String `tfsdk:"update_trigger"`
	NIOS          types.Object `tfsdk:"nios"`
}

var FixedaddresstemplateAttrTypes = map[string]attr.Type{
	"id":             types.StringType,
	"update_trigger": types.StringType,
	"nios":           types.ObjectType{AttrTypes: NIOSFixedaddresstemplateAttrTypes},
}

type NIOSFixedaddresstemplateModel struct {
	Bootfile                    types.String `tfsdk:"bootfile"`
	Bootserver                  types.String `tfsdk:"bootserver"`
	Comment                     types.String `tfsdk:"comment"`
	DdnsDomainname              types.String `tfsdk:"ddns_domainname"`
	DdnsHostname                types.String `tfsdk:"ddns_hostname"`
	DenyBootp                   types.Bool   `tfsdk:"deny_bootp"`
	EnableDdns                  types.Bool   `tfsdk:"enable_ddns"`
	EnablePxeLeaseTime          types.Bool   `tfsdk:"enable_pxe_lease_time"`
	ExtAttrs                    types.Map    `tfsdk:"ext_attrs"`
	ExtAttrsAll                 types.Map    `tfsdk:"ext_attrs_all"`
	IgnoreDhcpOptionListRequest types.Bool   `tfsdk:"ignore_dhcp_option_list_request"`
	LogicFilterRules            types.List   `tfsdk:"logic_filter_rules"`
	Name                        types.String `tfsdk:"name"`
	Nextserver                  types.String `tfsdk:"nextserver"`
	NumberOfAddresses           types.Int64  `tfsdk:"number_of_addresses"`
	Offset                      types.Int64  `tfsdk:"offset"`
	Options                     types.List   `tfsdk:"options"`
	PxeLeaseTime                types.Int64  `tfsdk:"pxe_lease_time"`
}

var NIOSFixedaddresstemplateAttrTypes = map[string]attr.Type{
	"bootfile":                        types.StringType,
	"bootserver":                      types.StringType,
	"comment":                         types.StringType,
	"ddns_domainname":                 types.StringType,
	"ddns_hostname":                   types.StringType,
	"deny_bootp":                      types.BoolType,
	"enable_ddns":                     types.BoolType,
	"enable_pxe_lease_time":           types.BoolType,
	"ext_attrs":                       types.MapType{ElemType: types.StringType},
	"ext_attrs_all":                   types.MapType{ElemType: types.StringType},
	"ignore_dhcp_option_list_request": types.BoolType,
	"logic_filter_rules":              types.ListType{ElemType: types.ObjectType{AttrTypes: FixedaddresstemplateLogicFilterRulesAttrTypes}},
	"name":                            types.StringType,
	"nextserver":                      types.StringType,
	"number_of_addresses":             types.Int64Type,
	"offset":                          types.Int64Type,
	"options":                         types.ListType{ElemType: types.ObjectType{AttrTypes: FixedaddresstemplateOptionsAttrTypes}},
	"pxe_lease_time":                  types.Int64Type,
}

const (
	FixedaddresstemplateReturnFields = "bootfile,bootserver,comment,ddns_domainname,ddns_hostname,deny_bootp,enable_ddns,enable_pxe_lease_time,extattrs,ignore_dhcp_option_list_request,logic_filter_rules,name,nextserver,number_of_addresses,offset,options,pxe_lease_time,use_bootfile,use_bootserver,use_ddns_domainname,use_deny_bootp,use_enable_ddns,use_ignore_dhcp_option_list_request,use_logic_filter_rules,use_nextserver,use_options,use_pxe_lease_time"
)

var FixedaddresstemplateResourceSchemaAttributes = map[string]schema.Attribute{
	"id": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The reference to the object.",
	},
	"update_trigger": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "An arbitrary value used to trigger an update. Not sent to the API. Change it when Terraform reports no infrastructure changes.",
	},
	"nios": schema.SingleNestedAttribute{
		Optional:            true,
		MarkdownDescription: "NIOS backend-specific fields.",
		Attributes:          FixedaddresstemplateResourceNiosSchemaAttributes,
	},
}

var FixedaddresstemplateResourceNiosSchemaAttributes = map[string]schema.Attribute{
	"bootfile": schema.StringAttribute{
		Optional: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
		},
		MarkdownDescription: "The boot file name for the fixed address. You can configure the DHCP server to support clients that use the boot file name option in their DHCPREQUEST messages.",
	},
	"bootserver": schema.StringAttribute{
		Optional: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
			customvalidator.IsValidIPv4OrFQDN(),
		},
		MarkdownDescription: "The boot server address for the fixed address. You can specify the name and/or IP address of the boot server that the host needs to boot. The boot server IPv4 Address or name in FQDN format.",
	},
	"comment": schema.StringAttribute{
		Optional: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
			stringvalidator.LengthBetween(0, 256),
		},
		MarkdownDescription: "A descriptive comment of a fixed address template object.",
	},
	"ddns_domainname": schema.StringAttribute{
		Optional: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
			customvalidator.ValidateTrimmedString(),
		},
		MarkdownDescription: "The dynamic DNS domain name the appliance uses specifically for DDNS updates for this fixed address.",
	},
	"ddns_hostname": schema.StringAttribute{
		Optional: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
			customvalidator.ValidateTrimmedString(),
		},
		MarkdownDescription: "The DDNS host name for this fixed address.",
	},
	"deny_bootp": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Determines if BOOTP settings are disabled and BOOTP requests will be denied.",
	},
	"enable_ddns": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Determines if the DHCP server sends DDNS updates to DNS servers in the same Grid, and to external DNS servers.",
	},
	"enable_pxe_lease_time": schema.BoolAttribute{
		Optional: true,
		Computed: true,
		Default:  booldefault.StaticBool(false),
		Validators: []validator.Bool{
			boolvalidator.AlsoRequires(path.MatchRelative().AtParent().AtName("pxe_lease_time")),
		},
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
		MarkdownDescription: "All ext_attrs including inherited values.",
	},
	"ignore_dhcp_option_list_request": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "If this field is set to False, the appliance returns all DHCP options the client is eligible to receive, rather than only the list of options the client has requested.",
	},
	"logic_filter_rules": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: FixedaddresstemplateLogicFilterRulesResourceSchemaAttributes,
		},
		Optional: true,
		Validators: []validator.List{
			customvalidator.ListNotEmpty(),
		},
		MarkdownDescription: "This field contains the logic filters to be applied on this fixed address. This list corresponds to the match rules that are written to the dhcpd configuration file.",
	},
	"name": schema.StringAttribute{
		Required: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
			customvalidator.ValidateTrimmedString(),
		},
		MarkdownDescription: "The name of a fixed address template object.",
	},
	"nextserver": schema.StringAttribute{
		Optional: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
			customvalidator.IsValidIPv4OrFQDN(),
		},
		MarkdownDescription: "The name in FQDN and/or IPv4 Address format of the next server that the host needs to boot.",
	},
	"number_of_addresses": schema.Int64Attribute{
		Optional: true,
		Computed: true,
		Validators: []validator.Int64{
			int64validator.AlsoRequires(path.MatchRelative().AtParent().AtName("offset")),
		},
		MarkdownDescription: "The number of addresses for this fixed address.",
	},
	"offset": schema.Int64Attribute{
		Optional: true,
		Computed: true,
		Validators: []validator.Int64{
			int64validator.AlsoRequires(path.MatchRelative().AtParent().AtName("number_of_addresses")),
		},
		MarkdownDescription: "The start address offset for this fixed address.",
	},
	"options": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: FixedaddresstemplateOptionsResourceSchemaAttributes,
		},
		Optional: true,
		Computed: true,
		Default:  listdefault.StaticValue(types.ListValueMust(types.ObjectType{AttrTypes: FixedaddresstemplateOptionsAttrTypes}, []attr.Value{})),
		Validators: []validator.List{
			customvalidator.ListNotEmpty(),
		},
		MarkdownDescription: "An array of DHCP option dhcpoption structs that lists the DHCP options associated with the object.",
	},
	"pxe_lease_time": schema.Int64Attribute{
		Optional: true,
		Computed: true,
		Validators: []validator.Int64{
			int64validator.Between(0, 399999999),
		},
		MarkdownDescription: "The PXE lease time value for a DHCP Fixed Address object. Some hosts use PXE (Preboot Execution Environment) to boot remotely from a server. To better manage your IP resources, set a different lease time for PXE boot requests. You can configure the DHCP server to allocate an IP address with a shorter lease time to hosts that send PXE boot requests, so IP addresses are not leased longer than necessary. A 32-bit unsigned integer that represents the duration, in seconds, for which the update is cached. Zero indicates that the update is not cached.",
	},
}

// Expand converts the TF model to the infoblox core model
func (m *FixedaddresstemplateModel) Expand(ctx context.Context, diags *diag.Diagnostics, isCreate bool) *coremodel.Fixedaddresstemplate {
	if m == nil {
		return nil
	}

	obj := &coremodel.Fixedaddresstemplate{}

	// Expand NIOS nested attribute (returns nil if not present)
	niosModel := flex.ExpandNestedObject[NIOSFixedaddresstemplateModel](ctx, m.NIOS, diags)
	if niosModel != nil {
		obj.NIOS = niosModel.Expand(ctx, diags)
	}

	return obj
}

// Expand converts the NIOS TF model to the core model.
func (m *NIOSFixedaddresstemplateModel) Expand(ctx context.Context, diags *diag.Diagnostics) *coremodel.NIOSFixedaddresstemplateExt {
	return &coremodel.NIOSFixedaddresstemplateExt{
		Bootfile:                    flex.ExpandStringPointerNullAsEmpty(m.Bootfile),
		Bootserver:                  flex.ExpandStringPointerNullAsEmpty(m.Bootserver),
		Comment:                     flex.ExpandStringPointerNullAsEmpty(m.Comment),
		DdnsDomainname:              flex.ExpandStringPointerNullAsEmpty(m.DdnsDomainname),
		DdnsHostname:                flex.ExpandStringPointerNullAsEmpty(m.DdnsHostname),
		DenyBootp:                   flex.ExpandBoolPointer(m.DenyBootp),
		EnableDdns:                  flex.ExpandBoolPointer(m.EnableDdns),
		EnablePxeLeaseTime:          flex.ExpandBoolPointer(m.EnablePxeLeaseTime),
		ExtAttrs:                    flex.ExpandMapStringAny(ctx, m.ExtAttrs, diags),
		IgnoreDhcpOptionListRequest: flex.ExpandBoolPointer(m.IgnoreDhcpOptionListRequest),
		LogicFilterRules:            flex.ExpandFrameworkListNestedBlock(ctx, m.LogicFilterRules, diags, ExpandFixedaddresstemplateLogicFilterRules),
		Name:                        flex.ExpandStringPointerNullAsEmpty(m.Name),
		Nextserver:                  flex.ExpandStringPointerNullAsEmpty(m.Nextserver),
		NumberOfAddresses:           flex.ExpandInt64Pointer(m.NumberOfAddresses),
		Offset:                      flex.ExpandInt64Pointer(m.Offset),
		Options:                     flex.ExpandFrameworkListNestedBlock(ctx, m.Options, diags, ExpandFixedaddresstemplateOptions),
		PxeLeaseTime:                flex.ExpandInt64Pointer(m.PxeLeaseTime),
	}
}

// ApplyFixedaddresstemplateNIOSUseFlags derives NIOS use flags from the raw config
// value(s) and writes them onto the core model. A flag is true when the user
// set any of its governed value fields in config.
func ApplyFixedaddresstemplateNIOSUseFlags(ctx context.Context, config tfsdk.Config, obj *coremodel.Fixedaddresstemplate, diags *diag.Diagnostics) {
	if obj == nil || obj.NIOS == nil {
		return
	}
	obj.NIOS.UseBootfile = flex.DeriveUseFlag(ctx, config, diags, path.Root("nios").AtName("bootfile"))
	obj.NIOS.UseBootserver = flex.DeriveUseFlag(ctx, config, diags, path.Root("nios").AtName("bootserver"))
	obj.NIOS.UseDdnsDomainname = flex.DeriveUseFlag(ctx, config, diags, path.Root("nios").AtName("ddns_domainname"))
	obj.NIOS.UseDenyBootp = flex.DeriveUseFlag(ctx, config, diags, path.Root("nios").AtName("deny_bootp"))
	obj.NIOS.UseEnableDdns = flex.DeriveUseFlag(ctx, config, diags, path.Root("nios").AtName("enable_ddns"))
	obj.NIOS.UseIgnoreDhcpOptionListRequest = flex.DeriveUseFlag(ctx, config, diags, path.Root("nios").AtName("ignore_dhcp_option_list_request"))
	obj.NIOS.UseLogicFilterRules = flex.DeriveUseFlag(ctx, config, diags, path.Root("nios").AtName("logic_filter_rules"))
	obj.NIOS.UseNextserver = flex.DeriveUseFlag(ctx, config, diags, path.Root("nios").AtName("nextserver"))
	obj.NIOS.UseOptions = flex.DeriveUseFlag(ctx, config, diags, path.Root("nios").AtName("options"))
	obj.NIOS.UsePxeLeaseTime = flex.DeriveUseFlag(ctx, config, diags, path.Root("nios").AtName("pxe_lease_time"))
}

// Flatten populates the TF model from a core response.
func (m *FixedaddresstemplateModel) Flatten(ctx context.Context, resp *coremodel.Fixedaddresstemplate, diags *diag.Diagnostics) {
	if resp == nil {
		return
	}

	m.Id = flex.FlattenStringPointer(resp.Id)

	// Extract existing NIOS model, flatten API response onto it, convert back
	niosModel := flex.ExpandNestedObject[NIOSFixedaddresstemplateModel](ctx, m.NIOS, diags)
	if niosModel == nil {
		niosModel = &NIOSFixedaddresstemplateModel{}
	}
	plannedNIOS := flex.ExpandNestedObject[NIOSFixedaddresstemplateModel](ctx, m.NIOS, diags)
	niosModel.Flatten(ctx, resp.NIOS, diags)
	if resp.NIOS != nil {
		PostFlattenFixedaddresstemplateNIOS(ctx, plannedNIOS, niosModel, diags)
		m.NIOS = flex.FlattenNestedObject(ctx, niosModel, NIOSFixedaddresstemplateAttrTypes, diags)
	} else {
		m.NIOS = types.ObjectNull(NIOSFixedaddresstemplateAttrTypes)
	}

}

// Flatten merges API response onto existing NIOS model.
func (m *NIOSFixedaddresstemplateModel) Flatten(ctx context.Context, from *coremodel.NIOSFixedaddresstemplateExt, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	planExtAttrs := m.ExtAttrs
	if planExtAttrs.IsUnknown() {
		planExtAttrs = types.MapNull(types.StringType)
	}
	m.Bootfile = flex.FlattenStringPointerEmptyAsNull(from.Bootfile)
	m.Bootserver = flex.FlattenStringPointerEmptyAsNull(from.Bootserver)
	m.Comment = flex.FlattenStringPointerEmptyAsNull(from.Comment)
	m.DdnsDomainname = flex.FlattenStringPointerEmptyAsNull(from.DdnsDomainname)
	m.DdnsHostname = flex.FlattenStringPointerEmptyAsNull(from.DdnsHostname)
	m.DenyBootp = flex.FlattenBoolPointer(from.DenyBootp)
	m.EnableDdns = flex.FlattenBoolPointer(from.EnableDdns)
	m.EnablePxeLeaseTime = flex.FlattenBoolPointer(from.EnablePxeLeaseTime)
	m.ExtAttrs, m.ExtAttrsAll = flex.FlattenEAs(planExtAttrs, from.ExtAttrs)
	m.IgnoreDhcpOptionListRequest = flex.FlattenBoolPointer(from.IgnoreDhcpOptionListRequest)
	m.LogicFilterRules = flex.FlattenFrameworkListNestedBlock(ctx, from.LogicFilterRules, FixedaddresstemplateLogicFilterRulesAttrTypes, diags, FlattenFixedaddresstemplateLogicFilterRules)
	m.Name = flex.FlattenStringPointerEmptyAsNull(from.Name)
	m.Nextserver = flex.FlattenStringPointerEmptyAsNull(from.Nextserver)
	m.NumberOfAddresses = flex.FlattenInt64Pointer(from.NumberOfAddresses)
	m.Offset = flex.FlattenInt64Pointer(from.Offset)
	m.Options = flex.FlattenFrameworkListNestedBlock(ctx, from.Options, FixedaddresstemplateOptionsAttrTypes, diags, FlattenFixedaddresstemplateOptions)
	m.PxeLeaseTime = flex.FlattenInt64Pointer(from.PxeLeaseTime)
}
