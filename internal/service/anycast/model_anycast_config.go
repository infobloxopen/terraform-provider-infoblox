package anycast

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-nettypes/iptypes"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/mapvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/mapdefault"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"

	coremodel "github.com/infobloxopen/terraform-provider-infoblox/internal/core/model/anycast"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
	customvalidator "github.com/infobloxopen/terraform-provider-infoblox/internal/validator"
)

type AnycastConfigModel struct {
	Id   types.Int64  `tfsdk:"id"`
	UDDI types.Object `tfsdk:"uddi"`
}

var AnycastConfigAttrTypes = map[string]attr.Type{
	"id":   types.Int64Type,
	"uddi": types.ObjectType{AttrTypes: UDDIAnycastConfigAttrTypes},
}

type UDDIAnycastConfigModel struct {
	AnycastIpAddress   iptypes.IPv4Address `tfsdk:"anycast_ip_address"`
	AnycastIpv6Address iptypes.IPv6Address `tfsdk:"anycast_ipv6_address"`
	CreatedAt          timetypes.RFC3339   `tfsdk:"created_at"`
	Description        types.String        `tfsdk:"description"`
	Fields             types.Object        `tfsdk:"fields"`
	IsConfigured       types.Bool          `tfsdk:"is_configured"`
	Name               types.String        `tfsdk:"name"`
	OnpremHosts        types.List          `tfsdk:"onprem_hosts"`
	RuntimeStatus      types.String        `tfsdk:"runtime_status"`
	Service            types.String        `tfsdk:"service"`
	Tags               types.Map           `tfsdk:"tags"`
	TagsAll            types.Map           `tfsdk:"tags_all"`
	UpdatedAt          timetypes.RFC3339   `tfsdk:"updated_at"`
}

var UDDIAnycastConfigAttrTypes = map[string]attr.Type{
	"anycast_ip_address":   iptypes.IPv4AddressType{},
	"anycast_ipv6_address": iptypes.IPv6AddressType{},
	"created_at":           timetypes.RFC3339Type{},
	"description":          types.StringType,
	"fields":               types.ObjectType{AttrTypes: ProtobufFieldMaskAttrTypes},
	"is_configured":        types.BoolType,
	"name":                 types.StringType,
	"onprem_hosts":         types.ListType{ElemType: types.ObjectType{AttrTypes: OnpremHostRefAttrTypes}},
	"runtime_status":       types.StringType,
	"service":              types.StringType,
	"tags":                 types.MapType{ElemType: types.StringType},
	"tags_all":             types.MapType{ElemType: types.StringType},
	"updated_at":           timetypes.RFC3339Type{},
}

const (
	AnycastConfigReturnFields = ""
)

var AnycastConfigResourceSchemaAttributes = map[string]schema.Attribute{
	"id": schema.Int64Attribute{
		Computed:            true,
		MarkdownDescription: "",
	},
	"uddi": schema.SingleNestedAttribute{
		Optional:            true,
		MarkdownDescription: "UDDI backend-specific fields.",
		Attributes:          AnycastConfigResourceUddiSchemaAttributes,
	},
}

var AnycastConfigResourceUddiSchemaAttributes = map[string]schema.Attribute{
	"anycast_ip_address": schema.StringAttribute{
		Required:            true,
		CustomType:          iptypes.IPv4AddressType{},
		MarkdownDescription: "IPv4 address of the host in string format.",
	},
	"anycast_ipv6_address": schema.StringAttribute{
		Optional:            true,
		CustomType:          iptypes.IPv6AddressType{},
		MarkdownDescription: "IPv6 address of the host in string format",
	},
	"created_at": schema.StringAttribute{
		Computed:            true,
		CustomType:          timetypes.RFC3339Type{},
		MarkdownDescription: "Time when the object has been created.",
	},
	"description": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The description for the address object. May contain 0 to 1024 characters. Can include UTF-8.",
	},
	"fields": schema.SingleNestedAttribute{
		Attributes:          ProtobufFieldMaskResourceSchemaAttributes,
		Optional:            true,
		MarkdownDescription: "paths: \"f.a\"     paths: \"f.b.d\"  Here `f` represents a field in some root message, `a` and `b` fields in the message found in `f`, and `d` a field found in the message in `f.b`.  Field masks are used to specify a subset of fields that should be returned by a get operation or modified by an update operation. Field masks also have a custom JSON encoding (see below).  # Field Masks in Projections  When used in the context of a projection, a response message or sub-message is filtered by the API to only contain those fields as specified in the mask. For example, if the mask in the previous example is applied to a response message as follows:      f {       a : 22       b {         d : 1         x : 2       }       y : 13     }     z: 8  The result will not contain specific values for fields x,y and z (their value will be set to the default, and omitted in proto text output):       f {       a : 22       b {         d : 1       }     }  A repeated field is not allowed except at the last position of a paths string.  If a FieldMask object is not present in a get operation, the operation applies to all fields (as if a FieldMask of all fields had been specified).  Note that a field mask does not necessarily apply to the top-level response message. In case of a REST get operation, the field mask applies directly to the response, but in case of a REST list operation, the mask instead applies to each individual message in the returned resource list. In case of a REST custom method, other definitions may be used. Where the mask applies will be clearly documented together with its declaration in the API.  In any case, the effect on the returned resource/resources is required behavior for APIs.  # Field Masks in Update Operations  A field mask in update operations specifies which fields of the targeted resource are going to be updated. The API is required to only change the values of the fields as specified in the mask and leave the others untouched. If a resource is passed in to describe the updated values, the API ignores the values of all fields not covered by the mask.  If a repeated field is specified for an update operation, the existing repeated values in the target resource will be overwritten by the new values. Note that a repeated field is only allowed in the last position of a `paths` string.  If a sub-message is specified in the last position of the field mask for an update operation, then the existing sub-message in the target resource is overwritten. Given the target message:      f {       b {         d : 1         x : 2       }       c : 1     }  And an update message:      f {       b {         d : 10       }     }  then if the field mask is:   paths: \"f.b\"  then the result will be:      f {       b {         d : 10       }       c : 1     }  However, if the update mask was:   paths: \"f.b.d\"  then the result would be:      f {       b {         d : 10         x : 2       }       c : 1     }  In order to reset a field's value to the default, the field must be in the mask and set to the default value in the provided resource. Hence, in order to reset all fields of a resource, provide a default instance of the resource and set all fields in the mask, or do not provide a mask as described below.  If a field mask is not present on update, the operation applies to all fields (as if a field mask of all fields has been specified). Note that in the presence of schema evolution, this may mean that fields the client does not know and has therefore not filled into the request will be reset to their default. If this is unwanted behavior, a specific service may require a client to always specify a field mask, producing an error if not.  As with get operations, the location of the resource which describes the updated values in the request message depends on the operation kind. In any case, the effect of the field mask is required to be honored by the API.  ## Considerations for HTTP REST  The HTTP kind of an update operation which uses a field mask must be set to PATCH instead of PUT in order to satisfy HTTP semantics (PUT must only be used for full updates).  # JSON Encoding of Field Masks  In JSON, a field mask is encoded as a single string where paths are separated by a comma. Fields name in each path are converted to/from lower-camel naming conventions.  As an example, consider the following message declarations:      message Profile {       User user = 1;       Photo photo = 2;     }     message User {       string display_name = 1;       string address = 2;     }  In proto a field mask for `Profile` may look as such:      mask {       paths: \"user.display_name\"       paths: \"photo\"     }  In JSON, the same mask is represented as below:      {       mask: \"user.displayName,photo\"     }  # Field Masks and Oneof Fields  Field masks treat fields in oneofs just as regular fields. Consider the following message:      message SampleMessage {       oneof test_oneof {         string name = 4;         SubMessage sub_message = 9;       }     }  The field mask can be:      mask {       paths: \"name\"     }  Or:      mask {       paths: \"sub_message\"     }  Note that oneof type names (\"test_oneof\" in this case) cannot be used in paths.  ## Field Mask Verification  The implementation of the all the API methods, which have any FieldMask type field in the request, should verify the included field paths, and return `INVALID_ARGUMENT` error if any path is duplicated or unmappable.",
	},
	"is_configured": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "",
	},
	"name": schema.StringAttribute{
		Required:            true,
		MarkdownDescription: "The name of the anycast configuration.",
	},
	"onprem_hosts": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: OnpremHostRefResourceSchemaAttributes,
		},
		Optional: true,
		Validators: []validator.List{
			customvalidator.ListNotEmpty(),
		},
		MarkdownDescription: "Struct on-prem host reference.",
	},
	"runtime_status": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "",
	},
	"service": schema.StringAttribute{
		Validators: []validator.String{
			stringvalidator.OneOf("DNS", "DFP", "NTP"),
		},
		Required:            true,
		MarkdownDescription: "The type of the Service used in anycast configuration, supports (`dns`, `ntp`, `dfp`).",
	},
	"tags": schema.MapAttribute{
		Optional:    true,
		Computed:    true,
		ElementType: types.StringType,
		Default:     mapdefault.StaticValue(types.MapNull(types.StringType)),
		Validators: []validator.Map{
			mapvalidator.SizeAtLeast(1),
		},
		MarkdownDescription: "The tags for the anycast configuration object.",
	},
	"tags_all": schema.MapAttribute{
		Computed:            true,
		ElementType:         types.StringType,
		MarkdownDescription: "All tags including inherited values.",
	},
	"updated_at": schema.StringAttribute{
		Computed:            true,
		CustomType:          timetypes.RFC3339Type{},
		MarkdownDescription: "Time when the object has been updated. Equals to _created_at_ if not updated after creation.",
	},
}

// Expand converts the TF model to the infoblox core model
func (m *AnycastConfigModel) Expand(ctx context.Context, diags *diag.Diagnostics, isCreate bool) *coremodel.AnycastConfig {
	if m == nil {
		return nil
	}

	obj := &coremodel.AnycastConfig{}

	// Expand UDDI nested attribute (returns nil if not present)
	uddiModel := flex.ExpandNestedObject[UDDIAnycastConfigModel](ctx, m.UDDI, diags)
	if uddiModel != nil {
		obj.UDDI = uddiModel.Expand(ctx, diags)
	}

	return obj
}

// Expand converts the UDDI TF model to the core model.
func (m *UDDIAnycastConfigModel) Expand(ctx context.Context, diags *diag.Diagnostics) *coremodel.UDDIAnycastConfigExt {
	return &coremodel.UDDIAnycastConfigExt{
		AnycastIpAddress:   flex.ExpandIPv4Address(m.AnycastIpAddress),
		AnycastIpv6Address: flex.ExpandIPv6Address(m.AnycastIpv6Address),
		CreatedAt:          flex.ExpandRFC3339(m.CreatedAt, diags),
		Description:        flex.ExpandStringPointer(m.Description),
		Fields:             ExpandProtobufFieldMask(ctx, m.Fields, diags),
		IsConfigured:       flex.ExpandBoolPointer(m.IsConfigured),
		Name:               flex.ExpandStringPointer(m.Name),
		OnpremHosts:        flex.ExpandFrameworkListNestedBlock(ctx, m.OnpremHosts, diags, ExpandOnpremHostRef),
		RuntimeStatus:      flex.ExpandStringPointer(m.RuntimeStatus),
		Service:            flex.ExpandStringPointer(m.Service),
		Tags:               flex.ExpandMapStringAny(ctx, m.Tags, diags),
		UpdatedAt:          flex.ExpandRFC3339(m.UpdatedAt, diags),
	}
}

// Flatten populates the TF model from a core response.
func (m *AnycastConfigModel) Flatten(ctx context.Context, resp *coremodel.AnycastConfig, diags *diag.Diagnostics) {
	if resp == nil {
		return
	}

	m.Id = flex.FlattenInt64Pointer(resp.Id)

	// Extract existing UDDI model, flatten API response onto it, convert back
	uddiModel := flex.ExpandNestedObject[UDDIAnycastConfigModel](ctx, m.UDDI, diags)
	if uddiModel == nil {
		uddiModel = &UDDIAnycastConfigModel{}
	}
	uddiModel.Flatten(ctx, resp.UDDI, diags)
	if resp.UDDI != nil {
		m.UDDI = flex.FlattenNestedObject(ctx, uddiModel, UDDIAnycastConfigAttrTypes, diags)
	} else {
		m.UDDI = types.ObjectNull(UDDIAnycastConfigAttrTypes)
	}
}

// Flatten merges API response onto existing UDDI model.
func (m *UDDIAnycastConfigModel) Flatten(ctx context.Context, from *coremodel.UDDIAnycastConfigExt, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	m.AnycastIpAddress = flex.FlattenIPv4Address(from.AnycastIpAddress)
	m.AnycastIpv6Address = flex.FlattenIPv6Address(from.AnycastIpv6Address)
	m.CreatedAt = flex.FlattenRFC3339(from.CreatedAt)
	m.Description = flex.FlattenStringPointer(from.Description)
	m.Fields = FlattenProtobufFieldMask(ctx, from.Fields, diags)
	m.IsConfigured = flex.FlattenBoolPointer(from.IsConfigured)
	m.Name = flex.FlattenStringPointer(from.Name)
	m.OnpremHosts = flex.FlattenFrameworkListNestedBlock(ctx, from.OnpremHosts, OnpremHostRefAttrTypes, diags, FlattenOnpremHostRef)
	m.RuntimeStatus = flex.FlattenStringPointer(from.RuntimeStatus)
	m.Service = flex.FlattenStringPointer(from.Service)
	tagsAll := flex.FlattenMapStringAny(ctx, from.Tags, diags)
	if m.Tags.IsNull() || m.Tags.IsUnknown() {
		m.Tags = tagsAll
	}
	m.TagsAll = tagsAll
	m.UpdatedAt = flex.FlattenRFC3339(from.UpdatedAt)
}
