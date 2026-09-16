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

type ZoneDelegatedModel struct {
	Id   types.String `tfsdk:"id"`
	NIOS types.Object `tfsdk:"nios"`
	UDDI types.Object `tfsdk:"uddi"`
}

var ZoneDelegatedAttrTypes = map[string]attr.Type{
	"id":   types.StringType,
	"nios": types.ObjectType{AttrTypes: NIOSZoneDelegatedAttrTypes},
	"uddi": types.ObjectType{AttrTypes: UDDIZoneDelegatedAttrTypes},
}

type NIOSZoneDelegatedModel struct {
	Comment                types.String                        `tfsdk:"comment"`
	DelegateTo             types.List                          `tfsdk:"delegate_to"`
	DelegatedTtl           types.Int64                         `tfsdk:"delegated_ttl"`
	Disable                types.Bool                          `tfsdk:"disable"`
	EnableRfc2317Exclusion types.Bool                          `tfsdk:"enable_rfc2317_exclusion"`
	ExtAttrs               types.Map                           `tfsdk:"ext_attrs"`
	ExtAttrsAll            types.Map                           `tfsdk:"ext_attrs_all"`
	Fqdn                   types.String                        `tfsdk:"fqdn"`
	Locked                 types.Bool                          `tfsdk:"locked"`
	MsAdIntegrated         types.Bool                          `tfsdk:"ms_ad_integrated"`
	MsDdnsMode             types.String                        `tfsdk:"ms_ddns_mode"`
	NsGroup                types.String                        `tfsdk:"ns_group"`
	Prefix                 internaltypes.CaseInsensitiveString `tfsdk:"prefix"`
	View                   types.String                        `tfsdk:"view"`
	ZoneFormat             types.String                        `tfsdk:"zone_format"`
}

var NIOSZoneDelegatedAttrTypes = map[string]attr.Type{
	"comment":                  types.StringType,
	"delegate_to":              types.ListType{ElemType: types.ObjectType{AttrTypes: ZoneDelegatedDelegateToAttrTypes}},
	"delegated_ttl":            types.Int64Type,
	"disable":                  types.BoolType,
	"enable_rfc2317_exclusion": types.BoolType,
	"ext_attrs":                types.MapType{ElemType: types.StringType},
	"ext_attrs_all":            types.MapType{ElemType: types.StringType},
	"fqdn":                     types.StringType,
	"locked":                   types.BoolType,
	"ms_ad_integrated":         types.BoolType,
	"ms_ddns_mode":             types.StringType,
	"ns_group":                 types.StringType,
	"prefix":                   internaltypes.CaseInsensitiveStringType{},
	"view":                     types.StringType,
	"zone_format":              types.StringType,
}

type UDDIZoneDelegatedModel struct {
	Comment           types.String `tfsdk:"comment"`
	CompartmentId     types.String `tfsdk:"compartment_id"`
	DelegationServers types.List   `tfsdk:"delegation_servers"`
	Disabled          types.Bool   `tfsdk:"disabled"`
	Fqdn              types.String `tfsdk:"fqdn"`
	Parent            types.String `tfsdk:"parent"`
	Tags              types.Map    `tfsdk:"tags"`
	TagsAll           types.Map    `tfsdk:"tags_all"`
	View              types.String `tfsdk:"view"`
}

var UDDIZoneDelegatedAttrTypes = map[string]attr.Type{
	"comment":            types.StringType,
	"compartment_id":     types.StringType,
	"delegation_servers": types.ListType{ElemType: types.ObjectType{AttrTypes: DelegationServerAttrTypes}},
	"disabled":           types.BoolType,
	"fqdn":               types.StringType,
	"parent":             types.StringType,
	"tags":               types.MapType{ElemType: types.StringType},
	"tags_all":           types.MapType{ElemType: types.StringType},
	"view":               types.StringType,
}

const (
	ZoneDelegatedReturnFields = "address,comment,delegate_to,delegated_ttl,disable,display_domain,dns_fqdn,enable_rfc2317_exclusion,extattrs,fqdn,locked,locked_by,mask_prefix,ms_ad_integrated,ms_ddns_mode,ms_managed,ms_read_only,ms_sync_master_name,ns_group,parent,prefix,use_delegated_ttl,using_srg_associations,view,zone_format"
)

var ZoneDelegatedResourceSchemaAttributes = map[string]schema.Attribute{
	"id": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The reference to the object.",
	},
	"nios": schema.SingleNestedAttribute{
		Optional:            true,
		MarkdownDescription: "NIOS backend-specific fields.",
		Attributes:          ZoneDelegatedResourceNiosSchemaAttributes,
	},
	"uddi": schema.SingleNestedAttribute{
		Optional:            true,
		MarkdownDescription: "UDDI backend-specific fields.",
		Attributes:          ZoneDelegatedResourceUddiSchemaAttributes,
	},
}

var ZoneDelegatedResourceNiosSchemaAttributes = map[string]schema.Attribute{
	"comment": schema.StringAttribute{
		Optional: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
			stringvalidator.LengthBetween(0, 256),
			customvalidator.ValidateTrimmedString(),
		},
		MarkdownDescription: "Comment for the zone; maximum 256 characters.",
	},
	"delegate_to": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: ZoneDelegatedDelegateToResourceSchemaAttributes,
		},
		Optional: true,
		Validators: []validator.List{
			customvalidator.ListNotEmpty(),
		},
		MarkdownDescription: "This provides information for the remote name server that maintains data for the delegated zone. The Infoblox appliance redirects queries for data for the delegated zone to this remote name server.",
	},
	"delegated_ttl": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "You can specify the Time to Live (TTL) values of auto-generated NS and glue records for a delegated zone. This value is the number of seconds that data is cached.",
	},
	"disable": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Determines whether a zone is disabled or not. When this is set to False, the zone is enabled.",
	},
	"enable_rfc2317_exclusion": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "This flag controls whether automatic generation of RFC 2317 CNAMEs for delegated reverse zones overwrite existing PTR records. The default behavior is to overwrite all the existing records in the range; this corresponds to \"allow_ptr_creation_in_parent\" set to False. However, when this flag is set to True the existing PTR records are not overwritten.",
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
	"fqdn": schema.StringAttribute{
		Required: true,
		PlanModifiers: []planmodifier.String{
			immutable.ImmutableString(),
		},
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
			customvalidator.IsValidNIOSDomainName(),
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
		MarkdownDescription: "The delegation NS group bound with delegated zone.",
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

var ZoneDelegatedResourceUddiSchemaAttributes = map[string]schema.Attribute{
	"comment": schema.StringAttribute{
		Default:             stringdefault.StaticString(""),
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "Optional. Comment for zone delegation.",
	},
	"compartment_id": schema.StringAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The access view associated with the object. If no access view is associated with the object, the value defaults to empty.",
	},
	"delegation_servers": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: DelegationServerResourceSchemaAttributes,
		},
		Required: true,
		Validators: []validator.List{
			customvalidator.ListNotEmpty(),
		},
		MarkdownDescription: "Required. DNS zone delegation servers. Order is not significant.",
	},
	"disabled": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Optional. _true_ to disable object. A disabled object is effectively non-existent when generating resource records.",
	},
	"fqdn": schema.StringAttribute{
		Required: true,
		PlanModifiers: []planmodifier.String{
			stringplanmodifier.RequiresReplaceIfConfigured(),
		},
		Validators: []validator.String{
			customvalidator.IsValidUDDIDomainName(),
		},
		MarkdownDescription: "Delegation FQDN. The FQDN supplied at creation will be converted to canonical form.  Read-only after creation.",
	},
	"parent": schema.StringAttribute{
		Optional:            true,
		Computed:            true,
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
func (m *ZoneDelegatedModel) Expand(ctx context.Context, diags *diag.Diagnostics, isCreate bool) *coremodel.ZoneDelegated {
	if m == nil {
		return nil
	}

	obj := &coremodel.ZoneDelegated{}

	// Expand NIOS nested attribute (returns nil if not present)
	niosModel := flex.ExpandNestedObject[NIOSZoneDelegatedModel](ctx, m.NIOS, diags)
	if niosModel != nil {
		obj.NIOS = niosModel.Expand(ctx, diags, isCreate)
	}

	// Expand UDDI nested attribute (returns nil if not present)
	uddiModel := flex.ExpandNestedObject[UDDIZoneDelegatedModel](ctx, m.UDDI, diags)
	if uddiModel != nil {
		obj.UDDI = uddiModel.Expand(ctx, diags, isCreate)
	}

	return obj
}

// Expand converts the NIOS TF model to the core model.
func (m *NIOSZoneDelegatedModel) Expand(ctx context.Context, diags *diag.Diagnostics, isCreate bool) *coremodel.NIOSZoneDelegatedExt {
	ext := &coremodel.NIOSZoneDelegatedExt{
		Comment:                flex.ExpandStringPointerNullAsEmpty(m.Comment),
		DelegateTo:             flex.ExpandFrameworkListNestedBlock(ctx, m.DelegateTo, diags, ExpandZoneDelegatedDelegateTo),
		DelegatedTtl:           flex.ExpandInt64Pointer(m.DelegatedTtl),
		Disable:                flex.ExpandBoolPointer(m.Disable),
		EnableRfc2317Exclusion: flex.ExpandBoolPointer(m.EnableRfc2317Exclusion),
		ExtAttrs:               flex.ExpandMapStringAny(ctx, m.ExtAttrs, diags),
		Locked:                 flex.ExpandBoolPointer(m.Locked),
		MsAdIntegrated:         flex.ExpandBoolPointer(m.MsAdIntegrated),
		MsDdnsMode:             flex.ExpandStringPointerNullAsEmpty(m.MsDdnsMode),
		NsGroup:                flex.ExpandStringPointer(m.NsGroup),
		Prefix:                 flex.ExpandStringPointer(m.Prefix.StringValue),
	}
	if isCreate {
		ext.Fqdn = flex.ExpandStringPointerNullAsEmpty(m.Fqdn)
		ext.View = flex.ExpandStringPointerNullAsEmpty(m.View)
		ext.ZoneFormat = flex.ExpandStringPointerNullAsEmpty(m.ZoneFormat)
	}
	return ext
}

// ApplyZoneDelegatedNIOSUseFlags derives NIOS use flags from the raw config
// value(s) and writes them onto the core model. A flag is true when the user
// set any of its governed value fields in config.
func ApplyZoneDelegatedNIOSUseFlags(ctx context.Context, config tfsdk.Config, obj *coremodel.ZoneDelegated, diags *diag.Diagnostics) {
	if obj == nil || obj.NIOS == nil {
		return
	}
	obj.NIOS.UseDelegatedTtl = flex.DeriveUseFlag(ctx, config, diags, path.Root("nios").AtName("delegated_ttl"))
}

// Expand converts the UDDI TF model to the core model.
func (m *UDDIZoneDelegatedModel) Expand(ctx context.Context, diags *diag.Diagnostics, isCreate bool) *coremodel.UDDIZoneDelegatedExt {
	ext := &coremodel.UDDIZoneDelegatedExt{
		Comment:           flex.ExpandStringPointer(m.Comment),
		CompartmentId:     flex.ExpandStringPointer(m.CompartmentId),
		DelegationServers: flex.ExpandFrameworkListNestedBlock(ctx, m.DelegationServers, diags, ExpandDelegationServer),
		Disabled:          flex.ExpandBoolPointer(m.Disabled),
		Parent:            flex.ExpandStringPointer(m.Parent),
		Tags:              flex.ExpandMapStringAny(ctx, m.Tags, diags),
	}
	if isCreate {
		ext.Fqdn = flex.ExpandStringPointer(m.Fqdn)
		ext.View = flex.ExpandStringPointer(m.View)
	}
	return ext
}

// Flatten populates the TF model from a core response.
func (m *ZoneDelegatedModel) Flatten(ctx context.Context, resp *coremodel.ZoneDelegated, diags *diag.Diagnostics) {
	if resp == nil {
		return
	}

	m.Id = flex.FlattenStringPointer(resp.Id)

	// Extract existing NIOS model, flatten API response onto it, convert back
	niosModel := flex.ExpandNestedObject[NIOSZoneDelegatedModel](ctx, m.NIOS, diags)
	if niosModel == nil {
		niosModel = &NIOSZoneDelegatedModel{}
	}
	niosModel.Flatten(ctx, resp.NIOS, diags)
	if resp.NIOS != nil {
		m.NIOS = flex.FlattenNestedObject(ctx, niosModel, NIOSZoneDelegatedAttrTypes, diags)
	} else {
		m.NIOS = types.ObjectNull(NIOSZoneDelegatedAttrTypes)
	}

	// Extract existing UDDI model, flatten API response onto it, convert back
	uddiModel := flex.ExpandNestedObject[UDDIZoneDelegatedModel](ctx, m.UDDI, diags)
	if uddiModel == nil {
		uddiModel = &UDDIZoneDelegatedModel{}
	}
	uddiModel.Flatten(ctx, resp.UDDI, diags)
	if resp.UDDI != nil {
		m.UDDI = flex.FlattenNestedObject(ctx, uddiModel, UDDIZoneDelegatedAttrTypes, diags)
	} else {
		m.UDDI = types.ObjectNull(UDDIZoneDelegatedAttrTypes)
	}
}

// Flatten merges API response onto existing NIOS model.
func (m *NIOSZoneDelegatedModel) Flatten(ctx context.Context, from *coremodel.NIOSZoneDelegatedExt, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	planExtAttrs := m.ExtAttrs
	if planExtAttrs.IsUnknown() {
		planExtAttrs = types.MapNull(types.StringType)
	}
	m.Comment = flex.FlattenStringPointerEmptyAsNull(from.Comment)
	m.DelegateTo = flex.FlattenFrameworkListNestedBlock(ctx, from.DelegateTo, ZoneDelegatedDelegateToAttrTypes, diags, FlattenZoneDelegatedDelegateTo)
	m.DelegatedTtl = flex.FlattenInt64Pointer(from.DelegatedTtl)
	m.Disable = flex.FlattenBoolPointer(from.Disable)
	m.EnableRfc2317Exclusion = flex.FlattenBoolPointer(from.EnableRfc2317Exclusion)
	m.ExtAttrs, m.ExtAttrsAll = flex.FlattenEAs(planExtAttrs, from.ExtAttrs)
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
func (m *UDDIZoneDelegatedModel) Flatten(ctx context.Context, from *coremodel.UDDIZoneDelegatedExt, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	m.Comment = flex.FlattenStringPointer(from.Comment)
	m.CompartmentId = flex.FlattenStringPointer(from.CompartmentId)
	m.DelegationServers = flex.FlattenFrameworkListNestedBlock(ctx, from.DelegationServers, DelegationServerAttrTypes, diags, FlattenDelegationServer)
	m.Disabled = flex.FlattenBoolPointer(from.Disabled)
	m.Fqdn = flex.FlattenStringPointer(from.Fqdn)
	m.Parent = flex.FlattenStringPointer(from.Parent)
	tagsAll := flex.FlattenMapStringAny(ctx, from.Tags, diags)
	if m.Tags.IsNull() || m.Tags.IsUnknown() {
		m.Tags = tagsAll
	}
	m.TagsAll = tagsAll
	m.View = flex.FlattenStringPointer(from.View)
}
