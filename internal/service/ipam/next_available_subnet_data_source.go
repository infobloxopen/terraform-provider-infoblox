package ipam

import (
	"context"
	"fmt"
	"io"
	"regexp"

	"github.com/hashicorp/terraform-plugin-framework-validators/int32validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"

	uddiclient "github.com/infobloxopen/bloxone-go-client/client"
	uddiipam "github.com/infobloxopen/bloxone-go-client/ipam"

	"github.com/infobloxopen/terraform-provider-unified/internal/core"
	"github.com/infobloxopen/terraform-provider-unified/internal/flex"
)

// Next-available data sources call the UDDI SDK directly and do not go through
// the core service/mapper layer: they are UDDI-only and return a flat list of
// allocated values, so there is no unified core model to map to.

var _ datasource.DataSource = &NextAvailableSubnetDataSource{}
var _ datasource.DataSourceWithConfigure = &NextAvailableSubnetDataSource{}

func NewNextAvailableSubnetDataSource() datasource.DataSource {
	return &NextAvailableSubnetDataSource{}
}

type NextAvailableSubnetDataSource struct {
	uddi *uddiclient.APIClient
}

type NextAvailableSubnetDataSourceModel struct {
	Id         types.String `tfsdk:"id"`
	Cidr       types.Int64  `tfsdk:"cidr"`
	Count      types.Int32  `tfsdk:"subnet_count"`
	TagFilters types.Map    `tfsdk:"tag_filters"`
	Results    types.List   `tfsdk:"results"`
}

func (m *NextAvailableSubnetDataSourceModel) FlattenResults(ctx context.Context, from []uddiipam.Subnet, diags *diag.Diagnostics) {
	values := make([]string, 0, len(from))
	for _, s := range from {
		if s.Address != nil {
			values = append(values, *s.Address)
		}
	}
	m.Results = flex.FlattenFrameworkListString(ctx, values, diags)
}

func (d *NextAvailableSubnetDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_next_available_subnets"
}

func (d *NextAvailableSubnetDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Retrieves the next available subnets in the specified address block. Only applicable for the UDDI backend.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "An application specific resource identity of a resource.",
				Validators: []validator.String{
					stringvalidator.RegexMatches(regexp.MustCompile(`^ipam/address_block/[0-9a-f-].*$`), "invalid resource ID specified"),
					stringvalidator.ConflictsWith(path.MatchRoot("tag_filters")),
				},
			},
			"cidr": schema.Int64Attribute{
				Required:            true,
				MarkdownDescription: "The cidr value of subnets to be created.",
			},
			"subnet_count": schema.Int32Attribute{
				Optional:            true,
				MarkdownDescription: "Number of subnets to generate. Default 1 if not set.",
				Validators: []validator.Int32{
					int32validator.AtLeast(1),
				},
			},
			"tag_filters": schema.MapAttribute{
				ElementType:         types.StringType,
				Optional:            true,
				MarkdownDescription: "Key-value pairs to filter subnets by tags.",
			},
			"results": schema.ListAttribute{
				ElementType:         types.StringType,
				Computed:            true,
				MarkdownDescription: "List of next available subnet addresses in the specified resource.",
			},
		},
	}
}

func (d *NextAvailableSubnetDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*core.UnifiedClient)
	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected DataSource Configure Type",
			fmt.Sprintf("Expected *core.UnifiedClient, got: %T.", req.ProviderData),
		)
		return
	}

	if client.UDDI == nil {
		resp.Diagnostics.AddError(
			"Unsupported Backend",
			"unified_next_available_subnets is only supported on the UDDI backend.",
		)
		return
	}

	d.uddi = client.UDDI
}

func (d *NextAvailableSubnetDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data NextAvailableSubnetDataSourceModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	if data.Count.IsNull() {
		data.Count = types.Int32Value(1)
	}
	count := data.Count.ValueInt32()
	cidr := int32(data.Cidr.ValueInt64())

	if len(data.TagFilters.Elements()) > 0 {
		results, err := d.findByTags(ctx, data.TagFilters, cidr, count, &resp.Diagnostics)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read next available subnets by tags, got error: %s", err))
			return
		}
		if int32(len(results)) < count {
			resp.Diagnostics.AddError(
				"Insufficient Available Subnets",
				fmt.Sprintf("Requested %d subnets with CIDR %d, but only %d were found across the matched address blocks.", count, cidr, len(results)),
			)
			return
		}
		data.FlattenResults(ctx, results, &resp.Diagnostics)
	} else if !data.Id.IsNull() {
		apiRes, _, err := d.uddi.IPAddressManagementAPI.AddressBlockAPI.
			ListNextAvailableSubnet(ctx, data.Id.ValueString()).
			Cidr(cidr).
			Count(count).
			Execute()
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read next available subnets, got error: %s", err))
			return
		}
		data.FlattenResults(ctx, apiRes.GetResults(), &resp.Diagnostics)
	} else {
		resp.Diagnostics.AddError("Missing Parameters", "Either id or tag_filters must be specified.")
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

// findByTags enumerates address blocks matching the tag filters and allocates
// the next available subnets from each in address order until count is reached.
func (d *NextAvailableSubnetDataSource) findByTags(ctx context.Context, tagFilters types.Map, cidr, count int32, diags *diag.Diagnostics) ([]uddiipam.Subnet, error) {
	tfilter := core.BuildTagFilter(flex.ExpandMapString(ctx, tagFilters, diags))

	blocks, err := core.ReadAllPagesUDDI(func(offset, limit int32) ([]uddiipam.AddressBlock, error) {
		apiRes, _, err := d.uddi.IPAddressManagementAPI.AddressBlockAPI.
			List(ctx).
			Tfilter(tfilter).
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

	var results []uddiipam.Subnet
	for _, ab := range blocks {
		if ab.Cidr != nil && *ab.Cidr >= int64(cidr) {
			continue
		}
		if int32(len(results)) >= count {
			break
		}
		remaining := count - int32(len(results))
		found, err := d.findSubnet(ctx, *ab.Id, cidr, remaining)
		if err != nil {
			return nil, err
		}
		results = append(results, found...)
	}
	return results, nil
}

// findSubnet allocates up to count subnets from a single address block. On a 400
// response it retries with the count the server reports as available, returning
// that partial result. A 400 with no recoverable count yields an empty result so
// the caller can skip the block any other error is returned to the caller.
func (d *NextAvailableSubnetDataSource) findSubnet(ctx context.Context, id string, cidr, count int32) ([]uddiipam.Subnet, error) {
	apiRes, httpRes, err := d.uddi.IPAddressManagementAPI.AddressBlockAPI.
		ListNextAvailableSubnet(ctx, id).
		Cidr(cidr).
		Count(count).
		Execute()
	if err != nil {
		if httpRes != nil && httpRes.StatusCode == 400 {
			body, _ := io.ReadAll(httpRes.Body)
			_ = httpRes.Body.Close()
			if available := core.ExtractAvailableCountFromError(body); available > 0 {
				retryRes, _, retryErr := d.uddi.IPAddressManagementAPI.AddressBlockAPI.
					ListNextAvailableSubnet(ctx, id).
					Cidr(cidr).
					Count(available).
					Execute()
				if retryErr != nil {
					return nil, retryErr
				}
				return retryRes.GetResults(), nil
			}
			// 400 with no available count: the block cannot satisfy the
			// request, so skip it rather than failing the whole read.
			return nil, nil
		}
		return nil, err
	}
	return apiRes.GetResults(), nil
}
