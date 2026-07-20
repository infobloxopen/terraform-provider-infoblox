package dns

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
	uddidns "github.com/infobloxopen/universal-ddi-go-client/dnsconfig"
)

// ECSBlockModel is the Terraform model for ECSBlock
type ECSBlockModel struct {
	EcsEnabled    types.Bool  `tfsdk:"ecs_enabled"`
	EcsForwarding types.Bool  `tfsdk:"ecs_forwarding"`
	EcsPrefixV4   types.Int64 `tfsdk:"ecs_prefix_v4"`
	EcsPrefixV6   types.Int64 `tfsdk:"ecs_prefix_v6"`
	EcsZones      types.List  `tfsdk:"ecs_zones"`
}

// ECSBlockAttrTypes contains the attribute types for ECSBlockModel
var ECSBlockAttrTypes = map[string]attr.Type{
	"ecs_enabled":    types.BoolType,
	"ecs_forwarding": types.BoolType,
	"ecs_prefix_v4":  types.Int64Type,
	"ecs_prefix_v6":  types.Int64Type,
	"ecs_zones":      types.ListType{ElemType: types.ObjectType{AttrTypes: ECSZoneAttrTypes}},
}

// ECSBlockResourceSchemaAttributes contains the schema attributes for ECSBlockModel
var ECSBlockResourceSchemaAttributes = map[string]schema.Attribute{
	"ecs_enabled": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Optional. Field config for _ecs_enabled_ field.",
	},
	"ecs_forwarding": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Optional. Field config for _ecs_forwarding_ field.",
	},
	"ecs_prefix_v4": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "Optional. Field config for _ecs_prefix_v4_ field.",
	},
	"ecs_prefix_v6": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "Optional. Field config for _ecs_prefix_v6_ field.",
	},
	"ecs_zones": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: ECSZoneResourceSchemaAttributes,
		},
		Optional: true,
		Validators: []validator.List{
			listvalidator.SizeAtLeast(1),
		},
		MarkdownDescription: "Optional. Field config for _ecs_zones_ field.",
	},
}

// ExpandECSBlock converts a Terraform Object to SDK type
func ExpandECSBlock(ctx context.Context, o types.Object, diags *diag.Diagnostics) *uddidns.ECSBlock {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m ECSBlockModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

// Expand converts the Terraform model to SDK type
func (m *ECSBlockModel) Expand(ctx context.Context, diags *diag.Diagnostics) *uddidns.ECSBlock {
	if m == nil {
		return nil
	}
	to := &uddidns.ECSBlock{
		EcsEnabled:    flex.ExpandBoolPointer(m.EcsEnabled),
		EcsForwarding: flex.ExpandBoolPointer(m.EcsForwarding),
		EcsPrefixV4:   flex.ExpandInt64Pointer(m.EcsPrefixV4),
		EcsPrefixV6:   flex.ExpandInt64Pointer(m.EcsPrefixV6),
		EcsZones:      flex.ExpandFrameworkListNestedBlock(ctx, m.EcsZones, diags, ExpandECSZone),
	}
	return to
}

// FlattenECSBlock converts an SDK type to Terraform Object
func FlattenECSBlock(ctx context.Context, from *uddidns.ECSBlock, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(ECSBlockAttrTypes)
	}
	m := &ECSBlockModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, ECSBlockAttrTypes, m)
	diags.Append(d...)
	return t
}

// Flatten populates the Terraform model from SDK type
func (m *ECSBlockModel) Flatten(ctx context.Context, from *uddidns.ECSBlock, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	m.EcsEnabled = flex.FlattenBoolPointer(from.EcsEnabled)
	m.EcsForwarding = flex.FlattenBoolPointer(from.EcsForwarding)
	m.EcsPrefixV4 = flex.FlattenInt64Pointer(from.EcsPrefixV4)
	m.EcsPrefixV6 = flex.FlattenInt64Pointer(from.EcsPrefixV6)
	m.EcsZones = flex.FlattenFrameworkListNestedBlock(ctx, from.EcsZones, ECSZoneAttrTypes, diags, FlattenECSZone)
}
