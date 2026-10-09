package ipam

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	niosipam "github.com/infobloxopen/infoblox-nios-go-client/ipam"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
	customvalidator "github.com/infobloxopen/terraform-provider-infoblox/internal/validator"
)

// NetworkviewFederatedRealmsModel is the Terraform model for NetworkviewFederatedRealms
type NetworkviewFederatedRealmsModel struct {
	Name types.String `tfsdk:"name"`
	Id   types.String `tfsdk:"id"`
}

// NetworkviewFederatedRealmsAttrTypes contains the attribute types for NetworkviewFederatedRealmsModel
var NetworkviewFederatedRealmsAttrTypes = map[string]attr.Type{
	"name": types.StringType,
	"id":   types.StringType,
}

// NetworkviewFederatedRealmsResourceSchemaAttributes contains the schema attributes for NetworkviewFederatedRealmsModel
var NetworkviewFederatedRealmsResourceSchemaAttributes = map[string]schema.Attribute{
	"name": schema.StringAttribute{
		Required: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
		},
		MarkdownDescription: "The federated realm name",
	},
	"id": schema.StringAttribute{
		Required: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
		},
		MarkdownDescription: "The federated realm id",
	},
}

// ExpandNetworkviewFederatedRealms converts a Terraform Object to SDK type
func ExpandNetworkviewFederatedRealms(ctx context.Context, o types.Object, diags *diag.Diagnostics) *niosipam.NetworkviewFederatedRealms {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m NetworkviewFederatedRealmsModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

// Expand converts the Terraform model to SDK type
func (m *NetworkviewFederatedRealmsModel) Expand(ctx context.Context, diags *diag.Diagnostics) *niosipam.NetworkviewFederatedRealms {
	if m == nil {
		return nil
	}
	to := &niosipam.NetworkviewFederatedRealms{
		Name: flex.ExpandStringPointerNullAsEmpty(m.Name),
		Id:   flex.ExpandStringPointerNullAsEmpty(m.Id),
	}
	return to
}

// FlattenNetworkviewFederatedRealms converts an SDK type to Terraform Object
func FlattenNetworkviewFederatedRealms(ctx context.Context, from *niosipam.NetworkviewFederatedRealms, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(NetworkviewFederatedRealmsAttrTypes)
	}
	m := &NetworkviewFederatedRealmsModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, NetworkviewFederatedRealmsAttrTypes, m)
	diags.Append(d...)
	return t
}

// Flatten populates the Terraform model from SDK type
func (m *NetworkviewFederatedRealmsModel) Flatten(ctx context.Context, from *niosipam.NetworkviewFederatedRealms, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	m.Name = flex.FlattenStringPointerEmptyAsNull(from.Name)
	m.Id = flex.FlattenStringPointerEmptyAsNull(from.Id)
}
