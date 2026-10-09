package dns

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	niosdns "github.com/infobloxopen/infoblox-nios-go-client/dns"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
	customvalidator "github.com/infobloxopen/terraform-provider-infoblox/internal/validator"
)

// ZoneForwardForwardingServersModel is the Terraform model for ZoneForwardForwardingServers
type ZoneForwardForwardingServersModel struct {
	Name                  types.String `tfsdk:"name"`
	ForwardersOnly        types.Bool   `tfsdk:"forwarders_only"`
	ForwardTo             types.List   `tfsdk:"forward_to"`
	UseOverrideForwarders types.Bool   `tfsdk:"use_override_forwarders"`
}

// ZoneForwardForwardingServersAttrTypes contains the attribute types for ZoneForwardForwardingServersModel
var ZoneForwardForwardingServersAttrTypes = map[string]attr.Type{
	"name":                    types.StringType,
	"forwarders_only":         types.BoolType,
	"forward_to":              types.ListType{ElemType: types.ObjectType{AttrTypes: ZoneforwardforwardingserversForwardToAttrTypes}},
	"use_override_forwarders": types.BoolType,
}

// ZoneForwardForwardingServersResourceSchemaAttributes contains the schema attributes for ZoneForwardForwardingServersModel
var ZoneForwardForwardingServersResourceSchemaAttributes = map[string]schema.Attribute{
	"name": schema.StringAttribute{
		Required: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
		},
		MarkdownDescription: "The name of this Grid member in FQDN format.",
	},
	"forwarders_only": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Determines if the appliance sends queries to forwarders only, and not to other internal or Internet root servers.",
	},
	"forward_to": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: ZoneforwardforwardingserversForwardToResourceSchemaAttributes,
		},
		Optional: true,
		Validators: []validator.List{
			customvalidator.ListNotEmpty(),
		},
		MarkdownDescription: "The information for the remote name server to which you want the Infoblox appliance to forward queries for a specified domain name.",
	},
	"use_override_forwarders": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Use flag for: forward_to",
	},
}

// ExpandZoneForwardForwardingServers converts a Terraform Object to SDK type
func ExpandZoneForwardForwardingServers(ctx context.Context, o types.Object, diags *diag.Diagnostics) *niosdns.ZoneForwardForwardingServers {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m ZoneForwardForwardingServersModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

// Expand converts the Terraform model to SDK type
func (m *ZoneForwardForwardingServersModel) Expand(ctx context.Context, diags *diag.Diagnostics) *niosdns.ZoneForwardForwardingServers {
	if m == nil {
		return nil
	}
	to := &niosdns.ZoneForwardForwardingServers{
		Name:                  flex.ExpandStringPointerNullAsEmpty(m.Name),
		ForwardersOnly:        flex.ExpandBoolPointer(m.ForwardersOnly),
		ForwardTo:             flex.ExpandFrameworkListNestedBlock(ctx, m.ForwardTo, diags, ExpandZoneforwardforwardingserversForwardTo),
		UseOverrideForwarders: flex.ExpandBoolPointer(m.UseOverrideForwarders),
	}
	return to
}

// FlattenZoneForwardForwardingServers converts an SDK type to Terraform Object
func FlattenZoneForwardForwardingServers(ctx context.Context, from *niosdns.ZoneForwardForwardingServers, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(ZoneForwardForwardingServersAttrTypes)
	}
	m := &ZoneForwardForwardingServersModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, ZoneForwardForwardingServersAttrTypes, m)
	diags.Append(d...)
	return t
}

// Flatten populates the Terraform model from SDK type
func (m *ZoneForwardForwardingServersModel) Flatten(ctx context.Context, from *niosdns.ZoneForwardForwardingServers, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	m.Name = flex.FlattenStringPointerEmptyAsNull(from.Name)
	m.ForwardersOnly = flex.FlattenBoolPointer(from.ForwardersOnly)
	m.ForwardTo = flex.FlattenFrameworkListNestedBlock(ctx, from.ForwardTo, ZoneforwardforwardingserversForwardToAttrTypes, diags, FlattenZoneforwardforwardingserversForwardTo)
	m.UseOverrideForwarders = flex.FlattenBoolPointer(from.UseOverrideForwarders)
}
