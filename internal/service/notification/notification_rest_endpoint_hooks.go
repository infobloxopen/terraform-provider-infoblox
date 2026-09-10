package notification

import (
	"context"
	"fmt"
	"net/url"
	"strconv"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/utils"
)

// ValidateNotificationRestEndpoint validates the NotificationRestEndpoint configuration.
func ValidateNotificationRestEndpoint(ctx context.Context, data NotificationRestEndpointModel, resp *resource.ValidateConfigResponse) {
	if nios := flex.ExpandNestedObject[NIOSNotificationRestEndpointModel](ctx, data.NIOS, &resp.Diagnostics); nios != nil {
		validateNotificationRestEndpointNIOSConfig(ctx, nios, resp)
	}
}

func validateNotificationRestEndpointNIOSConfig(ctx context.Context, m *NIOSNotificationRestEndpointModel, resp *resource.ValidateConfigResponse) {
	// outbound_members is required when outbound_member_type == "MEMBER", forbidden otherwise.
	if !m.OutboundMemberType.IsNull() && !m.OutboundMemberType.IsUnknown() {
		if m.OutboundMemberType.ValueString() == "MEMBER" {
			if m.OutboundMembers.IsNull() || m.OutboundMembers.IsUnknown() {
				resp.Diagnostics.AddAttributeError(
					path.Root("nios").AtName("outbound_members"),
					"Invalid Configuration",
					"Attribute 'outbound_members' must be specified when 'outbound_member_type' is set to 'MEMBER'.",
				)
			}
		} else if !m.OutboundMembers.IsNull() && !m.OutboundMembers.IsUnknown() {
			resp.Diagnostics.AddAttributeError(
				path.Root("nios").AtName("outbound_members"),
				"Invalid Configuration",
				"Attribute 'outbound_members' cannot be specified when 'outbound_member_type' is set to 'GM'.",
			)
		}
	}

	// uri must be a valid URL.
	if !m.Uri.IsNull() && !m.Uri.IsUnknown() {
		if _, err := url.ParseRequestURI(m.Uri.ValueString()); err != nil {
			resp.Diagnostics.AddAttributeError(
				path.Root("nios").AtName("uri"),
				"Invalid URI",
				"URI must contain a valid value.",
			)
		}
	}

	// Validate template_instance parameter syntax vs value types.
	if m.TemplateInstance.IsNull() || m.TemplateInstance.IsUnknown() {
		return
	}
	var ti NotificationRestEndpointTemplateInstanceModel
	resp.Diagnostics.Append(m.TemplateInstance.As(ctx, &ti, basetypes.ObjectAsOptions{})...)
	if resp.Diagnostics.HasError() {
		return
	}
	if ti.Parameters.IsNull() || ti.Parameters.IsUnknown() {
		return
	}
	var params []NotificationrestendpointtemplateinstanceParametersModel
	resp.Diagnostics.Append(ti.Parameters.ElementsAs(ctx, &params, false)...)
	if resp.Diagnostics.HasError() {
		return
	}
	for i, p := range params {
		if p.Syntax.IsNull() || p.Syntax.IsUnknown() || p.Value.IsNull() || p.Value.IsUnknown() {
			continue
		}
		syntax := p.Syntax.ValueString()
		value := p.Value.ValueString()
		switch syntax {
		case "INT":
			if _, err := strconv.Atoi(value); err != nil {
				resp.Diagnostics.AddAttributeError(
					path.Root("nios").AtName("template_instance").AtName("parameters").AtListIndex(i).AtName("value"),
					"Invalid Value for INT Syntax",
					fmt.Sprintf("The value of the parameter definition '%s' is incorrect. The value type should be %s. Got: %s", p.Name.ValueString(), syntax, value),
				)
			}
		case "BOOL":
			if value != "True" && value != "False" {
				resp.Diagnostics.AddAttributeError(
					path.Root("nios").AtName("template_instance").AtName("parameters").AtListIndex(i).AtName("value"),
					"Invalid Value for BOOL Syntax",
					fmt.Sprintf("The value of the parameter definition '%s' is incorrect. The value type should be %s, either True/False (Case Sensitive). Got: %s", p.Name.ValueString(), syntax, value),
				)
			}
		}
	}
}

// PostFlattenNotificationRestEndpointNIOS copies write-only fields (password, wapi_user_password,
// client_certificate_file) from the plan back to the flattened state, since the NIOS API never
// echoes these values back.
func PostFlattenNotificationRestEndpointNIOS(ctx context.Context, planned, flattened *NIOSNotificationRestEndpointModel, diags *diag.Diagnostics) {
	if planned != nil {
		flattened.Password = planned.Password
		flattened.WapiUserPassword = planned.WapiUserPassword
		flattened.ClientCertificateFile = planned.ClientCertificateFile
		// token is write-only; NIOS never echoes it back.
		// Preserve only if known (set explicitly by user or populated by the upload hook).
		// An Unknown value means no upload happened and no explicit value was given — use null.
		if !planned.ClientCertificateToken.IsUnknown() {
			flattened.ClientCertificateToken = planned.ClientCertificateToken
		} else {
			flattened.ClientCertificateToken = flex.FlattenStringPointerEmptyAsNull(nil)
		}
	} else {
		flattened.Password = flex.FlattenStringPointerEmptyAsNull(nil)
		flattened.WapiUserPassword = flex.FlattenStringPointerEmptyAsNull(nil)
		flattened.ClientCertificateFile = types.StringNull()
		flattened.ClientCertificateToken = flex.FlattenStringPointerEmptyAsNull(nil)
	}
}

// ProcessNIOSNotificationRestEndpointFileUpload uploads the certificate file (if set) and stores
// the resulting token in client_certificate_token before the resource is expanded and sent to NIOS.
func (r *NotificationRestEndpointResource) ProcessNIOSNotificationRestEndpointFileUpload(ctx context.Context, data *NotificationRestEndpointModel, diags *diag.Diagnostics) bool {
	nios := flex.ExpandNestedObject[NIOSNotificationRestEndpointModel](ctx, data.NIOS, diags)
	if nios == nil || nios.ClientCertificateFile.IsNull() || nios.ClientCertificateFile.IsUnknown() {
		return true
	}
	token, err := utils.UploadFileWithToken(ctx, r.niosHostURL, nios.ClientCertificateFile.ValueString(), r.niosUsername, r.niosPassword)
	if err != nil {
		diags.AddError("File Upload Error", fmt.Sprintf("Failed to upload certificate file: %s", err))
		return false
	}
	nios.ClientCertificateToken = flex.FlattenStringPointer(&token)
	data.NIOS = flex.FlattenNestedObject(ctx, nios, NIOSNotificationRestEndpointAttrTypes, diags)
	return !diags.HasError()
}
