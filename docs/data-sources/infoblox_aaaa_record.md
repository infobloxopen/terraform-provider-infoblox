# AAAA-record Data Source

Use the `infoblox_aaaa_record` data source to retrieve the following information for list of AAAA-records if any, which are managed by a NIOS server:

* `dns_view`: the DNS view which the record's zone belongs to. Example: `nondefault_dnsview`
* `ipv6_addr`: the IPv6 address associated with the AAAA-record. Example: `2001::14`
* `fqdn`: the fully qualified domain name which the IP address is assigned to. Example: `foo1.test.com`
* `zone`: the zone that contains the record in the specified DNS view. Example: `test.com`.
* `ttl`: the "time to live" value of the record, in seconds. Example: `1800`.
* `comment`: the description of the record. This is a regular comment. Example: `Temporary AAAA-record`.
* `ext_attrs`: the set of extensible attributes of the record, if any. The content is formatted as string of JSON map. Example: `"{\"TestEA\":56,\"TestEA1\":\"kickoff\"}"`.

As there is new feature filters , the previous usage of combination of DNS view, IPv6 address and FQDN, has been removed.

For usage of filters, add the fields as keys and appropriate values to be passed to the keys like `name`, `view` corresponding to object. Only searchable fields
from below list of supported arguments for filters, are allowed to use in filters, for retrieving one or more records or objects matching
filters.

### Supported Arguments for filters

-----
| Field    | Alias    | Type   | Searchable |
|----------|----------|--------|------------|
| name     | fqdn     | string | yes        |
| view     | dns_view | string | yes        |
| zone     | ---      | string | yes        |
| ttl      | ---      | uint   | no         |
| comment  | ---      | string | yes        |
| ipv6addr | ip_addr  | string | yes        |

!> Any of the combination from searchable fields in supported arguments list for fields are allowed.

!> Please consider using only fields as the keys in terraform datasource filters, kindly don't use alias names as keys from the above table.

### Example for using the filters:
 ```hcl
 data "infoblox_aaaa_record" "aaaa_rec_filter" {
    filters = {
        name = "debug.test.com"
        ipv6addr = "2002::100"
        view = "nondefault_dnsview" // associated DNS view
    }
 }
 ```

!> From the above example, if the 'view' alias 'dns_view' value is not specified, if same record exists in one or more different DNS views, those
all records will be fetched in results.

!> If `null` or empty filters are passed, then all the records or objects associated with datasource like here `infoblox_aaaa_record` will be fetched in results.

### Example of an AAAA-record Data Source Block

This example defines a data source of type `infoblox_aaaa_record` and the name "qa_rec_temp", which is configured in a Terraform file.
You can reference this resource and retrieve information about it.

```hcl
resource "infoblox_aaaa_record" "vip_host" {
  fqdn = "very-interesting-host.example.com"
  ipv6_addr = "2a05:d014:275:cb00:ec0d:12e2:df27:aa60"
  comment = "some comment"
  ttl = 120 // 120s
  ext_attrs = jsonencode({
    "Location" = "65.8665701230204, -37.00791763398113"
  })
}

data "infoblox_aaaa_record" "qa_rec_temp" {
  filters = {
    name ="very-interesting-host.example.com"
    ipv6addr ="2a05:d014:275:cb00:ec0d:12e2:df27:aa60"
  }

  // This is just to ensure that the record has been be created
  // using 'infoblox_aaaa_record' resource block before the data source will be queried.
  depends_on = [infoblox_aaaa_record.vip_host]
}

output "qa_rec_res" {
  value = data.infoblox_aaaa_record.qa_rec_temp
}

// accessing ip addr field in results
output "qa_rec_addr" {
  value = data.infoblox_aaaa_record.qa_rec_temp.results.0.ip_addr //zero represents index of json object from results list
}

// accessing AAAA-Record through EA's
data "infoblox_aaaa_record" "qa_rec_ea" {
  filters = {
    "*Site" = "sample test site"
    "*Location" = "65.8665701230204, -37.00791763398113"
  }
}

output "qa_rec_out" {
  value = data.infoblox_aaaa_record.qa_rec_ea
}
```
