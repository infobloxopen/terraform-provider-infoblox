package common

import (
	"fmt"
	"reflect"
	"strings"
)

// MapTo maps src to a new instance of type D using the provided field map.
// Returns the mapped destination and any error.
func MapTo[D any](src any, fieldMap map[string]string) (D, error) {
	var dst D
	err := MapFields(src, &dst, fieldMap)
	return dst, err
}

func MapFields(src any, dst any, fieldMap map[string]string) error {
	srcVal := reflect.ValueOf(src).Elem()
	dstVal := reflect.ValueOf(dst).Elem()

	// Track top-level fields that have explicit mappings
	mappedSrcFields := make(map[string]bool)

	// First pass: explicit mappings from fieldMap
	for srcPath, dstPath := range fieldMap {
		// Only track direct top-level mappings (no nested paths)
		if !strings.Contains(srcPath, ".") {
			mappedSrcFields[srcPath] = true
		}

		s, err := getByPath(srcVal, srcPath)
		if err != nil {
			continue // skip missing source paths
		}

		if !s.IsValid() || isZeroValue(s) {
			continue
		}

		err = setByPath(dstVal, dstPath, s)
		if err != nil {
			return fmt.Errorf("mapping failed for %s -> %s : %v", srcPath, dstPath, err)
		}
	}

	// Second pass: auto-map same-name top-level fields not mentioned in fieldMap.
	// This is a fallback for fields with identical names in src and dst.
	// Nested fields and fields with different names must be explicitly mapped in fieldMap and are not handled here.
	if srcVal.Kind() == reflect.Struct {
		srcType := srcVal.Type()
		for i := 0; i < srcVal.NumField(); i++ {
			field := srcType.Field(i)
			fieldVal := srcVal.Field(i)

			// Skip if explicitly mapped
			if mappedSrcFields[field.Name] {
				continue
			}

			// Skip unexported fields
			if !field.IsExported() {
				continue
			}

			// Skip zero/nil values
			if !fieldVal.IsValid() || isZeroValue(fieldVal) {
				continue
			}

			// Try to set same-name field in destination (silently skip if not found)
			_ = setByPath(dstVal, field.Name, fieldVal)
		}
	}

	return nil
}

func isZeroValue(v reflect.Value) bool {
	switch v.Kind() {
	case reflect.Pointer, reflect.Interface, reflect.Slice, reflect.Map, reflect.Chan, reflect.Func:
		return v.IsNil()
	default:
		return v.IsZero()
	}
}

func setByPath(dst reflect.Value, path string, value reflect.Value) error {
	parts := strings.Split(path, ".")

	current := dst

	for i := 0; i < len(parts); i++ {
		part := parts[i]
		last := i == len(parts)-1

		// unwrap pointer/interface
		for current.Kind() == reflect.Pointer || current.Kind() == reflect.Interface {
			if current.IsNil() {
				if current.Kind() == reflect.Pointer {
					current.Set(reflect.New(current.Type().Elem()))
				} else {
					return fmt.Errorf("nil interface at %s", part)
				}
			}
			current = current.Elem()
		}

		switch current.Kind() {

		case reflect.Struct:
			field := current.FieldByName(part)
			if !field.IsValid() {
				return fmt.Errorf("field %s not found", part)
			}

			if last {
				return setValue(field, value)
			}

			current = field

		case reflect.Map:
			if current.Type().Key().Kind() != reflect.String {
				return fmt.Errorf("only map[string] supported, got %s", current.Type())
			}

			key := reflect.ValueOf(part)

			if current.IsNil() {
				current.Set(reflect.MakeMap(current.Type()))
			}

			if last {
				return setMapValue(current, key, value)
			}

			next := current.MapIndex(key)

			if !next.IsValid() {
				created, err := createContainer(current.Type().Elem())
				if err != nil {
					return err
				}
				current.SetMapIndex(key, created)
				next = created
			}

			current = next

		default:
			return fmt.Errorf("unsupported type %s at %s", current.Kind(), part)
		}
	}

	return nil
}

func setValue(dst reflect.Value, src reflect.Value) error {
	if !dst.CanSet() {
		return fmt.Errorf("cannot set %s", dst.Type())
	}

	if src.Kind() == reflect.Interface {
		if src.IsNil() {
			return nil
		}
		src = src.Elem()
	}

	// Dereference src pointer if dst is non-pointer
	if src.Kind() == reflect.Pointer && dst.Kind() != reflect.Pointer {
		if src.IsNil() {
			return nil // nil pointer, don't set
		}
		src = src.Elem()
	}

	// direct assign
	if src.Type().AssignableTo(dst.Type()) {
		dst.Set(src)
		return nil
	}

	// convert
	if src.Type().ConvertibleTo(dst.Type()) {
		dst.Set(src.Convert(dst.Type()))
		return nil
	}

	// pointer assignment: dst is pointer, src is non-pointer
	if dst.Kind() == reflect.Pointer {
		ptr := reflect.New(dst.Type().Elem())

		if src.Type().AssignableTo(dst.Type().Elem()) {
			ptr.Elem().Set(src)
			dst.Set(ptr)
			return nil
		}

		if src.Type().ConvertibleTo(dst.Type().Elem()) {
			ptr.Elem().Set(src.Convert(dst.Type().Elem()))
			dst.Set(ptr)
			return nil
		}
	}

	return fmt.Errorf("cannot assign %s to %s", src.Type(), dst.Type())
}

func setMapValue(m reflect.Value, key reflect.Value, val reflect.Value) error {
	elemType := m.Type().Elem()

	if val.Type().AssignableTo(elemType) {
		m.SetMapIndex(key, val)
		return nil
	}

	if val.Type().ConvertibleTo(elemType) {
		m.SetMapIndex(key, val.Convert(elemType))
		return nil
	}

	return fmt.Errorf("cannot assign %s to map value %s", val.Type(), elemType)
}

func createContainer(t reflect.Type) (reflect.Value, error) {
	switch t.Kind() {

	case reflect.Map:
		return reflect.MakeMap(t), nil

	case reflect.Interface:
		// assume map[string]any
		return reflect.ValueOf(map[string]any{}), nil

	case reflect.Pointer:
		if t.Elem().Kind() == reflect.Map {
			m := reflect.MakeMap(t.Elem())
			ptr := reflect.New(t.Elem())
			ptr.Elem().Set(m)
			return ptr, nil
		}
		return reflect.New(t.Elem()), nil

	case reflect.Struct:
		return reflect.New(t).Elem(), nil
	}

	return reflect.Value{}, fmt.Errorf("unsupported container type %s", t)
}

func getByPath(src reflect.Value, path string) (reflect.Value, error) {
	parts := strings.Split(path, ".")

	current := src

	for _, part := range parts {

		// unwrap
		for current.Kind() == reflect.Pointer || current.Kind() == reflect.Interface {
			if current.IsNil() {
				return reflect.Value{}, fmt.Errorf("nil while reading %s", part)
			}
			current = current.Elem()
		}

		switch current.Kind() {

		case reflect.Struct:
			field := current.FieldByName(part)
			if !field.IsValid() {
				return reflect.Value{}, fmt.Errorf("field %s not found", part)
			}
			current = field

		case reflect.Map:
			key := reflect.ValueOf(part)

			val := current.MapIndex(key)
			if !val.IsValid() {
				return reflect.Value{}, fmt.Errorf("key %s not found", part)
			}
			current = val

		default:
			return reflect.Value{}, fmt.Errorf("unsupported type %s in getByPath", current.Kind())
		}
	}

	return current, nil
}
