package infoblox

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"sort"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	ibclient "github.com/infobloxopen/infoblox-go-client/v2"
)

func resourceZoneDelegated() *schema.Resource {
	return &schema.Resource{
		Create: resourceZoneDelegatedCreate,
		Read:   resourceZoneDelegatedRead,
		Update: resourceZoneDelegatedUpdate,
		Delete: resourceZoneDelegatedDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"fqdn": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The FQDN of the delegated zone.",
			},
			"delegate_to": {
				Type:     schema.TypeSet,
				Required: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"address": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "IP of Name Server",
						},
						"name": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "FQDN of Name Server",
						},
					},
				},
			},
			"ext_attrs": {
				Type:        schema.TypeString,
				Default:     "",
				Optional:    true,
				Description: "Extensible attributes, as a map in JSON format",
			},
		},
	}
}

func resourceNameServer() *schema.Schema {
	return &schema.Schema{
		Type:     schema.TypeSet,
		Required: true,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"address": {
					Type:        schema.TypeString,
					Computed:    true,
					Description: "IP of Name Server",
				},
				"name": {
					Type:        schema.TypeString,
					Required:    true,
					Description: "FQDN of Name Server",
				},
			},
		},
	}
}

func computeDelegations(delegations []interface{}) ([]ibclient.NameServer, []map[string]interface{}, error) {
	var nameServers []ibclient.NameServer
	computedDelegations := make([]map[string]interface{}, 0)
	for _, delegation := range delegations {
		var ns ibclient.NameServer
		delegationMap := delegation.(map[string]interface{})
		ns.Name = delegationMap["name"].(string)
		lookupHosts, err := net.LookupHost(delegationMap["name"].(string))
		if err != nil {
			return nil, nil, fmt.Errorf("Failed to resolve delegate_to: %w", err)
		}
		sort.Strings(lookupHosts)
		ns.Address = lookupHosts[0]
		delegationMap["address"] = ns.Address
		nameServers = append(nameServers, ns)
		computedDelegations = append(computedDelegations, delegationMap)
	}
	return nameServers, computedDelegations, nil
}

func resourceZoneDelegatedCreate(d *schema.ResourceData, m interface{}) error {
	log.Printf("[DEBUG] %s: Beginning to create Zone Delegated", resourceZoneDelegatedIDString(d))

	extAttrJSON := d.Get("ext_attrs").(string)
	extAttrs := make(map[string]interface{})
	if extAttrJSON != "" {
		if err := json.Unmarshal([]byte(extAttrJSON), &extAttrs); err != nil {
			return fmt.Errorf("cannot process 'ext_attrs' field: %w", err)
		}
	}

	delegatedFQDN := d.Get("fqdn").(string)

	delegations := d.Get("delegate_to").(*schema.Set).List()

	nameServers, computedDelegations, err := computeDelegations(delegations)
	if err != nil {
		return err
	}

	var tenantID string
	if tempVal, ok := extAttrs["Tenant ID"]; ok {
		tenantID = tempVal.(string)
	}

	connector := m.(*ibclient.Connector)

	objMgr := ibclient.NewObjectManager(connector, "Terraform", tenantID)

	zoneDelegated, err := objMgr.CreateZoneDelegated(
		delegatedFQDN,
		nameServers)
	if err != nil {
		return fmt.Errorf("Error creating Zone Delegated: %w", err)
	}

	d.Set("delegate_to", computedDelegations)

	d.SetId(zoneDelegated.Ref)

	log.Printf("[DEBUG] %s: Creation of Zone Delegated complete", resourceZoneDelegatedIDString(d))
	return nil
}

func resourceZoneDelegatedRead(d *schema.ResourceData, m interface{}) error {
	log.Printf("[DEBUG] %s: Begining to Get Zone Delegated", resourceZoneDelegatedIDString(d))

	extAttrJSON := d.Get("ext_attrs").(string)
	extAttrs := make(map[string]interface{})
	if extAttrJSON != "" {
		if err := json.Unmarshal([]byte(extAttrJSON), &extAttrs); err != nil {
			return fmt.Errorf("cannot process 'ext_attrs' field: %w", err)
		}
	}

	var tenantID string
	if tempVal, ok := extAttrs["Tenant ID"]; ok {
		tenantID = tempVal.(string)
	}

	connector := m.(*ibclient.Connector)

	objMgr := ibclient.NewObjectManager(connector, "Terraform", tenantID)

	// first attempt to read by ref, otherwise assume import and support fqdn
	zoneDelegatedObj, err := objMgr.GetZoneDelegated(d.Id())
	if err != nil {
		return fmt.Errorf("Getting Zone Delegated failed: %w", err)
	}

	var delegations []map[string]interface{}
	for _, delegation := range zoneDelegatedObj.DelegateTo {
		ns := make(map[string]interface{})
		ns["address"] = delegation.Address
		ns["name"] = delegation.Name
		delegations = append(delegations, ns)
	}

	d.Set("fqdn", zoneDelegatedObj.Fqdn)
	d.Set("delegate_to", delegations)

	d.SetId(zoneDelegatedObj.Ref)

	log.Printf("[DEBUG] %s: Completed reading Zone Delegated ", resourceZoneDelegatedIDString(d))
	return nil
}

func resourceZoneDelegatedUpdate(d *schema.ResourceData, m interface{}) error {
	log.Printf("[DEBUG] %s: Beginning to update Zone Delegated", resourceZoneDelegatedIDString(d))

	extAttrJSON := d.Get("ext_attrs").(string)
	extAttrs := make(map[string]interface{})
	if extAttrJSON != "" {
		if err := json.Unmarshal([]byte(extAttrJSON), &extAttrs); err != nil {
			return fmt.Errorf("cannot process 'ext_attrs' field: %w", err)
		}
	}

	delegations := d.Get("delegate_to").(*schema.Set).List()

	nameServers, computedDelegations, err := computeDelegations(delegations)
	if err != nil {
		return err
	}

	var tenantID string
	if tempVal, ok := extAttrs["Tenant ID"]; ok {
		tenantID = tempVal.(string)
	}

	connector := m.(*ibclient.Connector)

	objMgr := ibclient.NewObjectManager(connector, "Terraform", tenantID)

	zoneDelegatedUpdated, err := objMgr.UpdateZoneDelegated(d.Id(), nameServers)
	if err != nil {
		return fmt.Errorf("Updating of Zone Delegated failed : %w", err)
	}

	d.Set("delegate_to", computedDelegations)

	d.SetId(zoneDelegatedUpdated.Ref)
	return nil
}

func resourceZoneDelegatedDelete(d *schema.ResourceData, m interface{}) error {
	log.Printf("[DEBUG] %s: Beginning Deletion of Zone Delegated", resourceZoneDelegatedIDString(d))

	extAttrJSON := d.Get("ext_attrs").(string)
	extAttrs := make(map[string]interface{})
	if extAttrJSON != "" {
		if err := json.Unmarshal([]byte(extAttrJSON), &extAttrs); err != nil {
			return fmt.Errorf("cannot process 'ext_attrs' field: %w", err)
		}
	}

	var tenantID string
	if tempVal, ok := extAttrs["Tenant ID"]; ok {
		tenantID = tempVal.(string)
	}

	connector := m.(*ibclient.Connector)

	objMgr := ibclient.NewObjectManager(connector, "Terraform", tenantID)

	_, err := objMgr.DeleteZoneDelegated(d.Id())
	if err != nil {
		return fmt.Errorf("Deletion of Zone Delegated failed : %w", err)
	}
	d.SetId("")

	log.Printf("[DEBUG] %s: Deletion of Zone Delegated complete", resourceZoneDelegatedIDString(d))
	return nil
}

type resourceZoneDelegatedIDStringInterface interface {
	Id() string
}

func resourceZoneDelegatedIDString(d resourceZoneDelegatedIDStringInterface) string {
	id := d.Id()
	if id == "" {
		id = "<new resource>"
	}
	return fmt.Sprintf("infoblox_zone_delegated (ID = %s)", id)
}
