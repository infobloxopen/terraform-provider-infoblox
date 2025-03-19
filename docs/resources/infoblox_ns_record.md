# NS-record Resource

The `infoblox_ns_record` resource enables you to perform `create`, `update` and `delete` operations on NS Record in a NIOS appliance.
The resource represents the ‘record:ns’ WAPI object in NIOS.

The following list describes the parameters you can define in the resource block of the NS Record object:

* `name`: required, The name of the NS record in FQDN format. This value can be in unicode format. Values with leading or trailing white space are not valid for this field. Example: `test.com`
* `nameserver`: required, The domain name of an authoritative server for the redirected zone. Values with leading or trailing white space are not valid for this field. Example: `ns1.test.com`
* `addresses`: required, The list of zone name servers. 
```terraform
addresses {
  address = "2.3.3.3"
  auto_create_ptr = true
}
```
* `dns_view`: optional, The name of the DNS view in which the record resides.The default value is The default DNS view. Example: `external`

### Examples of a NS Record Block

```hcl
//creating NS record 
resource "infoblox_ns_record" "ns" {
  name = "test.com"
  nameserver = "name_server.test.com"
  addresses{
    address = "2.3.2.5"
    auto_create_ptr=true
  }
  dns_view = "default"
}
```