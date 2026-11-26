package ipam

import (
	"context"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	uddiipam "github.com/infobloxopen/universal-ddi-go-client/ipam"
)

func strPtr(s string) *string { return &s }

func TestInheritedDHCPOptionRoundTrip(t *testing.T) {
	ctx := context.Background()
	var diags diag.Diagnostics

	from := &uddiipam.InheritedDHCPOptionList{
		Action: strPtr("inherit"),
		Value: []uddiipam.InheritedDHCPOption{
			{
				Action:      strPtr("inherit"),
				DisplayName: strPtr("parent-block"),
				Source:      strPtr("ipam/address_block/abc"),
				Value: &uddiipam.InheritedDHCPOptionItem{
					Option: &uddiipam.OptionItem{
						OptionCode:  strPtr("dhcp/option_code/123"),
						OptionValue: strPtr("10.0.0.1"),
						Type:        strPtr("option"),
					},
					OverridingGroup: strPtr("dhcp/option_group/xyz"),
				},
			},
		},
	}

	obj := FlattenInheritedDHCPOptionList(ctx, from, &diags)
	if diags.HasError() {
		t.Fatalf("flatten: %v", diags)
	}

	got := ExpandInheritedDHCPOptionList(ctx, obj, &diags)
	if diags.HasError() {
		t.Fatalf("expand: %v", diags)
	}
	if got == nil || len(got.Value) != 1 {
		t.Fatalf("expected one inherited option, got %+v", got)
	}
	opt := got.Value[0]
	if opt.GetAction() != "inherit" {
		t.Errorf("action = %q", opt.GetAction())
	}
	if opt.Value == nil || opt.Value.Option == nil {
		t.Fatalf("expected value to be sent on expand, got %+v", opt.Value)
	}
	if opt.Value.Option.GetOptionCode() != "dhcp/option_code/123" ||
		opt.Value.Option.GetOptionValue() != "10.0.0.1" ||
		opt.Value.Option.GetType() != "option" {
		t.Errorf("unexpected option item %+v", opt.Value.Option)
	}
	if opt.Value.GetOverridingGroup() != "dhcp/option_group/xyz" {
		t.Errorf("overriding_group = %q", opt.Value.GetOverridingGroup())
	}
	// display_name and source are read-only and must not be sent back.
	if opt.DisplayName != nil || opt.Source != nil {
		t.Errorf("read-only fields should not be expanded: %+v", opt)
	}
}

func TestInheritedDHCPOptionFlattenReadOnlyFields(t *testing.T) {
	ctx := context.Background()
	var diags diag.Diagnostics

	obj := FlattenInheritedDHCPOption(ctx, &uddiipam.InheritedDHCPOption{
		Action:      strPtr("block"),
		DisplayName: strPtr("parent"),
		Source:      strPtr("ipam/ip_space/1"),
	}, &diags)
	if diags.HasError() {
		t.Fatalf("flatten: %v", diags)
	}
	attrs := obj.Attributes()
	if v := attrs["display_name"].(types.String).ValueString(); v != "parent" {
		t.Errorf("display_name = %q", v)
	}
	if v := attrs["source"].(types.String).ValueString(); v != "ipam/ip_space/1" {
		t.Errorf("source = %q", v)
	}
	if !attrs["value"].(types.Object).IsNull() {
		t.Errorf("value should be null when absent")
	}
}
