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

// ZoneAuthDnssecKeysModel is the Terraform model for ZoneAuthDnssecKeys
type ZoneAuthDnssecKeysModel struct {
	Tag           types.Int64  `tfsdk:"tag"`
	Status        types.String `tfsdk:"status"`
	NextEventDate types.Int64  `tfsdk:"next_event_date"`
	Type          types.String `tfsdk:"type"`
	Algorithm     types.String `tfsdk:"algorithm"`
	PublicKey     types.String `tfsdk:"public_key"`
}

// ZoneAuthDnssecKeysAttrTypes contains the attribute types for ZoneAuthDnssecKeysModel
var ZoneAuthDnssecKeysAttrTypes = map[string]attr.Type{
	"tag":             types.Int64Type,
	"status":          types.StringType,
	"next_event_date": types.Int64Type,
	"type":            types.StringType,
	"algorithm":       types.StringType,
	"public_key":      types.StringType,
}

// ZoneAuthDnssecKeysResourceSchemaAttributes contains the schema attributes for ZoneAuthDnssecKeysModel
var ZoneAuthDnssecKeysResourceSchemaAttributes = map[string]schema.Attribute{
	"tag": schema.Int64Attribute{
		Required:            true,
		MarkdownDescription: "The tag of the key for the zone.",
	},
	"status": schema.StringAttribute{
		Validators: []validator.String{
			stringvalidator.OneOf("ACTIVE", "PUBLISHED", "ROLLED", "IMPORTED"),
		},
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The status of the key for the zone.",
	},
	"next_event_date": schema.Int64Attribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The next event date for the key, the rollover date for an active key or the removal date for an already rolled one.",
	},
	"type": schema.StringAttribute{
		Validators: []validator.String{
			stringvalidator.OneOf("KSK", "ZSK"),
		},
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The key type.",
	},
	"algorithm": schema.StringAttribute{
		Validators: []validator.String{
			stringvalidator.OneOf("10", "5", "7", "8", "13", "14"),
		},
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The public-key encryption algorithm. Values 1, 3 and 6 are deprecated from NIOS 9.0.",
	},
	"public_key": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
		},
		MarkdownDescription: "The Base-64 encoding of the public key.",
	},
}

// ExpandZoneAuthDnssecKeys converts a Terraform Object to SDK type
func ExpandZoneAuthDnssecKeys(ctx context.Context, o types.Object, diags *diag.Diagnostics) *niosdns.ZoneAuthDnssecKeys {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m ZoneAuthDnssecKeysModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

// Expand converts the Terraform model to SDK type
func (m *ZoneAuthDnssecKeysModel) Expand(ctx context.Context, diags *diag.Diagnostics) *niosdns.ZoneAuthDnssecKeys {
	if m == nil {
		return nil
	}
	to := &niosdns.ZoneAuthDnssecKeys{
		Tag:           flex.ExpandInt64Pointer(m.Tag),
		Status:        flex.ExpandStringPointer(m.Status),
		NextEventDate: flex.ExpandInt64Pointer(m.NextEventDate),
		Type:          flex.ExpandStringPointer(m.Type),
		Algorithm:     flex.ExpandStringPointer(m.Algorithm),
		PublicKey:     flex.ExpandStringPointerNullAsEmpty(m.PublicKey),
	}
	return to
}

// FlattenZoneAuthDnssecKeys converts an SDK type to Terraform Object
func FlattenZoneAuthDnssecKeys(ctx context.Context, from *niosdns.ZoneAuthDnssecKeys, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(ZoneAuthDnssecKeysAttrTypes)
	}
	m := &ZoneAuthDnssecKeysModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, ZoneAuthDnssecKeysAttrTypes, m)
	diags.Append(d...)
	return t
}

// Flatten populates the Terraform model from SDK type
func (m *ZoneAuthDnssecKeysModel) Flatten(ctx context.Context, from *niosdns.ZoneAuthDnssecKeys, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	m.Tag = flex.FlattenInt64Pointer(from.Tag)
	m.Status = flex.FlattenStringPointerEmptyAsNull(from.Status)
	m.NextEventDate = flex.FlattenInt64Pointer(from.NextEventDate)
	m.Type = flex.FlattenStringPointerEmptyAsNull(from.Type)
	m.Algorithm = flex.FlattenStringPointerEmptyAsNull(from.Algorithm)
	m.PublicKey = flex.FlattenStringPointerEmptyAsNull(from.PublicKey)
}
