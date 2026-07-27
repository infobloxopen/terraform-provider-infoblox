package dns

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	niosdns "github.com/infobloxopen/infoblox-nios-go-client/dns"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
	customvalidator "github.com/infobloxopen/terraform-provider-infoblox/internal/validator"
)

// ViewFixedRrsetOrderFqdnsModel is the Terraform model for ViewFixedRrsetOrderFqdns
type ViewFixedRrsetOrderFqdnsModel struct {
	Fqdn       types.String `tfsdk:"fqdn"`
	RecordType types.String `tfsdk:"record_type"`
}

// ViewFixedRrsetOrderFqdnsAttrTypes contains the attribute types for ViewFixedRrsetOrderFqdnsModel
var ViewFixedRrsetOrderFqdnsAttrTypes = map[string]attr.Type{
	"fqdn":        types.StringType,
	"record_type": types.StringType,
}

// ViewFixedRrsetOrderFqdnsResourceSchemaAttributes contains the schema attributes for ViewFixedRrsetOrderFqdnsModel
var ViewFixedRrsetOrderFqdnsResourceSchemaAttributes = map[string]schema.Attribute{
	"fqdn": schema.StringAttribute{
		Required: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
			customvalidator.IsValidDomainName(),
		},
		MarkdownDescription: "The FQDN of the fixed RRset configuration item.",
	},
	"record_type": schema.StringAttribute{
		Default: stringdefault.StaticString("A"),
		Validators: []validator.String{
			stringvalidator.OneOf("A", "AAAA", "BOTH"),
		},
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The record type for the specified FQDN in the fixed RRset configuration.",
	},
}

// ExpandViewFixedRrsetOrderFqdns converts a Terraform Object to SDK type
func ExpandViewFixedRrsetOrderFqdns(ctx context.Context, o types.Object, diags *diag.Diagnostics) *niosdns.ViewFixedRrsetOrderFqdns {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m ViewFixedRrsetOrderFqdnsModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

// Expand converts the Terraform model to SDK type
func (m *ViewFixedRrsetOrderFqdnsModel) Expand(ctx context.Context, diags *diag.Diagnostics) *niosdns.ViewFixedRrsetOrderFqdns {
	if m == nil {
		return nil
	}
	to := &niosdns.ViewFixedRrsetOrderFqdns{
		Fqdn:       flex.ExpandStringPointerNullAsEmpty(m.Fqdn),
		RecordType: flex.ExpandStringPointerNullAsEmpty(m.RecordType),
	}
	return to
}

// FlattenViewFixedRrsetOrderFqdns converts an SDK type to Terraform Object
func FlattenViewFixedRrsetOrderFqdns(ctx context.Context, from *niosdns.ViewFixedRrsetOrderFqdns, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(ViewFixedRrsetOrderFqdnsAttrTypes)
	}
	m := &ViewFixedRrsetOrderFqdnsModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, ViewFixedRrsetOrderFqdnsAttrTypes, m)
	diags.Append(d...)
	return t
}

// Flatten populates the Terraform model from SDK type
func (m *ViewFixedRrsetOrderFqdnsModel) Flatten(ctx context.Context, from *niosdns.ViewFixedRrsetOrderFqdns, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	m.Fqdn = flex.FlattenStringPointerEmptyAsNull(from.Fqdn)
	m.RecordType = flex.FlattenStringPointerEmptyAsNull(from.RecordType)
}
