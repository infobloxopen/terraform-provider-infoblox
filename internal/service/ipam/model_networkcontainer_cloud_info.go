package ipam

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
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

// NetworkcontainerCloudInfoModel is the Terraform model for NetworkcontainerCloudInfo
type NetworkcontainerCloudInfoModel struct {
	DelegatedMember types.Object `tfsdk:"delegated_member"`
	DelegatedScope  types.String `tfsdk:"delegated_scope"`
	DelegatedRoot   types.String `tfsdk:"delegated_root"`
	OwnedByAdaptor  types.Bool   `tfsdk:"owned_by_adaptor"`
	Usage           types.String `tfsdk:"usage"`
	Tenant          types.String `tfsdk:"tenant"`
	MgmtPlatform    types.String `tfsdk:"mgmt_platform"`
	AuthorityType   types.String `tfsdk:"authority_type"`
}

// NetworkcontainerCloudInfoAttrTypes contains the attribute types for NetworkcontainerCloudInfoModel
var NetworkcontainerCloudInfoAttrTypes = map[string]attr.Type{
	"delegated_member": types.ObjectType{AttrTypes: NetworkcontainercloudinfoDelegatedMemberAttrTypes},
	"delegated_scope":  types.StringType,
	"delegated_root":   types.StringType,
	"owned_by_adaptor": types.BoolType,
	"usage":            types.StringType,
	"tenant":           types.StringType,
	"mgmt_platform":    types.StringType,
	"authority_type":   types.StringType,
}

// NetworkcontainerCloudInfoResourceSchemaAttributes contains the schema attributes for NetworkcontainerCloudInfoModel
var NetworkcontainerCloudInfoResourceSchemaAttributes = map[string]schema.Attribute{
	"delegated_member": schema.SingleNestedAttribute{
		Attributes:          NetworkcontainercloudinfoDelegatedMemberResourceSchemaAttributes,
		Optional:            true,
		MarkdownDescription: "The Cloud Platform Appliance to which authority of the object is delegated.",
	},
	"delegated_scope": schema.StringAttribute{
		Validators: []validator.String{
			stringvalidator.OneOf("NONE", "ROOT", "SUBTREE", "RECLAIMING"),
		},
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "Indicates the scope of delegation for the object. This can be one of the following: NONE (outside any delegation), ROOT (the delegation point), SUBTREE (within the scope of a delegation), RECLAIMING (within the scope of a delegation being reclaimed, either as the delegation point or in the subtree).",
	},
	"delegated_root": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
		},
		MarkdownDescription: "Indicates the root of the delegation if delegated_scope is SUBTREE or RECLAIMING. This is not set otherwise.",
	},
	"owned_by_adaptor": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "Determines whether the object was created by the cloud adapter or not.",
	},
	"usage": schema.StringAttribute{
		Validators: []validator.String{
			stringvalidator.OneOf("NONE", "ADAPTER", "USED_BY", "DELEGATED"),
		},
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "Indicates the cloud origin of the object.",
	},
	"tenant": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
		},
		MarkdownDescription: "Reference to the tenant object associated with the object, if any.",
	},
	"mgmt_platform": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
		},
		MarkdownDescription: "Indicates the specified cloud management platform.",
	},
	"authority_type": schema.StringAttribute{
		Validators: []validator.String{
			stringvalidator.OneOf("NONE", "GM", "CP"),
		},
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "Type of authority over the object.",
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
		DelegatedScope:  flex.ExpandStringPointerNullAsEmpty(m.DelegatedScope),
		DelegatedRoot:   flex.ExpandStringPointerNullAsEmpty(m.DelegatedRoot),
		OwnedByAdaptor:  flex.ExpandBoolPointer(m.OwnedByAdaptor),
		Usage:           flex.ExpandStringPointerNullAsEmpty(m.Usage),
		Tenant:          flex.ExpandStringPointerNullAsEmpty(m.Tenant),
		MgmtPlatform:    flex.ExpandStringPointerNullAsEmpty(m.MgmtPlatform),
		AuthorityType:   flex.ExpandStringPointerNullAsEmpty(m.AuthorityType),
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
	m.DelegatedScope = flex.FlattenStringPointerEmptyAsNull(from.DelegatedScope)
	m.DelegatedRoot = flex.FlattenStringPointerEmptyAsNull(from.DelegatedRoot)
	m.OwnedByAdaptor = flex.FlattenBoolPointer(from.OwnedByAdaptor)
	m.Usage = flex.FlattenStringPointerEmptyAsNull(from.Usage)
	m.Tenant = flex.FlattenStringPointerEmptyAsNull(from.Tenant)
	m.MgmtPlatform = flex.FlattenStringPointerEmptyAsNull(from.MgmtPlatform)
	m.AuthorityType = flex.FlattenStringPointerEmptyAsNull(from.AuthorityType)
}
