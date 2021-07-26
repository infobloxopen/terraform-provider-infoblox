# Developing the provider
Contribution to the plugin can be done at various levels and we would embrace it.
Working at code level would require to setup the environment, clone the code, build and test the code.


## Requirements
Golang and Terraform in the system are basic requirements to build and test the plugin.
* Install and configure apt version of [Golang](https://golang.org/doc/install). At this point 1.12.x version has been used for development.
* [Terraform](https://www.terraform.io/downloads.html) 0.11.x used for development and testing.

## Building the Provider
Once environment has been setup we are ready to clone the code and build it. Run the build before and after the changes and make sure it passes without failures.

The project uses [Go Modules](https://blog.golang.org/using-go-modules) making it safe to work with it outside of your existing [GOPATH](http://golang.org/doc/code.html#GOPATH).
```sh
$ mkdir -p $HOME/development/infoblox/
$ cd $HOME/development/infoblox/
$ git clone https://github.com/infobloxopen/terraform-provider-infoblox
$ cd terraform-provider-infoblox
$ make build
```


## Testing the provider
To test the provider and to run the full suite of acceptance tests run below commands respectively,
```sh
$ make test
$ make testacc
```
While running acceptance tests, 
* Create appropriate DNS views and zones being defined in test files.
* Create appropriate EAs

Writing acceptance tests for the bugs fixed/features developed would be recommended.

## Using the Provider
The bug fixes and features developed can be tested using the binary built after source code changes.

When `terraform init` is run on terraform v0.11 it creates a `.terraform` directory, in the parent directory. The newly created directory consists of plugin binaries specified in tf files. Replace the infoblox binary in this directory with the one you have built with the changes.

The location of binary would be as below, when infoblox plugin version given is 1.1.0 ,
`.terraform/plugins/registry.terraform.io/terraform-providers/infoblox/1.1.0/linux_amd64/`

