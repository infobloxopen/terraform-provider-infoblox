package infra

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
	uddiinfra "github.com/infobloxopen/universal-ddi-go-client/inframgmt"
)

// ServiceHostConfigModel is the Terraform model for ServiceHostConfig
type ServiceHostConfigModel struct {
	CurrentVersion types.String      `tfsdk:"current_version"`
	ExtraData      types.String      `tfsdk:"extra_data"`
	HostId         types.String      `tfsdk:"host_id"`
	Id             types.String      `tfsdk:"id"`
	ServiceId      types.String      `tfsdk:"service_id"`
	ServiceType    types.String      `tfsdk:"service_type"`
	UpgradedAt     timetypes.RFC3339 `tfsdk:"upgraded_at"`
}

// ServiceHostConfigAttrTypes contains the attribute types for ServiceHostConfigModel
var ServiceHostConfigAttrTypes = map[string]attr.Type{
	"current_version": types.StringType,
	"extra_data":      types.StringType,
	"host_id":         types.StringType,
	"id":              types.StringType,
	"service_id":      types.StringType,
	"service_type":    types.StringType,
	"upgraded_at":     timetypes.RFC3339Type{},
}

// ServiceHostConfigResourceSchemaAttributes contains the schema attributes for ServiceHostConfigModel
var ServiceHostConfigResourceSchemaAttributes = map[string]schema.Attribute{
	"current_version": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The current version of the Service deployed on the Host.",
	},
	"extra_data": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The field to carry any extra data specific to this configuration.",
	},
	"host_id": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The resource identifier.",
	},
	"id": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The resource identifier.",
	},
	"service_id": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The resource identifier.",
	},
	"service_type": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The type of the Service deployed on the Host (`dns`, `cdc`, etc.).",
	},
	"upgraded_at": schema.StringAttribute{
		Optional:            true,
		CustomType:          timetypes.RFC3339Type{},
		MarkdownDescription: "The timestamp of the latest upgrade of the Host-specific Service configuration.",
	},
}

// ExpandServiceHostConfig converts a Terraform Object to SDK type
func ExpandServiceHostConfig(ctx context.Context, o types.Object, diags *diag.Diagnostics) *uddiinfra.ServiceHostConfig {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m ServiceHostConfigModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

// Expand converts the Terraform model to SDK type
func (m *ServiceHostConfigModel) Expand(ctx context.Context, diags *diag.Diagnostics) *uddiinfra.ServiceHostConfig {
	if m == nil {
		return nil
	}
	to := &uddiinfra.ServiceHostConfig{
		CurrentVersion: flex.ExpandStringPointer(m.CurrentVersion),
		ExtraData:      flex.ExpandStringPointer(m.ExtraData),
		HostId:         flex.ExpandStringPointer(m.HostId),
		Id:             flex.ExpandStringPointer(m.Id),
		ServiceId:      flex.ExpandStringPointer(m.ServiceId),
		ServiceType:    flex.ExpandStringPointer(m.ServiceType),
		UpgradedAt:     flex.ExpandRFC3339(m.UpgradedAt, diags),
	}
	return to
}

// FlattenServiceHostConfig converts an SDK type to Terraform Object
func FlattenServiceHostConfig(ctx context.Context, from *uddiinfra.ServiceHostConfig, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(ServiceHostConfigAttrTypes)
	}
	m := &ServiceHostConfigModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, ServiceHostConfigAttrTypes, m)
	diags.Append(d...)
	return t
}

// Flatten populates the Terraform model from SDK type
func (m *ServiceHostConfigModel) Flatten(ctx context.Context, from *uddiinfra.ServiceHostConfig, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	m.CurrentVersion = flex.FlattenStringPointer(from.CurrentVersion)
	m.ExtraData = flex.FlattenStringPointer(from.ExtraData)
	m.HostId = flex.FlattenStringPointer(from.HostId)
	m.Id = flex.FlattenStringPointer(from.Id)
	m.ServiceId = flex.FlattenStringPointer(from.ServiceId)
	m.ServiceType = flex.FlattenStringPointer(from.ServiceType)
	m.UpgradedAt = flex.FlattenRFC3339(from.UpgradedAt)
}
