// Create a NS Group Stub Member with Basic Fields
resource "infoblox_nsgroup_stubmember" "nsgroup_stubmember_with_basic_fields" {
  nios = {
    name = "stubmember1"
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
    name = "stubmember2"
    stub_members = [
      {
        name = "member.com"
      }
    ]
    comment = "This is a comment"
    ext_attrs = {
      Site = "location-1"
    }
  }
}
