# Developing the Provider Plugin
Contribution to the plugin can be done at various levels and we would embrace it. 

Working at code level requires you  to set up the environment, clone the code, build, and test the code.

## Building the Binary from Source Code and Using it
Golang and Terraform installed in the system are basic requirements to build and test the plugin.

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
  
* To install the resulting binary as a plugin, follow the instructions on page:
  `https://www.terraform.io/docs/cli/config/config-file.html#development-overrides-for-provider-developers`

## Testing the Provider
To run the full suite of acceptance tests, run the following commands:
  ```
    $ export INFOBLOX_SERVER=<nios_ip-addr> or <hostname>
    $ export INFOBLOX_USERNAME=<nios_username>
    $ export INFOBLOX_PASSWORD=<nios_password>
    $ make test
    $ export TF_ACC=true # without this only unit tests (not acceptance tests) run
    $ make testacc
  ```

Refer to the comments included in the code for running the tests, and make sure that the mentioned conditions are met. 
For example, you may have to create objects such as DNS zones and views before running the tests.
