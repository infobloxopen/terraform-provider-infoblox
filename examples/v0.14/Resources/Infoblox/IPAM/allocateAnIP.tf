terraform {
  # Required providers block for Terraform v0.14.7
  required_providers {
    infoblox = {
      source = "infobloxopen/infoblox"
      version = ">=2.0"
    }
  }
}

# Will be depreciated, use infoblox_ip_allocation resource for the same
# Allocate IP from IPv4 network
resource "infoblox_ipv4_allocation" "ipv4_allocation"{
  network_view= "default"
  cidr = "10.0.0.0/24"

  #Create Host Record with DNS and DHCP flags
  dns_view="default"
  fqdn="testipv4.aws.com"
  enable_dns = "true"
  enable_dhcp = "true"
  
  comment = "tf IPv4 allocation"
  ext_attrs = jsonencode({
    "Tenant ID" = "tf-plugin"
    "Network Name" = "ipv4-tf-network"
    "VM Name" =  "tf-ec2-instance"
    "Location" = "Test loc."
    "Site" = "Test site"
  })
}

# Will be depreciated, use infoblox_ip_allocation resource for the same
# Allocate IP from IPv6 network
resource "infoblox_ipv6_allocation" "ipv6_allocation" {
  network_view= "default"
  cidr = "2000:00/64"
  duid = "00:00:00:00:00:00:00:00"

  #Create Host Record with DNS and DHCP flags
  dns_view="default"
  fqdn="testipv6.aws.com"
  enable_dns = "true"
  enable_dhcp = "true"

  comment = "tf IPv6 allocation"
  ext_attrs = jsonencode({
    "Tenant ID" = "tf-plugin"
    "Network Name" = "ipv6-tf-network"
    "VM Name" =  "tf-ec2-instance-ipv6"
    "Location" = "Test loc."
    "Site" = "Test site"
  })
}

# Allocate an IP from IPv4 and or IPv6 network
resource "infoblox_ip_allocation" "ip_allocation" {
  network_view= "default"
  ipv4_cidr = "10.0.0.0/24"
  ipv6_cidr = "2000::00/64"
  duid = "00:00:00:00:00:00:00:01"

  #Create Host Record with DNS and DHCP flags
  dns_view="default"
  fqdn="testip.example.com"
  enable_dns = "false"

  comment = "tf IPv4 and IPv6 allocation"
  ext_attrs = jsonencode({
    "Tenant ID" = "tf-plugin"
    "Network Name" = "tf-network"
    "VM Name" =  "tf-ec2-instance"
    "Location" = "Test loc."
    "Site" = "Test site"
  })
}
