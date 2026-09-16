// Create a Shared Record Group with Basic Fields
resource "infoblox_sharedrecordgroup" "shared_record_group_with_basic_fields" {
  nios = {
    name = "example-shared-record-group"
  }
}

// Create a Shared Record Group with Additional Fields
resource "infoblox_sharedrecordgroup" "shared_record_group_with_additional_fields" {
  nios = {
    name = "example-shared-record-group-2"

    // Additional Fields
    ext_attrs = {
      Site = "location-1"
    }
    record_name_policy = "Allow Any"
    comment            = "Shared Record Group created by Terraform"
  }
}
