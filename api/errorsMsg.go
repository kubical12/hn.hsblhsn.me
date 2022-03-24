package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"go.uber.org/zap"
)

// ErrorMsg represents an JSON error message.
type ErrorMsg struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
}

// String implements fmt.Stringer interface.
func (m ErrorMsg) String() string {
	m.Error = true
	b, err := json.Marshal(m)
	if err != nil {
		panic(err)
	}
	return string(b)
}

// HTTPError returns an error message as a JSON response.
func HTTPError(w http.ResponseWriter, err error, code int, msg string) {
	if err != nil {
		zap.L().Error(
			"error occurred in handler",
			zap.Error(err),
			zap.Int("code", code),
			zap.String("msg", msg),
		)
	}
	w.WriteHeader(code)
	fmt.Fprintln(w, ErrorMsg{Message: msg})
}
