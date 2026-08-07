package dhcp

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
)

// ValidateFilteroption validates the Filteroption configuration.
func ValidateFilteroption(ctx context.Context, data FilteroptionModel, resp *resource.ValidateConfigResponse) {
	if nios := flex.ExpandNestedObject[NIOSFilteroptionModel](ctx, data.NIOS, &resp.Diagnostics); nios != nil {
		validateFilteroptionNIOSConfig(ctx, nios, resp)
	}
	if uddi := flex.ExpandNestedObject[UDDIFilteroptionModel](ctx, data.UDDI, &resp.Diagnostics); uddi != nil {
		validateFilteroptionUDDIConfig(ctx, uddi, resp)
	}
}

func validateFilteroptionNIOSConfig(ctx context.Context, m *NIOSFilteroptionModel, resp *resource.ValidateConfigResponse) {
}

func validateFilteroptionUDDIConfig(ctx context.Context, m *UDDIFilteroptionModel, resp *resource.ValidateConfigResponse) {
}

func PostFlattenFilteroptionNIOS(ctx context.Context, planned, flattened *NIOSFilteroptionModel, diags *diag.Diagnostics) {
	if planned == nil || flattened == nil {
		return
	}

	// NIOS returns the option list in its own order and includes options the
	// configuration never asked for, so reconcile it against the plan. This also
	// restores a planned empty list when NIOS reports no options at all, which
	// otherwise flattens to null and breaks apply consistency.
	if !planned.OptionList.IsUnknown() {
		if reordered, d := reorderAndFilterDHCPOptions(ctx, planned.OptionList, flattened.OptionList); !d.HasError() {
			flattened.OptionList = reordered.(basetypes.ListValue)
		}
	}
}

// reorderAndFilterDHCPOptions rebuilds the DHCP option list returned by NIOS so it
// follows the order and membership of the plan. Options are matched on "name" and
// fall back to "num", mirroring how NIOS identifies a DHCP option.
func reorderAndFilterDHCPOptions(ctx context.Context, planValue, stateValue attr.Value) (attr.Value, *diag.Diagnostics) {
	var diags diag.Diagnostics

	const (
		primaryKey   = "name"
		secondaryKey = "num"
	)

	if planValue.IsNull() || planValue.IsUnknown() {
		return stateValue, &diags
	}
	if stateValue.IsNull() || stateValue.IsUnknown() {
		return planValue, &diags
	}

	planList, ok := planValue.(basetypes.ListValue)
	if !ok {
		diags.AddError("Type Error", "planValue must be a basetypes.ListValue")
		return stateValue, &diags
	}
	stateList, ok := stateValue.(basetypes.ListValue)
	if !ok {
		diags.AddError("Type Error", "stateValue must be a basetypes.ListValue")
		return planValue, &diags
	}

	// Index the options NIOS returned by both keys so either can resolve a match.
	nameToState := make(map[string]attr.Value)
	numToState := make(map[int64]attr.Value)

	for _, elem := range stateList.Elements() {
		obj, ok := elem.(basetypes.ObjectValue)
		if !ok {
			continue
		}
		attrs := obj.Attributes()

		if nameAttr, has := attrs[primaryKey]; has && nameAttr != nil && !nameAttr.IsNull() && !nameAttr.IsUnknown() {
			if strVal, ok := nameAttr.(basetypes.StringValue); ok {
				nameToState[strVal.ValueString()] = elem
			}
		}
		if numAttr, has := attrs[secondaryKey]; has && numAttr != nil && !numAttr.IsNull() && !numAttr.IsUnknown() {
			if intVal, ok := numAttr.(basetypes.Int64Value); ok {
				numToState[intVal.ValueInt64()] = elem
			}
		}
	}

	reordered := make([]attr.Value, 0, len(planList.Elements()))
	for _, planElem := range planList.Elements() {
		planObj, ok := planElem.(basetypes.ObjectValue)
		if !ok {
			reordered = append(reordered, planElem)
			continue
		}
		planAttrs := planObj.Attributes()

		var matchedState attr.Value

		if nameAttr, has := planAttrs[primaryKey]; has && nameAttr != nil && !nameAttr.IsNull() && !nameAttr.IsUnknown() {
			if strVal, ok := nameAttr.(basetypes.StringValue); ok {
				if s, exists := nameToState[strVal.ValueString()]; exists {
					matchedState = s
				}
			}
		}

		if matchedState == nil {
			if numAttr, has := planAttrs[secondaryKey]; has && numAttr != nil && !numAttr.IsNull() && !numAttr.IsUnknown() {
				if intVal, ok := numAttr.(basetypes.Int64Value); ok {
					if s, exists := numToState[intVal.ValueInt64()]; exists {
						matchedState = s
					}
				}
			}
		}

		// Fall back to the planned option when NIOS returned no counterpart.
		if matchedState != nil {
			reordered = append(reordered, matchedState)
		} else {
			reordered = append(reordered, planElem)
		}
	}

	newList, d := basetypes.NewListValue(planList.ElementType(ctx), reordered)
	diags.Append(d...)

	return newList, &diags
}
