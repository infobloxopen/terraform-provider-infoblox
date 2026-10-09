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

// NotificationRestEndpointTemplateInstanceModel is the Terraform model for NotificationRestEndpointTemplateInstance
type NotificationRestEndpointTemplateInstanceModel struct {
	Template   types.String `tfsdk:"template"`
	Parameters types.List   `tfsdk:"parameters"`
}

// NotificationRestEndpointTemplateInstanceAttrTypes contains the attribute types for NotificationRestEndpointTemplateInstanceModel
var NotificationRestEndpointTemplateInstanceAttrTypes = map[string]attr.Type{
	"template":   types.StringType,
	"parameters": types.ListType{ElemType: types.ObjectType{AttrTypes: NotificationrestendpointtemplateinstanceParametersAttrTypes}},
}

// NotificationRestEndpointTemplateInstanceResourceSchemaAttributes contains the schema attributes for NotificationRestEndpointTemplateInstanceModel
var NotificationRestEndpointTemplateInstanceResourceSchemaAttributes = map[string]schema.Attribute{
	"template": schema.StringAttribute{
		Required: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
		},
		MarkdownDescription: "The name of the REST API template parameter.",
	},
	"parameters": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: NotificationrestendpointtemplateinstanceParametersResourceSchemaAttributes,
		},
		Optional: true,
		Validators: []validator.List{
			customvalidator.ListNotEmpty(),
		},
		MarkdownDescription: "The notification REST template parameters.",
	},
}

// ExpandNotificationRestEndpointTemplateInstance converts a Terraform Object to SDK type
func ExpandNotificationRestEndpointTemplateInstance(ctx context.Context, o types.Object, diags *diag.Diagnostics) *niosnotification.NotificationRestEndpointTemplateInstance {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m NotificationRestEndpointTemplateInstanceModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

// Expand converts the Terraform model to SDK type
func (m *NotificationRestEndpointTemplateInstanceModel) Expand(ctx context.Context, diags *diag.Diagnostics) *niosnotification.NotificationRestEndpointTemplateInstance {
	if m == nil {
		return nil
	}
	to := &niosnotification.NotificationRestEndpointTemplateInstance{
		Template:   flex.ExpandStringPointerNullAsEmpty(m.Template),
		Parameters: flex.ExpandFrameworkListNestedBlock(ctx, m.Parameters, diags, ExpandNotificationrestendpointtemplateinstanceParameters),
	}
	return to
}

// FlattenNotificationRestEndpointTemplateInstance converts an SDK type to Terraform Object
func FlattenNotificationRestEndpointTemplateInstance(ctx context.Context, from *niosnotification.NotificationRestEndpointTemplateInstance, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(NotificationRestEndpointTemplateInstanceAttrTypes)
	}
	m := &NotificationRestEndpointTemplateInstanceModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, NotificationRestEndpointTemplateInstanceAttrTypes, m)
	diags.Append(d...)
	return t
}

// Flatten populates the Terraform model from SDK type
func (m *NotificationRestEndpointTemplateInstanceModel) Flatten(ctx context.Context, from *niosnotification.NotificationRestEndpointTemplateInstance, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	m.Template = flex.FlattenStringPointerEmptyAsNull(from.Template)
	m.Parameters = flex.FlattenFrameworkListNestedBlock(ctx, from.Parameters, NotificationrestendpointtemplateinstanceParametersAttrTypes, diags, FlattenNotificationrestendpointtemplateinstanceParameters)
}
