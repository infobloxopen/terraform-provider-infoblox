
# Getting the provider plugin
There are two ways of getting the plugin and using it:

* Use the published provider plugin from the Terraform registry: You do not need to do anything special, the plugin will be installed by Terraform automatically once it is mentioned as a requirement in your .tf files.
* Build from GitHub repository source code: This way is suitable for those who want to introduce some customizations or get the latest features under development, which are not in the official plugin yet.

## Getting it from the Terraform Official Registry
Specify the plugin version in the .tf file as follows:
  ```
    terraform {
      required_providers {
        infoblox = {
          source = "infobloxopen/infoblox"
          version = ">= 2.7.0"
        }
      }
    }

    provider "infoblox" {
      # Configuration options
      server = "nios_ip-addr"
      username = "username"
      password = "password"
    }
  ```

Configure the credentials as environment variables as follows:
  ```
    $ export INFOBLOX_SERVER=<nios_ip-addr> or <hostname>
    $ export INFOBLOX_USERNAME=<nios_username>
    $ export INFOBLOX_PASSWORD=<nios_password>
  ```

Terraform installs the specified version of the plugin when a `terraform init` is run.

Refer to the official documentation for more information https://registry.terraform.io/providers/infobloxopen/infoblox/latest/docs .

## Building a Binary from the GitHub Source Code and Using it
Complete the following steps to build the binary:
* Install and set up Golang  version 1.21 or later from:
  `https://golang.org/doc/install`
* Install Terraform CLI v1.8.1+ from:  
  `https://www.terraform.io/downloads.html`
* Clone the repo and build it as follows:
  ```
    $ cd `go env GOPATH`/src
    $ mkdir -p github.com/infobloxopen
    $ cd github.com/infobloxopen
    $ git clone https://github.com/infobloxopen/terraform-provider-infoblox
    $ cd terraform-provider-infoblox
    $ make build
  ```  

To install the resulting binary as a plugin, follow the instructions on page:
  ```
    https://www.terraform.io/docs/cli/config/config-file.html#development-overrides-for-provider-developers
  ```
