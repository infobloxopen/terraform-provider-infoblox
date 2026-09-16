# Natgroup — nios list cases
case "basic" {
  backend        = "nios"
  min_tf_version = "1.14.0"

  step {
    nios {
      name = "{{random}}"
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

  step {
    nios {
      name    = "{{random}}"
      comment = "This is a test natgroup"
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
