package dns

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	niosdns "github.com/infobloxopen/infoblox-nios-go-client/dns"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
	customvalidator "github.com/infobloxopen/terraform-provider-infoblox/internal/validator"
)

// ZoneAuthMsSecondariesModel is the Terraform model for ZoneAuthMsSecondaries
type ZoneAuthMsSecondariesModel struct {
	Address                      types.String `tfsdk:"address"`
	IsMaster                     types.Bool   `tfsdk:"is_master"`
	NsIp                         types.String `tfsdk:"ns_ip"`
	NsName                       types.String `tfsdk:"ns_name"`
	Stealth                      types.Bool   `tfsdk:"stealth"`
	SharedWithMsParentDelegation types.Bool   `tfsdk:"shared_with_ms_parent_delegation"`
}

// ZoneAuthMsSecondariesAttrTypes contains the attribute types for ZoneAuthMsSecondariesModel
var ZoneAuthMsSecondariesAttrTypes = map[string]attr.Type{
	"address":                          types.StringType,
	"is_master":                        types.BoolType,
	"ns_ip":                            types.StringType,
	"ns_name":                          types.StringType,
	"stealth":                          types.BoolType,
	"shared_with_ms_parent_delegation": types.BoolType,
}

// ZoneAuthMsSecondariesResourceSchemaAttributes contains the schema attributes for ZoneAuthMsSecondariesModel
var ZoneAuthMsSecondariesResourceSchemaAttributes = map[string]schema.Attribute{
	"address": schema.StringAttribute{
		Required: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
			customvalidator.IsValidIPOrFQDN(),
		},
		MarkdownDescription: "The address of the server.",
	},
	"is_master": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "This flag indicates if this server is a synchronization master.",
	},
	"ns_ip": schema.StringAttribute{
		Required: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
		},
		MarkdownDescription: "This address is used when generating the NS record in the zone, which can be different in case of multihomed hosts.",
	},
	"ns_name": schema.StringAttribute{
		Required: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
			customvalidator.ValidateTrimmedString(),
		},
		MarkdownDescription: "This name is used when generating the NS record in the zone, which can be different in case of multihomed hosts.",
	},
	"stealth": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Set this flag to hide the NS record for the primary name server from DNS queries.",
	},
	"shared_with_ms_parent_delegation": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "This flag represents whether the name server is shared with the parent Microsoft primary zone's delegation server.",
	},
}

// ExpandZoneAuthMsSecondaries converts a Terraform Object to SDK type
func ExpandZoneAuthMsSecondaries(ctx context.Context, o types.Object, diags *diag.Diagnostics) *niosdns.ZoneAuthMsSecondaries {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m ZoneAuthMsSecondariesModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

// Expand converts the Terraform model to SDK type
func (m *ZoneAuthMsSecondariesModel) Expand(ctx context.Context, diags *diag.Diagnostics) *niosdns.ZoneAuthMsSecondaries {
	if m == nil {
		return nil
	}
	to := &niosdns.ZoneAuthMsSecondaries{
		Address:                      flex.ExpandStringPointerNullAsEmpty(m.Address),
		IsMaster:                     flex.ExpandBoolPointer(m.IsMaster),
		NsIp:                         flex.ExpandStringPointerNullAsEmpty(m.NsIp),
		NsName:                       flex.ExpandStringPointerNullAsEmpty(m.NsName),
		Stealth:                      flex.ExpandBoolPointer(m.Stealth),
		SharedWithMsParentDelegation: flex.ExpandBoolPointer(m.SharedWithMsParentDelegation),
	}
	return to
}

// FlattenZoneAuthMsSecondaries converts an SDK type to Terraform Object
func FlattenZoneAuthMsSecondaries(ctx context.Context, from *niosdns.ZoneAuthMsSecondaries, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(ZoneAuthMsSecondariesAttrTypes)
	}
	m := &ZoneAuthMsSecondariesModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, ZoneAuthMsSecondariesAttrTypes, m)
	diags.Append(d...)
	return t
}

// Flatten populates the Terraform model from SDK type
func (m *ZoneAuthMsSecondariesModel) Flatten(ctx context.Context, from *niosdns.ZoneAuthMsSecondaries, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	m.Address = flex.FlattenStringPointerEmptyAsNull(from.Address)
	m.IsMaster = flex.FlattenBoolPointer(from.IsMaster)
	m.NsIp = flex.FlattenStringPointerEmptyAsNull(from.NsIp)
	m.NsName = flex.FlattenStringPointerEmptyAsNull(from.NsName)
	m.Stealth = flex.FlattenBoolPointer(from.Stealth)
	m.SharedWithMsParentDelegation = flex.FlattenBoolPointer(from.SharedWithMsParentDelegation)
}
