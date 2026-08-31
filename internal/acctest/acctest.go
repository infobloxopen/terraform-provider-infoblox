package acctest

import (
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"

	niosclient "github.com/infobloxopen/infoblox-nios-go-client/client"
	niosoption "github.com/infobloxopen/infoblox-nios-go-client/option"
	uddiclient "github.com/infobloxopen/universal-ddi-go-client/client"
	uddioption "github.com/infobloxopen/universal-ddi-go-client/option"

	"github.com/infobloxopen/terraform-provider-infoblox/internal/provider"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/utils"
)

const (
	letterBytes = "abcdefghijklmnopqrstuvwxyz"

	envNIOSPassthru = "INFOBLOX_ACC_NIOS_PASSTHRU"
)

func NIOSPassthruEnabled() bool {
	return os.Getenv(envNIOSPassthru) == "true"
}

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

	// ProtoV6ProviderFactories instantiates the infoblox provider for acceptance testing.
	ProtoV6ProviderFactories = map[string]func() (tfprotov6.ProviderServer, error){
		"infoblox": providerserver.NewProtocol6WithError(provider.New("test", "test")()),
	}
)

// PreCheckNIOS validates NIOS environment and initializes the client.
// With INFOBLOX_ACC_NIOS_PASSTHRU=true the Grid is reached through the Infoblox Portal, and the
// verification client takes the same route as the provider rather than connecting to the Grid directly.
func PreCheckNIOS(t *testing.T) {
	if NIOSPassthruEnabled() {
		preCheckNIOSPassthru(t)
		return
	}

	t.Logf("NIOS transport: direct Grid (set %s=true for Infoblox Portal passthrough)", envNIOSPassthru)

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

// preCheckNIOSPassthru validates the Infoblox Portal environment and initializes a NIOS client
// that reaches the Grid through the Portal. Grid credentials are not used on this route.
func preCheckNIOSPassthru(t *testing.T) {
	t.Logf("NIOS transport: Infoblox Portal passthrough (%s=true)", envNIOSPassthru)

	portalURL := os.Getenv("INFOBLOX_PORTAL_URL")
	if portalURL == "" {
		t.Fatalf("INFOBLOX_PORTAL_URL must be set when %s=true", envNIOSPassthru)
	}

	portalKey := os.Getenv("INFOBLOX_PORTAL_KEY")
	if portalKey == "" {
		t.Fatalf("INFOBLOX_PORTAL_KEY must be set when %s=true", envNIOSPassthru)
	}

	licenseUID := os.Getenv("NIOS_LICENSE_UID")
	if licenseUID == "" {
		t.Fatalf("NIOS_LICENSE_UID must be set when %s=true", envNIOSPassthru)
	}

	NIOSClient = niosclient.NewAPIClient(
		niosoption.WithClientName("terraform-acceptance-tests"),
		niosoption.WithNIOSPassthrough(true),
		niosoption.WithPortalUrl(portalURL),
		niosoption.WithPortalAPIKey(portalKey),
		niosoption.WithNIOSLicenseUID(licenseUID),
		niosoption.WithDebug(true),
	)
}

// PreCheckUDDI validates UDDI environment and initializes the client.
func PreCheckUDDI(t *testing.T) {
	cspURL := os.Getenv("INFOBLOX_PORTAL_URL")
	if cspURL == "" {
		t.Fatal("INFOBLOX_PORTAL_URL must be set for UDDI acceptance tests")
	}

	apiKey := os.Getenv("INFOBLOX_PORTAL_KEY")
	if apiKey == "" {
		t.Fatal("INFOBLOX_PORTAL_KEY must be set for UDDI acceptance tests")
	}

	UDDIClient = uddiclient.NewAPIClient(
		uddioption.WithClientName("terraform-acceptance-tests"),
		uddioption.WithCSPUrl(cspURL),
		uddioption.WithAPIKey(apiKey),
		uddioption.WithDebug(true),
	)
}

// PreCheck validates the test environment for the given backend ("nios" or "uddi").
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

// RandomIP generates a random private IPv4 address in the 10.x.x.x range.
func RandomIP() string {
	return fmt.Sprintf("10.%d.%d.%d", rand.Intn(256), rand.Intn(256), 1+rand.Intn(254))
}

// RandomIPv6 generates a random IPv6 address under 2001:db8::/32.
func RandomIPv6() string {
	third := 1 + rand.Intn(65535)
	fourth := 1 + rand.Intn(65535)
	host := 1 + rand.Intn(65535)

	return fmt.Sprintf("2001:db8:%x:%x::%x", third, fourth, host)
}

// RandomOctet generates a random octet (0-255).
func RandomOctet() int {
	return rand.Intn(256)
}

// RandomIPWithSpecificOctetsSet returns an IP with the given prefix and a random last octet (1-254).
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

// RandomCIDRNetwork generates a random IPv4 network in CIDR notation (/16-/24).
func RandomCIDRNetwork() string {
	base := 10 + rand.Intn(246) // 10-255 for first octet
	second := rand.Intn(256)    // 0-255 for second octet
	cidr := 16 + rand.Intn(9)   // /16 to /24 (common for network containers)

	return fmt.Sprintf("%d.%d.0.0/%d", base, second, cidr)
}

// RandomIPv4Network generates a random /16-aligned IPv4 network address, without a prefix length.
func RandomIPv4Network() string {
	return fmt.Sprintf("%d.%d.0.0", 10+rand.Intn(246), rand.Intn(256))
}

// RandomIPv6Network generates a random IPv6 network under 2001:db8::/32 (RFC 3849 test range).
func RandomIPv6Network() string {
	third := rand.Intn(65536)  // 0-FFFF for third hextet
	fourth := rand.Intn(65536) // 0-FFFF for fourth hextet
	cidr := 64 + rand.Intn(60)

	return fmt.Sprintf("2001:db8:%x:%x::/%d", third, fourth, cidr)
}

// RandomIPv6NetworkAddress generates a random /64-aligned IPv6 network address, without a prefix length.
func RandomIPv6NetworkAddress() string {
	return fmt.Sprintf("2001:db8:%x:%x::", rand.Intn(65536), rand.Intn(65536))
}

// RandomIPv6NetworkWith4BitBoundary generates a random IPv6 network with a CIDR
// that is a 4-bit boundary (multiple of 4). This is required for operations like
// auto_create_reversezone which only supports 4-bit boundary CIDRs.
func RandomIPv6NetworkWith4BitBoundary() string {
	third := rand.Intn(65536)  // 0-FFFF for third hextet
	fourth := rand.Intn(65536) // 0-FFFF for fourth hextet
	// Valid 4-bit boundary CIDRs for IPv6: multiples of 4 between 64 and 124
	validCidrs := []int{64, 68, 72, 76, 80, 84, 88, 92, 96, 100, 104, 108, 112, 116, 120, 124}
	cidr := validCidrs[rand.Intn(len(validCidrs))]

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

// RandomMACAddress generates a random MAC address.
func RandomMACAddress() string {
	return fmt.Sprintf("%02x:%02x:%02x:%02x:%02x:%02x",
		rand.Intn(256),
		rand.Intn(256),
		rand.Intn(256),
		rand.Intn(256),
		rand.Intn(256),
		rand.Intn(256))
}

// Random32Hexadecimal generates a random 32-character hexadecimal string.
func Random32Hexadecimal() string {
	return fmt.Sprintf("%016x%016x", rand.Uint64(), rand.Uint64())
}

// CaseFileExists checks if a case file exists.
func CaseFileExists(relativePath string) bool {
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

// ResolvePlaceholder returns a concrete value for a single placeholder token (e.g. "{{random_ip}}").
// Each token maps to its matching Random* generator; unrecognised tokens produce a random name.
// More-specific prefixes must appear before shorter ones (e.g. random_ipv6_network before random_ip).
func ResolvePlaceholder(placeholder string) string {
	name := strings.TrimSuffix(strings.TrimPrefix(placeholder, "{{"), "}}")
	switch {
	case name == "random_octet":
		return fmt.Sprintf("%d", 1+rand.Intn(254)) // 1-254 valid IP host octet
	case name == "random_hextet":
		return fmt.Sprintf("%x", rand.Intn(65536)) // 0-FFFF single IPv6 hextet
	case name == "grid_master_hostname":
		return os.Getenv("NIOS_GRID_MASTER_HOSTNAME")
	case name == "grid_member_hostname":
		return os.Getenv("NIOS_GRID_MEMBER_HOSTNAME")
	case name == "grid_member_2_hostname":
		return os.Getenv("NIOS_GRID_MEMBER_2_HOSTNAME")
	case name == "discovery_member_hostname":
		return os.Getenv("NIOS_DISCOVERY_MEMBER_HOSTNAME")
	case name == "pxgrid_endpoint_ref":
		return os.Getenv("NIOS_PXGRID_ENDPOINT_REF")
	case strings.HasPrefix(name, "random_int"):
		return fmt.Sprintf("%d", 1+rand.Intn(9999))
	case strings.HasPrefix(name, "random_ipv6_network_address"):
		return RandomIPv6NetworkAddress()
	case strings.HasPrefix(name, "random_ipv6_network_4bit_boundary"):
		return RandomIPv6NetworkWith4BitBoundary()
	case strings.HasPrefix(name, "random_ipv6_network"):
		return RandomIPv6Network()
	case strings.HasPrefix(name, "random_ipv6"):
		return RandomIPv6()
	case strings.HasPrefix(name, "random_ipv4_network"):
		return RandomIPv4Network()
	case strings.HasPrefix(name, "random_cidr_network"):
		return RandomCIDRNetwork()
	case strings.HasPrefix(name, "random_mac"):
		return RandomMACAddress()
	case strings.HasPrefix(name, "random_hex32"):
		return Random32Hexadecimal()
	case strings.HasPrefix(name, "random_ip"):
		return RandomIP()
	case strings.HasPrefix(name, "future_time"):
		return FutureTime(name)
	default:
		return RandomNameWithPrefix("tf-acc-test")
	}
}

// FutureTime resolves a "future_time_<N>h" token to a timestamp N hours from now,
// formatted per utils.NaiveDatetimeLayout. Falls back to a 24-hour offset if N is missing/invalid.
func FutureTime(name string) string {
	hours := 24
	if suffix, ok := strings.CutSuffix(strings.TrimPrefix(name, "future_time_"), "h"); ok {
		if n, err := strconv.Atoi(suffix); err == nil {
			hours = n
		}
	}
	return time.Now().Add(time.Duration(hours) * time.Hour).UTC().Format(utils.NaiveDatetimeLayout)
}

// ReplacePlaceholders substitutes all {{token}} placeholders in content with random or env-sourced values.
// See ResolvePlaceholder for the full token list data.
func ReplacePlaceholders(content string) string {
	result := content
	for _, ph := range placeholderPattern.FindAllString(content, -1) {
		result = strings.ReplaceAll(result, ph, ResolvePlaceholder(ph))
	}
	return result
}

// ProviderConfigHCL returns provider HCL for backend ("nios" or "uddi"), reading credentials from env.
func ProviderConfigHCL(backend string) string {
	switch backend {
	case "nios":
		if NIOSPassthruEnabled() {
			return fmt.Sprintf(`
provider "infoblox" {
  uddi = {
    portal_url           = %q
    portal_key           = %q
    nios_license_uid     = %q
    enable_nios_passthru = true
  }
}
`, os.Getenv("INFOBLOX_PORTAL_URL"), os.Getenv("INFOBLOX_PORTAL_KEY"), os.Getenv("NIOS_LICENSE_UID"))
		}

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
    portal_url = %q
    portal_key = %q
  }
}
`, os.Getenv("INFOBLOX_PORTAL_URL"), os.Getenv("INFOBLOX_PORTAL_KEY"))

	default:
		return ""
	}
}
