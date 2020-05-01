package error

type ErrorJSON struct {
	Message string `json:"message"`
}

func NewError(e error) ErrorJSON {
	return ErrorJSON{e.Error()}
}

const (
	InternalServerError = "internal server error"
)
