terraform {
  required_providers {
    infoblox = {
      source  = "infobloxopen/infoblox"
      version = "0.0.1"
    }
  }
}

provider "infoblox" {
  nios = {
    host_url = "<NIOS_HOST_URL>"
    username = "<NIOS_USERNAME>"
    password = "<NIOS_PASSWORD>"
  }
}