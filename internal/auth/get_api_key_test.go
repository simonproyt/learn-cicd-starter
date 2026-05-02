package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	tests := []struct {
		name       string
		headers    http.Header
		wantKey    string
		wantErr    bool
		errMessage string
	}{
		{
			name:    "valid API key header",
			headers: http.Header{"Authorization": []string{"ApiKey abc123"}},
			wantKey: "abc123",
			wantErr: false,
		},
		{
			name:       "missing authorization header",
			headers:    http.Header{},
			wantKey:    "",
			wantErr:    true,
			errMessage: ErrNoAuthHeaderIncluded.Error(),
		},
		{
			name:       "malformed authorization header",
			headers:    http.Header{"Authorization": []string{"Bearer abc123"}},
			wantKey:    "",
			wantErr:    true,
			errMessage: "malformed authorization header",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotKey, err := GetAPIKey(tt.headers)
			if (err != nil) != tt.wantErr {
				t.Fatalf("GetAPIKey() error = %v, wantErr %v", err, tt.wantErr)
			}
			if gotKey != tt.wantKey {
				t.Fatalf("GetAPIKey() = %q, want %q", gotKey, tt.wantKey)
			}
			if err != nil && err.Error() != tt.errMessage {
				t.Fatalf("GetAPIKey() error = %q, want %q", err.Error(), tt.errMessage)
			}
		})
	}
}
