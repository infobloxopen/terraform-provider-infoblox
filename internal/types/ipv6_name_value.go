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
	_ basetypes.StringValuable                   = (*IPv6Name)(nil)
	_ basetypes.StringValuableWithSemanticEquals = (*IPv6Name)(nil)
	_ xattr.ValidateableAttribute                = (*IPv6Name)(nil)
	_ function.ValidateableParameter             = (*IPv6Name)(nil)
)

type IPv6Name struct {
	basetypes.StringValue
}

// Type returns an IPv6NameType.
func (v IPv6Name) Type(_ context.Context) attr.Type {
	return IPv6NameType{}
}

// Equal returns true if the given value is equivalent.
func (v IPv6Name) Equal(o attr.Value) bool {
	other, ok := o.(IPv6Name)

	if !ok {
		return false
	}

	return v.StringValue.Equal(other.StringValue)
}

// StringSemanticEquals returns true if the given IPv6 name string value is semantically equal to the current IPv6 name string value.
// This comparison utilizes netip.ParseAddr and then compares the resulting netip.Addr representations. This means `compressed` IPv6 address values
// are considered semantically equal to `non-compressed` IPv6 address values.
func (v IPv6Name) StringSemanticEquals(_ context.Context, newValuable basetypes.StringValuable) (bool, diag.Diagnostics) {
	var diags diag.Diagnostics

	newValue, ok := newValuable.(IPv6Name)
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

	ipv6, zone, ok := strings.Cut(v.ValueString(), ".")
	if !ok {
		diags.AddError(
			"Invalid IPv6 Name Value",
			"The current value is not a valid IPv6 name. "+
				"Expected format '<ipv6>.<rp-zone>' or '<ipv6>/<prefix>.<rp-zone>'.\n\n"+
				"Given Value: "+v.ValueString(),
		)
		return false, diags
	}

	newIPv6, newZone, ok := strings.Cut(newValue.ValueString(), ".")
	if !ok {
		diags.AddError(
			"Invalid IPv6 Name Value",
			"The new value is not a valid IPv6 name. "+
				"Expected format '<ipv6>.<rp-zone>' or '<ipv6>/<prefix>.<rp-zone>'.\n\n"+
				"Given Value: "+newValue.ValueString(),
		)
		return false, diags
	}

	// IPv6 addresses are already validated at this point, ignoring errors
	if strings.Contains(ipv6, "/") {
		// CIDR
		newPrefix, _ := netip.ParsePrefix(newIPv6)
		currentPrefix, _ := netip.ParsePrefix(ipv6)

		return currentPrefix == newPrefix && zone == newZone, diags
	} else {
		// Plain address
		newIPAddr, _ := netip.ParseAddr(newIPv6)
		currentIPAddr, _ := netip.ParseAddr(ipv6)

		return currentIPAddr == newIPAddr && zone == newZone, diags
	}
}

// ValidateAttribute implements attribute value validation. This type requires the value provided to be a String
// value that is a valid IPv6 name.
func (v IPv6Name) ValidateAttribute(ctx context.Context, req xattr.ValidateAttributeRequest, resp *xattr.ValidateAttributeResponse) {
	var (
		ipAddr netip.Addr
		err    error
	)

	if v.IsUnknown() || v.IsNull() {
		return
	}

	ipv6, _, ok := strings.Cut(v.ValueString(), ".")
	if !ok {
		resp.Diagnostics.AddAttributeError(
			req.Path,
			"Invalid IPv6 Name Value",
			"A string value was provided that is not a valid IPv6 name. "+
				"Expected format: '<ipv6>.<rp-zone>' or '<ipv6>/<prefix>.<rp-zone>'.\n\n"+
				"Given Value: "+v.ValueString(),
		)
		return
	}

	if strings.Contains(ipv6, "/") {
		// CIDR
		prefix, err := netip.ParsePrefix(ipv6)
		if err != nil {
			resp.Diagnostics.AddAttributeError(
				req.Path,
				"Invalid IPv6 Name Value",
				"A string value was provided that is not a valid IPv6 name. "+
					"Expected format: '<ipv6>.<rp-zone>' or '<ipv6>/<prefix>.<rp-zone>'.\n\n"+
					"Given Value: "+v.ValueString(),
			)

			return
		}
		ipAddr = prefix.Addr()
	} else {
		// Plain address
		ipAddr, err = netip.ParseAddr(ipv6)
		if err != nil {
			resp.Diagnostics.AddAttributeError(
				req.Path,
				"Invalid IPv6 Name Value",
				"A string value was provided that is not a valid IPv6 name. "+
					"Expected format: '<ipv6>.<rp-zone>' or '<ipv6>/<prefix>.<rp-zone>'.\n\n"+
					"Given Value: "+v.ValueString(),
			)

			return
		}
	}

	if !ipAddr.Is6() {
		resp.Diagnostics.AddAttributeError(
			req.Path,
			"Invalid IPv6 Name Value",
			"A string value was provided that is not a valid IPv6 name. "+
				"The IPv6 portion must be a valid IPv6 address or IPv6 CIDR prefix.\n\n"+
				"Expected format: '<ipv6>.<rp-zone>' or '<ipv6>/<prefix>.<rp-zone>'.\n\n"+
				"Given Value: "+v.ValueString(),
		)
		return
	}
}

// ValidateParameter implements provider-defined function parameter value validation. This type requires the value provided
// to be a String value that is a valid IPv6 name.
func (v IPv6Name) ValidateParameter(ctx context.Context, req function.ValidateParameterRequest, resp *function.ValidateParameterResponse) {
	var (
		ipAddr netip.Addr
		err    error
	)

	if v.IsUnknown() || v.IsNull() {
		return
	}

	ipv6, _, ok := strings.Cut(v.ValueString(), ".")
	if !ok {
		resp.Error = function.NewArgumentFuncError(
			req.Position,
			"Invalid IPv6 Name Value: "+
				"A string value was provided that is not a valid IPv6 name. "+
				"Expected format: '<ipv6>.<rp-zone>' or '<ipv6>/<prefix>.<rp-zone>'.\n\n"+
				"Given Value: "+v.ValueString(),
		)
		return
	}

	if strings.Contains(ipv6, "/") {
		// CIDR
		prefix, err := netip.ParsePrefix(ipv6)
		if err != nil {
			resp.Error = function.NewArgumentFuncError(
				req.Position,
				"Invalid IPv6 Name Value: "+
					"A string value was provided that is not a valid IPv6 name. "+
					"Expected format: '<ipv6>.<rp-zone>' or '<ipv6>/<prefix>.<rp-zone>'.\n\n"+
					"Given Value: "+v.ValueString()+"\n"+
					"Error: "+err.Error(),
			)

			return
		}
		ipAddr = prefix.Addr()
	} else {
		// Plain address
		ipAddr, err = netip.ParseAddr(ipv6)
		if err != nil {
			resp.Error = function.NewArgumentFuncError(
				req.Position,
				"Invalid IPv6 Name Value: "+
					"A string value was provided that is not a valid IPv6 name.\n\n"+
					"Expected format: '<ipv6>.<rp-zone>' or '<ipv6>/<prefix>.<rp-zone>'.\n\n"+
					"Given Value: "+v.ValueString()+"\n"+
					"Error: "+err.Error(),
			)

			return
		}
	}

	if !ipAddr.Is6() {
		resp.Error = function.NewArgumentFuncError(
			req.Position,
			"Invalid IPv6 Name Value: "+
				"A string value was provided that is not a valid IPv6 name.\n\n"+
				"Expected format: '<ipv6>.<rp-zone>' or '<ipv6>/<prefix>.<rp-zone>'.\n\n"+
				"Given Value: "+v.ValueString(),
		)
		return
	}
}
