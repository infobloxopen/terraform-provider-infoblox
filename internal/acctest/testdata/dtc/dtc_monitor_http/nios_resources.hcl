# TODO: grid prereqs — two dtc:certificate objects must exist on the test appliance
# Auto-generated resource acceptance-test cases for DtcMonitorHttp.
case "basic" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name = "{{random}}"
    }
    check = {
      "nios.name"                  = "{{random}}"
      "nios.content_check"         = "NONE"
      "nios.content_check_input"   = "ALL"
      "nios.content_extract_group" = "0"
      "nios.content_extract_type"  = "STRING"
      "nios.enable_sni"            = "false"
      "nios.interval"              = "5"
      "nios.port"                  = "80"
      "nios.result"                = "ANY"
      "nios.result_code"           = "200"
      "nios.retry_down"            = "1"
      "nios.retry_up"              = "1"
      "nios.secure"                = "false"
      "nios.timeout"               = "15"
      "nios.validate_cert"         = "true"
    }
  }

}

case "disappears" {
  backend               = "nios"
  disappears            = true
  expect_non_empty_plan = true
  parallel              = true

  step {
    nios {
      name = "{{random}}"
    }
  }

}

case "ciphers" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name    = "{{random}}"
      ciphers = "DHE-RSA-AES256-SHA"
    }
    check = {
      "nios.ciphers" = "DHE-RSA-AES256-SHA"
    }
  }

  step {
    nios {
      name    = "{{random}}"
      ciphers = "DEFAULT"
    }
    check = {
      "nios.ciphers" = "DEFAULT"
    }
  }

}

case "client_cert" {
  # TODO: grid prereqs — dtc:certificate refs 0504e596...abad058a and 4b0bcb4d...55ad5c must exist on the test appliance
  backend  = "nios"
  parallel = true

  step {
    nios {
      name        = "{{random}}"
      client_cert = "dtc:certificate/ZG5zLmlkbnNfY2VydGlmaWNhdGUkYjIzMDM0NDhhODlhMjRmMGNlNWE3OTRiMWRiYWI2NjkwM2IyNmYwNGY0MzNhODA4YmExZDNiNjY3NzU2NTY5NTA5MmJjYTAzZTA0MjIxN2ZkOWFlOWI4YzE1N2I0MmQyNWEzYWJjNzA4MGZiYWRiYWRmY2I3NjkwYzgxN2NlODY:0504e596e496491145f9946315092522abad058a"
    }
    check = {
      "nios.client_cert" = "dtc:certificate/ZG5zLmlkbnNfY2VydGlmaWNhdGUkYjIzMDM0NDhhODlhMjRmMGNlNWE3OTRiMWRiYWI2NjkwM2IyNmYwNGY0MzNhODA4YmExZDNiNjY3NzU2NTY5NTA5MmJjYTAzZTA0MjIxN2ZkOWFlOWI4YzE1N2I0MmQyNWEzYWJjNzA4MGZiYWRiYWRmY2I3NjkwYzgxN2NlODY:0504e596e496491145f9946315092522abad058a"
    }
  }

  step {
    nios {
      name        = "{{random}}"
      client_cert = "dtc:certificate/ZG5zLmlkbnNfY2VydGlmaWNhdGUkNmQxN2YzODc5MjYxZDhkM2U2NGM3MTEwYmU4ZTU3Nzk2ZjY5Y2EzYTc1ZmYxYzM0MWI3NTZkNGEyZWFkYWFmNWI4NzVhNjdlZmU4ZjU5ZjczZmRjODMyY2U5MTlhMzQzYmI1OGMyNTQxOGFkN2RmMWEyYTY3NzA5YWRlNWIxMGM:4b0bcb4d0f766414393f1d649bab3e6cb255ad5c"
    }
    check = {
      "nios.client_cert" = "dtc:certificate/ZG5zLmlkbnNfY2VydGlmaWNhdGUkNmQxN2YzODc5MjYxZDhkM2U2NGM3MTEwYmU4ZTU3Nzk2ZjY5Y2EzYTc1ZmYxYzM0MWI3NTZkNGEyZWFkYWFmNWI4NzVhNjdlZmU4ZjU5ZjczZmRjODMyY2U5MTlhMzQzYmI1OGMyNTQxOGFkN2RmMWEyYTY3NzA5YWRlNWIxMGM:4b0bcb4d0f766414393f1d649bab3e6cb255ad5c"
    }
  }

}

case "comment" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name    = "{{random}}"
      comment = "This is a comment"
    }
    check = {
      "nios.comment" = "This is a comment"
    }
  }

  step {
    nios {
      name    = "{{random}}"
      comment = "This comment is updated"
    }
    check = {
      "nios.comment" = "This comment is updated"
    }
  }

}

case "content_check" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name                  = "{{random}}"
      content_check         = "EXTRACT"
      content_check_op      = "EQ"
      content_check_regex   = "The current load is ([0-9]+)"
      content_extract_type  = "STRING"
      content_extract_value = "default value"
    }
    check = {
      "nios.content_check" = "EXTRACT"
    }
  }

  step {
    nios {
      name                  = "{{random}}"
      content_check         = "MATCH"
      content_check_op      = "EQ"
      content_check_regex   = "The current load is ([0-9]+)"
      content_extract_type  = "STRING"
      content_extract_value = "default value"
    }
    check = {
      "nios.content_check" = "MATCH"
    }
  }

}

case "content_check_input" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name                = "{{random}}"
      content_check_input = "BODY"
    }
    check = {
      "nios.content_check_input" = "BODY"
    }
  }

  step {
    nios {
      name                = "{{random}}"
      content_check_input = "HEADERS"
    }
    check = {
      "nios.content_check_input" = "HEADERS"
    }
  }

}

case "content_check_op" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name             = "{{random}}"
      content_check_op = "GEQ"
    }
    check = {
      "nios.content_check_op" = "GEQ"
    }
  }

  step {
    nios {
      name             = "{{random}}"
      content_check_op = "LEQ"
    }
    check = {
      "nios.content_check_op" = "LEQ"
    }
  }

  step {
    nios {
      name             = "{{random}}"
      content_check_op = "EQ"
    }
    check = {
      "nios.content_check_op" = "EQ"
    }
  }

  step {
    nios {
      name             = "{{random}}"
      content_check_op = "NEQ"
    }
    check = {
      "nios.content_check_op" = "NEQ"
    }
  }

}

case "content_check_regex" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name                = "{{random}}"
      content_check_regex = "HTTP/1\\\\.[01] (200|201|204)"
    }
    check = {
      "nios.content_check_regex" = "HTTP/1\\\\.[01] (200|201|204)"
    }
  }

  step {
    nios {
      name                = "{{random}}"
      content_check_regex = "Status: (2[0-9]{2})"
    }
    check = {
      "nios.content_check_regex" = "Status: (2[0-9]{2})"
    }
  }

}

case "content_extract_group" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name                  = "{{random}}"
      content_extract_group = 5
    }
    check = {
      "nios.content_extract_group" = "5"
    }
  }

  step {
    nios {
      name                  = "{{random}}"
      content_extract_group = 8
    }
    check = {
      "nios.content_extract_group" = "8"
    }
  }

}

case "content_extract_type" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name                 = "{{random}}"
      content_extract_type = "STRING"
    }
    check = {
      "nios.content_extract_type" = "STRING"
    }
  }

  step {
    nios {
      name                 = "{{random}}"
      content_extract_type = "INTEGER"
    }
    check = {
      "nios.content_extract_type" = "INTEGER"
    }
  }

}

case "content_extract_value" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name                  = "{{random}}"
      content_extract_value = "SUCCESS"
    }
    check = {
      "nios.content_extract_value" = "SUCCESS"
    }
  }

  step {
    nios {
      name                  = "{{random}}"
      content_extract_value = "ACTIVE"
    }
    check = {
      "nios.content_extract_value" = "ACTIVE"
    }
  }

}

case "enable_sni" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name       = "{{random}}"
      enable_sni = false
    }
    check = {
      "nios.enable_sni" = "false"
    }
  }

  step {
    nios {
      name       = "{{random}}"
      enable_sni = true
    }
    check = {
      "nios.enable_sni" = "true"
    }
  }

}

case "ext_attrs" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name      = "{{random}}"
      ext_attrs = { Site = "{{random2}}" }
    }
    check = {
      "nios.ext_attrs.Site" = "{{random2}}"
    }
  }

  step {
    nios {
      name      = "{{random}}"
      ext_attrs = { Site = "{{random3}}" }
    }
    check = {
      "nios.ext_attrs.Site" = "{{random3}}"
    }
  }

}

case "interval" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name     = "{{random}}"
      interval = 4
    }
    check = {
      "nios.interval" = "4"
    }
  }

  step {
    nios {
      name     = "{{random}}"
      interval = 10
    }
    check = {
      "nios.interval" = "10"
    }
  }

}

case "name" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name = "{{random}}"
    }
    check = {
      "nios.name" = "{{random}}"
    }
  }

  step {
    nios {
      name = "{{random2}}"
    }
    check = {
      "nios.name" = "{{random2}}"
    }
  }

}

case "port" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name = "{{random}}"
      port = 80
    }
    check = {
      "nios.port" = "80"
    }
  }

  step {
    nios {
      name = "{{random}}"
      port = 8080
    }
    check = {
      "nios.port" = "8080"
    }
  }

}

case "request" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name    = "{{random}}"
      request = "GET /api/health HTTP/1.1\nHost: example.com\nUser-Agent: NIOS-Monitor"
    }
    check = {
      "nios.request" = "GET /api/health HTTP/1.1\nHost: example.com\nUser-Agent: NIOS-Monitor"
    }
  }

  step {
    nios {
      name    = "{{random}}"
      request = "HEAD /resource HTTP/1.1\nHost: example.com\nAccept: */*"
    }
    check = {
      "nios.request" = "HEAD /resource HTTP/1.1\nHost: example.com\nAccept: */*"
    }
  }

  step {
    nios {
      name    = "{{random}}"
      request = "POST /submit HTTP/1.1\nHost: example.com\nContent-Type: application/json"
    }
    check = {
      "nios.request" = "POST /submit HTTP/1.1\nHost: example.com\nContent-Type: application/json"
    }
  }

}

case "result" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name   = "{{random}}"
      result = "CODE_IS"
    }
    check = {
      "nios.result" = "CODE_IS"
    }
  }

  step {
    nios {
      name   = "{{random}}"
      result = "CODE_IS_NOT"
    }
    check = {
      "nios.result" = "CODE_IS_NOT"
    }
  }

}

case "result_code" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name        = "{{random}}"
      result_code = 200
    }
    check = {
      "nios.result_code" = "200"
    }
  }

  step {
    nios {
      name        = "{{random}}"
      result_code = 404
    }
    check = {
      "nios.result_code" = "404"
    }
  }

}

case "retry_down" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name       = "{{random}}"
      retry_down = 4
    }
    check = {
      "nios.retry_down" = "4"
    }
  }

  step {
    nios {
      name       = "{{random}}"
      retry_down = 5
    }
    check = {
      "nios.retry_down" = "5"
    }
  }

}

case "retry_up" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name     = "{{random}}"
      retry_up = 4
    }
    check = {
      "nios.retry_up" = "4"
    }
  }

  step {
    nios {
      name     = "{{random}}"
      retry_up = 5
    }
    check = {
      "nios.retry_up" = "5"
    }
  }

}

case "secure" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name   = "{{random}}"
      secure = false
    }
    check = {
      "nios.secure" = "false"
    }
  }

  step {
    nios {
      name   = "{{random}}"
      secure = true
    }
    check = {
      "nios.secure" = "true"
    }
  }

}

case "timeout" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name    = "{{random}}"
      timeout = 20
    }
    check = {
      "nios.timeout" = "20"
    }
  }

  step {
    nios {
      name    = "{{random}}"
      timeout = 30
    }
    check = {
      "nios.timeout" = "30"
    }
  }

}

case "validate_cert" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name          = "{{random}}"
      validate_cert = true
    }
    check = {
      "nios.validate_cert" = "true"
    }
  }

  step {
    nios {
      name          = "{{random}}"
      validate_cert = false
    }
    check = {
      "nios.validate_cert" = "false"
    }
  }

}
