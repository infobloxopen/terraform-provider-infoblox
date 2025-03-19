resource "infoblox_ns_record" "ns1"{
  name = "test.com"
  nameserver = "name.test.com"
   addresses{
     address = "2.3.2.5"
     auto_create_ptr=true
   }
  addresses{
    address = "2.3.23.3"
    auto_create_ptr=false
  }
  addresses{
    address = "2.3.1.2"
    auto_create_ptr=true
  }
  dns_view = "default"
}

//IPV4 reverse mapping zone
resource "infoblox_zone_auth" "zone_test" {
  fqdn = "10.0.0.0/24"
  view = "default"
  zone_format = "IPV4"
}
//ns record created in a reverse mapping zone
resource "infoblox_ns_record" "ns3"{
  name = "0.0.10.in-addr.arpa"
  nameserver = "name.test.com"
  addresses{
    address = "2.3.2.5"
    auto_create_ptr=true
  }
  dns_view = "default"
  depends_on = ["infoblox_zone_auth.zone_test"]
}