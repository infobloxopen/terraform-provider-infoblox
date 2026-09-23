package dtc

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/mapvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/mapdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"

	coremodel "github.com/infobloxopen/terraform-provider-infoblox/internal/core/model/dtc"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
	importmod "github.com/infobloxopen/terraform-provider-infoblox/internal/planmodifiers/import"
	customvalidator "github.com/infobloxopen/terraform-provider-infoblox/internal/validator"
)

type DtcMonitorSnmpModel struct {
	Id            types.String `tfsdk:"id"`
	UpdateTrigger types.String `tfsdk:"update_trigger"`
	NIOS          types.Object `tfsdk:"nios"`
	UDDI          types.Object `tfsdk:"uddi"`
}

var DtcMonitorSnmpAttrTypes = map[string]attr.Type{
	"id":             types.StringType,
	"update_trigger": types.StringType,
	"nios":           types.ObjectType{AttrTypes: NIOSDtcMonitorSnmpAttrTypes},
	"uddi":           types.ObjectType{AttrTypes: UDDIDtcMonitorSnmpAttrTypes},
}

type NIOSDtcMonitorSnmpModel struct {
	Comment     types.String `tfsdk:"comment"`
	Community   types.String `tfsdk:"community"`
	Context     types.String `tfsdk:"context"`
	EngineId    types.String `tfsdk:"engine_id"`
	ExtAttrs    types.Map    `tfsdk:"ext_attrs"`
	ExtAttrsAll types.Map    `tfsdk:"ext_attrs_all"`
	Interval    types.Int64  `tfsdk:"interval"`
	Name        types.String `tfsdk:"name"`
	Oids        types.List   `tfsdk:"oids"`
	Port        types.Int64  `tfsdk:"port"`
	RetryDown   types.Int64  `tfsdk:"retry_down"`
	RetryUp     types.Int64  `tfsdk:"retry_up"`
	Timeout     types.Int64  `tfsdk:"timeout"`
	User        types.String `tfsdk:"user"`
	Version     types.String `tfsdk:"version"`
}

var NIOSDtcMonitorSnmpAttrTypes = map[string]attr.Type{
	"comment":       types.StringType,
	"community":     types.StringType,
	"context":       types.StringType,
	"engine_id":     types.StringType,
	"ext_attrs":     types.MapType{ElemType: types.StringType},
	"ext_attrs_all": types.MapType{ElemType: types.StringType},
	"interval":      types.Int64Type,
	"name":          types.StringType,
	"oids":          types.ListType{ElemType: types.ObjectType{AttrTypes: MonitorSnmpOidsAttrTypes}},
	"port":          types.Int64Type,
	"retry_down":    types.Int64Type,
	"retry_up":      types.Int64Type,
	"timeout":       types.Int64Type,
	"user":          types.StringType,
	"version":       types.StringType,
}

type UDDIDtcMonitorSnmpModel struct {
	CheckList         types.List   `tfsdk:"check_list"`
	Comment           types.String `tfsdk:"comment"`
	Community         types.String `tfsdk:"community"`
	ContextEngineId   types.String `tfsdk:"context_engine_id"`
	ContextName       types.String `tfsdk:"context_name"`
	Disabled          types.Bool   `tfsdk:"disabled"`
	Interval          types.Int64  `tfsdk:"interval"`
	Name              types.String `tfsdk:"name"`
	Port              types.Int64  `tfsdk:"port"`
	RetryDown         types.Int64  `tfsdk:"retry_down"`
	RetryUp           types.Int64  `tfsdk:"retry_up"`
	Tags              types.Map    `tfsdk:"tags"`
	TagsAll           types.Map    `tfsdk:"tags_all"`
	Timeout           types.Int64  `tfsdk:"timeout"`
	UserSecurityModel types.String `tfsdk:"user_security_model"`
	Version           types.String `tfsdk:"version"`
}

var UDDIDtcMonitorSnmpAttrTypes = map[string]attr.Type{
	"check_list":          types.ListType{ElemType: types.ObjectType{AttrTypes: SNMPHealthCheckEntryCheckAttrTypes}},
	"comment":             types.StringType,
	"community":           types.StringType,
	"context_engine_id":   types.StringType,
	"context_name":        types.StringType,
	"disabled":            types.BoolType,
	"interval":            types.Int64Type,
	"name":                types.StringType,
	"port":                types.Int64Type,
	"retry_down":          types.Int64Type,
	"retry_up":            types.Int64Type,
	"tags":                types.MapType{ElemType: types.StringType},
	"tags_all":            types.MapType{ElemType: types.StringType},
	"timeout":             types.Int64Type,
	"user_security_model": types.StringType,
	"version":             types.StringType,
}

const (
	DtcMonitorSnmpReturnFields = "comment,community,context,engine_id,extattrs,interval,name,oids,port,retry_down,retry_up,timeout,user,version"
)

var DtcMonitorSnmpResourceSchemaAttributes = map[string]schema.Attribute{
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
		Attributes:          DtcMonitorSnmpResourceNiosSchemaAttributes,
	},
	"uddi": schema.SingleNestedAttribute{
		Optional:            true,
		MarkdownDescription: "UDDI backend-specific fields.",
		Attributes:          DtcMonitorSnmpResourceUddiSchemaAttributes,
	},
}

var DtcMonitorSnmpResourceNiosSchemaAttributes = map[string]schema.Attribute{
	"comment": schema.StringAttribute{
		Optional: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
			stringvalidator.LengthBetween(0, 256),
			customvalidator.ValidateTrimmedString(),
		},
		MarkdownDescription: "Comment for this DTC monitor; maximum 256 characters.",
	},
	"community": schema.StringAttribute{
		Default:  stringdefault.StaticString("public"),
		Optional: true,
		Computed: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
		},
		MarkdownDescription: "The SNMP community string for SNMP authentication.",
	},
	"context": schema.StringAttribute{
		Optional: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
			customvalidator.ValidateTrimmedString(),
		},
		MarkdownDescription: "The SNMPv3 context.",
	},
	"engine_id": schema.StringAttribute{
		Optional: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
			customvalidator.ValidateTrimmedString(),
		},
		MarkdownDescription: "The SNMPv3 engine identifier.",
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
	"interval": schema.Int64Attribute{
		Optional:            true,
		Computed:            true,
		Default:             int64default.StaticInt64(5),
		MarkdownDescription: "The interval for TCP health check.",
	},
	"name": schema.StringAttribute{
		Required: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
			customvalidator.ValidateTrimmedString(),
		},
		MarkdownDescription: "The display name for this DTC monitor.",
	},
	"oids": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: MonitorSnmpOidsResourceSchemaAttributes,
		},
		Optional: true,
		Validators: []validator.List{
			customvalidator.ListNotEmpty(),
		},
		MarkdownDescription: "A list of OIDs for SNMP monitoring.",
	},
	"port": schema.Int64Attribute{
		Optional: true,
		Computed: true,
		Default:  int64default.StaticInt64(161),
		Validators: []validator.Int64{
			int64validator.Between(1, 65535),
		},
		MarkdownDescription: "The port value for SNMP requests.",
	},
	"retry_down": schema.Int64Attribute{
		Optional:            true,
		Computed:            true,
		Default:             int64default.StaticInt64(1),
		MarkdownDescription: "The value of how many times the server should appear as down to be treated as dead after it was alive.",
	},
	"retry_up": schema.Int64Attribute{
		Optional:            true,
		Computed:            true,
		Default:             int64default.StaticInt64(1),
		MarkdownDescription: "The value of how many times the server should appear as up to be treated as alive after it was dead.",
	},
	"timeout": schema.Int64Attribute{
		Optional:            true,
		Computed:            true,
		Default:             int64default.StaticInt64(15),
		MarkdownDescription: "The timeout for TCP health check in seconds.",
	},
	"user": schema.StringAttribute{
		Optional: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
		},
		MarkdownDescription: "The SNMPv3 user setting.",
	},
	"version": schema.StringAttribute{
		Default: stringdefault.StaticString("V2C"),
		Validators: []validator.String{
			stringvalidator.OneOf("V1", "V2C", "V3"),
		},
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The SNMP protocol version for the SNMP health check.",
	},
}

var DtcMonitorSnmpResourceUddiSchemaAttributes = map[string]schema.Attribute{
	"check_list": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: SNMPHealthCheckEntryCheckResourceSchemaAttributes,
		},
		Optional: true,
		Validators: []validator.List{
			customvalidator.ListNotEmpty(),
		},
		MarkdownDescription: "List of specific checks for SNMP entries and their values in MIB hierarchy. Supported up to 15 checks.",
	},
	"comment": schema.StringAttribute{
		Default:             stringdefault.StaticString(""),
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "Optional. Comment for __SNMPHealthCheck__.",
	},
	"community": schema.StringAttribute{
		Default:             stringdefault.StaticString("public"),
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "Optional. SNMP community string used for authentication. Mandatory for __v1__ and __v2c__ versions, ignored for __v3__.  Defaults to __public__.",
	},
	"context_engine_id": schema.StringAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "Optional. Uniquely identifies an SNMP entity that may realize an instance of a context with a particular context name.  Format is an arbitrary string that can contain from 10 to 64 hexadecimal digits (5 to 32 octet numbers).  Ignored for __v1__ and __v2c__ versions.",
	},
	"context_name": schema.StringAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "Optional. Name of administratively unique context for __v3__ version. Ignored for __v1__ and __v2c__ versions.",
	},
	"disabled": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Optional. Flag which enables/disables __SNMPHealthCheck__. Defaults to _false_.",
	},
	"interval": schema.Int64Attribute{
		Optional:            true,
		Computed:            true,
		Default:             int64default.StaticInt64(15),
		MarkdownDescription: "Optional. Interval value in seconds. The health check runs only for the specified interval and it is measured from the beginning of the previous check cycle. Defaults to _15_.",
	},
	"name": schema.StringAttribute{
		Required:            true,
		MarkdownDescription: "Display name of __SNMPHealthCheck__.",
	},
	"port": schema.Int64Attribute{
		Optional:            true,
		Computed:            true,
		Default:             int64default.StaticInt64(161),
		MarkdownDescription: "Optional. Destination UDP port of __SNMPHealthCheck__. Defaults to _161_.",
	},
	"retry_down": schema.Int64Attribute{
		Optional:            true,
		Computed:            true,
		Default:             int64default.StaticInt64(1),
		MarkdownDescription: "Optional. Retry down count. The value determines how many bad health checks in a row must be received by the onprem host from the DTC Server for treating the health check as failed. Defaults to _1_.",
	},
	"retry_up": schema.Int64Attribute{
		Optional:            true,
		Computed:            true,
		Default:             int64default.StaticInt64(1),
		MarkdownDescription: "Optional. Retry up count. The value determines how many good health checks in a row must be received by the onprem host from the DTC Server for treating the health check as successful. Defaults to _1_.",
	},
	"tags": schema.MapAttribute{
		Optional:    true,
		Computed:    true,
		ElementType: types.StringType,
		Default:     mapdefault.StaticValue(types.MapNull(types.StringType)),
		Validators: []validator.Map{
			mapvalidator.SizeAtLeast(1),
		},
		MarkdownDescription: "Optional. The tags for __SNMPHealthCheck__ in JSON format.",
	},
	"tags_all": schema.MapAttribute{
		Computed:            true,
		ElementType:         types.StringType,
		MarkdownDescription: "All tags including inherited values.",
	},
	"timeout": schema.Int64Attribute{
		Optional:            true,
		Computed:            true,
		Default:             int64default.StaticInt64(10),
		MarkdownDescription: "Optional. Timeout value in seconds. The health check waits for the specified number of seconds after sending a request. If it does not receive a response within the number of seconds, then the health check is considered as failed. Defaults to _10_.",
	},
	"user_security_model": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The resource identifier.",
	},
	"version": schema.StringAttribute{
		Validators: []validator.String{
			stringvalidator.OneOf("v1", "v2c", "v3"),
		},
		Required:            true,
		MarkdownDescription: "SNMP version.  Allowed values: * v1  - version 1 * v2c - version 2 community * v3  - version 3",
	},
}

// Expand converts the TF model to the infoblox core model
func (m *DtcMonitorSnmpModel) Expand(ctx context.Context, diags *diag.Diagnostics, isCreate bool) *coremodel.DtcMonitorSnmp {
	if m == nil {
		return nil
	}

	obj := &coremodel.DtcMonitorSnmp{}

	// Expand NIOS nested attribute (returns nil if not present)
	niosModel := flex.ExpandNestedObject[NIOSDtcMonitorSnmpModel](ctx, m.NIOS, diags)
	if niosModel != nil {
		obj.NIOS = niosModel.Expand(ctx, diags)
	}

	// Expand UDDI nested attribute (returns nil if not present)
	uddiModel := flex.ExpandNestedObject[UDDIDtcMonitorSnmpModel](ctx, m.UDDI, diags)
	if uddiModel != nil {
		obj.UDDI = uddiModel.Expand(ctx, diags)
	}

	return obj
}

// Expand converts the NIOS TF model to the core model.
func (m *NIOSDtcMonitorSnmpModel) Expand(ctx context.Context, diags *diag.Diagnostics) *coremodel.NIOSDtcMonitorSnmpExt {
	return &coremodel.NIOSDtcMonitorSnmpExt{
		Comment:   flex.ExpandStringPointerNullAsEmpty(m.Comment),
		Community: flex.ExpandStringPointerNullAsEmpty(m.Community),
		Context:   flex.ExpandStringPointerNullAsEmpty(m.Context),
		EngineId:  flex.ExpandStringPointerNullAsEmpty(m.EngineId),
		ExtAttrs:  flex.ExpandMapStringAny(ctx, m.ExtAttrs, diags),
		Interval:  flex.ExpandInt64Pointer(m.Interval),
		Name:      flex.ExpandStringPointerNullAsEmpty(m.Name),
		Oids:      flex.ExpandFrameworkListNestedBlock(ctx, m.Oids, diags, ExpandMonitorSnmpOids),
		Port:      flex.ExpandInt64Pointer(m.Port),
		RetryDown: flex.ExpandInt64Pointer(m.RetryDown),
		RetryUp:   flex.ExpandInt64Pointer(m.RetryUp),
		Timeout:   flex.ExpandInt64Pointer(m.Timeout),
		User:      flex.ExpandStringPointer(m.User),
		Version:   flex.ExpandStringPointerNullAsEmpty(m.Version),
	}
}

// Expand converts the UDDI TF model to the core model.
func (m *UDDIDtcMonitorSnmpModel) Expand(ctx context.Context, diags *diag.Diagnostics) *coremodel.UDDIDtcMonitorSnmpExt {
	return &coremodel.UDDIDtcMonitorSnmpExt{
		CheckList:         flex.ExpandFrameworkListNestedBlock(ctx, m.CheckList, diags, ExpandSNMPHealthCheckEntryCheck),
		Comment:           flex.ExpandStringPointer(m.Comment),
		Community:         flex.ExpandStringPointer(m.Community),
		ContextEngineId:   flex.ExpandStringPointer(m.ContextEngineId),
		ContextName:       flex.ExpandStringPointer(m.ContextName),
		Disabled:          flex.ExpandBoolPointer(m.Disabled),
		Interval:          flex.ExpandInt64Pointer(m.Interval),
		Name:              flex.ExpandString(m.Name),
		Port:              flex.ExpandInt64Pointer(m.Port),
		RetryDown:         flex.ExpandInt64Pointer(m.RetryDown),
		RetryUp:           flex.ExpandInt64Pointer(m.RetryUp),
		Tags:              flex.ExpandMapStringAny(ctx, m.Tags, diags),
		Timeout:           flex.ExpandInt64Pointer(m.Timeout),
		UserSecurityModel: flex.ExpandStringPointer(m.UserSecurityModel),
		Version:           flex.ExpandString(m.Version),
	}
}

// Flatten populates the TF model from a core response.
func (m *DtcMonitorSnmpModel) Flatten(ctx context.Context, resp *coremodel.DtcMonitorSnmp, diags *diag.Diagnostics) {
	if resp == nil {
		return
	}

	m.Id = flex.FlattenStringPointer(resp.Id)

	// Extract existing NIOS model, flatten API response onto it, convert back
	niosModel := flex.ExpandNestedObject[NIOSDtcMonitorSnmpModel](ctx, m.NIOS, diags)
	if niosModel == nil {
		niosModel = &NIOSDtcMonitorSnmpModel{}
	}
	plannedNIOS := flex.ExpandNestedObject[NIOSDtcMonitorSnmpModel](ctx, m.NIOS, diags)
	niosModel.Flatten(ctx, resp.NIOS, diags)
	if resp.NIOS != nil {
		PostFlattenDtcMonitorSnmpNIOS(ctx, plannedNIOS, niosModel, diags)
		m.NIOS = flex.FlattenNestedObject(ctx, niosModel, NIOSDtcMonitorSnmpAttrTypes, diags)
	} else {
		m.NIOS = types.ObjectNull(NIOSDtcMonitorSnmpAttrTypes)
	}

	// Extract existing UDDI model, flatten API response onto it, convert back
	uddiModel := flex.ExpandNestedObject[UDDIDtcMonitorSnmpModel](ctx, m.UDDI, diags)
	if uddiModel == nil {
		uddiModel = &UDDIDtcMonitorSnmpModel{}
	}
	uddiModel.Flatten(ctx, resp.UDDI, diags)
	if resp.UDDI != nil {
		m.UDDI = flex.FlattenNestedObject(ctx, uddiModel, UDDIDtcMonitorSnmpAttrTypes, diags)
	} else {
		m.UDDI = types.ObjectNull(UDDIDtcMonitorSnmpAttrTypes)
	}
}

// Flatten merges API response onto existing NIOS model.
func (m *NIOSDtcMonitorSnmpModel) Flatten(ctx context.Context, from *coremodel.NIOSDtcMonitorSnmpExt, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	planExtAttrs := m.ExtAttrs
	if planExtAttrs.IsUnknown() {
		planExtAttrs = types.MapNull(types.StringType)
	}
	m.Comment = flex.FlattenStringPointerEmptyAsNull(from.Comment)
	m.Community = flex.FlattenStringPointerEmptyAsNull(from.Community)
	m.Context = flex.FlattenStringPointerEmptyAsNull(from.Context)
	m.EngineId = flex.FlattenStringPointerEmptyAsNull(from.EngineId)
	m.ExtAttrs, m.ExtAttrsAll = flex.FlattenEAs(planExtAttrs, from.ExtAttrs)
	m.Interval = flex.FlattenInt64Pointer(from.Interval)
	m.Name = flex.FlattenStringPointerEmptyAsNull(from.Name)
	m.Oids = flex.FlattenFrameworkListNestedBlock(ctx, from.Oids, MonitorSnmpOidsAttrTypes, diags, FlattenMonitorSnmpOids)
	m.Port = flex.FlattenInt64Pointer(from.Port)
	m.RetryDown = flex.FlattenInt64Pointer(from.RetryDown)
	m.RetryUp = flex.FlattenInt64Pointer(from.RetryUp)
	m.Timeout = flex.FlattenInt64Pointer(from.Timeout)
	m.User = flex.FlattenStringPointerEmptyAsNull(from.User)
	m.Version = flex.FlattenStringPointerEmptyAsNull(from.Version)
}

// Flatten merges API response onto existing UDDI model.
func (m *UDDIDtcMonitorSnmpModel) Flatten(ctx context.Context, from *coremodel.UDDIDtcMonitorSnmpExt, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	m.CheckList = flex.FlattenFrameworkListNestedBlock(ctx, from.CheckList, SNMPHealthCheckEntryCheckAttrTypes, diags, FlattenSNMPHealthCheckEntryCheck)
	m.Comment = flex.FlattenStringPointer(from.Comment)
	m.Community = flex.FlattenStringPointer(from.Community)
	m.ContextEngineId = flex.FlattenStringPointer(from.ContextEngineId)
	m.ContextName = flex.FlattenStringPointer(from.ContextName)
	m.Disabled = flex.FlattenBoolPointer(from.Disabled)
	m.Interval = flex.FlattenInt64Pointer(from.Interval)
	m.Name = flex.FlattenString(from.Name)
	m.Port = flex.FlattenInt64Pointer(from.Port)
	m.RetryDown = flex.FlattenInt64Pointer(from.RetryDown)
	m.RetryUp = flex.FlattenInt64Pointer(from.RetryUp)
	tagsAll := flex.FlattenMapStringAny(ctx, from.Tags, diags)
	if m.Tags.IsNull() || m.Tags.IsUnknown() {
		m.Tags = tagsAll
	}
	m.TagsAll = tagsAll
	m.Timeout = flex.FlattenInt64Pointer(from.Timeout)
	m.UserSecurityModel = flex.FlattenStringPointer(from.UserSecurityModel)
	m.Version = flex.FlattenString(from.Version)
}
