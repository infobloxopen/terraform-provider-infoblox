package acctest

import (
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"

	uddiclient "github.com/infobloxopen/bloxone-go-client/client"
	uddioption "github.com/infobloxopen/bloxone-go-client/option"
	niosclient "github.com/infobloxopen/infoblox-nios-go-client/client"
	niosoption "github.com/infobloxopen/infoblox-nios-go-client/option"

	"github.com/infobloxopen/terraform-provider-infoblox/internal/provider"
)

const (
	letterBytes = "abcdefghijklmnopqrstuvwxyz"
)

var packageDir string

func init() {
	_, currentFile, _, ok := runtime.Caller(0)
	if ok {
		packageDir = filepath.Dir(currentFile)
	}
}

var (
	// NIOSClient is used for NIOS verification tests
	NIOSClient *niosclient.APIClient

	// UDDIClient is used for UDDI verification tests
	UDDIClient *uddiclient.APIClient

	// ProtoV6ProviderFactories are used to instantiate a provider during
	// acceptance testing.
	ProtoV6ProviderFactories = map[string]func() (tfprotov6.ProviderServer, error){
		"infoblox": providerserver.NewProtocol6WithError(provider.New("test", "test")()),
	}
)

// PreCheckNIOS validates NIOS environment and initializes the client.
func PreCheckNIOS(t *testing.T) {
	hostURL := os.Getenv("NIOS_HOST_URL")
	if hostURL == "" {
		t.Fatal("NIOS_HOST_URL must be set for NIOS acceptance tests")
	}

	username := os.Getenv("NIOS_USERNAME")
	if username == "" {
		t.Fatal("NIOS_USERNAME must be set for NIOS acceptance tests")
	}

	password := os.Getenv("NIOS_PASSWORD")
	if password == "" {
		t.Fatal("NIOS_PASSWORD must be set for NIOS acceptance tests")
	}

	NIOSClient = niosclient.NewAPIClient(
		niosoption.WithClientName("terraform-acceptance-tests"),
		niosoption.WithNIOSHostUrl(hostURL),
		niosoption.WithNIOSUsername(username),
		niosoption.WithNIOSPassword(password),
		niosoption.WithDebug(true),
	)
}

// PreCheckUDDI validates UDDI environment and initializes the client.
func PreCheckUDDI(t *testing.T) {
	cspURL := os.Getenv("BLOXONE_CSP_URL")
	if cspURL == "" {
		t.Fatal("BLOXONE_CSP_URL must be set for UDDI acceptance tests")
	}

	apiKey := os.Getenv("BLOXONE_API_KEY")
	if apiKey == "" {
		t.Fatal("BLOXONE_API_KEY must be set for UDDI acceptance tests")
	}

	UDDIClient = uddiclient.NewAPIClient(
		uddioption.WithClientName("terraform-acceptance-tests"),
		uddioption.WithCSPUrl(cspURL),
		uddioption.WithAPIKey(apiKey),
		uddioption.WithDebug(true),
	)
}

// PreCheck validates the test environment based on backend.
// Use PreCheckNIOS or PreCheckUDDI directly when backend is known.
func PreCheck(t *testing.T, backend string) {
	switch backend {
	case "nios":
		PreCheckNIOS(t)
	case "uddi":
		PreCheckUDDI(t)
	default:
		t.Fatalf("Unknown backend: %s", backend)
	}
}

// RandomName generates a random lowercase string.
func RandomName() string {
	b := make([]byte, 6)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

// RandomNameWithPrefix generates a random name with the given prefix.
func RandomNameWithPrefix(prefix string) string {
	return fmt.Sprintf("%s-%s", prefix, RandomName())
}

// RandomIP generates a random private IP address with valid host octet (1-254).
func RandomIP() string {
	// Use 10.x.x.x private range with valid host octet
	return fmt.Sprintf("10.%d.%d.%d", rand.Intn(256), rand.Intn(256), 1+rand.Intn(254))
}

// RandomOctet generates a random octet (0-255).
func RandomOctet() int {
	return rand.Intn(256)
}

// RandomIPWithSpecificOctetsSet generates a random IP address with the first three octets set to the given prefix.
// Last octet is 1-254 (valid host range, excludes 0=network and 255=broadcast).
func RandomIPWithSpecificOctetsSet(prefix string) string {
	return fmt.Sprintf("%s.%d", prefix, 1+rand.Intn(254))
}

// RandomNumber generates a random number up to maxLimit.
func RandomNumber(maxLimit int) int {
	if maxLimit <= 0 {
		return 0
	}
	return rand.Intn(maxLimit)
}

// RandomCIDRNetwork generates a random network with specific CIDR
func RandomCIDRNetwork() string {
	// Generate test-suitable private networks
	base := 10 + rand.Intn(246) // 10-255 for first octet
	second := rand.Intn(256)    // 0-255 for second octet
	cidr := 16 + rand.Intn(9)   // /16 to /24 (common for network containers)

	return fmt.Sprintf("%d.%d.0.0/%d", base, second, cidr)
}

// RandomIPv6Network generates a random IPv6 network with specific CIDR
func RandomIPv6Network() string {
	// Generate a random IPv6 network using the documentation prefix 2001:db8::/32
	// This is reserved for documentation and testing purposes (RFC 3849)
	third := rand.Intn(65536)  // 0-FFFF for third hextet
	fourth := rand.Intn(65536) // 0-FFFF for fourth hextet
	cidr := 64 + rand.Intn(60)

	return fmt.Sprintf("2001:db8:%x:%x::/%d", third, fourth, cidr)
}

// RandomAlphaNumeric generates a random alphanumeric string of the specified length.
func RandomAlphaNumeric(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}

// RandomMACAddress generates a random MAC address
func RandomMACAddress() string {
	return fmt.Sprintf("%02x:%02x:%02x:%02x:%02x:%02x",
		rand.Intn(256),
		rand.Intn(256),
		rand.Intn(256),
		rand.Intn(256),
		rand.Intn(256),
		rand.Intn(256))
}

// Random32Hexadecimal generates a random 32-character hexadecimal string
func Random32Hexadecimal() string {
	// Two 64-bit random values = 128 bits = 32 hex characters
	return fmt.Sprintf("%016x%016x", rand.Uint64(), rand.Uint64())
}

// TfvarsExists checks if a tfvars file exists.
func TfvarsExists(relativePath string) bool {
	if packageDir == "" {
		return false
	}
	path := GetTestdataPath(relativePath)
	_, err := os.Stat(path)
	return err == nil
}

// GetTestdataPath returns the absolute path to a testdata file.
func GetTestdataPath(relativePath string) string {
	return filepath.Join(packageDir, "testdata", relativePath)
}

// ResolvePlaceholder returns a freshly generated concrete value for a single
// placeholder token (e.g. "{{random_ipv6_network}}"). Function-specific
// placeholders round-trip to the matching acctest.Random* generator so a value
// produced by a function in the legacy test keeps that function's format
// (IPv6 network, CIDR, IP, MAC, hex, ...) instead of collapsing to a plain
// name token. A trailing numeric disambiguator (e.g. "{{random2}}") is matched
// against its base so distinct vars of the same type still resolve correctly.
//
// Ordering note: more specific bases must precede their prefixes (e.g.
// random_ipv6_network before random_ip, random_int before random_ip).
func ResolvePlaceholder(placeholder string) string {
	name := strings.TrimSuffix(strings.TrimPrefix(placeholder, "{{"), "}}")
	switch {
	case name == "random_octet":
		return fmt.Sprintf("%d", 1+rand.Intn(254)) // 1-254 valid IP host octet
	case name == "grid_master_hostname":
		return os.Getenv("NIOS_GRID_MASTER_HOSTNAME")
	case name == "discovery_member_hostname":
		return os.Getenv("NIOS_DISCOVERY_MEMBER_HOSTNAME")
	case name == "pxgrid_endpoint_ref":
		return os.Getenv("NIOS_PXGRID_ENDPOINT_REF")
	case strings.HasPrefix(name, "random_int"):
		return fmt.Sprintf("%d", 1+rand.Intn(9999))
	case strings.HasPrefix(name, "random_ipv6_network"):
		return RandomIPv6Network()
	case strings.HasPrefix(name, "random_cidr_network"):
		return RandomCIDRNetwork()
	case strings.HasPrefix(name, "random_mac"):
		return RandomMACAddress()
	case strings.HasPrefix(name, "random_hex32"):
		return Random32Hexadecimal()
	case strings.HasPrefix(name, "random_ip"):
		return RandomIP()
	default:
		return RandomNameWithPrefix("tf-acc-test")
	}
}

// ReplacePlaceholders replaces template placeholders with random values.
// Supported placeholders:
//   - {{random}} - Random string (e.g., "tf-acc-test-abc123")
//   - {{random_octet}} - Random 1-254 (valid IP host range, excludes 0=network and 255=broadcast)
//   - {{random_int}} - Random integer 1-9999
//   - {{random_ip}} - Random private IPv4 host address
//   - {{random_cidr_network}} - Random IPv4 network in CIDR notation
//   - {{random_ipv6_network}} - Random IPv6 network in CIDR notation
//   - {{random_mac}} - Random MAC address
//   - {{random_hex32}} - Random 32-character hexadecimal string
func ReplacePlaceholders(content string) string {
	result := content
	for _, ph := range placeholderPattern.FindAllString(content, -1) {
		result = strings.ReplaceAll(result, ph, ResolvePlaceholder(ph))
	}
	return result
}

// ProviderConfigHCL returns provider configuration HCL for the given backend.
// It reads credentials from environment variables.
func ProviderConfigHCL(backend string) string {
	switch backend {
	case "nios":
		return fmt.Sprintf(`
provider "infoblox" {
  nios = {
    host_url = %q
    username = %q
    password = %q
  }
}
`, os.Getenv("NIOS_HOST_URL"), os.Getenv("NIOS_USERNAME"), os.Getenv("NIOS_PASSWORD"))

	case "uddi":
		return fmt.Sprintf(`
provider "infoblox" {
  uddi = {
    csp_url = %q
    api_key = %q
  }
}
`, os.Getenv("BLOXONE_CSP_URL"), os.Getenv("BLOXONE_API_KEY"))

	default:
		return ""
	}
}
