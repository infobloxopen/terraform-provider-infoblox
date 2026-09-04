package fw

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	stringplanmodifier "github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"

	coremodel "github.com/infobloxopen/terraform-provider-infoblox/internal/core/model/fw"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
	customvalidator "github.com/infobloxopen/terraform-provider-infoblox/internal/validator"
)

type AccessCodeModel struct {
	Id   types.String `tfsdk:"id"`
	UDDI types.Object `tfsdk:"uddi"`
}

var AccessCodeAttrTypes = map[string]attr.Type{
	"id":   types.StringType,
	"uddi": types.ObjectType{AttrTypes: UDDIAccessCodeAttrTypes},
}

type UDDIAccessCodeModel struct {
	AccessKey   types.String      `tfsdk:"access_key"`
	Activation  timetypes.RFC3339 `tfsdk:"activation"`
	CreatedTime timetypes.RFC3339 `tfsdk:"created_time"`
	Description types.String      `tfsdk:"description"`
	Expiration  timetypes.RFC3339 `tfsdk:"expiration"`
	Name        types.String      `tfsdk:"name"`
	PolicyIds   types.List        `tfsdk:"policy_ids"`
	Rules       types.List        `tfsdk:"rules"`
	UpdatedTime timetypes.RFC3339 `tfsdk:"updated_time"`
}

var UDDIAccessCodeAttrTypes = map[string]attr.Type{
	"access_key":   types.StringType,
	"activation":   timetypes.RFC3339Type{},
	"created_time": timetypes.RFC3339Type{},
	"description":  types.StringType,
	"expiration":   timetypes.RFC3339Type{},
	"name":         types.StringType,
	"policy_ids":   types.ListType{ElemType: types.Int32Type},
	"rules":        types.ListType{ElemType: types.ObjectType{AttrTypes: AccessCodeRuleAttrTypes}},
	"updated_time": timetypes.RFC3339Type{},
}

const (
	AccessCodeReturnFields = ""
)

var AccessCodeResourceSchemaAttributes = map[string]schema.Attribute{
	"id": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "Auto generated unique Bypass Code value",
	},
	"uddi": schema.SingleNestedAttribute{
		Optional:            true,
		MarkdownDescription: "UDDI backend-specific fields.",
		Attributes:          AccessCodeResourceUddiSchemaAttributes,
	},
}

var AccessCodeResourceUddiSchemaAttributes = map[string]schema.Attribute{
	"access_key": schema.StringAttribute{
		Computed: true,
		PlanModifiers: []planmodifier.String{
			stringplanmodifier.UseStateForUnknown(),
		},
		MarkdownDescription: "Auto generated unique Bypass Code value",
	},
	"activation": schema.StringAttribute{
		Required:            true,
		CustomType:          timetypes.RFC3339Type{},
		MarkdownDescription: "The time when the Bypass Code object was activated.",
	},
	"created_time": schema.StringAttribute{
		Computed:            true,
		CustomType:          timetypes.RFC3339Type{},
		MarkdownDescription: "The time when the Bypass Code object was created.",
	},
	"description": schema.StringAttribute{
		Default:             stringdefault.StaticString(""),
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "",
	},
	"expiration": schema.StringAttribute{
		Required:            true,
		CustomType:          timetypes.RFC3339Type{},
		MarkdownDescription: "The time when the Bypass Code object was expired.",
	},
	"name": schema.StringAttribute{
		Required:            true,
		MarkdownDescription: "The name of Bypass Code",
	},
	"policy_ids": schema.ListAttribute{
		ElementType:         types.Int32Type,
		Computed:            true,
		MarkdownDescription: "The list of SecurityPolicy object identifiers.",
	},
	"rules": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: AccessCodeRuleResourceSchemaAttributes,
		},
		Required: true,
		Validators: []validator.List{
			customvalidator.ListNotEmpty(),
		},
		MarkdownDescription: "The list of selected security rules",
	},
	"updated_time": schema.StringAttribute{
		Computed:            true,
		CustomType:          timetypes.RFC3339Type{},
		MarkdownDescription: "The time when the Bypass Code object was last updated.",
	},
}

// Expand converts the TF model to the infoblox core model
func (m *AccessCodeModel) Expand(ctx context.Context, diags *diag.Diagnostics, isCreate bool) *coremodel.AccessCode {
	if m == nil {
		return nil
	}

	obj := &coremodel.AccessCode{}

	// Expand UDDI nested attribute (returns nil if not present)
	uddiModel := flex.ExpandNestedObject[UDDIAccessCodeModel](ctx, m.UDDI, diags)
	if uddiModel != nil {
		obj.UDDI = uddiModel.Expand(ctx, diags)
	}

	return obj
}

// Expand converts the UDDI TF model to the core model.
func (m *UDDIAccessCodeModel) Expand(ctx context.Context, diags *diag.Diagnostics) *coremodel.UDDIAccessCodeExt {
	return &coremodel.UDDIAccessCodeExt{
		AccessKey:   flex.ExpandStringPointer(m.AccessKey),
		Activation:  flex.ExpandRFC3339(m.Activation, diags),
		Description: flex.ExpandStringPointer(m.Description),
		Expiration:  flex.ExpandRFC3339(m.Expiration, diags),
		Name:        flex.ExpandStringPointer(m.Name),
		PolicyIds:   flex.ExpandFrameworkListInt32(m.PolicyIds),
		Rules:       flex.ExpandFrameworkListNestedBlock(ctx, m.Rules, diags, ExpandAccessCodeRule),
	}
}

// Flatten populates the TF model from a core response.
func (m *AccessCodeModel) Flatten(ctx context.Context, resp *coremodel.AccessCode, diags *diag.Diagnostics) {
	if resp == nil {
		return
	}

	m.Id = flex.FlattenStringPointer(resp.Id)

	// Extract existing UDDI model, flatten API response onto it, convert back
	uddiModel := flex.ExpandNestedObject[UDDIAccessCodeModel](ctx, m.UDDI, diags)
	if uddiModel == nil {
		uddiModel = &UDDIAccessCodeModel{}
	}
	uddiModel.Flatten(ctx, resp.UDDI, diags)
	if resp.UDDI != nil {
		m.UDDI = flex.FlattenNestedObject(ctx, uddiModel, UDDIAccessCodeAttrTypes, diags)
	} else {
		m.UDDI = types.ObjectNull(UDDIAccessCodeAttrTypes)
	}
}

// Flatten merges API response onto existing UDDI model.
func (m *UDDIAccessCodeModel) Flatten(ctx context.Context, from *coremodel.UDDIAccessCodeExt, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	m.AccessKey = flex.FlattenStringPointer(from.AccessKey)
	m.Activation = flex.FlattenRFC3339(from.Activation)
	m.CreatedTime = flex.FlattenRFC3339(from.CreatedTime)
	m.Description = flex.FlattenStringPointer(from.Description)
	m.Expiration = flex.FlattenRFC3339(from.Expiration)
	m.Name = flex.FlattenStringPointer(from.Name)
	m.PolicyIds = flex.FlattenFrameworkListInt32(from.PolicyIds)
	m.Rules = flex.FlattenFrameworkListNestedBlock(ctx, from.Rules, AccessCodeRuleAttrTypes, diags, FlattenAccessCodeRule)
	m.UpdatedTime = flex.FlattenRFC3339(from.UpdatedTime)
}
