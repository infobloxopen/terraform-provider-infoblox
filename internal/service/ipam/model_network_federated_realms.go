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

// NetworkFederatedRealmsModel is the Terraform model for NetworkFederatedRealms
type NetworkFederatedRealmsModel struct {
	Name types.String `tfsdk:"name"`
	Id   types.String `tfsdk:"id"`
}

// NetworkFederatedRealmsAttrTypes contains the attribute types for NetworkFederatedRealmsModel
var NetworkFederatedRealmsAttrTypes = map[string]attr.Type{
	"name": types.StringType,
	"id":   types.StringType,
}

// NetworkFederatedRealmsResourceSchemaAttributes contains the schema attributes for NetworkFederatedRealmsModel
var NetworkFederatedRealmsResourceSchemaAttributes = map[string]schema.Attribute{
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

// ExpandNetworkFederatedRealms converts a Terraform Object to SDK type
func ExpandNetworkFederatedRealms(ctx context.Context, o types.Object, diags *diag.Diagnostics) *niosipam.NetworkFederatedRealms {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m NetworkFederatedRealmsModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

// Expand converts the Terraform model to SDK type
func (m *NetworkFederatedRealmsModel) Expand(ctx context.Context, diags *diag.Diagnostics) *niosipam.NetworkFederatedRealms {
	if m == nil {
		return nil
	}
	to := &niosipam.NetworkFederatedRealms{
		Name: flex.ExpandStringPointerNullAsEmpty(m.Name),
		Id:   flex.ExpandStringPointerNullAsEmpty(m.Id),
	}
	return to
}

// FlattenNetworkFederatedRealms converts an SDK type to Terraform Object
func FlattenNetworkFederatedRealms(ctx context.Context, from *niosipam.NetworkFederatedRealms, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(NetworkFederatedRealmsAttrTypes)
	}
	m := &NetworkFederatedRealmsModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, NetworkFederatedRealmsAttrTypes, m)
	diags.Append(d...)
	return t
}

// Flatten populates the Terraform model from SDK type
func (m *NetworkFederatedRealmsModel) Flatten(ctx context.Context, from *niosipam.NetworkFederatedRealms, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	m.Name = flex.FlattenStringPointerEmptyAsNull(from.Name)
	m.Id = flex.FlattenStringPointerEmptyAsNull(from.Id)
}
