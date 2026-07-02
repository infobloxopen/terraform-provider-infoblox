package dns

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	uddidns "github.com/infobloxopen/bloxone-go-client/dnsdata"
)

// RecordInheritanceModel is the Terraform model for RecordInheritance
type RecordInheritanceModel struct {
	Ttl types.Object `tfsdk:"ttl"`
}

// RecordInheritanceAttrTypes contains the attribute types for RecordInheritanceModel
var RecordInheritanceAttrTypes = map[string]attr.Type{
	"ttl": types.ObjectType{AttrTypes: Inheritance2InheritedUInt32AttrTypes},
}

// RecordInheritanceResourceSchemaAttributes contains the schema attributes for RecordInheritanceModel
var RecordInheritanceResourceSchemaAttributes = map[string]schema.Attribute{
	"ttl": schema.SingleNestedAttribute{
		Attributes:          Inheritance2InheritedUInt32ResourceSchemaAttributes,
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "",
	},
}

// ExpandRecordInheritance converts a Terraform Object to SDK type
func ExpandRecordInheritance(ctx context.Context, o types.Object, diags *diag.Diagnostics) *uddidns.RecordInheritance {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m RecordInheritanceModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

// Expand converts the Terraform model to SDK type
func (m *RecordInheritanceModel) Expand(ctx context.Context, diags *diag.Diagnostics) *uddidns.RecordInheritance {
	if m == nil {
		return nil
	}
	to := &uddidns.RecordInheritance{
		Ttl: ExpandInheritance2InheritedUInt32(ctx, m.Ttl, diags),
	}
	return to
}

// FlattenRecordInheritance converts an SDK type to Terraform Object
func FlattenRecordInheritance(ctx context.Context, from *uddidns.RecordInheritance, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(RecordInheritanceAttrTypes)
	}
	m := &RecordInheritanceModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, RecordInheritanceAttrTypes, m)
	diags.Append(d...)
	return t
}

// Flatten populates the Terraform model from SDK type
func (m *RecordInheritanceModel) Flatten(ctx context.Context, from *uddidns.RecordInheritance, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	m.Ttl = FlattenInheritance2InheritedUInt32(ctx, from.Ttl, diags)
}
