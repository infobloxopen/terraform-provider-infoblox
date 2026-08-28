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

type DtcPoolModel struct {
	Id   types.String `tfsdk:"id"`
	NIOS types.Object `tfsdk:"nios"`
	UDDI types.Object `tfsdk:"uddi"`
}

var DtcPoolAttrTypes = map[string]attr.Type{
	"id":   types.StringType,
	"nios": types.ObjectType{AttrTypes: NIOSDtcPoolAttrTypes},
	"uddi": types.ObjectType{AttrTypes: UDDIDtcPoolAttrTypes},
}

type NIOSDtcPoolModel struct {
	AutoConsolidatedMonitors types.Bool                       `tfsdk:"auto_consolidated_monitors"`
	Availability             types.String                     `tfsdk:"availability"`
	Comment                  types.String                     `tfsdk:"comment"`
	ConsolidatedMonitors     types.List                       `tfsdk:"consolidated_monitors"`
	Disable                  types.Bool                       `tfsdk:"disable"`
	ExtAttrs                 types.Map                        `tfsdk:"ext_attrs"`
	ExtAttrsAll              types.Map                        `tfsdk:"ext_attrs_all"`
	LbAlternateMethod        types.String                     `tfsdk:"lb_alternate_method"`
	LbAlternateTopology      types.String                     `tfsdk:"lb_alternate_topology"`
	LbDynamicRatioAlternate  types.Object                     `tfsdk:"lb_dynamic_ratio_alternate"`
	LbDynamicRatioPreferred  types.Object                     `tfsdk:"lb_dynamic_ratio_preferred"`
	LbPreferredMethod        types.String                     `tfsdk:"lb_preferred_method"`
	LbPreferredTopology      types.String                     `tfsdk:"lb_preferred_topology"`
	Monitors                 internaltypes.UnorderedListValue `tfsdk:"monitors"`
	Name                     types.String                     `tfsdk:"name"`
	Quorum                   types.Int64                      `tfsdk:"quorum"`
	Servers                  types.List                       `tfsdk:"servers"`
	Ttl                      types.Int64                      `tfsdk:"ttl"`
}

var NIOSDtcPoolAttrTypes = map[string]attr.Type{
	"auto_consolidated_monitors": types.BoolType,
	"availability":               types.StringType,
	"comment":                    types.StringType,
	"consolidated_monitors":      types.ListType{ElemType: types.ObjectType{AttrTypes: PoolConsolidatedMonitorsAttrTypes}},
	"disable":                    types.BoolType,
	"ext_attrs":                  types.MapType{ElemType: types.StringType},
	"ext_attrs_all":              types.MapType{ElemType: types.StringType},
	"lb_alternate_method":        types.StringType,
	"lb_alternate_topology":      types.StringType,
	"lb_dynamic_ratio_alternate": types.ObjectType{AttrTypes: PoolLbDynamicRatioAlternateAttrTypes},
	"lb_dynamic_ratio_preferred": types.ObjectType{AttrTypes: PoolLbDynamicRatioPreferredAttrTypes},
	"lb_preferred_method":        types.StringType,
	"lb_preferred_topology":      types.StringType,
	"monitors":                   internaltypes.UnorderedListOfStringType,
	"name":                       types.StringType,
	"quorum":                     types.Int64Type,
	"servers":                    types.ListType{ElemType: types.ObjectType{AttrTypes: PoolServersAttrTypes}},
	"ttl":                        types.Int64Type,
}

type UDDIDtcPoolModel struct {
	Comment                   types.String `tfsdk:"comment"`
	ConsolidatedHealthEnabled types.Bool   `tfsdk:"consolidated_health_enabled"`
	Disabled                  types.Bool   `tfsdk:"disabled"`
	HealthChecks              types.List   `tfsdk:"health_checks"`
	InheritanceSources        types.Object `tfsdk:"inheritance_sources"`
	Method                    types.String `tfsdk:"method"`
	Name                      types.String `tfsdk:"name"`
	PoolAvailability          types.String `tfsdk:"pool_availability"`
	PoolServersQuorum         types.Int64  `tfsdk:"pool_servers_quorum"`
	ServerAvailability        types.String `tfsdk:"server_availability"`
	ServerHealthChecksQuorum  types.Int64  `tfsdk:"server_health_checks_quorum"`
	Servers                   types.List   `tfsdk:"servers"`
	Tags                      types.Map    `tfsdk:"tags"`
	TagsAll                   types.Map    `tfsdk:"tags_all"`
	Ttl                       types.Int64  `tfsdk:"ttl"`
}

var UDDIDtcPoolAttrTypes = map[string]attr.Type{
	"comment":                     types.StringType,
	"consolidated_health_enabled": types.BoolType,
	"disabled":                    types.BoolType,
	"health_checks":               types.ListType{ElemType: types.ObjectType{AttrTypes: PoolHealthCheckAttrTypes}},
	"inheritance_sources":         types.ObjectType{AttrTypes: TTLInheritanceAttrTypes},
	"method":                      types.StringType,
	"name":                        types.StringType,
	"pool_availability":           types.StringType,
	"pool_servers_quorum":         types.Int64Type,
	"server_availability":         types.StringType,
	"server_health_checks_quorum": types.Int64Type,
	"servers":                     types.ListType{ElemType: types.ObjectType{AttrTypes: PoolServerAttrTypes}},
	"tags":                        types.MapType{ElemType: types.StringType},
	"tags_all":                    types.MapType{ElemType: types.StringType},
	"ttl":                         types.Int64Type,
}

const (
	DtcPoolInheritanceType = "full"
	DtcPoolReturnFields    = "auto_consolidated_monitors,availability,comment,consolidated_monitors,disable,extattrs,health,lb_alternate_method,lb_alternate_topology,lb_dynamic_ratio_alternate,lb_dynamic_ratio_preferred,lb_preferred_method,lb_preferred_topology,monitors,name,quorum,servers,ttl,use_ttl"
)

var DtcPoolResourceSchemaAttributes = map[string]schema.Attribute{
	"id": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The reference to the object.",
	},
	"nios": schema.SingleNestedAttribute{
		Optional:            true,
		MarkdownDescription: "NIOS backend-specific fields.",
		Attributes:          DtcPoolResourceNiosSchemaAttributes,
	},
	"uddi": schema.SingleNestedAttribute{
		Optional:            true,
		MarkdownDescription: "UDDI backend-specific fields.",
		Attributes:          DtcPoolResourceUddiSchemaAttributes,
	},
}

var DtcPoolResourceNiosSchemaAttributes = map[string]schema.Attribute{
	"auto_consolidated_monitors": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Flag for enabling auto managing DTC Consolidated Monitors in DTC Pool.",
	},
	"availability": schema.StringAttribute{
		Default: stringdefault.StaticString("ALL"),
		Validators: []validator.String{
			stringvalidator.OneOf("ALL", "ANY", "QUORUM"),
		},
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "A resource in the pool is available if ANY, at least QUORUM, or ALL monitors for the pool say that it is up.",
	},
	"comment": schema.StringAttribute{
		Optional: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
			stringvalidator.LengthBetween(0, 256),
			customvalidator.ValidateTrimmedString(),
		},
		MarkdownDescription: "The comment for the DTC Pool; maximum 256 characters.",
	},
	"consolidated_monitors": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: PoolConsolidatedMonitorsResourceSchemaAttributes,
		},
		Optional: true,
		Validators: []validator.List{
			customvalidator.ListNotEmpty(),
		},
		MarkdownDescription: "List of monitors and associated members statuses of which are shared across members and consolidated in server availability determination.",
	},
	"disable": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Determines whether the DTC Pool is disabled or not. When this is set to False, the fixed address is enabled.",
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
	"lb_alternate_method": schema.StringAttribute{
		Default: stringdefault.StaticString("NONE"),
		Validators: []validator.String{
			stringvalidator.OneOf("ALL_AVAILABLE", "DYNAMIC_RATIO", "GLOBAL_AVAILABILITY", "NONE", "RATIO", "ROUND_ROBIN", "SOURCE_IP_HASH", "TOPOLOGY"),
		},
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The alternate load balancing method. Use this to select a method type from the pool if the preferred method does not return any results.",
	},
	"lb_alternate_topology": schema.StringAttribute{
		Optional: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
		},
		MarkdownDescription: "The alternate topology for load balancing.",
	},
	"lb_dynamic_ratio_alternate": schema.SingleNestedAttribute{
		Attributes:          PoolLbDynamicRatioAlternateResourceSchemaAttributes,
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The DTC Pool settings for dynamic ratio when its selected as alternate method.",
	},
	"lb_dynamic_ratio_preferred": schema.SingleNestedAttribute{
		Attributes:          PoolLbDynamicRatioPreferredResourceSchemaAttributes,
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The DTC Pool settings for dynamic ratio when its selected as preferred method.",
	},
	"lb_preferred_method": schema.StringAttribute{
		Validators: []validator.String{
			stringvalidator.OneOf("ALL_AVAILABLE", "DYNAMIC_RATIO", "GLOBAL_AVAILABILITY", "RATIO", "ROUND_ROBIN", "SOURCE_IP_HASH", "TOPOLOGY"),
		},
		Required:            true,
		MarkdownDescription: "The preferred load balancing method. Use this to select a method type from the pool.",
	},
	"lb_preferred_topology": schema.StringAttribute{
		Optional: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
		},
		MarkdownDescription: "The preferred topology for load balancing.",
	},
	"monitors": schema.ListAttribute{
		ElementType: types.StringType,
		Optional:    true,
		CustomType:  internaltypes.UnorderedListOfStringType,
		Validators: []validator.List{
			customvalidator.ListNotEmpty(),
		},
		MarkdownDescription: "The monitors related to pool.",
	},
	"name": schema.StringAttribute{
		Required: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
			customvalidator.ValidateTrimmedString(),
		},
		MarkdownDescription: "The DTC Pool display name.",
	},
	"quorum": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "For availability mode QUORUM, at least this many monitors must report the resource as up for it to be available",
	},
	"servers": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: PoolServersResourceSchemaAttributes,
		},
		Optional: true,
		Validators: []validator.List{
			customvalidator.ListNotEmpty(),
		},
		MarkdownDescription: "The servers related to the pool.",
	},
	"ttl": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "The Time To Live (TTL) value for the DTC Pool. A 32-bit unsigned integer that represents the duration, in seconds, for which the record is valid (cached). Zero indicates that the record should not be cached.",
	},
}

var DtcPoolResourceUddiSchemaAttributes = map[string]schema.Attribute{
	"comment": schema.StringAttribute{
		Default:             stringdefault.StaticString(""),
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "Optional. Comment for __Pool__.",
	},
	"consolidated_health_enabled": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "Optional. Pool-level switch that enables consolidated health probing for this __Pool__.  Defaults to _false_ (consolidated health disabled). Set to _true_ to enable consolidated probing on this __Pool__. When _false_, any per-__PoolHealthCheck__ consolidation configuration is preserved in storage but suppressed at runtime.",
	},
	"disabled": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Optional. Flag which enables/disables __Pool__.  Defaults to _false_.",
	},
	"health_checks": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: PoolHealthCheckResourceSchemaAttributes,
		},
		Optional: true,
		Validators: []validator.List{
			customvalidator.ListNotEmpty(),
		},
		MarkdownDescription: "Optional. List of __HealthCheck__ objects IDs assigned to __Pool__.  Defaults to _empty_.",
	},
	"inheritance_sources": schema.SingleNestedAttribute{
		Attributes:          TTLInheritanceResourceSchemaAttributes,
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The inheritance configuration specifies how the object inherits the _ttl_ field.",
	},
	"method": schema.StringAttribute{
		Required:            true,
		MarkdownDescription: "Load balancing method used for selecting __Server__ assigned to __Pool__.  Valid values are: * _round_robin_ If the _round_robin_ load balancing method is selected, Universal DDI adjusts the response to a query in a sequential and circular manner, directing clients to pools.  * _ratio_ If _ratio_ load balancing method is selected, Universal DDI adjusts the response to a query so that clients are directed to pool using weighted round robin, a load-balancing pattern in which requests are distributed among several resources based on weight assigned to each resource. The distribution of responses over time will be equal for all available pools but the sequence of the responses won't be guaranteed. When equal weights are assigned for resources (pools) it effectively leads to basic round robin which directs clients to pools in sequential and circular manner.  * _global_availability_ If _global_availability_ load balancing method is selected clients are directed to the first server that is up in the _servers_ list.  Defaults to _round_robin_.",
	},
	"name": schema.StringAttribute{
		Required:            true,
		MarkdownDescription: "Display name of __Pool__.",
	},
	"pool_availability": schema.StringAttribute{
		Default:             stringdefault.StaticString("any"),
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "Optional. Pool Availability setting defines how __Pool__ health is calculated.  Valid values are: * _all_ If _all_ availability selected then __Pool__ is treated healthy when all pool's servers are healthy. * _quorum_ If _quorum_ availability selected then __Pool__ is treated healthy when at least N pool's servers are healthy. N is configurable via the value from _pool_servers_quorum_ setting. * _any_ If _any_ availability selected then __Pool__ is treated healthy when at least one pool's server is healthy.  Defaults to _any_.",
	},
	"pool_servers_quorum": schema.Int64Attribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "Pool Servers Quorum defines a minimal number of pool's healthy servers required for treating __Pool__ as healthy when Pool Availability is set to _quorum_.",
	},
	"server_availability": schema.StringAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "Optional. Server Availability setting defines how __Server__ health is calculated.  Valid values are: * _all_ If _all_ availability selected then __Server__ is treated healthy when all pool's health checks are positive. * _quorum_ If _quorum_ availability selected then __Server__ is treated healthy when at least N pool's health checks are positive. N is configurable via the value from _server_health_checks_quorum_ setting. * _any_ If _any_ availability selected then __Server__ is treated healthy when at least one pool's health check is positive  Defaults to _all_.",
	},
	"server_health_checks_quorum": schema.Int64Attribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "Server Health Checks Quorum defines a minimal number of pool's positive health checks required for treating __Server__ as healthy when Server Availability is set to _quorum_.",
	},
	"servers": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: PoolServerResourceSchemaAttributes,
		},
		Optional: true,
		Validators: []validator.List{
			customvalidator.ListNotEmpty(),
		},
		MarkdownDescription: "Optional. List of __Server__ objects assigned to __Pool__.  Defaults to _empty_.",
	},
	"tags": schema.MapAttribute{
		Optional:    true,
		Computed:    true,
		ElementType: types.StringType,
		Default:     mapdefault.StaticValue(types.MapNull(types.StringType)),
		Validators: []validator.Map{
			mapvalidator.SizeAtLeast(1),
		},
		MarkdownDescription: "Optional. The tags for __Pool__ in JSON format.",
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
}

// Expand converts the TF model to the infoblox core model
func (m *DtcPoolModel) Expand(ctx context.Context, diags *diag.Diagnostics, isCreate bool) *coremodel.DtcPool {
	if m == nil {
		return nil
	}

	obj := &coremodel.DtcPool{}

	// Expand NIOS nested attribute (returns nil if not present)
	niosModel := flex.ExpandNestedObject[NIOSDtcPoolModel](ctx, m.NIOS, diags)
	if niosModel != nil {
		obj.NIOS = niosModel.Expand(ctx, diags)
	}

	// Expand UDDI nested attribute (returns nil if not present)
	uddiModel := flex.ExpandNestedObject[UDDIDtcPoolModel](ctx, m.UDDI, diags)
	if uddiModel != nil {
		obj.UDDI = uddiModel.Expand(ctx, diags)
	}

	return obj
}

// Expand converts the NIOS TF model to the core model.
func (m *NIOSDtcPoolModel) Expand(ctx context.Context, diags *diag.Diagnostics) *coremodel.NIOSDtcPoolExt {
	return &coremodel.NIOSDtcPoolExt{
		AutoConsolidatedMonitors: flex.ExpandBoolPointer(m.AutoConsolidatedMonitors),
		Availability:             flex.ExpandStringPointerNullAsEmpty(m.Availability),
		Comment:                  flex.ExpandStringPointerNullAsEmpty(m.Comment),
		ConsolidatedMonitors:     flex.ExpandFrameworkListNestedBlock(ctx, m.ConsolidatedMonitors, diags, ExpandPoolConsolidatedMonitors),
		Disable:                  flex.ExpandBoolPointer(m.Disable),
		ExtAttrs:                 flex.ExpandMapStringAny(ctx, m.ExtAttrs, diags),
		LbAlternateMethod:        flex.ExpandStringPointerNullAsEmpty(m.LbAlternateMethod),
		LbAlternateTopology:      flex.ExpandStringPointer(m.LbAlternateTopology),
		LbDynamicRatioAlternate:  ExpandPoolLbDynamicRatioAlternate(ctx, m.LbDynamicRatioAlternate, diags),
		LbDynamicRatioPreferred:  ExpandPoolLbDynamicRatioPreferred(ctx, m.LbDynamicRatioPreferred, diags),
		LbPreferredMethod:        flex.ExpandStringPointerNullAsEmpty(m.LbPreferredMethod),
		LbPreferredTopology:      flex.ExpandStringPointer(m.LbPreferredTopology),
		Monitors:                 flex.ExpandFrameworkListString(ctx, m.Monitors, diags),
		Name:                     flex.ExpandStringPointerNullAsEmpty(m.Name),
		Quorum:                   flex.ExpandInt64Pointer(m.Quorum),
		Servers:                  flex.ExpandFrameworkListNestedBlock(ctx, m.Servers, diags, ExpandPoolServers),
		Ttl:                      flex.ExpandInt64Pointer(m.Ttl),
	}
}

// ApplyDtcPoolNIOSUseFlags derives NIOS use flags from the raw config
// value(s) and writes them onto the core model. A flag is true when the user
// set any of its governed value fields in config.
func ApplyDtcPoolNIOSUseFlags(ctx context.Context, config tfsdk.Config, obj *coremodel.DtcPool, diags *diag.Diagnostics) {
	if obj == nil || obj.NIOS == nil {
		return
	}
	obj.NIOS.UseTtl = flex.DeriveUseFlag(ctx, config, diags, path.Root("nios").AtName("ttl"))
}

// Expand converts the UDDI TF model to the core model.
func (m *UDDIDtcPoolModel) Expand(ctx context.Context, diags *diag.Diagnostics) *coremodel.UDDIDtcPoolExt {
	return &coremodel.UDDIDtcPoolExt{
		Comment:                   flex.ExpandStringPointer(m.Comment),
		ConsolidatedHealthEnabled: flex.ExpandBoolPointer(m.ConsolidatedHealthEnabled),
		Disabled:                  flex.ExpandBoolPointer(m.Disabled),
		HealthChecks:              flex.ExpandFrameworkListNestedBlock(ctx, m.HealthChecks, diags, ExpandPoolHealthCheck),
		InheritanceSources:        ExpandTTLInheritance(ctx, m.InheritanceSources, diags),
		Method:                    flex.ExpandString(m.Method),
		Name:                      flex.ExpandString(m.Name),
		PoolAvailability:          flex.ExpandStringPointer(m.PoolAvailability),
		PoolServersQuorum:         flex.ExpandInt64Pointer(m.PoolServersQuorum),
		ServerAvailability:        flex.ExpandStringPointer(m.ServerAvailability),
		ServerHealthChecksQuorum:  flex.ExpandInt64Pointer(m.ServerHealthChecksQuorum),
		Servers:                   flex.ExpandFrameworkListNestedBlock(ctx, m.Servers, diags, ExpandPoolServer),
		Tags:                      flex.ExpandMapStringAny(ctx, m.Tags, diags),
		Ttl:                       flex.ExpandInt64Pointer(m.Ttl),
	}
}

// Flatten populates the TF model from a core response.
func (m *DtcPoolModel) Flatten(ctx context.Context, resp *coremodel.DtcPool, diags *diag.Diagnostics) {
	if resp == nil {
		return
	}

	m.Id = flex.FlattenStringPointer(resp.Id)

	// Extract existing NIOS model, flatten API response onto it, convert back
	niosModel := flex.ExpandNestedObject[NIOSDtcPoolModel](ctx, m.NIOS, diags)
	if niosModel == nil {
		niosModel = &NIOSDtcPoolModel{}
	}
	niosModel.Flatten(ctx, resp.NIOS, diags)
	if resp.NIOS != nil {
		m.NIOS = flex.FlattenNestedObject(ctx, niosModel, NIOSDtcPoolAttrTypes, diags)
	} else {
		m.NIOS = types.ObjectNull(NIOSDtcPoolAttrTypes)
	}

	// Extract existing UDDI model, flatten API response onto it, convert back
	uddiModel := flex.ExpandNestedObject[UDDIDtcPoolModel](ctx, m.UDDI, diags)
	if uddiModel == nil {
		uddiModel = &UDDIDtcPoolModel{}
	}
	uddiModel.Flatten(ctx, resp.UDDI, diags)
	if resp.UDDI != nil {
		m.UDDI = flex.FlattenNestedObject(ctx, uddiModel, UDDIDtcPoolAttrTypes, diags)
	} else {
		m.UDDI = types.ObjectNull(UDDIDtcPoolAttrTypes)
	}
}

// Flatten merges API response onto existing NIOS model.
func (m *NIOSDtcPoolModel) Flatten(ctx context.Context, from *coremodel.NIOSDtcPoolExt, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	planExtAttrs := m.ExtAttrs
	if planExtAttrs.IsUnknown() {
		planExtAttrs = types.MapNull(types.StringType)
	}
	m.AutoConsolidatedMonitors = flex.FlattenBoolPointer(from.AutoConsolidatedMonitors)
	m.Availability = flex.FlattenStringPointerEmptyAsNull(from.Availability)
	m.Comment = flex.FlattenStringPointerEmptyAsNull(from.Comment)
	m.ConsolidatedMonitors = flex.FlattenFrameworkListNestedBlock(ctx, from.ConsolidatedMonitors, PoolConsolidatedMonitorsAttrTypes, diags, FlattenPoolConsolidatedMonitors)
	m.Disable = flex.FlattenBoolPointer(from.Disable)
	m.ExtAttrs, m.ExtAttrsAll = flex.FlattenEAs(planExtAttrs, from.ExtAttrs)
	m.LbAlternateMethod = flex.FlattenStringPointerEmptyAsNull(from.LbAlternateMethod)
	m.LbAlternateTopology = flex.FlattenStringPointerEmptyAsNull(from.LbAlternateTopology)
	m.LbDynamicRatioAlternate = FlattenPoolLbDynamicRatioAlternate(ctx, from.LbDynamicRatioAlternate, diags)
	m.LbDynamicRatioPreferred = FlattenPoolLbDynamicRatioPreferred(ctx, from.LbDynamicRatioPreferred, diags)
	m.LbPreferredMethod = flex.FlattenStringPointerEmptyAsNull(from.LbPreferredMethod)
	m.LbPreferredTopology = flex.FlattenStringPointerEmptyAsNull(from.LbPreferredTopology)
	m.Monitors = flex.FlattenFrameworkUnorderedListString(ctx, from.Monitors, diags)
	m.Name = flex.FlattenStringPointerEmptyAsNull(from.Name)
	m.Quorum = flex.FlattenInt64Pointer(from.Quorum)
	m.Servers = flex.FlattenFrameworkListNestedBlock(ctx, from.Servers, PoolServersAttrTypes, diags, FlattenPoolServers)
	m.Ttl = flex.FlattenInt64Pointer(from.Ttl)
}

// Flatten merges API response onto existing UDDI model.
func (m *UDDIDtcPoolModel) Flatten(ctx context.Context, from *coremodel.UDDIDtcPoolExt, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	m.Comment = flex.FlattenStringPointer(from.Comment)
	m.ConsolidatedHealthEnabled = flex.FlattenBoolPointer(from.ConsolidatedHealthEnabled)
	m.Disabled = flex.FlattenBoolPointer(from.Disabled)
	m.HealthChecks = flex.FlattenFrameworkListNestedBlock(ctx, from.HealthChecks, PoolHealthCheckAttrTypes, diags, FlattenPoolHealthCheck)
	m.InheritanceSources = FlattenTTLInheritance(ctx, from.InheritanceSources, diags)
	m.Method = flex.FlattenString(from.Method)
	m.Name = flex.FlattenString(from.Name)
	m.PoolAvailability = flex.FlattenStringPointer(from.PoolAvailability)
	m.PoolServersQuorum = flex.FlattenInt64Pointer(from.PoolServersQuorum)
	m.ServerAvailability = flex.FlattenStringPointer(from.ServerAvailability)
	m.ServerHealthChecksQuorum = flex.FlattenInt64Pointer(from.ServerHealthChecksQuorum)
	m.Servers = flex.FlattenFrameworkListNestedBlock(ctx, from.Servers, PoolServerAttrTypes, diags, FlattenPoolServer)
	tagsAll := flex.FlattenMapStringAny(ctx, from.Tags, diags)
	if m.Tags.IsNull() || m.Tags.IsUnknown() {
		m.Tags = tagsAll
	}
	m.TagsAll = tagsAll
	m.Ttl = flex.FlattenInt64Pointer(from.Ttl)
}
