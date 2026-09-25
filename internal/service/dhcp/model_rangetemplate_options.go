package dhcp

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	niosdhcp "github.com/infobloxopen/infoblox-nios-go-client/dhcp"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
	customvalidator "github.com/infobloxopen/terraform-provider-infoblox/internal/validator"
)

// RangetemplateOptionsModel is the Terraform model for RangetemplateOptions
type RangetemplateOptionsModel struct {
	Name        types.String `tfsdk:"name"`
	Num         types.Int64  `tfsdk:"num"`
	VendorClass types.String `tfsdk:"vendor_class"`
	Value       types.String `tfsdk:"value"`
	UseOption   types.Bool   `tfsdk:"use_option"`
}

// RangetemplateOptionsAttrTypes contains the attribute types for RangetemplateOptionsModel
var RangetemplateOptionsAttrTypes = map[string]attr.Type{
	"name":         types.StringType,
	"num":          types.Int64Type,
	"vendor_class": types.StringType,
	"value":        types.StringType,
	"use_option":   types.BoolType,
}

// RangetemplateOptionsResourceSchemaAttributes contains the schema attributes for RangetemplateOptionsModel
var RangetemplateOptionsResourceSchemaAttributes = map[string]schema.Attribute{
	"name": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
			customvalidator.ValidateTrimmedString(),
		},
		MarkdownDescription: "Name of the DHCP option.",
	},
	"num": schema.Int64Attribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The code of the DHCP option.",
	},
	"vendor_class": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
			customvalidator.ValidateTrimmedString(),
		},
		MarkdownDescription: "The name of the space this DHCP option is associated to.",
	},
	"value": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
			customvalidator.ValidateTrimmedString(),
		},
		MarkdownDescription: "Value of the DHCP option",
	},
	"use_option": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "Only applies to special options that are displayed separately from other options and have a use flag. These options are: * routers * router-templates * domain-name-servers * domain-name * broadcast-address * broadcast-address-offset * dhcp-lease-time * dhcp6.name-servers",
	},
}

// ExpandRangetemplateOptions converts a Terraform Object to SDK type
func ExpandRangetemplateOptions(ctx context.Context, o types.Object, diags *diag.Diagnostics) *niosdhcp.RangetemplateOptions {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m RangetemplateOptionsModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

// Expand converts the Terraform model to SDK type
func (m *RangetemplateOptionsModel) Expand(ctx context.Context, diags *diag.Diagnostics) *niosdhcp.RangetemplateOptions {
	if m == nil {
		return nil
	}
	to := &niosdhcp.RangetemplateOptions{
		Name:        flex.ExpandStringPointer(m.Name),
		Num:         flex.ExpandInt64Pointer(m.Num),
		VendorClass: flex.ExpandStringPointer(m.VendorClass),
		Value:       flex.ExpandStringPointerNullAsEmpty(m.Value),
		UseOption:   flex.ExpandBoolPointer(m.UseOption),
	}
	return to
}

// FlattenRangetemplateOptions converts an SDK type to Terraform Object
func FlattenRangetemplateOptions(ctx context.Context, from *niosdhcp.RangetemplateOptions, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(RangetemplateOptionsAttrTypes)
	}
	m := &RangetemplateOptionsModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, RangetemplateOptionsAttrTypes, m)
	diags.Append(d...)
	return t
}

// Flatten populates the Terraform model from SDK type
func (m *RangetemplateOptionsModel) Flatten(ctx context.Context, from *niosdhcp.RangetemplateOptions, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	m.Name = flex.FlattenStringPointerEmptyAsNull(from.Name)
	m.Num = flex.FlattenInt64Pointer(from.Num)
	m.VendorClass = flex.FlattenStringPointerEmptyAsNull(from.VendorClass)
	m.Value = flex.FlattenStringPointerEmptyAsNull(from.Value)
	m.UseOption = flex.FlattenBoolPointer(from.UseOption)
}
