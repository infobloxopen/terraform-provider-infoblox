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

// NetworkcontainerLogicFilterRulesModel is the Terraform model for NetworkcontainerLogicFilterRules
type NetworkcontainerLogicFilterRulesModel struct {
	Filter types.String `tfsdk:"filter"`
	Type   types.String `tfsdk:"type"`
}

// NetworkcontainerLogicFilterRulesAttrTypes contains the attribute types for NetworkcontainerLogicFilterRulesModel
var NetworkcontainerLogicFilterRulesAttrTypes = map[string]attr.Type{
	"filter": types.StringType,
	"type":   types.StringType,
}

// NetworkcontainerLogicFilterRulesResourceSchemaAttributes contains the schema attributes for NetworkcontainerLogicFilterRulesModel
var NetworkcontainerLogicFilterRulesResourceSchemaAttributes = map[string]schema.Attribute{
	"filter": schema.StringAttribute{
		Optional: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
		},
		MarkdownDescription: "The filter name.",
	},
	"type": schema.StringAttribute{
		Optional: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
		},
		MarkdownDescription: "The filter type. Valid values are: * MAC * NAC * Option",
	},
}

// ExpandNetworkcontainerLogicFilterRules converts a Terraform Object to SDK type
func ExpandNetworkcontainerLogicFilterRules(ctx context.Context, o types.Object, diags *diag.Diagnostics) *niosipam.NetworkcontainerLogicFilterRules {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m NetworkcontainerLogicFilterRulesModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

// Expand converts the Terraform model to SDK type
func (m *NetworkcontainerLogicFilterRulesModel) Expand(ctx context.Context, diags *diag.Diagnostics) *niosipam.NetworkcontainerLogicFilterRules {
	if m == nil {
		return nil
	}
	to := &niosipam.NetworkcontainerLogicFilterRules{
		Filter: flex.ExpandStringPointerNullAsEmpty(m.Filter),
		Type:   flex.ExpandStringPointerNullAsEmpty(m.Type),
	}
	return to
}

// FlattenNetworkcontainerLogicFilterRules converts an SDK type to Terraform Object
func FlattenNetworkcontainerLogicFilterRules(ctx context.Context, from *niosipam.NetworkcontainerLogicFilterRules, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(NetworkcontainerLogicFilterRulesAttrTypes)
	}
	m := &NetworkcontainerLogicFilterRulesModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, NetworkcontainerLogicFilterRulesAttrTypes, m)
	diags.Append(d...)
	return t
}

// Flatten populates the Terraform model from SDK type
func (m *NetworkcontainerLogicFilterRulesModel) Flatten(ctx context.Context, from *niosipam.NetworkcontainerLogicFilterRules, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	m.Filter = flex.FlattenStringPointerEmptyAsNull(from.Filter)
	m.Type = flex.FlattenStringPointerEmptyAsNull(from.Type)
}
