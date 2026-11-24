package flex

import (
	"context"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func TestGetInternalID(t *testing.T) {
	mapOf := func(elems map[string]attr.Value) types.Map {
		m, d := types.MapValue(types.StringType, elems)
		if d.HasError() {
			t.Fatalf("building map: %v", d)
		}
		return m
	}

	tests := []struct {
		name string
		in   types.Map
		want string
	}{
		{"null map", types.MapNull(types.StringType), ""},
		{"unknown map", types.MapUnknown(types.StringType), ""},
		{"no entry", mapOf(map[string]attr.Value{"Site": types.StringValue("a")}), ""},
		{"null entry", mapOf(map[string]attr.Value{TerraformInternalID: types.StringNull()}), ""},
		{"set", mapOf(map[string]attr.Value{TerraformInternalID: types.StringValue("id-1")}), "id-1"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetInternalID(tt.in); got != tt.want {
				t.Errorf("GetInternalID() = %q, want %q", got, tt.want)
			}
		})
	}
}

// TestGetInternalID_AfterSetInternalID checks the value SetInternalID adds is read back.
func TestGetInternalID_AfterSetInternalID(t *testing.T) {
	var diags diag.Diagnostics
	m := SetInternalID(context.Background(), types.MapNull(types.StringType), &diags)
	if diags.HasError() {
		t.Fatalf("SetInternalID: %v", diags)
	}
	id := GetInternalID(m)
	if len(id) != 36 {
		t.Errorf("expected a UUID, got %q", id)
	}
}
