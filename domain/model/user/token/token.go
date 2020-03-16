package token

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

type Token struct {
	token string
}

func NewToken(t string) Token {
	return Token{t}
}

func GenToken() (Token, error) {
	n, err := rand.Int(rand.Reader, max())
	if err != nil {
		return NewToken(""), err
	}
	return NewToken(fmt.Sprintf("%x", n)), nil
}

func max() *big.Int {
	x := big.NewInt(2)
	y := big.NewInt(80)
	n := new(big.Int).Exp(x, y, nil)
	return n // 0x100000000000000000000
}

func (t Token) Token() string {
	return t.token
}
