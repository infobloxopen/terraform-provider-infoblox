package core

import (
	"encoding/json"
	"regexp"
	"strconv"
	"strings"
)

const (
	// DefaultListLimit is the default page size for list operations
	DefaultListLimit int32 = 1000
)

var availableNetworksRe = regexp.MustCompile(`The available networks are: (\d+)`)

// ExtractNIOSRef extracts the object identifier from a full NIOS _ref.
// e.g., "record:a/ZG5z..." -> "ZG5z..."
// Returns the input unchanged if no prefix is found.
func ExtractNIOSRef(ref string) string {
	v := strings.SplitN(strings.Trim(ref, "/"), "/", 2)
	if len(v) < 2 {
		return ref
	}
	return v[1]
}

// ParseEAValue parses a string value that may represent a JSON array.
// Returns parsed []any if valid JSON array, otherwise returns original string.
// TODO: handle single-quoted arrays if needed.
func ParseEAValue(valStr string) any {
	trimmed := strings.TrimSpace(valStr)
	if !strings.HasPrefix(trimmed, "[") || !strings.HasSuffix(trimmed, "]") {
		return valStr
	}

	var parsed []any
	if err := json.Unmarshal([]byte(trimmed), &parsed); err == nil {
		return parsed
	}

	return valStr
}

// StringifyEAValue converts an EA value to string for TF state.
// Arrays are converted to JSON format (double quotes).
func StringifyEAValue(val any) string {
	switch v := val.(type) {
	case []any:
		jsonBytes, err := json.Marshal(v)
		if err != nil {
			return "[]"
		}
		return string(jsonBytes)
	case string:
		return v
	default:
		if val == nil {
			return ""
		}
		jsonBytes, err := json.Marshal(val)
		if err != nil {
			return ""
		}
		return string(jsonBytes)
	}
}

// JoinFilters joins filter strings with " and "
func JoinFilters(filters []string) string {
	return strings.Join(filters, " and ")
}

func FilterExpr(key, value string) string {
	if _, err := strconv.Atoi(value); err == nil {
		return key + "==" + value
	}
	if _, err := strconv.ParseFloat(value, 64); err == nil {
		return key + "==" + value
	}
	return key + "=='" + value + "'"
}

// BuildTagFilter builds a UDDI tfilter expression from tag key/value pairs.
// e.g. {"a": "1", "b": "2"} -> "'a'=='1' and 'b'=='2'".
func BuildTagFilter(tags map[string]string) string {
	filters := make([]string, 0, len(tags))
	for k, v := range tags {
		filters = append(filters, "'"+k+"'=='"+v+"'")
	}
	return JoinFilters(filters)
}

// ExtractAvailableCountFromError parses a UDDI next-available 400 error body and
// returns the number of resources actually available, enabling partial-fill
// allocation across multiple tag-matched scopes. Returns 0 when not parseable.
func ExtractAvailableCountFromError(body []byte) int32 {
	var errorResponse struct {
		Error []struct {
			Message string `json:"message"`
		} `json:"error"`
	}

	if err := json.Unmarshal(body, &errorResponse); err != nil {
		return 0
	}

	for _, e := range errorResponse.Error {
		if match := availableNetworksRe.FindStringSubmatch(e.Message); len(match) > 1 {
			if count, err := strconv.ParseInt(match[1], 10, 32); err == nil {
				return int32(count)
			}
		}
	}
	return 0
}

// TranslateFilterKeys translates unified filter keys to backend-specific API field names.
// Keys not found in the mapping are passed through unchanged.
func TranslateFilterKeys(filters map[string]string, fieldMap map[string]string) map[string]string {
	if filters == nil || len(fieldMap) == 0 {
		return filters
	}
	result := make(map[string]string, len(filters))
	for k, v := range filters {
		if mappedKey, ok := fieldMap[k]; ok {
			result[mappedKey] = v
		} else {
			result[k] = v
		}
	}
	return result
}

// ReadAllPagesNIOS fetches all pages using next_page_id pagination
func ReadAllPagesNIOS[T any](fetchPage func(pageID string) ([]T, string, error)) ([]T, error) {
	var allResults []T
	var pageID string
	for {
		results, nextPageID, err := fetchPage(pageID)
		if err != nil {
			return nil, err
		}
		allResults = append(allResults, results...)
		if nextPageID == "" {
			break
		}
		pageID = nextPageID
	}
	return allResults, nil
}

// ReadAllPagesUDDI fetches all pages using offset/limit pagination.
// pagingOpts, in order: page size (default DefaultListLimit), paging (0 = first page only).
func ReadAllPagesUDDI[T any](fetchPage func(offset, limit int32) ([]T, error), pagingOpts ...int32) ([]T, error) {

	pageSize, paging := DefaultListLimit, int32(1)

	if len(pagingOpts) > 0 && pagingOpts[0] > 0 {
		pageSize = pagingOpts[0]
	}
	if len(pagingOpts) > 1 {
		paging = pagingOpts[1]
	}

	var allResults []T
	var offset int32 = 0
	for {
		results, err := fetchPage(offset, pageSize)
		if err != nil {
			return nil, err
		}
		allResults = append(allResults, results...)
		if paging == 0 || int32(len(results)) < pageSize {
			break
		}
		offset += pageSize
	}
	return allResults, nil
}
