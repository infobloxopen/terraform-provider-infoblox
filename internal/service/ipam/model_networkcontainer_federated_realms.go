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

// NetworkcontainerFederatedRealmsModel is the Terraform model for NetworkcontainerFederatedRealms
type NetworkcontainerFederatedRealmsModel struct {
	Name types.String `tfsdk:"name"`
	Id   types.String `tfsdk:"id"`
}

// NetworkcontainerFederatedRealmsAttrTypes contains the attribute types for NetworkcontainerFederatedRealmsModel
var NetworkcontainerFederatedRealmsAttrTypes = map[string]attr.Type{
	"name": types.StringType,
	"id":   types.StringType,
}

// NetworkcontainerFederatedRealmsResourceSchemaAttributes contains the schema attributes for NetworkcontainerFederatedRealmsModel
var NetworkcontainerFederatedRealmsResourceSchemaAttributes = map[string]schema.Attribute{
	"name": schema.StringAttribute{
		Optional: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
		},
		MarkdownDescription: "The federated realm name",
	},
	"id": schema.StringAttribute{
		Computed: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
		},
		MarkdownDescription: "The federated realm id",
	},
}

// ExpandNetworkcontainerFederatedRealms converts a Terraform Object to SDK type
func ExpandNetworkcontainerFederatedRealms(ctx context.Context, o types.Object, diags *diag.Diagnostics) *niosipam.NetworkcontainerFederatedRealms {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m NetworkcontainerFederatedRealmsModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

// Expand converts the Terraform model to SDK type
func (m *NetworkcontainerFederatedRealmsModel) Expand(ctx context.Context, diags *diag.Diagnostics) *niosipam.NetworkcontainerFederatedRealms {
	if m == nil {
		return nil
	}
	to := &niosipam.NetworkcontainerFederatedRealms{
		Name: flex.ExpandStringPointerNullAsEmpty(m.Name),
		Id:   flex.ExpandStringPointerNullAsEmpty(m.Id),
	}
	return to
}

// FlattenNetworkcontainerFederatedRealms converts an SDK type to Terraform Object
func FlattenNetworkcontainerFederatedRealms(ctx context.Context, from *niosipam.NetworkcontainerFederatedRealms, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(NetworkcontainerFederatedRealmsAttrTypes)
	}
	m := &NetworkcontainerFederatedRealmsModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, NetworkcontainerFederatedRealmsAttrTypes, m)
	diags.Append(d...)
	return t
}

// Flatten populates the Terraform model from SDK type
func (m *NetworkcontainerFederatedRealmsModel) Flatten(ctx context.Context, from *niosipam.NetworkcontainerFederatedRealms, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	m.Name = flex.FlattenStringPointerEmptyAsNull(from.Name)
	m.Id = flex.FlattenStringPointerEmptyAsNull(from.Id)
}
