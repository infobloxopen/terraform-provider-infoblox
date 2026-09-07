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

// RangetemplateMsOptionsModel is the Terraform model for RangetemplateMsOptions
type RangetemplateMsOptionsModel struct {
	Num         types.Int64  `tfsdk:"num"`
	Value       types.String `tfsdk:"value"`
	Name        types.String `tfsdk:"name"`
	VendorClass types.String `tfsdk:"vendor_class"`
	UserClass   types.String `tfsdk:"user_class"`
	Type        types.String `tfsdk:"type"`
}

// RangetemplateMsOptionsAttrTypes contains the attribute types for RangetemplateMsOptionsModel
var RangetemplateMsOptionsAttrTypes = map[string]attr.Type{
	"num":          types.Int64Type,
	"value":        types.StringType,
	"name":         types.StringType,
	"vendor_class": types.StringType,
	"user_class":   types.StringType,
	"type":         types.StringType,
}

// RangetemplateMsOptionsResourceSchemaAttributes contains the schema attributes for RangetemplateMsOptionsModel
var RangetemplateMsOptionsResourceSchemaAttributes = map[string]schema.Attribute{
	"num": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "The code of the DHCP option.",
	},
	"value": schema.StringAttribute{
		Optional: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
		},
		MarkdownDescription: "Value of the DHCP option.",
	},
	"name": schema.StringAttribute{
		Optional: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
		},
		MarkdownDescription: "The name of the DHCP option.",
	},
	"vendor_class": schema.StringAttribute{
		Optional: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
		},
		MarkdownDescription: "The name of the vendor class with which this DHCP option is associated.",
	},
	"user_class": schema.StringAttribute{
		Optional: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
		},
		MarkdownDescription: "The name of the user class with which this DHCP option is associated.",
	},
	"type": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
		},
		MarkdownDescription: "The DHCP option type. Valid values are: * \"16-bit signed integer\" * \"16-bit unsigned integer\" * \"32-bit signed integer\" * \"32-bit unsigned integer\" * \"64-bit unsigned integer\" * \"8-bit signed integer\" * \"8-bit unsigned integer (1,2,4,8)\" * \"8-bit unsigned integer\" * \"array of 16-bit integer\" * \"array of 16-bit unsigned integer\" * \"array of 32-bit integer\" * \"array of 32-bit unsigned integer\" * \"array of 64-bit unsigned integer\" * \"array of 8-bit integer\" * \"array of 8-bit unsigned integer\" * \"array of ip-address pair\" * \"array of ip-address\" * \"array of string\" * \"binary\" * \"boolean array of ip-address\" * \"boolean\" * \"boolean-text\" * \"domain-list\" * \"domain-name\" * \"encapsulated\" * \"ip-address\" * \"string\" * \"text\"",
	},
}

// ExpandRangetemplateMsOptions converts a Terraform Object to SDK type
func ExpandRangetemplateMsOptions(ctx context.Context, o types.Object, diags *diag.Diagnostics) *niosdhcp.RangetemplateMsOptions {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m RangetemplateMsOptionsModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

// Expand converts the Terraform model to SDK type
func (m *RangetemplateMsOptionsModel) Expand(ctx context.Context, diags *diag.Diagnostics) *niosdhcp.RangetemplateMsOptions {
	if m == nil {
		return nil
	}
	to := &niosdhcp.RangetemplateMsOptions{
		Num:         flex.ExpandInt64Pointer(m.Num),
		Value:       flex.ExpandStringPointerNullAsEmpty(m.Value),
		Name:        flex.ExpandStringPointerNullAsEmpty(m.Name),
		VendorClass: flex.ExpandStringPointerNullAsEmpty(m.VendorClass),
		UserClass:   flex.ExpandStringPointerNullAsEmpty(m.UserClass),
		Type:        flex.ExpandStringPointerNullAsEmpty(m.Type),
	}
	return to
}

// FlattenRangetemplateMsOptions converts an SDK type to Terraform Object
func FlattenRangetemplateMsOptions(ctx context.Context, from *niosdhcp.RangetemplateMsOptions, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(RangetemplateMsOptionsAttrTypes)
	}
	m := &RangetemplateMsOptionsModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, RangetemplateMsOptionsAttrTypes, m)
	diags.Append(d...)
	return t
}

// Flatten populates the Terraform model from SDK type
func (m *RangetemplateMsOptionsModel) Flatten(ctx context.Context, from *niosdhcp.RangetemplateMsOptions, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	m.Num = flex.FlattenInt64Pointer(from.Num)
	m.Value = flex.FlattenStringPointerEmptyAsNull(from.Value)
	m.Name = flex.FlattenStringPointerEmptyAsNull(from.Name)
	m.VendorClass = flex.FlattenStringPointerEmptyAsNull(from.VendorClass)
	m.UserClass = flex.FlattenStringPointerEmptyAsNull(from.UserClass)
	m.Type = flex.FlattenStringPointerEmptyAsNull(from.Type)
}
