data "infoblox_record_srv" "get_srv_record_using_filters" {
  filters = {
    name = "_sip._tcp.example.com"
  }
}

data "infoblox_record_srv" "get_all_srv_records" {}
