package infoblox

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	ibclient "github.com/infobloxopen/infoblox-go-client/v2"
	"github.com/infobloxopen/infoblox-go-client/v2/utils"
)

var testResourceDtcMonitorHttp = `
	resource "infoblox_dtc_monitor_http" "test-monitor_http1" {
		name = "test-monitor_http1"
		port = 81
	}
`

var testResourceDtcMonitorHttp2 = `
	resource "infoblox_dtc_monitor_http" "test-monitor_http2" {
		name       = "test-monitor_http2"
		comment    = "test dtc monitor_http with max params"
		ext_attrs  = jsonencode({ "Site" = "India" })
		interval   = 7
		port       = 8080
		retry_down = 2
		retry_up   = 2
		timeout    = 5
		ciphers    = "AES256-GCM-SHA384"
		// client_cert = "dtc:certificate object reference"
		content_check         = "EXTRACT"
		content_check_input   = "BODY"
		content_check_op      = "EQ"
		content_check_regex   = "healthy: ([0-9]+)"
		content_extract_group = 1
		content_extract_type  = "STRING"
		content_extract_value = "all systems healthy"
		enable_sni            = true
		request               = "GET / HTTP/1.1\nConnection: close\n\n"
		result                = "CODE_IS"
		result_code           = 204
		secure                = true
		validate_cert         = true
	}
`

// fill all fields with golang default values with API default values (as defined in the resource schema)
func setDtcMonitorHttpObjectDefaults(obj *ibclient.DtcMonitorHttp) *ibclient.DtcMonitorHttp {
	if obj.ContentCheck == "" {
		obj.ContentCheck = "NONE"
	}
	if obj.ContentCheckInput == "" {
		obj.ContentCheckInput = "ALL"
	}
	if obj.ContentExtractType == "" {
		obj.ContentExtractType = "STRING"
	}
	if obj.Request != nil && *obj.Request == "" {
		obj.Request = utils.StringPtr("GET /\n\n")
	}
	if obj.Result == "" {
		obj.Result = "ANY"
	}
	if obj.ResultCode != nil && *obj.ResultCode == 0 {
		obj.ResultCode = utils.Uint32Ptr(200)
	}
	return obj
}

func testDtcMonitorHttpCompare(t *testing.T, resourceName string, expectedMonitorHttp *ibclient.DtcMonitorHttp) resource.TestCheckFunc {
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
		dtcMonitorHttp, err := objMgr.SearchObjectByAltId("DtcMonitorHttp", ref, internalId, eaNameForInternalId)
		if err != nil {
			if isNotFoundError(err) {
				if expectedMonitorHttp == nil {
					return nil
				}
				return fmt.Errorf("object with Terraform ID '%s' not found, but expected to exist", internalId)
			}
		}
		// Assertion of object type and error handling
		var rec *ibclient.DtcMonitorHttp
		recJson, _ := json.Marshal(dtcMonitorHttp)
		err = json.Unmarshal(recJson, &rec)

		if rec.Ciphers != nil && expectedMonitorHttp.Ciphers != nil {
			if *rec.Ciphers != *expectedMonitorHttp.Ciphers {
				return fmt.Errorf(
					"the value of 'ciphers' field is '%s', but expected '%s'",
					*rec.Ciphers, *expectedMonitorHttp.Ciphers)
			}
		}

		if rec.ClientCert != nil && expectedMonitorHttp.ClientCert != nil {
			if *rec.ClientCert != *expectedMonitorHttp.ClientCert {
				return fmt.Errorf(
					"the value of 'client_cert' field is '%s', but expected '%s'",
					*rec.ClientCert, *expectedMonitorHttp.ClientCert)
			}
		}

		if rec.Comment != nil && expectedMonitorHttp.Comment != nil {
			if *rec.Comment != *expectedMonitorHttp.Comment {
				return fmt.Errorf(
					"the value of 'comment' field is '%s', but expected '%s'",
					*rec.Comment, *expectedMonitorHttp.Comment)
			}
		}

		if rec.ContentCheck != expectedMonitorHttp.ContentCheck {
			return fmt.Errorf(
				"the value of 'content_check' field is '%s', but expected '%s'",
				rec.ContentCheck, expectedMonitorHttp.ContentCheck)
		}

		if rec.ContentCheckInput != expectedMonitorHttp.ContentCheckInput {
			return fmt.Errorf(
				"the value of 'content_check_input' field is '%s', but expected '%s'",
				rec.ContentCheckInput, expectedMonitorHttp.ContentCheckInput)
		}

		if rec.ContentCheckOp != expectedMonitorHttp.ContentCheckOp {
			return fmt.Errorf(
				"the value of 'content_check_op' field is '%s', but expected '%s'",
				rec.ContentCheckOp, expectedMonitorHttp.ContentCheckOp)
		}

		if rec.ContentCheckRegex != nil && expectedMonitorHttp.ContentCheckRegex != nil {
			if *rec.ContentCheckRegex != *expectedMonitorHttp.ContentCheckRegex {
				return fmt.Errorf(
					"the value of 'content_check_regex' field is '%s', but expected '%s'",
					*rec.ContentCheckRegex, *expectedMonitorHttp.ContentCheckRegex)
			}
		}

		if rec.ContentExtractGroup != nil && expectedMonitorHttp.ContentExtractGroup != nil {
			if *rec.ContentExtractGroup != *expectedMonitorHttp.ContentExtractGroup {
				return fmt.Errorf(
					"the value of 'content_extract_group' field is '%d', but expected '%d'",
					*rec.ContentExtractGroup, *expectedMonitorHttp.ContentExtractGroup)
			}
		}

		if rec.ContentExtractType != expectedMonitorHttp.ContentExtractType {
			return fmt.Errorf(
				"the value of 'content_extract_type' field is '%s', but expected '%s'",
				rec.ContentExtractType, expectedMonitorHttp.ContentExtractType)
		}

		if rec.ContentExtractValue != nil && expectedMonitorHttp.ContentExtractValue != nil {
			if *rec.ContentExtractValue != *expectedMonitorHttp.ContentExtractValue {
				return fmt.Errorf(
					"the value of 'content_extract_value' field is '%s', but expected '%s'",
					*rec.ContentExtractValue, *expectedMonitorHttp.ContentExtractValue)
			}
		}

		if rec.EnableSni != nil && expectedMonitorHttp.EnableSni != nil {
			if *rec.EnableSni != *expectedMonitorHttp.EnableSni {
				return fmt.Errorf(
					"the value of 'enable_sni' field is '%t', but expected '%t'",
					*rec.EnableSni, *expectedMonitorHttp.EnableSni)
			}
		}

		if rec.Interval != nil && expectedMonitorHttp.Interval != nil {
			if *rec.Interval != *expectedMonitorHttp.Interval {
				return fmt.Errorf(
					"the value of 'interval' field is '%d', but expected '%d'",
					*rec.Interval, *expectedMonitorHttp.Interval)
			}
		}

		if rec.Name != nil && expectedMonitorHttp.Name != nil {
			if *rec.Name != *expectedMonitorHttp.Name {
				return fmt.Errorf(
					"the value of 'name' field is '%s', but expected '%s'",
					*rec.Name, *expectedMonitorHttp.Name)
			}
		}

		if rec.Port != nil && expectedMonitorHttp.Port != nil {
			if *rec.Port != *expectedMonitorHttp.Port {
				return fmt.Errorf(
					"the value of 'port' field is '%d', but expected '%d'",
					*rec.Port, *expectedMonitorHttp.Port)
			}
		}

		if rec.Request != nil && expectedMonitorHttp.Request != nil {
			if *rec.Request != *expectedMonitorHttp.Request {
				return fmt.Errorf(
					"the value of 'request' field is '%s', but expected '%s'",
					*rec.Request, *expectedMonitorHttp.Request)
			}
		}

		if rec.Result != expectedMonitorHttp.Result {
			return fmt.Errorf(
				"the value of 'result' field is '%s', but expected '%s'",
				rec.Result, expectedMonitorHttp.Result)
		}

		if rec.ResultCode != nil && expectedMonitorHttp.ResultCode != nil {
			if *rec.ResultCode != *expectedMonitorHttp.ResultCode {
				return fmt.Errorf(
					"the value of 'result_code' field is '%d', but expected '%d'",
					*rec.ResultCode, *expectedMonitorHttp.ResultCode)
			}
		}

		if rec.RetryDown != nil && expectedMonitorHttp.RetryDown != nil {
			if *rec.RetryDown != *expectedMonitorHttp.RetryDown {
				return fmt.Errorf(
					"the value of 'retry_down' field is '%d', but expected '%d'",
					*rec.RetryDown, *expectedMonitorHttp.RetryDown)
			}
		}

		if rec.RetryUp != nil && expectedMonitorHttp.RetryUp != nil {
			if *rec.RetryUp != *expectedMonitorHttp.RetryUp {
				return fmt.Errorf(
					"the value of 'retry_up' field is '%d', but expected '%d'",
					*rec.RetryUp, *expectedMonitorHttp.RetryUp)
			}
		}

		if rec.Secure != nil && expectedMonitorHttp.Secure != nil {
			if *rec.Secure != *expectedMonitorHttp.Secure {
				return fmt.Errorf(
					"the value of 'secure' field is '%t', but expected '%t'",
					*rec.Secure, *expectedMonitorHttp.Secure)
			}
		}

		if rec.Timeout != nil && expectedMonitorHttp.Timeout != nil {
			if *rec.Timeout != *expectedMonitorHttp.Timeout {
				return fmt.Errorf(
					"the value of 'timeout' field is '%d', but expected '%d'",
					*rec.Timeout, *expectedMonitorHttp.Timeout)
			}
		}

		if rec.ValidateCert != nil && expectedMonitorHttp.ValidateCert != nil {
			if *rec.ValidateCert != *expectedMonitorHttp.ValidateCert {
				return fmt.Errorf(
					"the value of 'validate_cert' field is '%t', but expected '%t'",
					*rec.ValidateCert, *expectedMonitorHttp.ValidateCert)
			}
		}

		return validateEAs(rec.Ea, expectedMonitorHttp.Ea)
	}
}

func testDtcdtcMonitorHttpDestroy(s *terraform.State) error {
	meta := testAccProvider.Meta()

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "resource_dtc_monitor_http" {
			continue
		}
		connector := meta.(ibclient.IBConnector)
		objMgr := ibclient.NewObjectManager(connector, "terraform_test", "test")
		rec, _ := objMgr.GetDtcMonitorHttpByRef(rs.Primary.ID)
		if rec != nil {
			return fmt.Errorf("record not found")
		}
	}
	return nil
}

func TestAccResourceDtcMonitorHttp(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testDtcdtcMonitorHttpDestroy,
		Steps: []resource.TestStep{
			// minimum params
			{
				Config: testResourceDtcMonitorHttp,
				Check: testDtcMonitorHttpCompare(t,
					"infoblox_dtc_monitor_http.test-monitor_http1",
					setDtcMonitorHttpObjectDefaults(&ibclient.DtcMonitorHttp{
						Name: utils.StringPtr("test-monitor_http1"),
						Port: utils.Uint32Ptr(81),
					})),
			},
			// maximum params
			{
				Config: testResourceDtcMonitorHttp2,
				Check: testDtcMonitorHttpCompare(t,
					"infoblox_dtc_monitor_http.test-monitor_http2",
					setDtcMonitorHttpObjectDefaults(&ibclient.DtcMonitorHttp{
						Name:                utils.StringPtr("test-monitor_http2"),
						Comment:             utils.StringPtr("test dtc monitor_http with max params"),
						Ea:                  map[string]interface{}{"Site": "India"},
						Interval:            utils.Uint32Ptr(7),
						Port:                utils.Uint32Ptr(8080),
						RetryDown:           utils.Uint32Ptr(2),
						RetryUp:             utils.Uint32Ptr(2),
						Timeout:             utils.Uint32Ptr(5),
						Ciphers:             utils.StringPtr("AES256-GCM-SHA384"),
						ClientCert:          utils.StringPtr(""),
						ContentCheck:        "EXTRACT",
						ContentCheckInput:   "BODY",
						ContentCheckOp:      "EQ",
						ContentCheckRegex:   utils.StringPtr("healthy: ([0-9]+)"),
						ContentExtractGroup: utils.Uint32Ptr(1),
						ContentExtractType:  "STRING",
						ContentExtractValue: utils.StringPtr("all systems healthy"),
						EnableSni:           utils.BoolPtr(true),
						Request:             utils.StringPtr("GET / HTTP/1.1\nConnection: close\n\n"),
						Result:              "CODE_IS",
						ResultCode:          utils.Uint32Ptr(204),
						Secure:              utils.BoolPtr(true),
						ValidateCert:        utils.BoolPtr(true),
					})),
			},
		},
	})
}
