package controllers

type Error struct {
	Message string
}

// Error構造体にmessageが入ったものを返却。
func NewError(err error) *Error {
	return &Error{
		Message: err.Error(),
	}
}