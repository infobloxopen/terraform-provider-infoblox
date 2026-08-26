# RecordRpzNaptr — nios list cases
# TODO: The following prerequisites MUST exist on the grid before running these tests:
#   - RPZ zone : tf-acc-rpz-naptr.example.com         (view: default)
#   - RPZ zone : tf-acc-rpz-naptr-custom.example.com  (view: tf-acc-rpz-naptr-view)
#   - DNS view : tf-acc-rpz-naptr-view

case "basic" {
  backend        = "nios"
  min_tf_version = "1.14.0"
  parallel       = true

  step {
    nios {
      name        = "{{random2}}.tf-acc-rpz-naptr.example.com"
      rp_zone     = "tf-acc-rpz-naptr.example.com"
      order       = 10
      preference  = 10
      replacement = "."
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
  parallel       = true

  step {
    nios {
      name        = "{{random2}}.tf-acc-rpz-naptr.example.com"
      rp_zone     = "tf-acc-rpz-naptr.example.com"
      order       = 10
      preference  = 10
      replacement = "."
    }
  }

  step {
    query            = true
    provider         = infoblox
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
  parallel       = true

  step {
    nios {
      name        = "{{random2}}.tf-acc-rpz-naptr.example.com"
      rp_zone     = "tf-acc-rpz-naptr.example.com"
      order       = 10
      preference  = 10
      replacement = "."
      ext_attrs   = { Site = "{{random3}}" }
    }
  }

  step {
    query            = true
    provider         = infoblox
    include_resource = true
    filter {
      type   = "ext_attr_filters"
      values = {
        Site = "nios.ext_attrs.Site"
      }
    }
  }

}
