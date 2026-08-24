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

	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
	uddidtc "github.com/infobloxopen/universal-ddi-go-client/dtc"
)

// RecordModel is the Terraform model for Record
type RecordModel struct {
	Rdata types.Map    `tfsdk:"rdata"`
	Type  types.String `tfsdk:"type"`
}

// RecordAttrTypes contains the attribute types for RecordModel
var RecordAttrTypes = map[string]attr.Type{
	"rdata": types.MapType{ElemType: types.StringType},
	"type":  types.StringType,
}

// RecordResourceSchemaAttributes contains the schema attributes for RecordModel
var RecordResourceSchemaAttributes = map[string]schema.Attribute{
	"rdata": schema.MapAttribute{
		ElementType:         types.StringType,
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "JSON representation of resource record data.",
	},
	"type": schema.StringAttribute{
		Validators: []validator.String{
			stringvalidator.OneOf("A", "AAAA", "CNAME", "HTTPS", "SRV", "SVCB"),
		},
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "Resource record type.  List of supported types: * _A_ (_TYPE1_) * _AAAA_ (_TYPE28_) * _CNAME_ (_TYPE5_) * _HTTPS_ (_TYPE65_) * _SRV_ (_TYPE33_) * _SVCB_ (_TYPE64_)",
	},
}

// ExpandRecord converts a Terraform Object to SDK type
func ExpandRecord(ctx context.Context, o types.Object, diags *diag.Diagnostics) *uddidtc.Record {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m RecordModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

// Expand converts the Terraform model to SDK type
func (m *RecordModel) Expand(ctx context.Context, diags *diag.Diagnostics) *uddidtc.Record {
	if m == nil {
		return nil
	}
	to := &uddidtc.Record{
		Rdata: flex.ExpandMapStringAny(ctx, m.Rdata, diags),
		Type:  flex.ExpandString(m.Type),
	}
	return to
}

// FlattenRecord converts an SDK type to Terraform Object
func FlattenRecord(ctx context.Context, from *uddidtc.Record, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(RecordAttrTypes)
	}
	m := &RecordModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, RecordAttrTypes, m)
	diags.Append(d...)
	return t
}

// Flatten populates the Terraform model from SDK type
func (m *RecordModel) Flatten(ctx context.Context, from *uddidtc.Record, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	m.Rdata = flex.FlattenMapStringAny(ctx, from.Rdata, diags)
	m.Type = flex.FlattenString(from.Type)
}
