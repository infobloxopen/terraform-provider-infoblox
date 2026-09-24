package clouddiscovery

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	stringplanmodifier "github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
	uddiclouddiscovery "github.com/infobloxopen/universal-ddi-go-client/clouddiscovery"
)

// CredentialConfigModel is the Terraform model for CredentialConfig
type CredentialConfigModel struct {
	AccessIdentifier types.String `tfsdk:"access_identifier"`
	Enclave          types.String `tfsdk:"enclave"`
	Region           types.String `tfsdk:"region"`
}

// CredentialConfigAttrTypes contains the attribute types for CredentialConfigModel
var CredentialConfigAttrTypes = map[string]attr.Type{
	"access_identifier": types.StringType,
	"enclave":           types.StringType,
	"region":            types.StringType,
}

// CredentialConfigResourceSchemaAttributes contains the schema attributes for CredentialConfigModel
var CredentialConfigResourceSchemaAttributes = map[string]schema.Attribute{
	"access_identifier": schema.StringAttribute{
		Optional: true,
		PlanModifiers: []planmodifier.String{
			stringplanmodifier.RequiresReplaceIfConfigured(),
		},
		MarkdownDescription: "",
	},
	"enclave": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "",
	},
	"region": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "",
	},
}

// ExpandCredentialConfig converts a Terraform Object to SDK type
func ExpandCredentialConfig(ctx context.Context, o types.Object, diags *diag.Diagnostics) *uddiclouddiscovery.CredentialConfig {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m CredentialConfigModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

// Expand converts the Terraform model to SDK type
func (m *CredentialConfigModel) Expand(ctx context.Context, diags *diag.Diagnostics) *uddiclouddiscovery.CredentialConfig {
	if m == nil {
		return nil
	}
	to := &uddiclouddiscovery.CredentialConfig{
		AccessIdentifier: flex.ExpandStringPointer(m.AccessIdentifier),
		Enclave:          flex.ExpandStringPointer(m.Enclave),
		Region:           flex.ExpandStringPointer(m.Region),
	}
	return to
}

// FlattenCredentialConfig converts an SDK type to Terraform Object
func FlattenCredentialConfig(ctx context.Context, from *uddiclouddiscovery.CredentialConfig, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(CredentialConfigAttrTypes)
	}
	m := &CredentialConfigModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, CredentialConfigAttrTypes, m)
	diags.Append(d...)
	return t
}

// Flatten populates the Terraform model from SDK type
func (m *CredentialConfigModel) Flatten(ctx context.Context, from *uddiclouddiscovery.CredentialConfig, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	m.AccessIdentifier = flex.FlattenStringPointer(from.AccessIdentifier)
	m.Enclave = flex.FlattenStringPointer(from.Enclave)
	m.Region = flex.FlattenStringPointer(from.Region)
}
