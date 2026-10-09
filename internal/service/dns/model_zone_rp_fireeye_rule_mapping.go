package dns

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
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

// ZoneRpFireeyeRuleMappingModel is the Terraform model for ZoneRpFireeyeRuleMapping
type ZoneRpFireeyeRuleMappingModel struct {
	AptOverride           types.String `tfsdk:"apt_override"`
	FireeyeAlertMapping   types.List   `tfsdk:"fireeye_alert_mapping"`
	SubstitutedDomainName types.String `tfsdk:"substituted_domain_name"`
}

// ZoneRpFireeyeRuleMappingAttrTypes contains the attribute types for ZoneRpFireeyeRuleMappingModel
var ZoneRpFireeyeRuleMappingAttrTypes = map[string]attr.Type{
	"apt_override":            types.StringType,
	"fireeye_alert_mapping":   types.ListType{ElemType: types.ObjectType{AttrTypes: ZonerpfireeyerulemappingFireeyeAlertMappingAttrTypes}},
	"substituted_domain_name": types.StringType,
}

// ZoneRpFireeyeRuleMappingResourceSchemaAttributes contains the schema attributes for ZoneRpFireeyeRuleMappingModel
var ZoneRpFireeyeRuleMappingResourceSchemaAttributes = map[string]schema.Attribute{
	"apt_override": schema.StringAttribute{
		Validators: []validator.String{
			stringvalidator.OneOf("PASSTHRU", "NXDOMAIN", "NODATA", "SUBSTITUTE", "NOOVERRIDE"),
		},
		Optional:            true,
		MarkdownDescription: "The override setting for APT alerts.",
	},
	"fireeye_alert_mapping": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: ZonerpfireeyerulemappingFireeyeAlertMappingResourceSchemaAttributes,
		},
		Optional: true,
		Computed: true,
		Validators: []validator.List{
			customvalidator.ListNotEmpty(),
		},
		MarkdownDescription: "The FireEye alert mapping.",
	},
	"substituted_domain_name": schema.StringAttribute{
		Optional: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
		},
		MarkdownDescription: "The domain name to be substituted, this is applicable only when apt_override is set to \"SUBSTITUTE\".",
	},
}

// ExpandZoneRpFireeyeRuleMapping converts a Terraform Object to SDK type
func ExpandZoneRpFireeyeRuleMapping(ctx context.Context, o types.Object, diags *diag.Diagnostics) *niosdns.ZoneRpFireeyeRuleMapping {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m ZoneRpFireeyeRuleMappingModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

// Expand converts the Terraform model to SDK type
func (m *ZoneRpFireeyeRuleMappingModel) Expand(ctx context.Context, diags *diag.Diagnostics) *niosdns.ZoneRpFireeyeRuleMapping {
	if m == nil {
		return nil
	}
	to := &niosdns.ZoneRpFireeyeRuleMapping{
		AptOverride:           flex.ExpandStringPointerNullAsEmpty(m.AptOverride),
		FireeyeAlertMapping:   flex.ExpandFrameworkListNestedBlock(ctx, m.FireeyeAlertMapping, diags, ExpandZonerpfireeyerulemappingFireeyeAlertMapping),
		SubstitutedDomainName: flex.ExpandStringPointerNullAsEmpty(m.SubstitutedDomainName),
	}
	return to
}

// FlattenZoneRpFireeyeRuleMapping converts an SDK type to Terraform Object
func FlattenZoneRpFireeyeRuleMapping(ctx context.Context, from *niosdns.ZoneRpFireeyeRuleMapping, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(ZoneRpFireeyeRuleMappingAttrTypes)
	}
	m := &ZoneRpFireeyeRuleMappingModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, ZoneRpFireeyeRuleMappingAttrTypes, m)
	diags.Append(d...)
	return t
}

// Flatten populates the Terraform model from SDK type
func (m *ZoneRpFireeyeRuleMappingModel) Flatten(ctx context.Context, from *niosdns.ZoneRpFireeyeRuleMapping, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	m.AptOverride = flex.FlattenStringPointerEmptyAsNull(from.AptOverride)
	m.FireeyeAlertMapping = flex.FlattenFrameworkListNestedBlock(ctx, from.FireeyeAlertMapping, ZonerpfireeyerulemappingFireeyeAlertMappingAttrTypes, diags, FlattenZonerpfireeyerulemappingFireeyeAlertMapping)
	m.SubstitutedDomainName = flex.FlattenStringPointerEmptyAsNull(from.SubstitutedDomainName)
}
