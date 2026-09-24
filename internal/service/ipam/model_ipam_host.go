package ipam

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/mapvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	listplanmodifier "github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/mapdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"

	coremodel "github.com/infobloxopen/terraform-provider-infoblox/internal/core/model/ipam"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
	customvalidator "github.com/infobloxopen/terraform-provider-infoblox/internal/validator"
)

type IpamHostModel struct {
	Id            types.String `tfsdk:"id"`
	UpdateTrigger types.String `tfsdk:"update_trigger"`
	UDDI          types.Object `tfsdk:"uddi"`
}

var IpamHostAttrTypes = map[string]attr.Type{
	"id":             types.StringType,
	"update_trigger": types.StringType,
	"uddi":           types.ObjectType{AttrTypes: UDDIIpamHostAttrTypes},
}

type UDDIIpamHostModel struct {
	Addresses           types.List   `tfsdk:"addresses"`
	AutoGenerateRecords types.Bool   `tfsdk:"auto_generate_records"`
	Comment             types.String `tfsdk:"comment"`
	HostNames           types.List   `tfsdk:"host_names"`
	Name                types.String `tfsdk:"name"`
	Tags                types.Map    `tfsdk:"tags"`
	TagsAll             types.Map    `tfsdk:"tags_all"`
}

var UDDIIpamHostAttrTypes = map[string]attr.Type{
	"addresses":             types.ListType{ElemType: types.ObjectType{AttrTypes: HostAddressAttrTypes}},
	"auto_generate_records": types.BoolType,
	"comment":               types.StringType,
	"host_names":            types.ListType{ElemType: types.ObjectType{AttrTypes: HostNameAttrTypes}},
	"name":                  types.StringType,
	"tags":                  types.MapType{ElemType: types.StringType},
	"tags_all":              types.MapType{ElemType: types.StringType},
}

const (
	IpamHostReturnFields = ""
)

var IpamHostResourceSchemaAttributes = map[string]schema.Attribute{
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
		Attributes:          IpamHostResourceUddiSchemaAttributes,
	},
}

var IpamHostResourceUddiSchemaAttributes = map[string]schema.Attribute{
	"addresses": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: HostAddressResourceSchemaAttributes,
		},
		Optional: true,
		PlanModifiers: []planmodifier.List{
			listplanmodifier.RequiresReplaceIfConfigured(),
		},
		Validators: []validator.List{
			customvalidator.ListNotEmpty(),
		},
		MarkdownDescription: "The list of all addresses associated with the IPAM host, which may be in different IP spaces.",
	},
	"auto_generate_records": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "This flag specifies if resource records have to be auto generated for the host.",
	},
	"comment": schema.StringAttribute{
		Default:             stringdefault.StaticString(""),
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The description for the IPAM host. May contain 0 to 1024 characters. Can include UTF-8.",
	},
	"host_names": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: HostNameResourceSchemaAttributes,
		},
		Optional: true,
		Validators: []validator.List{
			customvalidator.ListNotEmpty(),
		},
		MarkdownDescription: "The name records to be generated for the host.  This field is required if _auto_generate_records_ is true.",
	},
	"name": schema.StringAttribute{
		Required:            true,
		MarkdownDescription: "The name of the IPAM host. Must contain 1 to 256 characters. Can include UTF-8.",
	},
	"tags": schema.MapAttribute{
		Optional:    true,
		Computed:    true,
		ElementType: types.StringType,
		Default:     mapdefault.StaticValue(types.MapNull(types.StringType)),
		Validators: []validator.Map{
			mapvalidator.SizeAtLeast(1),
		},
		MarkdownDescription: "The tags for the IPAM host in JSON format.",
	},
	"tags_all": schema.MapAttribute{
		Computed:            true,
		ElementType:         types.StringType,
		MarkdownDescription: "All tags including inherited values.",
	},
}

// Expand converts the TF model to the infoblox core model
func (m *IpamHostModel) Expand(ctx context.Context, diags *diag.Diagnostics, isCreate bool) *coremodel.IpamHost {
	if m == nil {
		return nil
	}

	obj := &coremodel.IpamHost{}

	// Expand UDDI nested attribute (returns nil if not present)
	uddiModel := flex.ExpandNestedObject[UDDIIpamHostModel](ctx, m.UDDI, diags)
	if uddiModel != nil {
		obj.UDDI = uddiModel.Expand(ctx, diags)
	}

	return obj
}

// Expand converts the UDDI TF model to the core model.
func (m *UDDIIpamHostModel) Expand(ctx context.Context, diags *diag.Diagnostics) *coremodel.UDDIIpamHostExt {
	return &coremodel.UDDIIpamHostExt{
		Addresses:           flex.ExpandFrameworkListNestedBlock(ctx, m.Addresses, diags, ExpandHostAddress),
		AutoGenerateRecords: flex.ExpandBoolPointer(m.AutoGenerateRecords),
		Comment:             flex.ExpandStringPointer(m.Comment),
		HostNames:           flex.ExpandFrameworkListNestedBlock(ctx, m.HostNames, diags, ExpandHostName),
		Name:                flex.ExpandString(m.Name),
		Tags:                flex.ExpandMapStringAny(ctx, m.Tags, diags),
	}
}

// Flatten populates the TF model from a core response.
func (m *IpamHostModel) Flatten(ctx context.Context, resp *coremodel.IpamHost, diags *diag.Diagnostics) {
	if resp == nil {
		return
	}

	m.Id = flex.FlattenStringPointer(resp.Id)

	// Extract existing UDDI model, flatten API response onto it, convert back
	uddiModel := flex.ExpandNestedObject[UDDIIpamHostModel](ctx, m.UDDI, diags)
	if uddiModel == nil {
		uddiModel = &UDDIIpamHostModel{}
	}
	uddiModel.Flatten(ctx, resp.UDDI, diags)
	if resp.UDDI != nil {
		m.UDDI = flex.FlattenNestedObject(ctx, uddiModel, UDDIIpamHostAttrTypes, diags)
	} else {
		m.UDDI = types.ObjectNull(UDDIIpamHostAttrTypes)
	}
}

// Flatten merges API response onto existing UDDI model.
func (m *UDDIIpamHostModel) Flatten(ctx context.Context, from *coremodel.UDDIIpamHostExt, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	m.Addresses = flex.FlattenFrameworkListNestedBlock(ctx, from.Addresses, HostAddressAttrTypes, diags, FlattenHostAddress)
	m.AutoGenerateRecords = flex.FlattenBoolPointer(from.AutoGenerateRecords)
	m.Comment = flex.FlattenStringPointer(from.Comment)
	m.HostNames = flex.FlattenFrameworkListNestedBlock(ctx, from.HostNames, HostNameAttrTypes, diags, FlattenHostName)
	m.Name = flex.FlattenString(from.Name)
	tagsAll := flex.FlattenMapStringAny(ctx, from.Tags, diags)
	if m.Tags.IsNull() || m.Tags.IsUnknown() {
		m.Tags = tagsAll
	}
	m.TagsAll = tagsAll
}
