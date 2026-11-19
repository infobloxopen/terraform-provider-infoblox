package ipamfederation

import (
	"context"
	"fmt"
	"io"
	"regexp"

	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"

	uddiclient "github.com/infobloxopen/universal-ddi-go-client/client"
	uddiipamfederation "github.com/infobloxopen/universal-ddi-go-client/ipamfederation"

	"github.com/infobloxopen/terraform-provider-infoblox/internal/core"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
)

// Next-available data sources call the UDDI SDK directly and do not go through the core
// service/mapper layer: they are UDDI-only and return a flat list of addresses, so there is
// no infoblox core model to map to.

var _ datasource.DataSource = &NextAvailableReservedBlockDataSource{}
var _ datasource.DataSourceWithConfigure = &NextAvailableReservedBlockDataSource{}

func NewNextAvailableReservedBlockDataSource() datasource.DataSource {
	return &NextAvailableReservedBlockDataSource{}
}

type NextAvailableReservedBlockDataSource struct {
	uddi *uddiclient.APIClient
}

type NextAvailableReservedBlockDataSourceModel struct {
	Id                 types.String `tfsdk:"id"`
	Cidr               types.Int64  `tfsdk:"cidr"`
	ReservedBlockCount types.Int64  `tfsdk:"reserved_block_count"`
	Name               types.String `tfsdk:"name"`
	Comment            types.String `tfsdk:"comment"`
	TagFilters         types.Map    `tfsdk:"tag_filters"`
	Results            types.List   `tfsdk:"results"`
}

func (m *NextAvailableReservedBlockDataSourceModel) FlattenResults(ctx context.Context, from []uddiipamfederation.ReservedBlock, diags *diag.Diagnostics) {
	values := make([]string, 0, len(from))
	for _, b := range from {
		if b.Address != nil {
			values = append(values, *b.Address)
		}
	}
	m.Results = flex.FlattenFrameworkListString(ctx, values, diags)
}

func (d *NextAvailableReservedBlockDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_next_available_reserved_block"
}

func (d *NextAvailableReservedBlockDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Retrieves the next available Reserved Block addresses from the specified parent Federated Block, without allocating them. Only applicable for the UDDI backend.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "The resource identifier of the parent Federated Block.",
				Validators: []validator.String{
					stringvalidator.RegexMatches(regexp.MustCompile(`^federation/federated_block/[0-9a-f-].*$`), "invalid resource ID specified"),
					stringvalidator.ConflictsWith(path.MatchRoot("tag_filters")),
					stringvalidator.AtLeastOneOf(path.MatchRoot("id"), path.MatchRoot("tag_filters")),
				},
			},
			"cidr": schema.Int64Attribute{
				Required:            true,
				MarkdownDescription: "The CIDR of the reserved blocks to be allocated.",
			},
			"reserved_block_count": schema.Int64Attribute{
				Optional:            true,
				MarkdownDescription: "The count of reserved blocks required. Defaults to 1 if not set.",
				Validators: []validator.Int64{
					int64validator.AtLeast(1),
				},
			},
			"name": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "The name to be provided.",
			},
			"comment": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "The description for the reserved block. May contain 0 to 1024 characters. Can include UTF-8.",
			},
			"tag_filters": schema.MapAttribute{
				ElementType:         types.StringType,
				Optional:            true,
				MarkdownDescription: "Key-value pairs to filter the candidate parent Federated Blocks by tags.",
			},
			"results": schema.ListAttribute{
				ElementType:         types.StringType,
				Computed:            true,
				MarkdownDescription: "List of next available reserved block addresses in the specified parent Federated Block.",
			},
		},
	}
}

func (d *NextAvailableReservedBlockDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

	if client.UDDI == nil {
		resp.Diagnostics.AddError(
			"Unsupported Backend",
			"infoblox_next_available_reserved_block is only supported on the UDDI backend.",
		)
		return
	}

	d.uddi = client.UDDI
}

func (d *NextAvailableReservedBlockDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data NextAvailableReservedBlockDataSourceModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	if data.ReservedBlockCount.IsNull() {
		data.ReservedBlockCount = types.Int64Value(1)
	}
	count := data.ReservedBlockCount.ValueInt64()
	cidr := data.Cidr.ValueInt64()

	if len(data.TagFilters.Elements()) > 0 {
		results, err := d.findByTags(ctx, &data, cidr, count, &resp.Diagnostics)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read next available reserved blocks by tags, got error: %s", err))
			return
		}
		if resp.Diagnostics.HasError() {
			return
		}
		if int64(len(results)) < count {
			resp.Diagnostics.AddError(
				"Insufficient Available Reserved Blocks",
				fmt.Sprintf("Requested %d reserved blocks with CIDR %d, but only %d were found across the matched Federated Blocks.", count, cidr, len(results)),
			)
			return
		}
		data.FlattenResults(ctx, results, &resp.Diagnostics)
	} else if !data.Id.IsNull() && !data.Id.IsUnknown() {
		results, err := d.findReservedBlocks(ctx, data.Id.ValueString(), cidr, count, data.Name, data.Comment)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read next available reserved blocks, got error: %s", err))
			return
		}
		data.FlattenResults(ctx, results, &resp.Diagnostics)
	} else {
		resp.Diagnostics.AddError("Missing Parameters", "Either id or tag_filters must be specified.")
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

// findByTags enumerates the Federated Blocks matching the tag filters and takes the next
// available reserved blocks from each, in order, until count is reached.
func (d *NextAvailableReservedBlockDataSource) findByTags(ctx context.Context, data *NextAvailableReservedBlockDataSourceModel, cidr, count int64, diags *diag.Diagnostics) ([]uddiipamfederation.ReservedBlock, error) {
	tfilter := core.BuildTagFilter(flex.ExpandMapString(ctx, data.TagFilters, diags))

	// Only id and cidr are needed from the parents. Restricting the fields also keeps the
	// List response clear of utilization_v6, which the SDK cannot currently unmarshal.
	// federated_realm must stay in the list because the SDK model requires it.
	blocks, err := core.ReadAllPagesUDDI(func(offset, limit int32) ([]uddiipamfederation.FederatedBlock, error) {
		apiRes, _, err := d.uddi.IPAMFederationAPI.FederatedBlockAPI.
			List(ctx).
			Tfilter(tfilter).
			Fields("id,cidr,federated_realm").
			Offset(offset).
			Limit(limit).
			Execute()
		if err != nil {
			return nil, err
		}
		return apiRes.GetResults(), nil
	})
	if err != nil {
		return nil, err
	}
	if len(blocks) == 0 {
		diags.AddError("No Federated Blocks Found", "No federated blocks found with the given tags.")
		return nil, nil
	}

	var results []uddiipamfederation.ReservedBlock
	for _, fb := range blocks {
		if fb.Id == nil || (fb.Cidr != nil && *fb.Cidr >= cidr) {
			continue
		}
		if int64(len(results)) >= count {
			break
		}
		found, err := d.findReservedBlocks(ctx, *fb.Id, cidr, count-int64(len(results)), data.Name, data.Comment)
		if err != nil {
			// This parent has no (or insufficient) space; try the next candidate instead of
			// failing the whole search.
			continue
		}
		results = append(results, found...)
	}
	return results, nil
}

// findReservedBlocks asks one parent Federated Block for up to count reserved blocks. On a 400
// it retries once with the count the server reports as available and returns that partial
// result; a 400 with no recoverable count is returned as an error.
func (d *NextAvailableReservedBlockDataSource) findReservedBlocks(ctx context.Context, id string, cidr, count int64, name, comment types.String) ([]uddiipamfederation.ReservedBlock, error) {
	build := func(c int64) uddiipamfederation.NextAvailableReservedBlockAPIListNextAvailableReservedBlocksRequest {
		r := d.uddi.IPAMFederationAPI.NextAvailableReservedBlockAPI.
			ListNextAvailableReservedBlocks(ctx, id).
			Cidr(cidr).
			Count(c)
		if !name.IsNull() && !name.IsUnknown() {
			r = r.Name(name.ValueString())
		}
		if !comment.IsNull() && !comment.IsUnknown() {
			r = r.Comment(comment.ValueString())
		}
		return r
	}

	apiRes, httpRes, err := build(count).Execute()
	if err != nil {
		if httpRes != nil && httpRes.StatusCode == 400 {
			body, _ := io.ReadAll(httpRes.Body)
			_ = httpRes.Body.Close()
			if available := int64(core.ExtractAvailableCountFromError(body)); available > 0 {
				retryRes, _, retryErr := build(available).Execute()
				if retryErr == nil {
					return retryRes.GetResults(), nil
				}
			}
		}
		return nil, err
	}
	return apiRes.GetResults(), nil
}
