package keys

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/mapvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/mapdefault"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"

	coremodel "github.com/infobloxopen/terraform-provider-infoblox/internal/core/model/keys"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
)

type KerberosKeyModel struct {
	Id   types.String `tfsdk:"id"`
	UDDI types.Object `tfsdk:"uddi"`
}

var KerberosKeyAttrTypes = map[string]attr.Type{
	"id":   types.StringType,
	"uddi": types.ObjectType{AttrTypes: UDDIKerberosKeyAttrTypes},
}

type UDDIKerberosKeyModel struct {
	Algorithm  types.String `tfsdk:"algorithm"`
	Comment    types.String `tfsdk:"comment"`
	Domain     types.String `tfsdk:"domain"`
	Principal  types.String `tfsdk:"principal"`
	Tags       types.Map    `tfsdk:"tags"`
	TagsAll    types.Map    `tfsdk:"tags_all"`
	UploadedAt types.String `tfsdk:"uploaded_at"`
	Version    types.Int64  `tfsdk:"version"`
}

var UDDIKerberosKeyAttrTypes = map[string]attr.Type{
	"algorithm":   types.StringType,
	"comment":     types.StringType,
	"domain":      types.StringType,
	"principal":   types.StringType,
	"tags":        types.MapType{ElemType: types.StringType},
	"tags_all":    types.MapType{ElemType: types.StringType},
	"uploaded_at": types.StringType,
	"version":     types.Int64Type,
}

const (
	KerberosKeyReturnFields = ""
)

var KerberosKeyResourceSchemaAttributes = map[string]schema.Attribute{
	"id": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The resource identifier.",
	},
	"uddi": schema.SingleNestedAttribute{
		Optional:            true,
		MarkdownDescription: "UDDI backend-specific fields.",
		Attributes:          KerberosKeyResourceUddiSchemaAttributes,
	},
}

var KerberosKeyResourceUddiSchemaAttributes = map[string]schema.Attribute{
	"algorithm": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "Encryption algorithm of the key in accordance with RFC 3961.",
	},
	"comment": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The description for Kerberos key. May contain 0 to 1024 characters. Can include UTF-8.",
	},
	"domain": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "Kerberos realm of the principal.",
	},
	"principal": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "Kerberos principal associated with key.",
	},
	"tags": schema.MapAttribute{
		Optional:    true,
		Computed:    true,
		ElementType: types.StringType,
		Default:     mapdefault.StaticValue(types.MapNull(types.StringType)),
		Validators: []validator.Map{
			mapvalidator.SizeAtLeast(1),
		},
		MarkdownDescription: "The tags for the Kerberos key in JSON format.",
	},
	"tags_all": schema.MapAttribute{
		Computed:            true,
		ElementType:         types.StringType,
		MarkdownDescription: "All tags including inherited values.",
	},
	"uploaded_at": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "Upload time for the key.",
	},
	"version": schema.Int64Attribute{
		Computed:            true,
		MarkdownDescription: "The version number (KVNO) of the key.",
	},
}

// Expand converts the TF model to the infoblox core model
func (m *KerberosKeyModel) Expand(ctx context.Context, diags *diag.Diagnostics, isCreate bool) *coremodel.KerberosKey {
	if m == nil {
		return nil
	}

	obj := &coremodel.KerberosKey{}

	// Expand UDDI nested attribute (returns nil if not present)
	uddiModel := flex.ExpandNestedObject[UDDIKerberosKeyModel](ctx, m.UDDI, diags)
	if uddiModel != nil {
		obj.UDDI = uddiModel.Expand(ctx, diags)
	}

	return obj
}

// Expand converts the UDDI TF model to the core model.
func (m *UDDIKerberosKeyModel) Expand(ctx context.Context, diags *diag.Diagnostics) *coremodel.UDDIKerberosKeyExt {
	return &coremodel.UDDIKerberosKeyExt{
		Comment: flex.ExpandStringPointer(m.Comment),
		Tags:    flex.ExpandMapStringAny(ctx, m.Tags, diags),
	}
}

// Flatten populates the TF model from a core response.
func (m *KerberosKeyModel) Flatten(ctx context.Context, resp *coremodel.KerberosKey, diags *diag.Diagnostics) {
	if resp == nil {
		return
	}

	m.Id = flex.FlattenStringPointer(resp.Id)

	// Extract existing UDDI model, flatten API response onto it, convert back
	uddiModel := flex.ExpandNestedObject[UDDIKerberosKeyModel](ctx, m.UDDI, diags)
	if uddiModel == nil {
		uddiModel = &UDDIKerberosKeyModel{}
	}
	uddiModel.Flatten(ctx, resp.UDDI, diags)
	if resp.UDDI != nil {
		m.UDDI = flex.FlattenNestedObject(ctx, uddiModel, UDDIKerberosKeyAttrTypes, diags)
	} else {
		m.UDDI = types.ObjectNull(UDDIKerberosKeyAttrTypes)
	}
}

// Flatten merges API response onto existing UDDI model.
func (m *UDDIKerberosKeyModel) Flatten(ctx context.Context, from *coremodel.UDDIKerberosKeyExt, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	m.Algorithm = flex.FlattenStringPointer(from.Algorithm)
	m.Comment = flex.FlattenStringPointer(from.Comment)
	m.Domain = flex.FlattenStringPointer(from.Domain)
	m.Principal = flex.FlattenStringPointer(from.Principal)
	tagsAll := flex.FlattenMapStringAny(ctx, from.Tags, diags)
	if m.Tags.IsNull() || m.Tags.IsUnknown() {
		m.Tags = tagsAll
	}
	m.TagsAll = tagsAll
	m.UploadedAt = flex.FlattenStringPointer(from.UploadedAt)
	m.Version = flex.FlattenInt64Pointer(from.Version)
}
