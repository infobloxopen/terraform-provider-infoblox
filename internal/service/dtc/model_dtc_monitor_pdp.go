package dtc

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/mapvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/mapdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
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
}

var DtcMonitorPdpAttrTypes = map[string]attr.Type{
	"id":   types.StringType,
	"nios": types.ObjectType{AttrTypes: NIOSDtcMonitorPdpAttrTypes},
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
