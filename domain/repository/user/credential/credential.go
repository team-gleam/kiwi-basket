package credential

type ICredentialRepository interface {
	Append() error
	Delete() error
	Exists() bool
}
