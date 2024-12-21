package infoblox

import (
	"encoding/json"
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	ibclient "github.com/infobloxopen/infoblox-go-client/v2"
	"github.com/infobloxopen/infoblox-go-client/v2/utils"
	"reflect"
	"regexp"
	"testing"
)

var testResourceDtcLbdn = `resource "infoblox_dtc_lbdn" "testLbdn1" {
    name = "testLbdn123"
    auth_zones = ["test.com"]
  	comment = "test lbdn with max params"
  	ext_attrs = jsonencode({
    	"Location" = "65.8665701230204, -37.00791763398113"
  	})
  	lb_method = "TOPOLOGY"
  	patterns = ["test.com*", "info.com*"]
  	pools {
    	pool = "pool2"
    	ratio = 2
  	}
  	pools {
    	pool = "rrpool"
    	ratio = 3
  	}
  	pools {
    	pool = "test-pool"
    	ratio = 6
  	}
    topology = "test-topo"
  	ttl = 120
  	disable = true
  	types = ["A", "AAAA", "CNAME"]
  	persistence = 60
  	priority = 1
}`

var testResourceDtcLbdn2 = `resource "infoblox_dtc_lbdn" "testLbdn2" {
    name = "testLbdn456"
  	lb_method = "RATIO"
    types = ["A", "AAAA"]
}`

var testResourceDtcLbdn3 = `resource "infoblox_dtc_lbdn" "testLbdn3" {
    name = "testLbdn789"
  	lb_method = "TOPOLOGY"
}`

func testDtcLbdnDestroy(s *terraform.State) error {
	meta := testAccProvider.Meta()

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "resource_dtc_lbdn" {
			continue
		}
		connector := meta.(ibclient.IBConnector)
		objMgr := ibclient.NewObjectManager(connector, "terraform_test", "test")
		rec, _ := objMgr.GetDtcLbdnByRef(rs.Primary.ID)
		if rec != nil {
			return fmt.Errorf("record not found")
		}
	}
	return nil
}

func testDtcLbdnCompare(t *testing.T, resourceName string, expectedLbdn *ibclient.DtcLbdn) resource.TestCheckFunc {
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
		lbdn, err := objMgr.SearchObjectByAltId("DtcLbdn", ref, internalId, eaNameForInternalId)
		if err != nil {
			if isNotFoundError(err) {
				if expectedLbdn == nil {
					return nil
				}
				return fmt.Errorf("object with Terraform ID '%s' not found, but expected to exist", internalId)
			}
		}
		// Assertion of object type and error handling
		var rec *ibclient.DtcLbdn
		recJson, _ := json.Marshal(lbdn)
		err = json.Unmarshal(recJson, &rec)

		if *rec.Name != *expectedLbdn.Name {
			return fmt.Errorf(
				"the value of 'name' field is '%s', but expected '%s'",
				*rec.Name, *expectedLbdn.Name)
		}
		if rec.Comment != nil && expectedLbdn.Comment != nil {
			if *rec.Comment != *expectedLbdn.Comment {
				return fmt.Errorf(
					"the value of 'comment' field is '%s', but expected '%s'",
					*rec.Comment, *expectedLbdn.Comment)
			}
		}
		if rec.Disable != nil && expectedLbdn.Disable != nil {
			if *rec.Disable != *expectedLbdn.Disable {
				return fmt.Errorf(
					"the value of 'disable' field is '%t', but expected '%t'",
					*rec.Disable, *expectedLbdn.Disable)
			}
		}
		if rec.AuthZones != nil && expectedLbdn.AuthZones != nil {
			if !reflect.DeepEqual(rec.AuthZones, expectedLbdn.AuthZones) {
				return fmt.Errorf("the value of 'auth_zones' field is '%v', but expected '%v'", rec.AuthZones, expectedLbdn.AuthZones)
			}
		}
		if rec.AutoConsolidatedMonitors != nil && expectedLbdn.AutoConsolidatedMonitors != nil {
			if *rec.AutoConsolidatedMonitors != *expectedLbdn.AutoConsolidatedMonitors {
				return fmt.Errorf("the value of 'auto_consolidated_monitors' field is '%t', but expected '%t'", *rec.AutoConsolidatedMonitors, *expectedLbdn.AutoConsolidatedMonitors)
			}
		}
		if rec.LbMethod != expectedLbdn.LbMethod {
			return fmt.Errorf("the value of 'lb_method' field is '%s', but expected '%s'", rec.LbMethod, expectedLbdn.LbMethod)
		}
		if rec.Patterns != nil && expectedLbdn.Patterns != nil {
			if !reflect.DeepEqual(rec.Patterns, expectedLbdn.Patterns) {
				return fmt.Errorf("the value of 'patterns' field is '%v', but expected '%v'", rec.Patterns, expectedLbdn.Patterns)
			}
		}
		if rec.Persistence != nil && expectedLbdn.Persistence != nil {
			if *rec.Persistence != *expectedLbdn.Persistence {
				return fmt.Errorf("the value of 'persistence' field is '%d', but expected '%d'", *rec.Persistence, *expectedLbdn.Persistence)
			}
		}
		if rec.Pools != nil && expectedLbdn.Pools != nil {
			if !comparePools(rec.Pools, expectedLbdn.Pools, connector) {
				return fmt.Errorf("the value of 'pools' field is '%v', but expected '%v'", rec.Pools, expectedLbdn.Pools)
			}
		}
		if rec.Priority != nil && expectedLbdn.Priority != nil {
			if *rec.Priority != *expectedLbdn.Priority {
				return fmt.Errorf("the value of 'priority' field is '%d', but expected '%d'", *rec.Priority, *expectedLbdn.Priority)
			}
		}
		if rec.Topology != nil && expectedLbdn.Topology != nil {
			var topology ibclient.DtcTopology
			err := connector.GetObject(&ibclient.DtcTopology{}, *rec.Topology, nil, &topology)
			if err != nil {
				return fmt.Errorf("error getting topology object: %s", *rec.Topology)
			}
			if *expectedLbdn.Topology != *topology.Name {
				return fmt.Errorf("the value of 'topology' field is '%s', but expected '%s'", *topology.Name, *expectedLbdn.Topology)
			}
		}
		if rec.Types != nil && expectedLbdn.Types != nil {
			if !reflect.DeepEqual(rec.Types, expectedLbdn.Types) {
				return fmt.Errorf("the value of 'types' field is '%v', but expected '%v'", rec.Types, expectedLbdn.Types)
			}
		}
		if rec.Ttl != nil && expectedLbdn.Ttl != nil {
			if *rec.Ttl != *expectedLbdn.Ttl {
				return fmt.Errorf("the value of 'ttl' field is '%d', but expected '%d'", *rec.Ttl, *expectedLbdn.Ttl)
			}
		}
		if rec.UseTtl != nil && expectedLbdn.UseTtl != nil {
			if *rec.UseTtl != *expectedLbdn.UseTtl {
				return fmt.Errorf("the value of 'use_ttl' field is '%t', but expected '%t'", *rec.UseTtl, *expectedLbdn.UseTtl)
			}
		}
		return validateEAs(rec.Ea, expectedLbdn.Ea)
	}
}

func comparePools(pools1, pools2 []*ibclient.DtcPoolLink, connector ibclient.IBConnector) bool {
	if len(pools1) != len(pools2) {
		return false
	}
	for i := range pools1 {
		var pool1 ibclient.DtcPool
		err1 := connector.GetObject(&ibclient.DtcPool{}, pools1[i].Pool, nil, &pool1)
		if err1 != nil {
			return false
		}
		if *pool1.Name != pools2[i].Pool || pools1[i].Ratio != pools2[i].Ratio {
			return false
		}
	}
	return true
}

func TestAccResourceDtcLbdn(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testDtcLbdnDestroy,
		Steps: []resource.TestStep{
			// minimum params
			{
				Config: testResourceDtcLbdn2,
				Check: testDtcLbdnCompare(t, "infoblox_dtc_lbdn.testLbdn2", &ibclient.DtcLbdn{
					Name:     utils.StringPtr("testLbdn456"),
					LbMethod: "RATIO",
					Types:    []string{"A", "AAAA"},
				}),
			},
			// maximum params
			{
				Config: testResourceDtcLbdn,
				Check: testDtcLbdnCompare(t, "infoblox_dtc_lbdn.testLbdn1", &ibclient.DtcLbdn{
					Name:      utils.StringPtr("testLbdn123"),
					LbMethod:  "TOPOLOGY",
					Topology:  utils.StringPtr("test-topo"),
					AuthZones: []*ibclient.ZoneAuth{{Fqdn: "test.com"}},
					Comment:   utils.StringPtr("test lbdn with max params"),
					Types:     []string{"A", "AAAA", "CNAME"},
					Pools: []*ibclient.DtcPoolLink{
						{Pool: "pool2", Ratio: 2},
						{Pool: "rrpool", Ratio: 3},
						{Pool: "test-pool", Ratio: 6}},
					Patterns:    []string{"test.com*", "info.com*"},
					Persistence: utils.Uint32Ptr(60),
					Priority:    utils.Uint32Ptr(1),
					Ttl:         utils.Uint32Ptr(120),
					Disable:     utils.BoolPtr(true),
					Ea:          map[string]interface{}{"Location": "65.8665701230204, -37.00791763398113"},
				}),
			},
			// negative test case
			{
				Config:      testResourceDtcLbdn3,
				ExpectError: regexp.MustCompile("topology field is required when lbMethod is TOPOLOGY"),
			},
		},
	})
}
