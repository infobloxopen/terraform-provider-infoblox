# TODO: The following prerequisites MUST exist on the grid before running these tests:
#   - admin-group : admingroup with name "admin-group"

case "basic" {
  backend  = "nios"
  parallel = true
  step {
    nios {
      name         = "{{random}}"
      password     = "Example-Admin123!"
      admin_groups = ["admin-group"]
    }
  }
  step {
    query    = true
    provider = infoblox
    limit    = 5
  }
}

case "filters" {
  backend  = "nios"
  parallel = true
  step {
    nios {
      name         = "{{random}}"
      password     = tostring("Example-Admin123!")
      admin_groups = ["admin-group"]
    }
  }
  step {
    query            = true
    provider         = infoblox
    include_resource = true
    filter {
      type   = "filters"
      values = { name = "nios.name" }
    }
  }
}

case "ext_attr_filters" {
  backend  = "nios"
  parallel = true
  step {
    nios {
      name         = "{{random}}"
      password     = tostring("Example-Admin123!")
      admin_groups = ["admin-group"]
      ext_attrs    = { Site = "{{random2}}" }
    }
  }
  step {
    query            = true
    provider         = infoblox
    include_resource = true
    filter {
      type   = "ext_attr_filters"
      values = { Site = "nios.ext_attrs.Site" }
    }
  }
}
