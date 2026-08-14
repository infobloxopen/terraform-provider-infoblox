package ipam

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
	uddiipam "github.com/infobloxopen/universal-ddi-go-client/ipam"
)

// DDNSUpdateBlockModel is the Terraform model for DDNSUpdateBlock
type DDNSUpdateBlockModel struct {
	DdnsDomain      types.String `tfsdk:"ddns_domain"`
	DdnsSendUpdates types.Bool   `tfsdk:"ddns_send_updates"`
}

// DDNSUpdateBlockAttrTypes contains the attribute types for DDNSUpdateBlockModel
var DDNSUpdateBlockAttrTypes = map[string]attr.Type{
	"ddns_domain":       types.StringType,
	"ddns_send_updates": types.BoolType,
}

// DDNSUpdateBlockResourceSchemaAttributes contains the schema attributes for DDNSUpdateBlockModel
var DDNSUpdateBlockResourceSchemaAttributes = map[string]schema.Attribute{
	"ddns_domain": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The domain name for DDNS.",
	},
	"ddns_send_updates": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Determines if DDNS updates are enabled at this level.",
	},
}

// ExpandDDNSUpdateBlock converts a Terraform Object to SDK type
func ExpandDDNSUpdateBlock(ctx context.Context, o types.Object, diags *diag.Diagnostics) *uddiipam.DDNSUpdateBlock {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m DDNSUpdateBlockModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

// Expand converts the Terraform model to SDK type
func (m *DDNSUpdateBlockModel) Expand(ctx context.Context, diags *diag.Diagnostics) *uddiipam.DDNSUpdateBlock {
	if m == nil {
		return nil
	}
	to := &uddiipam.DDNSUpdateBlock{
		DdnsDomain:      flex.ExpandStringPointer(m.DdnsDomain),
		DdnsSendUpdates: flex.ExpandBoolPointer(m.DdnsSendUpdates),
	}
	return to
}

// FlattenDDNSUpdateBlock converts an SDK type to Terraform Object
func FlattenDDNSUpdateBlock(ctx context.Context, from *uddiipam.DDNSUpdateBlock, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(DDNSUpdateBlockAttrTypes)
	}
	m := &DDNSUpdateBlockModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, DDNSUpdateBlockAttrTypes, m)
	diags.Append(d...)
	return t
}

// Flatten populates the Terraform model from SDK type
func (m *DDNSUpdateBlockModel) Flatten(ctx context.Context, from *uddiipam.DDNSUpdateBlock, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	m.DdnsDomain = flex.FlattenStringPointer(from.DdnsDomain)
	m.DdnsSendUpdates = flex.FlattenBoolPointer(from.DdnsSendUpdates)
}
