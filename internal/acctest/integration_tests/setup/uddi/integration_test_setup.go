// Objects read by this setup program (not created, IDs stored as env vars):
//
// Infra Hosts (up to 2 online hosts):
//   - UDDI_INFRA_HOST_DISPLAY_NAME_1, UDDI_INFRA_HOST_LEGACY_ID_1
//   - UDDI_INFRA_HOST_DISPLAY_NAME_2, UDDI_INFRA_HOST_LEGACY_ID_2
//
// DNS Services (the Service named "DNS <display_name>" for each infra host above,
// tagged with a random "location" value):
//   - UDDI_DNS_SERVICE_TAG_KEY_1, UDDI_DNS_SERVICE_TAG_VALUE_1
//   - UDDI_DNS_SERVICE_TAG_KEY_2, UDDI_DNS_SERVICE_TAG_VALUE_2
//
// DNS Hosts (up to 2):
//   - UDDI_DNS_HOST_ID_1, UDDI_DNS_HOST_ID_2
//
// DHCP Hosts (up to 2):
//   - UDDI_DHCP_HOST_ID_1, UDDI_DHCP_HOST_ID_2
//
// Objects created by this setup program (IDs stored as env vars):
//
// Anycast Service (on first online host):
//   - tf_anycast_service_1 (UDDI_ANYCAST_SERVICE_ID_1)
//
// DHCP Option Groups:
//   - tf_option_group_1 (UDDI_OPTION_GROUP_1_ID)
//   - tf_option_group_2 (UDDI_OPTION_GROUP_2_ID)
//
// DHCP Option Space:
//   - tf_option_space_1 (UDDI_OPTION_SPACE_1_ID)
//
// DHCP Option Code:
//   - tf_option_code_1 (UDDI_OPTION_CODE_1_ID)
//
// DNS Auth Zone:
//   - example_zone_250 (UDDI_AUTH_ZONE_1_ID)
//
// IPAM IP Spaces:
//   - tf_ip_space_1 (UDDI_IP_SPACE_ID_1)
//   - tf_ip_space_2 (UDDI_IP_SPACE_ID_2)
//
// DNS Auth NSGs:
//   - tf_test_auth_nsg_1 (UDDI_AUTH_NSG_ID_1)
//   - tf_test_auth_nsg_2 (UDDI_AUTH_NSG_ID_2)
//
// DNS Views:
//   - tf_test_view_1 (UDDI_VIEW_ID_1)
//
// DTC Policies:
//   - tf_dtc_policy_1 (UDDI_DTC_POLICY_ID_1)
//   - tf_dtc_policy_2 (UDDI_DTC_POLICY_ID_2)
//
// DTC SNMP User Security Models:
//   - tf_snmp_usm_1 (UDDI_SNMP_USER_SECURITY_MODEL_ID_1)
//   - tf_snmp_usm_2 (UDDI_SNMP_USER_SECURITY_MODEL_ID_2)

package main

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	uddiclient "github.com/infobloxopen/universal-ddi-go-client/client"
	"github.com/infobloxopen/universal-ddi-go-client/dnsconfig"
	"github.com/infobloxopen/universal-ddi-go-client/dtc"
	"github.com/infobloxopen/universal-ddi-go-client/inframgmt"
	"github.com/infobloxopen/universal-ddi-go-client/ipam"
	uddioption "github.com/infobloxopen/universal-ddi-go-client/option"

	"github.com/infobloxopen/terraform-provider-infoblox/internal/acctest"
)

// dnsServiceTagKey is the tag key set on the DNS Service associated with each infra host.
const dnsServiceTagKey = "location"

var pipelineEnvFile *os.File

func writePipelineEnvVar(key, value string) error {
	if pipelineEnvFile == nil {
		return fmt.Errorf("pipeline_uddi.env file is not initialized")
	}
	if _, err := fmt.Fprintf(pipelineEnvFile, "%s=%s\n", key, value); err != nil {
		return fmt.Errorf("write env var %s to pipeline_uddi.env: %w", key, err)
	}
	return nil
}

// StoreInfraHostDetails fetches up to two online Infra Hosts via the Detail API
// and stores each host's display_name and legacy_id into pipeline_uddi.env as
// UDDI_INFRA_HOST_DISPLAY_NAME_1/2 and UDDI_INFRA_HOST_LEGACY_ID_1/2.
// For each host, it then looks up the DNS Service named "DNS <display_name>" and,
// if found, tags it with a random "location" value, storing that tag into
// pipeline_uddi.env as UDDI_DNS_SERVICE_TAG_KEY_1/2 and UDDI_DNS_SERVICE_TAG_VALUE_1/2.
// It returns the fetched hosts so callers can use them without a second API call.
func StoreInfraHostDetails(ctx context.Context, client *uddiclient.APIClient) ([]inframgmt.DetailHost, error) {
	resp, _, err := client.InfraManagementAPI.DetailAPI.HostsList(ctx).
		Filter("composite_status=='online'").
		Limit(2).
		Execute()
	if err != nil {
		return nil, fmt.Errorf("store infra host details: list detail hosts: %w", err)
	}

	if resp == nil || len(resp.GetResults()) == 0 {
		fmt.Println("No online infra hosts found, skipping infra host detail storage")
		return nil, nil
	}

	hosts := resp.GetResults()
	tagN := 1
	for i, host := range hosts {
		if i >= 2 {
			break
		}
		n := i + 1
		displayNameVar := fmt.Sprintf("UDDI_INFRA_HOST_DISPLAY_NAME_%d", n)
		legacyIDVar := fmt.Sprintf("UDDI_INFRA_HOST_LEGACY_ID_%d", n)

		displayName := host.GetDisplayName()
		if displayName != "" {
			if err := writePipelineEnvVar(displayNameVar, displayName); err != nil {
				return nil, fmt.Errorf("store infra host details: write %s: %w", displayNameVar, err)
			}
			fmt.Printf("Stored infra host display_name %q as %s\n", displayName, displayNameVar)
		}

		if v := host.GetLegacyId(); v != "" {
			if err := writePipelineEnvVar(legacyIDVar, v); err != nil {
				return nil, fmt.Errorf("store infra host details: write %s: %w", legacyIDVar, err)
			}
			fmt.Printf("Stored infra host legacy_id %q as %s\n", v, legacyIDVar)
		}

		if displayName == "" {
			continue
		}

		tagged, err := tagDNSServiceForHost(ctx, client, displayName, tagN)
		if err != nil {
			return nil, err
		}
		if tagged {
			tagN++
		}
	}

	return hosts, nil
}

// tagDNSServiceForHost looks up the DNS Service named "DNS <displayName>" (the Service
// associated with the infra host identified by displayName) and, if one exists, sets a
// random "location" tag on it via the Services API. The tag key/value are stored into
// pipeline_uddi.env as UDDI_DNS_SERVICE_TAG_KEY_<n> and UDDI_DNS_SERVICE_TAG_VALUE_<n>.
// It returns false (with no error) if no matching DNS service is found.
func tagDNSServiceForHost(ctx context.Context, client *uddiclient.APIClient, displayName string, n int) (bool, error) {
	serviceName := "DNS_" + displayName

	listResp, _, err := client.InfraManagementAPI.ServicesAPI.List(ctx).
		Filter(fmt.Sprintf("name=='%s'", serviceName)).
		Execute()
	if err != nil {
		return false, fmt.Errorf("tag DNS service for host %q: list services: %w", displayName, err)
	}

	if listResp == nil || len(listResp.GetResults()) == 0 {
		fmt.Printf("No DNS service named %q found, skipping tag update\n", serviceName)
		return false, nil
	}

	svc := listResp.GetResults()[0]
	if svc.Id == nil || *svc.Id == "" {
		return false, fmt.Errorf("tag DNS service for host %q: service %q has no ID", displayName, serviceName)
	}

	tagValue := acctest.RandomName()
	tags := make(map[string]interface{}, len(svc.Tags)+1)
	for k, v := range svc.Tags {
		tags[k] = v
	}
	tags[dnsServiceTagKey] = tagValue
	svc.Tags = tags

	if _, _, err := client.InfraManagementAPI.ServicesAPI.Update(ctx, *svc.Id).Body(svc).Execute(); err != nil {
		return false, fmt.Errorf("tag DNS service for host %q: update service %q: %w", displayName, serviceName, err)
	}

	tagKeyVar := fmt.Sprintf("UDDI_DNS_SERVICE_TAG_KEY_%d", n)
	tagValueVar := fmt.Sprintf("UDDI_DNS_SERVICE_TAG_VALUE_%d", n)

	if err := writePipelineEnvVar(tagKeyVar, dnsServiceTagKey); err != nil {
		return false, fmt.Errorf("tag DNS service for host %q: write %s: %w", displayName, tagKeyVar, err)
	}
	fmt.Printf("Stored DNS service tag key %q as %s\n", dnsServiceTagKey, tagKeyVar)

	if err := writePipelineEnvVar(tagValueVar, tagValue); err != nil {
		return false, fmt.Errorf("tag DNS service for host %q: write %s: %w", displayName, tagValueVar, err)
	}
	fmt.Printf("Stored DNS service tag value %q as %s (service %q)\n", tagValue, tagValueVar, serviceName)

	return true, nil
}

// CreateAnyCastService creates an anycast Service on the given host and stores its ID
// into pipeline_uddi.env as UDDI_ANYCAST_SERVICE_ID_1. If an anycast service already
// exists on the host the call is skipped gracefully.
func CreateAnyCastService(ctx context.Context, client *uddiclient.APIClient, host inframgmt.DetailHost) error {
	const serviceName = "tf_anycast_service_1"

	// Skip if anycast is already deployed on this host.
	for _, svc := range host.GetServices() {
		if svc.GetServiceType() == "anycast" {
			fmt.Printf("Anycast service already exists on host %q, skipping creation\n", host.GetDisplayName())
			return nil
		}
	}

	pool := host.GetPool()
	poolId := pool.GetPoolId()
	if poolId == "" {
		return fmt.Errorf("create anycast service: host %q has no pool_id", host.GetDisplayName())
	}

	body := inframgmt.Service{
		Name:         serviceName,
		ServiceType:  "anycast",
		DesiredState: inframgmt.PtrString("start"),
		PoolId:       poolId,
	}

	resp, _, err := client.InfraManagementAPI.ServicesAPI.Create(ctx).Body(body).Execute()
	if err != nil {
		if strings.Contains(err.Error(), "Cannot have duplicate service") {
			fmt.Printf("Anycast service %q already exists (duplicate error), skipping\n", serviceName)
			return nil
		}
		return fmt.Errorf("create anycast service: %w", err)
	}

	if resp == nil || resp.Result == nil || resp.Result.Id == nil {
		return fmt.Errorf("create anycast service: create response missing ID")
	}

	createdID := *resp.Result.Id
	if err := writePipelineEnvVar("UDDI_ANYCAST_SERVICE_ID_1", createdID); err != nil {
		return fmt.Errorf("create anycast service: write UDDI_ANYCAST_SERVICE_ID_1: %w", err)
	}

	fmt.Printf("Anycast service %q created successfully (ID: %q, env: UDDI_ANYCAST_SERVICE_ID_1)\n", serviceName, createdID)
	return nil
}

// StoreDNSHostIDs lists DNS Host objects and stores the IDs of the first two
// into pipeline_uddi.env as UDDI_DNS_HOST_ID_1 and UDDI_DNS_HOST_ID_2.
func StoreDNSHostIDs(ctx context.Context, client *uddiclient.APIClient) error {
	resp, _, err := client.DNSConfigurationAPI.HostAPI.List(ctx).Tfilter(`"used_for"=="Terraform Provider Acceptance Tests"`).Execute()
	if err != nil {
		return fmt.Errorf("store DNS host IDs: list DNS hosts: %w", err)
	}

	if resp == nil || resp.Results == nil || len(resp.Results) == 0 {
		fmt.Println("No DNS hosts found, skipping DNS host ID storage")
		return nil
	}

	envVars := []string{"UDDI_DNS_HOST_ID_1", "UDDI_DNS_HOST_ID_2"}
	stored := 0

	for i, host := range resp.Results {
		if i >= 2 {
			break
		}
		if host.Id == nil || *host.Id == "" {
			fmt.Printf("DNS host at index %d has no ID, skipping\n", i)
			continue
		}
		if err := writePipelineEnvVar(envVars[stored], *host.Id); err != nil {
			return fmt.Errorf("store DNS host IDs: write %s: %w", envVars[stored], err)
		}
		fmt.Printf("Stored DNS host ID %q as %s\n", *host.Id, envVars[stored])
		stored++
	}

	if stored == 0 {
		fmt.Println("No DNS host IDs could be stored (all hosts missing ID)")
	}

	return nil
}

// StoreDHCPHostIDs lists DHCP Host objects and stores the IDs of the first two
// into pipeline_uddi.env as UDDI_DHCP_HOST_ID_1 and UDDI_DHCP_HOST_ID_2.
func StoreDHCPHostIDs(ctx context.Context, client *uddiclient.APIClient) error {
	resp, _, err := client.IPAddressManagementAPI.DhcpHostAPI.List(ctx).Tfilter(`"used_for"=="Terraform Provider Acceptance Tests"`).Execute()
	if err != nil {
		return fmt.Errorf("store DHCP host IDs: list DHCP hosts: %w", err)
	}

	if resp == nil || resp.Results == nil || len(resp.Results) == 0 {
		fmt.Println("No DHCP hosts found, skipping DHCP host ID storage")
		return nil
	}

	envVars := []string{"UDDI_DHCP_HOST_ID_1", "UDDI_DHCP_HOST_ID_2"}
	stored := 0

	for i, host := range resp.Results {
		if stored >= 2 {
			break
		}
		if host.Id == nil || *host.Id == "" {
			fmt.Printf("DHCP host at index %d has no ID, skipping\n", i)
			continue
		}
		if val, ok := host.Tags["host/deployment_type"]; ok {
			if s, ok := val.(string); ok && s == "CNIOS" {
				fmt.Printf("DHCP host at index %d (%q) has deployment_type CNIOS, skipping\n", i, *host.Id)
				continue
			}
		}
		if err := writePipelineEnvVar(envVars[stored], *host.Id); err != nil {
			return fmt.Errorf("store DHCP host IDs: write %s: %w", envVars[stored], err)
		}
		fmt.Printf("Stored DHCP host ID %q as %s\n", *host.Id, envVars[stored])
		stored++
	}

	if stored == 0 {
		fmt.Println("No DHCP host IDs could be stored (all hosts missing ID)")
	}

	return nil
}

// CreateOptionGroups creates two DHCP option groups and stores their IDs into
// pipeline_uddi.env as UDDI_OPTION_GROUP_1_ID and UDDI_OPTION_GROUP_2_ID.
// If a group already exists, its existing ID is stored instead.
func CreateOptionGroups(ctx context.Context, client *uddiclient.APIClient) error {
	optionGroups := []struct {
		name     string
		protocol string
		idVar    string
	}{
		{name: "tf_option_group_1", protocol: "ip4", idVar: "UDDI_OPTION_GROUP_1_ID"},
		{name: "tf_option_group_2", protocol: "ip6", idVar: "UDDI_OPTION_GROUP_2_ID"},
	}

	for _, og := range optionGroups {
		body := ipam.OptionGroup{
			Name:     og.name,
			Protocol: dnsconfig.PtrString(og.protocol),
		}

		resp, _, err := client.IPAddressManagementAPI.OptionGroupAPI.Create(ctx).Body(body).Execute()
		if err != nil {
			if strings.Contains(err.Error(), "is already an existing") || strings.Contains(err.Error(), "conflict") {
				// Fetch the existing group's ID
				listResp, _, listErr := client.IPAddressManagementAPI.OptionGroupAPI.List(ctx).Filter("name==\"" + og.name + "\"").Execute()
				if listErr != nil {
					return fmt.Errorf("create option groups: list existing groups to find %q: %w", og.name, listErr)
				}

				var existingID string
				if listResp != nil {
					for _, existing := range listResp.Results {
						if existing.Name == og.name && existing.Id != nil {
							existingID = *existing.Id
							break
						}
					}
				}

				if existingID == "" {
					return fmt.Errorf("create option groups: option group %q already exists but ID could not be resolved", og.name)
				}

				if err := writePipelineEnvVar(og.idVar, existingID); err != nil {
					return fmt.Errorf("create option groups: write %s for existing group: %w", og.idVar, err)
				}

				fmt.Printf("Option group %q already exists, using existing ID %q (env: %s)\n", og.name, existingID, og.idVar)
				continue
			}
			return fmt.Errorf("create option groups: create %q: %w", og.name, err)
		}

		if resp == nil || resp.Result == nil || resp.Result.Id == nil {
			return fmt.Errorf("create option groups: create response for %q missing ID", og.name)
		}

		createdID := *resp.Result.Id
		if err := writePipelineEnvVar(og.idVar, createdID); err != nil {
			return fmt.Errorf("create option groups: write %s: %w", og.idVar, err)
		}

		fmt.Printf("Option group %q created successfully (ID: %q, env: %s)\n", og.name, createdID, og.idVar)
	}

	return nil
}

// CreateOptionCode creates a DHCP option space and an option code within it,
// storing their IDs into pipeline_uddi.env as UDDI_OPTION_SPACE_1_ID and UDDI_OPTION_CODE_1_ID.
// If either object already exists, its existing ID is stored instead.
func CreateOptionCode(ctx context.Context, client *uddiclient.APIClient) error {
	const (
		optionSpaceName = "tf_option_space_1"
		optionCodeName  = "tf_option_code_1"
		optionCodeCode  = int64(234)
		optionCodeType  = "boolean"
	)

	// Create or find the option space.
	optionSpaceID, err := createOrFindOptionSpace(ctx, client, optionSpaceName)
	if err != nil {
		return err
	}
	if err := writePipelineEnvVar("UDDI_OPTION_SPACE_1_ID", optionSpaceID); err != nil {
		return fmt.Errorf("create option code: write UDDI_OPTION_SPACE_1_ID: %w", err)
	}

	// Create or find the option code within that option space.
	body := ipam.OptionCode{
		Name:        optionCodeName,
		Code:        optionCodeCode,
		OptionSpace: optionSpaceID,
		Type:        optionCodeType,
	}

	resp, _, err := client.IPAddressManagementAPI.OptionCodeAPI.Create(ctx).Body(body).Execute()
	if err != nil {
		if strings.Contains(err.Error(), "is already an existing") || strings.Contains(err.Error(), "conflict") {
			listResp, _, listErr := client.IPAddressManagementAPI.OptionCodeAPI.List(ctx).Filter("name==\"" + optionCodeName + "\"").Execute()
			if listErr != nil {
				return fmt.Errorf("create option code: list existing option codes to find %q: %w", optionCodeName, listErr)
			}

			var existingID string
			if listResp != nil {
				for _, existing := range listResp.Results {
					if existing.Name == optionCodeName && existing.Id != nil {
						existingID = *existing.Id
						break
					}
				}
			}

			if existingID == "" {
				return fmt.Errorf("create option code: option code %q already exists but ID could not be resolved", optionCodeName)
			}

			if err := writePipelineEnvVar("UDDI_OPTION_CODE_1_ID", existingID); err != nil {
				return fmt.Errorf("create option code: write UDDI_OPTION_CODE_1_ID for existing code: %w", err)
			}

			fmt.Printf("Option code %q already exists, using existing ID %q (env: UDDI_OPTION_CODE_1_ID)\n", optionCodeName, existingID)
			return nil
		}
		return fmt.Errorf("create option code: create %q: %w", optionCodeName, err)
	}

	if resp == nil || resp.Result == nil || resp.Result.Id == nil {
		return fmt.Errorf("create option code: create response for %q missing ID", optionCodeName)
	}

	createdID := *resp.Result.Id
	if err := writePipelineEnvVar("UDDI_OPTION_CODE_1_ID", createdID); err != nil {
		return fmt.Errorf("create option code: write UDDI_OPTION_CODE_1_ID: %w", err)
	}

	fmt.Printf("Option code %q created successfully (ID: %q, env: UDDI_OPTION_CODE_1_ID)\n", optionCodeName, createdID)
	return nil
}

func createOrFindOptionSpace(ctx context.Context, client *uddiclient.APIClient, name string) (string, error) {
	body := ipam.OptionSpace{
		Name: name,
	}

	resp, _, err := client.IPAddressManagementAPI.OptionSpaceAPI.Create(ctx).Body(body).Execute()
	if err != nil {
		if strings.Contains(err.Error(), "is already an existing") || strings.Contains(err.Error(), "conflict") {
			listResp, _, listErr := client.IPAddressManagementAPI.OptionSpaceAPI.List(ctx).Filter("name==\"" + name + "\"").Execute()
			if listErr != nil {
				return "", fmt.Errorf("create or find option space: list existing spaces to find %q: %w", name, listErr)
			}

			if listResp != nil {
				for _, existing := range listResp.Results {
					if existing.Name == name && existing.Id != nil {
						fmt.Printf("Option space %q already exists, using existing ID %q\n", name, *existing.Id)
						return *existing.Id, nil
					}
				}
			}

			return "", fmt.Errorf("create or find option space: option space %q already exists but ID could not be resolved", name)
		}
		return "", fmt.Errorf("create or find option space: create %q: %w", name, err)
	}

	if resp == nil || resp.Result == nil || resp.Result.Id == nil {
		return "", fmt.Errorf("create or find option space: create response for %q missing ID", name)
	}

	fmt.Printf("Option space %q created successfully (ID: %q)\n", name, *resp.Result.Id)
	return *resp.Result.Id, nil
}

// readDefaultDNSViewID lists DNS views and returns the ID of the view named "default".
func readDefaultDNSViewID(ctx context.Context, client *uddiclient.APIClient) (string, error) {
	listResp, _, err := client.DNSConfigurationAPI.ViewAPI.List(ctx).Execute()
	if err != nil {
		return "", fmt.Errorf("read default DNS view: list views: %w", err)
	}

	if listResp != nil {
		for _, view := range listResp.Results {
			if view.Name == "default" && view.Id != nil {
				return *view.Id, nil
			}
		}
	}

	return "", fmt.Errorf("read default DNS view: no view named %q found", "default")
}

// CreateAuthZone creates a DNS auth zone and stores its ID into
// pipeline_uddi.env as UDDI_AUTH_ZONE_1_ID.
// If the zone already exists, its existing ID is stored instead.
func CreateAuthZone(ctx context.Context, client *uddiclient.APIClient) error {
	defaultViewID, err := readDefaultDNSViewID(ctx, client)
	if err != nil {
		return fmt.Errorf("create auth zone: %w", err)
	}
	fmt.Printf("Using default DNS view ID %q\n", defaultViewID)

	authZones := []struct {
		fqdn        string
		primaryType string
		idVar       string
	}{
		{fqdn: "example_zone_250.", primaryType: "cloud", idVar: "UDDI_AUTH_ZONE_ID_1"},
	}

	for _, az := range authZones {
		body := dnsconfig.AuthZone{
			Fqdn:        dnsconfig.PtrString(az.fqdn),
			PrimaryType: dnsconfig.PtrString(az.primaryType),
			View:        dnsconfig.PtrString(defaultViewID),
		}

		resp, _, err := client.DNSConfigurationAPI.AuthZoneAPI.Create(ctx).Body(body).Execute()
		if err != nil {
			if strings.Contains(err.Error(), "is already an existing") || strings.Contains(err.Error(), "conflict") || strings.Contains(err.Error(), "already exists") {
				// Fetch the existing zone's ID
				listResp, _, listErr := client.DNSConfigurationAPI.AuthZoneAPI.List(ctx).Filter("fqdn==\"" + az.fqdn + "\"").Execute()
				if listErr != nil {
					return fmt.Errorf("create auth zone: list existing zones to find %q: %w", az.fqdn, listErr)
				}

				var existingID string
				if listResp != nil {
					for _, existing := range listResp.Results {
						// UDDI stores fqdn with a trailing dot (e.g. "example_zone_250."), so compare with it trimmed.
						if existing.Fqdn != nil && strings.TrimSuffix(*existing.Fqdn, ".") == az.fqdn && existing.Id != nil {
							existingID = *existing.Id
							break
						}
					}
				}

				if existingID == "" {
					return fmt.Errorf("create auth zone: auth zone %q already exists but ID could not be resolved", az.fqdn)
				}

				if err := writePipelineEnvVar(az.idVar, existingID); err != nil {
					return fmt.Errorf("create auth zone: write %s for existing zone: %w", az.idVar, err)
				}

				fmt.Printf("Auth zone %q already exists, using existing ID %q (env: %s)\n", az.fqdn, existingID, az.idVar)
				continue
			}
			return fmt.Errorf("create auth zone: create %q: %w", az.fqdn, err)
		}

		if resp == nil || resp.Result == nil || resp.Result.Id == nil {
			return fmt.Errorf("create auth zone: create response for %q missing ID", az.fqdn)
		}

		createdID := *resp.Result.Id
		if err := writePipelineEnvVar(az.idVar, createdID); err != nil {
			return fmt.Errorf("create auth zone: write %s: %w", az.idVar, err)
		}

		fmt.Printf("Auth zone %q created successfully (ID: %q, env: %s)\n", az.fqdn, createdID, az.idVar)
	}

	return nil
}

// CreateIPSpaces creates two IPAM IP spaces and stores their IDs into
// pipeline_uddi.env as UDDI_IP_SPACE_ID_1 and UDDI_IP_SPACE_ID_2.
// If an IP space already exists, its existing ID is stored instead.
func CreateIPSpaces(ctx context.Context, client *uddiclient.APIClient) error {
	ipSpaces := []struct {
		name  string
		idVar string
	}{
		{name: "tf_ip_space_1", idVar: "UDDI_IP_SPACE_ID_1"},
		{name: "tf_ip_space_2", idVar: "UDDI_IP_SPACE_ID_2"},
	}

	for _, is := range ipSpaces {
		body := ipam.IPSpace{
			Name: is.name,
		}

		resp, _, err := client.IPAddressManagementAPI.IpSpaceAPI.Create(ctx).Body(body).Execute()
		if err != nil {
			if strings.Contains(err.Error(), "is already an existing") || strings.Contains(err.Error(), "conflict") || strings.Contains(err.Error(), "already exists") {
				listResp, _, listErr := client.IPAddressManagementAPI.IpSpaceAPI.List(ctx).Execute()
				if listErr != nil {
					return fmt.Errorf("create IP spaces: list existing spaces to find %q: %w", is.name, listErr)
				}

				var existingID string
				if listResp != nil {
					for _, existing := range listResp.Results {
						if existing.Name == is.name && existing.Id != nil {
							existingID = *existing.Id
							break
						}
					}
				}

				if existingID == "" {
					return fmt.Errorf("create IP spaces: IP space %q already exists but ID could not be resolved", is.name)
				}

				if err := writePipelineEnvVar(is.idVar, existingID); err != nil {
					return fmt.Errorf("create IP spaces: write %s for existing space: %w", is.idVar, err)
				}

				fmt.Printf("IP space %q already exists, using existing ID %q (env: %s)\n", is.name, existingID, is.idVar)
				continue
			}
			return fmt.Errorf("create IP spaces: create %q: %w", is.name, err)
		}

		if resp == nil || resp.Result == nil || resp.Result.Id == nil {
			return fmt.Errorf("create IP spaces: create response for %q missing ID", is.name)
		}

		createdID := *resp.Result.Id
		if err := writePipelineEnvVar(is.idVar, createdID); err != nil {
			return fmt.Errorf("create IP spaces: write %s: %w", is.idVar, err)
		}

		fmt.Printf("IP space %q created successfully (ID: %q, env: %s)\n", is.name, createdID, is.idVar)
	}

	return nil
}

// CreateAuthNSGs creates two DNS auth NSGs and stores their IDs into
// pipeline_uddi.env as UDDI_AUTH_NSG_ID_1 and UDDI_AUTH_NSG_ID_2.
// If an NSG already exists, its existing ID is stored instead.
func CreateAuthNSGs(ctx context.Context, client *uddiclient.APIClient) error {
	authNSGs := []struct {
		name  string
		idVar string
	}{
		{name: "tf_test_auth_nsg_1", idVar: "UDDI_AUTH_NSG_ID_1"},
		{name: "tf_test_auth_nsg_2", idVar: "UDDI_AUTH_NSG_ID_2"},
	}

	for _, nsg := range authNSGs {
		body := dnsconfig.AuthNSG{Name: nsg.name}
		resp, _, err := client.DNSConfigurationAPI.AuthNsgAPI.Create(ctx).Body(body).Execute()
		if err != nil {
			if strings.Contains(err.Error(), "is already an existing") || strings.Contains(err.Error(), "conflict") || strings.Contains(err.Error(), "already exists") {
				listResp, _, listErr := client.DNSConfigurationAPI.AuthNsgAPI.List(ctx).Execute()
				if listErr != nil {
					return fmt.Errorf("create auth NSGs: list existing NSGs to find %q: %w", nsg.name, listErr)
				}

				var existingID string
				if listResp != nil {
					for _, existing := range listResp.Results {
						if existing.Name == nsg.name && existing.Id != nil {
							existingID = *existing.Id
							break
						}
					}
				}

				if existingID == "" {
					return fmt.Errorf("create auth NSGs: NSG %q already exists but ID could not be resolved", nsg.name)
				}

				if err := writePipelineEnvVar(nsg.idVar, existingID); err != nil {
					return fmt.Errorf("create auth NSGs: write %s for existing NSG: %w", nsg.idVar, err)
				}

				fmt.Printf("Auth NSG %q already exists, using existing ID %q (env: %s)\n", nsg.name, existingID, nsg.idVar)
				continue
			}
			return fmt.Errorf("create auth NSGs: create %q: %w", nsg.name, err)
		}

		if resp == nil || resp.Result == nil || resp.Result.Id == nil {
			return fmt.Errorf("create auth NSGs: create response for %q missing ID", nsg.name)
		}

		createdID := *resp.Result.Id
		if err := writePipelineEnvVar(nsg.idVar, createdID); err != nil {
			return fmt.Errorf("create auth NSGs: write %s: %w", nsg.idVar, err)
		}

		fmt.Printf("Auth NSG %q created successfully (ID: %q, env: %s)\n", nsg.name, createdID, nsg.idVar)
	}

	return nil
}

// CreateView creates a DNS view and stores its ID into
// pipeline_uddi.env as UDDI_VIEW_ID_1.
// If the view already exists, its existing ID is stored instead.
func CreateView(ctx context.Context, client *uddiclient.APIClient) error {
	const viewName = "tf_test_view_1"
	const viewIDVar = "UDDI_VIEW_ID_1"

	body := dnsconfig.View{Name: viewName}
	resp, _, err := client.DNSConfigurationAPI.ViewAPI.Create(ctx).Body(body).Execute()
	if err != nil {
		if strings.Contains(err.Error(), "is already an existing") || strings.Contains(err.Error(), "conflict") || strings.Contains(err.Error(), "already exists") {
			listResp, _, listErr := client.DNSConfigurationAPI.ViewAPI.List(ctx).Execute()
			if listErr != nil {
				return fmt.Errorf("create view: list existing views to find %q: %w", viewName, listErr)
			}

			var existingID string
			if listResp != nil {
				for _, existing := range listResp.Results {
					if existing.Name == viewName && existing.Id != nil {
						existingID = *existing.Id
						break
					}
				}
			}

			if existingID == "" {
				return fmt.Errorf("create view: view %q already exists but ID could not be resolved", viewName)
			}

			if err := writePipelineEnvVar(viewIDVar, existingID); err != nil {
				return fmt.Errorf("create view: write %s for existing view: %w", viewIDVar, err)
			}

			fmt.Printf("View %q already exists, using existing ID %q (env: %s)\n", viewName, existingID, viewIDVar)
			return nil
		}
		return fmt.Errorf("create view: create %q: %w", viewName, err)
	}

	if resp == nil || resp.Result == nil || resp.Result.Id == nil {
		return fmt.Errorf("create view: create response for %q missing ID", viewName)
	}

	createdID := *resp.Result.Id
	if err := writePipelineEnvVar(viewIDVar, createdID); err != nil {
		return fmt.Errorf("create view: write %s: %w", viewIDVar, err)
	}

	fmt.Printf("View %q created successfully (ID: %q, env: %s)\n", viewName, createdID, viewIDVar)
	return nil
}

// CreateDtcPolicies creates two DTC policies and stores their IDs into
// pipeline_uddi.env as UDDI_DTC_POLICY_ID_1 and UDDI_DTC_POLICY_ID_2.
// If a policy already exists, its existing ID is stored instead.
func CreateDtcPolicies(ctx context.Context, client *uddiclient.APIClient) error {
	policies := []struct {
		name  string
		idVar string
	}{
		{name: "tf_dtc_policy_1", idVar: "UDDI_DTC_POLICY_ID_1"},
		{name: "tf_dtc_policy_2", idVar: "UDDI_DTC_POLICY_ID_2"},
	}

	for _, p := range policies {
		body := dtc.Policy{
			Name:   p.name,
			Method: "round_robin",
		}

		resp, _, err := client.DNSTrafficControlAPI.PolicyAPI.Create(ctx).Body(body).Execute()
		if err != nil {
			if strings.Contains(err.Error(), "is already an existing") || strings.Contains(err.Error(), "conflict") {
				listResp, _, listErr := client.DNSTrafficControlAPI.PolicyAPI.List(ctx).Execute()
				if listErr != nil {
					return fmt.Errorf("create dtc policies: list existing policies to find %q: %w", p.name, listErr)
				}

				var existingID string
				if listResp != nil {
					for _, existing := range listResp.Results {
						if existing.Name == p.name && existing.Id != nil {
							existingID = *existing.Id
							break
						}
					}
				}

				if existingID == "" {
					return fmt.Errorf("create dtc policies: policy %q already exists but ID could not be resolved", p.name)
				}

				if err := writePipelineEnvVar(p.idVar, existingID); err != nil {
					return fmt.Errorf("create dtc policies: write %s for existing policy: %w", p.idVar, err)
				}

				fmt.Printf("DTC policy %q already exists, using existing ID %q (env: %s)\n", p.name, existingID, p.idVar)
				continue
			}
			return fmt.Errorf("create dtc policies: create %q: %w", p.name, err)
		}

		if resp == nil || resp.Result == nil || resp.Result.Id == nil {
			return fmt.Errorf("create dtc policies: create response for %q missing ID", p.name)
		}

		createdID := *resp.Result.Id
		if err := writePipelineEnvVar(p.idVar, createdID); err != nil {
			return fmt.Errorf("create dtc policies: write %s: %w", p.idVar, err)
		}

		fmt.Printf("DTC policy %q created successfully (ID: %q, env: %s)\n", p.name, createdID, p.idVar)
	}

	return nil
}

// CreateSnmpUserSecurityModels creates two DTC SNMP User Security Model objects and
// stores their IDs into pipeline_uddi.env as UDDI_SNMP_USER_SECURITY_MODEL_ID_1 and
// UDDI_SNMP_USER_SECURITY_MODEL_ID_2. If a model already exists, its existing ID is
// stored instead.
func CreateSnmpUserSecurityModels(ctx context.Context, client *uddiclient.APIClient) error {
	models := []struct {
		username string
		idVar    string
	}{
		{username: "tf_snmp_usm_1", idVar: "UDDI_SNMP_USER_SECURITY_MODEL_ID_1"},
		{username: "tf_snmp_usm_2", idVar: "UDDI_SNMP_USER_SECURITY_MODEL_ID_2"},
	}

	for _, m := range models {
		body := dtc.SNMPUserSecurityModel{
			Username:        dtc.PtrString(m.username),
			AuthProtocol:    dtc.PtrString("NoAuth"),
			PrivacyProtocol: dtc.PtrString("NoPrivacy"),
		}

		resp, _, err := client.DNSTrafficControlAPI.SnmpUserSecurityAPI.Create(ctx).Body(body).Execute()
		if err != nil {
			if strings.Contains(err.Error(), "is already an existing") || strings.Contains(err.Error(), "conflict") {
				listResp, _, listErr := client.DNSTrafficControlAPI.SnmpUserSecurityAPI.List(ctx).Execute()
				if listErr != nil {
					return fmt.Errorf("create snmp user security models: list existing models to find %q: %w", m.username, listErr)
				}

				var existingID string
				if listResp != nil {
					for _, existing := range listResp.Results {
						if existing.Username != nil && *existing.Username == m.username && existing.Id != nil {
							existingID = *existing.Id
							break
						}
					}
				}

				if existingID == "" {
					return fmt.Errorf("create snmp user security models: model %q already exists but ID could not be resolved", m.username)
				}

				if err := writePipelineEnvVar(m.idVar, existingID); err != nil {
					return fmt.Errorf("create snmp user security models: write %s for existing model: %w", m.idVar, err)
				}

				fmt.Printf("SNMP user security model %q already exists, using existing ID %q (env: %s)\n", m.username, existingID, m.idVar)
				continue
			}
			return fmt.Errorf("create snmp user security models: create %q: %w", m.username, err)
		}

		if resp == nil || resp.Result == nil || resp.Result.Id == nil {
			return fmt.Errorf("create snmp user security models: create response for %q missing ID", m.username)
		}

		createdID := *resp.Result.Id
		if err := writePipelineEnvVar(m.idVar, createdID); err != nil {
			return fmt.Errorf("create snmp user security models: write %s: %w", m.idVar, err)
		}

		fmt.Printf("SNMP user security model %q created successfully (ID: %q, env: %s)\n", m.username, createdID, m.idVar)
	}

	return nil
}

func main() {
	cspURL := strings.TrimSpace(os.Getenv("INFOBLOX_PORTAL_URL"))
	apiKey := strings.TrimSpace(os.Getenv("INFOBLOX_PORTAL_KEY"))

	if cspURL == "" || apiKey == "" {
		fmt.Println("Missing required UDDI configuration.")
		fmt.Println("Supported env vars: INFOBLOX_PORTAL_URL, INFOBLOX_PORTAL_KEY")
		return
	}

	cwd, err := os.Getwd()
	if err != nil {
		fmt.Printf("Error getting current working directory: %v\n", err)
		return
	}

	pipelineEnvPath := filepath.Join(cwd, "pipeline_uddi.env")
	f, err := os.OpenFile(pipelineEnvPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0600)
	if err != nil {
		fmt.Printf("Error opening pipeline_uddi.env: %v\n", err)
		return
	}
	pipelineEnvFile = f
	defer func() {
		_ = pipelineEnvFile.Close()
		pipelineEnvFile = nil
	}()

	client := uddiclient.NewAPIClient(
		uddioption.WithClientName("terraform-integration-test-setup"),
		uddioption.WithCSPUrl(cspURL),
		uddioption.WithAPIKey(apiKey),
	)

	ctx := context.Background()

	infraHosts, err := StoreInfraHostDetails(ctx, client)
	if err != nil {
		fmt.Printf("Error storing infra host details: %v\n", err)
		return
	}
	fmt.Println("Infra host details stored successfully")

	if len(infraHosts) > 0 {
		if err := CreateAnyCastService(ctx, client, infraHosts[0]); err != nil {
			fmt.Printf("Error creating anycast service: %v\n", err)
			return
		}
		fmt.Println("Anycast service created successfully")
	}

	if err := StoreDNSHostIDs(ctx, client); err != nil {
		fmt.Printf("Error storing DNS host IDs: %v\n", err)
		return
	}
	fmt.Println("DNS host IDs stored successfully")

	if err := StoreDHCPHostIDs(ctx, client); err != nil {
		fmt.Printf("Error storing DHCP host IDs: %v\n", err)
		return
	}
	fmt.Println("DHCP host IDs stored successfully")

	if err := CreateOptionGroups(ctx, client); err != nil {
		fmt.Printf("Error creating option groups: %v\n", err)
		return
	}
	fmt.Println("Option groups created successfully")

	if err := CreateOptionCode(ctx, client); err != nil {
		fmt.Printf("Error creating option code: %v\n", err)
		return
	}
	fmt.Println("Option code created successfully")

	if err := CreateDtcPolicies(ctx, client); err != nil {
		fmt.Printf("Error creating dtc policies: %v\n", err)
		return
	}
	fmt.Println("DTC policies created successfully")

	if err := CreateSnmpUserSecurityModels(ctx, client); err != nil {
		fmt.Printf("Error creating snmp user security models: %v\n", err)
		return
	}
	fmt.Println("SNMP user security models created successfully")

	if err := CreateAuthZone(ctx, client); err != nil {
		fmt.Printf("Error creating auth zone: %v\n", err)
		return
	}
	fmt.Println("Auth zone created successfully")

	if err := CreateIPSpaces(ctx, client); err != nil {
		fmt.Printf("Error creating IP spaces: %v\n", err)
		return
	}
	fmt.Println("IP spaces created successfully")

	if err := CreateAuthNSGs(ctx, client); err != nil {
		fmt.Printf("Error creating auth NSGs: %v\n", err)
		return
	}
	fmt.Println("Auth NSGs created successfully")

	if err := CreateView(ctx, client); err != nil {
		fmt.Printf("Error creating view: %v\n", err)
		return
	}
	fmt.Println("View created successfully")
}
