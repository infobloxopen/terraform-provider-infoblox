package dns

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	niosdns "github.com/infobloxopen/infoblox-nios-go-client/dns"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
)

// ViewscavengingsettingsEaExpressionListModel is the Terraform model for ViewscavengingsettingsEaExpressionList
type ViewscavengingsettingsEaExpressionListModel struct {
	Op      types.String `tfsdk:"op"`
	Op1     types.String `tfsdk:"op1"`
	Op1Type types.String `tfsdk:"op1_type"`
	Op2     types.String `tfsdk:"op2"`
	Op2Type types.String `tfsdk:"op2_type"`
}

// ViewscavengingsettingsEaExpressionListAttrTypes contains the attribute types for ViewscavengingsettingsEaExpressionListModel
var ViewscavengingsettingsEaExpressionListAttrTypes = map[string]attr.Type{
	"op":       types.StringType,
	"op1":      types.StringType,
	"op1_type": types.StringType,
	"op2":      types.StringType,
	"op2_type": types.StringType,
}

// ViewscavengingsettingsEaExpressionListResourceSchemaAttributes contains the schema attributes for ViewscavengingsettingsEaExpressionListModel
var ViewscavengingsettingsEaExpressionListResourceSchemaAttributes = map[string]schema.Attribute{
	"op": schema.StringAttribute{
		Validators: []validator.String{
			stringvalidator.OneOf("AND", "ENDLIST", "GT", "LT", "LE", "GE", "MATCH_IP", "MATCH_RANGE", "MATCH_CIDR", "EQ", "EXISTS", "NOT_EQ", "NOT_EXISTS", "OR"),
		},
		Required:            true,
		MarkdownDescription: "The operation name.",
	},
	"op1": schema.StringAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The name of the Extensible Attribute Definition object which is used as the first operand value.",
	},
	"op1_type": schema.StringAttribute{
		Validators: []validator.String{
			stringvalidator.OneOf("FIELD", "LIST", "STRING"),
		},
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The first operand type.",
	},
	"op2": schema.StringAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The second operand value.",
	},
	"op2_type": schema.StringAttribute{
		Validators: []validator.String{
			stringvalidator.OneOf("FIELD", "LIST", "STRING"),
		},
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The second operand type.",
	},
}

// ExpandViewscavengingsettingsEaExpressionList converts a Terraform Object to SDK type
func ExpandViewscavengingsettingsEaExpressionList(ctx context.Context, o types.Object, diags *diag.Diagnostics) *niosdns.ViewscavengingsettingsEaExpressionList {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m ViewscavengingsettingsEaExpressionListModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

// Expand converts the Terraform model to SDK type
func (m *ViewscavengingsettingsEaExpressionListModel) Expand(ctx context.Context, diags *diag.Diagnostics) *niosdns.ViewscavengingsettingsEaExpressionList {
	if m == nil {
		return nil
	}
	to := &niosdns.ViewscavengingsettingsEaExpressionList{
		Op:      flex.ExpandStringPointerNullAsEmpty(m.Op),
		Op1:     flex.ExpandStringPointerNullAsEmpty(m.Op1),
		Op1Type: flex.ExpandStringPointerNullAsEmpty(m.Op1Type),
		Op2:     flex.ExpandStringPointerNullAsEmpty(m.Op2),
		Op2Type: flex.ExpandStringPointerNullAsEmpty(m.Op2Type),
	}
	return to
}

// FlattenViewscavengingsettingsEaExpressionList converts an SDK type to Terraform Object
func FlattenViewscavengingsettingsEaExpressionList(ctx context.Context, from *niosdns.ViewscavengingsettingsEaExpressionList, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(ViewscavengingsettingsEaExpressionListAttrTypes)
	}
	m := &ViewscavengingsettingsEaExpressionListModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, ViewscavengingsettingsEaExpressionListAttrTypes, m)
	diags.Append(d...)
	return t
}

// Flatten populates the Terraform model from SDK type
func (m *ViewscavengingsettingsEaExpressionListModel) Flatten(ctx context.Context, from *niosdns.ViewscavengingsettingsEaExpressionList, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	m.Op = flex.FlattenStringPointerEmptyAsNull(from.Op)
	m.Op1 = flex.FlattenStringPointerEmptyAsNull(from.Op1)
	m.Op1Type = flex.FlattenStringPointerEmptyAsNull(from.Op1Type)
	m.Op2 = flex.FlattenStringPointerEmptyAsNull(from.Op2)
	m.Op2Type = flex.FlattenStringPointerEmptyAsNull(from.Op2Type)
}
