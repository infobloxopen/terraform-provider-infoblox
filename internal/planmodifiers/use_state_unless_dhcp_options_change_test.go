package planmodifiers

import (
	"context"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
)

// Minimal UDDI-shaped schema: uddi.dhcp_options and uddi.inheritance_sources.dhcp_options.
var testOptionItemType = tftypes.Object{AttributeTypes: map[string]tftypes.Type{
	"option_code": tftypes.String,
}}

var testInheritedOptionType = tftypes.Object{AttributeTypes: map[string]tftypes.Type{
	"action": tftypes.String,
	"value": tftypes.Object{AttributeTypes: map[string]tftypes.Type{
		"option": testOptionItemType,
	}},
}}

var testInheritedListType = tftypes.Object{AttributeTypes: map[string]tftypes.Type{
	"action": tftypes.String,
	"value":  tftypes.List{ElementType: testInheritedOptionType},
}}

var testInheritanceType = tftypes.Object{AttributeTypes: map[string]tftypes.Type{
	"dhcp_options": testInheritedListType,
}}

var testUddiType = tftypes.Object{AttributeTypes: map[string]tftypes.Type{
	"dhcp_options":        tftypes.List{ElementType: testOptionItemType},
	"inheritance_sources": testInheritanceType,
}}

var testRootType = tftypes.Object{AttributeTypes: map[string]tftypes.Type{
	"uddi": testUddiType,
}}

func testSchema() schema.Schema {
	optionItem := map[string]schema.Attribute{
		"option_code": schema.StringAttribute{Optional: true},
	}
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"uddi": schema.SingleNestedAttribute{
				Optional: true,
				Attributes: map[string]schema.Attribute{
					"dhcp_options": schema.ListNestedAttribute{
						Optional:     true,
						NestedObject: schema.NestedAttributeObject{Attributes: optionItem},
					},
					"inheritance_sources": schema.SingleNestedAttribute{
						Optional: true,
						Computed: true,
						Attributes: map[string]schema.Attribute{
							"dhcp_options": schema.SingleNestedAttribute{
								Optional: true,
								Computed: true,
								Attributes: map[string]schema.Attribute{
									"action": schema.StringAttribute{Optional: true, Computed: true},
									"value": schema.ListNestedAttribute{
										Optional: true,
										Computed: true,
										NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{
											"action": schema.StringAttribute{Optional: true, Computed: true},
											"value": schema.SingleNestedAttribute{
												Computed: true,
												Attributes: map[string]schema.Attribute{
													"option": schema.SingleNestedAttribute{
														Computed:   true,
														Attributes: optionItem,
													},
												},
											},
										}},
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func optionItem(code string) tftypes.Value {
	return tftypes.NewValue(testOptionItemType, map[string]tftypes.Value{
		"option_code": tftypes.NewValue(tftypes.String, code),
	})
}

func optionList(codes ...string) tftypes.Value {
	if codes == nil {
		return tftypes.NewValue(tftypes.List{ElementType: testOptionItemType}, nil)
	}
	vals := make([]tftypes.Value, 0, len(codes))
	for _, c := range codes {
		vals = append(vals, optionItem(c))
	}
	return tftypes.NewValue(tftypes.List{ElementType: testOptionItemType}, vals)
}

// inheritance builds inheritance_sources with one entry per code=>action pair.
func inheritance(entries ...[2]string) tftypes.Value {
	vals := make([]tftypes.Value, 0, len(entries))
	for _, e := range entries {
		vals = append(vals, tftypes.NewValue(testInheritedOptionType, map[string]tftypes.Value{
			"action": tftypes.NewValue(tftypes.String, e[1]),
			"value": tftypes.NewValue(testInheritedOptionType.AttributeTypes["value"], map[string]tftypes.Value{
				"option": optionItem(e[0]),
			}),
		}))
	}
	return tftypes.NewValue(testInheritanceType, map[string]tftypes.Value{
		"dhcp_options": tftypes.NewValue(testInheritedListType, map[string]tftypes.Value{
			"action": tftypes.NewValue(tftypes.String, "inherit"),
			"value":  tftypes.NewValue(tftypes.List{ElementType: testInheritedOptionType}, vals),
		}),
	})
}

func root(opts, inh tftypes.Value) tftypes.Value {
	return tftypes.NewValue(testRootType, map[string]tftypes.Value{
		"uddi": tftypes.NewValue(testUddiType, map[string]tftypes.Value{
			"dhcp_options":        opts,
			"inheritance_sources": inh,
		}),
	})
}

func runModifier(t *testing.T, stateOpts, configOpts tftypes.Value, stateInh tftypes.Value) (types.Object, types.Object) {
	t.Helper()
	ctx := context.Background()
	s := testSchema()
	unknownInh := tftypes.NewValue(testInheritanceType, tftypes.UnknownValue)
	nullInh := tftypes.NewValue(testInheritanceType, nil)

	state := tfsdk.State{Schema: s, Raw: root(stateOpts, stateInh)}
	plan := tfsdk.Plan{Schema: s, Raw: root(configOpts, unknownInh)}
	config := tfsdk.Config{Schema: s, Raw: root(configOpts, nullInh)}

	p := path.Root("uddi").AtName("inheritance_sources")
	var stateVal, planVal, configVal types.Object
	if d := state.GetAttribute(ctx, p, &stateVal); d.HasError() {
		t.Fatalf("state: %v", d)
	}
	if d := plan.GetAttribute(ctx, p, &planVal); d.HasError() {
		t.Fatalf("plan: %v", d)
	}
	if d := config.GetAttribute(ctx, p, &configVal); d.HasError() {
		t.Fatalf("config: %v", d)
	}

	req := planmodifier.ObjectRequest{
		Path:        p,
		Config:      config,
		ConfigValue: configVal,
		Plan:        plan,
		PlanValue:   planVal,
		State:       state,
		StateValue:  stateVal,
	}
	resp := &planmodifier.ObjectResponse{PlanValue: req.PlanValue}
	UseStateForUnknownUnlessDhcpOptionsChange().PlanModifyObject(ctx, req, resp)
	if resp.Diagnostics.HasError() {
		t.Fatalf("diagnostics: %v", resp.Diagnostics)
	}
	return resp.PlanValue, stateVal
}

func firstEntryAction(t *testing.T, v types.Object) string {
	t.Helper()
	tv, err := v.ToTerraformValue(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	got, _, err := tftypes.WalkAttributePath(tv, tftypes.NewAttributePath().
		WithAttributeName("dhcp_options").WithAttributeName("value").WithElementKeyInt(0).WithAttributeName("action"))
	if err != nil {
		t.Fatal(err)
	}
	var s string
	if err := got.(tftypes.Value).As(&s); err != nil {
		t.Fatal(err)
	}
	return s
}

func TestUseStateForUnknownUnlessDhcpOptionsChange_NoChangeUsesState(t *testing.T) {
	plan, state := runModifier(t, optionList("code-a"), optionList("code-a"), inheritance([2]string{"code-b", "inherit"}))
	if !plan.Equal(state) {
		t.Fatalf("expected state value, got %s", plan)
	}
}

func TestUseStateForUnknownUnlessDhcpOptionsChange_InheritToOverride(t *testing.T) {
	plan, _ := runModifier(t, optionList(), optionList("code-b"), inheritance([2]string{"code-b", "inherit"}))
	if plan.IsUnknown() {
		t.Fatal("expected a known plan value")
	}
	if got := firstEntryAction(t, plan); got != "override" {
		t.Fatalf("action = %q, want override", got)
	}
}

func TestUseStateForUnknownUnlessDhcpOptionsChange_UnrelatedOptionKeepsInherit(t *testing.T) {
	plan, _ := runModifier(t, optionList(), optionList("code-a"), inheritance([2]string{"code-b", "inherit"}))
	if plan.IsUnknown() {
		t.Fatal("expected a known plan value")
	}
	if got := firstEntryAction(t, plan); got != "inherit" {
		t.Fatalf("action = %q, want inherit", got)
	}
}

func TestUseStateForUnknownUnlessDhcpOptionsChange_RemovedOverrideStaysUnknown(t *testing.T) {
	plan, _ := runModifier(t, optionList("code-b"), optionList(), inheritance([2]string{"code-b", "override"}))
	if !plan.IsUnknown() {
		t.Fatalf("expected unknown, got %s", plan)
	}
}
