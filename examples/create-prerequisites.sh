#!/bin/bash

NIOS_SERVER="${NIOS_SERVER:-192.168.1.2:443}"
NIOS_USER="${NIOS_USER:-admin}"
NIOS_PASSWORD="${NIOS_PASSWORD:-infoblox}"

WAPI_URL="https://${NIOS_SERVER}/wapi/v2.11.1"
CURL_AUTH="-u ${NIOS_USER}:${NIOS_PASSWORD}"

curl -k -X POST -H 'Content-Type: application/json' $CURL_AUTH "${WAPI_URL}/networkview" -d '{"name":"nondefault_netview"}'
echo
curl -k -X POST -H 'Content-Type: application/json' $CURL_AUTH "${WAPI_URL}/view" -d '{"network_view":"default","name":"nondefault_dnsview1"}'
echo
curl -k -X POST -H 'Content-Type: application/json' $CURL_AUTH "${WAPI_URL}/view" -d '{"network_view":"nondefault_netview","name":"nondefault_dnsview2"}'
echo

curl -k -X POST -H 'Content-Type: application/json' $CURL_AUTH "${WAPI_URL}/zone_auth" -d '{"view":"default","fqdn":"example1.org"}'
echo
curl -k -X POST -H 'Content-Type: application/json' $CURL_AUTH "${WAPI_URL}/zone_auth" -d '{"view":"nondefault_dnsview1","fqdn":"example2.org"}'
echo
curl -k -X POST -H 'Content-Type: application/json' $CURL_AUTH "${WAPI_URL}/zone_auth" -d '{"view":"default.nondefault_netview","fqdn":"example3.org"}'
echo
curl -k -X POST -H 'Content-Type: application/json' $CURL_AUTH "${WAPI_URL}/zone_auth" -d '{"view":"nondefault_dnsview2","fqdn":"example4.org"}'
echo

# Create a custom network view and a zone for Dynamic allocation using next available IP and network.
curl -k -X POST -H 'Content-Type: application/json' $CURL_AUTH "${WAPI_URL}/networkview" -d '{"name":"custom"}'
echo
curl -k -X POST -H 'Content-Type: application/json' $CURL_AUTH "${WAPI_URL}/networkview" -d '{"name":"test"}'
echo

curl -k -X POST -H 'Content-Type: application/json' $CURL_AUTH "${WAPI_URL}/zone_auth" -d '{"view":"custom","fqdn":"test.com"}'
echo
curl -k -X POST -H 'Content-Type: application/json' $CURL_AUTH "${WAPI_URL}/zone_auth" -d '{"view":"custom","fqdn":"ex.org"}'
echo
curl -k -X POST -H 'Content-Type: application/json' $CURL_AUTH "${WAPI_URL}/zone_auth" -d '{"view":"custom","fqdn":"test_fwzone"}'
echo
curl -k -X POST -H 'Content-Type: application/json' $CURL_AUTH "${WAPI_URL}/zone_auth" -d '{"view":"custom","fqdn":"example.com"}'
echo
curl -k -X POST -H 'Content-Type: application/json' $CURL_AUTH "${WAPI_URL}/zone_auth" -d '{"view":"test","fqdn":"test.com"}'
echo

# Create a network and network container with extensible attributes for Next available functionality.
curl -k -X POST -H 'Content-Type: application/json' $CURL_AUTH "${WAPI_URL}/network" -d '{"network": "10.1.0.0/24", "extattrs": {"Site": {"value": "Turkey"}}}'
echo
curl -k -X POST -H 'Content-Type: application/json' $CURL_AUTH "${WAPI_URL}/ipv6network" -d '{"network": "551:0db8:85a3::/64", "extattrs": {"Site": {"value": "Turkey"}}}'
echo
curl -k -X POST -H 'Content-Type: application/json' $CURL_AUTH "${WAPI_URL}/network" -d '{"network": "10.10.11.1", "extattrs": {"Site": {"value": "Nainital"}}}'
echo
curl -k -X POST -H 'Content-Type: application/json' $CURL_AUTH "${WAPI_URL}/networkcontainer" -d '{"network": "10.10.11.0/28", "extattrs": {"Site": {"value": "Blr"}}}'
echo
curl -k -X POST -H 'Content-Type: application/json' $CURL_AUTH "${WAPI_URL}/ipv6network" -d '{"network": "556:0db8:85a3::/64", "extattrs": {"Site": {"value": "Blr"}}}'
echo
curl -k -X POST -H 'Content-Type: application/json' $CURL_AUTH "${WAPI_URL}/ipv6networkcontainer" -d '{"network": "555:0db8:85a3::/64", "extattrs": {"Site": {"value": "Uzbekistan"}}}'
echo

for view in default default.nondefault_netview nondefault_dnsview1 nondefault_dnsview2; do
  for zone in 2002:1f93::/64; do
    curl -k -X POST -H 'Content-Type: application/json' $CURL_AUTH "${WAPI_URL}/zone_auth" -d "{\"view\":\"${view}\",\"fqdn\":\"${zone}\",\"zone_format\":\"IPV6\"}"
    echo
  done
done

for view in default default.nondefault_netview nondefault_dnsview1 nondefault_dnsview2; do
  for zone in 10.0.0.0/8; do
    curl -k -X POST -H 'Content-Type: application/json' $CURL_AUTH "${WAPI_URL}/zone_auth" -d "{\"view\":\"${view}\",\"fqdn\":\"${zone}\",\"zone_format\":\"IPV4\"}"
    echo
  done
done

curl -k -X POST -H 'Content-Type: application/json' $CURL_AUTH "${WAPI_URL}/extensibleattributedef" -d '{"name":"Terraform Internal ID","type":"STRING"}'
echo

# Uncomment the lines below if you haven't got a Cloud Network Automation license installed, and there are no appropriate EAs defined at your appliance.
# curl -k -X POST -H 'Content-Type: application/json' $CURL_AUTH "${WAPI_URL}/extensibleattributedef" -d '{"name":"CMP Type","type":"STRING"}'
# echo
# curl -k -X POST -H 'Content-Type: application/json' $CURL_AUTH "${WAPI_URL}/extensibleattributedef" -d '{"name":"Tenant ID","type":"STRING"}'
# echo
# curl -k -X POST -H 'Content-Type: application/json' $CURL_AUTH "${WAPI_URL}/extensibleattributedef" -d '{"name":"Cloud API Owned","type":"ENUM","list_values":[{"value": "True"},{"value": "False"}]}'
# echo
