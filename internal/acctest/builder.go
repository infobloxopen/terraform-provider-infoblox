package acctest

import (
	"fmt"
	"strings"
)

// BuildResourceHCL generates HCL config string for a resource.
func BuildResourceHCL(resourceType, resourceLabel string, tv *CaseConfig) string {
	var sb strings.Builder

	fmt.Fprintf(&sb, "resource %q %q {\n", resourceType, resourceLabel)

	for k, v := range tv.Common {
		fmt.Fprintf(&sb, "  %s = %s\n", k, formatHCLValue(v))
	}

	// Always add backend block (even if empty) - required by resource validation
	if tv.Backend == "nios" {
		if len(tv.NIOS) > 0 {
			sb.WriteString("  nios = {\n")
			for k, v := range tv.NIOS {
				fmt.Fprintf(&sb, "    %s = %s\n", k, formatHCLValue(v))
			}
			sb.WriteString("  }\n")
		} else {
			sb.WriteString("  nios = {}\n")
		}
	}

	if tv.Backend == "uddi" {
		if len(tv.UDDI) > 0 {
			sb.WriteString("  uddi = {\n")
			for k, v := range tv.UDDI {
				fmt.Fprintf(&sb, "    %s = %s\n", k, formatHCLValue(v))
			}
			sb.WriteString("  }\n")
		} else {
			sb.WriteString("  uddi = {}\n")
		}
	}

	sb.WriteString("}\n")
	return sb.String()
}

// RawExpr is a verbatim HCL expression (e.g. a resource reference like ${infoblox_zone_auth.prereq.fqdn})
// emitted unquoted so Terraform resolves it at apply time instead of treating it as a string literal.
type RawExpr string

func formatHCLValue(v any) string {
	switch val := v.(type) {
	case RawExpr:
		return string(val)
	case string:
		return fmt.Sprintf("%q", val)
	case bool:
		return fmt.Sprintf("%t", val)
	case int:
		return fmt.Sprintf("%d", val)
	case int64:
		return fmt.Sprintf("%d", val)
	case float64:
		if val == float64(int(val)) {
			return fmt.Sprintf("%d", int(val))
		}
		return fmt.Sprintf("%g", val)
	case map[string]any:
		return formatHCLMap(val)
	case []any:
		return formatHCLList(val)
	case nil:
		return "null"
	default:
		return fmt.Sprintf("%q", fmt.Sprintf("%v", val))
	}
}

func formatHCLMap(m map[string]any) string {
	if len(m) == 0 {
		return "{}"
	}
	var parts []string
	for k, v := range m {
		parts = append(parts, fmt.Sprintf("%q = %s", k, formatHCLValue(v)))
	}
	return "{\n    " + strings.Join(parts, "\n    ") + "\n  }"
}

func formatHCLList(arr []any) string {
	if len(arr) == 0 {
		return "[]"
	}
	var parts []string
	for _, v := range arr {
		parts = append(parts, formatHCLValue(v))
	}
	return "[" + strings.Join(parts, ", ") + "]"
}
