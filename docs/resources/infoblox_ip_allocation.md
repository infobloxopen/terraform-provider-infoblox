# IP Allocation Resource

The `infoblox_ip_allocation` resource allows allocation of a new IP address from a network that already exists as a NIOS object. The IP address can be allocated statically by specifying an address or dynamically as the next available IP address from the specified IPv4 and/or IPv6 network blocks. The allocation is done by creating a host record in NIOS with an IPv4 address, an IPv6 address, or both assigned to the record. The allocated IP address is marked as ‘used’ in the appropriate network block.

-> As a prerequisite for creation of Host records using the `infoblox_ip_allocation` and `infoblox_ip_association` resources, you must create the extensible attribute `Terraform Internal ID` of string type in Infoblox NIOS Grid Manager. For steps, refer to the [Infoblox NIOS Documentation] (<https://docs.infoblox.com/space/NIOS/35400616/NIOS>).

The following list describes the parameters you can define in the `infoblox_ip_allocation` resource block:

* `fqdn`: required, specifies the name (in FQDN format) of a host with which an IP address needs to be allocated.
  In a cloud environment, a VM name could be used as a host name.
  For a host record with its name in FQDN format and the
  `enable_dns` flag enabled, if you disable the flag,
  you must remove the zone part from the record name and
  keep only the host name.
  For example, hostname1.zone.com must be changed to `hostname1`.
  Example: `ip-12-34-56-78.us-west-2.compute.internal`.
* `network_view`: optional, specifies the network view from which to get the specified network block.
  If a value is not specified, the name `default` is set as the network view. Example: `dmz_netview`.
* `dns_view`: optional, specifies the DNS view in which to create the DNS
  resource records that are associated with the IP address.
  * If `enable_dns` is set to `true`, you must configure this parameter.
  * If `enable_dns` is set to `false`, you must remove this parameter from the resource block.

  For more information, see the description of the enable_dns parameter.
  Example: `external`.
* `enable_dns`: optional, a flag that specifies whether DNS records associated with the resource must be created. The default value is `true`.
  When you update the enable_dns parameter, consider the following points:
  * If you set the parameter to `false` when **creating** a resource, you must not specify the `dns_view` parameter.
  * If you set the parameter to `false` when **updating** a resource, you must:
    * Remove the `dns_view` parameter. Not removing it can result in unexpected errors.
    * Remove the zone part in the FQDN and keep only the host record name.
      For example, the FQDN with `host1` as the record name and zone as `example.org`, must be changed from `host1.example.org` to `host1`.
  * If you set the parameter from `false` to `true`, you must:
    * Specify the `dns_view` parameter because if a value is not specified, the name `default` is NOT configured.
    * Change the `fqdn` value to FQDN without a leading dot, that means, add a zone that exists in the specified DNS view to the
       name of the host record. For example, if the `fqdn` value is `host1` and the zone
       selected from the specified DNS view is `example.com`, then the `fqdn` must be changed to `host1.example.com`.
* `ipv4_cidr`: required only for dynamic allocation, specifies the IPv4 network block (in CIDR format) from where to allocate the next available IP address.
  Use this parameter only when `ipv4_addr` is not specified. Example: `10.0.0.0/24`.
* `ipv6_cidr`: required only for dynamic allocation, specifies the IPv6 network block (in CIDR format) from where to allocate the next available IP address.
  Use this parameter only when `ipv6_addr` is not specified. Example: `2000:1148::/32`.
* `ipv4_addr`: required only for static allocation, specifies an IPv4 address to allocate.
  Use this parameter only when `ipv4_cidr` is not specified. The allocated IP address will be marked as ‘Used’ in NIOS Grid Manager.
  The default value is an empty string. If you specify both `ipv4_addr` and `ipv4_cidr`, then `ipv4_addr` is ignored.
  Example: `10.0.0.10`.
* `ipv6_addr`: required only for static allocation, specifies an IPv6 address to allocate.
  Use this parameter only when `ipv6_cidr` is not specified. The allocated IP address will be marked as ‘Used’ in NIOS Grid Manager.
  The default value is an empty string. If you specify both `ipv6_addr` and `ipv6_cidr`, then the `ipv6_addr` address is allocated and `ipv6_cidr` is ignored.
  Example: `2000:1148::10`.
* `filter_params`: required for dynamic allocation only if `ipv4_addr`, `ipv4_cidr`, `ipv6_addr` and `ipv6_cidr` are not set, specifies the extensible attributes of the parent network that must be used as filters to retrieve the next available IP address for creating the host record object.
  The content is formatted as a string of a JSON map. Example: `jsonencode({"*Site": "Turkey"})`.
* `ip_address_type`: required only when filter_params is used, Specifies the type of IP address to allocate. The valid values are, `IPV4`, `IPV6`, and `Both`. The default value is `IPv4`.
* `ttl`: optional, specifies the 'time to live' value for the DNS record. This parameter is relevant only when `enable_dns` is set to `true`.
  If a value is not specified, then in NIOS, the value is inherited from the parent zone of the DNS records for this resource. Example: `3600`.
* `disable`: optional,specifies whether the record disabled or not. The default value is `false`. Example: `true`.
* `comment`: optional, specifies the human-readable description of the resource. Example: `Front-end cloud node`.
* `aliases`: optional, specifies the list of aliases for the host record. Example: `["alias1", "alias2"]`.
* `ext_attrs`: optional, specifies the set of NIOS extensible attributes that are attached to the NIOS resource.
  An extensible attribute must be a JSON map translated into a string value. Example:

```hcl
jsonencode({
  "Tenant ID" = "tf-plugin"
  "Location"  = "Test loc."
  "Site"      = "Test site"
})
```

When you use the `infoblox_ip_allocation` resource block to allocate or deallocate a static IP address from a Host record,
you must configure appropriate dependencies so that workflows run in the correct order. In the following example,
dependencies have been configured for network view and the extensible attribute, `Network Name`:

```hcl
resource "infoblox_ip_allocation" "ip_allocation" {
  network_view = infoblox_ipv6_network.ipv6_network.network_view
  ipv4_addr    = "10.0.0.32"
  ipv6_addr    = "2001:1890:1959:2710::32"

  #Create Host Record with DNS flags
  dns_view   = "default"
  fqdn       = "testipv4v6"
  enable_dns = "false"

  ext_attrs = jsonencode({
    "Tenant ID"    = "tf-plugin"
    "Network Name" = lookup(jsondecode(infoblox_ipv4_network.ipv4_network.ext_attrs), "Network Name")
    "VM Name"      = "tf-vmware-ipv4-ipv61"
    "Location"     = "Test loc."
    "Site"         = "Test site"
  })
}
```

When you perform a `create` or an `update` operation using this allocation resource, the following read-only parameters are computed:

* `allocated_ipv4_addr`: if you allocated a dynamic IP address, this value is the IP address allocated from the specified IPv4 CIDR or `filter_params`.
  If you allocated a static IP address, this value is the IP address that you specified in the `ipv4_addr` field.
  You can reference this field for the IP address when using other resources. Example:

```hcl
resource "infoblox_ip_aloocation" "allocation1" {
  dns_view  = "default"
  ipv4_cidr = infoblox_ipv4_network.cidr
  fqdn      = local.vm_name
}
# You can add a reference for the IP address as follows: infoblox_ip_allocation.allocation1.allocated_ipv4_addr
```
```hcl
resource "infoblox_ip_allocation" "ipv4_allocation_with_ea" {
  fqdn = local.vm_name
  //Extensible attributes of parent network 
  filter_params = jsonencode({
    "*Site": local.site
  })
}
```

* `allocated_ipv6_addr`: if you allocated a dynamic IP address, this value is the IP address allocated from the specified IPv6 CIDR or `filter_params`.
  If you allocated a static IP address, this value is the IP address that you specified `ipv6_addr` field.
  You can reference this field for the IP address when using other resources. See the previous description for an example.

### Examples of a Resource Block

You can use static and dynamic allocation methods independently to assign IPv4 and IPv6 addresses within the same
resource.

```hcl
// IP address allocation, minimal set of parameters
resource "infoblox_ip_allocation" "allocation1" {
  dns_view = "default" // this time the parameter is not optional, because...
  # enable_dns = true // ... this is 'true', by default
  fqdn      = "host1.example1.org" // this resource type is implemented as a host record on NIOS side
  ipv4_addr = "1.2.3.4"
}

// Wide set of parameters
resource "infoblox_ip_allocation" "allocation2" {
  dns_view  = "nondefault_dnsview1"
  fqdn      = "host2.example2.org"
  ipv6_addr = "2002:1f93::1234"
  ttl       = 120
  comment   = "another host record, statically allocated"
  ext_attrs = jsonencode({
    "Tenant ID" = "tenant_3261798"
  })
}

// IPv4 and IPv6 at the same time
resource "infoblox_ip_allocation" "allocation3" {
  dns_view     = "nondefault_dnsview2"
  network_view = "nondefault_netview" // we have to mention the exact network view which the DNS view belongs to
  fqdn         = "host3.example4.org"
  ipv6_addr    = "2002:1f93::12:40"
  ipv4_addr    = "1.2.3.40"
  ttl          = 0
  comment      = "still statically allocated, but IPv4 and IPv6 together"
  ext_attrs = jsonencode({
    "Tenant ID" = "tenant_3261798"
  })
}

resource "infoblox_ip_allocation" "allocation4" {
  enable_dns = false
  # dns_view = "nondefault_dnsview2" // dns_view makes no sense when enable_dns = false and
  // you MUST remove it or comment out
  network_view = "nondefault_netview" // we have to mention the exact network view which the DNS view belongs to
  fqdn         = "host4"              // just one name component, because there is no host-to-DNS-zone assignment

  // either of the IP addresses must belong to an existing
  // network (not a network container) in the GIVEN network view,
  // because of enable_dns = false
  ipv6_addr = "2002:1f93::12:40"
  ipv4_addr = "10.1.0.60"

  // we must ensure that appropriate network exists by the time this resource is being created
  depends_on = [infoblox_ipv4_network.net1]
}

// dynamic allocation
resource "infoblox_ip_allocation" "allocation5" {
  dns_view     = "nondefault_dnsview2"
  network_view = "nondefault_netview"
  fqdn         = "host5.example4.org"
  ipv6_cidr    = infoblox_ipv6_network.net2.cidr
  ipv4_cidr    = infoblox_ipv4_network.net2.cidr
}

// dynamic allocation of both IPv4 and IPv6 host records using filter_params with aliases
resource "infoblox_ip_allocation" "rec_host17" {
  fqdn = "new777.test.com"
  aliases = ["www.test.com"]
  disable = false
  //Extensible attributes of parent network 
  filter_params = jsonencode({
    "*Site": "Turkey"
  })
  ip_address_type = "Both"
  enable_dns = true
  ttl = 60
}
```
