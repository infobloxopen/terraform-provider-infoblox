package flex

import (
	"context"
	"fmt"
	"maps"

	"github.com/hashicorp/go-uuid"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// TerraformInternalID is the extensible attribute key used to store TF generated UUID
// for uniquely identifying resources across imports and drift detection.
const TerraformInternalID = "Terraform Internal ID"

// AssociateInternalIDKey is the private state key used during import flow
// to signal that the resource needs a Terraform Internal ID to be associated.
const AssociateInternalIDKey = "associate_internal_id"

// SetInternalID adds a Terraform Internal ID to the ext_attrs map.
// Used during Create to associate a unique identifier with the resource.
func SetInternalID(ctx context.Context, extAttrs types.Map, diags *diag.Diagnostics) types.Map {
	elements := make(map[string]attr.Value)
	if !extAttrs.IsNull() && !extAttrs.IsUnknown() {
		elements = extAttrs.Elements()
	}

	internalID, err := uuid.GenerateUUID()
	if err != nil {
		diags.AddError("Error generating UUID", fmt.Sprintf("Unable to generate internal ID: %s", err))
		return extAttrs
	}

	elements[TerraformInternalID] = types.StringValue(internalID)
	newMap, d := types.MapValue(types.StringType, elements)
	diags.Append(d...)
	return newMap
}

// MergeEAs merges plan ext_attrs with state ext_attrs_all.
// Ensures inherited EAs and TF Internal ID are preserved during Update.
func MergeEAs(planExtAttrs, stateExtAttrsAll types.Map) types.Map {
	if stateExtAttrsAll.IsNull() || stateExtAttrsAll.IsUnknown() {
		return planExtAttrs
	}

	merged := make(map[string]attr.Value)
	if !planExtAttrs.IsNull() && !planExtAttrs.IsUnknown() {
		maps.Copy(merged, planExtAttrs.Elements())
	}

	// Add state ext_attrs_all entries not in plan (inherited + TF ID)
	for k, v := range stateExtAttrsAll.Elements() {
		if _, exists := merged[k]; !exists {
			merged[k] = v
		}
	}

	if len(merged) == 0 {
		return types.MapNull(types.StringType)
	}

	result, _ := types.MapValue(types.StringType, merged)
	return result
}

// FlattenEAs splits API response ext_attrs (already stringified) into:
// - ext_attrs: user-defined values (what user specified in plan)
// - ext_attrs_all: TF Internal ID + inherited (computed, read-only)
func FlattenEAs(
	planExtAttrs types.Map,
	respExtAttrs map[string]any,
) (extAttrs, extAttrsAll types.Map) {
	if len(respExtAttrs) == 0 {
		return types.MapNull(types.StringType), types.MapNull(types.StringType)
	}

	planKeys := make(map[string]bool)
	if !planExtAttrs.IsNull() && !planExtAttrs.IsUnknown() {
		for k := range planExtAttrs.Elements() {
			planKeys[k] = true
		}
	}

	userEAs := make(map[string]attr.Value)
	computedEAs := make(map[string]attr.Value)

	for key, val := range respExtAttrs {
		strVal, _ := val.(string) // val will already be stringified(core is handling it)
		attrVal := types.StringValue(strVal)

		// TF Internal ID always goes to ext_attrs_all
		if key == TerraformInternalID {
			computedEAs[key] = attrVal
			continue
		}

		// If user specified this EA in plan, it goes to ext_attrs
		// Otherwise it's inherited(not overridden) and goes to ext_attrs_all
		if planKeys[key] {
			userEAs[key] = attrVal
		} else {
			computedEAs[key] = attrVal
		}
	}

	if len(userEAs) > 0 {
		extAttrs, _ = types.MapValue(types.StringType, userEAs)
	} else {
		extAttrs = types.MapNull(types.StringType)
	}

	if len(computedEAs) > 0 {
		extAttrsAll, _ = types.MapValue(types.StringType, computedEAs)
	} else {
		extAttrsAll = types.MapNull(types.StringType)
	}

	return extAttrs, extAttrsAll
}
