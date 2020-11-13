package username

import (
	"testing"
)

func TestNewUsername(t *testing.T) {
	tts := []struct {
		name       string
		input      string
		expected   Username
		shouldFail bool
	}{
		{"success", "user", Username{"user"}, false},
		{"empty string", "", Username{}, true},
	}

	for _, tt := range tts {
		t.Run(tt.name, func(t *testing.T) {
			v, e := NewUsername(tt.input)
			if !tt.shouldFail && e != nil {
				t.Fatalf("unexpected error: %v", e)
			} else if tt.shouldFail && e == nil {
				t.Fatalf("expected error but got nil")
			}

			if tt.expected != v {
				t.Fatalf("expected: %v; got: %v\n", tt.expected, v)
			}
		})
	}
}

func TestUsernameGetter(t *testing.T) {
	username, _ := NewUsername("user")
	if username.name != username.Name() {
		t.Fatalf("expected: %v; got: %v\n", username.name, username.Name())
	}
}
