# IP Association Resource

The `infoblox_ip_association` resource enables you to manage DHCP related properties of the Host record object that was created using the `infoblox_ip_allocation` resource. You can update the Host record created using the allocation resource with details of the VM created in the cloud environment. The VM details include a MAC address in case of an IPv4 address and a DUID in case of an IPv6 address.

The following list describes the parameters you can define in the `infoblox_ip_association` resource block:

* `internal_id`: required,  specifies the reference of the Terraform internal ID that is generated for the Host record by the IP address allocation operation. Example: `infoblox_ip_allocation.foo.internal_id`.
* `enable_dhcp`: optional, a flag that specifies whether the Host record must be created in the DHCP server side. The default value is `false`.
* `mac_addr`: required only if the Host record has an IPv4 address, applies only to IPv4 addresses. It specifies the MAC address to associate the IP address with. It is the MAC address of the network interface of the cloud instance that corresponds to the Host record created using the `infoblox_ip_allocation` resource. The default value is an empty string, which is internally transformed to` 00:00:00:00:00:00`. Example: `02:42:97:87:70:f9`.

-> When you keep the default value, the `enable_dhcp` flag in NIOS for the referenced IPv4 address, is automatically set to `false` even though the field is set to `true` in the Terraform .state file.

* `duid`: required only if the Host record has an IPv6 address, applies only for IPv6 addresses. It specifies the DHCPv6 Unique Identifier (DUID) of the address object. The default is an empty string. Example: `34:df:37:1a:d9:7f` (The DUID could be the same type of value as the MAC address of a network interface).

!> If the Host record contains of an IPv6 address, you must enter a value in this field. Otherwise, NIOS returns an error when the association resource is created.

-> Currently, for a Host record with both IPv4 and IPv6 addresses, you can disable or enable DHCP for both the address types at the same time using the Terraform plug-in. If you want to disable DHCP for any one of the addresses when the flag is enabled (`enable_dhcp = "true"`), in the `infoblox_ip_association` resource block, comment out `mac_addr` to disable only the IPv4 address, or comment out `duid` to disable only the IPv6 address. If you want to enable DHCP for any one of the addresses when the flag is disabled (`enable_dhcp = "false"`), in the `infoblox_ip_association` resource block, set `enable_dhcp = "true"`, then comment out `mac_addr` to keep only the IPv6 address, or comment out the `duid` to keep only the IPv4 address. This causes the MAC address or DUID to be disassociated from the Host record at the NIOS side.

### Example of the Resource Block

The following examples show the creation of two networks (one IPv4 and one IPv6), allocation of IP addresses from those networks, and then associating the network interface IDs with the IP addresses.

```hcl
resource infoblox_ipv4_network "net1" {
  cidr = "10.0.0.0/24"
}

resource infoblox_ipv6_network "net2" {
  cidr = "2001::/56"
}

resource "infoblox_ip_allocation" "foo" {
  fqdn="hostname1.test.com"
  ipv4_addr="10.0.0.12"
  ipv6_addr="2001::10"
  enable_dns = "false"
  comment = "10.0.0.12 IP is allocated"
  ttl=0
  ext_attrs = jsonencode({
    "Tenant ID" = "terraform_test_tenant"
    "VM Name" =  "tf-ec2-instance"
    "Location" = "Test loc."
    "Site" = "Test site"
  })
  depends_on = [infoblox_ipv4_network.net1, infoblox_ipv6_network.net2]
}

resource "infoblox_ip_association" "foo" {
  enable_dhcp = true
  mac_addr = "11:22:33:44:55:66"
  duid = "22:44:66"
  internal_id = infoblox_ip_allocation.foo.internal_id
}
```

The following example shows static allocation of IPv4 address and dynamic allocation of IPv6 address:

```hcl
resource infoblox_ipv4_network "net1" {
  cidr = "10.0.0.0/24"
}

resource infoblox_ipv6_network "net2" {
  cidr = "2001::/56"
}

resource "infoblox_ip_allocation" "foo" {
  fqdn="hostname1.test.com"
  ipv4_addr="10.0.0.12"
  ipv6_cidr = infoblox_ipv6_network.net2.cidr
  enable_dns = "false"
  comment = "10.0.0.12 IP is allocated"
  ttl=0
  ext_attrs = jsonencode({
    "Tenant ID" = "terraform_test_tenant"
    "VM Name" =  "tf-ec2-instance"
    "Location" = "Test loc."
    "Site" = "Test site"
  })

  depends_on = [infoblox_ipv4_network.net1]

}

resource "infoblox_ip_association" "foo" {
  enable_dhcp = true
  mac_addr = "11:22:33:44:55:66"
  duid = "22:44:66"
  internal_id = infoblox_ip_allocation.foo.internal_id
}
```
