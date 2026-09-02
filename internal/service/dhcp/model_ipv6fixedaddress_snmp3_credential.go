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

// Ipv6fixedaddressSnmp3CredentialModel is the Terraform model for Ipv6fixedaddressSnmp3Credential
type Ipv6fixedaddressSnmp3CredentialModel struct {
	User                   types.String `tfsdk:"user"`
	AuthenticationProtocol types.String `tfsdk:"authentication_protocol"`
	AuthenticationPassword types.String `tfsdk:"authentication_password"`
	PrivacyProtocol        types.String `tfsdk:"privacy_protocol"`
	PrivacyPassword        types.String `tfsdk:"privacy_password"`
	Comment                types.String `tfsdk:"comment"`
	CredentialGroup        types.String `tfsdk:"credential_group"`
}

// Ipv6fixedaddressSnmp3CredentialAttrTypes contains the attribute types for Ipv6fixedaddressSnmp3CredentialModel
var Ipv6fixedaddressSnmp3CredentialAttrTypes = map[string]attr.Type{
	"user":                    types.StringType,
	"authentication_protocol": types.StringType,
	"authentication_password": types.StringType,
	"privacy_protocol":        types.StringType,
	"privacy_password":        types.StringType,
	"comment":                 types.StringType,
	"credential_group":        types.StringType,
}

// Ipv6fixedaddressSnmp3CredentialResourceSchemaAttributes contains the schema attributes for Ipv6fixedaddressSnmp3CredentialModel
var Ipv6fixedaddressSnmp3CredentialResourceSchemaAttributes = map[string]schema.Attribute{
	"user": schema.StringAttribute{
		Required: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
			customvalidator.ValidateTrimmedString(),
		},
		MarkdownDescription: "The SNMPv3 user name.",
	},
	"authentication_protocol": schema.StringAttribute{
		Validators: []validator.String{
			stringvalidator.OneOf("NONE", "MD5", "SHA", "SHA-224", "SHA-256", "SHA-384", "SHA-512"),
		},
		Required:            true,
		MarkdownDescription: "Authentication protocol for the SNMPv3 user.",
	},
	"authentication_password": schema.StringAttribute{
		Optional: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
			customvalidator.ValidateTrimmedString(),
		},
		MarkdownDescription: "Authentication password for the SNMPv3 user.",
	},
	"privacy_protocol": schema.StringAttribute{
		Validators: []validator.String{
			stringvalidator.OneOf("NONE", "DES", "3DES", "AES", "AES-192", "AES-192C", "AES-256", "AES-256C"),
		},
		Required:            true,
		MarkdownDescription: "Privacy protocol for the SNMPv3 user.",
	},
	"privacy_password": schema.StringAttribute{
		Optional: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
			customvalidator.ValidateTrimmedString(),
		},
		MarkdownDescription: "Privacy password for the SNMPv3 user.",
	},
	"comment": schema.StringAttribute{
		Optional: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
			customvalidator.ValidateTrimmedString(),
		},
		MarkdownDescription: "Comments for the SNMPv3 user.",
	},
	"credential_group": schema.StringAttribute{
		Default:  stringdefault.StaticString("default"),
		Optional: true,
		Computed: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
		},
		MarkdownDescription: "Group for the SNMPv3 credential.",
	},
}

// ExpandIpv6fixedaddressSnmp3Credential converts a Terraform Object to SDK type
func ExpandIpv6fixedaddressSnmp3Credential(ctx context.Context, o types.Object, diags *diag.Diagnostics) *niosdhcp.Ipv6fixedaddressSnmp3Credential {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m Ipv6fixedaddressSnmp3CredentialModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

// Expand converts the Terraform model to SDK type
func (m *Ipv6fixedaddressSnmp3CredentialModel) Expand(ctx context.Context, diags *diag.Diagnostics) *niosdhcp.Ipv6fixedaddressSnmp3Credential {
	if m == nil {
		return nil
	}
	to := &niosdhcp.Ipv6fixedaddressSnmp3Credential{
		User:                   flex.ExpandStringPointerNullAsEmpty(m.User),
		AuthenticationProtocol: flex.ExpandStringPointerNullAsEmpty(m.AuthenticationProtocol),
		AuthenticationPassword: flex.ExpandStringPointerNullAsEmpty(m.AuthenticationPassword),
		PrivacyProtocol:        flex.ExpandStringPointerNullAsEmpty(m.PrivacyProtocol),
		PrivacyPassword:        flex.ExpandStringPointerNullAsEmpty(m.PrivacyPassword),
		Comment:                flex.ExpandStringPointerNullAsEmpty(m.Comment),
		CredentialGroup:        flex.ExpandStringPointerNullAsEmpty(m.CredentialGroup),
	}
	return to
}

// FlattenIpv6fixedaddressSnmp3Credential converts an SDK type to Terraform Object
func FlattenIpv6fixedaddressSnmp3Credential(ctx context.Context, from *niosdhcp.Ipv6fixedaddressSnmp3Credential, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(Ipv6fixedaddressSnmp3CredentialAttrTypes)
	}
	m := &Ipv6fixedaddressSnmp3CredentialModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, Ipv6fixedaddressSnmp3CredentialAttrTypes, m)
	diags.Append(d...)
	return t
}

// Flatten populates the Terraform model from SDK type
func (m *Ipv6fixedaddressSnmp3CredentialModel) Flatten(ctx context.Context, from *niosdhcp.Ipv6fixedaddressSnmp3Credential, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	m.User = flex.FlattenStringPointerEmptyAsNull(from.User)
	m.AuthenticationProtocol = flex.FlattenStringPointerEmptyAsNull(from.AuthenticationProtocol)
	m.AuthenticationPassword = flex.FlattenStringPointerEmptyAsNull(from.AuthenticationPassword)
	m.PrivacyProtocol = flex.FlattenStringPointerEmptyAsNull(from.PrivacyProtocol)
	m.PrivacyPassword = flex.FlattenStringPointerEmptyAsNull(from.PrivacyPassword)
	m.Comment = flex.FlattenStringPointerEmptyAsNull(from.Comment)
	m.CredentialGroup = flex.FlattenStringPointerEmptyAsNull(from.CredentialGroup)
}
