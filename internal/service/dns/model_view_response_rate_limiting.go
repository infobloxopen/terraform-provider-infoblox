package dns

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	niosdns "github.com/infobloxopen/infoblox-nios-go-client/dns"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
)

// ViewResponseRateLimitingModel is the Terraform model for ViewResponseRateLimiting
type ViewResponseRateLimitingModel struct {
	EnableRrl          types.Bool  `tfsdk:"enable_rrl"`
	LogOnly            types.Bool  `tfsdk:"log_only"`
	ResponsesPerSecond types.Int64 `tfsdk:"responses_per_second"`
	Window             types.Int64 `tfsdk:"window"`
	Slip               types.Int64 `tfsdk:"slip"`
}

// ViewResponseRateLimitingAttrTypes contains the attribute types for ViewResponseRateLimitingModel
var ViewResponseRateLimitingAttrTypes = map[string]attr.Type{
	"enable_rrl":           types.BoolType,
	"log_only":             types.BoolType,
	"responses_per_second": types.Int64Type,
	"window":               types.Int64Type,
	"slip":                 types.Int64Type,
}

// ViewResponseRateLimitingResourceSchemaAttributes contains the schema attributes for ViewResponseRateLimitingModel
var ViewResponseRateLimitingResourceSchemaAttributes = map[string]schema.Attribute{
	"enable_rrl": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "Determines if the response rate limiting is enabled or not.",
	},
	"log_only": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "Determines if logging for response rate limiting without dropping any requests is enabled or not.",
	},
	"responses_per_second": schema.Int64Attribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The number of responses per client per second.",
	},
	"window": schema.Int64Attribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The time interval in seconds over which responses are tracked.",
	},
	"slip": schema.Int64Attribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The response rate limiting slip. Note that if slip is not equal to 0 every n-th rate-limited UDP request is sent a truncated response instead of being dropped.",
	},
}

// ExpandViewResponseRateLimiting converts a Terraform Object to SDK type
func ExpandViewResponseRateLimiting(ctx context.Context, o types.Object, diags *diag.Diagnostics) *niosdns.ViewResponseRateLimiting {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m ViewResponseRateLimitingModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

// Expand converts the Terraform model to SDK type
func (m *ViewResponseRateLimitingModel) Expand(ctx context.Context, diags *diag.Diagnostics) *niosdns.ViewResponseRateLimiting {
	if m == nil {
		return nil
	}
	to := &niosdns.ViewResponseRateLimiting{
		EnableRrl:          flex.ExpandBoolPointer(m.EnableRrl),
		LogOnly:            flex.ExpandBoolPointer(m.LogOnly),
		ResponsesPerSecond: flex.ExpandInt64Pointer(m.ResponsesPerSecond),
		Window:             flex.ExpandInt64Pointer(m.Window),
		Slip:               flex.ExpandInt64Pointer(m.Slip),
	}
	return to
}

// FlattenViewResponseRateLimiting converts an SDK type to Terraform Object
func FlattenViewResponseRateLimiting(ctx context.Context, from *niosdns.ViewResponseRateLimiting, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(ViewResponseRateLimitingAttrTypes)
	}
	m := &ViewResponseRateLimitingModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, ViewResponseRateLimitingAttrTypes, m)
	diags.Append(d...)
	return t
}

// Flatten populates the Terraform model from SDK type
func (m *ViewResponseRateLimitingModel) Flatten(ctx context.Context, from *niosdns.ViewResponseRateLimiting, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	m.EnableRrl = flex.FlattenBoolPointer(from.EnableRrl)
	m.LogOnly = flex.FlattenBoolPointer(from.LogOnly)
	m.ResponsesPerSecond = flex.FlattenInt64Pointer(from.ResponsesPerSecond)
	m.Window = flex.FlattenInt64Pointer(from.Window)
	m.Slip = flex.FlattenInt64Pointer(from.Slip)
}
