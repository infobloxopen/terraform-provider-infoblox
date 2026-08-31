package dns

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
	uddidns "github.com/infobloxopen/universal-ddi-go-client/dnsconfig"
)

// KerberosKeyModel is the Terraform model for KerberosKey
type KerberosKeyModel struct {
	Algorithm  types.String `tfsdk:"algorithm"`
	Domain     types.String `tfsdk:"domain"`
	Key        types.String `tfsdk:"key"`
	Principal  types.String `tfsdk:"principal"`
	UploadedAt types.String `tfsdk:"uploaded_at"`
	Version    types.Int64  `tfsdk:"version"`
}

// KerberosKeyAttrTypes contains the attribute types for KerberosKeyModel
var KerberosKeyAttrTypes = map[string]attr.Type{
	"algorithm":   types.StringType,
	"domain":      types.StringType,
	"key":         types.StringType,
	"principal":   types.StringType,
	"uploaded_at": types.StringType,
	"version":     types.Int64Type,
}

// KerberosKeyResourceSchemaAttributes contains the schema attributes for KerberosKeyModel
var KerberosKeyResourceSchemaAttributes = map[string]schema.Attribute{
	"algorithm": schema.StringAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "Encryption algorithm of the key in accordance with RFC 3961.",
	},
	"domain": schema.StringAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "Kerberos realm of the principal.",
	},
	"key": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The resource identifier.",
	},
	"principal": schema.StringAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "Kerberos principal associated with key.",
	},
	"uploaded_at": schema.StringAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "Upload time for the key.",
	},
	"version": schema.Int64Attribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The version number (KVNO) of the key.",
	},
}

// ExpandKerberosKey converts a Terraform Object to SDK type
func ExpandKerberosKey(ctx context.Context, o types.Object, diags *diag.Diagnostics) *uddidns.KerberosKey {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m KerberosKeyModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

// Expand converts the Terraform model to SDK type
func (m *KerberosKeyModel) Expand(ctx context.Context, diags *diag.Diagnostics) *uddidns.KerberosKey {
	if m == nil {
		return nil
	}
	to := &uddidns.KerberosKey{
		Algorithm:  flex.ExpandStringPointer(m.Algorithm),
		Domain:     flex.ExpandStringPointer(m.Domain),
		Key:        flex.ExpandString(m.Key),
		Principal:  flex.ExpandStringPointer(m.Principal),
		UploadedAt: flex.ExpandStringPointer(m.UploadedAt),
		Version:    flex.ExpandInt64Pointer(m.Version),
	}
	return to
}

// FlattenKerberosKey converts an SDK type to Terraform Object
func FlattenKerberosKey(ctx context.Context, from *uddidns.KerberosKey, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(KerberosKeyAttrTypes)
	}
	m := &KerberosKeyModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, KerberosKeyAttrTypes, m)
	diags.Append(d...)
	return t
}

// Flatten populates the Terraform model from SDK type
func (m *KerberosKeyModel) Flatten(ctx context.Context, from *uddidns.KerberosKey, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	m.Algorithm = flex.FlattenStringPointer(from.Algorithm)
	m.Domain = flex.FlattenStringPointer(from.Domain)
	m.Key = flex.FlattenString(from.Key)
	m.Principal = flex.FlattenStringPointer(from.Principal)
	m.UploadedAt = flex.FlattenStringPointer(from.UploadedAt)
	m.Version = flex.FlattenInt64Pointer(from.Version)
}
