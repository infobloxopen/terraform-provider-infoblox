package clouddiscovery

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
	uddiclouddiscovery "github.com/infobloxopen/universal-ddi-go-client/clouddiscovery"
)

// AccountModel is the Terraform model for Account
type AccountModel struct {
	CompositeStatus        types.String      `tfsdk:"composite_status"`
	CompositeStatusMessage types.String      `tfsdk:"composite_status_message"`
	CreatedAt              timetypes.RFC3339 `tfsdk:"created_at"`
	DeletedAt              timetypes.RFC3339 `tfsdk:"deleted_at"`
	DhcpServerId           types.String      `tfsdk:"dhcp_server_id"`
	DnsServerId            types.String      `tfsdk:"dns_server_id"`
	Id                     types.String      `tfsdk:"id"`
	LastSuccessfulSync     timetypes.RFC3339 `tfsdk:"last_successful_sync"`
	LastSync               timetypes.RFC3339 `tfsdk:"last_sync"`
	Name                   types.String      `tfsdk:"name"`
	ParentId               types.String      `tfsdk:"parent_id"`
	PercentComplete        types.Int32       `tfsdk:"percent_complete"`
	ProviderAccountId      types.String      `tfsdk:"provider_account_id"`
	ScheduleId             types.String      `tfsdk:"schedule_id"`
	State                  types.String      `tfsdk:"state"`
	Status                 types.String      `tfsdk:"status"`
	StatusMessage          types.String      `tfsdk:"status_message"`
	UpdatedAt              timetypes.RFC3339 `tfsdk:"updated_at"`
}

// AccountAttrTypes contains the attribute types for AccountModel
var AccountAttrTypes = map[string]attr.Type{
	"composite_status":         types.StringType,
	"composite_status_message": types.StringType,
	"created_at":               timetypes.RFC3339Type{},
	"deleted_at":               timetypes.RFC3339Type{},
	"dhcp_server_id":           types.StringType,
	"dns_server_id":            types.StringType,
	"id":                       types.StringType,
	"last_successful_sync":     timetypes.RFC3339Type{},
	"last_sync":                timetypes.RFC3339Type{},
	"name":                     types.StringType,
	"parent_id":                types.StringType,
	"percent_complete":         types.Int32Type,
	"provider_account_id":      types.StringType,
	"schedule_id":              types.StringType,
	"state":                    types.StringType,
	"status":                   types.StringType,
	"status_message":           types.StringType,
	"updated_at":               timetypes.RFC3339Type{},
}

// AccountResourceSchemaAttributes contains the schema attributes for AccountModel
var AccountResourceSchemaAttributes = map[string]schema.Attribute{
	"composite_status": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "",
	},
	"composite_status_message": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "Status message of the sync operation.",
	},
	"created_at": schema.StringAttribute{
		Optional:            true,
		Computed:            true,
		CustomType:          timetypes.RFC3339Type{},
		MarkdownDescription: "Timestamp when the object has been created.",
	},
	"deleted_at": schema.StringAttribute{
		Optional:            true,
		Computed:            true,
		CustomType:          timetypes.RFC3339Type{},
		MarkdownDescription: "Timestamp when the object has been deleted.",
	},
	"dhcp_server_id": schema.StringAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "",
	},
	"dns_server_id": schema.StringAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "DNS Server ID.",
	},
	"id": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "Auto-generated unique source account ID. Format BloxID.",
	},
	"last_successful_sync": schema.StringAttribute{
		Optional:            true,
		Computed:            true,
		CustomType:          timetypes.RFC3339Type{},
		MarkdownDescription: "Last successful sync timestamp.",
	},
	"last_sync": schema.StringAttribute{
		Optional:            true,
		Computed:            true,
		CustomType:          timetypes.RFC3339Type{},
		MarkdownDescription: "Last sync timestamp.",
	},
	"name": schema.StringAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "Name of the source account.",
	},
	"parent_id": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "Parent ID.",
	},
	"percent_complete": schema.Int32Attribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "Sync progress as a percentage.",
	},
	"provider_account_id": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "",
	},
	"schedule_id": schema.StringAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "Schedule ID.",
	},
	"state": schema.StringAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "",
	},
	"status": schema.StringAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "Status of the sync operation.",
	},
	"status_message": schema.StringAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "Status message of the sync operation.",
	},
	"updated_at": schema.StringAttribute{
		Optional:            true,
		Computed:            true,
		CustomType:          timetypes.RFC3339Type{},
		MarkdownDescription: "Timestamp when the object has been updated.",
	},
}

// ExpandAccount converts a Terraform Object to SDK type
func ExpandAccount(ctx context.Context, o types.Object, diags *diag.Diagnostics) *uddiclouddiscovery.Account {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m AccountModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

// Expand converts the Terraform model to SDK type
func (m *AccountModel) Expand(ctx context.Context, diags *diag.Diagnostics) *uddiclouddiscovery.Account {
	if m == nil {
		return nil
	}
	to := &uddiclouddiscovery.Account{
		CompositeStatus:        flex.ExpandStringPointer(m.CompositeStatus),
		CompositeStatusMessage: flex.ExpandStringPointer(m.CompositeStatusMessage),
		CreatedAt:              flex.ExpandRFC3339(m.CreatedAt, diags),
		DeletedAt:              flex.ExpandRFC3339(m.DeletedAt, diags),
		DhcpServerId:           flex.ExpandStringPointer(m.DhcpServerId),
		DnsServerId:            flex.ExpandStringPointer(m.DnsServerId),
		Id:                     flex.ExpandStringPointer(m.Id),
		LastSuccessfulSync:     flex.ExpandRFC3339(m.LastSuccessfulSync, diags),
		LastSync:               flex.ExpandRFC3339(m.LastSync, diags),
		Name:                   flex.ExpandString(m.Name),
		ParentId:               flex.ExpandStringPointer(m.ParentId),
		PercentComplete:        flex.ExpandInt32Pointer(m.PercentComplete),
		ProviderAccountId:      flex.ExpandStringPointer(m.ProviderAccountId),
		ScheduleId:             flex.ExpandStringPointer(m.ScheduleId),
		State:                  flex.ExpandStringPointer(m.State),
		Status:                 flex.ExpandStringPointer(m.Status),
		StatusMessage:          flex.ExpandStringPointer(m.StatusMessage),
		UpdatedAt:              flex.ExpandRFC3339(m.UpdatedAt, diags),
	}
	return to
}

// FlattenAccount converts an SDK type to Terraform Object
func FlattenAccount(ctx context.Context, from *uddiclouddiscovery.Account, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(AccountAttrTypes)
	}
	m := &AccountModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, AccountAttrTypes, m)
	diags.Append(d...)
	return t
}

// Flatten populates the Terraform model from SDK type
func (m *AccountModel) Flatten(ctx context.Context, from *uddiclouddiscovery.Account, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	m.CompositeStatus = flex.FlattenStringPointer(from.CompositeStatus)
	m.CompositeStatusMessage = flex.FlattenStringPointer(from.CompositeStatusMessage)
	m.CreatedAt = flex.FlattenRFC3339(from.CreatedAt)
	m.DeletedAt = flex.FlattenRFC3339(from.DeletedAt)
	m.DhcpServerId = flex.FlattenStringPointer(from.DhcpServerId)
	m.DnsServerId = flex.FlattenStringPointer(from.DnsServerId)
	m.Id = flex.FlattenStringPointer(from.Id)
	m.LastSuccessfulSync = flex.FlattenRFC3339(from.LastSuccessfulSync)
	m.LastSync = flex.FlattenRFC3339(from.LastSync)
	m.Name = flex.FlattenString(from.Name)
	m.ParentId = flex.FlattenStringPointer(from.ParentId)
	m.PercentComplete = flex.FlattenInt32Pointer(from.PercentComplete)
	m.ProviderAccountId = flex.FlattenStringPointer(from.ProviderAccountId)
	m.ScheduleId = flex.FlattenStringPointer(from.ScheduleId)
	m.State = flex.FlattenStringPointer(from.State)
	m.Status = flex.FlattenStringPointer(from.Status)
	m.StatusMessage = flex.FlattenStringPointer(from.StatusMessage)
	m.UpdatedAt = flex.FlattenRFC3339(from.UpdatedAt)
}
