package dns

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
	customvalidator "github.com/infobloxopen/terraform-provider-infoblox/internal/validator"
	uddidns "github.com/infobloxopen/universal-ddi-go-client/dnsconfig"
)

// ForwarderModel is the Terraform model for Forwarder
type ForwarderModel struct {
	Address types.String `tfsdk:"address"`
	Fqdn    types.String `tfsdk:"fqdn"`
}

// ForwarderAttrTypes contains the attribute types for ForwarderModel
var ForwarderAttrTypes = map[string]attr.Type{
	"address": types.StringType,
	"fqdn":    types.StringType,
}

// ForwarderResourceSchemaAttributes contains the schema attributes for ForwarderModel
func ForwarderResourceSchemaAttributes(fqdnOptional bool) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"address": schema.StringAttribute{
			Required:            true,
			MarkdownDescription: "Server IP address.",
		},
		"fqdn": schema.StringAttribute{
			Optional: fqdnOptional,
			Required: !fqdnOptional,
			Validators: []validator.String{
				customvalidator.IsValidUDDIDomainName(),
			},
			MarkdownDescription: "Server FQDN.",
		},
	}
}

// ExpandForwarder converts a Terraform Object to SDK type
func ExpandForwarder(ctx context.Context, o types.Object, diags *diag.Diagnostics) *uddidns.Forwarder {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m ForwarderModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

// Expand converts the Terraform model to SDK type
func (m *ForwarderModel) Expand(ctx context.Context, diags *diag.Diagnostics) *uddidns.Forwarder {
	if m == nil {
		return nil
	}
	to := &uddidns.Forwarder{
		Address: flex.ExpandString(m.Address),
		Fqdn:    flex.ExpandStringPointer(m.Fqdn),
	}
	return to
}

// FlattenForwarder converts an SDK type to Terraform Object
func FlattenForwarder(ctx context.Context, from *uddidns.Forwarder, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(ForwarderAttrTypes)
	}
	m := &ForwarderModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, ForwarderAttrTypes, m)
	diags.Append(d...)
	return t
}

// Flatten populates the Terraform model from SDK type
func (m *ForwarderModel) Flatten(ctx context.Context, from *uddidns.Forwarder, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	m.Address = flex.FlattenString(from.Address)
	m.Fqdn = flex.FlattenStringPointerEmptyAsNull(from.Fqdn)
}
