// Create a NS Group Stub Member with Basic Fields
resource "infoblox_nsgroup_stubmember" "nsgroup_stubmember_with_basic_fields" {
  nios = {
    name = "example_stubmember"
    stub_members = [
      {
        name = "member.com"
      }
    ]
  }
}

// Create a NS Group Stub Member with Additional Fields
resource "infoblox_nsgroup_stubmember" "nsgroup_stubmember_with_additional_fields" {
  nios = {
    name = "example_stubmember_with_additional_fields"
    stub_members = [
      {
        name = "member.com"
      }
    ]
    comment = "Example comment for NS Group Stub Member"
    ext_attrs = {
      Site = "location-1"
    }
  }
}
