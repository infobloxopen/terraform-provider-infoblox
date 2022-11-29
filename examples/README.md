This directory contains example files for the current version of Terraform Infoblox Provider plugin:

* 'resources' and 'datasources' directories are about the resources and data sources, supported by the plugin; except for the following resources (deprecated):
  * `infoblox_ipv4_allocation`
  * `infoblox_ipv6_allocation`
  * `infoblox_ipv4_association`
  * `infoblox_ipv6_association`
* 'integration' directory contains example files to show the cases of integration with cloud environments; files under `archived` sub-directory are of older version and are not maintained anymore.

For the examples in 'resources' and 'datasources' directories, the prerequisites (on NIOS side) are:

   * network views: `default` (exists by default after NIOS is installed), `nondefault_netview`
   * DNS views:
     * in the default network view:
       * `default` (should already exist)
       * `nondefault_dnsview1`
     * in the network view `nondefault_netview`:
       * `default.nondefault_netview` (will be created automatically via UI when creating the network view `nondefault_netview`)
       * `nondefault_dnsview2`
   * empty forward-mapping DNS zones:
     * DNS view `default`: `example1.org`
     * DNS view `nondefault_dnsview1`: `example2.org`
     * DNS view `default.nondefault_netview`: `example3.org`
     * DNS view `nondefault_dnsview2`: `example4.org`
   * empty reverse-mapping DNS zones, in all the above-mentioned DNS views:
     * 10.in-addr.arpa
     * 0.0.0.0.0.0.0.0.3.9.f.1.2.0.0.2.ip6.arpa
   * there must be no networks, network containers, subnets, ranges, etc, which fall under the CIDRs 10.0.0.0/8 and 2002:1f93::/64.
   * EA of type 'Terraform Internal ID' of type 'string'.
   * Cloud Platform Management license or appropriate cloud-related extensible attributes, created manually.

You may use the script `create-prerequisites.sh` in this directory to create them. Please, do not forget to review which lines you do need in it,
and which are to be commented out. Environment variable NIOS_SERVER, NIOS_USER and NIOS_PASSWORD are for defining access credentials and
the server's address and port number. Other settings can be altered by editing the script.

The examples are written in the way that they can be used all together or partially. In the latter case, not all the prerequisites are required.
'provider.tf' file contains credentials to access NIOS server to run the examples on. You have to change appropriate values to those which correspond to your server.
