# PTR-record Resource

The `infoblox_ptr_record` resource allows you to create PTR-records in forward-mapping and reverse-mapping zones. In case of reverse-mapping zone, the PTR-record maps IP addresses with domain names.

The following list describes the parameters you can define for the `infoblox_ptr_record` resource block:

* `ptrdname`: required, specifies the domain name in the FQDN format to which the record should point to. Example: `host1.example.com`.
* `ip_addr`: required only for static allocation in reverse-mapping zones, specifies the IPv4 or IPv6 address for record creation in reverse-mapping zone. Example: `82.50.36.8`.
    * For allocating a static IP address, specify a valid IP address.
    * For allocating a dynamic IP address, do not use this field. Instead, define the `cidr` field.
* `cidr`: required only for dynamic allocation in reverse-mapping zones, specifies the network address in CIDR format, under which the record must be created. For static allocation, do not use this field. Instead, define the `ip_addr` field. Example: `10.3.128.0/20`.
* `network_view`: optional, specifies the network view to use when allocating an IP address from a network dynamically. If a value is not specified, the name `default` is used as the network view. For static allocation, do not use this field. Example: `netview1`.
* `dns_view`: optional, specifies the DNS view in which the zone exists. If a value is not specified, the name `default` is used as the DNS view. Example: `external_dnsview`.
* `ttl`: optional, specifies the "time to live" value for the PTR-record. The parameter does not have a default value. If a value is not specified, then in NIOS, the value is inherited from the parent zone of the DNS record for this resource. A TTL value of 0 (zero) means caching should be disabled for this record. Example: `10`.
* `record_name`: required only in case of forward-mapping zones, specifies the domain name in FQDN format; it is the name of the DNS PTR-record. Example: `service1.zone21.org`.
* `comment`: optional, describes the PTR-record. Example: `some unknown host`.
* `ext_attrs`: optional, a set of NIOS extensible attributes that are attached to the PTR-record. Example: `jsonencode({})`.

-> When creating the PTR-record in a forward-mapping zone, `ptrdname` and `record_name` parameters are required, and `network_view` is optional. The corresponding forward-mapping zone must have been already created at the appropriate DNS view.

->  When creating a PTR-record in a reverse-mapping zone, you must specify the `ptrdname` attribute with either the `ip_addr`, `cidr`, or `record_name` attribute. If you configure all three parameters, that is `ip_addr`, `cidr`, and `record_name`, the order of precedence for the record creation is `record_name` > `ip_addr` > `cidr`. The values of parameters with lower precedence are ignored.

### Example of a PTR-record Resource

```hcl
// PTR-record, minimal set of parameters
// Actually, either way may be used in reverse-mapping
// zones to specify an IP address:
//   1) 'ip_addr' (yes, literally) and
//   2) 'record_name' - in the form of a domain name (ex. 1.0.0.10.in-addr.arpa)
resource "infoblox_ptr_record" "ptr1" {
  ptrdname = "rec1.example1.org"
  ip_addr = "10.0.0.1"
}

resource "infoblox_ptr_record" "ptr2" {
  ptrdname = "rec2.example1.org"
  record_name = "2.0.0.10.in-addr.arpa"
}

// statically allocated PTR-record, full set of parameters
resource "infoblox_ptr_record" "ptr3" {
  ptrdname = "rec3.example2.org"
  dns_view = "nondefault_dnsview1"
  ip_addr = "2002:1f93::3"
  comment = "workstation #3"
  ttl = 300 # 5 minutes
  ext_attrs = jsonencode({
    "Location" = "the main office"
  })
}

// dynamically allocated PTR-record, minimal set of parameters
resource "infoblox_ptr_record" "ptr4" {
  ptrdname = "rec4.example2.org"
  cidr = infoblox_ipv4_network.net1.cidr
}

// statically allocated PTR-record, full set of parameters, non-default network view
resource "infoblox_ptr_record" "ptr5" {
  ptrdname = "rec5.example2.org"
  dns_view = "nondefault_dnsview2"
  network_view = "nondefault_netview"
  ip_addr = "2002:1f93::5"
  comment = "workstation #5"
  ttl = 300 # 5 minutes
  ext_attrs = jsonencode({
    "Location" = "the main office"
  })
}

// PTR-record in a forward-mapping zone
resource "infoblox_ptr_record" "ptr6_forward" {
  ptrdname = "example1.org"
  record_name = "www.example1.org"
}
```
