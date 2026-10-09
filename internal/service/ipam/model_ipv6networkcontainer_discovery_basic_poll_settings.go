package ipam

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	niosipam "github.com/infobloxopen/infoblox-nios-go-client/ipam"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
	customvalidator "github.com/infobloxopen/terraform-provider-infoblox/internal/validator"
)

// Ipv6networkcontainerDiscoveryBasicPollSettingsModel is the Terraform model for Ipv6networkcontainerDiscoveryBasicPollSettings
type Ipv6networkcontainerDiscoveryBasicPollSettingsModel struct {
	PortScanning                            types.Bool   `tfsdk:"port_scanning"`
	DeviceProfile                           types.Bool   `tfsdk:"device_profile"`
	SnmpCollection                          types.Bool   `tfsdk:"snmp_collection"`
	CliCollection                           types.Bool   `tfsdk:"cli_collection"`
	NetbiosScanning                         types.Bool   `tfsdk:"netbios_scanning"`
	CompletePingSweep                       types.Bool   `tfsdk:"complete_ping_sweep"`
	SmartSubnetPingSweep                    types.Bool   `tfsdk:"smart_subnet_ping_sweep"`
	AutoArpRefreshBeforeSwitchPortPolling   types.Bool   `tfsdk:"auto_arp_refresh_before_switch_port_polling"`
	SwitchPortDataCollectionPolling         types.String `tfsdk:"switch_port_data_collection_polling"`
	SwitchPortDataCollectionPollingSchedule types.Object `tfsdk:"switch_port_data_collection_polling_schedule"`
	SwitchPortDataCollectionPollingInterval types.Int64  `tfsdk:"switch_port_data_collection_polling_interval"`
	CredentialGroup                         types.String `tfsdk:"credential_group"`
	PollingFrequencyModifier                types.String `tfsdk:"polling_frequency_modifier"`
	UseGlobalPollingFrequencyModifier       types.Bool   `tfsdk:"use_global_polling_frequency_modifier"`
}

// Ipv6networkcontainerDiscoveryBasicPollSettingsAttrTypes contains the attribute types for Ipv6networkcontainerDiscoveryBasicPollSettingsModel
var Ipv6networkcontainerDiscoveryBasicPollSettingsAttrTypes = map[string]attr.Type{
	"port_scanning":                                types.BoolType,
	"device_profile":                               types.BoolType,
	"snmp_collection":                              types.BoolType,
	"cli_collection":                               types.BoolType,
	"netbios_scanning":                             types.BoolType,
	"complete_ping_sweep":                          types.BoolType,
	"smart_subnet_ping_sweep":                      types.BoolType,
	"auto_arp_refresh_before_switch_port_polling":  types.BoolType,
	"switch_port_data_collection_polling":          types.StringType,
	"switch_port_data_collection_polling_schedule": types.ObjectType{AttrTypes: Ipv6networkcontainerdiscoverybasicpollsettingsSwitchPortDataCollectionPollingScheduleAttrTypes},
	"switch_port_data_collection_polling_interval": types.Int64Type,
	"credential_group":                             types.StringType,
	"polling_frequency_modifier":                   types.StringType,
	"use_global_polling_frequency_modifier":        types.BoolType,
}

// Ipv6networkcontainerDiscoveryBasicPollSettingsResourceSchemaAttributes contains the schema attributes for Ipv6networkcontainerDiscoveryBasicPollSettingsModel
var Ipv6networkcontainerDiscoveryBasicPollSettingsResourceSchemaAttributes = map[string]schema.Attribute{
	"port_scanning": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Determines whether port scanning is enabled or not.",
	},
	"device_profile": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Determines whether device profile is enabled or not.",
	},
	"snmp_collection": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(true),
		MarkdownDescription: "Determines whether SNMP collection is enabled or not.",
	},
	"cli_collection": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(true),
		MarkdownDescription: "Determines whether CLI collection is enabled or not.",
	},
	"netbios_scanning": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Determines whether netbios scanning is enabled or not.",
	},
	"complete_ping_sweep": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Determines whether complete ping sweep is enabled or not.",
	},
	"smart_subnet_ping_sweep": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Determines whether smart subnet ping sweep is enabled or not.",
	},
	"auto_arp_refresh_before_switch_port_polling": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(true),
		MarkdownDescription: "Determines whether auto ARP refresh before switch port polling is enabled or not.",
	},
	"switch_port_data_collection_polling": schema.StringAttribute{
		Default: stringdefault.StaticString("PERIODIC"),
		Validators: []validator.String{
			stringvalidator.OneOf("DISABLED", "PERIODIC", "SCHEDULED"),
		},
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "A switch port data collection polling mode.",
	},
	"switch_port_data_collection_polling_schedule": schema.SingleNestedAttribute{
		Attributes:          Ipv6networkcontainerdiscoverybasicpollsettingsSwitchPortDataCollectionPollingScheduleResourceSchemaAttributes,
		Optional:            true,
		MarkdownDescription: "A Schedule Setting struct that determines switch port data collection polling schedule.",
	},
	"switch_port_data_collection_polling_interval": schema.Int64Attribute{
		Optional:            true,
		Computed:            true,
		Default:             int64default.StaticInt64(3600),
		MarkdownDescription: "Indicates the interval for switch port data collection polling.",
	},
	"credential_group": schema.StringAttribute{
		Default:  stringdefault.StaticString("default"),
		Optional: true,
		Computed: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
		},
		MarkdownDescription: "Credential group.",
	},
	"polling_frequency_modifier": schema.StringAttribute{
		Default:  stringdefault.StaticString("1"),
		Optional: true,
		Computed: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
		},
		MarkdownDescription: "Polling Frequency Modifier.",
	},
	"use_global_polling_frequency_modifier": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(true),
		MarkdownDescription: "Use Global Polling Frequency Modifier.",
	},
}

// ExpandIpv6networkcontainerDiscoveryBasicPollSettings converts a Terraform Object to SDK type
func ExpandIpv6networkcontainerDiscoveryBasicPollSettings(ctx context.Context, o types.Object, diags *diag.Diagnostics) *niosipam.Ipv6networkcontainerDiscoveryBasicPollSettings {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m Ipv6networkcontainerDiscoveryBasicPollSettingsModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

// Expand converts the Terraform model to SDK type
func (m *Ipv6networkcontainerDiscoveryBasicPollSettingsModel) Expand(ctx context.Context, diags *diag.Diagnostics) *niosipam.Ipv6networkcontainerDiscoveryBasicPollSettings {
	if m == nil {
		return nil
	}
	to := &niosipam.Ipv6networkcontainerDiscoveryBasicPollSettings{
		PortScanning:                            flex.ExpandBoolPointer(m.PortScanning),
		DeviceProfile:                           flex.ExpandBoolPointer(m.DeviceProfile),
		SnmpCollection:                          flex.ExpandBoolPointer(m.SnmpCollection),
		CliCollection:                           flex.ExpandBoolPointer(m.CliCollection),
		NetbiosScanning:                         flex.ExpandBoolPointer(m.NetbiosScanning),
		CompletePingSweep:                       flex.ExpandBoolPointer(m.CompletePingSweep),
		SmartSubnetPingSweep:                    flex.ExpandBoolPointer(m.SmartSubnetPingSweep),
		AutoArpRefreshBeforeSwitchPortPolling:   flex.ExpandBoolPointer(m.AutoArpRefreshBeforeSwitchPortPolling),
		SwitchPortDataCollectionPolling:         flex.ExpandStringPointerNullAsEmpty(m.SwitchPortDataCollectionPolling),
		SwitchPortDataCollectionPollingSchedule: ExpandIpv6networkcontainerdiscoverybasicpollsettingsSwitchPortDataCollectionPollingSchedule(ctx, m.SwitchPortDataCollectionPollingSchedule, diags),
		SwitchPortDataCollectionPollingInterval: flex.ExpandInt64Pointer(m.SwitchPortDataCollectionPollingInterval),
		CredentialGroup:                         flex.ExpandStringPointerNullAsEmpty(m.CredentialGroup),
		PollingFrequencyModifier:                flex.ExpandStringPointerNullAsEmpty(m.PollingFrequencyModifier),
		UseGlobalPollingFrequencyModifier:       flex.ExpandBoolPointer(m.UseGlobalPollingFrequencyModifier),
	}
	return to
}

// FlattenIpv6networkcontainerDiscoveryBasicPollSettings converts an SDK type to Terraform Object
func FlattenIpv6networkcontainerDiscoveryBasicPollSettings(ctx context.Context, from *niosipam.Ipv6networkcontainerDiscoveryBasicPollSettings, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(Ipv6networkcontainerDiscoveryBasicPollSettingsAttrTypes)
	}
	m := &Ipv6networkcontainerDiscoveryBasicPollSettingsModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, Ipv6networkcontainerDiscoveryBasicPollSettingsAttrTypes, m)
	diags.Append(d...)
	return t
}

// Flatten populates the Terraform model from SDK type
func (m *Ipv6networkcontainerDiscoveryBasicPollSettingsModel) Flatten(ctx context.Context, from *niosipam.Ipv6networkcontainerDiscoveryBasicPollSettings, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	m.PortScanning = flex.FlattenBoolPointer(from.PortScanning)
	m.DeviceProfile = flex.FlattenBoolPointer(from.DeviceProfile)
	m.SnmpCollection = flex.FlattenBoolPointer(from.SnmpCollection)
	m.CliCollection = flex.FlattenBoolPointer(from.CliCollection)
	m.NetbiosScanning = flex.FlattenBoolPointer(from.NetbiosScanning)
	m.CompletePingSweep = flex.FlattenBoolPointer(from.CompletePingSweep)
	m.SmartSubnetPingSweep = flex.FlattenBoolPointer(from.SmartSubnetPingSweep)
	m.AutoArpRefreshBeforeSwitchPortPolling = flex.FlattenBoolPointer(from.AutoArpRefreshBeforeSwitchPortPolling)
	m.SwitchPortDataCollectionPolling = flex.FlattenStringPointerEmptyAsNull(from.SwitchPortDataCollectionPolling)
	m.SwitchPortDataCollectionPollingSchedule = FlattenIpv6networkcontainerdiscoverybasicpollsettingsSwitchPortDataCollectionPollingSchedule(ctx, from.SwitchPortDataCollectionPollingSchedule, diags)
	m.SwitchPortDataCollectionPollingInterval = flex.FlattenInt64Pointer(from.SwitchPortDataCollectionPollingInterval)
	m.CredentialGroup = flex.FlattenStringPointerEmptyAsNull(from.CredentialGroup)
	m.PollingFrequencyModifier = flex.FlattenStringPointerEmptyAsNull(from.PollingFrequencyModifier)
	m.UseGlobalPollingFrequencyModifier = flex.FlattenBoolPointer(from.UseGlobalPollingFrequencyModifier)
}
