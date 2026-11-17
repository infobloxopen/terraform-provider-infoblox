package ipamfederation

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/mapvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/mapdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	stringplanmodifier "github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"

	coremodel "github.com/infobloxopen/terraform-provider-infoblox/internal/core/model/ipamfederation"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/dynamicallocation"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
)

type ReservedBlockModel struct {
	Id            types.String `tfsdk:"id"`
	UpdateTrigger types.String `tfsdk:"update_trigger"`
	UDDI          types.Object `tfsdk:"uddi"`
}

var ReservedBlockAttrTypes = map[string]attr.Type{
	"id":             types.StringType,
	"update_trigger": types.StringType,
	"uddi":           types.ObjectType{AttrTypes: UDDIReservedBlockAttrTypes},
}

type UDDIReservedBlockModel struct {
	Address           types.String      `tfsdk:"address"`
	Cidr              types.Int64       `tfsdk:"cidr"`
	Comment           types.String      `tfsdk:"comment"`
	CreatedAt         timetypes.RFC3339 `tfsdk:"created_at"`
	FederatedPoolId   types.String      `tfsdk:"federated_pool_id"`
	FederatedRealm    types.String      `tfsdk:"federated_realm"`
	Metadata          types.Map         `tfsdk:"metadata"`
	Name              types.String      `tfsdk:"name"`
	NetworkCompliant  types.Bool        `tfsdk:"network_compliant"`
	Parent            types.String      `tfsdk:"parent"`
	Protocol          types.String      `tfsdk:"protocol"`
	Region            types.String      `tfsdk:"region"`
	Tags              types.Map         `tfsdk:"tags"`
	TagsAll           types.Map         `tfsdk:"tags_all"`
	UpdatedAt         timetypes.RFC3339 `tfsdk:"updated_at"`
	DynamicAllocation types.Object      `tfsdk:"dynamic_allocation"`
}

var UDDIReservedBlockAttrTypes = map[string]attr.Type{
	"address":            types.StringType,
	"cidr":               types.Int64Type,
	"comment":            types.StringType,
	"created_at":         timetypes.RFC3339Type{},
	"federated_pool_id":  types.StringType,
	"federated_realm":    types.StringType,
	"metadata":           types.MapType{ElemType: types.StringType},
	"name":               types.StringType,
	"network_compliant":  types.BoolType,
	"parent":             types.StringType,
	"protocol":           types.StringType,
	"region":             types.StringType,
	"tags":               types.MapType{ElemType: types.StringType},
	"tags_all":           types.MapType{ElemType: types.StringType},
	"updated_at":         timetypes.RFC3339Type{},
	"dynamic_allocation": types.ObjectType{AttrTypes: dynamicallocation.NextAvailableReservedBlockAttrTypes},
}

const (
	ReservedBlockReturnFields = ""
)

var ReservedBlockResourceSchemaAttributes = map[string]schema.Attribute{
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
		Attributes:          ReservedBlockResourceUddiSchemaAttributes,
	},
}

var ReservedBlockResourceUddiSchemaAttributes = map[string]schema.Attribute{
	"address": schema.StringAttribute{
		Optional: true,
		Computed: true,
		PlanModifiers: []planmodifier.String{
			stringplanmodifier.RequiresReplaceIfConfigured(),
		},
		Validators: []validator.String{
			stringvalidator.ExactlyOneOf(
				path.MatchRelative().AtParent().AtName("dynamic_allocation"),
			),
		},
		MarkdownDescription: "The address field in form “a.b.c.d/n” where the “/n” may be omitted. In this case, the CIDR value must be defined in the _cidr_ field. When reading, the _address_ field is always in the form “a.b.c.d”.",
	},
	"cidr": schema.Int64Attribute{
		Required:            true,
		MarkdownDescription: "The CIDR of the reserved block. This is required field, if _address_ does not specify it in its input.",
	},
	"comment": schema.StringAttribute{
		Default:  stringdefault.StaticString(""),
		Optional: true,
		Computed: true,
		Validators: []validator.String{
			stringvalidator.LengthBetween(0, 1024),
		},
		MarkdownDescription: "The description for the reserved block. May contain 0 to 1024 characters. Can include UTF-8.",
	},
	"created_at": schema.StringAttribute{
		Computed:            true,
		CustomType:          timetypes.RFC3339Type{},
		MarkdownDescription: "Time when the object has been created.",
	},
	"federated_pool_id": schema.StringAttribute{
		Optional: true,
		PlanModifiers: []planmodifier.String{
			stringplanmodifier.RequiresReplaceIfConfigured(),
		},
		MarkdownDescription: "The resource identifier.",
	},
	"federated_realm": schema.StringAttribute{
		Required:            true,
		MarkdownDescription: "The resource identifier.",
	},
	"metadata": schema.MapAttribute{
		ElementType:         types.StringType,
		Computed:            true,
		MarkdownDescription: "The metadata for the reserved block in JSON format.",
	},
	"name": schema.StringAttribute{
		Default:  stringdefault.StaticString(""),
		Optional: true,
		Computed: true,
		Validators: []validator.String{
			stringvalidator.LengthBetween(1, 256),
		},
		MarkdownDescription: "The name of the reserved block. May contain 1 to 256 characters. Can include UTF-8.",
	},
	"network_compliant": schema.BoolAttribute{
		Computed:            true,
		MarkdownDescription: "The compliance status of the reserved block, as determined by the federation service.",
	},
	"parent": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The resource identifier.",
	},
	"protocol": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The type of protocol of reserved block (_ip4_ or _ip6_).",
	},
	"region": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The region where the reserved block is located.",
	},
	"tags": schema.MapAttribute{
		Optional:    true,
		Computed:    true,
		ElementType: types.StringType,
		Default:     mapdefault.StaticValue(types.MapNull(types.StringType)),
		Validators: []validator.Map{
			mapvalidator.SizeAtLeast(1),
		},
		MarkdownDescription: "The tags for the reserved block in JSON format.",
	},
	"tags_all": schema.MapAttribute{
		Computed:            true,
		ElementType:         types.StringType,
		MarkdownDescription: "All tags including inherited values.",
	},
	"updated_at": schema.StringAttribute{
		Computed:            true,
		CustomType:          timetypes.RFC3339Type{},
		MarkdownDescription: "Time when the object has been updated. Equals to _created_at_ if not updated after creation.",
	},
	"dynamic_allocation": schema.SingleNestedAttribute{
		Attributes:          dynamicallocation.NextAvailableReservedBlockResourceSchemaAttributes,
		Optional:            true,
		MarkdownDescription: "Dynamically allocate the next available reserved block from a parent scope. Mutually exclusive with the static \"address\" field.",
	},
}

// Expand converts the TF model to the infoblox core model
func (m *ReservedBlockModel) Expand(ctx context.Context, diags *diag.Diagnostics, isCreate bool) *coremodel.ReservedBlock {
	if m == nil {
		return nil
	}

	obj := &coremodel.ReservedBlock{}

	// Expand UDDI nested attribute (returns nil if not present)
	uddiModel := flex.ExpandNestedObject[UDDIReservedBlockModel](ctx, m.UDDI, diags)
	if uddiModel != nil {
		obj.UDDI = uddiModel.Expand(ctx, diags, isCreate)
	}

	return obj
}

// Expand converts the UDDI TF model to the core model.
func (m *UDDIReservedBlockModel) Expand(ctx context.Context, diags *diag.Diagnostics, isCreate bool) *coremodel.UDDIReservedBlockExt {
	ext := &coremodel.UDDIReservedBlockExt{
		Cidr:            flex.ExpandInt64Pointer(m.Cidr),
		Comment:         flex.ExpandStringPointer(m.Comment),
		FederatedPoolId: flex.ExpandStringPointer(m.FederatedPoolId),
		FederatedRealm:  flex.ExpandString(m.FederatedRealm),
		Name:            flex.ExpandStringPointer(m.Name),
		Tags:            flex.ExpandMapStringAny(ctx, m.Tags, diags),
	}
	if isCreate {
		ext.Address = flex.ExpandStringPointer(m.Address)
		if alloc := BuildReservedBlockAllocation(ctx, m.DynamicAllocation, diags); alloc != nil {
			ext.Address = alloc
		}
	}
	return ext
}

// Flatten populates the TF model from a core response.
func (m *ReservedBlockModel) Flatten(ctx context.Context, resp *coremodel.ReservedBlock, diags *diag.Diagnostics) {
	if resp == nil {
		return
	}

	m.Id = flex.FlattenStringPointer(resp.Id)

	// Extract existing UDDI model, flatten API response onto it, convert back
	uddiModel := flex.ExpandNestedObject[UDDIReservedBlockModel](ctx, m.UDDI, diags)
	if uddiModel == nil {
		uddiModel = &UDDIReservedBlockModel{}
	}
	uddiModel.Flatten(ctx, resp.UDDI, diags)
	if resp.UDDI != nil {
		m.UDDI = flex.FlattenNestedObject(ctx, uddiModel, UDDIReservedBlockAttrTypes, diags)
	} else {
		m.UDDI = types.ObjectNull(UDDIReservedBlockAttrTypes)
	}
}

// Flatten merges API response onto existing UDDI model.
func (m *UDDIReservedBlockModel) Flatten(ctx context.Context, from *coremodel.UDDIReservedBlockExt, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	m.Address = flex.FlattenStringPointer(from.Address)
	m.Cidr = flex.FlattenInt64Pointer(from.Cidr)
	m.Comment = flex.FlattenStringPointer(from.Comment)
	m.CreatedAt = flex.FlattenRFC3339(from.CreatedAt)
	m.FederatedPoolId = flex.FlattenStringPointer(from.FederatedPoolId)
	m.FederatedRealm = flex.FlattenString(from.FederatedRealm)
	m.Metadata = flex.FlattenMapStringAny(ctx, from.Metadata, diags)
	m.Name = flex.FlattenStringPointer(from.Name)
	m.NetworkCompliant = flex.FlattenBoolPointer(from.NetworkCompliant)
	m.Parent = flex.FlattenStringPointer(from.Parent)
	m.Protocol = flex.FlattenStringPointer(from.Protocol)
	m.Region = flex.FlattenStringPointer(from.Region)
	tagsAll := flex.FlattenMapStringAny(ctx, from.Tags, diags)
	if m.Tags.IsNull() || m.Tags.IsUnknown() {
		m.Tags = tagsAll
	}
	m.TagsAll = tagsAll
	m.UpdatedAt = flex.FlattenRFC3339(from.UpdatedAt)
	if len(m.DynamicAllocation.AttributeTypes(ctx)) == 0 {
		m.DynamicAllocation = types.ObjectNull(dynamicallocation.NextAvailableReservedBlockAttrTypes)
	}
}
