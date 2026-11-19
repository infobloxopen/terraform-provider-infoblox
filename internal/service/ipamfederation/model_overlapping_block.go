package ipamfederation

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-nettypes/iptypes"
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
)

type OverlappingBlockModel struct {
	Id            types.String `tfsdk:"id"`
	UpdateTrigger types.String `tfsdk:"update_trigger"`
	UDDI          types.Object `tfsdk:"uddi"`
}

var OverlappingBlockAttrTypes = map[string]attr.Type{
	"id":             types.StringType,
	"update_trigger": types.StringType,
	"uddi":           types.ObjectType{AttrTypes: UDDIOverlappingBlockAttrTypes},
}

type UDDIOverlappingBlockModel struct {
	Address          iptypes.IPAddress `tfsdk:"address"`
	Cidr             types.Int64       `tfsdk:"cidr"`
	Comment          types.String      `tfsdk:"comment"`
	FederatedPoolId  types.String      `tfsdk:"federated_pool_id"`
	FederatedRealm   types.String      `tfsdk:"federated_realm"`
	Name             types.String      `tfsdk:"name"`
	NetworkCompliant types.Bool        `tfsdk:"network_compliant"`
	Parent           types.String      `tfsdk:"parent"`
	Protocol         types.String      `tfsdk:"protocol"`
	Tags             types.Map         `tfsdk:"tags"`
	TagsAll          types.Map         `tfsdk:"tags_all"`
}

var UDDIOverlappingBlockAttrTypes = map[string]attr.Type{
	"address":           iptypes.IPAddressType{},
	"cidr":              types.Int64Type,
	"comment":           types.StringType,
	"federated_pool_id": types.StringType,
	"federated_realm":   types.StringType,
	"name":              types.StringType,
	"network_compliant": types.BoolType,
	"parent":            types.StringType,
	"protocol":          types.StringType,
	"tags":              types.MapType{ElemType: types.StringType},
	"tags_all":          types.MapType{ElemType: types.StringType},
}

const (
	OverlappingBlockReturnFields = ""
)

var OverlappingBlockResourceSchemaAttributes = map[string]schema.Attribute{
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
		Attributes:          OverlappingBlockResourceUddiSchemaAttributes,
	},
}

var OverlappingBlockResourceUddiSchemaAttributes = map[string]schema.Attribute{
	"address": schema.StringAttribute{
		Required:   true,
		CustomType: iptypes.IPAddressType{},
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
		MarkdownDescription: "The CIDR of the overlapping block. This is required, if _address_ does not specify it in its input.",
	},
	"comment": schema.StringAttribute{
		Default:  stringdefault.StaticString(""),
		Optional: true,
		Computed: true,
		Validators: []validator.String{
			stringvalidator.LengthBetween(0, 1024),
		},
		MarkdownDescription: "The description for the overlapping block. May contain 0 to 1024 characters. Can include UTF-8.",
	},
	"federated_pool_id": schema.StringAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The resource identifier.",
	},
	"federated_realm": schema.StringAttribute{
		Required:            true,
		MarkdownDescription: "The resource identifier.",
	},
	"name": schema.StringAttribute{
		Optional: true,
		Validators: []validator.String{
			stringvalidator.LengthBetween(1, 256),
		},
		MarkdownDescription: "The name of the overlapping block. May contain 1 to 256 characters. Can include UTF-8.",
	},
	"network_compliant": schema.BoolAttribute{
		Computed:            true,
		MarkdownDescription: "The compliance status of the overlapping block, as determined by the federation service.",
	},
	"parent": schema.StringAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The resource identifier.",
	},
	"protocol": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The type of protocol of overlapping block (_ip4_ or _ip6_).",
	},
	"tags": schema.MapAttribute{
		Optional:    true,
		Computed:    true,
		ElementType: types.StringType,
		Default:     mapdefault.StaticValue(types.MapNull(types.StringType)),
		Validators: []validator.Map{
			mapvalidator.SizeAtLeast(1),
		},
		MarkdownDescription: "The tags for the overlapping block in JSON format.",
	},
	"tags_all": schema.MapAttribute{
		Computed:            true,
		ElementType:         types.StringType,
		MarkdownDescription: "All tags including inherited values.",
	},
}

// Expand converts the TF model to the infoblox core model
func (m *OverlappingBlockModel) Expand(ctx context.Context, diags *diag.Diagnostics, isCreate bool) *coremodel.OverlappingBlock {
	if m == nil {
		return nil
	}

	obj := &coremodel.OverlappingBlock{}

	// Expand UDDI nested attribute (returns nil if not present)
	uddiModel := flex.ExpandNestedObject[UDDIOverlappingBlockModel](ctx, m.UDDI, diags)
	if uddiModel != nil {
		obj.UDDI = uddiModel.Expand(ctx, diags, isCreate)
	}

	return obj
}

// Expand converts the UDDI TF model to the core model.
func (m *UDDIOverlappingBlockModel) Expand(ctx context.Context, diags *diag.Diagnostics, isCreate bool) *coremodel.UDDIOverlappingBlockExt {
	ext := &coremodel.UDDIOverlappingBlockExt{
		Cidr:            flex.ExpandInt64Pointer(m.Cidr),
		Comment:         flex.ExpandStringPointer(m.Comment),
		FederatedPoolId: flex.ExpandStringPointer(m.FederatedPoolId),
		FederatedRealm:  flex.ExpandString(m.FederatedRealm),
		Name:            flex.ExpandStringPointer(m.Name),
		Parent:          flex.ExpandStringPointer(m.Parent),
		Tags:            flex.ExpandMapStringAny(ctx, m.Tags, diags),
	}
	if isCreate {
		ext.Address = flex.ExpandIPAddress(m.Address)
	}
	return ext
}

// Flatten populates the TF model from a core response.
func (m *OverlappingBlockModel) Flatten(ctx context.Context, resp *coremodel.OverlappingBlock, diags *diag.Diagnostics) {
	if resp == nil {
		return
	}

	m.Id = flex.FlattenStringPointer(resp.Id)

	// Extract existing UDDI model, flatten API response onto it, convert back
	uddiModel := flex.ExpandNestedObject[UDDIOverlappingBlockModel](ctx, m.UDDI, diags)
	if uddiModel == nil {
		uddiModel = &UDDIOverlappingBlockModel{}
	}
	uddiModel.Flatten(ctx, resp.UDDI, diags)
	if resp.UDDI != nil {
		m.UDDI = flex.FlattenNestedObject(ctx, uddiModel, UDDIOverlappingBlockAttrTypes, diags)
	} else {
		m.UDDI = types.ObjectNull(UDDIOverlappingBlockAttrTypes)
	}
}

// Flatten merges API response onto existing UDDI model.
func (m *UDDIOverlappingBlockModel) Flatten(ctx context.Context, from *coremodel.UDDIOverlappingBlockExt, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	m.Address = flex.FlattenIPAddress(from.Address)
	m.Cidr = flex.FlattenInt64Pointer(from.Cidr)
	m.Comment = flex.FlattenStringPointer(from.Comment)
	m.FederatedPoolId = flex.FlattenStringPointer(from.FederatedPoolId)
	m.FederatedRealm = flex.FlattenString(from.FederatedRealm)
	m.Name = flex.FlattenStringPointerEmptyAsNull(from.Name)
	m.NetworkCompliant = flex.FlattenBoolPointer(from.NetworkCompliant)
	m.Parent = flex.FlattenStringPointer(from.Parent)
	m.Protocol = flex.FlattenStringPointer(from.Protocol)
	tagsAll := flex.FlattenMapStringAny(ctx, from.Tags, diags)
	if m.Tags.IsNull() || m.Tags.IsUnknown() {
		m.Tags = tagsAll
	}
	m.TagsAll = tagsAll
}
