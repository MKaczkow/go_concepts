package utils

import (
	"testing"
)

func TestGetRandomUserAgent(t *testing.T) {
	// Test multiple times to ensure randomness
	seen := make(map[string]bool)
	iterations := 100

	for i := 0; i < iterations; i++ {
		agent := GetRandomUserAgent()

		// Verify returned user agent is not empty
		if agent == "" {
			t.Error("GetRandomUserAgent returned empty string")
		}

		// Track unique user agents seen
		seen[agent] = true
	}

	// Verify we get different user agents (randomness check)
	if len(seen) == 1 {
		t.Error("GetRandomUserAgent always returns same value")
	}
}
