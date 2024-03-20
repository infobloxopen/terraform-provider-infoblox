package ibclient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"reflect"
	"time"
)

const MACADDR_ZERO = "00:00:00:00:00:00"

type Bool bool

func (b Bool) MarshalJSON() ([]byte, error) {
	if b {
		return json.Marshal("True")
	}

	return json.Marshal("False")
}

type EA map[string]interface{}

func (ea EA) Count() int {
	return len(ea)
}

func (ea EA) MarshalJSON() ([]byte, error) {
	m := make(map[string]interface{})
	for k, v := range ea {
		value := make(map[string]interface{})
		value["value"] = v
		m[k] = value
	}

	return json.Marshal(m)
}

func (ea *EA) UnmarshalJSON(b []byte) (err error) {
	var m map[string]map[string]interface{}

	decoder := json.NewDecoder(bytes.NewBuffer(b))
	decoder.UseNumber()
	err = decoder.Decode(&m)
	if err != nil {
		return
	}

	*ea = make(EA)
	for k, v := range m {
		val := v["value"]
		switch valType := reflect.TypeOf(val).String(); valType {
		case "json.Number":
			var i64 int64
			i64, err = val.(json.Number).Int64()
			val = int(i64)
		case "string":
			if val.(string) == "True" {
				val = Bool(true)
			} else if val.(string) == "False" {
				val = Bool(false)
			}
		case "[]interface {}":
			nval := val.([]interface{})
			nVals := make([]string, len(nval))
			for i, v := range nval {
				nVals[i] = fmt.Sprintf("%v", v)
			}
			val = nVals
		default:
			val = fmt.Sprintf("%v", val)
		}

		(*ea)[k] = val
	}

	return
}

type EASearch map[string]interface{}

func (eas EASearch) MarshalJSON() ([]byte, error) {
	m := make(map[string]interface{})
	for k, v := range eas {
		m["*"+k] = v
	}

	return json.Marshal(m)
}

type IBBase struct {
	returnFields []string
	eaSearch     EASearch
}

type IBObject interface {
	ObjectType() string
	ReturnFields() []string
	EaSearch() EASearch
	SetReturnFields([]string)
}

func (obj *IBBase) ReturnFields() []string {
	return obj.returnFields
}

func (obj *IBBase) SetReturnFields(rf []string) {
	obj.returnFields = rf
}

func (obj *IBBase) EaSearch() EASearch {
	return obj.eaSearch
}

// QueryParams is a general struct to add query params used in makeRequest
type QueryParams struct {
	forceProxy bool

	searchFields map[string]string
}

func NewQueryParams(forceProxy bool, searchFields map[string]string) *QueryParams {
	qp := QueryParams{forceProxy: forceProxy}
	if searchFields != nil {
		qp.searchFields = searchFields
	} else {
		qp.searchFields = make(map[string]string)
	}

	return &qp
}

type RequestBody struct {
	Data               map[string]interface{} `json:"data,omitempty"`
	Args               map[string]string      `json:"args,omitempty"`
	Method             string                 `json:"method"`
	Object             string                 `json:"object,omitempty"`
	EnableSubstitution bool                   `json:"enable_substitution,omitempty"`
	AssignState        map[string]string      `json:"assign_state,omitempty"`
	Discard            bool                   `json:"discard,omitempty"`
}

type SingleRequest struct {
	IBBase `json:"-"`
	Body   *RequestBody
}

type MultiRequest struct {
	IBBase `json:"-"`
	Body   []*RequestBody
}

func (r *MultiRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.Body)
}

func NewMultiRequest(body []*RequestBody) *MultiRequest {
	req := &MultiRequest{Body: body}
	return req
}

func (MultiRequest) ObjectType() string {
	return "request"
}

func NewRequest(body *RequestBody) *SingleRequest {
	req := &SingleRequest{Body: body}
	return req
}

func (SingleRequest) ObjectType() string {
	return "request"
}

type NetworkContainerNextAvailable struct {
	IBBase      `json:"-"`
	objectType  string
	Network     *NetworkContainerNextAvailableInfo `json:"network"`
	NetviewName string                             `json:"network_view,omitempty"`
	Comment     string                             `json:"comment"`
	Ea          EA                                 `json:"extattrs"`
}

func (nc *NetworkContainerNextAvailable) ObjectType() string {
	return nc.objectType
}

type NetworkContainerNextAvailableInfo struct {
	Function     string            `json:"_object_function"`
	ResultField  string            `json:"_result_field"`
	Object       string            `json:"_object"`
	ObjectParams map[string]string `json:"_object_parameters"`
	Params       map[string]uint   `json:"_parameters"`
	NetviewName  string            `json:"network_view,omitempty"`
}

func NewNetworkContainerNextAvailableInfo(netview, cidr string, prefixLen uint, isIPv6 bool) *NetworkContainerNextAvailableInfo {
	containerInfo := NetworkContainerNextAvailableInfo{
		Function:     "next_available_network",
		ResultField:  "networks",
		ObjectParams: map[string]string{"network": cidr},
		Params:       map[string]uint{"cidr": prefixLen},
		NetviewName:  netview,
	}

	if isIPv6 {
		containerInfo.Object = "ipv6networkcontainer"
	} else {
		containerInfo.Object = "networkcontainer"
	}

	return &containerInfo
}

func NewNetworkContainerNextAvailable(ncav *NetworkContainerNextAvailableInfo, isIPv6 bool, comment string, ea EA) *NetworkContainerNextAvailable {
	nc := &NetworkContainerNextAvailable{
		Network:     ncav,
		NetviewName: ncav.NetviewName,
		Ea:          ea,
		Comment:     comment,
	}

	if isIPv6 {
		nc.objectType = "ipv6networkcontainer"
	} else {
		nc.objectType = "networkcontainer"
	}
	nc.returnFields = []string{"extattrs", "network", "network_view", "comment"}

	return nc
}

type FixedAddress struct {
	IBBase      `json:"-"`
	objectType  string
	Ref         string `json:"_ref,omitempty"`
	NetviewName string `json:"network_view,omitempty"`
	Cidr        string `json:"network,omitempty"`
	Comment     string `json:"comment"`
	IPv4Address string `json:"ipv4addr,omitempty"`
	IPv6Address string `json:"ipv6addr,omitempty"`
	Duid        string `json:"duid,omitempty"`
	Mac         string `json:"mac,omitempty"`
	Name        string `json:"name,omitempty"`
	MatchClient string `json:"match_client,omitempty"`
	Ea          EA     `json:"extattrs"`
}

func (fa FixedAddress) ObjectType() string {
	return fa.objectType
}

func NewEmptyFixedAddress(isIPv6 bool) *FixedAddress {
	res := &FixedAddress{}
	res.Ea = make(EA)
	if isIPv6 {
		res.objectType = "ipv6fixedaddress"
		res.returnFields = []string{"extattrs", "ipv6addr", "duid", "name", "network", "network_view", "comment"}
	} else {
		res.objectType = "fixedaddress"
		res.returnFields = []string{"extattrs", "ipv4addr", "mac", "name", "network", "network_view", "comment"}
	}
	return res
}

func NewFixedAddress(
	netView string,
	name string,
	ipAddr string,
	cidr string,
	macOrDuid string,
	clients string,
	eas EA,
	ref string,
	isIPv6 bool,
	comment string) *FixedAddress {

	res := NewEmptyFixedAddress(isIPv6)
	res.NetviewName = netView
	res.Name = name
	res.Cidr = cidr
	res.MatchClient = clients
	res.Ea = eas
	res.Ref = ref
	res.Comment = comment
	if isIPv6 {
		res.IPv6Address = ipAddr
		res.Duid = macOrDuid
	} else {
		res.IPv4Address = ipAddr
		res.Mac = macOrDuid
	}
	return res
}

func NewEmptyDNSView() *View {
	res := &View{}
	res.returnFields = []string{"extattrs", "name", "network_view", "comment"}
	return res
}

type Network struct {
	IBBase      `json:"-"`
	objectType  string
	Ref         string `json:"_ref,omitempty"`
	NetviewName string `json:"network_view,omitempty"`
	Cidr        string `json:"network,omitempty"`
	Ea          EA     `json:"extattrs"`
	Comment     string `json:"comment"`
}

func (n Network) ObjectType() string {
	return n.objectType
}

func NewNetwork(netview string, cidr string, isIPv6 bool, comment string, ea EA) *Network {
	var res Network
	res.NetviewName = netview
	res.Cidr = cidr
	res.Ea = ea
	res.Comment = comment
	if isIPv6 {
		res.objectType = "ipv6network"
	} else {
		res.objectType = "network"
	}
	res.returnFields = []string{"extattrs", "network", "network_view", "comment"}

	return &res
}

type NetworkContainer struct {
	IBBase      `json:"-"`
	objectType  string
	Ref         string `json:"_ref,omitempty"`
	NetviewName string `json:"network_view,omitempty"`
	Cidr        string `json:"network,omitempty"`
	Comment     string `json:"comment"`
	Ea          EA     `json:"extattrs"`
}

func (nc NetworkContainer) ObjectType() string {
	return nc.objectType
}

func NewNetworkContainer(netview, cidr string, isIPv6 bool, comment string, ea EA) *NetworkContainer {
	nc := NetworkContainer{
		NetviewName: netview,
		Cidr:        cidr,
		Ea:          ea,
		Comment:     comment,
	}

	if isIPv6 {
		nc.objectType = "ipv6networkcontainer"
	} else {
		nc.objectType = "networkcontainer"
	}
	nc.returnFields = []string{"extattrs", "network", "network_view", "comment"}

	return &nc
}

// License represents license wapi object
type License struct {
	IBBase           `json:"-"`
	objectType       string
	Ref              string `json:"_ref,omitempty"`
	ExpirationStatus string `json:"expiration_status,omitempty"`
	ExpiryDate       int    `json:"expiry_date,omitempty"`
	HwID             string `json:"hwid,omitempty"`
	Key              string `json:"key,omitempty"`
	Kind             string `json:"kind,omitempty"`
	Limit            string `json:"limit,omitempty"`
	LimitContext     string `json:"limit_context,omitempty"`
	Licensetype      string `json:"type,omitempty"`
}

func (l License) ObjectType() string {
	return l.objectType
}

func NewGridLicense(license License) *License {
	result := license
	result.objectType = "license:gridwide"
	returnFields := []string{"expiration_status",
		"expiry_date",
		"key",
		"limit",
		"limit_context",
		"type"}
	result.returnFields = returnFields
	return &result
}

func NewLicense(license License) *License {
	result := license
	returnFields := []string{"expiration_status",
		"expiry_date",
		"hwid",
		"key",
		"kind",
		"limit",
		"limit_context",
		"type"}
	result.objectType = "member:license"
	result.returnFields = returnFields
	return &result
}

var mxRecordReturnFieldsList = []string{"mail_exchanger", "view", "name", "preference", "ttl", "use_ttl", "comment", "extattrs", "zone"}

func NewEmptyRecordMX() *RecordMX {
	res := &RecordMX{}
	res.returnFields = mxRecordReturnFieldsList

	return res
}

func NewRecordMX(rm RecordMX) *RecordMX {
	res := rm
	res.returnFields = mxRecordReturnFieldsList

	return &res
}

var srvRecordReturnFieldsList = []string{"name", "view", "priority", "weight", "port", "target", "ttl", "use_ttl", "comment", "extattrs", "zone"}

func NewEmptyRecordSRV() *RecordSRV {
	res := RecordSRV{}
	res.returnFields = srvRecordReturnFieldsList

	return &res
}

func NewRecordSRV(rv RecordSRV) *RecordSRV {
	res := rv
	res.returnFields = srvRecordReturnFieldsList

	return &res
}

// UnixTime is used to marshall/unmarshall epoch seconds
// presented in different parts of WAPI objects
type UnixTime struct {
	time.Time
}

func (u *UnixTime) UnmarshalJSON(b []byte) error {
	var timestamp int64
	err := json.Unmarshal(b, &timestamp)
	if err != nil {
		return err
	}
	u.Time = time.Unix(timestamp, 0)
	return nil
}

func (u UnixTime) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("%d", u.Time.Unix())), nil
}

type Dns struct {
	IBBase      `json:"-"`
	objectType  string
	Ref         string `json:"_ref,omitempty"`
	Ea          EA     `json:"extattrs"`
	Comment     string `json:"comment,omitempty"`
	HostName    string `json:"host_name,omitempty"`
	IPv4Address string `json:"ipv4addr,omitempty"`
	EnableDns   bool   `json:"enable_dns"`
}

func (d Dns) ObjectType() string {
	return d.objectType
}

func NewDns(dns Dns) *Dns {
	result := dns
	result.objectType = "member:dns"
	returnFields := []string{"enable_dns", "host_name"}
	result.returnFields = returnFields
	return &result
}

type Dhcp struct {
	IBBase      `json:"-"`
	objectType  string
	Ref         string `json:"_ref,omitempty"`
	Ea          EA     `json:"extattrs"`
	Comment     string `json:"comment,omitempty"`
	HostName    string `json:"host_name,omitempty"`
	IPv4Address string `json:"ipv4addr,omitempty"`
	EnableDhcp  bool   `json:"enable_dhcp"`
}

func (d Dhcp) ObjectType() string {
	return d.objectType
}

func NewDhcp(dhcp Dhcp) *Dhcp {
	result := dhcp
	result.objectType = "member:dhcpproperties"
	returnFields := []string{"enable_dhcp", "host_name"}
	result.returnFields = returnFields
	return &result
}
