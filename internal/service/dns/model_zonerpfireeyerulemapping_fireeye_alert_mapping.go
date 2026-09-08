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
)

// ZonerpfireeyerulemappingFireeyeAlertMappingModel is the Terraform model for ZonerpfireeyerulemappingFireeyeAlertMapping
type ZonerpfireeyerulemappingFireeyeAlertMappingModel struct {
	AlertType types.String `tfsdk:"alert_type"`
	RpzRule   types.String `tfsdk:"rpz_rule"`
	Lifetime  types.Int64  `tfsdk:"lifetime"`
}

// ZonerpfireeyerulemappingFireeyeAlertMappingAttrTypes contains the attribute types for ZonerpfireeyerulemappingFireeyeAlertMappingModel
var ZonerpfireeyerulemappingFireeyeAlertMappingAttrTypes = map[string]attr.Type{
	"alert_type": types.StringType,
	"rpz_rule":   types.StringType,
	"lifetime":   types.Int64Type,
}

// ZonerpfireeyerulemappingFireeyeAlertMappingResourceSchemaAttributes contains the schema attributes for ZonerpfireeyerulemappingFireeyeAlertMappingModel
var ZonerpfireeyerulemappingFireeyeAlertMappingResourceSchemaAttributes = map[string]schema.Attribute{
	"alert_type": schema.StringAttribute{
		Validators: []validator.String{
			stringvalidator.OneOf("INFECTION_MATCH", "WEB_INFECTION", "MALWARE_OBJECT", "DOMAIN_MATCH", "MALWARE_CALLBACK"),
		},
		Optional:            true,
		MarkdownDescription: "The type of Fireeye Alert.",
	},
	"rpz_rule": schema.StringAttribute{
		Validators: []validator.String{
			stringvalidator.OneOf("PASSTHRU", "NXDOMAIN", "NODATA", "SUBSTITUTE", "NONE"),
		},
		Optional:            true,
		MarkdownDescription: "The RPZ rule for the alert.",
	},
	"lifetime": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "The expiration Lifetime of alert type. The 32-bit unsigned integer represents the amount of seconds this alert type will live for. 0 means the alert will never expire.",
	},
}

// ExpandZonerpfireeyerulemappingFireeyeAlertMapping converts a Terraform Object to SDK type
func ExpandZonerpfireeyerulemappingFireeyeAlertMapping(ctx context.Context, o types.Object, diags *diag.Diagnostics) *niosdns.ZonerpfireeyerulemappingFireeyeAlertMapping {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m ZonerpfireeyerulemappingFireeyeAlertMappingModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

// Expand converts the Terraform model to SDK type
func (m *ZonerpfireeyerulemappingFireeyeAlertMappingModel) Expand(ctx context.Context, diags *diag.Diagnostics) *niosdns.ZonerpfireeyerulemappingFireeyeAlertMapping {
	if m == nil {
		return nil
	}
	to := &niosdns.ZonerpfireeyerulemappingFireeyeAlertMapping{
		AlertType: flex.ExpandStringPointerNullAsEmpty(m.AlertType),
		RpzRule:   flex.ExpandStringPointerNullAsEmpty(m.RpzRule),
		Lifetime:  flex.ExpandInt64Pointer(m.Lifetime),
	}
	return to
}

// FlattenZonerpfireeyerulemappingFireeyeAlertMapping converts an SDK type to Terraform Object
func FlattenZonerpfireeyerulemappingFireeyeAlertMapping(ctx context.Context, from *niosdns.ZonerpfireeyerulemappingFireeyeAlertMapping, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(ZonerpfireeyerulemappingFireeyeAlertMappingAttrTypes)
	}
	m := &ZonerpfireeyerulemappingFireeyeAlertMappingModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, ZonerpfireeyerulemappingFireeyeAlertMappingAttrTypes, m)
	diags.Append(d...)
	return t
}

// Flatten populates the Terraform model from SDK type
func (m *ZonerpfireeyerulemappingFireeyeAlertMappingModel) Flatten(ctx context.Context, from *niosdns.ZonerpfireeyerulemappingFireeyeAlertMapping, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	m.AlertType = flex.FlattenStringPointerEmptyAsNull(from.AlertType)
	m.RpzRule = flex.FlattenStringPointerEmptyAsNull(from.RpzRule)
	m.Lifetime = flex.FlattenInt64Pointer(from.Lifetime)
}
