package dhcp

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/mapvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/mapdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"

	coremodel "github.com/infobloxopen/terraform-provider-infoblox/internal/core/model/dhcp"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
	customvalidator "github.com/infobloxopen/terraform-provider-infoblox/internal/validator"
)

type Ipv6DhcpOptionspaceModel struct {
	Id   types.String `tfsdk:"id"`
	NIOS types.Object `tfsdk:"nios"`
	UDDI types.Object `tfsdk:"uddi"`
}

var Ipv6DhcpOptionspaceAttrTypes = map[string]attr.Type{
	"id":   types.StringType,
	"nios": types.ObjectType{AttrTypes: NIOSIpv6DhcpOptionspaceAttrTypes},
	"uddi": types.ObjectType{AttrTypes: UDDIIpv6DhcpOptionspaceAttrTypes},
}

type NIOSIpv6DhcpOptionspaceModel struct {
	Comment          types.String `tfsdk:"comment"`
	EnterpriseNumber types.Int64  `tfsdk:"enterprise_number"`
	Name             types.String `tfsdk:"name"`
}

var NIOSIpv6DhcpOptionspaceAttrTypes = map[string]attr.Type{
	"comment":           types.StringType,
	"enterprise_number": types.Int64Type,
	"name":              types.StringType,
}

type UDDIIpv6DhcpOptionspaceModel struct {
	Comment  types.String `tfsdk:"comment"`
	Name     types.String `tfsdk:"name"`
	Protocol types.String `tfsdk:"protocol"`
	Tags     types.Map    `tfsdk:"tags"`
	TagsAll  types.Map    `tfsdk:"tags_all"`
}

var UDDIIpv6DhcpOptionspaceAttrTypes = map[string]attr.Type{
	"comment":  types.StringType,
	"name":     types.StringType,
	"protocol": types.StringType,
	"tags":     types.MapType{ElemType: types.StringType},
	"tags_all": types.MapType{ElemType: types.StringType},
}

const (
	Ipv6DhcpOptionspaceReturnFields = "comment,enterprise_number,name,option_definitions"
)

var Ipv6DhcpOptionspaceResourceSchemaAttributes = map[string]schema.Attribute{
	"id": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The reference to the object.",
	},
	"nios": schema.SingleNestedAttribute{
		Optional:            true,
		MarkdownDescription: "NIOS backend-specific fields.",
		Attributes:          Ipv6DhcpOptionspaceResourceNiosSchemaAttributes,
	},
	"uddi": schema.SingleNestedAttribute{
		Optional:            true,
		MarkdownDescription: "UDDI backend-specific fields.",
		Attributes:          Ipv6DhcpOptionspaceResourceUddiSchemaAttributes,
	},
}

var Ipv6DhcpOptionspaceResourceNiosSchemaAttributes = map[string]schema.Attribute{
	"comment": schema.StringAttribute{
		Optional: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
			stringvalidator.LengthBetween(0, 256),
		},
		MarkdownDescription: "A descriptive comment of a DHCP IPv6 option space object.",
	},
	"enterprise_number": schema.Int64Attribute{
		Required: true,
		Validators: []validator.Int64{
			int64validator.Between(0, 4294967295),
		},
		MarkdownDescription: "The enterprise number of a DHCP IPv6 option space object.",
	},
	"name": schema.StringAttribute{
		Required: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
		},
		MarkdownDescription: "The name of a DHCP IPv6 option space object.",
	},
}

var Ipv6DhcpOptionspaceResourceUddiSchemaAttributes = map[string]schema.Attribute{
	"comment": schema.StringAttribute{
		Default:  stringdefault.StaticString(""),
		Optional: true,
		Computed: true,
		Validators: []validator.String{
			stringvalidator.LengthBetween(0, 1024),
		},
		MarkdownDescription: "The description for the option space. May contain 0 to 1024 characters. Can include UTF-8.",
	},
	"name": schema.StringAttribute{
		Required: true,
		Validators: []validator.String{
			stringvalidator.LengthBetween(1, 256),
		},
		MarkdownDescription: "The name of the option space. Must contain 1 to 256 characters. Can include UTF-8.",
	},
	"protocol": schema.StringAttribute{
		Default:             stringdefault.StaticString("ip6"),
		Computed:            true,
		MarkdownDescription: "The type of protocol for the option space (_ip4_ or _ip6_).",
	},
	"tags": schema.MapAttribute{
		Optional:    true,
		Computed:    true,
		ElementType: types.StringType,
		Default:     mapdefault.StaticValue(types.MapNull(types.StringType)),
		Validators: []validator.Map{
			mapvalidator.SizeAtLeast(1),
		},
		MarkdownDescription: "The tags for the option space in JSON format.",
	},
	"tags_all": schema.MapAttribute{
		Computed:            true,
		ElementType:         types.StringType,
		MarkdownDescription: "All tags including inherited values.",
	},
}

// Expand converts the TF model to the infoblox core model
func (m *Ipv6DhcpOptionspaceModel) Expand(ctx context.Context, diags *diag.Diagnostics, isCreate bool) *coremodel.Ipv6DhcpOptionspace {
	if m == nil {
		return nil
	}

	obj := &coremodel.Ipv6DhcpOptionspace{}

	// Expand NIOS nested attribute (returns nil if not present)
	niosModel := flex.ExpandNestedObject[NIOSIpv6DhcpOptionspaceModel](ctx, m.NIOS, diags)
	if niosModel != nil {
		obj.NIOS = niosModel.Expand(ctx, diags)
	}

	// Expand UDDI nested attribute (returns nil if not present)
	uddiModel := flex.ExpandNestedObject[UDDIIpv6DhcpOptionspaceModel](ctx, m.UDDI, diags)
	if uddiModel != nil {
		obj.UDDI = uddiModel.Expand(ctx, diags, isCreate)
	}

	return obj
}

// Expand converts the NIOS TF model to the core model.
func (m *NIOSIpv6DhcpOptionspaceModel) Expand(ctx context.Context, diags *diag.Diagnostics) *coremodel.NIOSIpv6DhcpOptionspaceExt {
	return &coremodel.NIOSIpv6DhcpOptionspaceExt{
		Comment:          flex.ExpandStringPointerNullAsEmpty(m.Comment),
		EnterpriseNumber: flex.ExpandInt64Pointer(m.EnterpriseNumber),
		Name:             flex.ExpandStringPointerNullAsEmpty(m.Name),
	}
}

// Expand converts the UDDI TF model to the core model.
func (m *UDDIIpv6DhcpOptionspaceModel) Expand(ctx context.Context, diags *diag.Diagnostics, isCreate bool) *coremodel.UDDIIpv6DhcpOptionspaceExt {
	ext := &coremodel.UDDIIpv6DhcpOptionspaceExt{
		Comment: flex.ExpandStringPointer(m.Comment),
		Name:    flex.ExpandString(m.Name),
		Tags:    flex.ExpandMapStringAny(ctx, m.Tags, diags),
	}
	if isCreate {
		ext.Protocol = flex.ExpandStringPointer(m.Protocol)
	}
	return ext
}

// Flatten populates the TF model from a core response.
func (m *Ipv6DhcpOptionspaceModel) Flatten(ctx context.Context, resp *coremodel.Ipv6DhcpOptionspace, diags *diag.Diagnostics) {
	if resp == nil {
		return
	}

	m.Id = flex.FlattenStringPointer(resp.Id)

	// Extract existing NIOS model, flatten API response onto it, convert back
	niosModel := flex.ExpandNestedObject[NIOSIpv6DhcpOptionspaceModel](ctx, m.NIOS, diags)
	if niosModel == nil {
		niosModel = &NIOSIpv6DhcpOptionspaceModel{}
	}
	niosModel.Flatten(ctx, resp.NIOS, diags)
	if resp.NIOS != nil {
		m.NIOS = flex.FlattenNestedObject(ctx, niosModel, NIOSIpv6DhcpOptionspaceAttrTypes, diags)
	} else {
		m.NIOS = types.ObjectNull(NIOSIpv6DhcpOptionspaceAttrTypes)
	}

	// Extract existing UDDI model, flatten API response onto it, convert back
	uddiModel := flex.ExpandNestedObject[UDDIIpv6DhcpOptionspaceModel](ctx, m.UDDI, diags)
	if uddiModel == nil {
		uddiModel = &UDDIIpv6DhcpOptionspaceModel{}
	}
	uddiModel.Flatten(ctx, resp.UDDI, diags)
	if resp.UDDI != nil {
		m.UDDI = flex.FlattenNestedObject(ctx, uddiModel, UDDIIpv6DhcpOptionspaceAttrTypes, diags)
	} else {
		m.UDDI = types.ObjectNull(UDDIIpv6DhcpOptionspaceAttrTypes)
	}
}

// Flatten merges API response onto existing NIOS model.
func (m *NIOSIpv6DhcpOptionspaceModel) Flatten(ctx context.Context, from *coremodel.NIOSIpv6DhcpOptionspaceExt, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	m.Comment = flex.FlattenStringPointerEmptyAsNull(from.Comment)
	m.EnterpriseNumber = flex.FlattenInt64Pointer(from.EnterpriseNumber)
	m.Name = flex.FlattenStringPointerEmptyAsNull(from.Name)
}

// Flatten merges API response onto existing UDDI model.
func (m *UDDIIpv6DhcpOptionspaceModel) Flatten(ctx context.Context, from *coremodel.UDDIIpv6DhcpOptionspaceExt, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	m.Comment = flex.FlattenStringPointer(from.Comment)
	m.Name = flex.FlattenString(from.Name)
	m.Protocol = flex.FlattenStringPointer(from.Protocol)
	tagsAll := flex.FlattenMapStringAny(ctx, from.Tags, diags)
	if m.Tags.IsNull() || m.Tags.IsUnknown() {
		m.Tags = tagsAll
	}
	m.TagsAll = tagsAll
}
