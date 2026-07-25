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

	"github.com/hashicorp/terraform-plugin-framework/path"
	niosdns "github.com/infobloxopen/infoblox-nios-go-client/dns"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
	customvalidator "github.com/infobloxopen/terraform-provider-infoblox/internal/validator"
)

// ZoneAuthAllowUpdateModel is the Terraform model for ZoneAuthAllowUpdate
type ZoneAuthAllowUpdateModel struct {
	Struct         types.String `tfsdk:"struct"`
	Ref            types.String `tfsdk:"ref"`
	Address        types.String `tfsdk:"address"`
	Permission     types.String `tfsdk:"permission"`
	TsigKey        types.String `tfsdk:"tsig_key"`
	TsigKeyAlg     types.String `tfsdk:"tsig_key_alg"`
	TsigKeyName    types.String `tfsdk:"tsig_key_name"`
	UseTsigKeyName types.Bool   `tfsdk:"use_tsig_key_name"`
}

// ZoneAuthAllowUpdateAttrTypes contains the attribute types for ZoneAuthAllowUpdateModel
var ZoneAuthAllowUpdateAttrTypes = map[string]attr.Type{
	"struct":            types.StringType,
	"ref":               types.StringType,
	"address":           types.StringType,
	"permission":        types.StringType,
	"tsig_key":          types.StringType,
	"tsig_key_alg":      types.StringType,
	"tsig_key_name":     types.StringType,
	"use_tsig_key_name": types.BoolType,
}

// ZoneAuthAllowUpdateResourceSchemaAttributes contains the schema attributes for ZoneAuthAllowUpdateModel
var ZoneAuthAllowUpdateResourceSchemaAttributes = map[string]schema.Attribute{
	"struct": schema.StringAttribute{
		Optional: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
		},
		MarkdownDescription: "The struct type of the object. The value must be one of 'addressac' and 'tsigac'.",
	},
	"ref": schema.StringAttribute{
		Optional: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
			stringvalidator.ConflictsWith(
				path.MatchRelative().AtParent().AtName("struct"),
				path.MatchRelative().AtParent().AtName("address"),
				path.MatchRelative().AtParent().AtName("permission"),
				path.MatchRelative().AtParent().AtName("tsig_key"),
				path.MatchRelative().AtParent().AtName("tsig_key_alg"),
				path.MatchRelative().AtParent().AtName("tsig_key_name"),
			),
		},
		MarkdownDescription: "The reference to the object.",
	},
	"address": schema.StringAttribute{
		Optional: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
			stringvalidator.ConflictsWith(
				path.MatchRelative().AtParent().AtName("tsig_key"),
				path.MatchRelative().AtParent().AtName("tsig_key_alg"),
				path.MatchRelative().AtParent().AtName("tsig_key_name"),
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
				path.MatchRelative().AtParent().AtName("tsig_key_name"),
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
			customvalidator.ValidateTrimmedString(),
		},
		MarkdownDescription: "A generated TSIG key. If the external primary server is a NIOS appliance running DNS One 2.x code, this can be set to :2xCOMPAT.",
	},
	"tsig_key_alg": schema.StringAttribute{
		Validators: []validator.String{
			stringvalidator.OneOf("HMAC-MD5", "HMAC-SHA256"),
		},
		Optional:            true,
		MarkdownDescription: "The TSIG key algorithm.",
	},
	"tsig_key_name": schema.StringAttribute{
		Optional: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
			customvalidator.ValidateTrimmedString(),
		},
		MarkdownDescription: "The name of the TSIG key. If 2.x TSIG compatibility is used, this is set to 'tsig_xfer' on retrieval, and ignored on insert or update.",
	},
	"use_tsig_key_name": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Use flag for: tsig_key_name",
	},
}

// ExpandZoneAuthAllowUpdate converts a Terraform Object to SDK type
func ExpandZoneAuthAllowUpdate(ctx context.Context, o types.Object, diags *diag.Diagnostics) *niosdns.ZoneAuthAllowUpdate {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m ZoneAuthAllowUpdateModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

// Expand converts the Terraform model to SDK type
func (m *ZoneAuthAllowUpdateModel) Expand(ctx context.Context, diags *diag.Diagnostics) *niosdns.ZoneAuthAllowUpdate {
	if m == nil {
		return nil
	}
	to := &niosdns.ZoneAuthAllowUpdate{
		Struct:         flex.ExpandStringPointerNullAsEmpty(m.Struct),
		Ref:            flex.ExpandStringPointerNullAsEmpty(m.Ref),
		Address:        flex.ExpandStringPointerNullAsEmpty(m.Address),
		Permission:     flex.ExpandStringPointerNullAsEmpty(m.Permission),
		TsigKey:        flex.ExpandStringPointerNullAsEmpty(m.TsigKey),
		TsigKeyAlg:     flex.ExpandStringPointerNullAsEmpty(m.TsigKeyAlg),
		TsigKeyName:    flex.ExpandStringPointerNullAsEmpty(m.TsigKeyName),
		UseTsigKeyName: flex.ExpandBoolPointer(m.UseTsigKeyName),
	}
	return to
}

// FlattenZoneAuthAllowUpdate converts an SDK type to Terraform Object
func FlattenZoneAuthAllowUpdate(ctx context.Context, from *niosdns.ZoneAuthAllowUpdate, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(ZoneAuthAllowUpdateAttrTypes)
	}
	m := &ZoneAuthAllowUpdateModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, ZoneAuthAllowUpdateAttrTypes, m)
	diags.Append(d...)
	return t
}

// Flatten populates the Terraform model from SDK type
func (m *ZoneAuthAllowUpdateModel) Flatten(ctx context.Context, from *niosdns.ZoneAuthAllowUpdate, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	m.Struct = flex.FlattenStringPointerEmptyAsNull(from.Struct)
	m.Ref = flex.FlattenStringPointerEmptyAsNull(from.Ref)
	m.Address = flex.FlattenStringPointerEmptyAsNull(from.Address)
	m.Permission = flex.FlattenStringPointerEmptyAsNull(from.Permission)
	m.TsigKey = flex.FlattenStringPointerEmptyAsNull(from.TsigKey)
	m.TsigKeyAlg = flex.FlattenStringPointerEmptyAsNull(from.TsigKeyAlg)
	m.TsigKeyName = flex.FlattenStringPointerEmptyAsNull(from.TsigKeyName)
	m.UseTsigKeyName = flex.FlattenBoolPointer(from.UseTsigKeyName)
}
