package login

import (
	"crypto/rand"
	"math/big"
	"testing"
)

func TestHashPassword(t *testing.T) {
	p7, err := genRandomPassword(7)
	if err != nil {
		t.Error(err)
	}
	p72, err := genRandomPassword(72)
	if err != nil {
		t.Error(err)
	}

	ps := []string{
		p7,
		p72,
	}

	for _, p := range ps {
		t.Run("equivalent", func(t *testing.T) {
			h1 := hashPassword(p)
			h2 := hashPassword(p)

			if h1 != h2 {
				t.Errorf("Failed# unexpected output\nraw: %s\n1: %s\n2: %s", p, h1, h2)
			}
		})
	}
}

func genRandomPassword(l int) (string, error) {
	str := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"

	pass := make([]byte, l)

	for i := range pass {
		n, err := rand.Int(rand.Reader, big.NewInt(int64(len(str))))
		if err != nil {
			return "", err
		}

		pass[i] = str[n.Int64()]
	}

	return string(pass), nil
}