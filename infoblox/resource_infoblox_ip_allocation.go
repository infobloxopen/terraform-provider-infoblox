package infoblox

import (
	"encoding/json"
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"sort"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	log "github.com/sirupsen/logrus"

	ibclient "github.com/infobloxopen/infoblox-go-client/v2"
)

func normalizeAlias(alias, domain string) string {
	if !strings.HasSuffix(alias, "."+domain) {
		return alias + "." + domain
	}
	return alias
}

// normalizeAndSortAliases normalizes each alias by appending the domain if missing and sorts the resulting list.
func normalizeAndSortAliases(aliases []interface{}, domain string, enableDNS bool) []string {
	var normalizedAliases []string
	for _, alias := range aliases {
		aliasStr := alias.(string)
		if enableDNS {
			aliasStr = normalizeAlias(aliasStr, domain)
		}
		normalizedAliases = append(normalizedAliases, aliasStr)
	}
	sort.Strings(normalizedAliases)
	return normalizedAliases
}

func resourceIPAllocation() *schema.Resource {
	// TODO: move towards context-aware equivalents of these fields, as these are deprecated.
	return &schema.Resource{
		Create: resourceAllocationRequest,
		Read:   resourceAllocationGet,
		Update: resourceAllocationUpdate,
		Delete: resourceAllocationRelease,

		Importer: &schema.ResourceImporter{
			State: ipAllocationImporter,
		},

		Schema: map[string]*schema.Schema{
			"network_view": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     defaultNetView,
				Description: "network view name on NIOS server.",
			},
			"dns_view": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "DNS view under which the zone has been created.",
			},
			"enable_dns": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     true,
				Description: "flag that defines if the host record is to be used for DNS purposes.",
			},
			"ipv4_cidr": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The IPv4 cidr from which an IPv4 address will be allocated.",
			},
			"ipv6_cidr": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The IPv6 cidr from which an IPv6 address will be allocated.",
				StateFunc: func(val interface{}) string {
					return normalizeIPAddress(val)
				},
			},
			"ipv4_addr": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "",
				Description: "IPv4 address of cloud instance." +
					"Set a valid IP address for static allocation and leave empty if dynamically allocated.",
			},
			"allocated_ipv4_addr": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "Value which comes from 'ipv4_addr' (if specified) or from auto-allocation function (using 'ipv4_cidr' or 'filter_params').",
			},
			"ipv6_addr": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "",
				Description: "IPv6 address of cloud instance." +
					"Set a valid IP address for static allocation and leave empty if dynamically allocated.",
				StateFunc: func(val interface{}) string {
					if val == "" {
						return ""
					}
					return normalizeIPAddress(val)
				},
			},
			"allocated_ipv6_addr": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "Value which comes from 'ipv6_addr' (if specified) or from auto-allocation function (using 'ipv6_cidr' or 'filter_params').",
			},
			"fqdn": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The host name for Host Record in FQDN format.",
			},
			"filter_params": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The parent network block's extensible attributes. This field is used for dynamic allocation along with 'ip_address_type' field.",
			},
			"ip_address_type": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The type of IP address to allocate. This filed is used only when 'filter_params' field is used. Valid values are: IPV4, IPV6, Both. Default value is IPV4",
				ValidateFunc: validation.StringInSlice([]string{
					"IPV4", "IPV6", "Both",
				}, false),
				Default: "IPV4",
				DiffSuppressFunc: func(k, oldValue, newValue string, d *schema.ResourceData) bool {
					if d.Get("filter_params") != "" && newValue == "" {
						if oldValue == "IPV4" {
							return true
						}
					} else if d.Get("filter_params") == "" {
						return true
					}
					return oldValue == newValue
				},
			},
			"ttl": {
				Type:        schema.TypeInt,
				Optional:    true,
				Default:     ttlUndef,
				Description: "TTL attribute value for the record.",
			},
			"comment": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "",
				Description: "A description of IP address allocation.",
			},
			"ext_attrs": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "",
				Description: "The extensible attributes for IP address allocation, as a map in JSON format",
			},
			"disable": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: "Disables the Host record if set to 'true'.",
			},
			"internal_id": {
				Type:     schema.TypeString,
				Computed: true,
				Description: "Internal ID of an object at NIOS side," +
					" used by Infoblox Terraform plugin to search for a NIOS's object" +
					" which corresponds to the Terraform resource.",
			},
			"ref": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "NIOS object's reference, not to be set by a user.",
			},
			"aliases": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "A set of IP allocation aliases",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				DiffSuppressFunc: func(k, oldValue, newValue string, d *schema.ResourceData) bool {
					if newValue == "0" {
						return false
					}
					if oldValue == newValue {
						return true
					}
					enableDNS := d.Get("enable_dns").(bool)
					fqdn := d.Get("fqdn").(string)
					domain := strings.Join(strings.Split(fqdn, ".")[1:], ".")
					oldAliases, newAliases := d.GetChange("aliases")
					oldAliasesNew := normalizeAndSortAliases(oldAliases.([]interface{}), domain, enableDNS)
					newAliasesNew := normalizeAndSortAliases(newAliases.([]interface{}), domain, enableDNS)
					// Compare the sorted aliases
					return strings.Join(oldAliasesNew, ",") == strings.Join(newAliasesNew, ",")
				},
			},
		},
	}
}

// This function is for retrieving a host record by either known reference or,
// if the reference points to nothing (returns 'not found'),
// by internal_id. It returns the host record itself.
//
// If err == nil then hostRec != nil,
// other cases must be considered as a serious bug.
// err == nil also means that 'internal_id' exists and
// is of a proper format.
// This function MUST NOT set any resource's properties,
// other behaviour is a bug.
func getOrFindHostRec(d *schema.ResourceData, m interface{}) (
	hostRec *ibclient.HostRecord,
	err error) {

	var (
		ref         string
		actualIntId *internalResourceId
	)

	if r, found := d.GetOk("ref"); found {
		ref = r.(string)
	} else {
		_, ref = getAltIdFields(d.Id())
	}

	if id, found := d.GetOk("internal_id"); !found {
		return nil, fmt.Errorf("internal_id value is required for the resource but it is not defined")
	} else {
		actualIntId = newInternalResourceIdFromString(id.(string))
		if actualIntId == nil {
			return nil, fmt.Errorf("internal_id value is not in a proper format")
		}
	}

	// TODO: use proper Tenant ID
	objMgr := ibclient.NewObjectManager(m.(ibclient.IBConnector), "Terraform", "")
	return objMgr.SearchHostRecordByAltId(actualIntId.String(), ref, eaNameForInternalId)
}

func resourceAllocationRequest(d *schema.ResourceData, m interface{}) error {
	networkView := d.Get("network_view").(string)
	dnsView := d.Get("dns_view").(string)
	enableDns := d.Get("enable_dns").(bool)
	fqdn := d.Get("fqdn").(string)
	if intId := d.Get("internal_id"); intId.(string) != "" {
		return fmt.Errorf("the value of 'internal_id' field must not be set manually")
	}

	ipv4Cidr := d.Get("ipv4_cidr").(string)
	ipv6Cidr := d.Get("ipv6_cidr").(string)
	ipv4Addr := d.Get("ipv4_addr").(string)
	ipv6Addr := d.Get("ipv6_addr").(string)
	nextAvailableFilter := d.Get("filter_params").(string)
	ipAdressType := d.Get("ip_address_type").(string)
	if nextAvailableFilter == "" {
		if err := d.Set("ip_address_type", ""); err != nil {
			return err

		}
	}
	if (ipv4Cidr == "" && ipv6Cidr == "" && ipv4Addr == "" && ipv6Addr == "") && nextAvailableFilter == "" {
		return fmt.Errorf("allocation through host address record creation needs an IPv4/IPv6 address" +
			" or IPv4/IPv6 cidr or filter_params")
	}

	ZeroMacAddr := "00:00:00:00:00:00"
	var macAddr string
	if ipv4Cidr != "" || ipv4Addr != "" {
		macAddr = ZeroMacAddr
	}

	aliases := d.Get("aliases").([]interface{})
	aliasStrs := make([]string, len(aliases))
	for i, alias := range aliases {
		aliasStrs[i] = alias.(string)
	}
	var ttl uint32
	useTtl := false
	tempVal := d.Get("ttl")
	tempTTL := tempVal.(int)
	if tempTTL >= 0 {
		useTtl = true
		ttl = uint32(tempTTL)
	} else if tempTTL != ttlUndef {
		return fmt.Errorf("TTL value must be 0 or higher")
	}

	comment := d.Get("comment").(string)
	disable := d.Get("disable").(bool)

	extAttrJSON := d.Get("ext_attrs").(string)
	extAttrs, err := terraformDeserializeEAs(extAttrJSON)
	if err != nil {
		return fmt.Errorf("failed to allocate IP: %w", err)
	}

	var tenantID string
	// TODO: where will we get this value from? What is its source?
	if tempVal, ok := extAttrs[eaNameForTenantId]; ok {
		tenantID = tempVal.(string)
	}

	connector := m.(ibclient.IBConnector)
	objMgr := ibclient.NewObjectManager(connector, "Terraform", tenantID)

	internalId := generateInternalId()
	extAttrs[eaNameForInternalId] = internalId.String()

	var (
		newRecordHost interface{}
		eaMap         map[string]string
	)

	if ipv4Addr == "" && ipv4Cidr == "" && ipv6Cidr == "" && ipv6Addr == "" && nextAvailableFilter != "" {
		err = json.Unmarshal([]byte(nextAvailableFilter), &eaMap)
		eaMap["network_view"] = networkView
		if err != nil {
			return fmt.Errorf("error unmarshalling extra attributes of network: %s", err)
		}
		newRecordHost, err = objMgr.AllocateNextAvailableIp(fqdn, "record:host", eaMap, nil, false, extAttrs,
			comment, disable, nil, ipAdressType, enableDns, false, "", "", networkView, dnsView, useTtl, ttl, aliasStrs)
		d.Set("ip_address_type", ipAdressType)
	} else {

		// enableDns and enableDhcp flags used to create host record with respective flags.
		// By default, enableDns is true.
		newRecordHost, err = objMgr.CreateHostRecord(enableDns, false, fqdn, networkView, dnsView, ipv4Cidr,
			ipv6Cidr, ipv4Addr, ipv6Addr, macAddr, "", useTtl, ttl, comment, extAttrs, aliasStrs, disable)
	}

	if err != nil {
		return fmt.Errorf("error while creating a host record: %s", err.Error())
	}
	hostRec := newRecordHost.(*ibclient.HostRecord)

	d.SetId(internalId.String())
	if err = d.Set("ref", hostRec.Ref); err != nil {
		return err
	}

	// For compatibility reason. This field should be deprecated in the future.
	if err = d.Set("internal_id", internalId.String()); err != nil {
		return err
	}

	if hostRec.Ipv6Addrs == nil || len(hostRec.Ipv6Addrs) < 1 {
		if err := d.Set("allocated_ipv6_addr", ""); err != nil {
			return err
		}
	} else {
		if err := d.Set("allocated_ipv6_addr", hostRec.Ipv6Addrs[0].Ipv6Addr); err != nil {
			return err
		}
	}

	alias := hostRec.Aliases
	aliasesInterface := make([]interface{}, len(alias))
	for i, a := range alias {
		aliasesInterface[i] = a
	}

	if err = d.Set("aliases", aliasesInterface); err != nil {
		return err
	}
	if hostRec.Ipv4Addrs == nil || len(hostRec.Ipv4Addrs) < 1 {
		if err := d.Set("allocated_ipv4_addr", ""); err != nil {
			return err
		}
	} else {
		if err := d.Set("allocated_ipv4_addr", hostRec.Ipv4Addrs[0].Ipv4Addr); err != nil {
			return err
		}
	}

	return resourceAllocationGet(d, m)
}

func resourceAllocationGet(d *schema.ResourceData, m interface{}) error {
	var ttl int
	obj, err := getOrFindHostRec(d, m)
	if err != nil {
		if _, ok := err.(*ibclient.NotFoundError); ok {
			d.SetId("")
			return nil
		}

		return err
	}

	_, nextAvailableFilterOk := d.GetOk("filter_params")
	if obj.Ipv6Addrs == nil || len(obj.Ipv6Addrs) < 1 {
		if err := d.Set("allocated_ipv6_addr", ""); err != nil {
			return err
		}
	} else {
		if err := d.Set("allocated_ipv6_addr", obj.Ipv6Addrs[0].Ipv6Addr); err != nil {
			return err
		}
		_, found := d.GetOk("ipv6_cidr")
		if !found && !nextAvailableFilterOk {
			if err := d.Set("ipv6_addr", obj.Ipv6Addrs[0].Ipv6Addr); err != nil {
				return err
			}
		}
	}
	if obj.Ipv4Addrs == nil || len(obj.Ipv4Addrs) < 1 {
		if err := d.Set("allocated_ipv4_addr", ""); err != nil {
			return err
		}
	} else {
		if err := d.Set("allocated_ipv4_addr", obj.Ipv4Addrs[0].Ipv4Addr); err != nil {
			return err
		}
		_, found := d.GetOk("ipv4_cidr")
		if !found && !nextAvailableFilterOk {
			if err := d.Set("ipv4_addr", obj.Ipv4Addrs[0].Ipv4Addr); err != nil {
				return err
			}
		}
	}

	alias := obj.Aliases
	aliasesInterface := make([]interface{}, len(alias))
	for i, a := range alias {
		aliasesInterface[i] = a
	}

	if err = d.Set("aliases", aliasesInterface); err != nil {
		return err
	}
	extAttrJSON := d.Get("ext_attrs").(string)
	extAttrs, err := terraformDeserializeEAs(extAttrJSON)
	if err != nil {
		return err
	}

	delete(obj.Ea, eaNameForInternalId)

	omittedEAs := omitEAs(obj.Ea, extAttrs)

	if omittedEAs != nil && len(omittedEAs) > 0 {
		eaJSON, err := terraformSerializeEAs(omittedEAs)
		if err != nil {
			return err
		}

		if err = d.Set("ext_attrs", eaJSON); err != nil {
			return err
		}
	}

	if err = d.Set("comment", obj.Comment); err != nil {
		return err
	}

	if err = d.Set("dns_view", obj.View); err != nil {
		return err
	}

	if err = d.Set("network_view", obj.NetworkView); err != nil {
		return err
	}

	if err = d.Set("enable_dns", obj.EnableDns); err != nil {
		return err
	}

	if err = d.Set("fqdn", obj.Name); err != nil {
		return err
	}

	if err = d.Set("disable", obj.Disable); err != nil {
		return err
	}

	if obj.Ttl != nil {
		ttl = int(*obj.Ttl)
	}
	if !*obj.UseTtl {
		ttl = ttlUndef
	}
	if err = d.Set("ttl", ttl); err != nil {
		return err
	}

	if err = d.Set("ref", obj.Ref); err != nil {
		return err
	}

	return nil
}

func resourceAllocationUpdate(d *schema.ResourceData, m interface{}) (err error) {
	var updateSuccessful bool
	defer func() {
		// Reverting the state back, in case of a failure,
		// otherwise Terraform will keep the values, which leaded to the failure,
		// in the state file.
		if !updateSuccessful {
			prevNetView, _ := d.GetChange("network_view")
			prevDNSView, _ := d.GetChange("dns_view")
			prevFQDN, _ := d.GetChange("fqdn")
			prevIPv4Addr, _ := d.GetChange("ipv4_addr")
			prevIPv6Addr, _ := d.GetChange("ipv6_addr")
			prevIPv4CIDR, _ := d.GetChange("ipv4_cidr")
			prevIPv6CIDR, _ := d.GetChange("ipv6_cidr")
			prevNextAvailableFilter, _ := d.GetChange("filter_params")
			prevIpAdressType, _ := d.GetChange("ip_address_type")
			prevEnableDNS, _ := d.GetChange("enable_dns")
			prevAlias, _ := d.GetChange("aliases")
			prevTTL, _ := d.GetChange("ttl")
			prevComment, _ := d.GetChange("comment")
			prevDisable, _ := d.GetChange("disable")
			prevEa, _ := d.GetChange("ext_attrs")

			_ = d.Set("network_view", prevNetView.(string))
			_ = d.Set("dns_view", prevDNSView.(string))
			_ = d.Set("fqdn", prevFQDN.(string))
			_ = d.Set("ipv4_addr", prevIPv4Addr.(string))
			_ = d.Set("ipv6_addr", prevIPv6Addr.(string))
			_ = d.Set("ipv4_cidr", prevIPv4CIDR.(string))
			_ = d.Set("ipv6_cidr", prevIPv6CIDR.(string))
			_ = d.Set("filter_params", prevNextAvailableFilter.(string))
			_ = d.Set("ip_address_type", prevIpAdressType.(string))
			_ = d.Set("enable_dns", prevEnableDNS.(bool))
			_ = d.Set("aliases", prevAlias)
			_ = d.Set("ttl", prevTTL.(int))
			_ = d.Set("comment", prevComment.(string))
			_ = d.Set("disable", prevDisable.(bool))
			_ = d.Set("ext_attrs", prevEa.(string))
		}
	}()

	hostRecObj, err := getOrFindHostRec(d, m)
	if err != nil {
		if _, ok := err.(*ibclient.NotFoundError); ok {
			d.SetId("")
			return ibclient.NewNotFoundError(fmt.Sprintf(
				"cannot find apropriate object on NIOS side for resource with ID '%s': %s;"+
					" removing the resource from Terraform state",
				d.Id(), err))
		}

		return err
	}

	if d.HasChange("internal_id") {
		return fmt.Errorf("changing the value of 'internal_id' field is not allowed")
	}
	if d.HasChange("network_view") {
		return fmt.Errorf("changing the value of 'network_view' field is not allowed")
	}
	if d.HasChange("filter_params") {
		return fmt.Errorf("changing the value of 'filter_params' field is not allowed")
	}
	if d.HasChange("ip_address_type") {
		return fmt.Errorf("changing the value of 'ip_address_type' field is not allowed")
	}

	enableDNS := d.Get("enable_dns").(bool)
	dnsView := d.Get("dns_view").(string)
	dnsView = strings.TrimSpace(dnsView)
	fqdn := d.Get("fqdn").(string)
	aliases := d.Get("aliases").([]interface{})
	aliasStrs := make([]string, len(aliases))
	for i, alias := range aliases {
		aliasStrs[i] = alias.(string)
	}
	if d.HasChange("dns_view") && !d.HasChange("enable_dns") {
		return fmt.Errorf(
			"changing the value of 'dns_view' field is allowed only for the case of changing 'enable_dns' option")
	}
	if enableDNS {
		if dnsView == disabledDNSView {
			return fmt.Errorf("a valid DNS view's name MUST be defined ('dns_view' property) once 'enable_dns' has been changed from 'false' to 'true'")
		}
		if !strings.ContainsRune(fqdn, '.') {
			return fmt.Errorf("'fqdn' value must be an FQDN without a trailing dot")
		}

	}

	// internalId != nil here, because getOrFindHostRec() checks for this and returns an error otherwise.
	internalId := newInternalResourceIdFromString(d.Get("internal_id").(string))

	ipv4Cidr := d.Get("ipv4_cidr").(string)
	ipv6Cidr := d.Get("ipv6_cidr").(string)
	ipv4Addr := d.Get("ipv4_addr").(string)
	ipv6Addr := d.Get("ipv6_addr").(string)
	_, nextAvailableFilterOk := d.GetOk("filter_params")

	// If 'ipv4_cidr' or 'ipv6_cidr' are unchanged, then nothing to update here.
	// making them empty to skip dynamic allocation of a new IP address again.
	// (This is to prevent record renewal for the case when 'cidr' is used for IP address allocation,
	// otherwise the address will be changing during every 'update' operation).
	if !d.HasChange("ipv4_cidr") {
		ipv4Cidr = ""
	}
	if !d.HasChange("ipv6_cidr") {
		ipv6Cidr = ""
	}

	var ttl uint32
	useTtl := false
	tempVal := d.Get("ttl")
	tempTTL := tempVal.(int)
	if tempTTL >= 0 {
		useTtl = true
		ttl = uint32(tempTTL)
	} else if tempTTL != ttlUndef {
		return fmt.Errorf("TTL value must be 0 or higher")
	}

	comment := d.Get("comment").(string)
	disable := d.Get("disable").(bool)

	oldExtAttrsJSON, newExtAttrsJSON := d.GetChange("ext_attrs")

	newExtAttrs, err := terraformDeserializeEAs(newExtAttrsJSON.(string))
	if err != nil {
		return err
	}

	oldExtAttrs, err := terraformDeserializeEAs(oldExtAttrsJSON.(string))
	if err != nil {
		return err
	}

	var tenantID string
	if tempVal, ok := newExtAttrs[eaNameForTenantId]; ok {
		tenantID = tempVal.(string)
	}

	connector := m.(ibclient.IBConnector)
	objMgr := ibclient.NewObjectManager(connector, "Terraform", tenantID)

	// Retrieve the IP of Host or Fixed Address record,
	// when IP is allocated using CIDR and an empty IP is passed for update.
	needIpv4Addr := ipv4Cidr == "" && ipv4Addr == ""
	needIpv6Addr := ipv6Cidr == "" && ipv6Addr == ""
	var (
		macAddr, duid string
	)
	if needIpv4Addr || needIpv6Addr {
		_, ipv4CidrFlag := d.GetOk("ipv4_cidr")
		if (ipv4CidrFlag || nextAvailableFilterOk) && len(hostRecObj.Ipv4Addrs) > 0 {
			ipv4Addr = *hostRecObj.Ipv4Addrs[0].Ipv4Addr
			if hostRecObj.Ipv4Addrs[0].Mac != nil {
				macAddr = *hostRecObj.Ipv4Addrs[0].Mac
			}
		}
		_, ipv6CidrFlag := d.GetOk("ipv6_cidr")
		if (ipv6CidrFlag || nextAvailableFilterOk) && len(hostRecObj.Ipv6Addrs) > 0 {
			ipv6Addr = *hostRecObj.Ipv6Addrs[0].Ipv6Addr
			if hostRecObj.Ipv6Addrs[0].Duid != nil {
				duid = *hostRecObj.Ipv6Addrs[0].Duid
			}
		}
	}

	newExtAttrs[eaNameForInternalId] = internalId.String()

	var (
		recIpV4Addr *ibclient.HostRecordIpv4Addr
		recIpV6Addr *ibclient.HostRecordIpv6Addr
	)
	if len(hostRecObj.Ipv4Addrs) > 0 {
		recIpV4Addr = &hostRecObj.Ipv4Addrs[0]
	}
	if len(hostRecObj.Ipv6Addrs) > 0 {
		recIpV6Addr = &hostRecObj.Ipv6Addrs[0]
	}

	enableDhcp := false

	if recIpV4Addr != nil && recIpV4Addr.EnableDhcp != nil {
		if recIpV4Addr.Mac != nil {
			macAddr = *recIpV4Addr.Mac
			enableDhcp = *recIpV4Addr.EnableDhcp
		}
	}

	if recIpV6Addr != nil && recIpV6Addr.EnableDhcp != nil {
		if recIpV6Addr.Duid != nil {
			duid = *recIpV6Addr.Duid
			enableDhcp = *recIpV6Addr.EnableDhcp
		}
	}

	hr, err := objMgr.GetHostRecordByRef(hostRecObj.Ref)
	if err != nil {
		return fmt.Errorf("failed to update IP allocation: %w", err)
	}

	mergedEAs, err := mergeEAs(hr.Ea, newExtAttrs, oldExtAttrs, connector)
	if err != nil {
		return err
	}

	hostRecObj, err = objMgr.UpdateHostRecord(
		hostRecObj.Ref,
		enableDNS,
		enableDhcp,
		fqdn,
		hostRecObj.NetworkView,
		dnsView,
		ipv4Cidr, ipv6Cidr,
		ipv4Addr, ipv6Addr,
		macAddr, duid,
		useTtl, ttl,
		comment,
		mergedEAs,
		aliasStrs, disable)
	if err != nil {
		return fmt.Errorf(
			"error while updating the host record with ID '%s': %s", d.Id(), err.Error())
	}
	updateSuccessful = true
	if err = d.Set("ref", hostRecObj.Ref); err != nil {
		return err
	}
	if err = d.Set("dns_view", hostRecObj.View); err != nil {
		return err
	}
	if err = d.Set("fqdn", hostRecObj.Name); err != nil {
		return err
	}

	if hostRecObj.Ipv6Addrs == nil || len(hostRecObj.Ipv6Addrs) < 1 {
		if err := d.Set("allocated_ipv6_addr", ""); err != nil {
			return err
		}
	} else {
		if err := d.Set("allocated_ipv6_addr", hostRecObj.Ipv6Addrs[0].Ipv6Addr); err != nil {
			return err
		}
	}
	alias := hostRecObj.Aliases
	aliasesInterface := make([]interface{}, len(alias))
	for i, a := range alias {
		aliasesInterface[i] = a
	}

	if err = d.Set("aliases", aliasesInterface); err != nil {
		return err
	}

	if hostRecObj.Ipv4Addrs == nil || len(hostRecObj.Ipv4Addrs) < 1 {
		if err := d.Set("allocated_ipv4_addr", ""); err != nil {
			return err
		}
	} else {
		if err := d.Set("allocated_ipv4_addr", hostRecObj.Ipv4Addrs[0].Ipv4Addr); err != nil {
			return err
		}
	}

	return nil
}

func resourceAllocationRelease(d *schema.ResourceData, m interface{}) error {
	if d.HasChange("network_view") {
		return fmt.Errorf("changing the value of 'network_view' field is not allowed")
	}
	if d.HasChange("dns_view") {
		return fmt.Errorf("changing the value of 'dns_view' field is not allowed")
	}
	extAttrJSON := d.Get("ext_attrs").(string)
	extAttrs, err := terraformDeserializeEAs(extAttrJSON)
	if err != nil {
		return fmt.Errorf("failed to delete network container: %w", err)
	}

	var tenantID string
	if tempVal, ok := extAttrs[eaNameForTenantId]; ok {
		tenantID = tempVal.(string)
	}

	hostRec, err := getOrFindHostRec(d, m)
	if err != nil {
		if _, ok := err.(*ibclient.NotFoundError); !ok {
			return fmt.Errorf("cannot retrieve existing record from NIOS server for the resource ID %q: %s", d.Id(), err)
		}

		// The resource seems to be deleted already,
		// let's not fail the plan's execution,
		// the corresponding NIOS object doesn't exist anyway.
		// TODO: re-align this with ip_association.
		log.Warningf(
			"unsuccessfull attempt to delete a host record for the resource ID '%s': the object cannot be found; nevertheless, the resource is still to be deleted from Terraform state", d.Id())
		d.SetId("")

		return nil
	}

	connector := m.(ibclient.IBConnector)
	objMgr := ibclient.NewObjectManager(connector, "Terraform", tenantID)
	_, err = objMgr.DeleteHostRecord(hostRec.Ref)
	if err != nil {
		return fmt.Errorf("error while releasing the resource with ID '%s': %s", d.Id(), err.Error())
	}
	d.SetId("")

	return nil
}

func ipAllocationImporter(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	var ttl int
	internalId := newInternalResourceIdFromString(d.Id())
	if internalId == nil {
		return nil, fmt.Errorf("ID value provided is not in a proper format")
	}

	d.SetId(internalId.String())
	if err := d.Set("internal_id", internalId.String()); err != nil {
		return nil, err
	}
	obj, err := getOrFindHostRec(d, m)
	if err != nil {
		return nil, err
	}

	if obj.Ipv6Addrs == nil || len(obj.Ipv6Addrs) < 1 {
		if err := d.Set("allocated_ipv6_addr", ""); err != nil {
			return nil, err
		}
	} else {
		if err := d.Set("allocated_ipv6_addr", obj.Ipv6Addrs[0].Ipv6Addr); err != nil {
			return nil, err
		}
		if _, found := d.GetOk("ipv6_cidr"); !found {
			if err := d.Set("ipv6_addr", obj.Ipv6Addrs[0].Ipv6Addr); err != nil {
				return nil, err
			}
		}
	}
	if obj.Ipv4Addrs == nil || len(obj.Ipv4Addrs) < 1 {
		if err := d.Set("allocated_ipv4_addr", ""); err != nil {
			return nil, err
		}
	} else {
		if err := d.Set("allocated_ipv4_addr", obj.Ipv4Addrs[0].Ipv4Addr); err != nil {
			return nil, err
		}
		if _, found := d.GetOk("ipv4_cidr"); !found {
			if err := d.Set("ipv4_addr", obj.Ipv4Addrs[0].Ipv4Addr); err != nil {
				return nil, err
			}
		}
	}

	extAttrJSON := d.Get("ext_attrs").(string)
	_, err = terraformDeserializeEAs(extAttrJSON)
	if err != nil {
		return nil, err
	}

	delete(obj.Ea, eaNameForInternalId)

	if obj.Ea != nil && len(obj.Ea) > 0 {
		eaJSON, err := terraformSerializeEAs(obj.Ea)
		if err != nil {
			return nil, err
		}

		if err = d.Set("ext_attrs", eaJSON); err != nil {
			return nil, err
		}
	}

	alias := obj.Aliases
	aliasesInterface := make([]interface{}, len(alias))
	for i, a := range alias {
		aliasesInterface[i] = a
	}

	if err = d.Set("aliases", aliasesInterface); err != nil {
		return nil, err
	}

	if err = d.Set("comment", obj.Comment); err != nil {
		return nil, err
	}

	if err = d.Set("dns_view", obj.View); err != nil {
		return nil, err
	}

	if err = d.Set("network_view", obj.NetworkView); err != nil {
		return nil, err
	}

	if err = d.Set("enable_dns", obj.EnableDns); err != nil {
		return nil, err
	}

	if err = d.Set("fqdn", obj.Name); err != nil {
		return nil, err
	}

	if err = d.Set("disable", obj.Disable); err != nil {
		return nil, err
	}

	if err = d.Set("ip_address_type", ""); err != nil {
		return nil, err
	}

	if obj.Ttl != nil {
		ttl = int(*obj.Ttl)
	}
	if !*obj.UseTtl {
		ttl = ttlUndef
	}
	if err = d.Set("ttl", ttl); err != nil {
		return nil, err
	}

	if err = d.Set("ref", obj.Ref); err != nil {
		return nil, err
	}

	return []*schema.ResourceData{d}, nil
}
