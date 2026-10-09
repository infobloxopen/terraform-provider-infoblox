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

// Ipv6networkcontainerCloudInfoModel is the Terraform model for Ipv6networkcontainerCloudInfo
type Ipv6networkcontainerCloudInfoModel struct {
	DelegatedMember types.Object `tfsdk:"delegated_member"`
}

// Ipv6networkcontainerCloudInfoAttrTypes contains the attribute types for Ipv6networkcontainerCloudInfoModel
var Ipv6networkcontainerCloudInfoAttrTypes = map[string]attr.Type{
	"delegated_member": types.ObjectType{AttrTypes: Ipv6networkcontainercloudinfoDelegatedMemberAttrTypes},
}

// Ipv6networkcontainerCloudInfoResourceSchemaAttributes contains the schema attributes for Ipv6networkcontainerCloudInfoModel
var Ipv6networkcontainerCloudInfoResourceSchemaAttributes = map[string]schema.Attribute{
	"delegated_member": schema.SingleNestedAttribute{
		Attributes:          Ipv6networkcontainercloudinfoDelegatedMemberResourceSchemaAttributes,
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The Cloud Platform Appliance to which authority of the object is delegated.",
	},
}

// ExpandIpv6networkcontainerCloudInfo converts a Terraform Object to SDK type
func ExpandIpv6networkcontainerCloudInfo(ctx context.Context, o types.Object, diags *diag.Diagnostics) *niosipam.Ipv6networkcontainerCloudInfo {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m Ipv6networkcontainerCloudInfoModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

// Expand converts the Terraform model to SDK type
func (m *Ipv6networkcontainerCloudInfoModel) Expand(ctx context.Context, diags *diag.Diagnostics) *niosipam.Ipv6networkcontainerCloudInfo {
	if m == nil {
		return nil
	}
	to := &niosipam.Ipv6networkcontainerCloudInfo{
		DelegatedMember: ExpandIpv6networkcontainercloudinfoDelegatedMember(ctx, m.DelegatedMember, diags),
	}
	return to
}

// FlattenIpv6networkcontainerCloudInfo converts an SDK type to Terraform Object
func FlattenIpv6networkcontainerCloudInfo(ctx context.Context, from *niosipam.Ipv6networkcontainerCloudInfo, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(Ipv6networkcontainerCloudInfoAttrTypes)
	}
	m := &Ipv6networkcontainerCloudInfoModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, Ipv6networkcontainerCloudInfoAttrTypes, m)
	diags.Append(d...)
	return t
}

// Flatten populates the Terraform model from SDK type
func (m *Ipv6networkcontainerCloudInfoModel) Flatten(ctx context.Context, from *niosipam.Ipv6networkcontainerCloudInfo, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	m.DelegatedMember = FlattenIpv6networkcontainercloudinfoDelegatedMember(ctx, from.DelegatedMember, diags)
}
