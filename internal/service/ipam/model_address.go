package ipam

import (
	"context"

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

	coremodel "github.com/infobloxopen/terraform-provider-infoblox/internal/core/model/ipam"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/dynamicallocation"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
	customvalidator "github.com/infobloxopen/terraform-provider-infoblox/internal/validator"
)

type AddressModel struct {
	Id   types.String `tfsdk:"id"`
	UDDI types.Object `tfsdk:"uddi"`
}

var AddressAttrTypes = map[string]attr.Type{
	"id":   types.StringType,
	"uddi": types.ObjectType{AttrTypes: UDDIAddressAttrTypes},
}

type UDDIAddressModel struct {
	Address           types.String `tfsdk:"address"`
	Comment           types.String `tfsdk:"comment"`
	ExternalKeys      types.Map    `tfsdk:"external_keys"`
	Host              types.String `tfsdk:"host"`
	Hwaddr            types.String `tfsdk:"hwaddr"`
	Interface         types.String `tfsdk:"interface"`
	Names             types.List   `tfsdk:"names"`
	Range             types.String `tfsdk:"range"`
	Space             types.String `tfsdk:"space"`
	Tags              types.Map    `tfsdk:"tags"`
	TagsAll           types.Map    `tfsdk:"tags_all"`
	DynamicAllocation types.Object `tfsdk:"dynamic_allocation"`
}

var UDDIAddressAttrTypes = map[string]attr.Type{
	"address":            types.StringType,
	"comment":            types.StringType,
	"external_keys":      types.MapType{ElemType: types.StringType},
	"host":               types.StringType,
	"hwaddr":             types.StringType,
	"interface":          types.StringType,
	"names":              types.ListType{ElemType: types.ObjectType{AttrTypes: NameAttrTypes}},
	"range":              types.StringType,
	"space":              types.StringType,
	"tags":               types.MapType{ElemType: types.StringType},
	"tags_all":           types.MapType{ElemType: types.StringType},
	"dynamic_allocation": types.ObjectType{AttrTypes: dynamicallocation.NextAvailableAddressAttrTypes},
}

const (
	AddressReturnFields = ""
)

var AddressResourceSchemaAttributes = map[string]schema.Attribute{
	"id": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The resource identifier.",
	},
	"uddi": schema.SingleNestedAttribute{
		Optional:            true,
		MarkdownDescription: "UDDI backend-specific fields.",
		Attributes:          AddressResourceUddiSchemaAttributes,
	},
}

var AddressResourceUddiSchemaAttributes = map[string]schema.Attribute{
	"address": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Validators: []validator.String{
			stringvalidator.ExactlyOneOf(
				path.MatchRelative().AtParent().AtName("dynamic_allocation"),
			),
		},
		MarkdownDescription: "The address in form \"a.b.c.d\".",
	},
	"comment": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The description for the address object. May contain 0 to 1024 characters. Can include UTF-8.",
	},
	"external_keys": schema.MapAttribute{
		ElementType:         types.StringType,
		Optional:            true,
		MarkdownDescription: "The external keys (source key) for this address in JSON format.",
	},
	"host": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The resource identifier.",
	},
	"hwaddr": schema.StringAttribute{
		Default:             stringdefault.StaticString(""),
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The hardware address associated with this IP address.",
	},
	"interface": schema.StringAttribute{
		Default:             stringdefault.StaticString(""),
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The name of the network interface card (NIC) associated with the address, if any.",
	},
	"names": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: NameResourceSchemaAttributes,
		},
		Optional: true,
		Validators: []validator.List{
			customvalidator.ListNotEmpty(),
		},
		MarkdownDescription: "The list of all names associated with this address.",
	},
	"range": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The resource identifier.",
	},
	"space": schema.StringAttribute{
		Required: true,
		PlanModifiers: []planmodifier.String{
			stringplanmodifier.RequiresReplaceIfConfigured(),
		},
		MarkdownDescription: "The resource identifier.",
	},
	"tags": schema.MapAttribute{
		Optional:    true,
		Computed:    true,
		ElementType: types.StringType,
		Default:     mapdefault.StaticValue(types.MapNull(types.StringType)),
		Validators: []validator.Map{
			mapvalidator.SizeAtLeast(1),
		},
		MarkdownDescription: "The tags for this address in JSON format.",
	},
	"tags_all": schema.MapAttribute{
		Computed:            true,
		ElementType:         types.StringType,
		MarkdownDescription: "All tags including inherited values.",
	},
	"dynamic_allocation": schema.SingleNestedAttribute{
		Attributes:          dynamicallocation.NextAvailableAddressResourceSchemaAttributes,
		Optional:            true,
		MarkdownDescription: "Dynamically allocate the next available address from a parent scope. Mutually exclusive with the static \"address\" field.",
	},
}

// Expand converts the TF model to the infoblox core model
func (m *AddressModel) Expand(ctx context.Context, diags *diag.Diagnostics, isCreate bool) *coremodel.Address {
	if m == nil {
		return nil
	}

	obj := &coremodel.Address{}

	// Expand UDDI nested attribute (returns nil if not present)
	uddiModel := flex.ExpandNestedObject[UDDIAddressModel](ctx, m.UDDI, diags)
	if uddiModel != nil {
		obj.UDDI = uddiModel.Expand(ctx, diags, isCreate)
	}

	return obj
}

// Expand converts the UDDI TF model to the core model.
func (m *UDDIAddressModel) Expand(ctx context.Context, diags *diag.Diagnostics, isCreate bool) *coremodel.UDDIAddressExt {
	ext := &coremodel.UDDIAddressExt{
		Address:      flex.ExpandString(m.Address),
		Comment:      flex.ExpandStringPointer(m.Comment),
		ExternalKeys: flex.ExpandMapStringAny(ctx, m.ExternalKeys, diags),
		Host:         flex.ExpandStringPointer(m.Host),
		Hwaddr:       flex.ExpandStringPointer(m.Hwaddr),
		Interface:    flex.ExpandStringPointer(m.Interface),
		Names:        flex.ExpandFrameworkListNestedBlock(ctx, m.Names, diags, ExpandName),
		Range:        flex.ExpandStringPointer(m.Range),
		Space:        flex.ExpandStringPointer(m.Space),
		Tags:         flex.ExpandMapStringAny(ctx, m.Tags, diags),
	}
	if isCreate {
		if alloc := BuildAddressAllocation(ctx, m.DynamicAllocation, diags); alloc != nil {
			ext.Address = *alloc
		}
	}
	return ext
}

// Flatten populates the TF model from a core response.
func (m *AddressModel) Flatten(ctx context.Context, resp *coremodel.Address, diags *diag.Diagnostics) {
	if resp == nil {
		return
	}

	m.Id = flex.FlattenStringPointer(resp.Id)

	// Extract existing UDDI model, flatten API response onto it, convert back
	uddiModel := flex.ExpandNestedObject[UDDIAddressModel](ctx, m.UDDI, diags)
	if uddiModel == nil {
		uddiModel = &UDDIAddressModel{}
	}
	uddiModel.Flatten(ctx, resp.UDDI, diags)
	if resp.UDDI != nil {
		m.UDDI = flex.FlattenNestedObject(ctx, uddiModel, UDDIAddressAttrTypes, diags)
	} else {
		m.UDDI = types.ObjectNull(UDDIAddressAttrTypes)
	}
}

// Flatten merges API response onto existing UDDI model.
func (m *UDDIAddressModel) Flatten(ctx context.Context, from *coremodel.UDDIAddressExt, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	m.Address = flex.FlattenString(from.Address)
	m.Comment = flex.FlattenStringPointer(from.Comment)
	m.ExternalKeys = flex.FlattenMapStringAny(ctx, from.ExternalKeys, diags)
	m.Host = flex.FlattenStringPointer(from.Host)
	m.Hwaddr = flex.FlattenStringPointer(from.Hwaddr)
	m.Interface = flex.FlattenStringPointer(from.Interface)
	m.Names = flex.FlattenFrameworkListNestedBlock(ctx, from.Names, NameAttrTypes, diags, FlattenName)
	m.Range = flex.FlattenStringPointer(from.Range)
	m.Space = flex.FlattenStringPointer(from.Space)
	tagsAll := flex.FlattenMapStringAny(ctx, from.Tags, diags)
	if m.Tags.IsNull() || m.Tags.IsUnknown() {
		m.Tags = tagsAll
	}
	m.TagsAll = tagsAll
	if len(m.DynamicAllocation.AttributeTypes(ctx)) == 0 {
		m.DynamicAllocation = types.ObjectNull(dynamicallocation.NextAvailableAddressAttrTypes)
	}
}
