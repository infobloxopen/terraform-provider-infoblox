// Create a Shared Record Group (Required as Parent)
resource "infoblox_sharedrecordgroup" "parent_sharedrecordgroup" {
  nios = {
    name = "example-sharedrecordgroup"
  }
}

// Create a Shared SRV Record with Basic Fields
resource "infoblox_sharedrecord_srv" "sharedrecord_srv_basic_fields" {
  nios = {
    name                = "sharedrecord_srv.example.com"
    port                = 443
    priority            = 10
    shared_record_group = infoblox_sharedrecordgroup.parent_sharedrecordgroup.nios.name
    target              = "server.example.com"
    weight              = 5
  }
}

// Create a Shared SRV Record with Additional Fields
resource "infoblox_sharedrecord_srv" "sharedrecord_srv_additional_fields" {
  nios = {
    name                = "_http._tcp.example.com"
    port                = 80
    priority            = 20
    shared_record_group = infoblox_sharedrecordgroup.parent_sharedrecordgroup.nios.name
    target              = "webserver.example.com"
    weight              = 10
    comment             = "Example Shared SRV Record"
    disable             = true
    ext_attrs = {
      Site = "location-1"
    }
    ttl = 7200
  }
}
