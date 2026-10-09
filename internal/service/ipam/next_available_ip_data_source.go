package ipam

import (
	"context"
	"fmt"
	"net/http"
	"regexp"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/mapvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"

	uddiipam "github.com/infobloxopen/universal-ddi-go-client/ipam"

	uddiclient "github.com/infobloxopen/universal-ddi-go-client/client"

	"github.com/infobloxopen/terraform-provider-infoblox/internal/core"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
)

// Next-available data sources call the UDDI SDK directly and do not go through
// the core service/mapper layer: they are UDDI-only and return a flat list of
// allocated values, so there is no infoblox core model to map to.

var _ datasource.DataSource = &NextAvailableIPDataSource{}
var _ datasource.DataSourceWithConfigure = &NextAvailableIPDataSource{}

func NewNextAvailableIPDataSource() datasource.DataSource {
	return &NextAvailableIPDataSource{}
}

type NextAvailableIPDataSource struct {
	uddi *uddiclient.APIClient
}

type NextAvailableIPDataSourceModel struct {
	Id           types.String `tfsdk:"id"`
	Contiguous   types.Bool   `tfsdk:"contiguous"`
	Count        types.Int64  `tfsdk:"ip_count"`
	TagFilters   types.Map    `tfsdk:"tag_filters"`
	ResourceType types.String `tfsdk:"resource_type"`
	Results      types.List   `tfsdk:"results"`
}

func (m *NextAvailableIPDataSourceModel) FlattenResults(ctx context.Context, from []uddiipam.Address, diags *diag.Diagnostics) {
	values := make([]string, 0, len(from))
	for _, a := range from {
		values = append(values, a.Address)
	}
	m.Results = flex.FlattenFrameworkListString(ctx, values, diags)
}

func (d *NextAvailableIPDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_next_available_ips"
}

func (d *NextAvailableIPDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Retrieves the next available IP addresses in the specified resource. The resource can be an address block, subnet or range. Only applicable for the UDDI backend.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "An application specific resource identity of a resource.",
				Validators: []validator.String{
					stringvalidator.RegexMatches(regexp.MustCompile(`^ipam/(range|subnet|address_block)/[0-9a-f-].*$`), "invalid resource ID specified"),
					stringvalidator.ConflictsWith(path.MatchRoot("tag_filters")),
					stringvalidator.ConflictsWith(path.MatchRoot("resource_type")),
				},
			},
			"contiguous": schema.BoolAttribute{
				Optional:            true,
				MarkdownDescription: "Indicates whether the IP addresses should belong to a contiguous block. Defaults to false.",
			},
			"ip_count": schema.Int64Attribute{
				Optional:            true,
				MarkdownDescription: "The number of IP addresses requested. Defaults to 1.",
				Validators: []validator.Int64{
					int64validator.Between(1, 20),
				},
			},
			"tag_filters": schema.MapAttribute{
				ElementType:         types.StringType,
				Optional:            true,
				MarkdownDescription: "Map of tag key/value pairs to filter resources.",
				Validators: []validator.Map{
					mapvalidator.AlsoRequires(path.MatchRoot("resource_type")),
				},
			},
			"resource_type": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Resource type to search when using tag filters (address_block, subnet, or range).",
				Validators: []validator.String{
					stringvalidator.OneOf("address_block", "subnet", "range"),
					stringvalidator.AlsoRequires(path.MatchRoot("tag_filters")),
				},
			},
			"results": schema.ListAttribute{
				ElementType:         types.StringType,
				Computed:            true,
				MarkdownDescription: "List of next available IP addresses in the specified resource.",
			},
		},
	}
}

func (d *NextAvailableIPDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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
			"infoblox_next_available_ips is only supported on the UDDI backend.",
		)
		return
	}

	d.uddi = client.UDDI
}

func (d *NextAvailableIPDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data NextAvailableIPDataSourceModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	if data.Count.IsNull() {
		data.Count = types.Int64Value(1)
	}
	count := data.Count.ValueInt64()

	if len(data.TagFilters.Elements()) > 0 {
		if data.ResourceType.IsNull() {
			resp.Diagnostics.AddError("Missing resource_type", "resource_type is required when using tag_filters.")
			return
		}
		addresses, err := d.findByTags(ctx, data, count, &resp.Diagnostics)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read next available IPs by tags, got error: %s", err))
			return
		}
		if int64(len(addresses)) < count {
			resp.Diagnostics.AddError(
				"Insufficient Available IPs",
				fmt.Sprintf("Requested %d IPs, but only %d were found in %ss matching the given tags.", count, len(addresses), data.ResourceType.ValueString()),
			)
			return
		}
		data.FlattenResults(ctx, addresses[:count], &resp.Diagnostics)
	} else if !data.Id.IsNull() {
		results, _, err := d.getByID(ctx, data.Id.ValueString(), int32(count), data.Contiguous.ValueBool())
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read next available IPs, got error: %s", err))
			return
		}
		data.FlattenResults(ctx, results, &resp.Diagnostics)
	} else {
		resp.Diagnostics.AddError("Missing Parameters", "Either id or tag_filters must be specified.")
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

// getByID allocates the next available IPs from a single resource, dispatching on
// the resource type encoded in the id prefix. It returns the HTTP status code
// alongside any error so callers can distinguish a 400 (resource cannot satisfy
// the request) from a genuine failure.
func (d *NextAvailableIPDataSource) getByID(ctx context.Context, id string, count int32, contiguous bool) ([]uddiipam.Address, int, error) {
	prefix := id[:strings.LastIndex(id, "/")]
	var (
		apiRes  *uddiipam.NextAvailableIPResponse
		httpRes *http.Response
		err     error
	)
	switch prefix {
	case "ipam/address_block":
		apiRes, httpRes, err = d.uddi.IPAddressManagementAPI.AddressBlockAPI.
			ListNextAvailableIP(ctx, id).Count(count).Contiguous(contiguous).Execute()
	case "ipam/subnet":
		apiRes, httpRes, err = d.uddi.IPAddressManagementAPI.SubnetAPI.
			ListNextAvailableIP(ctx, id).Count(count).Contiguous(contiguous).Execute()
	case "ipam/range":
		apiRes, httpRes, err = d.uddi.IPAddressManagementAPI.RangeAPI.
			ListNextAvailableIP(ctx, id).Count(count).Contiguous(contiguous).Execute()
	default:
		return nil, 0, fmt.Errorf("unsupported resource id: %s", id)
	}
	if err != nil {
		status := 0
		if httpRes != nil {
			status = httpRes.StatusCode
		}
		return nil, status, err
	}
	return apiRes.GetResults(), 0, nil
}

// findByTags enumerates resources of the requested type matching the tag filters
// and allocates the next available IPs from each until count is reached.
func (d *NextAvailableIPDataSource) findByTags(ctx context.Context, data NextAvailableIPDataSourceModel, count int64, diags *diag.Diagnostics) ([]uddiipam.Address, error) {
	tfilter := core.BuildTagFilter(flex.ExpandMapString(ctx, data.TagFilters, diags))

	resources, err := d.listResourcesByTag(ctx, data.ResourceType.ValueString(), tfilter)
	if err != nil {
		return nil, err
	}
	if len(resources) == 0 {
		return nil, fmt.Errorf("no %ss found with the given tags", data.ResourceType.ValueString())
	}

	contiguous := data.Contiguous.ValueBool()
	var all []uddiipam.Address
	for _, id := range resources {
		if int64(len(all)) >= count {
			break
		}
		remaining := int32(count - int64(len(all)))
		res, status, err := d.getByID(ctx, id, remaining, contiguous)
		if err != nil {
			if status != 400 {
				return nil, err
			}
			// 400: this resource cannot satisfy the full remaining count, so
			// fall back to progressively smaller counts to take what it has.
			for try := remaining - 1; try > 0; try-- {
				if retry, _, retryErr := d.getByID(ctx, id, try, contiguous); retryErr == nil && len(retry) > 0 {
					all = append(all, retry...)
					break
				}
			}
			continue
		}
		all = append(all, res...)
	}
	return all, nil
}

// listResourcesByTag returns the ids of resources of the given type matching the tfilter.
func (d *NextAvailableIPDataSource) listResourcesByTag(ctx context.Context, resourceType, tfilter string) ([]string, error) {
	var ids []string
	switch resourceType {
	case "address_block":
		blocks, err := core.ReadAllPagesUDDI(func(offset, limit int32) ([]uddiipam.AddressBlock, error) {
			res, _, err := d.uddi.IPAddressManagementAPI.AddressBlockAPI.
				List(ctx).Tfilter(tfilter).Offset(offset).Limit(limit).Execute()
			if err != nil {
				return nil, err
			}
			return res.GetResults(), nil
		})
		if err != nil {
			return nil, err
		}
		for _, b := range blocks {
			if b.Id != nil {
				ids = append(ids, *b.Id)
			}
		}
	case "subnet":
		subnets, err := core.ReadAllPagesUDDI(func(offset, limit int32) ([]uddiipam.Subnet, error) {
			res, _, err := d.uddi.IPAddressManagementAPI.SubnetAPI.
				List(ctx).Tfilter(tfilter).Offset(offset).Limit(limit).Execute()
			if err != nil {
				return nil, err
			}
			return res.GetResults(), nil
		})
		if err != nil {
			return nil, err
		}
		for _, s := range subnets {
			if s.Id != nil {
				ids = append(ids, *s.Id)
			}
		}
	case "range":
		ranges, err := core.ReadAllPagesUDDI(func(offset, limit int32) ([]uddiipam.Range, error) {
			res, _, err := d.uddi.IPAddressManagementAPI.RangeAPI.
				List(ctx).Tfilter(tfilter).Offset(offset).Limit(limit).Execute()
			if err != nil {
				return nil, err
			}
			return res.GetResults(), nil
		})
		if err != nil {
			return nil, err
		}
		for _, r := range ranges {
			if r.Id != nil {
				ids = append(ids, *r.Id)
			}
		}
	}
	return ids, nil
}
