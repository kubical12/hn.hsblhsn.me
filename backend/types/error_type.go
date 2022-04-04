package types

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
)

type encodeableError struct {
	err     error
	message string
}

func NewEncodeableError(err error, msg string) Error {
	return &encodeableError{
		err:     err,
		message: msg,
	}
}

func (e encodeableError) Message() string {
	return e.message
}

func (e encodeableError) Error() string {
	return e.err.Error()
}

type Error interface {
	Message() string
	Error() string
}

func EncodeErrorMessage(ctx context.Context, err error, w http.ResponseWriter) {
	log.Printf("types: encoding error message: %v", err)
	val, ok := err.(Error)
	if !ok {
		val = NewEncodeableError(err, "Internal server error. Please try again.")
	}
	w.WriteHeader(http.StatusInternalServerError)
	if err := json.NewEncoder(w).Encode(map[string]any{
		"error":   true,
		"message": val.Message(),
	}); err != nil {
		panic(err)
	}
}
