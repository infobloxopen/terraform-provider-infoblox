package clouddiscovery

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	stringplanmodifier "github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
	customvalidator "github.com/infobloxopen/terraform-provider-infoblox/internal/validator"
	uddiclouddiscovery "github.com/infobloxopen/universal-ddi-go-client/clouddiscovery"
)

// DestinationModel is the Terraform model for Destination
type DestinationModel struct {
	Config          types.Object      `tfsdk:"config"`
	DeletedAt       timetypes.RFC3339 `tfsdk:"deleted_at"`
	DestinationType types.String      `tfsdk:"destination_type"`
	Id              types.String      `tfsdk:"id"`
}

// DestinationAttrTypes contains the attribute types for DestinationModel
var DestinationAttrTypes = map[string]attr.Type{
	"config":           types.ObjectType{AttrTypes: DestinationConfigAttrTypes},
	"deleted_at":       timetypes.RFC3339Type{},
	"destination_type": types.StringType,
	"id":               types.StringType,
}

// DestinationResourceSchemaAttributes contains the schema attributes for DestinationModel
var DestinationResourceSchemaAttributes = map[string]schema.Attribute{
	"config": schema.SingleNestedAttribute{
		Attributes:          DestinationConfigResourceSchemaAttributes,
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "Destination configuration. Ex.: '{  \"dns\": {    \"view_name\": \"view 1\",    \"view_id\": \"dns/view/v1\",    \"consolidated_zone_data_enabled\": false,    \"sync_type\": \"read_only/read_write\"    \"split_view_enabled\": false  },  \"ipam\": {    \"ip_space\": \"\",  },  \"account\": {},  }'.",
	},
	"deleted_at": schema.StringAttribute{
		Optional:            true,
		Computed:            true,
		CustomType:          timetypes.RFC3339Type{},
		MarkdownDescription: "Timestamp when the object has been deleted.",
	},
	"destination_type": schema.StringAttribute{
		Validators: []validator.String{
			customvalidator.StringNotNull(),
			stringvalidator.OneOf("DNS", "IPAM/DHCP", "ACCOUNTS"),
		},
		Optional: true,
		Computed: true,
		PlanModifiers: []planmodifier.String{
			stringplanmodifier.RequiresReplaceIfConfigured(),
		},
		MarkdownDescription: "Destination type: DNS / IPAM / ACCOUNT.",
	},
	"id": schema.StringAttribute{
		Computed: true,
		PlanModifiers: []planmodifier.String{
			stringplanmodifier.UseStateForUnknown(),
		},
		MarkdownDescription: "Auto-generated unique destination ID. Format BloxID.",
	},
}

// ExpandDestination converts a Terraform Object to SDK type
func ExpandDestination(ctx context.Context, o types.Object, diags *diag.Diagnostics) *uddiclouddiscovery.Destination {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m DestinationModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

// Expand converts the Terraform model to SDK type
func (m *DestinationModel) Expand(ctx context.Context, diags *diag.Diagnostics) *uddiclouddiscovery.Destination {
	if m == nil {
		return nil
	}
	to := &uddiclouddiscovery.Destination{
		Config:          ExpandDestinationConfig(ctx, m.Config, diags),
		DeletedAt:       flex.ExpandRFC3339(m.DeletedAt, diags),
		DestinationType: flex.ExpandString(m.DestinationType),
		Id:              flex.ExpandStringPointer(m.Id),
	}
	return to
}

// FlattenDestination converts an SDK type to Terraform Object
func FlattenDestination(ctx context.Context, from *uddiclouddiscovery.Destination, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(DestinationAttrTypes)
	}
	m := &DestinationModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, DestinationAttrTypes, m)
	diags.Append(d...)
	return t
}

// Flatten populates the Terraform model from SDK type
func (m *DestinationModel) Flatten(ctx context.Context, from *uddiclouddiscovery.Destination, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	m.Config = FlattenDestinationConfig(ctx, from.Config, diags)
	m.DeletedAt = flex.FlattenRFC3339(from.DeletedAt)
	m.DestinationType = flex.FlattenString(from.DestinationType)
	m.Id = flex.FlattenStringPointer(from.Id)
}
