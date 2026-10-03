// Objects created by this setup program (via PreConfig and main):
//
// Network Views (IPAM):
//   - test_network_view, custom_view, test_network_view2, ms_server, ms_server2
//
// Networks (IPAM, default network view):
//   - 10.0.0.0/24, 15.0.0.0/24, 16.0.0.0/24
//
// DHCP Filter Options:
//   - example-option-filter-1, example-option-filter-2
//
// DHCP MAC Filters:
//   - mac_filter, mac_filter2, example-mac-filter-1
//
// DHCP NAC Filters:
//   - nac_filter, nac_filter_rule, ipv6_nac_filter
//
// DHCP IPv6 Filter Options:
//   - ipv6_option_filter, ipv6_option_filter1
//
// DHCP Relay Agent Filters:
//   - relay_agent_filter, relay_agent_filter2
//
// DHCP Fingerprint Filters:
//   - test_filter_fingerprint, test_filter_fingerprint1
//
// DHCP Failover Associations:
//   - example_failover_association, example_failover_association1
//
// DNS Views:
//   - custom_dns_view
//
// DNS Zone Auth:
//   - example.com (default view, forward)
//   - 192.168.10.0/24 (default view, IPV4 reverse)
//   - 2001::/64 (default view, IPV6 reverse)
//   - tf-srg-zone-1.com (default view, forward)
//   - tf-srg-zone-2.com (default view, forward)
//
// DNS Zone RP (RPZ):
//   - test-rpz.com (default view)
//
// DNS Rulesets:
//   - blacklist_ruleset_1 (type: BLACKLIST)
//   - blacklist_ruleset_2 (type: BLACKLIST)
//   - nxdomain_ruleset_1 (type: NXDOMAIN)
//   - nxdomain_ruleset_2 (type: NXDOMAIN)
//
// DNS DDNS Principal Cluster Groups:
//   - dynamic_update_grp_1, dynamic_update_grp_2
//
// DNS64 Groups:
//   - dns64_group (prefix: 64:FF9B::/96)
//
// DNS Stub NS Groups:
//   - stub_ns_group1, stub_ns_group2
//
// DNS Forward Stub Servers:
//   - ensg1, ensg2
//
// Grid Members:
//   - infoblox.member2 (VIP: 172.28.38.251, BOTH addr type)
//   - infoblox.member3 (VIP: 172.28.38.252, IPV4 addr type)
//
// Grid Upgrade Dependent Groups:
//   - example_upgrade_dependent_group1, example_upgrade_dependent_group2
//
// Microsoft Servers:
//   - 10.10.10.10, example_server (default network view)
//   - ms_example_server (ms_server network view), ms_example_server2 (ms_server2 network view)
//
// Security Admin Users:
//   - aws1, aws2 (SAML auth, cloud-api-only group)
//
// Security Active Directory Auth Services:
//   - active_dir, active_dir_test
//
// RIR Organizations:
//   - ORG-CB11-INTEST (name: rir-org-test1), ORG-CB12-INTEST (name: rir-org-test2)
//
// Notification REST Endpoint:
//   - notification_rest_endpoint999
//
// Syslog Endpoint:
//   - syslogendpoint123
//
// pxGrid Endpoint:
//   - Example_pxgrid_ISE_endpoint
//
// CA Certificates (uploaded via fileop):
//   - EAP_CA cert from nios_security_certificate_authservice/cert.pem
//   - EAP_CA cert from nios_notification_rest_endpoint/dummy-bundle.pem
//
// Ecosystem Templates (uploaded if present in nios_ecosystem_templates/):
//   - Version5_DXL_Session_Template.json, Version5_Syslog_Session_Template.json
//   - Version5_Syslog_Action_Template.json, Version5_DXL_action_template.json
//   - Version5_DNS_Zone_and_Records.json, Version5_REST_API_Session_Template.json
//   - event_dhcp_lease_template.json, IPAM_PxgridEvent.json
//   - Version5_Syslog_Session_Template1.json, Version5_DXL_Session_Template1.json
//   - Version5_REST_API_Session_Template1.json, event_template_schema_default.json
//
// Objects updated (not created):
//   - Grid DHCP properties: enable_roaming_hosts=true
//   - Grid settings: enable_rir_swip=true, enable_network_users=true, dns resolvers=["127.0.0.1"]
//   - Parental control subscriber: enable_parental_control=true
//   - Upgrade schedule and distribution schedule: start_time and group times set
//   - Member file distribution: TFTP ACL (Address: Any, Permission: ALLOW) appended
//   - Grid DNS properties: log_rpz=true, allow_recursive_query=true
//   - Discovery member properties: enable_service=true, is_sa=true (if NIOS_DISCOVERY_MEMBER_URL is set)

package main

import (
	"bytes"
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/infobloxopen/infoblox-nios-go-client/client"
	"github.com/infobloxopen/infoblox-nios-go-client/dhcp"
	"github.com/infobloxopen/infoblox-nios-go-client/discovery"
	"github.com/infobloxopen/infoblox-nios-go-client/dns"
	"github.com/infobloxopen/infoblox-nios-go-client/grid"
	"github.com/infobloxopen/infoblox-nios-go-client/ipam"
	"github.com/infobloxopen/infoblox-nios-go-client/microsoft"
	"github.com/infobloxopen/infoblox-nios-go-client/misc"
	"github.com/infobloxopen/infoblox-nios-go-client/notification"
	"github.com/infobloxopen/infoblox-nios-go-client/option"
	"github.com/infobloxopen/infoblox-nios-go-client/parentalcontrol"
	"github.com/infobloxopen/infoblox-nios-go-client/rir"
	"github.com/infobloxopen/infoblox-nios-go-client/security"

	"github.com/infobloxopen/terraform-provider-infoblox/internal/acctest"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/core"
)

// UploadInitResponse holds the fields returned by the NIOS uploadinit fileop endpoint.
type UploadInitResponse struct {
	Token string `json:"token"`
	URL   string `json:"url"`
}

var pipelineEnvFile *os.File

func writePipelineEnvVar(key, value string) error {
	if pipelineEnvFile == nil {
		return fmt.Errorf("pipeline_nios.env file is not initialized")
	}

	if _, err := fmt.Fprintf(pipelineEnvFile, "%s=%s\n", key, value); err != nil {
		return fmt.Errorf("write env var %s to pipeline_nios.env: %w", key, err)
	}

	return nil
}

func firstNonEmpty(values ...string) string {
	for _, v := range values {
		if strings.TrimSpace(v) != "" {
			return v
		}
	}
	return ""
}

func normalizeAddressFromEnv(value string) string {
	v := strings.TrimSpace(value)
	if v == "" {
		return ""
	}

	if u, err := url.Parse(v); err == nil && u.Hostname() != "" {
		return strings.TrimSpace(u.Hostname())
	}

	if host, _, err := net.SplitHostPort(v); err == nil {
		return strings.Trim(host, "[]")
	}

	return v
}

type GridHostnames struct {
	MasterHostname          string
	MemberHostname          string
	DiscoveryMemberHostname string
}

// ResolveAndStoreGridHostnames resolves grid hostnames from VIP addresses and
// persists them into pipeline_nios.env for downstream integration tests.
func ResolveAndStoreGridHostnames(gridClient *grid.APIClient) (GridHostnames, error) {
	if gridClient == nil {
		return GridHostnames{}, fmt.Errorf("resolve hostnames: GRID client is required")
	}

	masterAddr := normalizeAddressFromEnv(os.Getenv("NIOS_HOST_URL"))
	memberAddr := normalizeAddressFromEnv(os.Getenv("NIOS_MEMBER_URL"))
	discoveryMemberAddr := normalizeAddressFromEnv(os.Getenv("NIOS_DISCOVERY_MEMBER_URL"))

	memberListResp, _, err := gridClient.MemberAPI.List(context.Background()).
		ReturnAsObject(1).
		ReturnFieldsPlus("vip_setting,host_name,config_addr_type").
		Execute()
	if err != nil {
		return GridHostnames{}, fmt.Errorf("resolve hostnames: failed to list members: %w", err)
	}

	if memberListResp.ListMemberResponseObject == nil {
		return GridHostnames{}, fmt.Errorf("resolve hostnames: member list response object is nil")
	}

	memberResp := memberListResp.ListMemberResponseObject.Result

	resolved := GridHostnames{}
	gridMasterConfigAddrType := ""
	discoveryMemberConfigAddrType := ""

	for _, m := range memberResp {
		if m.VipSetting == nil || m.VipSetting.Address == nil || m.HostName == nil {
			continue
		}

		vipAddr := strings.TrimSpace(*m.VipSetting.Address)
		hostName := strings.TrimSpace(*m.HostName)

		if masterAddr != "" && vipAddr == masterAddr && resolved.MasterHostname == "" {
			resolved.MasterHostname = hostName
			if m.ConfigAddrType != nil {
				gridMasterConfigAddrType = *m.ConfigAddrType
			}
		}

		if memberAddr != "" && vipAddr == memberAddr && resolved.MemberHostname == "" {
			resolved.MemberHostname = hostName
		}

		if discoveryMemberAddr != "" && vipAddr == discoveryMemberAddr && resolved.DiscoveryMemberHostname == "" {
			resolved.DiscoveryMemberHostname = hostName
			if m.ConfigAddrType != nil {
				discoveryMemberConfigAddrType = *m.ConfigAddrType
			}
		}
	}

	if resolved.MasterHostname != "" {
		if err := writePipelineEnvVar("NIOS_GRID_MASTER_HOSTNAME", resolved.MasterHostname); err != nil {
			return GridHostnames{}, fmt.Errorf("resolve hostnames: failed to write NIOS_GRID_MASTER_HOSTNAME: %w", err)
		}
		fmt.Printf("Mapped master VIP %q to hostname %q\n", masterAddr, resolved.MasterHostname)
	}

	if resolved.MemberHostname != "" {
		if err := writePipelineEnvVar("NIOS_GRID_MEMBER_HOSTNAME", resolved.MemberHostname); err != nil {
			return GridHostnames{}, fmt.Errorf("resolve hostnames: failed to write NIOS_GRID_MEMBER_HOSTNAME: %w", err)
		}
		fmt.Printf("Mapped member VIP %q to hostname %q\n", memberAddr, resolved.MemberHostname)
	}

	if resolved.DiscoveryMemberHostname != "" {
		if err := writePipelineEnvVar("NIOS_DISCOVERY_MEMBER_HOSTNAME", resolved.DiscoveryMemberHostname); err != nil {
			return GridHostnames{}, fmt.Errorf("resolve hostnames: failed to write NIOS_DISCOVERY_MEMBER_HOSTNAME: %w", err)
		}
		fmt.Printf("Mapped discovery member VIP %q to hostname %q\n", discoveryMemberAddr, resolved.DiscoveryMemberHostname)
	}

	// Add Grid Master Config Addr Type as env variable
	if gridMasterConfigAddrType != "" {
		if err := writePipelineEnvVar("NIOS_GRID_MASTER_CONFIG_ADDR_TYPE", gridMasterConfigAddrType); err != nil {
			return GridHostnames{}, fmt.Errorf("resolve hostnames: failed to write NIOS_GRID_MASTER_CONFIG_ADDR_TYPE: %w", err)
		}
		fmt.Printf("Grid master config address type is %q\n", gridMasterConfigAddrType)
	}

	if discoveryMemberConfigAddrType != "" {
		if err := writePipelineEnvVar("NIOS_DISCOVERY_MEMBER_CONFIG_ADDR_TYPE", discoveryMemberConfigAddrType); err != nil {
			return GridHostnames{}, fmt.Errorf("resolve hostnames: failed to write NIOS_DISCOVERY_MEMBER_CONFIG_ADDR_TYPE: %w", err)
		}
		fmt.Printf("Discovery member config address type is %q\n", discoveryMemberConfigAddrType)
	}

	if masterAddr != "" && resolved.MasterHostname == "" {
		fmt.Printf("No member found with VIP address matching NIOS_HOST_URL (%q)\n", masterAddr)
	}

	if memberAddr != "" && resolved.MemberHostname == "" {
		fmt.Printf("No member found with VIP address matching NIOS_MEMBER_URL (%q)\n", memberAddr)
	}

	if discoveryMemberAddr != "" && resolved.DiscoveryMemberHostname == "" {
		fmt.Printf("No member found with VIP address matching NIOS_DISCOVERY_MEMBER_URL (%q)\n", discoveryMemberAddr)
	}

	return resolved, nil
}

// UploadInit calls the NIOS fileop uploadinit endpoint and returns the upload
// token and URL needed for subsequent file-upload operations.
func UploadInit(host, wapiVer, username, password string) (*UploadInitResponse, error) {
	url := fmt.Sprintf("%s/wapi/%s/fileop?_function=uploadinit", host, wapiVer)

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBufferString("{}"))
	if err != nil {
		return nil, fmt.Errorf("uploadinit: create request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.SetBasicAuth(username, password)

	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true}, //nolint:gosec // intentional for lab/CI grids with self-signed certs
		},
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("uploadinit: execute request: %w", err)
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("uploadinit: unexpected status %s", resp.Status)
	}

	var result UploadInitResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("uploadinit: decode response: %w", err)
	}

	return &result, nil
}

// UploadCertificateRequest holds the body fields for the uploadcertificate fileop call.
type UploadCertificateRequest struct {
	CertificateUsage string `json:"certificate_usage"`
	Member           string `json:"member"`
	Token            string `json:"token"`
}

type UploadFileRequest struct {
	File string `json:"file"`
	Url  string `json:"url"`
}

func UploadFile(host, wapiVer, username, password, url, filePath string) error {
	file, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("uploadfile: open file %q: %w", filePath, err)
	}
	defer func() { _ = file.Close() }()

	var requestBody bytes.Buffer
	writer := multipart.NewWriter(&requestBody)

	formFile, err := writer.CreateFormFile("file", filepath.Base(filePath))
	if err != nil {
		return fmt.Errorf("uploadfile: create form file: %w", err)
	}

	if _, err := io.Copy(formFile, file); err != nil {
		return fmt.Errorf("uploadfile: copy file contents: %w", err)
	}

	if err := writer.Close(); err != nil {
		return fmt.Errorf("uploadfile: close multipart writer: %w", err)
	}

	req, err := http.NewRequest(http.MethodPost, url, &requestBody)
	if err != nil {
		return fmt.Errorf("uploadfile: create request: %w", err)
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())
	req.SetBasicAuth(username, password)

	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true}, //nolint:gosec // intentional for lab/CI grids with self-signed certs
		},
	}

	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("uploadfile: execute request: %w", err)
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("uploadfile: unexpected status %s", resp.Status)
	}

	return nil

}

// UploadCertificate uses the token from UploadInit to upload a certificate to NIOS.
func UploadCertificate(host, wapiVer, username, password, certificateUsage, member, token string) error {
	endpoint := fmt.Sprintf("%s/wapi/%s/fileop?_function=uploadcertificate", host, wapiVer)

	cleanToken := strings.TrimSpace(token)

	body := UploadCertificateRequest{
		CertificateUsage: certificateUsage,
		Member:           member,
		Token:            cleanToken,
	}
	bodyBytes, err := json.Marshal(body)
	if err != nil {
		return fmt.Errorf("uploadcertificate: marshal request body: %w", err)
	}

	req, err := http.NewRequest(http.MethodPost, endpoint, bytes.NewBuffer(bodyBytes))
	if err != nil {
		return fmt.Errorf("uploadcertificate: create request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.SetBasicAuth(username, password)

	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true}, //nolint:gosec // intentional for lab/CI grids with self-signed certs
		},
	}

	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("uploadcertificate: execute request: %w", err)
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode == 400 {
		fmt.Printf("uploadcertificate: bad request. This may indicate that that either `disable_strict_ca_cert_check` is not set on NIOS or Certificate is already uploaded. Skipping certificate upload.\n")
		return nil
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("uploadcertificate: unexpected status %s", resp.Status)
	}

	return nil
}

type UploadEcoSystemTemplatesRequest struct {
	Token string `json:"token"`
}

func UploadEcoSystemTemplates(host, wapiVer, username, password, token string) error {
	endpoint := fmt.Sprintf("%s/wapi/%s/fileop?_function=restapi_template_import", host, wapiVer)

	cleanToken := strings.TrimSpace(token)

	body := UploadEcoSystemTemplatesRequest{
		Token: cleanToken,
	}
	bodyBytes, err := json.Marshal(body)
	if err != nil {
		return fmt.Errorf("uploadecosystemtemplates: marshal request body: %w", err)
	}

	req, err := http.NewRequest(http.MethodPost, endpoint, bytes.NewBuffer(bodyBytes))
	if err != nil {
		return fmt.Errorf("uploadecosystemtemplates: create request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.SetBasicAuth(username, password)

	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true}, //nolint:gosec // intentional for lab/CI grids with self-signed certs
		},
	}

	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("uploadecosystemtemplates: execute request: %w", err)
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("uploadecosystemtemplates: unexpected status %s", resp.Status)
	}

	return nil
}

// ConfigureEcoSystemTemplates uploads the ecosystem template file and triggers
// template import using the upload token lifecycle.
func ConfigureEcoSystemTemplates(host, wapiVer, username, password, ecosystemTemplatePath string) error {
	certUpload, err := UploadInit(host, wapiVer, username, password)
	if err != nil {
		return fmt.Errorf("configure ecosystem templates: upload init: %w", err)
	}

	if err := UploadFile(host, wapiVer, username, password, certUpload.URL, ecosystemTemplatePath); err != nil {
		return fmt.Errorf("configure ecosystem templates: upload file: %w", err)
	}

	if err := UploadEcoSystemTemplates(host, wapiVer, username, password, certUpload.Token); err != nil {
		return fmt.Errorf("configure ecosystem templates: import templates: %w", err)
	}

	return nil
}

// ConfigureCACertificates uploads a CA certificate file, imports it into NIOS,
// and stores the created certificate ref in the provided environment variable.
func ConfigureCACertificates(host, wapiVer, username, password, certificateUsage, member, certificateFilePath, certRefEnvVar, certSerialEnvVar string) error {
	certUpload, err := UploadInit(host, wapiVer, username, password)
	if err != nil {
		return fmt.Errorf("configure CA certificates: upload init: %w", err)
	}

	if err := UploadFile(host, wapiVer, username, password, certUpload.URL, certificateFilePath); err != nil {
		return fmt.Errorf("configure CA certificates: upload file: %w", err)
	}

	if err := UploadCertificate(host, wapiVer, username, password, certificateUsage, member, certUpload.Token); err != nil {
		return fmt.Errorf("configure CA certificates: upload certificate: %w", err)
	}

	if err := FetchAndStoreCertificateRef(host, wapiVer, username, password, certRefEnvVar, certSerialEnvVar); err != nil {
		return fmt.Errorf("configure CA certificates: fetch/store cert ref: %w", err)
	}

	return nil
}

func ConfigurePxgridEndpoint(host, wapiVer, username, password, certUploadFilePath string, miscClient *misc.APIClient) error {
	if miscClient == nil {
		return fmt.Errorf("configure pxgrid endpoint: MISC client is required")
	}

	uploadInitResp, err := UploadInit(host, wapiVer, username, password)
	if err != nil {
		return fmt.Errorf("configure pxgrid endpoint: upload init: %w", err)
	}

	if err := UploadFile(host, wapiVer, username, password, uploadInitResp.URL, certUploadFilePath); err != nil {
		return fmt.Errorf("configure pxgrid endpoint: upload file: %w", err)
	}

	endpointName := "Example_pxgrid_ISE_endpoint"
	resolvedOutboundMember := firstNonEmpty("infoblox.member2")

	pxGridEndpointBody := misc.PxgridEndpoint{
		Address:                misc.PtrString("1.1.1.1"),
		Name:                   misc.PtrString(endpointName),
		NetworkView:            misc.PtrString("test_network_view"),
		ClientCertificateToken: misc.PtrString(uploadInitResp.Token),
		OutboundMemberType:     misc.PtrString("MEMBER"),
		OutboundMembers:        []string{resolvedOutboundMember},
		SubscribeSettings: &misc.PxgridEndpointSubscribeSettings{
			EnabledAttributes: []string{"DOMAINNAME"},
		},
		PublishSettings: &misc.PxgridEndpointPublishSettings{
			EnabledAttributes: []string{"IPADDRESS"},
		},
	}

	apiResp, _, err := miscClient.PxgridEndpointAPI.Create(context.Background()).
		PxgridEndpoint(pxGridEndpointBody).
		ReturnAsObject(1).
		ReturnFieldsPlus("name").
		Execute()
	if err != nil {
		if strings.Contains(err.Error(), "already exists") || strings.Contains(err.Error(), "is already in use for another endpoint") {
			existingResp, _, listErr := miscClient.PxgridEndpointAPI.List(context.Background()).
				ReturnAsObject(1).
				ReturnFieldsPlus("name").
				Execute()
			if listErr != nil {
				return fmt.Errorf("configure pxgrid endpoint: failed to list existing pxgrid endpoint: %w", listErr)
			}
			if existingResp.ListPxgridEndpointResponseObject == nil || len(existingResp.ListPxgridEndpointResponseObject.Result) == 0 {
				return fmt.Errorf("configure pxgrid endpoint: pxgrid endpoint %q already exists but no matching object was returned", endpointName)
			}

			existingRef := existingResp.ListPxgridEndpointResponseObject.Result[0].GetRef()
			if existingRef == "" {
				return fmt.Errorf("configure pxgrid endpoint: pxgrid endpoint %q already exists but ref could not be resolved", endpointName)
			}

			if err := writePipelineEnvVar("NIOS_PXGRID_ENDPOINT_REF", existingRef); err != nil {
				return fmt.Errorf("configure pxgrid endpoint: failed to write existing pxgrid endpoint ref: %w", err)
			}

			fmt.Printf("pxGrid endpoint %q already exists, using ref %s\n", endpointName, existingRef)
			return nil
		}

		return fmt.Errorf("configure pxgrid endpoint: failed to create pxgrid endpoint: %w", err)
	}

	if apiResp.CreatePxgridEndpointResponseAsObject == nil || apiResp.CreatePxgridEndpointResponseAsObject.Result == nil {
		return fmt.Errorf("configure pxgrid endpoint: create response did not include a result")
	}

	createdRef := apiResp.CreatePxgridEndpointResponseAsObject.Result.GetRef()
	if createdRef == "" {
		return fmt.Errorf("configure pxgrid endpoint: created pxgrid endpoint ref is empty")
	}

	if err := writePipelineEnvVar("NIOS_PXGRID_ENDPOINT_REF", createdRef); err != nil {
		return fmt.Errorf("configure pxgrid endpoint: failed to write created pxgrid endpoint ref: %w", err)
	}

	fmt.Printf("pxGrid endpoint %q created successfully with ref %s\n", endpointName, createdRef)
	return nil
}

// CACertificate represents a single CA certificate object from the NIOS WAPI.
type CACertificate struct {
	Ref               string `json:"_ref"`
	DistinguishedName string `json:"distinguished_name"`
	Issuer            string `json:"issuer"`
	Serial            string `json:"serial"`
	ValidNotAfter     int64  `json:"valid_not_after"`
	ValidNotBefore    int64  `json:"valid_not_before"`
}

// FetchAndStoreCertificateRef fetches the CA certificate ref from NIOS WAPI and writes it
// into the opened pipeline_nios.env file using the provided envVarName key.
func FetchAndStoreCertificateRef(host, wapiVer, username, password, envVarName, certSerialEnvVar string) error {
	endpoint := fmt.Sprintf("%s/wapi/%s/cacertificate", host, wapiVer)

	req, err := http.NewRequest(http.MethodGet, endpoint, nil)
	if err != nil {
		return fmt.Errorf("fetchcertref: create request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.SetBasicAuth(username, password)

	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true}, //nolint:gosec // intentional for lab/CI grids with self-signed certs
		},
	}

	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("fetchcertref: execute request: %w", err)
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("fetchcertref: unexpected status %s", resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("fetchcertref: read response body: %w", err)
	}

	var certificates []CACertificate
	if err := json.Unmarshal(body, &certificates); err != nil {
		return fmt.Errorf("fetchcertref: unmarshal response: %w", err)
	}

	if len(certificates) == 0 {
		return fmt.Errorf("fetchcertref: no certificates found in response")
	}

	certRef := certificates[0].Ref
	if certRef == "" {
		return fmt.Errorf("fetchcertref: certificate ref is empty")
	}

	if err := writePipelineEnvVar(envVarName, certRef); err != nil {
		return fmt.Errorf("fetchcertref: failed to write env variable %s: %w", envVarName, err)
	}

	certSerial := certificates[0].Serial
	if certSerial == "" {
		return fmt.Errorf("fetchcertref: certificate serial is empty")
	}

	if err := writePipelineEnvVar(certSerialEnvVar, certSerial); err != nil {
		return fmt.Errorf("fetchcertref: failed to write env variable %s: %w", certSerialEnvVar, err)
	}

	return nil
}

type PreConfigClients struct {
	IPAM         *ipam.APIClient
	DHCP         *dhcp.APIClient
	DNS          *dns.APIClient
	GRID         *grid.APIClient
	MISC         *misc.APIClient
	MICROSOFT    *microsoft.APIClient
	NOTIFICATION *notification.APIClient
	PARENTAL     *parentalcontrol.APIClient
	SECURITY     *security.APIClient
	RIR          *rir.APIClient
	DISCOVERY    *discovery.APIClient
}

// PreConfig creates the network views required for integration testing.
// If a network view already exists (error contains "already exists"), it skips creation and continues.
// For any other error, it returns the error immediately.
func PreConfig(clients PreConfigClients, hostnames GridHostnames) error {
	if clients.IPAM == nil {
		return fmt.Errorf("preconfig: IPAM client is required")
	}
	if clients.DHCP == nil {
		return fmt.Errorf("preconfig: DHCP client is required")
	}
	if clients.DNS == nil {
		return fmt.Errorf("preconfig: DNS client is required")
	}
	if clients.GRID == nil {
		return fmt.Errorf("preconfig: GRID client is required")
	}
	if clients.MISC == nil {
		return fmt.Errorf("preconfig: MISC client is required")
	}
	if clients.MICROSOFT == nil {
		return fmt.Errorf("preconfig: MICROSOFT client is required")
	}
	if clients.NOTIFICATION == nil {
		return fmt.Errorf("preconfig: NOTIFICATION client is required")
	}
	if clients.PARENTAL == nil {
		return fmt.Errorf("preconfig: PARENTAL client is required")
	}
	if clients.SECURITY == nil {
		return fmt.Errorf("preconfig: SECURITY client is required")
	}
	if clients.RIR == nil {
		return fmt.Errorf("preconfig: RIR client is required")
	}

	// Ensure roaming hosts support is enabled at grid DHCP properties level.
	gridDHCPResp, _, err := clients.GRID.GridDhcppropertiesAPI.List(context.Background()).
		ReturnAsObject(1).
		ReturnFieldsPlus("enable_roaming_hosts").
		Execute()
	if err != nil {
		return fmt.Errorf("failed to list grid DHCP properties: %w", err)
	}

	if gridDHCPResp.ListGridDhcppropertiesResponseObject == nil || len(gridDHCPResp.ListGridDhcppropertiesResponseObject.Result) == 0 {
		return fmt.Errorf("grid DHCP properties list returned no results")
	}

	gridDHCPProperties := gridDHCPResp.ListGridDhcppropertiesResponseObject.Result[0]
	if gridDHCPProperties.Ref == nil || *gridDHCPProperties.Ref == "" {
		return fmt.Errorf("grid DHCP properties ref is empty")
	}

	enableRoamingHosts := false
	if gridDHCPProperties.EnableRoamingHosts != nil {
		enableRoamingHosts = *gridDHCPProperties.EnableRoamingHosts
	}

	if !enableRoamingHosts {
		gridDHCPPropertiesBody := grid.GridDhcpproperties{
			EnableRoamingHosts: grid.PtrBool(true),
		}

		gridDhcpPropertiesRef := core.ExtractNIOSRef(*gridDHCPProperties.Ref)
		_, _, err = clients.GRID.GridDhcppropertiesAPI.Update(context.Background(), gridDhcpPropertiesRef).
			GridDhcpproperties(gridDHCPPropertiesBody).
			Execute()
		if err != nil {
			return fmt.Errorf("failed to enable roaming hosts on grid DHCP properties %q: %w", *gridDHCPProperties.Ref, err)
		}

		fmt.Printf("Roaming hosts enabled for grid DHCP properties %q\n", *gridDHCPProperties.Ref)
	} else {
		fmt.Printf("Roaming hosts already enabled for grid DHCP properties %q, skipping update\n", *gridDHCPProperties.Ref)
	}

	// Update Grid settings: enable RIR SWIP and MS network users.
	gridListResp, _, err := clients.GRID.GridAPI.List(context.Background()).
		ReturnAsObject(1).
		ReturnFieldsPlus("enable_rir_swip,ms_setting").
		Execute()
	if err != nil {
		return fmt.Errorf("failed to list grid objects: %w", err)
	}
	if gridListResp.ListGridResponseObject == nil || len(gridListResp.ListGridResponseObject.Result) == 0 {
		return fmt.Errorf("grid list returned no results")
	}
	gridRef := gridListResp.ListGridResponseObject.Result[0].Ref
	if gridRef == nil || *gridRef == "" {
		return fmt.Errorf("grid ref is empty")
	}
	gridBody := grid.Grid{
		MsSetting: &grid.GridMsSetting{
			EnableNetworkUsers: grid.PtrBool(true),
		},
		EnableRirSwip: grid.PtrBool(true),
		DnsResolverSetting: &grid.GridDnsResolverSetting{
			Resolvers: []string{"127.0.0.1"},
		},
	}
	_, _, err = clients.GRID.GridAPI.Update(context.Background(), core.ExtractNIOSRef(*gridRef)).
		Grid(gridBody).
		Execute()
	if err != nil {
		return fmt.Errorf("failed to update grid settings for %q: %w", *gridRef, err)
	}
	fmt.Printf("Grid settings updated (enable_rir_swip=true, enable_network_users=true) for %q\n", *gridRef)

	// Ensure parental control is enabled.
	parentalResp, _, err := clients.PARENTAL.ParentalcontrolSubscriberAPI.List(context.Background()).
		ReturnAsObject(1).
		ReturnFieldsPlus("enable_parental_control").
		Execute()
	if err != nil {
		return fmt.Errorf("failed to list parentalcontrol subscribers: %w", err)
	}
	if parentalResp.ListParentalcontrolSubscriberResponseObject == nil || len(parentalResp.ListParentalcontrolSubscriberResponseObject.Result) == 0 {
		return fmt.Errorf("parentalcontrol subscriber list returned no results")
	}

	parentalSubscriber := parentalResp.ListParentalcontrolSubscriberResponseObject.Result[0]
	if parentalSubscriber.Ref == nil || *parentalSubscriber.Ref == "" {
		return fmt.Errorf("parentalcontrol subscriber ref is empty")
	}

	parentalControlStatus := false
	if parentalSubscriber.EnableParentalControl != nil {
		parentalControlStatus = *parentalSubscriber.EnableParentalControl
	}

	if !parentalControlStatus {
		parentalControlRef := core.ExtractNIOSRef(*parentalSubscriber.Ref)
		parentalControlSubscriberBody := parentalcontrol.ParentalcontrolSubscriber{
			EnableParentalControl: parentalcontrol.PtrBool(true),
			CatAcctname:           parentalcontrol.PtrString("test_account"),
			PcZoneName:            parentalcontrol.PtrString("PcZone"),
		}

		_, _, err = clients.PARENTAL.ParentalcontrolSubscriberAPI.Update(context.Background(), parentalControlRef).
			ParentalcontrolSubscriber(parentalControlSubscriberBody).
			Execute()
		if err != nil {
			return fmt.Errorf("failed to enable parental control on %q: %w", *parentalSubscriber.Ref, err)
		}

		fmt.Printf("Parental control enabled for subscriber %q\n", *parentalSubscriber.Ref)
	} else {
		fmt.Printf("Parental control already enabled for subscriber %q, skipping update\n", *parentalSubscriber.Ref)
	}

	// Check if ip_space_discriminator is available as a readable field.
	ipDiscResp, _, err := clients.PARENTAL.ParentalcontrolSubscriberAPI.List(context.Background()).
		ReturnAsObject(1).
		ReturnFieldsPlus("ip_space_discriminator").
		Execute()
	if err != nil {
		return fmt.Errorf("failed to list parentalcontrol subscribers for ip_space_discriminator check: %w", err)
	}
	if ipDiscResp.ListParentalcontrolSubscriberResponseObject == nil || len(ipDiscResp.ListParentalcontrolSubscriberResponseObject.Result) == 0 {
		return fmt.Errorf("parentalcontrol subscriber list returned no results for ip_space_discriminator check")
	}

	ipDiscSubscriber := ipDiscResp.ListParentalcontrolSubscriberResponseObject.Result[0]
	if ipDiscSubscriber.IpSpaceDiscriminator != nil && *ipDiscSubscriber.IpSpaceDiscriminator != "" {
		// ip_space_discriminator already present — field is readable/editable
		fmt.Printf("ip_space_discriminator already present (%q) for subscriber %q, marking editable\n",
			*ipDiscSubscriber.IpSpaceDiscriminator, *ipDiscSubscriber.Ref)
		if err := writePipelineEnvVar("SUBSCRIBER_BLOCK_SIZE_EDITABLE", "true"); err != nil {
			return fmt.Errorf("failed to write SUBSCRIBER_BLOCK_SIZE_EDITABLE: %w", err)
		}
	} else {
		// ip_space_discriminator absent — attempt an Update to probe editability
		if ipDiscSubscriber.Ref == nil || *ipDiscSubscriber.Ref == "" {
			return fmt.Errorf("parentalcontrol subscriber ref is empty for ip_space_discriminator update")
		}
		ipDiscRef := core.ExtractNIOSRef(*ipDiscSubscriber.Ref)
		_, _, updateErr := clients.PARENTAL.ParentalcontrolSubscriberAPI.Update(context.Background(), ipDiscRef).
			ParentalcontrolSubscriber(parentalcontrol.ParentalcontrolSubscriber{
				IpSpaceDiscriminator: parentalcontrol.PtrString("Deterministic-NAT-Port"),
			}).
			Execute()
		if updateErr != nil {
			fmt.Printf("ip_space_discriminator update failed for subscriber %q (API error: %v), marking not editable\n",
				*ipDiscSubscriber.Ref, updateErr)
			if err := writePipelineEnvVar("SUBSCRIBER_BLOCK_SIZE_EDITABLE", "false"); err != nil {
				return fmt.Errorf("failed to write SUBSCRIBER_BLOCK_SIZE_EDITABLE: %w", err)
			}
		} else {
			fmt.Printf("ip_space_discriminator updated to \"Deterministic-NAT-Port\" for subscriber %q, marking editable\n",
				*ipDiscSubscriber.Ref)
			if err := writePipelineEnvVar("SUBSCRIBER_BLOCK_SIZE_EDITABLE", "true"); err != nil {
				return fmt.Errorf("failed to write SUBSCRIBER_BLOCK_SIZE_EDITABLE: %w", err)
			}
		}
	}

	networkViews := []string{"test_network_view", "custom_view", "test_network_view2", "ms_server", "ms_server2"}

	for _, viewName := range networkViews {
		networkViewBody := ipam.Networkview{
			Name: ipam.PtrString(viewName),
		}

		_, _, err := clients.IPAM.NetworkviewAPI.Create(context.Background()).
			Networkview(networkViewBody).
			Execute()

		if err != nil {
			// Check if the error is because the object already exists
			if strings.Contains(err.Error(), "already exists") {
				fmt.Printf("Network view %q already exists, skipping creation\n", viewName)
				continue
			}
			// Return any other error
			return fmt.Errorf("failed to create network view %q: %w", viewName, err)
		}

		fmt.Printf("Network view %q created successfully\n", viewName)
	}

	if hostnames.MasterHostname == "" || hostnames.MemberHostname == "" {
		resolvedHostnames, err := ResolveAndStoreGridHostnames(clients.GRID)
		if err != nil {
			return err
		}

		if hostnames.MasterHostname == "" {
			hostnames.MasterHostname = resolvedHostnames.MasterHostname
		}

		if hostnames.MemberHostname == "" {
			hostnames.MemberHostname = resolvedHostnames.MemberHostname
		}
	}

	memberHostname := strings.TrimSpace(hostnames.MemberHostname)

	// Create networks
	type networkEntry struct {
		cidr        string
		networkView string
	}
	networks := []networkEntry{
		{"10.0.0.0/24", "default"},
		{"15.0.0.0/24", "default"},
		{"16.0.0.0/24", "default"},
	}

	for _, n := range networks {
		networkBody := ipam.Network{
			Network: &ipam.NetworkNetwork{
				String: ipam.PtrString(n.cidr),
			},
			NetworkView: ipam.PtrString(n.networkView),
			Comment:     ipam.PtrString("Created For Integration Testing"),
		}

		_, _, err := clients.IPAM.NetworkAPI.Create(context.Background()).
			Network(networkBody).
			Execute()

		if err != nil {
			if strings.Contains(err.Error(), "already exists") {
				fmt.Printf("Network %q already exists, skipping creation\n", n.cidr)
				continue
			}
			return fmt.Errorf("failed to create network %q: %w", n.cidr, err)
		}

		fmt.Printf("Network %q created successfully\n", n.cidr)
	}

	// Create filter options
	filterOptions := []string{"example-option-filter-1", "example-option-filter-2"}

	for _, filterName := range filterOptions {
		filterOptionBody := dhcp.Filteroption{
			Name: dhcp.PtrString(filterName),
		}

		_, _, err := clients.DHCP.FilteroptionAPI.Create(context.Background()).
			Filteroption(filterOptionBody).
			Execute()

		if err != nil {
			if strings.Contains(err.Error(), "already exists") {
				fmt.Printf("Filter option %q already exists, skipping creation\n", filterName)
				continue
			}
			return fmt.Errorf("failed to create filter option %q: %w", filterName, err)
		}

		fmt.Printf("Filter option %q created successfully\n", filterName)
	}

	// Create MAC filters
	macFilters := []string{"mac_filter", "mac_filter2", "example-mac-filter-1"}

	for _, macFilterName := range macFilters {
		filterMACBody := dhcp.Filtermac{
			Name: dhcp.PtrString(macFilterName),
		}

		_, _, err := clients.DHCP.FiltermacAPI.Create(context.Background()).
			Filtermac(filterMACBody).
			Execute()

		if err != nil {
			if strings.Contains(err.Error(), "already exists") {
				fmt.Printf("MAC filter %q already exists, skipping creation\n", macFilterName)
				continue
			}
			return fmt.Errorf("failed to create MAC filter %q: %w", macFilterName, err)
		}

		fmt.Printf("MAC filter %q created successfully\n", macFilterName)
	}

	// Create NAC filters
	nacFilters := []string{"nac_filter", "nac_filter_rule", "ipv6_nac_filter"}

	for _, nacFilterName := range nacFilters {
		filterNACBody := dhcp.Filternac{
			Name: dhcp.PtrString(nacFilterName),
		}

		_, _, err := clients.DHCP.FilternacAPI.Create(context.Background()).
			Filternac(filterNACBody).
			Execute()

		if err != nil {
			if strings.Contains(err.Error(), "already exists") {
				fmt.Printf("NAC filter %q already exists, skipping creation\n", nacFilterName)
				continue
			}
			return fmt.Errorf("failed to create NAC filter %q: %w", nacFilterName, err)
		}

		fmt.Printf("NAC filter %q created successfully\n", nacFilterName)
	}

	// Create IPv6 filter options
	ipv6FilterOptions := []string{"ipv6_option_filter", "ipv6_option_filter1"}

	for _, ipv6FilterName := range ipv6FilterOptions {
		ipv6FilterOptionBody := dhcp.Ipv6filteroption{
			Name: dhcp.PtrString(ipv6FilterName),
		}

		_, _, err := clients.DHCP.Ipv6filteroptionAPI.Create(context.Background()).
			Ipv6filteroption(ipv6FilterOptionBody).
			Execute()

		if err != nil {
			if strings.Contains(err.Error(), "already exists") {
				fmt.Printf("IPv6 filter option %q already exists, skipping creation\n", ipv6FilterName)
				continue
			}
			return fmt.Errorf("failed to create IPv6 filter option %q: %w", ipv6FilterName, err)
		}

		fmt.Printf("IPv6 filter option %q created successfully\n", ipv6FilterName)
	}

	// Create relay agent filters
	relayAgentFilters := []string{"relay_agent_filter", "relay_agent_filter2"}

	for _, relayFilterName := range relayAgentFilters {
		filterRelayAgentBody := dhcp.Filterrelayagent{
			Name:        dhcp.PtrString(relayFilterName),
			IsCircuitId: dhcp.PtrString("NOT_SET"),
		}

		_, _, err := clients.DHCP.FilterrelayagentAPI.Create(context.Background()).
			Filterrelayagent(filterRelayAgentBody).
			Execute()

		if err != nil {
			if strings.Contains(err.Error(), "already exists") {
				fmt.Printf("Relay agent filter %q already exists, skipping creation\n", relayFilterName)
				continue
			}
			return fmt.Errorf("failed to create relay agent filter %q: %w", relayFilterName, err)
		}

		fmt.Printf("Relay agent filter %q created successfully\n", relayFilterName)
	}

	// Create fingerprint filters
	fingerprintFilters := []string{"test_filter_fingerprint", "test_filter_fingerprint1"}

	resp2, _, err := clients.DHCP.FingerprintAPI.List(context.Background()).ReturnAsObject(1).Execute()
	if err != nil {
		return fmt.Errorf("failed to list fingerprints: %w", err)
	}
	if resp2 == nil || resp2.ListFingerprintResponseObject == nil || len(resp2.ListFingerprintResponseObject.Result) == 0 {
		return fmt.Errorf("no fingerprints found in response")
	}
	fingerPrintName := resp2.ListFingerprintResponseObject.Result[0].Name

	for _, fpFilterName := range fingerprintFilters {
		filterFingerprintBody := dhcp.Filterfingerprint{
			Name:        dhcp.PtrString(fpFilterName),
			Fingerprint: []string{*fingerPrintName},
		}

		_, _, err := clients.DHCP.FilterfingerprintAPI.Create(context.Background()).
			Filterfingerprint(filterFingerprintBody).
			Execute()

		if err != nil {
			if strings.Contains(err.Error(), "already exists") {
				fmt.Printf("Fingerprint filter %q already exists, skipping creation\n", fpFilterName)
				continue
			}
			return fmt.Errorf("failed to create fingerprint filter %q: %w", fpFilterName, err)
		}

		fmt.Printf("Fingerprint filter %q created successfully\n", fpFilterName)
	}

	// Create admin users
	adminUsers := []string{"aws1", "aws2"}

	for _, adminUserName := range adminUsers {
		adminUser := security.Adminuser{
			Name:        security.PtrString(adminUserName),
			AuthType:    security.PtrString("SAML"),
			AdminGroups: []string{"cloud-api-only"},
		}

		_, _, err := clients.SECURITY.AdminuserAPI.Create(context.Background()).
			Adminuser(adminUser).
			Execute()

		if err != nil {
			if strings.Contains(err.Error(), "already exists") {
				fmt.Printf("Admin user %q already exists, skipping creation\n", adminUserName)
				continue
			}
			return fmt.Errorf("failed to create admin user %q: %w", adminUserName, err)
		}

		fmt.Printf("Admin user %q created successfully\n", adminUserName)
	}

	// Create DNS views
	dnsViews := []string{"custom_dns_view"}

	for _, dnsViewName := range dnsViews {
		dnsViewBody := dns.View{
			Name: dns.PtrString(dnsViewName),
		}

		_, _, err := clients.DNS.ViewAPI.Create(context.Background()).
			View(dnsViewBody).
			Execute()

		if err != nil {
			if strings.Contains(err.Error(), "already exists") {
				fmt.Printf("DNS view %q already exists, skipping creation\n", dnsViewName)
				continue
			}
			return fmt.Errorf("failed to create DNS view %q: %w", dnsViewName, err)
		}

		fmt.Printf("DNS view %q created successfully\n", dnsViewName)
	}

	// Create DDNS principal cluster groups
	ddnsPrincipalClusterGroups := []string{"dynamic_update_grp_1", "dynamic_update_grp_2"}

	for _, groupName := range ddnsPrincipalClusterGroups {
		ddnsPrincipalClusterGroupBody := dns.DdnsPrincipalclusterGroup{
			Name: dns.PtrString(groupName),
		}

		_, _, err := clients.DNS.DdnsPrincipalclusterGroupAPI.Create(context.Background()).
			DdnsPrincipalclusterGroup(ddnsPrincipalClusterGroupBody).
			Execute()

		if err != nil {
			if strings.Contains(err.Error(), "already exists") {
				fmt.Printf("DDNS principal cluster group %q already exists, skipping creation\n", groupName)
				continue
			}
			return fmt.Errorf("failed to create DDNS principal cluster group %q: %w", groupName, err)
		}

		fmt.Printf("DDNS principal cluster group %q created successfully\n", groupName)
	}

	// Create DNS64 groups
	dns64Groups := []struct {
		name   string
		prefix string
	}{
		{name: "dns64_group", prefix: "64:FF9B::/96"},
	}

	for _, g := range dns64Groups {
		dns64GroupBody := dns.Dns64group{
			Name:   dns.PtrString(g.name),
			Prefix: dns.PtrString(g.prefix),
		}

		_, _, err := clients.DNS.Dns64groupAPI.Create(context.Background()).
			Dns64group(dns64GroupBody).
			Execute()

		if err != nil {
			if strings.Contains(err.Error(), "already exists") {
				fmt.Printf("DNS64 group %q already exists, skipping creation\n", g.name)
				continue
			}
			return fmt.Errorf("failed to create DNS64 group %q: %w", g.name, err)
		}

		fmt.Printf("DNS64 group %q created successfully\n", g.name)
	}

	// Create stub NS groups
	stubNSGroups := []string{"stub_ns_group1", "stub_ns_group2"}

	for _, groupName := range stubNSGroups {
		stubMemberBody := dns.NsgroupStubmember{
			Name: dns.PtrString(groupName),
			StubMembers: []dns.NsgroupStubmemberStubMembers{
				{
					Name: dns.PtrString(memberHostname),
				},
			},
		}

		_, _, err := clients.DNS.NsgroupStubmemberAPI.Create(context.Background()).
			NsgroupStubmember(stubMemberBody).
			Execute()

		if err != nil {
			if strings.Contains(err.Error(), "already exists") {
				fmt.Printf("Stub NS group %q already exists, skipping creation\n", groupName)
				continue
			}
			return fmt.Errorf("failed to create stub NS group %q: %w", groupName, err)
		}

		fmt.Printf("Stub NS group %q created successfully\n", groupName)
	}

	// Create forward stub servers
	forwardStubServers := []string{"ensg1", "ensg2"}

	for _, serverName := range forwardStubServers {
		forwardStubServerBody := dns.NsgroupForwardstubserver{
			Name: dns.PtrString(serverName),
			ExternalServers: []dns.NsgroupForwardstubserverExternalServers{
				{
					Name:    dns.PtrString("example.com"),
					Address: dns.PtrString("2.3.4.4"),
				},
			},
		}

		_, _, err := clients.DNS.NsgroupForwardstubserverAPI.Create(context.Background()).
			NsgroupForwardstubserver(forwardStubServerBody).
			Execute()

		if err != nil {
			if strings.Contains(err.Error(), "already exists") {
				fmt.Printf("Forward stub server %q already exists, skipping creation\n", serverName)
				continue
			}
			return fmt.Errorf("failed to create forward stub server %q: %w", serverName, err)
		}

		fmt.Printf("Forward stub server %q created successfully\n", serverName)
	}

	// Create zone auth
	zoneAuthBody := dns.ZoneAuth{
		Fqdn: dns.PtrString("example.com"),
		View: dns.PtrString("default"),
	}

	_, _, err = clients.DNS.ZoneAuthAPI.Create(context.Background()).
		ZoneAuth(zoneAuthBody).
		Execute()

	if err != nil {
		if strings.Contains(err.Error(), "exists") {
			fmt.Printf("Zone auth %q already exists, skipping creation\n", "example.com")
		} else {
			return fmt.Errorf("failed to create zone auth %q: %w", "example.com", err)
		}
	} else {
		fmt.Printf("Zone auth %q created successfully\n", "example.com")
	}

	// Create IPv4 reverse mapping zone auth
	zoneAuthIPv4Body := dns.ZoneAuth{
		Fqdn:       dns.PtrString("192.168.10.0/24"),
		View:       dns.PtrString("default"),
		ZoneFormat: dns.PtrString("IPV4"),
	}

	_, _, err = clients.DNS.ZoneAuthAPI.Create(context.Background()).
		ZoneAuth(zoneAuthIPv4Body).
		Execute()

	if err != nil {
		if strings.Contains(err.Error(), "exists") {
			fmt.Printf("Zone auth %q already exists, skipping creation\n", "192.168.10.0/24")
		} else {
			return fmt.Errorf("failed to create zone auth %q: %w", "192.168.10.0/24", err)
		}
	} else {
		fmt.Printf("Zone auth %q created successfully\n", "192.168.10.0/24")
	}

	// Create IPv6 reverse mapping zone auth
	ipv6ZoneAuthBody := dns.ZoneAuth{
		Fqdn:       dns.PtrString("2001::/64"),
		View:       dns.PtrString("default"),
		ZoneFormat: dns.PtrString("IPV6"),
	}

	_, _, err = clients.DNS.ZoneAuthAPI.Create(context.Background()).
		ZoneAuth(ipv6ZoneAuthBody).
		Execute()

	if err != nil {
		if strings.Contains(err.Error(), "exists") {
			fmt.Printf("Zone auth %q already exists, skipping creation\n", "2001::/64")
		} else {
			return fmt.Errorf("failed to create zone auth %q: %w", "2001::/64", err)
		}
	} else {
		fmt.Printf("Zone auth %q created successfully\n", "2001::/64")
	}

	// Create RPZ zone
	zoneRpBody := dns.ZoneRp{
		Fqdn: dns.PtrString("test-rpz.com"),
		View: dns.PtrString("default"),
	}

	_, _, err = clients.DNS.ZoneRpAPI.Create(context.Background()).
		ZoneRp(zoneRpBody).
		Execute()

	if err != nil {
		if strings.Contains(err.Error(), "exists") {
			fmt.Printf("Zone RP %q already exists, skipping creation\n", "test-rpz.com")
		} else {
			return fmt.Errorf("failed to create zone RP %q: %w", "test-rpz.com", err)
		}
	} else {
		fmt.Printf("Zone RP %q created successfully\n", "test-rpz.com")
	}

	// Create shared record group zones
	srgZones := []string{"tf-srg-zone-1.com", "tf-srg-zone-2.com"}
	for _, fqdn := range srgZones {
		srgZoneBody := dns.ZoneAuth{
			Fqdn: dns.PtrString(fqdn),
			View: dns.PtrString("default"),
		}
		_, _, err = clients.DNS.ZoneAuthAPI.Create(context.Background()).
			ZoneAuth(srgZoneBody).
			Execute()
		if err != nil {
			if strings.Contains(err.Error(), "exists") {
				fmt.Printf("Zone auth %q already exists, skipping creation\n", fqdn)
			} else {
				return fmt.Errorf("failed to create zone auth %q: %w", fqdn, err)
			}
		} else {
			fmt.Printf("Zone auth %q created successfully\n", fqdn)
		}
	}

	// Create blacklist rulesets
	blacklistRulesets := []string{"blacklist_ruleset_1", "blacklist_ruleset_2"}
	for _, name := range blacklistRulesets {
		rulesetBody := misc.Ruleset{
			Name: misc.PtrString(name),
			Type: misc.PtrString("BLACKLIST"),
		}
		_, _, err = clients.MISC.RulesetAPI.Create(context.Background()).
			Ruleset(rulesetBody).
			Execute()
		if err != nil {
			if strings.Contains(err.Error(), "exists") || strings.Contains(err.Error(), "already exists") {
				fmt.Printf("Ruleset %q already exists, skipping creation\n", name)
			} else {
				return fmt.Errorf("failed to create blacklist ruleset %q: %w", name, err)
			}
		} else {
			fmt.Printf("Blacklist ruleset %q created successfully\n", name)
		}
	}

	// Create nxdomain rulesets
	nxdomainRulesets := []string{"nxdomain_ruleset_1", "nxdomain_ruleset_2"}
	for _, name := range nxdomainRulesets {
		rulesetBody := misc.Ruleset{
			Name: misc.PtrString(name),
			Type: misc.PtrString("NXDOMAIN"),
		}
		_, _, err = clients.MISC.RulesetAPI.Create(context.Background()).
			Ruleset(rulesetBody).
			Execute()
		if err != nil {
			if strings.Contains(err.Error(), "exists") || strings.Contains(err.Error(), "already exists") {
				fmt.Printf("Ruleset %q already exists, skipping creation\n", name)
			} else {
				return fmt.Errorf("failed to create nxdomain ruleset %q: %w", name, err)
			}
		} else {
			fmt.Printf("NXDOMAIN ruleset %q created successfully\n", name)
		}
	}

	// Create DHCP failovers
	failovers := []struct {
		name      string
		secondary string
	}{
		{name: "example_failover_association", secondary: "250.251.2.2"},
		{name: "example_failover_association1", secondary: "250.252.2.2"},
	}

	for _, f := range failovers {
		failoverBody := dhcp.Dhcpfailover{
			Name:                dhcp.PtrString(f.name),
			PrimaryServerType:   dhcp.PtrString("GRID"),
			Primary:             dhcp.PtrString(memberHostname),
			SecondaryServerType: dhcp.PtrString("EXTERNAL"),
			Secondary:           dhcp.PtrString(f.secondary),
		}

		_, _, err := clients.DHCP.DhcpfailoverAPI.Create(context.Background()).
			Dhcpfailover(failoverBody).
			Execute()

		if err != nil {
			if strings.Contains(err.Error(), "already exists") {
				fmt.Printf("DHCP failover %q already exists, skipping creation\n", f.name)
				continue
			}
			return fmt.Errorf("failed to create DHCP failover %q: %w", f.name, err)
		}

		fmt.Printf("DHCP failover %q created successfully\n", f.name)
	}

	// Create Microsoft servers
	msServers := []struct {
		address     string
		dnsView     *string
		networkView string
	}{
		{address: "10.10.10.10", dnsView: microsoft.PtrString("default"), networkView: "default"},
		{address: "10.10.10.11", dnsView: microsoft.PtrString("default"), networkView: "default"},
		{address: "example_server", dnsView: microsoft.PtrString("default"), networkView: "default"},
		{address: "ms_example_server", dnsView: nil, networkView: "ms_server"},
		{address: "ms_example_server2", dnsView: nil, networkView: "ms_server2"},
	}

	for _, server := range msServers {
		msServerBody := microsoft.Msserver{
			Address:                 microsoft.PtrString(server.address),
			DnsView:                 server.dnsView,
			NetworkView:             microsoft.PtrString(server.networkView),
			ReadOnly:                microsoft.PtrBool(false),
			SynchronizationMinDelay: microsoft.PtrInt64(2),
			LoginName:               microsoft.PtrString("admin"),
		}

		_, _, err := clients.MICROSOFT.MsserverAPI.Create(context.Background()).
			Msserver(msServerBody).
			Execute()

		if err != nil {
			if strings.Contains(err.Error(), "is using") || strings.Contains(err.Error(), "is already using") {
				fmt.Printf("Microsoft server %q already exists, skipping creation\n", server.address)
				continue
			}
			return fmt.Errorf("failed to create Microsoft server %q: %w", server.address, err)
		}

		fmt.Printf("Microsoft server %q created successfully\n", server.address)
	}

	// Create grid members

	networkSettingAddress6 := fmt.Sprintf("2001:db8:%x:%x::%x", acctest.RandomNumber(65535), acctest.RandomNumber(65535), acctest.RandomNumber(65535))
	ipv6SettingVal2 := grid.MemberIpv6Setting{
		AutoRouterConfigEnabled: grid.PtrBool(false),
		Dscp:                    grid.PtrInt64(0),
		Enabled:                 grid.PtrBool(true),
		UseDscp:                 grid.PtrBool(false),
		VirtualIp:               grid.PtrString(networkSettingAddress6),
		Gateway:                 grid.PtrString("2001::1"),
		Primary:                 grid.PtrBool(true),
		CidrPrefix:              grid.PtrInt64(8),
	}

	ipv6SettingVal3 := grid.MemberIpv6Setting{
		AutoRouterConfigEnabled: grid.PtrBool(false),
		Dscp:                    grid.PtrInt64(0),
		Enabled:                 grid.PtrBool(false),
		UseDscp:                 grid.PtrBool(false),
		Primary:                 grid.PtrBool(true),
	}

	members := []struct {
		hostName        string
		vipAddress      string
		configAddrType  string
		ipv6SettingVal  grid.MemberIpv6Setting
		masterCandidate bool
	}{
		{hostName: "infoblox.member2", vipAddress: "172.28.38.251", configAddrType: "BOTH", ipv6SettingVal: ipv6SettingVal2, masterCandidate: true},
		{hostName: "infoblox.member3", vipAddress: "172.28.38.252", configAddrType: "IPV4", ipv6SettingVal: ipv6SettingVal3, masterCandidate: false},
	}

	for _, member := range members {
		memberBody := grid.Member{
			HostName:                 grid.PtrString(member.hostName),
			ConfigAddrType:           grid.PtrString(member.configAddrType),
			Platform:                 grid.PtrString("VNIOS"),
			ServiceTypeConfiguration: grid.PtrString("ALL_V4"),
			MasterCandidate:          grid.PtrBool(member.masterCandidate),
			VipSetting: &grid.MemberVipSetting{
				Address:    grid.PtrString(member.vipAddress),
				Gateway:    grid.PtrString("172.28.38.1"),
				SubnetMask: grid.PtrString("255.255.254.0"),
				Primary:    grid.PtrBool(true),
				Dscp:       grid.PtrInt64(0),
				UseDscp:    grid.PtrBool(false),
			},
			Ipv6Setting: &member.ipv6SettingVal,
		}

		_, _, err := clients.GRID.MemberAPI.Create(context.Background()).
			Member(memberBody).
			Execute()

		if err != nil {
			if strings.Contains(err.Error(), "is already using") {
				fmt.Printf("Member %q already exists, skipping creation\n", member.hostName)
				continue
			}
			return fmt.Errorf("failed to create member %q: %w", member.hostName, err)
		}

		fmt.Printf("Member %q created successfully\n", member.hostName)
	}

	// Ensure upgrade schedule has required start time and group upgrade times.
	upgradeScheduleResp, _, err := clients.GRID.UpgradescheduleAPI.List(context.Background()).
		ReturnAsObject(1).
		ReturnFieldsPlus("start_time,upgrade_groups").
		Execute()
	if err != nil {
		return fmt.Errorf("failed to list upgrade schedule: %w", err)
	}
	if upgradeScheduleResp.ListUpgradescheduleResponseObject == nil || len(upgradeScheduleResp.ListUpgradescheduleResponseObject.Result) == 0 {
		return fmt.Errorf("upgrade schedule list returned no results")
	}

	upgradeSchedule := upgradeScheduleResp.ListUpgradescheduleResponseObject.Result[0]
	if upgradeSchedule.Ref == nil || *upgradeSchedule.Ref == "" {
		return fmt.Errorf("upgrade schedule ref is empty")
	}

	const defaultUpgradeScheduleStartTime int64 = 1829280600
	const defaultUpgradeGroupTime int64 = 1830144600

	upgradeScheduleBody := grid.Upgradeschedule{
		Active: dns.PtrBool(false),
	}

	upgradeScheduleBody.StartTime = grid.PtrInt64(defaultUpgradeScheduleStartTime)

	if upgradeSchedule.HasUpgradeGroups() {
		scheduleUpgradeGroups := upgradeSchedule.GetUpgradeGroups()
		hasMissingGroupUpgradeTime := false

		var upgradeGroupsForBody []grid.UpgradescheduleUpgradeGroups
		for _, ug := range scheduleUpgradeGroups {
			groupBody := grid.UpgradescheduleUpgradeGroups{}
			if ug.HasName() {
				groupBody.SetName(ug.GetName())
			}
			groupBody.SetUpgradeTime(defaultUpgradeGroupTime)
			hasMissingGroupUpgradeTime = true
			upgradeGroupsForBody = append(upgradeGroupsForBody, groupBody)
		}

		if hasMissingGroupUpgradeTime {
			upgradeScheduleBody.SetUpgradeGroups(upgradeGroupsForBody)
		}
	}

	upgradeScheduleRef := core.ExtractNIOSRef(*upgradeSchedule.Ref)
	_, _, err = clients.GRID.UpgradescheduleAPI.Update(context.Background(), upgradeScheduleRef).
		Upgradeschedule(upgradeScheduleBody).
		Execute()
	if err != nil {
		return fmt.Errorf("failed to update upgrade schedule %q: %w", *upgradeSchedule.Ref, err)
	}

	fmt.Printf("Upgrade schedule %q updated\n", *upgradeSchedule.Ref)

	// Ensure distribution schedule has required start time and group upgrade times.
	distributionScheduleResp, _, err := clients.GRID.DistributionscheduleAPI.List(context.Background()).
		ReturnAsObject(1).
		ReturnFieldsPlus("start_time,upgrade_groups").
		Execute()
	if err != nil {
		return fmt.Errorf("failed to list distribution schedule: %w", err)
	}
	if distributionScheduleResp.ListDistributionscheduleResponseObject == nil || len(distributionScheduleResp.ListDistributionscheduleResponseObject.Result) == 0 {
		return fmt.Errorf("distribution schedule list returned no results")
	}

	distributionSchedule := distributionScheduleResp.ListDistributionscheduleResponseObject.Result[0]
	if distributionSchedule.Ref == nil || *distributionSchedule.Ref == "" {
		return fmt.Errorf("distribution schedule ref is empty")
	}

	defaultDistributionScheduleStartTime := time.Now().Add(12 * time.Hour).Unix()
	defaultDistributionScheduleGroupTime := time.Now().Add(18 * time.Hour).Unix()

	distributionScheduleBody := grid.Distributionschedule{
		Active: dns.PtrBool(false),
	}

	distributionScheduleBody.StartTime = grid.PtrInt64(defaultDistributionScheduleStartTime)

	if distributionSchedule.HasUpgradeGroups() {
		scheduleUpgradeGroups := distributionSchedule.GetUpgradeGroups()
		hasMissingGroupUpgradeTime := false

		var upgradeGroupsForBody []grid.DistributionscheduleUpgradeGroups
		for _, ug := range scheduleUpgradeGroups {
			groupBody := grid.DistributionscheduleUpgradeGroups{}
			if ug.HasName() {
				groupBody.SetName(ug.GetName())
			}
			groupBody.SetDistributionTime(defaultDistributionScheduleGroupTime)
			hasMissingGroupUpgradeTime = true
			upgradeGroupsForBody = append(upgradeGroupsForBody, groupBody)
		}

		if hasMissingGroupUpgradeTime {
			distributionScheduleBody.SetUpgradeGroups(upgradeGroupsForBody)
		}
	}

	distributionScheduleRef := core.ExtractNIOSRef(*distributionSchedule.Ref)
	_, _, err = clients.GRID.DistributionscheduleAPI.Update(context.Background(), distributionScheduleRef).
		Distributionschedule(distributionScheduleBody).
		Execute()
	if err != nil {
		return fmt.Errorf("failed to update distribution schedule %q: %w", *distributionSchedule.Ref, err)
	}

	fmt.Printf("Distribution schedule %q updated\n", *distributionSchedule.Ref)

	// Create upgrade dependent groups
	upgradeGroups := []string{"example_upgrade_dependent_group1", "example_upgrade_dependent_group2"}

	for _, groupName := range upgradeGroups {
		upgradeGroupBody := grid.Upgradegroup{
			Name: grid.PtrString(groupName),
		}

		_, _, err := clients.GRID.UpgradegroupAPI.Create(context.Background()).
			Upgradegroup(upgradeGroupBody).
			Execute()

		if err != nil {
			if strings.Contains(err.Error(), "already exists") {
				fmt.Printf("Upgrade dependent group %q already exists, skipping creation\n", groupName)
				continue
			}
			return fmt.Errorf("failed to create upgrade dependent group %q: %w", groupName, err)
		}

		fmt.Printf("Upgrade dependent group %q created successfully\n", groupName)
	}

	// Create RIR organizations
	rirExtAttrs := map[string]rir.ExtAttrs{
		"RIPE Admin Contact":     {Value: "ib-contact"},
		"RIPE Country":           {Value: "United Kingdom (GB)"},
		"RIPE Technical Contact": {Value: "TEST123-IB"},
		"RIPE Email":             {Value: "support@infoblox.com"},
	}

	rirOrgs := []struct {
		id   string
		name string
	}{
		{id: "ORG-CB11-INTEST", name: "rir-org-test1"},
		{id: "ORG-CB12-INTEST", name: "rir-org-test2"},
	}

	for _, org := range rirOrgs {
		rirOrgBody := rir.RirOrganization{
			Id:          rir.PtrString(org.id),
			Name:        rir.PtrString(org.name),
			Maintainer:  rir.PtrString("infoblox"),
			Password:    rir.PtrString("test-pass"),
			SenderEmail: rir.PtrString("support@infoblox.com"),
			Rir:         rir.PtrString("RIPE"),
			ExtAttrs:    &rirExtAttrs,
		}

		_, _, err := clients.RIR.RirOrganizationAPI.Create(context.Background()).
			RirOrganization(rirOrgBody).
			Execute()
		if err != nil {
			if strings.Contains(err.Error(), "already exists") {
				fmt.Printf("RIR organization %q already exists, skipping creation\n", org.id)
				continue
			}
			return fmt.Errorf("failed to create RIR organization %q: %w", org.id, err)
		}

		fmt.Printf("RIR organization %q created successfully\n", org.id)
	}

	// Create Active Directory auth services
	adAuthServices := []struct {
		name      string
		refEnvVar string
	}{
		{name: "active_dir", refEnvVar: "NIOS_AD_AUTH_SERVICE_ACTIVE_DIR_REF"},
		{name: "active_dir_test", refEnvVar: "NIOS_AD_AUTH_SERVICE_ACTIVE_DIR_TEST_REF"},
	}

	adAuthServiceRefs := make(map[string]string)

	for _, ad := range adAuthServices {
		adAuthServiceBody := security.AdAuthService{
			Name:     security.PtrString(ad.name),
			AdDomain: security.PtrString("example.com"),
			DomainControllers: []security.AdAuthServiceDomainControllers{
				{
					FqdnOrIp: security.PtrString("1.1.1.1"),
					Disabled: security.PtrBool(false),
					AuthPort: security.PtrInt64(389),
				},
			},
			Timeout: security.PtrInt64(5),
		}

		resp, _, err := clients.SECURITY.AdAuthServiceAPI.Create(context.Background()).
			AdAuthService(adAuthServiceBody).
			Execute()

		if err != nil {
			if strings.Contains(err.Error(), "already exists") {
				filtersAdAuth := map[string]interface{}{
					"name": ad.name,
				}

				existingResp, _, listErr := clients.SECURITY.AdAuthServiceAPI.List(context.Background()).
					ReturnAsObject(1).
					Filters(filtersAdAuth).
					ReturnFieldsPlus("name").
					Execute()
				if listErr != nil {
					return fmt.Errorf("failed to list AD auth services to capture existing ref: %w", listErr)
				}

				if existingResp.ListAdAuthServiceResponseObject == nil || len(existingResp.ListAdAuthServiceResponseObject.Result) == 0 {
					return fmt.Errorf("AD auth service %q already exists but no matching object was returned", ad.name)
				}

				var existingAdRef string
				for _, existingAd := range existingResp.ListAdAuthServiceResponseObject.Result {
					if existingAd.Name != nil && *existingAd.Name == ad.name {
						existingAdRef = existingAd.GetRef()
						break
					}
				}

				if existingAdRef == "" {
					return fmt.Errorf("AD auth service %q already exists but ref could not be resolved", ad.name)
				}

				adAuthServiceRefs[ad.name] = existingAdRef
				if err := writePipelineEnvVar(ad.refEnvVar, existingAdRef); err != nil {
					return fmt.Errorf("failed to write env var %s for existing AD auth service: %w", ad.refEnvVar, err)
				}

				fmt.Printf("Captured ref for existing Active Directory auth service %q: %s (env: %s)\n", ad.name, existingAdRef, ad.refEnvVar)
				fmt.Printf("Active Directory auth service %q already exists, skipping creation\n", ad.name)
				continue
			}
			return fmt.Errorf("failed to create Active Directory auth service %q: %w", ad.name, err)
		}

		// Extract ref from response (plain string in default mode, nested object in _return_as_object mode)
		var adRef string
		if resp.String != nil {
			adRef = *resp.String
		} else if resp.CreateAdAuthServiceResponseAsObject != nil && resp.CreateAdAuthServiceResponseAsObject.Result != nil {
			adRef = resp.CreateAdAuthServiceResponseAsObject.Result.GetRef()
		}
		if adRef != "" {
			adAuthServiceRefs[ad.name] = adRef
		}

		if err := writePipelineEnvVar(ad.refEnvVar, adRef); err != nil {
			return fmt.Errorf("failed to write env var %s: %w", ad.refEnvVar, err)
		}

		fmt.Printf("Active Directory auth service %q created successfully (ref: %s, env: %s)\n", ad.name, adRef, ad.refEnvVar)
	}

	// TODO : Remote Lookup Services need to be removed to run Member UTs
	// Update auth policy to append the first AD auth service ref
	// Only proceed if at least one new ref was captured
	//firstADRef := adAuthServiceRefs["active_dir"]
	//if firstADRef != "" {
	//	// GET current auth policy
	//	listResp, _, err := clients.SECURITY.AuthpolicyAPI.List(context.Background()).
	//		ReturnAsObject(1).
	//		ReturnFieldsPlus("auth_services").
	//		Execute()
	//	if err != nil {
	//		return fmt.Errorf("failed to list auth policy: %w", err)
	//	}
	//
	//	// Resolve the auth policy slice and ref from either response shape
	//	var authPolicies []security.Authpolicy
	//	if listResp.ListAuthpolicyResponseObject != nil {
	//		authPolicies = listResp.ListAuthpolicyResponseObject.Result
	//	} else if listResp.ArrayOfAuthpolicy != nil {
	//		authPolicies = *listResp.ArrayOfAuthpolicy
	//	}
	//
	//	if len(authPolicies) == 0 {
	//		return fmt.Errorf("auth policy list returned no results")
	//	}
	//
	//	authPolicyRef := authPolicies[0].Ref
	//
	//	if authPolicyRef == nil || *authPolicyRef == "" {
	//		return fmt.Errorf("auth policy ref is empty")
	//	}
	//
	//	extractedAuthPolicyRef := core.ExtractNIOSRef(*authPolicyRef)
	//
	//	// Append first AD auth service ref to existing auth services
	//	authServices := authPolicies[0].GetAuthServices()
	//	authServices = append(authServices, firstADRef)
	//
	//	authPolicyBody := security.Authpolicy{
	//		AuthServices: authServices,
	//	}
	//
	//	_, _, err = clients.SECURITY.AuthpolicyAPI.Update(context.Background(), extractedAuthPolicyRef).
	//		Authpolicy(authPolicyBody).
	//		Execute()
	//	if err != nil {
	//		return fmt.Errorf("failed to update auth policy %q: %w", *authPolicyRef, err)
	//	}
	//
	//	fmt.Printf("Auth policy %q updated with AD auth service ref %q\n", *authPolicyRef, firstADRef)
	//}

	// Update MemberFileDistribution to append a new TFTP ACL entry
	memberFiledistResp, _, err := clients.GRID.MemberFiledistributionAPI.List(context.Background()).ReturnAsObject(1).Execute()
	if err != nil {
		return fmt.Errorf("failed to list member file distribution: %w", err)
	}

	// Resolve the member file distribution slice from either response shape
	var memberFiledistributions []grid.MemberFiledistribution
	if memberFiledistResp.ListMemberFiledistributionResponseObject != nil {
		memberFiledistributions = memberFiledistResp.ListMemberFiledistributionResponseObject.Result
	} else if memberFiledistResp.ArrayOfMemberFiledistribution != nil {
		memberFiledistributions = *memberFiledistResp.ArrayOfMemberFiledistribution
	}

	if len(memberFiledistributions) > 0 {
		memberRef := memberFiledistributions[0].Ref
		if memberRef != nil && *memberRef != "" {
			// Get existing TFTP ACLs
			memberTftpACL := memberFiledistributions[0].GetTftpAcls()

			// Create new TFTP ACL entry
			newTftpACL := grid.MemberFiledistributionTftpAcls{
				Address:    grid.PtrString("Any"),
				Permission: grid.PtrString("ALLOW"),
			}
			memberTftpACL = append(memberTftpACL, newTftpACL)

			memberFileDistributionBody := grid.MemberFiledistribution{
				TftpAcls: memberTftpACL,
			}

			extractedMemberRef := core.ExtractNIOSRef(*memberRef)

			_, _, err := clients.GRID.MemberFiledistributionAPI.Update(context.Background(), extractedMemberRef).
				MemberFiledistribution(memberFileDistributionBody).
				Execute()
			if err != nil {
				return fmt.Errorf("failed to update member file distribution %q: %w", *memberRef, err)
			}

			fmt.Printf("Member file distribution %q updated with new TFTP ACL (Address: Any, Permission: ALLOW)\n", *memberRef)
		}
	}

	// Update RPZ logging in DNS grid properties
	gridDNSResp, _, err := clients.GRID.GridDnsAPI.List(context.Background()).
		ReturnAsObject(1).
		ReturnFields("logging_categories").
		Execute()
	if err != nil {
		return fmt.Errorf("failed to list grid DNS properties: %w", err)
	}

	if gridDNSResp.ListGridDnsResponseObject == nil || len(gridDNSResp.ListGridDnsResponseObject.Result) == 0 {
		return fmt.Errorf("grid DNS properties list returned no results")
	}

	gridDNSRef := gridDNSResp.ListGridDnsResponseObject.Result[0].Ref
	if gridDNSRef == nil || *gridDNSRef == "" {
		return fmt.Errorf("grid DNS properties ref is empty")
	}

	dnsGridPropertiesRef := core.ExtractNIOSRef(*gridDNSRef)
	rpzLoggingLevel := gridDNSResp.ListGridDnsResponseObject.Result[0].LoggingCategories.LogRpz
	allowRecursiveQuery := gridDNSResp.ListGridDnsResponseObject.Result[0].AllowRecursiveQuery

	if rpzLoggingLevel == nil || !*rpzLoggingLevel || allowRecursiveQuery == nil || !*allowRecursiveQuery {
		fmt.Printf("RPZ logging level is %v, updating to true\n", rpzLoggingLevel)
		dnsGridUpdateBody := grid.GridDns{
			LoggingCategories: &grid.GridDnsLoggingCategories{
				LogRpz: grid.PtrBool(true),
			},
			AllowRecursiveQuery: grid.PtrBool(true),
		}

		_, _, err = clients.GRID.GridDnsAPI.Update(context.Background(), dnsGridPropertiesRef).
			GridDns(dnsGridUpdateBody).
			Execute()
		if err != nil {
			return fmt.Errorf("failed to update RPZ logging in DNS grid properties %q: %w", dnsGridPropertiesRef, err)
		}

		fmt.Printf("RPZ logging enabled for DNS grid properties %q\n", dnsGridPropertiesRef)
	}

	// Create notification REST endpoint and persist its ref
	notificationRestEndpointBody := notification.NotificationRestEndpoint{
		Name:               notification.PtrString("notification_rest_endpoint999"),
		OutboundMemberType: notification.PtrString("GM"),
		Uri:                notification.PtrString("https://example.com"),
	}

	notificationResp, _, err := clients.NOTIFICATION.NotificationRestEndpointAPI.Create(context.Background()).
		ReturnAsObject(1).
		ReturnFieldsPlus("name").
		NotificationRestEndpoint(notificationRestEndpointBody).
		Execute()
	if err != nil {
		if strings.Contains(err.Error(), "already exists") {
			filtersNotifRest := map[string]interface{}{
				"name": *notificationRestEndpointBody.Name,
			}
			// if the notification REST endpoint already exists, we can attempt to capture its ref for later use
			existingResp, _, listErr := clients.NOTIFICATION.NotificationRestEndpointAPI.List(context.Background()).
				ReturnAsObject(1).
				Filters(filtersNotifRest).
				ReturnFieldsPlus("name").
				Execute()
			if listErr != nil {
				return fmt.Errorf("failed to list notification REST endpoints to capture existing ref: %w", listErr)
			}

			if existingResp.ListNotificationRestEndpointResponseObject != nil &&
				len(existingResp.ListNotificationRestEndpointResponseObject.Result) > 0 {
				for _, existingEndpoint := range existingResp.ListNotificationRestEndpointResponseObject.Result {
					if existingEndpoint.Name != nil && *existingEndpoint.Name == *notificationRestEndpointBody.Name {
						existingRef := existingEndpoint.GetRef()
						if err := writePipelineEnvVar("NIOS_NOTIFICATION_REST_ENDPOINT_REF", existingRef); err != nil {
							return fmt.Errorf("failed to write NIOS_NOTIFICATION_REST_ENDPOINT_REF for existing endpoint: %w", err)
						}
						fmt.Printf("Captured ref for existing notification REST endpoint %q: %s\n",
							*notificationRestEndpointBody.Name, existingRef)
						break
					}
				}
			}
			fmt.Printf("Notification REST endpoint %q already exists, skipping creation\n", *notificationRestEndpointBody.Name)
		} else {
			return fmt.Errorf("failed to create notification REST endpoint %q: %w", *notificationRestEndpointBody.Name, err)
		}
	} else {
		if notificationResp.CreateNotificationRestEndpointResponseAsObject == nil ||
			notificationResp.CreateNotificationRestEndpointResponseAsObject.Result == nil ||
			notificationResp.CreateNotificationRestEndpointResponseAsObject.Result.Ref == nil {
			return fmt.Errorf("notification REST endpoint create response missing ref")
		}

		notifRestEndpointRef := *notificationResp.CreateNotificationRestEndpointResponseAsObject.Result.Ref
		if err := writePipelineEnvVar("NIOS_NOTIFICATION_REST_ENDPOINT_REF", notifRestEndpointRef); err != nil {
			return fmt.Errorf("failed to write NIOS_NOTIFICATION_REST_ENDPOINT_REF: %w", err)
		}

		fmt.Printf("Notification REST endpoint %q created successfully (ref: %s)\n",
			*notificationRestEndpointBody.Name, notifRestEndpointRef)
	}

	// Create syslog endpoint and persist its ref.
	syslogEndpointBody := misc.SyslogEndpoint{
		Name:               misc.PtrString("syslogendpoint123"),
		OutboundMemberType: misc.PtrString("GM"),
		SyslogServers: []misc.SyslogEndpointSyslogServers{
			{
				Address:        misc.PtrString("127.0.0.1"),
				Port:           misc.PtrInt64(514),
				ConnectionType: misc.PtrString("udp"),
				Format:         misc.PtrString("formatted"),
			},
		},
	}

	syslogResp, _, err := clients.MISC.SyslogEndpointAPI.Create(context.Background()).
		SyslogEndpoint(syslogEndpointBody).
		ReturnAsObject(1).
		ReturnFieldsPlus("name").
		Execute()
	if err != nil {
		if strings.Contains(err.Error(), "already exists") {
			filtersSyslog := map[string]interface{}{
				"name": *syslogEndpointBody.Name,
			}

			existingResp, _, listErr := clients.MISC.SyslogEndpointAPI.List(context.Background()).
				ReturnAsObject(1).
				Filters(filtersSyslog).
				ReturnFieldsPlus("name").
				Execute()
			if listErr != nil {
				return fmt.Errorf("failed to list syslog endpoints to capture existing ref: %w", listErr)
			}

			if existingResp.ListSyslogEndpointResponseObject != nil && len(existingResp.ListSyslogEndpointResponseObject.Result) > 0 {
				for _, existingEndpoint := range existingResp.ListSyslogEndpointResponseObject.Result {
					if existingEndpoint.Name != nil && *existingEndpoint.Name == *syslogEndpointBody.Name {
						existingRef := existingEndpoint.GetRef()
						if err := writePipelineEnvVar("NIOS_SYSLOG_ENDPOINT_REF", existingRef); err != nil {
							return fmt.Errorf("failed to write NIOS_SYSLOG_ENDPOINT_REF for existing endpoint: %w", err)
						}
						fmt.Printf("Captured ref for existing syslog endpoint %q: %s\n", *syslogEndpointBody.Name, existingRef)
						break
					}
				}
			}

			fmt.Printf("Syslog endpoint %q already exists, skipping creation\n", *syslogEndpointBody.Name)
		} else {
			return fmt.Errorf("failed to create syslog endpoint %q: %w", *syslogEndpointBody.Name, err)
		}
	} else {
		if syslogResp.CreateSyslogEndpointResponseAsObject == nil || syslogResp.CreateSyslogEndpointResponseAsObject.Result == nil ||
			syslogResp.CreateSyslogEndpointResponseAsObject.Result.Ref == nil {
			return fmt.Errorf("syslog endpoint create response missing ref")
		}

		syslogEndpointRef := *syslogResp.CreateSyslogEndpointResponseAsObject.Result.Ref
		if err := writePipelineEnvVar("NIOS_SYSLOG_ENDPOINT_REF", syslogEndpointRef); err != nil {
			return fmt.Errorf("failed to write NIOS_SYSLOG_ENDPOINT_REF: %w", err)
		}

		fmt.Printf("Syslog endpoint %q created successfully (ref: %s)\n", *syslogEndpointBody.Name, syslogEndpointRef)
	}

	// Enable Discovery service on the discovery member if NIOS_DISCOVERY_MEMBER_URL is set.
	discoveryMemberAddr := normalizeAddressFromEnv(os.Getenv("NIOS_DISCOVERY_MEMBER_URL"))
	if discoveryMemberAddr != "" {
		if clients.DISCOVERY == nil {
			return fmt.Errorf("preconfig: DISCOVERY client is required when NIOS_DISCOVERY_MEMBER_URL is set")
		}

		discoveryListResp, _, err := clients.DISCOVERY.DiscoveryMemberpropertiesAPI.List(context.Background()).
			ReturnAsObject(1).
			ReturnFieldsPlus("address,enable_service,is_sa").
			Execute()
		if err != nil {
			return fmt.Errorf("preconfig: failed to list discovery member properties: %w", err)
		}
		if discoveryListResp.ListDiscoveryMemberpropertiesResponseObject == nil {
			return fmt.Errorf("preconfig: discovery member properties list response object is nil")
		}

		discoveryMemberRef := ""
		for _, dm := range discoveryListResp.ListDiscoveryMemberpropertiesResponseObject.Result {
			if dm.Address != nil && strings.TrimSpace(*dm.Address) == discoveryMemberAddr {
				discoveryMemberRef = dm.GetRef()
				break
			}
		}

		if discoveryMemberRef == "" {
			fmt.Printf("No discovery member found with address %q, skipping discovery service enablement\n", discoveryMemberAddr)
		} else {
			extractedRef := core.ExtractNIOSRef(discoveryMemberRef)
			discoveryMemberBody := discovery.DiscoveryMemberproperties{
				EnableService: discovery.PtrBool(true),
				IsSa:          discovery.PtrBool(true),
			}

			_, _, err = clients.DISCOVERY.DiscoveryMemberpropertiesAPI.Update(context.Background(), extractedRef).
				DiscoveryMemberproperties(discoveryMemberBody).
				Execute()
			if err != nil {
				return fmt.Errorf("preconfig: failed to enable discovery service on member %q: %w", discoveryMemberRef, err)
			}

			fmt.Printf("Discovery service enabled (enable_service=true, is_sa=true) for member %q\n", discoveryMemberRef)
		}
	}

	return nil
}

func main() {
	host := strings.TrimSpace(firstNonEmpty(os.Getenv("NIOS_HOST_URL")))
	wapiVer := strings.TrimSpace(firstNonEmpty(os.Getenv("NIOS_WAPI_VERSION"), "v2.13.6"))
	username := strings.TrimSpace(firstNonEmpty(os.Getenv("NIOS_USERNAME")))
	password := strings.TrimSpace(firstNonEmpty(os.Getenv("NIOS_PASSWORD")))

	if host == "" || wapiVer == "" || username == "" || password == "" {
		fmt.Println("Missing required NIOS configuration. Ensure host, WAPI version, username, and password are set.")
		fmt.Println("Supported env vars: NIOS_HOST_URL (or NIOS_HOST), NIOS_WAPI_VERSION, NIOS_USERNAME, NIOS_PASSWORD")
		return
	}

	if !strings.HasPrefix(host, "https://") {
		fmt.Printf("Invalid NIOS host %q: must include https://\n", host)
		return
	}

	// Construct absolute path to ecosystem template file
	cwd, err := os.Getwd()
	if err != nil {
		fmt.Printf("Error getting current working directory: %v\n", err)
		return
	}

	pipelineEnvPath := filepath.Join(cwd, "pipeline_nios.env")
	f, err := os.Create(pipelineEnvPath)
	if err != nil {
		fmt.Printf("Error creating pipeline_nios.env: %v\n", err)
		return
	}
	pipelineEnvFile = f
	defer func() {
		_ = pipelineEnvFile.Close()
		pipelineEnvFile = nil
	}()

	apiClient := client.NewAPIClient(
		option.WithNIOSHostUrl(host),
		option.WithNIOSUsername(username),
		option.WithNIOSPassword(password),
		option.WithDebug(true),
	)

	hostnames, err := ResolveAndStoreGridHostnames(apiClient.GridAPI)
	if err != nil {
		fmt.Printf("Error resolving grid hostnames: %v\n", err)
		return
	}

	certMember := strings.TrimSpace(hostnames.MasterHostname)

	caCertUploadFilePath := filepath.Join(cwd, "internal/acctest/integration_tests/nios_security_certificate_authservice", "cert.pem")

	err = ConfigureCACertificates(host, wapiVer, username, password, "EAP_CA", certMember, caCertUploadFilePath, "NIOS_CA_CERT1_REF", "NIOS_CA_CERT1_SERIAL")
	if err != nil {
		fmt.Printf("Error configuring CA certificates: %v\n", err)
		return
	}
	fmt.Println("CA certificate 1 configured successfully")

	caCertUploadFilePath2 := filepath.Join(cwd, "internal/acctest/integration_tests/nios_notification_rest_endpoint", "dummy-bundle.pem")

	err = ConfigureCACertificates(host, wapiVer, username, password, "EAP_CA", certMember, caCertUploadFilePath2, "NIOS_CA_CERT2_REF", "NIOS_CA_CERT2_SERIAL")
	if err != nil {
		fmt.Printf("Error configuring CA certificates: %v\n", err)
		return
	}
	fmt.Println("CA certificate 2 configured successfully")

	ecosystemTemplatePath := filepath.Join(cwd, "internal/acctest/integration_tests/nios_ecosystem_templates", "Version5_DXL_Session_Template.json")

	if _, statErr := os.Stat(ecosystemTemplatePath); statErr == nil {
		err = ConfigureEcoSystemTemplates(host, wapiVer, username, password, ecosystemTemplatePath)
		if err != nil {
			fmt.Printf("Error uploading ecosystem templates: %v\n", err)
			return
		}
		fmt.Println("Ecosystem template 1 uploaded successfully")
	} else if os.IsNotExist(statErr) {
		fmt.Printf("Ecosystem template 1 not found at %s, skipping upload\n", ecosystemTemplatePath)
	} else {
		fmt.Printf("Error checking ecosystem template 1 path %s: %v\n", ecosystemTemplatePath, statErr)
		return
	}

	ecosystemTemplatePath2 := filepath.Join(cwd, "internal/acctest/integration_tests/nios_ecosystem_templates", "Version5_Syslog_Session_Template.json")

	if _, statErr := os.Stat(ecosystemTemplatePath2); statErr == nil {
		err = ConfigureEcoSystemTemplates(host, wapiVer, username, password, ecosystemTemplatePath2)
		if err != nil {
			fmt.Printf("Error uploading ecosystem templates: %v\n", err)
			return
		}
		fmt.Println("Ecosystem template 2 uploaded successfully")
	} else if os.IsNotExist(statErr) {
		fmt.Printf("Ecosystem template 2 not found at %s, skipping upload\n", ecosystemTemplatePath2)
	} else {
		fmt.Printf("Error checking ecosystem template 2 path %s: %v\n", ecosystemTemplatePath2, statErr)
		return
	}

	ecosystemTemplatePath3 := filepath.Join(cwd, "internal/acctest/integration_tests/nios_ecosystem_templates", "Version5_Syslog_Action_Template.json")

	if _, statErr := os.Stat(ecosystemTemplatePath3); statErr == nil {
		err = ConfigureEcoSystemTemplates(host, wapiVer, username, password, ecosystemTemplatePath3)
		if err != nil {
			fmt.Printf("Error uploading ecosystem templates: %v\n", err)
			return
		}
		fmt.Println("Ecosystem template 3 uploaded successfully")
	} else if os.IsNotExist(statErr) {
		fmt.Printf("Ecosystem template 3 not found at %s, skipping upload\n", ecosystemTemplatePath3)
	} else {
		fmt.Printf("Error checking ecosystem template 3 path %s: %v\n", ecosystemTemplatePath3, statErr)
		return
	}

	ecosystemTemplatePath4 := filepath.Join(cwd, "internal/acctest/integration_tests/nios_ecosystem_templates", "Version5_DXL_action_template.json")

	if _, statErr := os.Stat(ecosystemTemplatePath4); statErr == nil {
		err = ConfigureEcoSystemTemplates(host, wapiVer, username, password, ecosystemTemplatePath4)
		if err != nil {
			fmt.Printf("Error uploading ecosystem templates: %v\n", err)
			return
		}
		fmt.Println("Ecosystem template 4 uploaded successfully")
	} else if os.IsNotExist(statErr) {
		fmt.Printf("Ecosystem template 4 not found at %s, skipping upload\n", ecosystemTemplatePath4)
	} else {
		fmt.Printf("Error checking ecosystem template 4 path %s: %v\n", ecosystemTemplatePath4, statErr)
		return
	}

	ecosystemTemplatePath5 := filepath.Join(cwd, "internal/acctest/integration_tests/nios_ecosystem_templates", "Version5_DNS_Zone_and_Records.json")

	if _, statErr := os.Stat(ecosystemTemplatePath5); statErr == nil {
		err = ConfigureEcoSystemTemplates(host, wapiVer, username, password, ecosystemTemplatePath5)
		if err != nil {
			fmt.Printf("Error uploading ecosystem templates: %v\n", err)
			return
		}
		fmt.Println("Ecosystem template 5 uploaded successfully")
	} else if os.IsNotExist(statErr) {
		fmt.Printf("Ecosystem template 5 not found at %s, skipping upload\n", ecosystemTemplatePath5)
	} else {
		fmt.Printf("Error checking ecosystem template 5 path %s: %v\n", ecosystemTemplatePath5, statErr)
		return
	}

	ecosystemTemplatePath6 := filepath.Join(cwd, "internal/acctest/integration_tests/nios_ecosystem_templates", "Version5_REST_API_Session_Template.json")

	if _, statErr := os.Stat(ecosystemTemplatePath6); statErr == nil {
		err = ConfigureEcoSystemTemplates(host, wapiVer, username, password, ecosystemTemplatePath6)
		if err != nil {
			fmt.Printf("Error uploading ecosystem templates: %v\n", err)
			return
		}
		fmt.Println("Ecosystem template 6 uploaded successfully")
	} else if os.IsNotExist(statErr) {
		fmt.Printf("Ecosystem template 6 not found at %s, skipping upload\n", ecosystemTemplatePath6)
	} else {
		fmt.Printf("Error checking ecosystem template 6 path %s: %v\n", ecosystemTemplatePath6, statErr)
		return
	}

	ecosystemTemplatePath7 := filepath.Join(cwd, "internal/acctest/integration_tests/nios_ecosystem_templates", "event_dhcp_lease_template.json")

	if _, statErr := os.Stat(ecosystemTemplatePath7); statErr == nil {
		err = ConfigureEcoSystemTemplates(host, wapiVer, username, password, ecosystemTemplatePath7)
		if err != nil {
			fmt.Printf("Error uploading ecosystem templates: %v\n", err)
			return
		}
		fmt.Println("Ecosystem template 7 uploaded successfully")
	} else if os.IsNotExist(statErr) {
		fmt.Printf("Ecosystem template 7 not found at %s, skipping upload\n", ecosystemTemplatePath7)
	} else {
		fmt.Printf("Error checking ecosystem template 7 path %s: %v\n", ecosystemTemplatePath7, statErr)
		return
	}

	ecosystemTemplatePath8 := filepath.Join(cwd, "internal/acctest/integration_tests/nios_ecosystem_templates", "Version5_DNS_Zone_and_Records.json")

	if _, statErr := os.Stat(ecosystemTemplatePath8); statErr == nil {
		err = ConfigureEcoSystemTemplates(host, wapiVer, username, password, ecosystemTemplatePath8)
		if err != nil {
			fmt.Printf("Error uploading ecosystem templates: %v\n", err)
			return
		}
		fmt.Println("Ecosystem template 8 uploaded successfully")
	} else if os.IsNotExist(statErr) {
		fmt.Printf("Ecosystem template 8 not found at %s, skipping upload\n", ecosystemTemplatePath8)
	} else {
		fmt.Printf("Error checking ecosystem template 8 path %s: %v\n", ecosystemTemplatePath8, statErr)
		return
	}

	ecosystemTemplatePath9 := filepath.Join(cwd, "internal/acctest/integration_tests/nios_ecosystem_templates", "IPAM_PxgridEvent.json")

	if _, statErr := os.Stat(ecosystemTemplatePath9); statErr == nil {
		err = ConfigureEcoSystemTemplates(host, wapiVer, username, password, ecosystemTemplatePath9)
		if err != nil {
			fmt.Printf("Error uploading ecosystem templates: %v\n", err)
			return
		}
		fmt.Println("Ecosystem template 9 uploaded successfully")
	} else if os.IsNotExist(statErr) {
		fmt.Printf("Ecosystem template 9 not found at %s, skipping upload\n", ecosystemTemplatePath9)
	} else {
		fmt.Printf("Error checking ecosystem template 9 path %s: %v\n", ecosystemTemplatePath9, statErr)
		return
	}

	ecosystemTemplatePath10 := filepath.Join(cwd, "internal/acctest/integration_tests/nios_ecosystem_templates", "Version5_Syslog_Session_Template1.json")

	if _, statErr := os.Stat(ecosystemTemplatePath10); statErr == nil {
		err = ConfigureEcoSystemTemplates(host, wapiVer, username, password, ecosystemTemplatePath10)
		if err != nil {
			fmt.Printf("Error uploading ecosystem templates: %v\n", err)
			return
		}
		fmt.Println("Ecosystem template 10 uploaded successfully")
	} else if os.IsNotExist(statErr) {
		fmt.Printf("Ecosystem template 10 not found at %s, skipping upload\n", ecosystemTemplatePath10)
	} else {
		fmt.Printf("Error checking ecosystem template 10 path %s: %v\n", ecosystemTemplatePath10, statErr)
		return
	}

	ecosystemTemplatePath11 := filepath.Join(cwd, "internal/acctest/integration_tests/nios_ecosystem_templates", "Version5_DXL_Session_Template1.json")

	if _, statErr := os.Stat(ecosystemTemplatePath11); statErr == nil {
		err = ConfigureEcoSystemTemplates(host, wapiVer, username, password, ecosystemTemplatePath11)
		if err != nil {
			fmt.Printf("Error uploading ecosystem templates: %v\n", err)
			return
		}
		fmt.Println("Ecosystem template 11 uploaded successfully")
	} else if os.IsNotExist(statErr) {
		fmt.Printf("Ecosystem template 11 not found at %s, skipping upload\n", ecosystemTemplatePath11)
	} else {
		fmt.Printf("Error checking ecosystem template 11 path %s: %v\n", ecosystemTemplatePath11, statErr)
		return
	}

	ecosystemTemplatePath12 := filepath.Join(cwd, "internal/acctest/integration_tests/nios_ecosystem_templates", "Version5_REST_API_Session_Template1.json")

	if _, statErr := os.Stat(ecosystemTemplatePath12); statErr == nil {
		err = ConfigureEcoSystemTemplates(host, wapiVer, username, password, ecosystemTemplatePath12)
		if err != nil {
			fmt.Printf("Error uploading ecosystem templates: %v\n", err)
			return
		}
		fmt.Println("Ecosystem template 12 uploaded successfully")
	} else if os.IsNotExist(statErr) {
		fmt.Printf("Ecosystem template 12 not found at %s, skipping upload\n", ecosystemTemplatePath12)
	} else {
		fmt.Printf("Error checking ecosystem template 12 path %s: %v\n", ecosystemTemplatePath12, statErr)
		return
	}

	ecosystemTemplatePath13 := filepath.Join(cwd, "internal/acctest/integration_tests/nios_ecosystem_templates", "event_template_schema_default.json")

	if _, statErr := os.Stat(ecosystemTemplatePath13); statErr == nil {
		err = ConfigureEcoSystemTemplates(host, wapiVer, username, password, ecosystemTemplatePath13)
		if err != nil {
			fmt.Printf("Error uploading ecosystem templates: %v\n", err)
			return
		}
		fmt.Println("Ecosystem template 13 uploaded successfully")
	} else if os.IsNotExist(statErr) {
		fmt.Printf("Ecosystem template 13 not found at %s, skipping upload\n", ecosystemTemplatePath13)
	} else {
		fmt.Printf("Error checking ecosystem template 13 path %s: %v\n", ecosystemTemplatePath13, statErr)
		return
	}

	clients := PreConfigClients{
		IPAM:         apiClient.IPAMAPI,
		DHCP:         apiClient.DHCPAPI,
		DNS:          apiClient.DNSAPI,
		GRID:         apiClient.GridAPI,
		MISC:         apiClient.MiscAPI,
		MICROSOFT:    apiClient.MicrosoftAPI,
		NOTIFICATION: apiClient.NotificationAPI,
		PARENTAL:     apiClient.ParentalControlAPI,
		SECURITY:     apiClient.SecurityAPI,
		RIR:          apiClient.RIRAPI,
		DISCOVERY:    apiClient.DiscoveryAPI,
	}

	err = PreConfig(clients, hostnames)
	if err != nil {
		fmt.Printf("Error during pre-configuration: %v\n", err)
		return
	}
	fmt.Println("Pre-configuration completed successfully")

	pxgridCertUploadFilePath := filepath.Join(cwd, "internal/acctest/integration_tests/nios_notification_rest_endpoint", "dummy-bundle.pem")
	err = ConfigurePxgridEndpoint(host, wapiVer, username, password, pxgridCertUploadFilePath, apiClient.MiscAPI)
	if err != nil {
		fmt.Printf("Error configuring pxGrid endpoint: %v\n", err)
		return
	}

	fmt.Println("PXGRID endpoint configured successfully")

	fmt.Printf("Environment setup complete. Variables written to %s\n", pipelineEnvPath)

}
