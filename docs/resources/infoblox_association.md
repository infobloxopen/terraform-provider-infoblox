# IP address association

'infoblox_ipv4_association' and 'infoblox_ipv6_association'
resources are for mapping an IP address to a particular VM, by MAC
address or DUID (for IPv4 and IPv6 respectively). This is pretty much
the same as allocation-resources but works with already existing host
records (ex. created as 'infoblox_*_allocation' resources). More
differences: 'cidr' attribute is useless and 'ip_addr' is mandatory.

## Examples

    resource "infoblox_ipv4_association" "association1" {
      network_view="edge"
    
      # The value is taken from the previous example.
      ip_addr = infoblox_ipv4_allocation.alloc1
    
      fqdn="honeypot-vm.edge.example.com"
      enable_dns = "false"
      enable_dhcp = "false"
    
      # The VM network interface’s MAC address is used.
      mac_addr = aws_network_interface.vm1_ni0.mac_address
    }
    
    resource "infoblox_ipv6_association" "association2" {
      network_view="edge"
    
      # The value is taken from the previous example.
      ip_addr = infoblox_ipv6_allocation.alloc2
    
      fqdn="honeypot-vm.edge.example.com"
      enable_dns = "false"
      enable_dhcp = "false"
    
      # The VM network interface’s MAC address is used.
      mac_addr = aws_network_interface.vm1_ni1.mac_address
    }
