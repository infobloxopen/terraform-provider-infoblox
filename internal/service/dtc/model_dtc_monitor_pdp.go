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

type DtcMonitorPdpModel struct {
	Id   types.String `tfsdk:"id"`
	NIOS types.Object `tfsdk:"nios"`
	UDDI types.Object `tfsdk:"uddi"`
}

var DtcMonitorPdpAttrTypes = map[string]attr.Type{
	"id":   types.StringType,
	"nios": types.ObjectType{AttrTypes: NIOSDtcMonitorPdpAttrTypes},
	"uddi": types.ObjectType{AttrTypes: UDDIDtcMonitorPdpAttrTypes},
}

type NIOSDtcMonitorPdpModel struct {
	Comment     types.String `tfsdk:"comment"`
	ExtAttrs    types.Map    `tfsdk:"ext_attrs"`
	ExtAttrsAll types.Map    `tfsdk:"ext_attrs_all"`
	Interval    types.Int64  `tfsdk:"interval"`
	Name        types.String `tfsdk:"name"`
	Port        types.Int64  `tfsdk:"port"`
	RetryDown   types.Int64  `tfsdk:"retry_down"`
	RetryUp     types.Int64  `tfsdk:"retry_up"`
	Timeout     types.Int64  `tfsdk:"timeout"`
}

var NIOSDtcMonitorPdpAttrTypes = map[string]attr.Type{
	"comment":       types.StringType,
	"ext_attrs":     types.MapType{ElemType: types.StringType},
	"ext_attrs_all": types.MapType{ElemType: types.StringType},
	"interval":      types.Int64Type,
	"name":          types.StringType,
	"port":          types.Int64Type,
	"retry_down":    types.Int64Type,
	"retry_up":      types.Int64Type,
	"timeout":       types.Int64Type,
}

type UDDIDtcMonitorPdpModel struct {
	Comment   types.String `tfsdk:"comment"`
	Disabled  types.Bool   `tfsdk:"disabled"`
	Interval  types.Int64  `tfsdk:"interval"`
	Name      types.String `tfsdk:"name"`
	Port      types.Int64  `tfsdk:"port"`
	RetryDown types.Int64  `tfsdk:"retry_down"`
	RetryUp   types.Int64  `tfsdk:"retry_up"`
	Tags      types.Map    `tfsdk:"tags"`
	TagsAll   types.Map    `tfsdk:"tags_all"`
	Timeout   types.Int64  `tfsdk:"timeout"`
}

var UDDIDtcMonitorPdpAttrTypes = map[string]attr.Type{
	"comment":    types.StringType,
	"disabled":   types.BoolType,
	"interval":   types.Int64Type,
	"name":       types.StringType,
	"port":       types.Int64Type,
	"retry_down": types.Int64Type,
	"retry_up":   types.Int64Type,
	"tags":       types.MapType{ElemType: types.StringType},
	"tags_all":   types.MapType{ElemType: types.StringType},
	"timeout":    types.Int64Type,
}

const (
	DtcMonitorPdpReturnFields = "comment,extattrs,interval,name,port,retry_down,retry_up,timeout"
)

var DtcMonitorPdpResourceSchemaAttributes = map[string]schema.Attribute{
	"id": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The reference to the object.",
	},
	"nios": schema.SingleNestedAttribute{
		Optional:            true,
		MarkdownDescription: "NIOS backend-specific fields.",
		Attributes:          DtcMonitorPdpResourceNiosSchemaAttributes,
	},
	"uddi": schema.SingleNestedAttribute{
		Optional:            true,
		MarkdownDescription: "UDDI backend-specific fields.",
		Attributes:          DtcMonitorPdpResourceUddiSchemaAttributes,
	},
}

var DtcMonitorPdpResourceNiosSchemaAttributes = map[string]schema.Attribute{
	"comment": schema.StringAttribute{
		Optional: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
			stringvalidator.LengthBetween(0, 256),
			customvalidator.ValidateTrimmedString(),
		},
		MarkdownDescription: "Comment for this DTC monitor; maximum 256 characters.",
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
	"port": schema.Int64Attribute{
		Optional: true,
		Computed: true,
		Default:  int64default.StaticInt64(2123),
		Validators: []validator.Int64{
			int64validator.Between(1, 65535),
		},
		MarkdownDescription: "The port value for PDP requests.",
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
}

var DtcMonitorPdpResourceUddiSchemaAttributes = map[string]schema.Attribute{
	"comment": schema.StringAttribute{
		Default:             stringdefault.StaticString(""),
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "Optional. Comment for __PDPHealthCheck__.",
	},
	"disabled": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Optional. Flag which enables/disables __PDPHealthCheck__. Defaults to _false_.",
	},
	"interval": schema.Int64Attribute{
		Optional:            true,
		Computed:            true,
		Default:             int64default.StaticInt64(15),
		MarkdownDescription: "Optional. Interval value in seconds. The health check runs only for the specified interval and it is measured from the beginning of the previous check cycle. Defaults to _15_.",
	},
	"name": schema.StringAttribute{
		Required:            true,
		MarkdownDescription: "Display name of __PDPHealthCheck__.",
	},
	"port": schema.Int64Attribute{
		Optional:            true,
		Computed:            true,
		Default:             int64default.StaticInt64(2123),
		MarkdownDescription: "Optional. Destination UDP port of __PDPHealthCheck__ for the GTP Echo Request. Defaults to _2123_.",
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
		MarkdownDescription: "Optional. The tags for __PDPHealthCheck__ in JSON format.",
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
}

// Expand converts the TF model to the infoblox core model
func (m *DtcMonitorPdpModel) Expand(ctx context.Context, diags *diag.Diagnostics, isCreate bool) *coremodel.DtcMonitorPdp {
	if m == nil {
		return nil
	}

	obj := &coremodel.DtcMonitorPdp{}

	// Expand NIOS nested attribute (returns nil if not present)
	niosModel := flex.ExpandNestedObject[NIOSDtcMonitorPdpModel](ctx, m.NIOS, diags)
	if niosModel != nil {
		obj.NIOS = niosModel.Expand(ctx, diags)
	}

	// Expand UDDI nested attribute (returns nil if not present)
	uddiModel := flex.ExpandNestedObject[UDDIDtcMonitorPdpModel](ctx, m.UDDI, diags)
	if uddiModel != nil {
		obj.UDDI = uddiModel.Expand(ctx, diags)
	}

	return obj
}

// Expand converts the NIOS TF model to the core model.
func (m *NIOSDtcMonitorPdpModel) Expand(ctx context.Context, diags *diag.Diagnostics) *coremodel.NIOSDtcMonitorPdpExt {
	return &coremodel.NIOSDtcMonitorPdpExt{
		Comment:   flex.ExpandStringPointerNullAsEmpty(m.Comment),
		ExtAttrs:  flex.ExpandMapStringAny(ctx, m.ExtAttrs, diags),
		Interval:  flex.ExpandInt64Pointer(m.Interval),
		Name:      flex.ExpandStringPointerNullAsEmpty(m.Name),
		Port:      flex.ExpandInt64Pointer(m.Port),
		RetryDown: flex.ExpandInt64Pointer(m.RetryDown),
		RetryUp:   flex.ExpandInt64Pointer(m.RetryUp),
		Timeout:   flex.ExpandInt64Pointer(m.Timeout),
	}
}

// Expand converts the UDDI TF model to the core model.
func (m *UDDIDtcMonitorPdpModel) Expand(ctx context.Context, diags *diag.Diagnostics) *coremodel.UDDIDtcMonitorPdpExt {
	return &coremodel.UDDIDtcMonitorPdpExt{
		Comment:   flex.ExpandStringPointer(m.Comment),
		Disabled:  flex.ExpandBoolPointer(m.Disabled),
		Interval:  flex.ExpandInt64Pointer(m.Interval),
		Name:      flex.ExpandString(m.Name),
		Port:      flex.ExpandInt64Pointer(m.Port),
		RetryDown: flex.ExpandInt64Pointer(m.RetryDown),
		RetryUp:   flex.ExpandInt64Pointer(m.RetryUp),
		Tags:      flex.ExpandMapStringAny(ctx, m.Tags, diags),
		Timeout:   flex.ExpandInt64Pointer(m.Timeout),
	}
}

// Flatten populates the TF model from a core response.
func (m *DtcMonitorPdpModel) Flatten(ctx context.Context, resp *coremodel.DtcMonitorPdp, diags *diag.Diagnostics) {
	if resp == nil {
		return
	}

	m.Id = flex.FlattenStringPointer(resp.Id)

	// Extract existing NIOS model, flatten API response onto it, convert back
	niosModel := flex.ExpandNestedObject[NIOSDtcMonitorPdpModel](ctx, m.NIOS, diags)
	if niosModel == nil {
		niosModel = &NIOSDtcMonitorPdpModel{}
	}
	niosModel.Flatten(ctx, resp.NIOS, diags)
	if resp.NIOS != nil {
		m.NIOS = flex.FlattenNestedObject(ctx, niosModel, NIOSDtcMonitorPdpAttrTypes, diags)
	} else {
		m.NIOS = types.ObjectNull(NIOSDtcMonitorPdpAttrTypes)
	}

	// Extract existing UDDI model, flatten API response onto it, convert back
	uddiModel := flex.ExpandNestedObject[UDDIDtcMonitorPdpModel](ctx, m.UDDI, diags)
	if uddiModel == nil {
		uddiModel = &UDDIDtcMonitorPdpModel{}
	}
	uddiModel.Flatten(ctx, resp.UDDI, diags)
	if resp.UDDI != nil {
		m.UDDI = flex.FlattenNestedObject(ctx, uddiModel, UDDIDtcMonitorPdpAttrTypes, diags)
	} else {
		m.UDDI = types.ObjectNull(UDDIDtcMonitorPdpAttrTypes)
	}
}

// Flatten merges API response onto existing NIOS model.
func (m *NIOSDtcMonitorPdpModel) Flatten(ctx context.Context, from *coremodel.NIOSDtcMonitorPdpExt, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	planExtAttrs := m.ExtAttrs
	if planExtAttrs.IsUnknown() {
		planExtAttrs = types.MapNull(types.StringType)
	}
	m.Comment = flex.FlattenStringPointerEmptyAsNull(from.Comment)
	m.ExtAttrs, m.ExtAttrsAll = flex.FlattenEAs(planExtAttrs, from.ExtAttrs)
	m.Interval = flex.FlattenInt64Pointer(from.Interval)
	m.Name = flex.FlattenStringPointerEmptyAsNull(from.Name)
	m.Port = flex.FlattenInt64Pointer(from.Port)
	m.RetryDown = flex.FlattenInt64Pointer(from.RetryDown)
	m.RetryUp = flex.FlattenInt64Pointer(from.RetryUp)
	m.Timeout = flex.FlattenInt64Pointer(from.Timeout)
}

// Flatten merges API response onto existing UDDI model.
func (m *UDDIDtcMonitorPdpModel) Flatten(ctx context.Context, from *coremodel.UDDIDtcMonitorPdpExt, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	m.Comment = flex.FlattenStringPointer(from.Comment)
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
}
