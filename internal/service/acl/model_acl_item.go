package acl

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
	planmod "github.com/infobloxopen/terraform-provider-infoblox/internal/planmodifiers"
	customvalidator "github.com/infobloxopen/terraform-provider-infoblox/internal/validator"
	uddiacl "github.com/infobloxopen/universal-ddi-go-client/dnsconfig"
)

// ACLItemModel is the Terraform model for ACLItem
type ACLItemModel struct {
	Access  types.String `tfsdk:"access"`
	Acl     types.String `tfsdk:"acl"`
	Address types.String `tfsdk:"address"`
	Element types.String `tfsdk:"element"`
	TsigKey types.Object `tfsdk:"tsig_key"`
}

// ACLItemAttrTypes contains the attribute types for ACLItemModel
var ACLItemAttrTypes = map[string]attr.Type{
	"access":   types.StringType,
	"acl":      types.StringType,
	"address":  types.StringType,
	"element":  types.StringType,
	"tsig_key": types.ObjectType{AttrTypes: TSIGKeyAttrTypes},
}

// ACLItemResourceSchemaAttributes contains the schema attributes for ACLItemModel
var ACLItemResourceSchemaAttributes = map[string]schema.Attribute{
	"access": schema.StringAttribute{
		Optional: true,
		Computed: true,
		PlanModifiers: []planmodifier.String{
			planmod.UseEmptyStringForNull(),
		},
		MarkdownDescription: "Access permission for _element_.  Allowed values:  * _allow_,  * _deny_.",
	},
	"acl": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The resource identifier.",
	},
	"address": schema.StringAttribute{
		Optional: true,
		Computed: true,
		PlanModifiers: []planmodifier.String{
			planmod.UseEmptyStringForNull(),
		},
		MarkdownDescription: "Optional. Data for _ip_ _element_.  Must be empty if _element_ is not _ip_.",
	},
	"element": schema.StringAttribute{
		Validators: []validator.String{
			customvalidator.StringNotNull(),
			stringvalidator.OneOf("any", "ip", "acl", "tsig_key"),
		},
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "Type of element.  Allowed values:  * _any_,  * _ip_,  * _acl_,  * _tsig_key_.",
	},
	"tsig_key": schema.SingleNestedAttribute{
		Attributes:          TSIGKeyResourceSchemaAttributes,
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "Optional. TSIG key.  Must be empty if _element_ is not _tsig_key_.",
	},
}

// ExpandACLItem converts a Terraform Object to SDK type
func ExpandACLItem(ctx context.Context, o types.Object, diags *diag.Diagnostics) *uddiacl.ACLItem {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m ACLItemModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

// Expand converts the Terraform model to SDK type
func (m *ACLItemModel) Expand(ctx context.Context, diags *diag.Diagnostics) *uddiacl.ACLItem {
	if m == nil {
		return nil
	}
	to := &uddiacl.ACLItem{
		Access:  flex.ExpandString(m.Access),
		Acl:     flex.ExpandStringPointer(m.Acl),
		Address: flex.ExpandStringPointer(m.Address),
		Element: flex.ExpandString(m.Element),
		TsigKey: ExpandTSIGKey(ctx, m.TsigKey, diags),
	}
	return to
}

// FlattenACLItem converts an SDK type to Terraform Object
func FlattenACLItem(ctx context.Context, from *uddiacl.ACLItem, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(ACLItemAttrTypes)
	}
	m := &ACLItemModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, ACLItemAttrTypes, m)
	diags.Append(d...)
	return t
}

// Flatten populates the Terraform model from SDK type
func (m *ACLItemModel) Flatten(ctx context.Context, from *uddiacl.ACLItem, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	m.Access = flex.FlattenString(from.Access)
	m.Acl = flex.FlattenStringPointer(from.Acl)
	m.Address = flex.FlattenStringPointer(from.Address)
	m.Element = flex.FlattenString(from.Element)
	m.TsigKey = FlattenTSIGKey(ctx, from.TsigKey, diags)
}
