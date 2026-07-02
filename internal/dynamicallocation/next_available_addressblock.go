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

type NextAvailableAddressBlockModel struct {
	NextAvailableId basetypes.StringValue `tfsdk:"next_available_id"`
}

var NextAvailableAddressBlockAttrTypes = map[string]attr.Type{
	"next_available_id": basetypes.StringType{},
}

var NextAvailableAddressBlockResourceSchemaAttributes = map[string]schema.Attribute{
	"next_available_id": schema.StringAttribute{
		Required: true,
		Validators: []validator.String{
			stringvalidator.RegexMatches(
				regexp.MustCompile(`^ipam/address_block/[0-9a-f-].*$`),
				"must be the resource identifier of an address block (e.g. \"ipam/address_block/<uuid>\").",
			),
		},
		PlanModifiers: []planmodifier.String{
			stringplanmodifier.RequiresReplace(),
		},
		MarkdownDescription: "The resource identifier of the address block from which the next available address block should be allocated.",
	},
}

func (m NextAvailableAddressBlockModel) Suffixed(suffix string) string {
	return m.NextAvailableId.ValueString() + suffix
}
