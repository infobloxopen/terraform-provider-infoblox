# Zone Forward Data Source

Use the `infoblox_zone_forward` data source to retrieve the following information for Forward Zone if any, which are managed by a NIOS server:

* `fqdn`: The name of this DNS zone. For a reverse zone, this is in “address/cidr” format. Example: `11.10.0.0/24`. For other zones, this is in FQDN format. Example: `demozone.com` This value can be in unicode format.
* `view`: The name of the DNS view in which the zone resides. Example: `external`.
* `zone_format`: Determines the format of corresponding zone. Valid values are `FORWARD`, `IPV4` and `IPV6`.
* `comment`: The Description of Forward Zone Object. Example: `random forward zone`.
* `ext_attrs`: The set of extensible attributes of the record, if any. The content is formatted as string of JSON map. Example: `"{\"Location\":\"unknown\",\"TestEA\":\"ZoneTesting\"}"`.
* `zone_format`: Determines the format of corresponding zone. Valid values are `FORWARD`, `IPV4` and `IPV6`. Default value: `FORWARD`.
* `ns_group`: Specifies the name server group that serves DNS for this zone. Example: `demoGrp`.
* `external_ns_group`: Specifies the name of the forward stub server. Example: `stubGroup`.
* `disable`: Specifies whether the zone is disabled. Default value: `false`.
* `forwarders_only`: Specifies whether the appliance sends queries to forwarders only, and not to other internal or Internet root servers. Default value: `false`.
* `forward_to`: Determines the information for the remote name servers to which you want the Infoblox appliance to forward queries for a specified domain name. Example:
```terraform
forward_to {
    name = "te32.dz.ex.com"
    address = "10.0.0.1"
  }
```
* `forwarding_servers`: optional, determines the information for the Grid members to which you want the Infoblox appliance to forward queries for a specified domain name. Example:
```terraform
forwarding_servers {
    name = "infoblox.172_28_83_0"
    forwarders_only = true
    use_override_forwarders = true
    forward_to {
      name = "kk.fwd.com"
      address = "10.2.1.31"
    }
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
data "infoblox_zone_forward" "data_zone_forward" {
  filters = {
    fqdn = "zone_forward.ex.org"
    view = "default"
  }
}
 ```
!> From the above example, if the 'view' value is not specified, if same zone name exists in one or more different DNS views, those
all zones will be fetched in results.

!> If `null` or empty filters are passed, then all the zones or objects associated with datasource like here `infoblox_zone_forward` will be fetched in results.

### Example of the Zone Forward Data Source Block

```hcl
resource "infoblox_zone_forward" "forwardzone_forwardTo" {
  fqdn = "zone_forward.ex.org"
  forward_to {
    name = "test22.dz.ex.com"
    address = "10.0.0.1"
  }
  forward_to {
    name = "test2.dz.ex.com"
    address = "10.0.0.2"
  }
  ext_attrs = jsonencode({
    "Site" = "Antarctica"
  })
}

// accessing Zone Forward by specifying fqdn, view and extra attribute Site
data "infoblox_zone_forward" "data_zone_forward" {
  filters = {
    fqdn = "zone_forward.ex.org"
    view = "default"
    "*Site" = "Antarctica"
  }
  // This is just to ensure that the record has been be created
  depends_on = [infoblox_zone_forward.forwardzone_forwardTo]
}

// returns matching Zone Forward with fqdn and view, if any
output "zone_forward_data3" {
  value = data.infoblox_zone_forward.data_zone_forward
}


resource "infoblox_zone_forward" "forwardzone_IPV4_nsGroup_externalNsGroup" {
  fqdn = "195.1.0.0/24"
  comment = "Forward zone IPV4"
  external_ns_group = "stub server"
  zone_format = "IPV4"
  ns_group = "test"
}

// accessing Zone Forward by specifying fqdn, view and comment
data "infoblox_zone_forward" "datazone_foward_fqdn_view_comment" {
  filters = {
    fqdn = "195.1.0.0/24"
    view = "default"
    comment = "Forward zone IPV4"
  }
  // This is just to ensure that the record has been be created
  depends_on = [infoblox_zone_forward.forwardzone_IPV4_nsGroup_externalNsGroup]
}

// returns matching Zone Forward with fqdn, view and comment, if any
output "zone_forward_data4" {
  value = data.infoblox_zone_forward.datazone_foward_fqdn_view_comment
}
```