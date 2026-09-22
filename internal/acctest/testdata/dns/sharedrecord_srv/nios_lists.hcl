# Auto-generated list acceptance-test cases for SharedrecordSrv.
case "basic" {
  backend        = "nios"
  min_tf_version = "1.14.0"
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_sharedrecordgroup" "parent_sharedrecord_group" {
    nios = {
      name = "{{random3}}"
    }
  }
  PREREQ

  step {
    nios {
      name                = "{{random}}.example.com"
      port                = 10
      priority            = 80
      shared_record_group = infoblox_sharedrecordgroup.parent_sharedrecord_group.nios.name
      target              = "{{random2}}.target.com"
      weight              = 10
    }
  }

  step {
    query    = true
    provider = infoblox
    limit    = 5
  }

}

case "filters" {
  backend        = "nios"
  min_tf_version = "1.14.0"
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_sharedrecordgroup" "parent_sharedrecord_group" {
    nios = {
      name = "{{random3}}"
    }
  }
  PREREQ

  step {
    nios {
      name                = "{{random}}.example.com"
      port                = 10
      priority            = 80
      shared_record_group = infoblox_sharedrecordgroup.parent_sharedrecord_group.nios.name
      target              = "{{random2}}.target.com"
      weight              = 10
    }
  }

  step {
    query    = true
    provider = infoblox
    include_resource = true
    filter {
      type   = "filters"
      values = {
        name = "nios.name"
      }
    }
  }

}

case "ext_attr_filters" {
  backend        = "nios"
  min_tf_version = "1.14.0"
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_sharedrecordgroup" "parent_sharedrecord_group" {
    nios = {
      name = "{{random4}}"
    }
  }
  PREREQ

  step {
    nios {
      name                = "{{random2}}.example.com"
      port                = 10
      priority            = 80
      shared_record_group = infoblox_sharedrecordgroup.parent_sharedrecord_group.nios.name
      target              = "{{random3}}.target.com"
      weight              = 10
      ext_attrs           = { Site = "{{random}}" }
    }
  }

  step {
    query    = true
    provider = infoblox
    include_resource = true
    filter {
      type   = "ext_attr_filters"
      values = {
        Site = "nios.ext_attrs.Site"
      }
    }
  }

}
