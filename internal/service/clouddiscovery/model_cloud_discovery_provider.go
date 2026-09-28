package clouddiscovery

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/mapvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/mapdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	stringplanmodifier "github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"

	coremodel "github.com/infobloxopen/terraform-provider-infoblox/internal/core/model/clouddiscovery"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
	customvalidator "github.com/infobloxopen/terraform-provider-infoblox/internal/validator"
)

type CloudDiscoveryProviderModel struct {
	Id            types.String `tfsdk:"id"`
	UpdateTrigger types.String `tfsdk:"update_trigger"`
	UDDI          types.Object `tfsdk:"uddi"`
}

var CloudDiscoveryProviderAttrTypes = map[string]attr.Type{
	"id":             types.StringType,
	"update_trigger": types.StringType,
	"uddi":           types.ObjectType{AttrTypes: UDDICloudDiscoveryProviderAttrTypes},
}

type UDDICloudDiscoveryProviderModel struct {
	AccountPreference       types.String `tfsdk:"account_preference"`
	AdditionalConfig        types.Object `tfsdk:"additional_config"`
	CredentialPreference    types.Object `tfsdk:"credential_preference"`
	Description             types.String `tfsdk:"description"`
	DesiredState            types.String `tfsdk:"desired_state"`
	DestinationTypesEnabled types.List   `tfsdk:"destination_types_enabled"`
	Destinations            types.List   `tfsdk:"destinations"`
	IsDisabled              types.Bool   `tfsdk:"is_disabled"`
	LabsProvider            types.Bool   `tfsdk:"labs_provider"`
	Name                    types.String `tfsdk:"name"`
	ProviderType            types.String `tfsdk:"provider_type"`
	SourceConfigs           types.List   `tfsdk:"source_configs"`
	SyncInterval            types.String `tfsdk:"sync_interval"`
	Tags                    types.Map    `tfsdk:"tags"`
	TagsAll                 types.Map    `tfsdk:"tags_all"`
}

var UDDICloudDiscoveryProviderAttrTypes = map[string]attr.Type{
	"account_preference":        types.StringType,
	"additional_config":         types.ObjectType{AttrTypes: AdditionalConfigAttrTypes},
	"credential_preference":     types.ObjectType{AttrTypes: CredentialPreferenceAttrTypes},
	"description":               types.StringType,
	"desired_state":             types.StringType,
	"destination_types_enabled": types.ListType{ElemType: types.StringType},
	"destinations":              types.ListType{ElemType: types.ObjectType{AttrTypes: DestinationAttrTypes}},
	"is_disabled":               types.BoolType,
	"labs_provider":             types.BoolType,
	"name":                      types.StringType,
	"provider_type":             types.StringType,
	"source_configs":            types.ListType{ElemType: types.ObjectType{AttrTypes: SourceConfigAttrTypes}},
	"sync_interval":             types.StringType,
	"tags":                      types.MapType{ElemType: types.StringType},
	"tags_all":                  types.MapType{ElemType: types.StringType},
}

const (
	CloudDiscoveryProviderReturnFields = ""
)

var CloudDiscoveryProviderResourceSchemaAttributes = map[string]schema.Attribute{
	"id": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "Auto-generated unique discovery config ID. Format BloxID.",
	},
	"update_trigger": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "An arbitrary value used to trigger an update. Not sent to the API. Change it when Terraform reports no infrastructure changes.",
	},
	"uddi": schema.SingleNestedAttribute{
		Optional:            true,
		MarkdownDescription: "UDDI backend-specific fields.",
		Attributes:          CloudDiscoveryProviderResourceUddiSchemaAttributes,
	},
}

var CloudDiscoveryProviderResourceUddiSchemaAttributes = map[string]schema.Attribute{
	"account_preference": schema.StringAttribute{
		Required:            true,
		MarkdownDescription: "Account preference. For ex.: single, multiple, auto-discover-multiple.",
	},
	"additional_config": schema.SingleNestedAttribute{
		Attributes:          AdditionalConfigResourceSchemaAttributes,
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "Additional configuration. Ex.: '{    \"excluded_object_types\": [],    \"exclusion_account_list\": [],    \"zone_forwarding\": \"true\" or \"false\" }'.",
	},
	"credential_preference": schema.SingleNestedAttribute{
		Attributes:          CredentialPreferenceResourceSchemaAttributes,
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "Credential preference. Ex.: '{    \"type\": \"static\" or \"delegated\",    \"access_identifier_type\": \"role_arn\" or \"tenant_id\" or \"project_id\"  }'.",
	},
	"description": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "Description of the discovery config. Optional.",
	},
	"desired_state": schema.StringAttribute{
		Default:             stringdefault.StaticString("enabled"),
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "Desired state. Default is \"enabled\".",
	},
	"destination_types_enabled": schema.ListAttribute{
		ElementType: types.StringType,
		Optional:    true,
		Computed:    true,
		Validators: []validator.List{
			customvalidator.ListNotEmpty(),
		},
		MarkdownDescription: "Destinations types enabled: Ex.: DNS, IPAM and ACCOUNT.",
	},
	"destinations": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: DestinationResourceSchemaAttributes,
		},
		Optional: true,
		Computed: true,
		Validators: []validator.List{
			customvalidator.ListNotEmpty(),
		},
		MarkdownDescription: "Destinations.",
	},
	"is_disabled": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "is_disabled. Enables/Disables provider. Newer version of desired_state.",
	},
	"labs_provider": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "labs_provider. Indicates if a provider is enabled through Infoblox Labs.",
	},
	"name": schema.StringAttribute{
		Required: true,
		PlanModifiers: []planmodifier.String{
			stringplanmodifier.RequiresReplaceIfConfigured(),
		},
		MarkdownDescription: "Name of the discovery config.",
	},
	"provider_type": schema.StringAttribute{
		Validators: []validator.String{
			stringvalidator.OneOf("Amazon Web Services", "Google Cloud Platform", "Microsoft Azure"),
		},
		Required: true,
		PlanModifiers: []planmodifier.String{
			stringplanmodifier.RequiresReplaceIfConfigured(),
		},
		MarkdownDescription: "Provider type. Ex.: Amazon Web Services, Google Cloud Platform, Microsoft Azure.",
	},
	"source_configs": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: SourceConfigResourceSchemaAttributes,
		},
		Optional: true,
		Computed: true,
		Validators: []validator.List{
			customvalidator.ListNotEmpty(),
		},
		MarkdownDescription: "Source configs.",
	},
	"sync_interval": schema.StringAttribute{
		Default:             stringdefault.StaticString("Auto"),
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "",
	},
	"tags": schema.MapAttribute{
		Optional:    true,
		Computed:    true,
		ElementType: types.StringType,
		Default:     mapdefault.StaticValue(types.MapNull(types.StringType)),
		Validators: []validator.Map{
			mapvalidator.SizeAtLeast(1),
		},
		MarkdownDescription: "Tagging specifics.",
	},
	"tags_all": schema.MapAttribute{
		Computed:            true,
		ElementType:         types.StringType,
		MarkdownDescription: "All tags including inherited values.",
	},
}

// Expand converts the TF model to the infoblox core model
func (m *CloudDiscoveryProviderModel) Expand(ctx context.Context, diags *diag.Diagnostics, isCreate bool) *coremodel.CloudDiscoveryProvider {
	if m == nil {
		return nil
	}

	obj := &coremodel.CloudDiscoveryProvider{}

	// Expand UDDI nested attribute (returns nil if not present)
	uddiModel := flex.ExpandNestedObject[UDDICloudDiscoveryProviderModel](ctx, m.UDDI, diags)
	if uddiModel != nil {
		obj.UDDI = uddiModel.Expand(ctx, diags)
	}

	return obj
}

// Expand converts the UDDI TF model to the core model.
func (m *UDDICloudDiscoveryProviderModel) Expand(ctx context.Context, diags *diag.Diagnostics) *coremodel.UDDICloudDiscoveryProviderExt {
	return &coremodel.UDDICloudDiscoveryProviderExt{
		AccountPreference:       flex.ExpandString(m.AccountPreference),
		AdditionalConfig:        ExpandAdditionalConfig(ctx, m.AdditionalConfig, diags),
		CredentialPreference:    ExpandCredentialPreference(ctx, m.CredentialPreference, diags),
		Description:             flex.ExpandStringPointer(m.Description),
		DesiredState:            flex.ExpandStringPointer(m.DesiredState),
		DestinationTypesEnabled: flex.ExpandFrameworkListString(ctx, m.DestinationTypesEnabled, diags),
		Destinations:            flex.ExpandFrameworkListNestedBlock(ctx, m.Destinations, diags, ExpandDestination),
		IsDisabled:              flex.ExpandBoolPointer(m.IsDisabled),
		LabsProvider:            flex.ExpandBoolPointer(m.LabsProvider),
		Name:                    flex.ExpandString(m.Name),
		ProviderType:            flex.ExpandString(m.ProviderType),
		SourceConfigs:           flex.ExpandFrameworkListNestedBlock(ctx, m.SourceConfigs, diags, ExpandSourceConfig),
		SyncInterval:            flex.ExpandStringPointer(m.SyncInterval),
		Tags:                    flex.ExpandMapStringAny(ctx, m.Tags, diags),
	}
}

// Flatten populates the TF model from a core response.
func (m *CloudDiscoveryProviderModel) Flatten(ctx context.Context, resp *coremodel.CloudDiscoveryProvider, diags *diag.Diagnostics) {
	if resp == nil {
		return
	}

	m.Id = flex.FlattenStringPointer(resp.Id)

	// Extract existing UDDI model, flatten API response onto it, convert back
	uddiModel := flex.ExpandNestedObject[UDDICloudDiscoveryProviderModel](ctx, m.UDDI, diags)
	if uddiModel == nil {
		uddiModel = &UDDICloudDiscoveryProviderModel{}
	}
	uddiModel.Flatten(ctx, resp.UDDI, diags)
	if resp.UDDI != nil {
		m.UDDI = flex.FlattenNestedObject(ctx, uddiModel, UDDICloudDiscoveryProviderAttrTypes, diags)
	} else {
		m.UDDI = types.ObjectNull(UDDICloudDiscoveryProviderAttrTypes)
	}
}

// Flatten merges API response onto existing UDDI model.
func (m *UDDICloudDiscoveryProviderModel) Flatten(ctx context.Context, from *coremodel.UDDICloudDiscoveryProviderExt, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	m.AccountPreference = flex.FlattenString(from.AccountPreference)
	m.AdditionalConfig = FlattenAdditionalConfig(ctx, from.AdditionalConfig, diags)
	m.CredentialPreference = FlattenCredentialPreference(ctx, from.CredentialPreference, diags)
	m.Description = flex.FlattenStringPointer(from.Description)
	m.DesiredState = flex.FlattenStringPointer(from.DesiredState)
	m.DestinationTypesEnabled = flex.FlattenFrameworkListString(ctx, from.DestinationTypesEnabled, diags)
	m.Destinations = flex.FlattenFrameworkListNestedBlock(ctx, from.Destinations, DestinationAttrTypes, diags, FlattenDestination)
	m.IsDisabled = flex.FlattenBoolPointer(from.IsDisabled)
	m.LabsProvider = flex.FlattenBoolPointer(from.LabsProvider)
	m.Name = flex.FlattenString(from.Name)
	m.ProviderType = flex.FlattenString(from.ProviderType)
	m.SourceConfigs = flex.FlattenFrameworkListNestedBlock(ctx, from.SourceConfigs, SourceConfigAttrTypes, diags, FlattenSourceConfig)
	m.SyncInterval = flex.FlattenStringPointer(from.SyncInterval)
	tagsAll := flex.FlattenMapStringAny(ctx, from.Tags, diags)
	if m.Tags.IsNull() || m.Tags.IsUnknown() {
		m.Tags = tagsAll
	}
	m.TagsAll = tagsAll
}
