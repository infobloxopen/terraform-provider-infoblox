# SRV-record Data Source

Use the data source to retrieve the following information for SRV-record from the corresponding object in NIOS:

* `dns_view`: the DNS view which the record's zone belongs to.
* `name`: the record's name in the format, defined in RFC2782 document. Example: `_http._tcp.acme.com`
* `target`: an FQDN of the host which is responsible for providing the service specified by `name`. Example: `www.acme.com`
* `port`: a port number (0..65535) on the `target` host which the service expects requests on.
* `priority`: a priority number, as described in RFC2782.
* `weight`: a weight number, as described in RFC2782.
* `zone`: the zone which the record belongs to.
* `ttl`: the "time to live" value of the record, in seconds. Example: `1800`.
* `comment`: the description of the record. This is a regular comment. Example: `spare node for the service`.
* `ext_attrs`: the set of extensible attributes of the record, if any. The content is formatted as string of JSON map. Example: `"{\"Owner\":\"State Library\", \"Expires\":\"never\"}"`.

For usage of filters, add the fields as keys and appropriate values to be passed to the keys like `name`, `view` corresponding to object.
From the below list of supported arguments for filters,  use only the searchable fields for retriving the matching records.

### Supported Arguments for filters

-----

| Field    | Alias    | Type   | Searchable |
|----------|----------|--------|------------|
| name     | fqdn     | string | yes        |
| priority | priority | uint32 | yes        |
| view     | dns_view | string | yes        |
| weight   | weight   | uint32 | yes        |
| port     | port     | uint32 | yes        |
| target   | target   | string | yes        |
| ttl      | ttl      | uint32 | no         |
| comment  | comment  | string | yes        |
| zone     | zone     | string | yes        |

!> Any of the combination from searchable fields in supported arguments list for fields are allowed.

!> Please consider using only fields as the keys in terraform datasource filters, kindly don't use alias names as keys from the above table.

### Example for using the filters:
 ```hcl
 data "infoblox_srv_record" "srv_rec_filter" {
    filters = {
        name = "testsrv.demo.com"
        view = "nondefault_dnsview" // associated DNS view
        priority = 15
    }
 }
 ```

!> From the above example, if the 'view' alias 'dns_view' value is not specified, if same record exists in one or more different DNS views, those
all records will be fetched in results.

!> If `null` or empty filters are passed, then all the records or objects associated with datasource like here `infoblox_srv_record` will be fetched in results.

### Example of the SRV-record Data Source Block

```hcl
resource "infoblox_srv_record" "rec2" {
    dns_view = "nondefault_dnsview1"
    name = "_sip._udp.example2.org"
    priority = 12
    weight = 10
    port = 5060
    target = "sip.example2.org"
    ttl = 3600
    comment = "example SRV record"
    ext_attrs = jsonencode({
        "Location" = "65.8665701230204, -37.00791763398113"
    })
}

data "infoblox_srv_record" "ds1" {
    filters = {
      dns_view = "nondefault_dnsview1"
      name = "_sip._udp.example2.org"
      port = 5060
      target = "sip.example2.org"
    }

  // This is just to ensure that the record has been be created
  // using 'infoblox_srv_record' resource block before the data source will be queried.
    depends_on = [infoblox_srv_record.rec2]
}

output "srv_rec_res" {
  value = data.infoblox_srv_record.ds1
}

// accessing individual field in results
output "srv_rec_name" {
  value = data.infoblox_srv_record.ds1.results.0.name //zero represents index of json object from results list
}

// accessing SRV-Record through EA's
data "infoblox_srv_record" "srv_rec_ea" {
  filters = {
    "*Owner" = "State Library"
    "*Expires" = "never"
  }
}

// throws matching SRV-Records with EA, if any
output "srv_rec_out" {
  value = data.infoblox_srv_record.srv_rec_ea
}
```
