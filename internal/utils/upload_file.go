package utils

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// WapiVersion is the WAPI version the provider is built against. It must match the base path of the
// vendored infoblox-nios-go-client (WAPI 2.13.6, NIOS 9.0.6+), which pins the same version.
const WapiVersion = "v2.13.6"

// uploadInitResponse holds the token and upload URL returned by the NIOS uploadinit API.
type uploadInitResponse struct {
	Token string `json:"token"`
	URL   string `json:"url"`
}

func generateUploadToken(ctx context.Context, httpClient *http.Client, baseURL string) (*uploadInitResponse, error) {
	uploadInitURL := fmt.Sprintf("%s/wapi/%s/fileop?_function=uploadinit", strings.TrimRight(baseURL, "/"), WapiVersion)
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, uploadInitURL, bytes.NewReader([]byte("{}")))
	if err != nil {
		return nil, fmt.Errorf("error creating uploadinit request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")

	tflog.Debug(ctx, fmt.Sprintf("Making uploadinit request to: %s", uploadInitURL))
	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error making uploadinit request: %w", err)
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != http.StatusOK {
		bodyBytes, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("uploadinit failed with status %d: %s", resp.StatusCode, string(bodyBytes))
	}

	var result uploadInitResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("error decoding uploadinit response: %w", err)
	}
	tflog.Debug(ctx, "Upload token generated successfully")
	return &result, nil
}

func uploadFile(ctx context.Context, httpClient *http.Client, uploadURL, filePath string) error {
	file, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("error opening file: %w", err)
	}
	defer func() { _ = file.Close() }()

	var requestBody bytes.Buffer
	writer := multipart.NewWriter(&requestBody)
	part, err := writer.CreateFormFile("file", filepath.Base(filePath))
	if err != nil {
		return fmt.Errorf("error creating form file: %w", err)
	}
	if _, err = io.Copy(part, file); err != nil {
		return fmt.Errorf("error copying file content: %w", err)
	}
	if err = writer.Close(); err != nil {
		return fmt.Errorf("error finalizing multipart form: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, uploadURL, &requestBody)
	if err != nil {
		return fmt.Errorf("error creating upload request: %w", err)
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())

	tflog.Debug(ctx, fmt.Sprintf("Uploading file to: %s", uploadURL))
	resp, err := httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("error uploading file: %w", err)
	}
	defer func() { _ = resp.Body.Close() }()
	if resp.StatusCode != http.StatusOK {
		bodyBytes, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("file upload failed with status %d: %s", resp.StatusCode, string(bodyBytes))
	}
	tflog.Info(ctx, fmt.Sprintf("File %s uploaded successfully", filePath))
	return nil
}

// UploadFileWithToken uploads filePath to the NIOS grid and returns the resulting
// upload token. The token can be used in subsequent WAPI API calls that accept a
// client_certificate_token field.
//
// httpClient must be the NIOS go-client's configured HTTP client, so the upload honours the
// provider's ssl_verify, CA certificate and proxy settings and reuses its authentication.
// Uploads are not supported through the Infoblox Portal passthrough, because the upload URL
// returned by NIOS points at the Grid itself.
func UploadFileWithToken(ctx context.Context, httpClient *http.Client, baseURL string, passthrough bool, filePath string) (string, error) {
	if passthrough {
		return "", errors.New("file uploads are not supported when the provider uses the Infoblox Portal passthrough (enable_nios_passthru); connect to the Grid directly with host_url, or upload the file in NIOS and set client_certificate_token instead")
	}
	if httpClient == nil {
		return "", errors.New("the NIOS client is not configured")
	}
	init, err := generateUploadToken(ctx, httpClient, baseURL)
	if err != nil {
		return "", fmt.Errorf("unable to generate upload token: %w", err)
	}
	if err = uploadFile(ctx, httpClient, init.URL, filePath); err != nil {
		return "", fmt.Errorf("unable to upload file: %w", err)
	}
	return init.Token, nil
}
