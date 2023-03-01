package response

import (
	"encoding/json"
	"net/http"
)

func WriteSuccess(w http.ResponseWriter, resp HasWrapper) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	return json.NewEncoder(w).Encode(resp.Wrap())
}

func WriteError(w http.ResponseWriter, code int, msg string) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	return json.NewEncoder(w).Encode(Wrapper{
		ErrorCode: code,
		Message:   msg,
		Data:      nil,
	})
}
