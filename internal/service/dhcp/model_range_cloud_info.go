package dhcp

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	niosdhcp "github.com/infobloxopen/infoblox-nios-go-client/dhcp"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
	customvalidator "github.com/infobloxopen/terraform-provider-infoblox/internal/validator"
)

// RangeCloudInfoModel is the Terraform model for RangeCloudInfo
type RangeCloudInfoModel struct {
	DelegatedMember types.Object `tfsdk:"delegated_member"`
	DelegatedScope  types.String `tfsdk:"delegated_scope"`
	DelegatedRoot   types.String `tfsdk:"delegated_root"`
	OwnedByAdaptor  types.Bool   `tfsdk:"owned_by_adaptor"`
	Usage           types.String `tfsdk:"usage"`
	Tenant          types.String `tfsdk:"tenant"`
	MgmtPlatform    types.String `tfsdk:"mgmt_platform"`
	AuthorityType   types.String `tfsdk:"authority_type"`
}

// RangeCloudInfoAttrTypes contains the attribute types for RangeCloudInfoModel
var RangeCloudInfoAttrTypes = map[string]attr.Type{
	"delegated_member": types.ObjectType{AttrTypes: RangecloudinfoDelegatedMemberAttrTypes},
	"delegated_scope":  types.StringType,
	"delegated_root":   types.StringType,
	"owned_by_adaptor": types.BoolType,
	"usage":            types.StringType,
	"tenant":           types.StringType,
	"mgmt_platform":    types.StringType,
	"authority_type":   types.StringType,
}

// RangeCloudInfoResourceSchemaAttributes contains the schema attributes for RangeCloudInfoModel
var RangeCloudInfoResourceSchemaAttributes = map[string]schema.Attribute{
	"delegated_member": schema.SingleNestedAttribute{
		Attributes:          RangecloudinfoDelegatedMemberResourceSchemaAttributes,
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

// ExpandRangeCloudInfo converts a Terraform Object to SDK type
func ExpandRangeCloudInfo(ctx context.Context, o types.Object, diags *diag.Diagnostics) *niosdhcp.RangeCloudInfo {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m RangeCloudInfoModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

// Expand converts the Terraform model to SDK type
func (m *RangeCloudInfoModel) Expand(ctx context.Context, diags *diag.Diagnostics) *niosdhcp.RangeCloudInfo {
	if m == nil {
		return nil
	}
	to := &niosdhcp.RangeCloudInfo{
		DelegatedMember: ExpandRangecloudinfoDelegatedMember(ctx, m.DelegatedMember, diags),
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

// FlattenRangeCloudInfo converts an SDK type to Terraform Object
func FlattenRangeCloudInfo(ctx context.Context, from *niosdhcp.RangeCloudInfo, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(RangeCloudInfoAttrTypes)
	}
	m := &RangeCloudInfoModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, RangeCloudInfoAttrTypes, m)
	diags.Append(d...)
	return t
}

// Flatten populates the Terraform model from SDK type
func (m *RangeCloudInfoModel) Flatten(ctx context.Context, from *niosdhcp.RangeCloudInfo, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	m.DelegatedMember = FlattenRangecloudinfoDelegatedMember(ctx, from.DelegatedMember, diags)
	m.DelegatedScope = flex.FlattenStringPointerEmptyAsNull(from.DelegatedScope)
	m.DelegatedRoot = flex.FlattenStringPointerEmptyAsNull(from.DelegatedRoot)
	m.OwnedByAdaptor = flex.FlattenBoolPointer(from.OwnedByAdaptor)
	m.Usage = flex.FlattenStringPointerEmptyAsNull(from.Usage)
	m.Tenant = flex.FlattenStringPointerEmptyAsNull(from.Tenant)
	m.MgmtPlatform = flex.FlattenStringPointerEmptyAsNull(from.MgmtPlatform)
	m.AuthorityType = flex.FlattenStringPointerEmptyAsNull(from.AuthorityType)
}
