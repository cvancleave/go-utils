package utils

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
)

func SetCorsHeaders(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Methods", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
}

func GetRequest(endpoint string) ([]byte, error) {

	resp, err := http.Get(endpoint)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func PostFormRequest(endpoint string, formData url.Values) ([]byte, error) {

	resp, err := http.PostForm(endpoint, formData)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func PostRequest(endpoint string, data any) ([]byte, error) {

	body, _ := json.Marshal(data)
	resp, err := http.Post(endpoint, "application/json", bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err = io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func RespondJson(w http.ResponseWriter, code int, data any) {
	body, _ := json.Marshal(data)
	w.WriteHeader(code)
	w.Write(body)
}

type ErrorResponse struct {
	ErrorCode    int    `json:"error_code"`
	ErrorMessage string `json:"error_message"`
}

func RespondError(w http.ResponseWriter, code int, message string) {
	response := ErrorResponse{
		ErrorCode:    code,
		ErrorMessage: message,
	}
	body, _ := json.Marshal(response)
	w.WriteHeader(code)
	w.Write(body)
}
