package dhcp

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	coremodel "github.com/infobloxopen/terraform-provider-infoblox/internal/core/model/dhcp"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
)

// ValidateHardwareFilter validates the HardwareFilter configuration.
func ValidateHardwareFilter(ctx context.Context, data HardwareFilterModel, resp *resource.ValidateConfigResponse) {
	if uddi := flex.ExpandNestedObject[UDDIHardwareFilterModel](ctx, data.UDDI, &resp.Diagnostics); uddi != nil {
		validateHardwareFilterUDDIConfig(ctx, uddi, resp)
	}
}

func validateHardwareFilterUDDIConfig(ctx context.Context, m *UDDIHardwareFilterModel, resp *resource.ValidateConfigResponse) {
}

// PostExpandHardwareFilterUDDI clears the OptionItem fields that don't apply to
// its type. The framework doesn't mark unconfigured Optional+Computed attributes
// nested inside a list (dhcp_options[*].option_value, option_code, group) as
// unknown when the sibling discriminator (type) switches, so a stale plan value
// can otherwise be sent for the no-longer-applicable field, which the API rejects.
func PostExpandHardwareFilterUDDI(ctx context.Context, ext *coremodel.UDDIHardwareFilterExt, diags *diag.Diagnostics) *coremodel.UDDIHardwareFilterExt {
	if ext == nil {
		return nil
	}
	for i := range ext.DhcpOptions {
		opt := &ext.DhcpOptions[i]
		if opt.Type == nil {
			continue
		}
		switch *opt.Type {
		case "group":
			opt.OptionCode = nil
			opt.OptionValue = nil
		case "option":
			opt.Group = nil
		}
	}
	return ext
}

// PostFlattenHardwareFilterUDDI nulls out the OptionItem fields that don't apply
// to its type. The API echoes "" for the unset fields on a group-type entry
// (and vice versa) instead of omitting them, which would otherwise conflict with
// the null value PostExpandHardwareFilterUDDI caused to be sent.
func PostFlattenHardwareFilterUDDI(ctx context.Context, planned, flattened *UDDIHardwareFilterModel, diags *diag.Diagnostics) {
	if flattened == nil || flattened.DhcpOptions.IsNull() || flattened.DhcpOptions.IsUnknown() {
		return
	}
	elements := flattened.DhcpOptions.Elements()
	newElements := make([]attr.Value, len(elements))
	for i, elem := range elements {
		obj, ok := elem.(types.Object)
		if !ok {
			newElements[i] = elem
			continue
		}
		attrs := obj.Attributes()
		typeVal, ok := attrs["type"].(types.String)
		if !ok || typeVal.IsNull() || typeVal.IsUnknown() {
			newElements[i] = elem
			continue
		}
		newAttrs := make(map[string]attr.Value, len(attrs))
		for k, v := range attrs {
			newAttrs[k] = v
		}
		switch typeVal.ValueString() {
		case "group":
			newAttrs["option_code"] = types.StringNull()
			newAttrs["option_value"] = types.StringNull()
		case "option":
			newAttrs["group"] = types.StringNull()
		}
		newObj, d := types.ObjectValue(OptionItemAttrTypes, newAttrs)
		diags.Append(d...)
		newElements[i] = newObj
	}
	newList, d := types.ListValue(types.ObjectType{AttrTypes: OptionItemAttrTypes}, newElements)
	diags.Append(d...)
	flattened.DhcpOptions = newList
}
