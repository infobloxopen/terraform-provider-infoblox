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

// RecordcnamecloudinfoDelegatedMemberModel is the Terraform model for RecordcnamecloudinfoDelegatedMember
type RecordcnamecloudinfoDelegatedMemberModel struct {
	Ipv4addr types.String `tfsdk:"ipv4addr"`
	Ipv6addr types.String `tfsdk:"ipv6addr"`
	Name     types.String `tfsdk:"name"`
}

// RecordcnamecloudinfoDelegatedMemberAttrTypes contains the attribute types for RecordcnamecloudinfoDelegatedMemberModel
var RecordcnamecloudinfoDelegatedMemberAttrTypes = map[string]attr.Type{
	"ipv4addr": types.StringType,
	"ipv6addr": types.StringType,
	"name":     types.StringType,
}

// RecordcnamecloudinfoDelegatedMemberResourceSchemaAttributes contains the schema attributes for RecordcnamecloudinfoDelegatedMemberModel
var RecordcnamecloudinfoDelegatedMemberResourceSchemaAttributes = map[string]schema.Attribute{
	"ipv4addr": schema.StringAttribute{
		Optional: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
		},
		MarkdownDescription: "The IPv4 Address of the Grid Member.",
	},
	"ipv6addr": schema.StringAttribute{
		Optional: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
		},
		MarkdownDescription: "The IPv6 Address of the Grid Member.",
	},
	"name": schema.StringAttribute{
		Optional: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
		},
		MarkdownDescription: "The Grid member name",
	},
}

// ExpandRecordcnamecloudinfoDelegatedMember converts a Terraform Object to SDK type
func ExpandRecordcnamecloudinfoDelegatedMember(ctx context.Context, o types.Object, diags *diag.Diagnostics) *niosdns.RecordcnamecloudinfoDelegatedMember {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m RecordcnamecloudinfoDelegatedMemberModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

// Expand converts the Terraform model to SDK type
func (m *RecordcnamecloudinfoDelegatedMemberModel) Expand(ctx context.Context, diags *diag.Diagnostics) *niosdns.RecordcnamecloudinfoDelegatedMember {
	if m == nil {
		return nil
	}
	to := &niosdns.RecordcnamecloudinfoDelegatedMember{
		Ipv4addr: flex.ExpandStringPointerNullAsEmpty(m.Ipv4addr),
		Ipv6addr: flex.ExpandStringPointerNullAsEmpty(m.Ipv6addr),
		Name:     flex.ExpandStringPointerNullAsEmpty(m.Name),
	}
	return to
}

// FlattenRecordcnamecloudinfoDelegatedMember converts an SDK type to Terraform Object
func FlattenRecordcnamecloudinfoDelegatedMember(ctx context.Context, from *niosdns.RecordcnamecloudinfoDelegatedMember, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(RecordcnamecloudinfoDelegatedMemberAttrTypes)
	}
	m := &RecordcnamecloudinfoDelegatedMemberModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, RecordcnamecloudinfoDelegatedMemberAttrTypes, m)
	diags.Append(d...)
	return t
}

// Flatten populates the Terraform model from SDK type
func (m *RecordcnamecloudinfoDelegatedMemberModel) Flatten(ctx context.Context, from *niosdns.RecordcnamecloudinfoDelegatedMember, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	m.Ipv4addr = flex.FlattenStringPointerEmptyAsNull(from.Ipv4addr)
	m.Ipv6addr = flex.FlattenStringPointerEmptyAsNull(from.Ipv6addr)
	m.Name = flex.FlattenStringPointerEmptyAsNull(from.Name)
}
