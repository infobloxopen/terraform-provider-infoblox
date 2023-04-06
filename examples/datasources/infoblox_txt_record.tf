data "infoblox_txt_record" "ds1"{
  // the arguments are taken from the examples for infoblox_txt_record resource

  // as we use a reference to a resource's field, we do not know if
  // it is 'default' (may be omitted) or not.
  dns_view=infoblox_txt_record.rec3.dns_view

  fqdn = infoblox_txt_record.rec3.fqdn

  // zone, text, ttl, comment and ext_attrs values may be retrieved using this data source.

  depends_on = [infoblox_txt_record.rec1]
}
