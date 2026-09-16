package dhcp

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
	coremodel "github.com/infobloxopen/terraform-provider-infoblox/internal/core/model/dhcp"
	coresvc "github.com/infobloxopen/terraform-provider-infoblox/internal/core/service/dhcp"

	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/utils"
	customvalidator "github.com/infobloxopen/terraform-provider-infoblox/internal/validator"
)

var _ datasource.DataSource = &HaGroupDataSource{}
var _ datasource.DataSourceWithValidateConfig = &HaGroupDataSource{}
var _ datasource.DataSourceWithConfigure = &HaGroupDataSource{}

func NewHaGroupDataSource() datasource.DataSource {
	return &HaGroupDataSource{}
}

type HaGroupDataSource struct {
	backend core.BackendType
	service coresvc.HaGroupService
}

func (d *HaGroupDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_ha_group"
}

// HaGroupDataSourceModel is the filter model for the datasource
type HaGroupDataSourceModel struct {
	Filters    types.Map   `tfsdk:"filters"`
	TagFilters types.Map   `tfsdk:"tag_filters"`
	Results    types.List  `tfsdk:"results"`
	Paging     types.Int32 `tfsdk:"paging"`
	Limit      types.Int32 `tfsdk:"limit"`
}

// FlattenResults flattens core records to the Results list using existing Flatten method.
func (m *HaGroupDataSourceModel) FlattenResults(ctx context.Context, from []*coremodel.HaGroup, diags *diag.Diagnostics) {
	if len(from) == 0 {
		m.Results = types.ListNull(types.ObjectType{AttrTypes: HaGroupAttrTypes})
		return
	}
	elements := make([]attr.Value, 0, len(from))
	for _, obj := range from {
		model := &HaGroupModel{}
		model.Flatten(ctx, obj, diags)
		objValue, d := types.ObjectValueFrom(ctx, HaGroupAttrTypes, model)
		diags.Append(d...)
		elements = append(elements, objValue)
	}
	list, d := types.ListValue(types.ObjectType{AttrTypes: HaGroupAttrTypes}, elements)
	diags.Append(d...)
	m.Results = list
}

func (d *HaGroupDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Retrieves information about existing Infoblox DHCP HA Groups.",
		Attributes: map[string]schema.Attribute{
			"filters": schema.MapAttribute{
				Description: "Filter are used to return a more specific list of results. Filters can be used to match resources by specific attributes.",
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
					Attributes: utils.DataSourceResultAttributes(HaGroupResourceSchemaAttributes),
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
			"limit": schema.Int32Attribute{
				Optional:    true,
				Description: "Number of results to return per page. Defaults to 1000. Only applicable for UDDI backend.",
			},
		},
	}
}

func (d *HaGroupDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

	d.service = coresvc.NewHaGroupService(d.backend, client.NIOS, client.UDDI)
}

func (d *HaGroupDataSource) ValidateConfig(ctx context.Context, req datasource.ValidateConfigRequest, resp *datasource.ValidateConfigResponse) {
	var data HaGroupDataSourceModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	customvalidator.ValidateDataSourceFilters(d.backend, types.MapNull(types.StringType), data.TagFilters, types.Int32Null(), data.Limit, &resp.Diagnostics)
}

func (d *HaGroupDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data HaGroupDataSourceModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Build list options
	opts := &core.ListOptions{
		Filters:      flex.ExpandMapString(ctx, data.Filters, &resp.Diagnostics),
		TagFilter:    flex.ExpandMapString(ctx, data.TagFilters, &resp.Diagnostics),
		ReturnFields: HaGroupReturnFields,
		Paging:       1,
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

	var allResults []*coremodel.HaGroup
	var err error

	switch d.backend {
	case core.BackendNIOS:
		allResults, err = core.ReadAllPagesNIOS(func(pageID string) ([]*coremodel.HaGroup, string, error) {
			opts.PageID = pageID
			recs, _, nextPageID, e := d.service.List(ctx, opts)
			return recs, nextPageID, e
		})
	case core.BackendUDDI:
		allResults, err = core.ReadAllPagesUDDI(func(offset, limit int32) ([]*coremodel.HaGroup, error) {
			opts.Offset = offset
			opts.Limit = limit
			recs, _, _, e := d.service.List(ctx, opts)
			return recs, e
		}, opts.Limit, opts.Paging)
	}

	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to list HaGroup records: %s", err))
		return
	}

	tflog.Info(ctx, fmt.Sprintf("Retrieved %d results", len(allResults)))

	// Flatten results
	data.FlattenResults(ctx, allResults, &resp.Diagnostics)

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
