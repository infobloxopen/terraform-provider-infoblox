package ipam

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	niosipam "github.com/infobloxopen/infoblox-nios-go-client/ipam"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
	customvalidator "github.com/infobloxopen/terraform-provider-infoblox/internal/validator"
)

// Ipv6networkcloudinfoDelegatedMemberModel is the Terraform model for Ipv6networkcloudinfoDelegatedMember
type Ipv6networkcloudinfoDelegatedMemberModel struct {
	Ipv4addr types.String `tfsdk:"ipv4addr"`
	Ipv6addr types.String `tfsdk:"ipv6addr"`
	Name     types.String `tfsdk:"name"`
}

// Ipv6networkcloudinfoDelegatedMemberAttrTypes contains the attribute types for Ipv6networkcloudinfoDelegatedMemberModel
var Ipv6networkcloudinfoDelegatedMemberAttrTypes = map[string]attr.Type{
	"ipv4addr": types.StringType,
	"ipv6addr": types.StringType,
	"name":     types.StringType,
}

// Ipv6networkcloudinfoDelegatedMemberResourceSchemaAttributes contains the schema attributes for Ipv6networkcloudinfoDelegatedMemberModel
var Ipv6networkcloudinfoDelegatedMemberResourceSchemaAttributes = map[string]schema.Attribute{
	"ipv4addr": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
		},
		MarkdownDescription: "The IPv4 Address of the Grid Member.",
	},
	"ipv6addr": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
		},
		MarkdownDescription: "The IPv6 Address of the Grid Member.",
	},
	"name": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
		},
		MarkdownDescription: "The Grid member name",
	},
}

// ExpandIpv6networkcloudinfoDelegatedMember converts a Terraform Object to SDK type
func ExpandIpv6networkcloudinfoDelegatedMember(ctx context.Context, o types.Object, diags *diag.Diagnostics) *niosipam.Ipv6networkcloudinfoDelegatedMember {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m Ipv6networkcloudinfoDelegatedMemberModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

// Expand converts the Terraform model to SDK type
func (m *Ipv6networkcloudinfoDelegatedMemberModel) Expand(ctx context.Context, diags *diag.Diagnostics) *niosipam.Ipv6networkcloudinfoDelegatedMember {
	if m == nil {
		return nil
	}
	to := &niosipam.Ipv6networkcloudinfoDelegatedMember{
		Ipv4addr: flex.ExpandStringPointerNullAsEmpty(m.Ipv4addr),
		Ipv6addr: flex.ExpandStringPointerNullAsEmpty(m.Ipv6addr),
		Name:     flex.ExpandStringPointerNullAsEmpty(m.Name),
	}
	return to
}

// FlattenIpv6networkcloudinfoDelegatedMember converts an SDK type to Terraform Object
func FlattenIpv6networkcloudinfoDelegatedMember(ctx context.Context, from *niosipam.Ipv6networkcloudinfoDelegatedMember, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(Ipv6networkcloudinfoDelegatedMemberAttrTypes)
	}
	m := &Ipv6networkcloudinfoDelegatedMemberModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, Ipv6networkcloudinfoDelegatedMemberAttrTypes, m)
	diags.Append(d...)
	return t
}

// Flatten populates the Terraform model from SDK type
func (m *Ipv6networkcloudinfoDelegatedMemberModel) Flatten(ctx context.Context, from *niosipam.Ipv6networkcloudinfoDelegatedMember, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	m.Ipv4addr = flex.FlattenStringPointerEmptyAsNull(from.Ipv4addr)
	m.Ipv6addr = flex.FlattenStringPointerEmptyAsNull(from.Ipv6addr)
	m.Name = flex.FlattenStringPointerEmptyAsNull(from.Name)
}
