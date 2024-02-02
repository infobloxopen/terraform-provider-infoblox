package infoblox

const testCasePtrRecordTestErrData01 = `
resource "infoblox_ptr_record" "err_ptr_rec1" {
  ptrdname = "test.com"
  ip_addr = "10.0.0.1"
  record_name = "1.0.0.10.in-addr.arpa"
  cidr = "10.0.0.0/8"
}
`

const testCasePtrRecordTestErrData02 = `
resource "infoblox_ptr_record" "err_ptr_rec1" {
  ptrdname = "test.com"
  record_name = "1.0.0.10.in-addr.arpa"
  cidr = "10.0.0.0/8"
}
`

const testCasePtrRecordTestErrData03 = `
resource "infoblox_ptr_record" "err_ptr_rec1" {
  ptrdname = "test.com"
  ip_addr = "10.0.0.1"
  cidr = "10.0.0.0/8"
}
`

const testCasePtrRecordTestErrData04 = `
resource "infoblox_ptr_record" "err_ptr_rec1" {
  ptrdname = "test.com"
  ip_addr = "10.0.0.1"
  record_name = "1.0.0.10.in-addr.arpa"
}
`

const testCasePtrRecordTestErrData05Pre = `
resource "infoblox_ptr_record" "err_ptr_rec1" {
  ptrdname = "test.com"
  ip_addr = "10.0.0.1"
}
`

const testCasePtrRecordTestErrData05 = `
resource "infoblox_ptr_record" "err_ptr_rec1" {
  ptrdname = "test.com"
  ip_addr = "10.0.0.2"
  record_name = "2.0.0.10.in-addr.arpa"
}
`

const testCasePtrRecordTestErrData06 = `
resource "infoblox_ptr_record" "err_ptr_rec1" {
  ptrdname = "test.com"
  ip_addr = "10.10.0.1"
  cidr = "10.10.0.0/24"
}

resource "infoblox_ptr_record" "err_ptr_rec2" {
  ptrdname = "test.com"
  cidr = "172.18.18.0/24"
}
`

const testCasePtrRecordTestErrData07 = `
resource "infoblox_ptr_record" "err_ptr_rec2" {
  ptrdname = "test.com"
  cidr = "10.0.0.0/24"
  ip_addr = "10.0.0.1"
}
`

const testCasePtrRecordTestErrData08 = `
resource "infoblox_ptr_record" "err_ptr_rec2" {
  ptrdname = "test.com"
  cidr = "10.10.0.0/24"
  record_name = "2.0.10.10.in-addr.arpa"
}

resource "infoblox_ptr_record" "err_ptr_rec3" {
  ptrdname = "test.com"
  record_name = "2.0.0.10.in-addr.arpa"
}
`

const testCasePtrRecordTestErrData09 = `
resource "infoblox_ptr_record" "err_ptr_rec3" {
  ptrdname = "test.com"
  record_name = "1.0.0.10.in-addr.arpa"
  ip_addr = "10.0.0.1"
}
`

const testCasePtrRecordTestErrData10 = `
resource "infoblox_ptr_record" "err_ptr_rec3" {
  ptrdname = "test.com"
  record_name = "1.0.10.10.in-addr.arpa"
  cidr = "10.10.0.0/24"
}
`
