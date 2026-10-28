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
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/hashicorp/terraform-plugin-go/tftypes"

	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
)

// Legacy describes how a Legacy Infoblox Provider (v2.x) resource converts into a v3 resource.
type Legacy struct {
	// Type is the v2 resource type, e.g. infoblox_a_record.
	Type string
	// Defaults seeds nios attributes that NIOS never returns, e.g. restart_if_needed.
	Defaults map[string]any
	// FromV2 copies v2 string or bool attributes into nios attributes the v3 Read keeps from state (v3 name -> v2 name).
	FromV2 map[string]string
}

// Mover handles `moved` blocks from a renamed v2 type, e.g. infoblox_a_record -> infoblox_record_a.
func (l Legacy) Mover() resource.StateMover {
	return resource.StateMover{StateMover: l.move}
}

// Upgrader handles v2 state of a type that kept its name in v3, e.g. infoblox_zone_auth.
// Terraform never calls MoveState for those, only UpgradeState. Register it at schema version 0
// and bump the v3 schema to version 1.
func (l Legacy) Upgrader() resource.StateUpgrader {
	return resource.StateUpgrader{StateUpgrader: l.upgrade}
}

func (l Legacy) move(ctx context.Context, req resource.MoveStateRequest, resp *resource.MoveStateResponse) {
	// Returning without a state lets Terraform report the unsupported source.
	if req.SourceTypeName != l.Type || !strings.HasSuffix(req.SourceProviderAddress, "infobloxopen/infoblox") {
		return
	}
	ref := l.convert(ctx, req.SourceRawState.JSON, &resp.TargetState, &resp.Diagnostics)
	if ref != "" && resp.TargetIdentity != nil {
		resp.Diagnostics.Append(resp.TargetIdentity.SetAttribute(ctx, path.Root("id"), ref)...)
	}
}

func (l Legacy) upgrade(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
	schemaType := resp.State.Schema.Type().TerraformType(ctx)

	// 3.0.0-beta state is also version 0 but already has the v3 shape; keep it as is.
	if isV3(req.RawState.JSON) {
		v, err := req.RawState.UnmarshalWithOpts(schemaType, tfprotov6.UnmarshalOpts{
			ValueFromJSONOpts: tftypes.ValueFromJSONOpts{IgnoreUndefinedAttributes: true},
		})
		if err != nil {
			resp.Diagnostics.AddError("Unable to Upgrade State", fmt.Sprintf("Cannot read %s state: %s", l.Type, err))
			return
		}
		resp.State.Raw = v
		return
	}

	resp.State.Raw = tftypes.NewValue(schemaType, nil)
	l.convert(ctx, req.RawState.JSON, &resp.State, &resp.Diagnostics)
}

// convert writes only what the v3 Read cannot rebuild from NIOS: the ref, the Terraform
// Internal ID and the configured EAs. The refresh that follows fills in the rest.
func (l Legacy) convert(ctx context.Context, raw []byte, state *tfsdk.State, diags *diag.Diagnostics) string {
	var v2 map[string]json.RawMessage
	if err := json.Unmarshal(raw, &v2); err != nil {
		diags.AddError("Unable to Convert State", fmt.Sprintf("Cannot parse %s state: %s", l.Type, err))
		return ""
	}

	// ip_allocation keeps its internal ID in id and the host ref in ref.
	ref := stringAttr(v2, "ref")
	if ref == "" {
		ref = stringAttr(v2, "id")
	}
	if ref == "" {
		diags.AddError("Unable to Convert State", l.Type+" state has no object reference.")
		return ""
	}

	set := func(p path.Path, v any) { diags.Append(state.SetAttribute(ctx, p, v)...) }
	nios := path.Root("nios")

	set(path.Root("id"), ref)
	for name, v := range l.Defaults {
		set(nios.AtName(name), v)
	}
	for v3Name, v2Name := range l.FromV2 {
		var v any
		if json.Unmarshal(v2[v2Name], &v) != nil {
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

	if _, d := state.Schema.AttributeAtPath(ctx, nios.AtName("ext_attrs_all")); d.HasError() {
		return ref
	}

	// v3 Read finds the object through this EA and fails without ext_attrs_all.
	eaAll := map[string]string{}
	if id := stringAttr(v2, "internal_id"); id != "" {
		eaAll[flex.TerraformInternalID] = id
	} else {
		diags.AddWarning("Missing Internal ID", l.Type+" has no internal_id; v3 cannot find it again if its ref changes.")
	}
	set(nios.AtName("ext_attrs_all"), eaAll)

	// v2 state keeps only the EAs set in config, so they carry over as configured EAs.
	eas, err := legacyEAs(stringAttr(v2, "ext_attrs"))
	if err != nil {
		diags.AddError("Unable to Convert State", fmt.Sprintf("%s ext_attrs: %s", l.Type, err))
		return ""
	}
	if len(eas) > 0 {
		set(nios.AtName("ext_attrs"), eas)
	}
	return ref
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

// legacyEAs converts the v2 ext_attrs JSON string into the v3 string map; non-string values stay compact JSON.
func legacyEAs(raw string) (map[string]string, error) {
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
