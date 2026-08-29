package ipamfederation

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/mapvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/mapdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"

	coremodel "github.com/infobloxopen/terraform-provider-infoblox/internal/core/model/ipamfederation"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
)

type FederatedRealmModel struct {
	Id   types.String `tfsdk:"id"`
	UDDI types.Object `tfsdk:"uddi"`
}

var FederatedRealmAttrTypes = map[string]attr.Type{
	"id":   types.StringType,
	"uddi": types.ObjectType{AttrTypes: UDDIFederatedRealmAttrTypes},
}

type UDDIFederatedRealmModel struct {
	Comment types.String `tfsdk:"comment"`
	Name    types.String `tfsdk:"name"`
	Tags    types.Map    `tfsdk:"tags"`
	TagsAll types.Map    `tfsdk:"tags_all"`
}

var UDDIFederatedRealmAttrTypes = map[string]attr.Type{
	"comment":  types.StringType,
	"name":     types.StringType,
	"tags":     types.MapType{ElemType: types.StringType},
	"tags_all": types.MapType{ElemType: types.StringType},
}

const (
	FederatedRealmReturnFields = ""
)

var FederatedRealmResourceSchemaAttributes = map[string]schema.Attribute{
	"id": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The resource identifier.",
	},
	"uddi": schema.SingleNestedAttribute{
		Optional:            true,
		MarkdownDescription: "UDDI backend-specific fields.",
		Attributes:          FederatedRealmResourceUddiSchemaAttributes,
	},
}

var FederatedRealmResourceUddiSchemaAttributes = map[string]schema.Attribute{
	"comment": schema.StringAttribute{
		Default:  stringdefault.StaticString(""),
		Optional: true,
		Computed: true,
		Validators: []validator.String{
			stringvalidator.LengthBetween(0, 1024),
		},
		MarkdownDescription: "The description of the federated realm. May contain 0 to 1024 characters. Can include UTF-8.",
	},
	"name": schema.StringAttribute{
		Required: true,
		Validators: []validator.String{
			stringvalidator.LengthBetween(1, 256),
		},
		MarkdownDescription: "The name of the federated realm. May contain 1 to 256 characters; can include UTF-8.",
	},
	"tags": schema.MapAttribute{
		Optional:    true,
		Computed:    true,
		ElementType: types.StringType,
		Default:     mapdefault.StaticValue(types.MapNull(types.StringType)),
		Validators: []validator.Map{
			mapvalidator.SizeAtLeast(1),
		},
		MarkdownDescription: "The tags for the federated realm in JSON format.",
	},
	"tags_all": schema.MapAttribute{
		Computed:            true,
		ElementType:         types.StringType,
		MarkdownDescription: "All tags including inherited values.",
	},
}

// Expand converts the TF model to the infoblox core model
func (m *FederatedRealmModel) Expand(ctx context.Context, diags *diag.Diagnostics, isCreate bool) *coremodel.FederatedRealm {
	if m == nil {
		return nil
	}

	obj := &coremodel.FederatedRealm{}

	// Expand UDDI nested attribute (returns nil if not present)
	uddiModel := flex.ExpandNestedObject[UDDIFederatedRealmModel](ctx, m.UDDI, diags)
	if uddiModel != nil {
		obj.UDDI = uddiModel.Expand(ctx, diags)
	}

	return obj
}

// Expand converts the UDDI TF model to the core model.
func (m *UDDIFederatedRealmModel) Expand(ctx context.Context, diags *diag.Diagnostics) *coremodel.UDDIFederatedRealmExt {
	return &coremodel.UDDIFederatedRealmExt{
		Comment: flex.ExpandStringPointer(m.Comment),
		Name:    flex.ExpandString(m.Name),
		Tags:    flex.ExpandMapStringAny(ctx, m.Tags, diags),
	}
}

// Flatten populates the TF model from a core response.
func (m *FederatedRealmModel) Flatten(ctx context.Context, resp *coremodel.FederatedRealm, diags *diag.Diagnostics) {
	if resp == nil {
		return
	}

	m.Id = flex.FlattenStringPointer(resp.Id)

	// Extract existing UDDI model, flatten API response onto it, convert back
	uddiModel := flex.ExpandNestedObject[UDDIFederatedRealmModel](ctx, m.UDDI, diags)
	if uddiModel == nil {
		uddiModel = &UDDIFederatedRealmModel{}
	}
	uddiModel.Flatten(ctx, resp.UDDI, diags)
	if resp.UDDI != nil {
		m.UDDI = flex.FlattenNestedObject(ctx, uddiModel, UDDIFederatedRealmAttrTypes, diags)
	} else {
		m.UDDI = types.ObjectNull(UDDIFederatedRealmAttrTypes)
	}
}

// Flatten merges API response onto existing UDDI model.
func (m *UDDIFederatedRealmModel) Flatten(ctx context.Context, from *coremodel.UDDIFederatedRealmExt, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	m.Comment = flex.FlattenStringPointer(from.Comment)
	m.Name = flex.FlattenString(from.Name)
	tagsAll := flex.FlattenMapStringAny(ctx, from.Tags, diags)
	if m.Tags.IsNull() || m.Tags.IsUnknown() {
		m.Tags = tagsAll
	}
	m.TagsAll = tagsAll
}
