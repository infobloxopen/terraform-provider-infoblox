# Auto-generated datasource acceptance-test cases for SharedrecordAaaa.
case "filters" {
  backend           = "nios"
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_sharedrecordgroup" "parent_sharedrecord_group" {
    nios = {
      name = "{{random2}}"
    }
  }
  PREREQ

  filter {
    type = "filters"
    values = {
      name = "nios.name"
    }
  }

  pair_checks = ["nios.comment", "nios.disable", "nios.ipv6addr", "nios.name", "nios.shared_record_group", "nios.ttl", "nios.use_ttl"]

  step {
    nios {
      name                = "{{random}}"
      ipv6addr            = "2001:db8::1"
      shared_record_group = infoblox_sharedrecordgroup.parent_sharedrecord_group.nios.name
    }
  }

}

case "ext_attr_filters" {
  backend           = "nios"
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_sharedrecordgroup" "parent_sharedrecord_group" {
    nios = {
      name = "{{random2}}"
    }
  }
  PREREQ

  filter {
    type = "ext_attr_filters"
    values = {
      Site = "nios.ext_attrs.Site"
    }
  }

  pair_checks = ["nios.comment", "nios.disable", "nios.ipv6addr", "nios.name", "nios.shared_record_group", "nios.ttl", "nios.use_ttl"]

  step {
    nios {
      name                = "{{random}}.example.com"
      ipv6addr            = "2001:db8::1"
      shared_record_group = infoblox_sharedrecordgroup.parent_sharedrecord_group.nios.name
      ext_attrs           = { Site = "{{random3}}" }
    }
  }

}
