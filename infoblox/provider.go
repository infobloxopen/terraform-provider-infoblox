package infoblox

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	log "github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	ibclient "github.com/infobloxopen/infoblox-go-client/v2"
	"math"
	"reflect"
	"sort"
	"strings"
	"time"
)

// Common parameters
const (
	ttlUndef            = math.MinInt32
	eaNameForInternalId = "Terraform Internal ID"
	eaNameForTenantId   = "Tenant ID"
	altIdSeparator      = "|"

	defaultDNSView  = "default"
	disabledDNSView = "  "
	defaultNetView  = "default"
)

// Internal ID represents an immutable ID during resource's lifecycle.
// NIOS object's reference may get changed, sometimes this is a problem:
//
//	when more than one TF resources have the same NIOS WAPI object as a backend,
//	changing reference to the object invalidates the old reference,
//	which needs to be changed for all appropriate TF resources.
//	Doing this is problematic.
//	An example of such resources: a pair of infoblox_ipvX_allocation/infoblox_ipvX_association.
//	They both must relate to a single host record on NIOS side.
//
// Important requirement: the text representing an internal ID must not contain '|' sign,
//
//	or in general: the sign (or a sequence of) which is defined by altIdSeparator constant.
type internalResourceId struct {
	value uuid.UUID
}

func (id *internalResourceId) Equal(id2 *internalResourceId) bool {
	if id2 == nil {
		panic("the argument must not be nil")
	}
	if id == nil {
		return false
	}
	return id.value.String() == id2.value.String()
}

func (id *internalResourceId) String() string {
	if id == nil {
		return ""
	}
	return id.value.String()
}

// Returns a pointer to parsed internal resource ID, nil otherwise.
func newInternalResourceIdFromString(id string) *internalResourceId {
	newUUID, err := uuid.Parse(id)
	if err != nil {
		log.Error(context.Background(), "cannot parse internal ID", map[string]interface{}{
			"internal ID": id,
			"error":       err.Error()})
		return nil
	}

	return &internalResourceId{value: newUUID}
}

func generateInternalId() *internalResourceId {
	uuid_new, err := uuid.NewRandom()
	if err != nil {
		panic(err)
	}
	return &internalResourceId{value: uuid_new}
}

// A separate function to abstract from the nature of internal ID,
// from particular format.
func isValidInternalId(internalId string) bool {
	_, err := uuid.Parse(internalId)
	if err != nil {
		return false
	}

	return true
}

func generateAltId(internalId *internalResourceId, ref string) string {
	if internalId == nil {
		panic("the argument must not be nil")
	}
	return fmt.Sprintf(
		"%s%s%s",
		internalId.String(), altIdSeparator, ref)
}

func getAltIdFields(altId string) (internalId *internalResourceId, ref string) {
	idParts := strings.Split(altId, altIdSeparator)
	switch len(idParts) {
	case 1:
		if isValidInternalId(idParts[0]) {
			internalId = newInternalResourceIdFromString(idParts[0])
		} else {
			ref = strings.TrimSpace(idParts[0])
		}
	case 2:
		if isValidInternalId(idParts[0]) {
			internalId = newInternalResourceIdFromString(idParts[0])
			ref = strings.TrimSpace(idParts[1])
		} else {
			ref = strings.TrimSpace(idParts[0])
		}
	}

	return
}

// This function checks if the text string has any trailing or leading spaces.
func checkAndTrimSpaces(text string) (string, bool) {
	newText := strings.TrimSpace(text)
	return newText, text != newText
}

const errMsgFormatLeadingTrailingSpaces = "leading or trailing spaces are not allowed for the '%s' field"

func isNotFoundError(err error) bool {
	if _, notFoundErr := err.(*ibclient.NotFoundError); notFoundErr {
		return true
	}

	return false
}

func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"server": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("INFOBLOX_SERVER", nil),
				Description: "Infoblox server IP address.",
			},
			"username": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("INFOBLOX_USERNAME", nil),
				Description: "User to authenticate with Infoblox server.",
			},
			"password": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("INFOBLOX_PASSWORD", nil),
				Description: "Password to authenticate with Infoblox server.",
			},
			"wapi_version": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("WAPI_VERSION", "2.12.3"),
				Description: "WAPI Version of Infoblox server defaults to v2.12.3",
			},
			"port": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("PORT", "443"),
				Description: "Port number used for connection for Infoblox Server.",
			},

			"sslmode": {
				Type:        schema.TypeBool,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("SSLMODE", "false"),
				Description: "If set, Infoblox client will permit unverifiable SSL certificates.",
			},
			"connect_timeout": {
				Type:        schema.TypeInt,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("CONNECT_TIMEOUT", 60),
				Description: "Maximum wait for connection, in seconds. Zero or not specified means wait indefinitely.",
			},
			"pool_connections": {
				Type:        schema.TypeInt,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("POOL_CONNECTIONS", "10"),
				Description: "Maximum number of connections to establish to the Infoblox server. Zero means unlimited.",
			},
		},

		ResourcesMap: map[string]*schema.Resource{
			"infoblox_network_view":           resourceNetworkView(),
			"infoblox_ipv4_network_container": resourceIPv4NetworkContainer(),
			"infoblox_ipv6_network_container": resourceIPv6NetworkContainer(),
			"infoblox_ipv4_network":           resourceIPv4Network(),
			"infoblox_ipv6_network":           resourceIPv6Network(),
			"infoblox_ip_allocation":          resourceIPAllocation(),
			"infoblox_ip_association":         resourceIpAssociationInit(),
			"infoblox_a_record":               resourceARecord(),
			"infoblox_aaaa_record":            resourceAAAARecord(),
			"infoblox_cname_record":           resourceCNAMERecord(),
			"infoblox_ptr_record":             resourcePTRRecord(),
			"infoblox_zone_delegated":         resourceZoneDelegated(),
			"infoblox_txt_record":             resourceTXTRecord(),
			"infoblox_mx_record":              resourceMXRecord(),
			"infoblox_srv_record":             resourceSRVRecord(),
			"infoblox_dns_view":               resourceDNSView(),
			"infoblox_zone_auth":              resourceZoneAuth(),
			"infoblox_zone_forward":           resourceZoneForward(),
			"infoblox_dtc_lbdn":               resourceDtcLbdnRecord(),
			"infoblox_dtc_pool":               resourceDtcPool(),
			"infoblox_dtc_server":             resourceDtcServer(),
			"infoblox_ipv4_fixed_address":     resourceFixedRecord(),
			"infoblox_alias_record":           resourceAliasRecord(),
			"infoblox_ns_record":              resourceNSRecord(),
			"infoblox_ipv4_range":             resourceRange(),
			"infoblox_ipv4_range_template":    resourceRangeTemplate(),
			"infoblox_ipv4_shared_network":    resourceIpv4SharedNetwork(),
		},
		DataSourcesMap: map[string]*schema.Resource{
			"infoblox_ipv4_network":           dataSourceIPv4Network(),
			"infoblox_ipv6_network":           dataSourceIPv6Network(),
			"infoblox_ipv4_network_container": dataSourceIpv4NetworkContainer(),
			"infoblox_ipv6_network_container": dataSourceIpv6NetworkContainer(),
			"infoblox_network_view":           dataSourceNetworkView(),
			"infoblox_a_record":               dataSourceARecord(),
			"infoblox_aaaa_record":            dataSourceAAAARecord(),
			"infoblox_cname_record":           dataSourceCNameRecord(),
			"infoblox_ptr_record":             dataSourcePtrRecord(),
			"infoblox_zone_delegated":         dataSourceZoneDelegated(),
			"infoblox_txt_record":             dataSourceTXTRecord(),
			"infoblox_mx_record":              dataSourceMXRecord(),
			"infoblox_srv_record":             dataSourceSRVRecord(),
			"infoblox_host_record":            dataSourceHostRecord(),
			"infoblox_zone_auth":              dataSourceZoneAuth(),
			"infoblox_dns_view":               dataSourceDNSView(),
			"infoblox_zone_forward":           dataSourceZoneForward(),
			"infoblox_dtc_lbdn":               dataSourceDtcLbdnRecord(),
			"infoblox_dtc_pool":               datasourceDtcPool(),
			"infoblox_dtc_server":             dataSourceDtcServer(),
			"infoblox_ipv4_fixed_address":     dataSourceFixedAddress(),
			"infoblox_alias_record":           dataSourceAliasRecord(),
			"infoblox_ns_record":              dataSourceNSRecord(),
			"infoblox_ipv4_range":             dataSourceRange(),
			"infoblox_ipv4_range_template":    dataSourceRangeTemplate(),
			"infoblox_ipv4_shared_network":    dataSourceIpv4SharedNetwork(),
		},
		ConfigureContextFunc: providerConfigure,
	}
}

func providerConfigure(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {

	if d.Get("password") == "" {
		return nil, diag.Diagnostics{diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Export the required INFOBLOX_PASSWORD environment variable to set the password.",
		}}
	}

	seconds := int64(d.Get("connect_timeout").(int))
	hostConfig := ibclient.HostConfig{
		Host:    d.Get("server").(string),
		Port:    d.Get("port").(string),
		Version: d.Get("wapi_version").(string),
	}

	authConfig := ibclient.AuthConfig{
		Username: d.Get("username").(string),
		Password: d.Get("password").(string),
	}

	transportConfig := ibclient.TransportConfig{
		SslVerify:           d.Get("sslmode").(bool),
		HttpRequestTimeout:  time.Duration(seconds),
		HttpPoolConnections: d.Get("pool_connections").(int),
	}

	requestBuilder := &ibclient.WapiRequestBuilder{}
	requestor := &ibclient.WapiHttpRequestor{}

	// TODO: reconsider. For the case when there is a need to keep more data than just a go-client's Connector.
	conn, err := ibclient.NewConnector(hostConfig, authConfig, transportConfig, requestBuilder, requestor)
	if err != nil {
		return nil, diag.Diagnostics{diag.Diagnostic{Summary: err.Error()}}
	}

	// Check and Create Pre-requisites
	err = checkAndCreatePreRequisites(conn)
	if err != nil {
		return nil, diag.Diagnostics{diag.Diagnostic{Summary: err.Error()}}
	}
	return conn, nil
}

// filterFromMap generates filter map for NIOS query parameters from a terraform map[string]interface{}
func filterFromMap(filtersMap map[string]interface{}) map[string]string {
	filters := make(map[string]string, len(filtersMap))

	for k, v := range filtersMap {
		filters[k] = v.(string)
	}

	return filters
}

// terraformSerializeEAs will convert ibclient.EA to a JSON-formatted string,
// which is generally used as a value for 'ext_attrs' terraform fields.
func terraformSerializeEAs(ea ibclient.EA) (string, error) {
	delete(ea, eaNameForInternalId)
	eaMap := (map[string]interface{})(ea)
	if len(eaMap) == 0 {
		return "", nil
	}
	eaJSON, err := json.Marshal(eaMap)
	if err != nil {
		return "", err
	}

	return string(eaJSON), nil
}

// terraformDeserializeEAs converts JSON-formatted string
// of extensible attributes to a map
func terraformDeserializeEAs(extAttrJSON string) (map[string]interface{}, error) {
	extAttrs := make(map[string]interface{})
	if extAttrJSON != "" {
		if err := json.Unmarshal([]byte(extAttrJSON), &extAttrs); err != nil {
			return nil, fmt.Errorf("cannot process 'ext_attrs' field: %w", err)
		}
	}
	if extAttrs == nil {
		extAttrs = make(map[string]interface{})
	}
	return extAttrs, nil
}

// omitEAs will omit NIOS-side EAs that are not present on the terraform-provider side.
// Should be used for read operations.
func omitEAs(niosEAs, terraformEAs map[string]interface{}) map[string]interface{} {
	// ToDo: When EA inheritance is implemented on the go-client side, only inherited EAs should be omitted here.
	res := niosEAs
	for attrName, _ := range niosEAs {
		if _, ok := terraformEAs[attrName]; !ok {
			delete(res, attrName)
		}
	}

	return res
}

// mergeEAs merges omitted NIOS-side EAs with EAs specified in terraform configuration.
// Should be used in update functions.
func mergeEAs(niosEAs, newTerraformEAs, oldTerraformEAs map[string]interface{}, conn ibclient.IBConnector) (ibclient.EA, error) {
	res := map[string]interface{}{}
	for key, niosVal := range niosEAs {
		// If EA is present on the NIOS side, and there's no attempt to
		// change a value of this EA by the terraform user, use EA value from NIOS

		// If EA is required returns true, else returns false
		req := checkEARequirement(key, conn)

		if newTfVal, newTfValFound := newTerraformEAs[key]; !newTfValFound {
			if _, oldTfValFound := oldTerraformEAs[key]; !oldTfValFound {
				res[key] = niosVal
			}
			_, oldTfValFound := oldTerraformEAs[key]
			if req && oldTfValFound {
				return nil, fmt.Errorf("%s is required attribute, can't be removed", key)
			}

		} else {
			if req && newTfVal == "" {
				return nil, fmt.Errorf("%s is required attribute, can't be empty", key)
			}
			res[key] = newTfVal
		}
	}

	// Merge EAs, added to the terraform configuration
	for key, newTfVal := range newTerraformEAs {
		if _, ok := res[key]; !ok {
			res[key] = newTfVal
		}
	}

	return res, nil
}

func checkEARequirement(name string, conn ibclient.IBConnector) bool {
	eadef := &ibclient.EADefinition{}
	eadef.SetReturnFields(append(eadef.ReturnFields(), "flags"))

	sf := map[string]string{
		"name": name,
	}
	qp := ibclient.NewQueryParams(false, sf)
	var res []ibclient.EADefinition

	err := conn.GetObject(eadef, "", qp, &res)
	if err != nil {
		fmt.Errorf("failed to get EA definition")
	}
	result := &res[0]
	if result.Flags != nil {
		if strings.Contains(*result.Flags, "M") {
			return true
		}
	}
	return false
}

// Check Pre-requisites for the provider and create if not present
func checkAndCreatePreRequisites(conn ibclient.IBConnector) error {
	// 1. Create EA Definition for Internal ID if not present.

	objMgr := ibclient.NewObjectManager(conn, "Terraform", "")

	// Check if EA Definition for Internal ID is present
	_, err := objMgr.GetEADefinition(eaNameForInternalId)
	// Check for 404 error and create EA Definition if not present
	if isNotFoundError(err) {
		// Create EA Definition
		var EA ibclient.EADefinition
		var ea_string = eaNameForInternalId
		var flags = "CR"
		var comment = "Internal ID for Terraform Resource"
		EA.Name = &ea_string
		EA.Type = "STRING"
		EA.Flags = &flags
		EA.Comment = &comment
		_, err = objMgr.CreateEADefinition(EA)
		if err != nil {
			return err
		}
	}
	return nil
}

// Fetch Resource using the Ref | Terraform Internal ID

//Func to search the object using the ref or internal_id

func searchObjectByRefOrInternalId(objType string, d *schema.ResourceData, m interface{}) (
	record interface{},
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

	if id, found := d.GetOk("internal_id"); found {
		actualIntId = newInternalResourceIdFromString(id.(string))
		if actualIntId == nil {
			return nil, fmt.Errorf("internal_id value is not in a proper format")
		}
	}

	extAttrJSON := d.Get("ext_attrs").(string)
	extAttrs, err := terraformDeserializeEAs(extAttrJSON)
	if err != nil {
		return nil, err
	}

	var tenantID string
	tempVal, found := extAttrs[eaNameForTenantId]
	if found {
		tenantID = tempVal.(string)
	}

	objMgr := ibclient.NewObjectManager(m.(ibclient.IBConnector), "Terraform", tenantID)
	return objMgr.SearchObjectByAltId(objType, ref, actualIntId.String(), eaNameForInternalId)
}

func CompareSortedList(oldList interface{}, newList interface{}, key1 string, key2 string) bool {
	oldListSlice, okOld := oldList.([]interface{})
	newListSlice, okNew := newList.([]interface{})
	if !okOld || !okNew {
		return false
	}
	// If both lists are empty, they are equal
	if len(oldListSlice) == 0 && len(newListSlice) == 0 {
		return true
	}
	if len(oldListSlice) == 0 || len(newListSlice) == 0 {
		return false
	}
	// Determine the type of the first element
	switch oldListSlice[0].(type) {
	case string:
		return sortAndCompareStringSlices(oldListSlice, newListSlice)

	case map[string]interface{}:
		sortByKeys(oldListSlice, key1, key2)
		sortByKeys(newListSlice, key1, key2)
	}
	return reflect.DeepEqual(oldListSlice, newListSlice)
}

// sortAndCompareStringSlices sorts two slices of strings and compares them
func sortAndCompareStringSlices(oldListSlice, newListSlice []interface{}) bool {
	oldStrs := make([]string, len(oldListSlice))
	newStrs := make([]string, len(newListSlice))

	for i, v := range oldListSlice {
		oldStrs[i] = v.(string)
	}
	for i, v := range newListSlice {
		newStrs[i] = v.(string)
	}

	sort.Strings(oldStrs)
	sort.Strings(newStrs)

	return reflect.DeepEqual(oldStrs, newStrs)
}

func sortByKeys(list []interface{}, key1, key2 string) {
	sort.Slice(list, func(i, j int) bool {
		slice1, ok1 := list[i].(map[string]interface{})
		slice2, ok2 := list[j].(map[string]interface{})
		if !ok1 || !ok2 {
			return false
		}

		// Compare key1 first, then key2
		if slice1[key1].(string) == slice2[key1].(string) {
			return slice1[key2].(string) < slice2[key2].(string)
		}
		return slice1[key1].(string) < slice2[key1].(string)
	})
}

// sortOptions sorts a slice of DHCP options by the specified field.
func sortOptions(options []interface{}, field string) {
	sort.SliceStable(options, func(i, j int) bool {
		return options[i].(map[string]interface{})[field].(string) < options[j].(map[string]interface{})[field].(string)
	})
}

// isDefault checks if the given option is a default DHCP option.
func isDefault(opt map[string]interface{}) bool {
	return opt["name"] == "dhcp-lease-time" && opt["num"] == 51 && opt["use_option"] == false && opt["value"] == "43200" && opt["vendor_class"] == "DHCP"
}

func optimizeDhcpOptions(list1 []interface{}, list2 []interface{}) []interface{} {

	sortOptions(list1, "name")
	sortOptions(list2, "name")
	var optimizedList []interface{}

	// Create a map of new options for quick lookup
	newOptionsMap := make(map[string]map[string]interface{})
	for _, newOpt := range list2 {
		optMap, ok := newOpt.(map[string]interface{})
		if ok {
			if name, exists := optMap["name"].(string); exists {
				newOptionsMap[name] = optMap
			}
		}
	}

	// Create a map of existing options in oldList for quick lookup
	oldOptionsMap := make(map[string]map[string]interface{})
	for _, oldOpt := range list1 {
		oldOptMap, ok := oldOpt.(map[string]interface{})
		if ok {
			if name, exists := oldOptMap["name"].(string); exists {
				oldOptionsMap[name] = oldOptMap
			}
		}
	}

	// Iterate through oldList to update subfields if there are changes
	specialNames := map[string]bool{
		"routers":                  true,
		"router-templates":         true,
		"domain-name-servers":      true,
		"domain-name":              true,
		"broadcast-address-offset": true,
		"dhcp6.name-servers":       true,
		"broadcast-address":        true,
		"dhcp-lease-time":          true,
	}

	for _, oldOpt := range list1 {
		oldOptMap, ok := oldOpt.(map[string]interface{})
		if !ok {
			continue
		}
		name, exists := oldOptMap["name"].(string)
		if !exists {
			continue
		}

		if specialNames[name] {
			// Handle special options
			if newOptMap, found := newOptionsMap[name]; found {
				// Update subfields in oldOptMap with values from newOptMap
				for key, value := range newOptMap {
					oldOptMap[key] = value
				}
			} else {
				// If the option is not found in newList(default dhcp-lease-time), don't do anything
				if name == "dhcp-lease-time" {
					oldOptMap["value"] = "43200"
					oldOptMap["use_option"] = false
				} else {
					// if Option is removed from tf file, set its value to an empty string and use_option to false
					oldOptMap["value"] = ""
					oldOptMap["use_option"] = false
				}
			}
			optimizedList = append(optimizedList, oldOptMap)
		} else {
			// Handle custom DHCP options
			if newOptMap, found := newOptionsMap[name]; found {
				// Check for changes in subfields and update oldOptMap with new values
				for key, newValue := range newOptMap {
					// if there are changes in the subfields of custom DHCP options or new DHCP options are provided in the tf file, update the oldOptMap
					if oldValue, exists := oldOptMap[key]; !exists || !reflect.DeepEqual(oldValue, newValue) {
						oldOptMap[key] = newValue
					}
				}
				optimizedList = append(optimizedList, oldOptMap)
			}
		}
	}

	// Iterate through newList to find options present in newList and not in oldList
	for _, newOpt := range list2 {
		newOptMap, ok := newOpt.(map[string]interface{})
		if ok {
			if name, exists := newOptMap["name"].(string); exists {
				if _, found := oldOptionsMap[name]; !found && newOptMap["name"] != "" {
					// Option is in newList but not in oldList, add it to oldList (options newly added to tf file)
					optimizedList = append(optimizedList, newOptMap)
				}
			}
		}
	}

	return optimizedList
}
