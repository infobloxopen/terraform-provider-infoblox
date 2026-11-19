package ipamfederation

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-nettypes/iptypes"
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
	immutable "github.com/infobloxopen/terraform-provider-infoblox/internal/planmodifiers/immutable"
)

type FederatedBlockModel struct {
	Id            types.String `tfsdk:"id"`
	UpdateTrigger types.String `tfsdk:"update_trigger"`
	UDDI          types.Object `tfsdk:"uddi"`
}

var FederatedBlockAttrTypes = map[string]attr.Type{
	"id":             types.StringType,
	"update_trigger": types.StringType,
	"uddi":           types.ObjectType{AttrTypes: UDDIFederatedBlockAttrTypes},
}

type UDDIFederatedBlockModel struct {
	Address           iptypes.IPAddress `tfsdk:"address"`
	AllocationV4      types.Object      `tfsdk:"allocation_v4"`
	Cidr              types.Int64       `tfsdk:"cidr"`
	Comment           types.String      `tfsdk:"comment"`
	CreatedAt         timetypes.RFC3339 `tfsdk:"created_at"`
	FederatedPoolId   types.String      `tfsdk:"federated_pool_id"`
	FederatedRealm    types.String      `tfsdk:"federated_realm"`
	Metadata          types.Map         `tfsdk:"metadata"`
	Name              types.String      `tfsdk:"name"`
	NetworkCompliance types.Object      `tfsdk:"network_compliance"`
	NetworkCompliant  types.Bool        `tfsdk:"network_compliant"`
	Parent            types.String      `tfsdk:"parent"`
	Protocol          types.String      `tfsdk:"protocol"`
	Region            types.String      `tfsdk:"region"`
	State             types.String      `tfsdk:"state"`
	Tags              types.Map         `tfsdk:"tags"`
	TagsAll           types.Map         `tfsdk:"tags_all"`
	UpdatedAt         timetypes.RFC3339 `tfsdk:"updated_at"`
	Utilization       types.Int64       `tfsdk:"utilization"`
	UtilizationV6     types.Object      `tfsdk:"utilization_v6"`
	DynamicAllocation types.Object      `tfsdk:"dynamic_allocation"`
}

var UDDIFederatedBlockAttrTypes = map[string]attr.Type{
	"address":            iptypes.IPAddressType{},
	"allocation_v4":      types.ObjectType{AttrTypes: AllocationAttrTypes},
	"cidr":               types.Int64Type,
	"comment":            types.StringType,
	"created_at":         timetypes.RFC3339Type{},
	"federated_pool_id":  types.StringType,
	"federated_realm":    types.StringType,
	"metadata":           types.MapType{ElemType: types.StringType},
	"name":               types.StringType,
	"network_compliance": types.ObjectType{AttrTypes: NetworkComplianceAttrTypes},
	"network_compliant":  types.BoolType,
	"parent":             types.StringType,
	"protocol":           types.StringType,
	"region":             types.StringType,
	"state":              types.StringType,
	"tags":               types.MapType{ElemType: types.StringType},
	"tags_all":           types.MapType{ElemType: types.StringType},
	"updated_at":         timetypes.RFC3339Type{},
	"utilization":        types.Int64Type,
	"utilization_v6":     types.ObjectType{AttrTypes: UtilizationV6AttrTypes},
	"dynamic_allocation": types.ObjectType{AttrTypes: dynamicallocation.NextAvailableFederatedBlockAttrTypes},
}

const (
	FederatedBlockReturnFields = ""
)

var FederatedBlockResourceSchemaAttributes = map[string]schema.Attribute{
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
		Attributes:          FederatedBlockResourceUddiSchemaAttributes,
	},
}

var FederatedBlockResourceUddiSchemaAttributes = map[string]schema.Attribute{
	"address": schema.StringAttribute{
		Optional:   true,
		Computed:   true,
		CustomType: iptypes.IPAddressType{},
		PlanModifiers: []planmodifier.String{
			immutable.ImmutableString(),
		},
		Validators: []validator.String{
			stringvalidator.ExactlyOneOf(
				path.MatchRelative().AtParent().AtName("dynamic_allocation"),
			),
		},
		MarkdownDescription: "The address field in form “a.b.c.d/n” where the “/n” may be omitted. In this case, the CIDR value must be defined in the _cidr_ field. When reading, the _address_ field is always in the form “a.b.c.d”.",
	},
	"allocation_v4": schema.SingleNestedAttribute{
		Attributes:          AllocationResourceSchemaAttributes,
		Computed:            true,
		MarkdownDescription: "The __Allocation__ tracks the distribution of Federated Blocks within each of the supported objects.",
	},
	"cidr": schema.Int64Attribute{
		Required:            true,
		MarkdownDescription: "The CIDR of the federated block. This is required, if _address_ does not specify it in its input.",
	},
	"comment": schema.StringAttribute{
		Default:  stringdefault.StaticString(""),
		Optional: true,
		Computed: true,
		Validators: []validator.String{
			stringvalidator.LengthBetween(0, 1024),
		},
		MarkdownDescription: "The description for the federated block. May contain 0 to 1024 characters. Can include UTF-8.",
	},
	"created_at": schema.StringAttribute{
		Computed:            true,
		CustomType:          timetypes.RFC3339Type{},
		MarkdownDescription: "Time when the object has been created.",
	},
	"federated_pool_id": schema.StringAttribute{
		Optional: true,
		Computed: true,
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
		MarkdownDescription: "The metadata for the federated block in JSON format.",
	},
	"name": schema.StringAttribute{
		Default:  stringdefault.StaticString(""),
		Optional: true,
		Computed: true,
		Validators: []validator.String{
			stringvalidator.LengthBetween(1, 256),
		},
		MarkdownDescription: "The name of the federated block. May contain 1 to 256 characters. Can include UTF-8.",
	},
	"network_compliance": schema.SingleNestedAttribute{
		Attributes:          NetworkComplianceResourceSchemaAttributes,
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The __NetworkCompliance__ represents the network compliance of federated pool and federated block.",
	},
	"network_compliant": schema.BoolAttribute{
		Computed:            true,
		MarkdownDescription: "Indicates if this block is compliant with its parent's network compliance policy. When false, a trouble dot should be displayed in the UI to indicate non-compliance.",
	},
	"parent": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The resource identifier.",
	},
	"protocol": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The type of protocol of federated block (_ip4_ or _ip6_).",
	},
	"region": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The region where the federated block is located.",
	},
	"state": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The state of the federated block.",
	},
	"tags": schema.MapAttribute{
		Optional:    true,
		Computed:    true,
		ElementType: types.StringType,
		Default:     mapdefault.StaticValue(types.MapNull(types.StringType)),
		Validators: []validator.Map{
			mapvalidator.SizeAtLeast(1),
		},
		MarkdownDescription: "The tags for the federated block in JSON format.",
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
	"utilization": schema.Int64Attribute{
		Computed:            true,
		MarkdownDescription: "The percentage of Federated Block utilization.",
	},
	"utilization_v6": schema.SingleNestedAttribute{
		Attributes:          UtilizationV6ResourceSchemaAttributes,
		Computed:            true,
		MarkdownDescription: "The __UtilizationV6__ represents IPv6 utilization metrics.",
	},
	"dynamic_allocation": schema.SingleNestedAttribute{
		Attributes:          dynamicallocation.NextAvailableFederatedBlockResourceSchemaAttributes,
		Optional:            true,
		MarkdownDescription: "Dynamically allocate the next available federated block from a parent scope. Mutually exclusive with the static \"address\" field.",
	},
}

// Expand converts the TF model to the infoblox core model
func (m *FederatedBlockModel) Expand(ctx context.Context, diags *diag.Diagnostics, isCreate bool) *coremodel.FederatedBlock {
	if m == nil {
		return nil
	}

	obj := &coremodel.FederatedBlock{}

	// Expand UDDI nested attribute (returns nil if not present)
	uddiModel := flex.ExpandNestedObject[UDDIFederatedBlockModel](ctx, m.UDDI, diags)
	if uddiModel != nil {
		obj.UDDI = uddiModel.Expand(ctx, diags, isCreate)
	}

	return obj
}

// Expand converts the UDDI TF model to the core model.
func (m *UDDIFederatedBlockModel) Expand(ctx context.Context, diags *diag.Diagnostics, isCreate bool) *coremodel.UDDIFederatedBlockExt {
	ext := &coremodel.UDDIFederatedBlockExt{
		Cidr:              flex.ExpandInt64Pointer(m.Cidr),
		Comment:           flex.ExpandStringPointer(m.Comment),
		FederatedPoolId:   flex.ExpandStringPointer(m.FederatedPoolId),
		FederatedRealm:    flex.ExpandString(m.FederatedRealm),
		Name:              flex.ExpandStringPointer(m.Name),
		NetworkCompliance: ExpandNetworkCompliance(ctx, m.NetworkCompliance, diags),
		Tags:              flex.ExpandMapStringAny(ctx, m.Tags, diags),
	}
	if isCreate {
		ext.Address = flex.ExpandIPAddress(m.Address)
		if alloc := BuildFederatedBlockAllocation(ctx, m.DynamicAllocation, diags); alloc != nil {
			ext.Address = alloc
		}
	}
	return ext
}

// Flatten populates the TF model from a core response.
func (m *FederatedBlockModel) Flatten(ctx context.Context, resp *coremodel.FederatedBlock, diags *diag.Diagnostics) {
	if resp == nil {
		return
	}

	m.Id = flex.FlattenStringPointer(resp.Id)

	// Extract existing UDDI model, flatten API response onto it, convert back
	uddiModel := flex.ExpandNestedObject[UDDIFederatedBlockModel](ctx, m.UDDI, diags)
	if uddiModel == nil {
		uddiModel = &UDDIFederatedBlockModel{}
	}
	uddiModel.Flatten(ctx, resp.UDDI, diags)
	if resp.UDDI != nil {
		m.UDDI = flex.FlattenNestedObject(ctx, uddiModel, UDDIFederatedBlockAttrTypes, diags)
	} else {
		m.UDDI = types.ObjectNull(UDDIFederatedBlockAttrTypes)
	}
}

// Flatten merges API response onto existing UDDI model.
func (m *UDDIFederatedBlockModel) Flatten(ctx context.Context, from *coremodel.UDDIFederatedBlockExt, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	m.Address = flex.FlattenIPAddress(from.Address)
	m.AllocationV4 = FlattenAllocation(ctx, from.AllocationV4, diags)
	m.Cidr = flex.FlattenInt64Pointer(from.Cidr)
	m.Comment = flex.FlattenStringPointer(from.Comment)
	m.CreatedAt = flex.FlattenRFC3339(from.CreatedAt)
	m.FederatedPoolId = flex.FlattenStringPointer(from.FederatedPoolId)
	m.FederatedRealm = flex.FlattenString(from.FederatedRealm)
	m.Metadata = flex.FlattenMapStringAny(ctx, from.Metadata, diags)
	m.Name = flex.FlattenStringPointer(from.Name)
	m.NetworkCompliance = FlattenNetworkCompliance(ctx, from.NetworkCompliance, diags)
	m.NetworkCompliant = flex.FlattenBoolPointer(from.NetworkCompliant)
	m.Parent = flex.FlattenStringPointer(from.Parent)
	m.Protocol = flex.FlattenStringPointer(from.Protocol)
	m.Region = flex.FlattenStringPointer(from.Region)
	m.State = flex.FlattenStringPointer(from.State)
	tagsAll := flex.FlattenMapStringAny(ctx, from.Tags, diags)
	if m.Tags.IsNull() || m.Tags.IsUnknown() {
		m.Tags = tagsAll
	}
	m.TagsAll = tagsAll
	m.UpdatedAt = flex.FlattenRFC3339(from.UpdatedAt)
	m.Utilization = flex.FlattenInt64Pointer(from.Utilization)
	m.UtilizationV6 = FlattenUtilizationV6(ctx, from.UtilizationV6, diags)
	if len(m.DynamicAllocation.AttributeTypes(ctx)) == 0 {
		m.DynamicAllocation = types.ObjectNull(dynamicallocation.NextAvailableFederatedBlockAttrTypes)
	}
}
