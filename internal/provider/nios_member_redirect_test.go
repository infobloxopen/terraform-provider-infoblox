package provider

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	niosclient "github.com/infobloxopen/infoblox-nios-go-client/client"
	niosoption "github.com/infobloxopen/infoblox-nios-go-client/option"
)

func TestNiosMemberRedirect(t *testing.T) {
	cases := map[string]struct {
		status      int
		contentType string
		body        string
		wantURL     string
		wantMatch   bool
	}{
		"grid member redirect page": {
			status:      http.StatusOK,
			contentType: "text/html",
			body:        `<HTML><HEAD><META HTTP-EQUIV="REFRESH" CONTENT = "0; URL=https://10.0.0.1"/></HEAD><BODY>Click <A HREF="https://10.0.0.1">here</A></BODY></HTML>`,
			wantURL:     "https://10.0.0.1",
			wantMatch:   true,
		},
		"grid master 401 page": {
			status:      http.StatusUnauthorized,
			contentType: "text/html; charset=iso-8859-1",
			body:        `<!DOCTYPE HTML PUBLIC "-//IETF//DTD HTML 2.0//EN"><html><head><title>401 Unauthorized</title></head></html>`,
		},
		"wapi json error": {
			status:      http.StatusBadRequest,
			contentType: "application/json",
			body:        `{"Error": "AdmConProtoError: bad request"}`,
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
				w.Header().Set("Content-Type", tc.contentType)
				w.WriteHeader(tc.status)
				_, _ = w.Write([]byte(tc.body))
			}))
			defer srv.Close()

			client := niosclient.NewAPIClient(
				niosoption.WithNIOSHostUrl(srv.URL),
				niosoption.WithNIOSUsername("u"),
				niosoption.WithNIOSPassword("p"),
			)
			err := checkAndCreatePreRequisitesForNIOS(context.Background(), client)
			if err == nil {
				t.Fatal("expected an error")
			}

			gotURL, ok := niosMemberRedirect(err)
			if ok != tc.wantMatch || gotURL != tc.wantURL {
				t.Fatalf("niosMemberRedirect() = (%q, %v), want (%q, %v); err: %v", gotURL, ok, tc.wantURL, tc.wantMatch, err)
			}
		})
	}
}
