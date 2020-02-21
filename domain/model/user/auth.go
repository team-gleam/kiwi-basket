package user

type Auth struct {
	username Username
	token    string
}

func (a Auth) Token() string {
	return a.token
}
