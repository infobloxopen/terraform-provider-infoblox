# Auto-generated datasource acceptance-test cases for SharedrecordSrv.
case "filters" {
  backend = "nios"
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_sharedrecordgroup" "parent_sharedrecord_group" {
    nios = {
      name = "{{random2}}"
    }
  }
  PREREQ

  filter {
    type   = "filters"
    values = {
      name = "nios.name"
    }
  }

  pair_checks = ["nios.comment", "nios.disable", "nios.name", "nios.port", "nios.priority", "nios.shared_record_group", "nios.target", "nios.ttl", "nios.weight"]

  step {
    nios {
      name                = "{{random}}.example.com"
      port                = 80
      priority            = 10
      shared_record_group = infoblox_sharedrecordgroup.parent_sharedrecord_group.nios.name
      target              = "{{random3}}.target.com"
      weight              = 10
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

  pair_checks = ["nios.comment", "nios.disable", "nios.name", "nios.port", "nios.priority", "nios.shared_record_group", "nios.target", "nios.ttl", "nios.use_ttl", "nios.weight"]

  step {
    nios {
      name                = "{{random}}.example.com"
      port                = 80
      priority            = 10
      shared_record_group = infoblox_sharedrecordgroup.parent_sharedrecord_group.nios.name
      target              = "{{random2}}.target.com"
      weight              = 10
      ext_attrs           = { Site = "{{random4}}" }
    }
  }

}
