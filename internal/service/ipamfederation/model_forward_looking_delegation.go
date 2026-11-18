package ipamfederation

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/mapvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/mapdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	stringplanmodifier "github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"

	coremodel "github.com/infobloxopen/terraform-provider-infoblox/internal/core/model/ipamfederation"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
	internaltypes "github.com/infobloxopen/terraform-provider-infoblox/internal/types"
	customvalidator "github.com/infobloxopen/terraform-provider-infoblox/internal/validator"
)

type ForwardLookingDelegationModel struct {
	Id            types.String `tfsdk:"id"`
	UpdateTrigger types.String `tfsdk:"update_trigger"`
	UDDI          types.Object `tfsdk:"uddi"`
}

var ForwardLookingDelegationAttrTypes = map[string]attr.Type{
	"id":             types.StringType,
	"update_trigger": types.StringType,
	"uddi":           types.ObjectType{AttrTypes: UDDIForwardLookingDelegationAttrTypes},
}

type UDDIForwardLookingDelegationModel struct {
	Address         types.String                     `tfsdk:"address"`
	Cidr            types.Int64                      `tfsdk:"cidr"`
	Comment         types.String                     `tfsdk:"comment"`
	FederatedPoolId types.String                     `tfsdk:"federated_pool_id"`
	FederatedRealms internaltypes.UnorderedListValue `tfsdk:"federated_realms"`
	Name            types.String                     `tfsdk:"name"`
	Tags            types.Map                        `tfsdk:"tags"`
	TagsAll         types.Map                        `tfsdk:"tags_all"`
}

var UDDIForwardLookingDelegationAttrTypes = map[string]attr.Type{
	"address":           types.StringType,
	"cidr":              types.Int64Type,
	"comment":           types.StringType,
	"federated_pool_id": types.StringType,
	"federated_realms":  internaltypes.UnorderedListOfStringType,
	"name":              types.StringType,
	"tags":              types.MapType{ElemType: types.StringType},
	"tags_all":          types.MapType{ElemType: types.StringType},
}

const (
	ForwardLookingDelegationReturnFields = ""
)

var ForwardLookingDelegationResourceSchemaAttributes = map[string]schema.Attribute{
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
		Attributes:          ForwardLookingDelegationResourceUddiSchemaAttributes,
	},
}

var ForwardLookingDelegationResourceUddiSchemaAttributes = map[string]schema.Attribute{
	"address": schema.StringAttribute{
		Required: true,
		PlanModifiers: []planmodifier.String{
			stringplanmodifier.RequiresReplaceIfConfigured(),
		},
		MarkdownDescription: "The address field in form “a.b.c.d/n” where the “/n” may be omitted. In this case, the CIDR value must be defined in the _cidr_ field. When reading, the _address_ field is always in the form “a.b.c.d”.",
	},
	"cidr": schema.Int64Attribute{
		Optional: true,
		Computed: true,
		Validators: []validator.Int64{
			int64validator.Between(1, 128),
		},
		MarkdownDescription: "The CIDR of the delegation. This is required, if _address_ does not specify it in its input.",
	},
	"comment": schema.StringAttribute{
		Default:  stringdefault.StaticString(""),
		Optional: true,
		Computed: true,
		Validators: []validator.String{
			stringvalidator.LengthBetween(0, 1024),
		},
		MarkdownDescription: "The description for the delegation. May contain 0 to 1024 characters. Can include UTF-8.",
	},
	"federated_pool_id": schema.StringAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The resource identifier.",
	},
	"federated_realms": schema.ListAttribute{
		ElementType: types.StringType,
		Required:    true,
		CustomType:  internaltypes.UnorderedListOfStringType,
		Validators: []validator.List{
			customvalidator.ListNotEmpty(),
		},
		MarkdownDescription: "The resource identifier.",
	},
	"name": schema.StringAttribute{
		Optional: true,
		Validators: []validator.String{
			stringvalidator.LengthBetween(1, 256),
		},
		MarkdownDescription: "The name of the delegation. May contain 1 to 256 characters. Can include UTF-8.",
	},
	"tags": schema.MapAttribute{
		Optional:    true,
		Computed:    true,
		ElementType: types.StringType,
		Default:     mapdefault.StaticValue(types.MapNull(types.StringType)),
		Validators: []validator.Map{
			mapvalidator.SizeAtLeast(1),
		},
		MarkdownDescription: "The tags for the delegation in JSON format.",
	},
	"tags_all": schema.MapAttribute{
		Computed:            true,
		ElementType:         types.StringType,
		MarkdownDescription: "All tags including inherited values.",
	},
}

// Expand converts the TF model to the infoblox core model
func (m *ForwardLookingDelegationModel) Expand(ctx context.Context, diags *diag.Diagnostics, isCreate bool) *coremodel.ForwardLookingDelegation {
	if m == nil {
		return nil
	}

	obj := &coremodel.ForwardLookingDelegation{}

	// Expand UDDI nested attribute (returns nil if not present)
	uddiModel := flex.ExpandNestedObject[UDDIForwardLookingDelegationModel](ctx, m.UDDI, diags)
	if uddiModel != nil {
		obj.UDDI = uddiModel.Expand(ctx, diags, isCreate)
	}

	return obj
}

// Expand converts the UDDI TF model to the core model.
func (m *UDDIForwardLookingDelegationModel) Expand(ctx context.Context, diags *diag.Diagnostics, isCreate bool) *coremodel.UDDIForwardLookingDelegationExt {
	ext := &coremodel.UDDIForwardLookingDelegationExt{
		Cidr:            flex.ExpandInt64Pointer(m.Cidr),
		Comment:         flex.ExpandStringPointer(m.Comment),
		FederatedPoolId: flex.ExpandStringPointer(m.FederatedPoolId),
		FederatedRealms: flex.ExpandFrameworkListString(ctx, m.FederatedRealms, diags),
		Name:            flex.ExpandStringPointer(m.Name),
		Tags:            flex.ExpandMapStringAny(ctx, m.Tags, diags),
	}
	if isCreate {
		ext.Address = flex.ExpandStringPointer(m.Address)
	}
	return ext
}

// Flatten populates the TF model from a core response.
func (m *ForwardLookingDelegationModel) Flatten(ctx context.Context, resp *coremodel.ForwardLookingDelegation, diags *diag.Diagnostics) {
	if resp == nil {
		return
	}

	m.Id = flex.FlattenStringPointer(resp.Id)

	// Extract existing UDDI model, flatten API response onto it, convert back
	uddiModel := flex.ExpandNestedObject[UDDIForwardLookingDelegationModel](ctx, m.UDDI, diags)
	if uddiModel == nil {
		uddiModel = &UDDIForwardLookingDelegationModel{}
	}
	uddiModel.Flatten(ctx, resp.UDDI, diags)
	if resp.UDDI != nil {
		m.UDDI = flex.FlattenNestedObject(ctx, uddiModel, UDDIForwardLookingDelegationAttrTypes, diags)
	} else {
		m.UDDI = types.ObjectNull(UDDIForwardLookingDelegationAttrTypes)
	}
}

// Flatten merges API response onto existing UDDI model.
func (m *UDDIForwardLookingDelegationModel) Flatten(ctx context.Context, from *coremodel.UDDIForwardLookingDelegationExt, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	m.Address = flex.FlattenStringPointer(from.Address)
	m.Cidr = flex.FlattenInt64Pointer(from.Cidr)
	m.Comment = flex.FlattenStringPointer(from.Comment)
	m.FederatedPoolId = flex.FlattenStringPointer(from.FederatedPoolId)
	m.FederatedRealms = flex.FlattenFrameworkUnorderedListString(ctx, from.FederatedRealms, diags)
	m.Name = flex.FlattenStringPointerEmptyAsNull(from.Name)
	tagsAll := flex.FlattenMapStringAny(ctx, from.Tags, diags)
	if m.Tags.IsNull() || m.Tags.IsUnknown() {
		m.Tags = tagsAll
	}
	m.TagsAll = tagsAll
}
