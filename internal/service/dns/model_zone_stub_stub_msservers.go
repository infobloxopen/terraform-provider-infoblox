package dns

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/hashicorp/terraform-plugin-framework-nettypes/iptypes"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	niosdns "github.com/infobloxopen/infoblox-nios-go-client/dns"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
	customvalidator "github.com/infobloxopen/terraform-provider-infoblox/internal/validator"
)

// ZoneStubStubMsserversModel is the Terraform model for ZoneStubStubMsservers
type ZoneStubStubMsserversModel struct {
	Address  iptypes.IPAddress `tfsdk:"address"`
	IsMaster types.Bool        `tfsdk:"is_master"`
	NsIp     types.String      `tfsdk:"ns_ip"`
	NsName   types.String      `tfsdk:"ns_name"`
	Stealth  types.Bool        `tfsdk:"stealth"`
}

// ZoneStubStubMsserversAttrTypes contains the attribute types for ZoneStubStubMsserversModel
var ZoneStubStubMsserversAttrTypes = map[string]attr.Type{
	"address":   iptypes.IPAddressType{},
	"is_master": types.BoolType,
	"ns_ip":     types.StringType,
	"ns_name":   types.StringType,
	"stealth":   types.BoolType,
}

// ZoneStubStubMsserversResourceSchemaAttributes contains the schema attributes for ZoneStubStubMsserversModel
var ZoneStubStubMsserversResourceSchemaAttributes = map[string]schema.Attribute{
	"address": schema.StringAttribute{
		Required:   true,
		CustomType: iptypes.IPAddressType{},
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
		},
		MarkdownDescription: "The address of the server.",
	},
	"is_master": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
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
}

// ExpandZoneStubStubMsservers converts a Terraform Object to SDK type
func ExpandZoneStubStubMsservers(ctx context.Context, o types.Object, diags *diag.Diagnostics) *niosdns.ZoneStubStubMsservers {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m ZoneStubStubMsserversModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

// Expand converts the Terraform model to SDK type
func (m *ZoneStubStubMsserversModel) Expand(ctx context.Context, diags *diag.Diagnostics) *niosdns.ZoneStubStubMsservers {
	if m == nil {
		return nil
	}
	to := &niosdns.ZoneStubStubMsservers{
		Address:  flex.ExpandIPAddress(m.Address),
		IsMaster: flex.ExpandBoolPointer(m.IsMaster),
		NsIp:     flex.ExpandStringPointerNullAsEmpty(m.NsIp),
		NsName:   flex.ExpandStringPointerNullAsEmpty(m.NsName),
		Stealth:  flex.ExpandBoolPointer(m.Stealth),
	}
	return to
}

// FlattenZoneStubStubMsservers converts an SDK type to Terraform Object
func FlattenZoneStubStubMsservers(ctx context.Context, from *niosdns.ZoneStubStubMsservers, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(ZoneStubStubMsserversAttrTypes)
	}
	m := &ZoneStubStubMsserversModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, ZoneStubStubMsserversAttrTypes, m)
	diags.Append(d...)
	return t
}

// Flatten populates the Terraform model from SDK type
func (m *ZoneStubStubMsserversModel) Flatten(ctx context.Context, from *niosdns.ZoneStubStubMsservers, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	m.Address = flex.FlattenIPAddress(from.Address)
	m.IsMaster = flex.FlattenBoolPointer(from.IsMaster)
	m.NsIp = flex.FlattenStringPointerEmptyAsNull(from.NsIp)
	m.NsName = flex.FlattenStringPointerEmptyAsNull(from.NsName)
	m.Stealth = flex.FlattenBoolPointer(from.Stealth)
}
