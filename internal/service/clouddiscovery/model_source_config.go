package clouddiscovery

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	listplanmodifier "github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
	stringplanmodifier "github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
	customvalidator "github.com/infobloxopen/terraform-provider-infoblox/internal/validator"
	uddiclouddiscovery "github.com/infobloxopen/universal-ddi-go-client/clouddiscovery"
)

// SourceConfigModel is the Terraform model for SourceConfig
type SourceConfigModel struct {
	AccountScheduleId    types.String      `tfsdk:"account_schedule_id"`
	Accounts             types.List        `tfsdk:"accounts"`
	CloudCredentialId    types.String      `tfsdk:"cloud_credential_id"`
	CredentialConfig     types.Object      `tfsdk:"credential_config"`
	DeletedAt            timetypes.RFC3339 `tfsdk:"deleted_at"`
	Id                   types.String      `tfsdk:"id"`
	RestrictedToAccounts types.List        `tfsdk:"restricted_to_accounts"`
}

// SourceConfigAttrTypes contains the attribute types for SourceConfigModel
var SourceConfigAttrTypes = map[string]attr.Type{
	"account_schedule_id":    types.StringType,
	"accounts":               types.ListType{ElemType: types.ObjectType{AttrTypes: AccountAttrTypes}},
	"cloud_credential_id":    types.StringType,
	"credential_config":      types.ObjectType{AttrTypes: CredentialConfigAttrTypes},
	"deleted_at":             timetypes.RFC3339Type{},
	"id":                     types.StringType,
	"restricted_to_accounts": types.ListType{ElemType: types.StringType},
}

// SourceConfigResourceSchemaAttributes contains the schema attributes for SourceConfigModel
var SourceConfigResourceSchemaAttributes = map[string]schema.Attribute{
	"account_schedule_id": schema.StringAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "Account Schedule ID.",
	},
	"accounts": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: AccountResourceSchemaAttributes,
		},
		Computed: true,
		Validators: []validator.List{
			customvalidator.ListNotEmpty(),
		},
		MarkdownDescription: "List of accounts to be discovered.",
	},
	"cloud_credential_id": schema.StringAttribute{
		Optional: true,
		PlanModifiers: []planmodifier.String{
			stringplanmodifier.RequiresReplaceIfConfigured(),
		},
		MarkdownDescription: "Cloud Credential ID.",
	},
	"credential_config": schema.SingleNestedAttribute{
		Attributes:          CredentialConfigResourceSchemaAttributes,
		Optional:            true,
		MarkdownDescription: "Credential configuration. Ex.: '{    \"access_identifier\": \"arn:aws:iam::1234:role/access_for_discovery\",    \"region\": \"us-east-1\",    \"enclave\": \"commercial/gov\"  }'.",
	},
	"deleted_at": schema.StringAttribute{
		Optional:            true,
		Computed:            true,
		CustomType:          timetypes.RFC3339Type{},
		MarkdownDescription: "Timestamp when the object has been deleted.",
	},
	"id": schema.StringAttribute{
		Computed: true,
		PlanModifiers: []planmodifier.String{
			stringplanmodifier.UseStateForUnknown(),
		},
		MarkdownDescription: "Auto-generated unique source config ID. Format BloxID.",
	},
	"restricted_to_accounts": schema.ListAttribute{
		ElementType: types.StringType,
		Optional:    true,
		Computed:    true,
		PlanModifiers: []planmodifier.List{
			listplanmodifier.RequiresReplaceIfConfigured(),
		},
		Validators: []validator.List{
			customvalidator.ListNotEmpty(),
		},
		MarkdownDescription: "Provider account IDs such as accountID/ SubscriptionID to be restricted for a given source_config.",
	},
}

// ExpandSourceConfig converts a Terraform Object to SDK type
func ExpandSourceConfig(ctx context.Context, o types.Object, diags *diag.Diagnostics) *uddiclouddiscovery.SourceConfig {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m SourceConfigModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

// Expand converts the Terraform model to SDK type
func (m *SourceConfigModel) Expand(ctx context.Context, diags *diag.Diagnostics) *uddiclouddiscovery.SourceConfig {
	if m == nil {
		return nil
	}
	to := &uddiclouddiscovery.SourceConfig{
		AccountScheduleId:    flex.ExpandStringPointer(m.AccountScheduleId),
		Accounts:             flex.ExpandFrameworkListNestedBlock(ctx, m.Accounts, diags, ExpandAccount),
		CloudCredentialId:    flex.ExpandStringPointer(m.CloudCredentialId),
		CredentialConfig:     ExpandCredentialConfig(ctx, m.CredentialConfig, diags),
		DeletedAt:            flex.ExpandRFC3339(m.DeletedAt, diags),
		Id:                   flex.ExpandStringPointer(m.Id),
		RestrictedToAccounts: flex.ExpandFrameworkListString(ctx, m.RestrictedToAccounts, diags),
	}
	return to
}

// FlattenSourceConfig converts an SDK type to Terraform Object
func FlattenSourceConfig(ctx context.Context, from *uddiclouddiscovery.SourceConfig, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(SourceConfigAttrTypes)
	}
	m := &SourceConfigModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, SourceConfigAttrTypes, m)
	diags.Append(d...)
	return t
}

// Flatten populates the Terraform model from SDK type
func (m *SourceConfigModel) Flatten(ctx context.Context, from *uddiclouddiscovery.SourceConfig, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	m.AccountScheduleId = flex.FlattenStringPointer(from.AccountScheduleId)
	m.Accounts = flex.FlattenFrameworkListNestedBlock(ctx, from.Accounts, AccountAttrTypes, diags, FlattenAccount)
	m.CloudCredentialId = flex.FlattenStringPointer(from.CloudCredentialId)
	m.CredentialConfig = FlattenCredentialConfig(ctx, from.CredentialConfig, diags)
	m.DeletedAt = flex.FlattenRFC3339(from.DeletedAt)
	m.Id = flex.FlattenStringPointer(from.Id)
	m.RestrictedToAccounts = flex.FlattenFrameworkListString(ctx, from.RestrictedToAccounts, diags)
}
