// Package main provides a standalone cleanup utility that removes dangling
// integration-test resources left behind in the NIOS grid after test failures.
// Only objects whose primary keys are hardcoded (IPs, FQDNs, well-known names)
// are targeted. Randomly-named objects are not touched.
//
// Objects cleared by this utility:
//
//	Cloud
//	  - AWS Route53 Task Groups whose name starts with "test-taskgroup"
//	  - Cloud AWS Users whose name starts with "aws-user"
//
//	DHCP
//	  - All DHCP Failover associations (any failover found in the grid)
//	  - Shared Networks whose name starts with "shared_network"
//
//	IPAM / Networks
//	  - Networks: 10.0.0.0/24, 15.0.0.0/24, 16.0.0.0/24, 85.85.0.0/16 (exact match)
//	  - Networks: any network matching 201.*/24 (regex filter)
//	  - Network Views whose name starts with "network-view" (regex filter)
//
//	DNS
//	  - Authoritative Zone "example.com" in the default view
//
//	Grid Members
//	  - Members whose vip_setting.address matches 172.28.38.* that are NOT running
//
//	Parental Control
//	  - Parental Control AVPs whose name starts with "parentalcontrol-avp"
//
//	Microsoft
//	  - Microsoft Servers whose address matches 10\.10.* (regex filter)
package main

import (
	"context"
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/infobloxopen/infoblox-nios-go-client/client"
	"github.com/infobloxopen/infoblox-nios-go-client/option"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/core"
)

func cleanupFirstNonEmpty(values ...string) string {
	for _, v := range values {
		if strings.TrimSpace(v) != "" {
			return v
		}
	}
	return ""
}

func cleanupAWSRoute53TaskGroups(ctx context.Context, apiClient *client.APIClient) {
	filters := map[string]interface{}{"name~": "^test-taskgroup"}
	resp, _, err := apiClient.CloudAPI.Awsrte53taskgroupAPI.List(ctx).
		Filters(filters).
		ReturnAsObject(1).
		Execute()
	if err != nil {
		fmt.Printf("cleanup: failed to list AWS Route53 task groups: %v\n", err)
		return
	}
	if resp == nil || resp.ListAwsrte53taskgroupResponseObject == nil || len(resp.ListAwsrte53taskgroupResponseObject.Result) == 0 {
		fmt.Println("cleanup: no AWS Route53 task groups matching '^test-taskgroup' found")
		return
	}
	for _, tg := range resp.ListAwsrte53taskgroupResponseObject.Result {
		ref := core.ExtractNIOSRef(tg.GetRef())
		if ref == "" {
			continue
		}
		_, err := apiClient.CloudAPI.Awsrte53taskgroupAPI.Delete(ctx, ref).Execute()
		if err != nil {
			fmt.Printf("cleanup: failed to delete AWS Route53 task group (ref=%q): %v\n", ref, err)
		} else {
			fmt.Printf("cleanup: deleted AWS Route53 task group (ref=%q)\n", ref)
		}
	}
}

func cleanupCloudAWSUsers(ctx context.Context, apiClient *client.APIClient) {
	filters := map[string]interface{}{"name~": "^aws-user"}
	resp, _, err := apiClient.CloudAPI.AwsuserAPI.List(ctx).
		Filters(filters).
		ReturnAsObject(1).
		Execute()
	if err != nil {
		fmt.Printf("cleanup: failed to list AWS users: %v\n", err)
		return
	}
	if resp == nil || resp.ListAwsuserResponseObject == nil || len(resp.ListAwsuserResponseObject.Result) == 0 {
		fmt.Println("cleanup: no AWS users matching '^aws-user' found")
		return
	}
	for _, user := range resp.ListAwsuserResponseObject.Result {
		ref := core.ExtractNIOSRef(user.GetRef())
		if ref == "" {
			continue
		}
		_, err := apiClient.CloudAPI.AwsuserAPI.Delete(ctx, ref).Execute()
		if err != nil {
			fmt.Printf("cleanup: failed to delete AWS user (ref=%q): %v\n", ref, err)
		} else {
			fmt.Printf("cleanup: deleted AWS user (ref=%q)\n", ref)
		}
	}
}

func cleanupDHCPFailovers(ctx context.Context, apiClient *client.APIClient) {
	resp, _, err := apiClient.DHCPAPI.DhcpfailoverAPI.List(ctx).
		ReturnAsObject(1).
		Execute()
	if err != nil {
		fmt.Printf("cleanup: failed to list DHCP failovers: %v\n", err)
		return
	}
	if resp == nil || resp.ListDhcpfailoverResponseObject == nil || len(resp.ListDhcpfailoverResponseObject.Result) == 0 {
		fmt.Println("cleanup: no DHCP failovers found")
		return
	}
	for _, failover := range resp.ListDhcpfailoverResponseObject.Result {
		ref := core.ExtractNIOSRef(failover.GetRef())
		if ref == "" {
			continue
		}
		_, err := apiClient.DHCPAPI.DhcpfailoverAPI.Delete(ctx, ref).Execute()
		if err != nil {
			fmt.Printf("cleanup: failed to delete DHCP failover (ref=%q): %v\n", ref, err)
		} else {
			fmt.Printf("cleanup: deleted DHCP failover (ref=%q)\n", ref)
		}
	}
}

func cleanupSharedNetworks(ctx context.Context, apiClient *client.APIClient) {
	filters := map[string]interface{}{"name~": "^shared_network"}
	resp, _, err := apiClient.DHCPAPI.SharednetworkAPI.List(ctx).
		Filters(filters).
		ReturnAsObject(1).
		Execute()
	if err != nil {
		fmt.Printf("cleanup: failed to list shared networks: %v\n", err)
		return
	}
	if resp == nil || resp.ListSharednetworkResponseObject == nil || len(resp.ListSharednetworkResponseObject.Result) == 0 {
		fmt.Println("cleanup: no shared networks matching '^shared_network' found")
		return
	}
	for _, sn := range resp.ListSharednetworkResponseObject.Result {
		ref := core.ExtractNIOSRef(sn.GetRef())
		if ref == "" {
			continue
		}
		_, err := apiClient.DHCPAPI.SharednetworkAPI.Delete(ctx, ref).Execute()
		if err != nil {
			fmt.Printf("cleanup: failed to delete shared network (ref=%q): %v\n", ref, err)
		} else {
			fmt.Printf("cleanup: deleted shared network (ref=%q)\n", ref)
		}
	}
}

func cleanupNetworks(ctx context.Context, apiClient *client.APIClient) {
	hardcoded := []string{"10.0.0.0/24", "15.0.0.0/24", "16.0.0.0/24", "85.85.0.0/16"}
	for _, cidr := range hardcoded {
		filters := map[string]interface{}{"network": cidr}
		resp, _, err := apiClient.IPAMAPI.NetworkAPI.List(ctx).
			Filters(filters).
			ReturnAsObject(1).
			Execute()
		if err != nil {
			fmt.Printf("cleanup: failed to list network %q: %v\n", cidr, err)
			continue
		}
		if resp == nil || resp.ListNetworkResponseObject == nil || len(resp.ListNetworkResponseObject.Result) == 0 {
			fmt.Printf("cleanup: network %q not found, skipping\n", cidr)
			continue
		}
		for _, net := range resp.ListNetworkResponseObject.Result {
			ref := core.ExtractNIOSRef(net.GetRef())
			if ref == "" {
				continue
			}
			_, err := apiClient.IPAMAPI.NetworkAPI.Delete(ctx, ref).Execute()
			if err != nil {
				fmt.Printf("cleanup: failed to delete network %q (ref=%q): %v\n", cidr, ref, err)
			} else {
				fmt.Printf("cleanup: deleted network %q (ref=%q)\n", cidr, ref)
			}
		}
	}

	regexFilters := map[string]interface{}{"network~": `201.*/24`}
	resp201, _, err := apiClient.IPAMAPI.NetworkAPI.List(ctx).
		Filters(regexFilters).
		ReturnAsObject(1).
		Execute()
	if err != nil {
		fmt.Printf("cleanup: failed to list 201.*/24 networks: %v\n", err)
		return
	}
	if resp201 == nil || resp201.ListNetworkResponseObject == nil || len(resp201.ListNetworkResponseObject.Result) == 0 {
		fmt.Println("cleanup: no 201.*/24 networks found")
		return
	}
	for _, net := range resp201.ListNetworkResponseObject.Result {
		fullRef := net.GetRef()
		ref := core.ExtractNIOSRef(fullRef)
		if ref == "" {
			continue
		}
		_, err := apiClient.IPAMAPI.NetworkAPI.Delete(ctx, ref).Execute()
		if err != nil {
			fmt.Printf("cleanup: failed to delete 201.*/24 network %q (ref=%q): %v\n", fullRef, ref, err)
		} else {
			fmt.Printf("cleanup: deleted 201.*/24 network %q (ref=%q)\n", fullRef, ref)
		}
	}
}

func cleanupIpv6Networks(ctx context.Context, apiClient *client.APIClient) {
	// Delete hardcoded IPv6 CIDRs used by tests across all network views.
	type ipv6NetworkEntry struct {
		cidr        string
		networkView string
	}
	entries := []ipv6NetworkEntry{
		{"119::/64", "test_network_view"},
		{"140::/64", "default"},
		{"141::/64", "default"},
		{"219::/64", "test_network_view"},
		{"221::/64", "test_network_view"},
	}
	for _, e := range entries {
		filters := map[string]interface{}{"network": e.cidr, "network_view": e.networkView}
		resp, _, err := apiClient.IPAMAPI.Ipv6networkAPI.List(ctx).
			Filters(filters).
			ReturnAsObject(1).
			Execute()
		if err != nil {
			fmt.Printf("cleanup: failed to list ipv6 network %q (view=%q): %v\n", e.cidr, e.networkView, err)
			continue
		}
		if resp == nil || resp.ListIpv6networkResponseObject == nil || len(resp.ListIpv6networkResponseObject.Result) == 0 {
			fmt.Printf("cleanup: ipv6 network %q (view=%q) not found, skipping\n", e.cidr, e.networkView)
			continue
		}
		for _, net := range resp.ListIpv6networkResponseObject.Result {
			ref := core.ExtractNIOSRef(net.GetRef())
			if ref == "" {
				continue
			}
			_, err := apiClient.IPAMAPI.Ipv6networkAPI.Delete(ctx, ref).Execute()
			if err != nil {
				fmt.Printf("cleanup: failed to delete ipv6 network %q (view=%q, ref=%q): %v\n", e.cidr, e.networkView, ref, err)
			} else {
				fmt.Printf("cleanup: deleted ipv6 network %q (view=%q, ref=%q)\n", e.cidr, e.networkView, ref)
			}
		}
	}
}

func cleanupNetworkViews(ctx context.Context, apiClient *client.APIClient) {
	filters := map[string]interface{}{"name~": "^network-view"}
	resp, _, err := apiClient.IPAMAPI.NetworkviewAPI.List(ctx).
		Filters(filters).
		ReturnAsObject(1).
		Execute()
	if err != nil {
		fmt.Printf("cleanup: failed to list network views: %v\n", err)
		return
	}
	if resp == nil || resp.ListNetworkviewResponseObject == nil || len(resp.ListNetworkviewResponseObject.Result) == 0 {
		fmt.Println("cleanup: no network views matching '^network-view' found")
		return
	}
	for _, nv := range resp.ListNetworkviewResponseObject.Result {
		ref := core.ExtractNIOSRef(nv.GetRef())
		if ref == "" {
			continue
		}
		_, err := apiClient.IPAMAPI.NetworkviewAPI.Delete(ctx, ref).Execute()
		if err != nil {
			fmt.Printf("cleanup: failed to delete network view (ref=%q): %v\n", ref, err)
		} else {
			fmt.Printf("cleanup: deleted network view (ref=%q)\n", ref)
		}
	}
}

func cleanupDNSZone(ctx context.Context, apiClient *client.APIClient) {
	filters := map[string]interface{}{
		"fqdn": "example.com",
		"view": "default",
	}
	resp, _, err := apiClient.DNSAPI.ZoneAuthAPI.List(ctx).
		Filters(filters).
		ReturnAsObject(1).
		Execute()
	if err != nil {
		fmt.Printf("cleanup: failed to list zone auth for 'example.com': %v\n", err)
		return
	}
	if resp == nil || resp.ListZoneAuthResponseObject == nil || len(resp.ListZoneAuthResponseObject.Result) == 0 {
		fmt.Println("cleanup: zone auth 'example.com' in default view not found")
		return
	}
	for _, zone := range resp.ListZoneAuthResponseObject.Result {
		ref := core.ExtractNIOSRef(zone.GetRef())
		if ref == "" {
			continue
		}
		_, err := apiClient.DNSAPI.ZoneAuthAPI.Delete(ctx, ref).Execute()
		if err != nil {
			fmt.Printf("cleanup: failed to delete zone 'example.com' (ref=%q): %v\n", ref, err)
		} else {
			fmt.Printf("cleanup: deleted zone auth 'example.com' (ref=%q)\n", ref)
		}
	}
}

func cleanupMembers(ctx context.Context, apiClient *client.APIClient) {
	resp, _, err := apiClient.GridAPI.MemberAPI.List(ctx).
		ReturnAsObject(1).
		ReturnFieldsPlus("service_status,vip_setting").
		Execute()
	if err != nil {
		fmt.Printf("cleanup: failed to list members: %v\n", err)
		return
	}
	if resp == nil || resp.ListMemberResponseObject == nil || len(resp.ListMemberResponseObject.Result) == 0 {
		fmt.Println("cleanup: no members found")
		return
	}

	vipRegex := regexp.MustCompile(`^172\.28\.38\.`)

	for _, member := range resp.ListMemberResponseObject.Result {
		fullRef := member.GetRef()
		if fullRef == "" {
			continue
		}

		vipSetting := member.GetVipSetting()
		vipAddr := vipSetting.GetAddress()
		if !vipRegex.MatchString(vipAddr) {
			continue
		}

		cloudDNSSyncWorking := false
		for _, svc := range member.GetServiceStatus() {
			if svc.GetService() == "CLOUD_DNS_SYNC" {
				if svc.GetStatus() == "WORKING" {
					cloudDNSSyncWorking = true
				}
				break
			}
		}

		if cloudDNSSyncWorking {
			fmt.Printf("cleanup: skipping member %q (vip=%s) — CLOUD_DNS_SYNC is WORKING\n", fullRef, vipAddr)
			continue
		}

		ref := core.ExtractNIOSRef(fullRef)
		_, err := apiClient.GridAPI.MemberAPI.Delete(ctx, ref).Execute()
		if err != nil {
			fmt.Printf("cleanup: failed to delete member %q (ref=%q): %v\n", fullRef, ref, err)
		} else {
			fmt.Printf("cleanup: deleted member %q (vip=%s, ref=%q)\n", fullRef, vipAddr, ref)
		}
	}
}

func cleanupParentalControlAVPs(ctx context.Context, apiClient *client.APIClient) {
	filters := map[string]interface{}{"name~": "^parentalcontrol-avp"}
	resp, _, err := apiClient.ParentalControlAPI.ParentalcontrolAvpAPI.List(ctx).
		Filters(filters).
		ReturnAsObject(1).
		Execute()
	if err != nil {
		fmt.Printf("cleanup: failed to list parental control AVPs: %v\n", err)
		return
	}
	if resp == nil || resp.ListParentalcontrolAvpResponseObject == nil || len(resp.ListParentalcontrolAvpResponseObject.Result) == 0 {
		fmt.Println("cleanup: no parental control AVPs matching '^parentalcontrol-avp' found")
		return
	}
	for _, avp := range resp.ListParentalcontrolAvpResponseObject.Result {
		ref := core.ExtractNIOSRef(avp.GetRef())
		if ref == "" {
			continue
		}
		_, err := apiClient.ParentalControlAPI.ParentalcontrolAvpAPI.Delete(ctx, ref).Execute()
		if err != nil {
			fmt.Printf("cleanup: failed to delete parental control AVP (ref=%q): %v\n", ref, err)
		} else {
			fmt.Printf("cleanup: deleted parental control AVP (ref=%q)\n", ref)
		}
	}
}

func cleanupMicrosoftServers(ctx context.Context, apiClient *client.APIClient) {
	filters := map[string]interface{}{"address~": `10\.10.*`}
	resp, _, err := apiClient.MicrosoftAPI.MsserverAPI.List(ctx).
		Filters(filters).
		ReturnAsObject(1).
		Execute()
	if err != nil {
		fmt.Printf("cleanup: failed to list Microsoft servers: %v\n", err)
		return
	}
	if resp == nil || resp.ListMsserverResponseObject == nil || len(resp.ListMsserverResponseObject.Result) == 0 {
		fmt.Println("cleanup: no Microsoft servers matching '10.10.*' found")
		return
	}
	for _, server := range resp.ListMsserverResponseObject.Result {
		ref := core.ExtractNIOSRef(server.GetRef())
		if ref == "" {
			continue
		}
		_, err := apiClient.MicrosoftAPI.MsserverAPI.Delete(ctx, ref).Execute()
		if err != nil {
			fmt.Printf("cleanup: failed to delete Microsoft server (ref=%q): %v\n", ref, err)
		} else {
			fmt.Printf("cleanup: deleted Microsoft server (ref=%q)\n", ref)
		}
	}
}

func Cleanup(apiClient *client.APIClient) {
	ctx := context.Background()

	fmt.Println("--- Cleaning up AWS Route53 Task Groups (prefix: test-taskgroup) ---")
	cleanupAWSRoute53TaskGroups(ctx, apiClient)

	fmt.Println("--- Cleaning up Cloud AWS Users ---")
	cleanupCloudAWSUsers(ctx, apiClient)

	fmt.Println("--- Cleaning up DHCP Failovers ---")
	cleanupDHCPFailovers(ctx, apiClient)

	fmt.Println("--- Cleaning up Shared Networks (prefix: shared_network) ---")
	cleanupSharedNetworks(ctx, apiClient)

	fmt.Println("--- Cleaning up Networks (10.0.0.0/24, 15.0.0.0/24, 16.0.0.0/24, 85.85.0.0/16, 201.*/24) ---")
	cleanupNetworks(ctx, apiClient)

	fmt.Println("--- Cleaning up IPv6 Networks (test CIDRs) ---")
	cleanupIpv6Networks(ctx, apiClient)

	fmt.Println("--- Cleaning up Network Views (prefix: network-view) ---")
	cleanupNetworkViews(ctx, apiClient)

	fmt.Println("--- Cleaning up DNS Zone (example.com / default) ---")
	cleanupDNSZone(ctx, apiClient)

	fmt.Println("--- Cleaning up Members (prefix: 172.28.38, non-running only) ---")
	cleanupMembers(ctx, apiClient)

	fmt.Println("--- Cleaning up Parental Control AVPs ---")
	cleanupParentalControlAVPs(ctx, apiClient)

	fmt.Println("--- Cleaning up Microsoft Servers (address prefix: 10.10) ---")
	cleanupMicrosoftServers(ctx, apiClient)
}

func main() {
	host := strings.TrimSpace(cleanupFirstNonEmpty(os.Getenv("NIOS_HOST_URL")))
	username := strings.TrimSpace(cleanupFirstNonEmpty(os.Getenv("NIOS_USERNAME")))
	password := strings.TrimSpace(cleanupFirstNonEmpty(os.Getenv("NIOS_PASSWORD")))

	if host == "" || username == "" || password == "" {
		fmt.Println("Missing required NIOS configuration. Ensure host, username, and password are set.")
		fmt.Println("Supported env vars: NIOS_HOST_URL, NIOS_USERNAME, NIOS_PASSWORD")
		os.Exit(1)
	}

	if !strings.HasPrefix(host, "https://") {
		fmt.Printf("Invalid NIOS host %q: must include https://\n", host)
		os.Exit(1)
	}

	apiClient := client.NewAPIClient(
		option.WithNIOSHostUrl(host),
		option.WithNIOSUsername(username),
		option.WithNIOSPassword(password),
		option.WithDebug(true),
	)

	fmt.Println("Starting cleanup of dangling integration test resources...")
	Cleanup(apiClient)
	fmt.Println("Cleanup complete.")
}
