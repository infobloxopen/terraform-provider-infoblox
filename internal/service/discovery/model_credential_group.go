package discovery

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"

	coremodel "github.com/infobloxopen/terraform-provider-infoblox/internal/core/model/discovery"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
	customvalidator "github.com/infobloxopen/terraform-provider-infoblox/internal/validator"
)

type CredentialGroupModel struct {
	Id   types.String `tfsdk:"id"`
	NIOS types.Object `tfsdk:"nios"`
}

var CredentialGroupAttrTypes = map[string]attr.Type{
	"id":   types.StringType,
	"nios": types.ObjectType{AttrTypes: NIOSCredentialGroupAttrTypes},
}

type NIOSCredentialGroupModel struct {
	Name types.String `tfsdk:"name"`
}

var NIOSCredentialGroupAttrTypes = map[string]attr.Type{
	"name": types.StringType,
}

const (
	CredentialGroupReturnFields = "name"
)

var CredentialGroupResourceSchemaAttributes = map[string]schema.Attribute{
	"id": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The reference to the object.",
	},
	"nios": schema.SingleNestedAttribute{
		Optional:            true,
		MarkdownDescription: "NIOS backend-specific fields.",
		Attributes:          CredentialGroupResourceNiosSchemaAttributes,
	},
}

var CredentialGroupResourceNiosSchemaAttributes = map[string]schema.Attribute{
	"name": schema.StringAttribute{
		Required: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
			customvalidator.ValidateTrimmedString(),
		},
		MarkdownDescription: "The name of the Credential group.",
	},
}

// Expand converts the TF model to the infoblox core model
func (m *CredentialGroupModel) Expand(ctx context.Context, diags *diag.Diagnostics, isCreate bool) *coremodel.CredentialGroup {
	if m == nil {
		return nil
	}

	obj := &coremodel.CredentialGroup{}

	// Expand NIOS nested attribute (returns nil if not present)
	niosModel := flex.ExpandNestedObject[NIOSCredentialGroupModel](ctx, m.NIOS, diags)
	if niosModel != nil {
		obj.NIOS = niosModel.Expand(ctx, diags)
	}

	return obj
}

// Expand converts the NIOS TF model to the core model.
func (m *NIOSCredentialGroupModel) Expand(ctx context.Context, diags *diag.Diagnostics) *coremodel.NIOSCredentialGroupExt {
	return &coremodel.NIOSCredentialGroupExt{
		Name: flex.ExpandStringPointerNullAsEmpty(m.Name),
	}
}

// Flatten populates the TF model from a core response.
func (m *CredentialGroupModel) Flatten(ctx context.Context, resp *coremodel.CredentialGroup, diags *diag.Diagnostics) {
	if resp == nil {
		return
	}

	m.Id = flex.FlattenStringPointer(resp.Id)

	// Extract existing NIOS model, flatten API response onto it, convert back
	niosModel := flex.ExpandNestedObject[NIOSCredentialGroupModel](ctx, m.NIOS, diags)
	if niosModel == nil {
		niosModel = &NIOSCredentialGroupModel{}
	}
	niosModel.Flatten(ctx, resp.NIOS, diags)
	if resp.NIOS != nil {
		m.NIOS = flex.FlattenNestedObject(ctx, niosModel, NIOSCredentialGroupAttrTypes, diags)
	} else {
		m.NIOS = types.ObjectNull(NIOSCredentialGroupAttrTypes)
	}

}

// Flatten merges API response onto existing NIOS model.
func (m *NIOSCredentialGroupModel) Flatten(ctx context.Context, from *coremodel.NIOSCredentialGroupExt, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	m.Name = flex.FlattenStringPointerEmptyAsNull(from.Name)
}
