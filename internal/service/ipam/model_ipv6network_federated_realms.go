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

// Ipv6networkFederatedRealmsModel is the Terraform model for Ipv6networkFederatedRealms
type Ipv6networkFederatedRealmsModel struct {
	Name types.String `tfsdk:"name"`
	Id   types.String `tfsdk:"id"`
}

// Ipv6networkFederatedRealmsAttrTypes contains the attribute types for Ipv6networkFederatedRealmsModel
var Ipv6networkFederatedRealmsAttrTypes = map[string]attr.Type{
	"name": types.StringType,
	"id":   types.StringType,
}

// Ipv6networkFederatedRealmsResourceSchemaAttributes contains the schema attributes for Ipv6networkFederatedRealmsModel
var Ipv6networkFederatedRealmsResourceSchemaAttributes = map[string]schema.Attribute{
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

// ExpandIpv6networkFederatedRealms converts a Terraform Object to SDK type
func ExpandIpv6networkFederatedRealms(ctx context.Context, o types.Object, diags *diag.Diagnostics) *niosipam.Ipv6networkFederatedRealms {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m Ipv6networkFederatedRealmsModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

// Expand converts the Terraform model to SDK type
func (m *Ipv6networkFederatedRealmsModel) Expand(ctx context.Context, diags *diag.Diagnostics) *niosipam.Ipv6networkFederatedRealms {
	if m == nil {
		return nil
	}
	to := &niosipam.Ipv6networkFederatedRealms{
		Name: flex.ExpandStringPointerNullAsEmpty(m.Name),
		Id:   flex.ExpandStringPointerNullAsEmpty(m.Id),
	}
	return to
}

// FlattenIpv6networkFederatedRealms converts an SDK type to Terraform Object
func FlattenIpv6networkFederatedRealms(ctx context.Context, from *niosipam.Ipv6networkFederatedRealms, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(Ipv6networkFederatedRealmsAttrTypes)
	}
	m := &Ipv6networkFederatedRealmsModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, Ipv6networkFederatedRealmsAttrTypes, m)
	diags.Append(d...)
	return t
}

// Flatten populates the Terraform model from SDK type
func (m *Ipv6networkFederatedRealmsModel) Flatten(ctx context.Context, from *niosipam.Ipv6networkFederatedRealms, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	m.Name = flex.FlattenStringPointerEmptyAsNull(from.Name)
	m.Id = flex.FlattenStringPointerEmptyAsNull(from.Id)
}
