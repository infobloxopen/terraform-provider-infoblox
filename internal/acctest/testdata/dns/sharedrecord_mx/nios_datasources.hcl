# Auto-generated datasource acceptance-test cases for SharedrecordMx.
#
# TODO: These cases use the shared record group "shared_group", which must already
#       exist on the grid. The generated prerequisite is commented out because
#       infoblox_shared_record_group is not implemented in the provider yet.
#       Once it is, restore the prerequisite block and remove this note.
case "filters" {
  backend = "nios"
  # prerequisites_hcl = <<-PREREQ
  # resource "infoblox_shared_record_group_unknown" "parent_sharedrecord_group" {
  # nios = {
  # name = "{{random3}}"
  # }
  # }
  # PREREQ

  filter {
    type   = "filters"
    values = {
      name = "nios.name"
    }
  }

  pair_checks = ["nios.comment", "nios.disable", "nios.mail_exchanger", "nios.name", "nios.preference", "nios.shared_record_group", "nios.ttl", "nios.use_ttl"]

  step {
    nios {
      mail_exchanger      = "{{random2}}.example.com"
      name                = "{{random}}.example.com"
      preference          = 10
      shared_record_group = "shared_group"
    }
  }

}

case "ext_attr_filters" {
  backend = "nios"
  # prerequisites_hcl = <<-PREREQ
  # resource "infoblox_shared_record_group_unknown" "parent_sharedrecord_group" {
  # nios = {
  # name = "{{random3}}"
  # }
  # }
  # PREREQ

  filter {
    type   = "ext_attr_filters"
    values = {
      Site = "nios.ext_attrs.Site"
    }
  }

  pair_checks = ["nios.comment", "nios.disable", "nios.mail_exchanger", "nios.name", "nios.preference", "nios.shared_record_group", "nios.ttl", "nios.use_ttl"]

  step {
    nios {
      mail_exchanger      = "{{random2}}.example.com"
      name                = "{{random}}.example.com"
      preference          = 10
      shared_record_group = "shared_group"
      ext_attrs           = { Site = "{{random4}}" }
    }
  }

}
