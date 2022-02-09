# PTR-record Resource

The `infoblox_ptr_record` resource allows you to create PTR-records in forward- and reverse-mapping zones. In case of reverse-mapping zone, the PTR-record maps IP addresses with domain names.

The following list describes the parameters you can define for the `infoblox_ptr_record` resource block:

* `ptrdname`: required, specifies the domain name in the FQDN format to which the record should point to. Example: `host1.example.com`.
* `ip_addr`: required only for static allocation, specifies the IPv4 or IPv6 address for record creation. Example: `82.50.36.8`.
    * For allocating a static IP address, specify a valid IP address.
    * For allocating a dynamic IP address, do not use this field. Instead, define the `cidr` field.
* `cidr`: required only for dynamic allocation, specifies the network address in CIDR format, under which the record must be created. For static allocation, do not use this field. Instead, define the `ip_addr` field. Example: `10.3.128.0/20`.
* `network_view`: required only for dynamic allocation, specifies the network view to use when allocating an IP address from a network dynamically. For static allocation, do not use this field. Example: `netview1`.
* `dns_view`: optional, specifies the DNS view in which the zone exists. If a value is not specified, the default DNS view is considered. Example: `external_dnsview`.
* `ttl`: optional, time to live value for the PTR-record. The parameter does not have a default value. If you do not specify a value, the TTL value is inherited from Grid DNS properties. A TTL value of 0 (zero) means caching should be disabled for this record. Example: `10`.
* `record_name`: required only in case of forward-mapping zones, the domain name in FQDN; actual name of the record. Example: `service1.zone21.org`.
* `comment`: optional, describes the PTR-record. Example: `some unknown host`.
* `ext_attrs`: optional, a set of NIOS extensible attributes that are attached to the PTR-record. Example: `jsonencode({})`.

-> When creating the PTR-record in a forward-mapping zone, `ptrdname` and `record_name` parameters are required, and `network_view` is optional. The corresponding forward-mapping zone must have been already created at the appropriate DNS view.

-> When creating the PTR-record in a reverse-mapping zone, the combination of `ptrdname`, `ip_addr`, and `network_view` parameters, or `ptrdname`, `cidr`, and `network_view` parameters is required.

### Example of the PTR-record Resource

```hcl
resource "infoblox_ptr_record" "ptr_rr_1" {
  ptrdname = "vm1.test.com"
  dns_view = "default" # the same as omitting it

  # Record in forward mapping zone
  record_name = "vm1.test.com"

  # Record in reverse mapping zone

  # network_view = "default" # we can comment it out,
                             # because this is the default value.

  # ip_addr is not defined, thus cidr is required.
  cidr = "infoblox_ipv4_network.ipv4_network.cidr" # an IPv4 address will be allocated.
  # an IPv6 address will be allocated. 
  # cidr = "infoblox_ipv6_network.ipv6_network.cidr" 

  comment = "PTR record 1"
  ext_attrs = jsonencode({
    "Tenant ID" = "tf-plugin"
    "Cloud API Owned" = "True"
    "CMP Type"= "VMware"
    "Site" = "Nevada" 
  }) 
}
```
