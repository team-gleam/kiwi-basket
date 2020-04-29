package error

type Error struct {
	Message string `json:"message"`
}

func NewError(e error) Error {
	return Error{e.Error()}
}
