package dns

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	uddidns "github.com/infobloxopen/universal-ddi-go-client/dnsconfig"
)

// HostInheritanceModel is the Terraform model for HostInheritance
type HostInheritanceModel struct {
	KerberosKeys types.Object `tfsdk:"kerberos_keys"`
}

// HostInheritanceAttrTypes contains the attribute types for HostInheritanceModel
var HostInheritanceAttrTypes = map[string]attr.Type{
	"kerberos_keys": types.ObjectType{AttrTypes: InheritedKerberosKeysAttrTypes},
}

// HostInheritanceResourceSchemaAttributes contains the schema attributes for HostInheritanceModel
var HostInheritanceResourceSchemaAttributes = map[string]schema.Attribute{
	"kerberos_keys": schema.SingleNestedAttribute{
		Attributes:          InheritedKerberosKeysResourceSchemaAttributes,
		Computed:            true,
		MarkdownDescription: "Optional. Field config for _kerberos_keys_ field from _Host_ object.",
	},
}

// ExpandHostInheritance converts a Terraform Object to SDK type
func ExpandHostInheritance(ctx context.Context, o types.Object, diags *diag.Diagnostics) *uddidns.HostInheritance {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m HostInheritanceModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

// Expand converts the Terraform model to SDK type
func (m *HostInheritanceModel) Expand(ctx context.Context, diags *diag.Diagnostics) *uddidns.HostInheritance {
	if m == nil {
		return nil
	}
	to := &uddidns.HostInheritance{
		KerberosKeys: ExpandInheritedKerberosKeys(ctx, m.KerberosKeys, diags),
	}
	return to
}

// FlattenHostInheritance converts an SDK type to Terraform Object
func FlattenHostInheritance(ctx context.Context, from *uddidns.HostInheritance, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(HostInheritanceAttrTypes)
	}
	m := &HostInheritanceModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, HostInheritanceAttrTypes, m)
	diags.Append(d...)
	return t
}

// Flatten populates the Terraform model from SDK type
func (m *HostInheritanceModel) Flatten(ctx context.Context, from *uddidns.HostInheritance, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	m.KerberosKeys = FlattenInheritedKerberosKeys(ctx, from.KerberosKeys, diags)
}
