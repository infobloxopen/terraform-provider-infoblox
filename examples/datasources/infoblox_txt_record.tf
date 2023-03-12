data "infoblox_txt_record" "ds1"{
  // the arguments are taken from the examples for infoblox_txt_record resource
  dns_view=infoblox_txt_record.rec3.dns_view // this is a required parameter
  fqdn = infoblox_txt_record.rec3.fqdn

  // zone, text, ttl, comment and ext_attrs values may be retrieved using this data source.
}
