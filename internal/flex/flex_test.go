package flex

import (
	"context"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type testRealm struct {
	Name string
}

var testRealmAttrTypes = map[string]attr.Type{"name": types.StringType}

func expandTestRealm(_ context.Context, o types.Object, _ *diag.Diagnostics) *testRealm {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	return &testRealm{Name: o.Attributes()["name"].(types.String).ValueString()}
}

func TestExpandFrameworkListNestedBlockOmitEmpty(t *testing.T) {
	ctx := context.Background()
	elemType := types.ObjectType{AttrTypes: testRealmAttrTypes}

	realm, d := types.ObjectValue(testRealmAttrTypes, map[string]attr.Value{"name": types.StringValue("r1")})
	if d.HasError() {
		t.Fatalf("building object: %v", d)
	}
	populated, d := types.ListValue(elemType, []attr.Value{realm})
	if d.HasError() {
		t.Fatalf("building list: %v", d)
	}
	empty, _ := types.ListValue(elemType, []attr.Value{})

	cases := map[string]struct {
		in      types.List
		wantNil bool
		wantLen int
	}{
		"null":      {in: types.ListNull(elemType), wantNil: true},
		"unknown":   {in: types.ListUnknown(elemType), wantNil: true},
		"empty":     {in: empty, wantNil: true},
		"populated": {in: populated, wantLen: 1},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			var diags diag.Diagnostics
			got := ExpandFrameworkListNestedBlockOmitEmpty(ctx, tc.in, &diags, expandTestRealm)
			if diags.HasError() {
				t.Fatalf("unexpected diagnostics: %v", diags)
			}
			if tc.wantNil {
				if got != nil {
					t.Fatalf("expected nil slice so the field is omitted, got %#v", got)
				}
				return
			}
			if len(got) != tc.wantLen {
				t.Fatalf("expected %d elements, got %d", tc.wantLen, len(got))
			}
		})
	}

	// The non-omitting helper keeps its existing contract: null expands to a non-nil empty slice.
	var diags diag.Diagnostics
	if got := ExpandFrameworkListNestedBlock(ctx, types.ListNull(elemType), &diags, expandTestRealm); got == nil {
		t.Fatal("ExpandFrameworkListNestedBlock should still return a non-nil empty slice for null")
	}
}
