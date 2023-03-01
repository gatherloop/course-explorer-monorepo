package main

import (
	"bytes"
	"course-explorer-monorepo/libs/api/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io"
	"log"
	"net/http"
)

type Response struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

var mapRouting = map[string]string{
	"/hello": "http://localhost:3000/hello", // sesuai port dari servicenya
}

func main() {
	mux := mux.NewRouter()

	mux.HandleFunc("/hello", handleFuntion).Methods("GET")

	mux.HandleFunc("/hello", handleFuntion).Methods("POST")

	muxWithMiddlewares := utils.NewCorsMiddleware(mux.ServeHTTP)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", "5000"), muxWithMiddlewares))
}

func handleFuntion(writer http.ResponseWriter, request *http.Request) {
	url := mapRouting[request.RequestURI]
	statusCode := http.StatusInternalServerError

	response, data := newHTTPRequest(request, request.Method, url)

	if response != nil {
		statusCode = response.StatusCode
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(statusCode)
	json.NewEncoder(writer).Encode(data)
}

func newHTTPRequest(request *http.Request, method, url string) (*http.Response, Response) {
	var (
		res            Response
		requestBody    io.Reader
		client         = &http.Client{}
		data, body     map[string]interface{}
		methodWithBody = map[string]bool{
			"POST": true,
			"PUT":  true,
		}
	)

	if methodWithBody[method] {
		err := json.NewDecoder(request.Body).Decode(&body)
		if err != nil {
			res.Message = err.Error()
			return nil, res
		}

		bodyByte, err := json.Marshal(body)
		if err != nil {
			res.Message = err.Error()
			return nil, res
		}

		requestBody = bytes.NewBuffer(bodyByte)
	}

	req, err := http.NewRequest(method, url, requestBody)
	if err != nil {
		res.Message = err.Error()
		return nil, res
	}

	req.Header.Set("Content-Type", "application/json")

	for key, header := range request.Header {
		req.Header.Set(key, header[0])
	}

	response, err := client.Do(req)
	if err != nil {
		res.Message = err.Error()
		return nil, res
	}
	defer response.Body.Close()

	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		res.Message = err.Error()
		return nil, res
	}

	res.Data = data["data"]
	res.Message = data["message"].(string)

	return response, res
}
