# Auto-generated datasource acceptance-test cases for SharedrecordCname.
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

  pair_checks = ["nios.canonical", "nios.comment", "nios.disable", "nios.name", "nios.shared_record_group", "nios.ttl"]

  step {
    nios {
      name                = "{{random}}"
      canonical           = "{{random2}}.com"
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

  pair_checks = ["nios.canonical", "nios.comment", "nios.disable", "nios.name", "nios.shared_record_group", "nios.ttl", "nios.use_ttl"]

  step {
    nios {
      name                = "{{random}}"
      canonical           = "{{random2}}.com"
      shared_record_group = infoblox_sharedrecordgroup.parent_sharedrecord_group.nios.name
      ext_attrs           = { Site = "{{random}}" }
    }
  }

}
