package dns

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/hashicorp/terraform-plugin-framework-nettypes/iptypes"
	niosdns "github.com/infobloxopen/infoblox-nios-go-client/dns"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
	customvalidator "github.com/infobloxopen/terraform-provider-infoblox/internal/validator"
)

// NsgroupgridsecondariesPreferredPrimariesModel is the Terraform model for NsgroupgridsecondariesPreferredPrimaries
type NsgroupgridsecondariesPreferredPrimariesModel struct {
	Address iptypes.IPAddress `tfsdk:"address"`
	Name    types.String      `tfsdk:"name"`
	Stealth types.Bool        `tfsdk:"stealth"`
}

// NsgroupgridsecondariesPreferredPrimariesAttrTypes contains the attribute types for NsgroupgridsecondariesPreferredPrimariesModel
var NsgroupgridsecondariesPreferredPrimariesAttrTypes = map[string]attr.Type{
	"address": iptypes.IPAddressType{},
	"name":    types.StringType,
	"stealth": types.BoolType,
}

// NsgroupgridsecondariesPreferredPrimariesResourceSchemaAttributes contains the schema attributes for NsgroupgridsecondariesPreferredPrimariesModel
var NsgroupgridsecondariesPreferredPrimariesResourceSchemaAttributes = map[string]schema.Attribute{
	"address": schema.StringAttribute{
		Required:   true,
		CustomType: iptypes.IPAddressType{},
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
			customvalidator.ValidateTrimmedString(),
		},
		MarkdownDescription: "The IPv4 Address or IPv6 Address of the server.",
	},
	"name": schema.StringAttribute{
		Required: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
			customvalidator.ValidateTrimmedString(),
		},
		MarkdownDescription: "A resolvable domain name for the external DNS server.",
	},
	"stealth": schema.BoolAttribute{
		Computed:            true,
		MarkdownDescription: "Set this flag to hide the NS record for the primary name server from DNS queries.",
	},
}

// ExpandNsgroupgridsecondariesPreferredPrimaries converts a Terraform Object to SDK type
func ExpandNsgroupgridsecondariesPreferredPrimaries(ctx context.Context, o types.Object, diags *diag.Diagnostics) *niosdns.NsgroupgridsecondariesPreferredPrimaries {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m NsgroupgridsecondariesPreferredPrimariesModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

// Expand converts the Terraform model to SDK type
func (m *NsgroupgridsecondariesPreferredPrimariesModel) Expand(ctx context.Context, diags *diag.Diagnostics) *niosdns.NsgroupgridsecondariesPreferredPrimaries {
	if m == nil {
		return nil
	}
	to := &niosdns.NsgroupgridsecondariesPreferredPrimaries{
		Address: flex.ExpandIPAddress(m.Address),
		Name:    flex.ExpandStringPointerNullAsEmpty(m.Name),
		Stealth: flex.ExpandBoolPointer(m.Stealth),
	}
	return to
}

// FlattenNsgroupgridsecondariesPreferredPrimaries converts an SDK type to Terraform Object
func FlattenNsgroupgridsecondariesPreferredPrimaries(ctx context.Context, from *niosdns.NsgroupgridsecondariesPreferredPrimaries, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(NsgroupgridsecondariesPreferredPrimariesAttrTypes)
	}
	m := &NsgroupgridsecondariesPreferredPrimariesModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, NsgroupgridsecondariesPreferredPrimariesAttrTypes, m)
	diags.Append(d...)
	return t
}

// Flatten populates the Terraform model from SDK type
func (m *NsgroupgridsecondariesPreferredPrimariesModel) Flatten(ctx context.Context, from *niosdns.NsgroupgridsecondariesPreferredPrimaries, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	m.Address = flex.FlattenIPAddress(from.Address)
	m.Name = flex.FlattenStringPointerEmptyAsNull(from.Name)
	m.Stealth = flex.FlattenBoolPointer(from.Stealth)
}
