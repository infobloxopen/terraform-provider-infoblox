# Security Policy

## Reporting a Vulnerability

Please do not report security vulnerabilities through public GitHub issues, pull requests or discussions.

Report suspected vulnerabilities in this provider, or in the Go clients it uses, to the Infoblox Product Security team at [PSIRT@infoblox.com](mailto:PSIRT@infoblox.com). Include the provider version, the backend (NIOS, UDDI, or NIOS through the Infoblox Portal), steps to reproduce, and the impact you observed.

Reports are handled under the [Infoblox Vulnerability Responsible Disclosure Policy](https://www.infoblox.com/company/legal/vulnerability-responsible-disclosure-policy/).

## Supported Versions

Security fixes are released for the latest release of the provider. Upgrade to the latest version before reporting an issue.

## Handling Credentials

Never commit NIOS passwords, Infoblox Portal API keys or other credentials in Terraform configuration. Supply them through environment variables (for example `NIOS_PASSWORD` and `INFOBLOX_PORTAL_KEY`) or a secrets manager, and mark Terraform variables that hold them as `sensitive`.
