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

// NetworkCloudInfoModel is the Terraform model for NetworkCloudInfo
type NetworkCloudInfoModel struct {
	DelegatedMember types.Object `tfsdk:"delegated_member"`
}

// NetworkCloudInfoAttrTypes contains the attribute types for NetworkCloudInfoModel
var NetworkCloudInfoAttrTypes = map[string]attr.Type{
	"delegated_member": types.ObjectType{AttrTypes: NetworkcloudinfoDelegatedMemberAttrTypes},
}

// NetworkCloudInfoResourceSchemaAttributes contains the schema attributes for NetworkCloudInfoModel
var NetworkCloudInfoResourceSchemaAttributes = map[string]schema.Attribute{
	"delegated_member": schema.SingleNestedAttribute{
		Attributes:          NetworkcloudinfoDelegatedMemberResourceSchemaAttributes,
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The Cloud Platform Appliance to which authority of the object is delegated.",
	},
}

// ExpandNetworkCloudInfo converts a Terraform Object to SDK type
func ExpandNetworkCloudInfo(ctx context.Context, o types.Object, diags *diag.Diagnostics) *niosipam.NetworkCloudInfo {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m NetworkCloudInfoModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

// Expand converts the Terraform model to SDK type
func (m *NetworkCloudInfoModel) Expand(ctx context.Context, diags *diag.Diagnostics) *niosipam.NetworkCloudInfo {
	if m == nil {
		return nil
	}
	to := &niosipam.NetworkCloudInfo{
		DelegatedMember: ExpandNetworkcloudinfoDelegatedMember(ctx, m.DelegatedMember, diags),
	}
	return to
}

// FlattenNetworkCloudInfo converts an SDK type to Terraform Object
func FlattenNetworkCloudInfo(ctx context.Context, from *niosipam.NetworkCloudInfo, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(NetworkCloudInfoAttrTypes)
	}
	m := &NetworkCloudInfoModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, NetworkCloudInfoAttrTypes, m)
	diags.Append(d...)
	return t
}

// Flatten populates the Terraform model from SDK type
func (m *NetworkCloudInfoModel) Flatten(ctx context.Context, from *niosipam.NetworkCloudInfo, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	m.DelegatedMember = FlattenNetworkcloudinfoDelegatedMember(ctx, from.DelegatedMember, diags)
}
