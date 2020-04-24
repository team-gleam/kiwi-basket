package username

import "fmt"

const (
	InvalidUsername = "invalid empty string"
)

type Username struct {
	name string
}

func NewUsername(u string) (Username, error) {
	if u == "" {
		return Username{}, fmt.Errorf(InvalidUsername)
	}
	return Username{u}, nil
}

func (u Username) Name() string {
	return u.name
}
