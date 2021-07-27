# PTR-record

PTR records can be created in forward and reverse mapping zones.

When a PTR record is created in forward mapping zones we need both
'ptrdname' and 'record_name'. Respective forward mapping zone has to be
in place at the appropriate DNS view.

When a PTR record is created in a reverse mapping zone it maps between
IP addresses and domain name; in other words, which domain name
corresponds to the given IP address. The attributes for the resource
are: network_view, cidr, ip_address, dns_view, ptrdname,
record_name, ttl. 'ptrdname' is required, others are optional from
Terraform's point of view. Consider the table of the attributes:

| Attribute | Required/optional | Description | Example |
| --- | --- | --- | --- |
| ptrdname | required | The domain name in FQDN to which the record should point to. | host1.example.com |
| ip_address | required for static allocation, see the description | IPv4/IPv6 address for record creation. Set the field with valid IP for static allocation. For dynamic allocation, leave this field empty and set the 'cidr' field. This field should be left empty in case PTR-record is to be created in a forward-mapping zone.| 82.50.36.8 |
cidr | required for dynamic allocation, see the description | The network address, in CIDR format, under which the record has to be created. | 10.3.128.0/20 |
network_view | see the description | Network view to use when allocating an IP address from a network dynamically. For static allocation, leave this field empty. | tenant 16 |
dns_view | optional | DNS view which the zone does exist within. | district-4 |
record_name | see the description | The domain name in FQDN, actual name of the record. For creating a PTR-record in a forward-mapping zone. | service1.zone21.org |

## Example

    resource "infoblox_ptr_record" "ptr_rr_1" {
      ptrdname = "vm1.test.com"
      dns_view = "default" # the same as omitting it
    
      # Record in forward mapping zone
      record_name = "vm1.test.com"
    
      # Record in reverse mapping zone
    
      # network_view = "default" # we can comment it out,
                                 # because this is the default value.
    
      # ip_addr is not defined, thus cidr is required.
      cidr = "10.2.3.0/24" # an IP address will be allocated.
    
      comment = "PTR record 1"
      ext_attrs = jsonencode({}) # Just an example of an empty EA set.
    }
