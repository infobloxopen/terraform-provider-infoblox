# AAAA-record Resource

The `infoblox_aaaa_record` resource associates a domain name with an IPv6 address.

The following list describes the parameters you can define in the resource block of the record:

* `fqdn`: required, specifies the fully qualified domain name for which you want to assign the IP address to. Example: `host43.zone12.org`
* `network_view`: optional, specifies the network view to use when allocating an IP address from a network dynamically. If a value is not specified, the name `default` is used for the network view. For static allocation, do not use this field. Example: `networkview1`
* `dns_view`: optional, specifies the DNS view in which the zone exists. If a value is not specified, the name `default` is used for DNS view. Example: `dns_view_1`
* `ttl`: optional, specifies the "time to live" value for the record. There is no default value for this parameter. If you do not specify a value, the TTL value is inherited from Grid DNS properties. A TTL value of 0 (zero) means caching should be disabled for this record. Example: `600`
* `comment`: optional, describes the record. Example: `static record #1`
* `ext_attrs`: koptional, a set of NIOS extensible attributes that are attached to the record. Example: `jsonencode({})`
* `ipv6_addr`: required only for static allocation, specifies the IPv6 address to associate with the AAAA-record. Example: `2001:db8::ff00:42:8329`.
  * For allocating a static IP address, specify a valid IP address.
  * For allocating a dynamic IP address, configure the `cidr` field instead of `ipv6_addr` . Optionally, specify a `network_view` if you do not want to allocate it in the network view `default`.
* `cidr`: required only for dynamic allocation, specifies the network from which to allocate an IP address when the `ipv6_addr` field is empty. The address is in CIDR format. For static allocation, use `ipv6_addr` instead of `cidr`. Example: `2001::/64`.

### Examples of an AAAA-record Block

```hcl
// static AAAA-record, minimal set of parameters
resource "infoblox_aaaa_record" "aaaa_rec1" {
  fqdn = "static1.example1.org"
  ipv6_addr = "2002:1111::1401" // not necessarily from a network existing in NIOS DB
}

// all the parameters for a static AAAA-record
resource "infoblox_aaaa_record" "aaaa_rec2" {
  fqdn = "static2.example4.org"
  ipv6_addr = "2002:1111::1402"
  comment = "example static AAAA-record aaaa_rec2"
  dns_view = "nondefault_dnsview2"
  ttl = 120 // 120s
  ext_attrs = jsonencode({
    "Location" = "65.8665701230204, -37.00791763398113"
  })
}

// all the parameters for a dynamic AAAA-record
resource "infoblox_aaaa_record" "aaaa_rec3" {
  fqdn = "dynamic1.example2.org"
  cidr = infoblox_ipv6_network.net2.cidr // the network  must exist, you may use the example for infoblox_ipv6_network resource.
  network_view = infoblox_ipv6_network.net2.network_view // not necessarily in the same network view as the DNS view resides in.
  comment = "example dynamic AAAA-record aaaa_rec3"
  dns_view = "nondefault_dnsview1"
  ttl = 0 // 0 = disable caching
  ext_attrs = jsonencode({})
}
```
