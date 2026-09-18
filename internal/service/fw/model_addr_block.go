package fw

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
	uddifw "github.com/infobloxopen/universal-ddi-go-client/fw"
)

// AddrBlockModel is the Terraform model for AddrBlock
type AddrBlockModel struct {
	Address     types.String `tfsdk:"address"`
	Description types.String `tfsdk:"description"`
}

// AddrBlockAttrTypes contains the attribute types for AddrBlockModel
var AddrBlockAttrTypes = map[string]attr.Type{
	"address":     types.StringType,
	"description": types.StringType,
}

// AddrBlockResourceSchemaAttributes contains the schema attributes for AddrBlockModel
var AddrBlockResourceSchemaAttributes = map[string]schema.Attribute{
	"address": schema.StringAttribute{
		Required:            true,
		MarkdownDescription: "The CIDR of the address block.",
	},
	"description": schema.StringAttribute{
		Default:             stringdefault.StaticString(""),
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "End-user description for the address block.",
	},
}

// ExpandAddrBlock converts a Terraform Object to SDK type
func ExpandAddrBlock(ctx context.Context, o types.Object, diags *diag.Diagnostics) *uddifw.AddrBlock {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m AddrBlockModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

// Expand converts the Terraform model to SDK type
func (m *AddrBlockModel) Expand(ctx context.Context, diags *diag.Diagnostics) *uddifw.AddrBlock {
	if m == nil {
		return nil
	}
	to := &uddifw.AddrBlock{
		Address:     flex.ExpandStringPointer(m.Address),
		Description: flex.ExpandStringPointer(m.Description),
	}
	return to
}

// FlattenAddrBlock converts an SDK type to Terraform Object
func FlattenAddrBlock(ctx context.Context, from *uddifw.AddrBlock, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(AddrBlockAttrTypes)
	}
	m := &AddrBlockModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, AddrBlockAttrTypes, m)
	diags.Append(d...)
	return t
}

// Flatten populates the Terraform model from SDK type
func (m *AddrBlockModel) Flatten(ctx context.Context, from *uddifw.AddrBlock, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	m.Address = flex.FlattenStringPointer(from.Address)
	m.Description = flex.FlattenStringPointer(from.Description)
}
