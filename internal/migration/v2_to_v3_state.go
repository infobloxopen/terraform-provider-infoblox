package migration

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/hashicorp/terraform-plugin-go/tftypes"

	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
)

// V2Resource describes how a v2 (2.x) resource converts into a v3 resource.
type V2Resource struct {
	// Type is the v2 resource type, e.g. infoblox_a_record.
	Type string
	// Defaults seeds nios attributes that NIOS never returns, e.g. restart_if_needed.
	Defaults map[string]any
	// FromV2 copies v2 string or bool attributes into nios attributes the v3 Read keeps from state (v3 name -> v2 name).
	FromV2 map[string]string
	// EmptyOptions starts nios.options as an empty list when v2 configured none. v3 Read keeps only
	// the options already in state, or all of them when state has none, and NIOS always returns
	// a built-in dhcp-lease-time option.
	EmptyOptions bool
}

const (
	ipAllocationType  = "infoblox_ip_allocation"
	ipAssociationType = "infoblox_ip_association"

	// A newer v2 stores both `ref` and `internal_id` on apply.
	upgradeV2Hint = "Upgrade to v2.6.0 or later, run `terraform apply` to store the missing values, then migrate again."
)

// Mover handles `moved` blocks from a renamed v2 type, e.g. infoblox_a_record -> infoblox_record_a.
func (r V2Resource) Mover() resource.StateMover {
	return resource.StateMover{StateMover: r.move}
}

// Upgrader handles v2 state of a type that kept its name in v3, e.g. infoblox_zone_auth.
// Terraform never calls MoveState for those, only UpgradeState.
func (r V2Resource) Upgrader() resource.StateUpgrader {
	return resource.StateUpgrader{StateUpgrader: r.upgrade}
}

func (r V2Resource) move(ctx context.Context, req resource.MoveStateRequest, resp *resource.MoveStateResponse) {
	// Returning without a state lets Terraform report the unsupported source.
	// The address is <host>/<namespace>/<type>; any host, so private registry mirrors work too.
	if req.SourceTypeName != r.Type || !strings.HasSuffix(req.SourceProviderAddress, "/infobloxopen/infoblox") {
		return
	}
	ref := r.convert(ctx, req.SourceRawState.JSON, &resp.TargetState, &resp.Diagnostics)
	if ref != "" && resp.TargetIdentity != nil {
		resp.Diagnostics.Append(resp.TargetIdentity.SetAttribute(ctx, path.Root("id"), ref)...)
	}
}

func (r V2Resource) upgrade(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
	schemaType := resp.State.Schema.Type().TerraformType(ctx)

	// 3.0.0-beta state is also version 0 but already has the v3 shape; keep it as is.
	if isV3(req.RawState.JSON) {
		v, err := req.RawState.UnmarshalWithOpts(schemaType, tfprotov6.UnmarshalOpts{
			ValueFromJSONOpts: tftypes.ValueFromJSONOpts{IgnoreUndefinedAttributes: true},
		})
		if err != nil {
			resp.Diagnostics.AddError("Unable to Upgrade State", fmt.Sprintf("Cannot read %s state: %s", r.Type, err))
			return
		}
		resp.State.Raw = v
		return
	}

	resp.State.Raw = tftypes.NewValue(schemaType, nil)
	r.convert(ctx, req.RawState.JSON, &resp.State, &resp.Diagnostics)
}

// convert writes only what the v3 Read cannot rebuild from NIOS: the ref, the Terraform
// Internal ID and the configured EAs. The refresh that follows fills in the rest.
func (r V2Resource) convert(ctx context.Context, raw []byte, state *tfsdk.State, diags *diag.Diagnostics) string {
	var v2 map[string]json.RawMessage
	if err := json.Unmarshal(raw, &v2); err != nil {
		diags.AddError("Unable to Convert State", fmt.Sprintf("Cannot parse %s state: %s", r.Type, err))
		return ""
	}
	if r.Type == ipAssociationType {
		return r.convertAssociation(ctx, v2, state, diags)
	}

	nios := path.Root("nios")

	// v3 Read finds the object through this EA. Checked before the ref, because state from
	// Legacy Infoblox Provider versions before v2.6.0 has neither. record_ns has no EAs.
	hasEAs := hasAttr(ctx, state, nios.AtName("ext_attrs_all"))
	internalID := stringAttr(v2, "internal_id")
	if hasEAs && internalID == "" {
		diags.AddError("Missing Internal ID", fmt.Sprintf(
			"%s state has no internal_id, so it was written by a Legacy Infoblox Provider (v2.x.x) "+
				"older than v2.6.0 and cannot be migrated. %s", r.Type, upgradeV2Hint))
		return ""
	}

	// Every v2 resource keeps the NIOS ref in `ref`. Its `id` is the same ref, except for
	// ip_allocation: v2.1.0 stores "<internal ID><separator><ref>", later versions the Internal ID.
	ref := stringAttr(v2, "ref")
	if ref == "" && r.Type != ipAllocationType {
		ref = stringAttr(v2, "id")
	}
	if ref == "" {
		diags.AddError("Unable to Convert State", fmt.Sprintf(
			"%s state has no object reference (ref). %s", r.Type, upgradeV2Hint))
		return ""
	}

	set := func(p path.Path, v any) { diags.Append(state.SetAttribute(ctx, p, v)...) }

	set(path.Root("id"), ref)
	// Skip attributes the target schema no longer has, like the writes below.
	for name, v := range r.Defaults {
		if hasAttr(ctx, state, nios.AtName(name)) {
			set(nios.AtName(name), v)
		}
	}
	// Only v2 zone_auth has it; a v2 `true` would otherwise show as a change.
	if b, ok := boolAttr(v2, "restart_if_needed"); ok && hasAttr(ctx, state, nios.AtName("restart_if_needed")) {
		set(nios.AtName("restart_if_needed"), b)
	}
	for v3Name, v2Name := range r.FromV2 {
		var v any
		if !hasAttr(ctx, state, nios.AtName(v3Name)) || json.Unmarshal(v2[v2Name], &v) != nil {
			continue
		}
		switch v := v.(type) {
		case bool:
			set(nios.AtName(v3Name), v)
		case string:
			if strings.TrimSpace(v) != "" {
				set(nios.AtName(v3Name), v)
			}
		}
	}

	// fixed_address, range, range_template and shared_network: v2 `options` blocks.
	if hasAttr(ctx, state, nios.AtName("options")) {
		if opts := configuredOptions(v2["options"]); len(opts) > 0 || r.EmptyOptions {
			set(nios.AtName("options"), opts)
		}
	}

	if !hasEAs {
		return ref
	}
	set(nios.AtName("ext_attrs_all"), map[string]string{flex.TerraformInternalID: internalID})

	// v2 state keeps only the EAs set in config, so they carry over as configured EAs.
	eas, err := parseEAs(stringAttr(v2, "ext_attrs"))
	if err != nil {
		diags.AddError("Unable to Convert State", fmt.Sprintf("%s ext_attrs: %s", r.Type, err))
		return ""
	}
	if len(eas) > 0 {
		set(nios.AtName("ext_attrs"), eas)
	}
	return ref
}

// convertAssociation handles ip_association: v3 has only `nios`, the host record link plus the
// DHCP settings, which Read refreshes from the host. They are seeded so an address family the
// host lacks matches the defaults.
func (r V2Resource) convertAssociation(ctx context.Context, v2 map[string]json.RawMessage, state *tfsdk.State, diags *diag.Diagnostics) string {
	ref := stringAttr(v2, "ref")
	if ref == "" {
		diags.AddError("Unable to Convert State", fmt.Sprintf(
			"%s state has no host record reference (ref). %s", r.Type, upgradeV2Hint))
		return ""
	}
	dhcp, _ := boolAttr(v2, "enable_dhcp")
	nios := path.Root("nios")
	for name, v := range map[string]any{
		"ref":                ref,
		"internal_id":        stringAttr(v2, "internal_id"),
		"mac":                stringAttr(v2, "mac_addr"),
		"duid":               stringAttr(v2, "duid"),
		"configure_for_dhcp": dhcp,
		"match_client":       "DUID",
	} {
		diags.Append(state.SetAttribute(ctx, nios.AtName(name), v)...)
	}
	return ref
}

// dhcpOption is the v2 `options` block, and the v3 options element of every type that has one.
type dhcpOption struct {
	Name        types.String `tfsdk:"name"`
	Num         types.Int64  `tfsdk:"num"`
	UseOption   types.Bool   `tfsdk:"use_option"`
	Value       types.String `tfsdk:"value"`
	VendorClass types.String `tfsdk:"vendor_class"`
}

// configuredOptions returns the DHCP options a v2 configuration set, from the v2 state.
// v2 stores every option NIOS returns and hides two kinds from plans: the built-in
// dhcp-lease-time, and options removed from the config, which v2 leaves in NIOS with an
// empty value. v3 Read keeps only the options already in state, so it must hold exactly
// the configured ones, in NIOS order (the config order needs a plan, which a mover doesn't have).
func configuredOptions(raw json.RawMessage) []dhcpOption {
	var in []struct {
		Name        string `json:"name"`
		Num         int64  `json:"num"`
		UseOption   bool   `json:"use_option"`
		Value       string `json:"value"`
		VendorClass string `json:"vendor_class"`
	}
	_ = json.Unmarshal(raw, &in)
	out := []dhcpOption{}
	for _, o := range in {
		if o.Name == "dhcp-lease-time" && o.Num == 51 && !o.UseOption && o.Value == "43200" && o.VendorClass == "DHCP" {
			continue
		}
		if !o.UseOption && o.Value == "" {
			continue
		}
		out = append(out, dhcpOption{
			Name:        types.StringValue(o.Name),
			Num:         types.Int64Value(o.Num),
			UseOption:   types.BoolValue(o.UseOption),
			Value:       types.StringValue(o.Value),
			VendorClass: types.StringValue(o.VendorClass),
		})
	}
	return out
}

func hasAttr(ctx context.Context, state *tfsdk.State, p path.Path) bool {
	_, d := state.Schema.AttributeAtPath(ctx, p)
	return !d.HasError()
}

func boolAttr(attrs map[string]json.RawMessage, name string) (bool, bool) {
	var b bool
	return b, json.Unmarshal(attrs[name], &b) == nil
}

func isV3(raw []byte) bool {
	var attrs map[string]json.RawMessage
	if json.Unmarshal(raw, &attrs) != nil {
		return false
	}
	_, nios := attrs["nios"]
	_, uddi := attrs["uddi"]
	return nios || uddi
}

func stringAttr(attrs map[string]json.RawMessage, name string) string {
	var s string
	_ = json.Unmarshal(attrs[name], &s)
	return s
}

// parseEAs converts the v2 ext_attrs JSON string into the v3 string map; non-string values stay compact JSON.
func parseEAs(raw string) (map[string]string, error) {
	if raw == "" {
		return nil, nil
	}
	var in map[string]json.RawMessage
	if err := json.Unmarshal([]byte(raw), &in); err != nil {
		return nil, err
	}
	out := make(map[string]string, len(in))
	for k, v := range in {
		if k == flex.TerraformInternalID {
			continue
		}
		var s string
		if json.Unmarshal(v, &s) == nil {
			out[k] = s
			continue
		}
		var buf bytes.Buffer
		if err := json.Compact(&buf, v); err != nil {
			return nil, err
		}
		out[k] = buf.String()
	}
	return out, nil
}
