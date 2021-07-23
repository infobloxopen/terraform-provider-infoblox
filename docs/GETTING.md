
# Getting the provider plugin
There are two ways of getting the plugin and using it:

* Use the published provider plugin from the Terraform registry: You do not need to do anything special, the plugin will be installed by Terraform automatically once it is mentioned as a requirement in your .tf files.
* Build from GitHub repository source code: This way is suitable for those who want to introduce some customizations or get the latest features under development, which are not in the official plugin yet.

## Getting it from the Terraform Official Registry
To use the published plugin available in the Terraform Registry:
  * Specify the version of the plugin
  * Configure it with the proper credentials.
Terraform installs the specified version of the plugin when a `terraform init` is run.

Below is a sample block. Refer to the official documentation for more information https://registry.terraform.io/providers/infobloxopen/infoblox/latest/docs .

### Using it in Terraform v0.14
Specify the plugin version in the .tf file as follows:
```
terraform {
  required_providers {
    infoblox = {
      source = “infobloxopen/infoblox”
      version = ">= 2.0.0"
    }
  }
}
provider "infoblox" {
  # Configuration options
}
```

Configure the credentials as environment variables as follows:
```
$ export INFOBLOX_SERVER=nios_ip-addr_or_hostname
$ export INFOBLOX_USERNAME=nios_username_on_the_server
$ export INFOBLOX_PASSWORD=appropriate_nios_password
```

## Building a Binary from the GitHub Source Code and Using it
To build the binary from the source code available in the GitHub repository, you must set up the environment, clone the code, and build it. Then, place the generated binary at an appropriate location and write suitable terraform configuration files to run. To build, complete the following steps::
* Install and set up Golang  version 1.16 or later from:
  `https://golang.org/doc/install`
* Install Terraform CLI v0.14.x from:  
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
* To install the resulting binary as a plugin, follow the instructions on page:
```
  https://www.terraform.io/docs/cli/config/config-file.html#development-overrides-for-provider-developers
```
