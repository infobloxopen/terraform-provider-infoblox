package security

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	niossecurity "github.com/infobloxopen/infoblox-nios-go-client/security"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
	customvalidator "github.com/infobloxopen/terraform-provider-infoblox/internal/validator"
)

// AdminuserSshKeysModel is the Terraform model for AdminuserSshKeys
type AdminuserSshKeysModel struct {
	KeyName  types.String `tfsdk:"key_name"`
	KeyType  types.String `tfsdk:"key_type"`
	KeyValue types.String `tfsdk:"key_value"`
}

// AdminuserSshKeysAttrTypes contains the attribute types for AdminuserSshKeysModel
var AdminuserSshKeysAttrTypes = map[string]attr.Type{
	"key_name":  types.StringType,
	"key_type":  types.StringType,
	"key_value": types.StringType,
}

// AdminuserSshKeysResourceSchemaAttributes contains the schema attributes for AdminuserSshKeysModel
var AdminuserSshKeysResourceSchemaAttributes = map[string]schema.Attribute{
	"key_name": schema.StringAttribute{
		Optional: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
		},
		MarkdownDescription: "Unique identifier for the key",
	},
	"key_type": schema.StringAttribute{
		Validators: []validator.String{
			stringvalidator.OneOf("RSA", "ECDSA", "ED25519"),
		},
		Optional:            true,
		MarkdownDescription: "ssh_key_types",
	},
	"key_value": schema.StringAttribute{
		Optional: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
		},
		MarkdownDescription: "ssh key text",
	},
}

// ExpandAdminuserSshKeys converts a Terraform Object to SDK type
func ExpandAdminuserSshKeys(ctx context.Context, o types.Object, diags *diag.Diagnostics) *niossecurity.AdminuserSshKeys {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m AdminuserSshKeysModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

// Expand converts the Terraform model to SDK type
func (m *AdminuserSshKeysModel) Expand(ctx context.Context, diags *diag.Diagnostics) *niossecurity.AdminuserSshKeys {
	if m == nil {
		return nil
	}
	to := &niossecurity.AdminuserSshKeys{
		KeyName:  flex.ExpandStringPointerNullAsEmpty(m.KeyName),
		KeyType:  flex.ExpandStringPointerNullAsEmpty(m.KeyType),
		KeyValue: flex.ExpandStringPointerNullAsEmpty(m.KeyValue),
	}
	return to
}

// FlattenAdminuserSshKeys converts an SDK type to Terraform Object
func FlattenAdminuserSshKeys(ctx context.Context, from *niossecurity.AdminuserSshKeys, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(AdminuserSshKeysAttrTypes)
	}
	m := &AdminuserSshKeysModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, AdminuserSshKeysAttrTypes, m)
	diags.Append(d...)
	return t
}

// Flatten populates the Terraform model from SDK type
func (m *AdminuserSshKeysModel) Flatten(ctx context.Context, from *niossecurity.AdminuserSshKeys, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	m.KeyName = flex.FlattenStringPointerEmptyAsNull(from.KeyName)
	m.KeyType = flex.FlattenStringPointerEmptyAsNull(from.KeyType)
	m.KeyValue = flex.FlattenStringPointerEmptyAsNull(from.KeyValue)
}
