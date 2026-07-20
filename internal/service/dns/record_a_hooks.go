package dns

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	niosdns "github.com/infobloxopen/infoblox-nios-go-client/dns"
	coremodel "github.com/infobloxopen/terraform-provider-infoblox/internal/core/model/dns"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/dynamicallocation"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
)

// boolRecordAOptionKeys are the only valid option keys for an A record, and they must be boolean strings.
var boolRecordAOptionKeys = []string{"create_ptr", "check_rmz"}

// ValidateRecordA validates the RecordA configuration.
func ValidateRecordA(ctx context.Context, data RecordAModel, resp *resource.ValidateConfigResponse) {
	if nios := flex.ExpandNestedObject[NIOSRecordAModel](ctx, data.NIOS, &resp.Diagnostics); nios != nil {
		validateRecordANIOSConfig(ctx, nios, resp)
	}
	if uddi := flex.ExpandNestedObject[UDDIRecordAModel](ctx, data.UDDI, &resp.Diagnostics); uddi != nil {
		validateRecordAUDDIConfig(ctx, uddi, resp)
	}
}

func validateRecordANIOSConfig(ctx context.Context, m *NIOSRecordAModel, resp *resource.ValidateConfigResponse) {
}

func validateRecordAUDDIConfig(ctx context.Context, m *UDDIRecordAModel, resp *resource.ValidateConfigResponse) {
	// rdata: the address subfield is the only allowed and required field.
	if !m.Rdata.IsNull() && !m.Rdata.IsUnknown() {
		elems := m.Rdata.Elements()
		address, present := elems["address"]
		if !present {
			resp.Diagnostics.AddAttributeError(
				path.Root("uddi").AtName("rdata"),
				"Missing Required Subfield",
				"The `address` subfield is required in `rdata` for an A record.",
			)
		} else if addrStr, ok := address.(types.String); ok && !addrStr.IsUnknown() {
			if addrStr.IsNull() || strings.TrimSpace(addrStr.ValueString()) == "" {
				resp.Diagnostics.AddAttributeError(
					path.Root("uddi").AtName("rdata").AtMapKey("address"),
					"Invalid Subfield Value",
					"The `address` subfield in `rdata` must be a non-empty IPv4 address for an A record.",
				)
			}
		}
	}

	// options: only create_ptr and check_rmz are valid, and values must be boolean strings.
	if !m.Options.IsNull() && !m.Options.IsUnknown() {
		allowed := make(map[string]struct{}, len(boolRecordAOptionKeys))
		for _, k := range boolRecordAOptionKeys {
			allowed[k] = struct{}{}
		}
		for key, val := range m.Options.Elements() {
			if _, valid := allowed[key]; !valid {
				resp.Diagnostics.AddAttributeError(
					path.Root("uddi").AtName("options").AtMapKey(key),
					"Invalid Option",
					fmt.Sprintf("`%s` is not a valid option for an A record. Valid options are: create_ptr, check_rmz.", key),
				)
				continue
			}
			if valStr, ok := val.(types.String); ok && !valStr.IsUnknown() && !valStr.IsNull() {
				if _, err := strconv.ParseBool(valStr.ValueString()); err != nil {
					resp.Diagnostics.AddAttributeError(
						path.Root("uddi").AtName("options").AtMapKey(key),
						"Invalid Option Value",
						fmt.Sprintf("`%s` must be a boolean value (\"true\" or \"false\"), got %q.", key, valStr.ValueString()),
					)
				}
			}
		}
	}
}

func BuildRecordAFuncCall(ctx context.Context, data types.Object, diags *diag.Diagnostics) *niosdns.FuncCall {
	if data.IsNull() || data.IsUnknown() {
		return nil
	}

	var m dynamicallocation.NextAvailableIpModel
	diags.Append(data.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}

	return m.FuncCall(ctx, "Ipv4addr", "network", diags)
}

func PostExpandRecordAUDDI(ctx context.Context, ext *coremodel.UDDIRecordAExt, diags *diag.Diagnostics) *coremodel.UDDIRecordAExt {
	if ext == nil {
		return ext
	}
	if ext.Options != nil {
		for _, k := range boolRecordAOptionKeys {
			if v, ok := ext.Options[k].(string); ok {
				if b, err := strconv.ParseBool(v); err == nil {
					ext.Options[k] = b
				}
			}
		}
	}
	return ext
}

func PostFlattenRecordAUDDI(ctx context.Context, planned, flattened *UDDIRecordAModel, diags *diag.Diagnostics) {
	if flattened == nil {
		return
	}
	if planned != nil {
		flattened.Options = planned.Options
	} else {
		flattened.Options = types.MapNull(types.StringType)
	}
}
