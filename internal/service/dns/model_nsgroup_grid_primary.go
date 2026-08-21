package dns

import (
	"context"

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

// NsgroupGridPrimaryModel is the Terraform model for NsgroupGridPrimary
type NsgroupGridPrimaryModel struct {
	Name    types.String `tfsdk:"name"`
	Stealth types.Bool   `tfsdk:"stealth"`
}

// NsgroupGridPrimaryAttrTypes contains the attribute types for NsgroupGridPrimaryModel
var NsgroupGridPrimaryAttrTypes = map[string]attr.Type{
	"name":    types.StringType,
	"stealth": types.BoolType,
}

// NsgroupGridPrimaryResourceSchemaAttributes contains the schema attributes for NsgroupGridPrimaryModel
var NsgroupGridPrimaryResourceSchemaAttributes = map[string]schema.Attribute{
	"name": schema.StringAttribute{
		Required: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
		},
		MarkdownDescription: "The grid member name.",
	},
	"stealth": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "This flag governs whether the specified Grid member is in stealth mode or not. If set to True, the member is in stealth mode. This flag is ignored if the struct is specified as part of a stub zone.",
	},
}

// ExpandNsgroupGridPrimary converts a Terraform Object to SDK type
func ExpandNsgroupGridPrimary(ctx context.Context, o types.Object, diags *diag.Diagnostics) *niosdns.NsgroupGridPrimary {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m NsgroupGridPrimaryModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

// Expand converts the Terraform model to SDK type
func (m *NsgroupGridPrimaryModel) Expand(ctx context.Context, diags *diag.Diagnostics) *niosdns.NsgroupGridPrimary {
	if m == nil {
		return nil
	}
	to := &niosdns.NsgroupGridPrimary{
		Name:    flex.ExpandStringPointerNullAsEmpty(m.Name),
		Stealth: flex.ExpandBoolPointer(m.Stealth),
	}
	return to
}

// FlattenNsgroupGridPrimary converts an SDK type to Terraform Object
func FlattenNsgroupGridPrimary(ctx context.Context, from *niosdns.NsgroupGridPrimary, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(NsgroupGridPrimaryAttrTypes)
	}
	m := &NsgroupGridPrimaryModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, NsgroupGridPrimaryAttrTypes, m)
	diags.Append(d...)
	return t
}

// Flatten populates the Terraform model from SDK type
func (m *NsgroupGridPrimaryModel) Flatten(ctx context.Context, from *niosdns.NsgroupGridPrimary, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	m.Name = flex.FlattenStringPointerEmptyAsNull(from.Name)
	m.Stealth = flex.FlattenBoolPointer(from.Stealth)
}
