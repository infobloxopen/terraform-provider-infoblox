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
	_ basetypes.StringValuable                   = (*IPv4Name)(nil)
	_ basetypes.StringValuableWithSemanticEquals = (*IPv4Name)(nil)
	_ xattr.ValidateableAttribute                = (*IPv4Name)(nil)
	_ function.ValidateableParameter             = (*IPv4Name)(nil)
)

type IPv4Name struct {
	basetypes.StringValue
}

// Type returns an IPv4NameType.
func (v IPv4Name) Type(_ context.Context) attr.Type {
	return IPv4NameType{}
}

// Equal returns true if the given value is equivalent.
func (v IPv4Name) Equal(o attr.Value) bool {
	other, ok := o.(IPv4Name)

	if !ok {
		return false
	}

	return v.StringValue.Equal(other.StringValue)
}

// StringSemanticEquals returns true if the given IPv4 name string value is semantically equal to the current IPv4 name string value.
// This comparison utilizes netip.ParseAddr and then compares the resulting netip.Addr representations.
func (v IPv4Name) StringSemanticEquals(_ context.Context, newValuable basetypes.StringValuable) (bool, diag.Diagnostics) {
	var diags diag.Diagnostics

	newValue, ok := newValuable.(IPv4Name)
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

	ipv4, zone, ok := splitIPv4AndZone(v.ValueString())
	if !ok {
		diags.AddError(
			"Invalid IPv4 Name Value",
			"The current value is not a valid IPv4 name. "+
				"Expected format '<ipv4>.<rp-zone>' or '<ipv4>/<prefix>.<rp-zone>'.\n\n"+
				"Given Value: "+v.ValueString(),
		)
		return false, diags
	}

	newIPv4, newZone, ok := splitIPv4AndZone(newValue.ValueString())
	if !ok {
		diags.AddError(
			"Invalid IPv4 Name Value",
			"The new value is not a valid IPv4 name. "+
				"Expected format '<ipv4>.<rp-zone>' or '<ipv4>/<prefix>.<rp-zone>'.\n\n"+
				"Given Value: "+newValue.ValueString(),
		)
		return false, diags
	}

	// IPv4 addresses are already validated at this point, ignoring errors
	if strings.Contains(ipv4, "/") {
		// CIDR
		newPrefix, _ := netip.ParsePrefix(newIPv4)
		currentPrefix, _ := netip.ParsePrefix(ipv4)

		return currentPrefix == newPrefix && zone == newZone, diags
	} else {
		// Plain address
		newIPAddr, _ := netip.ParseAddr(newIPv4)
		currentIPAddr, _ := netip.ParseAddr(ipv4)

		return currentIPAddr == newIPAddr && zone == newZone, diags
	}
}

// ValidateAttribute implements attribute value validation. This type requires the value provided to be a String
// value that is a valid IPv4 name.
func (v IPv4Name) ValidateAttribute(ctx context.Context, req xattr.ValidateAttributeRequest, resp *xattr.ValidateAttributeResponse) {
	var (
		ipAddr netip.Addr
		err    error
	)

	if v.IsUnknown() || v.IsNull() {
		return
	}

	ipv4, _, ok := splitIPv4AndZone(v.ValueString())
	if !ok {
		resp.Diagnostics.AddAttributeError(
			req.Path,
			"Invalid IPv4 Name Value",
			"A string value was provided that is not a valid IPv4 name. "+
				"Expected format: '<ipv4>.<rp-zone>' or '<ipv4>/<prefix>.<rp-zone>'.\n\n"+
				"Given Value: "+v.ValueString(),
		)
		return
	}

	if strings.Contains(ipv4, "/") {
		// CIDR
		prefix, err := netip.ParsePrefix(ipv4)
		if err != nil {
			resp.Diagnostics.AddAttributeError(
				req.Path,
				"Invalid IPv4 Name Value",
				"A string value was provided that is not a valid IPv4 name. "+
					"Expected format: '<ipv4>.<rp-zone>' or '<ipv4>/<prefix>.<rp-zone>'.\n\n"+
					"Given Value: "+v.ValueString(),
			)

			return
		}
		ipAddr = prefix.Addr()
	} else {
		// Plain address
		ipAddr, err = netip.ParseAddr(ipv4)
		if err != nil {
			resp.Diagnostics.AddAttributeError(
				req.Path,
				"Invalid IPv4 Name Value",
				"A string value was provided that is not a valid IPv4 name. "+
					"Expected format: '<ipv4>.<rp-zone>' or '<ipv4>/<prefix>.<rp-zone>'.\n\n"+
					"Given Value: "+v.ValueString(),
			)

			return
		}
	}

	if !ipAddr.Is4() {
		resp.Diagnostics.AddAttributeError(
			req.Path,
			"Invalid IPv4 Name Value",
			"A string value was provided that is not a valid IPv4 name. "+
				"The IPv4 portion must be a valid IPv4 address or IPv4 CIDR prefix.\n\n"+
				"Expected format: '<ipv4>.<rp-zone>' or '<ipv4>/<prefix>.<rp-zone>'.\n\n"+
				"Given Value: "+v.ValueString(),
		)

		return
	}
}

// ValidateParameter implements provider-defined function parameter value validation. This type requires the value provided
// to be a String value that is a valid IPv4 name.
func (v IPv4Name) ValidateParameter(ctx context.Context, req function.ValidateParameterRequest, resp *function.ValidateParameterResponse) {
	var (
		ipAddr netip.Addr
		err    error
	)

	if v.IsUnknown() || v.IsNull() {
		return
	}

	ipv4, _, ok := splitIPv4AndZone(v.ValueString())
	if !ok {
		resp.Error = function.NewArgumentFuncError(
			req.Position,
			"Invalid IPv4 Name Value: "+
				"A string value was provided that is not a valid IPv4 name. "+
				"Expected format: '<ipv4>.<rp-zone>' or '<ipv4>/<prefix>.<rp-zone>'.\n\n"+
				"Given Value: "+v.ValueString(),
		)
		return
	}

	if strings.Contains(ipv4, "/") {
		// CIDR
		prefix, err := netip.ParsePrefix(ipv4)
		if err != nil {
			resp.Error = function.NewArgumentFuncError(
				req.Position,
				"Invalid IPv4 Name Value: "+
					"A string value was provided that is not a valid IPv4 name. "+
					"Expected format: '<ipv4>.<rp-zone>' or '<ipv4>/<prefix>.<rp-zone>'.\n\n"+
					"Given Value: "+v.ValueString()+"\n"+
					"Error: "+err.Error(),
			)

			return
		}
		ipAddr = prefix.Addr()
	} else {
		// Plain address
		ipAddr, err = netip.ParseAddr(ipv4)
		if err != nil {
			resp.Error = function.NewArgumentFuncError(
				req.Position,
				"Invalid IPv4 Name Value: "+
					"A string value was provided that is not a valid IPv4 name.\n\n"+
					"Expected format: '<ipv4>.<rp-zone>' or '<ipv4>/<prefix>.<rp-zone>'.\n\n"+
					"Given Value: "+v.ValueString()+"\n"+
					"Error: "+err.Error(),
			)

			return
		}
	}

	if !ipAddr.Is4() {
		resp.Error = function.NewArgumentFuncError(
			req.Position,
			"Invalid IPv4 Name Value: "+
				"A string value was provided that is not a valid IPv4 name.\n\n"+
				"Expected format: '<ipv4>.<rp-zone>' or '<ipv4>/<prefix>.<rp-zone>'.\n\n"+
				"Given Value: "+v.ValueString(),
		)

		return
	}
}

// splitIPv4AndZone splits an IPv4 name into its IP and zone components.
func splitIPv4AndZone(value string) (ip, zone string, ok bool) {
	for i := 0; i < len(value); i++ {
		if value[i] != '.' {
			continue
		}

		left := value[:i]
		right := value[i+1:]

		// Try IPv4 address
		if addr, err := netip.ParseAddr(left); err == nil && addr.Is4() {
			return left, right, true
		}

		// Try IPv4 CIDR
		if pfx, err := netip.ParsePrefix(left); err == nil && pfx.Addr().Is4() {
			return left, right, true
		}
	}

	return "", "", false
}
