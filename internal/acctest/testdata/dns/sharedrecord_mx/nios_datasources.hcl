# Auto-generated datasource acceptance-test cases for SharedrecordMx.
case "filters" {
  backend = "nios"
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_sharedrecordgroup" "parent_sharedrecord_group" {
    nios = {
      name = "{{random3}}"
    }
  }
  PREREQ

  filter {
    type   = "filters"
    values = {
      name = "nios.name"
    }
  }

  pair_checks = ["nios.comment", "nios.disable", "nios.mail_exchanger", "nios.name", "nios.preference", "nios.shared_record_group", "nios.ttl"]

  step {
    nios {
      mail_exchanger      = "{{random2}}.example.com"
      name                = "{{random}}.example.com"
      preference          = 10
      shared_record_group = infoblox_sharedrecordgroup.parent_sharedrecord_group.nios.name
    }
  }

}

case "ext_attr_filters" {
  backend = "nios"
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_sharedrecordgroup" "parent_sharedrecord_group" {
    nios = {
      name = "{{random3}}"
    }
  }
  PREREQ

  filter {
    type   = "ext_attr_filters"
    values = {
      Site = "nios.ext_attrs.Site"
    }
  }

  pair_checks = ["nios.comment", "nios.disable", "nios.mail_exchanger", "nios.name", "nios.preference", "nios.shared_record_group", "nios.ttl"]

  step {
    nios {
      mail_exchanger      = "{{random2}}.example.com"
      name                = "{{random}}.example.com"
      preference          = 10
      shared_record_group = infoblox_sharedrecordgroup.parent_sharedrecord_group.nios.name
      ext_attrs           = { Site = "{{random4}}" }
    }
  }

}
