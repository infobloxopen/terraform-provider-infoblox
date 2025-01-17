package infoblox

import (
	"encoding/json"
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	ibclient "github.com/infobloxopen/infoblox-go-client/v2"
	"github.com/infobloxopen/infoblox-go-client/v2/utils"
	"regexp"
	"testing"
)

var testResourceDtcServer = `resource "infoblox_dtc_server" "test-server1" {
		name = "test-server1"
		host = "12.12.1.1"
}`

var testResourceDtcServer2 = `resource "infoblox_dtc_server" "test-server2" {
		name = "test-server2"
		host = "abc.com"
		comment = "test dtc server with max params"
		disable = true
		auto_create_host_record = false
		use_sni_hostname = true
		sni_hostname = "sni_name"
		ext_attrs = jsonencode({"Site" = "India"})
		monitors {
			monitor_name = "snmp"
			host = "12.13.14.15"
			monitor_type = "snmp"
		}
}`

var testResourceDtcServer3 = `resource "infoblox_dtc_server" "test-server3" {
		name = "test-server3"
		host = ""
}`

func testDtcServerCompare(t *testing.T, resourceName string, expectedServer *ibclient.DtcServer) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		res, found := s.RootModule().Resources[resourceName]
		if !found {
			return fmt.Errorf("not found: %s", resourceName)
		}
		if res.Primary.ID == "" {
			return fmt.Errorf("ID is not set")
		}

		internalId := res.Primary.Attributes["internal_id"]
		if internalId == "" {
			return fmt.Errorf("ID is not set")
		}

		ref, found := res.Primary.Attributes["ref"]
		if !found {
			return fmt.Errorf("'ref' attribute is not set")
		}

		connector := testAccProvider.Meta().(ibclient.IBConnector)
		objMgr := ibclient.NewObjectManager(
			connector,
			"terraform_test",
			"test")
		dtcServer, err := objMgr.SearchObjectByAltId("DtcServer", ref, internalId, eaNameForInternalId)
		if err != nil {
			if isNotFoundError(err) {
				if expectedServer == nil {
					return nil
				}
				return fmt.Errorf("object with Terraform ID '%s' not found, but expected to exist", internalId)
			}
		}
		// Assertion of object type and error handling
		var rec *ibclient.DtcServer
		recJson, _ := json.Marshal(dtcServer)
		err = json.Unmarshal(recJson, &rec)

		if *rec.Name != *expectedServer.Name {
			return fmt.Errorf(
				"the value of 'name' field is '%s', but expected '%s'",
				*rec.Name, *expectedServer.Name)
		}
		if rec.Comment != nil && expectedServer.Comment != nil {
			if *rec.Comment != *expectedServer.Comment {
				return fmt.Errorf(
					"the value of 'comment' field is '%s', but expected '%s'",
					*rec.Comment, *expectedServer.Comment)
			}
		}
		if rec.Disable != nil && expectedServer.Disable != nil {
			if *rec.Disable != *expectedServer.Disable {
				return fmt.Errorf(
					"the value of 'disable' field is '%t', but expected '%t'",
					*rec.Disable, *expectedServer.Disable)
			}
		}
		if rec.AutoCreateHostRecord != nil && expectedServer.AutoCreateHostRecord != nil {
			if *rec.AutoCreateHostRecord != *expectedServer.AutoCreateHostRecord {
				return fmt.Errorf("the value of 'auto_create_host_record' field is '%t', but expected '%t'", *rec.AutoCreateHostRecord, *expectedServer.AutoCreateHostRecord)
			}
		}
		if rec.Host != nil && expectedServer.Host != nil {
			if *rec.Host != *expectedServer.Host {
				return fmt.Errorf(
					"the value of 'host' field is '%s', but expected '%s'",
					*rec.Host, *expectedServer.Host)
			}
		}
		if rec.SniHostname != nil && expectedServer.SniHostname != nil {
			if *rec.SniHostname != *expectedServer.SniHostname {
				return fmt.Errorf(
					"the value of 'sni_hostname' field is '%s', but expected '%s'",
					*rec.SniHostname, *expectedServer.SniHostname)
			}
		}
		if rec.UseSniHostname != nil && expectedServer.UseSniHostname != nil {
			if *rec.UseSniHostname != *expectedServer.UseSniHostname {
				return fmt.Errorf(
					"the value of 'use_sni_hostname' field is '%t', but expected '%t'",
					*rec.UseSniHostname, *expectedServer.UseSniHostname)
			}
		}
		if rec.Monitors != nil && expectedServer.Monitors != nil {
			if !compareMonitors(rec.Monitors, expectedServer.Monitors, connector) {
				return fmt.Errorf("the value of 'monitors' field is '%v', but expected '%v'", rec.Monitors, expectedServer.Monitors)
			}
		}
		return validateEAs(rec.Ea, expectedServer.Ea)
	}
}

func compareMonitors(monitors1, monitors2 []*ibclient.DtcServerMonitor, connector ibclient.IBConnector) bool {
	if len(monitors1) != len(monitors2) {
		return false
	}
	for i := range monitors1 {
		if monitors1[i].Monitor != monitors2[i].Monitor || monitors1[i].Host != monitors2[i].Host {
			return false
		}
	}
	return true
}

func testDtcdtcServerDestroy(s *terraform.State) error {
	meta := testAccProvider.Meta()

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "resource_dtc_server" {
			continue
		}
		connector := meta.(ibclient.IBConnector)
		objMgr := ibclient.NewObjectManager(connector, "terraform_test", "test")
		rec, _ := objMgr.GetDtcServerByRef(rs.Primary.ID)
		if rec != nil {
			return fmt.Errorf("record not found")
		}
	}
	return nil
}

func TestAccResourceDtcServer(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testDtcdtcServerDestroy,
		Steps: []resource.TestStep{
			// minimum params
			{
				Config: testResourceDtcServer,
				Check: testDtcServerCompare(t, "infoblox_dtc_server.test-server1", &ibclient.DtcServer{
					Name: utils.StringPtr("test-server1"),
					Host: utils.StringPtr("12.12.1.1"),
				}),
			},
			// maximum params
			{
				Config: testResourceDtcServer2,
				Check: testDtcServerCompare(t, "infoblox_dtc_server.test-server2", &ibclient.DtcServer{
					Name:                 utils.StringPtr("test-server2"),
					Comment:              utils.StringPtr("test dtc server with max params"),
					Host:                 utils.StringPtr("abc.com"),
					Disable:              utils.BoolPtr(true),
					Ea:                   map[string]interface{}{"Site": "India"},
					Monitors:             []*ibclient.DtcServerMonitor{{Monitor: "dtc:monitor:snmp/ZG5zLmlkbnNfbW9uaXRvcl9zbm1wJHNubXA:snmp", Host: "12.13.14.15"}},
					SniHostname:          utils.StringPtr("sni_name"),
					UseSniHostname:       utils.BoolPtr(true),
					AutoCreateHostRecord: utils.BoolPtr(false),
				}),
			},
			// negative test case
			{
				Config:      testResourceDtcServer3,
				ExpectError: regexp.MustCompile("name and host fields are required to create a dtc server"),
			},
		},
	})
}
