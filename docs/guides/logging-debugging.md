---
page_title: "Logging and Debugging"
subcategory: "Guides"
description: |-
  Enable Terraform logs to debug issues with the Infoblox provider on the NIOS and UDDI backends.
---

# Logging and Debugging

## Terraform Logging

Terraform has detailed logs that can help debug provider issues. To enable them, set the `TF_LOG` environment variable to one of the log levels: `TRACE`, `DEBUG`, `INFO`, `WARN`, or `ERROR`:

```bash
# For Linux/macOS
export TF_LOG=DEBUG
terraform plan

# For Windows PowerShell
$env:TF_LOG="DEBUG"
terraform plan
```

The `TRACE` level is the most verbose and will include all API calls made by the provider to the configured backend, either the NIOS WAPI or the UDDI API.

### Provider-Specific Debugging

For debugging specific issues with the Infoblox provider:

1. Use `DEBUG` or `TRACE` log levels to see the API requests and responses
2. Check the request body and response status codes for API errors
3. Ensure correct credentials and permissions for the backend you configured

For more information on debugging Terraform providers, refer to the [Terraform debugging documentation](https://developer.hashicorp.com/terraform/internals/debugging).
