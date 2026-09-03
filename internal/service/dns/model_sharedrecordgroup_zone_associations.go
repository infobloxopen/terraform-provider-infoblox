package dns

import (
	"context"

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

// SharedrecordgroupZoneAssociationsModel is the Terraform model for SharedrecordgroupZoneAssociations
type SharedrecordgroupZoneAssociationsModel struct {
	Fqdn types.String `tfsdk:"fqdn"`
	View types.String `tfsdk:"view"`
}

// SharedrecordgroupZoneAssociationsAttrTypes contains the attribute types for SharedrecordgroupZoneAssociationsModel
var SharedrecordgroupZoneAssociationsAttrTypes = map[string]attr.Type{
	"fqdn": types.StringType,
	"view": types.StringType,
}

// SharedrecordgroupZoneAssociationsResourceSchemaAttributes contains the schema attributes for SharedrecordgroupZoneAssociationsModel
var SharedrecordgroupZoneAssociationsResourceSchemaAttributes = map[string]schema.Attribute{
	"fqdn": schema.StringAttribute{
		Optional: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
			customvalidator.IsValidNIOSDomainName(),
			customvalidator.IsNotArpa(),
		},
		MarkdownDescription: "The FQDN of the authoritative forward zone.",
	},
	"view": schema.StringAttribute{
		Default:  stringdefault.StaticString("default"),
		Optional: true,
		Computed: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
		},
		MarkdownDescription: "The view to which the zone belongs. If a view is not specified, the default view is used.",
	},
}

// ExpandSharedrecordgroupZoneAssociations converts a Terraform Object to SDK type
func ExpandSharedrecordgroupZoneAssociations(ctx context.Context, o types.Object, diags *diag.Diagnostics) *niosdns.SharedrecordgroupZoneAssociations {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m SharedrecordgroupZoneAssociationsModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

// Expand converts the Terraform model to SDK type
func (m *SharedrecordgroupZoneAssociationsModel) Expand(ctx context.Context, diags *diag.Diagnostics) *niosdns.SharedrecordgroupZoneAssociations {
	if m == nil {
		return nil
	}
	to := &niosdns.SharedrecordgroupZoneAssociations{
		Fqdn: flex.ExpandStringPointerNullAsEmpty(m.Fqdn),
		View: flex.ExpandStringPointerNullAsEmpty(m.View),
	}
	return to
}

// FlattenSharedrecordgroupZoneAssociations converts an SDK type to Terraform Object
func FlattenSharedrecordgroupZoneAssociations(ctx context.Context, from *niosdns.SharedrecordgroupZoneAssociations, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(SharedrecordgroupZoneAssociationsAttrTypes)
	}
	m := &SharedrecordgroupZoneAssociationsModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, SharedrecordgroupZoneAssociationsAttrTypes, m)
	diags.Append(d...)
	return t
}

// Flatten populates the Terraform model from SDK type
func (m *SharedrecordgroupZoneAssociationsModel) Flatten(ctx context.Context, from *niosdns.SharedrecordgroupZoneAssociations, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	m.Fqdn = flex.FlattenStringPointerEmptyAsNull(from.Fqdn)
	m.View = flex.FlattenStringPointerEmptyAsNull(from.View)
}
