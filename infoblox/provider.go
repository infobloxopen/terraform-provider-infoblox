package infoblox

import (
	"context"
	"fmt"
	"math"
	"strings"
	"time"

	"github.com/google/uuid"
	log "github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	ibclient "github.com/infobloxopen/infoblox-go-client/v2"
)

// Common parameters
const (
	ttlUndef            = math.MinInt32
	eaNameForInternalId = "Terraform Internal ID"
	eaNameForTenantId   = "Tenant ID"
	altIdSeparator      = "|"
)

// Internal ID represents an immutable ID during resource's lifecycle.
// NIOS object's reference may get changed, sometimes this is a problem:
//   when more than one TF resources have the same NIOS WAPI object as a backend,
//   changing reference to the object invalidates the old reference,
//   which needs to be changed for all appropriate TF resources.
//   Doing this is problematic.
//   An example of such resources: a pair of infoblox_ipvX_allocation/infoblox_ipvX_association.
//   They both must relate to a single host record on NIOS side.
// Important requirement: the text representing an internal ID must not contain '|' sign,
//   or in general: the sign (or a sequence of) which is defined by altIdSeparator constant.
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
	uuid, err := uuid.NewRandom()
	if err != nil {
		panic(err)
	}
	return &internalResourceId{value: uuid}
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
			"infoblox_zone_delegated":         resourceZoneDelegated(),
		},
		DataSourcesMap: map[string]*schema.Resource{
			"infoblox_ipv4_network":           dataSourceIPv4Network(),
			"infoblox_ipv4_network_container": dataSourceIpv4NetworkContainer(),
			"infoblox_network_view":           dataSourceNetworkView(),
			"infoblox_a_record":               dataSourceARecord(),
			"infoblox_aaaa_record":            dataSourceAAAARecord(),
			"infoblox_cname_record":           dataSourceCNameRecord(),
			"infoblox_ptr_record":             dataSourcePtrRecord(),
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
	return conn, nil
}
