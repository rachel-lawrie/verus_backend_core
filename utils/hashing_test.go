package utils_test

import (
	"testing"

	"github.com/rachel-lawrie/verus_backend_core/utils"
)

func TestHashAPIKey(t *testing.T) {
	tests := []struct {
		name     string
		apiKey   string
		expected string
	}{
		{
			name:     "Empty API Key",
			apiKey:   "",
			expected: "e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855",
		},
		{
			name:     "Valid API Key",
			apiKey:   "my-secret-api-key",
			expected: "325ededd6c3b9988f623c7f964abb9b016b76b0f8b3474df0f7d7c23b941381f",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := utils.HashAPIKey(tt.apiKey)
			if got != tt.expected {
				t.Errorf("HashAPIKey() = %v, want %v", got, tt.expected)
			}
		})
	}
}
