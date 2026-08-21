package dns

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	niosdns "github.com/infobloxopen/infoblox-nios-go-client/dns"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
	customvalidator "github.com/infobloxopen/terraform-provider-infoblox/internal/validator"
)

// NsgroupGridSecondariesModel is the Terraform model for NsgroupGridSecondaries
type NsgroupGridSecondariesModel struct {
	Name                     types.String `tfsdk:"name"`
	Stealth                  types.Bool   `tfsdk:"stealth"`
	GridReplicate            types.Bool   `tfsdk:"grid_replicate"`
	Lead                     types.Bool   `tfsdk:"lead"`
	PreferredPrimaries       types.List   `tfsdk:"preferred_primaries"`
	EnablePreferredPrimaries types.Bool   `tfsdk:"enable_preferred_primaries"`
}

// NsgroupGridSecondariesAttrTypes contains the attribute types for NsgroupGridSecondariesModel
var NsgroupGridSecondariesAttrTypes = map[string]attr.Type{
	"name":                       types.StringType,
	"stealth":                    types.BoolType,
	"grid_replicate":             types.BoolType,
	"lead":                       types.BoolType,
	"preferred_primaries":        types.ListType{ElemType: types.ObjectType{AttrTypes: NsgroupgridsecondariesPreferredPrimariesAttrTypes}},
	"enable_preferred_primaries": types.BoolType,
}

// NsgroupGridSecondariesResourceSchemaAttributes contains the schema attributes for NsgroupGridSecondariesModel
var NsgroupGridSecondariesResourceSchemaAttributes = map[string]schema.Attribute{
	"name": schema.StringAttribute{
		Required: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
		},
		MarkdownDescription: "The grid member name.",
	},
	"stealth": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "This flag governs whether the specified Grid member is in stealth mode or not. If set to True, the member is in stealth mode. This flag is ignored if the struct is specified as part of a stub zone.",
	},
	"grid_replicate": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The flag represents DNS zone transfers if set to False, and ID Grid Replication if set to True. This flag is ignored if the struct is specified as part of a stub zone or if it is set as grid_member in an authoritative zone.",
	},
	"lead": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "This flag controls whether the Grid lead secondary server performs zone transfers to non lead secondaries. This flag is ignored if the struct is specified as grid_member in an authoritative zone.",
	},
	"preferred_primaries": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: NsgroupgridsecondariesPreferredPrimariesResourceSchemaAttributes,
		},
		Optional: true,
		Computed: true,
		Validators: []validator.List{
			customvalidator.ListNotEmpty(),
		},
		MarkdownDescription: "The primary preference list with Grid member names and\\or External Server extserver structs for this member.",
	},
	"enable_preferred_primaries": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "This flag represents whether the preferred_primaries field values of this member are used.",
	},
}

// ExpandNsgroupGridSecondaries converts a Terraform Object to SDK type
func ExpandNsgroupGridSecondaries(ctx context.Context, o types.Object, diags *diag.Diagnostics) *niosdns.NsgroupGridSecondaries {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m NsgroupGridSecondariesModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

// Expand converts the Terraform model to SDK type
func (m *NsgroupGridSecondariesModel) Expand(ctx context.Context, diags *diag.Diagnostics) *niosdns.NsgroupGridSecondaries {
	if m == nil {
		return nil
	}
	to := &niosdns.NsgroupGridSecondaries{
		Name:                     flex.ExpandStringPointerNullAsEmpty(m.Name),
		Stealth:                  flex.ExpandBoolPointer(m.Stealth),
		GridReplicate:            flex.ExpandBoolPointer(m.GridReplicate),
		Lead:                     flex.ExpandBoolPointer(m.Lead),
		PreferredPrimaries:       flex.ExpandFrameworkListNestedBlock(ctx, m.PreferredPrimaries, diags, ExpandNsgroupgridsecondariesPreferredPrimaries),
		EnablePreferredPrimaries: flex.ExpandBoolPointer(m.EnablePreferredPrimaries),
	}
	return to
}

// FlattenNsgroupGridSecondaries converts an SDK type to Terraform Object
func FlattenNsgroupGridSecondaries(ctx context.Context, from *niosdns.NsgroupGridSecondaries, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(NsgroupGridSecondariesAttrTypes)
	}
	m := &NsgroupGridSecondariesModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, NsgroupGridSecondariesAttrTypes, m)
	diags.Append(d...)
	return t
}

// Flatten populates the Terraform model from SDK type
func (m *NsgroupGridSecondariesModel) Flatten(ctx context.Context, from *niosdns.NsgroupGridSecondaries, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	m.Name = flex.FlattenStringPointerEmptyAsNull(from.Name)
	m.Stealth = flex.FlattenBoolPointer(from.Stealth)
	m.GridReplicate = flex.FlattenBoolPointer(from.GridReplicate)
	m.Lead = flex.FlattenBoolPointer(from.Lead)
	m.PreferredPrimaries = flex.FlattenFrameworkListNestedBlock(ctx, from.PreferredPrimaries, NsgroupgridsecondariesPreferredPrimariesAttrTypes, diags, FlattenNsgroupgridsecondariesPreferredPrimaries)
	m.EnablePreferredPrimaries = flex.FlattenBoolPointer(from.EnablePreferredPrimaries)
}
