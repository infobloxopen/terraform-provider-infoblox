# Zone Forward Resource

The `infoblox_zone_forward` resource associates forward zone with a DNS View.The resource represents the ‘zone_forward’ WAPI object in NIOS.

The following list describes the parameters you can define in the resource block of the zone forward object:

* `fqdn`: required, specifies the name of this DNS zone. For a reverse zone, this is in “address/cidr” format.
  For other zones, this is in FQDN format. This value can be in unicode format.
  Example: `10.1.0.0/24` for reverse zone and `zone1.com` for forward zone.
* `view`: optional, specifies The name of the DNS view in which the zone resides. If value is not specified, `default` will be considered as default DNS view. Example: `external`.
* `zone_format`: optional, determines the format of corresponding zone. Valid values are `FORWARD`, `IPV4` and `IPV6`. Default value: `FORWARD`.
* `ns_group`: optional, specifies the name server group that serves DNS for this zone. Example: `demoGrp`.
* `external_ns_group`: Required if forward_to is not configured. Specifies the name of the forward stub server. Example: `stubGroup`.
* `disable`: optional, specifies whether the zone is disabled. Default value: `false`.
* `forwarders_only`: optional, specifies whether the appliance sends queries to forwarders only, and not to other internal or Internet root servers. Default value: `false`.
* `forward_to`: Required if external_ns_group is not configured. Determines the information for the remote name servers to which you want the Infoblox appliance to forward queries for a specified domain name. Example:
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
* `comment`: optional, description of the zone. Example: `custom forward zone`.
* `ext_attrs`: optional, set of the Extensible attributes of the zone, as a map in JSON format. Example: `jsonencode({})`.

!> For a reverse zone, the corresponding 'zone_format' value should be set. And 'fqdn' once set cannot be updated.
>**Note**: Either define forwarding_servers or ns_group. 
> If both the parameters are configured, settings of ns_group takes precedence.


### Examples of a Zone Forward Block

```hcl
//forward mapping zone, with minimum set of parameters
resource "infoblox_zone_forward" "forward_zone_forwardTo" {
  fqdn = "min_params.ex.org"
  forward_to {
    name = "test22.dz.ex.com"
    address = "10.0.0.1"
  }
  forward_to {
    name = "test2.dz.ex.com"
    address = "10.0.0.2"
  }
}

//forward zone with full set of parameters
resource "infoblox_zone_forward" "forward_zone_full_parameters" {
  fqdn = "max_params.ex.org"
  view = "nondefault_view"
  zone_format = "FORWARD"
  comment = "test sample forward zone"
  forward_to {
    name = "te32.dz.ex.com"
    address = "10.0.0.1"
  }
  forwarding_servers {
    name = "infoblox.172_28_83_216"
    forwarders_only = true
    use_override_forwarders = false
    forward_to {
      name = "cc.fwd.com"
      address = "10.1.1.1"
    }
  }
  forwarding_servers {
    name = "infoblox.172_28_83_0"
    forwarders_only = true
    use_override_forwarders = true
    forward_to {
      name = "kk.fwd.com"
      address = "10.2.1.31"
    }
  }
}

//forward zone with ns_group, external_ns_group and extra attribute Site
resource "infoblox_zone_forward" "forward_zone_nsGroup_externalNsGroup" {
  fqdn = "params_ns_ens.ex.org"
  ns_group = "test"
  external_ns_group = "stub server"
  ext_attrs = jsonencode({
    "Site" = "Antarctica"
  })
}
```

