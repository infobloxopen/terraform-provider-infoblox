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

type NextAvailableOverlappingBlockModel struct {
	NextAvailableId basetypes.StringValue `tfsdk:"next_available_id"`
}

var NextAvailableOverlappingBlockAttrTypes = map[string]attr.Type{
	"next_available_id": basetypes.StringType{},
}

var NextAvailableOverlappingBlockResourceSchemaAttributes = map[string]schema.Attribute{
	"next_available_id": schema.StringAttribute{
		Required: true,
		Validators: []validator.String{
			stringvalidator.RegexMatches(
				regexp.MustCompile(`^federation/federated_block/[0-9a-f-].*$`),
				"must be the resource identifier of a Federated Block (e.g. \"federation/federated_block/<uuid>\").",
			),
		},
		PlanModifiers: []planmodifier.String{
			stringplanmodifier.RequiresReplace(),
		},
		MarkdownDescription: "The resource identifier of the parent Federated Block under which the next available overlapping block should be allocated.",
	},
}

func (m NextAvailableOverlappingBlockModel) Suffixed(suffix string) string {
	return m.NextAvailableId.ValueString() + suffix
}
