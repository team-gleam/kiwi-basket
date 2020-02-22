package auth

type IAuthRepository interface {
	Append() error
	Delete() error
	IsExist() bool
}
