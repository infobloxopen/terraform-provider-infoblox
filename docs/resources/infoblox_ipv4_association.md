# IPv4 Association Resource

-> This resource will be deprecated in an upcoming release. Infoblox strongly recommends that you use `infoblox_ip_association` resource for allocation of IP addresses.

-> If you are not using other Terraform providers with the Infoblox provider to deploy virtual machines and allocate IP addresses from NIOS, then ignore the `infoblox_ipv4_association` resource block, which is used for updating the properties of virtual machines.

With the IP address association operation, you can update the Host record created using the IP address allocation operation with details of the VM created in the cloud environment. The VM details include VM tags (as extensible attributes) and MAC address.

The `infoblox_ipv4_association` resource maps the IP address of a Host record created in NIOS to a VM, by MAC address.

The following list describes the parameters you can define in the resource block:

- `fqdn`: required, specifies the name (in FQDN format) of a host for which an IP address needs to be allocated. When `enable_dns` is set to `true`, specify the zone name along with the host name in format: <hostname>.<zone>.
  When `enable_dns` is set to `false`, specify only the host name: <hostname>. Example: `ip-12-34-56-78.us-west-2.compute.internal`.
- `network_view`: optional, specifies the network view from which to get the specified network block. If a value is not specified, the default network view is considered. Example: `nview2`.
- `dns_view`: optional, specifies the DNS view in which to create DNS resource records that are associated with the IP address. If a value is not specified, the default DNS view is considered. This parameter is relevant only when `enable_dns` is set to `true`. Example: `internal_view`.
- `enable_dns`: optional, a flag that specifies whether DNS records associated with the resource must be created. The default value is `true`.
- `enable_dhcp`: optional, a flag that specifies whether to enable DHCP-related functionality for this resource. The default value is false.
- `ip_addr`: required, specifies an IP address that should be allocated (marked as ‘Used’ in NIOS Grid Manager). Example: `10.4.3.138`.
- `ttl`: optional, specifies the "time to live" value for the DNS record. This parameter is relevant only when `enable_dns` is to `true`. If a value is not specified, the value is same as that of the parent zone of the DNS records for this resource. Example: `3600`.
- `comment`: optional, describes the resource. Example: `QA cloud instance`.
- `ext_attrs`: optional, set of NIOS extensible attributes that are attached to the resource.
- `mac_addr`: optional, specifies the MAC address to associate the IP address with. The default value is `00:00:00:00:00:00`. Example: `02:42:97:87:70:f9`.

### Example of the Resource Block

As the IP address association operation is dependent on the allocation operation, the following examples for IPv4 demonstrate how to define the resource blocks in the Terraform configuration file:

```hcl
resource "infoblox_ipv4_allocation" "ipv4_allocation" {
  network_view = "default"
  cidr         = infoblox_ipv4_network.ipv4_network.cidr

  #Create Host Record with DNS and DHCP flags
  dns_view   = "default"
  fqdn       = "testipv4.aws.com"
  enable_dns = "true"
  # The VM network interface’s MAC address is used.
  comment = "Allocating an IP address"

  ext_attrs = jsonencode({
    "Tenant ID"       = "tf-plugin"
    "Cloud API Owned" = "True"
    "CMP Type"        = "AWS"
    "Site"            = "Nevada"
  })
}

resource "infoblox_ipv4_association" "ipv4_associate" {
  network_view = infoblox_ipv4_network.ipv4_network.network_view
  dns_view     = infoblox_ipv4_allocation.ipv4_allocation.dns_view
  ip_addr      = infoblox_ipv4_allocation.ipv4_allocation.ip_addr
  fqdn         = infoblox_ipv4_allocation.ipv4_allocation.fqdn
  enable_dns   = infoblox_ipv4_allocation.ipv4_allocation.enable_dns
  enable_dhcp  = infoblox_ipv4_allocation.ipv4_allocation.enable_dhcp
  mac_addr     = aws_network_interface.ni.mac_address
  comment      = "Associating of an IPv4 address"
  ext_attrs    = infoblox_ipv4_allocation.ipv4_allocation.ext_attrs
}
```
