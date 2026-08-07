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

// NetworkVlansModel is the Terraform model for NetworkVlans
type NetworkVlansModel struct {
	Vlan types.Map    `tfsdk:"vlan"`
	Id   types.Int64  `tfsdk:"id"`
	Name types.String `tfsdk:"name"`
}

// NetworkVlansAttrTypes contains the attribute types for NetworkVlansModel
var NetworkVlansAttrTypes = map[string]attr.Type{
	"vlan": types.MapType{ElemType: types.StringType},
	"id":   types.Int64Type,
	"name": types.StringType,
}

// NetworkVlansResourceSchemaAttributes contains the schema attributes for NetworkVlansModel
var NetworkVlansResourceSchemaAttributes = map[string]schema.Attribute{
	"vlan": schema.MapAttribute{
		ElementType:         types.StringType,
		Optional:            true,
		MarkdownDescription: "Reference to the underlying StaticVlan object vlan.",
	},
	"id": schema.Int64Attribute{
		Computed:            true,
		MarkdownDescription: "VLAN ID value.",
	},
	"name": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
		},
		MarkdownDescription: "Name of the VLAN.",
	},
}

// ExpandNetworkVlans converts a Terraform Object to SDK type
func ExpandNetworkVlans(ctx context.Context, o types.Object, diags *diag.Diagnostics) *niosipam.NetworkVlans {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m NetworkVlansModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

// Expand converts the Terraform model to SDK type
func (m *NetworkVlansModel) Expand(ctx context.Context, diags *diag.Diagnostics) *niosipam.NetworkVlans {
	if m == nil {
		return nil
	}
	to := &niosipam.NetworkVlans{
		Vlan: flex.ExpandMapStringAny(ctx, m.Vlan, diags),
		Id:   flex.ExpandInt64Pointer(m.Id),
		Name: flex.ExpandStringPointerNullAsEmpty(m.Name),
	}
	return to
}

// FlattenNetworkVlans converts an SDK type to Terraform Object
func FlattenNetworkVlans(ctx context.Context, from *niosipam.NetworkVlans, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(NetworkVlansAttrTypes)
	}
	m := &NetworkVlansModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, NetworkVlansAttrTypes, m)
	diags.Append(d...)
	return t
}

// Flatten populates the Terraform model from SDK type
func (m *NetworkVlansModel) Flatten(ctx context.Context, from *niosipam.NetworkVlans, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	m.Vlan = flex.FlattenMapStringAny(ctx, from.Vlan, diags)
	m.Id = flex.FlattenInt64Pointer(from.Id)
	m.Name = flex.FlattenStringPointerEmptyAsNull(from.Name)
}
