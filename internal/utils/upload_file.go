package utils

import (
	"bytes"
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"

	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// uploadInitResponse holds the token and upload URL returned by the NIOS uploadinit API.
type uploadInitResponse struct {
	Token string `json:"token"`
	URL   string `json:"url"`
}

func generateUploadToken(ctx context.Context, baseURL, username, password string) (*uploadInitResponse, error) {
	httpClient := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true}, //nolint:gosec
		},
	}

	uploadInitURL := fmt.Sprintf("%s/wapi/v2.13.6/fileop?_function=uploadinit", baseURL)
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, uploadInitURL, bytes.NewReader([]byte("{}")))
	if err != nil {
		return nil, fmt.Errorf("error creating uploadinit request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.SetBasicAuth(username, password)

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
	tflog.Debug(ctx, fmt.Sprintf("Generated upload token: %s with URL: %s", result.Token, result.URL))
	return &result, nil
}

func uploadFile(ctx context.Context, uploadURL, filePath, username, password string) error {
	httpClient := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true}, //nolint:gosec
		},
	}

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
	req.SetBasicAuth(username, password)

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
func UploadFileWithToken(ctx context.Context, baseURL, filePath, username, password string) (string, error) {
	init, err := generateUploadToken(ctx, baseURL, username, password)
	if err != nil {
		return "", fmt.Errorf("unable to generate upload token: %w", err)
	}
	if err = uploadFile(ctx, init.URL, filePath, username, password); err != nil {
		return "", fmt.Errorf("unable to upload file: %w", err)
	}
	return init.Token, nil
}
