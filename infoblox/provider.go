package infoblox

import (
	"context"
	"fmt"
	"math"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	ibclient "github.com/infobloxopen/infoblox-go-client/v2"
)

// Common parameters
const (
	ttlUndef            = math.MinInt32
	eaNameForInternalId = "Terraform Internal ID"
	altIdSeparator      = "|"
)

type internalResourceId struct {
	value uuid.UUID
}

func (id *internalResourceId) Equal(id2 *internalResourceId) bool {
	if id2 == nil {
		panic("the argument must not be nil")
	}
	return id.value.String() == id2.value.String()
}

func (id *internalResourceId) String() string {
	return id.value.String()
}

// Returns a pointer to parsed internal resource ID, nil otherwise.
func newInternalResourceIdFromString(id string) *internalResourceId {
	newUUID, err := uuid.Parse(id)
	if err != nil {
		return nil
	}

	return &internalResourceId{value: newUUID}
}

func generateInternalId() *internalResourceId {
	uuid, err := uuid.NewRandom()
	if err != nil {
		panic(err)
	}
	return &internalResourceId{value: uuid}
}

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

// valid = true:
//   - exactly 2 parts found
//   - ... and separated by the delimiter
//   - ... and the 1st one is a valid internal ID
//   - ... and the 2nd one is not empty
func getAltIdFields(altId string) (internalId *internalResourceId, ref string, valid bool) {
	idParts := strings.SplitN(altId, altIdSeparator, 2)
	switch len(idParts) {
	case 1:
		internalId = newInternalResourceIdFromString(idParts[0])
	case 2:
		internalId = newInternalResourceIdFromString(idParts[0])
		ref = idParts[1]
		valid = internalId != nil && ref != ""
	}

	return
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
				DefaultFunc: schema.EnvDefaultFunc("WAPI_VERSION", "2.7"),
				Description: "WAPI Version of Infoblox server defaults to v2.7.",
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
			"infoblox_ipv4_allocation":        resourceIPv4Allocation(),
			"infoblox_ipv6_allocation":        resourceIPv6Allocation(),
			"infoblox_ip_allocation":          resourceIPAllocation(),
			"infoblox_ipv4_association":       resourceIPv4AssociationInit(),
			"infoblox_ipv6_association":       resourceIPv6AssociationInit(),
			"infoblox_ip_association":         resourceIpAssociationInit(),
			"infoblox_a_record":               resourceARecord(),
			"infoblox_aaaa_record":            resourceAAAARecord(),
			"infoblox_cname_record":           resourceCNAMERecord(),
			"infoblox_ptr_record":             resourcePTRRecord(),
		},
		DataSourcesMap: map[string]*schema.Resource{
			"infoblox_ipv4_network": dataSourceIPv4Network(),
			"infoblox_a_record":     dataSourceARecord(),
			"infoblox_cname_record": dataSourceCNameRecord(),
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

	conn, err := ibclient.NewConnector(hostConfig, authConfig, transportConfig, requestBuilder, requestor)
	if err != nil {
		return nil, diag.Diagnostics{diag.Diagnostic{Summary: err.Error()}}
	}
	return conn, nil
}

func stateImporter(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	return []*schema.ResourceData{d}, nil
}
