package dns

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
	customvalidator "github.com/infobloxopen/terraform-provider-infoblox/internal/validator"
	uddidns "github.com/infobloxopen/universal-ddi-go-client/dnsconfig"
)

// InheritedKerberosKeysModel is the Terraform model for InheritedKerberosKeys
type InheritedKerberosKeysModel struct {
	Action      types.String `tfsdk:"action"`
	DisplayName types.String `tfsdk:"display_name"`
	Source      types.String `tfsdk:"source"`
	Value       types.List   `tfsdk:"value"`
}

// InheritedKerberosKeysAttrTypes contains the attribute types for InheritedKerberosKeysModel
var InheritedKerberosKeysAttrTypes = map[string]attr.Type{
	"action":       types.StringType,
	"display_name": types.StringType,
	"source":       types.StringType,
	"value":        types.ListType{ElemType: types.ObjectType{AttrTypes: KerberosKeyAttrTypes}},
}

// InheritedKerberosKeysResourceSchemaAttributes contains the schema attributes for InheritedKerberosKeysModel
var InheritedKerberosKeysResourceSchemaAttributes = map[string]schema.Attribute{
	"action": schema.StringAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "Optional. Inheritance setting for a field. Defaults to _inherit_.",
	},
	"display_name": schema.StringAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "Human-readable display name for the object referred to by _source_.",
	},
	"source": schema.StringAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The resource identifier.",
	},
	"value": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: KerberosKeyResourceSchemaAttributes,
		},
		Optional: true,
		Computed: true,
		Validators: []validator.List{
			customvalidator.ListNotEmpty(),
		},
		MarkdownDescription: "Inherited value.",
	},
}

// ExpandInheritedKerberosKeys converts a Terraform Object to SDK type
func ExpandInheritedKerberosKeys(ctx context.Context, o types.Object, diags *diag.Diagnostics) *uddidns.InheritedKerberosKeys {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m InheritedKerberosKeysModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

// Expand converts the Terraform model to SDK type
func (m *InheritedKerberosKeysModel) Expand(ctx context.Context, diags *diag.Diagnostics) *uddidns.InheritedKerberosKeys {
	if m == nil {
		return nil
	}
	to := &uddidns.InheritedKerberosKeys{
		Action:      flex.ExpandStringPointer(m.Action),
		DisplayName: flex.ExpandStringPointer(m.DisplayName),
		Source:      flex.ExpandStringPointer(m.Source),
		Value:       flex.ExpandFrameworkListNestedBlock(ctx, m.Value, diags, ExpandKerberosKey),
	}
	return to
}

// FlattenInheritedKerberosKeys converts an SDK type to Terraform Object
func FlattenInheritedKerberosKeys(ctx context.Context, from *uddidns.InheritedKerberosKeys, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(InheritedKerberosKeysAttrTypes)
	}
	m := &InheritedKerberosKeysModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, InheritedKerberosKeysAttrTypes, m)
	diags.Append(d...)
	return t
}

// Flatten populates the Terraform model from SDK type
func (m *InheritedKerberosKeysModel) Flatten(ctx context.Context, from *uddidns.InheritedKerberosKeys, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	m.Action = flex.FlattenStringPointer(from.Action)
	m.DisplayName = flex.FlattenStringPointer(from.DisplayName)
	m.Source = flex.FlattenStringPointer(from.Source)
	m.Value = flex.FlattenFrameworkListNestedBlock(ctx, from.Value, KerberosKeyAttrTypes, diags, FlattenKerberosKey)
}
