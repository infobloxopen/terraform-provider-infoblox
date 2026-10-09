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

// NsgroupStubmemberStubMembersModel is the Terraform model for NsgroupStubmemberStubMembers
type NsgroupStubmemberStubMembersModel struct {
	Name types.String `tfsdk:"name"`
}

// NsgroupStubmemberStubMembersAttrTypes contains the attribute types for NsgroupStubmemberStubMembersModel
var NsgroupStubmemberStubMembersAttrTypes = map[string]attr.Type{
	"name": types.StringType,
}

// NsgroupStubmemberStubMembersResourceSchemaAttributes contains the schema attributes for NsgroupStubmemberStubMembersModel
var NsgroupStubmemberStubMembersResourceSchemaAttributes = map[string]schema.Attribute{
	"name": schema.StringAttribute{
		Required: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
		},
		MarkdownDescription: "The grid member name.",
	},
}

// ExpandNsgroupStubmemberStubMembers converts a Terraform Object to SDK type
func ExpandNsgroupStubmemberStubMembers(ctx context.Context, o types.Object, diags *diag.Diagnostics) *niosdns.NsgroupStubmemberStubMembers {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m NsgroupStubmemberStubMembersModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

// Expand converts the Terraform model to SDK type
func (m *NsgroupStubmemberStubMembersModel) Expand(ctx context.Context, diags *diag.Diagnostics) *niosdns.NsgroupStubmemberStubMembers {
	if m == nil {
		return nil
	}
	to := &niosdns.NsgroupStubmemberStubMembers{
		Name: flex.ExpandStringPointerNullAsEmpty(m.Name),
	}
	return to
}

// FlattenNsgroupStubmemberStubMembers converts an SDK type to Terraform Object
func FlattenNsgroupStubmemberStubMembers(ctx context.Context, from *niosdns.NsgroupStubmemberStubMembers, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(NsgroupStubmemberStubMembersAttrTypes)
	}
	m := &NsgroupStubmemberStubMembersModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, NsgroupStubmemberStubMembersAttrTypes, m)
	diags.Append(d...)
	return t
}

// Flatten populates the Terraform model from SDK type
func (m *NsgroupStubmemberStubMembersModel) Flatten(ctx context.Context, from *niosdns.NsgroupStubmemberStubMembers, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	m.Name = flex.FlattenStringPointerEmptyAsNull(from.Name)
}
