case "basic" {
  backend  = "uddi"
  parallel = true

  step {
    uddi {
      name    = "dtc-monitor-http-{{random}}"
      port    = 80
      request = "GET / HTTP/1.0\r\n\r\n"
    }
    check = {
      "uddi.name"       = "dtc-monitor-http-{{random}}"
      "uddi.port"       = "80"
      "uddi.disabled"   = "false"
      "uddi.https"      = "false"
      "uddi.interval"   = "15"
      "uddi.timeout"    = "10"
      "uddi.retry_down" = "1"
      "uddi.retry_up"   = "1"
    }
  }

}

case "disappears" {
  backend               = "uddi"
  disappears            = true
  expect_non_empty_plan = true
  parallel              = true

  step {
    uddi {
      name    = "dtc-monitor-http-{{random}}"
      port    = 80
      request = "GET / HTTP/1.0\r\n\r\n"
    }
  }

}

case "comment" {
  backend  = "uddi"
  parallel = true

  step {
    uddi {
      name    = "dtc-monitor-http-{{random}}"
      port    = 80
      request = "GET / HTTP/1.0\r\n\r\n"
      comment = "First comment"
    }
    check = {
      "uddi.comment" = "First comment"
    }
  }

  step {
    uddi {
      name    = "dtc-monitor-http-{{random}}"
      port    = 80
      request = "GET / HTTP/1.0\r\n\r\n"
      comment = "Updated comment"
    }
    check = {
      "uddi.comment" = "Updated comment"
    }
  }

}

case "disabled" {
  backend  = "uddi"
  parallel = true

  step {
    uddi {
      name     = "dtc-monitor-http-{{random}}"
      port     = 80
      request  = "GET / HTTP/1.0\r\n\r\n"
      disabled = false
    }
    check = {
      "uddi.disabled" = "false"
    }
  }

  step {
    uddi {
      name     = "dtc-monitor-http-{{random}}"
      port     = 80
      request  = "GET / HTTP/1.0\r\n\r\n"
      disabled = true
    }
    check = {
      "uddi.disabled" = "true"
    }
  }

}

case "https" {
  backend  = "uddi"
  parallel = true

  step {
    uddi {
      name    = "dtc-monitor-http-{{random}}"
      port    = 80
      request = "GET / HTTP/1.0\r\n\r\n"
      https   = false
    }
    check = {
      "uddi.https" = "false"
    }
  }

  step {
    uddi {
      name    = "dtc-monitor-http-{{random}}"
      port    = 443
      request = "GET / HTTP/1.0\r\n\r\n"
      https   = true
    }
    check = {
      "uddi.https" = "true"
    }
  }

}

case "interval" {
  backend  = "uddi"
  parallel = true

  step {
    uddi {
      name     = "dtc-monitor-http-{{random}}"
      port     = 80
      request  = "GET / HTTP/1.0\r\n\r\n"
      interval = 30
    }
    check = {
      "uddi.interval" = "30"
    }
  }

  step {
    uddi {
      name     = "dtc-monitor-http-{{random}}"
      port     = 80
      request  = "GET / HTTP/1.0\r\n\r\n"
      interval = 60
    }
    check = {
      "uddi.interval" = "60"
    }
  }

}

case "timeout" {
  backend  = "uddi"
  parallel = true

  step {
    uddi {
      name     = "dtc-monitor-http-{{random}}"
      port     = 80
      request  = "GET / HTTP/1.0\r\n\r\n"
      interval = 30
      timeout  = 5
    }
    check = {
      "uddi.timeout" = "5"
    }
  }

  step {
    uddi {
      name     = "dtc-monitor-http-{{random}}"
      port     = 80
      request  = "GET / HTTP/1.0\r\n\r\n"
      interval = 30
      timeout  = 20
    }
    check = {
      "uddi.timeout" = "20"
    }
  }

}

case "retry_down" {
  backend  = "uddi"
  parallel = true

  step {
    uddi {
      name       = "dtc-monitor-http-{{random}}"
      port       = 80
      request    = "GET / HTTP/1.0\r\n\r\n"
      retry_down = 2
    }
    check = {
      "uddi.retry_down" = "2"
    }
  }

  step {
    uddi {
      name       = "dtc-monitor-http-{{random}}"
      port       = 80
      request    = "GET / HTTP/1.0\r\n\r\n"
      retry_down = 3
    }
    check = {
      "uddi.retry_down" = "3"
    }
  }

}

case "retry_up" {
  backend  = "uddi"
  parallel = true

  step {
    uddi {
      name     = "dtc-monitor-http-{{random}}"
      port     = 80
      request  = "GET / HTTP/1.0\r\n\r\n"
      retry_up = 2
    }
    check = {
      "uddi.retry_up" = "2"
    }
  }

  step {
    uddi {
      name     = "dtc-monitor-http-{{random}}"
      port     = 80
      request  = "GET / HTTP/1.0\r\n\r\n"
      retry_up = 3
    }
    check = {
      "uddi.retry_up" = "3"
    }
  }

}

case "request" {
  backend  = "uddi"
  parallel = true

  step {
    uddi {
      name    = "dtc-monitor-http-{{random}}"
      port    = 80
      request = "GET /health HTTP/1.1\r\nHost: example.com\r\n\r\n"
    }
    check = {
      "uddi.request" = "GET /health HTTP/1.1\r\nHost: example.com\r\n\r\n"
    }
  }

  step {
    uddi {
      name    = "dtc-monitor-http-{{random}}"
      port    = 80
      request = "HEAD /status HTTP/1.1\r\nHost: example.com\r\n\r\n"
    }
    check = {
      "uddi.request" = "HEAD /status HTTP/1.1\r\nHost: example.com\r\n\r\n"
    }
  }

}

case "codes" {
  backend  = "uddi"
  parallel = true

  step {
    uddi {
      name    = "dtc-monitor-http-{{random}}"
      port    = 80
      request = "GET / HTTP/1.0\r\n\r\n"
      codes   = "200,201"
    }
    check = {
      "uddi.codes" = "200,201"
    }
  }

  step {
    uddi {
      name    = "dtc-monitor-http-{{random}}"
      port    = 80
      request = "GET / HTTP/1.0\r\n\r\n"
      codes   = "200,201,204"
    }
    check = {
      "uddi.codes" = "200,201,204"
    }
  }

}

case "check_response_body" {
  backend  = "uddi"
  parallel = true

  step {
    uddi {
      name                      = "dtc-monitor-http-{{random}}"
      port                      = 80
      request                   = "GET / HTTP/1.0\r\n\r\n"
      check_response_body       = true
      check_response_body_regex = "OK"
    }
    check = {
      "uddi.check_response_body"       = "true"
      "uddi.check_response_body_regex" = "OK"
    }
  }

  step {
    uddi {
      name                      = "dtc-monitor-http-{{random}}"
      port                      = 80
      request                   = "GET / HTTP/1.0\r\n\r\n"
      check_response_body       = true
      check_response_body_regex = "SUCCESS"
    }
    check = {
      "uddi.check_response_body"       = "true"
      "uddi.check_response_body_regex" = "SUCCESS"
    }
  }

}

case "check_response_header" {
  backend  = "uddi"
  parallel = true

  step {
    uddi {
      name                  = "dtc-monitor-http-{{random}}"
      port                  = 80
      request               = "GET / HTTP/1.0\r\n\r\n"
      check_response_header = true
      check_response_header_regexes = [
        { header = "Content-Type", regex = "application/json" }
      ]
    }
    check = {
      "uddi.check_response_header"                  = "true"
      "uddi.check_response_header_regexes.#"        = "1"
      "uddi.check_response_header_regexes.0.header" = "Content-Type"
      "uddi.check_response_header_regexes.0.regex"  = "application/json"
    }
  }

  step {
    uddi {
      name                           = "dtc-monitor-http-{{random}}"
      port                           = 80
      request                        = "GET / HTTP/1.0\r\n\r\n"
      check_response_header          = true
      check_response_header_negative = true
      check_response_header_regexes = [
        { header = "Content-Type", regex = "text/plain" },
        { header = "X-Custom-Header", regex = "value" }
      ]
    }
    check = {
      "uddi.check_response_header"          = "true"
      "uddi.check_response_header_negative" = "true"
      "uddi.check_response_header_regexes.#" = "2"
    }
  }

}

case "name" {
  backend  = "uddi"
  parallel = true

  step {
    uddi {
      name    = "dtc-monitor-http-{{random}}"
      port    = 80
      request = "GET / HTTP/1.0\r\n\r\n"
    }
    check = {
      "uddi.name" = "dtc-monitor-http-{{random}}"
    }
  }

  step {
    uddi {
      name    = "dtc-monitor-http-{{random2}}"
      port    = 80
      request = "GET / HTTP/1.0\r\n\r\n"
    }
    check = {
      "uddi.name" = "dtc-monitor-http-{{random2}}"
    }
  }

}

case "port" {
  backend  = "uddi"
  parallel = true

  step {
    uddi {
      name    = "dtc-monitor-http-{{random}}"
      port    = 8080
      request = "GET / HTTP/1.0\r\n\r\n"
    }
    check = {
      "uddi.port" = "8080"
    }
  }

  step {
    uddi {
      name    = "dtc-monitor-http-{{random}}"
      port    = 8443
      request = "GET / HTTP/1.0\r\n\r\n"
    }
    check = {
      "uddi.port" = "8443"
    }
  }

}

case "tags" {
  backend  = "uddi"
  parallel = true

  step {
    uddi {
      name    = "dtc-monitor-http-{{random}}"
      port    = 80
      request = "GET / HTTP/1.0\r\n\r\n"
      tags    = { Site = "{{random2}}" }
    }
    check = {
      "uddi.tags.Site" = "{{random2}}"
    }
  }

  step {
    uddi {
      name    = "dtc-monitor-http-{{random}}"
      port    = 80
      request = "GET / HTTP/1.0\r\n\r\n"
      tags    = { Site = "{{random3}}" }
    }
    check = {
      "uddi.tags.Site" = "{{random3}}"
    }
  }

}
