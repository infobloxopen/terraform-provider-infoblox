// Package main provides a standalone cleanup utility that removes the
// well-known integration-test prerequisites created by
// integration_tests/setup/uddi/integration_test_setup.go, plus any dangling
// objects left behind in those prerequisites by failed test runs (e.g. an
// auth zone stranded inside the test DNS view). Only objects with hardcoded,
// well-known names (or IDs read from env vars written by setup) are
// targeted; randomly-named test resources are not touched.
//
// Objects cleared by this utility:
//
//	DNS
//	  - Auth Zone "example_zone_250" (default view)
//	  - Any Auth Zones left inside the "tf_test_view_1" DNS view
//	  - DNS View "tf_test_view_1"
//	  - Auth NSGs "tf_test_auth_nsg_1", "tf_test_auth_nsg_2"
//
//	IPAM
//	  - IP Spaces "tf_ip_space_1", "tf_ip_space_2"
//	  - DHCP Option Code "tf_option_code_1" (deleted before its option space)
//	  - DHCP Option Space "tf_option_space_1"
//	  - DHCP Option Groups "tf_option_group_1", "tf_option_group_2"
//
//	Infra Management
//	  - Anycast Service whose ID is stored in UDDI_ANYCAST_SERVICE_ID_1
package main

import (
	"context"
	"fmt"
	"os"
	"strings"

	uddiclient "github.com/infobloxopen/universal-ddi-go-client/client"
	uddioption "github.com/infobloxopen/universal-ddi-go-client/option"
)

func cleanupFirstNonEmpty(values ...string) string {
	for _, v := range values {
		if strings.TrimSpace(v) != "" {
			return v
		}
	}
	return ""
}

func cleanupAuthZonesInView(ctx context.Context, client *uddiclient.APIClient, viewID string) {
	resp, _, err := client.DNSConfigurationAPI.AuthZoneAPI.List(ctx).Execute()
	if err != nil {
		fmt.Printf("cleanup: failed to list auth zones to find zones in view %q: %v\n", viewID, err)
		return
	}
	if resp == nil || len(resp.Results) == 0 {
		return
	}
	for _, zone := range resp.Results {
		if zone.View == nil || *zone.View != viewID || zone.Id == nil {
			continue
		}
		fqdn := ""
		if zone.Fqdn != nil {
			fqdn = *zone.Fqdn
		}
		_, err := client.DNSConfigurationAPI.AuthZoneAPI.Delete(ctx, *zone.Id).Execute()
		if err != nil {
			fmt.Printf("cleanup: failed to delete auth zone %q (id=%q) in view %q: %v\n", fqdn, *zone.Id, viewID, err)
		} else {
			fmt.Printf("cleanup: deleted auth zone %q (id=%q) in view %q\n", fqdn, *zone.Id, viewID)
		}
	}
}

func cleanupAuthZoneByFQDN(ctx context.Context, client *uddiclient.APIClient, fqdn string) {
	resp, _, err := client.DNSConfigurationAPI.AuthZoneAPI.List(ctx).Execute()
	if err != nil {
		fmt.Printf("cleanup: failed to list auth zones to find %q: %v\n", fqdn, err)
		return
	}
	if resp == nil || len(resp.Results) == 0 {
		fmt.Printf("cleanup: auth zone %q not found\n", fqdn)
		return
	}
	found := false
	for _, zone := range resp.Results {
		if zone.Fqdn == nil || zone.Id == nil {
			continue
		}
		if strings.TrimSuffix(*zone.Fqdn, ".") != fqdn {
			continue
		}
		found = true
		_, err := client.DNSConfigurationAPI.AuthZoneAPI.Delete(ctx, *zone.Id).Execute()
		if err != nil {
			fmt.Printf("cleanup: failed to delete auth zone %q (id=%q): %v\n", fqdn, *zone.Id, err)
		} else {
			fmt.Printf("cleanup: deleted auth zone %q (id=%q)\n", fqdn, *zone.Id)
		}
	}
	if !found {
		fmt.Printf("cleanup: auth zone %q not found\n", fqdn)
	}
}

func cleanupView(ctx context.Context, client *uddiclient.APIClient, name string) {
	resp, _, err := client.DNSConfigurationAPI.ViewAPI.List(ctx).Execute()
	if err != nil {
		fmt.Printf("cleanup: failed to list views to find %q: %v\n", name, err)
		return
	}
	if resp == nil || len(resp.Results) == 0 {
		fmt.Printf("cleanup: view %q not found\n", name)
		return
	}
	for _, view := range resp.Results {
		if view.Name != name || view.Id == nil {
			continue
		}
		// Delete any auth zones left inside this view first, otherwise the
		// view delete will fail with a conflict.
		cleanupAuthZonesInView(ctx, client, *view.Id)

		_, err := client.DNSConfigurationAPI.ViewAPI.Delete(ctx, *view.Id).Execute()
		if err != nil {
			fmt.Printf("cleanup: failed to delete view %q (id=%q): %v\n", name, *view.Id, err)
		} else {
			fmt.Printf("cleanup: deleted view %q (id=%q)\n", name, *view.Id)
		}
		return
	}
	fmt.Printf("cleanup: view %q not found\n", name)
}

func cleanupAuthNSGs(ctx context.Context, client *uddiclient.APIClient, names []string) {
	resp, _, err := client.DNSConfigurationAPI.AuthNsgAPI.List(ctx).Execute()
	if err != nil {
		fmt.Printf("cleanup: failed to list auth NSGs: %v\n", err)
		return
	}
	if resp == nil || len(resp.Results) == 0 {
		fmt.Println("cleanup: no auth NSGs found")
		return
	}
	for _, name := range names {
		found := false
		for _, nsg := range resp.Results {
			if nsg.Name != name || nsg.Id == nil {
				continue
			}
			found = true
			_, err := client.DNSConfigurationAPI.AuthNsgAPI.Delete(ctx, *nsg.Id).Execute()
			if err != nil {
				fmt.Printf("cleanup: failed to delete auth NSG %q (id=%q): %v\n", name, *nsg.Id, err)
			} else {
				fmt.Printf("cleanup: deleted auth NSG %q (id=%q)\n", name, *nsg.Id)
			}
		}
		if !found {
			fmt.Printf("cleanup: auth NSG %q not found\n", name)
		}
	}
}

func cleanupIPSpaces(ctx context.Context, client *uddiclient.APIClient, names []string) {
	resp, _, err := client.IPAddressManagementAPI.IpSpaceAPI.List(ctx).Execute()
	if err != nil {
		fmt.Printf("cleanup: failed to list IP spaces: %v\n", err)
		return
	}
	if resp == nil || len(resp.Results) == 0 {
		fmt.Println("cleanup: no IP spaces found")
		return
	}
	for _, name := range names {
		found := false
		for _, is := range resp.Results {
			if is.Name != name || is.Id == nil {
				continue
			}
			found = true
			_, err := client.IPAddressManagementAPI.IpSpaceAPI.Delete(ctx, *is.Id).Execute()
			if err != nil {
				fmt.Printf("cleanup: failed to delete IP space %q (id=%q): %v\n", name, *is.Id, err)
			} else {
				fmt.Printf("cleanup: deleted IP space %q (id=%q)\n", name, *is.Id)
			}
		}
		if !found {
			fmt.Printf("cleanup: IP space %q not found\n", name)
		}
	}
}

func cleanupOptionCode(ctx context.Context, client *uddiclient.APIClient, name string) {
	resp, _, err := client.IPAddressManagementAPI.OptionCodeAPI.List(ctx).Execute()
	if err != nil {
		fmt.Printf("cleanup: failed to list option codes to find %q: %v\n", name, err)
		return
	}
	if resp == nil || len(resp.Results) == 0 {
		fmt.Printf("cleanup: option code %q not found\n", name)
		return
	}
	for _, oc := range resp.Results {
		if oc.Name != name || oc.Id == nil {
			continue
		}
		_, err := client.IPAddressManagementAPI.OptionCodeAPI.Delete(ctx, *oc.Id).Execute()
		if err != nil {
			fmt.Printf("cleanup: failed to delete option code %q (id=%q): %v\n", name, *oc.Id, err)
		} else {
			fmt.Printf("cleanup: deleted option code %q (id=%q)\n", name, *oc.Id)
		}
		return
	}
	fmt.Printf("cleanup: option code %q not found\n", name)
}

func cleanupOptionSpace(ctx context.Context, client *uddiclient.APIClient, name string) {
	resp, _, err := client.IPAddressManagementAPI.OptionSpaceAPI.List(ctx).Execute()
	if err != nil {
		fmt.Printf("cleanup: failed to list option spaces to find %q: %v\n", name, err)
		return
	}
	if resp == nil || len(resp.Results) == 0 {
		fmt.Printf("cleanup: option space %q not found\n", name)
		return
	}
	for _, os_ := range resp.Results {
		if os_.Name != name || os_.Id == nil {
			continue
		}
		_, err := client.IPAddressManagementAPI.OptionSpaceAPI.Delete(ctx, *os_.Id).Execute()
		if err != nil {
			fmt.Printf("cleanup: failed to delete option space %q (id=%q): %v\n", name, *os_.Id, err)
		} else {
			fmt.Printf("cleanup: deleted option space %q (id=%q)\n", name, *os_.Id)
		}
		return
	}
	fmt.Printf("cleanup: option space %q not found\n", name)
}

func cleanupOptionGroups(ctx context.Context, client *uddiclient.APIClient, names []string) {
	resp, _, err := client.IPAddressManagementAPI.OptionGroupAPI.List(ctx).Execute()
	if err != nil {
		fmt.Printf("cleanup: failed to list option groups: %v\n", err)
		return
	}
	if resp == nil || len(resp.Results) == 0 {
		fmt.Println("cleanup: no option groups found")
		return
	}
	for _, name := range names {
		found := false
		for _, og := range resp.Results {
			if og.Name != name || og.Id == nil {
				continue
			}
			found = true
			_, err := client.IPAddressManagementAPI.OptionGroupAPI.Delete(ctx, *og.Id).Execute()
			if err != nil {
				fmt.Printf("cleanup: failed to delete option group %q (id=%q): %v\n", name, *og.Id, err)
			} else {
				fmt.Printf("cleanup: deleted option group %q (id=%q)\n", name, *og.Id)
			}
		}
		if !found {
			fmt.Printf("cleanup: option group %q not found\n", name)
		}
	}
}

func cleanupAnyCastService(ctx context.Context, client *uddiclient.APIClient) {
	serviceID := strings.TrimSpace(os.Getenv("UDDI_ANYCAST_SERVICE_ID_1"))
	if serviceID == "" {
		fmt.Println("cleanup: UDDI_ANYCAST_SERVICE_ID_1 is not set, skipping anycast service deletion")
		return
	}

	_, err := client.InfraManagementAPI.ServicesAPI.Delete(ctx, serviceID).Execute()
	if err != nil {
		fmt.Printf("cleanup: failed to delete anycast service (ID=%q): %v\n", serviceID, err)
		return
	}

	fmt.Printf("cleanup: deleted anycast service (ID=%q)\n", serviceID)
}

func Cleanup(client *uddiclient.APIClient) {
	ctx := context.Background()

	fmt.Println("--- Cleaning up Auth Zone (example_zone_250) ---")
	cleanupAuthZoneByFQDN(ctx, client, "example_zone_250")

	fmt.Println("--- Cleaning up DNS View (tf_test_view_1) and any zones left inside it ---")
	cleanupView(ctx, client, "tf_test_view_1")

	fmt.Println("--- Cleaning up Auth NSGs (tf_test_auth_nsg_1, tf_test_auth_nsg_2) ---")
	cleanupAuthNSGs(ctx, client, []string{"tf_test_auth_nsg_1", "tf_test_auth_nsg_2"})

	fmt.Println("--- Cleaning up IP Spaces (tf_ip_space_1, tf_ip_space_2) ---")
	cleanupIPSpaces(ctx, client, []string{"tf_ip_space_1", "tf_ip_space_2"})

	fmt.Println("--- Cleaning up Option Code (tf_option_code_1) ---")
	cleanupOptionCode(ctx, client, "tf_option_code_1")

	fmt.Println("--- Cleaning up Option Space (tf_option_space_1) ---")
	cleanupOptionSpace(ctx, client, "tf_option_space_1")

	fmt.Println("--- Cleaning up Option Groups (tf_option_group_1, tf_option_group_2) ---")
	cleanupOptionGroups(ctx, client, []string{"tf_option_group_1", "tf_option_group_2"})

	fmt.Println("--- Cleaning up Anycast Service (UDDI_ANYCAST_SERVICE_ID_1) ---")
	cleanupAnyCastService(ctx, client)
}

func main() {
	cspURL := strings.TrimSpace(cleanupFirstNonEmpty(os.Getenv("INFOBLOX_PORTAL_URL")))
	apiKey := strings.TrimSpace(cleanupFirstNonEmpty(os.Getenv("INFOBLOX_PORTAL_KEY")))

	if cspURL == "" || apiKey == "" {
		fmt.Println("Missing required UDDI configuration.")
		fmt.Println("Supported env vars: INFOBLOX_PORTAL_URL, INFOBLOX_PORTAL_KEY")
		os.Exit(1)
	}

	client := uddiclient.NewAPIClient(
		uddioption.WithClientName("terraform-integration-test-cleanup"),
		uddioption.WithCSPUrl(cspURL),
		uddioption.WithAPIKey(apiKey),
	)

	fmt.Println("Starting cleanup of UDDI integration test prerequisites...")
	Cleanup(client)
	fmt.Println("Cleanup complete.")
}
