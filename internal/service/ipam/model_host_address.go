package ipam

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	stringplanmodifier "github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/dynamicallocation"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
	uddiipam "github.com/infobloxopen/universal-ddi-go-client/ipam"
)

// HostAddressModel is the Terraform model for HostAddress
type HostAddressModel struct {
	Address           types.String `tfsdk:"address"`
	Ref               types.String `tfsdk:"ref"`
	Space             types.String `tfsdk:"space"`
	DynamicAllocation types.Object `tfsdk:"dynamic_allocation"`
}

// HostAddressAttrTypes contains the attribute types for HostAddressModel
var HostAddressAttrTypes = map[string]attr.Type{
	"address":            types.StringType,
	"ref":                types.StringType,
	"space":              types.StringType,
	"dynamic_allocation": types.ObjectType{AttrTypes: dynamicallocation.NextAvailableAddressAttrTypes},
}

// HostAddressResourceSchemaAttributes contains the schema attributes for HostAddressModel
var HostAddressResourceSchemaAttributes = map[string]schema.Attribute{
	"address": schema.StringAttribute{
		Optional: true,
		PlanModifiers: []planmodifier.String{
			stringplanmodifier.RequiresReplaceIfConfigured(),
		},
		Validators: []validator.String{
			stringvalidator.ExactlyOneOf(
				path.MatchRelative().AtParent().AtName("dynamic_allocation"),
			),
		},
		MarkdownDescription: "Field usage depends on the operation:  * For read operation, _address_ of the _Address_ corresponding to the _ref_ resource.  * For write operation, _address_ to be created if the _Address_ does not exist. Required if _ref_ is not set on write:     * If the _Address_ already exists and is already pointing to the right _Host_, the operation proceeds.     * If the _Address_ already exists and is pointing to a different _Host, the operation must abort.     * If the _Address_ already exists and is not pointing to any _Host_, it is linked to the _Host_.",
	},
	"ref": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The resource identifier.",
	},
	"space": schema.StringAttribute{
		Optional: true,
		Validators: []validator.String{
			stringvalidator.AlsoRequires(path.MatchRelative().AtParent().AtName("address")),
		},
		MarkdownDescription: "The resource identifier.",
	},
	"dynamic_allocation": schema.SingleNestedAttribute{
		Attributes:          dynamicallocation.NextAvailableAddressResourceSchemaAttributes,
		Optional:            true,
		MarkdownDescription: "Dynamically allocate the address using the NIOS next_available_address function call. Mutually exclusive with the static value field.",
	},
}

// ExpandHostAddress converts a Terraform Object to SDK type
func ExpandHostAddress(ctx context.Context, o types.Object, diags *diag.Diagnostics) *uddiipam.HostAddress {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m HostAddressModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

// Expand converts the Terraform model to SDK type
func (m *HostAddressModel) Expand(ctx context.Context, diags *diag.Diagnostics) *uddiipam.HostAddress {
	if m == nil {
		return nil
	}
	to := &uddiipam.HostAddress{
		Address: flex.ExpandStringPointer(m.Address),
		Ref:     flex.ExpandStringPointer(m.Ref),
		Space:   flex.ExpandStringPointer(m.Space),
	}
	return to
}

// FlattenHostAddress converts an SDK type to Terraform Object
func FlattenHostAddress(ctx context.Context, from *uddiipam.HostAddress, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(HostAddressAttrTypes)
	}
	m := &HostAddressModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, HostAddressAttrTypes, m)
	diags.Append(d...)
	return t
}

// Flatten populates the Terraform model from SDK type
func (m *HostAddressModel) Flatten(ctx context.Context, from *uddiipam.HostAddress, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	m.Address = flex.FlattenStringPointer(from.Address)
	m.Ref = flex.FlattenStringPointer(from.Ref)
	m.Space = flex.FlattenStringPointer(from.Space)
	if len(m.DynamicAllocation.AttributeTypes(ctx)) == 0 {
		m.DynamicAllocation = types.ObjectNull(dynamicallocation.NextAvailableAddressAttrTypes)
	}
}
