# Auto-generated list acceptance-test cases for SharedrecordA.
#
# TODO: These cases use the shared record group "shared_group", which must already
#       exist on the grid. The generated prerequisite is commented out because
#       infoblox_shared_record_group is not implemented in the provider yet.
#       Once it is, restore the prerequisite block and remove this note.
case "basic" {
  backend        = "nios"
  min_tf_version = "1.14.0"
  # prerequisites_hcl = <<-PREREQ
  # resource "infoblox_shared_record_group_unknown" "parent_sharedrecord_group" {
  #   nios = {
  #     name = "{{random2}}"
  #   }
  # }
  # PREREQ

  step {
    nios {
      name                = "{{random}}"
      ipv4addr            = "10.0.0.0"
      shared_record_group = "shared_group"
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
  # prerequisites_hcl = <<-PREREQ
  # resource "infoblox_shared_record_group_unknown" "parent_sharedrecord_group" {
  #   nios = {
  #     name = "{{random2}}"
  #   }
  # }
  # PREREQ

  step {
    nios {
      name                = "{{random}}"
      ipv4addr            = "10.0.0.0"
      shared_record_group = "shared_group"
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
  # prerequisites_hcl = <<-PREREQ
  # resource "infoblox_shared_record_group_unknown" "parent_sharedrecord_group" {
  #   nios = {
  #     name = "{{random3}}"
  #   }
  # }
  # PREREQ

  step {
    nios {
      name                = "{{random2}}.example.com"
      ipv4addr            = "10.0.0.1"
      shared_record_group = "shared_group"
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
