# CNAME-record Data Source

Use the `infoblox_cname_record` data resource for the CNAME object to retrieve the following information for a CNAME record:

* `dns_view`: the DNS view which the record's zone belongs to. Example: `nondefault_dnsview`
* `canonical`: the canonical name of the record in the FQDN format. Example: `debug.point.somewhere.in`
* `alias`: the alias name of the record in the FQDN format. Example: `foo1.test.com`
* `zone`: the zone that contains the record in the specified DNS view. Example: `test.com`.
* `ttl`: the "time to live" value of the record, in seconds. Example: `3600`.
* `comment`: the text describing the record. This is a regular comment. Example: `Temporary CNAME-record`.
* `ext_attrs`: the set of extensible attributes of the record, if any. The content is formatted as string of JSON map. Example: `"{\"Owner\":\"State Library\",\"Expiry\":\"Never\"}"`

As there is new feature filters , the previous usage of combination of DNS view, alias and canonical name, has been removed.

For usage of filters, add the fields as keys and appropriate values to be passed to the keys like `name`, `view` corresponding to object.
From the below list of supported arguments for filters,  use only the searchable fields for retriving the matching records.

### Supported Arguments for filters

-----
| Field     | Alias     | Type   | Searchable |
|-----------|-----------|--------|------------|
| name      | alias     | string | yes        |
| view      | dns_view  | string | yes        |
| canonical | canonical | string | yes        |
| ttl       | ttl       | uint   | no         |
| comment   | comment   | string | yes        |
| zone      | zone      | string | yes        |

!> Any of the combination from searchable fields in supported arguments list for fields are allowed.

!> Please consider using only fields as the keys in terraform datasource filters, kindly don't use alias names as keys from the above table.

### Example for using the filters:
 ```hcl
 data "infoblox_cname_record" "cname_rec_filter" {
    filters = {
        name = "testing.demo1.com"
        canonical = "delivery.random.street.in"
    }
 }
 ```

!> From the above example, if the 'view' alias 'dns_view' value is not specified, if same record exists in one or more different DNS views, those
all records will be fetched in results.

!> If `null` or empty filters are passed, then all the records or objects associated with datasource like here `infoblox_cname_record`, will be fetched in results.

### Example of the CNAME-record Data Source Block

This example defines a data source of type `infoblox_cname_record` and the name "cname_rec", which is configured in a Terraform file.
You can reference this resource and retrieve information about it.

```hcl
resource "infoblox_cname_record" "foo" {
  dns_view = "default.nondefault_netview"
  canonical = "strange-place.somewhere.in.the.net"
  alias = "foo.test.com"
  comment = "we need to keep an eye on this strange host"
  ttl = 0 // disable caching
  ext_attrs = jsonencode({
    Site = "unknown"
    Location = "TBD"
  })
}

data "infoblox_cname_record" "cname_rec"{
  filters = {
    name = "foo.test.com"
    canonical = "strange-place.somewhere.in.the.net"
    view = "default.nondefault_netview"
  }

  // This is just to ensure that the record has been be created
  // using 'infoblox_cname_record' resource block before the data source will be queried.
  depends_on = [infoblox_cname_record.foo]
}

output "cname_rec_out" {
  value = data.infoblox_cname_record.cname_rec
}

// accessing individual field in results
output "cname_rec_alias" {
  value = data.infoblox_cname_record.cname_rec.results.0.alias //zero represents index of json object from results list
}

// accessing CNAME-Record through EA's
data "infoblox_cname_record" "cname_rec_ea" {
  filters = {
    "*Location" = "Cali"
  }
}

// throws matching CNAME records with EA, if any
output "cname_rec_res" {
  value = data.infoblox_cname_record.cname_rec_ea
}
```
