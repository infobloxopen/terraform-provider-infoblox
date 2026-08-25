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

// Ipv6networkcontainerFederatedRealmsModel is the Terraform model for Ipv6networkcontainerFederatedRealms
type Ipv6networkcontainerFederatedRealmsModel struct {
	Name types.String `tfsdk:"name"`
	Id   types.String `tfsdk:"id"`
}

// Ipv6networkcontainerFederatedRealmsAttrTypes contains the attribute types for Ipv6networkcontainerFederatedRealmsModel
var Ipv6networkcontainerFederatedRealmsAttrTypes = map[string]attr.Type{
	"name": types.StringType,
	"id":   types.StringType,
}

// Ipv6networkcontainerFederatedRealmsResourceSchemaAttributes contains the schema attributes for Ipv6networkcontainerFederatedRealmsModel
var Ipv6networkcontainerFederatedRealmsResourceSchemaAttributes = map[string]schema.Attribute{
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

// ExpandIpv6networkcontainerFederatedRealms converts a Terraform Object to SDK type
func ExpandIpv6networkcontainerFederatedRealms(ctx context.Context, o types.Object, diags *diag.Diagnostics) *niosipam.Ipv6networkcontainerFederatedRealms {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m Ipv6networkcontainerFederatedRealmsModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

// Expand converts the Terraform model to SDK type
func (m *Ipv6networkcontainerFederatedRealmsModel) Expand(ctx context.Context, diags *diag.Diagnostics) *niosipam.Ipv6networkcontainerFederatedRealms {
	if m == nil {
		return nil
	}
	to := &niosipam.Ipv6networkcontainerFederatedRealms{
		Name: flex.ExpandStringPointerNullAsEmpty(m.Name),
		Id:   flex.ExpandStringPointerNullAsEmpty(m.Id),
	}
	return to
}

// FlattenIpv6networkcontainerFederatedRealms converts an SDK type to Terraform Object
func FlattenIpv6networkcontainerFederatedRealms(ctx context.Context, from *niosipam.Ipv6networkcontainerFederatedRealms, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(Ipv6networkcontainerFederatedRealmsAttrTypes)
	}
	m := &Ipv6networkcontainerFederatedRealmsModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, Ipv6networkcontainerFederatedRealmsAttrTypes, m)
	diags.Append(d...)
	return t
}

// Flatten populates the Terraform model from SDK type
func (m *Ipv6networkcontainerFederatedRealmsModel) Flatten(ctx context.Context, from *niosipam.Ipv6networkcontainerFederatedRealms, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	m.Name = flex.FlattenStringPointerEmptyAsNull(from.Name)
	m.Id = flex.FlattenStringPointerEmptyAsNull(from.Id)
}
