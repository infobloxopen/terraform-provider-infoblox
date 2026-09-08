package cloud

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"

	coremodel "github.com/infobloxopen/terraform-provider-infoblox/internal/core/model/cloud"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
	customvalidator "github.com/infobloxopen/terraform-provider-infoblox/internal/validator"
)

type AwsuserModel struct {
	Id   types.String `tfsdk:"id"`
	NIOS types.Object `tfsdk:"nios"`
}

var AwsuserAttrTypes = map[string]attr.Type{
	"id":   types.StringType,
	"nios": types.ObjectType{AttrTypes: NIOSAwsuserAttrTypes},
}

type NIOSAwsuserModel struct {
	AccessKeyId     types.String `tfsdk:"access_key_id"`
	AccountId       types.String `tfsdk:"account_id"`
	GovcloudEnabled types.Bool   `tfsdk:"govcloud_enabled"`
	Name            types.String `tfsdk:"name"`
	NiosUserName    types.String `tfsdk:"nios_user_name"`
	SecretAccessKey types.String `tfsdk:"secret_access_key"`
}

var NIOSAwsuserAttrTypes = map[string]attr.Type{
	"access_key_id":     types.StringType,
	"account_id":        types.StringType,
	"govcloud_enabled":  types.BoolType,
	"name":              types.StringType,
	"nios_user_name":    types.StringType,
	"secret_access_key": types.StringType,
}

const (
	AwsuserReturnFields = "access_key_id,account_id,govcloud_enabled,last_used,name,nios_user_name,status"
)

var AwsuserResourceSchemaAttributes = map[string]schema.Attribute{
	"id": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The reference to the object.",
	},
	"nios": schema.SingleNestedAttribute{
		Optional:            true,
		MarkdownDescription: "NIOS backend-specific fields.",
		Attributes:          AwsuserResourceNiosSchemaAttributes,
	},
}

var AwsuserResourceNiosSchemaAttributes = map[string]schema.Attribute{
	"access_key_id": schema.StringAttribute{
		Required: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
			customvalidator.ValidateTrimmedString(),
			stringvalidator.LengthAtMost(255),
		},
		MarkdownDescription: "The unique Access Key ID of this AWS user. Maximum 255 characters.",
	},
	"account_id": schema.StringAttribute{
		Required: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
			stringvalidator.LengthAtMost(64),
		},
		MarkdownDescription: "The AWS Account ID of this AWS user. Maximum 64 characters.",
	},
	"govcloud_enabled": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Indicates if gov cloud is enabled or disabled.",
	},
	"name": schema.StringAttribute{
		Required: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
			stringvalidator.LengthAtMost(64),
		},
		MarkdownDescription: "The AWS user name. Maximum 64 characters.",
	},
	"nios_user_name": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
			stringvalidator.LengthAtMost(64),
		},
		MarkdownDescription: "The NIOS user name mapped to this AWS user. Maximum 64 characters.",
	},
	"secret_access_key": schema.StringAttribute{
		Sensitive: true,
		Required:  true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
			customvalidator.ValidateTrimmedString(),
		},
		MarkdownDescription: "The Secret Access Key for the Access Key ID of this user. Maximum 255 characters.",
	},
}

// Expand converts the TF model to the infoblox core model
func (m *AwsuserModel) Expand(ctx context.Context, diags *diag.Diagnostics, isCreate bool) *coremodel.Awsuser {
	if m == nil {
		return nil
	}

	obj := &coremodel.Awsuser{}

	// Expand NIOS nested attribute (returns nil if not present)
	niosModel := flex.ExpandNestedObject[NIOSAwsuserModel](ctx, m.NIOS, diags)
	if niosModel != nil {
		obj.NIOS = niosModel.Expand(ctx, diags)
	}

	return obj
}

// Expand converts the NIOS TF model to the core model.
func (m *NIOSAwsuserModel) Expand(ctx context.Context, diags *diag.Diagnostics) *coremodel.NIOSAwsuserExt {
	return &coremodel.NIOSAwsuserExt{
		AccessKeyId:     flex.ExpandStringPointerNullAsEmpty(m.AccessKeyId),
		AccountId:       flex.ExpandStringPointerNullAsEmpty(m.AccountId),
		GovcloudEnabled: flex.ExpandBoolPointer(m.GovcloudEnabled),
		Name:            flex.ExpandStringPointerNullAsEmpty(m.Name),
		NiosUserName:    flex.ExpandStringPointer(m.NiosUserName),
		SecretAccessKey: flex.ExpandStringPointerNullAsEmpty(m.SecretAccessKey),
	}
}

// Flatten populates the TF model from a core response.
func (m *AwsuserModel) Flatten(ctx context.Context, resp *coremodel.Awsuser, diags *diag.Diagnostics) {
	if resp == nil {
		return
	}

	m.Id = flex.FlattenStringPointer(resp.Id)

	// Extract existing NIOS model, flatten API response onto it, convert back
	niosModel := flex.ExpandNestedObject[NIOSAwsuserModel](ctx, m.NIOS, diags)
	if niosModel == nil {
		niosModel = &NIOSAwsuserModel{}
	}
	niosModel.Flatten(ctx, resp.NIOS, diags)
	if resp.NIOS != nil {
		m.NIOS = flex.FlattenNestedObject(ctx, niosModel, NIOSAwsuserAttrTypes, diags)
	} else {
		m.NIOS = types.ObjectNull(NIOSAwsuserAttrTypes)
	}

}

// Flatten merges API response onto existing NIOS model.
func (m *NIOSAwsuserModel) Flatten(ctx context.Context, from *coremodel.NIOSAwsuserExt, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	m.AccessKeyId = flex.FlattenStringPointerEmptyAsNull(from.AccessKeyId)
	m.AccountId = flex.FlattenStringPointerEmptyAsNull(from.AccountId)
	m.GovcloudEnabled = flex.FlattenBoolPointer(from.GovcloudEnabled)
	m.Name = flex.FlattenStringPointerEmptyAsNull(from.Name)
	m.NiosUserName = flex.FlattenStringPointerEmptyAsNull(from.NiosUserName)
}
