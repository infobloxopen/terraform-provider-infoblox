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