package dns

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	niosdns "github.com/infobloxopen/infoblox-nios-go-client/dns"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
)

// ViewScavengingSettingsModel is the Terraform model for ViewScavengingSettings
type ViewScavengingSettingsModel struct {
	EnableScavenging          types.Bool `tfsdk:"enable_scavenging"`
	EnableRecurrentScavenging types.Bool `tfsdk:"enable_recurrent_scavenging"`
	EnableAutoReclamation     types.Bool `tfsdk:"enable_auto_reclamation"`
	EnableRrLastQueried       types.Bool `tfsdk:"enable_rr_last_queried"`
	EnableZoneLastQueried     types.Bool `tfsdk:"enable_zone_last_queried"`
	ReclaimAssociatedRecords  types.Bool `tfsdk:"reclaim_associated_records"`
	ExpressionList            types.List `tfsdk:"expression_list"`
	EaExpressionList          types.List `tfsdk:"ea_expression_list"`
}

// ViewScavengingSettingsAttrTypes contains the attribute types for ViewScavengingSettingsModel
var ViewScavengingSettingsAttrTypes = map[string]attr.Type{
	"enable_scavenging":           types.BoolType,
	"enable_recurrent_scavenging": types.BoolType,
	"enable_auto_reclamation":     types.BoolType,
	"enable_rr_last_queried":      types.BoolType,
	"enable_zone_last_queried":    types.BoolType,
	"reclaim_associated_records":  types.BoolType,
	"expression_list":             types.ListType{ElemType: types.ObjectType{AttrTypes: ViewscavengingsettingsExpressionListAttrTypes}},
	"ea_expression_list":          types.ListType{ElemType: types.ObjectType{AttrTypes: ViewscavengingsettingsEaExpressionListAttrTypes}},
}

// ViewScavengingSettingsResourceSchemaAttributes contains the schema attributes for ViewScavengingSettingsModel
var ViewScavengingSettingsResourceSchemaAttributes = map[string]schema.Attribute{
	"enable_scavenging": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "This flag indicates if the resource record scavenging is enabled or not.",
	},
	"enable_recurrent_scavenging": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "This flag indicates if the recurrent resource record scavenging is enabled or not.",
	},
	"enable_auto_reclamation": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "This flag indicates if the automatic resource record scavenging is enabled or not.",
	},
	"enable_rr_last_queried": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "This flag indicates if the resource record last queried monitoring in affected zones is enabled or not.",
	},
	"enable_zone_last_queried": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "This flag indicates if the last queried monitoring for affected zones is enabled or not.",
	},
	"reclaim_associated_records": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "This flag indicates if the associated resource record scavenging is enabled or not.",
	},
	"expression_list": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: ViewscavengingsettingsExpressionListResourceSchemaAttributes,
		},
		Optional: true,
		Computed: true,
		Validators: []validator.List{
			listvalidator.SizeAtLeast(1),
		},
		MarkdownDescription: "The expression list. The particular record is treated as reclaimable if expression condition evaluates to 'true' for given record if scavenging hasn't been manually disabled on a given resource record.",
	},
	"ea_expression_list": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: ViewscavengingsettingsEaExpressionListResourceSchemaAttributes,
		},
		Optional: true,
		Computed: true,
		Validators: []validator.List{
			listvalidator.SizeAtLeast(1),
		},
		MarkdownDescription: "The extensible attributes expression list. The particular record is treated as reclaimable if extensible attributes expression condition evaluates to 'true' for given record if scavenging hasn't been manually disabled on a given resource record.",
	},
}

// ExpandViewScavengingSettings converts a Terraform Object to SDK type
func ExpandViewScavengingSettings(ctx context.Context, o types.Object, diags *diag.Diagnostics) *niosdns.ViewScavengingSettings {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m ViewScavengingSettingsModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

// Expand converts the Terraform model to SDK type
func (m *ViewScavengingSettingsModel) Expand(ctx context.Context, diags *diag.Diagnostics) *niosdns.ViewScavengingSettings {
	if m == nil {
		return nil
	}
	to := &niosdns.ViewScavengingSettings{
		EnableScavenging:          flex.ExpandBoolPointer(m.EnableScavenging),
		EnableRecurrentScavenging: flex.ExpandBoolPointer(m.EnableRecurrentScavenging),
		EnableAutoReclamation:     flex.ExpandBoolPointer(m.EnableAutoReclamation),
		EnableRrLastQueried:       flex.ExpandBoolPointer(m.EnableRrLastQueried),
		EnableZoneLastQueried:     flex.ExpandBoolPointer(m.EnableZoneLastQueried),
		ReclaimAssociatedRecords:  flex.ExpandBoolPointer(m.ReclaimAssociatedRecords),
		ExpressionList:            flex.ExpandFrameworkListNestedBlock(ctx, m.ExpressionList, diags, ExpandViewscavengingsettingsExpressionList),
		EaExpressionList:          flex.ExpandFrameworkListNestedBlock(ctx, m.EaExpressionList, diags, ExpandViewscavengingsettingsEaExpressionList),
	}
	return to
}

// FlattenViewScavengingSettings converts an SDK type to Terraform Object
func FlattenViewScavengingSettings(ctx context.Context, from *niosdns.ViewScavengingSettings, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(ViewScavengingSettingsAttrTypes)
	}
	m := &ViewScavengingSettingsModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, ViewScavengingSettingsAttrTypes, m)
	diags.Append(d...)
	return t
}

// Flatten populates the Terraform model from SDK type
func (m *ViewScavengingSettingsModel) Flatten(ctx context.Context, from *niosdns.ViewScavengingSettings, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	m.EnableScavenging = flex.FlattenBoolPointer(from.EnableScavenging)
	m.EnableRecurrentScavenging = flex.FlattenBoolPointer(from.EnableRecurrentScavenging)
	m.EnableAutoReclamation = flex.FlattenBoolPointer(from.EnableAutoReclamation)
	m.EnableRrLastQueried = flex.FlattenBoolPointer(from.EnableRrLastQueried)
	m.EnableZoneLastQueried = flex.FlattenBoolPointer(from.EnableZoneLastQueried)
	m.ReclaimAssociatedRecords = flex.FlattenBoolPointer(from.ReclaimAssociatedRecords)
	m.ExpressionList = flex.FlattenFrameworkListNestedBlock(ctx, from.ExpressionList, ViewscavengingsettingsExpressionListAttrTypes, diags, FlattenViewscavengingsettingsExpressionList)
	m.EaExpressionList = flex.FlattenFrameworkListNestedBlock(ctx, from.EaExpressionList, ViewscavengingsettingsEaExpressionListAttrTypes, diags, FlattenViewscavengingsettingsEaExpressionList)
}
