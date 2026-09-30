# NetworkList — uddi list cases
case "basic" {
  backend = "uddi"

  step {
    uddi {
      name       = "{{random}}"
      addr_block = [{ address = "{{random_public_ip}}/32" }]
    }
  }

  step {
    query    = true
    provider = infoblox
    limit    = 5
  }

}

case "filters" {
  backend = "uddi"

  step {
    uddi {
      name        = "{{random}}"
      addr_block  = [{ address = "{{random_public_ip}}/32" }]
      description = "Test Description"
    }
  }

  step {
    query            = true
    provider         = infoblox
    include_resource = true

    filter {
      type   = "filters"
      values = {
        name = "uddi.name"
      }
    }
  }

}
