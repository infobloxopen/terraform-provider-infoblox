package dtc

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/mapvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/mapdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"

	coremodel "github.com/infobloxopen/terraform-provider-infoblox/internal/core/model/dtc"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
	importmod "github.com/infobloxopen/terraform-provider-infoblox/internal/planmodifiers/import"
	customvalidator "github.com/infobloxopen/terraform-provider-infoblox/internal/validator"
)

type DtcServerModel struct {
	Id   types.String `tfsdk:"id"`
	NIOS types.Object `tfsdk:"nios"`
	UDDI types.Object `tfsdk:"uddi"`
}

var DtcServerAttrTypes = map[string]attr.Type{
	"id":   types.StringType,
	"nios": types.ObjectType{AttrTypes: NIOSDtcServerAttrTypes},
	"uddi": types.ObjectType{AttrTypes: UDDIDtcServerAttrTypes},
}

type NIOSDtcServerModel struct {
	AutoCreateHostRecord types.Bool   `tfsdk:"auto_create_host_record"`
	Comment              types.String `tfsdk:"comment"`
	Disable              types.Bool   `tfsdk:"disable"`
	ExtAttrs             types.Map    `tfsdk:"ext_attrs"`
	ExtAttrsAll          types.Map    `tfsdk:"ext_attrs_all"`
	Host                 types.String `tfsdk:"host"`
	Monitors             types.List   `tfsdk:"monitors"`
	Name                 types.String `tfsdk:"name"`
	SniHostname          types.String `tfsdk:"sni_hostname"`
}

var NIOSDtcServerAttrTypes = map[string]attr.Type{
	"auto_create_host_record": types.BoolType,
	"comment":                 types.StringType,
	"disable":                 types.BoolType,
	"ext_attrs":               types.MapType{ElemType: types.StringType},
	"ext_attrs_all":           types.MapType{ElemType: types.StringType},
	"host":                    types.StringType,
	"monitors":                types.ListType{ElemType: types.ObjectType{AttrTypes: ServerMonitorsAttrTypes}},
	"name":                    types.StringType,
	"sni_hostname":            types.StringType,
}

type UDDIDtcServerModel struct {
	Address                   types.String `tfsdk:"address"`
	AutoCreateResponseRecords types.Bool   `tfsdk:"auto_create_response_records"`
	Comment                   types.String `tfsdk:"comment"`
	Disabled                  types.Bool   `tfsdk:"disabled"`
	EndpointType              types.String `tfsdk:"endpoint_type"`
	Fqdn                      types.String `tfsdk:"fqdn"`
	Name                      types.String `tfsdk:"name"`
	Records                   types.List   `tfsdk:"records"`
	Tags                      types.Map    `tfsdk:"tags"`
	TagsAll                   types.Map    `tfsdk:"tags_all"`
}

var UDDIDtcServerAttrTypes = map[string]attr.Type{
	"address":                      types.StringType,
	"auto_create_response_records": types.BoolType,
	"comment":                      types.StringType,
	"disabled":                     types.BoolType,
	"endpoint_type":                types.StringType,
	"fqdn":                         types.StringType,
	"name":                         types.StringType,
	"records":                      types.ListType{ElemType: types.ObjectType{AttrTypes: RecordAttrTypes}},
	"tags":                         types.MapType{ElemType: types.StringType},
	"tags_all":                     types.MapType{ElemType: types.StringType},
}

const (
	DtcServerReturnFields = "auto_create_host_record,comment,disable,extattrs,health,host,monitors,name,sni_hostname,use_sni_hostname"
)

var DtcServerResourceSchemaAttributes = map[string]schema.Attribute{
	"id": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The reference to the object.",
	},
	"nios": schema.SingleNestedAttribute{
		Optional:            true,
		MarkdownDescription: "NIOS backend-specific fields.",
		Attributes:          DtcServerResourceNiosSchemaAttributes,
	},
	"uddi": schema.SingleNestedAttribute{
		Optional:            true,
		MarkdownDescription: "UDDI backend-specific fields.",
		Attributes:          DtcServerResourceUddiSchemaAttributes,
	},
}

var DtcServerResourceNiosSchemaAttributes = map[string]schema.Attribute{
	"auto_create_host_record": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(true),
		MarkdownDescription: "Enabling this option will auto-create a single read-only A/AAAA/CNAME record corresponding to the configured hostname and update it if the hostname changes.",
	},
	"comment": schema.StringAttribute{
		Optional: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
			customvalidator.ValidateTrimmedString(),
			stringvalidator.LengthBetween(0, 256),
		},
		MarkdownDescription: "Comment for the DTC Server; maximum 256 characters.",
	},
	"disable": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Determines whether the DTC Server is disabled or not. When this is set to False, the fixed address is enabled.",
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
	"host": schema.StringAttribute{
		Required: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
			customvalidator.IsValidIPv4OrFQDN(),
		},
		MarkdownDescription: "The address or FQDN of the server.",
	},
	"monitors": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: ServerMonitorsResourceSchemaAttributes,
		},
		Optional: true,
		Validators: []validator.List{
			customvalidator.ListNotEmpty(),
		},
		MarkdownDescription: "List of IP/FQDN and monitor pairs to be used for additional monitoring.",
	},
	"name": schema.StringAttribute{
		Required: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
			customvalidator.ValidateTrimmedString(),
		},
		MarkdownDescription: "The DTC Server display name.",
	},
	"sni_hostname": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
			customvalidator.IsValidNIOSDomainName(),
		},
		MarkdownDescription: "The hostname for Server Name Indication (SNI) in FQDN format.",
	},
}

var DtcServerResourceUddiSchemaAttributes = map[string]schema.Attribute{
	"address": schema.StringAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "IP Address of the __Server__. Must be set to a valid IP address if __endpoint_type__ is set to __address__. Alternatively, it can be left blank.",
	},
	"auto_create_response_records": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Optional. If the flag is enabled, A, AAAA or CNAME __Record__ is automatically generated.  Defaults to _false_.",
	},
	"comment": schema.StringAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "Optional. Comment for __Server__.",
	},
	"disabled": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Optional. Flag which enables/disables __Server__.  Defaults to _false_.",
	},
	"endpoint_type": schema.StringAttribute{
		Validators: []validator.String{
			stringvalidator.OneOf("address", "fqdn"),
		},
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The endpoint type configured for the __Server__. Can be IP Address or FQDN. The values of both fields __address__ and __fqdn__ are preserved and are not mutually exclusive, and the __endpoint_type__ defines which one to use.  Allowed values: * address * fqdn  Defaults to __address__.",
	},
	"fqdn": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Validators: []validator.String{
			customvalidator.IsValidUDDIDomainName(),
		},
		MarkdownDescription: "Fully Qualified Domain name of the __Server__. Must be set to a valid FQDN if __endpoint_type__ is set to __fqdn__. Alternatively, it can be left blank.",
	},
	"name": schema.StringAttribute{
		Required:            true,
		MarkdownDescription: "Display name of __Server__.",
	},
	"records": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: RecordResourceSchemaAttributes,
		},
		Optional: true,
		Computed: true,
		Validators: []validator.List{
			customvalidator.ListNotEmpty(),
		},
		MarkdownDescription: "Optional. List of __Records__ of the __Server__.",
	},
	"tags": schema.MapAttribute{
		Optional:    true,
		Computed:    true,
		ElementType: types.StringType,
		Default:     mapdefault.StaticValue(types.MapNull(types.StringType)),
		Validators: []validator.Map{
			mapvalidator.SizeAtLeast(1),
		},
		MarkdownDescription: "Optional. The tags for __Server__ in JSON format.",
	},
	"tags_all": schema.MapAttribute{
		Computed:            true,
		ElementType:         types.StringType,
		MarkdownDescription: "All tags including inherited values.",
	},
}

// Expand converts the TF model to the infoblox core model
func (m *DtcServerModel) Expand(ctx context.Context, diags *diag.Diagnostics, isCreate bool) *coremodel.DtcServer {
	if m == nil {
		return nil
	}

	obj := &coremodel.DtcServer{}

	// Expand NIOS nested attribute (returns nil if not present)
	niosModel := flex.ExpandNestedObject[NIOSDtcServerModel](ctx, m.NIOS, diags)
	if niosModel != nil {
		obj.NIOS = niosModel.Expand(ctx, diags)
	}

	// Expand UDDI nested attribute (returns nil if not present)
	uddiModel := flex.ExpandNestedObject[UDDIDtcServerModel](ctx, m.UDDI, diags)
	if uddiModel != nil {
		obj.UDDI = uddiModel.Expand(ctx, diags)
	}

	return obj
}

// Expand converts the NIOS TF model to the core model.
func (m *NIOSDtcServerModel) Expand(ctx context.Context, diags *diag.Diagnostics) *coremodel.NIOSDtcServerExt {
	return &coremodel.NIOSDtcServerExt{
		AutoCreateHostRecord: flex.ExpandBoolPointer(m.AutoCreateHostRecord),
		Comment:              flex.ExpandStringPointerNullAsEmpty(m.Comment),
		Disable:              flex.ExpandBoolPointer(m.Disable),
		ExtAttrs:             flex.ExpandMapStringAny(ctx, m.ExtAttrs, diags),
		Host:                 flex.ExpandStringPointerNullAsEmpty(m.Host),
		Monitors:             flex.ExpandFrameworkListNestedBlock(ctx, m.Monitors, diags, ExpandServerMonitors),
		Name:                 flex.ExpandStringPointerNullAsEmpty(m.Name),
		SniHostname:          flex.ExpandStringPointerNullAsEmpty(m.SniHostname),
	}
}

// ApplyDtcServerNIOSUseFlags derives NIOS use flags from the raw config
// value(s) and writes them onto the core model. A flag is true when the user
// set any of its governed value fields in config.
func ApplyDtcServerNIOSUseFlags(ctx context.Context, config tfsdk.Config, obj *coremodel.DtcServer, diags *diag.Diagnostics) {
	if obj == nil || obj.NIOS == nil {
		return
	}
	obj.NIOS.UseSniHostname = flex.DeriveUseFlag(ctx, config, diags, path.Root("nios").AtName("sni_hostname"))
}

// Expand converts the UDDI TF model to the core model.
func (m *UDDIDtcServerModel) Expand(ctx context.Context, diags *diag.Diagnostics) *coremodel.UDDIDtcServerExt {
	return &coremodel.UDDIDtcServerExt{
		Address:                   flex.ExpandStringPointer(m.Address),
		AutoCreateResponseRecords: flex.ExpandBoolPointer(m.AutoCreateResponseRecords),
		Comment:                   flex.ExpandStringPointer(m.Comment),
		Disabled:                  flex.ExpandBoolPointer(m.Disabled),
		EndpointType:              flex.ExpandStringPointer(m.EndpointType),
		Fqdn:                      flex.ExpandStringPointer(m.Fqdn),
		Name:                      flex.ExpandString(m.Name),
		Records:                   flex.ExpandFrameworkListNestedBlock(ctx, m.Records, diags, ExpandRecord),
		Tags:                      flex.ExpandMapStringAny(ctx, m.Tags, diags),
	}
}

// Flatten populates the TF model from a core response.
func (m *DtcServerModel) Flatten(ctx context.Context, resp *coremodel.DtcServer, diags *diag.Diagnostics) {
	if resp == nil {
		return
	}

	m.Id = flex.FlattenStringPointer(resp.Id)

	// Extract existing NIOS model, flatten API response onto it, convert back
	niosModel := flex.ExpandNestedObject[NIOSDtcServerModel](ctx, m.NIOS, diags)
	if niosModel == nil {
		niosModel = &NIOSDtcServerModel{}
	}
	niosModel.Flatten(ctx, resp.NIOS, diags)
	if resp.NIOS != nil {
		m.NIOS = flex.FlattenNestedObject(ctx, niosModel, NIOSDtcServerAttrTypes, diags)
	} else {
		m.NIOS = types.ObjectNull(NIOSDtcServerAttrTypes)
	}

	// Extract existing UDDI model, flatten API response onto it, convert back
	uddiModel := flex.ExpandNestedObject[UDDIDtcServerModel](ctx, m.UDDI, diags)
	if uddiModel == nil {
		uddiModel = &UDDIDtcServerModel{}
	}
	uddiModel.Flatten(ctx, resp.UDDI, diags)
	if resp.UDDI != nil {
		m.UDDI = flex.FlattenNestedObject(ctx, uddiModel, UDDIDtcServerAttrTypes, diags)
	} else {
		m.UDDI = types.ObjectNull(UDDIDtcServerAttrTypes)
	}
}

// Flatten merges API response onto existing NIOS model.
func (m *NIOSDtcServerModel) Flatten(ctx context.Context, from *coremodel.NIOSDtcServerExt, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	planExtAttrs := m.ExtAttrs
	if planExtAttrs.IsUnknown() {
		planExtAttrs = types.MapNull(types.StringType)
	}
	m.AutoCreateHostRecord = flex.FlattenBoolPointer(from.AutoCreateHostRecord)
	m.Comment = flex.FlattenStringPointerEmptyAsNull(from.Comment)
	m.Disable = flex.FlattenBoolPointer(from.Disable)
	m.ExtAttrs, m.ExtAttrsAll = flex.FlattenEAs(planExtAttrs, from.ExtAttrs)
	m.Host = flex.FlattenStringPointerEmptyAsNull(from.Host)
	m.Monitors = flex.FlattenFrameworkListNestedBlock(ctx, from.Monitors, ServerMonitorsAttrTypes, diags, FlattenServerMonitors)
	m.Name = flex.FlattenStringPointerEmptyAsNull(from.Name)
	m.SniHostname = flex.FlattenStringPointerEmptyAsNull(from.SniHostname)
}

// Flatten merges API response onto existing UDDI model.
func (m *UDDIDtcServerModel) Flatten(ctx context.Context, from *coremodel.UDDIDtcServerExt, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	m.Address = flex.FlattenStringPointer(from.Address)
	m.AutoCreateResponseRecords = flex.FlattenBoolPointer(from.AutoCreateResponseRecords)
	m.Comment = flex.FlattenStringPointer(from.Comment)
	m.Disabled = flex.FlattenBoolPointer(from.Disabled)
	m.EndpointType = flex.FlattenStringPointer(from.EndpointType)
	m.Fqdn = flex.FlattenStringPointer(from.Fqdn)
	m.Name = flex.FlattenString(from.Name)
	m.Records = flex.FlattenFrameworkListNestedBlock(ctx, from.Records, RecordAttrTypes, diags, FlattenRecord)
	tagsAll := flex.FlattenMapStringAny(ctx, from.Tags, diags)
	if m.Tags.IsNull() || m.Tags.IsUnknown() {
		m.Tags = tagsAll
	}
	m.TagsAll = tagsAll
}
