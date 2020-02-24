package credential

type ICredentialRepository interface {
	Append() error
	Delete() error
	IsExist() bool
}
