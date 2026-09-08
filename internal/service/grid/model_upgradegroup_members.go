package grid

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	niosgrid "github.com/infobloxopen/infoblox-nios-go-client/grid"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
	customvalidator "github.com/infobloxopen/terraform-provider-infoblox/internal/validator"
)

// UpgradegroupMembersModel is the Terraform model for UpgradegroupMembers
type UpgradegroupMembersModel struct {
	Member types.String `tfsdk:"member"`
}

// UpgradegroupMembersAttrTypes contains the attribute types for UpgradegroupMembersModel
var UpgradegroupMembersAttrTypes = map[string]attr.Type{
	"member": types.StringType,
}

// UpgradegroupMembersResourceSchemaAttributes contains the schema attributes for UpgradegroupMembersModel
var UpgradegroupMembersResourceSchemaAttributes = map[string]schema.Attribute{
	"member": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
		},
		MarkdownDescription: "The upgrade group member name.",
	},
}

// ExpandUpgradegroupMembers converts a Terraform Object to SDK type
func ExpandUpgradegroupMembers(ctx context.Context, o types.Object, diags *diag.Diagnostics) *niosgrid.UpgradegroupMembers {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m UpgradegroupMembersModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

// Expand converts the Terraform model to SDK type
func (m *UpgradegroupMembersModel) Expand(ctx context.Context, diags *diag.Diagnostics) *niosgrid.UpgradegroupMembers {
	if m == nil {
		return nil
	}
	to := &niosgrid.UpgradegroupMembers{
		Member: flex.ExpandStringPointerNullAsEmpty(m.Member),
	}
	return to
}

// FlattenUpgradegroupMembers converts an SDK type to Terraform Object
func FlattenUpgradegroupMembers(ctx context.Context, from *niosgrid.UpgradegroupMembers, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(UpgradegroupMembersAttrTypes)
	}
	m := &UpgradegroupMembersModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, UpgradegroupMembersAttrTypes, m)
	diags.Append(d...)
	return t
}

// Flatten populates the Terraform model from SDK type
func (m *UpgradegroupMembersModel) Flatten(ctx context.Context, from *niosgrid.UpgradegroupMembers, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	m.Member = flex.FlattenStringPointerEmptyAsNull(from.Member)
}
