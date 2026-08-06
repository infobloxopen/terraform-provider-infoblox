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

// NetworkcontainerCloudInfoModel is the Terraform model for NetworkcontainerCloudInfo
type NetworkcontainerCloudInfoModel struct {
	DelegatedMember types.Object `tfsdk:"delegated_member"`
}

// NetworkcontainerCloudInfoAttrTypes contains the attribute types for NetworkcontainerCloudInfoModel
var NetworkcontainerCloudInfoAttrTypes = map[string]attr.Type{
	"delegated_member": types.ObjectType{AttrTypes: NetworkcontainercloudinfoDelegatedMemberAttrTypes},
}

// NetworkcontainerCloudInfoResourceSchemaAttributes contains the schema attributes for NetworkcontainerCloudInfoModel
var NetworkcontainerCloudInfoResourceSchemaAttributes = map[string]schema.Attribute{
	"delegated_member": schema.SingleNestedAttribute{
		Attributes:          NetworkcontainercloudinfoDelegatedMemberResourceSchemaAttributes,
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The Cloud Platform Appliance to which authority of the object is delegated.",
	},
}

// ExpandNetworkcontainerCloudInfo converts a Terraform Object to SDK type
func ExpandNetworkcontainerCloudInfo(ctx context.Context, o types.Object, diags *diag.Diagnostics) *niosipam.NetworkcontainerCloudInfo {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m NetworkcontainerCloudInfoModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

// Expand converts the Terraform model to SDK type
func (m *NetworkcontainerCloudInfoModel) Expand(ctx context.Context, diags *diag.Diagnostics) *niosipam.NetworkcontainerCloudInfo {
	if m == nil {
		return nil
	}
	to := &niosipam.NetworkcontainerCloudInfo{
		DelegatedMember: ExpandNetworkcontainercloudinfoDelegatedMember(ctx, m.DelegatedMember, diags),
	}
	return to
}

// FlattenNetworkcontainerCloudInfo converts an SDK type to Terraform Object
func FlattenNetworkcontainerCloudInfo(ctx context.Context, from *niosipam.NetworkcontainerCloudInfo, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(NetworkcontainerCloudInfoAttrTypes)
	}
	m := &NetworkcontainerCloudInfoModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, NetworkcontainerCloudInfoAttrTypes, m)
	diags.Append(d...)
	return t
}

// Flatten populates the Terraform model from SDK type
func (m *NetworkcontainerCloudInfoModel) Flatten(ctx context.Context, from *niosipam.NetworkcontainerCloudInfo, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	m.DelegatedMember = FlattenNetworkcontainercloudinfoDelegatedMember(ctx, from.DelegatedMember, diags)
}
