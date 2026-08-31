package dhcp

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
	uddidhcp "github.com/infobloxopen/universal-ddi-go-client/ipam"
)

// HAGroupHostModel is the Terraform model for HAGroupHost
type HAGroupHostModel struct {
	Address types.String `tfsdk:"address"`
	Host    types.String `tfsdk:"host"`
	Role    types.String `tfsdk:"role"`
}

// HAGroupHostAttrTypes contains the attribute types for HAGroupHostModel
var HAGroupHostAttrTypes = map[string]attr.Type{
	"address": types.StringType,
	"host":    types.StringType,
	"role":    types.StringType,
}

// HAGroupHostResourceSchemaAttributes contains the schema attributes for HAGroupHostModel
var HAGroupHostResourceSchemaAttributes = map[string]schema.Attribute{
	"address": schema.StringAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The address on which this host listens.",
	},
	"host": schema.StringAttribute{
		Required:            true,
		MarkdownDescription: "The resource identifier.",
	},
	"role": schema.StringAttribute{
		Validators: []validator.String{
			stringvalidator.OneOf("active", "passive"),
		},
		Required:            true,
		MarkdownDescription: "The role of this host in the HA relationship: _active_ or _passive_.",
	},
}

// ExpandHAGroupHost converts a Terraform Object to SDK type
func ExpandHAGroupHost(ctx context.Context, o types.Object, diags *diag.Diagnostics) *uddidhcp.HAGroupHost {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m HAGroupHostModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

// Expand converts the Terraform model to SDK type
func (m *HAGroupHostModel) Expand(ctx context.Context, diags *diag.Diagnostics) *uddidhcp.HAGroupHost {
	if m == nil {
		return nil
	}
	to := &uddidhcp.HAGroupHost{
		Address: flex.ExpandStringPointer(m.Address),
		Host:    flex.ExpandString(m.Host),
		Role:    flex.ExpandStringPointer(m.Role),
	}
	return to
}

// FlattenHAGroupHost converts an SDK type to Terraform Object
func FlattenHAGroupHost(ctx context.Context, from *uddidhcp.HAGroupHost, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(HAGroupHostAttrTypes)
	}
	m := &HAGroupHostModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, HAGroupHostAttrTypes, m)
	diags.Append(d...)
	return t
}

// Flatten populates the Terraform model from SDK type
func (m *HAGroupHostModel) Flatten(ctx context.Context, from *uddidhcp.HAGroupHost, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	m.Address = flex.FlattenStringPointer(from.Address)
	m.Host = flex.FlattenString(from.Host)
	m.Role = flex.FlattenStringPointer(from.Role)
}
