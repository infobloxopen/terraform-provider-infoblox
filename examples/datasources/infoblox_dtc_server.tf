// creating a DTC Server resource
resource "infoblox_dtc_server" "server_record" {
  name = "server3"
  host = "34.23.3.1"
  comment = "test DTC Server"
}

// accessing s DTC Server by specifying name and comment
data "infoblox_dtc_server" "read_server" {
  filters = {
    name = "server3"
    comment = "test DTC Server"
  }
  // This is just to ensure that the record has been be created
  // using 'infoblox_dtc_server' resource block before the data source will be queried.
  depends_on = [infoblox_dtc_server.server_record]
}

// returns matching DTC Server with name and comment, if any
output "server_res" {
  value = data.infoblox_dtc_server.read_server
}