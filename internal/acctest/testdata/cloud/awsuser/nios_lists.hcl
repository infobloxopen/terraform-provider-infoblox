# Awsuser — nios list cases
case "basic" {
  backend        = "nios"
  min_tf_version = "1.14.0"

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
  backend        = "nios"
  min_tf_version = "1.14.0"

  step {
    nios {
      access_key_id    = "AKIA{{random}}"
      account_id       = "337773173961"
      name             = "{{random2}}"
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
