terraform {
  required_providers {
    infoblox = {
      source  = "infobloxopen/infoblox"
      version = "0.0.1"
    }
  }
}

provider "infoblox" {
  uddi = {
    csp_url = "<CSP_URL>"
    api_key = "<API_KEY>"
  }
}