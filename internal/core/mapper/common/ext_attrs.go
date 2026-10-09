package common

import (
	"fmt"
	"reflect"

	"github.com/infobloxopen/terraform-provider-infoblox/internal/core"
)

// ProcessExtAttrs converts map[string]any (ExtAttrs) to NIOS ExtAttrs struct format.
// src should have an ExtAttrs field of type map[string]any
// dst should have an ExtAttrs field of type *map[string]dns.ExtAttrs
func ProcessExtAttrs(src any, dst any) error {
	srcVal := reflect.ValueOf(src).Elem()
	dstVal := reflect.ValueOf(dst).Elem()

	// read ExtAttrs from src
	attrField := srcVal.FieldByName("ExtAttrs")
	if !attrField.IsValid() || attrField.IsNil() {
		return fmt.Errorf("ExtAttrs field not found or nil in source")
	}

	// find ExtAttrs in dst
	dstField := dstVal.FieldByName("ExtAttrs")
	if !dstField.IsValid() || !dstField.CanSet() {
		return fmt.Errorf("ExtAttrs field not found or cannot be set in destination")
	}

	converted, err := convertToExtAttrs(attrField, dstField.Type())
	if err != nil {
		return fmt.Errorf("ExtAttrs pre-processing failed: %v", err)
	}

	dstField.Set(converted)

	return nil
}

func convertToExtAttrs(src reflect.Value, dstType reflect.Type) (reflect.Value, error) {
	if dstType.Kind() != reflect.Pointer {
		return reflect.Value{}, fmt.Errorf("expected pointer to map, got %s", dstType)
	}

	mapType := dstType.Elem()
	elemType := mapType.Elem()

	resultMap := reflect.MakeMap(mapType)

	iter := src.MapRange()
	for iter.Next() {
		k := iter.Key()
		v := iter.Value()

		// create ExtAttrs struct
		newVal := reflect.New(elemType).Elem()

		// set Value field
		valueField := newVal.FieldByName("Value")
		if !valueField.IsValid() || !valueField.CanSet() {
			return reflect.Value{}, fmt.Errorf("value field not found in %s", elemType)
		}

		// Parse the value - convert JSON array strings to actual arrays
		parsedVal := v.Interface()
		if strVal, ok := parsedVal.(string); ok {
			parsedVal = core.ParseEAValue(strVal)
		}

		valueField.Set(reflect.ValueOf(parsedVal))

		resultMap.SetMapIndex(k, newVal)
	}

	ptr := reflect.New(mapType)
	ptr.Elem().Set(resultMap)

	return ptr, nil
}
