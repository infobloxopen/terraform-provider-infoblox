package ipam

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	niosipam "github.com/infobloxopen/infoblox-nios-go-client/ipam"
)

// Ipv6networkCloudInfoModel is the Terraform model for Ipv6networkCloudInfo
type Ipv6networkCloudInfoModel struct {
	DelegatedMember types.Object `tfsdk:"delegated_member"`
}

// Ipv6networkCloudInfoAttrTypes contains the attribute types for Ipv6networkCloudInfoModel
var Ipv6networkCloudInfoAttrTypes = map[string]attr.Type{
	"delegated_member": types.ObjectType{AttrTypes: Ipv6networkcloudinfoDelegatedMemberAttrTypes},
}

// Ipv6networkCloudInfoResourceSchemaAttributes contains the schema attributes for Ipv6networkCloudInfoModel
var Ipv6networkCloudInfoResourceSchemaAttributes = map[string]schema.Attribute{
	"delegated_member": schema.SingleNestedAttribute{
		Attributes:          Ipv6networkcloudinfoDelegatedMemberResourceSchemaAttributes,
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The Cloud Platform Appliance to which authority of the object is delegated.",
	},
}

// ExpandIpv6networkCloudInfo converts a Terraform Object to SDK type
func ExpandIpv6networkCloudInfo(ctx context.Context, o types.Object, diags *diag.Diagnostics) *niosipam.Ipv6networkCloudInfo {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m Ipv6networkCloudInfoModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

// Expand converts the Terraform model to SDK type
func (m *Ipv6networkCloudInfoModel) Expand(ctx context.Context, diags *diag.Diagnostics) *niosipam.Ipv6networkCloudInfo {
	if m == nil {
		return nil
	}
	to := &niosipam.Ipv6networkCloudInfo{
		DelegatedMember: ExpandIpv6networkcloudinfoDelegatedMember(ctx, m.DelegatedMember, diags),
	}
	return to
}

// FlattenIpv6networkCloudInfo converts an SDK type to Terraform Object
func FlattenIpv6networkCloudInfo(ctx context.Context, from *niosipam.Ipv6networkCloudInfo, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(Ipv6networkCloudInfoAttrTypes)
	}
	m := &Ipv6networkCloudInfoModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, Ipv6networkCloudInfoAttrTypes, m)
	diags.Append(d...)
	return t
}

// Flatten populates the Terraform model from SDK type
func (m *Ipv6networkCloudInfoModel) Flatten(ctx context.Context, from *niosipam.Ipv6networkCloudInfo, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	m.DelegatedMember = FlattenIpv6networkcloudinfoDelegatedMember(ctx, from.DelegatedMember, diags)
}
