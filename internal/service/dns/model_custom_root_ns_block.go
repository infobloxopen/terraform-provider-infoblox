package dns

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
	uddidns "github.com/infobloxopen/universal-ddi-go-client/dnsconfig"
)

// CustomRootNSBlockModel is the Terraform model for CustomRootNSBlock
type CustomRootNSBlockModel struct {
	CustomRootNs        types.List `tfsdk:"custom_root_ns"`
	CustomRootNsEnabled types.Bool `tfsdk:"custom_root_ns_enabled"`
}

// CustomRootNSBlockAttrTypes contains the attribute types for CustomRootNSBlockModel
var CustomRootNSBlockAttrTypes = map[string]attr.Type{
	"custom_root_ns":         types.ListType{ElemType: types.ObjectType{AttrTypes: RootNSAttrTypes}},
	"custom_root_ns_enabled": types.BoolType,
}

// CustomRootNSBlockResourceSchemaAttributes contains the schema attributes for CustomRootNSBlockModel
var CustomRootNSBlockResourceSchemaAttributes = map[string]schema.Attribute{
	"custom_root_ns": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: RootNSResourceSchemaAttributes,
		},
		Optional: true,
		Validators: []validator.List{
			listvalidator.SizeAtLeast(1),
		},
		MarkdownDescription: "Optional. Field config for _custom_root_ns_ field.",
	},
	"custom_root_ns_enabled": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Optional. Field config for _custom_root_ns_enabled_ field.",
	},
}

// ExpandCustomRootNSBlock converts a Terraform Object to SDK type
func ExpandCustomRootNSBlock(ctx context.Context, o types.Object, diags *diag.Diagnostics) *uddidns.CustomRootNSBlock {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m CustomRootNSBlockModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

// Expand converts the Terraform model to SDK type
func (m *CustomRootNSBlockModel) Expand(ctx context.Context, diags *diag.Diagnostics) *uddidns.CustomRootNSBlock {
	if m == nil {
		return nil
	}
	to := &uddidns.CustomRootNSBlock{
		CustomRootNs:        flex.ExpandFrameworkListNestedBlock(ctx, m.CustomRootNs, diags, ExpandRootNS),
		CustomRootNsEnabled: flex.ExpandBoolPointer(m.CustomRootNsEnabled),
	}
	return to
}

// FlattenCustomRootNSBlock converts an SDK type to Terraform Object
func FlattenCustomRootNSBlock(ctx context.Context, from *uddidns.CustomRootNSBlock, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(CustomRootNSBlockAttrTypes)
	}
	m := &CustomRootNSBlockModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, CustomRootNSBlockAttrTypes, m)
	diags.Append(d...)
	return t
}

// Flatten populates the Terraform model from SDK type
func (m *CustomRootNSBlockModel) Flatten(ctx context.Context, from *uddidns.CustomRootNSBlock, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	m.CustomRootNs = flex.FlattenFrameworkListNestedBlock(ctx, from.CustomRootNs, RootNSAttrTypes, diags, FlattenRootNS)
	m.CustomRootNsEnabled = flex.FlattenBoolPointer(from.CustomRootNsEnabled)
}
