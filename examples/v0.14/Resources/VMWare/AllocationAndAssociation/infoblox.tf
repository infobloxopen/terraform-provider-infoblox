# Create a network container in Infoblox Grid
resource "infoblox_ipv4_network_container" "IPv4_nw_c" {
  network_view="default"

  cidr = "10.0.0.0/16"
  comment = "tf IPv4 network container"
  ext_attrs = jsonencode({
    "Tenant ID" = "tf-plugin"
    Location = "Test loc."
    Site = "Test site"
  })
}

resource "infoblox_ipv6_network_container" "IPv6_nw_c" {
  network_view="default"

  cidr = "2001:1890:1959:2710::/62"
  comment = "tf IPv6 network container"
  ext_attrs = jsonencode({
    "Tenant ID" = "tf-plugin"
    Location = "Test loc."
    Site = "Test site"
  })
}

# Allocate a network in Infoblox Grid under provided parent CIDR
resource "infoblox_ipv4_network" "ipv4_network"{
  network_view = "default"

  parent_cidr = infoblox_ipv4_network_container.IPv4_nw_c.cidr
  allocate_prefix_len = 24
  reserve_ip = 2

  comment = "tf IPv4 network"
  ext_attrs = jsonencode({
    "Tenant ID" = "tf-plugin"
    "Network Name" = "ipv4-tf-network"
    Location = "Test loc."
    Site = "Test site"
  })
}

resource "infoblox_ipv6_network" "ipv6_network"{
  network_view = "default"

  parent_cidr = infoblox_ipv6_network_container.IPv6_nw_c.cidr
  allocate_prefix_len = 64
  reserve_ipv6 = 3

  comment = "tf IPv6 network"
  ext_attrs = jsonencode({
    "Tenant ID" = "tf-plugin"
    "Network Name" = "ipv6-tf-network"
    Location = "Test loc."
    Site = "Test site"
  })
}

# Allocate IP from network
resource "infoblox_ipv4_allocation" "ipv4_allocation"{
  network_view= "default"
  cidr = infoblox_ipv4_network.ipv4_network.cidr

  #Create Host Record with DNS and DHCP flags
  dns_view="default"
  fqdn="testipv4.vmware.com"
  enable_dns = "false"
  enable_dhcp = "false"

  comment = "tf IPv4 allocation"
  ext_attrs = jsonencode({
    "Tenant ID" = "tf-plugin"
    "Network Name" = "ipv4-tf-network"
    "VM Name" =  "tf-vmware-ipv4"
    Location = "Test loc."
    Site = "Test site"
  })
}

resource "infoblox_ipv6_allocation" "ipv6_allocation" {
  network_view= "default"
  cidr = infoblox_ipv6_network.ipv6_network.cidr
  duid = "00:00:00:00:00:00:00:00"

  #Create Host Record with DNS and DHCP flags
  dns_view="default"
  fqdn="testipv6.vmware.com"
  enable_dns = "false"
  enable_dhcp = "false"

  comment = "tf IPv6 allocation"
  ext_attrs = jsonencode({
    "Tenant ID" = "tf-plugin"
    "Network Name" = "ipv6-tf-network"
    "VM Name" = "tf-vmware-ipv6"
    Location = "Test loc."
    Site = "Test site"
  })
}

# Update Grid with VM data
resource "infoblox_ipv4_association" "ipv4_associate"{
  network_view = "default"
  cidr = infoblox_ipv4_network.ipv4_network.cidr
  ip_addr = infoblox_ipv4_allocation.ipv4_allocation.ip_addr
  mac_addr = vsphere_virtual_machine.vm_ipv4.network_interface[0].mac_address

  #Create Host Record with DNS and DHCP flags
  dns_view="default"
  fqdn="testipv4.vmware.com"
  enable_dns = "false"
  enable_dhcp = "false"

  comment = "tf IPv4 Association"
  ext_attrs = jsonencode({
    "Tenant ID" = "tf-plugin"
    "Network Name" = "ipv6-tf-network"
    "VM Name" = vsphere_virtual_machine.vm_ipv4.name
    "VM ID" =  vsphere_virtual_machine.vm_ipv4.id
    Location = "Test loc."
    Site = "Test site"
  })
}

resource "infoblox_ipv6_association" "ipv6_associate"{
  network_view = "default"
  cidr = infoblox_ipv6_network.ipv6_network.cidr
  ip_addr = infoblox_ipv6_allocation.ipv6_allocation.ip_addr
  duid = vsphere_virtual_machine.vm_ipv6.network_interface[0].mac_address

  #Create Host Record with DNS and DHCP flags
  dns_view="default"
  fqdn="testipv6.vmware.com"
  enable_dns = "false"
  enable_dhcp = "false"

  comment = "tf IPv6 Association"
  ext_attrs = jsonencode({
    "Tenant ID" = "tf-plugin"
    "Network Name" = "ipv6-tf-network"
    "VM Name" =  vsphere_virtual_machine.vm_ipv6.name
    "VM ID" =  vsphere_virtual_machine.vm_ipv6.id
    Location = "Test loc."
    Site = "Test site"
  })
}
