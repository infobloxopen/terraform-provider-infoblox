terraform {
  required_providers {
    infoblox = {
      source  = "infobloxopen/infoblox"
      version = "3.0.0-beta1"
    }
  }
}

provider "infoblox" {
  uddi = {
    portal_url = "<INFOBLOX_PORTAL_URL>"
    portal_key = "<INFOBLOX_PORTAL_KEY>"
  }
}