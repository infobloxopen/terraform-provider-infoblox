# IPv4 Allocation Resource

-> This resource will be deprecated in an upcoming release. Infoblox strongly recommends that you use `infoblox_ip_allocation` resource for allocation of IP addresses.

The `infoblox_ipv4_allocation` resource allows allocation of the next available IPv4 address from a specified network block if an IP address is not explicitly specified. The IP address can be allocated by specifying it in the resource definition or can be allocated automatically by defining the CIDR field. The allocated IP address is marked as ‘Used’ in the appropriate network block.

The allocation is done by creating a Host record in NIOS. The Host record can be created at the DNS server side by using the `enable_dns` flag, which is enabled by default. If you disable the flag, the record is not created in the DNS server side, but is visible on the IPAM tab in NIOS Grid Manager.

When you create the Host record at the DNS server side, specify the host name in the FQDN format. The name must include the zone name and the host name as shown in the following example:

```
enable_dns=true
fqdn=hostname1.zone.com
```

If you disable the `enable_dns` flag, specify only the host name as FQDN. For example: `fqdn=hostname1`

The following list describes the parameters you can define in the IP address allocation resource blocks:

* `fqdn`: required, specifies the name (in FQDN format) of a host for which an IP address needs to be allocated. When `enable_dns` is set to `true`, specify the zone name along with the host name in format: <hostname>.<zone>.
  When `enable_dns` is set to `false`, specify only the host name: <hostname>. In a cloud environment, a VM name could be used as a host name. Example: `ip-12-34-56-78.us-west-2.compute.internal`.
* `network_view`: optional, specifies the network view from which to get the specified network block. If a value is not specified, the default network view is considered. Example: `nview2`.
* `dns_view`: optional, specifies the DNS view in which to create DNS resource records that are associated with the IP address. If a value is not specified, the default DNS view is considered. This parameter is relevant only when `enable_dns` is set to `true`. Example: `internal_view`.
* `enable_dns`: optional, a flag that specifies whether DNS records associated with the resource must be created. The default value is `true`.
* `cidr`: required only for dynamic allocation, specifies the network block (in CIDR format) from where to allocate an IP address. For static allocation, do not use this field. Example: `10.4.3.128/20`.
* `ip_addr`: required only for static allocation, specifies an IP address that should be allocated (marked as ‘Used’ in NIOS Grid Manager). For dynamic allocation, do not use this field. Example: `10.4.3.138`.
* `ttl`: optional, specifies the "time to live" value for the DNS record. This parameter is relevant only when `enable_dns` is to `true`. If a value is not specified, the value is same as that of the parent zone of the DNS records for this resource. Example: `3600`.
* `comment`: optional, describes the resource. Example: `QA cloud instance`.
* `ext_attrs`: optional, set of NIOS extensible attributes that are attached to the resource.

### Example of the Resource Block

```hcl
resource "infoblox_ipv4_allocation" "alloc1" {
  network_view = "edge"
  cidr         = "172.30.11.0/24" # this is to allocate
  # an IP address from the given network block
  dns_view   = "default" # may be commented out
  fqdn       = "test-vm.edge.example.com"
  enable_dns = "true"
  comment    = "Allocating an IPv4 address"
  ext_attrs = jsonencode({
    "Tenant ID"       = "tf-plugin"
    "Cloud API Owned" = "True"
    "CMP Type"        = "VMware"
    "Site"            = "Nevada"
  })
}
```
