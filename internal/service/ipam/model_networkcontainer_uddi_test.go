package ipam

import (
	"context"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
)

// When address comes from dynamic_allocation it is computed; without
// UseStateForUnknown every update planned it as unknown.
func TestUDDINetworkcontainerAddressUsesStateForUnknown(t *testing.T) {
	ctx := context.Background()
	want := stringplanmodifier.UseStateForUnknown().Description(ctx)

	a, ok := NetworkcontainerResourceUddiSchemaAttributes["address"].(schema.StringAttribute)
	if !ok {
		t.Fatal("uddi address is not a StringAttribute")
	}
	for _, m := range a.PlanModifiers {
		if m.Description(ctx) == want {
			return
		}
	}
	t.Fatal("uddi address is missing UseStateForUnknown")
}
