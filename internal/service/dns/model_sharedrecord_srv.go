package dns

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/mapvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/mapdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"

	coremodel "github.com/infobloxopen/terraform-provider-infoblox/internal/core/model/dns"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
	immutable "github.com/infobloxopen/terraform-provider-infoblox/internal/planmodifiers/immutable"
	customvalidator "github.com/infobloxopen/terraform-provider-infoblox/internal/validator"
)

type SharedrecordSrvModel struct {
	Id            types.String `tfsdk:"id"`
	UpdateTrigger types.String `tfsdk:"update_trigger"`
	NIOS          types.Object `tfsdk:"nios"`
}

var SharedrecordSrvAttrTypes = map[string]attr.Type{
	"id":             types.StringType,
	"update_trigger": types.StringType,
	"nios":           types.ObjectType{AttrTypes: NIOSSharedrecordSrvAttrTypes},
}

type NIOSSharedrecordSrvModel struct {
	Comment           types.String `tfsdk:"comment"`
	Disable           types.Bool   `tfsdk:"disable"`
	ExtAttrs          types.Map    `tfsdk:"ext_attrs"`
	ExtAttrsAll       types.Map    `tfsdk:"ext_attrs_all"`
	Name              types.String `tfsdk:"name"`
	Port              types.Int64  `tfsdk:"port"`
	Priority          types.Int64  `tfsdk:"priority"`
	SharedRecordGroup types.String `tfsdk:"shared_record_group"`
	Target            types.String `tfsdk:"target"`
	Ttl               types.Int64  `tfsdk:"ttl"`
	Weight            types.Int64  `tfsdk:"weight"`
}

var NIOSSharedrecordSrvAttrTypes = map[string]attr.Type{
	"comment":             types.StringType,
	"disable":             types.BoolType,
	"ext_attrs":           types.MapType{ElemType: types.StringType},
	"ext_attrs_all":       types.MapType{ElemType: types.StringType},
	"name":                types.StringType,
	"port":                types.Int64Type,
	"priority":            types.Int64Type,
	"shared_record_group": types.StringType,
	"target":              types.StringType,
	"ttl":                 types.Int64Type,
	"weight":              types.Int64Type,
}

const (
	SharedrecordSrvReturnFields = "comment,disable,dns_name,dns_target,extattrs,name,port,priority,shared_record_group,target,ttl,use_ttl,weight"
)

var SharedrecordSrvResourceSchemaAttributes = map[string]schema.Attribute{
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
		Attributes:          SharedrecordSrvResourceNiosSchemaAttributes,
	},
}

var SharedrecordSrvResourceNiosSchemaAttributes = map[string]schema.Attribute{
	"comment": schema.StringAttribute{
		Optional: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
			customvalidator.ValidateTrimmedString(),
			stringvalidator.LengthBetween(0, 256),
		},
		MarkdownDescription: "Comment for this shared record; maximum 256 characters.",
	},
	"disable": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Determines if this shared record is disabled or not. False means that the record is enabled.",
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
		MarkdownDescription: "All ext_attrs including inherited values.",
	},
	"name": schema.StringAttribute{
		Required: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
			customvalidator.ValidateTrimmedString(),
		},
		MarkdownDescription: "Name for this shared record. This value can be in unicode format.",
	},
	"port": schema.Int64Attribute{
		Required: true,
		Validators: []validator.Int64{
			int64validator.Between(0, 65535),
		},
		MarkdownDescription: "The port of the shared SRV record. Valid values are from 0 to 65535 (inclusive), in 32-bit unsigned integer format.",
	},
	"priority": schema.Int64Attribute{
		Required: true,
		Validators: []validator.Int64{
			int64validator.Between(0, 65535),
		},
		MarkdownDescription: "The priority of the shared SRV record. Valid values are from 0 to 65535 (inclusive), in 32-bit unsigned integer format.",
	},
	"shared_record_group": schema.StringAttribute{
		Required: true,
		PlanModifiers: []planmodifier.String{
			immutable.ImmutableString(),
		},
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
		},
		MarkdownDescription: "The name of the shared record group in which the record resides.",
	},
	"target": schema.StringAttribute{
		Required: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
			customvalidator.IsValidNIOSDomainName(),
		},
		MarkdownDescription: "The target of the shared SRV record in FQDN format. This value can be in unicode format.",
	},
	"ttl": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "The Time To Live (TTL) value for this shared record. A 32-bit unsigned integer that represents the duration, in seconds, for which the shared record is valid (cached). Zero indicates that the shared record should not be cached.",
	},
	"weight": schema.Int64Attribute{
		Required: true,
		Validators: []validator.Int64{
			int64validator.Between(0, 65535),
		},
		MarkdownDescription: "The weight of the shared SRV record. Valid values are from 0 to 65535 (inclusive), in 32-bit unsigned integer format.",
	},
}

// Expand converts the TF model to the infoblox core model
func (m *SharedrecordSrvModel) Expand(ctx context.Context, diags *diag.Diagnostics, isCreate bool) *coremodel.SharedrecordSrv {
	if m == nil {
		return nil
	}

	obj := &coremodel.SharedrecordSrv{}

	// Expand NIOS nested attribute (returns nil if not present)
	niosModel := flex.ExpandNestedObject[NIOSSharedrecordSrvModel](ctx, m.NIOS, diags)
	if niosModel != nil {
		obj.NIOS = niosModel.Expand(ctx, diags)
	}

	return obj
}

// Expand converts the NIOS TF model to the core model.
func (m *NIOSSharedrecordSrvModel) Expand(ctx context.Context, diags *diag.Diagnostics) *coremodel.NIOSSharedrecordSrvExt {
	return &coremodel.NIOSSharedrecordSrvExt{
		Comment:           flex.ExpandStringPointerNullAsEmpty(m.Comment),
		Disable:           flex.ExpandBoolPointer(m.Disable),
		ExtAttrs:          flex.ExpandMapStringAny(ctx, m.ExtAttrs, diags),
		Name:              flex.ExpandStringPointerNullAsEmpty(m.Name),
		Port:              flex.ExpandInt64Pointer(m.Port),
		Priority:          flex.ExpandInt64Pointer(m.Priority),
		SharedRecordGroup: flex.ExpandStringPointerNullAsEmpty(m.SharedRecordGroup),
		Target:            flex.ExpandStringPointerNullAsEmpty(m.Target),
		Ttl:               flex.ExpandInt64Pointer(m.Ttl),
		Weight:            flex.ExpandInt64Pointer(m.Weight),
	}
}

// ApplySharedrecordSrvNIOSUseFlags derives NIOS use flags from the raw config
// value(s) and writes them onto the core model. A flag is true when the user
// set any of its governed value fields in config.
func ApplySharedrecordSrvNIOSUseFlags(ctx context.Context, config tfsdk.Config, obj *coremodel.SharedrecordSrv, diags *diag.Diagnostics) {
	if obj == nil || obj.NIOS == nil {
		return
	}
	obj.NIOS.UseTtl = flex.DeriveUseFlag(ctx, config, diags, path.Root("nios").AtName("ttl"))
}

// Flatten populates the TF model from a core response.
func (m *SharedrecordSrvModel) Flatten(ctx context.Context, resp *coremodel.SharedrecordSrv, diags *diag.Diagnostics) {
	if resp == nil {
		return
	}

	m.Id = flex.FlattenStringPointer(resp.Id)

	// Extract existing NIOS model, flatten API response onto it, convert back
	niosModel := flex.ExpandNestedObject[NIOSSharedrecordSrvModel](ctx, m.NIOS, diags)
	if niosModel == nil {
		niosModel = &NIOSSharedrecordSrvModel{}
	}
	niosModel.Flatten(ctx, resp.NIOS, diags)
	if resp.NIOS != nil {
		m.NIOS = flex.FlattenNestedObject(ctx, niosModel, NIOSSharedrecordSrvAttrTypes, diags)
	} else {
		m.NIOS = types.ObjectNull(NIOSSharedrecordSrvAttrTypes)
	}

}

// Flatten merges API response onto existing NIOS model.
func (m *NIOSSharedrecordSrvModel) Flatten(ctx context.Context, from *coremodel.NIOSSharedrecordSrvExt, diags *diag.Diagnostics) {
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
	m.Name = flex.FlattenStringPointerEmptyAsNull(from.Name)
	m.Port = flex.FlattenInt64Pointer(from.Port)
	m.Priority = flex.FlattenInt64Pointer(from.Priority)
	m.SharedRecordGroup = flex.FlattenStringPointerEmptyAsNull(from.SharedRecordGroup)
	m.Target = flex.FlattenStringPointerEmptyAsNull(from.Target)
	m.Ttl = flex.FlattenInt64Pointer(from.Ttl)
	m.Weight = flex.FlattenInt64Pointer(from.Weight)
}
