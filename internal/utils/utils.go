package utils

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	datasourceschema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	resourceschema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

const (
	NaiveDatetimeLayout string = "2006-01-02T15:04:05"
)

// DataSourceResultAttributes combines resource attributes into datasource attributes.
// Converts all fields to computed for datasource use.
func DataSourceResultAttributes(attrs map[string]resourceschema.Attribute) map[string]datasourceschema.Attribute {
	result := make(map[string]datasourceschema.Attribute)
	for k, v := range attrs {
		result[k] = toDataSourceAttribute(v)
	}
	return result
}

func toDataSourceAttribute(val resourceschema.Attribute) datasourceschema.Attribute {
	switch a := val.(type) {
	case resourceschema.StringAttribute:
		return datasourceschema.StringAttribute{Computed: true, MarkdownDescription: a.MarkdownDescription}
	case resourceschema.BoolAttribute:
		return datasourceschema.BoolAttribute{Computed: true, MarkdownDescription: a.MarkdownDescription}
	case resourceschema.Int32Attribute:
		return datasourceschema.Int32Attribute{Computed: true, MarkdownDescription: a.MarkdownDescription}
	case resourceschema.Int64Attribute:
		return datasourceschema.Int64Attribute{Computed: true, MarkdownDescription: a.MarkdownDescription}
	case resourceschema.Float64Attribute:
		return datasourceschema.Float64Attribute{Computed: true, MarkdownDescription: a.MarkdownDescription}
	case resourceschema.MapAttribute:
		return datasourceschema.MapAttribute{Computed: true, ElementType: a.ElementType, MarkdownDescription: a.MarkdownDescription}
	case resourceschema.ListAttribute:
		return datasourceschema.ListAttribute{Computed: true, ElementType: a.ElementType, MarkdownDescription: a.MarkdownDescription}
	case resourceschema.SetAttribute:
		return datasourceschema.SetAttribute{Computed: true, ElementType: a.ElementType, MarkdownDescription: a.MarkdownDescription}
	case resourceschema.ListNestedAttribute:
		return datasourceschema.ListNestedAttribute{
			Computed:            true,
			MarkdownDescription: a.MarkdownDescription,
			NestedObject: datasourceschema.NestedAttributeObject{
				Attributes: nestedAttrsToDataSource(a.NestedObject.Attributes),
			},
		}
	case resourceschema.SetNestedAttribute:
		return datasourceschema.SetNestedAttribute{
			Computed:            true,
			MarkdownDescription: a.MarkdownDescription,
			NestedObject: datasourceschema.NestedAttributeObject{
				Attributes: nestedAttrsToDataSource(a.NestedObject.Attributes),
			},
		}
	case resourceschema.SingleNestedAttribute:
		return datasourceschema.SingleNestedAttribute{
			Computed:            true,
			MarkdownDescription: a.MarkdownDescription,
			Attributes:          nestedAttrsToDataSource(a.Attributes),
		}
	case resourceschema.MapNestedAttribute:
		return datasourceschema.MapNestedAttribute{
			Computed:            true,
			MarkdownDescription: a.MarkdownDescription,
			NestedObject: datasourceschema.NestedAttributeObject{
				Attributes: nestedAttrsToDataSource(a.NestedObject.Attributes),
			},
		}
	default:
		return nil
	}
}

func nestedAttrsToDataSource(attrs map[string]resourceschema.Attribute) map[string]datasourceschema.Attribute {
	result := make(map[string]datasourceschema.Attribute)
	for k, v := range attrs {
		result[k] = toDataSourceAttribute(v)
	}
	return result
}

// ReorderAndFilterNestedListResponse reorders the response list to match the
// order of the plan list, keyed by the given primary key, so that Terraform does
// not detect spurious ordering-only diffs on nested lists.
func ReorderAndFilterNestedListResponse(
	ctx context.Context,
	planValue attr.Value,
	stateValue attr.Value,
	primaryKey string,
) (attr.Value, *diag.Diagnostics) {

	var diags diag.Diagnostics

	if planValue.IsNull() || planValue.IsUnknown() {
		return stateValue, &diags
	}
	if stateValue.IsNull() || stateValue.IsUnknown() {
		return planValue, &diags
	}

	planList, ok := planValue.(basetypes.ListValue)
	if !ok {
		diags.AddError("Type Error", "planValue must be a ListValue")
		return stateValue, &diags
	}
	stateList, ok := stateValue.(basetypes.ListValue)
	if !ok {
		diags.AddError("Type Error", "stateValue must be a ListValue")
		return planValue, &diags
	}

	// Convert state list into a lookup by primary key
	stateMap := make(map[string]attr.Value)
	for _, elem := range stateList.Elements() {
		obj := elem.(basetypes.ObjectValue)
		keyAttr, ok := obj.Attributes()[primaryKey]
		if !ok {
			diags.AddError("Missing Primary Key", fmt.Sprintf("State object missing primary key: %s", primaryKey))
			continue
		}
		if keyAttr.IsNull() || keyAttr.IsUnknown() {
			continue
		}
		key := keyAttr.(basetypes.StringValue).ValueString()
		stateMap[key] = elem
	}

	// Rebuild state list in the same order as plan
	var reordered []attr.Value
	for _, elem := range planList.Elements() {
		obj := elem.(basetypes.ObjectValue)
		keyAttr := obj.Attributes()[primaryKey]
		if keyAttr.IsNull() || keyAttr.IsUnknown() {
			continue
		}
		key := keyAttr.(basetypes.StringValue).ValueString()

		// Use existing state object if found, else use planned object
		if stateObj, exists := stateMap[key]; exists {
			reordered = append(reordered, stateObj)
		} else {
			reordered = append(reordered, elem)
		}
	}

	// Build new ListValue
	newList, d := basetypes.NewListValue(planList.ElementType(ctx), reordered)
	diags.Append(d...)

	return newList, &diags
}

// CopyFieldFromPlanToRespList copies a specific field from each object in the plan list to the corresponding object in the response list.
// The lists must be of the same length and contain objects in the same order of the plan values.
func CopyFieldFromPlanToRespList(ctx context.Context, planValue, respValue attr.Value, fieldName string) (attr.Value, *diag.Diagnostics) {
	var diags diag.Diagnostics

	// Check if both values are null or unknown
	if planValue.IsNull() || planValue.IsUnknown() {
		return respValue, &diags
	}

	if respValue.IsNull() || respValue.IsUnknown() {
		return respValue, &diags
	}

	planList, ok := planValue.(types.List)
	if !ok {
		diags.AddError(
			"Invalid Plan Value Type",
			fmt.Sprintf("Expected types.List, got %T", planValue),
		)
		return respValue, &diags
	}

	respList, ok := respValue.(types.List)
	if !ok {
		diags.AddError(
			"Invalid Response Value Type",
			fmt.Sprintf("Expected types.List, got %T", respValue),
		)
		return respValue, &diags
	}

	// Get the elements from both lists
	planElements := planList.Elements()
	respElements := respList.Elements()

	if len(planElements) != len(respElements) {
		diags.AddError(
			"List Length Mismatch",
			fmt.Sprintf("Plan list has %d elements, response list has %d elements",
				len(planElements), len(respElements)),
		)
		return respValue, &diags
	}

	modifiedElements := make([]attr.Value, len(respElements))

	for i, respElement := range respElements {
		planElement := planElements[i]

		// Convert elements to objects
		planObj, ok := planElement.(types.Object)
		if !ok {
			diags.AddError(
				"Invalid Plan Element Type",
				fmt.Sprintf("Expected types.Object at index %d, got %T", i, planElement),
			)
			return respValue, &diags
		}

		respObj, ok := respElement.(types.Object)
		if !ok {
			diags.AddError(
				"Invalid Response Element Type",
				fmt.Sprintf("Expected types.Object at index %d, got %T", i, respElement),
			)
			return respValue, &diags
		}

		planAttrs := planObj.Attributes()
		respAttrs := respObj.Attributes()

		// Check if the field exists in the plan object
		planFieldValue, exists := planAttrs[fieldName]
		if !exists {
			diags.AddError(
				"Field Not Found in Plan",
				fmt.Sprintf("Field '%s' not found in plan object at index %d", fieldName, i),
			)
			return respValue, &diags
		}
		if planFieldValue.IsUnknown() || planFieldValue.IsNull() || (planFieldValue.Type(ctx) == types.StringType && planFieldValue.(basetypes.StringValue).ValueString() == "") {
			modifiedElements[i] = respElement
			continue
		}

		// Check if the field exists in the response object
		if _, exists := respAttrs[fieldName]; !exists {
			diags.AddError(
				"Field Not Found in Response",
				fmt.Sprintf("Field '%s' not found in response object at index %d", fieldName, i),
			)
			return respValue, &diags
		}

		newAttrs := make(map[string]attr.Value)
		for k, v := range respAttrs {
			newAttrs[k] = v
		}
		newAttrs[fieldName] = planFieldValue

		// Create a new object with the modified attributes
		newObj, objDiags := types.ObjectValue(respObj.AttributeTypes(ctx), newAttrs)
		diags.Append(objDiags...)
		if objDiags.HasError() {
			return respValue, &diags
		}

		modifiedElements[i] = newObj
	}

	// Create a new list with the modified elements
	newList, listDiags := types.ListValue(respList.ElementType(ctx), modifiedElements)
	diags.Append(listDiags...)
	if listDiags.HasError() {
		return respValue, &diags
	}

	return newList, &diags
}

// CopyFieldFromPlanToRespObject copies a specific field from the plan object to the response object.
func CopyFieldFromPlanToRespObject(ctx context.Context, planValue, respValue attr.Value, fieldName string) (attr.Value, *diag.Diagnostics) {
	var diags diag.Diagnostics

	if planValue.IsNull() || planValue.IsUnknown() {
		return respValue, &diags
	}

	if respValue.IsNull() || respValue.IsUnknown() {
		return respValue, &diags
	}

	planObject, ok := planValue.(types.Object)
	if !ok {
		diags.AddError(
			"Invalid Plan Value Type",
			fmt.Sprintf("Expected types.Object, got %T", planValue),
		)
		return respValue, &diags
	}

	respObject, ok := respValue.(types.Object)
	if !ok {
		diags.AddError(
			"Invalid Response Value Type",
			fmt.Sprintf("Expected types.Object, got %T", respValue),
		)
		return respValue, &diags
	}

	planAttrs := planObject.Attributes()
	respAttrs := respObject.Attributes()

	planFieldValue, exists := planAttrs[fieldName]
	if !exists {
		diags.AddError(
			"Field Not Found in Plan",
			fmt.Sprintf("Field '%s' not found in plan object", fieldName),
		)
		return respValue, &diags
	}

	if planFieldValue.IsUnknown() {
		return respValue, &diags
	}

	if _, exists := respAttrs[fieldName]; !exists {
		diags.AddError(
			"Field Not Found in Response",
			fmt.Sprintf("Field '%s' not found in response object", fieldName),
		)
		return respValue, &diags
	}

	respAttrs[fieldName] = planFieldValue

	newObj, objDiags := types.ObjectValue(respObject.AttributeTypes(ctx), respAttrs)
	diags.Append(objDiags...)
	if objDiags.HasError() {
		return respValue, &diags
	}

	return newObj, &diags
}

func ReorderAndFilterDHCPOptions(
	ctx context.Context,
	planValue attr.Value,
	stateValue attr.Value,
) (attr.Value, *diag.Diagnostics) {

	var diags diag.Diagnostics
	primaryKey := "name"
	secondaryKey := "num"

	// Handle null/unknown gracefully
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

	// Build lookup maps from state: name->element and num->element
	nameToState := make(map[string]attr.Value)
	numToState := make(map[int64]attr.Value)

	for _, elem := range stateList.Elements() {
		obj, ok := elem.(basetypes.ObjectValue)
		if !ok {
			continue
		}
		attrs := obj.Attributes()

		// name -> basetypes.StringValue (per your note state has both keys)
		if nameAttr, has := attrs[primaryKey]; has && nameAttr != nil && !nameAttr.IsNull() && !nameAttr.IsUnknown() {
			if strVal, ok := nameAttr.(basetypes.StringValue); ok {
				nameToState[strVal.ValueString()] = elem
			}
		}

		// num -> basetypes.Int64Value (per your note state has both keys)
		if numAttr, has := attrs[secondaryKey]; has && numAttr != nil && !numAttr.IsNull() && !numAttr.IsUnknown() {
			if intVal, ok := numAttr.(basetypes.Int64Value); ok {
				numToState[intVal.ValueInt64()] = elem
			}
		}
	}

	// Rebuild ordered slice based on plan order
	var reordered []attr.Value
	for _, planElem := range planList.Elements() {
		planObj, ok := planElem.(basetypes.ObjectValue)
		if !ok {
			// if plan contains something else, append it as fallback
			reordered = append(reordered, planElem)
			continue
		}
		planAttributes := planObj.Attributes()

		var matchedState attr.Value

		// Try primaryKey (name) first if present and valid
		if primaryKeyAttribute, has := planAttributes[primaryKey]; has && primaryKeyAttribute != nil && !primaryKeyAttribute.IsNull() && !primaryKeyAttribute.IsUnknown() {
			if primaryKeyAttributeValue, ok := primaryKeyAttribute.(basetypes.StringValue); ok {
				if s, exists := nameToState[primaryKeyAttributeValue.ValueString()]; exists {
					matchedState = s
				}
			}
		}

		// If not matched by name, try secondaryKey (num)
		if matchedState == nil {
			if secondaryKeyAttribute, has := planAttributes[secondaryKey]; has && secondaryKeyAttribute != nil && !secondaryKeyAttribute.IsNull() && !secondaryKeyAttribute.IsUnknown() {
				if secondaryKeyAttributeValue, ok := secondaryKeyAttribute.(basetypes.Int64Value); ok {
					if s, exists := numToState[secondaryKeyAttributeValue.ValueInt64()]; exists {
						matchedState = s
					}
				}
			}
		}

		// If matchedState found, use it; else fall back to plan element itself
		if matchedState != nil {
			reordered = append(reordered, matchedState)
		} else {
			reordered = append(reordered, planElem)
		}
	}

	// Create new ListValue with same element type as plan list
	newList, d := basetypes.NewListValue(planList.ElementType(ctx), reordered)
	diags.Append(d...)

	return newList, &diags
}
