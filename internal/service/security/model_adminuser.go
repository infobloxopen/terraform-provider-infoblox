package security

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
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/mapdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"

	coremodel "github.com/infobloxopen/terraform-provider-infoblox/internal/core/model/security"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
	importmod "github.com/infobloxopen/terraform-provider-infoblox/internal/planmodifiers/import"
	customvalidator "github.com/infobloxopen/terraform-provider-infoblox/internal/validator"
)

type AdminuserModel struct {
	Id            types.String `tfsdk:"id"`
	UpdateTrigger types.String `tfsdk:"update_trigger"`
	NIOS          types.Object `tfsdk:"nios"`
}

var AdminuserAttrTypes = map[string]attr.Type{
	"id":             types.StringType,
	"update_trigger": types.StringType,
	"nios":           types.ObjectType{AttrTypes: NIOSAdminuserAttrTypes},
}

type NIOSAdminuserModel struct {
	AdminGroups                     types.List   `tfsdk:"admin_groups"`
	AuthMethod                      types.String `tfsdk:"auth_method"`
	AuthType                        types.String `tfsdk:"auth_type"`
	CaCertificateIssuer             types.String `tfsdk:"ca_certificate_issuer"`
	ClientCertificateSerialNumber   types.String `tfsdk:"client_certificate_serial_number"`
	Comment                         types.String `tfsdk:"comment"`
	Disable                         types.Bool   `tfsdk:"disable"`
	Email                           types.String `tfsdk:"email"`
	EnableCertificateAuthentication types.Bool   `tfsdk:"enable_certificate_authentication"`
	ExtAttrs                        types.Map    `tfsdk:"ext_attrs"`
	ExtAttrsAll                     types.Map    `tfsdk:"ext_attrs_all"`
	Name                            types.String `tfsdk:"name"`
	Password                        types.String `tfsdk:"password"`
	SshKeys                         types.List   `tfsdk:"ssh_keys"`
	TimeZone                        types.String `tfsdk:"time_zone"`
}

var NIOSAdminuserAttrTypes = map[string]attr.Type{
	"admin_groups":                      types.ListType{ElemType: types.StringType},
	"auth_method":                       types.StringType,
	"auth_type":                         types.StringType,
	"ca_certificate_issuer":             types.StringType,
	"client_certificate_serial_number":  types.StringType,
	"comment":                           types.StringType,
	"disable":                           types.BoolType,
	"email":                             types.StringType,
	"enable_certificate_authentication": types.BoolType,
	"ext_attrs":                         types.MapType{ElemType: types.StringType},
	"ext_attrs_all":                     types.MapType{ElemType: types.StringType},
	"name":                              types.StringType,
	"password":                          types.StringType,
	"ssh_keys":                          types.ListType{ElemType: types.ObjectType{AttrTypes: AdminuserSshKeysAttrTypes}},
	"time_zone":                         types.StringType,
}

const (
	AdminuserReturnFields = "admin_groups,auth_method,auth_type,ca_certificate_issuer,client_certificate_serial_number,comment,disable,email,enable_certificate_authentication,extattrs,name,ssh_keys,status,time_zone,use_ssh_keys,use_time_zone"
)

var AdminuserResourceSchemaAttributes = map[string]schema.Attribute{
	"id": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The reference to the object.",
	},
	"update_trigger": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "An arbitrary value used to trigger an update. Not sent to the API. Change it when Terraform reports no infrastructure changes.",
	},
	"nios": schema.SingleNestedAttribute{
		Optional:            true,
		MarkdownDescription: "NIOS backend-specific fields.",
		Attributes:          AdminuserResourceNiosSchemaAttributes,
	},
}

var AdminuserResourceNiosSchemaAttributes = map[string]schema.Attribute{
	"admin_groups": schema.ListAttribute{
		ElementType: types.StringType,
		Required:    true,
		Validators: []validator.List{
			customvalidator.ListNotEmpty(),
			listvalidator.SizeBetween(1, 1),
		},
		MarkdownDescription: "The names of the Admin Groups to which this Admin User belongs. Currently, this is limited to only one Admin Group.",
	},
	"auth_method": schema.StringAttribute{
		Default: stringdefault.StaticString("KEYPAIR"),
		Validators: []validator.String{
			stringvalidator.OneOf("KEYPAIR", "KEYPAIR_PASSWORD"),
		},
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "Determines the way of authentication",
	},
	"auth_type": schema.StringAttribute{
		Default: stringdefault.StaticString("LOCAL"),
		Validators: []validator.String{
			stringvalidator.OneOf("LOCAL", "RADIUS", "REMOTE", "SAML", "SAML_LOCAL"),
		},
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The authentication type for the admin user.",
	},
	"ca_certificate_issuer": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
		},
		MarkdownDescription: "The CA certificate that is used for user lookup during authentication.",
	},
	"client_certificate_serial_number": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
		},
		MarkdownDescription: "The serial number of the client certificate.",
	},
	"comment": schema.StringAttribute{
		Optional: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
			stringvalidator.LengthBetween(0, 256),
			customvalidator.ValidateTrimmedString(),
		},
		MarkdownDescription: "Comment for the admin user; maximum 256 characters.",
	},
	"disable": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Determines whether the admin user is disabled or not. When this is set to False, the admin user is enabled.",
	},
	"email": schema.StringAttribute{
		Optional: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
		},
		MarkdownDescription: "The e-mail address for the admin user.",
	},
	"enable_certificate_authentication": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Determines whether the user is allowed to log in only with the certificate. Regular username/password authentication will be disabled for this user.",
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
	"name": schema.StringAttribute{
		Required: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
		},
		MarkdownDescription: "The name of the admin user.",
	},
	"password": schema.StringAttribute{
		Sensitive: true,
		Required:  true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
			customvalidator.IsStrongPassword(),
		},
		MarkdownDescription: "The password for the administrator to use when logging in.",
	},
	"ssh_keys": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: AdminuserSshKeysResourceSchemaAttributes,
		},
		Optional: true,
		Validators: []validator.List{
			customvalidator.ListNotEmpty(),
		},
		MarkdownDescription: "List of ssh keys for a particular user.",
	},
	"time_zone": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
		},
		MarkdownDescription: "The time zone for this admin user.",
	},
}

// Expand converts the TF model to the infoblox core model
func (m *AdminuserModel) Expand(ctx context.Context, diags *diag.Diagnostics, isCreate bool) *coremodel.Adminuser {
	if m == nil {
		return nil
	}

	obj := &coremodel.Adminuser{}

	// Expand NIOS nested attribute (returns nil if not present)
	niosModel := flex.ExpandNestedObject[NIOSAdminuserModel](ctx, m.NIOS, diags)
	if niosModel != nil {
		obj.NIOS = niosModel.Expand(ctx, diags)
	}

	return obj
}

// Expand converts the NIOS TF model to the core model.
func (m *NIOSAdminuserModel) Expand(ctx context.Context, diags *diag.Diagnostics) *coremodel.NIOSAdminuserExt {
	return &coremodel.NIOSAdminuserExt{
		AdminGroups:                     flex.ExpandFrameworkListString(ctx, m.AdminGroups, diags),
		AuthMethod:                      flex.ExpandStringPointerNullAsEmpty(m.AuthMethod),
		AuthType:                        flex.ExpandStringPointerNullAsEmpty(m.AuthType),
		CaCertificateIssuer:             flex.ExpandStringPointer(m.CaCertificateIssuer),
		ClientCertificateSerialNumber:   flex.ExpandStringPointer(m.ClientCertificateSerialNumber),
		Comment:                         flex.ExpandStringPointerNullAsEmpty(m.Comment),
		Disable:                         flex.ExpandBoolPointer(m.Disable),
		Email:                           flex.ExpandStringPointerNullAsEmpty(m.Email),
		EnableCertificateAuthentication: flex.ExpandBoolPointer(m.EnableCertificateAuthentication),
		ExtAttrs:                        flex.ExpandMapStringAny(ctx, m.ExtAttrs, diags),
		Name:                            flex.ExpandStringPointerNullAsEmpty(m.Name),
		Password:                        flex.ExpandStringPointerNullAsEmpty(m.Password),
		SshKeys:                         flex.ExpandFrameworkListNestedBlock(ctx, m.SshKeys, diags, ExpandAdminuserSshKeys),
		TimeZone:                        flex.ExpandStringPointer(m.TimeZone),
	}
}

// ApplyAdminuserNIOSUseFlags derives NIOS use flags from the raw config
// value(s) and writes them onto the core model. A flag is true when the user
// set any of its governed value fields in config.
func ApplyAdminuserNIOSUseFlags(ctx context.Context, config tfsdk.Config, obj *coremodel.Adminuser, diags *diag.Diagnostics) {
	if obj == nil || obj.NIOS == nil {
		return
	}
	obj.NIOS.UseSshKeys = flex.DeriveUseFlag(ctx, config, diags, path.Root("nios").AtName("ssh_keys"))
	obj.NIOS.UseTimeZone = flex.DeriveUseFlag(ctx, config, diags, path.Root("nios").AtName("time_zone"))
}

// Flatten populates the TF model from a core response.
func (m *AdminuserModel) Flatten(ctx context.Context, resp *coremodel.Adminuser, diags *diag.Diagnostics) {
	if resp == nil {
		return
	}

	m.Id = flex.FlattenStringPointer(resp.Id)

	// Extract existing NIOS model, flatten API response onto it, convert back
	niosModel := flex.ExpandNestedObject[NIOSAdminuserModel](ctx, m.NIOS, diags)
	if niosModel == nil {
		niosModel = &NIOSAdminuserModel{}
	}
	plannedNIOS := flex.ExpandNestedObject[NIOSAdminuserModel](ctx, m.NIOS, diags)
	niosModel.Flatten(ctx, resp.NIOS, diags)
	if resp.NIOS != nil {
		PostFlattenAdminuserNIOS(ctx, plannedNIOS, niosModel, diags)
		m.NIOS = flex.FlattenNestedObject(ctx, niosModel, NIOSAdminuserAttrTypes, diags)
	} else {
		m.NIOS = types.ObjectNull(NIOSAdminuserAttrTypes)
	}

}

// Flatten merges API response onto existing NIOS model.
func (m *NIOSAdminuserModel) Flatten(ctx context.Context, from *coremodel.NIOSAdminuserExt, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	planExtAttrs := m.ExtAttrs
	if planExtAttrs.IsUnknown() {
		planExtAttrs = types.MapNull(types.StringType)
	}
	m.AdminGroups = flex.FlattenFrameworkListString(ctx, from.AdminGroups, diags)
	m.AuthMethod = flex.FlattenStringPointerEmptyAsNull(from.AuthMethod)
	m.AuthType = flex.FlattenStringPointerEmptyAsNull(from.AuthType)
	m.CaCertificateIssuer = flex.FlattenStringPointerEmptyAsNull(from.CaCertificateIssuer)
	m.ClientCertificateSerialNumber = flex.FlattenStringPointerEmptyAsNull(from.ClientCertificateSerialNumber)
	m.Comment = flex.FlattenStringPointerEmptyAsNull(from.Comment)
	m.Disable = flex.FlattenBoolPointer(from.Disable)
	m.Email = flex.FlattenStringPointerEmptyAsNull(from.Email)
	m.EnableCertificateAuthentication = flex.FlattenBoolPointer(from.EnableCertificateAuthentication)
	m.ExtAttrs, m.ExtAttrsAll = flex.FlattenEAs(planExtAttrs, from.ExtAttrs)
	m.Name = flex.FlattenStringPointerEmptyAsNull(from.Name)
	m.SshKeys = flex.FlattenFrameworkListNestedBlock(ctx, from.SshKeys, AdminuserSshKeysAttrTypes, diags, FlattenAdminuserSshKeys)
	m.TimeZone = flex.FlattenStringPointerEmptyAsNull(from.TimeZone)
}
