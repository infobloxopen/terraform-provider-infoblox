package dtc

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework-validators/int32validator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/infobloxopen/terraform-provider-infoblox/internal/core"
	coremodel "github.com/infobloxopen/terraform-provider-infoblox/internal/core/model/dtc"
	coresvc "github.com/infobloxopen/terraform-provider-infoblox/internal/core/service/dtc"

	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/utils"
	customvalidator "github.com/infobloxopen/terraform-provider-infoblox/internal/validator"
)

var _ datasource.DataSource = &DtcMonitorTcpDataSource{}
var _ datasource.DataSourceWithValidateConfig = &DtcMonitorTcpDataSource{}
var _ datasource.DataSourceWithConfigure = &DtcMonitorTcpDataSource{}

func NewDtcMonitorTcpDataSource() datasource.DataSource {
	return &DtcMonitorTcpDataSource{}
}

type DtcMonitorTcpDataSource struct {
	backend core.BackendType
	service coresvc.DtcMonitorTcpService
}

func (d *DtcMonitorTcpDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_dtc_monitor_tcp"
}

// DtcMonitorTcpDataSourceModel is the filter model for the datasource
type DtcMonitorTcpDataSourceModel struct {
	Filters        types.Map   `tfsdk:"filters"`
	ExtAttrFilters types.Map   `tfsdk:"ext_attr_filters"`
	TagFilters     types.Map   `tfsdk:"tag_filters"`
	Results        types.List  `tfsdk:"results"`
	MaxResults     types.Int32 `tfsdk:"max_results"`
	Paging         types.Int32 `tfsdk:"paging"`
	Limit          types.Int32 `tfsdk:"limit"`
}

// FlattenResults flattens core records to the Results list using existing Flatten method.
func (m *DtcMonitorTcpDataSourceModel) FlattenResults(ctx context.Context, from []*coremodel.DtcMonitorTcp, diags *diag.Diagnostics) {
	if len(from) == 0 {
		m.Results = types.ListNull(types.ObjectType{AttrTypes: DtcMonitorTcpAttrTypes})
		return
	}
	elements := make([]attr.Value, 0, len(from))
	for _, obj := range from {
		model := &DtcMonitorTcpModel{}
		model.Flatten(ctx, obj, diags)
		objValue, d := types.ObjectValueFrom(ctx, DtcMonitorTcpAttrTypes, model)
		diags.Append(d...)
		elements = append(elements, objValue)
	}
	list, d := types.ListValue(types.ObjectType{AttrTypes: DtcMonitorTcpAttrTypes}, elements)
	diags.Append(d...)
	m.Results = list
}

func (d *DtcMonitorTcpDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Retrieves information about existing Infoblox DtcMonitorTcp from both the NIOS and UDDI backends.",
		Attributes: map[string]schema.Attribute{
			"filters": schema.MapAttribute{
				Description: "Filter are used to return a more specific list of results. Filters can be used to match resources by specific attributes.",
				ElementType: types.StringType,
				Optional:    true,
			},
			"ext_attr_filters": schema.MapAttribute{
				Description: "Extensible Attribute Filters are used to return a more specific list of results by filtering on extensible attributes. Only applicable for NIOS backend.",
				ElementType: types.StringType,
				Optional:    true,
			},
			"tag_filters": schema.MapAttribute{
				Description: "Tag Filters are used to return a more specific list of results filtered by tags. Only applicable for UDDI backend.",
				ElementType: types.StringType,
				Optional:    true,
			},
			"results": schema.ListNestedAttribute{
				NestedObject: schema.NestedAttributeObject{
					Attributes: utils.DataSourceResultAttributes(DtcMonitorTcpResourceSchemaAttributes),
				},
				Computed: true,
			},
			"paging": schema.Int32Attribute{
				Optional:    true,
				Description: "Enable (1) or disable (0) paging for the data source query. Enabled by default. When disabled, only a single page of results is retrieved.",
				Validators: []validator.Int32{
					int32validator.OneOf(0, 1),
				},
			},
			"max_results": schema.Int32Attribute{
				Optional:    true,
				Description: "Number of results to return per page. Defaults to 1000. Only applicable for NIOS backend.",
			},
			"limit": schema.Int32Attribute{
				Optional:    true,
				Description: "Number of results to return per page. Defaults to 1000. Only applicable for UDDI backend.",
			},
		},
	}
}

func (d *DtcMonitorTcpDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*core.InfobloxClient)
	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected DataSource Configure Type",
			fmt.Sprintf("Expected *core.InfobloxClient, got: %T.", req.ProviderData),
		)
		return
	}

	if client.NIOS != nil {
		d.backend = core.BackendNIOS
	} else {
		d.backend = core.BackendUDDI
	}

	d.service = coresvc.NewDtcMonitorTcpService(d.backend, client.NIOS, client.UDDI)
}

func (d *DtcMonitorTcpDataSource) ValidateConfig(ctx context.Context, req datasource.ValidateConfigRequest, resp *datasource.ValidateConfigResponse) {
	var data DtcMonitorTcpDataSourceModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	customvalidator.ValidateDataSourceFilters(d.backend, data.ExtAttrFilters, data.TagFilters, data.MaxResults, data.Limit, &resp.Diagnostics)
}

func (d *DtcMonitorTcpDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data DtcMonitorTcpDataSourceModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Build list options
	opts := &core.ListOptions{
		Filters:       flex.ExpandMapString(ctx, data.Filters, &resp.Diagnostics),
		ExtAttrFilter: flex.ExpandMapString(ctx, data.ExtAttrFilters, &resp.Diagnostics),
		TagFilter:     flex.ExpandMapString(ctx, data.TagFilters, &resp.Diagnostics),
		ReturnFields:  DtcMonitorTcpReturnFields,
		Paging:        1,
	}

	if !data.MaxResults.IsNull() {
		opts.MaxResults = data.MaxResults.ValueInt32()
	}
	if !data.Paging.IsNull() {
		opts.Paging = data.Paging.ValueInt32()
	}
	if !data.Limit.IsNull() {
		opts.Limit = data.Limit.ValueInt32()
	}

	if resp.Diagnostics.HasError() {
		return
	}

	var allResults []*coremodel.DtcMonitorTcp
	var err error

	switch d.backend {
	case core.BackendNIOS:
		allResults, err = core.ReadAllPagesNIOS(func(pageID string) ([]*coremodel.DtcMonitorTcp, string, error) {
			opts.PageID = pageID
			recs, _, nextPageID, e := d.service.List(ctx, opts)
			return recs, nextPageID, e
		})
	case core.BackendUDDI:
		allResults, err = core.ReadAllPagesUDDI(func(offset, limit int32) ([]*coremodel.DtcMonitorTcp, error) {
			opts.Offset = offset
			opts.Limit = limit
			recs, _, _, e := d.service.List(ctx, opts)
			return recs, e
		}, opts.Limit, opts.Paging)
	}

	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to list DtcMonitorTcp records: %s", err))
		return
	}

	tflog.Info(ctx, fmt.Sprintf("Retrieved %d results", len(allResults)))

	// Flatten results
	data.FlattenResults(ctx, allResults, &resp.Diagnostics)

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
