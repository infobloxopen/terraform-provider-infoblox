package ipam

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
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

// NetworkLogicFilterRulesModel is the Terraform model for NetworkLogicFilterRules
type NetworkLogicFilterRulesModel struct {
	Filter types.String `tfsdk:"filter"`
	Type   types.String `tfsdk:"type"`
}

// NetworkLogicFilterRulesAttrTypes contains the attribute types for NetworkLogicFilterRulesModel
var NetworkLogicFilterRulesAttrTypes = map[string]attr.Type{
	"filter": types.StringType,
	"type":   types.StringType,
}

// NetworkLogicFilterRulesResourceSchemaAttributes contains the schema attributes for NetworkLogicFilterRulesModel
var NetworkLogicFilterRulesResourceSchemaAttributes = map[string]schema.Attribute{
	"filter": schema.StringAttribute{
		Required: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
		},
		MarkdownDescription: "The filter name.",
	},
	"type": schema.StringAttribute{
		Validators: []validator.String{
			stringvalidator.OneOf("MAC", "NAC", "Option"),
		},
		Required:            true,
		MarkdownDescription: "The filter type. Valid values are: * MAC * NAC * Option",
	},
}

// ExpandNetworkLogicFilterRules converts a Terraform Object to SDK type
func ExpandNetworkLogicFilterRules(ctx context.Context, o types.Object, diags *diag.Diagnostics) *niosipam.NetworkLogicFilterRules {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m NetworkLogicFilterRulesModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

// Expand converts the Terraform model to SDK type
func (m *NetworkLogicFilterRulesModel) Expand(ctx context.Context, diags *diag.Diagnostics) *niosipam.NetworkLogicFilterRules {
	if m == nil {
		return nil
	}
	to := &niosipam.NetworkLogicFilterRules{
		Filter: flex.ExpandStringPointerNullAsEmpty(m.Filter),
		Type:   flex.ExpandStringPointerNullAsEmpty(m.Type),
	}
	return to
}

// FlattenNetworkLogicFilterRules converts an SDK type to Terraform Object
func FlattenNetworkLogicFilterRules(ctx context.Context, from *niosipam.NetworkLogicFilterRules, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(NetworkLogicFilterRulesAttrTypes)
	}
	m := &NetworkLogicFilterRulesModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, NetworkLogicFilterRulesAttrTypes, m)
	diags.Append(d...)
	return t
}

// Flatten populates the Terraform model from SDK type
func (m *NetworkLogicFilterRulesModel) Flatten(ctx context.Context, from *niosipam.NetworkLogicFilterRules, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	m.Filter = flex.FlattenStringPointerEmptyAsNull(from.Filter)
	m.Type = flex.FlattenStringPointerEmptyAsNull(from.Type)
}
