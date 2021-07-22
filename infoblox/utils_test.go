package infoblox

import (
	"fmt"

	"reflect"
	"sort"
	"strings"

	ibclient "github.com/infobloxopen/infoblox-go-client/v2"
)

const testNetView = "default"

var NotFoundTexts = []string{"404 Not Found", "not found"}

const (
	eaListTypeString = iota
	eaListTypeInt
)

func isNotFoundError(err error) bool {
	if _, notFoundErr := err.(*ibclient.NotFoundError); notFoundErr {
		return true
	}

	// TODO: uncomment when infoblox-go-client will handle NotFoundError separately.
	//return false

	errText := err.Error()
	for _, text := range NotFoundTexts {
		if strings.Contains(errText, text) {
			return true
		}
	}

	return false
}

func typesEqual(a, b interface{}) bool {
	return reflect.TypeOf(a) == reflect.TypeOf(b)
}

func sortList(s interface{}) error {
	switch slice := s.(type) {
	case []int:
		sort.Ints(slice)
	case []string:
		stringSlice := s.([]string)
		sort.Strings(stringSlice)
	default:
		return fmt.Errorf("expected value is of an unsupported type")
	}
	return nil
}

func validateValues(actual, expected interface{}) (bool, error) {
	switch expTyped := expected.(type) {
	case int:
		av := actual.(int)
		if expTyped != av {
			return false, nil
		}
	case bool:
		av := actual.(bool)
		if expTyped != av {
			return false, nil
		}
	case string:
		av := actual.(string)
		if expTyped != av {
			return false, nil
		}
	default:
		return false, fmt.Errorf("expected value '%+v' is of an unsupported type", expected)
	}

	return true, nil
}

func validateEAs(actualEAs, expectedEAs map[string]interface{}) error {
	for eaKey, expEaVal := range expectedEAs {
		actEaVal, found := actualEAs[eaKey]
		if !found {
			return fmt.Errorf(
				"a value for extensible attribute '%s' not found, but expected to exist", eaKey)
		}

		if !typesEqual(actEaVal, expEaVal) {
			return fmt.Errorf("actual and expected values for extensible attribute '%s' have unequal types", eaKey)
		}

		reflActEaVal := reflect.ValueOf(actEaVal)
		switch reflActEaVal.Kind() {
		case reflect.Slice:
			var eaListType int

			switch actEaVal.(type) {
			case []int:
				eaListType = eaListTypeInt
			case []string:
				eaListType = eaListTypeString
			default:
				return fmt.Errorf("unsupported type for 'extensible_attributes' field value: %+v", actEaVal)
			}

			reflExpEaVal := reflect.ValueOf(expEaVal)
			if reflActEaVal.Len() != reflExpEaVal.Len() {
				return fmt.Errorf(
					"the value of extensible attribute '%s' is not equal to the expected one", eaKey)
			}
			numItems := reflExpEaVal.Len()
			if numItems == 0 {
				return nil
			}
			if err := sortList(actEaVal.(interface{})); err != nil {
				return err
			}
			if err := sortList(expEaVal.(interface{})); err != nil {
				return err
			}

			getElemFunc := func(slice interface{}, idx int) interface{} {
				switch eaListType {
				case eaListTypeInt:
					return slice.([]int)[idx]
				case eaListTypeString:
					return slice.([]string)[idx]
				default:
					panic("unexpected slice item's type")
				}
			}

			for i := 0; i < numItems; i++ {
				expVal := getElemFunc(expEaVal, i)
				actVal := getElemFunc(actEaVal, i)
				equal, err := validateValues(actVal, expVal)
				if err != nil {
					return err
				}
				if !equal {
					return fmt.Errorf(
						"the value for extensible attribute '%v' is '%v' but expected to be '%v'",
						eaKey, actEaVal, expEaVal)
				}
				return nil
			}
			return nil
		default:
			equal, err := validateValues(actEaVal, expEaVal)
			if err != nil {
				return err
			}
			if !equal {
				return fmt.Errorf(
					"the value for extensible attribute '%v' is '%v' but expected to be '%v'",
					eaKey, actEaVal, expEaVal)
			}
		}
	}

	return nil
}
