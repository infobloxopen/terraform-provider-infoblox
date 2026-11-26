package dhcp

import (
	"context"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	uddiipam "github.com/infobloxopen/universal-ddi-go-client/ipam"

	coremodel "github.com/infobloxopen/terraform-provider-infoblox/internal/core/model/dhcp"
	internaltypes "github.com/infobloxopen/terraform-provider-infoblox/internal/types"
)

func testOptionItems(codes ...string) []uddiipam.OptionItem {
	out := make([]uddiipam.OptionItem, 0, len(codes))
	for _, c := range codes {
		out = append(out, uddiipam.OptionItem{
			OptionCode:  strPtr(c),
			OptionValue: strPtr("v-" + c),
			Type:        strPtr("option"),
		})
	}
	return out
}

func assertUnorderedDhcpOptions(t *testing.T, schemaType attr.Type, attrType attr.Type, flatten func(codes ...string) internaltypes.UnorderedListValue) {
	t.Helper()
	ctx := context.Background()

	if !schemaType.Equal(attrType) {
		t.Fatalf("schema type %s does not match attr type %s", schemaType, attrType)
	}

	a := flatten("dhcp/option_code/1", "dhcp/option_code/2")
	b := flatten("dhcp/option_code/2", "dhcp/option_code/1")
	eq, d := a.ListSemanticEquals(ctx, b)
	if d.HasError() {
		t.Fatalf("semantic equals: %v", d)
	}
	if !eq {
		t.Fatal("reordered dhcp_options should be semantically equal")
	}

	c := flatten("dhcp/option_code/1", "dhcp/option_code/3")
	eq, _ = a.ListSemanticEquals(ctx, c)
	if eq {
		t.Fatal("different dhcp_options should not be semantically equal")
	}
}

func TestUDDIRangeDhcpOptionsUnordered(t *testing.T) {
	assertUnorderedDhcpOptions(t,
		RangeResourceUddiSchemaAttributes["dhcp_options"].GetType(),
		UDDIRangeAttrTypes["dhcp_options"],
		func(codes ...string) internaltypes.UnorderedListValue {
			var diags diag.Diagnostics
			m := &UDDIRangeModel{}
			m.Flatten(context.Background(), &coremodel.UDDIRangeExt{DhcpOptions: testOptionItems(codes...)}, &diags)
			if diags.HasError() {
				t.Fatalf("flatten: %v", diags)
			}
			if _, d := types.ObjectValueFrom(context.Background(), UDDIRangeAttrTypes, m); d.HasError() {
				t.Fatalf("object from model: %v", d)
			}
			return m.DhcpOptions
		})
}

func TestUDDIIpv6rangeDhcpOptionsUnordered(t *testing.T) {
	assertUnorderedDhcpOptions(t,
		Ipv6rangeResourceUddiSchemaAttributes["dhcp_options"].GetType(),
		UDDIIpv6rangeAttrTypes["dhcp_options"],
		func(codes ...string) internaltypes.UnorderedListValue {
			var diags diag.Diagnostics
			m := &UDDIIpv6rangeModel{}
			m.Flatten(context.Background(), &coremodel.UDDIIpv6rangeExt{DhcpOptions: testOptionItems(codes...)}, &diags)
			if diags.HasError() {
				t.Fatalf("flatten: %v", diags)
			}
			if _, d := types.ObjectValueFrom(context.Background(), UDDIIpv6rangeAttrTypes, m); d.HasError() {
				t.Fatalf("object from model: %v", d)
			}
			return m.DhcpOptions
		})
}
