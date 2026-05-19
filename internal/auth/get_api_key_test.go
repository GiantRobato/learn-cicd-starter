package auth

import (
	"errors"
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	tests := []struct {
		name        string
		headers     http.Header
		wantAPIKey  string
		wantErr     bool
		wantErrIs   error
	}{
		{
			name: "valid api key",
			headers: http.Header{
				"Authorization": []string{"ApiKey test-key"},
			},
			wantAPIKey: "test-key",
		},
		{
			name:      "missing authorization header",
			headers:   http.Header{},
			wantErr:   true,
			wantErrIs: ErrNoAuthHeaderIncluded,
		},
		{
			name: "malformed authorization header",
			headers: http.Header{
				"Authorization": []string{"Bearer test-key"},
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotAPIKey, err := GetAPIKey(tt.headers)
			if tt.wantErr {
				if err == nil {
					t.Fatal("expected error, got nil")
				}
				if tt.wantErrIs != nil && !errors.Is(err, tt.wantErrIs) {
					t.Fatalf("expected error %v, got %v", tt.wantErrIs, err)
				}
				return
			}

			if err != nil {
				t.Fatalf("expected no error, got %v", err)
			}
			if gotAPIKey != tt.wantAPIKey {
				t.Fatalf("expected API key %q, got %q", tt.wantAPIKey, gotAPIKey)
			}
		})
	}
}
