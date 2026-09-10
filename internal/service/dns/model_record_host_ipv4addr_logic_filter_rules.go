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

// RecordHostIpv4addrLogicFilterRulesModel is the Terraform model for RecordHostIpv4addrLogicFilterRules
type RecordHostIpv4addrLogicFilterRulesModel struct {
	Filter types.String `tfsdk:"filter"`
	Type   types.String `tfsdk:"type"`
}

// RecordHostIpv4addrLogicFilterRulesAttrTypes contains the attribute types for RecordHostIpv4addrLogicFilterRulesModel
var RecordHostIpv4addrLogicFilterRulesAttrTypes = map[string]attr.Type{
	"filter": types.StringType,
	"type":   types.StringType,
}

// RecordHostIpv4addrLogicFilterRulesResourceSchemaAttributes contains the schema attributes for RecordHostIpv4addrLogicFilterRulesModel
var RecordHostIpv4addrLogicFilterRulesResourceSchemaAttributes = map[string]schema.Attribute{
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

// ExpandRecordHostIpv4addrLogicFilterRules converts a Terraform Object to SDK type
func ExpandRecordHostIpv4addrLogicFilterRules(ctx context.Context, o types.Object, diags *diag.Diagnostics) *niosdns.RecordHostIpv4addrLogicFilterRules {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m RecordHostIpv4addrLogicFilterRulesModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

// Expand converts the Terraform model to SDK type
func (m *RecordHostIpv4addrLogicFilterRulesModel) Expand(ctx context.Context, diags *diag.Diagnostics) *niosdns.RecordHostIpv4addrLogicFilterRules {
	if m == nil {
		return nil
	}
	to := &niosdns.RecordHostIpv4addrLogicFilterRules{
		Filter: flex.ExpandStringPointerNullAsEmpty(m.Filter),
		Type:   flex.ExpandStringPointerNullAsEmpty(m.Type),
	}
	return to
}

// FlattenRecordHostIpv4addrLogicFilterRules converts an SDK type to Terraform Object
func FlattenRecordHostIpv4addrLogicFilterRules(ctx context.Context, from *niosdns.RecordHostIpv4addrLogicFilterRules, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(RecordHostIpv4addrLogicFilterRulesAttrTypes)
	}
	m := &RecordHostIpv4addrLogicFilterRulesModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, RecordHostIpv4addrLogicFilterRulesAttrTypes, m)
	diags.Append(d...)
	return t
}

// Flatten populates the Terraform model from SDK type
func (m *RecordHostIpv4addrLogicFilterRulesModel) Flatten(ctx context.Context, from *niosdns.RecordHostIpv4addrLogicFilterRules, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	m.Filter = flex.FlattenStringPointerEmptyAsNull(from.Filter)
	m.Type = flex.FlattenStringPointerEmptyAsNull(from.Type)
}
