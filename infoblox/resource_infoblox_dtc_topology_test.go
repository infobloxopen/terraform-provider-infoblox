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
			dest_type   = "POOL"
			destination = "dtc:pool/ZG5zLmlkbnNfcG9vbCRQb29sNjM:Pool63"   // TODO: static test data: Pool
			return_type = "REGULAR"
			sources     {
				source_op    = "IS"
				source_type  = "SUBNET"
				source_value = "12.13.14.0/24"
			}
		}
}`

var testResourceDtcTopology3 = `resource "infoblox_dtc_topology" "test-topology3" {
		name = ""
}`

var testResourceDtcPoolId = "dtc:pool/ZG5zLmlkbnNfcG9vbCRQb29sNjM:Pool63"

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
		recJson, _ := json.Marshal(dtcTopology)
		var rec *ibclient.DtcTopology
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
			if len(rec.Rules) != len(expectedTopology.Rules) {
				return fmt.Errorf("the length of 'rules' is %d but expected %d", len(rec.Rules), len(expectedTopology.Rules))
			}
			for idx, have_rule := range rec.Rules {
				expected_rule := expectedTopology.Rules[idx]
				if have_rule.DestType != expected_rule.DestType ||
					have_rule.ReturnType != expected_rule.ReturnType ||
					!reflect.DeepEqual(have_rule.Sources, expected_rule.Sources) {

					// TODO: pretty printing of pointer fields, currently *string and []*Source  are printed as pointers
					return fmt.Errorf("difference found in 'rules' at index %d. got %+v expected %+v",
						idx, have_rule, expected_rule)
				}
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
						DestType:        "POOL",
						DestinationLink: &testResourceDtcPoolId,
						ReturnType:      "REGULAR",
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
