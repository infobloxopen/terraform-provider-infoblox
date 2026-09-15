package fw

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/mapvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/mapdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"

	coremodel "github.com/infobloxopen/terraform-provider-infoblox/internal/core/model/fw"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
	customvalidator "github.com/infobloxopen/terraform-provider-infoblox/internal/validator"
)

type SecurityPolicyModel struct {
	Id   types.Int32  `tfsdk:"id"`
	UDDI types.Object `tfsdk:"uddi"`
}

var SecurityPolicyAttrTypes = map[string]attr.Type{
	"id":   types.Int32Type,
	"uddi": types.ObjectType{AttrTypes: UDDISecurityPolicyAttrTypes},
}

type UDDISecurityPolicyModel struct {
	AccessCodes         types.List   `tfsdk:"access_codes"`
	DefaultAction       types.String `tfsdk:"default_action"`
	DefaultRedirectName types.String `tfsdk:"default_redirect_name"`
	Description         types.String `tfsdk:"description"`
	DfpServices         types.List   `tfsdk:"dfp_services"`
	Dfps                types.List   `tfsdk:"dfps"`
	Ecs                 types.Bool   `tfsdk:"ecs"`
	Name                types.String `tfsdk:"name"`
	NetAddressDfps      types.List   `tfsdk:"net_address_dfps"`
	NetworkLists        types.List   `tfsdk:"network_lists"`
	OnpremResolve       types.Bool   `tfsdk:"onprem_resolve"`
	Precedence          types.Int32  `tfsdk:"precedence"`
	RoamingDeviceGroups types.List   `tfsdk:"roaming_device_groups"`
	Rules               types.List   `tfsdk:"rules"`
	SafeSearch          types.Bool   `tfsdk:"safe_search"`
	Tags                types.Map    `tfsdk:"tags"`
	TagsAll             types.Map    `tfsdk:"tags_all"`
	UserGroups          types.List   `tfsdk:"user_groups"`
}

var UDDISecurityPolicyAttrTypes = map[string]attr.Type{
	"access_codes":          types.ListType{ElemType: types.StringType},
	"default_action":        types.StringType,
	"default_redirect_name": types.StringType,
	"description":           types.StringType,
	"dfp_services":          types.ListType{ElemType: types.StringType},
	"dfps":                  types.ListType{ElemType: types.Int32Type},
	"ecs":                   types.BoolType,
	"name":                  types.StringType,
	"net_address_dfps":      types.ListType{ElemType: types.ObjectType{AttrTypes: NetAddrDfpAssignmentAttrTypes}},
	"network_lists":         types.ListType{ElemType: types.Int64Type},
	"onprem_resolve":        types.BoolType,
	"precedence":            types.Int32Type,
	"roaming_device_groups": types.ListType{ElemType: types.Int32Type},
	"rules":                 types.ListType{ElemType: types.ObjectType{AttrTypes: SecurityPolicyRuleAttrTypes}},
	"safe_search":           types.BoolType,
	"tags":                  types.MapType{ElemType: types.StringType},
	"tags_all":              types.MapType{ElemType: types.StringType},
	"user_groups":           types.ListType{ElemType: types.StringType},
}

const (
	SecurityPolicyReturnFields = ""
)

var SecurityPolicyResourceSchemaAttributes = map[string]schema.Attribute{
	"id": schema.Int32Attribute{
		Computed:            true,
		MarkdownDescription: "The Security Policy object identifier.",
	},
	"uddi": schema.SingleNestedAttribute{
		Optional:            true,
		MarkdownDescription: "UDDI backend-specific fields.",
		Attributes:          SecurityPolicyResourceUddiSchemaAttributes,
	},
}

var SecurityPolicyResourceUddiSchemaAttributes = map[string]schema.Attribute{
	"access_codes": schema.ListAttribute{
		ElementType: types.StringType,
		Optional:    true,
		Validators: []validator.List{
			customvalidator.ListNotEmpty(),
		},
		MarkdownDescription: "Access codes assigned to Security Policy",
	},
	"default_action": schema.StringAttribute{
		Default:             stringdefault.StaticString("action_allow"),
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The policy-level action gets applied when none of the policy rules apply/match. The default value for default_action is \"action_allow\".",
	},
	"default_redirect_name": schema.StringAttribute{
		Default:             stringdefault.StaticString(""),
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "Name of the custom redirect, if the default_action is \"action_redirect\".",
	},
	"description": schema.StringAttribute{
		Default:             stringdefault.StaticString(""),
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The brief description for the security policy.",
	},
	"dfp_services": schema.ListAttribute{
		ElementType:         types.StringType,
		Computed:            true,
		MarkdownDescription: "The list of DNS Forwarding Proxy Services object identifiers. For Internal Use only.",
	},
	"dfps": schema.ListAttribute{
		ElementType:         types.Int32Type,
		Optional:            true,
		Computed:            true,
		Default:             listdefault.StaticValue(types.ListValueMust(types.Int32Type, []attr.Value{})),
		MarkdownDescription: "The list of DNS Forwarding Proxy object identifiers.",
	},
	"ecs": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Use ECS for handling policy",
	},
	"name": schema.StringAttribute{
		Required:            true,
		MarkdownDescription: "The name of the security policy.",
	},
	"net_address_dfps": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: NetAddrDfpAssignmentResourceSchemaAttributes,
		},
		Optional: true,
		Validators: []validator.List{
			customvalidator.ListNotEmpty(),
		},
		MarkdownDescription: "List of DFPs associated with this policy via network address (with corresponding network address)",
	},
	"network_lists": schema.ListAttribute{
		ElementType:         types.Int64Type,
		Optional:            true,
		Computed:            true,
		Default:             listdefault.StaticValue(types.ListValueMust(types.Int64Type, []attr.Value{})),
		MarkdownDescription: "The list of Network Lists identifiers that represents networks that you want to protect from malicious attacks.",
	},
	"onprem_resolve": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Use DNS resolve on onprem side",
	},
	"precedence": schema.Int32Attribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "Security precedence enable selection of the highest priority policy, in cases where a query matches multiple policies.",
	},
	"roaming_device_groups": schema.ListAttribute{
		ElementType:         types.Int32Type,
		Optional:            true,
		Computed:            true,
		Default:             listdefault.StaticValue(types.ListValueMust(types.Int32Type, []attr.Value{})),
		MarkdownDescription: "The list of Infoblox Endpoint groups identifiers.",
	},
	"rules": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: SecurityPolicyRuleResourceSchemaAttributes,
		},
		Optional: true,
		Validators: []validator.List{
			customvalidator.ListNotEmpty(),
		},
		MarkdownDescription: "The list of Security Policy Rules objects that represent the set of rules and actions that you define to balance access and constraints so you can mitigate malicious attacks and provide security for your networks.",
	},
	"safe_search": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Apply automated rules to enforce safe search",
	},
	"tags": schema.MapAttribute{
		Optional:    true,
		Computed:    true,
		ElementType: types.StringType,
		Default:     mapdefault.StaticValue(types.MapNull(types.StringType)),
		Validators: []validator.Map{
			mapvalidator.SizeAtLeast(1),
		},
		MarkdownDescription: "Enables tag support for resource where tags attribute contains user-defined key value pairs",
	},
	"tags_all": schema.MapAttribute{
		Computed:            true,
		ElementType:         types.StringType,
		MarkdownDescription: "All tags including inherited values.",
	},
	"user_groups": schema.ListAttribute{
		ElementType: types.StringType,
		Optional:    true,
		Validators: []validator.List{
			customvalidator.ListNotEmpty(),
		},
		MarkdownDescription: "List of user groups associated with this policy",
	},
}

// Expand converts the TF model to the infoblox core model
func (m *SecurityPolicyModel) Expand(ctx context.Context, diags *diag.Diagnostics, isCreate bool) *coremodel.SecurityPolicy {
	if m == nil {
		return nil
	}

	obj := &coremodel.SecurityPolicy{}

	// Expand UDDI nested attribute (returns nil if not present)
	uddiModel := flex.ExpandNestedObject[UDDISecurityPolicyModel](ctx, m.UDDI, diags)
	if uddiModel != nil {
		obj.UDDI = uddiModel.Expand(ctx, diags)
	}

	return obj
}

// Expand converts the UDDI TF model to the core model.
func (m *UDDISecurityPolicyModel) Expand(ctx context.Context, diags *diag.Diagnostics) *coremodel.UDDISecurityPolicyExt {
	return &coremodel.UDDISecurityPolicyExt{
		AccessCodes:         flex.ExpandFrameworkListString(ctx, m.AccessCodes, diags),
		DefaultAction:       flex.ExpandStringPointer(m.DefaultAction),
		DefaultRedirectName: flex.ExpandStringPointer(m.DefaultRedirectName),
		Description:         flex.ExpandStringPointer(m.Description),
		Dfps:                flex.ExpandFrameworkListInt32(ctx, m.Dfps, diags),
		Ecs:                 flex.ExpandBoolPointer(m.Ecs),
		Name:                flex.ExpandStringPointer(m.Name),
		NetAddressDfps:      flex.ExpandFrameworkListNestedBlock(ctx, m.NetAddressDfps, diags, ExpandNetAddrDfpAssignment),
		NetworkLists:        flex.ExpandFrameworkListInt64(ctx, m.NetworkLists, diags),
		OnpremResolve:       flex.ExpandBoolPointer(m.OnpremResolve),
		Precedence:          flex.ExpandInt32Pointer(m.Precedence),
		RoamingDeviceGroups: flex.ExpandFrameworkListInt32(ctx, m.RoamingDeviceGroups, diags),
		Rules:               flex.ExpandFrameworkListNestedBlock(ctx, m.Rules, diags, ExpandSecurityPolicyRule),
		SafeSearch:          flex.ExpandBoolPointer(m.SafeSearch),
		Tags:                flex.ExpandMapStringAny(ctx, m.Tags, diags),
		UserGroups:          flex.ExpandFrameworkListString(ctx, m.UserGroups, diags),
	}
}

// Flatten populates the TF model from a core response.
func (m *SecurityPolicyModel) Flatten(ctx context.Context, resp *coremodel.SecurityPolicy, diags *diag.Diagnostics) {
	if resp == nil {
		return
	}

	m.Id = flex.FlattenInt32Pointer(resp.Id)

	// Extract existing UDDI model, flatten API response onto it, convert back
	uddiModel := flex.ExpandNestedObject[UDDISecurityPolicyModel](ctx, m.UDDI, diags)
	if uddiModel == nil {
		uddiModel = &UDDISecurityPolicyModel{}
	}
	plannedUDDI := flex.ExpandNestedObject[UDDISecurityPolicyModel](ctx, m.UDDI, diags)
	uddiModel.Flatten(ctx, resp.UDDI, diags)
	if resp.UDDI != nil {
		PostFlattenSecurityPolicyUDDI(ctx, plannedUDDI, uddiModel, diags)
		m.UDDI = flex.FlattenNestedObject(ctx, uddiModel, UDDISecurityPolicyAttrTypes, diags)
	} else {
		m.UDDI = types.ObjectNull(UDDISecurityPolicyAttrTypes)
	}
}

// Flatten merges API response onto existing UDDI model.
func (m *UDDISecurityPolicyModel) Flatten(ctx context.Context, from *coremodel.UDDISecurityPolicyExt, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	m.AccessCodes = flex.FlattenFrameworkListString(ctx, from.AccessCodes, diags)
	m.DefaultAction = flex.FlattenStringPointer(from.DefaultAction)
	m.DefaultRedirectName = flex.FlattenStringPointer(from.DefaultRedirectName)
	m.Description = flex.FlattenStringPointer(from.Description)
	m.DfpServices = flex.FlattenFrameworkListString(ctx, from.DfpServices, diags)
	m.Dfps = flex.FlattenFrameworkListInt32(ctx, from.Dfps, diags)
	m.Ecs = flex.FlattenBoolPointer(from.Ecs)
	m.Name = flex.FlattenStringPointer(from.Name)
	m.NetAddressDfps = flex.FlattenFrameworkListNestedBlock(ctx, from.NetAddressDfps, NetAddrDfpAssignmentAttrTypes, diags, FlattenNetAddrDfpAssignment)
	m.NetworkLists = flex.FlattenFrameworkListInt64(ctx, from.NetworkLists, diags)
	m.OnpremResolve = flex.FlattenBoolPointer(from.OnpremResolve)
	m.Precedence = flex.FlattenInt32Pointer(from.Precedence)
	m.RoamingDeviceGroups = flex.FlattenFrameworkListInt32(ctx, from.RoamingDeviceGroups, diags)
	m.Rules = flex.FlattenFrameworkListNestedBlock(ctx, from.Rules, SecurityPolicyRuleAttrTypes, diags, FlattenSecurityPolicyRule)
	m.SafeSearch = flex.FlattenBoolPointer(from.SafeSearch)
	tagsAll := flex.FlattenMapStringAny(ctx, from.Tags, diags)
	if m.Tags.IsNull() || m.Tags.IsUnknown() {
		m.Tags = tagsAll
	}
	m.TagsAll = tagsAll
	m.UserGroups = flex.FlattenFrameworkListString(ctx, from.UserGroups, diags)
}
