package dns

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listdefault"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/utils"
)

// ValidateRecordSvcb validates the RecordSvcb configuration.
func ValidateRecordSvcb(ctx context.Context, data RecordSvcbModel, resp *resource.ValidateConfigResponse) {
	if uddi := flex.ExpandNestedObject[UDDIRecordSvcbModel](ctx, data.UDDI, &resp.Diagnostics); uddi != nil {
		validateRecordSvcbUDDIConfig(ctx, uddi, resp)
	}
}

func validateRecordSvcbUDDIConfig(ctx context.Context, m *UDDIRecordSvcbModel, resp *resource.ValidateConfigResponse) {
	rdata := flex.ExpandNestedObject[UDDIRecordSvcbRdataModel](ctx, m.Rdata, &resp.Diagnostics)
	if rdata == nil {
		return
	}

	svcParamsProvided := !rdata.SvcParams.IsNull() && !rdata.SvcParams.IsUnknown() && len(rdata.SvcParams.Elements()) > 0
	priorityKnown := !rdata.Priority.IsNull() && !rdata.Priority.IsUnknown()

	if svcParamsProvided && priorityKnown && rdata.Priority.ValueInt64() <= 0 {
		resp.Diagnostics.AddAttributeError(
			path.Root("uddi").AtName("rdata").AtName("svc_params"),
			"Invalid Configuration",
			"'svc_params' can only be specified when 'priority' is greater than 0.",
		)
	}

	if svcParamsProvided {
		var params []UDDIRecordSvcbSvcParamModel
		resp.Diagnostics.Append(rdata.SvcParams.ElementsAs(ctx, &params, false)...)
		for i, p := range params {
			if p.Key.IsNull() || p.Key.IsUnknown() {
				continue
			}
			if p.Key.ValueString() == "ohttp" {
				continue
			}
			if p.Value.IsNull() || p.Value.IsUnknown() || p.Value.ValueString() == "" {
				resp.Diagnostics.AddAttributeError(
					path.Root("uddi").AtName("rdata").AtName("svc_params").AtListIndex(i).AtName("value"),
					"Invalid Configuration",
					fmt.Sprintf("'value' is required for svc_params entry with key %q.", p.Key.ValueString()),
				)
			}
		}
	}
}

func PostFlattenRecordSvcbUDDI(ctx context.Context, planned, flattened *UDDIRecordSvcbModel, diags *diag.Diagnostics) {
	if !planned.Rdata.IsNull() {
		if result, d := utils.CopyFieldFromPlanToRespObject(ctx, planned.Rdata, flattened.Rdata, "priority"); !d.HasError() {
			flattened.Rdata = result.(basetypes.ObjectValue)
		}
	}
}

// UDDIRecordSvcbSvcParamModel represents a single service parameter key-value pair.
type UDDIRecordSvcbSvcParamModel struct {
	Key   types.String `tfsdk:"key"`
	Value types.String `tfsdk:"value"`
}

var UDDIRecordSvcbSvcParamAttrTypes = map[string]attr.Type{
	"key":   types.StringType,
	"value": types.StringType,
}

var UDDIRecordSvcbSvcParamResourceSchemaAttributes = map[string]schema.Attribute{
	"key": schema.StringAttribute{
		Required:            true,
		MarkdownDescription: "The service parameter key (e.g. \"port\", \"ipv4hint\", \"ipv6hint\", \"ech\", \"alpn\").",
	},
	"value": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The service parameter value.",
	},
}

type UDDIRecordSvcbRdataModel struct {
	Priority   types.Int64  `tfsdk:"priority"`
	SvcParams  types.List   `tfsdk:"svc_params"`
	TargetName types.String `tfsdk:"target_name"`
}

var UDDIRecordSvcbRdataAttrTypes = map[string]attr.Type{
	"priority":    types.Int64Type,
	"svc_params":  types.ListType{ElemType: types.ObjectType{AttrTypes: UDDIRecordSvcbSvcParamAttrTypes}},
	"target_name": types.StringType,
}

var UDDIRecordSvcbRdataResourceSchemaAttributes = map[string]schema.Attribute{
	"priority": schema.Int64Attribute{
		Optional:            true,
		Computed:            true,
		Default:             int64default.StaticInt64(0),
		MarkdownDescription: "An unsigned 16-bit integer in the range 0 to 65535 that indicates the priority of the SVCB record. Lower values are preferred.",
	},
	"svc_params": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: UDDIRecordSvcbSvcParamResourceSchemaAttributes,
		},
		Optional: true,
		Computed: true,
		Default:  listdefault.StaticValue(types.ListValueMust(types.ObjectType{AttrTypes: UDDIRecordSvcbSvcParamAttrTypes}, []attr.Value{})),
		MarkdownDescription: "A list of service parameters for the SVCB record. Each entry is a key-value pair " +
			"(e.g. port, ipv4hint, ipv6hint, ech, alpn).",
	},
	"target_name": schema.StringAttribute{
		Required:            true,
		MarkdownDescription: "The domain name of the SVCB target. Use \".\" to indicate the service is located at the owner name itself.",
	},
}

func ExpandUDDIRecordSvcbRdata(ctx context.Context, o types.Object, diags *diag.Diagnostics) map[string]any {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m UDDIRecordSvcbRdataModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}

	rdata := map[string]any{
		"priority":    flex.ExpandInt64(m.Priority),
		"target_name": flex.ExpandString(m.TargetName),
	}

	if !m.SvcParams.IsNull() && !m.SvcParams.IsUnknown() {
		var params []UDDIRecordSvcbSvcParamModel
		diags.Append(m.SvcParams.ElementsAs(ctx, &params, false)...)
		if !diags.HasError() {
			expanded := make([]map[string]any, 0, len(params))
			for _, p := range params {
				expanded = append(expanded, map[string]any{
					"key":   flex.ExpandString(p.Key),
					"value": flex.ExpandString(p.Value),
				})
			}
			rdata["svc_params"] = expanded
		}
	}

	return rdata
}

func FlattenUDDIRecordSvcbRdata(ctx context.Context, from map[string]any, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(UDDIRecordSvcbRdataAttrTypes)
	}

	svcParams := flattenSvcParams(ctx, from["svc_params"], diags)

	m := UDDIRecordSvcbRdataModel{
		Priority:   flex.FlattenInt64Pointer(flex.RDataInt64Ptr(from["priority"])),
		SvcParams:  svcParams,
		TargetName: flex.FlattenStringPointer(flex.RDataStringPtr(from["target_name"])),
	}
	obj, d := types.ObjectValueFrom(ctx, UDDIRecordSvcbRdataAttrTypes, m)
	diags.Append(d...)
	return obj
}

func flattenSvcParams(ctx context.Context, raw any, diags *diag.Diagnostics) types.List {
	nullList := types.ListValueMust(types.ObjectType{AttrTypes: UDDIRecordSvcbSvcParamAttrTypes}, []attr.Value{})
	if raw == nil {
		return nullList
	}

	items, ok := raw.([]any)
	if !ok || len(items) == 0 {
		return nullList
	}

	elems := make([]attr.Value, 0, len(items))
	for _, item := range items {
		m, ok := item.(map[string]any)
		if !ok {
			continue
		}
		param := UDDIRecordSvcbSvcParamModel{
			Key:   flex.FlattenStringPointer(flex.RDataStringPtr(m["key"])),
			Value: flex.FlattenStringPointer(flex.RDataStringPtr(m["value"])),
		}
		obj, d := types.ObjectValueFrom(ctx, UDDIRecordSvcbSvcParamAttrTypes, param)
		diags.Append(d...)
		if d.HasError() {
			return nullList
		}
		elems = append(elems, obj)
	}

	list, d := types.ListValue(types.ObjectType{AttrTypes: UDDIRecordSvcbSvcParamAttrTypes}, elems)
	diags.Append(d...)
	return list
}
