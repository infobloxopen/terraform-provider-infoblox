package dtc

import (
	"context"
	"fmt"

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
}

func validateDtcTopologyUDDIConfig(ctx context.Context, m *UDDIDtcTopologyModel, resp *resource.ValidateConfigResponse) {
}

// populateDtcTopologyNIOSRules resolves each rule ref in the topology response to its full details.
// The topology API returns rules as bare refs; this fetches each rule individually so Flatten
// has dest_type, destination_link, return_type, sources, and valid to work with.
func (r *DtcTopologyResource) populateDtcTopologyNIOSRules(ctx context.Context, resp *coremodel.DtcTopology, diags *diag.Diagnostics) {
	if r.niosClient == nil || resp == nil || resp.NIOS == nil {
		return
	}
	for i, rule := range resp.NIOS.Rules {
		if rule.DtcTopologyRulesInnerOneOf == nil || rule.DtcTopologyRulesInnerOneOf.Ref == nil {
			continue
		}
		resp.NIOS.Rules[i].DtcTopologyRulesInnerOneOf1 = fetchDtcTopologyNIOSRuleDetails(ctx, r.niosClient, *rule.DtcTopologyRulesInnerOneOf.Ref, diags)
		if diags.HasError() {
			return
		}
	}
}

func fetchDtcTopologyNIOSRuleDetails(ctx context.Context, client *niosclient.APIClient, ruleRef string, diags *diag.Diagnostics) *niosdtc.DtcTopologyRulesInnerOneOf1 {
	apiRes, _, err := client.DTCAPI.
		DtcTopologyRuleAPI.
		Read(ctx, core.ExtractNIOSRef(ruleRef)).
		ReturnFieldsPlus(dtcTopologyRuleReturnFields).
		ReturnAsObject(1).
		Execute()
	if err != nil {
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
