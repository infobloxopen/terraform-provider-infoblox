# PTR-record Data Source

Use the data source to retrieve the following information for list of PTR-records from the corresponding object in NIOS:

* `dns_view`: the DNS view which the record's zone belongs to.
* `ip_addr`: the IPv4 or IPv6 address associated with the PTR-record.
* `record_name`: the name of the PTR-record in FQDN format, which can be used instead of an IP address. Example: `1.0.0.10.in-addr.arpa`.
* `ptrdname`: the fully qualified domain name that the PTR-record points to. Example: `delivery.test.com`
* `zone`: the zone which the record belongs to.
* `ttl`: the "time to live" value of the record, in seconds. Example: `1800`.
* `comment`: the description of the record. This is a regular comment. Example: `manager's PC`.
* `ext_attrs`: the set of extensible attributes of the record, if any. The content is formatted as string of JSON map. Example: `"{\"Owner\":\"State Library\", \"Expires\": \"never\"}"`.

As new feature filters are introduced, specifying combination DNS view , IPv4 address or IPv6 address or record name used instead of IP address
and ptrdname is removed.

For usage of filters, add the fields as keys and appropriate values to be passed to the keys like `name`, `view` corresponding to object. Only searchable fields
from below list of supported arguments for filters, are allowed to use in filters, for retrieving one or more records or objects matching
filters.

### Supported Arguments for filters

-----
| Field       | Alias    | Type   | Searchable |
|-------------|----------|--------|------------|
| ptrdname    | ---      | string | yes        |
| record_name | name     | string | yes        |
| view        | dns_view | string | yes        |
| ipv4addr    | ip_addr  | string | yes        |
| ipv6addr    | ip_addr  | string | yes        |
| ttl         | ---      | uint32 | no         |
| comment     | ---      | string | yes        |
| zone        | ---      | string | yes        |

!> From above list, both ipv4addr and ipv6addr are not allowed together in filters. Apart from this any other combination is allowed.

!> Please consider using only fields as the keys in terraform datasource filters, kindly don't use alias names as keys from the above table.

### Example for using the filters:
 ```hcl
 data "infoblox_ptr_record" "ptr_rec_filter" {
    filters = {
        ptrdname = "testing.example.com"
        view = "default" // associated DNS view
        ipv4addr = "20.11.0.19"
    }
 }
 ```

!> From the above example, if the 'view' alias 'dns_view' value is not specified, if same record exists in one or more different DNS views, those
all records will be fetched in results.

!> If `null` or empty filters are passed, then all the records or objects associated with datasource like here `infoblox_ptr_record` will be fetched in results.

### Example of the PTR-record Data Source Block

This example defines a data source of type `infoblox_ptr_record` and the name "vip_host", which is configured in a Terraform file.
You can reference this resource and retrieve information about it.

```hcl
resource "infoblox_ptr_record" "host1" {
  ptrdname = "host.example.org"
  ip_addr = "2a05:d014:275:cb00:ec0d:12e2:df27:aa60"
  comment = "workstation #3"
  ttl = 300 # 5 minutes
  ext_attrs = jsonencode({
    "Location" = "the main office"
  })
}

data "infoblox_ptr_record" "host1" {
  filters = {
    ptrdname="host.example.org"
    ip_addr="2a05:d014:275:cb00:ec0d:12e2:df27:aa60"
  }

  // This is just to ensure that the record has been be created
  // using 'infoblox_ptr_record' resource block before the data source will be queried.
  depends_on = [infoblox_ptr_record.host1]
}

output "ptr_rec_res" {
  value = data.infoblox_ptr_record.host1
}

data "infoblox_ptr_record" "host2" {
  filters = {
    dns_view="default"
    ptrdname="host.example.org"
    record_name="0.6.a.a.7.2.f.d.2.e.2.1.d.0.c.e.0.0.b.c.5.7.2.0.4.1.0.d.5.0.a.2.ip6.arpa"
  }
}

output "ptr_host_res" {
  value = data.infoblox_ptr_record.host2
}

// accessing individual field in results
output "ptr_rec_name" {
  value = data.infoblox_ptr_record.host2.results.0.ptrdname //zero represents index of json object from results list
}

// accessing PTR-Record through EA's
data "infoblox_ptr_record" "ptr_rec_ea" {
  filters = {
    "*Owner" = "State Library"
  }
}

// throws PTR-Records with EA, if any
output "ptr_rec_out" {
  value = data.infoblox_ptr_record.ptr_rec_ea
}
```
