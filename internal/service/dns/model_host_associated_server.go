package dns

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
	uddidns "github.com/infobloxopen/universal-ddi-go-client/dnsconfig"
)

// HostAssociatedServerModel is the Terraform model for HostAssociatedServer
type HostAssociatedServerModel struct {
	Id   types.String `tfsdk:"id"`
	Name types.String `tfsdk:"name"`
}

// HostAssociatedServerAttrTypes contains the attribute types for HostAssociatedServerModel
var HostAssociatedServerAttrTypes = map[string]attr.Type{
	"id":   types.StringType,
	"name": types.StringType,
}

// HostAssociatedServerResourceSchemaAttributes contains the schema attributes for HostAssociatedServerModel
var HostAssociatedServerResourceSchemaAttributes = map[string]schema.Attribute{
	"id": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The resource identifier.",
	},
	"name": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "DNS server name.",
	},
}

// ExpandHostAssociatedServer converts a Terraform Object to SDK type
func ExpandHostAssociatedServer(ctx context.Context, o types.Object, diags *diag.Diagnostics) *uddidns.HostAssociatedServer {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m HostAssociatedServerModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

// Expand converts the Terraform model to SDK type
func (m *HostAssociatedServerModel) Expand(ctx context.Context, diags *diag.Diagnostics) *uddidns.HostAssociatedServer {
	if m == nil {
		return nil
	}
	to := &uddidns.HostAssociatedServer{
		Id: flex.ExpandStringPointer(m.Id),
	}
	return to
}

// FlattenHostAssociatedServer converts an SDK type to Terraform Object
func FlattenHostAssociatedServer(ctx context.Context, from *uddidns.HostAssociatedServer, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(HostAssociatedServerAttrTypes)
	}
	m := &HostAssociatedServerModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, HostAssociatedServerAttrTypes, m)
	diags.Append(d...)
	return t
}

// Flatten populates the Terraform model from SDK type
func (m *HostAssociatedServerModel) Flatten(ctx context.Context, from *uddidns.HostAssociatedServer, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	m.Id = flex.FlattenStringPointer(from.Id)
	m.Name = flex.FlattenStringPointer(from.Name)
}
