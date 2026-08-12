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

// ViewMatchDestinationsModel is the Terraform model for ViewMatchDestinations
type ViewMatchDestinationsModel struct {
	Struct         types.String `tfsdk:"struct"`
	Ref            types.String `tfsdk:"ref"`
	Address        types.String `tfsdk:"address"`
	Permission     types.String `tfsdk:"permission"`
	TsigKey        types.String `tfsdk:"tsig_key"`
	TsigKeyAlg     types.String `tfsdk:"tsig_key_alg"`
	TsigKeyName    types.String `tfsdk:"tsig_key_name"`
	UseTsigKeyName types.Bool   `tfsdk:"use_tsig_key_name"`
}

// ViewMatchDestinationsAttrTypes contains the attribute types for ViewMatchDestinationsModel
var ViewMatchDestinationsAttrTypes = map[string]attr.Type{
	"struct":            types.StringType,
	"ref":               types.StringType,
	"address":           types.StringType,
	"permission":        types.StringType,
	"tsig_key":          types.StringType,
	"tsig_key_alg":      types.StringType,
	"tsig_key_name":     types.StringType,
	"use_tsig_key_name": types.BoolType,
}

// ViewMatchDestinationsResourceSchemaAttributes contains the schema attributes for ViewMatchDestinationsModel
var ViewMatchDestinationsResourceSchemaAttributes = map[string]schema.Attribute{
	"struct": schema.StringAttribute{
		Validators: []validator.String{
			stringvalidator.OneOf("addressac", "tsigac"),
		},
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The struct type of the object. The value must be one of 'addressac' and 'tsigac'.",
	},
	"ref": schema.StringAttribute{
		Optional: true,
		Computed: true,
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
		Computed: true,
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
		Computed:            true,
		MarkdownDescription: "The permission to use for this address.",
	},
	"tsig_key": schema.StringAttribute{
		Optional: true,
		Computed: true,
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
		Computed:            true,
		MarkdownDescription: "The TSIG key algorithm.",
	},
	"tsig_key_name": schema.StringAttribute{
		Optional: true,
		Computed: true,
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
		Computed: true,
		Validators: []validator.Bool{
			boolvalidator.ConflictsWith(
				path.MatchRelative().AtParent().AtName("address"),
				path.MatchRelative().AtParent().AtName("permission"),
			),
		},
		MarkdownDescription: "Use flag for: tsig_key_name",
	},
}

// ExpandViewMatchDestinations converts a Terraform Object to SDK type
func ExpandViewMatchDestinations(ctx context.Context, o types.Object, diags *diag.Diagnostics) *niosdns.ViewMatchDestinations {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m ViewMatchDestinationsModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

// Expand converts the Terraform model to SDK type
func (m *ViewMatchDestinationsModel) Expand(ctx context.Context, diags *diag.Diagnostics) *niosdns.ViewMatchDestinations {
	if m == nil {
		return nil
	}
	to := &niosdns.ViewMatchDestinations{
		Struct:         flex.ExpandStringPointerNullAsEmpty(m.Struct),
		Ref:            flex.ExpandStringPointer(m.Ref),
		Address:        flex.ExpandStringPointer(m.Address),
		Permission:     flex.ExpandStringPointer(m.Permission),
		TsigKey:        flex.ExpandStringPointer(m.TsigKey),
		TsigKeyAlg:     flex.ExpandStringPointer(m.TsigKeyAlg),
		TsigKeyName:    flex.ExpandStringPointer(m.TsigKeyName),
		UseTsigKeyName: flex.ExpandBoolPointer(m.UseTsigKeyName),
	}
	return to
}

// FlattenViewMatchDestinations converts an SDK type to Terraform Object
func FlattenViewMatchDestinations(ctx context.Context, from *niosdns.ViewMatchDestinations, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(ViewMatchDestinationsAttrTypes)
	}
	m := &ViewMatchDestinationsModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, ViewMatchDestinationsAttrTypes, m)
	diags.Append(d...)
	return t
}

// Flatten populates the Terraform model from SDK type
func (m *ViewMatchDestinationsModel) Flatten(ctx context.Context, from *niosdns.ViewMatchDestinations, diags *diag.Diagnostics) {
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
