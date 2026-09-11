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

// NsgroupForwardstubserverExternalServersModel is the Terraform model for NsgroupForwardstubserverExternalServers
type NsgroupForwardstubserverExternalServersModel struct {
	Address iptypes.IPAddress `tfsdk:"address"`
	Name    types.String      `tfsdk:"name"`
}

// NsgroupForwardstubserverExternalServersAttrTypes contains the attribute types for NsgroupForwardstubserverExternalServersModel
var NsgroupForwardstubserverExternalServersAttrTypes = map[string]attr.Type{
	"address": iptypes.IPAddressType{},
	"name":    types.StringType,
}

// NsgroupForwardstubserverExternalServersResourceSchemaAttributes contains the schema attributes for NsgroupForwardstubserverExternalServersModel
var NsgroupForwardstubserverExternalServersResourceSchemaAttributes = map[string]schema.Attribute{
	"address": schema.StringAttribute{
		Required:   true,
		CustomType: iptypes.IPAddressType{},
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
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
}

// ExpandNsgroupForwardstubserverExternalServers converts a Terraform Object to SDK type
func ExpandNsgroupForwardstubserverExternalServers(ctx context.Context, o types.Object, diags *diag.Diagnostics) *niosdns.NsgroupForwardstubserverExternalServers {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m NsgroupForwardstubserverExternalServersModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

// Expand converts the Terraform model to SDK type
func (m *NsgroupForwardstubserverExternalServersModel) Expand(ctx context.Context, diags *diag.Diagnostics) *niosdns.NsgroupForwardstubserverExternalServers {
	if m == nil {
		return nil
	}
	to := &niosdns.NsgroupForwardstubserverExternalServers{
		Address: flex.ExpandIPAddress(m.Address),
		Name:    flex.ExpandStringPointerNullAsEmpty(m.Name),
	}
	return to
}

// FlattenNsgroupForwardstubserverExternalServers converts an SDK type to Terraform Object
func FlattenNsgroupForwardstubserverExternalServers(ctx context.Context, from *niosdns.NsgroupForwardstubserverExternalServers, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(NsgroupForwardstubserverExternalServersAttrTypes)
	}
	m := &NsgroupForwardstubserverExternalServersModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, NsgroupForwardstubserverExternalServersAttrTypes, m)
	diags.Append(d...)
	return t
}

// Flatten populates the Terraform model from SDK type
func (m *NsgroupForwardstubserverExternalServersModel) Flatten(ctx context.Context, from *niosdns.NsgroupForwardstubserverExternalServers, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	m.Address = flex.FlattenIPAddress(from.Address)
	m.Name = flex.FlattenStringPointerEmptyAsNull(from.Name)
}
