package provider

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/list"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"

	niosclient "github.com/infobloxopen/infoblox-nios-go-client/client"
	gridclient "github.com/infobloxopen/infoblox-nios-go-client/grid"
	niosoption "github.com/infobloxopen/infoblox-nios-go-client/option"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/core"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/retry"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/service/acl"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/service/dhcp"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/service/dns"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/service/dtc"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/service/grid"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/service/ipam"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/service/keys"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/service/misc"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/service/rpz"
	uddiclient "github.com/infobloxopen/universal-ddi-go-client/client"
	uddioption "github.com/infobloxopen/universal-ddi-go-client/option"
)

var (
	_ provider.Provider                  = &InfobloxProvider{}
	_ provider.ProviderWithListResources = &InfobloxProvider{}
)

type (
	InfobloxProvider struct {
		version string
		commit  string
	}

	InfobloxProviderConfig struct {
		NIOS               *NIOSConfig `tfsdk:"nios"`
		UDDI               *UDDIConfig `tfsdk:"uddi"`
		OperationTimeout   types.Int64 `tfsdk:"operation_timeout"`
		ManageInternalIdEA types.Bool  `tfsdk:"manage_internal_id_ea"`
	}

	NIOSConfig struct {
		HostUrl  types.String `tfsdk:"host_url"`
		Username types.String `tfsdk:"username"`
		Password types.String `tfsdk:"password"`
	}

	UDDIConfig struct {
		PortalURL          types.String `tfsdk:"portal_url"`
		PortalKey          types.String `tfsdk:"portal_key"`
		NIOSLicenseUID     types.String `tfsdk:"nios_license_uid"`
		EnableNIOSPassthru types.Bool   `tfsdk:"enable_nios_passthru"`
		DefaultTags        types.Map    `tfsdk:"default_tags"`
	}
)

func (p *InfobloxProvider) Metadata(_ context.Context, _ provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "infoblox"
	resp.Version = p.version
}

func (p *InfobloxProvider) Schema(_ context.Context, _ provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "The Infoblox provider is used to interact with Infoblox NIOS and UDDI backends.",
		Attributes: map[string]schema.Attribute{
			"nios": buildNIOSAttribute(),
			"uddi": buildUDDIAttribute(),
			"operation_timeout": schema.Int64Attribute{
				Optional: true,
				Validators: []validator.Int64{
					int64validator.AtLeast(1),
				},
				MarkdownDescription: "Total time (in seconds) allowed for an operation, including any retries of it. Default value: 60",
			},
			"manage_internal_id_ea": schema.BoolAttribute{
				Optional:            true,
				MarkdownDescription: "Determines whether the provider manages the Terraform Internal ID extensible attribute in NIOS. This attribute is required by the provider to store the Terraform resource ID corresponding to NIOS objects. When true, the provider ensures the attribute exists and manages its lifecycle. When false, the provider does not validate, create, update, or otherwise manage the attribute. Default value: true",
			},
		},
	}
}

func buildNIOSAttribute() schema.Attribute {
	return schema.SingleNestedAttribute{
		Description: "Configuration for NIOS backend.",
		Optional:    true,
		Attributes: map[string]schema.Attribute{
			"host_url": schema.StringAttribute{
				MarkdownDescription: "URL for the NIOS host",
				Optional:            true,
			},
			"username": schema.StringAttribute{
				MarkdownDescription: "Username for the NIOS host",
				Optional:            true,
			},
			"password": schema.StringAttribute{
				MarkdownDescription: "Password for the NIOS host",
				Optional:            true,
				Sensitive:           true,
			},
		},
	}
}

func buildUDDIAttribute() schema.Attribute {
	return schema.SingleNestedAttribute{
		Description: "Configuration for UDDI backend.",
		Optional:    true,
		Attributes: map[string]schema.Attribute{
			"portal_url": schema.StringAttribute{
				MarkdownDescription: "URL for the Infoblox Portal, or its WAPI endpoint when `enable_nios_passthru` is true.",
				Optional:            true,
			},
			"portal_key": schema.StringAttribute{
				MarkdownDescription: "API key for accessing the UDDI API.",
				Optional:            true,
				Sensitive:           true,
			},
			"nios_license_uid": schema.StringAttribute{
				MarkdownDescription: "License UID of the NIOS Grid to manage, required when `enable_nios_passthru` is true.",
				Optional:            true,
			},
			"enable_nios_passthru": schema.BoolAttribute{
				MarkdownDescription: "Enable NIOS WAPI passthrough to manage objects on a NIOS Grid through the Infoblox Portal. Requires the NIOS Grid to be connected to the Portal. Default value: false",
				Optional:            true,
			},
			"default_tags": schema.MapAttribute{
				ElementType:         types.StringType,
				MarkdownDescription: "Tags applied to every UDDI object the provider creates or updates. A tag set on the resource itself takes precedence over the default of the same name. Not applicable when `enable_nios_passthru` is true.",
				Optional:            true,
			},
		},
	}
}

func (p *InfobloxProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
	var data InfobloxProviderConfig

	// Read config
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Validation
	if data.NIOS == nil && data.UDDI == nil {
		resp.Diagnostics.AddError(
			"Missing Configuration",
			"One of 'nios' or 'uddi' must be configured.",
		)
		return
	}

	if data.NIOS != nil && data.UDDI != nil {
		resp.Diagnostics.AddError(
			"Invalid Configuration",
			"Only one of 'nios' or 'uddi' can be configured at a time.",
		)
		return
	}

	// Set the global operation timeout if specified
	if !data.OperationTimeout.IsUnknown() && !data.OperationTimeout.IsNull() {
		retry.SetOperationTimeout(data.OperationTimeout.ValueInt64())
	}

	var infobloxClient core.InfobloxClient

	// NIOS configurations
	if data.NIOS != nil {
		infobloxClient.NIOS = niosclient.NewAPIClient(
			niosoption.WithClientName(fmt.Sprintf("terraform/%s#%s", p.version, p.commit)),
			niosoption.WithNIOSUsername(data.NIOS.Username.ValueString()),
			niosoption.WithNIOSPassword(data.NIOS.Password.ValueString()),
			niosoption.WithNIOSHostUrl(data.NIOS.HostUrl.ValueString()),
			niosoption.WithDebug(true),
		)
	}

	// UDDI configurations
	if data.UDDI != nil {
		if data.UDDI.EnableNIOSPassthru.IsUnknown() {
			resp.Diagnostics.AddError(
				"Invalid Configuration",
				"'uddi.enable_nios_passthru' is not known until apply, but the provider needs it during planning to select the backend. Use a value that is known before apply.",
			)
			return
		}

		// Passthrough reaches NIOS through the Infoblox Portal, so the backend is NIOS.
		if data.UDDI.EnableNIOSPassthru.ValueBool() {
			if !data.UDDI.DefaultTags.IsNull() {
				resp.Diagnostics.AddError(
					"Invalid Configuration",
					"'uddi.default_tags' is not applicable when 'uddi.enable_nios_passthru' is true. Remove 'default_tags' to manage NIOS through the Infoblox Portal.",
				)
				return
			}

			if data.UDDI.PortalURL.IsUnknown() || data.UDDI.PortalKey.IsUnknown() || data.UDDI.NIOSLicenseUID.IsUnknown() {
				resp.Diagnostics.AddError(
					"Invalid Configuration",
					"The 'uddi' attributes for NIOS through the Infoblox Portal are not known until apply, but the provider needs them during planning. Use values that are known before apply.",
				)
				return
			}

			client := p.newNIOSPassthruClient(data.UDDI, resp)
			if client == nil {
				return
			}

			infobloxClient.NIOS = client
		} else {
			if data.UDDI.NIOSLicenseUID.ValueString() != "" {
				resp.Diagnostics.AddError(
					"Invalid Configuration",
					"'uddi.nios_license_uid' is set but 'uddi.enable_nios_passthru' is not true.\n\n"+
						"For NIOS via the Infoblox Portal: set 'enable_nios_passthru = true' and 'portal_url' to the Portal's WAPI passthrough endpoint.\n"+
						"For UDDI objects: remove 'uddi.nios_license_uid' and set 'portal_url' to the Portal's CSP endpoint.\n\n"+
						"These are different URLs, so 'portal_url' must match the mode.",
				)
				return
			}

			defaultTags := expandDefaultTags(ctx, data.UDDI.DefaultTags, resp)
			if resp.Diagnostics.HasError() {
				return
			}

			client := uddiclient.NewAPIClient(
				uddioption.WithClientName(fmt.Sprintf("terraform/%s#%s", p.version, p.commit)),
				uddioption.WithCSPUrl(data.UDDI.PortalURL.ValueString()),
				uddioption.WithAPIKey(data.UDDI.PortalKey.ValueString()),
				uddioption.WithDefaultTags(defaultTags),
				uddioption.WithDebug(true),
			)

			infobloxClient.UDDI = client
		}
	}

	// The Terraform Internal ID EA lives on the Grid, so it applies to NIOS reached either way.
	if infobloxClient.NIOS != nil && !ensureNIOSPreRequisites(ctx, infobloxClient.NIOS, data.ManageInternalIdEA, resp) {
		return
	}

	// Set infoblox client
	resp.DataSourceData = &infobloxClient
	resp.ResourceData = &infobloxClient
	resp.ListResourceData = &infobloxClient
}

func expandDefaultTags(ctx context.Context, tags types.Map, resp *provider.ConfigureResponse) map[string]string {
	if tags.IsNull() || tags.IsUnknown() {
		return nil
	}

	defaultTags := make(map[string]string, len(tags.Elements()))
	resp.Diagnostics.Append(tags.ElementsAs(ctx, &defaultTags, false)...)

	return defaultTags
}

// newNIOSPassthruClient builds a NIOS client that reaches a Grid through the Infoblox Portal.
func (p *InfobloxProvider) newNIOSPassthruClient(
	uddi *UDDIConfig,
	resp *provider.ConfigureResponse,
) *niosclient.APIClient {
	options := []niosoption.ClientOption{
		niosoption.WithClientName(fmt.Sprintf("terraform/%s#%s", p.version, p.commit)),
		niosoption.WithNIOSPassthrough(true),
		niosoption.WithPortalUrl(uddi.PortalURL.ValueString()),
		niosoption.WithPortalAPIKey(uddi.PortalKey.ValueString()),
		niosoption.WithNIOSLicenseUID(uddi.NIOSLicenseUID.ValueString()),
		niosoption.WithDebug(true),
	}

	// Validate the options before creating the client to catch missing required fields early.
	if err := niosoption.ValidatePassthrough(options...); err != nil {
		resp.Diagnostics.AddError("Missing Infoblox Portal Configuration", err.Error())
		return nil
	}

	return niosclient.NewAPIClient(options...)
}

// ensureNIOSPreRequisites creates the Terraform Internal ID extensible attribute unless the user opted out, reporting false once a failure is recorded.
func ensureNIOSPreRequisites(
	ctx context.Context,
	client *niosclient.APIClient,
	manage types.Bool,
	resp *provider.ConfigureResponse,
) bool {
	// Defaults to true, so only an explicit false opts out.
	if !manage.IsNull() && !manage.IsUnknown() && !manage.ValueBool() {
		resp.Diagnostics.AddWarning(
			"Terraform Internal ID Check Disabled",
			fmt.Sprintf("The %q extensible attribute check is disabled (manage_internal_id_ea=false). "+
				"Operations on NIOS-managed resources may fail if the extensible attribute does not exist in NIOS.",
				flex.TerraformInternalID),
		)
		return true
	}

	if err := checkAndCreatePreRequisitesForNIOS(ctx, client); err != nil {
		resp.Diagnostics.AddError(
			"Failed to ensure Terraform extensible attribute exists",
			err.Error(),
		)
		return false
	}

	return true
}

func (p *InfobloxProvider) Resources(_ context.Context) []func() resource.Resource {
	return []func() resource.Resource{
		acl.NewNamedaclResource,

		dhcp.NewDhcpOptiondefinitionResource,
		dhcp.NewDhcpOptionspaceResource,
		dhcp.NewFilteroptionResource,
		dhcp.NewHaGroupResource,
		dhcp.NewIpv6DhcpOptiondefinitionResource,
		dhcp.NewIpv6DhcpOptionspaceResource,
		dhcp.NewIpv6fixedaddressResource,
		dhcp.NewIpv6rangetemplateResource,

		dns.NewAuthNsgResource,
		dns.NewDnsServerResource,
		dns.NewForwardNsgResource,
		dns.NewNsgroupResource,
		dns.NewRecordAResource,
		dns.NewRecordAaaaResource,
		dns.NewRecordAliasResource,
		dns.NewRecordCaaResource,
		dns.NewRecordPtrResource,
		dns.NewRecordCnameResource,
		dns.NewRecordDnameResource,
		dns.NewRecordMxResource,
		dns.NewRecordNaptrResource,
		dns.NewRecordNsResource,
		dns.NewRecordSrvResource,
		dns.NewRecordTxtResource,
		dns.NewSharedrecordAResource,
		dns.NewSharedrecordAaaaResource,
		dns.NewViewResource,
		dns.NewZoneAuthResource,
		dns.NewZoneDelegatedResource,
		dns.NewZoneForwardResource,
		dns.NewZoneRpResource,

		dtc.NewDtcPoolResource,
		dtc.NewDtcServerResource,

		grid.NewExtensibleattributedefResource,
		grid.NewNatgroupResource,
		grid.NewUpgradegroupResource,

		ipam.NewAddressResource,
		ipam.NewIpv6networkResource,
		ipam.NewIpv6networkcontainerResource,
		ipam.NewNetworkResource,
		ipam.NewNetworkcontainerResource,
		ipam.NewNetworkviewResource,

		keys.NewTsigKeyResource,

		misc.NewBfdtemplateResource,
		misc.NewRulesetResource,

		rpz.NewRecordRpzCnameClientipaddressdnResource,
		rpz.NewRecordRpzNaptrResource,
		rpz.NewRecordRpzTxtResource,
		rpz.NewRecordRpzAaaaIpaddressResource,
	}
}

func (p *InfobloxProvider) DataSources(ctx context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{
		acl.NewNamedaclDataSource,

		dhcp.NewDhcpOptiondefinitionDataSource,
		dhcp.NewDhcpOptionspaceDataSource,
		dhcp.NewFilteroptionDataSource,
		dhcp.NewHaGroupDataSource,
		dhcp.NewIpv6DhcpOptiondefinitionDataSource,
		dhcp.NewIpv6DhcpOptionspaceDataSource,
		dhcp.NewIpv6fixedaddressDataSource,
		dhcp.NewIpv6rangetemplateDataSource,

		dns.NewAuthNsgDataSource,
		dns.NewDnsServerDataSource,
		dns.NewForwardNsgDataSource,
		dns.NewNsgroupDataSource,
		dns.NewRecordADataSource,
		dns.NewRecordAaaaDataSource,
		dns.NewRecordAliasDataSource,
		dns.NewRecordCaaDataSource,
		dns.NewRecordPtrDataSource,
		dns.NewRecordCnameDataSource,
		dns.NewRecordDnameDataSource,
		dns.NewRecordMxDataSource,
		dns.NewRecordNaptrDataSource,
		dns.NewRecordNsDataSource,
		dns.NewRecordSrvDataSource,
		dns.NewRecordTxtDataSource,
		dns.NewSharedrecordADataSource,
		dns.NewSharedrecordAaaaDataSource,
		dns.NewViewDataSource,
		dns.NewZoneAuthDataSource,
		dns.NewZoneDelegatedDataSource,
		dns.NewZoneForwardDataSource,
		dns.NewZoneRpDataSource,

		dtc.NewDtcPoolDataSource,
		dtc.NewDtcServerDataSource,

		grid.NewExtensibleattributedefDataSource,
		grid.NewNatgroupDataSource,
		grid.NewUpgradegroupDataSource,

		ipam.NewAddressDataSource,
		ipam.NewIpv6networkDataSource,
		ipam.NewIpv6networkcontainerDataSource,
		ipam.NewNetworkDataSource,
		ipam.NewNetworkcontainerDataSource,
		ipam.NewNetworkviewDataSource,
		ipam.NewNextAvailableAddressBlockDataSource,
		ipam.NewNextAvailableIPDataSource,
		ipam.NewNextAvailableSubnetDataSource,

		keys.NewTsigKeyDataSource,

		misc.NewBfdtemplateDataSource,
		misc.NewRulesetDataSource,

		rpz.NewRecordRpzCnameClientipaddressdnDataSource,
		rpz.NewRecordRpzNaptrDataSource,
		rpz.NewRecordRpzTxtDataSource,
		rpz.NewRecordRpzAaaaIpaddressDataSource,
	}
}

func (p *InfobloxProvider) ListResources(_ context.Context) []func() list.ListResource {
	return []func() list.ListResource{
		acl.NewNamedaclList,

		dhcp.NewDhcpOptiondefinitionList,
		dhcp.NewDhcpOptionspaceList,
		dhcp.NewFilteroptionList,
		dhcp.NewHaGroupList,
		dhcp.NewIpv6DhcpOptiondefinitionList,
		dhcp.NewIpv6DhcpOptionspaceList,
		dhcp.NewIpv6fixedaddressList,
		dhcp.NewIpv6rangetemplateList,

		dns.NewAuthNsgList,
		dns.NewDnsServerList,
		dns.NewForwardNsgList,
		dns.NewNsgroupList,
		dns.NewRecordAList,
		dns.NewRecordAaaaList,
		dns.NewRecordAliasList,
		dns.NewRecordCaaList,
		dns.NewRecordPtrList,
		dns.NewRecordCnameList,
		dns.NewRecordDnameList,
		dns.NewRecordMxList,
		dns.NewRecordNaptrList,
		dns.NewRecordNsList,
		dns.NewRecordSrvList,
		dns.NewRecordTxtList,
		dns.NewSharedrecordAList,
		dns.NewSharedrecordAaaaList,
		dns.NewViewList,
		dns.NewZoneAuthList,
		dns.NewZoneDelegatedList,
		dns.NewZoneForwardList,
		dns.NewZoneRpList,

		dtc.NewDtcPoolList,
		dtc.NewDtcServerList,

		grid.NewExtensibleattributedefList,
		grid.NewNatgroupList,
		grid.NewUpgradegroupList,

		ipam.NewAddressList,
		ipam.NewIpv6networkList,
		ipam.NewIpv6networkcontainerList,
		ipam.NewNetworkList,
		ipam.NewNetworkcontainerList,
		ipam.NewNetworkviewList,

		keys.NewTsigKeyList,

		misc.NewBfdtemplateList,
		misc.NewRulesetList,

		rpz.NewRecordRpzCnameClientipaddressdnList,
		rpz.NewRecordRpzNaptrList,
		rpz.NewRecordRpzTxtList,
		rpz.NewRecordRpzAaaaIpaddressList,
	}
}

func New(version, commit string) func() provider.Provider {
	return func() provider.Provider {
		return &InfobloxProvider{
			version: version,
			commit:  commit,
		}
	}
}

// checkAndCreatePreRequisitesForNIOS creates the Terraform Internal ID extensible
// attribute definition on NIOS if it does not already exist. This EA is used to
// uniquely identify resources managed by Terraform across imports and drift
// detection.
func checkAndCreatePreRequisitesForNIOS(ctx context.Context, client *niosclient.APIClient) error {
	var readableAttributesForEADefinition = "allowed_object_types,comment,default_value,flags,list_values,max,min,name,namespace,type"

	filters := map[string]any{
		"name": flex.TerraformInternalID,
	}

	err := retry.Do(ctx, retry.Transient(), func(ctx context.Context) (int, error) {
		// Check if EA already exists
		apiRes, httpRes, callErr := client.GridAPI.ExtensibleattributedefAPI.
			List(ctx).
			Filters(filters).
			ReturnFieldsPlus(readableAttributesForEADefinition).
			ReturnAsObject(1).
			Execute()
		if callErr != nil {
			if httpRes != nil {
				return httpRes.StatusCode, fmt.Errorf("error checking for existing extensible attribute: %w", callErr)
			}
			return 0, fmt.Errorf("error checking for existing extensible attribute: %w", callErr)
		}

		// If EA already exists, creation is not required
		if len(apiRes.ListExtensibleattributedefResponseObject.GetResult()) > 0 {
			return http.StatusOK, nil
		}

		// Create EA if it doesn't exist
		data := gridclient.Extensibleattributedef{
			Name:    gridclient.PtrString(flex.TerraformInternalID),
			Type:    gridclient.PtrString("STRING"),
			Comment: gridclient.PtrString("Internal ID for Terraform Resource"),
			Flags:   gridclient.PtrString("CR"),
		}

		_, httpRes, callErr = client.GridAPI.ExtensibleattributedefAPI.
			Create(ctx).
			Extensibleattributedef(data).
			ReturnFieldsPlus(readableAttributesForEADefinition).
			ReturnAsObject(1).
			Execute()
		if callErr != nil {
			callErr = fmt.Errorf("error creating Terraform extensible attribute: %w", callErr)
		}
		if httpRes != nil {
			return httpRes.StatusCode, callErr
		}
		return 0, callErr
	})

	return err
}
