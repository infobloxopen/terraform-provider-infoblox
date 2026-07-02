terraform {
  required_providers {
    unified = {
      source  = "infobloxopen/unified"
      version = "0.0.1"
    }
  }
}

provider "unified" {
  uddi = {
    csp_url = "<CSP_URL>"
    api_key = "<API_KEY>"
  }
}