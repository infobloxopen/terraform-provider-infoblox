# Google Cloud Marketplace Terraform Module

This module deploys a product from the Google Cloud Marketplace via the user interface using Terraform.

## Usage
The provided test configuration can be used by executing:

```
terraform plan 
```

## Inputs
| Name | Description | Type | Default | Required |
|------|-------------|------|---------|----------|
| project_id | The ID of the project in which to provision resources. | `string` | `null` | yes |
| goog_cm_deployment_name | The name of the deployment and VM instance. | `string` | `null` | yes |
| source_image | The image name for the disk for the VM instance. | `string` | `"projects/infoblox-public-436917/global/images/infoblox-nios-906-53318-byol"` | yes |
| zone | The zone for the solution to be deployed. | `string` | `"us-east1-b"` | no |
| machine_type | The machine type to create, e.g. e2-small | `string` | `"n1-highmem-8"` | no |
| boot_disk_type | The boot disk type for the VM instance. | `string` | `"pd-ssd"` | no |
| boot_disk_size | The boot disk size for the VM instance in GBs | `number` | `500` | no |
| networks | The network name to attach the VM instance. | `list(string)` | `["default"]` | no |
| sub_networks | The sub network name to attach the VM instance. | `list(string)` | `[]` | no |
| external_ips | The external IPs assigned to the VM for public access. | `list(string)` | `[]` | no |
| labels | A map of key/value label pairs to assign to the instance. | `map(string)` | `{}` | no |
| enable_cloud_api | Allow full access to all of Google Cloud Platform APIs on the VM | `bool` | `true` | no |

## Outputs

| Name | Description |
|------|-------------|
| instance_self_link | Self-link for the compute instance. |
| instance_zone | Zone for the compute instance. |
| instance_machine_type | Machine type for the compute instance. |
| instance_nat_ip | External IP of the compute instance. |
| instance_network | Self-link for the network of the compute instance. |

## Requirements
### Terraform

Be sure you have the correct Terraform version (1.2.0+), you can choose the binary here:

https://releases.hashicorp.com/terraform/

### Configure a Service Account
In order to execute this module you must have a Service Account with the following project roles:

- `roles/compute.admin`
- `roles/iam.serviceAccountUser`

If you are using a shared VPC:

- `roles/compute.networkAdmin` is required on the Shared VPC host project.

### Enable API
In order to operate with the Service Account you must activate the following APIs on the project where the Service Account was created:

- Compute Engine API - `compute.googleapis.com`
