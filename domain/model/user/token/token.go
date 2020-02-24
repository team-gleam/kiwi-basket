package token

type Token struct {
	token string
}

func NewToken(t string) Token {
	return Token{t}
}

func (t Token) Token() string {
	return t.token
}
