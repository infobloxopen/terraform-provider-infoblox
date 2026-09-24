package clouddiscovery

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
	internaltypes "github.com/infobloxopen/terraform-provider-infoblox/internal/types"
	customvalidator "github.com/infobloxopen/terraform-provider-infoblox/internal/validator"
	uddiclouddiscovery "github.com/infobloxopen/universal-ddi-go-client/clouddiscovery"
)

// ZoneFilterModel is the Terraform model for ZoneFilter
type ZoneFilterModel struct {
	Action    types.String                     `tfsdk:"action"`
	Wildcards internaltypes.UnorderedListValue `tfsdk:"wildcards"`
}

// ZoneFilterAttrTypes contains the attribute types for ZoneFilterModel
var ZoneFilterAttrTypes = map[string]attr.Type{
	"action":    types.StringType,
	"wildcards": internaltypes.UnorderedListOfStringType,
}

// ZoneFilterResourceSchemaAttributes contains the schema attributes for ZoneFilterModel
var ZoneFilterResourceSchemaAttributes = map[string]schema.Attribute{
	"action": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "",
	},
	"wildcards": schema.ListAttribute{
		ElementType: types.StringType,
		Optional:    true,
		CustomType:  internaltypes.UnorderedListOfStringType,
		Validators: []validator.List{
			customvalidator.ListNotEmpty(),
		},
		MarkdownDescription: "",
	},
}

// ExpandZoneFilter converts a Terraform Object to SDK type
func ExpandZoneFilter(ctx context.Context, o types.Object, diags *diag.Diagnostics) *uddiclouddiscovery.ZoneFilter {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m ZoneFilterModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

// Expand converts the Terraform model to SDK type
func (m *ZoneFilterModel) Expand(ctx context.Context, diags *diag.Diagnostics) *uddiclouddiscovery.ZoneFilter {
	if m == nil {
		return nil
	}
	to := &uddiclouddiscovery.ZoneFilter{
		Action:    flex.ExpandStringPointer(m.Action),
		Wildcards: flex.ExpandFrameworkListString(ctx, m.Wildcards, diags),
	}
	return to
}

// FlattenZoneFilter converts an SDK type to Terraform Object
func FlattenZoneFilter(ctx context.Context, from *uddiclouddiscovery.ZoneFilter, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(ZoneFilterAttrTypes)
	}
	m := &ZoneFilterModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, ZoneFilterAttrTypes, m)
	diags.Append(d...)
	return t
}

// Flatten populates the Terraform model from SDK type
func (m *ZoneFilterModel) Flatten(ctx context.Context, from *uddiclouddiscovery.ZoneFilter, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	m.Action = flex.FlattenStringPointer(from.Action)
	m.Wildcards = flex.FlattenFrameworkUnorderedListString(ctx, from.Wildcards, diags)
}
