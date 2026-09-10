package dtc

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	niosdtc "github.com/infobloxopen/infoblox-nios-go-client/dtc"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
	customvalidator "github.com/infobloxopen/terraform-provider-infoblox/internal/validator"
)

// MonitorSnmpOidsModel is the Terraform model for MonitorSnmpOids
type MonitorSnmpOidsModel struct {
	Oid       types.String `tfsdk:"oid"`
	Comment   types.String `tfsdk:"comment"`
	Type      types.String `tfsdk:"type"`
	Condition types.String `tfsdk:"condition"`
	First     types.String `tfsdk:"first"`
	Last      types.String `tfsdk:"last"`
}

// MonitorSnmpOidsAttrTypes contains the attribute types for MonitorSnmpOidsModel
var MonitorSnmpOidsAttrTypes = map[string]attr.Type{
	"oid":       types.StringType,
	"comment":   types.StringType,
	"type":      types.StringType,
	"condition": types.StringType,
	"first":     types.StringType,
	"last":      types.StringType,
}

// MonitorSnmpOidsResourceSchemaAttributes contains the schema attributes for MonitorSnmpOidsModel
var MonitorSnmpOidsResourceSchemaAttributes = map[string]schema.Attribute{
	"oid": schema.StringAttribute{
		Optional: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
		},
		MarkdownDescription: "The SNMP OID value for DTC SNMP Monitor health checks.",
	},
	"comment": schema.StringAttribute{
		Optional: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
		},
		MarkdownDescription: "The comment for a DTC SNMP Health Monitor OID object.",
	},
	"type": schema.StringAttribute{
		Validators: []validator.String{
			stringvalidator.OneOf("STRING", "INTEGER"),
		},
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The value of the condition type for DTC SNMP Monitor health check results.",
	},
	"condition": schema.StringAttribute{
		Validators: []validator.String{
			stringvalidator.OneOf("ANY", "EXACT", "LEQ", "GEQ", "RANGE"),
		},
		Optional:            true,
		MarkdownDescription: "The condition of the validation result for an SNMP health check. The following conditions can be applied to the health check results: 'ANY' accepts any response; 'EXACT' accepts result equal to 'first'; 'LEQ' accepts result which is less than 'first'; 'GEQ' accepts result which is greater than 'first'; 'RANGE' accepts result value of which is between 'first' and 'last'.",
	},
	"first": schema.StringAttribute{
		Optional: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
		},
		MarkdownDescription: "The condition's first term to match against the SNMP health check result.",
	},
	"last": schema.StringAttribute{
		Optional: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
		},
		MarkdownDescription: "The condition's second term to match against the SNMP health check result with 'RANGE' condition.",
	},
}

// ExpandMonitorSnmpOids converts a Terraform Object to SDK type
func ExpandMonitorSnmpOids(ctx context.Context, o types.Object, diags *diag.Diagnostics) *niosdtc.DtcMonitorSnmpOids {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m MonitorSnmpOidsModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

// Expand converts the Terraform model to SDK type
func (m *MonitorSnmpOidsModel) Expand(ctx context.Context, diags *diag.Diagnostics) *niosdtc.DtcMonitorSnmpOids {
	if m == nil {
		return nil
	}
	to := &niosdtc.DtcMonitorSnmpOids{
		Oid:       flex.ExpandStringPointerNullAsEmpty(m.Oid),
		Comment:   flex.ExpandStringPointerNullAsEmpty(m.Comment),
		Type:      flex.ExpandStringPointer(m.Type),
		Condition: flex.ExpandStringPointer(m.Condition),
		First:     flex.ExpandStringPointerNullAsEmpty(m.First),
		Last:      flex.ExpandStringPointerNullAsEmpty(m.Last),
	}
	return to
}

// FlattenMonitorSnmpOids converts an SDK type to Terraform Object
func FlattenMonitorSnmpOids(ctx context.Context, from *niosdtc.DtcMonitorSnmpOids, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(MonitorSnmpOidsAttrTypes)
	}
	m := &MonitorSnmpOidsModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, MonitorSnmpOidsAttrTypes, m)
	diags.Append(d...)
	return t
}

// Flatten populates the Terraform model from SDK type
func (m *MonitorSnmpOidsModel) Flatten(ctx context.Context, from *niosdtc.DtcMonitorSnmpOids, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	m.Oid = flex.FlattenStringPointerEmptyAsNull(from.Oid)
	m.Comment = flex.FlattenStringPointerEmptyAsNull(from.Comment)
	m.Type = flex.FlattenStringPointerEmptyAsNull(from.Type)
	m.Condition = flex.FlattenStringPointerEmptyAsNull(from.Condition)
	m.First = flex.FlattenStringPointerEmptyAsNull(from.First)
	m.Last = flex.FlattenStringPointerEmptyAsNull(from.Last)
}
