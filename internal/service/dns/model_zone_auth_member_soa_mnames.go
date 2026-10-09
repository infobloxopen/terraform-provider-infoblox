package dns

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	niosdns "github.com/infobloxopen/infoblox-nios-go-client/dns"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
	customvalidator "github.com/infobloxopen/terraform-provider-infoblox/internal/validator"
)

// ZoneAuthMemberSoaMnamesModel is the Terraform model for ZoneAuthMemberSoaMnames
type ZoneAuthMemberSoaMnamesModel struct {
	GridPrimary     types.String `tfsdk:"grid_primary"`
	MsServerPrimary types.String `tfsdk:"ms_server_primary"`
	Mname           types.String `tfsdk:"mname"`
}

// ZoneAuthMemberSoaMnamesAttrTypes contains the attribute types for ZoneAuthMemberSoaMnamesModel
var ZoneAuthMemberSoaMnamesAttrTypes = map[string]attr.Type{
	"grid_primary":      types.StringType,
	"ms_server_primary": types.StringType,
	"mname":             types.StringType,
}

// ZoneAuthMemberSoaMnamesResourceSchemaAttributes contains the schema attributes for ZoneAuthMemberSoaMnamesModel
var ZoneAuthMemberSoaMnamesResourceSchemaAttributes = map[string]schema.Attribute{
	"grid_primary": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
			stringvalidator.ConflictsWith(path.MatchRelative().AtParent().AtName("ms_server_primary")),
		},
		MarkdownDescription: "The grid primary server for the zone. Only one of \"grid_primary\" or \"ms_server_primary\" should be set when modifying or creating the object.",
	},
	"ms_server_primary": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
		},
		MarkdownDescription: "The primary MS server for the zone. Only one of \"grid_primary\" or \"ms_server_primary\" should be set when modifying or creating the object.",
	},
	"mname": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
		},
		MarkdownDescription: "Master's SOA MNAME. This value can be in unicode format.",
	},
}

// ExpandZoneAuthMemberSoaMnames converts a Terraform Object to SDK type
func ExpandZoneAuthMemberSoaMnames(ctx context.Context, o types.Object, diags *diag.Diagnostics) *niosdns.ZoneAuthMemberSoaMnames {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m ZoneAuthMemberSoaMnamesModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

// Expand converts the Terraform model to SDK type
func (m *ZoneAuthMemberSoaMnamesModel) Expand(ctx context.Context, diags *diag.Diagnostics) *niosdns.ZoneAuthMemberSoaMnames {
	if m == nil {
		return nil
	}
	to := &niosdns.ZoneAuthMemberSoaMnames{
		GridPrimary:     flex.ExpandStringPointer(m.GridPrimary),
		MsServerPrimary: flex.ExpandStringPointer(m.MsServerPrimary),
		Mname:           flex.ExpandStringPointerNullAsEmpty(m.Mname),
	}
	return to
}

// FlattenZoneAuthMemberSoaMnames converts an SDK type to Terraform Object
func FlattenZoneAuthMemberSoaMnames(ctx context.Context, from *niosdns.ZoneAuthMemberSoaMnames, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(ZoneAuthMemberSoaMnamesAttrTypes)
	}
	m := &ZoneAuthMemberSoaMnamesModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, ZoneAuthMemberSoaMnamesAttrTypes, m)
	diags.Append(d...)
	return t
}

// Flatten populates the Terraform model from SDK type
func (m *ZoneAuthMemberSoaMnamesModel) Flatten(ctx context.Context, from *niosdns.ZoneAuthMemberSoaMnames, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	m.GridPrimary = flex.FlattenStringPointerEmptyAsNull(from.GridPrimary)
	m.MsServerPrimary = flex.FlattenStringPointerEmptyAsNull(from.MsServerPrimary)
	m.Mname = flex.FlattenStringPointerEmptyAsNull(from.Mname)
}
