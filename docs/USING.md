# Using the Provider from Terraform Registry
The plugin has been published under Terraform registry. Users can install it through terraform, for lifecycle management of Infoblox NIOS DDI objects.

To use the published plugin,
* Specify the version of the plugin and
* Configure it with the proper credentials
Terraform would take care to install the specified version of plugin when a `terraform init` is run.


## Using it in Terraform v0.11
The version of plugin and required credentials can be provided as follows in the tf file.
```
provider "infoblox"{
  version="~> 1.0"
  username="infoblox_user"
  password="infoblox"
  server="nios_server"
}
```
* `version`     : plugin version published in the registry.
* `username`    : NIOS grid user name
* `password`    : NIOS grid user password
* `server`      : NIOS rrid master or CP member IP address

