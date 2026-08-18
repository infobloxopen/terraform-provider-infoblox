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

// NetworkcontainersubscribesettingsMappedEaAttributesModel is the Terraform model for NetworkcontainersubscribesettingsMappedEaAttributes
type NetworkcontainersubscribesettingsMappedEaAttributesModel struct {
	Name     types.String `tfsdk:"name"`
	MappedEa types.String `tfsdk:"mapped_ea"`
}

// NetworkcontainersubscribesettingsMappedEaAttributesAttrTypes contains the attribute types for NetworkcontainersubscribesettingsMappedEaAttributesModel
var NetworkcontainersubscribesettingsMappedEaAttributesAttrTypes = map[string]attr.Type{
	"name":      types.StringType,
	"mapped_ea": types.StringType,
}

// NetworkcontainersubscribesettingsMappedEaAttributesResourceSchemaAttributes contains the schema attributes for NetworkcontainersubscribesettingsMappedEaAttributesModel
var NetworkcontainersubscribesettingsMappedEaAttributesResourceSchemaAttributes = map[string]schema.Attribute{
	"name": schema.StringAttribute{
		Validators: []validator.String{
			stringvalidator.OneOf("ACCOUNT_SESSION_ID", "AUDIT_SESSION_ID", "EPS_STATUS", "IP_ADDRESS", "MAC", "NAS_IP_ADDRESS", "NAS_PORT_ID", "POSTURE_STATUS", "POSTURE_TIMESTAMP"),
		},
		Optional:            true,
		MarkdownDescription: "The Cisco ISE attribute name that is enabled for publishsing from a Cisco ISE endpoint.",
	},
	"mapped_ea": schema.StringAttribute{
		Optional: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
		},
		MarkdownDescription: "The name of the extensible attribute definition object the Cisco ISE attribute that is enabled for subscription is mapped on.",
	},
}

// ExpandNetworkcontainersubscribesettingsMappedEaAttributes converts a Terraform Object to SDK type
func ExpandNetworkcontainersubscribesettingsMappedEaAttributes(ctx context.Context, o types.Object, diags *diag.Diagnostics) *niosipam.NetworkcontainersubscribesettingsMappedEaAttributes {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m NetworkcontainersubscribesettingsMappedEaAttributesModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

// Expand converts the Terraform model to SDK type
func (m *NetworkcontainersubscribesettingsMappedEaAttributesModel) Expand(ctx context.Context, diags *diag.Diagnostics) *niosipam.NetworkcontainersubscribesettingsMappedEaAttributes {
	if m == nil {
		return nil
	}
	to := &niosipam.NetworkcontainersubscribesettingsMappedEaAttributes{
		Name:     flex.ExpandStringPointerNullAsEmpty(m.Name),
		MappedEa: flex.ExpandStringPointerNullAsEmpty(m.MappedEa),
	}
	return to
}

// FlattenNetworkcontainersubscribesettingsMappedEaAttributes converts an SDK type to Terraform Object
func FlattenNetworkcontainersubscribesettingsMappedEaAttributes(ctx context.Context, from *niosipam.NetworkcontainersubscribesettingsMappedEaAttributes, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(NetworkcontainersubscribesettingsMappedEaAttributesAttrTypes)
	}
	m := &NetworkcontainersubscribesettingsMappedEaAttributesModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, NetworkcontainersubscribesettingsMappedEaAttributesAttrTypes, m)
	diags.Append(d...)
	return t
}

// Flatten populates the Terraform model from SDK type
func (m *NetworkcontainersubscribesettingsMappedEaAttributesModel) Flatten(ctx context.Context, from *niosipam.NetworkcontainersubscribesettingsMappedEaAttributes, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	m.Name = flex.FlattenStringPointerEmptyAsNull(from.Name)
	m.MappedEa = flex.FlattenStringPointerEmptyAsNull(from.MappedEa)
}
