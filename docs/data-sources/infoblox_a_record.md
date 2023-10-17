# A-record Data Source

Use the `infoblox_a_record` data source to retrieve the following information for list of A-records if any, which are managed by a NIOS server:

* `dns_view`: the DNS view which the record's zone belongs to. Example: `default`
* `ip_addr`: the IPv4 address associated with the A-record. Example: `17.10.0.8`
* `fqdn`: the fully qualified domain name which the IP address is assigned to. `blues.test.com`
* `zone`: the zone that contains the record in the specified DNS view. Example: `test.com`.
* `ttl`: the "time to live" value of the record, in seconds. Example: `1800`.
* `comment`: the description of the record. This is a regular comment. Example: `Temporary A-record`.
* `ext_attrs`: the set of extensible attributes of the record, if any. The content is formatted as string of JSON map. Example: `"{\"TestEA\":56,\"TestEA1\":\"kickoff\"}"`

As there is new feature filters , the previous usage of combination of DNS view, IPv4 address and FQDN, has been removed.

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
| ipv4addr | ip_addr  | string | yes        |

!> Any of the combination from searchable fields in supported arguments list for fields are allowed.

!> Please consider using only fields as the keys in terraform datasource filters, kindly don't use alias names as keys from the above table.

### Example for using the filters:
 ```hcl
 data "infoblox_a_record" "a_rec_filter" {
    filters = {
        name = "testing.demo1.com"
        view = "nondefault_dnsview" // associated DNS view
    }
 }
 ```

!> From the above example, if the 'view' alias 'dns_view' value is not specified, if same record exists in one or more different DNS views, those
all records will be fetched in results.

!> If `null` or empty filters are passed, then all the records or objects associated with datasource like here `infoblox_a_record` will be fetched in results.

### Example of an A-record Data Source Block

This example defines a data source of type `infoblox_a_record` and the name "a_rec_temp", which is configured in a Terraform file.
You can reference this resource and retrieve information about it.

```hcl
resource "infoblox_a_record" "vip_host" {
  fqdn = "very-interesting-host.example.com"
  ip_addr = "10.3.1.65"
  comment = "special host"
  dns_view = "nondefault_dnsview2"
  ttl = 120 // 120s
  ext_attrs = jsonencode({
    "Location" = "65.8665701230204, -37.00791763398113"
  })
}


data "infoblox_a_record" "a_rec_temp" {
  filters = {
    name = "very-interesting-host.example.com"
    ipv4addr = "10.3.1.65" //alias is ip_addr
    view = "nondefault_dnsview2"
  }
  
  // This is just to ensure that the record has been be created
  // using 'infoblox_a_record' resource block before the data source will be queried.
  depends_on = [infoblox_a_record.vip_host]
}

output "a_rec_res" {
  value = data.infoblox_a_record.a_rec_temp
}

// accessing individual field in results
output "a_rec_name" {
  value = data.infoblox_a_record.a_rec_temp.results.0.fqdn //zero represents index of json object from results list
}

// accessing A-Record through EA's
data "infoblox_a_record" "a_rec_ea" {
  filters = {
    "*Site" = "some test site"
    "*Location" = "65.8665701230204, -37.00791763398113"
  }
}

output "a_rec_out" {
  value = data.infoblox_a_record.a_rec_ea
}
```

