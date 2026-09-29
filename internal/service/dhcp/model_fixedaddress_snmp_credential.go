package dhcp

import (
	"context"

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

// FixedaddressSnmpCredentialModel is the Terraform model for FixedaddressSnmpCredential
type FixedaddressSnmpCredentialModel struct {
	CommunityString types.String `tfsdk:"community_string"`
	Comment         types.String `tfsdk:"comment"`
	CredentialGroup types.String `tfsdk:"credential_group"`
}

// FixedaddressSnmpCredentialAttrTypes contains the attribute types for FixedaddressSnmpCredentialModel
var FixedaddressSnmpCredentialAttrTypes = map[string]attr.Type{
	"community_string": types.StringType,
	"comment":          types.StringType,
	"credential_group": types.StringType,
}

// FixedaddressSnmpCredentialResourceSchemaAttributes contains the schema attributes for FixedaddressSnmpCredentialModel
var FixedaddressSnmpCredentialResourceSchemaAttributes = map[string]schema.Attribute{
	"community_string": schema.StringAttribute{
		Optional: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
			customvalidator.ValidateTrimmedString(),
		},
		MarkdownDescription: "The public community string.",
	},
	"comment": schema.StringAttribute{
		Optional: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
			customvalidator.ValidateTrimmedString(),
		},
		MarkdownDescription: "Comments for the SNMPv1 and SNMPv2 users.",
	},
	"credential_group": schema.StringAttribute{
		Default:  stringdefault.StaticString("default"),
		Optional: true,
		Computed: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
		},
		MarkdownDescription: "Group for the SNMPv1 and SNMPv2 credential.",
	},
}

// ExpandFixedaddressSnmpCredential converts a Terraform Object to SDK type
func ExpandFixedaddressSnmpCredential(ctx context.Context, o types.Object, diags *diag.Diagnostics) *niosdhcp.FixedaddressSnmpCredential {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m FixedaddressSnmpCredentialModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

// Expand converts the Terraform model to SDK type
func (m *FixedaddressSnmpCredentialModel) Expand(ctx context.Context, diags *diag.Diagnostics) *niosdhcp.FixedaddressSnmpCredential {
	if m == nil {
		return nil
	}
	to := &niosdhcp.FixedaddressSnmpCredential{
		CommunityString: flex.ExpandStringPointerNullAsEmpty(m.CommunityString),
		Comment:         flex.ExpandStringPointerNullAsEmpty(m.Comment),
		CredentialGroup: flex.ExpandStringPointerNullAsEmpty(m.CredentialGroup),
	}
	return to
}

// FlattenFixedaddressSnmpCredential converts an SDK type to Terraform Object
func FlattenFixedaddressSnmpCredential(ctx context.Context, from *niosdhcp.FixedaddressSnmpCredential, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(FixedaddressSnmpCredentialAttrTypes)
	}
	m := &FixedaddressSnmpCredentialModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, FixedaddressSnmpCredentialAttrTypes, m)
	diags.Append(d...)
	return t
}

// Flatten populates the Terraform model from SDK type
func (m *FixedaddressSnmpCredentialModel) Flatten(ctx context.Context, from *niosdhcp.FixedaddressSnmpCredential, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	m.CommunityString = flex.FlattenStringPointerEmptyAsNull(from.CommunityString)
	m.Comment = flex.FlattenStringPointerEmptyAsNull(from.Comment)
	m.CredentialGroup = flex.FlattenStringPointerEmptyAsNull(from.CredentialGroup)
}
