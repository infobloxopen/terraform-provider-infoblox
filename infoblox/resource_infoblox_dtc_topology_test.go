package infoblox

import (
	"encoding/json"
	"fmt"
	"reflect"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	ibclient "github.com/infobloxopen/infoblox-go-client/v2"
	"github.com/infobloxopen/infoblox-go-client/v2/utils"
)

var testResourceDtcTopology = `resource "infoblox_dtc_topology" "test-topology1" {
		name = "test-topology1"
}`

var testResourceDtcTopology2 = `resource "infoblox_dtc_topology" "test-topology2" {
		name = "test-topology2"
		comment = "test dtc topology with max params"
		ext_attrs = jsonencode({"Site" = "India"})
		rules {
			dest_type   = "SERVER"
			return_type = "REGULAR"
			sources     = [{
				source_op    = "IS"
				source_type  = "SUBNET"
				source_value = "12.13.14.0/24"
			}]
		}
}`

var testResourceDtcTopology3 = `resource "infoblox_dtc_topology" "test-topology3" {
		name = ""
}`

func compareRules(have []*ibclient.DtcTopologyRule, expected []*ibclient.DtcTopologyRule) bool {
	if len(have) != len(expected) {
		return false
	}
	for i := range have {
		have_i := have[i]
		expected_i := expected[i]
		if have_i.DestType != expected_i.DestType {
			return false
		}
		if have_i.ReturnType != expected_i.ReturnType {
			return false
		}
		if have_i.Valid != expected_i.Valid {
			return false
		}
		if !reflect.DeepEqual(have_i.Sources, expected_i.Sources) {
			return false
		}
	}
	return true
}

func testDtcTopologyCompare(t *testing.T, resourceName string, expectedTopology *ibclient.DtcTopology) resource.TestCheckFunc {
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
		dtcTopology, err := objMgr.SearchObjectByAltId("DtcTopology", ref, internalId, eaNameForInternalId)
		if err != nil {
			if isNotFoundError(err) {
				if expectedTopology == nil {
					return nil
				}
				return fmt.Errorf("object with Terraform ID '%s' not found, but expected to exist", internalId)
			}
		}
		// Assertion of object type and error handling
		var rec *ibclient.DtcTopology
		recJson, _ := json.Marshal(dtcTopology)
		err = json.Unmarshal(recJson, &rec)

		if *rec.Name != *expectedTopology.Name {
			return fmt.Errorf(
				"the value of 'name' field is '%s', but expected '%s'",
				*rec.Name, *expectedTopology.Name)
		}
		if rec.Comment != nil && expectedTopology.Comment != nil {
			if *rec.Comment != *expectedTopology.Comment {
				return fmt.Errorf(
					"the value of 'comment' field is '%s', but expected '%s'",
					*rec.Comment, *expectedTopology.Comment)
			}
		}
		if rec.Rules != nil && expectedTopology.Rules != nil {
			if !compareRules(rec.Rules, expectedTopology.Rules) {
				return fmt.Errorf("the value of 'rules' field is '%v', but expected '%v'", rec.Rules, expectedTopology.Rules)
			}
		}
		return validateEAs(rec.Ea, expectedTopology.Ea)
	}
}

func testDtcdtcTopologyDestroy(s *terraform.State) error {
	meta := testAccProvider.Meta()

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "resource_dtc_topology" {
			continue
		}
		connector := meta.(ibclient.IBConnector)
		objMgr := ibclient.NewObjectManager(connector, "terraform_test", "test")
		rec, _ := objMgr.GetDtcTopologyByRef(rs.Primary.ID)
		if rec != nil {
			return fmt.Errorf("record not found")
		}
	}
	return nil
}

func TestAccResourceDtcTopology(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testDtcdtcTopologyDestroy,
		Steps: []resource.TestStep{
			// minimum params
			{
				Config: testResourceDtcTopology,
				Check: testDtcTopologyCompare(t, "infoblox_dtc_topology.test-topology1", &ibclient.DtcTopology{
					Name: utils.StringPtr("test-topology1"),
				}),
			},
			// maximum params
			{
				Config: testResourceDtcTopology2,
				Check: testDtcTopologyCompare(t, "infoblox_dtc_topology.test-topology2", &ibclient.DtcTopology{
					Name:    utils.StringPtr("test-topology2"),
					Comment: utils.StringPtr("test dtc topology with max params"),
					Ea:      map[string]interface{}{"Site": "India"},
					Rules: []*ibclient.DtcTopologyRule{{
						DestType:   "SERVER",
						ReturnType: "REGULAR",
						Sources: []*ibclient.DtcTopologyRuleSource{{
							SourceOp:    "IS",
							SourceType:  "SUBNET",
							SourceValue: "12.13.14.0/24",
						}},
					}},
				}),
			},
			// negative test case
			{
				Config:      testResourceDtcTopology3,
				ExpectError: regexp.MustCompile("name field is required to create a Dtc Topology object"),
			},
		},
	})
}
