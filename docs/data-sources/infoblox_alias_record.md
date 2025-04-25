# Alias-record Data Source

Use the `infoblox_alias_record` data resource for the Alias record to retrieve the following information for a Alias record:

* `name`: the alias name of the record in the FQDN format. Example: `foo1.test.com`
* `dns_view`: the DNS view which the record's zone belongs to. Example: `nondefault_dnsview`
* `target_name`: the name of the target in FQDN format. Example: `aa.test.com`
* `target_type`: the type of the target object. Valid values are: `A`, `AAAA`, `MX`, `NAPTR`, `PTR`, `SPF`, `SRV` and `TXT`.
* `zone`: the name of the zone in which the record exists. Example: `test.com`.
* `ttl`: the "time to live" value of the record, in seconds. Example: `3600`.
* `dns_name`: the name for an Alias record in punycode format. Example: `foo1.test.com`.
* `dns_target_name`: the DNS target name of the Alias Record in punycode format. Example: `aa.test.com`.
* `disable`: the flag to disable the record. Valid values are `true` and `false`.
* `comment`: the text describing the record. This is a regular comment. Example: `Temporary Alias-record`.
* `creator`: the creator of the record. Valid value is `STATIC`.
* `ext_attrs`: the set of extensible attributes of the record, if any. The content is formatted as string of JSON map. Example: `"{\"Site\":\"Greece\"}"`
* `cloud_info`: Structure containing all cloud API related information for this object. Example: `"{\"authority_type\":\"GM\",\"delegated_scope\":\"NONE\",\"owned_by_adaptor\":false}"`

For usage of filters, add the fields as keys and appropriate values to be passed to the keys like `name`, `view`, `zone`, `comment`, `target_name`, and `target_type`  corresponding to object.
From the below list of supported arguments for filters,  use only the searchable fields for retrieving the matching records.

### Supported Arguments for filters

-----
| Field       | Alias       | Type   | Searchable |
|-------------|-------------|--------|------------|
| name        | name        | string | yes        |
| view        | dns_view    | string | yes        |
| target_name | target_name | string | yes        |
| target_type | target_type | uint   | yes        |
| comment     | comment     | string | yes        |
| zone        | zone        | string | yes        |

!> Any of the combination from searchable fields in supported arguments list for fields are allowed.

!> Please consider using only fields as the keys in terraform datasource filters, kindly don't use alias names as keys from the above table.

### Example for using the filters:
 ```hcl
data "infoblox_alias_record" "alias_read" {
  filters = {
    name = "foo1.test.com"
    comment = "Temporary Alias-record"
  }
}
 ```

!> From the above example, if the 'dns_view' alias 'dns_view' value is not specified, if same record exists in one or more different DNS views, those
all records will be fetched in results.

!> If `null` or empty filters are passed, then all the records or objects associated with datasource like here `infoblox_alias_record`, will be fetched in results.

### Example of the Alias-record Data Source Block

This example defines a data source of type `infoblox_alias_record` and the name "alias_read", which is configured in a Terraform file.
You can reference this resource and retrieve information about it.

```hcl
resource "infoblox_alias_record" "alias_record" {
  name = "alias-record.test.com"
  target_name = "hh.ll.com"
  target_type = "NAPTR"
  comment = "example alias record"
  dns_view = "default"
  disable = true
  ttl = 1200
  ext_attrs = jsonencode({
    "Site" = "Ireland"
  })
}

data "infoblox_alias_record" "alias_read"{
  filters = {
    name = infoblox_alias_record.alias_record.name
    target_name = infoblox_alias_record.alias_record.target_name
    view = infoblox_alias_record.alias_record.dns_view
  }
}

output "alias_record_out" {
  value = data.infoblox_alias_record.alias_read
}

// accessing individual field in results
output "alias_target_type_out" {
  value = data.infoblox_alias_record.alias_read.results.0.target_type //zero represents index of json object from results list
}

// accessing Alias-Record through EA's
data "infoblox_alias_record" "alias_read_ea" {
  filters = {
    "*Site" = "Ireland"
  }
}

// throws matching Alias records with EA, if any
output "alias_read_ea_out" {
  value = data.infoblox_alias_record.alias_read_ea
}
```
