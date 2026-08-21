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

// NsgroupgridprimaryPreferredPrimariesModel is the Terraform model for NsgroupgridprimaryPreferredPrimaries
type NsgroupgridprimaryPreferredPrimariesModel struct {
	Address iptypes.IPAddress `tfsdk:"address"`
	Name    types.String      `tfsdk:"name"`
	Stealth types.Bool        `tfsdk:"stealth"`
}

// NsgroupgridprimaryPreferredPrimariesAttrTypes contains the attribute types for NsgroupgridprimaryPreferredPrimariesModel
var NsgroupgridprimaryPreferredPrimariesAttrTypes = map[string]attr.Type{
	"address": iptypes.IPAddressType{},
	"name":    types.StringType,
	"stealth": types.BoolType,
}

// NsgroupgridprimaryPreferredPrimariesResourceSchemaAttributes contains the schema attributes for NsgroupgridprimaryPreferredPrimariesModel
var NsgroupgridprimaryPreferredPrimariesResourceSchemaAttributes = map[string]schema.Attribute{
	"address": schema.StringAttribute{
		Computed:   true,
		CustomType: iptypes.IPAddressType{},
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
		},
		MarkdownDescription: "The IPv4 Address or IPv6 Address of the server.",
	},
	"name": schema.StringAttribute{
		Computed: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
		},
		MarkdownDescription: "A resolvable domain name for the external DNS server.",
	},
	"stealth": schema.BoolAttribute{
		Computed:            true,
		MarkdownDescription: "Set this flag to hide the NS record for the primary name server from DNS queries.",
	},
}

// ExpandNsgroupgridprimaryPreferredPrimaries converts a Terraform Object to SDK type
func ExpandNsgroupgridprimaryPreferredPrimaries(ctx context.Context, o types.Object, diags *diag.Diagnostics) *niosdns.NsgroupgridprimaryPreferredPrimaries {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m NsgroupgridprimaryPreferredPrimariesModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

// Expand converts the Terraform model to SDK type
func (m *NsgroupgridprimaryPreferredPrimariesModel) Expand(ctx context.Context, diags *diag.Diagnostics) *niosdns.NsgroupgridprimaryPreferredPrimaries {
	if m == nil {
		return nil
	}
	to := &niosdns.NsgroupgridprimaryPreferredPrimaries{
		Address: flex.ExpandIPAddress(m.Address),
		Name:    flex.ExpandStringPointerNullAsEmpty(m.Name),
		Stealth: flex.ExpandBoolPointer(m.Stealth),
	}
	return to
}

// FlattenNsgroupgridprimaryPreferredPrimaries converts an SDK type to Terraform Object
func FlattenNsgroupgridprimaryPreferredPrimaries(ctx context.Context, from *niosdns.NsgroupgridprimaryPreferredPrimaries, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(NsgroupgridprimaryPreferredPrimariesAttrTypes)
	}
	m := &NsgroupgridprimaryPreferredPrimariesModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, NsgroupgridprimaryPreferredPrimariesAttrTypes, m)
	diags.Append(d...)
	return t
}

// Flatten populates the Terraform model from SDK type
func (m *NsgroupgridprimaryPreferredPrimariesModel) Flatten(ctx context.Context, from *niosdns.NsgroupgridprimaryPreferredPrimaries, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	m.Address = flex.FlattenIPAddress(from.Address)
	m.Name = flex.FlattenStringPointerEmptyAsNull(from.Name)
	m.Stealth = flex.FlattenBoolPointer(from.Stealth)
}
