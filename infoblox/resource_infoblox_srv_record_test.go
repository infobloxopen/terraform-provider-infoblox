package infoblox

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	ibclient "github.com/infobloxopen/infoblox-go-client/v2"
)

func testAccCheckSRVRecordDestroy(s *terraform.State) error {
	meta := testAccProvider.Meta()
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "resource_srv_record" {
			continue
		}
		connector := meta.(ibclient.IBConnector)
		objMgr := ibclient.NewObjectManager(connector, "terraform_test", "test")
		rec, _ := objMgr.GetSRVRecordByRef(rs.Primary.ID)
		if rec != nil {
			return fmt.Errorf("record not found")
		}
	}
	return nil
}

func testAccSRVRecordCompare(t *testing.T, resPath string, expectedRec *ibclient.RecordSRV) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		res, found := s.RootModule().Resources[resPath]
		if !found {
			return fmt.Errorf("not found: %s", resPath)
		}
		if res.Primary.ID == "" {
			return fmt.Errorf("ID is not set")
		}
		meta := testAccProvider.Meta()
		connector := meta.(ibclient.IBConnector)
		objMgr := ibclient.NewObjectManager(connector, "terraform_test", "test")

		rec, _ := objMgr.GetSRVRecordByRef(res.Primary.ID)
		if rec == nil {
			return fmt.Errorf("record not found")
		}

		if rec.Name != expectedRec.Name {
			return fmt.Errorf(
				"'name' does not match: got '%s', expected '%s'",
				rec.Name, expectedRec.Name)
		}

		if rec.View != expectedRec.View {
			return fmt.Errorf(
				"'dns_view' does not match: got '%s', expected '%s'",
				rec.View, expectedRec.View)
		}

		if rec.Priority != expectedRec.Priority {
			return fmt.Errorf(
				"'priority' does not match: got '%d', expected '%d'",
				rec.Priority, expectedRec.Priority)
		}

		if rec.Weight != expectedRec.Weight {
			return fmt.Errorf(
				"'weight' does not match: got '%d', expected '%d'",
				rec.Weight, expectedRec.Weight)
		}

		if rec.Port != expectedRec.Port {
			return fmt.Errorf(
				"'port' does not match: got '%d', expected '%d'",
				rec.Port, expectedRec.Port)
		}

		if rec.Target != expectedRec.Target {
			return fmt.Errorf(
				"'target' does not match: got '%s', expected '%s'",
				rec.Target, expectedRec.Target)
		}

		if rec.UseTtl != expectedRec.UseTtl {
			return fmt.Errorf(
				"TTL usage does not match: got '%t', expected '%t'",
				rec.UseTtl, expectedRec.UseTtl)
		}
		if rec.UseTtl {
			if rec.Ttl != expectedRec.Ttl {
				return fmt.Errorf(
					"'Ttl' usage does not match: got '%d', expected '%d'",
					rec.Ttl, expectedRec.Ttl)
			}
		}
		if rec.Comment != expectedRec.Comment {
			return fmt.Errorf(
				"'comment' does not match: got '%s', expected '%s'",
				rec.Comment, expectedRec.Comment)
		}
		return validateEAs(rec.Ea, expectedRec.Ea)
	}
}

func TestAccResourceSRVRecord(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckSRVRecordDestroy,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
					resource "infoblox_srv_record" "foo"{
						name = "_sip._tcp.example1.com"
						priority = 50
						weight = 30
						port = 80
						target = "sample.target1.com"
					}`),
				Check: resource.ComposeTestCheckFunc(
					testAccSRVRecordCompare(t, "infoblox_srv_record.foo", &ibclient.RecordSRV{
						Name:     "_sip._tcp.example1.com",
						Priority: 50,
						Weight:   30,
						Port:     80,
						Target:   "sample.target1.com",
						View:     "default",
						Ttl:      0,
						UseTtl:   false,
						Comment:  "",
						Ea:       nil,
					}),
				),
			},
			{
				Config: fmt.Sprintf(`
					resource "infoblox_srv_record" "foo1" {
						name = "_sip._udp.example2.com"
						priority = 60
						weight = 40
						port = 36
						target = "sample.target2.com"
						ttl = 300 //300s
						dns_view = "nondefault_view"
						comment = "test comment 1"
						extattrs = jsonencode({
							"Location" = "France"
							"Site" = "DHQ"
						})
					}`),
				Check: resource.ComposeTestCheckFunc(
					testAccSRVRecordCompare(t, "infoblox_srv_record.foo1", &ibclient.RecordSRV{
						Name:     "_sip._udp.example2.com",
						Priority: 60,
						View:     "nondefault_view",
						Weight:   40,
						Port:     36,
						Target:   "sample.target2.com",
						Ttl:      300,
						UseTtl:   true,
						Comment:  "test comment 1",
						Ea: ibclient.EA{
							"Location": "France",
							"Site":     "DHQ",
						},
					}),
				),
			},
			{
				Config: fmt.Sprintf(`
					resource "infoblox_srv_record" "foo2"{
						name = "_http._tcp.demo.example3.com"
						priority = 100
						dns_view = "nondefault_view"
						weight = 50
						port = 88
						target = "sample.target3.com"
						ttl = 140
						comment = "test comment 2"
					}`),
				Check: resource.ComposeTestCheckFunc(
					testAccSRVRecordCompare(t, "infoblox_srv_record.foo2", &ibclient.RecordSRV{
						Name:     "_http._tcp.demo.example3.com",
						Priority: 100,
						View:     "nondefault_view",
						Weight:   50,
						Port:     88,
						Target:   "sample.target3.com",
						Ttl:      140,
						Comment:  "test comment 2",
					}),
				),
			},
			{
				Config: fmt.Sprintf(`
					resource "infoblox_srv_record" "foo3"{
						name = "_http._udp.demo.example3.com"
						dns_view = "nondefault_view"
						priority = 120
						weight = 80
						port = 8080
						target = "demo.target3.com"
					}`),
				Check: resource.ComposeTestCheckFunc(
					testAccSRVRecordCompare(t, "infoblox_srv_record.foo3", &ibclient.RecordSRV{
						Name:     "_http._udp.demo.example3.com",
						View:     "nondefault_view",
						Priority: 120,
						Weight:   80,
						Port:     8080,
						Target:   "demo.target3.com",
					}),
				),
			},
		},
	})
}
