package util

import (
	"encoding/json"
	"net/http"

	"go-demo/dto"
)

func ResponseErr(response http.ResponseWriter, status int, message string) {
	response.Header().Set("content-type", "application/json")
	response.WriteHeader(status)
	responseData := dto.Response{
		Status:  status,
		Message: message,
	}
	json.NewEncoder(response).Encode(responseData)
}

func ResponseOk(response http.ResponseWriter, status int, message string, data interface{}) {
	response.Header().Set("content-type", "application/json")
	response.WriteHeader(status)
	responseData := dto.Response{
		Status:  status,
		Message: message,
		Data:    data,
	}
	json.NewEncoder(response).Encode(responseData)
}
