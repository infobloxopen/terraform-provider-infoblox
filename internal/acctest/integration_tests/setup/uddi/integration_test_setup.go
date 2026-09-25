// Objects read by this setup program (not created, IDs stored as env vars):
//
// DNS Hosts (up to 2):
//   - UDDI_DNS_HOST_ID_1, UDDI_DNS_HOST_ID_2
//
// DHCP Hosts (up to 2):
//   - UDDI_DHCP_HOST_ID_1, UDDI_DHCP_HOST_ID_2
//
// Objects created by this setup program (IDs stored as env vars):
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

package main

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	uddiclient "github.com/infobloxopen/universal-ddi-go-client/client"
	"github.com/infobloxopen/universal-ddi-go-client/dnsconfig"
	"github.com/infobloxopen/universal-ddi-go-client/ipam"
	uddioption "github.com/infobloxopen/universal-ddi-go-client/option"
)

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
				listResp, _, listErr := client.IPAddressManagementAPI.OptionGroupAPI.List(ctx).Execute()
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
			listResp, _, listErr := client.IPAddressManagementAPI.OptionCodeAPI.List(ctx).Execute()
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
			listResp, _, listErr := client.IPAddressManagementAPI.OptionSpaceAPI.List(ctx).Execute()
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
				listResp, _, listErr := client.DNSConfigurationAPI.AuthZoneAPI.List(ctx).Execute()
				if listErr != nil {
					return fmt.Errorf("create auth zone: list existing zones to find %q: %w", az.fqdn, listErr)
				}

				var existingID string
				if listResp != nil {
					for _, existing := range listResp.Results {
						if existing.Fqdn != nil && *existing.Fqdn == az.fqdn && existing.Id != nil {
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

	if err := CreateAuthZone(ctx, client); err != nil {
		fmt.Printf("Error creating auth zone: %v\n", err)
		return
	}
	fmt.Println("Auth zone created successfully")
}
