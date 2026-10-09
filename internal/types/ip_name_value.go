package types

import (
	"context"
	"fmt"
	"net/netip"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/attr/xattr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/function"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

var (
	_ basetypes.StringValuable                   = (*IPName)(nil)
	_ basetypes.StringValuableWithSemanticEquals = (*IPName)(nil)
	_ xattr.ValidateableAttribute                = (*IPName)(nil)
	_ function.ValidateableParameter             = (*IPName)(nil)
)

type IPName struct {
	basetypes.StringValue
}

// Type returns an IPNameType.
func (v IPName) Type(_ context.Context) attr.Type {
	return IPNameType{}
}

// Equal returns true if the given value is equivalent.
func (v IPName) Equal(o attr.Value) bool {
	other, ok := o.(IPName)

	if !ok {
		return false
	}

	return v.StringValue.Equal(other.StringValue)
}

// StringSemanticEquals returns true if the given IP name string value is semantically equal to the current IP name string value.
// This comparison utilizes netip.ParseAddr and then compares the resulting netip.Addr representations. This means `compressed` IP address values
// are considered semantically equal to `non-compressed` IP address values.
func (v IPName) StringSemanticEquals(_ context.Context, newValuable basetypes.StringValuable) (bool, diag.Diagnostics) {
	var diags diag.Diagnostics

	newValue, ok := newValuable.(IPName)
	if !ok {
		diags.AddError(
			"Semantic Equality Check Error",
			"An unexpected value type was received while performing semantic equality checks. "+
				"Please report this to the provider developers.\n\n"+
				"Expected Value Type: "+fmt.Sprintf("%T", v)+"\n"+
				"Got Value Type: "+fmt.Sprintf("%T", newValuable),
		)

		return false, diags
	}

	ip, zone, ok := splitIPAndZone(v.ValueString())
	if !ok {
		diags.AddError(
			"Invalid IP Name Value",
			"The current value is not a valid IP name. "+
				"Expected format '<ip>.<rp-zone>' or '<ip>/<prefix>.<rp-zone>'.\n\n"+
				"Given Value: "+v.ValueString(),
		)
		return false, diags
	}

	newIP, newZone, ok := splitIPAndZone(newValue.ValueString())
	if !ok {
		diags.AddError(
			"Invalid IP Name Value",
			"The new value is not a valid IP name. "+
				"Expected format '<ip>.<rp-zone>' or '<ip>/<prefix>.<rp-zone>'.\n\n"+
				"Given Value: "+newValue.ValueString(),
		)
		return false, diags
	}

	// IP addresses are already validated at this point, ignoring errors
	if strings.Contains(ip, "/") {
		// CIDR
		newPrefix, _ := netip.ParsePrefix(newIP)
		currentPrefix, _ := netip.ParsePrefix(ip)

		return currentPrefix == newPrefix && zone == newZone, diags
	} else {
		// Plain address
		newIpAddr, _ := netip.ParseAddr(newIP)
		currentIpAddr, _ := netip.ParseAddr(ip)

		return currentIpAddr == newIpAddr && zone == newZone, diags
	}
}

// ValidateAttribute implements attribute value validation. This type requires the value provided to be a String
// value that is a valid IP name.
func (v IPName) ValidateAttribute(ctx context.Context, req xattr.ValidateAttributeRequest, resp *xattr.ValidateAttributeResponse) {
	var (
		ipAddr netip.Addr
		err    error
	)

	if v.IsUnknown() || v.IsNull() {
		return
	}

	ip, _, ok := splitIPAndZone(v.ValueString())
	if !ok {
		resp.Diagnostics.AddAttributeError(
			req.Path,
			"Invalid IP Name Value",
			"A string value was provided that is not a valid IP name. "+
				"Expected format '<ip>.<rp-zone>' or '<ip>/<prefix>.<rp-zone>'.\n\n"+
				"Given Value: "+v.ValueString(),
		)
		return
	}

	if strings.Contains(ip, "/") {
		// CIDR
		prefix, err := netip.ParsePrefix(ip)
		if err != nil {
			resp.Diagnostics.AddAttributeError(
				req.Path,
				"Invalid IP Name Value",
				"A string value was provided that is not a valid IP name. "+
					"Expected format '<ip>.<rp-zone>' or '<ip>/<prefix>.<rp-zone>'.\n\n"+
					"Given Value: "+v.ValueString(),
			)

			return
		}
		ipAddr = prefix.Addr()
	} else {
		// Plain address
		ipAddr, err = netip.ParseAddr(ip)
		if err != nil {
			resp.Diagnostics.AddAttributeError(
				req.Path,
				"Invalid IP Name Value",
				"A string value was provided that is not a valid IP name. "+
					"Expected format '<ip>.<rp-zone>' or '<ip>/<prefix>.<rp-zone>'.\n\n"+
					"Given Value: "+v.ValueString()+"\n"+
					"Error: "+err.Error(),
			)

			return
		}
	}

	if !ipAddr.Is4() && !ipAddr.Is6() {
		resp.Diagnostics.AddAttributeError(
			req.Path,
			"Invalid IP Name Value",
			"A string value was provided that is not a valid IP name. "+
				"Expected format '<ip>.<rp-zone>' or '<ip>/<prefix>.<rp-zone>'.\n\n"+
				"Given Value: "+v.ValueString(),
		)
		return
	}
}

// ValidateParameter implements provider-defined function parameter value validation. This type requires the value provided
// to be a String value that is a valid IP name (IPv4 or IPv6).
func (v IPName) ValidateParameter(ctx context.Context, req function.ValidateParameterRequest, resp *function.ValidateParameterResponse) {
	var (
		ipAddr netip.Addr
		err    error
	)

	if v.IsUnknown() || v.IsNull() {
		return
	}

	ip, _, ok := splitIPAndZone(v.ValueString())
	if !ok {
		resp.Error = function.NewArgumentFuncError(
			req.Position,
			"Invalid IP Name Value: "+
				"A string value was provided that is not a valid IP name. "+
				"Expected format '<ip>.<rp-zone>' or '<ip>/<prefix>.<rp-zone>'.\n\n"+
				"Given Value: "+v.ValueString(),
		)
		return
	}

	if strings.Contains(ip, "/") {
		// CIDR
		prefix, err := netip.ParsePrefix(ip)
		if err != nil {
			resp.Error = function.NewArgumentFuncError(
				req.Position,
				"Invalid IP Name Value: "+
					"A string value was provided that is not a valid IP name. "+
					"Expected format: '<ip>.<rp-zone>' or '<ip>/<prefix>.<rp-zone>'.\n\n"+
					"Given Value: "+v.ValueString()+"\n"+
					"Error: "+err.Error(),
			)

			return
		}
		ipAddr = prefix.Addr()
	} else {
		// Plain address
		ipAddr, err = netip.ParseAddr(ip)
		if err != nil {
			resp.Error = function.NewArgumentFuncError(
				req.Position,
				"Invalid IP Name Value: "+
					"A string value was provided that is not a valid IP name.\n\n"+
					"Expected format: '<ip>.<rp-zone>' or '<ip>/<prefix>.<rp-zone>'.\n\n"+
					"Given Value: "+v.ValueString()+"\n"+
					"Error: "+err.Error(),
			)

			return
		}
	}

	if !ipAddr.Is4() && !ipAddr.Is6() {
		resp.Error = function.NewArgumentFuncError(
			req.Position,
			"Invalid IP Name Value: "+
				"A string value was provided that is not a valid IP name. "+
				"Expected format '<ip>.<rp-zone>' or '<ip>/<prefix>.<rp-zone>'.\n\n"+
				"Given Value: "+v.ValueString(),
		)
		return
	}
}

// splitIPAndZone splits an IP name string into its IP and rp-zone components.
// It returns the IP, rp-zone, and a boolean indicating success.
func splitIPAndZone(value string) (string, string, bool) {
	// Find the split point where left is a valid IP or CIDR
	for i := 0; i < len(value); i++ {
		if value[i] != '.' {
			continue
		}

		left := value[:i]
		right := value[i+1:]

		// Try IPv4 / IPv6 address
		if addr, err := netip.ParseAddr(left); err == nil {
			if addr.Is4() || addr.Is6() {
				return left, right, true
			}
		}

		// Try IPv4 / IPv6 prefix
		if pfx, err := netip.ParsePrefix(left); err == nil {
			if pfx.Addr().Is4() || pfx.Addr().Is6() {
				return left, right, true
			}
		}
	}

	return "", "", false
}
