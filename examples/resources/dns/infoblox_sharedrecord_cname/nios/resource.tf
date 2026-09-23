// Create a Shared Record Group (Required as Parent)
resource "infoblox_sharedrecordgroup" "parent_sharedrecord_group" {
  nios = {
    name = "example-sharedrecordgroup"
  }
}

// Create a Shared CNAME Record with Basic Fields
resource "infoblox_sharedrecord_cname" "shared_cname_record_with_basic_fields" {
  nios = {
    name                = "example-shared-record-cname"
    shared_record_group = infoblox_sharedrecordgroup.parent_sharedrecord_group.nios.name
    canonical           = "example.canonical.com"
  }
}

// Create a Shared CNAME Record with Additional Fields
resource "infoblox_sharedrecord_cname" "shared_cname_record_with_additional_fields" {
  nios = {
    name                = "example-shared-record-cname2"
    shared_record_group = infoblox_sharedrecordgroup.parent_sharedrecord_group.nios.name
    canonical           = "example.canonical2.com"

    // Additional Fields
    ext_attrs = {
      Site = "location-1"
    }
    comment = "Shared CNAME Record created by Terraform"
    disable = false
    ttl     = 3600
  }
}
