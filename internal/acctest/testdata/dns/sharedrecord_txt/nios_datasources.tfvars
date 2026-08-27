# Auto-generated datasource acceptance-test cases for SharedrecordTxt.
case "filters" {
  backend = "nios"
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_shared_record_group_unknown" "parent_sharedrecord_group" {
    nios = {
      name = "This is a shared record TXT record"
    }
  }
  PREREQ

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
      shared_record_group = infoblox_shared_record_group_unknown.parent_sharedrecord_group.nios.name
    }
  }

}

case "ext_attr_filters" {
  backend = "nios"
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_shared_record_group_unknown" "parent_sharedrecord_group" {
    nios = {
      name = "{{random2}}"
    }
  }
  PREREQ

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
      shared_record_group = infoblox_shared_record_group_unknown.parent_sharedrecord_group.nios.name
      text                = "This is a shared record TXT record"
      ext_attrs           = { Site = "{{random}}" }
    }
  }

}
