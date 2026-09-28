package clouddiscovery

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
	uddiclouddiscovery "github.com/infobloxopen/universal-ddi-go-client/clouddiscovery"
)

// CredentialPreferenceModel is the Terraform model for CredentialPreference
type CredentialPreferenceModel struct {
	AccessIdentifierType types.String `tfsdk:"access_identifier_type"`
	CredentialType       types.String `tfsdk:"credential_type"`
}

// CredentialPreferenceAttrTypes contains the attribute types for CredentialPreferenceModel
var CredentialPreferenceAttrTypes = map[string]attr.Type{
	"access_identifier_type": types.StringType,
	"credential_type":        types.StringType,
}

// CredentialPreferenceResourceSchemaAttributes contains the schema attributes for CredentialPreferenceModel
var CredentialPreferenceResourceSchemaAttributes = map[string]schema.Attribute{
	"access_identifier_type": schema.StringAttribute{
		Validators: []validator.String{
			stringvalidator.OneOf("role_arn", "tenant_id", "project_id"),
		},
		Optional:            true,
		MarkdownDescription: "Access identifier type. Possible values: role_arn, tenant_id, project_id.",
	},
	"credential_type": schema.StringAttribute{
		Validators: []validator.String{
			stringvalidator.OneOf("dynamic", "static"),
		},
		Optional:            true,
		MarkdownDescription: "Credential type. Possible values: dynamic, static.",
	},
}

// ExpandCredentialPreference converts a Terraform Object to SDK type
func ExpandCredentialPreference(ctx context.Context, o types.Object, diags *diag.Diagnostics) *uddiclouddiscovery.CredentialPreference {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m CredentialPreferenceModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

// Expand converts the Terraform model to SDK type
func (m *CredentialPreferenceModel) Expand(ctx context.Context, diags *diag.Diagnostics) *uddiclouddiscovery.CredentialPreference {
	if m == nil {
		return nil
	}
	to := &uddiclouddiscovery.CredentialPreference{
		AccessIdentifierType: flex.ExpandStringPointer(m.AccessIdentifierType),
		CredentialType:       flex.ExpandStringPointer(m.CredentialType),
	}
	return to
}

// FlattenCredentialPreference converts an SDK type to Terraform Object
func FlattenCredentialPreference(ctx context.Context, from *uddiclouddiscovery.CredentialPreference, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(CredentialPreferenceAttrTypes)
	}
	m := &CredentialPreferenceModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, CredentialPreferenceAttrTypes, m)
	diags.Append(d...)
	return t
}

// Flatten populates the Terraform model from SDK type
func (m *CredentialPreferenceModel) Flatten(ctx context.Context, from *uddiclouddiscovery.CredentialPreference, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	m.AccessIdentifierType = flex.FlattenStringPointer(from.AccessIdentifierType)
	m.CredentialType = flex.FlattenStringPointer(from.CredentialType)
}
