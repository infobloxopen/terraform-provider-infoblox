package anycast

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listdefault"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
	customvalidator "github.com/infobloxopen/terraform-provider-infoblox/internal/validator"
	uddianycast "github.com/infobloxopen/universal-ddi-go-client/anycast"
)

// AnycastConfigRefModel is the Terraform model for AnycastConfigRef
type AnycastConfigRefModel struct {
	AnycastConfigName types.String `tfsdk:"anycast_config_name"`
	RoutingProtocols  types.List   `tfsdk:"routing_protocols"`
}

// AnycastConfigRefAttrTypes contains the attribute types for AnycastConfigRefModel
var AnycastConfigRefAttrTypes = map[string]attr.Type{
	"anycast_config_name": types.StringType,
	"routing_protocols":   types.ListType{ElemType: types.StringType},
}

// AnycastConfigRefResourceSchemaAttributes contains the schema attributes for AnycastConfigRefModel
var AnycastConfigRefResourceSchemaAttributes = map[string]schema.Attribute{
	"anycast_config_name": schema.StringAttribute{
		Required:            true,
		MarkdownDescription: "Anycast Config Name",
	},
	"routing_protocols": schema.ListAttribute{
		ElementType: types.StringType,
		Optional:    true,
		Computed:    true,
		Default:     listdefault.StaticValue(types.ListNull(types.StringType)),
		Validators: []validator.List{
			customvalidator.ListNotEmpty(),
		},
		MarkdownDescription: "Routing protocols enabled for this anycast configuration, on a particular host. Valid protocol names are \"BGP\", \"OSPF\"/\"OSPFv2\", \"OSPFv3\".",
	},
}

// ExpandAnycastConfigRef converts a Terraform Object to SDK type
func ExpandAnycastConfigRef(ctx context.Context, o types.Object, diags *diag.Diagnostics) *uddianycast.AnycastConfigRef {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m AnycastConfigRefModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

// Expand converts the Terraform model to SDK type
func (m *AnycastConfigRefModel) Expand(ctx context.Context, diags *diag.Diagnostics) *uddianycast.AnycastConfigRef {
	if m == nil {
		return nil
	}
	to := &uddianycast.AnycastConfigRef{
		AnycastConfigName: flex.ExpandStringPointer(m.AnycastConfigName),
		RoutingProtocols:  flex.ExpandFrameworkListString(ctx, m.RoutingProtocols, diags),
	}
	return to
}

// FlattenAnycastConfigRef converts an SDK type to Terraform Object
func FlattenAnycastConfigRef(ctx context.Context, from *uddianycast.AnycastConfigRef, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(AnycastConfigRefAttrTypes)
	}
	m := &AnycastConfigRefModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, AnycastConfigRefAttrTypes, m)
	diags.Append(d...)
	return t
}

// Flatten populates the Terraform model from SDK type
func (m *AnycastConfigRefModel) Flatten(ctx context.Context, from *uddianycast.AnycastConfigRef, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	m.AnycastConfigName = flex.FlattenStringPointer(from.AnycastConfigName)
	m.RoutingProtocols = flex.FlattenFrameworkListString(ctx, from.RoutingProtocols, diags)
}
