# Region being used to create the resources
provider "aws" {
  region  = "us-west-1"
}

# Create a Virtual Private Cloud
resource "aws_vpc" "vpc" {
  cidr_block = "10.0.0.0/16"
  tags = {
    Name = "tf-vpc"
  }
}

# Create a Subnet
resource "aws_subnet" "subnet" {
  vpc_id     = aws_vpc.vpc.id
  cidr_block = infoblox_network.ib_network.cidr
  availability_zone = "us-west-1b"

  tags = {
    Name   = "tf-subnet"
    Subnet = "tf-subnet"
  }
}

# Create Network Interface
resource "aws_network_interface" "ni" {
  subnet_id   = aws_subnet.subnet.id
  private_ips = [infoblox_ip_allocation.ib_ip_allocation.ip_addr]

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
    Name = infoblox_ip_allocation.ib_ip_allocation.vm_name
  }
}
