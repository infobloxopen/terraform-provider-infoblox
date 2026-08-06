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

// RecordsrvcloudinfoDelegatedMemberModel is the Terraform model for RecordsrvcloudinfoDelegatedMember
type RecordsrvcloudinfoDelegatedMemberModel struct {
	Ipv4addr types.String `tfsdk:"ipv4addr"`
	Ipv6addr types.String `tfsdk:"ipv6addr"`
	Name     types.String `tfsdk:"name"`
}

// RecordsrvcloudinfoDelegatedMemberAttrTypes contains the attribute types for RecordsrvcloudinfoDelegatedMemberModel
var RecordsrvcloudinfoDelegatedMemberAttrTypes = map[string]attr.Type{
	"ipv4addr": types.StringType,
	"ipv6addr": types.StringType,
	"name":     types.StringType,
}

// RecordsrvcloudinfoDelegatedMemberResourceSchemaAttributes contains the schema attributes for RecordsrvcloudinfoDelegatedMemberModel
var RecordsrvcloudinfoDelegatedMemberResourceSchemaAttributes = map[string]schema.Attribute{
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

// ExpandRecordsrvcloudinfoDelegatedMember converts a Terraform Object to SDK type
func ExpandRecordsrvcloudinfoDelegatedMember(ctx context.Context, o types.Object, diags *diag.Diagnostics) *niosdns.RecordsrvcloudinfoDelegatedMember {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m RecordsrvcloudinfoDelegatedMemberModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

// Expand converts the Terraform model to SDK type
func (m *RecordsrvcloudinfoDelegatedMemberModel) Expand(ctx context.Context, diags *diag.Diagnostics) *niosdns.RecordsrvcloudinfoDelegatedMember {
	if m == nil {
		return nil
	}
	to := &niosdns.RecordsrvcloudinfoDelegatedMember{
		Ipv4addr: flex.ExpandStringPointerNullAsEmpty(m.Ipv4addr),
		Ipv6addr: flex.ExpandStringPointerNullAsEmpty(m.Ipv6addr),
		Name:     flex.ExpandStringPointerNullAsEmpty(m.Name),
	}
	return to
}

// FlattenRecordsrvcloudinfoDelegatedMember converts an SDK type to Terraform Object
func FlattenRecordsrvcloudinfoDelegatedMember(ctx context.Context, from *niosdns.RecordsrvcloudinfoDelegatedMember, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(RecordsrvcloudinfoDelegatedMemberAttrTypes)
	}
	m := &RecordsrvcloudinfoDelegatedMemberModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, RecordsrvcloudinfoDelegatedMemberAttrTypes, m)
	diags.Append(d...)
	return t
}

// Flatten populates the Terraform model from SDK type
func (m *RecordsrvcloudinfoDelegatedMemberModel) Flatten(ctx context.Context, from *niosdns.RecordsrvcloudinfoDelegatedMember, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	m.Ipv4addr = flex.FlattenStringPointerEmptyAsNull(from.Ipv4addr)
	m.Ipv6addr = flex.FlattenStringPointerEmptyAsNull(from.Ipv6addr)
	m.Name = flex.FlattenStringPointerEmptyAsNull(from.Name)
}
