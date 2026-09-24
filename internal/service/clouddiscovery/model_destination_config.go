package clouddiscovery

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	uddiclouddiscovery "github.com/infobloxopen/universal-ddi-go-client/clouddiscovery"
)

// DestinationConfigModel is the Terraform model for DestinationConfig
type DestinationConfigModel struct {
	Dns  types.Object `tfsdk:"dns"`
	Ipam types.Object `tfsdk:"ipam"`
}

// DestinationConfigAttrTypes contains the attribute types for DestinationConfigModel
var DestinationConfigAttrTypes = map[string]attr.Type{
	"dns":  types.ObjectType{AttrTypes: DNSConfigAttrTypes},
	"ipam": types.ObjectType{AttrTypes: IPAMConfigAttrTypes},
}

// DestinationConfigResourceSchemaAttributes contains the schema attributes for DestinationConfigModel
var DestinationConfigResourceSchemaAttributes = map[string]schema.Attribute{
	"dns": schema.SingleNestedAttribute{
		Attributes:          DNSConfigResourceSchemaAttributes,
		Optional:            true,
		MarkdownDescription: "",
	},
	"ipam": schema.SingleNestedAttribute{
		Attributes:          IPAMConfigResourceSchemaAttributes,
		Optional:            true,
		MarkdownDescription: "",
	},
}

// ExpandDestinationConfig converts a Terraform Object to SDK type
func ExpandDestinationConfig(ctx context.Context, o types.Object, diags *diag.Diagnostics) *uddiclouddiscovery.DestinationConfig {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m DestinationConfigModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

// Expand converts the Terraform model to SDK type
func (m *DestinationConfigModel) Expand(ctx context.Context, diags *diag.Diagnostics) *uddiclouddiscovery.DestinationConfig {
	if m == nil {
		return nil
	}
	to := &uddiclouddiscovery.DestinationConfig{
		Dns:  ExpandDNSConfig(ctx, m.Dns, diags),
		Ipam: ExpandIPAMConfig(ctx, m.Ipam, diags),
	}
	return to
}

// FlattenDestinationConfig converts an SDK type to Terraform Object
func FlattenDestinationConfig(ctx context.Context, from *uddiclouddiscovery.DestinationConfig, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(DestinationConfigAttrTypes)
	}
	m := &DestinationConfigModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, DestinationConfigAttrTypes, m)
	diags.Append(d...)
	return t
}

// Flatten populates the Terraform model from SDK type
func (m *DestinationConfigModel) Flatten(ctx context.Context, from *uddiclouddiscovery.DestinationConfig, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	m.Dns = FlattenDNSConfig(ctx, from.Dns, diags)
	m.Ipam = FlattenIPAMConfig(ctx, from.Ipam, diags)
}
