package login

type ILoginRepository interface {
	Create() error
	Delete() error
	IsExist() bool
}
