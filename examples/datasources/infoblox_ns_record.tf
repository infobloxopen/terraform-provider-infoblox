resource "infoblox_ns_record" "ns"{
  name = "test.com"
  nameserver = "name.test.com"
  addresses{
    address = "2.3.2.5"
    auto_create_ptr=true
  }
  dns_view = "default"
}

data "infoblox_ns_record" "testNs_read" {
  filters = {
    name = infoblox_ns_record.ns.name
    view = infoblox_ns_record.ns.dns_view
    nameserver=infoblox_ns_record.ns.nameserver
  }
}

output "dtc_rec_res1" {
  value = data.infoblox_ns_record.testNs_read
}