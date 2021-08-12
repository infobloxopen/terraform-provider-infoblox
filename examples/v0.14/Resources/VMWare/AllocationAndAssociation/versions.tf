terraform {
  required_providers {
    infoblox = {
      source = "infobloxopen/infoblox"
      version = ">=2.0"
    }
    vsphere = {
      source = "hashicorp/vsphere"
      version = "1.12.0"
    }
  }
}
