package main

import (
	"encoding/json"
	"net/http"
)

func JSON(w http.ResponseWriter, body any, status int) {
	w.WriteHeader(status)
	if body != nil {
		json.NewEncoder(w).Encode(body)
	}
}

type ErrorDto struct {
	Message    string `json:"message"`
	StatusCode int    `json:"status_code"`
}

func ERROR(w http.ResponseWriter, err error) {
	var errDto ErrorDto
	switch tErr := err.(type) {
	case *ApiError:
		errDto = ErrorDto{StatusCode: tErr.Status(), Message: tErr.Error()}
	default:
		errDto = ErrorDto{StatusCode: http.StatusInternalServerError, Message: tErr.Error()}
	}

	JSON(w, errDto, errDto.StatusCode)
}
