package notification

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"

	coremodel "github.com/infobloxopen/terraform-provider-infoblox/internal/core/model/notification"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
	immutable "github.com/infobloxopen/terraform-provider-infoblox/internal/planmodifiers/immutable"
	internaltypes "github.com/infobloxopen/terraform-provider-infoblox/internal/types"
	customvalidator "github.com/infobloxopen/terraform-provider-infoblox/internal/validator"
)

type NotificationRuleModel struct {
	Id            types.String `tfsdk:"id"`
	UpdateTrigger types.String `tfsdk:"update_trigger"`
	NIOS          types.Object `tfsdk:"nios"`
}

var NotificationRuleAttrTypes = map[string]attr.Type{
	"id":             types.StringType,
	"update_trigger": types.StringType,
	"nios":           types.ObjectType{AttrTypes: NIOSNotificationRuleAttrTypes},
}

type NIOSNotificationRuleModel struct {
	AllMembers                       types.Bool                          `tfsdk:"all_members"`
	Comment                          types.String                        `tfsdk:"comment"`
	Disable                          types.Bool                          `tfsdk:"disable"`
	EnableEventDeduplication         types.Bool                          `tfsdk:"enable_event_deduplication"`
	EnableEventDeduplicationLog      types.Bool                          `tfsdk:"enable_event_deduplication_log"`
	EventDeduplicationFields         types.List                          `tfsdk:"event_deduplication_fields"`
	EventDeduplicationLookbackPeriod types.Int64                         `tfsdk:"event_deduplication_lookback_period"`
	EventPriority                    types.String                        `tfsdk:"event_priority"`
	EventType                        types.String                        `tfsdk:"event_type"`
	ExpressionList                   types.List                          `tfsdk:"expression_list"`
	Name                             types.String                        `tfsdk:"name"`
	NotificationAction               types.String                        `tfsdk:"notification_action"`
	NotificationTarget               internaltypes.CaseInsensitiveString `tfsdk:"notification_target"`
	PublishSettings                  types.Object                        `tfsdk:"publish_settings"`
	ScheduledEvent                   types.Object                        `tfsdk:"scheduled_event"`
	SelectedMembers                  types.List                          `tfsdk:"selected_members"`
	TemplateInstance                 types.Object                        `tfsdk:"template_instance"`
}

var NIOSNotificationRuleAttrTypes = map[string]attr.Type{
	"all_members":                         types.BoolType,
	"comment":                             types.StringType,
	"disable":                             types.BoolType,
	"enable_event_deduplication":          types.BoolType,
	"enable_event_deduplication_log":      types.BoolType,
	"event_deduplication_fields":          types.ListType{ElemType: types.StringType},
	"event_deduplication_lookback_period": types.Int64Type,
	"event_priority":                      types.StringType,
	"event_type":                          types.StringType,
	"expression_list":                     types.ListType{ElemType: types.ObjectType{AttrTypes: NotificationRuleExpressionListAttrTypes}},
	"name":                                types.StringType,
	"notification_action":                 types.StringType,
	"notification_target":                 internaltypes.CaseInsensitiveStringType{},
	"publish_settings":                    types.ObjectType{AttrTypes: NotificationRulePublishSettingsAttrTypes},
	"scheduled_event":                     types.ObjectType{AttrTypes: NotificationRuleScheduledEventAttrTypes},
	"selected_members":                    types.ListType{ElemType: types.StringType},
	"template_instance":                   types.ObjectType{AttrTypes: NotificationRuleTemplateInstanceAttrTypes},
}

const (
	NotificationRuleReturnFields = "all_members,comment,disable,enable_event_deduplication,enable_event_deduplication_log,event_deduplication_fields,event_deduplication_lookback_period,event_priority,event_type,expression_list,name,notification_action,notification_target,publish_settings,scheduled_event,selected_members,template_instance,use_publish_settings"
)

var NotificationRuleResourceSchemaAttributes = map[string]schema.Attribute{
	"id": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The reference to the object.",
	},
	"update_trigger": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "An arbitrary value used to trigger an update. Not sent to the API. Change it when Terraform reports no infrastructure changes.",
	},
	"nios": schema.SingleNestedAttribute{
		Optional:            true,
		MarkdownDescription: "NIOS backend-specific fields.",
		Attributes:          NotificationRuleResourceNiosSchemaAttributes,
	},
}

var NotificationRuleResourceNiosSchemaAttributes = map[string]schema.Attribute{
	"all_members": schema.BoolAttribute{
		Computed:            true,
		MarkdownDescription: "Determines whether the notification rule is applied on all members or not. When this is set to False, the notification rule is applied only on selected_members.",
	},
	"comment": schema.StringAttribute{
		Optional: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
			stringvalidator.LengthBetween(0, 256),
		},
		MarkdownDescription: "The notification rule descriptive comment.",
	},
	"disable": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Determines whether a notification rule is disabled or not. When this is set to False, the notification rule is enabled.",
	},
	"enable_event_deduplication": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Determines whether the notification rule for event deduplication is enabled. Note that to enable event deduplication, you must set at least one deduplication field.",
	},
	"enable_event_deduplication_log": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Determines whether the notification rule for the event deduplication syslog is enabled.",
	},
	"event_deduplication_fields": schema.ListAttribute{
		ElementType: types.StringType,
		Optional:    true,
		Validators: []validator.List{
			customvalidator.ListNotEmpty(),
			listvalidator.ValueStringsAre(stringvalidator.OneOf("DISCOVERER", "DUID", "DXL_TOPIC", "IP_ADDRESS", "MAC_ADDRESS", "NETWORK", "NETWORK_VIEW", "OPERATION_TYPE", "QUERY_FQDN", "QUERY_NAME", "QUERY_TYPE", "RPZ_POLICY", "RPZ_TYPE", "RULE_ACTION", "RULE_CATEGORY", "RULE_SEVERITY", "RULE_SID", "SOURCE_IP", "SOURCE_PORT")),
		},
		MarkdownDescription: "The list of fields that must be used in the notification rule for event deduplication.",
	},
	"event_deduplication_lookback_period": schema.Int64Attribute{
		Optional:            true,
		Computed:            true,
		Default:             int64default.StaticInt64(600),
		MarkdownDescription: "The lookback period for the notification rule for event deduplication.",
	},
	"event_priority": schema.StringAttribute{
		Default:  stringdefault.StaticString("NORMAL"),
		Optional: true,
		Computed: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
		},
		MarkdownDescription: "Event priority.",
	},
	"event_type": schema.StringAttribute{
		Validators: []validator.String{
			stringvalidator.OneOf("DXL_EVENT_SUBSCRIBER", "DB_CHANGE_DNS_RECORD", "DB_CHANGE_DNS_ZONE", "DNS_RPZ", "DHCP_LEASES", "SECURITY_ADP", "IPAM", "ANALYTICS_DNS_TUNNEL", "DB_CHANGE_DHCP_FIXED_ADDRESS_IPV4", "DB_CHANGE_DHCP_FIXED_ADDRESS_IPV6", "DB_CHANGE_DHCP_NETWORK_IPV4", "DB_CHANGE_DHCP_NETWORK_IPV6", "DB_CHANGE_DHCP_RANGE_IPV4", "DB_CHANGE_DHCP_RANGE_IPV6", "DB_CHANGE_DNS_HOST_ADDRESS_IPV4", "DB_CHANGE_DNS_HOST_ADDRESS_IPV6", "DB_CHANGE_DNS_DISCOVERY_DATA", "SCHEDULE"),
		},
		Required:            true,
		MarkdownDescription: "The notification rule event type.",
	},
	"expression_list": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: NotificationRuleExpressionListResourceSchemaAttributes,
		},
		Optional: true,
		Validators: []validator.List{
			customvalidator.ListNotEmpty(),
		},
		MarkdownDescription: "The notification rule expression list.",
	},
	"name": schema.StringAttribute{
		Required: true,
		PlanModifiers: []planmodifier.String{
			immutable.ImmutableString(),
		},
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
			customvalidator.ValidateTrimmedString(),
		},
		MarkdownDescription: "The notification rule name.",
	},
	"notification_action": schema.StringAttribute{
		Validators: []validator.String{
			stringvalidator.OneOf("CISCOISE_QUARANTINE", "CISCOISE_PUBLISH", "RESTAPI_TEMPLATE_INSTANCE"),
		},
		Required:            true,
		MarkdownDescription: "The notification rule action is applied if expression list evaluates to True.",
	},
	"notification_target": schema.StringAttribute{
		Required:   true,
		CustomType: internaltypes.CaseInsensitiveStringType{},
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
		},
		MarkdownDescription: "The notification target.",
	},
	"publish_settings": schema.SingleNestedAttribute{
		Attributes:          NotificationRulePublishSettingsResourceSchemaAttributes,
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "",
	},
	"scheduled_event": schema.SingleNestedAttribute{
		Attributes:          NotificationRuleScheduledEventResourceSchemaAttributes,
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "",
	},
	"selected_members": schema.ListAttribute{
		ElementType:         types.StringType,
		Computed:            true,
		MarkdownDescription: "The list of the members on which the notification rule is applied.",
	},
	"template_instance": schema.SingleNestedAttribute{
		Attributes:          NotificationRuleTemplateInstanceResourceSchemaAttributes,
		Required:            true,
		MarkdownDescription: "",
	},
}

// Expand converts the TF model to the infoblox core model
func (m *NotificationRuleModel) Expand(ctx context.Context, diags *diag.Diagnostics, isCreate bool) *coremodel.NotificationRule {
	if m == nil {
		return nil
	}

	obj := &coremodel.NotificationRule{}

	// Expand NIOS nested attribute (returns nil if not present)
	niosModel := flex.ExpandNestedObject[NIOSNotificationRuleModel](ctx, m.NIOS, diags)
	if niosModel != nil {
		obj.NIOS = niosModel.Expand(ctx, diags, isCreate)
	}

	return obj
}

// Expand converts the NIOS TF model to the core model.
func (m *NIOSNotificationRuleModel) Expand(ctx context.Context, diags *diag.Diagnostics, isCreate bool) *coremodel.NIOSNotificationRuleExt {
	ext := &coremodel.NIOSNotificationRuleExt{
		Comment:                          flex.ExpandStringPointerNullAsEmpty(m.Comment),
		Disable:                          flex.ExpandBoolPointer(m.Disable),
		EnableEventDeduplication:         flex.ExpandBoolPointer(m.EnableEventDeduplication),
		EnableEventDeduplicationLog:      flex.ExpandBoolPointer(m.EnableEventDeduplicationLog),
		EventDeduplicationFields:         flex.ExpandFrameworkListString(ctx, m.EventDeduplicationFields, diags),
		EventDeduplicationLookbackPeriod: flex.ExpandInt64Pointer(m.EventDeduplicationLookbackPeriod),
		EventPriority:                    flex.ExpandStringPointerNullAsEmpty(m.EventPriority),
		EventType:                        flex.ExpandStringPointerNullAsEmpty(m.EventType),
		ExpressionList:                   flex.ExpandFrameworkListNestedBlock(ctx, m.ExpressionList, diags, ExpandNotificationRuleExpressionList),
		NotificationAction:               flex.ExpandStringPointerNullAsEmpty(m.NotificationAction),
		NotificationTarget:               flex.ExpandStringPointer(m.NotificationTarget.StringValue),
		PublishSettings:                  ExpandNotificationRulePublishSettings(ctx, m.PublishSettings, diags),
		ScheduledEvent:                   ExpandNotificationRuleScheduledEvent(ctx, m.ScheduledEvent, diags),
		TemplateInstance:                 ExpandNotificationRuleTemplateInstance(ctx, m.TemplateInstance, diags),
	}
	if isCreate {
		ext.Name = flex.ExpandStringPointerNullAsEmpty(m.Name)
	}
	return ext
}

// ApplyNotificationRuleNIOSUseFlags derives NIOS use flags from the raw config
// value(s) and writes them onto the core model. A flag is true when the user
// set any of its governed value fields in config.
func ApplyNotificationRuleNIOSUseFlags(ctx context.Context, config tfsdk.Config, obj *coremodel.NotificationRule, diags *diag.Diagnostics) {
	if obj == nil || obj.NIOS == nil {
		return
	}
	obj.NIOS.UsePublishSettings = flex.DeriveUseFlag(ctx, config, diags, path.Root("nios").AtName("publish_settings"))
}

// Flatten populates the TF model from a core response.
func (m *NotificationRuleModel) Flatten(ctx context.Context, resp *coremodel.NotificationRule, diags *diag.Diagnostics) {
	if resp == nil {
		return
	}

	m.Id = flex.FlattenStringPointer(resp.Id)

	// Extract existing NIOS model, flatten API response onto it, convert back
	niosModel := flex.ExpandNestedObject[NIOSNotificationRuleModel](ctx, m.NIOS, diags)
	if niosModel == nil {
		niosModel = &NIOSNotificationRuleModel{}
	}
	niosModel.Flatten(ctx, resp.NIOS, diags)
	if resp.NIOS != nil {
		m.NIOS = flex.FlattenNestedObject(ctx, niosModel, NIOSNotificationRuleAttrTypes, diags)
	} else {
		m.NIOS = types.ObjectNull(NIOSNotificationRuleAttrTypes)
	}

}

// Flatten merges API response onto existing NIOS model.
func (m *NIOSNotificationRuleModel) Flatten(ctx context.Context, from *coremodel.NIOSNotificationRuleExt, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	m.AllMembers = flex.FlattenBoolPointer(from.AllMembers)
	m.Comment = flex.FlattenStringPointerEmptyAsNull(from.Comment)
	m.Disable = flex.FlattenBoolPointer(from.Disable)
	m.EnableEventDeduplication = flex.FlattenBoolPointer(from.EnableEventDeduplication)
	m.EnableEventDeduplicationLog = flex.FlattenBoolPointer(from.EnableEventDeduplicationLog)
	m.EventDeduplicationFields = flex.FlattenFrameworkListString(ctx, from.EventDeduplicationFields, diags)
	m.EventDeduplicationLookbackPeriod = flex.FlattenInt64Pointer(from.EventDeduplicationLookbackPeriod)
	m.EventPriority = flex.FlattenStringPointerEmptyAsNull(from.EventPriority)
	m.EventType = flex.FlattenStringPointerEmptyAsNull(from.EventType)
	m.ExpressionList = flex.FlattenFrameworkListNestedBlock(ctx, from.ExpressionList, NotificationRuleExpressionListAttrTypes, diags, FlattenNotificationRuleExpressionList)
	m.Name = flex.FlattenStringPointerEmptyAsNull(from.Name)
	m.NotificationAction = flex.FlattenStringPointerEmptyAsNull(from.NotificationAction)
	m.NotificationTarget.StringValue = flex.FlattenStringPointer(from.NotificationTarget)
	m.PublishSettings = FlattenNotificationRulePublishSettings(ctx, from.PublishSettings, diags)
	m.ScheduledEvent = FlattenNotificationRuleScheduledEvent(ctx, from.ScheduledEvent, diags)
	m.SelectedMembers = flex.FlattenFrameworkListString(ctx, from.SelectedMembers, diags)
	m.TemplateInstance = FlattenNotificationRuleTemplateInstance(ctx, from.TemplateInstance, diags)
}
