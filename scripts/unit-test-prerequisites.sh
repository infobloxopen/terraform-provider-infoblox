#!/bin/bash

NIOS_SERVER="${NIOS_SERVER:-10.197.147.146:443}"
NIOS_USER="${NIOS_USER:-admin}"
NIOS_PASSWORD="${NIOS_PASSWORD:-infoblox}"

WAPI_URL="https://${NIOS_SERVER}/wapi/v2.12.3"
CURL_AUTH="-u ${NIOS_USER}:${NIOS_PASSWORD}"


curl -k -X POST -H 'Content-Type: application/json' $CURL_AUTH "${WAPI_URL}/extensibleattributedef" -d '{"name":"Location","type":"STRING"}'
echo

curl -k -X POST -H 'Content-Type: application/json' $CURL_AUTH "${WAPI_URL}/extensibleattributedef" -d '{"name":"Tenant ID","type":"STRING"}'
echo

curl -k -X POST -H 'Content-Type: application/json' $CURL_AUTH "${WAPI_URL}/extensibleattributedef" -d '{"name":"VM Name","type":"STRING"}'
echo

curl -k -X POST -H 'Content-Type: application/json' $CURL_AUTH "${WAPI_URL}/extensibleattributedef" -d '{"name":"Network Name","type":"STRING"}'
echo

# create a pool, zone_auth, server and topology for DTC LBDN, Pool and servers
curl -k -X POST -H 'Content-Type: application/json' $CURL_AUTH "${WAPI_URL}/dtc:server" -d "{\"name\":\"dummy-server2.com\",\"host\":\"12.10.10.1\"}"

curl -k -X PUT -H 'Content-Type: application/json'  $CURL_AUTH "${WAPI_URL}/dtc:monitor:snmp/ZG5zLmlkbnNfbW9uaXRvcl9zbm1wJHNubXA:snmp" -d '{"oids":[{"condition":"EXACT","first":"4","type":"INTEGER","oid":".1.2"}]}'
echo

server=$(curl -k -X POST -H 'Content-Type: application/json' $CURL_AUTH "${WAPI_URL}/dtc:server" -d "{\"name\":\"dummy-server.com\",\"host\":\"12.10.10.1\"}")
echo $server

pool1=$(curl -k -X POST -H 'Content-Type: application/json' $CURL_AUTH "${WAPI_URL}/dtc:pool" -d '{"name":"rrpool","lb_preferred_method":"GLOBAL_AVAILABILITY"}')
pool2=$(curl -k -X POST -H 'Content-Type: application/json' $CURL_AUTH "${WAPI_URL}/dtc:pool" -d '{"name":"pool2","lb_preferred_method":"GLOBAL_AVAILABILITY"}')

pool=$(curl -k -X POST -H 'Content-Type: application/json' $CURL_AUTH "${WAPI_URL}/dtc:pool" -d "{\"name\":\"test-pool\",\"lb_preferred_method\":\"GLOBAL_AVAILABILITY\",\"servers\":[{\"ratio\":2,\"server\":${server}}]}")
echo $pool

curl -k -X POST -H 'Content-Type: application/json' $CURL_AUTH "${WAPI_URL}/dtc:topology" -d "{\"name\":\"test-topo\",\"rules\":[{\"dest_type\":\"POOL\",\"destination_link\":"${pool}"}]}"
echo

curl -k -X POST -H 'Content-Type: application/json' $CURL_AUTH "${WAPI_URL}/dtc:topology" -d "{\"name\":\"topology_ruleset1\",\"rules\":[{\"dest_type\":\"SERVER\",\"destination_link\":"${server}"}]}"
echo

members=$(curl -k -X GET -H 'Content-Type: application/json' $CURL_AUTH "${WAPI_URL}/member")
echo
host_name=$(echo $members | grep -o '"host_name": *"[^"]*' | head -1 | awk -F'"' '{print $4}')
echo $host_name

curl -k -X POST -H 'Content-Type: application/json' $CURL_AUTH "${WAPI_URL}/zone_auth" -d "{\"fqdn\":\"test.com\",\"grid_primary\":[{\"name\":\"$host_name\"}]}"
echo

curl -k -X POST -H 'Content-Type: application/json' $CURL_AUTH "${WAPI_URL}/network" -d "{\"network\":\"17.0.0.0/24\",\"members\":[{\"_struct\":\"dhcpmember\",\"name\":\"infoblox.localdomain\"}]}"
echo