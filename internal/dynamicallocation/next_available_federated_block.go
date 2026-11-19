package dynamicallocation

import (
	"regexp"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

type NextAvailableFederatedBlockModel struct {
	NextAvailableId basetypes.StringValue `tfsdk:"next_available_id"`
}

var NextAvailableFederatedBlockAttrTypes = map[string]attr.Type{
	"next_available_id": basetypes.StringType{},
}

var NextAvailableFederatedBlockResourceSchemaAttributes = map[string]schema.Attribute{
	"next_available_id": schema.StringAttribute{
		Required: true,
		Validators: []validator.String{
			stringvalidator.RegexMatches(
				regexp.MustCompile(`^federation/federated_(block|pool)/[0-9a-f-].*$`),
				"must be the resource identifier of a Federated Block or a Federated Pool (e.g. \"federation/federated_block/<uuid>\").",
			),
		},
		PlanModifiers: []planmodifier.String{
			stringplanmodifier.RequiresReplace(),
		},
		MarkdownDescription: "The resource identifier of the parent Federated Block or Federated Pool under which the next available federated block should be allocated. Allocating directly from a Federated Pool is only supported for AWS-backed pools.",
	},
}

func (m NextAvailableFederatedBlockModel) Suffixed(suffix string) string {
	return m.NextAvailableId.ValueString() + suffix
}
