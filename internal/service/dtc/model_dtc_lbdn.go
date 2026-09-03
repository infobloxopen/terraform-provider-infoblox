package dtc

import (
	"context"

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

	coremodel "github.com/infobloxopen/terraform-provider-infoblox/internal/core/model/dtc"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
	importmod "github.com/infobloxopen/terraform-provider-infoblox/internal/planmodifiers/import"
	internaltypes "github.com/infobloxopen/terraform-provider-infoblox/internal/types"
	customvalidator "github.com/infobloxopen/terraform-provider-infoblox/internal/validator"
)

type DtcLbdnModel struct {
	Id   types.String `tfsdk:"id"`
	NIOS types.Object `tfsdk:"nios"`
	UDDI types.Object `tfsdk:"uddi"`
}

var DtcLbdnAttrTypes = map[string]attr.Type{
	"id":   types.StringType,
	"nios": types.ObjectType{AttrTypes: NIOSDtcLbdnAttrTypes},
	"uddi": types.ObjectType{AttrTypes: UDDIDtcLbdnAttrTypes},
}

type NIOSDtcLbdnModel struct {
	AuthZones                internaltypes.UnorderedListValue `tfsdk:"auth_zones"`
	AutoConsolidatedMonitors types.Bool                       `tfsdk:"auto_consolidated_monitors"`
	Comment                  types.String                     `tfsdk:"comment"`
	Disable                  types.Bool                       `tfsdk:"disable"`
	ExtAttrs                 types.Map                        `tfsdk:"ext_attrs"`
	ExtAttrsAll              types.Map                        `tfsdk:"ext_attrs_all"`
	LbMethod                 types.String                     `tfsdk:"lb_method"`
	Name                     types.String                     `tfsdk:"name"`
	Patterns                 internaltypes.UnorderedListValue `tfsdk:"patterns"`
	Persistence              types.Int64                      `tfsdk:"persistence"`
	Pools                    types.List                       `tfsdk:"pools"`
	Priority                 types.Int64                      `tfsdk:"priority"`
	Topology                 types.String                     `tfsdk:"topology"`
	Ttl                      types.Int64                      `tfsdk:"ttl"`
	Types                    internaltypes.UnorderedListValue `tfsdk:"types"`
}

var NIOSDtcLbdnAttrTypes = map[string]attr.Type{
	"auth_zones":                 internaltypes.UnorderedListOfStringType,
	"auto_consolidated_monitors": types.BoolType,
	"comment":                    types.StringType,
	"disable":                    types.BoolType,
	"ext_attrs":                  types.MapType{ElemType: types.StringType},
	"ext_attrs_all":              types.MapType{ElemType: types.StringType},
	"lb_method":                  types.StringType,
	"name":                       types.StringType,
	"patterns":                   internaltypes.UnorderedListOfStringType,
	"persistence":                types.Int64Type,
	"pools":                      types.ListType{ElemType: types.ObjectType{AttrTypes: LbdnPoolsAttrTypes}},
	"priority":                   types.Int64Type,
	"topology":                   types.StringType,
	"ttl":                        types.Int64Type,
	"types":                      internaltypes.UnorderedListOfStringType,
}

type UDDIDtcLbdnModel struct {
	Comment            types.String `tfsdk:"comment"`
	Disabled           types.Bool   `tfsdk:"disabled"`
	DtcPolicy          types.Object `tfsdk:"dtc_policy"`
	InheritanceSources types.Object `tfsdk:"inheritance_sources"`
	Name               types.String `tfsdk:"name"`
	Precedence         types.Int64  `tfsdk:"precedence"`
	Tags               types.Map    `tfsdk:"tags"`
	TagsAll            types.Map    `tfsdk:"tags_all"`
	Ttl                types.Int64  `tfsdk:"ttl"`
	View               types.String `tfsdk:"view"`
}

var UDDIDtcLbdnAttrTypes = map[string]attr.Type{
	"comment":             types.StringType,
	"disabled":            types.BoolType,
	"dtc_policy":          types.ObjectType{AttrTypes: DTCPolicyAttrTypes},
	"inheritance_sources": types.ObjectType{AttrTypes: TTLInheritanceDnsconfigAttrTypes},
	"name":                types.StringType,
	"precedence":          types.Int64Type,
	"tags":                types.MapType{ElemType: types.StringType},
	"tags_all":            types.MapType{ElemType: types.StringType},
	"ttl":                 types.Int64Type,
	"view":                types.StringType,
}

const (
	DtcLbdnInheritanceType = "full"
	DtcLbdnReturnFields    = "auth_zones,auto_consolidated_monitors,comment,disable,extattrs,health,lb_method,name,patterns,persistence,pools,priority,topology,ttl,types,use_ttl"
)

var DtcLbdnResourceSchemaAttributes = map[string]schema.Attribute{
	"id": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The reference to the object.",
	},
	"nios": schema.SingleNestedAttribute{
		Optional:            true,
		MarkdownDescription: "NIOS backend-specific fields.",
		Attributes:          DtcLbdnResourceNiosSchemaAttributes,
	},
	"uddi": schema.SingleNestedAttribute{
		Optional:            true,
		MarkdownDescription: "UDDI backend-specific fields.",
		Attributes:          DtcLbdnResourceUddiSchemaAttributes,
	},
}

var DtcLbdnResourceNiosSchemaAttributes = map[string]schema.Attribute{
	"auth_zones": schema.ListAttribute{
		ElementType: types.StringType,
		Optional:    true,
		CustomType:  internaltypes.UnorderedListOfStringType,
		Validators: []validator.List{
			customvalidator.ListNotEmpty(),
		},
		MarkdownDescription: "List of linked auth zones.",
	},
	"auto_consolidated_monitors": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Flag for enabling auto managing DTC Consolidated Monitors on related DTC Pools.",
	},
	"comment": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
			customvalidator.ValidateTrimmedString(),
			stringvalidator.LengthBetween(0, 256),
		},
		MarkdownDescription: "Comment for the DTC LBDN; maximum 256 characters.",
	},
	"disable": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Determines whether the DTC LBDN is disabled or not. When this is set to False, the fixed address is enabled.",
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
	"lb_method": schema.StringAttribute{
		Validators: []validator.String{
			stringvalidator.OneOf("GLOBAL_AVAILABILITY", "RATIO", "ROUND_ROBIN", "SOURCE_IP_HASH", "TOPOLOGY"),
		},
		Required:            true,
		MarkdownDescription: "The load balancing method. Used to select pool.",
	},
	"name": schema.StringAttribute{
		Required: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
			customvalidator.ValidateTrimmedString(),
		},
		MarkdownDescription: "The display name of the DTC LBDN, not DNS related.",
	},
	"patterns": schema.ListAttribute{
		ElementType: types.StringType,
		Optional:    true,
		CustomType:  internaltypes.UnorderedListOfStringType,
		Validators: []validator.List{
			customvalidator.ListNotEmpty(),
		},
		MarkdownDescription: "LBDN wildcards for pattern match.",
	},
	"persistence": schema.Int64Attribute{
		Optional:            true,
		Computed:            true,
		Default:             int64default.StaticInt64(0),
		MarkdownDescription: "Maximum time, in seconds, for which client specific LBDN responses will be cached. Zero specifies no caching.",
	},
	"pools": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: LbdnPoolsResourceSchemaAttributes,
		},
		Optional: true,
		Computed: true,
		Validators: []validator.List{
			customvalidator.ListNotEmpty(),
		},
		MarkdownDescription: "The maximum time, in seconds, for which client specific LBDN responses will be cached. Zero specifies no caching.",
	},
	"priority": schema.Int64Attribute{
		Optional:            true,
		Computed:            true,
		Default:             int64default.StaticInt64(1),
		MarkdownDescription: "The LBDN pattern match priority for \"overlapping\" DTC LBDN objects. LBDNs are \"overlapping\" if they are simultaneously assigned to a zone and have patterns that can match the same FQDN. The matching LBDN with highest priority (lowest ordinal) will be used.",
	},
	"topology": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
		},
		MarkdownDescription: "The topology rules for TOPOLOGY method.",
	},
	"ttl": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "The Time To Live (TTL) value for the DTC LBDN. A 32-bit unsigned integer that represents the duration, in seconds, for which the record is valid (cached). Zero indicates that the record should not be cached.",
	},
	"types": schema.ListAttribute{
		ElementType: types.StringType,
		Optional:    true,
		Computed:    true,
		CustomType:  internaltypes.UnorderedListOfStringType,
		Default:     listdefault.StaticValue(types.ListValueMust(types.StringType, []attr.Value{types.StringValue("A"), types.StringValue("AAAA")})),
		Validators: []validator.List{
			customvalidator.ListNotEmpty(),
			customvalidator.StringsInSlice([]string{"A", "AAAA", "CNAME", "NAPTR", "SRV"}),
		},
		MarkdownDescription: "The list of resource record types supported by LBDN.",
	},
}

var DtcLbdnResourceUddiSchemaAttributes = map[string]schema.Attribute{
	"comment": schema.StringAttribute{
		Default:             stringdefault.StaticString(""),
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "Optional. Comment for __LBDN__.",
	},
	"disabled": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Optional. _true_ to disable object. A disabled object is effectively non-existent when generating configuration.",
	},
	"dtc_policy": schema.SingleNestedAttribute{
		Attributes:          DTCPolicyResourceSchemaAttributes,
		Optional:            true,
		MarkdownDescription: "The __DTC Policy__ object.",
	},
	"inheritance_sources": schema.SingleNestedAttribute{
		Attributes:          TTLInheritanceDnsconfigResourceSchemaAttributes,
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The inheritance configuration specifies how the object inherits the _ttl_ field.",
	},
	"name": schema.StringAttribute{
		Required: true,
		Validators: []validator.String{
			customvalidator.IsValidUDDIDomainName(),
		},
		MarkdownDescription: "Name of __LBDN__.",
	},
	"precedence": schema.Int64Attribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "Optional. Precedence.",
	},
	"tags": schema.MapAttribute{
		Optional:    true,
		Computed:    true,
		ElementType: types.StringType,
		Default:     mapdefault.StaticValue(types.MapNull(types.StringType)),
		Validators: []validator.Map{
			mapvalidator.SizeAtLeast(1),
		},
		MarkdownDescription: "Optional. The tags for __LBDN__ in JSON format.",
	},
	"tags_all": schema.MapAttribute{
		Computed:            true,
		ElementType:         types.StringType,
		MarkdownDescription: "All tags including inherited values.",
	},
	"ttl": schema.Int64Attribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "Optional. Time to live value (in seconds) to be used for records in DTC response. Unsigned integer, min: 0, max 2147483647 (31-bits per RFC-2181).",
	},
	"view": schema.StringAttribute{
		Required:            true,
		MarkdownDescription: "The resource identifier.",
	},
}

// Expand converts the TF model to the infoblox core model
func (m *DtcLbdnModel) Expand(ctx context.Context, diags *diag.Diagnostics, isCreate bool) *coremodel.DtcLbdn {
	if m == nil {
		return nil
	}

	obj := &coremodel.DtcLbdn{}

	// Expand NIOS nested attribute (returns nil if not present)
	niosModel := flex.ExpandNestedObject[NIOSDtcLbdnModel](ctx, m.NIOS, diags)
	if niosModel != nil {
		obj.NIOS = niosModel.Expand(ctx, diags)
	}

	// Expand UDDI nested attribute (returns nil if not present)
	uddiModel := flex.ExpandNestedObject[UDDIDtcLbdnModel](ctx, m.UDDI, diags)
	if uddiModel != nil {
		obj.UDDI = uddiModel.Expand(ctx, diags)
	}

	return obj
}

// Expand converts the NIOS TF model to the core model.
func (m *NIOSDtcLbdnModel) Expand(ctx context.Context, diags *diag.Diagnostics) *coremodel.NIOSDtcLbdnExt {
	return &coremodel.NIOSDtcLbdnExt{
		AuthZones:                flex.ExpandFrameworkListString(ctx, m.AuthZones, diags),
		AutoConsolidatedMonitors: flex.ExpandBoolPointer(m.AutoConsolidatedMonitors),
		Comment:                  flex.ExpandStringPointerNullAsEmpty(m.Comment),
		Disable:                  flex.ExpandBoolPointer(m.Disable),
		ExtAttrs:                 flex.ExpandMapStringAny(ctx, m.ExtAttrs, diags),
		LbMethod:                 flex.ExpandStringPointerNullAsEmpty(m.LbMethod),
		Name:                     flex.ExpandStringPointerNullAsEmpty(m.Name),
		Patterns:                 flex.ExpandFrameworkListString(ctx, m.Patterns, diags),
		Persistence:              flex.ExpandInt64Pointer(m.Persistence),
		Pools:                    flex.ExpandFrameworkListNestedBlock(ctx, m.Pools, diags, ExpandLbdnPools),
		Priority:                 flex.ExpandInt64Pointer(m.Priority),
		Topology:                 flex.ExpandStringPointer(m.Topology),
		Ttl:                      flex.ExpandInt64Pointer(m.Ttl),
		Types:                    flex.ExpandFrameworkListString(ctx, m.Types, diags),
	}
}

// ApplyDtcLbdnNIOSUseFlags derives NIOS use flags from the raw config
// value(s) and writes them onto the core model. A flag is true when the user
// set any of its governed value fields in config.
func ApplyDtcLbdnNIOSUseFlags(ctx context.Context, config tfsdk.Config, obj *coremodel.DtcLbdn, diags *diag.Diagnostics) {
	if obj == nil || obj.NIOS == nil {
		return
	}
	obj.NIOS.UseTtl = flex.DeriveUseFlag(ctx, config, diags, path.Root("nios").AtName("ttl"))
}

// Expand converts the UDDI TF model to the core model.
func (m *UDDIDtcLbdnModel) Expand(ctx context.Context, diags *diag.Diagnostics) *coremodel.UDDIDtcLbdnExt {
	return &coremodel.UDDIDtcLbdnExt{
		Comment:            flex.ExpandStringPointer(m.Comment),
		Disabled:           flex.ExpandBoolPointer(m.Disabled),
		DtcPolicy:          ExpandDTCPolicy(ctx, m.DtcPolicy, diags),
		InheritanceSources: ExpandTTLInheritanceDnsconfig(ctx, m.InheritanceSources, diags),
		Name:               flex.ExpandString(m.Name),
		Precedence:         flex.ExpandInt64Pointer(m.Precedence),
		Tags:               flex.ExpandMapStringAny(ctx, m.Tags, diags),
		Ttl:                flex.ExpandInt64Pointer(m.Ttl),
		View:               flex.ExpandString(m.View),
	}
}

// Flatten populates the TF model from a core response.
func (m *DtcLbdnModel) Flatten(ctx context.Context, resp *coremodel.DtcLbdn, diags *diag.Diagnostics) {
	if resp == nil {
		return
	}

	m.Id = flex.FlattenStringPointer(resp.Id)

	// Extract existing NIOS model, flatten API response onto it, convert back
	niosModel := flex.ExpandNestedObject[NIOSDtcLbdnModel](ctx, m.NIOS, diags)
	if niosModel == nil {
		niosModel = &NIOSDtcLbdnModel{}
	}
	niosModel.Flatten(ctx, resp.NIOS, diags)
	if resp.NIOS != nil {
		m.NIOS = flex.FlattenNestedObject(ctx, niosModel, NIOSDtcLbdnAttrTypes, diags)
	} else {
		m.NIOS = types.ObjectNull(NIOSDtcLbdnAttrTypes)
	}

	// Extract existing UDDI model, flatten API response onto it, convert back
	uddiModel := flex.ExpandNestedObject[UDDIDtcLbdnModel](ctx, m.UDDI, diags)
	if uddiModel == nil {
		uddiModel = &UDDIDtcLbdnModel{}
	}
	uddiModel.Flatten(ctx, resp.UDDI, diags)
	if resp.UDDI != nil {
		m.UDDI = flex.FlattenNestedObject(ctx, uddiModel, UDDIDtcLbdnAttrTypes, diags)
	} else {
		m.UDDI = types.ObjectNull(UDDIDtcLbdnAttrTypes)
	}
}

// Flatten merges API response onto existing NIOS model.
func (m *NIOSDtcLbdnModel) Flatten(ctx context.Context, from *coremodel.NIOSDtcLbdnExt, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	planExtAttrs := m.ExtAttrs
	if planExtAttrs.IsUnknown() {
		planExtAttrs = types.MapNull(types.StringType)
	}
	m.AuthZones = flex.FlattenFrameworkUnorderedListString(ctx, from.AuthZones, diags)
	m.AutoConsolidatedMonitors = flex.FlattenBoolPointer(from.AutoConsolidatedMonitors)
	m.Comment = flex.FlattenStringPointerEmptyAsNull(from.Comment)
	m.Disable = flex.FlattenBoolPointer(from.Disable)
	m.ExtAttrs, m.ExtAttrsAll = flex.FlattenEAs(planExtAttrs, from.ExtAttrs)
	m.LbMethod = flex.FlattenStringPointerEmptyAsNull(from.LbMethod)
	m.Name = flex.FlattenStringPointerEmptyAsNull(from.Name)
	m.Patterns = flex.FlattenFrameworkUnorderedListString(ctx, from.Patterns, diags)
	m.Persistence = flex.FlattenInt64Pointer(from.Persistence)
	m.Pools = flex.FlattenFrameworkListNestedBlock(ctx, from.Pools, LbdnPoolsAttrTypes, diags, FlattenLbdnPools)
	m.Priority = flex.FlattenInt64Pointer(from.Priority)
	m.Topology = flex.FlattenStringPointerEmptyAsNull(from.Topology)
	m.Ttl = flex.FlattenInt64Pointer(from.Ttl)
	m.Types = flex.FlattenFrameworkUnorderedListString(ctx, from.Types, diags)
}

// Flatten merges API response onto existing UDDI model.
func (m *UDDIDtcLbdnModel) Flatten(ctx context.Context, from *coremodel.UDDIDtcLbdnExt, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	m.Comment = flex.FlattenStringPointer(from.Comment)
	m.Disabled = flex.FlattenBoolPointer(from.Disabled)
	m.DtcPolicy = FlattenDTCPolicy(ctx, from.DtcPolicy, diags)
	m.InheritanceSources = FlattenTTLInheritanceDnsconfig(ctx, from.InheritanceSources, diags)
	m.Name = flex.FlattenString(from.Name)
	m.Precedence = flex.FlattenInt64Pointer(from.Precedence)
	tagsAll := flex.FlattenMapStringAny(ctx, from.Tags, diags)
	if m.Tags.IsNull() || m.Tags.IsUnknown() {
		m.Tags = tagsAll
	}
	m.TagsAll = tagsAll
	m.Ttl = flex.FlattenInt64Pointer(from.Ttl)
	m.View = flex.FlattenString(from.View)
}
