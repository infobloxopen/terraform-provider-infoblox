package dns

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
)

// ValidateRecordCaa validates the RecordCaa configuration.
func ValidateRecordCaa(ctx context.Context, data RecordCaaModel, resp *resource.ValidateConfigResponse) {
	if nios := flex.ExpandNestedObject[NIOSRecordCaaModel](ctx, data.NIOS, &resp.Diagnostics); nios != nil {
		validateRecordCaaNIOSConfig(ctx, nios, resp)
	}
	if uddi := flex.ExpandNestedObject[UDDIRecordCaaModel](ctx, data.UDDI, &resp.Diagnostics); uddi != nil {
		validateRecordCaaUDDIConfig(ctx, uddi, resp)
	}
}

func validateRecordCaaNIOSConfig(ctx context.Context, m *NIOSRecordCaaModel, resp *resource.ValidateConfigResponse) {
}

func validateRecordCaaUDDIConfig(ctx context.Context, m *UDDIRecordCaaModel, resp *resource.ValidateConfigResponse) {
}

type UDDIRecordCaaRdataModel struct {
	Flags types.Int64  `tfsdk:"flags"`
	Tag   types.String `tfsdk:"tag"`
	Value types.String `tfsdk:"value"`
}

var UDDIRecordCaaRdataAttrTypes = map[string]attr.Type{
	"flags": types.Int64Type,
	"tag":   types.StringType,
	"value": types.StringType,
}

var UDDIRecordCaaRdataResourceSchemaAttributes = map[string]schema.Attribute{
	"flags": schema.Int64Attribute{
		Optional: true,
		Computed: true,
		Default:  int64default.StaticInt64(0),
		Validators: []validator.Int64{
			int64validator.Between(0, 255),
		},
		MarkdownDescription: "An unsigned 8-bit integer which specifies the CAA record flags. RFC 6844 defines one (highest) bit in flag octet, remaining bits are deferred for future use. This bit is referenced as Critical. When the bit is set (flag value == 128), issuers must not issue certificates in case CAA records contain unknown property tags.",
	},
	"tag": schema.StringAttribute{
		Required:            true,
		MarkdownDescription: "The CAA record property tag string which indicates the type of CAA record. The following property tags are defined by RFC 6844:\nissue: Used to explicitly authorize CA to issue certificates for the domain in which the property is published.\nissuewild: Used to explicitly authorize a single CA to issue wildcard certificates for the domain in which the property is published.\niodef: Used to specify an email address or URL to report invalid certificate requests or issuers' certificate policy violations.\nNote: issuewild type takes precedence over issue.",
	},
	"value": schema.StringAttribute{
		Required:            true,
		MarkdownDescription: "A string which contains the CAA record property value.\n\nSpecifies the CA who is authorized to issue a certificate for the domain if the CAA record property tag is issue or issuewild.\n\nSpecifies the URL/email address to report CAA policy violation for the domain if the CAA record property tag is iodef.",
	},
}

func ExpandUDDIRecordCaaRdata(ctx context.Context, o types.Object, diags *diag.Diagnostics) map[string]any {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m UDDIRecordCaaRdataModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	rdata := map[string]any{
		"tag":   flex.ExpandString(m.Tag),
		"value": flex.ExpandString(m.Value),
		"flags": flex.ExpandInt64(m.Flags),
	}
	return rdata
}

func FlattenUDDIRecordCaaRdata(ctx context.Context, from map[string]any, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(UDDIRecordCaaRdataAttrTypes)
	}
	m := UDDIRecordCaaRdataModel{
		Flags: flex.FlattenInt64Pointer(flex.RDataInt64Ptr(from["flags"])),
		Tag:   flex.FlattenStringPointer(flex.RDataStringPtr(from["tag"])),
		Value: flex.FlattenStringPointer(flex.RDataStringPtr(from["value"])),
	}
	obj, d := types.ObjectValueFrom(ctx, UDDIRecordCaaRdataAttrTypes, m)
	diags.Append(d...)
	return obj
}
