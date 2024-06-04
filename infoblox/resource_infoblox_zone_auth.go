package infoblox

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/infobloxopen/infoblox-go-client/v2/utils"

	ibclient "github.com/infobloxopen/infoblox-go-client/v2"
)

func resourceZoneAuth() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceZoneAuthCreate,
		ReadContext:   resourceZoneAuthRead,
		UpdateContext: resourceZoneAuthUpdate,
		DeleteContext: resourceZoneAuthDelete,
		Importer: &schema.ResourceImporter{
			State: resourceZoneAuthImport,
		},
		CustomizeDiff: func(context context.Context, d *schema.ResourceDiff, meta interface{}) error {
			if internalID := d.Get("internal_id"); internalID == "" || internalID == nil {
				err := d.SetNewComputed("internal_id")
				if err != nil {
					return err
				}
			}
			return nil
		},

		Schema: map[string]*schema.Schema{
			"fqdn": {
				Type:     schema.TypeString,
				Required: true,
				Description: "The name of this DNS zone. For a reverse zone, this is in 'address/cidr' " +
					"format. For other zones, this is in FQDN format. This value can be in " +
					"unicode format. Note that for a reverse zone, the corresponding zone_format " +
					"value should be set.",
			},

			"view": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "The name of the DNS view in which the zone resides. Example: 'external'",
			},

			"zone_format": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "Determines the format of this zone. Valid values are: FORWARD, IPV4, IPV6; the default is FORWARD",
			},

			"comment": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Comment for the zone; maximum 256 characters.",
			},

			"ext_attrs": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The Extensible attributes of the zone, as a map in JSON format",
			},

			"ns_group": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The name server group that serves DNS for this zone.",
			},

			"restart_if_needed": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Restarts the member service.",
			},

			"soa_default_ttl": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  28800,
				Description: "The Time to Live (TTL) value of the SOA record of this zone. This value is " +
					"the number of seconds that data is cached.",
			},

			"soa_expire": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  2419200,
				Description: "This setting defines the amount of time, in seconds, after which the " +
					"secondary server stops giving out answers about the zone because the zone " +
					"data is too old to be useful. The default is one week.",
			},

			"soa_negative_ttl": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  900,
				Description: "The negative Time to Live (TTL) value of the SOA of the zone indicates how " +
					"long a secondary server can cache data for 'Does Not Respond' responses.",
			},

			"soa_refresh": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  10800,
				Description: "This indicates the interval at which a secondary server sends a message to " +
					"the primary server for a zone to check that its data is current, and " +
					"retrieve fresh data if it is not.",
			},

			"soa_retry": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  3600,
				Description: "This indicates how long a secondary server must wait before attempting to " +
					"recontact the primary server after a connection failure between the two " +
					"servers occurs.",
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
		},
	}
}

func checkZoneFormat(f string) diag.Diagnostics {
	for _, v := range []string{"FORWARD", "IPV4", "IPV6"} {
		if f == v {
			return nil
		}
	}
	return diag.FromErr(fmt.Errorf("zone's format must be either of FORWARD, IPV4, IPV6"))
}

func formZone(
	create bool, d *schema.ResourceData, m interface{}) (
	*ibclient.ZoneAuth, diag.Diagnostics) {

	extAttrJSON := d.Get("ext_attrs").(string)
	extAttrs, err := terraformDeserializeEAs(extAttrJSON)
	if err != nil {
		return nil, diag.FromErr(err)
	}

	zone := &ibclient.ZoneAuth{
		Ea: extAttrs,
	}

	if create {
		zone.Fqdn = d.Get("fqdn").(string)

		zone.View = utils.StringPtr(d.Get("view").(string))
		if *zone.View == "" {
			zone.View = utils.StringPtr("default")
		}

		zf := d.Get("zone_format").(string)
		if errs := checkZoneFormat(zf); errs != nil {
			if zf == "" {
				zf = "FORWARD"
			} else {
				return nil, errs
			}
		}
		zone.ZoneFormat = zf
	}

	if d.HasChange("comment") {
		zone.Comment = utils.StringPtr(d.Get("comment").(string))
	}

	nsGrp := d.Get("ns_group").(string)
	if nsGrp != "" {
		zone.NsGroup = utils.StringPtr(nsGrp)
	} else {
		zone.NsGroup = nil
	}

	if d.HasChange("restart_if_needed") {
		zone.RestartIfNeeded = utils.BoolPtr(d.Get("restart_if_needed").(bool))
	}

	if d.HasChange("soa_default_ttl") {
		zone.SoaDefaultTtl = utils.Uint32Ptr(uint32(d.Get("soa_default_ttl").(int)))
	}

	if d.HasChange("soa_expire") {
		zone.SoaExpire = utils.Uint32Ptr(uint32(d.Get("soa_expire").(int)))
	}

	if d.HasChange("soa_negative_ttl") {
		zone.SoaNegativeTtl = utils.Uint32Ptr(uint32(d.Get("soa_negative_ttl").(int)))
	}

	if d.HasChange("soa_refresh") {
		zone.SoaRefresh = utils.Uint32Ptr(uint32(d.Get("soa_refresh").(int)))
	}

	if d.HasChange("soa_retry") {
		zone.SoaRetry = utils.Uint32Ptr(uint32(d.Get("soa_retry").(int)))
	}

	return zone, nil
}

func resourceZoneAuthCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	if intId := d.Get("internal_id"); intId.(string) != "" {
		return diag.FromErr(fmt.Errorf("the value of 'internal_id' field must not be set manually"))
	}

	zone, errs := formZone(true, d, m)
	if errs != nil {
		return errs
	}

	// Generate internal ID and add it to the extensible attributes
	internalId := generateInternalId()
	zone.Ea[eaNameForInternalId] = internalId.String()

	connector := m.(ibclient.IBConnector)
	zoneRef, err := connector.CreateObject(zone)
	if err != nil {
		return diag.FromErr(fmt.Errorf("failed to create a zone: %w", err))
	}

	if err = d.Set("ref", zoneRef); err != nil {
		return diag.FromErr(err)
	}

	if err = d.Set("internal_id", internalId.String()); err != nil {
		return diag.FromErr(err)
	}

	d.SetId(zoneRef)

	return resourceZoneAuthRead(ctx, d, m)
}

func resourceZoneAuthRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	extAttrJSON := d.Get("ext_attrs").(string)
	extAttrs, err := terraformDeserializeEAs(extAttrJSON)
	if err != nil {
		return diag.FromErr(err)
	}

	var diags diag.Diagnostics

	rec, err := searchObjectByRefOrInternalId("ZoneAuth", d, m)
	if err != nil {
		if _, ok := err.(*ibclient.NotFoundError); !ok {
			return diag.FromErr(ibclient.NewNotFoundError(fmt.Sprintf(
				"cannot find appropriate object on NIOS side for resource with ID '%s': %s;", d.Id(), err)))
		} else {
			d.SetId("")
			return nil
		}
	}

	// Assertion of object type and error handling
	var zoneResult *ibclient.ZoneAuth
	recJson, _ := json.Marshal(rec)
	err = json.Unmarshal(recJson, &zoneResult)

	if err != nil && zoneResult.Ref != "" {
		return diag.FromErr(fmt.Errorf("getting DNS View with ID: %s failed: %w", d.Id(), err))
	}

	err = d.Set("fqdn", zoneResult.Fqdn)
	if err != nil {
		return diag.FromErr(err)
	}

	if zoneResult.Comment != nil {
		err = d.Set("comment", *zoneResult.Comment)
		if err != nil {
			return diag.FromErr(err)
		}
	}

	if zoneResult.NsGroup != nil {
		err = d.Set("ns_group", *zoneResult.NsGroup)
	} else {
		err = d.Set("ns_group", "")
	}
	if err != nil {
		return diag.FromErr(err)
	}

	if zoneResult.SoaDefaultTtl != nil {
		err = d.Set("soa_default_ttl", *zoneResult.SoaDefaultTtl)
		if err != nil {
			return diag.FromErr(err)
		}
	}

	if zoneResult.SoaExpire != nil {
		err = d.Set("soa_expire", *zoneResult.SoaExpire)
		if err != nil {
			return diag.FromErr(err)
		}
	}

	if zoneResult.SoaNegativeTtl != nil {
		err = d.Set("soa_negative_ttl", *zoneResult.SoaNegativeTtl)
		if err != nil {
			return diag.FromErr(err)
		}
	}

	if zoneResult.SoaRefresh != nil {
		err = d.Set("soa_refresh", *zoneResult.SoaRefresh)
		if err != nil {
			return diag.FromErr(err)
		}
	}

	if zoneResult.SoaRetry != nil {
		err = d.Set("soa_retry", *zoneResult.SoaRetry)
		if err != nil {
			return diag.FromErr(err)
		}
	}

	if zoneResult.View != nil {
		err = d.Set("view", *zoneResult.View)
		if err != nil {
			return diag.FromErr(err)
		}
	}

	err = d.Set("zone_format", zoneResult.ZoneFormat)
	if err != nil {
		return diag.FromErr(err)
	}

	delete(zoneResult.Ea, eaNameForInternalId)

	omittedEAs := omitEAs(zoneResult.Ea, extAttrs)

	if omittedEAs != nil && len(omittedEAs) > 0 {
		eaJSON, err := terraformSerializeEAs(omittedEAs)
		if err != nil {
			return diag.FromErr(err)
		}
		if err = d.Set("ext_attrs", eaJSON); err != nil {
			return diag.FromErr(err)
		}
	}

	d.SetId(zoneResult.Ref)
	if err = d.Set("ref", zoneResult.Ref); err != nil {
		return diag.FromErr(err)
	}

	return diags
}

func resourceZoneAuthUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) (diags diag.Diagnostics) {
	if d.HasChange("internal_id") {
		return diag.FromErr(fmt.Errorf("changing the value of 'internal_id' field is not allowed"))
	}

	if d.HasChange("fqdn") {
		return diag.FromErr(fmt.Errorf("field is not allowed for update: fqdn"))
	}

	if d.HasChange("zone_format") {
		return diag.FromErr(fmt.Errorf("field is not allowed for update: zone_format"))
	}

	if d.HasChange("view") {
		return diag.FromErr(fmt.Errorf("field is not allowed for update: view"))
	}

	zone, errs := formZone(false, d, m)
	if errs != nil {
		return errs
	}

	oldExtAttrsJSON, newExtAttrsJSON := d.GetChange("ext_attrs")

	newExtAttrs, err := terraformDeserializeEAs(newExtAttrsJSON.(string))
	if err != nil {
		return diag.FromErr(err)
	}

	oldExtAttrs, err := terraformDeserializeEAs(oldExtAttrsJSON.(string))
	if err != nil {
		return diag.FromErr(err)
	}

	connector := m.(ibclient.IBConnector)

	rec, err := searchObjectByRefOrInternalId("ZoneAuth", d, m)
	if err != nil {
		if _, ok := err.(*ibclient.NotFoundError); !ok {
			return diag.FromErr(ibclient.NewNotFoundError(fmt.Sprintf(
				"cannot find appropriate object on NIOS side for resource with ID '%s': %s;", d.Id(), err)))
		} else {
			d.SetId("")
			return nil
		}
	}

	// Assertion of object type and error handling
	var zoneVal *ibclient.ZoneAuth
	recJson, _ := json.Marshal(rec)
	err = json.Unmarshal(recJson, &zoneVal)

	if err != nil {
		return diag.FromErr(fmt.Errorf("failed to read Zone Auth for update operation: %w", err))
	}

	internalId := d.Get("internal_id").(string)

	if internalId == "" {
		internalId = generateInternalId().String()
	}

	newInternalId := newInternalResourceIdFromString(internalId)
	newExtAttrs[eaNameForInternalId] = newInternalId.String()

	zone.Ea, err = mergeEAs(zoneVal.Ea, newExtAttrs, oldExtAttrs, connector)
	if err != nil {
		return diag.FromErr(err)
	}

	zoneRef, err := connector.UpdateObject(zone, d.Id())
	if err != nil {
		return diag.FromErr(fmt.Errorf("failed to update a zone: %w", err))
	}

	d.SetId(zoneRef)

	if err = d.Set("ref", zoneRef); err != nil {
		return diag.FromErr(err)
	}
	if err = d.Set("internal_id", newInternalId.String()); err != nil {
		return diag.FromErr(err)
	}
	return resourceZoneAuthRead(ctx, d, m)
}

func resourceZoneAuthDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	connector := m.(ibclient.IBConnector)

	rec, err := searchObjectByRefOrInternalId("ZoneAuth", d, m)
	if err != nil {
		if _, ok := err.(*ibclient.NotFoundError); !ok {
			return diag.FromErr(ibclient.NewNotFoundError(fmt.Sprintf(
				"cannot find appropriate object on NIOS side for resource with ID '%s': %s;", d.Id(), err)))
		} else {
			d.SetId("")
			return nil
		}
	}

	// Assertion of object type and error handling
	var zoneResult *ibclient.ZoneAuth
	recJson, _ := json.Marshal(rec)
	err = json.Unmarshal(recJson, &zoneResult)

	if err != nil && zoneResult.Ref != "" {
		return diag.FromErr(fmt.Errorf("getting zone with ID: %s failed: %w", d.Id(), err))
	}

	if _, err := connector.DeleteObject(zoneResult.Ref); err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")
	return nil
}

func resourceZoneAuthImport(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	extAttrJSON := d.Get("ext_attrs").(string)
	_, err := terraformDeserializeEAs(extAttrJSON)
	if err != nil {
		return nil, err
	}

	connector := m.(ibclient.IBConnector)

	zoneRef := d.Id()

	zoneResult := ibclient.ZoneAuth{}

	zone := &ibclient.ZoneAuth{}
	zone.SetReturnFields(append(
		zoneResult.ReturnFields(),
		"comment",
		"ns_group",
		"soa_default_ttl",
		"soa_expire",
		"soa_negative_ttl",
		"soa_refresh",
		"soa_retry",
		"view",
		"zone_format",
		"extattrs",
	))

	err = connector.GetObject(zone, zoneRef, nil, &zoneResult)
	if err != nil {
		return nil, fmt.Errorf("failed to read zone: %w", err)
	}

	err = d.Set("fqdn", zoneResult.Fqdn)
	if err != nil {
		return nil, err
	}

	if zoneResult.Comment != nil {
		err = d.Set("comment", *zoneResult.Comment)
		if err != nil {
			return nil, err
		}
	}

	if zoneResult.NsGroup != nil {
		err = d.Set("ns_group", *zoneResult.NsGroup)
		if err != nil {
			return nil, err
		}
	}

	if zoneResult.SoaDefaultTtl != nil {
		err = d.Set("soa_default_ttl", *zoneResult.SoaDefaultTtl)
		if err != nil {
			return nil, err
		}
	}

	if zoneResult.SoaExpire != nil {
		err = d.Set("soa_expire", *zoneResult.SoaExpire)
		if err != nil {
			return nil, err
		}
	}

	if zoneResult.SoaNegativeTtl != nil {
		err = d.Set("soa_negative_ttl", *zoneResult.SoaNegativeTtl)
		if err != nil {
			return nil, err
		}
	}

	if zoneResult.SoaRefresh != nil {
		err = d.Set("soa_refresh", *zoneResult.SoaRefresh)
		if err != nil {
			return nil, err
		}
	}

	if zoneResult.SoaRetry != nil {
		err = d.Set("soa_retry", *zoneResult.SoaRetry)
		if err != nil {
			return nil, err
		}
	}

	if zoneResult.View != nil {
		err = d.Set("view", *zoneResult.View)
		if err != nil {
			return nil, err
		}
	}

	err = d.Set("zone_format", zoneResult.ZoneFormat)
	if err != nil {
		return nil, err
	}

	if zoneResult.Ea != nil && len(zoneResult.Ea) > 0 {
		eaJSON, err := terraformSerializeEAs(zoneResult.Ea)
		if err != nil {
			return nil, err
		}
		if err = d.Set("ext_attrs", eaJSON); err != nil {
			return nil, err
		}
	}

	d.SetId(zoneResult.Ref)

	var dErr diag.Diagnostics
	var ctx context.Context
	dErr = resourceZoneAuthUpdate(ctx, d, m)
	if dErr != nil {
		return nil, fmt.Errorf("failed to import Zone Auth: %v", dErr)
	}

	return []*schema.ResourceData{d}, nil
}
