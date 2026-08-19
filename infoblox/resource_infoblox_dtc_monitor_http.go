package infoblox

import (
	"encoding/json"
	"fmt"
	"maps"
	"regexp"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"

	ibclient "github.com/infobloxopen/infoblox-go-client/v2"
)

type dtcMonitorHttpStateBridge struct {
	DtcMonitorBase
	ciphers               string
	client_cert           string
	content_check         string
	content_check_input   string
	content_check_op      string
	content_check_regex   string
	content_extract_group uint32
	content_extract_type  string
	content_extract_value string
	enable_sni            bool
	port                  uint32
	request               string
	result                string
	result_code           uint32
	secure                bool
	validate_cert         bool
}

func getDtcMonitorHttpFromTfState(d *schema.ResourceData) (dtcMonitorHttpStateBridge, error) {
	b := dtcMonitorHttpStateBridge{}
	mb, err := resourceDtcMonitorGetFromTfState(d)
	if err != nil {
		return b, err
	}

	b.DtcMonitorBase = mb
	b.ciphers = d.Get("ciphers").(string)
	b.client_cert = d.Get("client_cert").(string)
	b.content_check = d.Get("content_check").(string)
	b.content_check_input = d.Get("content_check_input").(string)
	b.content_check_op = d.Get("content_check_op").(string)
	b.content_check_regex = d.Get("content_check_regex").(string)
	b.content_extract_group = uint32(d.Get("content_extract_group").(int))
	b.content_extract_type = d.Get("content_extract_type").(string)
	b.content_extract_value = d.Get("content_extract_value").(string)
	b.enable_sni = d.Get("enable_sni").(bool)
	b.port = uint32(d.Get("port").(int))
	b.request = d.Get("request").(string)
	b.result = d.Get("result").(string)
	b.result_code = uint32(d.Get("result_code").(int))
	b.secure = d.Get("secure").(bool)
	b.validate_cert = d.Get("validate_cert").(bool)
	return b, nil
}

// read API response [b] and set state in [d] accordingly
func resourceDtcMonitorHttpSetTfState(d *schema.ResourceData, b *ibclient.DtcMonitorHttp) error {

	// set state with dtc:monitor (generic base class) attributes
	if err := resourceDtcMonitorSetTfState(d, b); err != nil {
		return err
	}

	// set state with attributes specific to dtc:monitor:http

	var err error
	if err = d.Set("ciphers", b.Ciphers); err != nil {
		return err
	}
	if err = d.Set("client_cert", b.ClientCert); err != nil {
		return err
	}
	if err = d.Set("content_check", b.ContentCheck); err != nil {
		return err
	}
	if err = d.Set("content_check_input", b.ContentCheckInput); err != nil {
		return err
	}
	if err = d.Set("content_check_op", b.ContentCheckOp); err != nil {
		return err
	}
	if err = d.Set("content_check_regex", b.ContentCheckRegex); err != nil {
		return err
	}
	if err = d.Set("content_extract_group", b.ContentExtractGroup); err != nil {
		return err
	}
	if err = d.Set("content_extract_type", b.ContentExtractType); err != nil {
		return err
	}
	if err = d.Set("content_extract_value", b.ContentExtractValue); err != nil {
		return err
	}
	if err = d.Set("enable_sni", b.EnableSni); err != nil {
		return err
	}
	if err = d.Set("port", b.Port); err != nil {
		return err
	}
	if err = d.Set("request", b.Request); err != nil {
		return err
	}
	if err = d.Set("result", b.Result); err != nil {
		return err
	}
	if err = d.Set("result_code", b.ResultCode); err != nil {
		return err
	}
	if err = d.Set("secure", b.Secure); err != nil {
		return err
	}
	if err = d.Set("validate_cert", b.ValidateCert); err != nil {
		return err
	}
	return nil
}

func resourceDtcMonitorHttp() *schema.Resource {
	var (
		monitorSchema = resourceDtcMonitorSchema()
		httpSchema    = map[string]*schema.Schema{
			"ciphers": {
				Type:        schema.TypeString,
				Description: "An optional cipher list for a secure HTTP/S connection.",
				Optional:    true,
				Default:     "",
			},
			"client_cert": {
				Type:        schema.TypeString,
				Description: "An optional client certificate, supplied in a secure HTTP/S mode if present. Must be a dtc:certificate object reference",
				Optional:    true,
				Default:     "",
			},
			"content_check": {
				Type: schema.TypeString,
				Description: `The content check type.
				Valid values are: EXTRACT, MATCH, NONE.
				`,
				Optional: true,
				Default:  "NONE",
			},
			"content_check_input": {
				Type:        schema.TypeString,
				Description: `A portion of response to use as input for content check. Valid values are: ALL, BODY, HEADERS`,
				Default:     "ALL",
				Optional:    true,
			},
			"content_check_op": {
				Type:        schema.TypeString,
				Description: "A content check success criteria operator. Valid values are: EQ, GEQ, LEQ, NEQ",
				Optional:    true,
				Default:     "",
			},
			"content_check_regex": {
				Type:        schema.TypeString,
				Description: "A content check regular expression. Values with leading or trailing white space are not valid for this field.",
				Optional:    true,
				Default:     "",
			},
			"content_extract_group": {
				Type:        schema.TypeInt,
				Description: "A content extraction sub-expression to extract.",
				Optional:    true,
				Default:     0,
			},
			"content_extract_type": {
				Type:        schema.TypeString,
				Description: "A content extraction expected type for the extracted data. Valid values are: INTEGER, STRING",
				Optional:    true,
				Default:     "STRING",
			},
			"content_extract_value": {
				Type:        schema.TypeString,
				Description: "A content extraction value to compare with extracted result. Values with leading or trailing white space are not valid for this field.",
				Optional:    true,
				Default:     "",
			},
			"enable_sni": {
				Type:        schema.TypeBool,
				Description: "Determines whether the Server Name Indication (SNI) for HTTPS monitor is enabled.",
				Optional:    true,
				Default:     false,
			},
			"port": {
				Type:        schema.TypeInt,
				Description: "The health monitor port value.",
				Optional:    true,
				Default:     80,
			},
			"request": {
				Type:        schema.TypeString,
				Description: "An HTTP request to send.",
				Optional:    true,
				Default:     "GET /\n\n",
				ValidateFunc: validation.StringMatch(regexp.MustCompile(`\n\n$`),
					"Suffix \\n\\n required (WAPI always adds it)."),
			},
			"result": {
				Type:        schema.TypeString,
				Description: "The type of an expected result. Valid values are: ANY, CODE_IS, CODE_IS_NOT",
				Optional:    true,
				Default:     "ANY",
			},
			"result_code": {
				Type:        schema.TypeInt,
				Description: "The expected return code.",
				Optional:    true,
				Default:     200,
			},
			"secure": {
				Type:        schema.TypeBool,
				Description: "Use HTTPS, not default HTTP",
				Optional:    true,
				Default:     false,
			},
			"validate_cert": {
				Type:        schema.TypeBool,
				Description: "Determines whether the validation of the remote server’s certificate is enabled.",
				Optional:    true,
				Default:     true,
			},
		}
	)
	maps.Copy(monitorSchema, httpSchema) // keys in monitorSchema are overwritten

	return &schema.Resource{
		// TODO: move towards context-aware equivalents of these fields, as these are deprecated.
		Create: resourceDtcMonitorHttpCreate,
		Read:   resourceDtcMonitorHttpGet,
		Update: resourceDtcMonitorHttpUpdate,
		Delete: makeResourceDtcMonitorDelete("DtcMonitorHttp"),

		Importer: &schema.ResourceImporter{
			State: resourceDtcMonitorHttpImporter,
		},

		Schema: monitorSchema,
	}
}

func resourceDtcMonitorHttpCreate(d *schema.ResourceData, m interface{}) error {
	if intId := d.Get("internal_id"); intId.(string) != "" {
		return fmt.Errorf("the value of 'internal_id' field must not be set manually")
	}
	mb, err := getDtcMonitorHttpFromTfState(d)
	if err != nil {
		return fmt.Errorf("failed to convert dtc:monitor:http from TF state to API object: %w", err)
	}

	connector := m.(ibclient.IBConnector)
	objMgr := ibclient.NewObjectManager(connector, "Terraform", mb.tenantID)
	createdObject, err := objMgr.CreateDtcMonitorHttp(
		mb.comment,
		mb.name,
		mb.extAttrs,
		mb.port,
		mb.interval,
		mb.retry_down,
		mb.retry_up,
		mb.timeout,
		mb.ciphers,
		mb.client_cert,
		mb.content_check,
		mb.content_check_input,
		mb.content_check_op,
		mb.content_check_regex,
		mb.content_extract_group,
		mb.content_extract_type,
		mb.content_extract_value,
		mb.enable_sni,
		mb.request,
		mb.result,
		mb.result_code,
		mb.secure,
		mb.validate_cert,
	)
	if err != nil {
		return fmt.Errorf("error while creating a dtc:monitor:http: %s", err.Error())
	}

	d.SetId(createdObject.Ref)
	if err = d.Set("ref", createdObject.Ref); err != nil {
		return err
	}

	// For compatibility reason. This field should be deprecated in the future.
	if err = d.Set("internal_id", mb.internalID); err != nil {
		return err
	}

	return resourceDtcMonitorHttpGet(d, m)
}

func resourceDtcMonitorHttpGet(d *schema.ResourceData, m interface{}) error {
	rec, err := searchObjectByRefOrInternalId("DtcMonitorHttp", d, m)
	if err != nil {
		if _, ok := err.(*ibclient.NotFoundError); ok {
			d.SetId("")
			return nil
		}

		return ibclient.NewNotFoundError(fmt.Sprintf(
			"cannot find appropriate object on NIOS side for resource with ID '%s': %s;", d.Id(), err))
	}

	var dtcMonitorHttp *ibclient.DtcMonitorHttp
	recJson, err := json.Marshal(rec)
	if err != nil {
		return fmt.Errorf("failed to marshal DTC MonitorHttp : %s", err.Error())
	}
	err = json.Unmarshal(recJson, &dtcMonitorHttp)
	if err != nil {
		return fmt.Errorf("failed getting dtc:monitor:http with %s", err.Error())
	}

	err = resourceDtcMonitorHttpSetTfState(d, dtcMonitorHttp)
	if err != nil {
		return fmt.Errorf("failed dtc:monitor:http data conversion: %w", err)
	}

	return nil
}

func resourceDtcMonitorHttpUpdate(d *schema.ResourceData, m interface{}) (err error) {
	var updateSuccessful bool
	defer func() {
		// Reverting the state back, in case of a failure,
		// otherwise Terraform will keep the values, which leaded to the failure,
		// in the state file.
		if !updateSuccessful {
			resourceDtcMonitorUpdateFailedResetState(d)

			prevCiphers, _ := d.GetChange("ciphers")
			prevClientCert, _ := d.GetChange("client_cert")
			prevContentCheck, _ := d.GetChange("content_check")
			prevContentCheckInput, _ := d.GetChange("content_check_input")
			prevContentCheckOp, _ := d.GetChange("content_check_op")
			prevContentCheckRegex, _ := d.GetChange("content_check_regex")
			prevContentExtractGroup, _ := d.GetChange("content_extract_group")
			prevContentExtractType, _ := d.GetChange("content_extract_type")
			prevContentExtractValue, _ := d.GetChange("content_extract_value")
			prevEnableSNI, _ := d.GetChange("enable_sni")
			prevPort, _ := d.GetChange("port")
			prevRequest, _ := d.GetChange("request")
			prevResult, _ := d.GetChange("result")
			prevResultCode, _ := d.GetChange("result_code")
			prevSecure, _ := d.GetChange("secure")
			prevValidateCert, _ := d.GetChange("validate_cert")

			_ = d.Set("ciphers", prevCiphers.(string))
			_ = d.Set("client_cert", prevClientCert.(string))
			_ = d.Set("content_check", prevContentCheck.(string))
			_ = d.Set("content_check_input", prevContentCheckInput.(string))
			_ = d.Set("content_check_op", prevContentCheckOp.(string))
			_ = d.Set("content_check_regex", prevContentCheckRegex.(string))
			_ = d.Set("content_extract_group", prevContentExtractGroup.(int))
			_ = d.Set("content_extract_type", prevContentExtractType.(string))
			_ = d.Set("content_extract_value", prevContentExtractValue.(string))
			_ = d.Set("enable_sni", prevEnableSNI.(bool))
			_ = d.Set("port", prevPort.(int))
			_ = d.Set("request", prevRequest)
			_ = d.Set("result", prevResult.(string))
			_ = d.Set("result_code", prevResultCode.(int))
			_ = d.Set("secure", prevSecure.(bool))
			_ = d.Set("validate_cert", prevValidateCert.(bool))
		}
	}()

	mb, err := getDtcMonitorHttpFromTfState(d)
	if err != nil {
		return fmt.Errorf("failed to convert dtc:monitor:http from TF state to API object: %w", err)
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

	rec, err := searchObjectByRefOrInternalId("DtcMonitorHttp", d, m)
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
	var dtcMonitor *ibclient.DtcMonitorHttp
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
	dtcMonitor, err = objMgr.UpdateDtcMonitorHttp(d.Id(),
		mb.comment,
		mb.name,
		mb.extAttrs,
		mb.port,
		mb.interval,
		mb.retry_down,
		mb.retry_up,
		mb.timeout,
		mb.ciphers,
		mb.client_cert,
		mb.content_check,
		mb.content_check_input,
		mb.content_check_op,
		mb.content_check_regex,
		mb.content_extract_group,
		mb.content_extract_type,
		mb.content_extract_value,
		mb.enable_sni,
		mb.request,
		mb.result,
		mb.result_code,
		mb.secure,
		mb.validate_cert,
	)
	if err != nil {
		return fmt.Errorf("error updating dtc:monitor:http: %w", err)
	}
	updateSuccessful = true
	d.SetId(dtcMonitor.Ref)
	if err = d.Set("ref", dtcMonitor.Ref); err != nil {
		return err
	}
	if err = d.Set("internal_id", newInternalId.String()); err != nil {
		return err
	}

	return resourceDtcMonitorHttpGet(d, m)
}

func resourceDtcMonitorHttpImporter(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
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
	obj, err := objMgr.GetDtcMonitorHttpByRef(d.Id())
	if err != nil {
		return nil, fmt.Errorf("getting dtc:monitor:http with ID: %s failed: %w", d.Id(), err)
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

	err = resourceDtcMonitorHttpGet(d, m)
	if err != nil {
		return nil, err
	}

	return []*schema.ResourceData{d}, nil
}
