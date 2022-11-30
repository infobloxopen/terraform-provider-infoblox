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

for view in default default.nondefault_netview nondefault_dnsview1 nondefault_dnsview2; do
  for zone in 10.0.0.0/8 2002:1f93::/64; do
    curl -k -X POST -H 'Content-Type: application/json' $CURL_AUTH "${WAPI_URL}/zone_auth" -d "{\"view\":\"${view}\",\"fqdn\":\"${zone}\"}"
    echo
  done
done

curl -k -X POST -H 'Content-Type: application/json' $CURL_AUTH "${WAPI_URL}/extensibleattributedef" -d '{"name":"Terraform Internal ID","type":"STRING"}'
echo

# Uncomment the lines below if you haven't got a Cloud Network Automation license installed, and there are no appropriate EAs defined at your appliance.
#curl -k -X POST -H 'Content-Type: application/json' $CURL_AUTH "${WAPI_URL}/extensibleattributedef" -d '{"name":"CMP Type","type":"STRING"}'
#echo
#curl -k -X POST -H 'Content-Type: application/json' $CURL_AUTH "${WAPI_URL}/extensibleattributedef" -d '{"name":"Tenant ID","type":"STRING"}'
#echo
#curl -k -X POST -H 'Content-Type: application/json' $CURL_AUTH "${WAPI_URL}/extensibleattributedef" -d '{"name":"Cloud API Owned","type":"ENUM","list_values":[{"value": "True"},{"value": "False"}]}'
#echo
