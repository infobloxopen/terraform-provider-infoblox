terraform {
  required_providers {
    unified = {
      source  = "infobloxopen/unified"
      version = "0.0.1"
    }
  }
}

provider "unified" {
  nios = {
    host_url = "<NIOS_HOST_URL>"
    username = "<NIOS_USERNAME>"
    password = "<NIOS_PASSWORD>"
  }
}