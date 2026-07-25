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
	niosdns "github.com/infobloxopen/infoblox-nios-go-client/dns"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
	customvalidator "github.com/infobloxopen/terraform-provider-infoblox/internal/validator"
)

// ZoneAuthMsDcNsRecordCreationModel is the Terraform model for ZoneAuthMsDcNsRecordCreation
type ZoneAuthMsDcNsRecordCreationModel struct {
	Address iptypes.IPv4Address `tfsdk:"address"`
	Comment types.String        `tfsdk:"comment"`
}

// ZoneAuthMsDcNsRecordCreationAttrTypes contains the attribute types for ZoneAuthMsDcNsRecordCreationModel
var ZoneAuthMsDcNsRecordCreationAttrTypes = map[string]attr.Type{
	"address": iptypes.IPv4AddressType{},
	"comment": types.StringType,
}

// ZoneAuthMsDcNsRecordCreationResourceSchemaAttributes contains the schema attributes for ZoneAuthMsDcNsRecordCreationModel
var ZoneAuthMsDcNsRecordCreationResourceSchemaAttributes = map[string]schema.Attribute{
	"address": schema.StringAttribute{
		Required:   true,
		CustomType: iptypes.IPv4AddressType{},
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
		},
		MarkdownDescription: "The IPv4 address of the domain controller that is allowed to create NS records.",
	},
	"comment": schema.StringAttribute{
		Optional: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
			customvalidator.ValidateTrimmedString(),
		},
		MarkdownDescription: "Optional user comment.",
	},
}

// ExpandZoneAuthMsDcNsRecordCreation converts a Terraform Object to SDK type
func ExpandZoneAuthMsDcNsRecordCreation(ctx context.Context, o types.Object, diags *diag.Diagnostics) *niosdns.ZoneAuthMsDcNsRecordCreation {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m ZoneAuthMsDcNsRecordCreationModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

// Expand converts the Terraform model to SDK type
func (m *ZoneAuthMsDcNsRecordCreationModel) Expand(ctx context.Context, diags *diag.Diagnostics) *niosdns.ZoneAuthMsDcNsRecordCreation {
	if m == nil {
		return nil
	}
	to := &niosdns.ZoneAuthMsDcNsRecordCreation{
		Address: flex.ExpandIPv4Address(m.Address),
		Comment: flex.ExpandStringPointerNullAsEmpty(m.Comment),
	}
	return to
}

// FlattenZoneAuthMsDcNsRecordCreation converts an SDK type to Terraform Object
func FlattenZoneAuthMsDcNsRecordCreation(ctx context.Context, from *niosdns.ZoneAuthMsDcNsRecordCreation, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(ZoneAuthMsDcNsRecordCreationAttrTypes)
	}
	m := &ZoneAuthMsDcNsRecordCreationModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, ZoneAuthMsDcNsRecordCreationAttrTypes, m)
	diags.Append(d...)
	return t
}

// Flatten populates the Terraform model from SDK type
func (m *ZoneAuthMsDcNsRecordCreationModel) Flatten(ctx context.Context, from *niosdns.ZoneAuthMsDcNsRecordCreation, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	m.Address = flex.FlattenIPv4Address(from.Address)
	m.Comment = flex.FlattenStringPointerEmptyAsNull(from.Comment)
}
