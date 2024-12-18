package infoblox

import (
	"encoding/json"
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	ibclient "github.com/infobloxopen/infoblox-go-client/v2"
	"github.com/infobloxopen/infoblox-go-client/v2/utils"
	"testing"
)

func testDtcPoolDestroy(s *terraform.State) error {
	meta := testAccProvider.Meta()

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "resource_dtc_pool" {
			continue
		}
		connector := meta.(ibclient.IBConnector)
		objMgr := ibclient.NewObjectManager(connector, "terraform_test", "test")
		rec, _ := objMgr.GetDtcPoolByRef(rs.Primary.ID)
		if rec != nil {
			return fmt.Errorf("pool not found")
		}
	}
	return nil
}

func testAccDtcPoolCompare(
	t *testing.T,
	resPath string,
	expectedRec *ibclient.DtcPool) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		res, found := s.RootModule().Resources[resPath]
		if !found {
			return fmt.Errorf("not found: %s", resPath)
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

		dtcPool, err := objMgr.SearchObjectByAltId("DtcPool", ref, internalId, eaNameForInternalId)
		if err != nil {
			if isNotFoundError(err) {
				if expectedRec == nil {
					return nil
				}
				return fmt.Errorf("object with Terraform ID '%s' not found, but expected to exist", internalId)
			}
		}
		// Assertion of object type and error handling
		var rec *ibclient.DtcPool
		recJson, _ := json.Marshal(dtcPool)
		err = json.Unmarshal(recJson, &rec)

		if rec.Name == nil {
			return fmt.Errorf("'name' is expected to be defined but it is not ")
		}
		if *rec.Name != *expectedRec.Name {
			return fmt.Errorf("'fqdn' does not match: got '%s', expected '%s'", *rec.Name, *expectedRec.Name)
		}
		if rec.LbPreferredMethod == "" {
			return fmt.Errorf("'lb_preferred_method' is expected to be defined but it is not")
		}
		if rec.LbPreferredMethod != expectedRec.LbPreferredMethod {
			return fmt.Errorf("'lb_preferred_method' does not match: got '%s', expected '%s'", rec.LbPreferredMethod, expectedRec.LbPreferredMethod)
		}
		if rec.Comment != nil {
			if expectedRec.Comment == nil {
				return fmt.Errorf("'comment' is expected to be undefined but it is not")
			}
			if *rec.Comment != *expectedRec.Comment {
				return fmt.Errorf(
					"'comment' does not match: got '%s', expected '%s'",
					*rec.Comment, *expectedRec.Comment)
			}
		} else if expectedRec.Comment != nil {
			return fmt.Errorf("'comment' is expected to be defined but it is not")
		}
		return validateEAs(rec.Ea, expectedRec.Ea)
	}
}

//var regexRequiredMissingPool = regexp.MustCompile("name and lbPreferredMethod must be provided to create a pool")

func TestAccResourceDtcPool(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testDtcPoolDestroy,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
					resource "infoblox_dtc_pool" "pool1" {
						name                 = "dtc_pool"
						comment              = "pool creation"
						lb_preferred_method  = "ROUND_ROBIN"
					}
				`),
				Check: resource.ComposeTestCheckFunc(
					testAccDtcPoolCompare(t, "infoblox_dtc_pool.pool1", &ibclient.DtcPool{
						Name:              utils.StringPtr("dtc_pool"),
						Comment:           utils.StringPtr("pool creation"),
						LbPreferredMethod: "ROUND_ROBIN",
					}),
				),
			},
			{
				Config: fmt.Sprintf(`
				resource "infoblox_dtc_pool" "pool2"{
				name = "dtc_pool2"
				comment="pool creation"
				monitors{
					monitor_name = "snmp"
     				monitor_type="snmp"
				}
				lb_preferred_method= "DYNAMIC_RATIO"
				lb_dynamic_ratio_preferred=  jsonencode({
						"monitor_name"="snmp"
						"monitor_type"="snmp"
						"method"="MONITOR"
						"monitor_metric"=".1.2"
						"monitor_weighing"="PRIORITY"
						"invert_monitor_metric"=true
				})
}`),
				Check: resource.ComposeTestCheckFunc(
					testAccDtcPoolCompare(t, "infoblox_dtc_pool.pool2", &ibclient.DtcPool{
						Name:              utils.StringPtr("dtc_pool2"),
						Comment:           utils.StringPtr("pool creation"),
						LbPreferredMethod: "DYNAMIC_RATIO",
						Monitors: []*ibclient.DtcMonitorHttp{
							{
								Ref: "dtc:monitor:snmp/ZG5zLmlkbnNfbW9uaXRvcl9zbm1wJHNubXA:snmp",
							},
						},
						LbDynamicRatioPreferred: &ibclient.SettingDynamicratio{
							Monitor:             "dtc:monitor:snmp/ZG5zLmlkbnNfbW9uaXRvcl9zbm1wJHNubXA:snmp",
							Method:              "MONITOR",
							MonitorMetric:       ".1.2",
							MonitorWeighing:     "PRIORITY",
							InvertMonitorMetric: true,
						},
					})),
			}, {
				Config: fmt.Sprintf(`
					resource "infoblox_dtc_pool" "pool3"{
					name = "dtc_pool3"
					comment = "pool creation"
					lb_preferred_method= TOPOLOGY
					lb_preferred_topology= "topology_ruleset"
					lb_alternate_method = "DYNAMIC_RATIO"
					lb_dynamic_ratio_alternate =jsonencode({
						"monitor_name"="snmp"
						"monitor_type"="snmp"
						"method"="MONITOR"
						"monitor_metric"=".1.2"
						"monitor_weighing"="PRIORITY"
						"invert_monitor_metric"=true
				})
				monitors{
						monitor_name = "snmp"
     					monitor_type="snmp"
				}	
				}`),
			},
		},
	})
}
