terraform {
  # Required providers block for Terraform v0.14.7
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 3.0"
    }
    infoblox = {
      source  = "terraform-providers/infoblox"
      version = ">= 1.0"
    }
  }
}

# Create a network container in Infoblox Grid
resource "infoblox_ipv4_network_container" "IPv4_nw_c" {
  network_view="default"

  cidr = aws_vpc.vpc.cidr_block
  comment = "tf IPv4 network container"
  ext_attrs = jsonencode({
    "Tenant ID" = "tf-plugin"
    "Location" = "Test loc."
    "Site" = "Test site"
  })
}

resource "infoblox_ipv6_network_container" "IPv6_nw_c" {
  network_view="default"

  cidr = aws_vpc.vpc.ipv6_cidr_block
  comment = "tf IPv6 network container"
  ext_attrs = jsonencode({
    "Tenant ID" = "tf-plugin"
    "Location" = "Test loc."
    "Site" = "Test site"
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
    "Location" = "Test loc."
    "Site" = "Test site"
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
    "Location" = "Test loc."
    "Site" = "Test site"
  })
}


# Allocate IP from network
resource "infoblox_ipv4_allocation" "ipv4_allocation"{
  network_view= "default"
  cidr = infoblox_ipv4_network.ipv4_network.cidr

  #Create Host Record with DNS and DHCP flags
  dns_view="default"
  fqdn="testipv4.aws.com"
  enable_dns = "false"
  enable_dhcp = "false"
  
  comment = "tf IPv4 allocation"
  ext_attrs = jsonencode({
    "Tenant ID" = "tf-plugin"
    "Network Name" = "ipv4-tf-network"
    "VM Name" =  "tf-ec2-instance"
    "Location" = "Test loc."
    "Site" = "Test site"
  })
}

resource "infoblox_ipv6_allocation" "ipv6_allocation" {
  network_view= "default"
  cidr = infoblox_ipv6_network.ipv6_network.cidr
  duid = "00:00:00:00:00:00:00:00"

  #Create Host Record with DNS and DHCP flags
  dns_view="default"
  fqdn="testipv6.aws.com"
  enable_dns = "false"
  enable_dhcp = "false"

  comment = "tf IPv6 allocation"
  ext_attrs = jsonencode({
    "Tenant ID" = "tf-plugin"
    "Network Name" = "ipv6-tf-network"
    "VM Name" =  "tf-ec2-instance-ipv6"
    "Location" = "Test loc."
    "Site" = "Test site"
  })
}


# Update Grid with VM data
resource "infoblox_ipv4_association" "ipv4_associate"{
  network_view = "default"
  cidr = infoblox_ipv4_network.ipv4_network.cidr
  ip_addr = infoblox_ipv4_allocation.ipv4_allocation.ip_addr
  mac_addr = aws_network_interface.ni.mac_address

  #Create Host Record with DNS and DHCP flags
  dns_view="default"
  fqdn="testipv4.aws.com"
  enable_dns = "false"
  enable_dhcp = "false"

  comment = "tf IPv4 Association"
  ext_attrs = jsonencode({
    "Tenant ID" = "tf-plugin"
    "Network Name" = "ipv6-tf-network"
    "VM Name" =  "tf-ec2-instance"
    "VM ID" =  aws_instance.ec2-instance.id
    "Location" = "Test loc."
    "Site" = "Test site"
  })
}

resource "infoblox_ipv6_association" "ipv6_associate"{
  network_view = "default"
  cidr = infoblox_ipv6_network.ipv6_network.cidr
  ip_addr = infoblox_ipv6_allocation.ipv6_allocation.ip_addr
  duid = aws_network_interface.ni.mac_address

  #Create Host Record with DNS and DHCP flags
  dns_view="default"
  fqdn="testipv6.aws.com"
  enable_dns = "false"
  enable_dhcp = "false"

  comment = "tf IPv6 Association"
  ext_attrs = jsonencode({
    "Tenant ID" = "tf-plugin"
    "Network Name" = "ipv6-tf-network"
    "VM Name" =  "tf-ec2-instance-ipv6"
    "VM ID" =  aws_instance.ec2-instance.id
    "Location" = "Test loc."
    "Site" = "Test site"
  })
}

/*
Below are the examples for Allocation through host record creation.
The same pattern follows for IPv6 Allocation, IPv4 and IPv6 Association.
# Allocate with default network and dns views
# Create Host record with default dns flag enabled
resource "infoblox_ipv4_allocation" "ipv4_allocation"{
  ip_addr = "10.0.0.11"
  fqdn="test.aws.com"
}
# Allocate with default network and dns views
# Create Host record with dns flag and dhcp flag disabled
resource "infoblox_ipv4_allocation" "ipv4_allocation"{
  ip_addr = "10.0.0.11"
  fqdn="test.aws.com"
  enable_dns = "false"
  enable_dhcp = "false"
} 
# Allocate with default network and dns views
# Create Host record with dns flag and dhcp flag enabled
resource "infoblox_ipv4_allocation" "ipv4_allocation"{
  ip_addr = "10.0.0.11"
  fqdn="test.aws.com"
  enable_dns = "true"
  enable_dhcp = "true"
} 
*/