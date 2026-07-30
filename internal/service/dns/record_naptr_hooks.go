package dns

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
)

// ValidateRecordNaptr validates the RecordNaptr configuration.
func ValidateRecordNaptr(ctx context.Context, data RecordNaptrModel, resp *resource.ValidateConfigResponse) {
	if nios := flex.ExpandNestedObject[NIOSRecordNaptrModel](ctx, data.NIOS, &resp.Diagnostics); nios != nil {
		validateRecordNaptrNIOSConfig(ctx, nios, resp)
	}
	if uddi := flex.ExpandNestedObject[UDDIRecordNaptrModel](ctx, data.UDDI, &resp.Diagnostics); uddi != nil {
		validateRecordNaptrUDDIConfig(ctx, uddi, resp)
	}
}

func validateRecordNaptrNIOSConfig(ctx context.Context, m *NIOSRecordNaptrModel, resp *resource.ValidateConfigResponse) {
}

func validateRecordNaptrUDDIConfig(ctx context.Context, m *UDDIRecordNaptrModel, resp *resource.ValidateConfigResponse) {
}

type UDDIRecordNaptrRdataModel struct {
	Flags       types.String `tfsdk:"flags"`
	Order       types.Int64  `tfsdk:"order"`
	Preference  types.Int64  `tfsdk:"preference"`
	Regexp      types.String `tfsdk:"regexp"`
	Replacement types.String `tfsdk:"replacement"`
	Services    types.String `tfsdk:"services"`
}

var UDDIRecordNaptrRdataAttrTypes = map[string]attr.Type{
	"flags":       types.StringType,
	"order":       types.Int64Type,
	"preference":  types.Int64Type,
	"regexp":      types.StringType,
	"replacement": types.StringType,
	"services":    types.StringType,
}

var UDDIRecordNaptrRdataResourceSchemaAttributes = map[string]schema.Attribute{
	"flags": schema.StringAttribute{
		Optional: true,
		Validators: []validator.String{
			stringvalidator.LengthAtMost(1),
		},
		MarkdownDescription: "A character string containing flags to control aspects of the rewriting and interpretation of the fields in the DNS resource record. The flags that are currently used are:\n" +
			"  * `U`: Indicates that the output maps to a URI (Uniform Record Identifier).\n" +
			"  * `S`: Indicates that the output is a domain name that has at least one SRV record. The DNS client must then send a query for the SRV record of the resulting domain name.\n" +
			"  * `A`: Indicates that the output is a domain name that has at least one A or AAAA record. The DNS client must then send a query for the A or AAAA record of the resulting domain name.\n" +
			"  * `P`: Indicates that the protocol specified in the _services_ field defines the next step or phase.",
	},
	"order": schema.Int64Attribute{
		Required:            true,
		MarkdownDescription: "A 16-bit unsigned integer specifying the order in which the NAPTR records must be processed. Low numbers are processed before high numbers, and once a NAPTR is found whose rule \"matches\" the target, the client must not consider any NAPTRs with a higher value for order (except as noted below for the \"flags\" field. The range of the value is 0 to 65535.",
	},
	"preference": schema.Int64Attribute{
		Required:            true,
		MarkdownDescription: "A 16-bit unsigned integer that specifies the order in which NAPTR records with equal \"order\" values should be processed, low numbers being processed before high numbers. This is similar to the preference field in an MX record, and is used so domain administrators can direct clients towards more capable hosts or lighter weight protocols. A client may look at records with higher preference values if it has a good reason to do so such as not understanding the preferred protocol or service. The range of the value is 0 to 65535.",
	},
	"regexp": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "A string containing a substitution expression that is applied to the original string held by the client in order to construct the next domain name to lookup. Defaults to none.",
	},
	"replacement": schema.StringAttribute{
		Required:            true,
		MarkdownDescription: "The next name to query for NAPTR, SRV, or address records depending on the value of the _flags_ field. This can be an absolute or relative domain name. Can be empty.",
	},
	"services": schema.StringAttribute{
		Required:            true,
		MarkdownDescription: "Specifies the service(s) available down this rewrite path. It may also specify the particular protocol that is used to talk with a service. A protocol must be specified if the flags field states that the NAPTR is terminal. If a protocol is specified, but the flags field does not state that the NAPTR is terminal, the next lookup must be for a NAPTR. The client may choose not to perform the next lookup if the protocol is unknown, but that behavior must not be relied upon.",
	},
}

func ExpandUDDIRecordNaptrRdata(ctx context.Context, o types.Object, diags *diag.Diagnostics) map[string]any {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m UDDIRecordNaptrRdataModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	to := map[string]any{
		"order":       flex.ExpandInt64(m.Order),
		"preference":  flex.ExpandInt64(m.Preference),
		"replacement": flex.ExpandString(m.Replacement),
		"services":    flex.ExpandString(m.Services),
		"flags":       flex.ExpandString(m.Flags),
		"regexp":      flex.ExpandString(m.Regexp),
	}
	return to
}

func FlattenUDDIRecordNaptrRdata(ctx context.Context, from map[string]any, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(UDDIRecordNaptrRdataAttrTypes)
	}
	m := UDDIRecordNaptrRdataModel{
		Flags:       flex.FlattenStringPointer(flex.RDataStringPtr(from["flags"])),
		Order:       flex.FlattenInt64Pointer(flex.RDataInt64Ptr(from["order"])),
		Preference:  flex.FlattenInt64Pointer(flex.RDataInt64Ptr(from["preference"])),
		Regexp:      flex.FlattenStringPointer(flex.RDataStringPtr(from["regexp"])),
		Replacement: flex.FlattenString(flex.RDataString(from["replacement"])),
		Services:    flex.FlattenString(flex.RDataString(from["services"])),
	}
	obj, d := types.ObjectValueFrom(ctx, UDDIRecordNaptrRdataAttrTypes, m)
	diags.Append(d...)
	return obj
}
