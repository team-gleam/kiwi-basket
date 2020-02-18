package user

import (
	"testing"
)

func TestNewUsername(t *testing.T) {
	tts := []struct {
		name        string
		input       string
		expextValue Username
		expectError bool
	}{
		{"success", "user", Username{"user"}, false},
		{"empty string", "", Username{}, true},
	}

	for _, tt := range tts {
		t.Run(tt.name, func(t *testing.T) {
			v, e := NewUsername(tt.input)
			if !tt.expectError && e != nil {
				t.Fatalf("unexpected error: %v\n", e)
			} else if tt.expectError && e == nil {
				t.Fatalf("expected error but got nil")
			}

			if tt.expextValue != v {
				t.Errorf("Failed# expected: %v; got: %v\n", tt.expextValue, v)
			}
		})
	}
}
