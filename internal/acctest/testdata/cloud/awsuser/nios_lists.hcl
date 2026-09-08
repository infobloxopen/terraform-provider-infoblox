# Awsuser — nios list cases
case "basic" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      access_key_id     = "AKIA{{random}}"
      account_id        = "337773173961"
      name              = "{{random2}}"
      secret_access_key = "S1JGWfwcZWEYhSkfpyhxigL9A/uaJ6mY"
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
      access_key_id    = "AKIA{{random}}"
      account_id       = "337773173961"
      name             = "{{random2}}"
      # tostring() makes this a RawExpr so the framework skips the auto pair-check (writeOnly field)
      secret_access_key = tostring("S1JGWfwcZWEYhSkfpyhxigL9A/uaJ6mY")
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
