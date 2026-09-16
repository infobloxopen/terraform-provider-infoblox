package dns

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	niosdns "github.com/infobloxopen/infoblox-nios-go-client/dns"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
	customvalidator "github.com/infobloxopen/terraform-provider-infoblox/internal/validator"
)

// ZoneRpGridPrimaryModel is the Terraform model for ZoneRpGridPrimary
type ZoneRpGridPrimaryModel struct {
	Name    types.String `tfsdk:"name"`
	Stealth types.Bool   `tfsdk:"stealth"`
}

// ZoneRpGridPrimaryAttrTypes contains the attribute types for ZoneRpGridPrimaryModel
var ZoneRpGridPrimaryAttrTypes = map[string]attr.Type{
	"name":    types.StringType,
	"stealth": types.BoolType,
}

// ZoneRpGridPrimaryResourceSchemaAttributes contains the schema attributes for ZoneRpGridPrimaryModel
var ZoneRpGridPrimaryResourceSchemaAttributes = map[string]schema.Attribute{
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
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "This flag governs whether the specified Grid member is in stealth mode or not. If set to True, the member is in stealth mode. This flag is ignored if the struct is specified as part of a stub zone.",
	},
}

// ExpandZoneRpGridPrimary converts a Terraform Object to SDK type
func ExpandZoneRpGridPrimary(ctx context.Context, o types.Object, diags *diag.Diagnostics) *niosdns.ZoneRpGridPrimary {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m ZoneRpGridPrimaryModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

// Expand converts the Terraform model to SDK type
func (m *ZoneRpGridPrimaryModel) Expand(ctx context.Context, diags *diag.Diagnostics) *niosdns.ZoneRpGridPrimary {
	if m == nil {
		return nil
	}
	to := &niosdns.ZoneRpGridPrimary{
		Name:    flex.ExpandStringPointerNullAsEmpty(m.Name),
		Stealth: flex.ExpandBoolPointer(m.Stealth),
	}
	return to
}

// FlattenZoneRpGridPrimary converts an SDK type to Terraform Object
func FlattenZoneRpGridPrimary(ctx context.Context, from *niosdns.ZoneRpGridPrimary, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(ZoneRpGridPrimaryAttrTypes)
	}
	m := &ZoneRpGridPrimaryModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, ZoneRpGridPrimaryAttrTypes, m)
	diags.Append(d...)
	return t
}

// Flatten populates the Terraform model from SDK type
func (m *ZoneRpGridPrimaryModel) Flatten(ctx context.Context, from *niosdns.ZoneRpGridPrimary, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	m.Name = flex.FlattenStringPointerEmptyAsNull(from.Name)
	m.Stealth = flex.FlattenBoolPointer(from.Stealth)
}
