package rpz

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/list"
	listschema "github.com/hashicorp/terraform-plugin-framework/list/schema"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/infobloxopen/terraform-provider-infoblox/internal/core"
	coremodel "github.com/infobloxopen/terraform-provider-infoblox/internal/core/model/rpz"
	coresvc "github.com/infobloxopen/terraform-provider-infoblox/internal/core/service/rpz"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/validator"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ list.ListResource                   = &RecordRpzNaptrList{}
	_ list.ListResourceWithConfigure      = &RecordRpzNaptrList{}
	_ list.ListResourceWithValidateConfig = &RecordRpzNaptrList{}
)

func NewRecordRpzNaptrList() list.ListResource {
	return &RecordRpzNaptrList{}
}

type RecordRpzNaptrList struct {
	backend core.BackendType
	service coresvc.RecordRpzNaptrService
}

type RecordRpzNaptrListModel struct {
	Filters        types.Map `tfsdk:"filters"`
	ExtAttrFilters types.Map `tfsdk:"ext_attr_filters"`
}

func (l *RecordRpzNaptrList) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_record_rpz_naptr"
}

func (l *RecordRpzNaptrList) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*core.InfobloxClient)
	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected List Resource Configure Type",
			fmt.Sprintf("Expected *core.InfobloxClient, got: %T.", req.ProviderData),
		)
		return
	}

	if client.NIOS != nil {
		l.backend = core.BackendNIOS
	} else {
		l.backend = core.BackendUDDI
	}

	l.service = coresvc.NewRecordRpzNaptrService(l.backend, client.NIOS, client.UDDI)
}

func (l *RecordRpzNaptrList) ListResourceConfigSchema(_ context.Context, _ list.ListResourceSchemaRequest, resp *list.ListResourceSchemaResponse) {
	resp.Schema = listschema.Schema{
		MarkdownDescription: "Retrieves a list of Infoblox RecordRpzNaptr from the NIOS backend.",
		Attributes: map[string]listschema.Attribute{
			"filters": listschema.MapAttribute{
				MarkdownDescription: "Filters are used to return a more specific list of results. Filters can be used to match resources by specific attributes (e.g. name, view). If multiple filters are specified, only resources that match all of them are returned.",
				ElementType:         types.StringType,
				Optional:            true,
			},
			"ext_attr_filters": listschema.MapAttribute{
				MarkdownDescription: "Extensible Attribute Filters are used to filter results by NIOS extensible attributes. Only applicable for the NIOS backend.",
				ElementType:         types.StringType,
				Optional:            true,
			},
		},
	}
}

func (l *RecordRpzNaptrList) ValidateListResourceConfig(ctx context.Context, req list.ValidateConfigRequest, resp *list.ValidateConfigResponse) {
	var data RecordRpzNaptrListModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	validator.ValidateListFilters(l.backend, data.ExtAttrFilters, types.MapNull(types.StringType), &resp.Diagnostics)
}

func (l *RecordRpzNaptrList) List(ctx context.Context, req list.ListRequest, stream *list.ListResultsStream) {
	var data RecordRpzNaptrListModel

	diags := req.Config.Get(ctx, &data)
	if diags.HasError() {
		stream.Results = list.ListResultsStreamDiagnostics(diags)
		return
	}

	requestLimit := int32(req.Limit)
	tflog.Info(ctx, fmt.Sprintf("infoblox_record_rpz_naptr list: req.Limit=%d backend=%s includeResource=%t",
		req.Limit, l.backend, req.IncludeResource))

	opts := &core.ListOptions{
		Filters:       flex.ExpandMapString(ctx, data.Filters, &diags),
		ExtAttrFilter: flex.ExpandMapString(ctx, data.ExtAttrFilters, &diags),
		ReturnFields:  RecordRpzNaptrReturnFields,
		Paging:        1,
	}
	if diags.HasError() {
		stream.Results = list.ListResultsStreamDiagnostics(diags)
		return
	}

	var records []*coremodel.RecordRpzNaptr
	var err error
	var totalFetched int32
	pageCount := 0

	switch l.backend {
	case core.BackendNIOS:
		records, err = core.ReadAllPagesNIOS(
			func(pageID string) ([]*coremodel.RecordRpzNaptr, string, error) {
				// Shrink page size so we never over-fetch past the caller's limit.
				pageSize := min(requestLimit-totalFetched, core.DefaultListLimit)
				pageCount++

				opts.PageID = pageID
				opts.MaxResults = pageSize
				recs, _, nextPageID, e := l.service.List(ctx, opts)
				if e != nil {
					return nil, "", e
				}
				totalFetched += int32(len(recs))
				tflog.Info(ctx, fmt.Sprintf("NIOS list page %d: requested=%d got=%d nextPageID=%q",
					pageCount, pageSize, len(recs), nextPageID))
				// Stop early once the cap is reached.
				if totalFetched >= requestLimit {
					return recs, "", nil
				}
				return recs, nextPageID, nil
			})

	case core.BackendUDDI:
		records, err = core.ReadAllPagesUDDI(
			func(offset, _ int32) ([]*coremodel.RecordRpzNaptr, error) {
				// Once the cap is reached, return an empty page
				remaining := requestLimit - totalFetched
				if remaining <= 0 {
					return nil, nil
				}
				// Shrink page size so we never over-fetch past the caller's limit.
				pageSize := min(remaining, core.DefaultListLimit)

				pageCount++
				opts.Offset = offset
				opts.Limit = pageSize
				recs, _, _, e := l.service.List(ctx, opts)
				if e != nil {
					return nil, e
				}
				totalFetched += int32(len(recs))
				tflog.Info(ctx, fmt.Sprintf("UDDI list page %d: offset=%d requested=%d got=%d",
					pageCount, offset, pageSize, len(recs)))
				return recs, nil
			})
	}

	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to list RecordRpzNaptr records: %s", err))
		stream.Results = list.ListResultsStreamDiagnostics(diags)
		return
	}

	tflog.Info(ctx, fmt.Sprintf("Retrieved %d total results (limit=%d)", len(records), requestLimit))

	stream.Results = func(push func(list.ListResult) bool) {
		for _, item := range records {
			result := req.NewListResult(ctx)

			// Set the identity on each result.
			result.Diagnostics.Append(
				result.Identity.SetAttribute(ctx, path.Root("id"), &item.Id)...,
			)
			if result.Diagnostics.HasError() {
				if !push(result) {
					return
				}
				continue
			}

			// By default, list only returns the identity. If IncludeResource is true,
			// the full resource is flattened and set on result.Resource.
			if req.IncludeResource {
				model := &RecordRpzNaptrModel{}
				model.Flatten(ctx, item, &result.Diagnostics)
				if !result.Diagnostics.HasError() {
					result.Diagnostics.Append(result.Resource.Set(ctx, model)...)
				}
				if result.Diagnostics.HasError() {
					if !push(result) {
						return
					}
					continue
				}
			}

			if !push(result) {
				return
			}
		}
	}
}
