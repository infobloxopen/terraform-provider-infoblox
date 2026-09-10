package dns

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	niosdns "github.com/infobloxopen/infoblox-nios-go-client/dns"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
	customvalidator "github.com/infobloxopen/terraform-provider-infoblox/internal/validator"
)

// RecordHostIpv4addrOptionsModel is the Terraform model for RecordHostIpv4addrOptions
type RecordHostIpv4addrOptionsModel struct {
	Name        types.String `tfsdk:"name"`
	Num         types.Int64  `tfsdk:"num"`
	VendorClass types.String `tfsdk:"vendor_class"`
	Value       types.String `tfsdk:"value"`
	UseOption   types.Bool   `tfsdk:"use_option"`
}

// RecordHostIpv4addrOptionsAttrTypes contains the attribute types for RecordHostIpv4addrOptionsModel
var RecordHostIpv4addrOptionsAttrTypes = map[string]attr.Type{
	"name":         types.StringType,
	"num":          types.Int64Type,
	"vendor_class": types.StringType,
	"value":        types.StringType,
	"use_option":   types.BoolType,
}

// RecordHostIpv4addrOptionsResourceSchemaAttributes contains the schema attributes for RecordHostIpv4addrOptionsModel
var RecordHostIpv4addrOptionsResourceSchemaAttributes = map[string]schema.Attribute{
	"name": schema.StringAttribute{
		Optional: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
		},
		MarkdownDescription: "Name of the DHCP option.",
	},
	"num": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "The code of the DHCP option.",
	},
	"vendor_class": schema.StringAttribute{
		Optional: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
		},
		MarkdownDescription: "The name of the space this DHCP option is associated to.",
	},
	"value": schema.StringAttribute{
		Optional: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
		},
		MarkdownDescription: "Value of the DHCP option",
	},
	"use_option": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Only applies to special options that are displayed separately from other options and have a use flag. These options are: * routers * router-templates * domain-name-servers * domain-name * broadcast-address * broadcast-address-offset * dhcp-lease-time * dhcp6.name-servers",
	},
}

// ExpandRecordHostIpv4addrOptions converts a Terraform Object to SDK type
func ExpandRecordHostIpv4addrOptions(ctx context.Context, o types.Object, diags *diag.Diagnostics) *niosdns.RecordHostIpv4addrOptions {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m RecordHostIpv4addrOptionsModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

// Expand converts the Terraform model to SDK type
func (m *RecordHostIpv4addrOptionsModel) Expand(ctx context.Context, diags *diag.Diagnostics) *niosdns.RecordHostIpv4addrOptions {
	if m == nil {
		return nil
	}
	to := &niosdns.RecordHostIpv4addrOptions{
		Name:        flex.ExpandStringPointerNullAsEmpty(m.Name),
		Num:         flex.ExpandInt64Pointer(m.Num),
		VendorClass: flex.ExpandStringPointerNullAsEmpty(m.VendorClass),
		Value:       flex.ExpandStringPointerNullAsEmpty(m.Value),
		UseOption:   flex.ExpandBoolPointer(m.UseOption),
	}
	return to
}

// FlattenRecordHostIpv4addrOptions converts an SDK type to Terraform Object
func FlattenRecordHostIpv4addrOptions(ctx context.Context, from *niosdns.RecordHostIpv4addrOptions, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(RecordHostIpv4addrOptionsAttrTypes)
	}
	m := &RecordHostIpv4addrOptionsModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, RecordHostIpv4addrOptionsAttrTypes, m)
	diags.Append(d...)
	return t
}

// Flatten populates the Terraform model from SDK type
func (m *RecordHostIpv4addrOptionsModel) Flatten(ctx context.Context, from *niosdns.RecordHostIpv4addrOptions, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	m.Name = flex.FlattenStringPointerEmptyAsNull(from.Name)
	m.Num = flex.FlattenInt64Pointer(from.Num)
	m.VendorClass = flex.FlattenStringPointerEmptyAsNull(from.VendorClass)
	m.Value = flex.FlattenStringPointerEmptyAsNull(from.Value)
	m.UseOption = flex.FlattenBoolPointer(from.UseOption)
}
