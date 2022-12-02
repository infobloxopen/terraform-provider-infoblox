# Creates next available network from a given parent CIDR in NIOS grid
terraform {
  # Required providers block for Terraform v0.13 and later
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

# Allocate a network in Infoblox Grid under provided parent CIDR
resource "infoblox_network" "ib_network"{
  network_view_name = "default"
  network_name = "tf-network"
  tenant_id = "tf-AWS-tenant"
  allocate_prefix_len = 24
  parent_cidr = "10.0.0.0/16"
  reserve_ip = 2
}

# Allocate IP from network 
resource "infoblox_ip_allocation" "ib_ip_allocation"{
  network_view_name= "default"
  vm_name = "tf-ec2-instance"
  cidr = infoblox_network.ib_network.cidr
  tenant_id = "tf-AWS-tenant" 
}

# Update Grid with VM data
resource "infoblox_ip_association" "ib_ip_associate"{
  network_view_name = "default"
  vm_name = infoblox_ip_allocation.ib_ip_allocation.vm_name
  cidr = infoblox_network.ib_network.cidr
  mac_addr = aws_network_interface.ni.mac_address
  ip_addr = infoblox_ip_allocation.ib_ip_allocation.ip_addr
  vm_id = aws_instance.ec2-instance.id
  tenant_id = "tf-AWS-tenant"
}
