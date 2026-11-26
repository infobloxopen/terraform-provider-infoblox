package types

import (
	"context"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func TestCaseInsensitiveStringEqual(t *testing.T) {
	a := CaseInsensitiveString{basetypes.NewStringValue("Example.com")}
	if !a.Equal(CaseInsensitiveString{basetypes.NewStringValue("Example.com")}) {
		t.Fatal("identical values must be Equal")
	}
	if a.Equal(CaseInsensitiveString{basetypes.NewStringValue("example.com")}) {
		t.Fatal("Equal is exact; case differences are handled by semantic equality")
	}
	if a.Equal(basetypes.NewStringValue("Example.com")) {
		t.Fatal("a different value type must not be Equal")
	}
}

func TestCaseInsensitiveStringSemanticEquals(t *testing.T) {
	ctx := context.Background()
	old := CaseInsensitiveString{basetypes.NewStringValue("Example.COM")}

	cases := map[string]struct {
		in   basetypes.StringValuable
		want bool
	}{
		"custom type, different case":  {CaseInsensitiveString{basetypes.NewStringValue("example.com")}, true},
		"plain string, different case": {basetypes.NewStringValue("example.com"), true},
		"different value":              {basetypes.NewStringValue("other.com"), false},
		"null":                         {basetypes.NewStringNull(), false},
	}
	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			got, diags := old.StringSemanticEquals(ctx, tc.in)
			if diags.HasError() {
				t.Fatalf("diags: %v", diags)
			}
			if got != tc.want {
				t.Fatalf("got %v, want %v", got, tc.want)
			}
		})
	}
}
