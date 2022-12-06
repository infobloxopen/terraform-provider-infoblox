# IP Association Resource

The `infoblox_ip_association` resource enables you to manage DHCP related properties of the Host record object that was created
using the `infoblox_ip_allocation` resource. You can update the Host record created using the allocation resource with
details of the VM created in the cloud environment. The VM details include a MAC address in case of an IPv4 address and
a DUID in case of an IPv6 address.

The following list describes the parameters you can define in the `infoblox_ip_association` resource block:

* `internal_id`: required, specifies the value of the "Terraform Internal ID" extensible attribute (one of the pre-requisites for the plugin),
  it is generated for the host record when creating the `infoblox_ip_allocation` resource.
  It is located in the `id` field of the appropriate `infoblox_ip_allocation` resource. Please, note that:

  * The reference to the `internal_id` field of `infoblox_ip_allocation` will continue to work until the next release.
  * The `id` and `internal_id` fields of the `infoblox_ip_allocation` resource use the same values.

  Example: `infoblox_ip_allocation.foo.internal_id`.

* `enable_dhcp`: optional, specifies whether the host record must be created on the DHCP server side. The default value is `false`.
* `mac_addr`: required only if the Host record has an IPv4 address, applies only to IPv4 addresses.
  It specifies the MAC address to associate the IP address with. It is the MAC address of the network interface of the cloud instance that
  corresponds to the host record created using the `infoblox_ip_allocation` resource.
  The default value is an empty string, which is internally transformed to `00:00:00:00:00:00`.

  -> When you keep the default value, the `enable_dhcp` flag in NIOS for the referenced IPv4 address, is automatically set to `false` even though the field is set to `true` in the Terraform .state file.

  Example: `02:42:97:87:70:f9`.

* `duid`: required only if the Host record has an IPv6 address, applies only for IPv6 addresses. 
  It specifies the DHCPv6 Unique Identifier (DUID) of the address object.
  The default is an empty string. Example: `34:df:37:1a:d9:7f`
  (The DUID could be the same type of value as the MAC address of a network interface).

!> If the Host record contains an IPv6 address, you must enter a value in this field. Otherwise, NIOS returns an error when the association resource is created.

-> Currently, for a Host record with both IPv4 and IPv6 addresses, you can disable or enable DHCP for both the address types at the same time
   using the Terraform plug-in. If you want to disable DHCP for any one of the addresses when the flag is enabled (`enable_dhcp = "true"`),
   in the `infoblox_ip_association` resource block, comment out `mac_addr` to disable only the IPv4 address, or comment out `duid` to disable only
   the IPv6 address. If you want to enable DHCP for any one of the addresses when the flag is disabled (`enable_dhcp = "false"`),
   in the `infoblox_ip_association` resource block, set `enable_dhcp = "true"`, then comment out `mac_addr` to keep only the IPv6 address,
   or comment out the `duid` to keep only the IPv4 address. This causes the MAC address or
   DUID to be disassociated from the Host record at the NIOS side.

### Examples of a Resource Block

The following examples of `infoblox_ip_association`, reference the examples of
`infoblox_ip_allociation` in the `infoblox_ip_allocation` topic. You can link the examples by referring to the
allocation resource name used in the `internal_id` parameter.

```hcl
resource "infoblox_ip_association" "association1" {
  internal_id = infoblox_ip_allocation.allocation1.id // which of the existing host records to deal with

  # enable_dhcp = false // this is the default
  mac_addr = "12:00:43:fe:9a:8c" // the address will be assigned but DHCP configuration will not contain it
}

resource "infoblox_ip_association" "association2" {
  internal_id = infoblox_ip_allocation.allocation4.id // let's choose allocation resource with enable_dns = false

  # enable_dhcp = false // this is the default
  duid = "00:43:d2:0a:11:e6" // the address will be assigned but DHCP configuration will not contain it
}

resource "infoblox_ip_association" "association3" {
  internal_id = infoblox_ip_allocation.allocation5.id

  enable_dhcp = true // all systems go

  mac_addr = "12:43:fd:ba:9c:c9"
  duid = "00:43:d2:0a:11:e6"
}

resource "infoblox_ip_association" "association4" {
  internal_id = infoblox_ip_allocation.allocation3.id

  enable_dhcp = true // all systems go

  // DHCP will be enabled for IPv4 ...
  mac_addr = "09:01:03:d3:db:2a"

  // ... and disabled for IPv6
  # duid = "10:2a:9f:dd:3e:0a"
  // yes, DUID will be de-associated, but you can uncomment this later
  // if you will decide to enable DHCP for IPv6
}
```
