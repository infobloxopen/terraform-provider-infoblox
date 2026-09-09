package dtc

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/mapvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/mapdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"

	coremodel "github.com/infobloxopen/terraform-provider-infoblox/internal/core/model/dtc"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
	importmod "github.com/infobloxopen/terraform-provider-infoblox/internal/planmodifiers/import"
	customvalidator "github.com/infobloxopen/terraform-provider-infoblox/internal/validator"
)

type DtcMonitorHttpModel struct {
	Id   types.String `tfsdk:"id"`
	NIOS types.Object `tfsdk:"nios"`
	UDDI types.Object `tfsdk:"uddi"`
}

var DtcMonitorHttpAttrTypes = map[string]attr.Type{
	"id":   types.StringType,
	"nios": types.ObjectType{AttrTypes: NIOSDtcMonitorHttpAttrTypes},
	"uddi": types.ObjectType{AttrTypes: UDDIDtcMonitorHttpAttrTypes},
}

type NIOSDtcMonitorHttpModel struct {
	Ciphers             types.String `tfsdk:"ciphers"`
	ClientCert          types.String `tfsdk:"client_cert"`
	Comment             types.String `tfsdk:"comment"`
	ContentCheck        types.String `tfsdk:"content_check"`
	ContentCheckInput   types.String `tfsdk:"content_check_input"`
	ContentCheckOp      types.String `tfsdk:"content_check_op"`
	ContentCheckRegex   types.String `tfsdk:"content_check_regex"`
	ContentExtractGroup types.Int64  `tfsdk:"content_extract_group"`
	ContentExtractType  types.String `tfsdk:"content_extract_type"`
	ContentExtractValue types.String `tfsdk:"content_extract_value"`
	EnableSni           types.Bool   `tfsdk:"enable_sni"`
	ExtAttrs            types.Map    `tfsdk:"ext_attrs"`
	ExtAttrsAll         types.Map    `tfsdk:"ext_attrs_all"`
	Interval            types.Int64  `tfsdk:"interval"`
	Name                types.String `tfsdk:"name"`
	Port                types.Int64  `tfsdk:"port"`
	Request             types.String `tfsdk:"request"`
	Result              types.String `tfsdk:"result"`
	ResultCode          types.Int64  `tfsdk:"result_code"`
	RetryDown           types.Int64  `tfsdk:"retry_down"`
	RetryUp             types.Int64  `tfsdk:"retry_up"`
	Secure              types.Bool   `tfsdk:"secure"`
	Timeout             types.Int64  `tfsdk:"timeout"`
	ValidateCert        types.Bool   `tfsdk:"validate_cert"`
}

var NIOSDtcMonitorHttpAttrTypes = map[string]attr.Type{
	"ciphers":               types.StringType,
	"client_cert":           types.StringType,
	"comment":               types.StringType,
	"content_check":         types.StringType,
	"content_check_input":   types.StringType,
	"content_check_op":      types.StringType,
	"content_check_regex":   types.StringType,
	"content_extract_group": types.Int64Type,
	"content_extract_type":  types.StringType,
	"content_extract_value": types.StringType,
	"enable_sni":            types.BoolType,
	"ext_attrs":             types.MapType{ElemType: types.StringType},
	"ext_attrs_all":         types.MapType{ElemType: types.StringType},
	"interval":              types.Int64Type,
	"name":                  types.StringType,
	"port":                  types.Int64Type,
	"request":               types.StringType,
	"result":                types.StringType,
	"result_code":           types.Int64Type,
	"retry_down":            types.Int64Type,
	"retry_up":              types.Int64Type,
	"secure":                types.BoolType,
	"timeout":               types.Int64Type,
	"validate_cert":         types.BoolType,
}

type UDDIDtcMonitorHttpModel struct {
	CheckResponseBody           types.Bool   `tfsdk:"check_response_body"`
	CheckResponseBodyNegative   types.Bool   `tfsdk:"check_response_body_negative"`
	CheckResponseBodyRegex      types.String `tfsdk:"check_response_body_regex"`
	CheckResponseHeader         types.Bool   `tfsdk:"check_response_header"`
	CheckResponseHeaderNegative types.Bool   `tfsdk:"check_response_header_negative"`
	CheckResponseHeaderRegexes  types.List   `tfsdk:"check_response_header_regexes"`
	Codes                       types.String `tfsdk:"codes"`
	Comment                     types.String `tfsdk:"comment"`
	Disabled                    types.Bool   `tfsdk:"disabled"`
	Https                       types.Bool   `tfsdk:"https"`
	Interval                    types.Int64  `tfsdk:"interval"`
	Name                        types.String `tfsdk:"name"`
	Port                        types.Int64  `tfsdk:"port"`
	Request                     types.String `tfsdk:"request"`
	RetryDown                   types.Int64  `tfsdk:"retry_down"`
	RetryUp                     types.Int64  `tfsdk:"retry_up"`
	Tags                        types.Map    `tfsdk:"tags"`
	TagsAll                     types.Map    `tfsdk:"tags_all"`
	Timeout                     types.Int64  `tfsdk:"timeout"`
}

var UDDIDtcMonitorHttpAttrTypes = map[string]attr.Type{
	"check_response_body":            types.BoolType,
	"check_response_body_negative":   types.BoolType,
	"check_response_body_regex":      types.StringType,
	"check_response_header":          types.BoolType,
	"check_response_header_negative": types.BoolType,
	"check_response_header_regexes":  types.ListType{ElemType: types.ObjectType{AttrTypes: HeaderRegexAttrTypes}},
	"codes":                          types.StringType,
	"comment":                        types.StringType,
	"disabled":                       types.BoolType,
	"https":                          types.BoolType,
	"interval":                       types.Int64Type,
	"name":                           types.StringType,
	"port":                           types.Int64Type,
	"request":                        types.StringType,
	"retry_down":                     types.Int64Type,
	"retry_up":                       types.Int64Type,
	"tags":                           types.MapType{ElemType: types.StringType},
	"tags_all":                       types.MapType{ElemType: types.StringType},
	"timeout":                        types.Int64Type,
}

const (
	DtcMonitorHttpReturnFields = "ciphers,client_cert,comment,content_check,content_check_input,content_check_op,content_check_regex,content_extract_group,content_extract_type,content_extract_value,enable_sni,extattrs,interval,name,port,request,result,result_code,retry_down,retry_up,secure,timeout,validate_cert"
)

var DtcMonitorHttpResourceSchemaAttributes = map[string]schema.Attribute{
	"id": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The reference to the object.",
	},
	"nios": schema.SingleNestedAttribute{
		Optional:            true,
		MarkdownDescription: "NIOS backend-specific fields.",
		Attributes:          DtcMonitorHttpResourceNiosSchemaAttributes,
	},
	"uddi": schema.SingleNestedAttribute{
		Optional:            true,
		MarkdownDescription: "UDDI backend-specific fields.",
		Attributes:          DtcMonitorHttpResourceUddiSchemaAttributes,
	},
}

var DtcMonitorHttpResourceNiosSchemaAttributes = map[string]schema.Attribute{
	"ciphers": schema.StringAttribute{
		Optional: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
		},
		MarkdownDescription: "An optional cipher list for a secure HTTP/S connection.",
	},
	"client_cert": schema.StringAttribute{
		Optional: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
		},
		MarkdownDescription: "An optional client certificate, supplied in a secure HTTP/S mode if present.",
	},
	"comment": schema.StringAttribute{
		Optional: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
			stringvalidator.LengthBetween(0, 256),
			customvalidator.ValidateTrimmedString(),
		},
		MarkdownDescription: "Comment for this DTC monitor; maximum 256 characters.",
	},
	"content_check": schema.StringAttribute{
		Default: stringdefault.StaticString("NONE"),
		Validators: []validator.String{
			stringvalidator.OneOf("NONE", "EXTRACT", "MATCH"),
		},
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The content check type.",
	},
	"content_check_input": schema.StringAttribute{
		Default: stringdefault.StaticString("ALL"),
		Validators: []validator.String{
			stringvalidator.OneOf("ALL", "HEADERS", "BODY"),
		},
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "A portion of response to use as input for content check.",
	},
	"content_check_op": schema.StringAttribute{
		Validators: []validator.String{
			stringvalidator.OneOf("EQ", "GEQ", "LEQ", "NEQ"),
		},
		Optional:            true,
		MarkdownDescription: "A content check success criteria operator.",
	},
	"content_check_regex": schema.StringAttribute{
		Optional: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
			customvalidator.ValidateTrimmedString(),
		},
		MarkdownDescription: "A content check regular expression.",
	},
	"content_extract_group": schema.Int64Attribute{
		Optional: true,
		Computed: true,
		Default:  int64default.StaticInt64(0),
		Validators: []validator.Int64{
			int64validator.Between(0, 8),
		},
		MarkdownDescription: "A content extraction sub-expression to extract.",
	},
	"content_extract_type": schema.StringAttribute{
		Default: stringdefault.StaticString("STRING"),
		Validators: []validator.String{
			stringvalidator.OneOf("STRING", "INTEGER"),
		},
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "A content extraction expected type for the extracted data.",
	},
	"content_extract_value": schema.StringAttribute{
		Optional: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
			customvalidator.ValidateTrimmedString(),
		},
		MarkdownDescription: "A content extraction value to compare with extracted result.",
	},
	"enable_sni": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Determines whether the Server Name Indication (SNI) for HTTPS monitor is enabled.",
	},
	"ext_attrs": schema.MapAttribute{
		Optional:    true,
		Computed:    true,
		ElementType: types.StringType,
		Default:     mapdefault.StaticValue(types.MapNull(types.StringType)),
		Validators: []validator.Map{
			mapvalidator.SizeAtLeast(1),
		},
		MarkdownDescription: "Extensible attributes associated with the object. For valid values for extensible attributes, see {extattrs:values}.",
	},
	"ext_attrs_all": schema.MapAttribute{
		Computed:            true,
		ElementType:         types.StringType,
		MarkdownDescription: "All ext_attrs including Terraform Internal ID and inherited attributes.",
		PlanModifiers: []planmodifier.Map{
			importmod.AssociateInternalId(),
		},
	},
	"interval": schema.Int64Attribute{
		Optional:            true,
		Computed:            true,
		Default:             int64default.StaticInt64(5),
		MarkdownDescription: "The interval for TCP health check.",
	},
	"name": schema.StringAttribute{
		Required: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
			customvalidator.ValidateTrimmedString(),
		},
		MarkdownDescription: "The display name for this DTC monitor.",
	},
	"port": schema.Int64Attribute{
		Optional: true,
		Computed: true,
		Default:  int64default.StaticInt64(80),
		Validators: []validator.Int64{
			int64validator.Between(1, 65535),
		},
		MarkdownDescription: "Port for TCP requests.",
	},
	"request": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
		},
		MarkdownDescription: "An HTTP request to send.",
	},
	"result": schema.StringAttribute{
		Default: stringdefault.StaticString("ANY"),
		Validators: []validator.String{
			stringvalidator.OneOf("ANY", "CODE_IS", "CODE_IS_NOT"),
		},
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The type of an expected result.",
	},
	"result_code": schema.Int64Attribute{
		Optional:            true,
		Computed:            true,
		Default:             int64default.StaticInt64(200),
		MarkdownDescription: "The expected return code.",
	},
	"retry_down": schema.Int64Attribute{
		Optional:            true,
		Computed:            true,
		Default:             int64default.StaticInt64(1),
		MarkdownDescription: "The value of how many times the server should appear as down to be treated as dead after it was alive.",
	},
	"retry_up": schema.Int64Attribute{
		Optional:            true,
		Computed:            true,
		Default:             int64default.StaticInt64(1),
		MarkdownDescription: "The value of how many times the server should appear as up to be treated as alive after it was dead.",
	},
	"secure": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "The connection security status.",
	},
	"timeout": schema.Int64Attribute{
		Optional:            true,
		Computed:            true,
		Default:             int64default.StaticInt64(15),
		MarkdownDescription: "The timeout for TCP health check in seconds.",
	},
	"validate_cert": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(true),
		MarkdownDescription: "Determines whether the validation of the remote server's certificate is enabled.",
	},
}

var DtcMonitorHttpResourceUddiSchemaAttributes = map[string]schema.Attribute{
	"check_response_body": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Optional. Flag which enables checking of the HTTP response body content. Defaults to _false_.",
	},
	"check_response_body_negative": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Optional. Flag which changes the meaning of the regex match result. If set to _true_, the response is valid if regular expression matches not found. Defaults to _false_.  The flag is currently not supported.",
	},
	"check_response_body_regex": schema.StringAttribute{
		Optional:            true,
		Computed:            true,
		Default:             stringdefault.StaticString(""),
		MarkdownDescription: "Optional. Regular expression to search for a string in the HTTP response body. Error if empty while _check_response_body_ is _true_. Defaults to empty.",
	},
	"check_response_header": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Optional. Flag which enables checking of the HTTP response header(s) content. Defaults to _false_.",
	},
	"check_response_header_negative": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Optional. Flag which changes the meaning of the header regexes match result. If set to _true_, neither expression matches must be found in their respective headers for the headers to be considered valid. Defaults to _false_.",
	},
	"check_response_header_regexes": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: HeaderRegexResourceSchemaAttributes,
		},
		Optional: true,
		Validators: []validator.List{
			customvalidator.ListNotEmpty(),
		},
		MarkdownDescription: "Optional. List of (header, regular expression) pairs. All expression matches must be found in their respective headers for the headers to be considered valid. Error if empty while _check_response_header_ is _true_. Defaults to empty.",
	},
	"codes": schema.StringAttribute{
		Optional:            true,
		Computed:            true,
		Default:             stringdefault.StaticString(""),
		MarkdownDescription: "Optional. Response Status Codes meaning the health check is successful. If empty, any code means success. Individual codes and code ranges are supported, ex. \"102,105-107,109-110,120\".",
	},
	"comment": schema.StringAttribute{
		Optional:            true,
		Computed:            true,
		Default:             stringdefault.StaticString(""),
		MarkdownDescription: "Optional. Comment for __HTTPHealthCheck__.",
	},
	"disabled": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Optional. Flag which enables/disables __HTTPHealthCheck__. Defaults to _false_.",
	},
	"https": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Optional. Flag which enables Hypertext Transfer Protocol Secure (HTTPS) in a health check. Defaults to _false_.",
	},
	"interval": schema.Int64Attribute{
		Optional:            true,
		Computed:            true,
		Default:             int64default.StaticInt64(15),
		MarkdownDescription: "Optional. Interval value in seconds. The health check runs only for the specified interval and it is measured from the beginning of the previous check cycle. Defaults to _15_.",
	},
	"name": schema.StringAttribute{
		Required:            true,
		MarkdownDescription: "Display name of __HTTPHealthCheck__.",
	},
	"port": schema.Int64Attribute{
		Required:            true,
		MarkdownDescription: "Destination TCP port of __HTTPHealthCheck__.",
	},
	"request": schema.StringAttribute{
		Required:            true,
		MarkdownDescription: "HTTP request in a text format, it consists of HTTP method, request target, HTTP headers, request body.",
	},
	"retry_down": schema.Int64Attribute{
		Optional:            true,
		Computed:            true,
		Default:             int64default.StaticInt64(1),
		MarkdownDescription: "Optional. Retry down count. The value determines how many bad health checks in a row must be received by the onprem host from the DTC Server for treating the health check as failed. Defaults to _1_.",
	},
	"retry_up": schema.Int64Attribute{
		Optional:            true,
		Computed:            true,
		Default:             int64default.StaticInt64(1),
		MarkdownDescription: "Optional. Retry up count. The value determines how many good health checks in a row must be received by the onprem host from the DTC Server for treating the health check as successful. Defaults to _1_.",
	},
	"tags": schema.MapAttribute{
		Optional:    true,
		Computed:    true,
		ElementType: types.StringType,
		Default:     mapdefault.StaticValue(types.MapNull(types.StringType)),
		Validators: []validator.Map{
			mapvalidator.SizeAtLeast(1),
		},
		MarkdownDescription: "Optional. The tags for __HTTPHealthCheck__ in JSON format.",
	},
	"tags_all": schema.MapAttribute{
		Computed:            true,
		ElementType:         types.StringType,
		MarkdownDescription: "All tags including inherited values.",
	},
	"timeout": schema.Int64Attribute{
		Optional:            true,
		Computed:            true,
		Default:             int64default.StaticInt64(10),
		MarkdownDescription: "Optional. Timeout value in seconds. The health check waits for the specified number of seconds after sending a request. If it does not receive a response within the number of seconds, then the health check is considered as failed. Defaults to _10_.",
	},
}

// Expand converts the TF model to the infoblox core model
func (m *DtcMonitorHttpModel) Expand(ctx context.Context, diags *diag.Diagnostics, isCreate bool) *coremodel.DtcMonitorHttp {
	if m == nil {
		return nil
	}

	obj := &coremodel.DtcMonitorHttp{}

	// Expand NIOS nested attribute (returns nil if not present)
	niosModel := flex.ExpandNestedObject[NIOSDtcMonitorHttpModel](ctx, m.NIOS, diags)
	if niosModel != nil {
		obj.NIOS = niosModel.Expand(ctx, diags)
	}

	// Expand UDDI nested attribute (returns nil if not present)
	uddiModel := flex.ExpandNestedObject[UDDIDtcMonitorHttpModel](ctx, m.UDDI, diags)
	if uddiModel != nil {
		obj.UDDI = uddiModel.Expand(ctx, diags)
	}

	return obj
}

// Expand converts the NIOS TF model to the core model.
func (m *NIOSDtcMonitorHttpModel) Expand(ctx context.Context, diags *diag.Diagnostics) *coremodel.NIOSDtcMonitorHttpExt {
	return &coremodel.NIOSDtcMonitorHttpExt{
		Ciphers:             flex.ExpandStringPointerNullAsEmpty(m.Ciphers),
		ClientCert:          flex.ExpandStringPointer(m.ClientCert),
		Comment:             flex.ExpandStringPointerNullAsEmpty(m.Comment),
		ContentCheck:        flex.ExpandStringPointerNullAsEmpty(m.ContentCheck),
		ContentCheckInput:   flex.ExpandStringPointerNullAsEmpty(m.ContentCheckInput),
		ContentCheckOp:      flex.ExpandStringPointer(m.ContentCheckOp),
		ContentCheckRegex:   flex.ExpandStringPointerNullAsEmpty(m.ContentCheckRegex),
		ContentExtractGroup: flex.ExpandInt64Pointer(m.ContentExtractGroup),
		ContentExtractType:  flex.ExpandStringPointerNullAsEmpty(m.ContentExtractType),
		ContentExtractValue: flex.ExpandStringPointerNullAsEmpty(m.ContentExtractValue),
		EnableSni:           flex.ExpandBoolPointer(m.EnableSni),
		ExtAttrs:            flex.ExpandMapStringAny(ctx, m.ExtAttrs, diags),
		Interval:            flex.ExpandInt64Pointer(m.Interval),
		Name:                flex.ExpandStringPointerNullAsEmpty(m.Name),
		Port:                flex.ExpandInt64Pointer(m.Port),
		Request:             flex.ExpandStringPointerNullAsEmpty(m.Request),
		Result:              flex.ExpandStringPointerNullAsEmpty(m.Result),
		ResultCode:          flex.ExpandInt64Pointer(m.ResultCode),
		RetryDown:           flex.ExpandInt64Pointer(m.RetryDown),
		RetryUp:             flex.ExpandInt64Pointer(m.RetryUp),
		Secure:              flex.ExpandBoolPointer(m.Secure),
		Timeout:             flex.ExpandInt64Pointer(m.Timeout),
		ValidateCert:        flex.ExpandBoolPointer(m.ValidateCert),
	}
}

// Expand converts the UDDI TF model to the core model.
func (m *UDDIDtcMonitorHttpModel) Expand(ctx context.Context, diags *diag.Diagnostics) *coremodel.UDDIDtcMonitorHttpExt {
	return &coremodel.UDDIDtcMonitorHttpExt{
		CheckResponseBody:           flex.ExpandBoolPointer(m.CheckResponseBody),
		CheckResponseBodyNegative:   flex.ExpandBoolPointer(m.CheckResponseBodyNegative),
		CheckResponseBodyRegex:      flex.ExpandStringPointer(m.CheckResponseBodyRegex),
		CheckResponseHeader:         flex.ExpandBoolPointer(m.CheckResponseHeader),
		CheckResponseHeaderNegative: flex.ExpandBoolPointer(m.CheckResponseHeaderNegative),
		CheckResponseHeaderRegexes:  flex.ExpandFrameworkListNestedBlock(ctx, m.CheckResponseHeaderRegexes, diags, ExpandHeaderRegex),
		Codes:                       flex.ExpandStringPointer(m.Codes),
		Comment:                     flex.ExpandStringPointer(m.Comment),
		Disabled:                    flex.ExpandBoolPointer(m.Disabled),
		Https:                       flex.ExpandBoolPointer(m.Https),
		Interval:                    flex.ExpandInt64Pointer(m.Interval),
		Name:                        flex.ExpandString(m.Name),
		Port:                        flex.ExpandInt64(m.Port),
		Request:                     flex.ExpandStringPointer(m.Request),
		RetryDown:                   flex.ExpandInt64Pointer(m.RetryDown),
		RetryUp:                     flex.ExpandInt64Pointer(m.RetryUp),
		Tags:                        flex.ExpandMapStringAny(ctx, m.Tags, diags),
		Timeout:                     flex.ExpandInt64Pointer(m.Timeout),
	}
}

// Flatten populates the TF model from a core response.
func (m *DtcMonitorHttpModel) Flatten(ctx context.Context, resp *coremodel.DtcMonitorHttp, diags *diag.Diagnostics) {
	if resp == nil {
		return
	}

	m.Id = flex.FlattenStringPointer(resp.Id)

	// Extract existing NIOS model, flatten API response onto it, convert back
	niosModel := flex.ExpandNestedObject[NIOSDtcMonitorHttpModel](ctx, m.NIOS, diags)
	if niosModel == nil {
		niosModel = &NIOSDtcMonitorHttpModel{}
	}
	plannedNIOS := flex.ExpandNestedObject[NIOSDtcMonitorHttpModel](ctx, m.NIOS, diags)
	niosModel.Flatten(ctx, resp.NIOS, diags)
	if resp.NIOS != nil {
		PostFlattenDtcMonitorHttpNIOS(ctx, plannedNIOS, niosModel, diags)
		m.NIOS = flex.FlattenNestedObject(ctx, niosModel, NIOSDtcMonitorHttpAttrTypes, diags)
	} else {
		m.NIOS = types.ObjectNull(NIOSDtcMonitorHttpAttrTypes)
	}

	// Extract existing UDDI model, flatten API response onto it, convert back
	uddiModel := flex.ExpandNestedObject[UDDIDtcMonitorHttpModel](ctx, m.UDDI, diags)
	if uddiModel == nil {
		uddiModel = &UDDIDtcMonitorHttpModel{}
	}
	plannedUDDI := flex.ExpandNestedObject[UDDIDtcMonitorHttpModel](ctx, m.UDDI, diags)
	uddiModel.Flatten(ctx, resp.UDDI, diags)
	if resp.UDDI != nil {
		PostFlattenDtcMonitorHttpUDDI(ctx, plannedUDDI, uddiModel, diags)
		m.UDDI = flex.FlattenNestedObject(ctx, uddiModel, UDDIDtcMonitorHttpAttrTypes, diags)
	} else {
		m.UDDI = types.ObjectNull(UDDIDtcMonitorHttpAttrTypes)
	}
}

// Flatten merges API response onto existing NIOS model.
func (m *NIOSDtcMonitorHttpModel) Flatten(ctx context.Context, from *coremodel.NIOSDtcMonitorHttpExt, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	planExtAttrs := m.ExtAttrs
	if planExtAttrs.IsUnknown() {
		planExtAttrs = types.MapNull(types.StringType)
	}
	m.Ciphers = flex.FlattenStringPointerEmptyAsNull(from.Ciphers)
	m.ClientCert = flex.FlattenStringPointerEmptyAsNull(from.ClientCert)
	m.Comment = flex.FlattenStringPointerEmptyAsNull(from.Comment)
	m.ContentCheck = flex.FlattenStringPointerEmptyAsNull(from.ContentCheck)
	m.ContentCheckInput = flex.FlattenStringPointerEmptyAsNull(from.ContentCheckInput)
	m.ContentCheckOp = flex.FlattenStringPointerEmptyAsNull(from.ContentCheckOp)
	m.ContentCheckRegex = flex.FlattenStringPointerEmptyAsNull(from.ContentCheckRegex)
	m.ContentExtractGroup = flex.FlattenInt64Pointer(from.ContentExtractGroup)
	m.ContentExtractType = flex.FlattenStringPointerEmptyAsNull(from.ContentExtractType)
	m.ContentExtractValue = flex.FlattenStringPointerEmptyAsNull(from.ContentExtractValue)
	m.EnableSni = flex.FlattenBoolPointer(from.EnableSni)
	m.ExtAttrs, m.ExtAttrsAll = flex.FlattenEAs(planExtAttrs, from.ExtAttrs)
	m.Interval = flex.FlattenInt64Pointer(from.Interval)
	m.Name = flex.FlattenStringPointerEmptyAsNull(from.Name)
	m.Port = flex.FlattenInt64Pointer(from.Port)
	m.Request = flex.FlattenStringPointerEmptyAsNull(from.Request)
	m.Result = flex.FlattenStringPointerEmptyAsNull(from.Result)
	m.ResultCode = flex.FlattenInt64Pointer(from.ResultCode)
	m.RetryDown = flex.FlattenInt64Pointer(from.RetryDown)
	m.RetryUp = flex.FlattenInt64Pointer(from.RetryUp)
	m.Secure = flex.FlattenBoolPointer(from.Secure)
	m.Timeout = flex.FlattenInt64Pointer(from.Timeout)
	m.ValidateCert = flex.FlattenBoolPointer(from.ValidateCert)
}

// Flatten merges API response onto existing UDDI model.
func (m *UDDIDtcMonitorHttpModel) Flatten(ctx context.Context, from *coremodel.UDDIDtcMonitorHttpExt, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	m.CheckResponseBody = flex.FlattenBoolPointer(from.CheckResponseBody)
	m.CheckResponseBodyNegative = flex.FlattenBoolPointer(from.CheckResponseBodyNegative)
	m.CheckResponseBodyRegex = flex.FlattenStringPointer(from.CheckResponseBodyRegex)
	m.CheckResponseHeader = flex.FlattenBoolPointer(from.CheckResponseHeader)
	m.CheckResponseHeaderNegative = flex.FlattenBoolPointer(from.CheckResponseHeaderNegative)
	m.CheckResponseHeaderRegexes = flex.FlattenFrameworkListNestedBlock(ctx, from.CheckResponseHeaderRegexes, HeaderRegexAttrTypes, diags, FlattenHeaderRegex)
	m.Codes = flex.FlattenStringPointer(from.Codes)
	m.Comment = flex.FlattenStringPointer(from.Comment)
	m.Disabled = flex.FlattenBoolPointer(from.Disabled)
	m.Https = flex.FlattenBoolPointer(from.Https)
	m.Interval = flex.FlattenInt64Pointer(from.Interval)
	m.Name = flex.FlattenString(from.Name)
	m.Port = flex.FlattenInt64(from.Port)
	m.Request = flex.FlattenStringPointer(from.Request)
	m.RetryDown = flex.FlattenInt64Pointer(from.RetryDown)
	m.RetryUp = flex.FlattenInt64Pointer(from.RetryUp)
	tagsAll := flex.FlattenMapStringAny(ctx, from.Tags, diags)
	if m.Tags.IsNull() || m.Tags.IsUnknown() {
		m.Tags = tagsAll
	}
	m.TagsAll = tagsAll
	m.Timeout = flex.FlattenInt64Pointer(from.Timeout)
}
