# Zone Auth Data Source

Use the `infoblox_zone_auth` data source to retrieve the following information for Authoritative Zone if any, which are managed by a NIOS server:

* `fqdn`: The name of this DNS zone. For a reverse zone, this is in “address/cidr” format. Example: `11.10.0.0/24`. For other zones, this is in FQDN format. Example: `demozone.com` This value can be in unicode format.
* `view`: The name of the DNS view in which the zone resides. Example: `external`.
* `zone_format`: Determines the format of corresponding zone. Valid values are `FORWARD`, `IPV4` and `IPV6`.
* `ns_group`: The name server group that serves DNS for this zone. Example: `demoGroup`.
* `comment`: The Description of Authoritative Zone Object. Example: `random authoritative zone`.
* `ext_attrs`: the set of extensible attributes of the record, if any. The content is formatted as string of JSON map. Example: `"{\"Location\":\"unknown\",\"TestEA\":\"ZoneTesting\"}"`.

For usage of filters, add the fields as keys and appropriate values to be passed to the keys like `name`, `view` corresponding to object.
From the below list of supported arguments for filters,  use only the searchable fields for retriving the matching records.

### Supported Arguments for filters

-----
| Field       | Alias       | Type   | Searchable |
|-------------|-------------|--------|------------|
| fqdn        | fqdn        | string | yes        |
| view        | view        | string | yes        |
| zone_format | zone_format | string | yes        |
| comment     | comment     | string | yes        |

!> Any of the combination from searchable fields in supported arguments list for fields are allowed.

!> Please consider using only fields as the keys in terraform datasource filters, kindly don't use alias names as keys from the above table.

### Example for using the filters:
 ```hcl
 data "infoblox_zone_auth" "zone_filter" {
    filters = {
        fqdn = "testing.com"
        view = "nondefault_dnsview" // associated DNS view
        zone_format = "FORWARD"
    }
 }
 ```
!> From the above example, if the 'view' value is not specified, if same zone name exists in one or more different DNS views, those
all zones will be fetched in results.

!> If `null` or empty filters are passed, then all the zones or objects associated with datasource like here `infoblox_zone_auth` will be fetched in results.

### Example of the Zone Auth Data Source Block

```hcl
resource "infoblox_zone_auth" "zone1" {
  fqdn = "test3.com"
  view = "default"
  zone_format = "FORWARD"
  ns_group = "testGroup"
  comment = "Zone Auth created newly"
  ext_attrs = jsonencode({
    Location = "AcceptanceTerraform"
  })
}

data "infoblox_zone_auth" "dzone1" {
  filters = {
    view = "default"
    fqdn = "test3.com"
    zone_format = "FORWARD"
  }

  # This is just to ensure that the zone has been be created
  # using 'infoblox_zone_auth' resource block before the data source will be queried.
  depends_on = [infoblox_zone_auth.zone1]
}

output "zauth_res" {
  value = data.infoblox_zone_auth.dzone1
}

# Accessing individual field in results
output "zauth_name" {
  value = data.infoblox_zone_auth.dzone1.results.0.fqdn # zero represents index of json object from results list
}

# Accessing Zone Auth through EA's
data "infoblox_zone_auth" "zauth_ea" {
  filters = {
    "*Location" = "AcceptanceTerraform"
  }
}

# Throws matching Zones with EA, if any
output "zauth_out" {
  value = data.infoblox_zone_auth.zauth_ea
}
```
