package dns

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/hashicorp/terraform-plugin-framework-validators/boolvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	niosdns "github.com/infobloxopen/infoblox-nios-go-client/dns"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
	customvalidator "github.com/infobloxopen/terraform-provider-infoblox/internal/validator"
)

// ZoneAuthUpdateForwardingModel is the Terraform model for ZoneAuthUpdateForwarding
type ZoneAuthUpdateForwardingModel struct {
	Struct         types.String `tfsdk:"struct"`
	Address        types.String `tfsdk:"address"`
	Permission     types.String `tfsdk:"permission"`
	TsigKey        types.String `tfsdk:"tsig_key"`
	TsigKeyAlg     types.String `tfsdk:"tsig_key_alg"`
	TsigKeyName    types.String `tfsdk:"tsig_key_name"`
	UseTsigKeyName types.Bool   `tfsdk:"use_tsig_key_name"`
}

// ZoneAuthUpdateForwardingAttrTypes contains the attribute types for ZoneAuthUpdateForwardingModel
var ZoneAuthUpdateForwardingAttrTypes = map[string]attr.Type{
	"struct":            types.StringType,
	"address":           types.StringType,
	"permission":        types.StringType,
	"tsig_key":          types.StringType,
	"tsig_key_alg":      types.StringType,
	"tsig_key_name":     types.StringType,
	"use_tsig_key_name": types.BoolType,
}

// ZoneAuthUpdateForwardingResourceSchemaAttributes contains the schema attributes for ZoneAuthUpdateForwardingModel
var ZoneAuthUpdateForwardingResourceSchemaAttributes = map[string]schema.Attribute{
	"struct": schema.StringAttribute{
		Required: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
		},
		MarkdownDescription: "The struct type of the object. The value must be one of 'addressac' and 'tsigac'.",
	},
	"address": schema.StringAttribute{
		Optional: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
			stringvalidator.ConflictsWith(
				path.MatchRelative().AtParent().AtName("tsig_key"),
				path.MatchRelative().AtParent().AtName("tsig_key_alg"),
				path.MatchRelative().AtParent().AtName("use_tsig_key_name"),
			),
			customvalidator.ValidateTrimmedString(),
		},
		MarkdownDescription: "The address this rule applies to or \"Any\".",
	},
	"permission": schema.StringAttribute{
		Validators: []validator.String{
			stringvalidator.ConflictsWith(
				path.MatchRelative().AtParent().AtName("tsig_key"),
				path.MatchRelative().AtParent().AtName("tsig_key_alg"),
				path.MatchRelative().AtParent().AtName("use_tsig_key_name"),
			),
			stringvalidator.OneOf("ALLOW", "DENY"),
		},
		Optional:            true,
		MarkdownDescription: "The permission to use for this address.",
	},
	"tsig_key": schema.StringAttribute{
		Optional: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
			stringvalidator.ConflictsWith(
				path.MatchRelative().AtParent().AtName("address"),
				path.MatchRelative().AtParent().AtName("permission"),
			),
			customvalidator.ValidateTrimmedString(),
		},
		MarkdownDescription: "A generated TSIG key. If the external primary server is a NIOS appliance running DNS One 2.x code, this can be set to :2xCOMPAT.",
	},
	"tsig_key_alg": schema.StringAttribute{
		Validators: []validator.String{
			stringvalidator.ConflictsWith(
				path.MatchRelative().AtParent().AtName("address"),
				path.MatchRelative().AtParent().AtName("permission"),
			),
			stringvalidator.OneOf("HMAC-MD5", "HMAC-SHA256"),
		},
		Optional:            true,
		MarkdownDescription: "The TSIG key algorithm.",
	},
	"tsig_key_name": schema.StringAttribute{
		Optional: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
			stringvalidator.ConflictsWith(
				path.MatchRelative().AtParent().AtName("address"),
				path.MatchRelative().AtParent().AtName("permission"),
			),
			customvalidator.ValidateTrimmedString(),
		},
		MarkdownDescription: "The name of the TSIG key. If 2.x TSIG compatibility is used, this is set to 'tsig_xfer' on retrieval, and ignored on insert or update.",
	},
	"use_tsig_key_name": schema.BoolAttribute{
		Optional: true,
		Validators: []validator.Bool{
			boolvalidator.ConflictsWith(
				path.MatchRelative().AtParent().AtName("address"),
				path.MatchRelative().AtParent().AtName("permission"),
			),
		},
		MarkdownDescription: "Use flag for: tsig_key_name",
	},
}

// ExpandZoneAuthUpdateForwarding converts a Terraform Object to SDK type
func ExpandZoneAuthUpdateForwarding(ctx context.Context, o types.Object, diags *diag.Diagnostics) *niosdns.ZoneAuthUpdateForwarding {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m ZoneAuthUpdateForwardingModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

// Expand converts the Terraform model to SDK type
func (m *ZoneAuthUpdateForwardingModel) Expand(ctx context.Context, diags *diag.Diagnostics) *niosdns.ZoneAuthUpdateForwarding {
	if m == nil {
		return nil
	}
	to := &niosdns.ZoneAuthUpdateForwarding{
		Struct:         flex.ExpandStringPointerNullAsEmpty(m.Struct),
		Address:        flex.ExpandStringPointerNullAsEmpty(m.Address),
		Permission:     flex.ExpandStringPointerNullAsEmpty(m.Permission),
		TsigKey:        flex.ExpandStringPointerNullAsEmpty(m.TsigKey),
		TsigKeyAlg:     flex.ExpandStringPointerNullAsEmpty(m.TsigKeyAlg),
		TsigKeyName:    flex.ExpandStringPointerNullAsEmpty(m.TsigKeyName),
		UseTsigKeyName: flex.ExpandBoolPointer(m.UseTsigKeyName),
	}
	return to
}

// FlattenZoneAuthUpdateForwarding converts an SDK type to Terraform Object
func FlattenZoneAuthUpdateForwarding(ctx context.Context, from *niosdns.ZoneAuthUpdateForwarding, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(ZoneAuthUpdateForwardingAttrTypes)
	}
	m := &ZoneAuthUpdateForwardingModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, ZoneAuthUpdateForwardingAttrTypes, m)
	diags.Append(d...)
	return t
}

// Flatten populates the Terraform model from SDK type
func (m *ZoneAuthUpdateForwardingModel) Flatten(ctx context.Context, from *niosdns.ZoneAuthUpdateForwarding, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	m.Struct = flex.FlattenStringPointerEmptyAsNull(from.Struct)
	m.Address = flex.FlattenStringPointerEmptyAsNull(from.Address)
	m.Permission = flex.FlattenStringPointerEmptyAsNull(from.Permission)
	m.TsigKey = flex.FlattenStringPointerEmptyAsNull(from.TsigKey)
	m.TsigKeyAlg = flex.FlattenStringPointerEmptyAsNull(from.TsigKeyAlg)
	m.TsigKeyName = flex.FlattenStringPointerEmptyAsNull(from.TsigKeyName)
	m.UseTsigKeyName = flex.FlattenBoolPointer(from.UseTsigKeyName)
}
