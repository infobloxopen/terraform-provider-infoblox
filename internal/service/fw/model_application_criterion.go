package fw

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
	uddifw "github.com/infobloxopen/universal-ddi-go-client/fw"
)

// ApplicationCriterionModel is the Terraform model for ApplicationCriterion
type ApplicationCriterionModel struct {
	Category    types.String `tfsdk:"category"`
	Id          types.String `tfsdk:"id"`
	Name        types.String `tfsdk:"name"`
	Subcategory types.String `tfsdk:"subcategory"`
}

// ApplicationCriterionAttrTypes contains the attribute types for ApplicationCriterionModel
var ApplicationCriterionAttrTypes = map[string]attr.Type{
	"category":    types.StringType,
	"id":          types.StringType,
	"name":        types.StringType,
	"subcategory": types.StringType,
}

// ApplicationCriterionResourceSchemaAttributes contains the schema attributes for ApplicationCriterionModel
var ApplicationCriterionResourceSchemaAttributes = map[string]schema.Attribute{
	"category": schema.StringAttribute{
		Default:             stringdefault.StaticString(""),
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "",
	},
	"id": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "",
	},
	"name": schema.StringAttribute{
		Default:             stringdefault.StaticString(""),
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "Name for the application. Since the name of application is unique it may be used as alternate key for the application. The 'name' is used for import-export workflow and should be resolved to the 'id' before continue processing Create/Update operations.",
	},
	"subcategory": schema.StringAttribute{
		Default:             stringdefault.StaticString(""),
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "",
	},
}

// ExpandApplicationCriterion converts a Terraform Object to SDK type
func ExpandApplicationCriterion(ctx context.Context, o types.Object, diags *diag.Diagnostics) *uddifw.ApplicationCriterion {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m ApplicationCriterionModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

// Expand converts the Terraform model to SDK type
func (m *ApplicationCriterionModel) Expand(ctx context.Context, diags *diag.Diagnostics) *uddifw.ApplicationCriterion {
	if m == nil {
		return nil
	}
	to := &uddifw.ApplicationCriterion{
		Category:    flex.ExpandStringPointer(m.Category),
		Id:          flex.ExpandStringPointer(m.Id),
		Name:        flex.ExpandStringPointer(m.Name),
		Subcategory: flex.ExpandStringPointer(m.Subcategory),
	}
	return to
}

// FlattenApplicationCriterion converts an SDK type to Terraform Object
func FlattenApplicationCriterion(ctx context.Context, from *uddifw.ApplicationCriterion, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(ApplicationCriterionAttrTypes)
	}
	m := &ApplicationCriterionModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, ApplicationCriterionAttrTypes, m)
	diags.Append(d...)
	return t
}

// Flatten populates the Terraform model from SDK type
func (m *ApplicationCriterionModel) Flatten(ctx context.Context, from *uddifw.ApplicationCriterion, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	m.Category = flex.FlattenStringPointer(from.Category)
	m.Id = flex.FlattenStringPointer(from.Id)
	m.Name = flex.FlattenStringPointer(from.Name)
	m.Subcategory = flex.FlattenStringPointer(from.Subcategory)
}
