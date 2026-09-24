package clouddiscovery

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	boolplanmodifier "github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	stringplanmodifier "github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
	customvalidator "github.com/infobloxopen/terraform-provider-infoblox/internal/validator"
	uddiclouddiscovery "github.com/infobloxopen/universal-ddi-go-client/clouddiscovery"
)

// DNSConfigModel is the Terraform model for DNSConfig
type DNSConfigModel struct {
	ConsolidatedZoneDataEnabled  types.Bool   `tfsdk:"consolidated_zone_data_enabled"`
	ResolverEndpointsSyncEnabled types.Bool   `tfsdk:"resolver_endpoints_sync_enabled"`
	SplitViewEnabled             types.Bool   `tfsdk:"split_view_enabled"`
	SyncType                     types.String `tfsdk:"sync_type"`
	ViewId                       types.String `tfsdk:"view_id"`
	ViewName                     types.String `tfsdk:"view_name"`
	ZoneFilters                  types.List   `tfsdk:"zone_filters"`
}

// DNSConfigAttrTypes contains the attribute types for DNSConfigModel
var DNSConfigAttrTypes = map[string]attr.Type{
	"consolidated_zone_data_enabled":  types.BoolType,
	"resolver_endpoints_sync_enabled": types.BoolType,
	"split_view_enabled":              types.BoolType,
	"sync_type":                       types.StringType,
	"view_id":                         types.StringType,
	"view_name":                       types.StringType,
	"zone_filters":                    types.ListType{ElemType: types.ObjectType{AttrTypes: ZoneFilterAttrTypes}},
}

// DNSConfigResourceSchemaAttributes contains the schema attributes for DNSConfigModel
var DNSConfigResourceSchemaAttributes = map[string]schema.Attribute{
	"consolidated_zone_data_enabled": schema.BoolAttribute{
		Optional: true,
		PlanModifiers: []planmodifier.Bool{
			boolplanmodifier.RequiresReplaceIfConfigured(),
		},
		MarkdownDescription: "",
	},
	"resolver_endpoints_sync_enabled": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "resolver_endpoints_sync_enabled enables discovery of inbound and outbound endpoints from third party providers.",
	},
	"split_view_enabled": schema.BoolAttribute{
		Optional: true,
		PlanModifiers: []planmodifier.Bool{
			boolplanmodifier.RequiresReplaceIfConfigured(),
		},
		MarkdownDescription: "split_view_enabled consolidates private zones into a single view, which is separate from the public zone view.",
	},
	"sync_type": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "",
	},
	"view_id": schema.StringAttribute{
		Optional: true,
		PlanModifiers: []planmodifier.String{
			stringplanmodifier.RequiresReplaceIfConfigured(),
		},
		MarkdownDescription: "",
	},
	"view_name": schema.StringAttribute{
		Optional: true,
		PlanModifiers: []planmodifier.String{
			stringplanmodifier.RequiresReplaceIfConfigured(),
		},
		MarkdownDescription: "",
	},
	"zone_filters": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: ZoneFilterResourceSchemaAttributes,
		},
		Optional: true,
		Validators: []validator.List{
			customvalidator.ListNotEmpty(),
		},
		MarkdownDescription: "",
	},
}

// ExpandDNSConfig converts a Terraform Object to SDK type
func ExpandDNSConfig(ctx context.Context, o types.Object, diags *diag.Diagnostics) *uddiclouddiscovery.DNSConfig {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m DNSConfigModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

// Expand converts the Terraform model to SDK type
func (m *DNSConfigModel) Expand(ctx context.Context, diags *diag.Diagnostics) *uddiclouddiscovery.DNSConfig {
	if m == nil {
		return nil
	}
	to := &uddiclouddiscovery.DNSConfig{
		ConsolidatedZoneDataEnabled:  flex.ExpandBoolPointer(m.ConsolidatedZoneDataEnabled),
		ResolverEndpointsSyncEnabled: flex.ExpandBoolPointer(m.ResolverEndpointsSyncEnabled),
		SplitViewEnabled:             flex.ExpandBoolPointer(m.SplitViewEnabled),
		SyncType:                     flex.ExpandStringPointer(m.SyncType),
		ViewId:                       flex.ExpandStringPointer(m.ViewId),
		ViewName:                     flex.ExpandStringPointer(m.ViewName),
		ZoneFilters:                  flex.ExpandFrameworkListNestedBlock(ctx, m.ZoneFilters, diags, ExpandZoneFilter),
	}
	return to
}

// FlattenDNSConfig converts an SDK type to Terraform Object
func FlattenDNSConfig(ctx context.Context, from *uddiclouddiscovery.DNSConfig, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(DNSConfigAttrTypes)
	}
	m := &DNSConfigModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, DNSConfigAttrTypes, m)
	diags.Append(d...)
	return t
}

// Flatten populates the Terraform model from SDK type
func (m *DNSConfigModel) Flatten(ctx context.Context, from *uddiclouddiscovery.DNSConfig, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	m.ConsolidatedZoneDataEnabled = flex.FlattenBoolPointer(from.ConsolidatedZoneDataEnabled)
	m.ResolverEndpointsSyncEnabled = flex.FlattenBoolPointer(from.ResolverEndpointsSyncEnabled)
	m.SplitViewEnabled = flex.FlattenBoolPointer(from.SplitViewEnabled)
	m.SyncType = flex.FlattenStringPointer(from.SyncType)
	m.ViewId = flex.FlattenStringPointer(from.ViewId)
	m.ViewName = flex.FlattenStringPointer(from.ViewName)
	m.ZoneFilters = flex.FlattenFrameworkListNestedBlock(ctx, from.ZoneFilters, ZoneFilterAttrTypes, diags, FlattenZoneFilter)
}
