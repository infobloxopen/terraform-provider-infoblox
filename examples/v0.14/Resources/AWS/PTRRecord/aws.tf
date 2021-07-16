# Region being used to create the resources
provider "aws" {
  region  = "us-west-1"
}

# Create a Virtual Private Cloud
resource "aws_vpc" "vpc" {
  cidr_block = "10.0.0.0/16"
  # Allocates /56 IPv6 CIDR block From Amazon Global Unicast Address to VPC
  assign_generated_ipv6_cidr_block = true
  tags = {
    Name = "tf-vpc"
  }
}

# Create a Subnet
resource "aws_subnet" "subnet" {
  vpc_id     = aws_vpc.vpc.id
  cidr_block = infoblox_ipv4_network.ipv4_network.cidr
  ipv6_cidr_block =  infoblox_ipv6_network.ipv6_network.cidr
  availability_zone = "us-west-1a"
  assign_ipv6_address_on_creation = false
  map_public_ip_on_launch = false

  tags = {
    Name   = "tf-subnet"
    Subnet = "tf-subnet"
  }
}


#Create Network Interface
resource "aws_network_interface" "ni" {
  subnet_id   = aws_subnet.subnet.id
  private_ips = [infoblox_ipv4_allocation.ipv4_allocation.ip_addr]
  ipv6_addresses = [infoblox_ipv6_allocation.ipv6_allocation.ip_addr]

  tags = {
    Name = "tf-ni"
  }
}

# Create AWS Instance
resource "aws_instance" "ec2-instance" {
  # This ami is for us-west-1, change to Amazon Linux AMI for your region
  ami           = "ami-03130878b60947df3"
  instance_type = "t2.micro"

  network_interface {
    network_interface_id = aws_network_interface.ni.id
    device_index = 0
  }

  tags = {
    Name = "tf-ec2-instance"
  }
}