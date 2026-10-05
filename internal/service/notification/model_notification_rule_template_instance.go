package notification

import (
	"context"

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

// NotificationRuleTemplateInstanceModel is the Terraform model for NotificationRuleTemplateInstance
type NotificationRuleTemplateInstanceModel struct {
	Template   types.String `tfsdk:"template"`
	Parameters types.List   `tfsdk:"parameters"`
}

// NotificationRuleTemplateInstanceAttrTypes contains the attribute types for NotificationRuleTemplateInstanceModel
var NotificationRuleTemplateInstanceAttrTypes = map[string]attr.Type{
	"template":   types.StringType,
	"parameters": types.ListType{ElemType: types.ObjectType{AttrTypes: NotificationruletemplateinstanceParametersAttrTypes}},
}

// NotificationRuleTemplateInstanceResourceSchemaAttributes contains the schema attributes for NotificationRuleTemplateInstanceModel
var NotificationRuleTemplateInstanceResourceSchemaAttributes = map[string]schema.Attribute{
	"template": schema.StringAttribute{
		Required: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
		},
		MarkdownDescription: "The name of the REST API template parameter.",
	},
	"parameters": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: NotificationruletemplateinstanceParametersResourceSchemaAttributes,
		},
		Optional: true,
		Validators: []validator.List{
			customvalidator.ListNotEmpty(),
		},
		MarkdownDescription: "The notification REST template parameters.",
	},
}

// ExpandNotificationRuleTemplateInstance converts a Terraform Object to SDK type
func ExpandNotificationRuleTemplateInstance(ctx context.Context, o types.Object, diags *diag.Diagnostics) *niosnotification.NotificationRuleTemplateInstance {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m NotificationRuleTemplateInstanceModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

// Expand converts the Terraform model to SDK type
func (m *NotificationRuleTemplateInstanceModel) Expand(ctx context.Context, diags *diag.Diagnostics) *niosnotification.NotificationRuleTemplateInstance {
	if m == nil {
		return nil
	}
	to := &niosnotification.NotificationRuleTemplateInstance{
		Template:   flex.ExpandStringPointerNullAsEmpty(m.Template),
		Parameters: flex.ExpandFrameworkListNestedBlock(ctx, m.Parameters, diags, ExpandNotificationruletemplateinstanceParameters),
	}
	return to
}

// FlattenNotificationRuleTemplateInstance converts an SDK type to Terraform Object
func FlattenNotificationRuleTemplateInstance(ctx context.Context, from *niosnotification.NotificationRuleTemplateInstance, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(NotificationRuleTemplateInstanceAttrTypes)
	}
	m := &NotificationRuleTemplateInstanceModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, NotificationRuleTemplateInstanceAttrTypes, m)
	diags.Append(d...)
	return t
}

// Flatten populates the Terraform model from SDK type
func (m *NotificationRuleTemplateInstanceModel) Flatten(ctx context.Context, from *niosnotification.NotificationRuleTemplateInstance, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	m.Template = flex.FlattenStringPointerEmptyAsNull(from.Template)
	m.Parameters = flex.FlattenFrameworkListNestedBlock(ctx, from.Parameters, NotificationruletemplateinstanceParametersAttrTypes, diags, FlattenNotificationruletemplateinstanceParameters)
}
