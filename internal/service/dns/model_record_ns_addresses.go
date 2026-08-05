package dns

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/hashicorp/terraform-plugin-framework-nettypes/iptypes"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	niosdns "github.com/infobloxopen/infoblox-nios-go-client/dns"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
	customvalidator "github.com/infobloxopen/terraform-provider-infoblox/internal/validator"
)

// RecordNsAddressesModel is the Terraform model for RecordNsAddresses
type RecordNsAddressesModel struct {
	Address       iptypes.IPAddress `tfsdk:"address"`
	AutoCreatePtr types.Bool        `tfsdk:"auto_create_ptr"`
}

// RecordNsAddressesAttrTypes contains the attribute types for RecordNsAddressesModel
var RecordNsAddressesAttrTypes = map[string]attr.Type{
	"address":         iptypes.IPAddressType{},
	"auto_create_ptr": types.BoolType,
}

// RecordNsAddressesResourceSchemaAttributes contains the schema attributes for RecordNsAddressesModel
var RecordNsAddressesResourceSchemaAttributes = map[string]schema.Attribute{
	"address": schema.StringAttribute{
		Required:   true,
		CustomType: iptypes.IPAddressType{},
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
		},
		MarkdownDescription: "The address of the Zone Name Server.",
	},
	"auto_create_ptr": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(true),
		MarkdownDescription: "Flag to indicate if ptr records need to be auto created.",
	},
}

// ExpandRecordNsAddresses converts a Terraform Object to SDK type
func ExpandRecordNsAddresses(ctx context.Context, o types.Object, diags *diag.Diagnostics) *niosdns.RecordNsAddresses {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m RecordNsAddressesModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

// Expand converts the Terraform model to SDK type
func (m *RecordNsAddressesModel) Expand(ctx context.Context, diags *diag.Diagnostics) *niosdns.RecordNsAddresses {
	if m == nil {
		return nil
	}
	to := &niosdns.RecordNsAddresses{
		Address:       flex.ExpandIPAddress(m.Address),
		AutoCreatePtr: flex.ExpandBoolPointer(m.AutoCreatePtr),
	}
	return to
}

// FlattenRecordNsAddresses converts an SDK type to Terraform Object
func FlattenRecordNsAddresses(ctx context.Context, from *niosdns.RecordNsAddresses, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(RecordNsAddressesAttrTypes)
	}
	m := &RecordNsAddressesModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, RecordNsAddressesAttrTypes, m)
	diags.Append(d...)
	return t
}

// Flatten populates the Terraform model from SDK type
func (m *RecordNsAddressesModel) Flatten(ctx context.Context, from *niosdns.RecordNsAddresses, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	m.Address = flex.FlattenIPAddress(from.Address)
	m.AutoCreatePtr = flex.FlattenBoolPointer(from.AutoCreatePtr)
}
