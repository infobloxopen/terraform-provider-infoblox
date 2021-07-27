# CNAME-record

Similarly, there is a data source for CNAME-records. The attributes are
literally the same as for A-record data source. To get information about
a CNAME-record, you have to specify a selector which uniquely identifies
it: a combination of DNS view ('dns_view' field), canonical name
('canonical' field) and an alias ('alias' field) which the record points
to. All the fields are required.

## Example

    data "infoblox_cname_record" "foo"{
      dns_view="default"
      alias="foo.test.com"
      canonical="main.test.com"    
    }
