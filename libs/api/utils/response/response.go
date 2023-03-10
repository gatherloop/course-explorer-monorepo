package response

import (
	"encoding/json"
	"net/http"

	domain "github.com/gatherloop/course-explorer-monorepo"

)

func Error(w http.ResponseWriter, httpCode int, data interface{}, message string) {
	var errorResponse contract.ErrorResponse
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
