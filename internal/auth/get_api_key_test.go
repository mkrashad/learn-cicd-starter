package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	tests := []struct {
		name    string
		header  string
		wantKey string
		wantErr bool
	}{
		{
			name:    "valid api key",
			header:  "ApiKey abc123",
			wantKey: "abc123",
			wantErr: false,
		},
		{
			name:    "missing authorization header",
			header:  "",
			wantKey: "",
			wantErr: true,
		},
		{
			name:    "wrong authorization scheme",
			header:  "Bearer abc123",
			wantKey: "",
			wantErr: true,
		},
		{
			name:    "missing api key",
			header:  "ApiKey",
			wantKey: "",
			wantErr: true,
		},
		{
			name:    "extra values",
			header:  "ApiKey abc123 extra",
			wantKey: "abc123",
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			headers := http.Header{}
			if tt.header != "" {
				headers.Set("Authorization", tt.header)
			}

			got, err := GetAPIKey(headers)

			if tt.wantErr {
				if err == nil {
					t.Fatal("expected an error")
				}
				//return
			}

			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			if got != tt.wantKey {
				t.Fatalf("expected %q, got %q", tt.wantKey, got)
			}
		})
	}
}