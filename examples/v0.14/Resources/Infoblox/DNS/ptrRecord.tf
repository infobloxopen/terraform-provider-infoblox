terraform {
  # Required providers block for Terraform v0.14.7
  required_providers {
    infoblox = {
      source = "infobloxopen/infoblox"
      version = ">=2.0"
    }
  }
}

# Create PTR record for VM
resource "infoblox_ptr_record" "ib_ptr_record_ipv4" {
  ptrdname = "tf-vmware-ipv4"
  dns_view = "default"

  # Record in forward mapping zone
  record_name = "tf-vmware-ipv4.vmware.com"

  # Record in reverse mapping zone
  #network_view = "default"
  #cidr = infoblox_ipv4_network.ipv4_network.cidr
  #ip_addr = infoblox_ipv4_allocation.ipv4_allocation.ip_addr

  comment = "PTR record created"
  ext_attrs = jsonencode({
    "Tenant ID" = "tf-plugin"
    "Network Name" = "ipv6-tf-network"
    "VM Name" =  "tf-vmware-ipv4"
    "Location" = "Test loc."
    "Site" = "Test site"
  })
}

resource "infoblox_ptr_record" "ib_ptr_record_ipv6" {
  ptrdname = "tf-vmware-ipv6.vmware.com"
  dns_view = "default"

  # Record in forward mapping zone
  record_name = "tf-vmware-ipv6.vmware.com"

  # Record in reverse mapping zone
  #network_view = "default"
  #cidr = infoblox_ipv4_network.ipv4_network.cidr
  #ip_addr = infoblox_ipv4_allocation.ipv4_allocation.ip_addr

  comment = "PTR record created"
  ext_attrs = jsonencode({
    "Tenant ID" = "tf-plugin"
    "Network Name" = "ipv6-tf-network"
    "VM Name" =  "tf-vmware-ipv6"
    "Location" = "Test loc."
    "Site" = "Test site"
  })
}
