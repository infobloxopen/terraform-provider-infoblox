case "basic" {
  backend        = "uddi"
  parallel       = true
  min_tf_version = "1.14.0"

  step {
    uddi {
      name = "custom-redirect-{{random}}"
      data = "156.2.3.10"
    }
  }

  step {
    query    = true
    provider = infoblox
    limit    = 5
  }
}

# Note: filtering on custom_redirects is broken at the API level — tracked in TDDFW-435.
case "filters" {
  backend     = "uddi"
  skip        = true
  skip_reason = "Filtering on custom_redirects returns 401 Unauthorized — API-level permission issue. See TDDFW-435."
}
