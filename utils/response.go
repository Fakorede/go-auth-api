package utils

import (
	"encoding/json"
	"net/http"
)

type jsonError struct {
	Message string `json:"message"`
}

func WriteJSON(w http.ResponseWriter, status int, data interface{}, wrap string) error {
	wrapper := make(map[string]interface{})

	wrapper[wrap] = data

	js, err := json.Marshal(wrapper)
	if err != nil {
		return err
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(js)

	return nil
}

func ErrorJSON(w http.ResponseWriter, status int, err error) {
	genError := jsonError{
		Message: err.Error(),
	}

	WriteJSON(w, status, genError, "error")
}