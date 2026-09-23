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

// NetworkviewCloudInfoModel is the Terraform model for NetworkviewCloudInfo
type NetworkviewCloudInfoModel struct {
	DelegatedMember types.Object `tfsdk:"delegated_member"`
}

// NetworkviewCloudInfoAttrTypes contains the attribute types for NetworkviewCloudInfoModel
var NetworkviewCloudInfoAttrTypes = map[string]attr.Type{
	"delegated_member": types.ObjectType{AttrTypes: NetworkviewcloudinfoDelegatedMemberAttrTypes},
}

// NetworkviewCloudInfoResourceSchemaAttributes contains the schema attributes for NetworkviewCloudInfoModel
var NetworkviewCloudInfoResourceSchemaAttributes = map[string]schema.Attribute{
	"delegated_member": schema.SingleNestedAttribute{
		Attributes:          NetworkviewcloudinfoDelegatedMemberResourceSchemaAttributes,
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The Cloud Platform Appliance to which authority of the object is delegated.",
	},
}

// ExpandNetworkviewCloudInfo converts a Terraform Object to SDK type
func ExpandNetworkviewCloudInfo(ctx context.Context, o types.Object, diags *diag.Diagnostics) *niosipam.NetworkviewCloudInfo {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m NetworkviewCloudInfoModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

// Expand converts the Terraform model to SDK type
func (m *NetworkviewCloudInfoModel) Expand(ctx context.Context, diags *diag.Diagnostics) *niosipam.NetworkviewCloudInfo {
	if m == nil {
		return nil
	}
	to := &niosipam.NetworkviewCloudInfo{
		DelegatedMember: ExpandNetworkviewcloudinfoDelegatedMember(ctx, m.DelegatedMember, diags),
	}
	return to
}

// FlattenNetworkviewCloudInfo converts an SDK type to Terraform Object
func FlattenNetworkviewCloudInfo(ctx context.Context, from *niosipam.NetworkviewCloudInfo, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(NetworkviewCloudInfoAttrTypes)
	}
	m := &NetworkviewCloudInfoModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, NetworkviewCloudInfoAttrTypes, m)
	diags.Append(d...)
	return t
}

// Flatten populates the Terraform model from SDK type
func (m *NetworkviewCloudInfoModel) Flatten(ctx context.Context, from *niosipam.NetworkviewCloudInfo, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	m.DelegatedMember = FlattenNetworkviewcloudinfoDelegatedMember(ctx, from.DelegatedMember, diags)
}
