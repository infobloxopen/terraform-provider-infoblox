package fw

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
	customvalidator "github.com/infobloxopen/terraform-provider-infoblox/internal/validator"
	uddifw "github.com/infobloxopen/universal-ddi-go-client/fw"
)

// NetAddrDfpAssignmentModel is the Terraform model for NetAddrDfpAssignment
type NetAddrDfpAssignmentModel struct {
	AddrNet         types.String `tfsdk:"addr_net"`
	DfpIds          types.List   `tfsdk:"dfp_ids"`
	DfpServiceIds   types.List   `tfsdk:"dfp_service_ids"`
	End             types.String `tfsdk:"end"`
	ExternalScopeId types.String `tfsdk:"external_scope_id"`
	HostId          types.String `tfsdk:"host_id"`
	IpSpaceId       types.String `tfsdk:"ip_space_id"`
	ScopeType       types.String `tfsdk:"scope_type"`
	Start           types.String `tfsdk:"start"`
}

// NetAddrDfpAssignmentAttrTypes contains the attribute types for NetAddrDfpAssignmentModel
var NetAddrDfpAssignmentAttrTypes = map[string]attr.Type{
	"addr_net":          types.StringType,
	"dfp_ids":           types.ListType{ElemType: types.Int32Type},
	"dfp_service_ids":   types.ListType{ElemType: types.StringType},
	"end":               types.StringType,
	"external_scope_id": types.StringType,
	"host_id":           types.StringType,
	"ip_space_id":       types.StringType,
	"scope_type":        types.StringType,
	"start":             types.StringType,
}

// NetAddrDfpAssignmentResourceSchemaAttributes contains the schema attributes for NetAddrDfpAssignmentModel
var NetAddrDfpAssignmentResourceSchemaAttributes = map[string]schema.Attribute{
	"addr_net": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "network address in IPv4 CIDR (address/bitmask length) string format",
	},
	"dfp_ids": schema.ListAttribute{
		ElementType: types.Int32Type,
		Optional:    true,
		Computed:    true,
		Validators: []validator.List{
			customvalidator.ListNotEmpty(),
		},
		MarkdownDescription: "The list of identifiers of DFPs that have association with this scope.",
	},
	"dfp_service_ids": schema.ListAttribute{
		ElementType: types.StringType,
		Optional:    true,
		Computed:    true,
		Validators: []validator.List{
			customvalidator.ListNotEmpty(),
		},
		MarkdownDescription: "",
	},
	"end": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "",
	},
	"external_scope_id": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "external scope ID, UUID",
	},
	"host_id": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "Host reference, UUID",
	},
	"ip_space_id": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "IPSpace reference, UUID",
	},
	"scope_type": schema.StringAttribute{
		Default: stringdefault.StaticString("UNKNOWN"),
		Validators: []validator.String{
			stringvalidator.OneOf("UNKNOWN", "ADDRESS_BLOCK", "SUBNET", "ADDRESS", "RANGE"),
		},
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "scope type",
	},
	"start": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "Start and end pair of addresses used for range scope type",
	},
}

// ExpandNetAddrDfpAssignment converts a Terraform Object to SDK type
func ExpandNetAddrDfpAssignment(ctx context.Context, o types.Object, diags *diag.Diagnostics) *uddifw.NetAddrDfpAssignment {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m NetAddrDfpAssignmentModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

// Expand converts the Terraform model to SDK type
func (m *NetAddrDfpAssignmentModel) Expand(ctx context.Context, diags *diag.Diagnostics) *uddifw.NetAddrDfpAssignment {
	if m == nil {
		return nil
	}
	to := &uddifw.NetAddrDfpAssignment{
		AddrNet:         flex.ExpandStringPointer(m.AddrNet),
		DfpIds:          flex.ExpandFrameworkListInt32(ctx, m.DfpIds, diags),
		DfpServiceIds:   flex.ExpandFrameworkListString(ctx, m.DfpServiceIds, diags),
		End:             flex.ExpandStringPointer(m.End),
		ExternalScopeId: flex.ExpandStringPointer(m.ExternalScopeId),
		HostId:          flex.ExpandStringPointer(m.HostId),
		IpSpaceId:       flex.ExpandStringPointer(m.IpSpaceId),
		ScopeType:       (*uddifw.NetAddrDfpAssignmentScopeType)(flex.ExpandStringPointer(m.ScopeType)),
		Start:           flex.ExpandStringPointer(m.Start),
	}
	return to
}

// FlattenNetAddrDfpAssignment converts an SDK type to Terraform Object
func FlattenNetAddrDfpAssignment(ctx context.Context, from *uddifw.NetAddrDfpAssignment, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(NetAddrDfpAssignmentAttrTypes)
	}
	m := &NetAddrDfpAssignmentModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, NetAddrDfpAssignmentAttrTypes, m)
	diags.Append(d...)
	return t
}

// Flatten populates the Terraform model from SDK type
func (m *NetAddrDfpAssignmentModel) Flatten(ctx context.Context, from *uddifw.NetAddrDfpAssignment, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	m.AddrNet = flex.FlattenStringPointer(from.AddrNet)
	m.DfpIds = flex.FlattenFrameworkListInt32(ctx, from.DfpIds, diags)
	m.DfpServiceIds = flex.FlattenFrameworkListString(ctx, from.DfpServiceIds, diags)
	m.End = flex.FlattenStringPointer(from.End)
	m.ExternalScopeId = flex.FlattenStringPointer(from.ExternalScopeId)
	m.HostId = flex.FlattenStringPointer(from.HostId)
	m.IpSpaceId = flex.FlattenStringPointer(from.IpSpaceId)
	m.ScopeType = flex.FlattenStringPointer((*string)(from.ScopeType))
	m.Start = flex.FlattenStringPointer(from.Start)
}
