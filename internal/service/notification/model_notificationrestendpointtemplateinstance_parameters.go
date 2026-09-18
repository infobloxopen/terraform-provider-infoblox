package notification

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	niosnotification "github.com/infobloxopen/infoblox-nios-go-client/notification"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
	customvalidator "github.com/infobloxopen/terraform-provider-infoblox/internal/validator"
)

// NotificationrestendpointtemplateinstanceParametersModel is the Terraform model for NotificationrestendpointtemplateinstanceParameters
type NotificationrestendpointtemplateinstanceParametersModel struct {
	Name         types.String `tfsdk:"name"`
	Value        types.String `tfsdk:"value"`
	DefaultValue types.String `tfsdk:"default_value"`
	Syntax       types.String `tfsdk:"syntax"`
}

// NotificationrestendpointtemplateinstanceParametersAttrTypes contains the attribute types for NotificationrestendpointtemplateinstanceParametersModel
var NotificationrestendpointtemplateinstanceParametersAttrTypes = map[string]attr.Type{
	"name":          types.StringType,
	"value":         types.StringType,
	"default_value": types.StringType,
	"syntax":        types.StringType,
}

// NotificationrestendpointtemplateinstanceParametersResourceSchemaAttributes contains the schema attributes for NotificationrestendpointtemplateinstanceParametersModel
var NotificationrestendpointtemplateinstanceParametersResourceSchemaAttributes = map[string]schema.Attribute{
	"name": schema.StringAttribute{
		Required: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
		},
		MarkdownDescription: "The name of the REST API template parameter.",
	},
	"value": schema.StringAttribute{
		Optional: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
		},
		MarkdownDescription: "The value of the REST API template parameter.",
	},
	"default_value": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
		},
		MarkdownDescription: "The default value of the REST API template parameter.",
	},
	"syntax": schema.StringAttribute{
		Validators: []validator.String{
			stringvalidator.OneOf("STR", "BOOL", "INT"),
		},
		Required:            true,
		MarkdownDescription: "The syntax of the REST API template parameter.",
	},
}

// ExpandNotificationrestendpointtemplateinstanceParameters converts a Terraform Object to SDK type
func ExpandNotificationrestendpointtemplateinstanceParameters(ctx context.Context, o types.Object, diags *diag.Diagnostics) *niosnotification.NotificationrestendpointtemplateinstanceParameters {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m NotificationrestendpointtemplateinstanceParametersModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

// Expand converts the Terraform model to SDK type
func (m *NotificationrestendpointtemplateinstanceParametersModel) Expand(ctx context.Context, diags *diag.Diagnostics) *niosnotification.NotificationrestendpointtemplateinstanceParameters {
	if m == nil {
		return nil
	}
	to := &niosnotification.NotificationrestendpointtemplateinstanceParameters{
		Name:         flex.ExpandStringPointerNullAsEmpty(m.Name),
		Value:        flex.ExpandStringPointerNullAsEmpty(m.Value),
		DefaultValue: flex.ExpandStringPointerNullAsEmpty(m.DefaultValue),
		Syntax:       flex.ExpandStringPointerNullAsEmpty(m.Syntax),
	}
	return to
}

// FlattenNotificationrestendpointtemplateinstanceParameters converts an SDK type to Terraform Object
func FlattenNotificationrestendpointtemplateinstanceParameters(ctx context.Context, from *niosnotification.NotificationrestendpointtemplateinstanceParameters, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(NotificationrestendpointtemplateinstanceParametersAttrTypes)
	}
	m := &NotificationrestendpointtemplateinstanceParametersModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, NotificationrestendpointtemplateinstanceParametersAttrTypes, m)
	diags.Append(d...)
	return t
}

// Flatten populates the Terraform model from SDK type
func (m *NotificationrestendpointtemplateinstanceParametersModel) Flatten(ctx context.Context, from *niosnotification.NotificationrestendpointtemplateinstanceParameters, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	m.Name = flex.FlattenStringPointerEmptyAsNull(from.Name)
	m.Value = flex.FlattenStringPointerEmptyAsNull(from.Value)
	m.DefaultValue = flex.FlattenStringPointerEmptyAsNull(from.DefaultValue)
	m.Syntax = flex.FlattenStringPointerEmptyAsNull(from.Syntax)
}
