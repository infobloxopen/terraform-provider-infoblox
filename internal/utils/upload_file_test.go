package utils

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

func TestUploadFileWithToken(t *testing.T) {
	var uploaded bool
	mux := http.NewServeMux()
	var srv *httptest.Server
	mux.HandleFunc("/wapi/"+WapiVersion+"/fileop", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Query().Get("_function") != "uploadinit" || r.Method != http.MethodPost {
			http.Error(w, "unexpected request", http.StatusBadRequest)
			return
		}
		_ = json.NewEncoder(w).Encode(uploadInitResponse{Token: "tok-123", URL: srv.URL + "/upload"})
	})
	mux.HandleFunc("/upload", func(w http.ResponseWriter, r *http.Request) {
		if !strings.HasPrefix(r.Header.Get("Content-Type"), "multipart/form-data") {
			http.Error(w, "expected multipart", http.StatusBadRequest)
			return
		}
		f, _, err := r.FormFile("file")
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		_ = f.Close()
		uploaded = true
	})
	// A TLS server whose certificate is only trusted by srv.Client(): the upload must use the
	// supplied client (and its TLS settings) rather than building its own.
	srv = httptest.NewTLSServer(mux)
	defer srv.Close()

	file := t.TempDir() + "/cert.pem"
	if err := writeFile(file, "dummy"); err != nil {
		t.Fatal(err)
	}

	token, err := UploadFileWithToken(context.Background(), srv.Client(), srv.URL+"/", false, file)
	if err != nil {
		t.Fatalf("upload failed: %v", err)
	}
	if token != "tok-123" || !uploaded {
		t.Fatalf("got token %q uploaded=%v", token, uploaded)
	}

	// The default client does not trust the test certificate, proving TLS verification is not bypassed.
	if _, err := UploadFileWithToken(context.Background(), &http.Client{}, srv.URL, false, file); err == nil {
		t.Fatal("expected a TLS verification error with an untrusting client")
	}
}

func TestUploadFileWithTokenRejectsPassthrough(t *testing.T) {
	_, err := UploadFileWithToken(context.Background(), http.DefaultClient, "https://portal.example.com", true, "unused")
	if err == nil || !strings.Contains(err.Error(), "passthrough") {
		t.Fatalf("expected passthrough error, got %v", err)
	}
}

func writeFile(path, content string) error {
	return os.WriteFile(path, []byte(content), 0o600)
}
