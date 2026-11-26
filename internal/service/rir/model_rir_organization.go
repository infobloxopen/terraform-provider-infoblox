package rir

import (
	"context"
	"regexp"

	"github.com/hashicorp/terraform-plugin-framework-validators/mapvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/mapdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"

	coremodel "github.com/infobloxopen/terraform-provider-infoblox/internal/core/model/rir"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
	importmod "github.com/infobloxopen/terraform-provider-infoblox/internal/planmodifiers/import"
	customvalidator "github.com/infobloxopen/terraform-provider-infoblox/internal/validator"
)

type RirOrganizationModel struct {
	Id            types.String `tfsdk:"id"`
	UpdateTrigger types.String `tfsdk:"update_trigger"`
	NIOS          types.Object `tfsdk:"nios"`
}

var RirOrganizationAttrTypes = map[string]attr.Type{
	"id":             types.StringType,
	"update_trigger": types.StringType,
	"nios":           types.ObjectType{AttrTypes: NIOSRirOrganizationAttrTypes},
}

type NIOSRirOrganizationModel struct {
	ExtAttrs    types.Map    `tfsdk:"ext_attrs"`
	ExtAttrsAll types.Map    `tfsdk:"ext_attrs_all"`
	Id          types.String `tfsdk:"id"`
	Maintainer  types.String `tfsdk:"maintainer"`
	Name        types.String `tfsdk:"name"`
	Password    types.String `tfsdk:"password"`
	Rir         types.String `tfsdk:"rir"`
	SenderEmail types.String `tfsdk:"sender_email"`
}

var NIOSRirOrganizationAttrTypes = map[string]attr.Type{
	"ext_attrs":     types.MapType{ElemType: types.StringType},
	"ext_attrs_all": types.MapType{ElemType: types.StringType},
	"id":            types.StringType,
	"maintainer":    types.StringType,
	"name":          types.StringType,
	"password":      types.StringType,
	"rir":           types.StringType,
	"sender_email":  types.StringType,
}

const (
	RirOrganizationReturnFields = "extattrs,id,maintainer,name,rir,sender_email"
)

var RirOrganizationResourceSchemaAttributes = map[string]schema.Attribute{
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
		Attributes:          RirOrganizationResourceNiosSchemaAttributes,
	},
}

var RirOrganizationResourceNiosSchemaAttributes = map[string]schema.Attribute{
	"ext_attrs": schema.MapAttribute{
		Optional:    true,
		Computed:    true,
		ElementType: types.StringType,
		Default:     mapdefault.StaticValue(types.MapNull(types.StringType)),
		Validators: []validator.Map{
			mapvalidator.SizeAtLeast(1),
		},
		MarkdownDescription: "Extensible attributes associated with the object. For valid values for extensible attributes, see {extattrs:values}.",
	},
	"ext_attrs_all": schema.MapAttribute{
		Computed:            true,
		ElementType:         types.StringType,
		MarkdownDescription: "All ext_attrs including Terraform Internal ID and inherited attributes.",
		PlanModifiers: []planmodifier.Map{
			importmod.AssociateInternalId(),
		},
	},
	"id": schema.StringAttribute{
		Required: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
			stringvalidator.RegexMatches(regexp.MustCompile(`^ORG-[A-Za-z]{2,4}[1-9][0-9]{0,4}-[A-Za-z0-9]{1,9}$`), "- Invalid Organization ID. A Valid Organization ID starts with 'ORG-', followed by 2-4 letters, then a number between 1 and 99999, and ends with a hyphen and 1-9 alphanumeric characters. Valid Examples for ID are ORG-CA1-RIPE or ORG-CB2-TEST"),
		},
		MarkdownDescription: "The RIR organization identifier.",
	},
	"maintainer": schema.StringAttribute{
		Required: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
			stringvalidator.LengthBetween(0, 80),
			stringvalidator.RegexMatches(regexp.MustCompile(`^[A-Za-z][A-Za-z0-9_-]*[A-Za-z0-9]$`), "- A valid maintainer starts with a letter, followed by letters, numbers, underscores, or hyphens, and ends with a letter or number. Valid examples for maintainer are 'infoblox' and 'nios-support'"),
		},
		MarkdownDescription: "The RIR organization maintainer.",
	},
	"name": schema.StringAttribute{
		Required: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
			customvalidator.ValidateTrimmedString(),
			stringvalidator.LengthBetween(0, 256),
			stringvalidator.RegexMatches(regexp.MustCompile(`^[A-Za-z0-9_-]+$`), "- Invalid Organization Name. A valid organization name can only contain letters, numbers, underscores, or hyphens."),
		},
		MarkdownDescription: "The RIR organization name.",
	},
	"password": schema.StringAttribute{
		Sensitive: true,
		Optional:  true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
			stringvalidator.LengthBetween(0, 256),
			customvalidator.ValidateTrimmedString(),
		},
		MarkdownDescription: "The password for the maintainer of RIR organization.",
	},
	"rir": schema.StringAttribute{
		Default: stringdefault.StaticString("RIPE"),
		Validators: []validator.String{
			stringvalidator.OneOf("RIPE"),
		},
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The RIR associated with RIR organization.",
	},
	"sender_email": schema.StringAttribute{
		Required: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
			stringvalidator.RegexMatches(regexp.MustCompile(`^[^@]+@[^@]+\.com$`), "- must be a valid .com email address"),
		},
		MarkdownDescription: "The sender e-mail address for RIR organization.",
	},
}

// Expand converts the TF model to the infoblox core model
func (m *RirOrganizationModel) Expand(ctx context.Context, diags *diag.Diagnostics, isCreate bool) *coremodel.RirOrganization {
	if m == nil {
		return nil
	}

	obj := &coremodel.RirOrganization{}

	// Expand NIOS nested attribute (returns nil if not present)
	niosModel := flex.ExpandNestedObject[NIOSRirOrganizationModel](ctx, m.NIOS, diags)
	if niosModel != nil {
		obj.NIOS = niosModel.Expand(ctx, diags)
	}

	return obj
}

// Expand converts the NIOS TF model to the core model.
func (m *NIOSRirOrganizationModel) Expand(ctx context.Context, diags *diag.Diagnostics) *coremodel.NIOSRirOrganizationExt {
	return &coremodel.NIOSRirOrganizationExt{
		ExtAttrs:    flex.ExpandMapStringAny(ctx, m.ExtAttrs, diags),
		Id:          flex.ExpandStringPointerNullAsEmpty(m.Id),
		Maintainer:  flex.ExpandStringPointerNullAsEmpty(m.Maintainer),
		Name:        flex.ExpandStringPointerNullAsEmpty(m.Name),
		Password:    flex.ExpandStringPointerNullAsEmpty(m.Password),
		Rir:         flex.ExpandStringPointerNullAsEmpty(m.Rir),
		SenderEmail: flex.ExpandStringPointerNullAsEmpty(m.SenderEmail),
	}
}

// Flatten populates the TF model from a core response.
func (m *RirOrganizationModel) Flatten(ctx context.Context, resp *coremodel.RirOrganization, diags *diag.Diagnostics) {
	if resp == nil {
		return
	}

	m.Id = flex.FlattenStringPointer(resp.Id)

	// Extract existing NIOS model, flatten API response onto it, convert back
	niosModel := flex.ExpandNestedObject[NIOSRirOrganizationModel](ctx, m.NIOS, diags)
	if niosModel == nil {
		niosModel = &NIOSRirOrganizationModel{}
	}
	plannedNIOS := flex.ExpandNestedObject[NIOSRirOrganizationModel](ctx, m.NIOS, diags)
	niosModel.Flatten(ctx, resp.NIOS, diags)
	if resp.NIOS != nil {
		PostFlattenRirOrganizationNIOS(ctx, plannedNIOS, niosModel, diags)
		m.NIOS = flex.FlattenNestedObject(ctx, niosModel, NIOSRirOrganizationAttrTypes, diags)
	} else {
		m.NIOS = types.ObjectNull(NIOSRirOrganizationAttrTypes)
	}

}

// Flatten merges API response onto existing NIOS model.
func (m *NIOSRirOrganizationModel) Flatten(ctx context.Context, from *coremodel.NIOSRirOrganizationExt, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	planExtAttrs := m.ExtAttrs
	if planExtAttrs.IsUnknown() {
		planExtAttrs = types.MapNull(types.StringType)
	}
	m.ExtAttrs, m.ExtAttrsAll = flex.FlattenEAs(planExtAttrs, from.ExtAttrs)
	m.Id = flex.FlattenStringPointerEmptyAsNull(from.Id)
	m.Maintainer = flex.FlattenStringPointerEmptyAsNull(from.Maintainer)
	m.Name = flex.FlattenStringPointerEmptyAsNull(from.Name)
	m.Password = flex.FlattenStringPointerEmptyAsNull(from.Password)
	m.Rir = flex.FlattenStringPointerEmptyAsNull(from.Rir)
	m.SenderEmail = flex.FlattenStringPointerEmptyAsNull(from.SenderEmail)
}
