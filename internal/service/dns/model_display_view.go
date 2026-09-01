package dns

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
	uddidns "github.com/infobloxopen/universal-ddi-go-client/dnsconfig"
)

// DisplayViewModel is the Terraform model for DisplayView
type DisplayViewModel struct {
	Comment types.String `tfsdk:"comment"`
	Name    types.String `tfsdk:"name"`
	View    types.String `tfsdk:"view"`
}

// DisplayViewAttrTypes contains the attribute types for DisplayViewModel
var DisplayViewAttrTypes = map[string]attr.Type{
	"comment": types.StringType,
	"name":    types.StringType,
	"view":    types.StringType,
}

// DisplayViewResourceSchemaAttributes contains the schema attributes for DisplayViewModel
var DisplayViewResourceSchemaAttributes = map[string]schema.Attribute{
	"comment": schema.StringAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "DNS view description.",
	},
	"name": schema.StringAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "DNS view name.",
	},
	"view": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The resource identifier.",
	},
}

// ExpandDisplayView converts a Terraform Object to SDK type
func ExpandDisplayView(ctx context.Context, o types.Object, diags *diag.Diagnostics) *uddidns.DisplayView {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m DisplayViewModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

// Expand converts the Terraform model to SDK type
func (m *DisplayViewModel) Expand(ctx context.Context, diags *diag.Diagnostics) *uddidns.DisplayView {
	if m == nil {
		return nil
	}
	to := &uddidns.DisplayView{
		Comment: flex.ExpandStringPointer(m.Comment),
		Name:    flex.ExpandStringPointer(m.Name),
		View:    flex.ExpandStringPointer(m.View),
	}
	return to
}

// FlattenDisplayView converts an SDK type to Terraform Object
func FlattenDisplayView(ctx context.Context, from *uddidns.DisplayView, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(DisplayViewAttrTypes)
	}
	m := &DisplayViewModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, DisplayViewAttrTypes, m)
	diags.Append(d...)
	return t
}

// Flatten populates the Terraform model from SDK type
func (m *DisplayViewModel) Flatten(ctx context.Context, from *uddidns.DisplayView, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	m.Comment = flex.FlattenStringPointer(from.Comment)
	m.Name = flex.FlattenStringPointer(from.Name)
	m.View = flex.FlattenStringPointer(from.View)
}
