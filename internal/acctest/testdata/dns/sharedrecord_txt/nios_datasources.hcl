# Auto-generated datasource acceptance-test cases for SharedrecordTxt.
#
# TODO: These cases use the shared record group "shared_group", which must already
#       exist on the grid. The generated prerequisite is commented out because
#       infoblox_shared_record_group is not implemented in the provider yet.
#       Once it is, restore the prerequisite block and remove this note.
case "filters" {
  backend = "nios"
  # prerequisites_hcl = <<-PREREQ
  # resource "infoblox_shared_record_group_unknown" "parent_sharedrecord_group" {
  #   nios = {
  #     name = "This is a shared record TXT record"
  #   }
  # }
  # PREREQ

  filter {
    type   = "filters"
    values = {
      name = "nios.name"
    }
  }

  pair_checks = ["nios.comment", "nios.disable", "nios.name", "nios.shared_record_group", "nios.text", "nios.ttl", "nios.use_ttl"]

  step {
    nios {
      name                = "{{random}}"
      text                = "{{random2}}"
      shared_record_group = "shared_group"
    }
  }

}

case "ext_attr_filters" {
  backend = "nios"
  # prerequisites_hcl = <<-PREREQ
  # resource "infoblox_shared_record_group_unknown" "parent_sharedrecord_group" {
  #   nios = {
  #     name = "{{random2}}"
  #   }
  # }
  # PREREQ

  filter {
    type   = "ext_attr_filters"
    values = {
      Site = "nios.ext_attrs.Site"
    }
  }

  pair_checks = ["nios.comment", "nios.disable", "nios.name", "nios.shared_record_group", "nios.text", "nios.ttl", "nios.use_ttl"]

  step {
    nios {
      name                = "{{random}}"
      shared_record_group = "shared_group"
      text                = "This is a shared record TXT record"
      ext_attrs           = { Site = "{{random}}" }
    }
  }

}
