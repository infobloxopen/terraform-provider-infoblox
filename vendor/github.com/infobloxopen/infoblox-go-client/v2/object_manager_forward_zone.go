package ibclient

import (
	"fmt"
)

func (objMgr *ObjectManager) CreateZoneForward(
	comment string,
	disable bool,
	eas EA,
	forwardTo []NameServer,
	forwardersOnly bool,
	forwardingServers []*Forwardingmemberserver,
	fqdn string,
	nsGroup string,
	view string,
	zoneFormat string) (*ZoneForward, error) {
	if fqdn == "" || forwardTo == nil {
		return nil, fmt.Errorf("FQDN and forwardTo fields are required to create a forward zone")
	}

	zoneForward := NewZoneForward(comment, disable, eas, forwardTo, forwardersOnly, forwardingServers, fqdn, nsGroup, view, zoneFormat)
	ref, err := objMgr.connector.CreateObject(zoneForward)
	if err != nil {
		return nil, err
	}
	zoneForward.Ref = ref
	return zoneForward, nil
}

func (objMgr *ObjectManager) DeleteZoneForward(ref string) (string, error) {
	return objMgr.connector.DeleteObject(ref)

}

func (objMgr *ObjectManager) GetZoneForwardByRef(ref string) (*ZoneForward, error) {
	zoneForward := NewEmptyZoneForward()
	zoneForward.SetReturnFields(append(zoneForward.ReturnFields(), "zone_format", "ns_group", "comment", "disable", "extattrs", "forwarders_only", "forwarding_servers"))
	err := objMgr.connector.GetObject(zoneForward, ref, NewQueryParams(false, nil), &zoneForward)
	if err != nil {
		return nil, err
	}
	return zoneForward, nil
}

func (objMgr *ObjectManager) GetZoneForwardFilters(filters map[string]string) ([]ZoneForward, error) {

	var res []ZoneForward
	zoneForward := NewEmptyZoneForward()
	zoneForward.SetReturnFields(append(zoneForward.ReturnFields(), "zone_format", "ns_group", "comment", "disable", "extattrs", "forwarders_only", "forwarding_servers"))

	err := objMgr.connector.GetObject(
		zoneForward, "", NewQueryParams(false, filters), &res)
	if err != nil {
		return nil, err
	}

	return res, err
}

//func (objMgr *ObjectManager) GetAllZoneForwardRecords(comment string, ea EA, fqdn string, view string, zoneFormat string) ([]ZoneForward, error) {
//	if fqdn == "" || view == "" || zoneFormat == "" {
//		return nil, fmt.Errorf("FQDN, view and zoneFormat fields are required to retrieve a forward zone")
//	}
//	var res []ZoneForward
//
//	zoneForward := NewZoneForward(comment, false, ea, nil, false, nil, fqdn, "", view, zoneFormat)
//	if ea != nil && len(ea) > 0 {
//		zoneForward.eaSearch = EASearch(ea)
//	}
//
//	sf := map[string]string{
//		"comment":     comment,
//		"fqdn":        fqdn,
//		"view":        view,
//		"zone_format": zoneFormat,
//	}
//	queryParams := NewQueryParams(false, sf)
//	err := objMgr.connector.GetObject(zoneForward, "", queryParams, &res)
//
//	if err != nil {
//		return nil, err
//	} else if res == nil || len(res) == 0 {
//		return nil, NewNotFoundError(
//			fmt.Sprintf("Zoneforward with comment: %s , fqdn: %s, zone_format: %s in view: %s is not found.", comment, fqdn, zoneFormat, view))
//	}
//	return res, nil
//}

func (objMgr *ObjectManager) UpdateZoneForward(
	ref string,
	comment string,
	disable bool,
	eas EA,
	forwardTo []NameServer,
	forwardersOnly bool,
	forwardingServers []*Forwardingmemberserver,
	nsGroup string) (*ZoneForward, error) {

	zoneForward := NewEmptyZoneForward()
	//fz, err := objMgr.GetZoneForwardByRef(ref)
	//if err != nil {
	//	return nil, err
	//}

	zoneForward.Comment = &comment
	zoneForward.Disable = &disable
	zoneForward.Ea = eas
	zoneForward.ForwardTo = forwardTo
	zoneForward.ForwardersOnly = &forwardersOnly
	zoneForward.ForwardingServers = forwardingServers
	if nsGroup != "" {
		zoneForward.NsGroup = &nsGroup
	} else {
		zoneForward.NsGroup = nil
	}

	//fqdn := fz.Fqdn
	//fz.Fqdn = ""

	new_ref, err := objMgr.connector.UpdateObject(zoneForward, ref)
	if err != nil {
		return nil, err
	}
	zoneForward.Ref = new_ref
	//fz.Fqdn = fqdn
	return zoneForward, nil

}

func NewEmptyZoneForward() *ZoneForward {
	return &ZoneForward{}
}

func NewZoneForward(comment string,
	disable bool,
	eas EA,
	forwardTo []NameServer,
	forwardersOnly bool,
	forwardingServers []*Forwardingmemberserver,
	fqdn string,
	nsGroup string,
	view string,
	zoneFormat string) *ZoneForward {

	zoneForward := NewEmptyZoneForward()
	zoneForward.Comment = &comment
	zoneForward.Disable = &disable
	zoneForward.Ea = eas
	zoneForward.ForwardTo = forwardTo
	zoneForward.ForwardersOnly = &forwardersOnly
	zoneForward.ForwardingServers = forwardingServers
	zoneForward.Fqdn = fqdn
	if nsGroup == "" {
		zoneForward.NsGroup = nil
	} else {
		zoneForward.NsGroup = &nsGroup
	}
	if view == "" {
		view = "default"
	}
	zoneForward.View = &view
	if zoneFormat == "" {
		zoneFormat = "FORWARD"
	}
	zoneForward.ZoneFormat = zoneFormat

	return zoneForward
}
