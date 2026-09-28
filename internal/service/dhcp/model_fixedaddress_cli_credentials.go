package dhcp

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

	niosdhcp "github.com/infobloxopen/infoblox-nios-go-client/dhcp"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
	customvalidator "github.com/infobloxopen/terraform-provider-infoblox/internal/validator"
)

// FixedaddressCliCredentialsModel is the Terraform model for FixedaddressCliCredentials
type FixedaddressCliCredentialsModel struct {
	User            types.String `tfsdk:"user"`
	Password        types.String `tfsdk:"password"`
	CredentialType  types.String `tfsdk:"credential_type"`
	Comment         types.String `tfsdk:"comment"`
	Id              types.Int64  `tfsdk:"id"`
	CredentialGroup types.String `tfsdk:"credential_group"`
}

// FixedaddressCliCredentialsAttrTypes contains the attribute types for FixedaddressCliCredentialsModel
var FixedaddressCliCredentialsAttrTypes = map[string]attr.Type{
	"user":             types.StringType,
	"password":         types.StringType,
	"credential_type":  types.StringType,
	"comment":          types.StringType,
	"id":               types.Int64Type,
	"credential_group": types.StringType,
}

// FixedaddressCliCredentialsResourceSchemaAttributes contains the schema attributes for FixedaddressCliCredentialsModel
var FixedaddressCliCredentialsResourceSchemaAttributes = map[string]schema.Attribute{
	"user": schema.StringAttribute{
		Optional: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
			customvalidator.ValidateTrimmedString(),
		},
		MarkdownDescription: "The CLI user name.",
	},
	"password": schema.StringAttribute{
		Optional: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
			customvalidator.ValidateTrimmedString(),
		},
		MarkdownDescription: "The CLI password.",
	},
	"credential_type": schema.StringAttribute{
		Validators: []validator.String{
			stringvalidator.OneOf("SSH", "TELNET", "ENABLE_SSH", "ENABLE_TELNET"),
		},
		Required:            true,
		MarkdownDescription: "The type of the credential.",
	},
	"comment": schema.StringAttribute{
		Optional: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
			customvalidator.ValidateTrimmedString(),
		},
		MarkdownDescription: "The commment for the credential.",
	},
	"id": schema.Int64Attribute{
		Computed:            true,
		MarkdownDescription: "The Credentials ID.",
	},
	"credential_group": schema.StringAttribute{
		Default:  stringdefault.StaticString("default"),
		Optional: true,
		Computed: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
		},
		MarkdownDescription: "Group for the CLI credential.",
	},
}

// ExpandFixedaddressCliCredentials converts a Terraform Object to SDK type
func ExpandFixedaddressCliCredentials(ctx context.Context, o types.Object, diags *diag.Diagnostics) *niosdhcp.FixedaddressCliCredentials {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m FixedaddressCliCredentialsModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

// Expand converts the Terraform model to SDK type
func (m *FixedaddressCliCredentialsModel) Expand(ctx context.Context, diags *diag.Diagnostics) *niosdhcp.FixedaddressCliCredentials {
	if m == nil {
		return nil
	}
	to := &niosdhcp.FixedaddressCliCredentials{
		User:            flex.ExpandStringPointerNullAsEmpty(m.User),
		Password:        flex.ExpandStringPointerNullAsEmpty(m.Password),
		CredentialType:  flex.ExpandStringPointerNullAsEmpty(m.CredentialType),
		Comment:         flex.ExpandStringPointerNullAsEmpty(m.Comment),
		Id:              flex.ExpandInt64Pointer(m.Id),
		CredentialGroup: flex.ExpandStringPointerNullAsEmpty(m.CredentialGroup),
	}
	return to
}

// FlattenFixedaddressCliCredentials converts an SDK type to Terraform Object
func FlattenFixedaddressCliCredentials(ctx context.Context, from *niosdhcp.FixedaddressCliCredentials, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(FixedaddressCliCredentialsAttrTypes)
	}
	m := &FixedaddressCliCredentialsModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, FixedaddressCliCredentialsAttrTypes, m)
	diags.Append(d...)
	return t
}

// Flatten populates the Terraform model from SDK type
func (m *FixedaddressCliCredentialsModel) Flatten(ctx context.Context, from *niosdhcp.FixedaddressCliCredentials, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	m.User = flex.FlattenStringPointerEmptyAsNull(from.User)
	m.CredentialType = flex.FlattenStringPointerEmptyAsNull(from.CredentialType)
	m.Comment = flex.FlattenStringPointerEmptyAsNull(from.Comment)
	m.Id = flex.FlattenInt64Pointer(from.Id)
	m.CredentialGroup = flex.FlattenStringPointerEmptyAsNull(from.CredentialGroup)
}
