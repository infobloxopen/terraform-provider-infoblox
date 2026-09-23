package dns

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

	niosdns "github.com/infobloxopen/infoblox-nios-go-client/dns"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
	customvalidator "github.com/infobloxopen/terraform-provider-infoblox/internal/validator"
)

// RecordHostCliCredentialsModel is the Terraform model for RecordHostCliCredentials
type RecordHostCliCredentialsModel struct {
	User            types.String `tfsdk:"user"`
	Password        types.String `tfsdk:"password"`
	CredentialType  types.String `tfsdk:"credential_type"`
	Comment         types.String `tfsdk:"comment"`
	Id              types.Int64  `tfsdk:"id"`
	CredentialGroup types.String `tfsdk:"credential_group"`
}

// RecordHostCliCredentialsAttrTypes contains the attribute types for RecordHostCliCredentialsModel
var RecordHostCliCredentialsAttrTypes = map[string]attr.Type{
	"user":             types.StringType,
	"password":         types.StringType,
	"credential_type":  types.StringType,
	"comment":          types.StringType,
	"id":               types.Int64Type,
	"credential_group": types.StringType,
}

// RecordHostCliCredentialsResourceSchemaAttributes contains the schema attributes for RecordHostCliCredentialsModel
var RecordHostCliCredentialsResourceSchemaAttributes = map[string]schema.Attribute{
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

// ExpandRecordHostCliCredentials converts a Terraform Object to SDK type
func ExpandRecordHostCliCredentials(ctx context.Context, o types.Object, diags *diag.Diagnostics) *niosdns.RecordHostCliCredentials {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m RecordHostCliCredentialsModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

// Expand converts the Terraform model to SDK type
func (m *RecordHostCliCredentialsModel) Expand(ctx context.Context, diags *diag.Diagnostics) *niosdns.RecordHostCliCredentials {
	if m == nil {
		return nil
	}
	to := &niosdns.RecordHostCliCredentials{
		User:            flex.ExpandStringPointerNullAsEmpty(m.User),
		Password:        flex.ExpandStringPointerNullAsEmpty(m.Password),
		CredentialType:  flex.ExpandStringPointerNullAsEmpty(m.CredentialType),
		Comment:         flex.ExpandStringPointerNullAsEmpty(m.Comment),
		Id:              flex.ExpandInt64Pointer(m.Id),
		CredentialGroup: flex.ExpandStringPointerNullAsEmpty(m.CredentialGroup),
	}
	return to
}

// FlattenRecordHostCliCredentials converts an SDK type to Terraform Object
func FlattenRecordHostCliCredentials(ctx context.Context, from *niosdns.RecordHostCliCredentials, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(RecordHostCliCredentialsAttrTypes)
	}
	m := &RecordHostCliCredentialsModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, RecordHostCliCredentialsAttrTypes, m)
	diags.Append(d...)
	return t
}

// Flatten populates the Terraform model from SDK type
func (m *RecordHostCliCredentialsModel) Flatten(ctx context.Context, from *niosdns.RecordHostCliCredentials, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	m.User = flex.FlattenStringPointerEmptyAsNull(from.User)
	m.Password = flex.FlattenStringPointerEmptyAsNull(from.Password)
	m.CredentialType = flex.FlattenStringPointerEmptyAsNull(from.CredentialType)
	m.Comment = flex.FlattenStringPointerEmptyAsNull(from.Comment)
	m.Id = flex.FlattenInt64Pointer(from.Id)
	m.CredentialGroup = flex.FlattenStringPointerEmptyAsNull(from.CredentialGroup)
}
