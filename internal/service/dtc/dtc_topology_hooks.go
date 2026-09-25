package dtc

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	niosclient "github.com/infobloxopen/infoblox-nios-go-client/client"
	niosdtc "github.com/infobloxopen/infoblox-nios-go-client/dtc"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/core"
	coremodel "github.com/infobloxopen/terraform-provider-infoblox/internal/core/model/dtc"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
)

const dtcTopologyRuleReturnFields = "dest_type,destination_link,return_type,sources,valid"

// ValidateDtcTopology validates the DtcTopology configuration.
func ValidateDtcTopology(ctx context.Context, data DtcTopologyModel, resp *resource.ValidateConfigResponse) {
	if nios := flex.ExpandNestedObject[NIOSDtcTopologyModel](ctx, data.NIOS, &resp.Diagnostics); nios != nil {
		validateDtcTopologyNIOSConfig(ctx, nios, resp)
	}
	if uddi := flex.ExpandNestedObject[UDDIDtcTopologyModel](ctx, data.UDDI, &resp.Diagnostics); uddi != nil {
		validateDtcTopologyUDDIConfig(ctx, uddi, resp)
	}
}

func validateDtcTopologyNIOSConfig(ctx context.Context, m *NIOSDtcTopologyModel, resp *resource.ValidateConfigResponse) {
	if m.Rules.IsNull() || m.Rules.IsUnknown() {
		return
	}
	var rules []TopologyRulesInnerModel
	resp.Diagnostics.Append(m.Rules.ElementsAs(ctx, &rules, false)...)
	if resp.Diagnostics.HasError() {
		return
	}
	firstDestType := ""
	for _, rule := range rules {
		if rule.DestType.IsUnknown() {
			continue
		}
		destType := rule.DestType.ValueString()
		if firstDestType == "" {
			firstDestType = destType
		} else if firstDestType != destType {
			resp.Diagnostics.AddError(
				"Invalid Configuration",
				fmt.Sprintf("All topology rules must have the same 'dest_type'. Found %q and %q.", firstDestType, destType),
			)
			return
		}
	}
}

func validateDtcTopologyUDDIConfig(ctx context.Context, m *UDDIDtcTopologyModel, resp *resource.ValidateConfigResponse) {
	if m.Sources.IsNull() || m.Sources.IsUnknown() {
		return
	}
	var sources []TopologySourceModel
	resp.Diagnostics.Append(m.Sources.ElementsAs(ctx, &sources, false)...)
	if resp.Diagnostics.HasError() {
		return
	}
	for i, src := range sources {
		if src.Source.IsUnknown() {
			continue
		}
		sourceVal := src.Source.ValueString()
		hasSubnets := !src.Subnets.IsNull() && !src.Subnets.IsUnknown()
		hasTagRules := !src.TagRules.IsNull() && !src.TagRules.IsUnknown()
		if sourceVal == "tag_rule" && hasSubnets {
			resp.Diagnostics.AddError(
				"Invalid Configuration",
				fmt.Sprintf("sources[%d]: 'subnets' must not be set when 'source' is \"tag_rule\".", i),
			)
		}
		if sourceVal == "subnet" && hasTagRules {
			resp.Diagnostics.AddError(
				"Invalid Configuration",
				fmt.Sprintf("sources[%d]: 'tag_rules' must not be set when 'source' is \"subnet\".", i),
			)
		}
	}
}

// populateDtcTopologyNIOSRules resolves each rule ref in the topology response to its full details.
// The topology API returns rules as bare refs; this fetches each rule individually so Flatten
// has dest_type, destination_link, return_type, sources, and valid to work with.
func populateDtcTopologyNIOSRules(ctx context.Context, client *niosclient.APIClient, resp *coremodel.DtcTopology, diags *diag.Diagnostics) {
	if client == nil || resp == nil || resp.NIOS == nil {
		return
	}
	filtered := make([]niosdtc.DtcTopologyRulesInner, 0, len(resp.NIOS.Rules))
	for _, rule := range resp.NIOS.Rules {
		if rule.DtcTopologyRulesInnerOneOf == nil || rule.DtcTopologyRulesInnerOneOf.Ref == nil {
			filtered = append(filtered, rule)
			continue
		}
		details := fetchDtcTopologyNIOSRuleDetails(ctx, client, *rule.DtcTopologyRulesInnerOneOf.Ref, diags)
		if diags.HasError() {
			return
		}
		if details == nil {
			// 404 — rule child object is stale; omit from results
			continue
		}
		rule.DtcTopologyRulesInnerOneOf1 = details
		filtered = append(filtered, rule)
	}
	resp.NIOS.Rules = filtered
}

func fetchDtcTopologyNIOSRuleDetails(ctx context.Context, client *niosclient.APIClient, ruleRef string, diags *diag.Diagnostics) *niosdtc.DtcTopologyRulesInnerOneOf1 {
	apiRes, httpResp, err := client.DTCAPI.
		DtcTopologyRuleAPI.
		Read(ctx, core.ExtractNIOSRef(ruleRef)).
		ReturnFieldsPlus(dtcTopologyRuleReturnFields).
		ReturnAsObject(1).
		Execute()
	if err != nil {
		if httpResp != nil && httpResp.StatusCode == http.StatusNotFound {
			return nil
		}
		diags.AddError("Client Error", fmt.Sprintf("Unable to read DTC Topology Rule %s: %s", ruleRef, err))
		return nil
	}

	res := apiRes.GetDtcTopologyRuleResponseObjectAsResult.GetResult()
	ruleData := &niosdtc.DtcTopologyRulesInnerOneOf1{}

	if destType, ok := res.GetDestTypeOk(); ok {
		ruleData.SetDestType(*destType)
	}
	if destLink, ok := res.GetDestinationLinkOk(); ok {
		if destLink.DtcTopologyRuleDestinationLinkOneOf != nil && destLink.DtcTopologyRuleDestinationLinkOneOf.Ref != nil {
			ruleData.SetDestinationLink(*destLink.DtcTopologyRuleDestinationLinkOneOf.Ref)
		}
	}
	if returnType, ok := res.GetReturnTypeOk(); ok {
		ruleData.SetReturnType(*returnType)
	}
	if valid, ok := res.GetValidOk(); ok {
		ruleData.SetValid(*valid)
	}
	if sources, ok := res.GetSourcesOk(); ok {
		convertedSources := make([]niosdtc.DtcTopologyRulesInnerOneOf1SourcesInner, len(sources))
		for i, source := range sources {
			innerSource := niosdtc.DtcTopologyRulesInnerOneOf1SourcesInner{}
			if v, ok := source.GetSourceOpOk(); ok {
				innerSource.SetSourceOp(*v)
			}
			if v, ok := source.GetSourceTypeOk(); ok {
				innerSource.SetSourceType(*v)
			}
			if v, ok := source.GetSourceValueOk(); ok {
				innerSource.SetSourceValue(*v)
			}
			convertedSources[i] = innerSource
		}
		ruleData.SetSources(convertedSources)
	}
	return ruleData
}
