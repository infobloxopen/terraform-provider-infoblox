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

// NetworkcontainerZoneAssociationsModel is the Terraform model for NetworkcontainerZoneAssociations
type NetworkcontainerZoneAssociationsModel struct {
	Fqdn      types.String `tfsdk:"fqdn"`
	IsDefault types.Bool   `tfsdk:"is_default"`
	View      types.String `tfsdk:"view"`
}

// NetworkcontainerZoneAssociationsAttrTypes contains the attribute types for NetworkcontainerZoneAssociationsModel
var NetworkcontainerZoneAssociationsAttrTypes = map[string]attr.Type{
	"fqdn":       types.StringType,
	"is_default": types.BoolType,
	"view":       types.StringType,
}

// NetworkcontainerZoneAssociationsResourceSchemaAttributes contains the schema attributes for NetworkcontainerZoneAssociationsModel
var NetworkcontainerZoneAssociationsResourceSchemaAttributes = map[string]schema.Attribute{
	"fqdn": schema.StringAttribute{
		Optional: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
		},
		MarkdownDescription: "The FQDN of the authoritative forward zone.",
	},
	"is_default": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "True if this is the default zone.",
	},
	"view": schema.StringAttribute{
		Optional: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
		},
		MarkdownDescription: "The view to which the zone belongs. If a view is not specified, the default view is used.",
	},
}

// ExpandNetworkcontainerZoneAssociations converts a Terraform Object to SDK type
func ExpandNetworkcontainerZoneAssociations(ctx context.Context, o types.Object, diags *diag.Diagnostics) *niosipam.NetworkcontainerZoneAssociations {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m NetworkcontainerZoneAssociationsModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

// Expand converts the Terraform model to SDK type
func (m *NetworkcontainerZoneAssociationsModel) Expand(ctx context.Context, diags *diag.Diagnostics) *niosipam.NetworkcontainerZoneAssociations {
	if m == nil {
		return nil
	}
	to := &niosipam.NetworkcontainerZoneAssociations{
		Fqdn:      flex.ExpandStringPointerNullAsEmpty(m.Fqdn),
		IsDefault: flex.ExpandBoolPointer(m.IsDefault),
		View:      flex.ExpandStringPointerNullAsEmpty(m.View),
	}
	return to
}

// FlattenNetworkcontainerZoneAssociations converts an SDK type to Terraform Object
func FlattenNetworkcontainerZoneAssociations(ctx context.Context, from *niosipam.NetworkcontainerZoneAssociations, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(NetworkcontainerZoneAssociationsAttrTypes)
	}
	m := &NetworkcontainerZoneAssociationsModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, NetworkcontainerZoneAssociationsAttrTypes, m)
	diags.Append(d...)
	return t
}

// Flatten populates the Terraform model from SDK type
func (m *NetworkcontainerZoneAssociationsModel) Flatten(ctx context.Context, from *niosipam.NetworkcontainerZoneAssociations, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	m.Fqdn = flex.FlattenStringPointerEmptyAsNull(from.Fqdn)
	m.IsDefault = flex.FlattenBoolPointer(from.IsDefault)
	m.View = flex.FlattenStringPointerEmptyAsNull(from.View)
}
