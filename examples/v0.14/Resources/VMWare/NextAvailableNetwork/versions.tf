terraform {
  required_providers {
    infoblox = {
      source  = "terraform-providers/infoblox"
      version = ">= 1.0"
    }
    vsphere = {
      source = "hashicorp/vsphere"
      version = "1.12.0"
    }
  }
}
