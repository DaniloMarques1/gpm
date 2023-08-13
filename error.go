package main

type ApiError struct {
	msg    string
	status int
}

func NewApiError(msg string, status int) *ApiError {
	return &ApiError{msg, status}
}

func (a *ApiError) Error() string {
	return a.msg
}

func (a *ApiError) Status() int {
	return a.status
}
