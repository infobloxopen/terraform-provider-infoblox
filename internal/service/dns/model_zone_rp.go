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

type ZoneRpModel struct {
	Id   types.String `tfsdk:"id"`
	NIOS types.Object `tfsdk:"nios"`
}

var ZoneRpAttrTypes = map[string]attr.Type{
	"id":   types.StringType,
	"nios": types.ObjectType{AttrTypes: NIOSZoneRpAttrTypes},
}

type NIOSZoneRpModel struct {
	Comment                          types.String                        `tfsdk:"comment"`
	Disable                          types.Bool                          `tfsdk:"disable"`
	ExtAttrs                         types.Map                           `tfsdk:"ext_attrs"`
	ExtAttrsAll                      types.Map                           `tfsdk:"ext_attrs_all"`
	ExternalPrimaries                types.List                          `tfsdk:"external_primaries"`
	ExternalSecondaries              types.List                          `tfsdk:"external_secondaries"`
	FireeyeRuleMapping               types.Object                        `tfsdk:"fireeye_rule_mapping"`
	Fqdn                             types.String                        `tfsdk:"fqdn"`
	GridPrimary                      types.List                          `tfsdk:"grid_primary"`
	GridSecondaries                  types.List                          `tfsdk:"grid_secondaries"`
	Locked                           types.Bool                          `tfsdk:"locked"`
	LogRpz                           types.Bool                          `tfsdk:"log_rpz"`
	MemberSoaMnames                  types.List                          `tfsdk:"member_soa_mnames"`
	NsGroup                          types.String                        `tfsdk:"ns_group"`
	Prefix                           internaltypes.CaseInsensitiveString `tfsdk:"prefix"`
	RecordNamePolicy                 types.String                        `tfsdk:"record_name_policy"`
	RpzDropIpRuleEnabled             types.Bool                          `tfsdk:"rpz_drop_ip_rule_enabled"`
	RpzDropIpRuleMinPrefixLengthIpv4 types.Int64                         `tfsdk:"rpz_drop_ip_rule_min_prefix_length_ipv4"`
	RpzDropIpRuleMinPrefixLengthIpv6 types.Int64                         `tfsdk:"rpz_drop_ip_rule_min_prefix_length_ipv6"`
	RpzPolicy                        types.String                        `tfsdk:"rpz_policy"`
	RpzSeverity                      types.String                        `tfsdk:"rpz_severity"`
	RpzType                          types.String                        `tfsdk:"rpz_type"`
	SetSoaSerialNumber               types.Bool                          `tfsdk:"set_soa_serial_number"`
	SoaDefaultTtl                    types.Int64                         `tfsdk:"soa_default_ttl"`
	SoaEmail                         types.String                        `tfsdk:"soa_email"`
	SoaExpire                        types.Int64                         `tfsdk:"soa_expire"`
	SoaNegativeTtl                   types.Int64                         `tfsdk:"soa_negative_ttl"`
	SoaRefresh                       types.Int64                         `tfsdk:"soa_refresh"`
	SoaRetry                         types.Int64                         `tfsdk:"soa_retry"`
	SoaSerialNumber                  types.Int64                         `tfsdk:"soa_serial_number"`
	SubstituteName                   types.String                        `tfsdk:"substitute_name"`
	UseExternalPrimary               types.Bool                          `tfsdk:"use_external_primary"`
	View                             types.String                        `tfsdk:"view"`
}

var NIOSZoneRpAttrTypes = map[string]attr.Type{
	"comment":                  types.StringType,
	"disable":                  types.BoolType,
	"ext_attrs":                types.MapType{ElemType: types.StringType},
	"ext_attrs_all":            types.MapType{ElemType: types.StringType},
	"external_primaries":       types.ListType{ElemType: types.ObjectType{AttrTypes: ZoneRpExternalPrimariesAttrTypes}},
	"external_secondaries":     types.ListType{ElemType: types.ObjectType{AttrTypes: ZoneRpExternalSecondariesAttrTypes}},
	"fireeye_rule_mapping":     types.ObjectType{AttrTypes: ZoneRpFireeyeRuleMappingAttrTypes},
	"fqdn":                     types.StringType,
	"grid_primary":             types.ListType{ElemType: types.ObjectType{AttrTypes: ZoneRpGridPrimaryAttrTypes}},
	"grid_secondaries":         types.ListType{ElemType: types.ObjectType{AttrTypes: ZoneRpGridSecondariesAttrTypes}},
	"locked":                   types.BoolType,
	"log_rpz":                  types.BoolType,
	"member_soa_mnames":        types.ListType{ElemType: types.ObjectType{AttrTypes: ZoneRpMemberSoaMnamesAttrTypes}},
	"ns_group":                 types.StringType,
	"prefix":                   internaltypes.CaseInsensitiveStringType{},
	"record_name_policy":       types.StringType,
	"rpz_drop_ip_rule_enabled": types.BoolType,
	"rpz_drop_ip_rule_min_prefix_length_ipv4": types.Int64Type,
	"rpz_drop_ip_rule_min_prefix_length_ipv6": types.Int64Type,
	"rpz_policy":            types.StringType,
	"rpz_severity":          types.StringType,
	"rpz_type":              types.StringType,
	"set_soa_serial_number": types.BoolType,
	"soa_default_ttl":       types.Int64Type,
	"soa_email":             types.StringType,
	"soa_expire":            types.Int64Type,
	"soa_negative_ttl":      types.Int64Type,
	"soa_refresh":           types.Int64Type,
	"soa_retry":             types.Int64Type,
	"soa_serial_number":     types.Int64Type,
	"substitute_name":       types.StringType,
	"use_external_primary":  types.BoolType,
	"view":                  types.StringType,
}

const (
	ZoneRpReturnFields = "address,comment,disable,display_domain,dns_soa_email,extattrs,external_primaries,external_secondaries,fireeye_rule_mapping,fqdn,grid_primary,grid_secondaries,locked,locked_by,log_rpz,mask_prefix,member_soa_mnames,member_soa_serials,network_view,ns_group,parent,prefix,primary_type,record_name_policy,rpz_drop_ip_rule_enabled,rpz_drop_ip_rule_min_prefix_length_ipv4,rpz_drop_ip_rule_min_prefix_length_ipv6,rpz_last_updated_time,rpz_policy,rpz_priority,rpz_priority_end,rpz_severity,rpz_type,soa_default_ttl,soa_email,soa_expire,soa_negative_ttl,soa_refresh,soa_retry,soa_serial_number,substitute_name,use_external_primary,use_grid_zone_timer,use_log_rpz,use_record_name_policy,use_rpz_drop_ip_rule,use_soa_email,view"
)

var ZoneRpResourceSchemaAttributes = map[string]schema.Attribute{
	"id": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The reference to the object.",
	},
	"nios": schema.SingleNestedAttribute{
		Optional:            true,
		MarkdownDescription: "NIOS backend-specific fields.",
		Attributes:          ZoneRpResourceNiosSchemaAttributes,
	},
}

var ZoneRpResourceNiosSchemaAttributes = map[string]schema.Attribute{
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
			Attributes: ZoneRpExternalPrimariesResourceSchemaAttributes,
		},
		Optional: true,
		Validators: []validator.List{
			customvalidator.ListNotEmpty(),
			listvalidator.ConflictsWith(path.MatchRelative().AtParent().AtName("ns_group")),
		},
		MarkdownDescription: "The list of external primary servers.",
	},
	"external_secondaries": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: ZoneRpExternalSecondariesResourceSchemaAttributes,
		},
		Optional: true,
		Validators: []validator.List{
			customvalidator.ListNotEmpty(),
			listvalidator.ConflictsWith(path.MatchRelative().AtParent().AtName("ns_group")),
		},
		MarkdownDescription: "The list of external secondary servers.",
	},
	"fireeye_rule_mapping": schema.SingleNestedAttribute{
		Attributes:          ZoneRpFireeyeRuleMappingResourceSchemaAttributes,
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "",
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
		MarkdownDescription: "The name of this DNS zone in FQDN format.",
	},
	"grid_primary": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: ZoneRpGridPrimaryResourceSchemaAttributes,
		},
		Optional: true,
		Validators: []validator.List{
			customvalidator.ListNotEmpty(),
			listvalidator.ConflictsWith(
				path.MatchRelative().AtParent().AtName("ns_group"),
				path.MatchRelative().AtParent().AtName("external_primaries"),
			),
		},
		MarkdownDescription: "The grid primary servers for this zone.",
	},
	"grid_secondaries": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: ZoneRpGridSecondariesResourceSchemaAttributes,
		},
		Optional: true,
		Validators: []validator.List{
			customvalidator.ListNotEmpty(),
			listvalidator.ConflictsWith(path.MatchRelative().AtParent().AtName("ns_group")),
		},
		MarkdownDescription: "The list with Grid members that are secondary servers for this zone.",
	},
	"locked": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "If you enable this flag, other administrators cannot make conflicting changes. This is for administration purposes only. The zone will continue to serve DNS data even when it is locked.",
	},
	"log_rpz": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(true),
		MarkdownDescription: "Determines whether RPZ logging enabled or not at zone level. When this is set to False, the logging is disabled.",
	},
	"member_soa_mnames": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: ZoneRpMemberSoaMnamesResourceSchemaAttributes,
		},
		Optional: true,
		Validators: []validator.List{
			customvalidator.ListNotEmpty(),
		},
		MarkdownDescription: "The list of per-member SOA MNAME information.",
	},
	"ns_group": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
			stringvalidator.ConflictsWith(path.MatchRelative().AtParent().AtName("grid_primary")),
		},
		MarkdownDescription: "The name server group that serves DNS for this zone.",
	},
	"prefix": schema.StringAttribute{
		Optional:   true,
		Computed:   true,
		CustomType: internaltypes.CaseInsensitiveStringType{},
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
			customvalidator.ValidateTrimmedString(),
		},
		MarkdownDescription: "The RFC2317 prefix value of this DNS zone. Use this field only when the netmask is greater than 24 bits; that is, for a mask between 25 and 31 bits. Enter a prefix, such as the name of the allocated address block. The prefix can be alphanumeric characters, such as 128/26 , 128-189 , or sub-B.",
	},
	"record_name_policy": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
		},
		MarkdownDescription: "The hostname policy for records under this zone.",
	},
	"rpz_drop_ip_rule_enabled": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Enables the appliance to ignore RPZ-IP triggers with prefix lengths less than the specified minimum prefix length.",
	},
	"rpz_drop_ip_rule_min_prefix_length_ipv4": schema.Int64Attribute{
		Optional: true,
		Computed: true,
		Default:  int64default.StaticInt64(29),
		Validators: []validator.Int64{
			int64validator.Between(0, 4294967295),
		},
		MarkdownDescription: "The minimum prefix length for IPv4 RPZ-IP triggers. The appliance ignores RPZ-IP triggers with prefix lengths less than the specified minimum IPv4 prefix length.",
	},
	"rpz_drop_ip_rule_min_prefix_length_ipv6": schema.Int64Attribute{
		Optional:            true,
		Computed:            true,
		Default:             int64default.StaticInt64(112),
		MarkdownDescription: "The minimum prefix length for IPv6 RPZ-IP triggers. The appliance ignores RPZ-IP triggers with prefix lengths less than the specified minimum IPv6 prefix length.",
	},
	"rpz_policy": schema.StringAttribute{
		Default: stringdefault.StaticString("GIVEN"),
		Validators: []validator.String{
			stringvalidator.OneOf("DISABLED", "GIVEN", "NODATA", "NXDOMAIN", "PASSTHRU", "SUBSTITUTE"),
		},
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The response policy zone override policy.",
	},
	"rpz_severity": schema.StringAttribute{
		Default: stringdefault.StaticString("MAJOR"),
		Validators: []validator.String{
			stringvalidator.OneOf("CRITICAL", "MAJOR", "WARNING", "INFORMATIONAL"),
		},
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The severity of this response policy zone.",
	},
	"rpz_type": schema.StringAttribute{
		Validators: []validator.String{
			stringvalidator.OneOf("LOCAL", "FIREEYE", "FEED"),
		},
		Optional: true,
		Computed: true,
		PlanModifiers: []planmodifier.String{
			immutable.ImmutableString(),
		},
		MarkdownDescription: "The type of rpz zone.",
	},
	"set_soa_serial_number": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "The serial number in the SOA record incrementally changes every time the record is modified. The Infoblox appliance allows you to change the serial number (in the SOA record) for the primary server so it is higher than the secondary server, thereby ensuring zone transfers come from the primary server (as they should). To change the serial number you need to set a new value at \"soa_serial_number\" and pass \"set_soa_serial_number\" as True.",
	},
	"soa_default_ttl": schema.Int64Attribute{
		Optional: true,
		Computed: true,
		Validators: []validator.Int64{
			int64validator.AlsoRequires(
				path.MatchRelative().AtParent().AtName("soa_expire"),
				path.MatchRelative().AtParent().AtName("soa_negative_ttl"),
				path.MatchRelative().AtParent().AtName("soa_refresh"),
				path.MatchRelative().AtParent().AtName("soa_retry"),
				path.MatchRelative().AtParent().AtName("grid_primary"),
			),
		},
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
		Optional: true,
		Computed: true,
		Validators: []validator.Int64{
			int64validator.AlsoRequires(
				path.MatchRelative().AtParent().AtName("soa_default_ttl"),
				path.MatchRelative().AtParent().AtName("soa_negative_ttl"),
				path.MatchRelative().AtParent().AtName("soa_refresh"),
				path.MatchRelative().AtParent().AtName("soa_retry"),
				path.MatchRelative().AtParent().AtName("grid_primary"),
			),
		},
		MarkdownDescription: "This setting defines the amount of time, in seconds, after which the secondary server stops giving out answers about the zone because the zone data is too old to be useful. The default is one week.",
	},
	"soa_negative_ttl": schema.Int64Attribute{
		Optional: true,
		Computed: true,
		Validators: []validator.Int64{
			int64validator.AlsoRequires(
				path.MatchRelative().AtParent().AtName("soa_default_ttl"),
				path.MatchRelative().AtParent().AtName("soa_expire"),
				path.MatchRelative().AtParent().AtName("soa_refresh"),
				path.MatchRelative().AtParent().AtName("soa_retry"),
				path.MatchRelative().AtParent().AtName("grid_primary"),
			),
		},
		MarkdownDescription: "The negative Time to Live (TTL) value of the SOA of the zone indicates how long a secondary server can cache data for \"Does Not Respond\" responses.",
	},
	"soa_refresh": schema.Int64Attribute{
		Optional: true,
		Computed: true,
		Validators: []validator.Int64{
			int64validator.AlsoRequires(
				path.MatchRelative().AtParent().AtName("soa_default_ttl"),
				path.MatchRelative().AtParent().AtName("soa_expire"),
				path.MatchRelative().AtParent().AtName("soa_negative_ttl"),
				path.MatchRelative().AtParent().AtName("soa_retry"),
				path.MatchRelative().AtParent().AtName("grid_primary"),
			),
		},
		MarkdownDescription: "This indicates the interval at which a secondary server sends a message to the primary server for a zone to check that its data is current, and retrieve fresh data if it is not.",
	},
	"soa_retry": schema.Int64Attribute{
		Optional: true,
		Computed: true,
		Validators: []validator.Int64{
			int64validator.AlsoRequires(
				path.MatchRelative().AtParent().AtName("soa_default_ttl"),
				path.MatchRelative().AtParent().AtName("soa_expire"),
				path.MatchRelative().AtParent().AtName("soa_negative_ttl"),
				path.MatchRelative().AtParent().AtName("soa_refresh"),
				path.MatchRelative().AtParent().AtName("grid_primary"),
			),
		},
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
	"substitute_name": schema.StringAttribute{
		Optional: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
			customvalidator.ValidateTrimmedString(),
		},
		MarkdownDescription: "The canonical name of redirect target in substitute policy of response policy zone.",
	},
	"use_external_primary": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "This flag controls whether the zone is using an external primary.",
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
}

// Expand converts the TF model to the infoblox core model
func (m *ZoneRpModel) Expand(ctx context.Context, diags *diag.Diagnostics, isCreate bool) *coremodel.ZoneRp {
	if m == nil {
		return nil
	}

	obj := &coremodel.ZoneRp{}

	// Expand NIOS nested attribute (returns nil if not present)
	niosModel := flex.ExpandNestedObject[NIOSZoneRpModel](ctx, m.NIOS, diags)
	if niosModel != nil {
		obj.NIOS = niosModel.Expand(ctx, diags, isCreate)
	}

	return obj
}

// Expand converts the NIOS TF model to the core model.
func (m *NIOSZoneRpModel) Expand(ctx context.Context, diags *diag.Diagnostics, isCreate bool) *coremodel.NIOSZoneRpExt {
	ext := &coremodel.NIOSZoneRpExt{
		Comment:                          flex.ExpandStringPointerNullAsEmpty(m.Comment),
		Disable:                          flex.ExpandBoolPointer(m.Disable),
		ExtAttrs:                         flex.ExpandMapStringAny(ctx, m.ExtAttrs, diags),
		ExternalPrimaries:                flex.ExpandFrameworkListNestedBlock(ctx, m.ExternalPrimaries, diags, ExpandZoneRpExternalPrimaries),
		ExternalSecondaries:              flex.ExpandFrameworkListNestedBlock(ctx, m.ExternalSecondaries, diags, ExpandZoneRpExternalSecondaries),
		FireeyeRuleMapping:               ExpandZoneRpFireeyeRuleMapping(ctx, m.FireeyeRuleMapping, diags),
		GridPrimary:                      flex.ExpandFrameworkListNestedBlock(ctx, m.GridPrimary, diags, ExpandZoneRpGridPrimary),
		GridSecondaries:                  flex.ExpandFrameworkListNestedBlock(ctx, m.GridSecondaries, diags, ExpandZoneRpGridSecondaries),
		Locked:                           flex.ExpandBoolPointer(m.Locked),
		LogRpz:                           flex.ExpandBoolPointer(m.LogRpz),
		MemberSoaMnames:                  flex.ExpandFrameworkListNestedBlock(ctx, m.MemberSoaMnames, diags, ExpandZoneRpMemberSoaMnames),
		NsGroup:                          flex.ExpandStringPointer(m.NsGroup),
		Prefix:                           flex.ExpandStringPointer(m.Prefix.StringValue),
		RecordNamePolicy:                 flex.ExpandStringPointer(m.RecordNamePolicy),
		RpzDropIpRuleEnabled:             flex.ExpandBoolPointer(m.RpzDropIpRuleEnabled),
		RpzDropIpRuleMinPrefixLengthIpv4: flex.ExpandInt64Pointer(m.RpzDropIpRuleMinPrefixLengthIpv4),
		RpzDropIpRuleMinPrefixLengthIpv6: flex.ExpandInt64Pointer(m.RpzDropIpRuleMinPrefixLengthIpv6),
		RpzPolicy:                        flex.ExpandStringPointerNullAsEmpty(m.RpzPolicy),
		RpzSeverity:                      flex.ExpandStringPointerNullAsEmpty(m.RpzSeverity),
		SetSoaSerialNumber:               flex.ExpandBoolPointer(m.SetSoaSerialNumber),
		SoaDefaultTtl:                    flex.ExpandInt64Pointer(m.SoaDefaultTtl),
		SoaEmail:                         flex.ExpandStringPointer(m.SoaEmail),
		SoaExpire:                        flex.ExpandInt64Pointer(m.SoaExpire),
		SoaNegativeTtl:                   flex.ExpandInt64Pointer(m.SoaNegativeTtl),
		SoaRefresh:                       flex.ExpandInt64Pointer(m.SoaRefresh),
		SoaRetry:                         flex.ExpandInt64Pointer(m.SoaRetry),
		SoaSerialNumber:                  flex.ExpandInt64Pointer(m.SoaSerialNumber),
		SubstituteName:                   flex.ExpandStringPointerNullAsEmpty(m.SubstituteName),
		UseExternalPrimary:               flex.ExpandBoolPointer(m.UseExternalPrimary),
	}
	if isCreate {
		ext.Fqdn = flex.ExpandStringPointerNullAsEmpty(m.Fqdn)
		ext.RpzType = flex.ExpandStringPointer(m.RpzType)
		ext.View = flex.ExpandStringPointerNullAsEmpty(m.View)
	}
	return ext
}

// ApplyZoneRpNIOSUseFlags derives NIOS use flags from the raw config
// value(s) and writes them onto the core model. A flag is true when the user
// set any of its governed value fields in config.
func ApplyZoneRpNIOSUseFlags(ctx context.Context, config tfsdk.Config, obj *coremodel.ZoneRp, diags *diag.Diagnostics) {
	if obj == nil || obj.NIOS == nil {
		return
	}
	obj.NIOS.UseGridZoneTimer = flex.DeriveUseFlag(ctx, config, diags, path.Root("nios").AtName("soa_default_ttl"), path.Root("nios").AtName("soa_expire"), path.Root("nios").AtName("soa_negative_ttl"), path.Root("nios").AtName("soa_refresh"), path.Root("nios").AtName("soa_retry"))
	obj.NIOS.UseLogRpz = flex.DeriveUseFlag(ctx, config, diags, path.Root("nios").AtName("log_rpz"))
	obj.NIOS.UseRecordNamePolicy = flex.DeriveUseFlag(ctx, config, diags, path.Root("nios").AtName("record_name_policy"))
	obj.NIOS.UseRpzDropIpRule = flex.DeriveUseFlag(ctx, config, diags, path.Root("nios").AtName("rpz_drop_ip_rule_enabled"), path.Root("nios").AtName("rpz_drop_ip_rule_min_prefix_length_ipv4"), path.Root("nios").AtName("rpz_drop_ip_rule_min_prefix_length_ipv6"))
	obj.NIOS.UseSoaEmail = flex.DeriveUseFlag(ctx, config, diags, path.Root("nios").AtName("soa_email"))
}

// Flatten populates the TF model from a core response.
func (m *ZoneRpModel) Flatten(ctx context.Context, resp *coremodel.ZoneRp, diags *diag.Diagnostics) {
	if resp == nil {
		return
	}

	m.Id = flex.FlattenStringPointer(resp.Id)

	// Extract existing NIOS model, flatten API response onto it, convert back
	niosModel := flex.ExpandNestedObject[NIOSZoneRpModel](ctx, m.NIOS, diags)
	if niosModel == nil {
		niosModel = &NIOSZoneRpModel{}
	}
	plannedNIOS := flex.ExpandNestedObject[NIOSZoneRpModel](ctx, m.NIOS, diags)
	niosModel.Flatten(ctx, resp.NIOS, diags)
	if resp.NIOS != nil {
		PostFlattenZoneRpNIOS(ctx, plannedNIOS, niosModel, diags)
		m.NIOS = flex.FlattenNestedObject(ctx, niosModel, NIOSZoneRpAttrTypes, diags)
	} else {
		m.NIOS = types.ObjectNull(NIOSZoneRpAttrTypes)
	}

}

// Flatten merges API response onto existing NIOS model.
func (m *NIOSZoneRpModel) Flatten(ctx context.Context, from *coremodel.NIOSZoneRpExt, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	planExtAttrs := m.ExtAttrs
	if planExtAttrs.IsUnknown() {
		planExtAttrs = types.MapNull(types.StringType)
	}
	m.Comment = flex.FlattenStringPointerEmptyAsNull(from.Comment)
	m.Disable = flex.FlattenBoolPointer(from.Disable)
	m.ExtAttrs, m.ExtAttrsAll = flex.FlattenEAs(planExtAttrs, from.ExtAttrs)
	m.ExternalPrimaries = flex.FlattenFrameworkListNestedBlock(ctx, from.ExternalPrimaries, ZoneRpExternalPrimariesAttrTypes, diags, FlattenZoneRpExternalPrimaries)
	m.ExternalSecondaries = flex.FlattenFrameworkListNestedBlock(ctx, from.ExternalSecondaries, ZoneRpExternalSecondariesAttrTypes, diags, FlattenZoneRpExternalSecondaries)
	m.FireeyeRuleMapping = FlattenZoneRpFireeyeRuleMapping(ctx, from.FireeyeRuleMapping, diags)
	m.Fqdn = flex.FlattenStringPointerEmptyAsNull(from.Fqdn)
	m.GridPrimary = flex.FlattenFrameworkListNestedBlock(ctx, from.GridPrimary, ZoneRpGridPrimaryAttrTypes, diags, FlattenZoneRpGridPrimary)
	m.GridSecondaries = flex.FlattenFrameworkListNestedBlock(ctx, from.GridSecondaries, ZoneRpGridSecondariesAttrTypes, diags, FlattenZoneRpGridSecondaries)
	m.Locked = flex.FlattenBoolPointer(from.Locked)
	m.LogRpz = flex.FlattenBoolPointer(from.LogRpz)
	m.MemberSoaMnames = flex.FlattenFrameworkListNestedBlock(ctx, from.MemberSoaMnames, ZoneRpMemberSoaMnamesAttrTypes, diags, FlattenZoneRpMemberSoaMnames)
	m.NsGroup = flex.FlattenStringPointerEmptyAsNull(from.NsGroup)
	m.Prefix.StringValue = flex.FlattenStringPointer(from.Prefix)
	m.RecordNamePolicy = flex.FlattenStringPointerEmptyAsNull(from.RecordNamePolicy)
	m.RpzDropIpRuleEnabled = flex.FlattenBoolPointer(from.RpzDropIpRuleEnabled)
	m.RpzDropIpRuleMinPrefixLengthIpv4 = flex.FlattenInt64Pointer(from.RpzDropIpRuleMinPrefixLengthIpv4)
	m.RpzDropIpRuleMinPrefixLengthIpv6 = flex.FlattenInt64Pointer(from.RpzDropIpRuleMinPrefixLengthIpv6)
	m.RpzPolicy = flex.FlattenStringPointerEmptyAsNull(from.RpzPolicy)
	m.RpzSeverity = flex.FlattenStringPointerEmptyAsNull(from.RpzSeverity)
	m.RpzType = flex.FlattenStringPointerEmptyAsNull(from.RpzType)
	m.SoaDefaultTtl = flex.FlattenInt64Pointer(from.SoaDefaultTtl)
	m.SoaEmail = flex.FlattenStringPointerEmptyAsNull(from.SoaEmail)
	m.SoaExpire = flex.FlattenInt64Pointer(from.SoaExpire)
	m.SoaNegativeTtl = flex.FlattenInt64Pointer(from.SoaNegativeTtl)
	m.SoaRefresh = flex.FlattenInt64Pointer(from.SoaRefresh)
	m.SoaRetry = flex.FlattenInt64Pointer(from.SoaRetry)
	m.SoaSerialNumber = flex.FlattenInt64Pointer(from.SoaSerialNumber)
	m.SubstituteName = flex.FlattenStringPointerEmptyAsNull(from.SubstituteName)
	m.UseExternalPrimary = flex.FlattenBoolPointer(from.UseExternalPrimary)
	m.View = flex.FlattenStringPointerEmptyAsNull(from.View)
}
