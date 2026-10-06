package infra

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/mapvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/mapdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	stringplanmodifier "github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"

	coremodel "github.com/infobloxopen/terraform-provider-infoblox/internal/core/model/infra"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
	customvalidator "github.com/infobloxopen/terraform-provider-infoblox/internal/validator"
)

type JoinTokenModel struct {
	Id            types.String `tfsdk:"id"`
	UpdateTrigger types.String `tfsdk:"update_trigger"`
	UDDI          types.Object `tfsdk:"uddi"`
}

var JoinTokenAttrTypes = map[string]attr.Type{
	"id":             types.StringType,
	"update_trigger": types.StringType,
	"uddi":           types.ObjectType{AttrTypes: UDDIJoinTokenAttrTypes},
}

type UDDIJoinTokenModel struct {
	Description types.String      `tfsdk:"description"`
	ExpiresAt   timetypes.RFC3339 `tfsdk:"expires_at"`
	JoinToken   types.String      `tfsdk:"join_token"`
	Name        types.String      `tfsdk:"name"`
	Tags        types.Map         `tfsdk:"tags"`
	TagsAll     types.Map         `tfsdk:"tags_all"`
	TokenId     types.String      `tfsdk:"token_id"`
}

var UDDIJoinTokenAttrTypes = map[string]attr.Type{
	"description": types.StringType,
	"expires_at":  timetypes.RFC3339Type{},
	"join_token":  types.StringType,
	"name":        types.StringType,
	"tags":        types.MapType{ElemType: types.StringType},
	"tags_all":    types.MapType{ElemType: types.StringType},
	"token_id":    types.StringType,
}

const (
	JoinTokenReturnFields = ""
)

var JoinTokenResourceSchemaAttributes = map[string]schema.Attribute{
	"id": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The resource identifier.",
	},
	"update_trigger": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "An arbitrary value used to trigger an update. Not sent to the API. Change it when Terraform reports no infrastructure changes.",
	},
	"uddi": schema.SingleNestedAttribute{
		Optional:            true,
		MarkdownDescription: "UDDI backend-specific fields.",
		Attributes:          JoinTokenResourceUddiSchemaAttributes,
	},
}

var JoinTokenResourceUddiSchemaAttributes = map[string]schema.Attribute{
	"description": schema.StringAttribute{
		Optional: true,
		PlanModifiers: []planmodifier.String{
			stringplanmodifier.RequiresReplaceIfConfigured(),
		},
		MarkdownDescription: "User-provided description for the join token.",
	},
	"expires_at": schema.StringAttribute{
		Optional:            true,
		CustomType:          timetypes.RFC3339Type{},
		MarkdownDescription: "Timestamp after which the join token cannot be used, in RFC3339 format.",
	},
	"join_token": schema.StringAttribute{
		Sensitive: true,
		Computed:  true,
		PlanModifiers: []planmodifier.String{
			stringplanmodifier.UseStateForUnknown(),
		},
		MarkdownDescription: "The actual join token secret. Returned only on create; never surfaced again by the API.",
	},
	"name": schema.StringAttribute{
		Required: true,
		PlanModifiers: []planmodifier.String{
			stringplanmodifier.RequiresReplaceIfConfigured(),
		},
		Validators: []validator.String{
			customvalidator.ValidateTrimmedString(),
		},
		MarkdownDescription: "User-provided name for the join token.",
	},
	"tags": schema.MapAttribute{
		Optional:    true,
		Computed:    true,
		ElementType: types.StringType,
		Default:     mapdefault.StaticValue(types.MapNull(types.StringType)),
		Validators: []validator.Map{
			mapvalidator.SizeAtLeast(1),
		},
		MarkdownDescription: "Tags associated with the join token. For valid values, see {tags:values}.",
	},
	"tags_all": schema.MapAttribute{
		Computed:            true,
		ElementType:         types.StringType,
		MarkdownDescription: "All tags including inherited values.",
	},
	"token_id": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "First half of the token, usable to identify the token in logs.",
	},
}

// Expand converts the TF model to the infoblox core model
func (m *JoinTokenModel) Expand(ctx context.Context, diags *diag.Diagnostics, isCreate bool) *coremodel.JoinToken {
	if m == nil {
		return nil
	}

	obj := &coremodel.JoinToken{}

	// Expand UDDI nested attribute (returns nil if not present)
	uddiModel := flex.ExpandNestedObject[UDDIJoinTokenModel](ctx, m.UDDI, diags)
	if uddiModel != nil {
		obj.UDDI = uddiModel.Expand(ctx, diags, isCreate)
	}

	return obj
}

// Expand converts the UDDI TF model to the core model.
func (m *UDDIJoinTokenModel) Expand(ctx context.Context, diags *diag.Diagnostics, isCreate bool) *coremodel.UDDIJoinTokenExt {
	ext := &coremodel.UDDIJoinTokenExt{
		ExpiresAt: flex.ExpandRFC3339(m.ExpiresAt, diags),
		JoinToken: flex.ExpandStringPointer(m.JoinToken),
		Tags:      flex.ExpandMapStringAny(ctx, m.Tags, diags),
		TokenId:   flex.ExpandStringPointer(m.TokenId),
	}
	if isCreate {
		ext.Description = flex.ExpandStringPointer(m.Description)
		ext.Name = flex.ExpandStringPointer(m.Name)
	}
	return ext
}

// Flatten populates the TF model from a core response.
func (m *JoinTokenModel) Flatten(ctx context.Context, resp *coremodel.JoinToken, diags *diag.Diagnostics) {
	if resp == nil {
		return
	}

	m.Id = flex.FlattenStringPointer(resp.Id)

	// Extract existing UDDI model, flatten API response onto it, convert back
	uddiModel := flex.ExpandNestedObject[UDDIJoinTokenModel](ctx, m.UDDI, diags)
	if uddiModel == nil {
		uddiModel = &UDDIJoinTokenModel{}
	}
	plannedUDDI := flex.ExpandNestedObject[UDDIJoinTokenModel](ctx, m.UDDI, diags)
	uddiModel.Flatten(ctx, resp.UDDI, diags)
	if resp.UDDI != nil {
		PostFlattenJoinTokenUDDI(ctx, plannedUDDI, uddiModel, diags)
		m.UDDI = flex.FlattenNestedObject(ctx, uddiModel, UDDIJoinTokenAttrTypes, diags)
	} else {
		m.UDDI = types.ObjectNull(UDDIJoinTokenAttrTypes)
	}
}

// Flatten merges API response onto existing UDDI model.
func (m *UDDIJoinTokenModel) Flatten(ctx context.Context, from *coremodel.UDDIJoinTokenExt, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	m.Description = flex.FlattenStringPointer(from.Description)
	m.ExpiresAt = flex.FlattenRFC3339(from.ExpiresAt)
	m.JoinToken = flex.FlattenStringPointer(from.JoinToken)
	m.Name = flex.FlattenStringPointer(from.Name)
	tagsAll := flex.FlattenMapStringAny(ctx, from.Tags, diags)
	if m.Tags.IsNull() || m.Tags.IsUnknown() {
		m.Tags = tagsAll
	}
	m.TagsAll = tagsAll
	m.TokenId = flex.FlattenStringPointer(from.TokenId)
}
