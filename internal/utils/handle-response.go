package utils

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Success bool   `json:"success"`
	Code    int    `json:"code"`
	Message string `json:"message"`
	Payload any    `json:"payload,omitempty"`
}

func HandleResponse(w http.ResponseWriter, statusCode int, message string, payload any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	resp := &Response{
		Success: statusCode >= 200 && statusCode < 300,
		Code:    statusCode,
		Message: message,
		Payload: payload,
	}

	enc := json.NewEncoder(w)

	enc.SetEscapeHTML(false)

	err := enc.Encode(resp)

	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
}
