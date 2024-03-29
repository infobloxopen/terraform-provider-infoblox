package infoblox

import (
	"os"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var testAccProviders map[string]*schema.Provider
var testAccProvider *schema.Provider

type testCase struct {
	val         interface{}
	f           schema.SchemaValidateFunc
	expectedErr *regexp.Regexp
}

func init() {
	testAccProvider = Provider()
	testAccProviders = map[string]*schema.Provider{
		"infoblox": testAccProvider,
	}
}

func TestProvider(t *testing.T) {
	if err := Provider().InternalValidate(); err != nil {
		t.Fatalf("err: %s", err)
	}

}

func TestProvider_impl(t *testing.T) {
	var _ = Provider()
}

func testAccPreCheck(t *testing.T) {
	if v := os.Getenv("INFOBLOX_USERNAME"); v == "" {
		t.Fatal("INFOBLOX_USERNAME must be set for acceptance tests")
	}

	if v := os.Getenv("INFOBLOX_PASSWORD"); v == "" {
		t.Fatal("INFOBLOX_PASSWORD must be set for acceptance tests")
	}

	if v := os.Getenv("INFOBLOX_SERVER"); v == "" {
		t.Fatal("INFOBLOX_SERVER must be set for acceptance tests")
	}

	_ = os.Setenv("INFOBLOX_CLIENT_TIMEOUT", "60")
}

func runTestCases(t *testing.T, cases []testCase) {
	matchErr := func(errs []error, r *regexp.Regexp) bool {
		// err must match one provided
		for _, err := range errs {
			if r.MatchString(err.Error()) {
				return true
			}
		}

		return false
	}

	for i, tc := range cases {
		_, errs := tc.f(tc.val, "test_property")

		if len(errs) == 0 && tc.expectedErr == nil {
			continue
		}

		if len(errs) != 0 && tc.expectedErr == nil {
			t.Fatalf("expected test case %d to produce no errors, got %v", i, errs)
		}

		if !matchErr(errs, tc.expectedErr) {
			t.Fatalf("expected test case %d to produce error matching \"%s\", got %v", i, tc.expectedErr, errs)
		}
	}
}
