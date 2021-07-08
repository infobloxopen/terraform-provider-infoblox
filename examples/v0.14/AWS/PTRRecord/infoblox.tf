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
    "CMP Type" = "Terraform"
    "Cloud API Owned" = "True"
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
    "CMP Type" = "Terraform"
    "Cloud API Owned" = "True"
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
    "CMP Type" = "Terraform"
    "Cloud API Owned" = "True"
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
    "CMP Type" = "Terraform"
    "Cloud API Owned" = "True"
    "Network Name" = "ipv6-tf-network"
    "Location" = "Test loc."
    "Site" = "Test site"
  })
}


# Allocate IP from network
resource "infoblox_ipv4_allocation" "ipv4_allocation"{
  network_view= "default"
  cidr = infoblox_ipv4_network.ipv4_network.cidr
  host_name = "test"

  #Create Host Record with DNS and DHCP flags
  #dns_view="default"
  #zone="aws.com"
  #enable_dns = "false"
  #enable_dhcp = "false"
  
  comment = "tf IPv4 allocation"
  ext_attrs = jsonencode({
    "Tenant ID" = "tf-plugin"
    "CMP Type" = "Terraform"
    "Cloud API Owned" = "True"
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
  host_name = "test"

  #Create Host Record with DNS and DHCP flags
  #dns_view="default"
  #zone="aws.com"
  #enable_dns = "false"
  #enable_dhcp = "false"

  comment = "tf IPv6 allocation"
  ext_attrs = jsonencode({
    "Tenant ID" = "tf-plugin"
    "CMP Type" = "Terraform"
    "Cloud API Owned" = "True"
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
  host_name = "test"

  #Create Host Record with DNS and DHCP flags
  #dns_view="default"
  #zone="aws.com"
  #enable_dns = "false"
  #enable_dhcp = "false"

  comment = "tf IPv4 Association"
  ext_attrs = jsonencode({
    "Tenant ID" = "tf-plugin"
    "CMP Type" = "Terraform"
    "Cloud API Owned" = "True"
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
  host_name = "test"

  #Create Host Record with DNS and DHCP flags
  #dns_view="default"
  #zone="aws.com"
  #enable_dns = "false"
  #enable_dhcp = "false"

  comment = "tf IPv6 Association"
  ext_attrs = jsonencode({
    "Tenant ID" = "tf-plugin"
    "CMP Type" = "Terraform"
    "Cloud API Owned" = "True"
    "Network Name" = "ipv6-tf-network"
    "VM Name" =  "tf-ec2-instance-ipv6"
    "VM ID" =  aws_instance.ec2-instance.id
    "Location" = "Test loc."
    "Site" = "Test site"
  })
}

# Create PTR record for VM
resource "infoblox_ptr_record" "ib_ptr_record"{
  ptrdname = "tf-ec2-instance.aws.com"
  dns_view = "default"

  # Record in forward mapping zone
  record_name = "tf-ec2-instance-ipv4.aws.com"

  # Record in reverse mapping zone
  #network_view = "default"
  #cidr = infoblox_ipv4_network.ipv4_network.cidr
  #ip_addr = infoblox_ipv4_allocation.ipv4_allocation.ip_addr

  ttl = 3600

  comment = "PTR record created"
  ext_attrs = jsonencode({
    "Tenant ID" = "tf-plugin"
    "CMP Type" = "Terraform"
    "Cloud API Owned" = "True"
    "Location" = "Test loc."
    "Site" = "Test site"
  })
}
