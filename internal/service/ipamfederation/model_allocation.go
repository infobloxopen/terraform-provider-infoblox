package ipamfederation

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
	uddiipamfederation "github.com/infobloxopen/universal-ddi-go-client/ipamfederation"
)

// AllocationModel is the Terraform model for Allocation
type AllocationModel struct {
	Allocated                types.Int64 `tfsdk:"allocated"`
	Delegated                types.Int64 `tfsdk:"delegated"`
	ForwardLookingDelegation types.Int64 `tfsdk:"forward_looking_delegation"`
	Overlapping              types.Int64 `tfsdk:"overlapping"`
	Reserved                 types.Int64 `tfsdk:"reserved"`
}

// AllocationAttrTypes contains the attribute types for AllocationModel
var AllocationAttrTypes = map[string]attr.Type{
	"allocated":                  types.Int64Type,
	"delegated":                  types.Int64Type,
	"forward_looking_delegation": types.Int64Type,
	"overlapping":                types.Int64Type,
	"reserved":                   types.Int64Type,
}

// AllocationResourceSchemaAttributes contains the schema attributes for AllocationModel
var AllocationResourceSchemaAttributes = map[string]schema.Attribute{
	"allocated": schema.Int64Attribute{
		Computed:            true,
		MarkdownDescription: "Percent of total space allocated.",
	},
	"delegated": schema.Int64Attribute{
		Computed:            true,
		MarkdownDescription: "Percent of total space delegated.",
	},
	"forward_looking_delegation": schema.Int64Attribute{
		Computed:            true,
		MarkdownDescription: "Percent of total space in forward looking delegation blocks.",
	},
	"overlapping": schema.Int64Attribute{
		Computed:            true,
		MarkdownDescription: "Percent of total space in overlapping blocks.",
	},
	"reserved": schema.Int64Attribute{
		Computed:            true,
		MarkdownDescription: "Percent of total space reserved.",
	},
}

// ExpandAllocation converts a Terraform Object to SDK type
func ExpandAllocation(ctx context.Context, o types.Object, diags *diag.Diagnostics) *uddiipamfederation.Allocation {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m AllocationModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

// Expand converts the Terraform model to SDK type
func (m *AllocationModel) Expand(ctx context.Context, diags *diag.Diagnostics) *uddiipamfederation.Allocation {
	if m == nil {
		return nil
	}
	to := &uddiipamfederation.Allocation{
		Allocated:                flex.ExpandInt64Pointer(m.Allocated),
		Delegated:                flex.ExpandInt64Pointer(m.Delegated),
		ForwardLookingDelegation: flex.ExpandInt64Pointer(m.ForwardLookingDelegation),
		Overlapping:              flex.ExpandInt64Pointer(m.Overlapping),
		Reserved:                 flex.ExpandInt64Pointer(m.Reserved),
	}
	return to
}

// FlattenAllocation converts an SDK type to Terraform Object
func FlattenAllocation(ctx context.Context, from *uddiipamfederation.Allocation, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(AllocationAttrTypes)
	}
	m := &AllocationModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, AllocationAttrTypes, m)
	diags.Append(d...)
	return t
}

// Flatten populates the Terraform model from SDK type
func (m *AllocationModel) Flatten(ctx context.Context, from *uddiipamfederation.Allocation, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	m.Allocated = flex.FlattenInt64Pointer(from.Allocated)
	m.Delegated = flex.FlattenInt64Pointer(from.Delegated)
	m.ForwardLookingDelegation = flex.FlattenInt64Pointer(from.ForwardLookingDelegation)
	m.Overlapping = flex.FlattenInt64Pointer(from.Overlapping)
	m.Reserved = flex.FlattenInt64Pointer(from.Reserved)
}
