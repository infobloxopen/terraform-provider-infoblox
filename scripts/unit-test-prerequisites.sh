#!/bin/bash

NIOS_SERVER="${NIOS_SERVER:-192.168.1.2:443}"
NIOS_USER="${NIOS_USER:-admin}"
NIOS_PASSWORD="${NIOS_PASSWORD:-infoblox}"

WAPI_URL="https://${NIOS_SERVER}/wapi/v2.11.1"
CURL_AUTH="-u ${NIOS_USER}:${NIOS_PASSWORD}"


curl -k -X POST -H 'Content-Type: application/json' $CURL_AUTH "${WAPI_URL}/extensibleattributedef" -d '{"name":"Location","type":"STRING"}'
echo

curl -k -X POST -H 'Content-Type: application/json' $CURL_AUTH "${WAPI_URL}/extensibleattributedef" -d '{"name":"Tenant ID","type":"STRING"}'
echo

curl -k -X POST -H 'Content-Type: application/json' $CURL_AUTH "${WAPI_URL}/extensibleattributedef" -d '{"name":"VM Name","type":"STRING"}'
echo

curl -k -X POST -H 'Content-Type: application/json' $CURL_AUTH "${WAPI_URL}/extensibleattributedef" -d '{"name":"Network Name","type":"STRING"}'
echo




