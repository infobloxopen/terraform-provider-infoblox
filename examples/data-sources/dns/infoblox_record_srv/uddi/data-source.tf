data "infoblox_record_srv" "get_srv_record_using_filters" {
  filters = {
    "name_in_zone" = "record_srv.example.com"
  }
}

data "infoblox_record_srv" "get_srv_record_using_tag_filters" {
  tag_filters = {
    Site = "location-1"
  }
}

data "infoblox_record_srv" "get_all_srv_records" {}

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
    csp_url = "https://stage.csp.infoblox.com"
    api_key = "4a815e1e1c86a208efab3a5bfcc6f1f73259c009c43d30b13b337786ca9b3328"
  }
}