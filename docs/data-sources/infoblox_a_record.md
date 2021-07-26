# A-record

This data source allows to retrieve the following information
(attributes) for an A-record which is managed by NIOS server:

| Attribute | Description | Example |
| --- | --- | --- |
| zone | The zone which contains the record in the given DNS view. | test.com |
| ttl | TTL of the record, in seconds. | 3600 |
| comment | A text describing the record, a regular comment. | Temporary A-record |
| ext_attrs | A set of extensible attributes of the record, if any. The content is in JSON map format. |  {"Owner": "State Library", "Expires": "never"} |

To get information about an A-record, you have to specify a selector
which uniquely identifies it: a combination of DNS view ('dns_view'
field), IPv4 address, which the record points to ('ip_addr' field), and
FQDN which corresponds to the IP address. All the fields are required to
get information about an A-record.

## Example

    data "infoblox_a_record" "vip_host" {
      dns_view="default"
      fqdn="very-interesting-host.example.com"
      ip_addr="10.3.1.65"
    }
    
    output "id" {
      value = data.infoblox_a_record.vip_host
    }
    
    output "zone" {
      value = data.infoblox_a_record.vip_host.zone
    }
    
    output "ttl" {
      value = data.infoblox_a_record.vip_host.ttl
    }
    
    output "comment" {
      value = data.infoblox_a_record.vip_host.comment
    }
    
    output "ext_attrs" {
      value = data.infoblox_a_record.vip_host.ext_attrs
    }

This defines a data source of type 'infoblox_a_record' with the name
'vip_host' in a Terraform file. Further you can reference this resource
and retrieve some information about it. For example,
**data.infoblox_a_record.vip_host.comment** gives a text which is a
comment for the A-record.
