// Create a DTC HTTP Health Check with required fields only
resource "infoblox_dtc_monitor_http" "example" {
  uddi = {
    name    = "example-http-monitor"
    port    = 80
    request = "GET / HTTP/1.0\r\n"
  }
}

// Create a DTC HTTP Health Check with all optional fields
resource "infoblox_dtc_monitor_http" "example_all_fields" {
  uddi = {
    name                         = "example-http-monitor-full"
    port                         = 443
    comment                      = "HTTPS health check for web servers"
    disabled                     = false
    https                        = true
    interval                     = 30
    timeout                      = 10
    retry_down                   = 2
    retry_up                     = 2
    request                      = "GET /health HTTP/1.1\r\nHost: example.com\r\n"
    codes                        = "200,201,204"
    check_response_body          = true
    check_response_body_regex    = "OK"
    check_response_body_negative = false
    check_response_header        = true
    check_response_header_regexes = [
      {
        header = "Content-Type"
        regex  = "application/json"
      }
    ]
    check_response_header_negative = false
    tags = {
      env  = "production"
      Site = "us-east-1"
    }
  }
}
