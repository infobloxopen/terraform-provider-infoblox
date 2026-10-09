package notification

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	niosnotification "github.com/infobloxopen/infoblox-nios-go-client/notification"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
	customvalidator "github.com/infobloxopen/terraform-provider-infoblox/internal/validator"
)

// NotificationRulePublishSettingsModel is the Terraform model for NotificationRulePublishSettings
type NotificationRulePublishSettingsModel struct {
	EnabledAttributes types.List `tfsdk:"enabled_attributes"`
}

// NotificationRulePublishSettingsAttrTypes contains the attribute types for NotificationRulePublishSettingsModel
var NotificationRulePublishSettingsAttrTypes = map[string]attr.Type{
	"enabled_attributes": types.ListType{ElemType: types.StringType},
}

// NotificationRulePublishSettingsResourceSchemaAttributes contains the schema attributes for NotificationRulePublishSettingsModel
var NotificationRulePublishSettingsResourceSchemaAttributes = map[string]schema.Attribute{
	"enabled_attributes": schema.ListAttribute{
		ElementType: types.StringType,
		Required:    true,
		Validators: []validator.List{
			customvalidator.ListNotEmpty(),
			listvalidator.ValueStringsAre(stringvalidator.OneOf("CLIENT_ID", "FINGERPRINT", "HOSTNAME", "INFOBLOX_MEMBER", "IPADDRESS", "LEASE_END_TIME", "LEASE_START_TIME", "LEASE_STATE", "MAC_OR_DUID", "NETBIOS_NAME")),
		},
		MarkdownDescription: "The list of NIOS extensible attributes enalbed for publishsing to Cisco ISE endpoint.",
	},
}

// ExpandNotificationRulePublishSettings converts a Terraform Object to SDK type
func ExpandNotificationRulePublishSettings(ctx context.Context, o types.Object, diags *diag.Diagnostics) *niosnotification.NotificationRulePublishSettings {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m NotificationRulePublishSettingsModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

// Expand converts the Terraform model to SDK type
func (m *NotificationRulePublishSettingsModel) Expand(ctx context.Context, diags *diag.Diagnostics) *niosnotification.NotificationRulePublishSettings {
	if m == nil {
		return nil
	}
	to := &niosnotification.NotificationRulePublishSettings{
		EnabledAttributes: flex.ExpandFrameworkListString(ctx, m.EnabledAttributes, diags),
	}
	return to
}

// FlattenNotificationRulePublishSettings converts an SDK type to Terraform Object
func FlattenNotificationRulePublishSettings(ctx context.Context, from *niosnotification.NotificationRulePublishSettings, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(NotificationRulePublishSettingsAttrTypes)
	}
	m := &NotificationRulePublishSettingsModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, NotificationRulePublishSettingsAttrTypes, m)
	diags.Append(d...)
	return t
}

// Flatten populates the Terraform model from SDK type
func (m *NotificationRulePublishSettingsModel) Flatten(ctx context.Context, from *niosnotification.NotificationRulePublishSettings, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	m.EnabledAttributes = flex.FlattenFrameworkListString(ctx, from.EnabledAttributes, diags)
}
