package misc

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"

	coremodel "github.com/infobloxopen/terraform-provider-infoblox/internal/core/model/misc"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
	customvalidator "github.com/infobloxopen/terraform-provider-infoblox/internal/validator"
)

type BfdtemplateModel struct {
	Id   types.String `tfsdk:"id"`
	NIOS types.Object `tfsdk:"nios"`
}

var BfdtemplateAttrTypes = map[string]attr.Type{
	"id":   types.StringType,
	"nios": types.ObjectType{AttrTypes: NIOSBfdtemplateAttrTypes},
}

type NIOSBfdtemplateModel struct {
	AuthenticationKey   types.String `tfsdk:"authentication_key"`
	AuthenticationKeyId types.Int64  `tfsdk:"authentication_key_id"`
	AuthenticationType  types.String `tfsdk:"authentication_type"`
	DetectionMultiplier types.Int64  `tfsdk:"detection_multiplier"`
	MinRxInterval       types.Int64  `tfsdk:"min_rx_interval"`
	MinTxInterval       types.Int64  `tfsdk:"min_tx_interval"`
	Name                types.String `tfsdk:"name"`
}

var NIOSBfdtemplateAttrTypes = map[string]attr.Type{
	"authentication_key":    types.StringType,
	"authentication_key_id": types.Int64Type,
	"authentication_type":   types.StringType,
	"detection_multiplier":  types.Int64Type,
	"min_rx_interval":       types.Int64Type,
	"min_tx_interval":       types.Int64Type,
	"name":                  types.StringType,
}

const (
	BfdtemplateReturnFields = "authentication_key_id,authentication_type,detection_multiplier,min_rx_interval,min_tx_interval,name"
)

var BfdtemplateResourceSchemaAttributes = map[string]schema.Attribute{
	"id": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The reference to the object.",
	},
	"nios": schema.SingleNestedAttribute{
		Optional:            true,
		MarkdownDescription: "NIOS backend-specific fields.",
		Attributes:          BfdtemplateResourceNiosSchemaAttributes,
	},
}

var BfdtemplateResourceNiosSchemaAttributes = map[string]schema.Attribute{
	"authentication_key": schema.StringAttribute{
		Optional: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
		},
		MarkdownDescription: "The authentication key for BFD protocol message-digest authentication.",
	},
	"authentication_key_id": schema.Int64Attribute{
		Optional: true,
		Computed: true,
		Default:  int64default.StaticInt64(1),
		Validators: []validator.Int64{
			int64validator.Between(1, 255),
		},
		MarkdownDescription: "The authentication key identifier for BFD protocol authentication. Valid values are between 1 and 255.",
	},
	"authentication_type": schema.StringAttribute{
		Default: stringdefault.StaticString("NONE"),
		Validators: []validator.String{
			stringvalidator.OneOf("NONE", "MD5", "METICULOUS-MD5", "SHA1", "METICULOUS-SHA1"),
		},
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The authentication type for BFD protocol.",
	},
	"detection_multiplier": schema.Int64Attribute{
		Optional: true,
		Computed: true,
		Default:  int64default.StaticInt64(3),
		Validators: []validator.Int64{
			int64validator.Between(3, 50),
		},
		MarkdownDescription: "The detection time multiplier value for BFD protocol. The negotiated transmit interval, multiplied by this value, provides the detection time for the receiving system in asynchronous BFD mode. Valid values are between 3 and 50.",
	},
	"min_rx_interval": schema.Int64Attribute{
		Optional: true,
		Computed: true,
		Default:  int64default.StaticInt64(100),
		Validators: []validator.Int64{
			int64validator.Between(50, 9999),
		},
		MarkdownDescription: "The minimum receive time (in seconds) for BFD protocol. Valid values are between 50 and 9999.",
	},
	"min_tx_interval": schema.Int64Attribute{
		Optional: true,
		Computed: true,
		Default:  int64default.StaticInt64(100),
		Validators: []validator.Int64{
			int64validator.Between(50, 9999),
		},
		MarkdownDescription: "The minimum transmission time (in seconds) for BFD protocol. Valid values are between 50 and 9999.",
	},
	"name": schema.StringAttribute{
		Required: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
			customvalidator.ValidateTrimmedString(),
		},
		MarkdownDescription: "The name of the BFD template object.",
	},
}

// Expand converts the TF model to the infoblox core model
func (m *BfdtemplateModel) Expand(ctx context.Context, diags *diag.Diagnostics, isCreate bool) *coremodel.Bfdtemplate {
	if m == nil {
		return nil
	}

	obj := &coremodel.Bfdtemplate{}

	// Expand NIOS nested attribute (returns nil if not present)
	niosModel := flex.ExpandNestedObject[NIOSBfdtemplateModel](ctx, m.NIOS, diags)
	if niosModel != nil {
		obj.NIOS = niosModel.Expand(ctx, diags)
	}

	return obj
}

// Expand converts the NIOS TF model to the core model.
func (m *NIOSBfdtemplateModel) Expand(ctx context.Context, diags *diag.Diagnostics) *coremodel.NIOSBfdtemplateExt {
	return &coremodel.NIOSBfdtemplateExt{
		AuthenticationKey:   flex.ExpandStringPointerNullAsEmpty(m.AuthenticationKey),
		AuthenticationKeyId: flex.ExpandInt64Pointer(m.AuthenticationKeyId),
		AuthenticationType:  flex.ExpandStringPointerNullAsEmpty(m.AuthenticationType),
		DetectionMultiplier: flex.ExpandInt64Pointer(m.DetectionMultiplier),
		MinRxInterval:       flex.ExpandInt64Pointer(m.MinRxInterval),
		MinTxInterval:       flex.ExpandInt64Pointer(m.MinTxInterval),
		Name:                flex.ExpandStringPointerNullAsEmpty(m.Name),
	}
}

// Flatten populates the TF model from a core response.
func (m *BfdtemplateModel) Flatten(ctx context.Context, resp *coremodel.Bfdtemplate, diags *diag.Diagnostics) {
	if resp == nil {
		return
	}

	m.Id = flex.FlattenStringPointer(resp.Id)

	// Extract existing NIOS model, flatten API response onto it, convert back
	niosModel := flex.ExpandNestedObject[NIOSBfdtemplateModel](ctx, m.NIOS, diags)
	if niosModel == nil {
		niosModel = &NIOSBfdtemplateModel{}
	}
	niosModel.Flatten(ctx, resp.NIOS, diags)
	if resp.NIOS != nil {
		m.NIOS = flex.FlattenNestedObject(ctx, niosModel, NIOSBfdtemplateAttrTypes, diags)
	} else {
		m.NIOS = types.ObjectNull(NIOSBfdtemplateAttrTypes)
	}

}

// Flatten merges API response onto existing NIOS model.
func (m *NIOSBfdtemplateModel) Flatten(ctx context.Context, from *coremodel.NIOSBfdtemplateExt, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	m.AuthenticationKeyId = flex.FlattenInt64Pointer(from.AuthenticationKeyId)
	m.AuthenticationType = flex.FlattenStringPointerEmptyAsNull(from.AuthenticationType)
	m.DetectionMultiplier = flex.FlattenInt64Pointer(from.DetectionMultiplier)
	m.MinRxInterval = flex.FlattenInt64Pointer(from.MinRxInterval)
	m.MinTxInterval = flex.FlattenInt64Pointer(from.MinTxInterval)
	m.Name = flex.FlattenStringPointerEmptyAsNull(from.Name)
}
