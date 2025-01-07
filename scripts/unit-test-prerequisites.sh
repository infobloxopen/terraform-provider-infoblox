#!/bin/bash

NIOS_SERVER="${NIOS_SERVER:-192.168.1.2:443}"
NIOS_USER="${NIOS_USER:-admin}"
NIOS_PASSWORD="${NIOS_PASSWORD:-infoblox}"

WAPI_URL="https://${NIOS_SERVER}/wapi/v2.13.5"
CURL_AUTH="-u ${NIOS_USER}:${NIOS_PASSWORD}"


curl -k -X POST -H 'Content-Type: application/json' $CURL_AUTH "${WAPI_URL}/extensibleattributedef" -d '{"name":"Location","type":"STRING"}'
echo

curl -k -X POST -H 'Content-Type: application/json' $CURL_AUTH "${WAPI_URL}/extensibleattributedef" -d '{"name":"Tenant ID","type":"STRING"}'
echo

curl -k -X POST -H 'Content-Type: application/json' $CURL_AUTH "${WAPI_URL}/extensibleattributedef" -d '{"name":"VM Name","type":"STRING"}'
echo

curl -k -X POST -H 'Content-Type: application/json' $CURL_AUTH "${WAPI_URL}/extensibleattributedef" -d '{"name":"Network Name","type":"STRING"}'
echo

# create a pool, zone_auth with grid primary, topology for DTC LBDN
pool=$(curl -k -X POST -H 'Content-Type: application/json' -u $CURL_AUTH "${WAPI_URL}/dtc:pool" -d '{"name":"pool-test","lb_preferred_method":"GLOBAL_AVAILABILITY"}')

members=$(curl -k -X GET -H 'Content-Type: application/json' -u $CURL_AUTH "${WAPI_URL}/member")
echo
host_name=$(echo $members | grep -o '"host_name": *"[^"]*' | head -1 | awk -F'"' '{print $4}')
echo $host_name

curl -k -X POST -H 'Content-Type: application/json' -u $CURL_AUTH "${WAPI_URL}/zone_auth" -d '{"name":"testZone.com","grid_primary":{"name":"$hostname"}}'
echo

curl -k -X POST -H 'Content-Type: application/json' -u $CURL_AUTH "${WAPI_URL}/dtc:topology" -d '{"name":"test-topo","rules":[{"dest_type":"POOL","destination_link":"$pool"}]}'
echo

