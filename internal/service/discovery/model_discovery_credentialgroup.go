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

type DiscoveryCredentialgroupModel struct {
	Id            types.String `tfsdk:"id"`
	UpdateTrigger types.String `tfsdk:"update_trigger"`
	NIOS          types.Object `tfsdk:"nios"`
}

var DiscoveryCredentialgroupAttrTypes = map[string]attr.Type{
	"id":             types.StringType,
	"update_trigger": types.StringType,
	"nios":           types.ObjectType{AttrTypes: NIOSDiscoveryCredentialgroupAttrTypes},
}

type NIOSDiscoveryCredentialgroupModel struct {
	Name types.String `tfsdk:"name"`
}

var NIOSDiscoveryCredentialgroupAttrTypes = map[string]attr.Type{
	"name": types.StringType,
}

const (
	DiscoveryCredentialgroupReturnFields = "name"
)

var DiscoveryCredentialgroupResourceSchemaAttributes = map[string]schema.Attribute{
	"id": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The reference to the object.",
	},
	"update_trigger": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "An arbitrary value used to trigger an update. Not sent to the API. Change it when Terraform reports no infrastructure changes.",
	},
	"nios": schema.SingleNestedAttribute{
		Optional:            true,
		MarkdownDescription: "NIOS backend-specific fields.",
		Attributes:          DiscoveryCredentialgroupResourceNiosSchemaAttributes,
	},
}

var DiscoveryCredentialgroupResourceNiosSchemaAttributes = map[string]schema.Attribute{
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
func (m *DiscoveryCredentialgroupModel) Expand(ctx context.Context, diags *diag.Diagnostics, isCreate bool) *coremodel.DiscoveryCredentialgroup {
	if m == nil {
		return nil
	}

	obj := &coremodel.DiscoveryCredentialgroup{}

	// Expand NIOS nested attribute (returns nil if not present)
	niosModel := flex.ExpandNestedObject[NIOSDiscoveryCredentialgroupModel](ctx, m.NIOS, diags)
	if niosModel != nil {
		obj.NIOS = niosModel.Expand(ctx, diags)
	}

	return obj
}

// Expand converts the NIOS TF model to the core model.
func (m *NIOSDiscoveryCredentialgroupModel) Expand(ctx context.Context, diags *diag.Diagnostics) *coremodel.NIOSDiscoveryCredentialgroupExt {
	return &coremodel.NIOSDiscoveryCredentialgroupExt{
		Name: flex.ExpandStringPointerNullAsEmpty(m.Name),
	}
}

// Flatten populates the TF model from a core response.
func (m *DiscoveryCredentialgroupModel) Flatten(ctx context.Context, resp *coremodel.DiscoveryCredentialgroup, diags *diag.Diagnostics) {
	if resp == nil {
		return
	}

	m.Id = flex.FlattenStringPointer(resp.Id)

	// Extract existing NIOS model, flatten API response onto it, convert back
	niosModel := flex.ExpandNestedObject[NIOSDiscoveryCredentialgroupModel](ctx, m.NIOS, diags)
	if niosModel == nil {
		niosModel = &NIOSDiscoveryCredentialgroupModel{}
	}
	niosModel.Flatten(ctx, resp.NIOS, diags)
	if resp.NIOS != nil {
		m.NIOS = flex.FlattenNestedObject(ctx, niosModel, NIOSDiscoveryCredentialgroupAttrTypes, diags)
	} else {
		m.NIOS = types.ObjectNull(NIOSDiscoveryCredentialgroupAttrTypes)
	}

}

// Flatten merges API response onto existing NIOS model.
func (m *NIOSDiscoveryCredentialgroupModel) Flatten(ctx context.Context, from *coremodel.NIOSDiscoveryCredentialgroupExt, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	m.Name = flex.FlattenStringPointerEmptyAsNull(from.Name)
}
