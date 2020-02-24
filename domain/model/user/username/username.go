package username

import "fmt"

type Username struct {
	name string
}

func NewUsername(u string) (Username, error) {
	if u == "" {
		return Username{}, fmt.Errorf("invalid empty string")
	}
	return Username{u}, nil
}

func (u Username) Name() string {
	return u.name
}
