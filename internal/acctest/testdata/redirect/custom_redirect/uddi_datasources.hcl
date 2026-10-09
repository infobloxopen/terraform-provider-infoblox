# Note: filtering on custom_redirects is broken at the API level — tracked in TDDFW-435.
case "filters" {
  backend     = "uddi"
  skip        = true
  skip_reason = "Filtering on custom_redirects returns 401 Unauthorized — API-level permission issue. See TDDFW-435."
}
