case "filters" {
  backend     = "uddi"
  skip        = true
  skip_reason = "The redirect API returns 401 Unauthorized when using _filter on /api/atcfw/v1/custom_redirects. This is an environment-level permission restriction on the stage CSP — filtering requires elevated permissions not available with the current API key."
}

