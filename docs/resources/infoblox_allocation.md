# IP address allocation

'infoblox_ipv4_allocation' and 'infoblox_ipv6_allocation' resources
allow allocation of the next available IP address from a specified
network block. Creation of an 'infoblox_*_allocation' resource
actually creates a NIOS Host record, with an IP address assigned to it;
the IP address is thereby marked as 'used' in appropriate network block.
The IP address may be specified directly in the resource's definition,
or allocated automatically (using special form 'func:nextavailableip').
Its attributes are:

| Attribute | Required/optional | Description | Example |
| --- | --- | --- | --- |
| fqdn | required | Specifies a domain name of a host which an IP address is to be allocated for. In FQDN format. A VM name in a cloud environment usually may be used as such a host name. | ip-12-34-56-78.us-west-2.compute.internal | 
| network_view | optional | Network view to get the specified network block from. If not specified, ‘default’ network view is used. | dmz_netview |
| dns_view       | optional | DNS view to create DNS resource records associated with the IP address. If omitted, the value ‘default’ is used. Makes sense only if ‘enable_dns’ is ‘true’. | internal_network |
| enable_dns     | optional | A flag that specifies whether it is needed to create DNS records associated with the resource. The default value is ‘true’. | true |
| enable_dhcp    | optional | A flag that specifies whether to enable DHCP-related functionality for this resource. The default value is ‘false’. | false |
| cidr            | required for dynamic allocation, see the description | The network block (in CIDR format) where to allocate an IP address from. Used for dynamic allocation; in this case ‘ip_addr’ attribute is empty or omitted. | 10.4.3.128/20 2a00:1148::/32 |
| ip_addr         | required for static allocation, see the description | An IP address to be allocated (marked as ‘Used’). For static allocation, in which case the ‘cidr’ attribute has to be empty or omitted. | 10.4.3.138 |
| mac_addr       | optional        | Only for IPv4. The MAC address to associate the IP address with. The default value is ‘00:00:00:00:00:00’. | 02:42:97:87:70:f9 |
| duid            | required        | Only for IPv6. DHCPv6 Unique Identifier (DUID) of the address object. | 0c:c0:84:d3:03:00:09:12 | 
| ttl             | optional        | The same as for DNS-related resources. This attribute’s value makes sense only when ‘enable_dns’ is ‘true’. If omitted, the value of this attribute is the same as for the parent zone of the DNS records for the resource. | 3600 |

> **Warning! If a host record with enable_dns = true
> has a name as FQDN, and then a user does an update
> making enable_dns = false, then the name of the host
> record MUST be changed to the form of just a name,
> without the zone part. Example: test.example.com -> test**

> **Note: Currently there is no support for multiple
> IP addresses for a host record.**

## Example

    resource "infoblox_ipv4_allocation" "alloc1" {
      network_view="edge"
      cidr="172.30.11.0/24 # this is to allocate
                           # an IP address from the given network block
      dns_view="default" # may be commented out
      fqdn="honeypot-vm.edge.example.com"
      enable_dns = "false"
      enable_dhcp = "false"
      comment = "A honeypot VM for malicious queries."
    }
    
    resource "infoblox_ipv6_allocation" "alloc2" {
      network_view="edge"
      cidr="2a00:1148::/64" # this is to allocate
                            # an IP address from the given network block
      fqdn="honeypot-vm.edge.example.com"
      duid = "0c:c0:84:d3:03:00:09:12"
      enable_dns = "false"
      enable_dhcp = "false"
      comment = "A honeypot VM for malicious queries."
    }
