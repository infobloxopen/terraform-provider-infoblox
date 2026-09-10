package dns

import (
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"

	coremodel "github.com/infobloxopen/terraform-provider-infoblox/internal/core/model/dns"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
	internaltypes "github.com/infobloxopen/terraform-provider-infoblox/internal/types"
	customvalidator "github.com/infobloxopen/terraform-provider-infoblox/internal/validator"
)

// Values of match_client that decide which identifier an IPv6 address leases on.
const (
	matchClientDuid = "DUID"
	matchClientMac  = "MAC_ADDRESS"
)

type IPAssociationModel struct {
	NIOS types.Object `tfsdk:"nios"`
}

var IPAssociationAttrTypes = map[string]attr.Type{
	"nios": types.ObjectType{AttrTypes: NIOSIPAssociationAttrTypes},
}

var IPAssociationResourceSchemaAttributes = map[string]schema.Attribute{
	"nios": schema.SingleNestedAttribute{
		Required:            true,
		MarkdownDescription: "NIOS backend-specific fields.",
		Attributes:          IPAssociationResourceNiosSchemaAttributes,
	},
}

type NIOSIPAssociationModel struct {
	RecordHostId     types.String             `tfsdk:"record_host_id"`
	InternalID       types.String             `tfsdk:"internal_id"`
	MacAddr          internaltypes.MACAddress `tfsdk:"mac"`
	Duid             internaltypes.DUIDValue  `tfsdk:"duid"`
	ConfigureForDhcp types.Bool               `tfsdk:"configure_for_dhcp"`
	MatchClient      types.String             `tfsdk:"match_client"`
}

var NIOSIPAssociationAttrTypes = map[string]attr.Type{
	"record_host_id":     types.StringType,
	"internal_id":        types.StringType,
	"mac":                internaltypes.MACAddressType{},
	"duid":               internaltypes.DUIDType{},
	"configure_for_dhcp": types.BoolType,
	"match_client":       types.StringType,
}

var IPAssociationResourceNiosSchemaAttributes = map[string]schema.Attribute{
	"record_host_id": schema.StringAttribute{
		Required:            true,
		MarkdownDescription: "Reference to the host record to associate, normally `infoblox_record_host.<name>.id`. The host record must already exist; this resource never creates or destroys one.",
	},
	"internal_id": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "Internal ID of the associated host record.",
	},
	"mac": schema.StringAttribute{
		CustomType:          internaltypes.MACAddressType{},
		Optional:            true,
		Computed:            true,
		Default:             stringdefault.StaticString(""),
		MarkdownDescription: "The MAC address to lease the address to.",
		Validators: []validator.String{
			customvalidator.IsValidMacAddress(),
		},
	},
	"duid": schema.StringAttribute{
		CustomType:          internaltypes.DUIDType{},
		Optional:            true,
		Computed:            true,
		Default:             stringdefault.StaticString(""),
		MarkdownDescription: "The DUID to lease the IPv6 address to.",
		Validators: []validator.String{
			customvalidator.IsValidDUID(),
		},
	},
	"configure_for_dhcp": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Set to true to enable the DHCP configuration for the associated addresses.",
	},
	"match_client": schema.StringAttribute{
		Optional:            true,
		Computed:            true,
		Default:             stringdefault.StaticString(matchClientDuid),
		MarkdownDescription: "Which identifier the IPv6 address is leased to: `DUID` matches the DUID, `MAC_ADDRESS` matches the MAC address.",
		Validators: []validator.String{
			stringvalidator.OneOf(matchClientDuid, matchClientMac),
		},
	},
}

// Expand writes this association's DHCP settings onto a host record read from
// NIOS, leaving everything else on it untouched - the host record's own lifecycle
// and DNS settings belong to infoblox_record_host.
func (m *NIOSIPAssociationModel) Expand(host *coremodel.RecordHost) *coremodel.RecordHost {
	if m == nil || host == nil || host.NIOS == nil {
		return host
	}

	if len(host.NIOS.Ipv4addrs) > 0 {
		v4 := &host.NIOS.Ipv4addrs[0]
		v4.Mac = flex.ExpandStringPointer(m.MacAddr.StringValue)
		v4.ConfigureForDhcp = flex.ExpandBoolPointer(m.ConfigureForDhcp)
	}

	if len(host.NIOS.Ipv6addrs) > 0 {
		v6 := &host.NIOS.Ipv6addrs[0]
		switch m.MatchClient.ValueString() {
		case matchClientDuid:
			v6.Duid = flex.ExpandStringPointer(m.Duid.StringValue)
			v6.Mac = nil
		case matchClientMac:
			v6.Mac = flex.ExpandStringPointer(m.MacAddr.StringValue)
			v6.Duid = nil
		}
		if m.ConfigureForDhcp.ValueBool() && m.MatchClient.ValueString() != "" {
			v6.MatchClient = m.MatchClient.ValueStringPointer()
		}
		v6.ConfigureForDhcp = flex.ExpandBoolPointer(m.ConfigureForDhcp)
	}

	stripReadOnlyForAssociation(host)
	return host
}

// Flatten reads the DHCP settings back off the host record. Attributes the record
// does not carry are left as they were, so an address family the user did not
// configure keeps its planned value rather than reverting to null.
func (m *NIOSIPAssociationModel) Flatten(host *coremodel.RecordHost) {
	if m == nil || host == nil || host.NIOS == nil {
		return
	}

	if len(host.NIOS.Ipv4addrs) > 0 {
		v4 := host.NIOS.Ipv4addrs[0]
		if v4.Mac != nil {
			m.MacAddr = internaltypes.NewMACAddressValue(*v4.Mac)
		}
		if v4.ConfigureForDhcp != nil {
			m.ConfigureForDhcp = types.BoolValue(*v4.ConfigureForDhcp)
		}
		if v4.MatchClient != nil {
			m.MatchClient = types.StringValue(*v4.MatchClient)
		}
	}

	if len(host.NIOS.Ipv6addrs) > 0 {
		v6 := host.NIOS.Ipv6addrs[0]
		if v6.Duid != nil {
			m.Duid = internaltypes.NewDUIDValue(*v6.Duid)
		}
		if v6.Mac != nil {
			m.MacAddr = internaltypes.NewMACAddressValue(*v6.Mac)
		}
		if v6.ConfigureForDhcp != nil {
			m.ConfigureForDhcp = types.BoolValue(*v6.ConfigureForDhcp)
		}
		if v6.MatchClient != nil {
			m.MatchClient = types.StringValue(*v6.MatchClient)
		}
	}
}

// stripReadOnlyForAssociation clears fields NIOS rejects on a host record update.
// The core model already omits most read-only fields - these are the ones it keeps
// because they are writable on other paths.
func stripReadOnlyForAssociation(host *coremodel.RecordHost) {
	host.NIOS.CloudInfo = nil
	host.NIOS.DnsAliases = nil
	host.NIOS.NetworkView = nil

	if len(host.NIOS.Ipv4addrs) > 0 {
		host.NIOS.Ipv4addrs[0].Host = nil
	}
	if len(host.NIOS.Ipv6addrs) > 0 {
		host.NIOS.Ipv6addrs[0].Host = nil
	}

	if host.NIOS.ConfigureForDns != nil && !*host.NIOS.ConfigureForDns {
		host.NIOS.Name = nil
		host.NIOS.View = nil
	}
}
