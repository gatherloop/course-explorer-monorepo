package response

import (
	"encoding/json"
	"net/http"
	"starter-monorepo/libs/api/domain"
)

func Error(w http.ResponseWriter, httpCode int, data interface{}, message string) {
	var errorResponse domain.ErrorResponse
	errorResponse.Data = data
	errorResponse.Message = message

	returnResponse(w, httpCode, errorResponse)
}

func Success(w http.ResponseWriter, httpCode int, data interface{}) {
	returnResponse(w, httpCode, data)
}

func returnResponse(w http.ResponseWriter, httpCode int, response interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpCode)
	json.NewEncoder(w).Encode(response)
	return
}
