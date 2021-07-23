# Creates a CNAME record for an AWS instance being created
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

# Create a network in Infoblox Grid 
resource "infoblox_network" "ib_network"{
  network_view_name = "default"
  network_name = "tf-network"
  cidr = "10.0.0.0/24"
  tenant_id = "tf-AWS-tenant"
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

# Create CNAME record
resource "infoblox_cname_record" "ib_cname_record"{
  canonical = infoblox_ip_allocation.ib_ip_allocation.vm_name
  zone = "tf.aws.com"
  dns_view = "default"
  alias = "cloudsolution"
  tenant_id="tf-AWS-tenant"
}
