# AAAA-record Resource

The `infoblox_aaaa_record` resource associates a domain name with an IPv6 address.

The following list describes the parameters you can define in the resource block of the record:

* `fqdn`: required, specifies the fully qualified domain name for which you want to assign the IP address to. Example: `host43.zone12.org`
* `network_view`: optional, specifies the Network view to use when allocating an IP address from a network dynamically. For static allocation, do not use this field. Example: `networkview1`
* `dns_view`: optional, specifies the DNS view in which the zone exists. If a value is not specified, the default DNS view is considered. Example: `dns_view_1`
* `ttl`: optional, specifies the time to live value for the record. There is no default value for this parameter. If you do not specify a value, the TTL value is inherited from Grid DNS properties. A TTL value of 0 (zero) means caching should be disabled for this record. Example: `600`
* `comment`: optional, describes the record. Example: `static record #1`
* `ext_attrs`: optional, a set of NIOS extensible attributes that are attached to the record. Example: `jsonencode({})`
* `ip_addr`: required only for static allocation, specifies the IPv4 address to associate with the A-record. Example: `2001:db8::ff00:42:8329`.
    * For allocating a static IP address, specify a valid IP address.
    * For allocating a dynamic IP address, do not use this field. Instead, define the `cidr` field. Optionally, set the `network_view` if you do not want to allocate in the default network view.
* `cidr`: required only for dynamic allocation, specifies the network from which to allocate an IP address when the `ipv6_addr` field is empty. The address is in CIDR format. For static allocation, use `ipv6_addr` instead of this field. Example: `2001::/64`.

-> All parameters except `fqdn` are optional. If you do not specify the `ipv6_addr` parameter, you must use the `cidr` parameter for dynamic address allocation. `cidr` and `ipv6_addr` are mutually exclusive. Specify a custom network view with `cidr` if you do not want to use the default network view.

### Examples of the AAAA-record Block

```hcl
resource "infoblox_aaaa_record" "aaaa_record_static1"{

  # the zone 'example.com' MUST exist in the DNS view ('default')
  fqdn = "static1-v6.example.com"

  ipv6_addr = "2001:db8::ff00:42:8329"
  ttl = 10
  comment = "static AAAA-record #1"
  ext_attrs = jsonencode({
    "Tenant ID" = "tf-plugin"
    "Cloud API Owned" = "True"
    "CMP Type"= "VMware"
    "Location" = "New York"
    "Site" = "HQ"
  })
}

resource "infoblox_aaaa_record" "aaaa_record_dynamic1"{
  fqdn = "dynamic2-v6.example.com"

  # not specifying explicitly means using the default network view
  # network_view = "default"

  # not specifying explicitly means using the default DNS view
  # dns_view = "default"

  # appropriate network MUST exist in the network view
  cidr = "infoblox_ipv6_network.ipv6_network.cidr"
}
```
