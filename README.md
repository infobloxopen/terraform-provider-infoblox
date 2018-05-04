# terraform-provider-infoblox
Infoblox Provider for Terraform

## Requirments

* [Terraform](https://www.terraform.io/downloads.html) 
* [Go](https://golang.org/doc/install) 1.9(to build the provider plugin)
* [dep](https://github.com/golang/dep)

## Building the Provider

clone repository to ```$GOPATH/src/github.com/infobloxopen/terraform-provider-infoblox```
```
$ mkdir -p $GOPATH/src/github.com/infobloxopen; cd $GOPATH/src/github.com/infobloxopen
$ git clone git@github.com:infobloxopen/terraform-provider-infoblox
```
Enter the provider directory and build the provider
```
$ cd $GOPATH/src/github.com/infobloxopen/terraform-provider-infoblox
$ make build
```
## Using the Provider
If you're building the provider, follow the instructions to [install it as a plugin](https://www.terraform.io/docs/plugins/basics.html#installing-a-plugin). After placing it into your plugins directory, run ```terraform init``` to initialize it.

## Developing the Provider
If you wish to work on the provider, you'll first need Go installed on your machine (version 1.9+ is required). You'll also need to correctly setup a GOPATH, as well as adding ```$GOPATH/bin``` to your ```$PATH```.

To compile the provider, run ```make build```. This will build the provider and put the provider binary in the ```$GOPATH/bin``` directory.
```
$ make build
...
$ $GOPATH/bin/terraform-provider-infoblox
...
```
In order to test the provider, you can simply run make test.
```
$ make test
```
In order to run the full suite of Acceptance tests, run make testacc.

Note: Acceptance tests create real resources, and often cost money to run.
```
$ make testacc
```
