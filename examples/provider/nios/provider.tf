terraform {
  required_providers {
    infoblox = {
      source  = "infobloxopen/infoblox"
      version = "3.0.0-beta"
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