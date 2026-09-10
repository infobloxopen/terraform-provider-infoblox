package dns

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/infobloxopen/terraform-provider-infoblox/internal/core"
	coremodel "github.com/infobloxopen/terraform-provider-infoblox/internal/core/model/dns"
	coresvc "github.com/infobloxopen/terraform-provider-infoblox/internal/core/service/dns"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/retry"
	internaltypes "github.com/infobloxopen/terraform-provider-infoblox/internal/types"
)

var (
	_ resource.Resource                   = &IPAssociationResource{}
	_ resource.ResourceWithValidateConfig = &IPAssociationResource{}
	_ resource.ResourceWithConfigure      = &IPAssociationResource{}
	_ resource.ResourceWithImportState    = &IPAssociationResource{}
)

func NewIPAssociationResource() resource.Resource {
	return &IPAssociationResource{}
}

// IPAssociationResource manages the DHCP settings of a host record that
// infoblox_record_host owns. It never creates or destroys the record itself.
// Create and Update both write settings onto an existing one, and Delete clears
// them again.
type IPAssociationResource struct {
	backend core.BackendType
	service coresvc.RecordHostService
}

func (r *IPAssociationResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_ip_association"
}

func (r *IPAssociationResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Associates DHCP settings with a host record managed by `infoblox_record_host`, in the NIOS backend.",
		Attributes:          IPAssociationResourceSchemaAttributes,
	}
}

func (r *IPAssociationResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*core.InfobloxClient)
	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Resource Configure Type",
			fmt.Sprintf("Expected *core.InfobloxClient, got: %T.", req.ProviderData),
		)
		return
	}

	if client.NIOS == nil {
		resp.Diagnostics.AddError(
			"Unsupported Backend",
			"infoblox_ip_association is only available for the NIOS backend.",
		)
		return
	}

	r.backend = core.BackendNIOS
	r.service = coresvc.NewRecordHostService(r.backend, client.NIOS, client.UDDI)
}

func (r *IPAssociationResource) retryPolicy(op retry.Operation) retry.Policy {
	return retry.For[coremodel.RecordHost](r.backend, op)
}

func (r *IPAssociationResource) ValidateConfig(ctx context.Context, req resource.ValidateConfigRequest, resp *resource.ValidateConfigResponse) {
	var data IPAssociationModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	nios := flex.ExpandNestedObject[NIOSIPAssociationModel](ctx, data.NIOS, &resp.Diagnostics)
	if nios == nil {
		return
	}

	// DHCP has to lease the address to something.
	if nios.ConfigureForDhcp.IsUnknown() || !nios.ConfigureForDhcp.ValueBool() {
		return
	}
	macEmpty := !nios.MacAddr.IsUnknown() && nios.MacAddr.ValueString() == ""
	duidEmpty := !nios.Duid.IsUnknown() && nios.Duid.ValueString() == ""
	if macEmpty && duidEmpty {
		resp.Diagnostics.AddError(
			"Invalid Configuration",
			"At least one of 'mac' or 'duid' must be configured when 'configure_for_dhcp' is true.",
		)
	}
}

func (r *IPAssociationResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data IPAssociationModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	r.associate(ctx, &data, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *IPAssociationResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data IPAssociationModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	nios := flex.ExpandNestedObject[NIOSIPAssociationModel](ctx, data.NIOS, &resp.Diagnostics)
	if nios == nil {
		return
	}

	host, notFound := r.findHost(ctx, nios, &resp.Diagnostics)
	if notFound {
		// The host record is gone, so there is nothing left to be associated with.
		resp.State.RemoveResource(ctx)
		return
	}
	if host == nil {
		return
	}

	nios.Flatten(host)
	data.NIOS = flex.FlattenNestedObject(ctx, nios, NIOSIPAssociationAttrTypes, &resp.Diagnostics)
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *IPAssociationResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data IPAssociationModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var stateInternalID types.String
	resp.Diagnostics.Append(req.State.GetAttribute(ctx, path.Root("nios").AtName("internal_id"), &stateInternalID)...)
	if resp.Diagnostics.HasError() {
		return
	}
	if nios := flex.ExpandNestedObject[NIOSIPAssociationModel](ctx, data.NIOS, &resp.Diagnostics); nios != nil {
		nios.InternalID = stateInternalID
		data.NIOS = flex.FlattenNestedObject(ctx, nios, NIOSIPAssociationAttrTypes, &resp.Diagnostics)
	}

	r.associate(ctx, &data, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

// Delete clears the DHCP settings but leaves the host record in place: its
// lifecycle belongs to infoblox_record_host.
func (r *IPAssociationResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data IPAssociationModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	nios := flex.ExpandNestedObject[NIOSIPAssociationModel](ctx, data.NIOS, &resp.Diagnostics)
	if nios == nil {
		return
	}

	host, notFound := r.findHost(ctx, nios, &resp.Diagnostics)
	if notFound {
		// Host record already destroyed, so its DHCP settings went with it.
		return
	}
	if host == nil {
		return
	}

	cleared := NIOSIPAssociationModel{
		MacAddr:          internaltypes.NewMACAddressValue(""),
		Duid:             internaltypes.NewDUIDValue(""),
		ConfigureForDhcp: types.BoolValue(false),
		MatchClient:      nios.MatchClient,
	}
	r.update(ctx, nios.RecordHostId.ValueString(), cleared.Expand(host), &resp.Diagnostics)
}

func (r *IPAssociationResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("nios").AtName("record_host_id"), req, resp)
}

// associate reads the host record, writes this association's settings onto it and flattens the result back.
func (r *IPAssociationResource) associate(ctx context.Context, data *IPAssociationModel, diags *diag.Diagnostics) {
	nios := flex.ExpandNestedObject[NIOSIPAssociationModel](ctx, data.NIOS, diags)
	if nios == nil {
		return
	}

	host, notFound := r.findHost(ctx, nios, diags)
	if host == nil {
		if notFound && !diags.HasError() {
			diags.AddError(
				"Not Found",
				"Unable to locate the host record to associate. Create infoblox_record_host first, and on import, import it before this resource.",
			)
		}
		return
	}

	updated := r.update(ctx, nios.RecordHostId.ValueString(), nios.Expand(host), diags)
	if updated == nil {
		return
	}
	nios.Flatten(updated)
	data.NIOS = flex.FlattenNestedObject(ctx, nios, NIOSIPAssociationAttrTypes, diags)
}

func (r *IPAssociationResource) update(ctx context.Context, id string, host *coremodel.RecordHost, diags *diag.Diagnostics) *coremodel.RecordHost {
	var (
		updated  *coremodel.RecordHost
		httpResp *http.Response
	)

	err := retry.Do(ctx, r.retryPolicy(retry.OpUpdate), func(ctx context.Context) (int, error) {
		var apiErr error
		updated, httpResp, apiErr = r.service.Update(ctx, id, host, &core.Options{
			ReturnFields: RecordHostReturnFields,
		})
		if httpResp != nil {
			return httpResp.StatusCode, apiErr
		}
		return 0, apiErr
	})
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to update the host record's DHCP settings: %s", err))
		return nil
	}
	return updated
}

// findHost resolves the host record by reference, falling back to a search on the
// Terraform Internal ID when that reference has gone stale.
func (r *IPAssociationResource) findHost(ctx context.Context, data *NIOSIPAssociationModel, diags *diag.Diagnostics) (*coremodel.RecordHost, bool) {
	var (
		host     *coremodel.RecordHost
		httpResp *http.Response
	)

	err := retry.Do(ctx, r.retryPolicy(retry.OpRead), func(ctx context.Context) (int, error) {
		var apiErr error
		host, httpResp, apiErr = r.service.Read(ctx, data.RecordHostId.ValueString(), &core.Options{
			ReturnFields: RecordHostReturnFields,
		})
		if httpResp != nil {
			return httpResp.StatusCode, apiErr
		}
		return 0, apiErr
	})
	if err == nil {
		internalID, ok := internalIDOf(host)
		if !ok {
			diags.AddError(
				"Missing Internal ID",
				"The host record carries no 'Terraform Internal ID' extensible attribute. Manage it with infoblox_record_host before associating.",
			)
			return nil, false
		}
		data.InternalID = types.StringValue(internalID)
		return host, false
	}
	if httpResp == nil || httpResp.StatusCode != http.StatusNotFound {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read the host record: %s", err))
		return nil, false
	}

	if data.InternalID.IsNull() || data.InternalID.ValueString() == "" {
		return nil, true
	}

	records, _, _, listErr := r.service.List(ctx, &core.ListOptions{
		ReturnFields:  RecordHostReturnFields,
		ExtAttrFilter: map[string]string{flex.TerraformInternalID: data.InternalID.ValueString()},
	})
	if listErr != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to search for the host record by internal ID: %s", listErr))
		return nil, false
	}
	if len(records) == 0 || records[0].Id == nil {
		return nil, true
	}

	// Adopt the updated reference
	data.RecordHostId = types.StringValue(*records[0].Id)
	return records[0], false
}

// internalIDOf reads the Terraform Internal ID extensible attribute off a host
// record, reporting whether it carries one at all.
func internalIDOf(host *coremodel.RecordHost) (string, bool) {
	if host == nil || host.NIOS == nil {
		return "", false
	}
	s, ok := host.NIOS.ExtAttrs[flex.TerraformInternalID].(string)
	return s, ok && s != ""
}
