package grid

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"

	coremodel "github.com/infobloxopen/terraform-provider-infoblox/internal/core/model/grid"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
	customvalidator "github.com/infobloxopen/terraform-provider-infoblox/internal/validator"
)

type NatgroupModel struct {
	Id   types.String `tfsdk:"id"`
	NIOS types.Object `tfsdk:"nios"`
}

var NatgroupAttrTypes = map[string]attr.Type{
	"id":   types.StringType,
	"nios": types.ObjectType{AttrTypes: NIOSNatgroupAttrTypes},
}

type NIOSNatgroupModel struct {
	Comment types.String `tfsdk:"comment"`
	Name    types.String `tfsdk:"name"`
}

var NIOSNatgroupAttrTypes = map[string]attr.Type{
	"comment": types.StringType,
	"name":    types.StringType,
}

const (
	NatgroupReturnFields = "comment,name"
)

var NatgroupResourceSchemaAttributes = map[string]schema.Attribute{
	"id": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The reference to the object.",
	},
	"nios": schema.SingleNestedAttribute{
		Optional:            true,
		MarkdownDescription: "NIOS backend-specific fields.",
		Attributes:          NatgroupResourceNiosSchemaAttributes,
	},
}

var NatgroupResourceNiosSchemaAttributes = map[string]schema.Attribute{
	"comment": schema.StringAttribute{
		Optional: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
		},
		MarkdownDescription: "The NAT group descriptive comment.",
	},
	"name": schema.StringAttribute{
		Optional: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
		},
		MarkdownDescription: "The name of a NAT group object.",
	},
}

// Expand converts the TF model to the infoblox core model
func (m *NatgroupModel) Expand(ctx context.Context, diags *diag.Diagnostics, isCreate bool) *coremodel.Natgroup {
	if m == nil {
		return nil
	}

	obj := &coremodel.Natgroup{}

	// Expand NIOS nested attribute (returns nil if not present)
	niosModel := flex.ExpandNestedObject[NIOSNatgroupModel](ctx, m.NIOS, diags)
	if niosModel != nil {
		obj.NIOS = niosModel.Expand(ctx, diags)
	}

	return obj
}

// Expand converts the NIOS TF model to the core model.
func (m *NIOSNatgroupModel) Expand(ctx context.Context, diags *diag.Diagnostics) *coremodel.NIOSNatgroupExt {
	return &coremodel.NIOSNatgroupExt{
		Comment: flex.ExpandStringPointerNullAsEmpty(m.Comment),
		Name:    flex.ExpandStringPointerNullAsEmpty(m.Name),
	}
}

// Flatten populates the TF model from a core response.
func (m *NatgroupModel) Flatten(ctx context.Context, resp *coremodel.Natgroup, diags *diag.Diagnostics) {
	if resp == nil {
		return
	}

	m.Id = flex.FlattenStringPointer(resp.Id)

	// Extract existing NIOS model, flatten API response onto it, convert back
	niosModel := flex.ExpandNestedObject[NIOSNatgroupModel](ctx, m.NIOS, diags)
	if niosModel == nil {
		niosModel = &NIOSNatgroupModel{}
	}
	niosModel.Flatten(ctx, resp.NIOS, diags)
	if resp.NIOS != nil {
		m.NIOS = flex.FlattenNestedObject(ctx, niosModel, NIOSNatgroupAttrTypes, diags)
	} else {
		m.NIOS = types.ObjectNull(NIOSNatgroupAttrTypes)
	}

}

// Flatten merges API response onto existing NIOS model.
func (m *NIOSNatgroupModel) Flatten(ctx context.Context, from *coremodel.NIOSNatgroupExt, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	m.Comment = flex.FlattenStringPointerEmptyAsNull(from.Comment)
	m.Name = flex.FlattenStringPointerEmptyAsNull(from.Name)
}
