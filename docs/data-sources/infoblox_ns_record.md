# NS-record Data Source

Use the `infoblox_ns_record` data source to retrieve the following information for an NS-Record if any, which is managed by a NIOS server:

* `name`: The name of the NS record in FQDN format. This value can be in unicode format. Values with leading or trailing white space are not valid for this field. Example: `test.com`
* `nameserver`: The domain name of an authoritative server for the redirected zone. Values with leading or trailing white space are not valid for this field. Example: `ns1.test.com`
* `addresses`: The list of zone name servers.
```terraform
addresses {
  address = "2.3.3.3"
  auto_create_ptr = true
}
```
* `dns_view`: The name of the DNS view in which the record resides.The default value is The default DNS view. Example: `external`
* `ms_delegation_name`: The MS delegation point name. The default value is Empty string. Example: `delegation.com`
* `creator`: The record creator. Valid values are `STATIC` and `SYSTEM`. Example: `STATIC`
* `dns_name`: The name of the NS record in punycode format. Example: `test.com`
* `policy`: The host name policy for the record. Example: `Allow Underscore`
* `zone`: The name of the zone in which the record resides. Example: `test.com`
* `cloud_info`: Structure containing all cloud API related information for this object. Example: `"{\"authority_type\":\"GM\",\"delegated_scope\":\"NONE\",\"owned_by_adaptor\":false}"`

### Supported Arguments for filters

-----
| Field      | Alias      | Type   | Searchable |
|------------|------------|--------|------------|
| name       | name       | string | yes        |
| nameserver | nameserver | string | yes        |
| creator    | creator    | string | yes        |
| view       | dns_view   | string | yes        |
| zone       | zone       | string | yes        |

!> Any of the combination from searchable fields in supported arguments list for fields are allowed.

!> Please consider using only fields as the keys in terraform datasource filters, kindly don't use alias names as keys from the above table.

### Example for using the filters:
 ```hcl
 data "infoblox_ns_record" "testns_read" {
  filters = {
    name = "test.com"
  }
}
 ```

!> If `null` or empty filters are passed, then all the objects associated with datasource like here `infoblox_ns_record` will be fetched in results.

### Example of an NS-record Data Source Block
```hcl
resource "infoblox_ns_record" "ns"{
  name = "test.com"
  nameserver = "name27.test.com"
  addresses{
    address = "2.3.2.5"
    auto_create_ptr=true
  }
  dns_view = "default"
}

data "infoblox_ns_record" "testNs_read" {
  filters = {
    name = infoblox_ns_record.ns.name
    view = infoblox_ns_record.ns.dns_view
  }
}

output "dtc_rec_res1" {
  value = data.infoblox_ns_record.testNs_read
}
```