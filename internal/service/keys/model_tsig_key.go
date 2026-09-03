package keys

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/mapvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/mapdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"

	coremodel "github.com/infobloxopen/terraform-provider-infoblox/internal/core/model/keys"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
	customvalidator "github.com/infobloxopen/terraform-provider-infoblox/internal/validator"
)

type TsigKeyModel struct {
	Id   types.String `tfsdk:"id"`
	UDDI types.Object `tfsdk:"uddi"`
}

var TsigKeyAttrTypes = map[string]attr.Type{
	"id":   types.StringType,
	"uddi": types.ObjectType{AttrTypes: UDDITsigKeyAttrTypes},
}

type UDDITsigKeyModel struct {
	Algorithm types.String `tfsdk:"algorithm"`
	Comment   types.String `tfsdk:"comment"`
	Name      types.String `tfsdk:"name"`
	Secret    types.String `tfsdk:"secret"`
	Tags      types.Map    `tfsdk:"tags"`
	TagsAll   types.Map    `tfsdk:"tags_all"`
}

var UDDITsigKeyAttrTypes = map[string]attr.Type{
	"algorithm": types.StringType,
	"comment":   types.StringType,
	"name":      types.StringType,
	"secret":    types.StringType,
	"tags":      types.MapType{ElemType: types.StringType},
	"tags_all":  types.MapType{ElemType: types.StringType},
}

const (
	TsigKeyReturnFields = ""
)

var TsigKeyResourceSchemaAttributes = map[string]schema.Attribute{
	"id": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The resource identifier.",
	},
	"uddi": schema.SingleNestedAttribute{
		Optional:            true,
		MarkdownDescription: "UDDI backend-specific fields.",
		Attributes:          TsigKeyResourceUddiSchemaAttributes,
	},
}

var TsigKeyResourceUddiSchemaAttributes = map[string]schema.Attribute{
	"algorithm": schema.StringAttribute{
		Default: stringdefault.StaticString("hmac_sha256"),
		Validators: []validator.String{
			stringvalidator.OneOf("hmac_sha1", "hmac_sha224", "hmac_sha256", "hmac_sha384", "hmac_sha512"),
		},
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The TSIG key algorithm.  Valid values are: * _hmac_sha1_ * _hmac_sha224_ * _hmac_sha256_ * _hmac_sha384_ * _hmac_sha512_  Defaults to _hmac_sha256_.",
	},
	"comment": schema.StringAttribute{
		Default:  stringdefault.StaticString(""),
		Optional: true,
		Computed: true,
		Validators: []validator.String{
			stringvalidator.LengthBetween(0, 1024),
		},
		MarkdownDescription: "The description for the TSIG key. May contain 0 to 1024 characters. Can include UTF-8.",
	},
	"name": schema.StringAttribute{
		Required: true,
		Validators: []validator.String{
			customvalidator.IsValidUDDIDomainName(),
		},
		MarkdownDescription: "The TSIG key name in the absolute domain name format.",
	},
	"secret": schema.StringAttribute{
		Sensitive:           true,
		Required:            true,
		MarkdownDescription: "The TSIG key secret as a Base64 encoded string.",
	},
	"tags": schema.MapAttribute{
		Optional:    true,
		Computed:    true,
		ElementType: types.StringType,
		Default:     mapdefault.StaticValue(types.MapNull(types.StringType)),
		Validators: []validator.Map{
			mapvalidator.SizeAtLeast(1),
		},
		MarkdownDescription: "The tags for the TSIG key in JSON format.",
	},
	"tags_all": schema.MapAttribute{
		Computed:            true,
		ElementType:         types.StringType,
		MarkdownDescription: "All tags including inherited values.",
	},
}

// Expand converts the TF model to the infoblox core model
func (m *TsigKeyModel) Expand(ctx context.Context, diags *diag.Diagnostics, isCreate bool) *coremodel.TsigKey {
	if m == nil {
		return nil
	}

	obj := &coremodel.TsigKey{}

	// Expand UDDI nested attribute (returns nil if not present)
	uddiModel := flex.ExpandNestedObject[UDDITsigKeyModel](ctx, m.UDDI, diags)
	if uddiModel != nil {
		obj.UDDI = uddiModel.Expand(ctx, diags)
	}

	return obj
}

// Expand converts the UDDI TF model to the core model.
func (m *UDDITsigKeyModel) Expand(ctx context.Context, diags *diag.Diagnostics) *coremodel.UDDITsigKeyExt {
	return &coremodel.UDDITsigKeyExt{
		Algorithm: flex.ExpandStringPointer(m.Algorithm),
		Comment:   flex.ExpandStringPointer(m.Comment),
		Name:      flex.ExpandString(m.Name),
		Secret:    flex.ExpandString(m.Secret),
		Tags:      flex.ExpandMapStringAny(ctx, m.Tags, diags),
	}
}

// Flatten populates the TF model from a core response.
func (m *TsigKeyModel) Flatten(ctx context.Context, resp *coremodel.TsigKey, diags *diag.Diagnostics) {
	if resp == nil {
		return
	}

	m.Id = flex.FlattenStringPointer(resp.Id)

	// Extract existing UDDI model, flatten API response onto it, convert back
	uddiModel := flex.ExpandNestedObject[UDDITsigKeyModel](ctx, m.UDDI, diags)
	if uddiModel == nil {
		uddiModel = &UDDITsigKeyModel{}
	}
	uddiModel.Flatten(ctx, resp.UDDI, diags)
	if resp.UDDI != nil {
		m.UDDI = flex.FlattenNestedObject(ctx, uddiModel, UDDITsigKeyAttrTypes, diags)
	} else {
		m.UDDI = types.ObjectNull(UDDITsigKeyAttrTypes)
	}
}

// Flatten merges API response onto existing UDDI model.
func (m *UDDITsigKeyModel) Flatten(ctx context.Context, from *coremodel.UDDITsigKeyExt, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	m.Algorithm = flex.FlattenStringPointer(from.Algorithm)
	m.Comment = flex.FlattenStringPointer(from.Comment)
	m.Name = flex.FlattenString(from.Name)
	m.Secret = flex.FlattenString(from.Secret)
	tagsAll := flex.FlattenMapStringAny(ctx, from.Tags, diags)
	if m.Tags.IsNull() || m.Tags.IsUnknown() {
		m.Tags = tagsAll
	}
	m.TagsAll = tagsAll
}
