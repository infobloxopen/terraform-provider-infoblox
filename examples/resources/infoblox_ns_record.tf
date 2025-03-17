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
