package infoblox

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	log "github.com/sirupsen/logrus"

	ibclient "github.com/infobloxopen/infoblox-go-client/v2"
)

// DtcMonitor: This file implements code that can be shared across concrete dtc:monitor subclasses, e.g. DtcMonitorIcmp, DtcMonitorHttp, ...

func resourceDtcMonitorSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"comment": {
			Type:        schema.TypeString,
			Optional:    true,
			Default:     "",
			Description: "Comment for this DTC monitor; maximum 256 characters.",
		},
		"ext_attrs": {
			Type:        schema.TypeString,
			Optional:    true,
			Default:     "",
			Description: "Extensible attributes associated with the object, as a map in JSON format",
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
		"interval": {
			Type:        schema.TypeInt,
			Description: "The interval for a health check.",
			Optional:    true,
			Default:     5,
		},
		"name": {
			Type:        schema.TypeString,
			Description: "The display name for this DTC monitor. Values with leading or trailing white space are not valid for this field.",
			Required:    true,
		},
		"retry_down": {
			Type:        schema.TypeInt,
			Description: "The number of how many times the server should appear as “DOWN” to be treated as dead after it was alive.",
			Optional:    true,
			Default:     1,
		},
		"retry_up": {
			Type:        schema.TypeInt,
			Description: "The number of many times the server should appear as “UP” to be treated as alive after it was dead.",
			Optional:    true,
			Default:     1,
		},
		"timeout": {
			Type:        schema.TypeInt,
			Description: "The timeout for a health check.",
			Optional:    true,
			Default:     15,
		},
	}
}

type DtcMonitorBase struct {
	extAttrs   map[string]any
	comment    string
	interval   uint32
	name       string
	retry_down uint32
	retry_up   uint32
	timeout    uint32
	tenantID   string
	internalID string
}

func resourceDtcMonitorGetFromTfState(d *schema.ResourceData) (DtcMonitorBase, error) {
	extAttrJSON := d.Get("ext_attrs").(string)
	extAttrs, err := terraformDeserializeEAs(extAttrJSON)
	if err != nil {
		return DtcMonitorBase{}, err
	}

	var tenantID string
	if tempVal, ok := extAttrs[eaNameForTenantId]; ok {
		tenantID = tempVal.(string)
	}

	internalId := generateInternalId()
	extAttrs[eaNameForInternalId] = internalId.String()

	return DtcMonitorBase{
		extAttrs:   extAttrs,
		comment:    d.Get("comment").(string),
		interval:   uint32(d.Get("interval").(int)),
		name:       d.Get("name").(string),
		retry_down: uint32(d.Get("retry_down").(int)),
		retry_up:   uint32(d.Get("retry_up").(int)),
		timeout:    uint32(d.Get("timeout").(int)),
		tenantID:   tenantID,
		internalID: internalId.String(),
	}, nil

}

func resourceDtcMonitorSetTfState(d *schema.ResourceData, apiResponse any) error {
	recJson, err := json.Marshal(apiResponse)
	if err != nil {
		return fmt.Errorf("failed to marshal DTC Monitor : %s", err.Error())
	}
	var dtcMonitor *ibclient.DtcMonitor
	err = json.Unmarshal(recJson, &dtcMonitor)
	if err != nil {
		return fmt.Errorf("failed getting DTC Monitor : %s", err.Error())
	}

	extAttrJSON := d.Get("ext_attrs").(string)
	extAttrs, err := terraformDeserializeEAs(extAttrJSON)
	if err != nil {
		return err
	}
	delete(dtcMonitor.Ea, eaNameForInternalId)
	omittedEAs := omitEAs(dtcMonitor.Ea, extAttrs)
	if omittedEAs != nil && len(omittedEAs) > 0 {
		eaJSON, err := terraformSerializeEAs(omittedEAs)
		if err != nil {
			return err
		}
		if err = d.Set("ext_attrs", eaJSON); err != nil {
			return err
		}
	}
	if err = d.Set("comment", dtcMonitor.Comment); err != nil {
		return err
	}
	if err = d.Set("interval", dtcMonitor.Interval); err != nil {
		return err
	}
	if err = d.Set("name", dtcMonitor.Name); err != nil {
		return err
	}
	if err = d.Set("retry_down", dtcMonitor.RetryDown); err != nil {
		return err
	}
	if err = d.Set("retry_up", dtcMonitor.RetryUp); err != nil {
		return err
	}
	if err = d.Set("timeout", dtcMonitor.Timeout); err != nil {
		return err
	}
	// NOTE: .Ref contains a dtc:monitor (which must not be read), .Monitor is the actual ref to the concrete subclass
	if err = d.Set("ref", dtcMonitor.Monitor); err != nil {
		return err
	}
	return nil
}

func resourceDtcMonitorUpdateFailedResetState(d *schema.ResourceData) {
	prevComment, _ := d.GetChange("comment")
	prevEa, _ := d.GetChange("ext_attrs")
	prevInterval, _ := d.GetChange("interval")
	prevName, _ := d.GetChange("name")
	prevRetryDown, _ := d.GetChange("retry_down")
	prevRetryUp, _ := d.GetChange("retry_up")
	prevTimeout, _ := d.GetChange("timeout")

	_ = d.Set("comment", prevComment.(string))
	_ = d.Set("ext_attrs", prevEa.(string))
	_ = d.Set("interval", prevInterval.(int))
	_ = d.Set("name", prevName.(string))
	_ = d.Set("retry_down", prevRetryDown.(int))
	_ = d.Set("retry_up", prevRetryUp.(int))
	_ = d.Set("timeout", prevTimeout.(int))
}

func makeResourceDtcMonitorDelete(objType string) func(d *schema.ResourceData, m interface{}) error {
	return func(d *schema.ResourceData, m interface{}) error {
		extAttrJSON := d.Get("ext_attrs").(string)
		extAttrs, err := terraformDeserializeEAs(extAttrJSON)
		if err != nil {
			return fmt.Errorf("failed to delete dtc monitor: %w", err)
		}

		var tenantID string
		if tempVal, ok := extAttrs[eaNameForTenantId]; ok {
			tenantID = tempVal.(string)
		}
		connector := m.(ibclient.IBConnector)
		objMgr := ibclient.NewObjectManager(connector, "Terraform", tenantID)

		rec, err := searchObjectByRefOrInternalId(objType, d, m)
		if err != nil {
			if _, ok := err.(*ibclient.NotFoundError); !ok {
				return fmt.Errorf("cannot retrieve existing dtc:monitor from NIOS server for the resource ID %q: %s", d.Id(), err)
			}

			// The resource seems to be deleted already,
			// let's not fail the plan's execution,
			// the corresponding NIOS object doesn't exist anyway.
			log.Warningf(
				"unsuccessfull attempt to delete resource ID '%s': the object cannot be found; deleting from Terraform state", d.Id())
			d.SetId("")

			return nil
		}

		// Assertion of object type and error handling
		var dtcMonitor *ibclient.DtcMonitor
		recJson, _ := json.Marshal(rec)
		err = json.Unmarshal(recJson, &dtcMonitor)

		_, err = objMgr.DeleteDtcMonitor(dtcMonitor.Ref)
		if err != nil {
			return fmt.Errorf("error while releasing the resource with ID '%s': %s", d.Id(), err.Error())
		}
		d.SetId("")

		return nil
	}
}
