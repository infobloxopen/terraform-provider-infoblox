# A-record Resource

The `infoblox_a_record` resource associates a domain name with an IPv4 address.

The following list describes the parameters you can define in the resource block of the record:

* `fqdn`: required, specifies the fully qualified domain name for which you want to assign the IP address to. Example: `host43.zone12.org`
* `network_view`: optional, specifies the Network view to use when **allocating** an IP address from a network dynamically. For static allocation, do not use this field. Example: `networkview1`
* `dns_view`: optional, specifies the DNS view in which the zone exists. If a value is not specified, the default DNS view is considered. Example: `dns_view_1`
* `ttl`: optional, specifies the time to live value for the record. There is no default value for this parameter. If you do not specify a value, the TTL value is inherited from Grid DNS properties. A TTL value of 0 (zero) means caching should be disabled for this record. Example: `600`
* `comment`: optional, describes the record. Example: `static record #1`
* `ext_attrs`: optional, a set of NIOS extensible attributes that are attached to the record. Example: `jsonencode({})`
* `ip_addr`: required only for static allocation, specifies the IPv4 address to associate with the A-record. Example: `91.84.20.6`.
    * For allocating a static IP address, specify a valid IP address.
    * For allocating a dynamic IP address, do not use this field. Instead, define the `cidr` field. Optionally, set the `network_view` if you do not want to allocate in the default network view.
* `cidr`: required only for dynamic allocation, specifies the network from which to allocate an IP address when the `ip_addr` field is empty. The address is in CIDR format. For static allocation, use `ip_addr` instead of this field. Example: `192.168.10.4/30`.

-> All parameters except `fqdn` are optional. If you do not specify the `ip_addr` parameter, you must use the `cidr` parameter for dynamic address allocation. `cidr` and `ip_addr` are mutually exclusive. Specify a custom network view with `cidr` if you do not want to use the default network view.

### Examples of the A-record Block

```hcl
resource "infoblox_a_record" "a_record_static1"{
  # the zone 'example.com' MUST exist in the DNS view ('default')
  fqdn = "static1.example.com"

  ip_addr = "192.168.31.31"
  ttl = 10
  comment = "static A-record #1"
  ext_attrs = jsonencode({
    "Tenant ID" = "tf-plugin"
    "Cloud API Owned" = "True"
    "CMP Type"= "VMware"
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
            # using parent zone's TTL value.

  # In case of non-default network view,
  # you should specify a DNS view as well.
  network_view = "non_default" # corresponding network view MUST exist
  dns_view = "nondefault_view" # corresponding DNS view MUST exist

  # appropriate network MUST exist in the network view
  cidr = infoblox_ipv4_network.ipv4_network.cidr
}

resource "infoblox_a_record" "a_record_dynamic2"{
  fqdn = "dynamic2.example.com"
  # not specifying explicitly means using the default network view
  # network_view = "default"
  # not specifying explicitly means using the default DNS view
  # dns_view = "default"
  # appropriate network MUST exist in the network view
  cidr = infoblox_ipv4_network.ipv4_network.cidr
}
```
