// Create a DTC Monitor HTTP with required fields only
resource "infoblox_dtc_monitor_http" "example" {
  nios = {
    name = "example-monitor-http"
  }
}

// Create a DTC Monitor HTTP with all optional fields
resource "infoblox_dtc_monitor_http" "example_all_fields" {
  nios = {
    name                  = "example-monitor-http-all-fields"
    ciphers               = "DHE-RSA-AES256-SHA"
    comment               = "Example DTC Monitor HTTP"
    content_check         = "EXTRACT"
    content_check_input   = "BODY"
    content_check_op      = "EQ"
    content_check_regex   = "Load: ([0-9]+)"
    content_extract_group = 1
    content_extract_type  = "STRING"
    content_extract_value = "SUCCESS"
    enable_sni            = false
    interval              = 5
    port                  = 80
    request               = "GET /api/health HTTP/1.1\nHost: example.com"
    result                = "CODE_IS"
    result_code           = 200
    retry_down            = 2
    retry_up              = 5
    secure                = false
    timeout               = 30
    validate_cert         = false
    ext_attrs = {
      Site = "us-east-1"
    }
  }
}
