package infoblox

import (
	"fmt"
	"net"
	"sort"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	ibclient "github.com/infobloxopen/infoblox-go-client/v2"
)

func testAccCheckZoneDelegatedDestroy(s *terraform.State) error {
	meta := testAccProvider.Meta()

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "infoblox_zone_delegated" {
			return fmt.Errorf("Resource type %s is invalid after destroy", rs.Type)
		}
		connector := meta.(ibclient.IBConnector)
		objMgr := ibclient.NewObjectManager(connector, "terraform_test", "test")
		rec, _ := objMgr.GetZoneDelegated(rs.Primary.ID)
		if rec != nil {
			return fmt.Errorf("Zone Delegation record found after destroy")
		}
	}
	return nil
}

func testAccZoneDelegatedCompare(t *testing.T, resPath string, expectedRec *ibclient.RecordNS) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		res, found := s.RootModule().Resources[resPath]
		if !found {
			return fmt.Errorf("Not found: %s", resPath)
		}
		if res.Primary.ID == "" {
			return fmt.Errorf("ID is not set")
		}
		meta := testAccProvider.Meta()
		connector := meta.(ibclient.IBConnector)
		objMgr := ibclient.NewObjectManager(connector, "terraform_test", "test")

		lookupHosts, err := net.LookupHost(*expectedRec.Nameserver)
		if err != nil {
			return fmt.Errorf("Failed to resolve delegate_to: %s", err.Error())
		}
		sort.Strings(lookupHosts)
		expectedRec.Addresses = append(expectedRec.Addresses, &ibclient.ZoneNameServer{Address: lookupHosts[0]})

		rec, _ := objMgr.GetZoneDelegated(res.Primary.ID)
		if rec == nil {
			return fmt.Errorf("record not found")
		}

		if rec.Fqdn != expectedRec.Name {
			return fmt.Errorf(
				"'fqdn' does not match: got '%s', expected '%s'",
				rec.Fqdn, expectedRec.Name)
		}
		if rec.DelegateTo[0].Address != expectedRec.Addresses[0].Address {
			return fmt.Errorf(
				"'delegate_to['address']' does not match: got '%s', expected '%s'",
				rec.DelegateTo[0].Address, expectedRec.Addresses[0].Address)
		}
		if rec.DelegateTo[0].Name != *expectedRec.Nameserver {
			return fmt.Errorf(
				"'delegate_to['name']' does not match: got '%s', expected '%s'",
				rec.DelegateTo[0].Name, *expectedRec.Nameserver)
		}
		return nil
	}
}

func strPtr(s string) *string {
	return &s
}

func TestAccResourceZoneDelegated(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckZoneDelegatedDestroy,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
						resource "infoblox_zone_delegated" "foo"{
							fqdn="subdomain.test.com"
							delegate_to {
								name = "ns2.infoblox.com"
							}
							}`),
				Check: resource.ComposeTestCheckFunc(
					testAccZoneDelegatedCompare(t, "infoblox_zone_delegated.foo", &ibclient.RecordNS{
						Name:       "subdomain.test.com",
						Nameserver: strPtr("ns2.infoblox.com"),
					}),
				),
			},
		},
	})
}
