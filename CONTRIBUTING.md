# Contributing to the Terraform Provider for Infoblox

Thank you for your interest in improving the Terraform Provider for Infoblox. This guide explains how to contribute and what to expect.

## Table of Contents

- [How We Accept Contributions](#how-we-accept-contributions)
- [Reporting a Bug](#reporting-a-bug)
- [Requesting a Feature or New Object](#requesting-a-feature-or-new-object)
- [Pull Requests](#pull-requests)
- [Building and Testing Locally](#building-and-testing-locally)
- [Security Issues](#security-issues)
- [Code of Conduct](#code-of-conduct)
- [License](#license)

## How We Accept Contributions

This provider manages two backends, **NIOS** and **Universal DDI (UDDI)**, from a single set of resources and data sources. Because many objects share a common schema, every change needs to be designed and validated with both backends in mind.

Before we merge a change to provider code, we check it against both backends. We run the full acceptance test suite against live NIOS Grids and Infoblox Portal accounts, across the supported versions. Most of the provider code is also produced and maintained by internal tooling, so a direct edit to these files would be overwritten by the next release.

For these reasons, **the best way to contribute a code change is to open an issue.** A clear issue gives us what we need to make the change properly, test it on both backends, and ship it in a release.

| You want to… | Do this |
|---|---|
| Report a bug | [Open a bug report](#reporting-a-bug) |
| Ask for a new resource, data source, or attribute | [Open a feature request](#requesting-a-feature-or-new-object) |
| Fix a typo or improve docs, guides, or examples | [Open a pull request](#pull-requests) |
| Change provider behavior in Go code | Open an issue first. You may include a proposed patch in it. |

## Reporting a Bug

Search the [existing issues](https://github.com/infobloxopen/terraform-provider-infoblox/issues) first. If you find a match, add your details to it instead of opening a new one.

A good bug report includes:

- **Versions**: the output of `terraform version` and the provider version.
- **Backend**: either NIOS (with its NIOS and WAPI versions) or UDDI.
- **Configuration**: the smallest Terraform configuration that reproduces the problem.
- **Steps**: the commands you ran, for example `terraform plan` and then `terraform apply`.
- **Expected and actual behavior**: what you expected and what happened, including the full error message.
- **Debug log**: run with `TF_LOG=DEBUG` and attach the relevant part. See [Debugging Terraform](https://developer.hashicorp.com/terraform/internals/debugging).


> **Remove secrets before you report a bug.** Remove passwords, API keys, hostnames, and IP addresses from configurations and logs.

## Requesting a Feature or New Object

Open an issue that describes:

- The object or attribute you need, and which backend or backends it applies to.
- The NIOS WAPI object or UDDI API endpoint it maps to, if you know it.
- Your use case, and an example of the Terraform configuration you would like to write.

Check the [Resources and Data Sources](docs/guides/resources-datasources.md) guide first. The object may already exist under a different name. If you are moving from another Infoblox provider, the [Object Mapping](docs/guides/object-mapping.md) guide shows the unified name for each object.

## Pull Requests

We welcome pull requests for:

- **Documentation**: guides under `templates/guides/`, or fixes to descriptions and typos.
- **Examples**: the Terraform files under `examples/`.

For any other change, open an issue first so we can agree on the approach.

When you open a pull request:

1. Fork the repository and create a branch from `master`.
2. Keep the change focused on one topic.
3. If you change anything under `examples/` or `templates/`, run `go generate ./...` and commit the updated files under `docs/`. Do not edit `docs/` directly, because `go generate` overwrites it.
4. Link the related issue in the pull request description, if there is one.

## Building and Testing Locally

You need:

- [Go](https://go.dev/doc/install) 1.26.8 or later
- [Terraform](https://developer.hashicorp.com/terraform/install) 1.12.1 or later

Build and install the provider:

```sh
git clone https://github.com/infobloxopen/terraform-provider-infoblox
cd terraform-provider-infoblox
make build
```

To use your local build, set up [development overrides](https://developer.hashicorp.com/terraform/cli/config/config-file#development-overrides-for-provider-developers) in your `~/.terraformrc`:

```hcl
provider_installation {
  dev_overrides {
    "infobloxopen/infoblox" = "<path to your PROVIDER_BIN>"
  }
  direct {}
}
```

Acceptance tests create real objects. Run them only against a test Grid or a test Portal account:

```sh
# NIOS
export NIOS_HOST_URL="https://<grid-ip>"
export NIOS_USERNAME="<username>"
export NIOS_PASSWORD="<password>"

# UDDI
export INFOBLOX_PORTAL_URL="https://csp.infoblox.com"
export INFOBLOX_PORTAL_KEY="<api-key>"

TF_ACC=1 go test ./internal/service/<group>/... -run <TestName> -v
```

Some tests need extra objects or environment variables. Check the test file before you run it.

## Security Issues

Do not report security vulnerabilities in public GitHub issues. Contact [Infoblox Support](https://info.infoblox.com/contact-form/) instead.

## Code of Conduct

Be respectful and constructive. We want this project to be welcoming to everyone.

## License

By contributing, you agree that your contributions are licensed under the [Mozilla Public License 2.0](LICENSE).
