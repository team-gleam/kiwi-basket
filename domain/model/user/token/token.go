package token

import (
	"crypto/rand"
	"math/big"
)

type Token struct {
	token string
}

func NewToken(t string) Token {
	return Token{t}
}

func GenToken() (Token, error) {
	str := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"

	length := 32
	token := make([]byte, length)

	for i := range token {
		n, err := rand.Int(rand.Reader, big.NewInt(int64(length)))
		if err != nil {
			return Token{}, err
		}

		token[i] = str[n.Int64()]
	}

	return NewToken(string(token)), nil
}

func (t Token) Token() string {
	return t.token
}
