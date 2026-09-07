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

// DelegationServerModel is the Terraform model for DelegationServer
type DelegationServerModel struct {
	Address      types.String `tfsdk:"address"`
	Fqdn         types.String `tfsdk:"fqdn"`
	ProtocolFqdn types.String `tfsdk:"protocol_fqdn"`
}

// DelegationServerAttrTypes contains the attribute types for DelegationServerModel
var DelegationServerAttrTypes = map[string]attr.Type{
	"address":       types.StringType,
	"fqdn":          types.StringType,
	"protocol_fqdn": types.StringType,
}

// DelegationServerResourceSchemaAttributes contains the schema attributes for DelegationServerModel
var DelegationServerResourceSchemaAttributes = map[string]schema.Attribute{
	"address": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "Optional. IP Address of nameserver.  Only required when fqdn of a delegation server falls under delegation fqdn",
	},
	"fqdn": schema.StringAttribute{
		Required: true,
		Validators: []validator.String{
			customvalidator.IsValidUDDIDomainName(),
		},
		MarkdownDescription: "Required. FQDN of nameserver.",
	},
	"protocol_fqdn": schema.StringAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "FQDN of nameserver in punycode.",
	},
}

// ExpandDelegationServer converts a Terraform Object to SDK type
func ExpandDelegationServer(ctx context.Context, o types.Object, diags *diag.Diagnostics) *uddidns.DelegationServer {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m DelegationServerModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

// Expand converts the Terraform model to SDK type
func (m *DelegationServerModel) Expand(ctx context.Context, diags *diag.Diagnostics) *uddidns.DelegationServer {
	if m == nil {
		return nil
	}
	to := &uddidns.DelegationServer{
		Address:      flex.ExpandStringPointer(m.Address),
		Fqdn:         flex.ExpandString(m.Fqdn),
		ProtocolFqdn: flex.ExpandStringPointer(m.ProtocolFqdn),
	}
	return to
}

// FlattenDelegationServer converts an SDK type to Terraform Object
func FlattenDelegationServer(ctx context.Context, from *uddidns.DelegationServer, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(DelegationServerAttrTypes)
	}
	m := &DelegationServerModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, DelegationServerAttrTypes, m)
	diags.Append(d...)
	return t
}

// Flatten populates the Terraform model from SDK type
func (m *DelegationServerModel) Flatten(ctx context.Context, from *uddidns.DelegationServer, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	m.Address = flex.FlattenStringPointer(from.Address)
	m.Fqdn = flex.FlattenString(from.Fqdn)
	m.ProtocolFqdn = flex.FlattenStringPointer(from.ProtocolFqdn)
}
