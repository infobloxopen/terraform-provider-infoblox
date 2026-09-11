package fw

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
	uddifw "github.com/infobloxopen/universal-ddi-go-client/fw"
)

// ItemStructsModel is the Terraform model for ItemStructs
type ItemStructsModel struct {
	Description   types.String      `tfsdk:"description"`
	ExpiryTime    timetypes.RFC3339 `tfsdk:"expiry_time"`
	Item          types.String      `tfsdk:"item"`
	Status        types.String      `tfsdk:"status"`
	StatusDetails types.String      `tfsdk:"status_details"`
}

// ItemStructsAttrTypes contains the attribute types for ItemStructsModel
var ItemStructsAttrTypes = map[string]attr.Type{
	"description":    types.StringType,
	"expiry_time":    timetypes.RFC3339Type{},
	"item":           types.StringType,
	"status":         types.StringType,
	"status_details": types.StringType,
}

// ItemStructsResourceSchemaAttributes contains the schema attributes for ItemStructsModel
var ItemStructsResourceSchemaAttributes = map[string]schema.Attribute{
	"description": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The description of the item.",
	},
	"expiry_time": schema.StringAttribute{
		Optional:            true,
		CustomType:          timetypes.RFC3339Type{},
		MarkdownDescription: "The time at which this list item expires, as an RFC 3339 timestamp string. May be specified in any timezone. Unset (null) means no expiry.  Write semantics: - Insert/replace (POST /named_lists/{id}/items): set when present;   NULL when absent in payload. - Patch update (PATCH /named_lists/{id}/items, updated_items_described):   set when present; unchanged when absent in payload. Clearing an existing expiry_time via request field mask is not supported.",
	},
	"item": schema.StringAttribute{
		Required:            true,
		MarkdownDescription: "The data of the item.",
	},
	"status": schema.StringAttribute{
		Validators: []validator.String{
			stringvalidator.OneOf("ACTIVE", "INACTIVE"),
		},
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The status of the item. Applicable to TI domains only",
	},
	"status_details": schema.StringAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The status details of the item. Applicable to TI domains only",
	},
}

// ExpandItemStructs converts a Terraform Object to SDK type
func ExpandItemStructs(ctx context.Context, o types.Object, diags *diag.Diagnostics) *uddifw.ItemStructs {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m ItemStructsModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

// Expand converts the Terraform model to SDK type
func (m *ItemStructsModel) Expand(ctx context.Context, diags *diag.Diagnostics) *uddifw.ItemStructs {
	if m == nil {
		return nil
	}
	to := &uddifw.ItemStructs{
		Description:   flex.ExpandStringPointer(m.Description),
		ExpiryTime:    flex.ExpandRFC3339(m.ExpiryTime, diags),
		Item:          flex.ExpandStringPointer(m.Item),
		Status:        (*uddifw.ItemStructsItemStatus)(flex.ExpandStringPointer(m.Status)),
		StatusDetails: flex.ExpandStringPointer(m.StatusDetails),
	}
	return to
}

// FlattenItemStructs converts an SDK type to Terraform Object
func FlattenItemStructs(ctx context.Context, from *uddifw.ItemStructs, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(ItemStructsAttrTypes)
	}
	m := &ItemStructsModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, ItemStructsAttrTypes, m)
	diags.Append(d...)
	return t
}

// Flatten populates the Terraform model from SDK type
func (m *ItemStructsModel) Flatten(ctx context.Context, from *uddifw.ItemStructs, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	m.Description = flex.FlattenStringPointer(from.Description)
	m.ExpiryTime = flex.FlattenRFC3339(from.ExpiryTime)
	m.Item = flex.FlattenStringPointer(from.Item)
	m.Status = flex.FlattenStringPointer((*string)(from.Status))
	m.StatusDetails = flex.FlattenStringPointer(from.StatusDetails)
}
