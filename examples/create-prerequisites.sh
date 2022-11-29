#!/bin/bash

NIOS_SERVER="${NIOS_SERVER:-192.168.0.1:443}"
NIOS_USER="${NIOS_USER:-admin}"
NIOS_PASSWORD="${NIOS_PASSWORD:-infoblox}"

WAPI_URL="https://${NIOS_SERVER}/wapi/v2.11.1"
CURL_AUTH="-u ${NIOS_USER}:${NIOS_PASSWORD}"

curl -k -X POST -H 'Content-Type: application/json' $CURL_AUTH "${WAPI_URL}/networkview" -d '{"name":"nondefault_netview"}'
curl -k -X POST -H 'Content-Type: application/json' $CURL_AUTH "${WAPI_URL}/view" -d '{"network_view":"default","name":"nondefault_dnsview1"}'
curl -k -X POST -H 'Content-Type: application/json' $CURL_AUTH "${WAPI_URL}/view" -d '{"network_view":"nondefault_netview","name":"nondefault_dnsview2"}'

curl -k -X POST -H 'Content-Type: application/json' $CURL_AUTH "${WAPI_URL}/zone_auth" -d '{"view":"default","fqdn":"example1.org"}'
curl -k -X POST -H 'Content-Type: application/json' $CURL_AUTH "${WAPI_URL}/zone_auth" -d '{"view":"nondefault_dnsview1","fqdn":"example2.org"}'
curl -k -X POST -H 'Content-Type: application/json' $CURL_AUTH "${WAPI_URL}/zone_auth" -d '{"view":"default.nondefault_netview","fqdn":"example3.org"}'
curl -k -X POST -H 'Content-Type: application/json' $CURL_AUTH "${WAPI_URL}/zone_auth" -d '{"view":"nondefault_dnsview2","fqdn":"example4.org"}'

for view in default default.nondefault_netview nondefault_dnsview1 nondefault_dnsview2; do
  for zone in 10.0.0.0/8 2002:1f93::/64; do
    curl -k -X POST -H 'Content-Type: application/json' $CURL_AUTH "${WAPI_URL}/zone_auth" -d "{\"view\":\"${view}\",\"fqdn\":\"${zone}\"}"
  done
done

curl -k -X POST -H 'Content-Type: application/json' $CURL_AUTH "${WAPI_URL}/extensibleattributedef" -d '{"name":"Terraform Internal ID","type":"STRING"}'

# Uncomment the lines below if you haven't got a Cloud Network Automation license installed, and there are no appropriate EAs defined at your appliance.
#curl -k -X POST -H 'Content-Type: application/json' $CURL_AUTH "${WAPI_URL}/extensibleattributedef" -d '{"name":"CMP Type","type":"STRING"}'
#curl -k -X POST -H 'Content-Type: application/json' $CURL_AUTH "${WAPI_URL}/extensibleattributedef" -d '{"name":"Tenant ID","type":"STRING"}'
#curl -k -X POST -H 'Content-Type: application/json' $CURL_AUTH "${WAPI_URL}/extensibleattributedef" -d '{"name":"Cloud API Owned","type":"ENUM","list_values":[{"value": "True"},{"value": "False"}]}'
