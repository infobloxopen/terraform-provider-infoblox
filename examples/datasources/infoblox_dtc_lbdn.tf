// creating a LBDN resource
resource "infoblox_dtc_lbdn" "lbdn_record" {
  name = "testLbdn3"
  lb_method = "ROUND_ROBIN"
  comment = "test LBDN"
}

// accessing LBDN by specifying name and comment
data "infoblox_dtc_lbdn" "readlbdn" {
  filters = {
    name = "testLbdn3"
    comment = "test LBDN"
  }
  // This is just to ensure that the record has been be created
  // using 'infoblox_dtc_lbdn' resource block before the data source will be queried.
  depends_on = [infoblox_dtc_lbdn.lbdn_record]
}

// returns matching LBDN with name and comment, if any
output "lbdn_res" {
  value = data.infoblox_dtc_lbdn.readlbdn
}