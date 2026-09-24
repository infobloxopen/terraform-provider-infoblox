package clouddiscovery

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
	customvalidator "github.com/infobloxopen/terraform-provider-infoblox/internal/validator"
	uddiclouddiscovery "github.com/infobloxopen/universal-ddi-go-client/clouddiscovery"
)

// AdditionalConfigModel is the Terraform model for AdditionalConfig
type AdditionalConfigModel struct {
	ExcludedAccounts      types.List   `tfsdk:"excluded_accounts"`
	ForwardZoneEnabled    types.Bool   `tfsdk:"forward_zone_enabled"`
	InternalRangesEnabled types.Bool   `tfsdk:"internal_ranges_enabled"`
	ObjectType            types.Object `tfsdk:"object_type"`
}

// AdditionalConfigAttrTypes contains the attribute types for AdditionalConfigModel
var AdditionalConfigAttrTypes = map[string]attr.Type{
	"excluded_accounts":       types.ListType{ElemType: types.StringType},
	"forward_zone_enabled":    types.BoolType,
	"internal_ranges_enabled": types.BoolType,
	"object_type":             types.ObjectType{AttrTypes: ObjectTypeAttrTypes},
}

// AdditionalConfigResourceSchemaAttributes contains the schema attributes for AdditionalConfigModel
var AdditionalConfigResourceSchemaAttributes = map[string]schema.Attribute{
	"excluded_accounts": schema.ListAttribute{
		ElementType: types.StringType,
		Optional:    true,
		Computed:    true,
		Validators: []validator.List{
			customvalidator.ListNotEmpty(),
		},
		MarkdownDescription: "",
	},
	"forward_zone_enabled": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "",
	},
	"internal_ranges_enabled": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "",
	},
	"object_type": schema.SingleNestedAttribute{
		Attributes:          ObjectTypeResourceSchemaAttributes,
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "",
	},
}

// ExpandAdditionalConfig converts a Terraform Object to SDK type
func ExpandAdditionalConfig(ctx context.Context, o types.Object, diags *diag.Diagnostics) *uddiclouddiscovery.AdditionalConfig {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m AdditionalConfigModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

// Expand converts the Terraform model to SDK type
func (m *AdditionalConfigModel) Expand(ctx context.Context, diags *diag.Diagnostics) *uddiclouddiscovery.AdditionalConfig {
	if m == nil {
		return nil
	}
	to := &uddiclouddiscovery.AdditionalConfig{
		ExcludedAccounts:      flex.ExpandFrameworkListString(ctx, m.ExcludedAccounts, diags),
		ForwardZoneEnabled:    flex.ExpandBoolPointer(m.ForwardZoneEnabled),
		InternalRangesEnabled: flex.ExpandBoolPointer(m.InternalRangesEnabled),
		ObjectType:            ExpandObjectType(ctx, m.ObjectType, diags),
	}
	return to
}

// FlattenAdditionalConfig converts an SDK type to Terraform Object
func FlattenAdditionalConfig(ctx context.Context, from *uddiclouddiscovery.AdditionalConfig, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(AdditionalConfigAttrTypes)
	}
	m := &AdditionalConfigModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, AdditionalConfigAttrTypes, m)
	diags.Append(d...)
	return t
}

// Flatten populates the Terraform model from SDK type
func (m *AdditionalConfigModel) Flatten(ctx context.Context, from *uddiclouddiscovery.AdditionalConfig, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	m.ExcludedAccounts = flex.FlattenFrameworkListString(ctx, from.ExcludedAccounts, diags)
	m.ForwardZoneEnabled = flex.FlattenBoolPointer(from.ForwardZoneEnabled)
	m.InternalRangesEnabled = flex.FlattenBoolPointer(from.InternalRangesEnabled)
	m.ObjectType = FlattenObjectType(ctx, from.ObjectType, diags)
}
