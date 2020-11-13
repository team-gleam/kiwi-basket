package token

import "testing"

func TestGenToken(t *testing.T) {
	t.Run("length of token", func(t *testing.T) {
		token, err := GenToken()
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if l := len(token.Token()); l != Length {
			t.Fatalf("expected: %v; got: %v\n", Length, l)
		}
	})
}
