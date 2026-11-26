package rir

import (
	"context"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func ripeEAs(t *testing.T, kv map[string]string) types.Map {
	t.Helper()
	elems := map[string]attr.Value{}
	for k, v := range kv {
		elems[k] = types.StringValue(v)
	}
	m, d := types.MapValue(types.StringType, elems)
	if d.HasError() {
		t.Fatal(d)
	}
	return m
}

func missingKeys(resp *resource.ValidateConfigResponse) []string {
	var out []string
	for _, d := range resp.Diagnostics.Errors() {
		if d.Summary() == "Missing Required Extensible Attribute" {
			out = append(out, strings.Split(d.Detail(), "`")[1])
		}
	}
	return out
}

func TestValidateRirOrganizationRequiredRipeEAs(t *testing.T) {
	full := map[string]string{
		"RIPE Admin Contact":     "ib-contact",
		"RIPE Country":           "United Kingdom (GB)",
		"RIPE Technical Contact": "TEST123-IB",
		"RIPE Email":             "support@example.com",
	}
	cases := map[string]struct {
		rir      types.String
		extAttrs types.Map
		missing  int
	}{
		"rir omitted, ext_attrs omitted": {types.StringNull(), types.MapNull(types.StringType), 4},
		"rir omitted, partial":           {types.StringNull(), ripeEAs(t, map[string]string{"RIPE Country": "United Kingdom (GB)"}), 3},
		"rir omitted, complete":          {types.StringNull(), ripeEAs(t, full), 0},
		"rir RIPE, complete":             {types.StringValue("RIPE"), ripeEAs(t, full), 0},
		"rir unknown":                    {types.StringUnknown(), types.MapNull(types.StringType), 0},
		"ext_attrs unknown":              {types.StringNull(), types.MapUnknown(types.StringType), 0},
	}
	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			m := &NIOSRirOrganizationModel{Rir: tc.rir, ExtAttrs: tc.extAttrs}
			resp := &resource.ValidateConfigResponse{}
			validateRirOrganizationNIOSConfig(context.Background(), m, resp)
			if got := missingKeys(resp); len(got) != tc.missing {
				t.Fatalf("expected %d missing keys, got %v (all diags: %v)", tc.missing, got, resp.Diagnostics)
			}
		})
	}
}
