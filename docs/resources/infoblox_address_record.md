# A-record and AAAA-record

A-record resource associates a domain name with an IPv4 address,
AAAA-record does the same for IPv6 addresses. The attributes for this
resource are: network_view, cidr, dns_view, fqdn, ip_address for
IPv4, ipv6_addr for IPv6, ttl. All except 'fqdn' are optional from
Terraform's point of view. You may omit 'ip_address' ('ipv6_addr') but
in this case you must use 'cidr' for dynamic address allocation. 'cidr'
and 'ip_address' ('ipv6_addr') are mutually exclusive. 'network_view'
makes sense only along with 'cidr' if you want to specify a network view
other than 'default'. Here is the table of attributes with their meaning
and examples:

| Attribute       | Required/optional | Description     | Example         |
| --- | --- | --- | --- |
| fqdn            | required        | Fully-qualified domain name which you want to assign the IP-address to. | host43.zone12.org |
| ip_address     | required for static allocation, see the description | Only for A-record. IP address to associate with the A-record. For static allocation, set the field with a valid IP address. For dynamic allocation, leave this field empty and set 'cidr' and optionally 'network_view' fields. | 91.84.20.6 |
| ipv6_address     | required for dynamic allocation, see the description | Only for AAAA-record. IP address to associate with the AAAA-record. For static allocation, set the field with a valid IP address. For dynamic allocation, leave this field empty and set 'cidr' and optionally 'network_view' fields. | The following are equivalent forms of the same IPv6 address: 2001:0db8:0000:0000:0000:ff00:0042:8329, 2001:db8:0:0:0:ff00:42:8329, 2001:db8::ff00:42:8329 |
| cidr     | see the description | Network to allocate an IP address from, when the 'ip_addr' field is empty (dynamic allocation). The address is in CIDR format. For static allocation, leave this field empty. | For IPv4: 192.168.10.4/30, for IPv6: 2001:db8::/64 |
| network_view   | optional        | Network view to use when allocating an IP address from a network dynamically. For static allocation, leave this field empty. | dmz_netview |
| dns_view       | optional        | DNS view which the zone does exist within. If omitted, the value ‘default’ is used. | internal_network |

## Example A-record resource definitions

    resource "infoblox_a_record" "a_record_static1"{
    # the zone 'example.com' MUST exist in the DNS view ('default')
    fqdn = "static1.example.com"
    
    ip_addr = "192.168.31.31"
    ttl = 10
    comment = "static A-record #1"
    ext_attrs = jsonencode({
    "Location" = "New York"
    "Site" = "HQ"
    })
    }
    
    resource "infoblox_a_record" "a_record_static2"{
    fqdn = "static2.example.com"
    ip_addr = "192.168.31.32"
    ttl = 0 # ttl=0 means 'do not cache'
    dns_view = "non_default_dnsview" # corresponding DNS view MUST exist
    }
    
    resource "infoblox_a_record" "a_record_dynamic1"{
    fqdn = "dynamic1.example.com"
    # ip_addr = "" # CIDR is used for dynamic allocation
    # ttl = 0 # not mentioning TTL value means
                # using the parent zone's TTL value.
    
    # In case of non-default network view,
    # you should specify a DNS view as well.
    network_view = "non_default" # corresponding network view MUST exist
    dns_view = "nondefault_view" # corresponding DNS view MUST exist
    
    # appropriate network MUST exist in the network view
    cidr = "10.20.30.0/24"
    }
    
    resource "infoblox_a_record" "a_record_dynamic2"{
    fqdn = "dynamic2.example.com"
    # not specifying explicitly means using the default network view
    # network_view = "default"
    # not specifying explicitly means using the default DNS view
    # dns_view = "default"
    # appropriate network MUST exist in the network view
    cidr = "10.20.30.0/24"
    }


## Example AAAA-record resource definitions

    resource "infoblox_aaaa_record" "aaaa_record_static1"{
      # the zone 'example.com' MUST exist in the DNS view ('default')
      fqdn = "static1-v6.example.com"
    
      ipv6_addr = "2001:db8::ff00:42:8329"
      ttl = 10
      comment = "static AAAA-record #1"
      ext_attrs = jsonencode({
        "Location" = "New York"
        "Site" = "HQ"
      })
    }
    
    resource "infoblox_aaaa_record" "aaaa_record_dynamic1"{
      fqdn = "dynamic2-v6.example.com"
    
      # not specifying explicitly means using the default network view
      # network_view = "default"
    
      # not specifying explicitly means using the default DNS view
      # dns_view = "default"
    
      # appropriate network MUST exist in the network view
      cidr = "2001:db8::/96"
    }
