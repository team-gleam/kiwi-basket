package user

type User struct {
	username Username
}

func NewUser(u Username) *User {
	return &User{u}
}

func (u *User) Username() Username {
	return u.username
}
