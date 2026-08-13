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
	"github.com/infobloxopen/terraform-provider-infoblox/internal/service/dns"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/service/ipam"
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
		NIOS             *NIOSConfig `tfsdk:"nios"`
		UDDI             *UDDIConfig `tfsdk:"uddi"`
		OperationTimeout types.Int64 `tfsdk:"operation_timeout"`
	}

	NIOSConfig struct {
		HostUrl            types.String `tfsdk:"host_url"`
		Username           types.String `tfsdk:"username"`
		Password           types.String `tfsdk:"password"`
		ManageInternalIdEA types.Bool   `tfsdk:"manage_internal_id_ea"`
	}

	UDDIConfig struct {
		CSPUrl types.String `tfsdk:"csp_url"`
		APIKey types.String `tfsdk:"api_key"`
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
			"manage_internal_id_ea": schema.BoolAttribute{
				Optional:            true,
				MarkdownDescription: "Determines whether the provider manages the Terraform Internal ID extensible attribute in NIOS. This attribute is required by the provider to store the Terraform resource ID corresponding to NIOS objects. When true, the provider ensures the attribute exists and manages its lifecycle. When false, the provider does not validate, create, update, or otherwise manage the attribute. Default value: true",
			},
		},
	}
}

func buildUDDIAttribute() schema.Attribute {
	return schema.SingleNestedAttribute{
		Description: "Configuration for UDDI backend.",
		Optional:    true,
		Attributes: map[string]schema.Attribute{
			"csp_url": schema.StringAttribute{
				MarkdownDescription: "URL for UDDI Cloud Services Portal",
				Optional:            true,
			},
			"api_key": schema.StringAttribute{
				MarkdownDescription: "API key for accessing the UDDI API.",
				Optional:            true,
				Sensitive:           true,
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
		client := niosclient.NewAPIClient(
			niosoption.WithClientName(fmt.Sprintf("terraform/%s#%s", p.version, p.commit)),
			niosoption.WithNIOSUsername(data.NIOS.Username.ValueString()),
			niosoption.WithNIOSPassword(data.NIOS.Password.ValueString()),
			niosoption.WithNIOSHostUrl(data.NIOS.HostUrl.ValueString()),
			niosoption.WithDebug(true),
		)

		// Default manage_internal_id_ea to true when not set.
		if data.NIOS.ManageInternalIdEA.IsUnknown() || data.NIOS.ManageInternalIdEA.IsNull() {
			data.NIOS.ManageInternalIdEA = types.BoolValue(true)
		}

		if data.NIOS.ManageInternalIdEA.ValueBool() {
			if err := checkAndCreatePreRequisitesForNIOS(ctx, client); err != nil {
				resp.Diagnostics.AddError(
					"Failed to ensure Terraform extensible attribute exists",
					err.Error(),
				)
				return
			}
		} else {
			// Raise a warning if the provider is not managing the Terraform Internal ID EA,
			// as this may lead to issues with resource management.
			resp.Diagnostics.AddWarning(
				"Terraform Internal ID Check Disabled",
				fmt.Sprintf("The %q extensible attribute check is disabled (manage_internal_id_ea=false). "+
					"Operations on NIOS-managed resources may fail if the extensible attribute does not exist in NIOS.",
					flex.TerraformInternalID),
			)
		}

		infobloxClient.NIOS = client
	}

	// UDDI configurations
	if data.UDDI != nil {
		client := uddiclient.NewAPIClient(
			uddioption.WithClientName(fmt.Sprintf("terraform/%s#%s", p.version, p.commit)),
			uddioption.WithCSPUrl(data.UDDI.CSPUrl.ValueString()),
			uddioption.WithAPIKey(data.UDDI.APIKey.ValueString()),
			uddioption.WithDebug(true),
		)

		infobloxClient.UDDI = client
	}

	// Set infoblox client
	resp.DataSourceData = &infobloxClient
	resp.ResourceData = &infobloxClient
	resp.ListResourceData = &infobloxClient
}

func (p *InfobloxProvider) Resources(_ context.Context) []func() resource.Resource {
	return []func() resource.Resource{
		dns.NewRecordAliasResource,
		dns.NewRecordNaptrResource,
		dns.NewRecordCnameResource,
		dns.NewRecordAaaaResource,
		dns.NewRecordTxtResource,
		dns.NewRecordCaaResource,
		dns.NewRecordDnameResource,
		dns.NewRecordNsResource,
		dns.NewZoneAuthResource,
		dns.NewViewResource,
		dns.NewRecordAResource,

		ipam.NewAddressResource,
	}
}

func (p *InfobloxProvider) DataSources(ctx context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{
		dns.NewRecordAliasDataSource,
		dns.NewRecordNaptrDataSource,
		dns.NewRecordCnameDataSource,
		dns.NewRecordAaaaDataSource,
		dns.NewRecordTxtDataSource,
		dns.NewRecordCaaDataSource,
		dns.NewRecordDnameDataSource,
		dns.NewRecordNsDataSource,
		dns.NewZoneAuthDataSource,
		dns.NewViewDataSource,
		dns.NewRecordADataSource,

		ipam.NewNextAvailableIPDataSource,
		ipam.NewNextAvailableSubnetDataSource,
		ipam.NewNextAvailableAddressBlockDataSource,
		ipam.NewAddressDataSource,
	}
}

func (p *InfobloxProvider) ListResources(_ context.Context) []func() list.ListResource {
	return []func() list.ListResource{
		dns.NewRecordAliasList,
		dns.NewRecordNaptrList,
		dns.NewRecordCnameList,
		dns.NewRecordAaaaList,
		dns.NewRecordTxtList,
		dns.NewRecordCaaList,
		dns.NewRecordDnameList,
		dns.NewRecordNsList,
		dns.NewZoneAuthList,
		dns.NewViewList,
		dns.NewRecordAList,

		ipam.NewAddressList,
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
