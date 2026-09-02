package dns

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/mapvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	int64planmodifier "github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/mapdefault"
	objectplanmodifier "github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	stringplanmodifier "github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
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

type ZoneAuthModel struct {
	Id   types.String `tfsdk:"id"`
	NIOS types.Object `tfsdk:"nios"`
	UDDI types.Object `tfsdk:"uddi"`
}

var ZoneAuthAttrTypes = map[string]attr.Type{
	"id":   types.StringType,
	"nios": types.ObjectType{AttrTypes: NIOSZoneAuthAttrTypes},
	"uddi": types.ObjectType{AttrTypes: UDDIZoneAuthAttrTypes},
}

type NIOSZoneAuthModel struct {
	AllowActiveDir                   types.List                          `tfsdk:"allow_active_dir"`
	AllowFixedRrsetOrder             types.Bool                          `tfsdk:"allow_fixed_rrset_order"`
	AllowGssTsigForUnderscoreZone    types.Bool                          `tfsdk:"allow_gss_tsig_for_underscore_zone"`
	AllowGssTsigZoneUpdates          types.Bool                          `tfsdk:"allow_gss_tsig_zone_updates"`
	AllowQuery                       types.List                          `tfsdk:"allow_query"`
	AllowTransfer                    types.List                          `tfsdk:"allow_transfer"`
	AllowUpdate                      types.List                          `tfsdk:"allow_update"`
	AllowUpdateForwarding            types.Bool                          `tfsdk:"allow_update_forwarding"`
	Comment                          types.String                        `tfsdk:"comment"`
	CopyXferToNotify                 types.Bool                          `tfsdk:"copy_xfer_to_notify"`
	CreateUnderscoreZones            types.Bool                          `tfsdk:"create_underscore_zones"`
	DdnsForceCreationTimestampUpdate types.Bool                          `tfsdk:"ddns_force_creation_timestamp_update"`
	DdnsPrincipalGroup               types.String                        `tfsdk:"ddns_principal_group"`
	DdnsPrincipalTracking            types.Bool                          `tfsdk:"ddns_principal_tracking"`
	DdnsRestrictPatterns             types.Bool                          `tfsdk:"ddns_restrict_patterns"`
	DdnsRestrictPatternsList         internaltypes.UnorderedListValue    `tfsdk:"ddns_restrict_patterns_list"`
	DdnsRestrictProtected            types.Bool                          `tfsdk:"ddns_restrict_protected"`
	DdnsRestrictSecure               types.Bool                          `tfsdk:"ddns_restrict_secure"`
	DdnsRestrictStatic               types.Bool                          `tfsdk:"ddns_restrict_static"`
	Disable                          types.Bool                          `tfsdk:"disable"`
	DisableForwarding                types.Bool                          `tfsdk:"disable_forwarding"`
	DisplayDomain                    types.String                        `tfsdk:"display_domain"`
	DnsIntegrityEnable               types.Bool                          `tfsdk:"dns_integrity_enable"`
	DnsIntegrityFrequency            types.Int64                         `tfsdk:"dns_integrity_frequency"`
	DnsIntegrityMember               types.String                        `tfsdk:"dns_integrity_member"`
	DnsIntegrityVerboseLogging       types.Bool                          `tfsdk:"dns_integrity_verbose_logging"`
	DnssecKeyParams                  types.Object                        `tfsdk:"dnssec_key_params"`
	DnssecKeys                       types.List                          `tfsdk:"dnssec_keys"`
	EffectiveCheckNamesPolicy        types.String                        `tfsdk:"effective_check_names_policy"`
	ExtAttrs                         types.Map                           `tfsdk:"ext_attrs"`
	ExtAttrsAll                      types.Map                           `tfsdk:"ext_attrs_all"`
	ExternalPrimaries                types.List                          `tfsdk:"external_primaries"`
	ExternalSecondaries              types.List                          `tfsdk:"external_secondaries"`
	Fqdn                             types.String                        `tfsdk:"fqdn"`
	GridPrimary                      types.List                          `tfsdk:"grid_primary"`
	GridSecondaries                  types.List                          `tfsdk:"grid_secondaries"`
	LastQueriedAcl                   types.List                          `tfsdk:"last_queried_acl"`
	Locked                           types.Bool                          `tfsdk:"locked"`
	MemberSoaMnames                  types.List                          `tfsdk:"member_soa_mnames"`
	MsAdIntegrated                   types.Bool                          `tfsdk:"ms_ad_integrated"`
	MsAllowTransfer                  types.List                          `tfsdk:"ms_allow_transfer"`
	MsAllowTransferMode              types.String                        `tfsdk:"ms_allow_transfer_mode"`
	MsDcNsRecordCreation             types.List                          `tfsdk:"ms_dc_ns_record_creation"`
	MsDdnsMode                       types.String                        `tfsdk:"ms_ddns_mode"`
	MsPrimaries                      types.List                          `tfsdk:"ms_primaries"`
	MsSecondaries                    types.List                          `tfsdk:"ms_secondaries"`
	MsSyncDisabled                   types.Bool                          `tfsdk:"ms_sync_disabled"`
	NotifyDelay                      types.Int64                         `tfsdk:"notify_delay"`
	NsGroup                          types.String                        `tfsdk:"ns_group"`
	Prefix                           internaltypes.CaseInsensitiveString `tfsdk:"prefix"`
	RecordNamePolicy                 types.String                        `tfsdk:"record_name_policy"`
	RemoveSubzones                   types.Bool                          `tfsdk:"remove_subzones"`
	RestartIfNeeded                  types.Bool                          `tfsdk:"restart_if_needed"`
	ScavengingSettings               types.Object                        `tfsdk:"scavenging_settings"`
	SetSoaSerialNumber               types.Bool                          `tfsdk:"set_soa_serial_number"`
	SoaDefaultTtl                    types.Int64                         `tfsdk:"soa_default_ttl"`
	SoaEmail                         types.String                        `tfsdk:"soa_email"`
	SoaExpire                        types.Int64                         `tfsdk:"soa_expire"`
	SoaNegativeTtl                   types.Int64                         `tfsdk:"soa_negative_ttl"`
	SoaRefresh                       types.Int64                         `tfsdk:"soa_refresh"`
	SoaRetry                         types.Int64                         `tfsdk:"soa_retry"`
	SoaSerialNumber                  types.Int64                         `tfsdk:"soa_serial_number"`
	Srgs                             types.List                          `tfsdk:"srgs"`
	UpdateForwarding                 types.List                          `tfsdk:"update_forwarding"`
	UseCheckNamesPolicy              types.Bool                          `tfsdk:"use_check_names_policy"`
	UseExternalPrimary               types.Bool                          `tfsdk:"use_external_primary"`
	UseImportFrom                    types.Bool                          `tfsdk:"use_import_from"`
	View                             types.String                        `tfsdk:"view"`
	ZoneFormat                       types.String                        `tfsdk:"zone_format"`
}

var NIOSZoneAuthAttrTypes = map[string]attr.Type{
	"allow_active_dir":                     types.ListType{ElemType: types.ObjectType{AttrTypes: ZoneAuthAllowActiveDirAttrTypes}},
	"allow_fixed_rrset_order":              types.BoolType,
	"allow_gss_tsig_for_underscore_zone":   types.BoolType,
	"allow_gss_tsig_zone_updates":          types.BoolType,
	"allow_query":                          types.ListType{ElemType: types.ObjectType{AttrTypes: ZoneAuthAllowQueryAttrTypes}},
	"allow_transfer":                       types.ListType{ElemType: types.ObjectType{AttrTypes: ZoneAuthAllowTransferAttrTypes}},
	"allow_update":                         types.ListType{ElemType: types.ObjectType{AttrTypes: ZoneAuthAllowUpdateAttrTypes}},
	"allow_update_forwarding":              types.BoolType,
	"comment":                              types.StringType,
	"copy_xfer_to_notify":                  types.BoolType,
	"create_underscore_zones":              types.BoolType,
	"ddns_force_creation_timestamp_update": types.BoolType,
	"ddns_principal_group":                 types.StringType,
	"ddns_principal_tracking":              types.BoolType,
	"ddns_restrict_patterns":               types.BoolType,
	"ddns_restrict_patterns_list":          internaltypes.UnorderedListOfStringType,
	"ddns_restrict_protected":              types.BoolType,
	"ddns_restrict_secure":                 types.BoolType,
	"ddns_restrict_static":                 types.BoolType,
	"disable":                              types.BoolType,
	"disable_forwarding":                   types.BoolType,
	"display_domain":                       types.StringType,
	"dns_integrity_enable":                 types.BoolType,
	"dns_integrity_frequency":              types.Int64Type,
	"dns_integrity_member":                 types.StringType,
	"dns_integrity_verbose_logging":        types.BoolType,
	"dnssec_key_params":                    types.ObjectType{AttrTypes: ZoneAuthDnssecKeyParamsAttrTypes},
	"dnssec_keys":                          types.ListType{ElemType: types.ObjectType{AttrTypes: ZoneAuthDnssecKeysAttrTypes}},
	"effective_check_names_policy":         types.StringType,
	"ext_attrs":                            types.MapType{ElemType: types.StringType},
	"ext_attrs_all":                        types.MapType{ElemType: types.StringType},
	"external_primaries":                   types.ListType{ElemType: types.ObjectType{AttrTypes: ZoneAuthExternalPrimariesAttrTypes}},
	"external_secondaries":                 types.ListType{ElemType: types.ObjectType{AttrTypes: ZoneAuthExternalSecondariesAttrTypes}},
	"fqdn":                                 types.StringType,
	"grid_primary":                         types.ListType{ElemType: types.ObjectType{AttrTypes: ZoneAuthGridPrimaryAttrTypes}},
	"grid_secondaries":                     types.ListType{ElemType: types.ObjectType{AttrTypes: ZoneAuthGridSecondariesAttrTypes}},
	"last_queried_acl":                     types.ListType{ElemType: types.ObjectType{AttrTypes: ZoneAuthLastQueriedAclAttrTypes}},
	"locked":                               types.BoolType,
	"member_soa_mnames":                    types.ListType{ElemType: types.ObjectType{AttrTypes: ZoneAuthMemberSoaMnamesAttrTypes}},
	"ms_ad_integrated":                     types.BoolType,
	"ms_allow_transfer":                    types.ListType{ElemType: types.ObjectType{AttrTypes: ZoneAuthMsAllowTransferAttrTypes}},
	"ms_allow_transfer_mode":               types.StringType,
	"ms_dc_ns_record_creation":             types.ListType{ElemType: types.ObjectType{AttrTypes: ZoneAuthMsDcNsRecordCreationAttrTypes}},
	"ms_ddns_mode":                         types.StringType,
	"ms_primaries":                         types.ListType{ElemType: types.ObjectType{AttrTypes: ZoneAuthMsPrimariesAttrTypes}},
	"ms_secondaries":                       types.ListType{ElemType: types.ObjectType{AttrTypes: ZoneAuthMsSecondariesAttrTypes}},
	"ms_sync_disabled":                     types.BoolType,
	"notify_delay":                         types.Int64Type,
	"ns_group":                             types.StringType,
	"prefix":                               internaltypes.CaseInsensitiveStringType{},
	"record_name_policy":                   types.StringType,
	"remove_subzones":                      types.BoolType,
	"restart_if_needed":                    types.BoolType,
	"scavenging_settings":                  types.ObjectType{AttrTypes: ZoneAuthScavengingSettingsAttrTypes},
	"set_soa_serial_number":                types.BoolType,
	"soa_default_ttl":                      types.Int64Type,
	"soa_email":                            types.StringType,
	"soa_expire":                           types.Int64Type,
	"soa_negative_ttl":                     types.Int64Type,
	"soa_refresh":                          types.Int64Type,
	"soa_retry":                            types.Int64Type,
	"soa_serial_number":                    types.Int64Type,
	"srgs":                                 types.ListType{ElemType: types.StringType},
	"update_forwarding":                    types.ListType{ElemType: types.ObjectType{AttrTypes: ZoneAuthUpdateForwardingAttrTypes}},
	"use_check_names_policy":               types.BoolType,
	"use_external_primary":                 types.BoolType,
	"use_import_from":                      types.BoolType,
	"view":                                 types.StringType,
	"zone_format":                          types.StringType,
}

type UDDIZoneAuthModel struct {
	Comment                  types.String `tfsdk:"comment"`
	CompartmentId            types.String `tfsdk:"compartment_id"`
	Disabled                 types.Bool   `tfsdk:"disabled"`
	ExternalPrimaries        types.List   `tfsdk:"external_primaries"`
	ExternalSecondaries      types.List   `tfsdk:"external_secondaries"`
	Fqdn                     types.String `tfsdk:"fqdn"`
	GssTsigEnabled           types.Bool   `tfsdk:"gss_tsig_enabled"`
	InheritanceSources       types.Object `tfsdk:"inheritance_sources"`
	InitialSoaSerial         types.Int64  `tfsdk:"initial_soa_serial"`
	InternalSecondaries      types.List   `tfsdk:"internal_secondaries"`
	Notify                   types.Bool   `tfsdk:"notify"`
	Nsgs                     types.List   `tfsdk:"nsgs"`
	Parent                   types.String `tfsdk:"parent"`
	PrimaryType              types.String `tfsdk:"primary_type"`
	QueryAcl                 types.List   `tfsdk:"query_acl"`
	Tags                     types.Map    `tfsdk:"tags"`
	TagsAll                  types.Map    `tfsdk:"tags_all"`
	TransferAcl              types.List   `tfsdk:"transfer_acl"`
	UpdateAcl                types.List   `tfsdk:"update_acl"`
	UseForwardersForSubzones types.Bool   `tfsdk:"use_forwarders_for_subzones"`
	View                     types.String `tfsdk:"view"`
}

var UDDIZoneAuthAttrTypes = map[string]attr.Type{
	"comment":                     types.StringType,
	"compartment_id":              types.StringType,
	"disabled":                    types.BoolType,
	"external_primaries":          types.ListType{ElemType: types.ObjectType{AttrTypes: ExternalPrimaryAttrTypes}},
	"external_secondaries":        types.ListType{ElemType: types.ObjectType{AttrTypes: ExternalSecondaryAttrTypes}},
	"fqdn":                        types.StringType,
	"gss_tsig_enabled":            types.BoolType,
	"inheritance_sources":         types.ObjectType{AttrTypes: AuthZoneInheritanceAttrTypes},
	"initial_soa_serial":          types.Int64Type,
	"internal_secondaries":        types.ListType{ElemType: types.ObjectType{AttrTypes: InternalSecondaryAttrTypes}},
	"notify":                      types.BoolType,
	"nsgs":                        types.ListType{ElemType: types.StringType},
	"parent":                      types.StringType,
	"primary_type":                types.StringType,
	"query_acl":                   types.ListType{ElemType: types.ObjectType{AttrTypes: ACLItemAttrTypes}},
	"tags":                        types.MapType{ElemType: types.StringType},
	"tags_all":                    types.MapType{ElemType: types.StringType},
	"transfer_acl":                types.ListType{ElemType: types.ObjectType{AttrTypes: ACLItemAttrTypes}},
	"update_acl":                  types.ListType{ElemType: types.ObjectType{AttrTypes: ACLItemAttrTypes}},
	"use_forwarders_for_subzones": types.BoolType,
	"view":                        types.StringType,
}

const (
	ZoneAuthInheritanceType = "full"
	ZoneAuthReturnFields    = "address,allow_active_dir,allow_fixed_rrset_order,allow_gss_tsig_for_underscore_zone,allow_gss_tsig_zone_updates,allow_query,allow_transfer,allow_update,allow_update_forwarding,aws_rte53_zone_info,cloud_info,comment,copy_xfer_to_notify,create_underscore_zones,ddns_force_creation_timestamp_update,ddns_principal_group,ddns_principal_tracking,ddns_restrict_patterns,ddns_restrict_patterns_list,ddns_restrict_protected,ddns_restrict_secure,ddns_restrict_static,disable,disable_forwarding,display_domain,dns_fqdn,dns_integrity_enable,dns_integrity_frequency,dns_integrity_member,dns_integrity_verbose_logging,dns_soa_email,dnssec_key_params,dnssec_keys,dnssec_ksk_rollover_date,dnssec_zsk_rollover_date,effective_check_names_policy,effective_record_name_policy,extattrs,external_primaries,external_secondaries,fqdn,grid_primary,grid_primary_shared_with_ms_parent_delegation,grid_secondaries,is_dnssec_enabled,is_dnssec_signed,is_multimaster,last_queried,last_queried_acl,locked,locked_by,mask_prefix,member_soa_mnames,member_soa_serials,ms_ad_integrated,ms_allow_transfer,ms_allow_transfer_mode,ms_dc_ns_record_creation,ms_ddns_mode,ms_managed,ms_primaries,ms_read_only,ms_secondaries,ms_sync_disabled,ms_sync_master_name,network_associations,network_view,notify_delay,ns_group,parent,prefix,primary_type,record_name_policy,records_monitored,rr_not_queried_enabled_time,scavenging_settings,soa_default_ttl,soa_email,soa_expire,soa_negative_ttl,soa_refresh,soa_retry,soa_serial_number,srgs,update_forwarding,use_allow_active_dir,use_allow_query,use_allow_transfer,use_allow_update,use_allow_update_forwarding,use_check_names_policy,use_copy_xfer_to_notify,use_ddns_force_creation_timestamp_update,use_ddns_patterns_restriction,use_ddns_principal_security,use_ddns_restrict_protected,use_ddns_restrict_static,use_dnssec_key_params,use_external_primary,use_grid_zone_timer,use_import_from,use_notify_delay,use_record_name_policy,use_scavenging_settings,use_soa_email,using_srg_associations,view,zone_format,zone_not_queried_enabled_time"
)

var ZoneAuthResourceSchemaAttributes = map[string]schema.Attribute{
	"id": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The reference to the object.",
	},
	"nios": schema.SingleNestedAttribute{
		Optional:            true,
		MarkdownDescription: "NIOS backend-specific fields.",
		Attributes:          ZoneAuthResourceNiosSchemaAttributes,
	},
	"uddi": schema.SingleNestedAttribute{
		Optional:            true,
		MarkdownDescription: "UDDI backend-specific fields.",
		Attributes:          ZoneAuthResourceUddiSchemaAttributes,
	},
}

var ZoneAuthResourceNiosSchemaAttributes = map[string]schema.Attribute{
	"allow_active_dir": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: ZoneAuthAllowActiveDirResourceSchemaAttributes,
		},
		Optional: true,
		Validators: []validator.List{
			customvalidator.ListNotEmpty(),
		},
		MarkdownDescription: "This field allows the zone to receive GSS-TSIG authenticated DDNS updates from DHCP clients and servers in an AD domain. Note that addresses specified in this field ignore the permission set in the struct which will be set to 'ALLOW'.",
	},
	"allow_fixed_rrset_order": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "The flag that allows to enable or disable fixed RRset ordering for authoritative forward-mapping zones.",
	},
	"allow_gss_tsig_for_underscore_zone": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "The flag that allows DHCP clients to perform GSS-TSIG signed updates for underscore zones.",
	},
	"allow_gss_tsig_zone_updates": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "The flag that enables or disables the zone for GSS-TSIG updates.",
	},
	"allow_query": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: ZoneAuthAllowQueryResourceSchemaAttributes,
		},
		Optional: true,
		Computed: true,
		Default:  listdefault.StaticValue(types.ListNull(types.ObjectType{AttrTypes: ZoneAuthAllowQueryAttrTypes})),
		Validators: []validator.List{
			customvalidator.ListNotEmpty(),
		},
		MarkdownDescription: "Determines whether DNS queries are allowed from a named ACL, or from a list of IPv4/IPv6 addresses, networks, and TSIG keys for the hosts.",
	},
	"allow_transfer": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: ZoneAuthAllowTransferResourceSchemaAttributes,
		},
		Optional: true,
		Computed: true,
		Default:  listdefault.StaticValue(types.ListNull(types.ObjectType{AttrTypes: ZoneAuthAllowTransferAttrTypes})),
		Validators: []validator.List{
			customvalidator.ListNotEmpty(),
		},
		MarkdownDescription: "Determines whether zone transfers are allowed from a named ACL, or from a list of IPv4/IPv6 addresses, networks, and TSIG keys for the hosts.",
	},
	"allow_update": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: ZoneAuthAllowUpdateResourceSchemaAttributes,
		},
		Optional: true,
		Computed: true,
		Default:  listdefault.StaticValue(types.ListNull(types.ObjectType{AttrTypes: ZoneAuthAllowUpdateAttrTypes})),
		Validators: []validator.List{
			customvalidator.ListNotEmpty(),
		},
		MarkdownDescription: "Determines whether dynamic DNS updates are allowed from a named ACL, or from a list of IPv4/IPv6 addresses, networks, and TSIG keys for the hosts.",
	},
	"allow_update_forwarding": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "The list with IP addresses, networks or TSIG keys for clients, from which forwarded dynamic updates are allowed.",
	},
	"comment": schema.StringAttribute{
		Optional: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
			stringvalidator.LengthBetween(0, 256),
			customvalidator.ValidateTrimmedString(),
		},
		MarkdownDescription: "Comment for the zone; maximum 256 characters.",
	},
	"copy_xfer_to_notify": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "If this flag is set to True then copy allowed IPs from Allow Transfer to Also Notify.",
	},
	"create_underscore_zones": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Determines whether automatic creation of subzones is enabled or not.",
	},
	"ddns_force_creation_timestamp_update": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Defines whether creation timestamp of RR should be updated ' when DDNS update happens even if there is no change to ' the RR.",
	},
	"ddns_principal_group": schema.StringAttribute{
		Optional: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
		},
		MarkdownDescription: "The DDNS Principal cluster group name.",
	},
	"ddns_principal_tracking": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "The flag that indicates whether the DDNS principal track is enabled or disabled.",
	},
	"ddns_restrict_patterns": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "The flag that indicates whether an option to restrict DDNS update request based on FQDN patterns is enabled or disabled.",
	},
	"ddns_restrict_patterns_list": schema.ListAttribute{
		ElementType: types.StringType,
		Optional:    true,
		CustomType:  internaltypes.UnorderedListOfStringType,
		Validators: []validator.List{
			customvalidator.ListNotEmpty(),
		},
		MarkdownDescription: "The unordered list of restriction patterns for an option of to restrict DDNS updates based on FQDN patterns.",
	},
	"ddns_restrict_protected": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "The flag that indicates whether an option to restrict DDNS update request to protected resource records is enabled or disabled.",
	},
	"ddns_restrict_secure": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "The flag that indicates whether DDNS update request for principal other than target resource record's principal is restricted.",
	},
	"ddns_restrict_static": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "The flag that indicates whether an option to restrict DDNS update request to resource records which are marked as 'STATIC' is enabled or disabled.",
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
		MarkdownDescription: "Determines whether the name servers that host the zone should forward queries (ended with the domain name of the zone) to any configured forwarders.",
	},
	"display_domain": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The displayed name of the DNS zone.",
	},
	"dns_integrity_enable": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "If this is set to True, DNS integrity check is enabled for this zone.",
	},
	"dns_integrity_frequency": schema.Int64Attribute{
		Optional: true,
		Computed: true,
		Default:  int64default.StaticInt64(3600),
		Validators: []validator.Int64{
			int64validator.Between(0, 4294967295),
		},
		MarkdownDescription: "The frequency, in seconds, of DNS integrity checks for this zone.",
	},
	"dns_integrity_member": schema.StringAttribute{
		Optional: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
		},
		MarkdownDescription: "The Grid member that performs DNS integrity checks for this zone.",
	},
	"dns_integrity_verbose_logging": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "If this is set to True, more information is logged for DNS integrity checks for this zone.",
	},
	"dnssec_key_params": schema.SingleNestedAttribute{
		Attributes:          ZoneAuthDnssecKeyParamsResourceSchemaAttributes,
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "",
	},
	"dnssec_keys": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: ZoneAuthDnssecKeysResourceSchemaAttributes,
		},
		Optional: true,
		Validators: []validator.List{
			customvalidator.ListNotEmpty(),
		},
		MarkdownDescription: "A list of DNSSEC keys for the zone.",
	},
	"effective_check_names_policy": schema.StringAttribute{
		Default: stringdefault.StaticString("WARN"),
		Validators: []validator.String{
			stringvalidator.OneOf("FAIL", "WARN"),
		},
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The value of the check names policy, which indicates the action the appliance takes when it encounters host names that do not comply with the Strict Hostname Checking policy. This value applies only if the host name restriction policy is set to \"Strict Hostname Checking\".",
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
	"external_primaries": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: ZoneAuthExternalPrimariesResourceSchemaAttributes,
		},
		Optional: true,
		Validators: []validator.List{
			customvalidator.ListNotEmpty(),
		},
		MarkdownDescription: "The list of external primary servers.",
	},
	"external_secondaries": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: ZoneAuthExternalSecondariesResourceSchemaAttributes,
		},
		Optional: true,
		Validators: []validator.List{
			customvalidator.ListNotEmpty(),
		},
		MarkdownDescription: "The list of external secondary servers.",
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
	"grid_primary": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: ZoneAuthGridPrimaryResourceSchemaAttributes,
		},
		Optional: true,
		Validators: []validator.List{
			customvalidator.ListNotEmpty(),
		},
		MarkdownDescription: "The grid primary servers for this zone.",
	},
	"grid_secondaries": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: ZoneAuthGridSecondariesResourceSchemaAttributes,
		},
		Optional: true,
		Validators: []validator.List{
			customvalidator.ListNotEmpty(),
		},
		MarkdownDescription: "The list with Grid members that are secondary servers for this zone.",
	},
	"last_queried_acl": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: ZoneAuthLastQueriedAclResourceSchemaAttributes,
		},
		Optional: true,
		Validators: []validator.List{
			customvalidator.ListNotEmpty(),
		},
		MarkdownDescription: "Determines last queried ACL for the specified IPv4 or IPv6 addresses and networks in scavenging settings.",
	},
	"locked": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "If you enable this flag, other administrators cannot make conflicting changes. This is for administration purposes only. The zone will continue to serve DNS data even when it is locked.",
	},
	"member_soa_mnames": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: ZoneAuthMemberSoaMnamesResourceSchemaAttributes,
		},
		Optional: true,
		Computed: true,
		Validators: []validator.List{
			customvalidator.ListNotEmpty(),
		},
		MarkdownDescription: "The list of per-member SOA MNAME information.",
	},
	"ms_ad_integrated": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "The flag that determines whether Active Directory is integrated or not. This field is valid only when ms_managed is \"STUB\", \"AUTH_PRIMARY\", or \"AUTH_BOTH\".",
	},
	"ms_allow_transfer": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: ZoneAuthMsAllowTransferResourceSchemaAttributes,
		},
		Optional: true,
		Validators: []validator.List{
			customvalidator.ListNotEmpty(),
		},
		MarkdownDescription: "The list of DNS clients that are allowed to perform zone transfers from a Microsoft DNS server. This setting applies only to zones with Microsoft DNS servers that are either primary or secondary servers. This setting does not inherit any value from the Grid or from any member that defines an allow_transfer value. This setting does not apply to any grid member. Use the allow_transfer field to control which DNS clients are allowed to perform zone transfers on Grid members.",
	},
	"ms_allow_transfer_mode": schema.StringAttribute{
		Default: stringdefault.StaticString("NONE"),
		Validators: []validator.String{
			stringvalidator.OneOf("ADDRESS_AC", "ANY", "ANY_NS", "NONE"),
		},
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "Determines which DNS clients are allowed to perform zone transfers from a Microsoft DNS server. Valid values are: \"ADDRESS_AC\", to use ms_allow_transfer field for specifying IP addresses, networks and Transaction Signature (TSIG) keys for clients that are allowed to do zone transfers. \"ANY\", to allow any client. \"ANY_NS\", to allow only the nameservers listed in this zone. \"NONE\", to deny all zone transfer requests.",
	},
	"ms_dc_ns_record_creation": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: ZoneAuthMsDcNsRecordCreationResourceSchemaAttributes,
		},
		Optional: true,
		Validators: []validator.List{
			customvalidator.ListNotEmpty(),
		},
		MarkdownDescription: "The list of domain controllers that are allowed to create NS records for authoritative zones.",
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
	"ms_primaries": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: ZoneAuthMsPrimariesResourceSchemaAttributes,
		},
		Optional: true,
		Validators: []validator.List{
			customvalidator.ListNotEmpty(),
		},
		MarkdownDescription: "The list with the Microsoft DNS servers that are primary servers for the zone. Although a zone typically has just one primary name server, you can specify up to ten independent servers for a single zone.",
	},
	"ms_secondaries": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: ZoneAuthMsSecondariesResourceSchemaAttributes,
		},
		Optional: true,
		Validators: []validator.List{
			customvalidator.ListNotEmpty(),
		},
		MarkdownDescription: "The list with the Microsoft DNS servers that are secondary servers for the zone.",
	},
	"ms_sync_disabled": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "This flag controls whether this zone is synchronized with Microsoft DNS servers.",
	},
	"notify_delay": schema.Int64Attribute{
		Optional: true,
		Computed: true,
		Default:  int64default.StaticInt64(5),
		Validators: []validator.Int64{
			int64validator.Between(5, 86400),
		},
		MarkdownDescription: "The number of seconds in delay with which notify messages are sent to secondaries.",
	},
	"ns_group": schema.StringAttribute{
		Optional: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
		},
		MarkdownDescription: "The name server group that serves DNS for this zone.",
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
	"record_name_policy": schema.StringAttribute{
		Optional: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
		},
		MarkdownDescription: "The hostname policy for records under this zone.",
	},
	"remove_subzones": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Remove subzones delete option. Determines whether all child objects should be removed alongside with the parent zone or child objects should be assigned to another parental zone. By default child objects are deleted with the parent zone.",
	},
	"restart_if_needed": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Restarts the member service.",
	},
	"scavenging_settings": schema.SingleNestedAttribute{
		Attributes:          ZoneAuthScavengingSettingsResourceSchemaAttributes,
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "",
	},
	"set_soa_serial_number": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "The serial number in the SOA record incrementally changes every time the record is modified. The Infoblox appliance allows you to change the serial number (in the SOA record) for the primary server so it is higher than the secondary server, thereby ensuring zone transfers come from the primary server (as they should). To change the serial number you need to set a new value at \"soa_serial_number\" and pass \"set_soa_serial_number\" as True.",
	},
	"soa_default_ttl": schema.Int64Attribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The Time to Live (TTL) value of the SOA record of this zone. This value is the number of seconds that data is cached.",
	},
	"soa_email": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
			customvalidator.ValidateTrimmedString(),
		},
		MarkdownDescription: "The SOA email value for this zone. This value can be in unicode format.",
	},
	"soa_expire": schema.Int64Attribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "This setting defines the amount of time, in seconds, after which the secondary server stops giving out answers about the zone because the zone data is too old to be useful. The default is one week.",
	},
	"soa_negative_ttl": schema.Int64Attribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The negative Time to Live (TTL) value of the SOA of the zone indicates how long a secondary server can cache data for \"Does Not Respond\" responses.",
	},
	"soa_refresh": schema.Int64Attribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "This indicates the interval at which a secondary server sends a message to the primary server for a zone to check that its data is current, and retrieve fresh data if it is not.",
	},
	"soa_retry": schema.Int64Attribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "This indicates how long a secondary server must wait before attempting to recontact the primary server after a connection failure between the two servers occurs.",
	},
	"soa_serial_number": schema.Int64Attribute{
		Optional: true,
		Computed: true,
		Validators: []validator.Int64{
			int64validator.AlsoRequires(path.MatchRelative().AtParent().AtName("set_soa_serial_number")),
		},
		MarkdownDescription: "The serial number in the SOA record incrementally changes every time the record is modified. The Infoblox appliance allows you to change the serial number (in the SOA record) for the primary server so it is higher than the secondary server, thereby ensuring zone transfers come from the primary server (as they should). To change the serial number you need to set a new value at \"soa_serial_number\" and pass \"set_soa_serial_number\" as True.",
	},
	"srgs": schema.ListAttribute{
		ElementType: types.StringType,
		Optional:    true,
		Validators: []validator.List{
			customvalidator.ListNotEmpty(),
		},
		MarkdownDescription: "The associated shared record groups of a DNS zone. If a shared record group is associated with a zone, then all shared records in a shared record group will be shared in the zone.",
	},
	"update_forwarding": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: ZoneAuthUpdateForwardingResourceSchemaAttributes,
		},
		Optional: true,
		Validators: []validator.List{
			customvalidator.ListNotEmpty(),
			listvalidator.AlsoRequires(path.MatchRelative().AtParent().AtName("allow_update_forwarding")),
		},
		MarkdownDescription: "Use this field to allow or deny dynamic DNS updates that are forwarded from specific IPv4/IPv6 addresses, networks, or a named ACL. You can also provide TSIG keys for clients that are allowed or denied to perform zone updates. This setting overrides the member-level setting.",
	},
	"use_check_names_policy": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Apply policy to dynamic updates and inbound zone transfers (This value applies only if the host name restriction policy is set to \"Strict Hostname Checking\".)",
	},
	"use_external_primary": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "This flag controls whether the zone is using an external primary.",
	},
	"use_import_from": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Use flag for: import_from",
	},
	"view": schema.StringAttribute{
		Default:  stringdefault.StaticString("default"),
		Optional: true,
		Computed: true,
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

var ZoneAuthResourceUddiSchemaAttributes = map[string]schema.Attribute{
	"comment": schema.StringAttribute{
		Default:             stringdefault.StaticString(""),
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "Optional. Comment for zone configuration.",
	},
	"compartment_id": schema.StringAttribute{
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
	"external_primaries": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: ExternalPrimaryResourceSchemaAttributes,
		},
		Optional: true,
		Validators: []validator.List{
			customvalidator.ListNotEmpty(),
		},
		MarkdownDescription: "Optional. DNS primaries external to BloxOne DDI. Order is not significant.",
	},
	"external_secondaries": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: ExternalSecondaryResourceSchemaAttributes,
		},
		Optional: true,
		Validators: []validator.List{
			customvalidator.ListNotEmpty(),
		},
		MarkdownDescription: "DNS secondaries external to BloxOne DDI. Order is not significant.",
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
	"gss_tsig_enabled": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "_gss_tsig_enabled_ enables/disables GSS-TSIG signed dynamic updates.  Defaults to _false_.",
	},
	"inheritance_sources": schema.SingleNestedAttribute{
		Attributes: AuthZoneInheritanceResourceSchemaAttributes,
		Optional:   true,
		Computed:   true,
		PlanModifiers: []planmodifier.Object{
			objectplanmodifier.UseStateForUnknown(),
		},
		MarkdownDescription: "Optional. Inheritance configuration.",
	},
	"initial_soa_serial": schema.Int64Attribute{
		Optional: true,
		Computed: true,
		Default:  int64default.StaticInt64(1),
		PlanModifiers: []planmodifier.Int64{
			int64planmodifier.RequiresReplaceIfConfigured(),
		},
		MarkdownDescription: "On-create-only. SOA serial is allowed to be set when the authoritative zone is created.",
	},
	"internal_secondaries": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: InternalSecondaryResourceSchemaAttributes,
		},
		Optional: true,
		Validators: []validator.List{
			customvalidator.ListNotEmpty(),
		},
		MarkdownDescription: "Optional. BloxOne DDI hosts acting as internal secondaries. Order is not significant.",
	},
	"notify": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Also notify all external secondary DNS servers if enabled.  Defaults to _false_.",
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
		Computed:            true,
		MarkdownDescription: "The resource identifier.",
	},
	"primary_type": schema.StringAttribute{
		Validators: []validator.String{
			stringvalidator.OneOf("external", "cloud"),
		},
		Required: true,
		PlanModifiers: []planmodifier.String{
			stringplanmodifier.RequiresReplaceIfConfigured(),
		},
		MarkdownDescription: "Primary type for an authoritative zone. Read only after creation. Allowed values:  * _external_: zone data owned by an external nameserver,  * _cloud_: zone data is owned by a BloxOne DDI host.",
	},
	"query_acl": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: ACLItemResourceSchemaAttributes,
		},
		Optional: true,
		Validators: []validator.List{
			customvalidator.ListNotEmpty(),
		},
		MarkdownDescription: "Optional. Clients must match this ACL to make authoritative queries. Also used for recursive queries if that ACL is unset.  Defaults to empty.",
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
	"transfer_acl": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: ACLItemResourceSchemaAttributes,
		},
		Optional: true,
		Validators: []validator.List{
			customvalidator.ListNotEmpty(),
		},
		MarkdownDescription: "Optional. Clients must match this ACL to receive zone transfers.",
	},
	"update_acl": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: ACLItemResourceSchemaAttributes,
		},
		Optional: true,
		Validators: []validator.List{
			customvalidator.ListNotEmpty(),
		},
		MarkdownDescription: "Optional. Specifies which hosts are allowed to submit Dynamic DNS updates for authoritative zones of _primary_type_ _cloud_.  Defaults to empty.",
	},
	"use_forwarders_for_subzones": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(true),
		MarkdownDescription: "Optional. Use default forwarders to resolve queries for subzones.  Defaults to _true_.",
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
func (m *ZoneAuthModel) Expand(ctx context.Context, diags *diag.Diagnostics, isCreate bool) *coremodel.ZoneAuth {
	if m == nil {
		return nil
	}

	obj := &coremodel.ZoneAuth{}

	// Expand NIOS nested attribute (returns nil if not present)
	niosModel := flex.ExpandNestedObject[NIOSZoneAuthModel](ctx, m.NIOS, diags)
	if niosModel != nil {
		obj.NIOS = niosModel.Expand(ctx, diags, isCreate)
	}

	// Expand UDDI nested attribute (returns nil if not present)
	uddiModel := flex.ExpandNestedObject[UDDIZoneAuthModel](ctx, m.UDDI, diags)
	if uddiModel != nil {
		obj.UDDI = uddiModel.Expand(ctx, diags, isCreate)
	}

	return obj
}

// Expand converts the NIOS TF model to the core model.
func (m *NIOSZoneAuthModel) Expand(ctx context.Context, diags *diag.Diagnostics, isCreate bool) *coremodel.NIOSZoneAuthExt {
	ext := &coremodel.NIOSZoneAuthExt{
		AllowActiveDir:                   flex.ExpandFrameworkListNestedBlock(ctx, m.AllowActiveDir, diags, ExpandZoneAuthAllowActiveDir),
		AllowFixedRrsetOrder:             flex.ExpandBoolPointer(m.AllowFixedRrsetOrder),
		AllowGssTsigForUnderscoreZone:    flex.ExpandBoolPointer(m.AllowGssTsigForUnderscoreZone),
		AllowGssTsigZoneUpdates:          flex.ExpandBoolPointer(m.AllowGssTsigZoneUpdates),
		AllowQuery:                       flex.ExpandFrameworkListNestedBlock(ctx, m.AllowQuery, diags, ExpandZoneAuthAllowQuery),
		AllowTransfer:                    flex.ExpandFrameworkListNestedBlock(ctx, m.AllowTransfer, diags, ExpandZoneAuthAllowTransfer),
		AllowUpdate:                      flex.ExpandFrameworkListNestedBlock(ctx, m.AllowUpdate, diags, ExpandZoneAuthAllowUpdate),
		AllowUpdateForwarding:            flex.ExpandBoolPointer(m.AllowUpdateForwarding),
		Comment:                          flex.ExpandStringPointerNullAsEmpty(m.Comment),
		CopyXferToNotify:                 flex.ExpandBoolPointer(m.CopyXferToNotify),
		CreateUnderscoreZones:            flex.ExpandBoolPointer(m.CreateUnderscoreZones),
		DdnsForceCreationTimestampUpdate: flex.ExpandBoolPointer(m.DdnsForceCreationTimestampUpdate),
		DdnsPrincipalGroup:               flex.ExpandStringPointer(m.DdnsPrincipalGroup),
		DdnsPrincipalTracking:            flex.ExpandBoolPointer(m.DdnsPrincipalTracking),
		DdnsRestrictPatterns:             flex.ExpandBoolPointer(m.DdnsRestrictPatterns),
		DdnsRestrictPatternsList:         flex.ExpandFrameworkListString(ctx, m.DdnsRestrictPatternsList, diags),
		DdnsRestrictProtected:            flex.ExpandBoolPointer(m.DdnsRestrictProtected),
		DdnsRestrictSecure:               flex.ExpandBoolPointer(m.DdnsRestrictSecure),
		DdnsRestrictStatic:               flex.ExpandBoolPointer(m.DdnsRestrictStatic),
		Disable:                          flex.ExpandBoolPointer(m.Disable),
		DisableForwarding:                flex.ExpandBoolPointer(m.DisableForwarding),
		DnsIntegrityEnable:               flex.ExpandBoolPointer(m.DnsIntegrityEnable),
		DnsIntegrityFrequency:            flex.ExpandInt64Pointer(m.DnsIntegrityFrequency),
		DnsIntegrityMember:               flex.ExpandStringPointer(m.DnsIntegrityMember),
		DnsIntegrityVerboseLogging:       flex.ExpandBoolPointer(m.DnsIntegrityVerboseLogging),
		DnssecKeyParams:                  ExpandZoneAuthDnssecKeyParams(ctx, m.DnssecKeyParams, diags),
		DnssecKeys:                       flex.ExpandFrameworkListNestedBlock(ctx, m.DnssecKeys, diags, ExpandZoneAuthDnssecKeys),
		EffectiveCheckNamesPolicy:        flex.ExpandStringPointer(m.EffectiveCheckNamesPolicy),
		ExtAttrs:                         flex.ExpandMapStringAny(ctx, m.ExtAttrs, diags),
		ExternalPrimaries:                flex.ExpandFrameworkListNestedBlock(ctx, m.ExternalPrimaries, diags, ExpandZoneAuthExternalPrimaries),
		ExternalSecondaries:              flex.ExpandFrameworkListNestedBlock(ctx, m.ExternalSecondaries, diags, ExpandZoneAuthExternalSecondaries),
		GridPrimary:                      flex.ExpandFrameworkListNestedBlock(ctx, m.GridPrimary, diags, ExpandZoneAuthGridPrimary),
		GridSecondaries:                  flex.ExpandFrameworkListNestedBlock(ctx, m.GridSecondaries, diags, ExpandZoneAuthGridSecondaries),
		LastQueriedAcl:                   flex.ExpandFrameworkListNestedBlock(ctx, m.LastQueriedAcl, diags, ExpandZoneAuthLastQueriedAcl),
		Locked:                           flex.ExpandBoolPointer(m.Locked),
		MemberSoaMnames:                  flex.ExpandFrameworkListNestedBlock(ctx, m.MemberSoaMnames, diags, ExpandZoneAuthMemberSoaMnames),
		MsAdIntegrated:                   flex.ExpandBoolPointer(m.MsAdIntegrated),
		MsAllowTransfer:                  flex.ExpandFrameworkListNestedBlock(ctx, m.MsAllowTransfer, diags, ExpandZoneAuthMsAllowTransfer),
		MsAllowTransferMode:              flex.ExpandStringPointer(m.MsAllowTransferMode),
		MsDcNsRecordCreation:             flex.ExpandFrameworkListNestedBlock(ctx, m.MsDcNsRecordCreation, diags, ExpandZoneAuthMsDcNsRecordCreation),
		MsDdnsMode:                       flex.ExpandStringPointer(m.MsDdnsMode),
		MsPrimaries:                      flex.ExpandFrameworkListNestedBlock(ctx, m.MsPrimaries, diags, ExpandZoneAuthMsPrimaries),
		MsSecondaries:                    flex.ExpandFrameworkListNestedBlock(ctx, m.MsSecondaries, diags, ExpandZoneAuthMsSecondaries),
		MsSyncDisabled:                   flex.ExpandBoolPointer(m.MsSyncDisabled),
		NotifyDelay:                      flex.ExpandInt64Pointer(m.NotifyDelay),
		NsGroup:                          flex.ExpandStringPointer(m.NsGroup),
		Prefix:                           flex.ExpandStringPointer(m.Prefix.StringValue),
		RecordNamePolicy:                 flex.ExpandStringPointer(m.RecordNamePolicy),
		RemoveSubzones:                   flex.ExpandBoolPointer(m.RemoveSubzones),
		RestartIfNeeded:                  flex.ExpandBoolPointer(m.RestartIfNeeded),
		ScavengingSettings:               ExpandZoneAuthScavengingSettings(ctx, m.ScavengingSettings, diags),
		SetSoaSerialNumber:               flex.ExpandBoolPointer(m.SetSoaSerialNumber),
		SoaDefaultTtl:                    flex.ExpandInt64Pointer(m.SoaDefaultTtl),
		SoaEmail:                         flex.ExpandStringPointer(m.SoaEmail),
		SoaExpire:                        flex.ExpandInt64Pointer(m.SoaExpire),
		SoaNegativeTtl:                   flex.ExpandInt64Pointer(m.SoaNegativeTtl),
		SoaRefresh:                       flex.ExpandInt64Pointer(m.SoaRefresh),
		SoaRetry:                         flex.ExpandInt64Pointer(m.SoaRetry),
		SoaSerialNumber:                  flex.ExpandInt64Pointer(m.SoaSerialNumber),
		Srgs:                             flex.ExpandFrameworkListString(ctx, m.Srgs, diags),
		UpdateForwarding:                 flex.ExpandFrameworkListNestedBlock(ctx, m.UpdateForwarding, diags, ExpandZoneAuthUpdateForwarding),
		UseCheckNamesPolicy:              flex.ExpandBoolPointer(m.UseCheckNamesPolicy),
		UseExternalPrimary:               flex.ExpandBoolPointer(m.UseExternalPrimary),
		UseImportFrom:                    flex.ExpandBoolPointer(m.UseImportFrom),
		View:                             flex.ExpandStringPointerNullAsEmpty(m.View),
	}
	if isCreate {
		ext.Fqdn = flex.ExpandStringPointerNullAsEmpty(m.Fqdn)
		ext.ZoneFormat = flex.ExpandStringPointer(m.ZoneFormat)
	}
	return ext
}

// ApplyZoneAuthNIOSUseFlags derives NIOS use flags from the raw config
// value(s) and writes them onto the core model. A flag is true when the user
// set any of its governed value fields in config.
func ApplyZoneAuthNIOSUseFlags(ctx context.Context, config tfsdk.Config, obj *coremodel.ZoneAuth, diags *diag.Diagnostics) {
	if obj == nil || obj.NIOS == nil {
		return
	}
	obj.NIOS.UseAllowActiveDir = flex.DeriveUseFlag(ctx, config, diags, path.Root("nios").AtName("allow_active_dir"))
	obj.NIOS.UseAllowQuery = flex.DeriveUseFlag(ctx, config, diags, path.Root("nios").AtName("allow_query"))
	obj.NIOS.UseAllowTransfer = flex.DeriveUseFlag(ctx, config, diags, path.Root("nios").AtName("allow_transfer"))
	obj.NIOS.UseAllowUpdate = flex.DeriveUseFlag(ctx, config, diags, path.Root("nios").AtName("allow_update"))
	obj.NIOS.UseAllowUpdateForwarding = flex.DeriveUseFlag(ctx, config, diags, path.Root("nios").AtName("allow_update_forwarding"))
	obj.NIOS.UseCopyXferToNotify = flex.DeriveUseFlag(ctx, config, diags, path.Root("nios").AtName("copy_xfer_to_notify"))
	obj.NIOS.UseDdnsForceCreationTimestampUpdate = flex.DeriveUseFlag(ctx, config, diags, path.Root("nios").AtName("ddns_force_creation_timestamp_update"))
	obj.NIOS.UseDdnsPatternsRestriction = flex.DeriveUseFlag(ctx, config, diags, path.Root("nios").AtName("ddns_restrict_patterns_list"), path.Root("nios").AtName("ddns_restrict_patterns"))
	obj.NIOS.UseDdnsPrincipalSecurity = flex.DeriveUseFlag(ctx, config, diags, path.Root("nios").AtName("ddns_restrict_secure"), path.Root("nios").AtName("ddns_principal_tracking"), path.Root("nios").AtName("ddns_principal_group"))
	obj.NIOS.UseDdnsRestrictProtected = flex.DeriveUseFlag(ctx, config, diags, path.Root("nios").AtName("ddns_restrict_protected"))
	obj.NIOS.UseDdnsRestrictStatic = flex.DeriveUseFlag(ctx, config, diags, path.Root("nios").AtName("ddns_restrict_static"))
	obj.NIOS.UseDnssecKeyParams = flex.DeriveUseFlag(ctx, config, diags, path.Root("nios").AtName("dnssec_key_params"))
	obj.NIOS.UseGridZoneTimer = flex.DeriveUseFlag(ctx, config, diags, path.Root("nios").AtName("soa_default_ttl"), path.Root("nios").AtName("soa_expire"), path.Root("nios").AtName("soa_negative_ttl"), path.Root("nios").AtName("soa_refresh"), path.Root("nios").AtName("soa_retry"))
	obj.NIOS.UseNotifyDelay = flex.DeriveUseFlag(ctx, config, diags, path.Root("nios").AtName("notify_delay"))
	obj.NIOS.UseRecordNamePolicy = flex.DeriveUseFlag(ctx, config, diags, path.Root("nios").AtName("record_name_policy"))
	obj.NIOS.UseScavengingSettings = flex.DeriveUseFlag(ctx, config, diags, path.Root("nios").AtName("scavenging_settings"), path.Root("nios").AtName("last_queried_acl"))
	obj.NIOS.UseSoaEmail = flex.DeriveUseFlag(ctx, config, diags, path.Root("nios").AtName("soa_email"))
}

// Expand converts the UDDI TF model to the core model.
func (m *UDDIZoneAuthModel) Expand(ctx context.Context, diags *diag.Diagnostics, isCreate bool) *coremodel.UDDIZoneAuthExt {
	ext := &coremodel.UDDIZoneAuthExt{
		Comment:                  flex.ExpandStringPointer(m.Comment),
		CompartmentId:            flex.ExpandStringPointer(m.CompartmentId),
		Disabled:                 flex.ExpandBoolPointer(m.Disabled),
		ExternalPrimaries:        flex.ExpandFrameworkListNestedBlock(ctx, m.ExternalPrimaries, diags, ExpandExternalPrimary),
		ExternalSecondaries:      flex.ExpandFrameworkListNestedBlock(ctx, m.ExternalSecondaries, diags, ExpandExternalSecondary),
		GssTsigEnabled:           flex.ExpandBoolPointer(m.GssTsigEnabled),
		InheritanceSources:       ExpandAuthZoneInheritance(ctx, m.InheritanceSources, diags),
		InternalSecondaries:      flex.ExpandFrameworkListNestedBlock(ctx, m.InternalSecondaries, diags, ExpandInternalSecondary),
		Notify:                   flex.ExpandBoolPointer(m.Notify),
		Nsgs:                     flex.ExpandFrameworkListString(ctx, m.Nsgs, diags),
		Parent:                   flex.ExpandStringPointer(m.Parent),
		QueryAcl:                 flex.ExpandFrameworkListNestedBlock(ctx, m.QueryAcl, diags, ExpandACLItem),
		Tags:                     flex.ExpandMapStringAny(ctx, m.Tags, diags),
		TransferAcl:              flex.ExpandFrameworkListNestedBlock(ctx, m.TransferAcl, diags, ExpandACLItem),
		UpdateAcl:                flex.ExpandFrameworkListNestedBlock(ctx, m.UpdateAcl, diags, ExpandACLItem),
		UseForwardersForSubzones: flex.ExpandBoolPointer(m.UseForwardersForSubzones),
	}
	if isCreate {
		ext.Fqdn = flex.ExpandStringPointer(m.Fqdn)
		ext.InitialSoaSerial = flex.ExpandInt64Pointer(m.InitialSoaSerial)
		ext.PrimaryType = flex.ExpandStringPointer(m.PrimaryType)
		ext.View = flex.ExpandStringPointer(m.View)
	}
	return ext
}

// Flatten populates the TF model from a core response.
func (m *ZoneAuthModel) Flatten(ctx context.Context, resp *coremodel.ZoneAuth, diags *diag.Diagnostics) {
	if resp == nil {
		return
	}

	m.Id = flex.FlattenStringPointer(resp.Id)

	// Extract existing NIOS model, flatten API response onto it, convert back
	niosModel := flex.ExpandNestedObject[NIOSZoneAuthModel](ctx, m.NIOS, diags)
	if niosModel == nil {
		niosModel = &NIOSZoneAuthModel{}
	}
	plannedNIOS := flex.ExpandNestedObject[NIOSZoneAuthModel](ctx, m.NIOS, diags)
	niosModel.Flatten(ctx, resp.NIOS, diags)
	if resp.NIOS != nil {
		PostFlattenZoneAuthNIOS(ctx, plannedNIOS, niosModel, diags)
		m.NIOS = flex.FlattenNestedObject(ctx, niosModel, NIOSZoneAuthAttrTypes, diags)
	} else {
		m.NIOS = types.ObjectNull(NIOSZoneAuthAttrTypes)
	}

	// Extract existing UDDI model, flatten API response onto it, convert back
	uddiModel := flex.ExpandNestedObject[UDDIZoneAuthModel](ctx, m.UDDI, diags)
	if uddiModel == nil {
		uddiModel = &UDDIZoneAuthModel{}
	}
	uddiModel.Flatten(ctx, resp.UDDI, diags)
	if resp.UDDI != nil {
		m.UDDI = flex.FlattenNestedObject(ctx, uddiModel, UDDIZoneAuthAttrTypes, diags)
	} else {
		m.UDDI = types.ObjectNull(UDDIZoneAuthAttrTypes)
	}
}

// Flatten merges API response onto existing NIOS model.
func (m *NIOSZoneAuthModel) Flatten(ctx context.Context, from *coremodel.NIOSZoneAuthExt, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	planExtAttrs := m.ExtAttrs
	if planExtAttrs.IsUnknown() {
		planExtAttrs = types.MapNull(types.StringType)
	}
	m.AllowActiveDir = flex.FlattenFrameworkListNestedBlock(ctx, from.AllowActiveDir, ZoneAuthAllowActiveDirAttrTypes, diags, FlattenZoneAuthAllowActiveDir)
	m.AllowFixedRrsetOrder = flex.FlattenBoolPointer(from.AllowFixedRrsetOrder)
	m.AllowGssTsigForUnderscoreZone = flex.FlattenBoolPointer(from.AllowGssTsigForUnderscoreZone)
	m.AllowGssTsigZoneUpdates = flex.FlattenBoolPointer(from.AllowGssTsigZoneUpdates)
	m.AllowQuery = flex.FlattenFrameworkListNestedBlock(ctx, from.AllowQuery, ZoneAuthAllowQueryAttrTypes, diags, FlattenZoneAuthAllowQuery)
	m.AllowTransfer = flex.FlattenFrameworkListNestedBlock(ctx, from.AllowTransfer, ZoneAuthAllowTransferAttrTypes, diags, FlattenZoneAuthAllowTransfer)
	m.AllowUpdate = flex.FlattenFrameworkListNestedBlock(ctx, from.AllowUpdate, ZoneAuthAllowUpdateAttrTypes, diags, FlattenZoneAuthAllowUpdate)
	m.AllowUpdateForwarding = flex.FlattenBoolPointer(from.AllowUpdateForwarding)
	m.Comment = flex.FlattenStringPointerEmptyAsNull(from.Comment)
	m.CopyXferToNotify = flex.FlattenBoolPointer(from.CopyXferToNotify)
	m.CreateUnderscoreZones = flex.FlattenBoolPointer(from.CreateUnderscoreZones)
	m.DdnsForceCreationTimestampUpdate = flex.FlattenBoolPointer(from.DdnsForceCreationTimestampUpdate)
	m.DdnsPrincipalGroup = flex.FlattenStringPointerEmptyAsNull(from.DdnsPrincipalGroup)
	m.DdnsPrincipalTracking = flex.FlattenBoolPointer(from.DdnsPrincipalTracking)
	m.DdnsRestrictPatterns = flex.FlattenBoolPointer(from.DdnsRestrictPatterns)
	m.DdnsRestrictPatternsList = flex.FlattenFrameworkUnorderedListString(ctx, from.DdnsRestrictPatternsList, diags)
	m.DdnsRestrictProtected = flex.FlattenBoolPointer(from.DdnsRestrictProtected)
	m.DdnsRestrictSecure = flex.FlattenBoolPointer(from.DdnsRestrictSecure)
	m.DdnsRestrictStatic = flex.FlattenBoolPointer(from.DdnsRestrictStatic)
	m.Disable = flex.FlattenBoolPointer(from.Disable)
	m.DisableForwarding = flex.FlattenBoolPointer(from.DisableForwarding)
	m.DisplayDomain = flex.FlattenStringPointerEmptyAsNull(from.DisplayDomain)
	m.DnsIntegrityEnable = flex.FlattenBoolPointer(from.DnsIntegrityEnable)
	m.DnsIntegrityFrequency = flex.FlattenInt64Pointer(from.DnsIntegrityFrequency)
	m.DnsIntegrityMember = flex.FlattenStringPointerEmptyAsNull(from.DnsIntegrityMember)
	m.DnsIntegrityVerboseLogging = flex.FlattenBoolPointer(from.DnsIntegrityVerboseLogging)
	m.DnssecKeyParams = FlattenZoneAuthDnssecKeyParams(ctx, from.DnssecKeyParams, diags)
	m.DnssecKeys = flex.FlattenFrameworkListNestedBlock(ctx, from.DnssecKeys, ZoneAuthDnssecKeysAttrTypes, diags, FlattenZoneAuthDnssecKeys)
	m.EffectiveCheckNamesPolicy = flex.FlattenStringPointerEmptyAsNull(from.EffectiveCheckNamesPolicy)
	m.ExtAttrs, m.ExtAttrsAll = flex.FlattenEAs(planExtAttrs, from.ExtAttrs)
	m.ExternalPrimaries = flex.FlattenFrameworkListNestedBlock(ctx, from.ExternalPrimaries, ZoneAuthExternalPrimariesAttrTypes, diags, FlattenZoneAuthExternalPrimaries)
	m.ExternalSecondaries = flex.FlattenFrameworkListNestedBlock(ctx, from.ExternalSecondaries, ZoneAuthExternalSecondariesAttrTypes, diags, FlattenZoneAuthExternalSecondaries)
	m.Fqdn = flex.FlattenStringPointerEmptyAsNull(from.Fqdn)
	m.GridPrimary = flex.FlattenFrameworkListNestedBlock(ctx, from.GridPrimary, ZoneAuthGridPrimaryAttrTypes, diags, FlattenZoneAuthGridPrimary)
	m.GridSecondaries = flex.FlattenFrameworkListNestedBlock(ctx, from.GridSecondaries, ZoneAuthGridSecondariesAttrTypes, diags, FlattenZoneAuthGridSecondaries)
	m.LastQueriedAcl = flex.FlattenFrameworkListNestedBlock(ctx, from.LastQueriedAcl, ZoneAuthLastQueriedAclAttrTypes, diags, FlattenZoneAuthLastQueriedAcl)
	m.Locked = flex.FlattenBoolPointer(from.Locked)
	m.MemberSoaMnames = flex.FlattenFrameworkListNestedBlock(ctx, from.MemberSoaMnames, ZoneAuthMemberSoaMnamesAttrTypes, diags, FlattenZoneAuthMemberSoaMnames)
	m.MsAdIntegrated = flex.FlattenBoolPointer(from.MsAdIntegrated)
	m.MsAllowTransfer = flex.FlattenFrameworkListNestedBlock(ctx, from.MsAllowTransfer, ZoneAuthMsAllowTransferAttrTypes, diags, FlattenZoneAuthMsAllowTransfer)
	m.MsAllowTransferMode = flex.FlattenStringPointerEmptyAsNull(from.MsAllowTransferMode)
	m.MsDcNsRecordCreation = flex.FlattenFrameworkListNestedBlock(ctx, from.MsDcNsRecordCreation, ZoneAuthMsDcNsRecordCreationAttrTypes, diags, FlattenZoneAuthMsDcNsRecordCreation)
	m.MsDdnsMode = flex.FlattenStringPointerEmptyAsNull(from.MsDdnsMode)
	m.MsPrimaries = flex.FlattenFrameworkListNestedBlock(ctx, from.MsPrimaries, ZoneAuthMsPrimariesAttrTypes, diags, FlattenZoneAuthMsPrimaries)
	m.MsSecondaries = flex.FlattenFrameworkListNestedBlock(ctx, from.MsSecondaries, ZoneAuthMsSecondariesAttrTypes, diags, FlattenZoneAuthMsSecondaries)
	m.MsSyncDisabled = flex.FlattenBoolPointer(from.MsSyncDisabled)
	m.NotifyDelay = flex.FlattenInt64Pointer(from.NotifyDelay)
	m.NsGroup = flex.FlattenStringPointerEmptyAsNull(from.NsGroup)
	m.Prefix.StringValue = flex.FlattenStringPointer(from.Prefix)
	m.RecordNamePolicy = flex.FlattenStringPointerEmptyAsNull(from.RecordNamePolicy)
	m.RemoveSubzones = flex.FlattenBoolPointer(from.RemoveSubzones)
	m.ScavengingSettings = FlattenZoneAuthScavengingSettings(ctx, from.ScavengingSettings, diags)
	m.SoaDefaultTtl = flex.FlattenInt64Pointer(from.SoaDefaultTtl)
	m.SoaEmail = flex.FlattenStringPointerEmptyAsNull(from.SoaEmail)
	m.SoaExpire = flex.FlattenInt64Pointer(from.SoaExpire)
	m.SoaNegativeTtl = flex.FlattenInt64Pointer(from.SoaNegativeTtl)
	m.SoaRefresh = flex.FlattenInt64Pointer(from.SoaRefresh)
	m.SoaRetry = flex.FlattenInt64Pointer(from.SoaRetry)
	m.SoaSerialNumber = flex.FlattenInt64Pointer(from.SoaSerialNumber)
	m.Srgs = flex.FlattenFrameworkListString(ctx, from.Srgs, diags)
	m.UpdateForwarding = flex.FlattenFrameworkListNestedBlock(ctx, from.UpdateForwarding, ZoneAuthUpdateForwardingAttrTypes, diags, FlattenZoneAuthUpdateForwarding)
	m.UseCheckNamesPolicy = flex.FlattenBoolPointer(from.UseCheckNamesPolicy)
	m.UseExternalPrimary = flex.FlattenBoolPointer(from.UseExternalPrimary)
	m.UseImportFrom = flex.FlattenBoolPointer(from.UseImportFrom)
	m.View = flex.FlattenStringPointerEmptyAsNull(from.View)
	m.ZoneFormat = flex.FlattenStringPointerEmptyAsNull(from.ZoneFormat)
}

// Flatten merges API response onto existing UDDI model.
func (m *UDDIZoneAuthModel) Flatten(ctx context.Context, from *coremodel.UDDIZoneAuthExt, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	m.Comment = flex.FlattenStringPointer(from.Comment)
	m.CompartmentId = flex.FlattenStringPointer(from.CompartmentId)
	m.Disabled = flex.FlattenBoolPointer(from.Disabled)
	m.ExternalPrimaries = flex.FlattenFrameworkListNestedBlock(ctx, from.ExternalPrimaries, ExternalPrimaryAttrTypes, diags, FlattenExternalPrimary)
	m.ExternalSecondaries = flex.FlattenFrameworkListNestedBlock(ctx, from.ExternalSecondaries, ExternalSecondaryAttrTypes, diags, FlattenExternalSecondary)
	m.Fqdn = flex.FlattenStringPointer(from.Fqdn)
	m.GssTsigEnabled = flex.FlattenBoolPointer(from.GssTsigEnabled)
	m.InheritanceSources = FlattenAuthZoneInheritance(ctx, from.InheritanceSources, diags)
	m.InitialSoaSerial = flex.FlattenInt64Pointer(from.InitialSoaSerial)
	m.InternalSecondaries = flex.FlattenFrameworkListNestedBlock(ctx, from.InternalSecondaries, InternalSecondaryAttrTypes, diags, FlattenInternalSecondary)
	m.Notify = flex.FlattenBoolPointer(from.Notify)
	m.Nsgs = flex.FlattenFrameworkListString(ctx, from.Nsgs, diags)
	m.Parent = flex.FlattenStringPointer(from.Parent)
	m.PrimaryType = flex.FlattenStringPointer(from.PrimaryType)
	m.QueryAcl = flex.FlattenFrameworkListNestedBlock(ctx, from.QueryAcl, ACLItemAttrTypes, diags, FlattenACLItem)
	tagsAll := flex.FlattenMapStringAny(ctx, from.Tags, diags)
	if m.Tags.IsNull() || m.Tags.IsUnknown() {
		m.Tags = tagsAll
	}
	m.TagsAll = tagsAll
	m.TransferAcl = flex.FlattenFrameworkListNestedBlock(ctx, from.TransferAcl, ACLItemAttrTypes, diags, FlattenACLItem)
	m.UpdateAcl = flex.FlattenFrameworkListNestedBlock(ctx, from.UpdateAcl, ACLItemAttrTypes, diags, FlattenACLItem)
	m.UseForwardersForSubzones = flex.FlattenBoolPointer(from.UseForwardersForSubzones)
	m.View = flex.FlattenStringPointer(from.View)
}
