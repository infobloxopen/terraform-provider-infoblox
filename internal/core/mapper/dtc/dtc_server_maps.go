package dtc

import "github.com/infobloxopen/terraform-provider-infoblox/internal/core"

// DtcServerNIOSFieldMap maps infoblox model fields to NIOS struct fields
var DtcServerNIOSFieldMap = map[string]string{
	"Id":                        "Ref",
	"NIOS.AutoCreateHostRecord": "AutoCreateHostRecord",
	"NIOS.Comment":              "Comment",
	"NIOS.Disable":              "Disable",
	"NIOS.Host":                 "Host",
	"NIOS.Monitors":             "Monitors",
	"NIOS.Name":                 "Name",
	"NIOS.SniHostname":          "SniHostname",
	"NIOS.UseSniHostname":       "UseSniHostname",
}

// DtcServerUDDIFieldMap maps infoblox model fields to UDDI struct fields
var DtcServerUDDIFieldMap = map[string]string{
	"UDDI.Address":                   "Address",
	"UDDI.AutoCreateResponseRecords": "AutoCreateResponseRecords",
	"UDDI.Comment":                   "Comment",
	"UDDI.Disabled":                  "Disabled",
	"UDDI.EndpointType":              "EndpointType",
	"UDDI.Fqdn":                      "Fqdn",
	"UDDI.Name":                      "Name",
	"UDDI.Records":                   "Records",
	"UDDI.Tags":                      "Tags",
}

// TODO: only searchable fields should be included here
// DtcServerFilterFieldMap maps infoblox filter keys to backend-specific API filter field names
var DtcServerFilterFieldMap = map[core.BackendType]map[string]string{
	core.BackendNIOS: {
		"id":                           "_ref",
		"nios.auto_create_host_record": "auto_create_host_record",
		"nios.comment":                 "comment",
		"nios.disable":                 "disable",
		"nios.ext_attrs":               "extattrs",
		"nios.host":                    "host",
		"nios.monitors":                "monitors",
		"nios.name":                    "name",
		"nios.sni_hostname":            "sni_hostname",
		"nios.use_sni_hostname":        "use_sni_hostname",
	},
	core.BackendUDDI: {
		"uddi.address":                      "address",
		"uddi.auto_create_response_records": "auto_create_response_records",
		"uddi.comment":                      "comment",
		"uddi.disabled":                     "disabled",
		"uddi.endpoint_type":                "endpoint_type",
		"uddi.fqdn":                         "fqdn",
		"uddi.name":                         "name",
		"uddi.records":                      "records",
		"uddi.tags":                         "tags",
	},
}
