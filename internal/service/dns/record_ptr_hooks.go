package dns

import (
	"context"
	"net"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	niosdns "github.com/infobloxopen/infoblox-nios-go-client/dns"

	coremodel "github.com/infobloxopen/terraform-provider-infoblox/internal/core/model/dns"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/dynamicallocation"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
)

// ---- PostExpand NIOS hook ----
// Clears empty-string fields that must be nil (not "") for NIOS PTR create/update.
// The generated Expand uses ExpandStringPointerNullAsEmpty which sends "" for null
// optional strings; NIOS rejects empty values for name/view.
func PostExpandRecordPtrNIOS(ctx context.Context, ext *coremodel.NIOSRecordPtrExt, diags *diag.Diagnostics) *coremodel.NIOSRecordPtrExt {
	if ext == nil {
		return nil
	}
	// Clear fields that must NOT be sent as empty string to NIOS
	if ext.Name != nil && *ext.Name == "" {
		ext.Name = nil
	}
	if ext.DdnsPrincipal != nil && *ext.DdnsPrincipal == "" {
		ext.DdnsPrincipal = nil
	}
	return ext
}

// ---- UDDI Rdata custom schema (dname) ----

// UDDIRecordPtrRdataModel holds the PTR record rdata sub-fields.
type UDDIRecordPtrRdataModel struct {
	DName types.String `tfsdk:"dname"`
}

var UDDIRecordPtrRdataAttrTypes = map[string]attr.Type{
	"dname": types.StringType,
}

var UDDIRecordPtrRdataResourceSchemaAttributes = map[string]schema.Attribute{
	"dname": schema.StringAttribute{
		Required:            true,
		MarkdownDescription: "A domain name which points to some location in the domain name space. Can be an absolute or relative domain name and may include UTF-8.",
	},
}

func ExpandUDDIRecordPtrRdata(ctx context.Context, o types.Object, diags *diag.Diagnostics) map[string]any {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m UDDIRecordPtrRdataModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return map[string]any{
		"dname": flex.ExpandString(m.DName),
	}
}

func FlattenUDDIRecordPtrRdata(ctx context.Context, from map[string]any, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(UDDIRecordPtrRdataAttrTypes)
	}
	m := UDDIRecordPtrRdataModel{
		DName: flex.FlattenStringPointer(flex.RDataStringPtr(from["dname"])),
	}
	obj, d := types.ObjectValueFrom(ctx, UDDIRecordPtrRdataAttrTypes, m)
	diags.Append(d...)
	return obj
}

// ---- UDDI Options custom schema (address) ----
type UDDIRecordPtrOptionsModel struct {
	Address types.String `tfsdk:"address"`
}

var UDDIRecordPtrOptionsAttrTypes = map[string]attr.Type{
	"address": types.StringType,
}

var UDDIRecordPtrOptionsResourceSchemaAttributes = map[string]schema.Attribute{
	"address": schema.StringAttribute{
		Optional: true,
		Computed: true,
		PlanModifiers: []planmodifier.String{
			stringplanmodifier.UseStateForUnknown(),
		},
		MarkdownDescription: "For GET operation it contains the IPv4 or IPv6 address represented by the PTR record.\n\n" +
			"For POST and PATCH operations it can be used to create/update a PTR record based on the IP address " +
			"it represents. In this case, in addition to the _address_ in the options field, the _view_ field must also be specified.",
	},
}

func ExpandUDDIRecordPtrOptions(ctx context.Context, o types.Object, diags *diag.Diagnostics) map[string]any {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m UDDIRecordPtrOptionsModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	to := map[string]any{}
	if !m.Address.IsNull() && !m.Address.IsUnknown() {
		to["address"] = flex.ExpandString(m.Address)
	}
	return to
}

func FlattenUDDIRecordPtrOptions(ctx context.Context, from map[string]any, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(UDDIRecordPtrOptionsAttrTypes)
	}
	m := UDDIRecordPtrOptionsModel{
		Address: flex.FlattenStringPointer(flex.RDataStringPtr(from["address"])),
	}
	obj, d := types.ObjectValueFrom(ctx, UDDIRecordPtrOptionsAttrTypes, m)
	diags.Append(d...)
	return obj
}

// ---- ValidateConfig ----

// ValidateRecordPtr validates the RecordPtr configuration.
func ValidateRecordPtr(ctx context.Context, data RecordPtrModel, resp *resource.ValidateConfigResponse) {
	if nios := flex.ExpandNestedObject[NIOSRecordPtrModel](ctx, data.NIOS, &resp.Diagnostics); nios != nil {
		validateRecordPtrNIOSConfig(ctx, nios, resp)
	}
	if uddi := flex.ExpandNestedObject[UDDIRecordPtrModel](ctx, data.UDDI, &resp.Diagnostics); uddi != nil {
		validateRecordPtrUDDIConfig(ctx, uddi, resp)
	}
}

func validateRecordPtrNIOSConfig(ctx context.Context, m *NIOSRecordPtrModel, resp *resource.ValidateConfigResponse) {
}

// BuildRecordPtrFuncCall builds the NIOS next_available_ip FuncCall for a PTR record.
// It detects IPv4 vs IPv6 from the network CIDR (colons indicate IPv6) and routes
// the allocation to the appropriate WAPI attribute name and object type.
func BuildRecordPtrFuncCall(ctx context.Context, data types.Object, diags *diag.Diagnostics) *niosdns.FuncCall {
	if data.IsNull() || data.IsUnknown() {
		return nil
	}

	var m dynamicallocation.NextAvailableIpModel
	diags.Append(data.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}

	ip, _, _ := net.ParseCIDR(m.Network.ValueString())
	if ip != nil && ip.To4() == nil {
		return m.FuncCall(ctx, "Ipv6addr", "ipv6network", diags)
	}
	return m.FuncCall(ctx, "Ipv4addr", "network", diags)
}

func validateRecordPtrUDDIConfig(ctx context.Context, m *UDDIRecordPtrModel, resp *resource.ValidateConfigResponse) {
}
