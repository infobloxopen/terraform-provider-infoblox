# Zone Delegated Data Source

Use the `infoblox_zone_delegated` data source to retrieve the following information about a delegated DNS zone from the corresponding object in NIOS:

* `fqdn`: The name of this DNS zone. For a reverse zone, this is in “address/cidr” format. Example: `11.10.0.0/24`. For other zones, this is in FQDN format. Example: `demozone.com` This value can be in unicode format.
* `view`: The name of the DNS view in which the zone resides. Example: `external`.
* `comment`: The Description of Delegated Zone Object. Example: `random delegated zone`.
* `ext_attrs`: The set of extensible attributes of the record, if any. The content is formatted as string of JSON map. Example: `"{\"Location\":\"unknown\",\"TestEA\":\"ZoneTesting\"}"`.
* `zone_format`: Determines the format of corresponding zone. Valid values are `FORWARD`, `IPV4` and `IPV6`.
* `ns_group`: Specifies the name server group that serves DNS for this zone. Example: `demoGroup`.
* `disable`: Specifies whether the zone is disabled.
* `locked`: The flag that restricts other administrators from making any changes. Note that this flag is for administration purposes only. The zone will continue to serve DNS data even when it is locked. Example: `false`.
* `delegated_ttl`: The TTL value for the delegated zone. Example: `60`.
* `delegate_to`: The remote server to which the NIOS appliance redirects queries for data for the delegated zone. Example:
```terraform
delegate_to {
    name = "te32.dz.ex.com"
    address = "10.0.0.1"
  }
```

For usage of filters, add the fields as keys and appropriate values to be passed to the keys like `fqdn`, `view` corresponding to object.
From the below list of supported arguments for filters,  use only the searchable fields for retrieving the matching records.

### Supported Arguments for filters

-----
| Field       | Alias       | Type   | Searchable |
|-------------|-------------|--------|------------|
| fqdn        | fqdn        | string | yes        |
| view        | view        | string | yes        |
| zone_format | zone_format | string | yes        |
| comment     | comment     | string | yes        |


!> Any combination of searchable fields in the supported arguments list for fields is allowed.

!> "Aliases are the parameter names used in the prior releases of Infoblox IPAM Plug-In for Terraform. Do not use the alias names for parameters in the data source blocks. Using them can result in error scenarios."

### Example for using the filters:
 ```hcl
data "infoblox_zone_delegated" "data_zone_delegated" {
  filters = {
    fqdn = "zone_delegated.ex.org"
    view = "default"
  }
}
 ```
!> From the above example, if the 'view' value is not specified, if same zone name exists in one or more different DNS views, those
all zones will be fetched in results.

!> If `null` or empty filters are passed, then all the zones or objects associated with datasource like here `infoblox_zone_delegated` will be fetched in results.

### Example of the Zone Delegated Data Source Block

```hcl
resource "infoblox_zone_delegated" "delegatedzone_delegateTo" {
  fqdn = "zone_delegated.ex.org"
  delegate_to {
    name = "test22.dz.ex.com"
    address = "10.0.0.1"
  }
  delegate_to {
    name = "test2.dz.ex.com"
    address = "10.0.0.2"
  }
  ext_attrs = jsonencode({
    "Site" = "Antarctica"
  })
}

# Accessing Zone Delegated by specifying fqdn, view and extra attribute Site
data "infoblox_zone_delegated" "data_zone_delegated" {
  filters = {
    fqdn = "zone_delegated.ex.org"
    view = "default"
    "*Site" = "Antarctica"
  }
  # This is just to ensure that the record has been be created
  depends_on = [infoblox_zone_delegated.delegatedzone_delegateTo]
}

# Returns matching Zone Delegated with fqdn and view, if any
output "zone_delegated_data3" {
  value = data.infoblox_zone_delegated.data_zone_delegated
}

resource "infoblox_zone_delegated" "delegatedzone_IPV4_nsGroup" {
  fqdn = "195.1.0.0/24"
  comment = "Delegated zone IPV4"
  zone_format = "IPV4"
  ns_group = "test"
}

# Accessing Zone Delegated by specifying fqdn, view and comment
data "infoblox_zone_delegated" "datazone_delegated_fqdn_view_comment" {
  filters = {
    fqdn = "195.1.0.0/24"
    view = "default"
    comment = "Delegated zone IPV4"
  }
  # This is just to ensure that the record has been be created
  depends_on = [infoblox_zone_delegated.delegatedzone_IPV4_nsGroup]
}

# Returns matching Zone Delegated with fqdn, view and comment, if any
output "zone_delegated_data4" {
  value = data.infoblox_zone_delegated.datazone_delegated_fqdn_view_comment
}
```
