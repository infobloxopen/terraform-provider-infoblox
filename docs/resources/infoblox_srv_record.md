# SRV-record Resource

The `infoblox_srv_record` resource corresponds to SRV-record (service record) on NIOS side, and its purpose is to provide
information about a network endpoint (host and port) which provides particular network service for the specified DNS zone.

The following list describes the parameters you can define in the resource block of the record:

* `dns_view`: optional, specifies the DNS view which the zone exists in. If a value is not specified, the name `default` is used for DNS view. Example: `dns_view_1`
* `name`: required, specifies the record's name in the format, defined in RFC2782 document. Example: `_http._tcp.acme.com`
* `target`: required, specifies an FQDN of the host which is responsible for providing the service specified by `name`. Example: `www.acme.com`
* `port`: required, specifies a port number (0..65535) on the `target` host which the service expects requests on.
* `priority`: required, specifies a priority number, as described in RFC2782.
* `weight`: required, specifies a weight number, as described in RFC2782.
* `ttl`: optional, specifies the "time to live" value for the record. There is no default value for this parameter. If a value is not specified, then in NIOS, the value is inherited from the parent zone of the DNS record for this resource. A TTL value of 0 (zero) means caching should be disabled for this record. Example: `600`
* `comment`: optional, describes the record. Example: `auto-created test record #1`
* `ext_attrs`: optional, a set of NIOS extensible attributes that are attached to the record. Example: `jsonencode({})`

## Examples

```hcl
// minimal set of parameters
resource "infoblox_srv_record" "rec1" {
    name = "_http._tcp.example.org"
    priority = 100
    weight = 75
    port = 8080
    target = "www.example.org"
} 

// all set of parameters for SRV record
resource "infoblox_srv_record" "rec2" {
    dns_view = "nondefault_dnsview1" // not 'default' thus must be specified
    name = "_sip._udp.example2.org"
    priority = 12
    weight = 10
    port = 5060
    target = "sip.example2.org"
    ttl = 3600
    comment = "example SRV record"
    ext_attrs = jsonencode({
        "Location" = "65.8665701230204, -37.00791763398113"
    })
}
```
