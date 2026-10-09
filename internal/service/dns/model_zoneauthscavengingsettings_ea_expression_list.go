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
	customvalidator "github.com/infobloxopen/terraform-provider-infoblox/internal/validator"
)

// ZoneauthscavengingsettingsEaExpressionListModel is the Terraform model for ZoneauthscavengingsettingsEaExpressionList
type ZoneauthscavengingsettingsEaExpressionListModel struct {
	Op      types.String `tfsdk:"op"`
	Op1     types.String `tfsdk:"op1"`
	Op1Type types.String `tfsdk:"op1_type"`
	Op2     types.String `tfsdk:"op2"`
	Op2Type types.String `tfsdk:"op2_type"`
}

// ZoneauthscavengingsettingsEaExpressionListAttrTypes contains the attribute types for ZoneauthscavengingsettingsEaExpressionListModel
var ZoneauthscavengingsettingsEaExpressionListAttrTypes = map[string]attr.Type{
	"op":       types.StringType,
	"op1":      types.StringType,
	"op1_type": types.StringType,
	"op2":      types.StringType,
	"op2_type": types.StringType,
}

// ZoneauthscavengingsettingsEaExpressionListResourceSchemaAttributes contains the schema attributes for ZoneauthscavengingsettingsEaExpressionListModel
var ZoneauthscavengingsettingsEaExpressionListResourceSchemaAttributes = map[string]schema.Attribute{
	"op": schema.StringAttribute{
		Validators: []validator.String{
			stringvalidator.OneOf("AND", "ENDLIST", "GT", "LT", "LE", "GE", "MATCH_IP", "MATCH_RANGE", "MATCH_CIDR", "EQ", "EXISTS", "NOT_EQ", "NOT_EXISTS", "OR"),
		},
		Required:            true,
		MarkdownDescription: "The operation name.",
	},
	"op1": schema.StringAttribute{
		Optional: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
		},
		MarkdownDescription: "The name of the Extensible Attribute Definition object which is used as the first operand value.",
	},
	"op1_type": schema.StringAttribute{
		Validators: []validator.String{
			stringvalidator.OneOf("FIELD", "LIST", "STRING"),
		},
		Optional:            true,
		MarkdownDescription: "The first operand type.",
	},
	"op2": schema.StringAttribute{
		Optional: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
		},
		MarkdownDescription: "The second operand value.",
	},
	"op2_type": schema.StringAttribute{
		Validators: []validator.String{
			stringvalidator.OneOf("FIELD", "LIST", "STRING"),
		},
		Optional:            true,
		MarkdownDescription: "The second operand type.",
	},
}

// ExpandZoneauthscavengingsettingsEaExpressionList converts a Terraform Object to SDK type
func ExpandZoneauthscavengingsettingsEaExpressionList(ctx context.Context, o types.Object, diags *diag.Diagnostics) *niosdns.ZoneauthscavengingsettingsEaExpressionList {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m ZoneauthscavengingsettingsEaExpressionListModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

// Expand converts the Terraform model to SDK type
func (m *ZoneauthscavengingsettingsEaExpressionListModel) Expand(ctx context.Context, diags *diag.Diagnostics) *niosdns.ZoneauthscavengingsettingsEaExpressionList {
	if m == nil {
		return nil
	}
	to := &niosdns.ZoneauthscavengingsettingsEaExpressionList{
		Op:      flex.ExpandStringPointer(m.Op),
		Op1:     flex.ExpandStringPointerNullAsEmpty(m.Op1),
		Op1Type: flex.ExpandStringPointer(m.Op1Type),
		Op2:     flex.ExpandStringPointerNullAsEmpty(m.Op2),
		Op2Type: flex.ExpandStringPointer(m.Op2Type),
	}
	return to
}

// FlattenZoneauthscavengingsettingsEaExpressionList converts an SDK type to Terraform Object
func FlattenZoneauthscavengingsettingsEaExpressionList(ctx context.Context, from *niosdns.ZoneauthscavengingsettingsEaExpressionList, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(ZoneauthscavengingsettingsEaExpressionListAttrTypes)
	}
	m := &ZoneauthscavengingsettingsEaExpressionListModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, ZoneauthscavengingsettingsEaExpressionListAttrTypes, m)
	diags.Append(d...)
	return t
}

// Flatten populates the Terraform model from SDK type
func (m *ZoneauthscavengingsettingsEaExpressionListModel) Flatten(ctx context.Context, from *niosdns.ZoneauthscavengingsettingsEaExpressionList, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	m.Op = flex.FlattenStringPointerEmptyAsNull(from.Op)
	m.Op1 = flex.FlattenStringPointerEmptyAsNull(from.Op1)
	m.Op1Type = flex.FlattenStringPointerEmptyAsNull(from.Op1Type)
	m.Op2 = flex.FlattenStringPointerEmptyAsNull(from.Op2)
	m.Op2Type = flex.FlattenStringPointerEmptyAsNull(from.Op2Type)
}
