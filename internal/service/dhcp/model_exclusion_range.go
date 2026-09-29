package dhcp

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
	uddidhcp "github.com/infobloxopen/universal-ddi-go-client/ipam"
)

// ExclusionRangeModel is the Terraform model for ExclusionRange
type ExclusionRangeModel struct {
	Comment types.String `tfsdk:"comment"`
	End     types.String `tfsdk:"end"`
	Start   types.String `tfsdk:"start"`
}

// ExclusionRangeAttrTypes contains the attribute types for ExclusionRangeModel
var ExclusionRangeAttrTypes = map[string]attr.Type{
	"comment": types.StringType,
	"end":     types.StringType,
	"start":   types.StringType,
}

// ExclusionRangeResourceSchemaAttributes contains the schema attributes for ExclusionRangeModel
var ExclusionRangeResourceSchemaAttributes = map[string]schema.Attribute{
	"comment": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Validators: []validator.String{
			stringvalidator.LengthBetween(0, 1024),
		},
		MarkdownDescription: "The description for the exclusion range. May contain 0 to 1024 characters. Can include UTF-8.",
	},
	"end": schema.StringAttribute{
		Required:            true,
		MarkdownDescription: "The end address of the exclusion range.",
	},
	"start": schema.StringAttribute{
		Required:            true,
		MarkdownDescription: "The start address of the exclusion range.",
	},
}

// ExpandExclusionRange converts a Terraform Object to SDK type
func ExpandExclusionRange(ctx context.Context, o types.Object, diags *diag.Diagnostics) *uddidhcp.ExclusionRange {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m ExclusionRangeModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

// Expand converts the Terraform model to SDK type
func (m *ExclusionRangeModel) Expand(ctx context.Context, diags *diag.Diagnostics) *uddidhcp.ExclusionRange {
	if m == nil {
		return nil
	}
	to := &uddidhcp.ExclusionRange{
		Comment: flex.ExpandStringPointer(m.Comment),
		End:     flex.ExpandString(m.End),
		Start:   flex.ExpandString(m.Start),
	}
	return to
}

// FlattenExclusionRange converts an SDK type to Terraform Object
func FlattenExclusionRange(ctx context.Context, from *uddidhcp.ExclusionRange, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(ExclusionRangeAttrTypes)
	}
	m := &ExclusionRangeModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, ExclusionRangeAttrTypes, m)
	diags.Append(d...)
	return t
}

// Flatten populates the Terraform model from SDK type
func (m *ExclusionRangeModel) Flatten(ctx context.Context, from *uddidhcp.ExclusionRange, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	m.Comment = flex.FlattenStringPointer(from.Comment)
	m.End = flex.FlattenString(from.End)
	m.Start = flex.FlattenString(from.Start)
}
