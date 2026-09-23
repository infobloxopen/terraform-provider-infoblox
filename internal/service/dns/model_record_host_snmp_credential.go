package dns

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	niosdns "github.com/infobloxopen/infoblox-nios-go-client/dns"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
	customvalidator "github.com/infobloxopen/terraform-provider-infoblox/internal/validator"
)

// RecordHostSnmpCredentialModel is the Terraform model for RecordHostSnmpCredential
type RecordHostSnmpCredentialModel struct {
	CommunityString types.String `tfsdk:"community_string"`
	Comment         types.String `tfsdk:"comment"`
	CredentialGroup types.String `tfsdk:"credential_group"`
}

// RecordHostSnmpCredentialAttrTypes contains the attribute types for RecordHostSnmpCredentialModel
var RecordHostSnmpCredentialAttrTypes = map[string]attr.Type{
	"community_string": types.StringType,
	"comment":          types.StringType,
	"credential_group": types.StringType,
}

// RecordHostSnmpCredentialResourceSchemaAttributes contains the schema attributes for RecordHostSnmpCredentialModel
var RecordHostSnmpCredentialResourceSchemaAttributes = map[string]schema.Attribute{
	"community_string": schema.StringAttribute{
		Optional: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
		},
		MarkdownDescription: "The public community string.",
	},
	"comment": schema.StringAttribute{
		Optional: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
		},
		MarkdownDescription: "Comments for the SNMPv1 and SNMPv2 users.",
	},
	"credential_group": schema.StringAttribute{
		Optional: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
		},
		MarkdownDescription: "Group for the SNMPv1 and SNMPv2 credential.",
	},
}

// ExpandRecordHostSnmpCredential converts a Terraform Object to SDK type
func ExpandRecordHostSnmpCredential(ctx context.Context, o types.Object, diags *diag.Diagnostics) *niosdns.RecordHostSnmpCredential {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m RecordHostSnmpCredentialModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

// Expand converts the Terraform model to SDK type
func (m *RecordHostSnmpCredentialModel) Expand(ctx context.Context, diags *diag.Diagnostics) *niosdns.RecordHostSnmpCredential {
	if m == nil {
		return nil
	}
	to := &niosdns.RecordHostSnmpCredential{
		CommunityString: flex.ExpandStringPointerNullAsEmpty(m.CommunityString),
		Comment:         flex.ExpandStringPointerNullAsEmpty(m.Comment),
		CredentialGroup: flex.ExpandStringPointerNullAsEmpty(m.CredentialGroup),
	}
	return to
}

// FlattenRecordHostSnmpCredential converts an SDK type to Terraform Object
func FlattenRecordHostSnmpCredential(ctx context.Context, from *niosdns.RecordHostSnmpCredential, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(RecordHostSnmpCredentialAttrTypes)
	}
	m := &RecordHostSnmpCredentialModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, RecordHostSnmpCredentialAttrTypes, m)
	diags.Append(d...)
	return t
}

// Flatten populates the Terraform model from SDK type
func (m *RecordHostSnmpCredentialModel) Flatten(ctx context.Context, from *niosdns.RecordHostSnmpCredential, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	m.CommunityString = flex.FlattenStringPointerEmptyAsNull(from.CommunityString)
	m.Comment = flex.FlattenStringPointerEmptyAsNull(from.Comment)
	m.CredentialGroup = flex.FlattenStringPointerEmptyAsNull(from.CredentialGroup)
}
