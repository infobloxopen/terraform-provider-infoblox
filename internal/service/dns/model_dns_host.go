package dns

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/mapvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/mapdefault"
	objectplanmodifier "github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	stringplanmodifier "github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"

	coremodel "github.com/infobloxopen/terraform-provider-infoblox/internal/core/model/dns"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
	customvalidator "github.com/infobloxopen/terraform-provider-infoblox/internal/validator"
)

type DnsHostModel struct {
	Id            types.String `tfsdk:"id"`
	UpdateTrigger types.String `tfsdk:"update_trigger"`
	UDDI          types.Object `tfsdk:"uddi"`
}

var DnsHostAttrTypes = map[string]attr.Type{
	"id":             types.StringType,
	"update_trigger": types.StringType,
	"uddi":           types.ObjectType{AttrTypes: UDDIDnsHostAttrTypes},
}

type UDDIDnsHostModel struct {
	AbsoluteName       types.String `tfsdk:"absolute_name"`
	Address            types.String `tfsdk:"address"`
	AnycastAddresses   types.List   `tfsdk:"anycast_addresses"`
	AssociatedServer   types.Object `tfsdk:"associated_server"`
	DfpService         types.String `tfsdk:"dfp_service"`
	InheritanceSources types.Object `tfsdk:"inheritance_sources"`
	KerberosKeys       types.List   `tfsdk:"kerberos_keys"`
	Name               types.String `tfsdk:"name"`
	Ophid              types.String `tfsdk:"ophid"`
	ProviderId         types.String `tfsdk:"provider_id"`
	Server             types.String `tfsdk:"server"`
	Tags               types.Map    `tfsdk:"tags"`
	TagsAll            types.Map    `tfsdk:"tags_all"`
	Type               types.String `tfsdk:"type"`
}

var UDDIDnsHostAttrTypes = map[string]attr.Type{
	"absolute_name":       types.StringType,
	"address":             types.StringType,
	"anycast_addresses":   types.ListType{ElemType: types.StringType},
	"associated_server":   types.ObjectType{AttrTypes: HostAssociatedServerAttrTypes},
	"dfp_service":         types.StringType,
	"inheritance_sources": types.ObjectType{AttrTypes: HostInheritanceAttrTypes},
	"kerberos_keys":       types.ListType{ElemType: types.ObjectType{AttrTypes: KerberosKeyAttrTypes}},
	"name":                types.StringType,
	"ophid":               types.StringType,
	"provider_id":         types.StringType,
	"server":              types.StringType,
	"tags":                types.MapType{ElemType: types.StringType},
	"tags_all":            types.MapType{ElemType: types.StringType},
	"type":                types.StringType,
}

const (
	DnsHostInheritanceType = "full"
	DnsHostReturnFields    = ""
)

var DnsHostResourceSchemaAttributes = map[string]schema.Attribute{
	"id": schema.StringAttribute{
		Required:            true,
		MarkdownDescription: "The resource identifier.",
	},
	"update_trigger": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "An arbitrary value used to trigger an update. Not sent to the API. Change it when Terraform reports no infrastructure changes.",
	},
	"uddi": schema.SingleNestedAttribute{
		Optional:            true,
		MarkdownDescription: "UDDI backend-specific fields.",
		Attributes:          DnsHostResourceUddiSchemaAttributes,
	},
}

var DnsHostResourceUddiSchemaAttributes = map[string]schema.Attribute{
	"absolute_name": schema.StringAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "Host FQDN.",
	},
	"address": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "Host's primary IP Address.",
	},
	"anycast_addresses": schema.ListAttribute{
		ElementType:         types.StringType,
		Computed:            true,
		MarkdownDescription: "Anycast address configured to the host. Order is not significant.",
	},
	"associated_server": schema.SingleNestedAttribute{
		Attributes:          HostAssociatedServerResourceSchemaAttributes,
		Optional:            true,
		MarkdownDescription: "Host associated server configuration.",
	},
	"dfp_service": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "DFP service indicates whether or not Universal DDI DNS and Universal TD DFP are both active on the host. If so, Universal DDI DNS will augment recursive queries and forward them to Universal TD DFP. Allowed values:  * _unavailable_: Universal TD DFP application is not available,  * _enabled_: Universal TD DFP application is available and enabled,  * _disabled_: Universal TD DFP application is available but disabled.",
	},
	"inheritance_sources": schema.SingleNestedAttribute{
		Attributes: HostInheritanceResourceSchemaAttributes,
		Optional:   true,
		Computed:   true,
		PlanModifiers: []planmodifier.Object{
			objectplanmodifier.UseStateForUnknown(),
		},
		MarkdownDescription: "Inheritance configuration specifies how and which fields _Host_ object inherits from _Global_ or _Server_ parent.",
	},
	"kerberos_keys": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: KerberosKeyResourceSchemaAttributes,
		},
		Optional: true,
		Computed: true,
		Validators: []validator.List{
			customvalidator.ListNotEmpty(),
		},
		MarkdownDescription: "Optional. _kerberos_keys_ contains a list of keys for GSS-TSIG signed dynamic updates.  Defaults to empty.",
	},
	"name": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "Host display name.",
	},
	"ophid": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "On-Prem Host ID.",
	},
	"provider_id": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "External provider identifier.",
	},
	"server": schema.StringAttribute{
		Optional: true,
		PlanModifiers: []planmodifier.String{
			stringplanmodifier.RequiresReplaceIfConfigured(),
		},
		MarkdownDescription: "The resource identifier.",
	},
	"tags": schema.MapAttribute{
		Optional:    true,
		Computed:    true,
		ElementType: types.StringType,
		Default:     mapdefault.StaticValue(types.MapNull(types.StringType)),
		Validators: []validator.Map{
			mapvalidator.SizeAtLeast(1),
		},
		MarkdownDescription: "Host tagging specifics.",
	},
	"tags_all": schema.MapAttribute{
		Computed:            true,
		ElementType:         types.StringType,
		MarkdownDescription: "All tags including inherited values.",
	},
	"type": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "Defines the type of host. Allowed values:  * _bloxone_ddi_: host type is BloxOne DDI,  * _microsoft_azure_: host type is Microsoft Azure,  * _amazon_web_service_: host type is Amazon Web Services,  * _microsoft_active_directory_: host type is Microsoft Active Directory,  * _google_cloud_platform_: host type is Google Cloud Platform.",
	},
}

// Expand converts the TF model to the infoblox core model
func (m *DnsHostModel) Expand(ctx context.Context, diags *diag.Diagnostics, isCreate bool) *coremodel.DnsHost {
	if m == nil {
		return nil
	}

	obj := &coremodel.DnsHost{}

	// Expand UDDI nested attribute (returns nil if not present)
	uddiModel := flex.ExpandNestedObject[UDDIDnsHostModel](ctx, m.UDDI, diags)
	if uddiModel != nil {
		obj.UDDI = uddiModel.Expand(ctx, diags)
	}

	return obj
}

// Expand converts the UDDI TF model to the core model.
func (m *UDDIDnsHostModel) Expand(ctx context.Context, diags *diag.Diagnostics) *coremodel.UDDIDnsHostExt {
	return &coremodel.UDDIDnsHostExt{
		AbsoluteName:       flex.ExpandStringPointer(m.AbsoluteName),
		AssociatedServer:   ExpandHostAssociatedServer(ctx, m.AssociatedServer, diags),
		InheritanceSources: ExpandHostInheritance(ctx, m.InheritanceSources, diags),
		KerberosKeys:       flex.ExpandFrameworkListNestedBlock(ctx, m.KerberosKeys, diags, ExpandKerberosKey),
		Server:             flex.ExpandStringPointer(m.Server),
		Tags:               flex.ExpandMapStringAny(ctx, m.Tags, diags),
	}
}

// Flatten populates the TF model from a core response.
func (m *DnsHostModel) Flatten(ctx context.Context, resp *coremodel.DnsHost, diags *diag.Diagnostics) {
	if resp == nil {
		return
	}

	m.Id = flex.FlattenStringPointer(resp.Id)

	// Extract existing UDDI model, flatten API response onto it, convert back
	uddiModel := flex.ExpandNestedObject[UDDIDnsHostModel](ctx, m.UDDI, diags)
	if uddiModel == nil {
		uddiModel = &UDDIDnsHostModel{}
	}
	uddiModel.Flatten(ctx, resp.UDDI, diags)
	if resp.UDDI != nil {
		m.UDDI = flex.FlattenNestedObject(ctx, uddiModel, UDDIDnsHostAttrTypes, diags)
	} else {
		m.UDDI = types.ObjectNull(UDDIDnsHostAttrTypes)
	}
}

// Flatten merges API response onto existing UDDI model.
func (m *UDDIDnsHostModel) Flatten(ctx context.Context, from *coremodel.UDDIDnsHostExt, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	m.AbsoluteName = flex.FlattenStringPointer(from.AbsoluteName)
	m.Address = flex.FlattenStringPointer(from.Address)
	m.AnycastAddresses = flex.FlattenFrameworkListString(ctx, from.AnycastAddresses, diags)
	m.AssociatedServer = FlattenHostAssociatedServer(ctx, from.AssociatedServer, diags)
	m.DfpService = flex.FlattenStringPointer(from.DfpService)
	m.InheritanceSources = FlattenHostInheritance(ctx, from.InheritanceSources, diags)
	m.KerberosKeys = flex.FlattenFrameworkListNestedBlock(ctx, from.KerberosKeys, KerberosKeyAttrTypes, diags, FlattenKerberosKey)
	m.Name = flex.FlattenStringPointer(from.Name)
	m.Ophid = flex.FlattenStringPointer(from.Ophid)
	m.ProviderId = flex.FlattenStringPointer(from.ProviderId)
	m.Server = flex.FlattenStringPointer(from.Server)
	tagsAll := flex.FlattenMapStringAny(ctx, from.Tags, diags)
	if m.Tags.IsNull() || m.Tags.IsUnknown() {
		m.Tags = tagsAll
	}
	m.TagsAll = tagsAll
	m.Type = flex.FlattenStringPointer(from.Type)
}
