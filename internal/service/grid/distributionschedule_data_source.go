package grid

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
	coremodel "github.com/infobloxopen/terraform-provider-infoblox/internal/core/model/grid"
	coresvc "github.com/infobloxopen/terraform-provider-infoblox/internal/core/service/grid"

	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/utils"
	customvalidator "github.com/infobloxopen/terraform-provider-infoblox/internal/validator"
)

var _ datasource.DataSource = &DistributionscheduleDataSource{}
var _ datasource.DataSourceWithValidateConfig = &DistributionscheduleDataSource{}
var _ datasource.DataSourceWithConfigure = &DistributionscheduleDataSource{}

func NewDistributionscheduleDataSource() datasource.DataSource {
	return &DistributionscheduleDataSource{}
}

type DistributionscheduleDataSource struct {
	backend core.BackendType
	service coresvc.DistributionscheduleService
}

func (d *DistributionscheduleDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_distribution_schedule"
}

// DistributionscheduleDataSourceModel is the filter model for the datasource
type DistributionscheduleDataSourceModel struct {
	Filters    types.Map   `tfsdk:"filters"`
	Results    types.List  `tfsdk:"results"`
	MaxResults types.Int32 `tfsdk:"max_results"`
	Paging     types.Int32 `tfsdk:"paging"`
}

// FlattenResults flattens core records to the Results list using existing Flatten method.
func (m *DistributionscheduleDataSourceModel) FlattenResults(ctx context.Context, from []*coremodel.Distributionschedule, diags *diag.Diagnostics) {
	if len(from) == 0 {
		m.Results = types.ListNull(types.ObjectType{AttrTypes: DistributionscheduleAttrTypes})
		return
	}
	elements := make([]attr.Value, 0, len(from))
	for _, obj := range from {
		model := &DistributionscheduleModel{}
		model.Flatten(ctx, obj, diags)
		objValue, d := types.ObjectValueFrom(ctx, DistributionscheduleAttrTypes, model)
		diags.Append(d...)
		elements = append(elements, objValue)
	}
	list, d := types.ListValue(types.ObjectType{AttrTypes: DistributionscheduleAttrTypes}, elements)
	diags.Append(d...)
	m.Results = list
}

func (d *DistributionscheduleDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Retrieves information about existing Infoblox Distributionschedule from the NIOS backend.",
		Attributes: map[string]schema.Attribute{
			"filters": schema.MapAttribute{
				Description: "Filter are used to return a more specific list of results. Filters can be used to match resources by specific attributes.",
				ElementType: types.StringType,
				Optional:    true,
			},
			"results": schema.ListNestedAttribute{
				NestedObject: schema.NestedAttributeObject{
					Attributes: utils.DataSourceResultAttributes(DistributionscheduleResourceSchemaAttributes),
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
		},
	}
}

func (d *DistributionscheduleDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

	d.service = coresvc.NewDistributionscheduleService(d.backend, client.NIOS, client.UDDI)
}

func (d *DistributionscheduleDataSource) ValidateConfig(ctx context.Context, req datasource.ValidateConfigRequest, resp *datasource.ValidateConfigResponse) {
	var data DistributionscheduleDataSourceModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	customvalidator.ValidateDataSourceFilters(d.backend, types.MapNull(types.StringType), types.MapNull(types.StringType), data.MaxResults, types.Int32Null(), &resp.Diagnostics)
}

func (d *DistributionscheduleDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data DistributionscheduleDataSourceModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Build list options
	opts := &core.ListOptions{
		Filters:      flex.ExpandMapString(ctx, data.Filters, &resp.Diagnostics),
		ReturnFields: DistributionscheduleReturnFields,
		Paging:       1,
	}

	if !data.MaxResults.IsNull() {
		opts.MaxResults = data.MaxResults.ValueInt32()
	}
	if !data.Paging.IsNull() {
		opts.Paging = data.Paging.ValueInt32()
	}

	if resp.Diagnostics.HasError() {
		return
	}

	var allResults []*coremodel.Distributionschedule
	var err error

	switch d.backend {
	case core.BackendNIOS:
		allResults, err = core.ReadAllPagesNIOS(func(pageID string) ([]*coremodel.Distributionschedule, string, error) {
			opts.PageID = pageID
			recs, _, nextPageID, e := d.service.List(ctx, opts)
			return recs, nextPageID, e
		})
	case core.BackendUDDI:
		allResults, err = core.ReadAllPagesUDDI(func(offset, limit int32) ([]*coremodel.Distributionschedule, error) {
			opts.Offset = offset
			opts.Limit = limit
			recs, _, _, e := d.service.List(ctx, opts)
			return recs, e
		}, opts.Limit, opts.Paging)
	}

	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to list Distributionschedule records: %s", err))
		return
	}

	tflog.Info(ctx, fmt.Sprintf("Retrieved %d results", len(allResults)))

	// Flatten results
	data.FlattenResults(ctx, allResults, &resp.Diagnostics)

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
