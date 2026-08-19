package infoblox

import (
	"encoding/json"
	"fmt"
	"maps"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	ibclient "github.com/infobloxopen/infoblox-go-client/v2"
)

func resourceDtcMonitorIcmp() *schema.Resource {
	var (
		monitorSchema = resourceDtcMonitorSchema()
		icmpSchema    = map[string]*schema.Schema{} // ICMP has no parameters not in base schema
	)
	maps.Copy(monitorSchema, icmpSchema) // keys in monitorSchema are overwritten

	return &schema.Resource{
		// TODO: move towards context-aware equivalents of these fields, as these are deprecated.
		Create: resourceDtcMonitorIcmpCreate,
		Read:   resourceDtcMonitorIcmpGet,
		Update: resourceDtcMonitorIcmpUpdate,
		Delete: makeResourceDtcMonitorDelete("DtcMonitorIcmp"),

		Importer: &schema.ResourceImporter{
			State: resourceDtcMonitorIcmpImporter,
		},

		Schema: monitorSchema,
	}
}

func resourceDtcMonitorIcmpCreate(d *schema.ResourceData, m interface{}) error {
	if intId := d.Get("internal_id"); intId.(string) != "" {
		return fmt.Errorf("the value of 'internal_id' field must not be set manually")
	}
	mb, err := resourceDtcMonitorGetFromTfState(d)
	if err != nil {
		return err
	}

	connector := m.(ibclient.IBConnector)
	objMgr := ibclient.NewObjectManager(connector, "Terraform", mb.tenantID)

	createdObject, err := objMgr.CreateDtcMonitorIcmp(
		mb.comment,
		mb.name,
		mb.extAttrs,
		uint32(mb.interval),
		uint32(mb.retry_down),
		uint32(mb.retry_up),
		uint32(mb.timeout))
	if err != nil {
		return fmt.Errorf("error while creating a dtc:monitor:icmp: %s", err.Error())
	}

	d.SetId(createdObject.Ref)
	if err = d.Set("ref", createdObject.Ref); err != nil {
		return err
	}

	// For compatibility reason. This field should be deprecated in the future.
	if err = d.Set("internal_id", mb.internalID); err != nil {
		return err
	}

	return resourceDtcMonitorIcmpGet(d, m)
}

func resourceDtcMonitorIcmpGet(d *schema.ResourceData, m interface{}) error {
	rec, err := searchObjectByRefOrInternalId("DtcMonitorIcmp", d, m)
	if err != nil {
		if _, ok := err.(*ibclient.NotFoundError); ok {
			d.SetId("")
			return nil
		}

		return ibclient.NewNotFoundError(fmt.Sprintf(
			"cannot find appropriate object on NIOS side for resource with ID '%s': %s;", d.Id(), err))
	}

	// set state with dtc:monitor (generic base class) attributes
	if err = resourceDtcMonitorSetTfState(d, rec); err != nil {
		return err
	}

	var dtcMonitorIcmp *ibclient.DtcMonitorIcmp
	recJson, err := json.Marshal(rec)
	if err != nil {
		return fmt.Errorf("failed to marshal DTC MonitorIcmp : %s", err.Error())
	}
	err = json.Unmarshal(recJson, &dtcMonitorIcmp)
	if err != nil {
		return fmt.Errorf("failed getting DTC MonitorIcmp : %s", err.Error())
	}

	// NOTE: setting state with icmp specific attributes not necessary, ICMP is completely covered by the base dtc:monitor

	return nil
}

func resourceDtcMonitorIcmpUpdate(d *schema.ResourceData, m interface{}) (err error) {
	var updateSuccessful bool
	defer func() {
		// Reverting the state back, in case of a failure,
		// otherwise Terraform will keep the values, which leaded to the failure,
		// in the state file.
		if !updateSuccessful {
			resourceDtcMonitorUpdateFailedResetState(d)
		}
	}()

	mb, err := resourceDtcMonitorGetFromTfState(d)
	if err != nil {
		return err
	}

	oldExtAttrsJSON, newExtAttrsJSON := d.GetChange("ext_attrs")
	newExtAttrs, err := terraformDeserializeEAs(newExtAttrsJSON.(string))
	if err != nil {
		return err
	}
	oldExtAttrs, err := terraformDeserializeEAs(oldExtAttrsJSON.(string))
	if err != nil {
		return err
	}
	mb.extAttrs = newExtAttrs

	var tenantID string
	if tempVal, found := newExtAttrs[eaNameForTenantId]; found {
		tenantID = tempVal.(string)
	}

	connector := m.(ibclient.IBConnector)
	objMgr := ibclient.NewObjectManager(connector, "Terraform", tenantID)

	rec, err := searchObjectByRefOrInternalId("DtcMonitorIcmp", d, m)
	if err != nil {
		if _, ok := err.(*ibclient.NotFoundError); !ok {
			return ibclient.NewNotFoundError(fmt.Sprintf(
				"cannot find appropriate object on NIOS side for resource with ID '%s': %s;", d.Id(), err))
		} else {
			d.SetId("")
			return nil
		}
	}
	recJson, err := json.Marshal(rec)
	if err != nil {
		return fmt.Errorf("failed to marshal dtc:monitor: %s", err.Error())
	}
	var dtcMonitor *ibclient.DtcMonitor
	err = json.Unmarshal(recJson, &dtcMonitor)
	if err != nil {
		return fmt.Errorf("failed getting dtc:monitor: %s", err.Error())
	}

	// If 'internal_id' is not set, then generate a new one and set it to the EA.
	internalId := d.Get("internal_id").(string)
	if internalId == "" {
		internalId = generateInternalId().String()
	}
	newInternalId := newInternalResourceIdFromString(internalId)
	newExtAttrs[eaNameForInternalId] = newInternalId.String()

	newExtAttrs, err = mergeEAs(dtcMonitor.Ea, newExtAttrs, oldExtAttrs, connector)
	if err != nil {
		return err
	}
	dtcMonitorIcmp, err := objMgr.UpdateDtcMonitorIcmp(
		d.Id(),
		mb.comment,
		mb.name,
		mb.extAttrs,
		uint32(mb.interval),
		uint32(mb.retry_down),
		uint32(mb.retry_up),
		uint32(mb.timeout))
	if err != nil {
		return fmt.Errorf("error updating dtc:monitor:icmp: %w", err)
	}
	updateSuccessful = true
	d.SetId(dtcMonitorIcmp.Ref)
	if err = d.Set("ref", dtcMonitorIcmp.Ref); err != nil {
		return err
	}
	if err = d.Set("internal_id", newInternalId.String()); err != nil {
		return err
	}

	return resourceDtcMonitorIcmpGet(d, m)
}

func resourceDtcMonitorIcmpImporter(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	extAttrJSON := d.Get("ext_attrs").(string)
	extAttrs, err := terraformDeserializeEAs(extAttrJSON)
	if err != nil {
		return nil, err
	}

	var tenantID string
	if tempVal, ok := extAttrs[eaNameForTenantId]; ok {
		tenantID = tempVal.(string)
	}

	connector := m.(ibclient.IBConnector)
	objMgr := ibclient.NewObjectManager(connector, "Terraform", tenantID)
	obj, err := objMgr.GetDtcMonitorIcmpByRef(d.Id())
	if err != nil {
		return nil, fmt.Errorf("getting dtc:monitor:icmp with ID: %s failed: %w", d.Id(), err)
	}

	// Set ref
	if err = d.Set("ref", obj.Ref); err != nil {
		return nil, err
	}

	if obj.Ea != nil && len(obj.Ea) > 0 {
		eaJSON, err := terraformSerializeEAs(obj.Ea)
		if err != nil {
			return nil, err
		}

		if err = d.Set("ext_attrs", eaJSON); err != nil {
			return nil, err
		}
	}

	err = resourceDtcMonitorIcmpGet(d, m)
	if err != nil {
		return nil, err
	}

	return []*schema.ResourceData{d}, nil
}
