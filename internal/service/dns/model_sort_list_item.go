package dns

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
	planmod "github.com/infobloxopen/terraform-provider-infoblox/internal/planmodifiers"
	customvalidator "github.com/infobloxopen/terraform-provider-infoblox/internal/validator"
	uddidns "github.com/infobloxopen/universal-ddi-go-client/dnsconfig"
)

// SortListItemModel is the Terraform model for SortListItem
type SortListItemModel struct {
	Acl                 types.String `tfsdk:"acl"`
	Element             types.String `tfsdk:"element"`
	PrioritizedNetworks types.List   `tfsdk:"prioritized_networks"`
	Source              types.String `tfsdk:"source"`
}

// SortListItemAttrTypes contains the attribute types for SortListItemModel
var SortListItemAttrTypes = map[string]attr.Type{
	"acl":                  types.StringType,
	"element":              types.StringType,
	"prioritized_networks": types.ListType{ElemType: types.StringType},
	"source":               types.StringType,
}

// SortListItemResourceSchemaAttributes contains the schema attributes for SortListItemModel
var SortListItemResourceSchemaAttributes = map[string]schema.Attribute{
	"acl": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The resource identifier.",
	},
	"element": schema.StringAttribute{
		Required:            true,
		MarkdownDescription: "Type of element.  Allowed values:  * _any_,  * _ip_,  * _acl_,",
	},
	"prioritized_networks": schema.ListAttribute{
		ElementType: types.StringType,
		Optional:    true,
		Validators: []validator.List{
			customvalidator.ListNotEmpty(),
		},
		MarkdownDescription: "Optional. The prioritized networks. If empty, the value of _source_ or networks from _acl_ is used.",
	},
	"source": schema.StringAttribute{
		Optional: true,
		Computed: true,
		PlanModifiers: []planmodifier.String{
			planmod.UseEmptyStringForNull(),
		},
		MarkdownDescription: "Must be empty if _element_ is not _ip_.",
	},
}

// ExpandSortListItem converts a Terraform Object to SDK type
func ExpandSortListItem(ctx context.Context, o types.Object, diags *diag.Diagnostics) *uddidns.SortListItem {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m SortListItemModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

// Expand converts the Terraform model to SDK type
func (m *SortListItemModel) Expand(ctx context.Context, diags *diag.Diagnostics) *uddidns.SortListItem {
	if m == nil {
		return nil
	}
	to := &uddidns.SortListItem{
		Acl:                 flex.ExpandStringPointer(m.Acl),
		Element:             flex.ExpandString(m.Element),
		PrioritizedNetworks: flex.ExpandFrameworkListString(ctx, m.PrioritizedNetworks, diags),
		Source:              flex.ExpandStringPointer(m.Source),
	}
	return to
}

// FlattenSortListItem converts an SDK type to Terraform Object
func FlattenSortListItem(ctx context.Context, from *uddidns.SortListItem, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(SortListItemAttrTypes)
	}
	m := &SortListItemModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, SortListItemAttrTypes, m)
	diags.Append(d...)
	return t
}

// Flatten populates the Terraform model from SDK type
func (m *SortListItemModel) Flatten(ctx context.Context, from *uddidns.SortListItem, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	m.Acl = flex.FlattenStringPointer(from.Acl)
	m.Element = flex.FlattenString(from.Element)
	m.PrioritizedNetworks = flex.FlattenFrameworkListString(ctx, from.PrioritizedNetworks, diags)
	m.Source = flex.FlattenStringPointer(from.Source)
}
