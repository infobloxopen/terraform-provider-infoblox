package dns

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	niosdns "github.com/infobloxopen/infoblox-nios-go-client/dns"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
	customvalidator "github.com/infobloxopen/terraform-provider-infoblox/internal/validator"
)

// ZoneRpGridSecondariesModel is the Terraform model for ZoneRpGridSecondaries
type ZoneRpGridSecondariesModel struct {
	Name                     types.String `tfsdk:"name"`
	Stealth                  types.Bool   `tfsdk:"stealth"`
	GridReplicate            types.Bool   `tfsdk:"grid_replicate"`
	Lead                     types.Bool   `tfsdk:"lead"`
	PreferredPrimaries       types.List   `tfsdk:"preferred_primaries"`
	EnablePreferredPrimaries types.Bool   `tfsdk:"enable_preferred_primaries"`
}

// ZoneRpGridSecondariesAttrTypes contains the attribute types for ZoneRpGridSecondariesModel
var ZoneRpGridSecondariesAttrTypes = map[string]attr.Type{
	"name":                       types.StringType,
	"stealth":                    types.BoolType,
	"grid_replicate":             types.BoolType,
	"lead":                       types.BoolType,
	"preferred_primaries":        types.ListType{ElemType: types.ObjectType{AttrTypes: ZonerpgridsecondariesPreferredPrimariesAttrTypes}},
	"enable_preferred_primaries": types.BoolType,
}

// ZoneRpGridSecondariesResourceSchemaAttributes contains the schema attributes for ZoneRpGridSecondariesModel
var ZoneRpGridSecondariesResourceSchemaAttributes = map[string]schema.Attribute{
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
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "This flag governs whether the specified Grid member is in stealth mode or not. If set to True, the member is in stealth mode. This flag is ignored if the struct is specified as part of a stub zone.",
	},
	"grid_replicate": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "The flag represents DNS zone transfers if set to False, and ID Grid Replication if set to True. This flag is ignored if the struct is specified as part of a stub zone or if it is set as grid_member in an authoritative zone.",
	},
	"lead": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "This flag controls whether the Grid lead secondary server performs zone transfers to non lead secondaries. This flag is ignored if the struct is specified as grid_member in an authoritative zone.",
	},
	"preferred_primaries": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: ZonerpgridsecondariesPreferredPrimariesResourceSchemaAttributes,
		},
		Optional: true,
		Validators: []validator.List{
			customvalidator.ListNotEmpty(),
		},
		MarkdownDescription: "The primary preference list with Grid member names and\\or External Server extserver structs for this member.",
	},
	"enable_preferred_primaries": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "This flag represents whether the preferred_primaries field values of this member are used.",
	},
}

// ExpandZoneRpGridSecondaries converts a Terraform Object to SDK type
func ExpandZoneRpGridSecondaries(ctx context.Context, o types.Object, diags *diag.Diagnostics) *niosdns.ZoneRpGridSecondaries {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m ZoneRpGridSecondariesModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

// Expand converts the Terraform model to SDK type
func (m *ZoneRpGridSecondariesModel) Expand(ctx context.Context, diags *diag.Diagnostics) *niosdns.ZoneRpGridSecondaries {
	if m == nil {
		return nil
	}
	to := &niosdns.ZoneRpGridSecondaries{
		Name:                     flex.ExpandStringPointerNullAsEmpty(m.Name),
		Stealth:                  flex.ExpandBoolPointer(m.Stealth),
		GridReplicate:            flex.ExpandBoolPointer(m.GridReplicate),
		Lead:                     flex.ExpandBoolPointer(m.Lead),
		PreferredPrimaries:       flex.ExpandFrameworkListNestedBlock(ctx, m.PreferredPrimaries, diags, ExpandZonerpgridsecondariesPreferredPrimaries),
		EnablePreferredPrimaries: flex.ExpandBoolPointer(m.EnablePreferredPrimaries),
	}
	return to
}

// FlattenZoneRpGridSecondaries converts an SDK type to Terraform Object
func FlattenZoneRpGridSecondaries(ctx context.Context, from *niosdns.ZoneRpGridSecondaries, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(ZoneRpGridSecondariesAttrTypes)
	}
	m := &ZoneRpGridSecondariesModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, ZoneRpGridSecondariesAttrTypes, m)
	diags.Append(d...)
	return t
}

// Flatten populates the Terraform model from SDK type
func (m *ZoneRpGridSecondariesModel) Flatten(ctx context.Context, from *niosdns.ZoneRpGridSecondaries, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	m.Name = flex.FlattenStringPointerEmptyAsNull(from.Name)
	m.Stealth = flex.FlattenBoolPointer(from.Stealth)
	m.GridReplicate = flex.FlattenBoolPointer(from.GridReplicate)
	m.Lead = flex.FlattenBoolPointer(from.Lead)
	m.PreferredPrimaries = flex.FlattenFrameworkListNestedBlock(ctx, from.PreferredPrimaries, ZonerpgridsecondariesPreferredPrimariesAttrTypes, diags, FlattenZonerpgridsecondariesPreferredPrimaries)
	m.EnablePreferredPrimaries = flex.FlattenBoolPointer(from.EnablePreferredPrimaries)
}
