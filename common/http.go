package common

import (
	"encoding/json"
	"io"
	"net/http"
)

func JsonResponse(w http.ResponseWriter, statusCode int, body interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(body)
}

func ByteResponse(w http.ResponseWriter, statusCode int, contentType string, body []byte) {
	w.Header().Set("Content-Type", contentType)
	w.WriteHeader(statusCode)
	w.Write(body)
}

func TextResponse(w http.ResponseWriter, statusCode int, body string) {
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(statusCode)
	w.Write([]byte(body))
}

func StreamResponse(w http.ResponseWriter, statusCode int, headers map[string][]string, body io.Reader) {
	if headers["Content-Type"] == nil {
		headers["Content-Type"] = []string{"text/plain"}
	}

	for k, v := range headers {
		for _, vv := range v {
			w.Header().Set(k, vv)
		}
	}

	w.WriteHeader(statusCode)
	io.Copy(w, body)
}
