package notification

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/mapvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/mapdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"

	coremodel "github.com/infobloxopen/terraform-provider-infoblox/internal/core/model/notification"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
	importmod "github.com/infobloxopen/terraform-provider-infoblox/internal/planmodifiers/import"
	customvalidator "github.com/infobloxopen/terraform-provider-infoblox/internal/validator"
)

type NotificationRestEndpointModel struct {
	Id   types.String `tfsdk:"id"`
	NIOS types.Object `tfsdk:"nios"`
}

var NotificationRestEndpointAttrTypes = map[string]attr.Type{
	"id":   types.StringType,
	"nios": types.ObjectType{AttrTypes: NIOSNotificationRestEndpointAttrTypes},
}

type NIOSNotificationRestEndpointModel struct {
	ClientCertificateToken types.String `tfsdk:"client_certificate_token"`
	Comment                types.String `tfsdk:"comment"`
	ExtAttrs               types.Map    `tfsdk:"ext_attrs"`
	ExtAttrsAll            types.Map    `tfsdk:"ext_attrs_all"`
	LogLevel               types.String `tfsdk:"log_level"`
	Name                   types.String `tfsdk:"name"`
	OutboundMemberType     types.String `tfsdk:"outbound_member_type"`
	OutboundMembers        types.List   `tfsdk:"outbound_members"`
	Password               types.String `tfsdk:"password"`
	ServerCertValidation   types.String `tfsdk:"server_cert_validation"`
	SyncDisabled           types.Bool   `tfsdk:"sync_disabled"`
	TemplateInstance       types.Object `tfsdk:"template_instance"`
	Timeout                types.Int64  `tfsdk:"timeout"`
	Uri                    types.String `tfsdk:"uri"`
	Username               types.String `tfsdk:"username"`
	VendorIdentifier       types.String `tfsdk:"vendor_identifier"`
	WapiUserName           types.String `tfsdk:"wapi_user_name"`
	WapiUserPassword       types.String `tfsdk:"wapi_user_password"`
	ClientCertificateFile  types.String `tfsdk:"client_certificate_file"`
}

var NIOSNotificationRestEndpointAttrTypes = map[string]attr.Type{
	"client_certificate_token": types.StringType,
	"comment":                  types.StringType,
	"ext_attrs":                types.MapType{ElemType: types.StringType},
	"ext_attrs_all":            types.MapType{ElemType: types.StringType},
	"log_level":                types.StringType,
	"name":                     types.StringType,
	"outbound_member_type":     types.StringType,
	"outbound_members":         types.ListType{ElemType: types.StringType},
	"password":                 types.StringType,
	"server_cert_validation":   types.StringType,
	"sync_disabled":            types.BoolType,
	"template_instance":        types.ObjectType{AttrTypes: NotificationRestEndpointTemplateInstanceAttrTypes},
	"timeout":                  types.Int64Type,
	"uri":                      types.StringType,
	"username":                 types.StringType,
	"vendor_identifier":        types.StringType,
	"wapi_user_name":           types.StringType,
	"wapi_user_password":       types.StringType,
	"client_certificate_file":  types.StringType,
}

const (
	NotificationRestEndpointReturnFields = "client_certificate_subject,client_certificate_valid_from,client_certificate_valid_to,comment,extattrs,log_level,name,outbound_member_type,outbound_members,server_cert_validation,sync_disabled,template_instance,timeout,uri,username,vendor_identifier,wapi_user_name"
)

var NotificationRestEndpointResourceSchemaAttributes = map[string]schema.Attribute{
	"id": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The reference to the object.",
	},
	"nios": schema.SingleNestedAttribute{
		Optional:            true,
		MarkdownDescription: "NIOS backend-specific fields.",
		Attributes:          NotificationRestEndpointResourceNiosSchemaAttributes,
	},
}

var NotificationRestEndpointResourceNiosSchemaAttributes = map[string]schema.Attribute{
	"client_certificate_token": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
		},
		MarkdownDescription: "The token returned by the uploadinit function call in object fileop for a notification REST endpoit client certificate.",
	},
	"comment": schema.StringAttribute{
		Optional: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
			stringvalidator.LengthBetween(0, 256),
		},
		MarkdownDescription: "The comment of a notification REST endpoint.",
	},
	"ext_attrs": schema.MapAttribute{
		Optional:    true,
		Computed:    true,
		ElementType: types.StringType,
		Default:     mapdefault.StaticValue(types.MapNull(types.StringType)),
		Validators: []validator.Map{
			mapvalidator.SizeAtLeast(1),
		},
		MarkdownDescription: "Extensible attributes associated with the object. For valid values for extensible attributes, see {extattrs:values}.",
	},
	"ext_attrs_all": schema.MapAttribute{
		Computed:            true,
		ElementType:         types.StringType,
		MarkdownDescription: "All ext_attrs including Terraform Internal ID and inherited attributes.",
		PlanModifiers: []planmodifier.Map{
			importmod.AssociateInternalId(),
		},
	},
	"log_level": schema.StringAttribute{
		Default: stringdefault.StaticString("WARNING"),
		Validators: []validator.String{
			stringvalidator.OneOf("ERROR", "WARNING", "INFO", "DEBUG"),
		},
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The log level for a notification REST endpoint.",
	},
	"name": schema.StringAttribute{
		Required: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
			customvalidator.ValidateTrimmedString(),
		},
		MarkdownDescription: "The name of a notification REST endpoint.",
	},
	"outbound_member_type": schema.StringAttribute{
		Validators: []validator.String{
			stringvalidator.OneOf("MEMBER", "GM"),
		},
		Required:            true,
		MarkdownDescription: "The outbound member which will generate an event.",
	},
	"outbound_members": schema.ListAttribute{
		ElementType: types.StringType,
		Optional:    true,
		Validators: []validator.List{
			customvalidator.ListNotEmpty(),
			listvalidator.SizeAtMost(1),
		},
		MarkdownDescription: "The list of members for outbound events.",
	},
	"password": schema.StringAttribute{
		Sensitive: true,
		Optional:  true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
			stringvalidator.AlsoRequires(path.MatchRelative().AtParent().AtName("username")),
		},
		MarkdownDescription: "The password of the user that can log into a notification REST endpoint.",
	},
	"server_cert_validation": schema.StringAttribute{
		Default: stringdefault.StaticString("CA_CERT"),
		Validators: []validator.String{
			stringvalidator.OneOf("NO_VALIDATION", "CA_CERT", "CA_CERT_NO_HOSTNAME"),
		},
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The server certificate validation type.",
	},
	"sync_disabled": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Determines if the sync process is disabled for a notification REST endpoint.",
	},
	"template_instance": schema.SingleNestedAttribute{
		Attributes:          NotificationRestEndpointTemplateInstanceResourceSchemaAttributes,
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "",
	},
	"timeout": schema.Int64Attribute{
		Optional:            true,
		Computed:            true,
		Default:             int64default.StaticInt64(30),
		MarkdownDescription: "The timeout of session management (in seconds).",
	},
	"uri": schema.StringAttribute{
		Required: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
			customvalidator.ValidateTrimmedString(),
		},
		MarkdownDescription: "The URI of a notification REST endpoint.",
	},
	"username": schema.StringAttribute{
		Optional: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
			customvalidator.ValidateTrimmedString(),
			stringvalidator.AlsoRequires(path.MatchRelative().AtParent().AtName("password")),
		},
		MarkdownDescription: "The username of the user that can log into a notification REST endpoint.",
	},
	"vendor_identifier": schema.StringAttribute{
		Optional: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
		},
		MarkdownDescription: "The vendor identifier.",
	},
	"wapi_user_name": schema.StringAttribute{
		Optional: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
			stringvalidator.AlsoRequires(path.MatchRelative().AtParent().AtName("wapi_user_password")),
		},
		MarkdownDescription: "The user name for WAPI integration.",
	},
	"wapi_user_password": schema.StringAttribute{
		Sensitive: true,
		Optional:  true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
			stringvalidator.AlsoRequires(path.MatchRelative().AtParent().AtName("wapi_user_name")),
		},
		MarkdownDescription: "The user password for WAPI integration.",
	},
	"client_certificate_file": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "Local path to a PEM certificate file to upload to the NIOS grid. When set, the file is uploaded on create/update and the resulting token is stored in `client_certificate_token`.",
	},
}

// Expand converts the TF model to the infoblox core model
func (m *NotificationRestEndpointModel) Expand(ctx context.Context, diags *diag.Diagnostics, isCreate bool) *coremodel.NotificationRestEndpoint {
	if m == nil {
		return nil
	}

	obj := &coremodel.NotificationRestEndpoint{}

	// Expand NIOS nested attribute (returns nil if not present)
	niosModel := flex.ExpandNestedObject[NIOSNotificationRestEndpointModel](ctx, m.NIOS, diags)
	if niosModel != nil {
		obj.NIOS = niosModel.Expand(ctx, diags)
	}

	return obj
}

// Expand converts the NIOS TF model to the core model.
func (m *NIOSNotificationRestEndpointModel) Expand(ctx context.Context, diags *diag.Diagnostics) *coremodel.NIOSNotificationRestEndpointExt {
	return &coremodel.NIOSNotificationRestEndpointExt{
		ClientCertificateToken: flex.ExpandStringPointerNullAsEmpty(m.ClientCertificateToken),
		Comment:                flex.ExpandStringPointerNullAsEmpty(m.Comment),
		ExtAttrs:               flex.ExpandMapStringAny(ctx, m.ExtAttrs, diags),
		LogLevel:               flex.ExpandStringPointerNullAsEmpty(m.LogLevel),
		Name:                   flex.ExpandStringPointerNullAsEmpty(m.Name),
		OutboundMemberType:     flex.ExpandStringPointerNullAsEmpty(m.OutboundMemberType),
		OutboundMembers:        flex.ExpandFrameworkListString(ctx, m.OutboundMembers, diags),
		Password:               flex.ExpandStringPointerNullAsEmpty(m.Password),
		ServerCertValidation:   flex.ExpandStringPointerNullAsEmpty(m.ServerCertValidation),
		SyncDisabled:           flex.ExpandBoolPointer(m.SyncDisabled),
		TemplateInstance:       ExpandNotificationRestEndpointTemplateInstance(ctx, m.TemplateInstance, diags),
		Timeout:                flex.ExpandInt64Pointer(m.Timeout),
		Uri:                    flex.ExpandStringPointerNullAsEmpty(m.Uri),
		Username:               flex.ExpandStringPointerNullAsEmpty(m.Username),
		VendorIdentifier:       flex.ExpandStringPointerNullAsEmpty(m.VendorIdentifier),
		WapiUserName:           flex.ExpandStringPointerNullAsEmpty(m.WapiUserName),
		WapiUserPassword:       flex.ExpandStringPointerNullAsEmpty(m.WapiUserPassword),
	}
}

// Flatten populates the TF model from a core response.
func (m *NotificationRestEndpointModel) Flatten(ctx context.Context, resp *coremodel.NotificationRestEndpoint, diags *diag.Diagnostics) {
	if resp == nil {
		return
	}

	m.Id = flex.FlattenStringPointer(resp.Id)

	// Extract existing NIOS model, flatten API response onto it, convert back
	niosModel := flex.ExpandNestedObject[NIOSNotificationRestEndpointModel](ctx, m.NIOS, diags)
	if niosModel == nil {
		niosModel = &NIOSNotificationRestEndpointModel{}
	}
	plannedNIOS := flex.ExpandNestedObject[NIOSNotificationRestEndpointModel](ctx, m.NIOS, diags)
	niosModel.Flatten(ctx, resp.NIOS, diags)
	if resp.NIOS != nil {
		PostFlattenNotificationRestEndpointNIOS(ctx, plannedNIOS, niosModel, diags)
		m.NIOS = flex.FlattenNestedObject(ctx, niosModel, NIOSNotificationRestEndpointAttrTypes, diags)
	} else {
		m.NIOS = types.ObjectNull(NIOSNotificationRestEndpointAttrTypes)
	}

}

// Flatten merges API response onto existing NIOS model.
func (m *NIOSNotificationRestEndpointModel) Flatten(ctx context.Context, from *coremodel.NIOSNotificationRestEndpointExt, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	planExtAttrs := m.ExtAttrs
	if planExtAttrs.IsUnknown() {
		planExtAttrs = types.MapNull(types.StringType)
	}
	m.ClientCertificateToken = flex.FlattenStringPointerEmptyAsNull(from.ClientCertificateToken)
	m.Comment = flex.FlattenStringPointerEmptyAsNull(from.Comment)
	m.ExtAttrs, m.ExtAttrsAll = flex.FlattenEAs(planExtAttrs, from.ExtAttrs)
	m.LogLevel = flex.FlattenStringPointerEmptyAsNull(from.LogLevel)
	m.Name = flex.FlattenStringPointerEmptyAsNull(from.Name)
	m.OutboundMemberType = flex.FlattenStringPointerEmptyAsNull(from.OutboundMemberType)
	m.OutboundMembers = flex.FlattenFrameworkListString(ctx, from.OutboundMembers, diags)
	m.Password = flex.FlattenStringPointerEmptyAsNull(from.Password)
	m.ServerCertValidation = flex.FlattenStringPointerEmptyAsNull(from.ServerCertValidation)
	m.SyncDisabled = flex.FlattenBoolPointer(from.SyncDisabled)
	m.TemplateInstance = FlattenNotificationRestEndpointTemplateInstance(ctx, from.TemplateInstance, diags)
	m.Timeout = flex.FlattenInt64Pointer(from.Timeout)
	m.Uri = flex.FlattenStringPointerEmptyAsNull(from.Uri)
	m.Username = flex.FlattenStringPointerEmptyAsNull(from.Username)
	m.VendorIdentifier = flex.FlattenStringPointerEmptyAsNull(from.VendorIdentifier)
	m.WapiUserName = flex.FlattenStringPointerEmptyAsNull(from.WapiUserName)
	m.WapiUserPassword = flex.FlattenStringPointerEmptyAsNull(from.WapiUserPassword)
}
