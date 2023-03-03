package main

import (
	"bytes"
	"course-explorer-monorepo/libs/api/middlewares"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"io"
	"log"
	"net/http"
	"os"
)

type Response struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type config struct {
	ContentAPIBaseURL string
}

var mapRouting map[string]string

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Println("err load env", err)
	}

	var cfg = getConfigBaseURL()

	mapRouting = initMapRouting(cfg)
}

func main() {
	mux := mux.NewRouter()

	mux.HandleFunc("/hello", handleFunction).Methods("GET")

	mux.HandleFunc("/hello", handleFunction).Methods("POST")

	muxWithMiddlewares := middlewares.NewCorsMiddleware(mux.ServeHTTP)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("PORT")), muxWithMiddlewares))
}

func getConfigBaseURL() config {
	return config{
		ContentAPIBaseURL: os.Getenv("CONTENT_API_BASE_URL"),
	}
}

func initMapRouting(cfg config) map[string]string {
	return map[string]string{
		"/hello": fmt.Sprintf("%s/hello", cfg.ContentAPIBaseURL),
	}
}

func handleFunction(writer http.ResponseWriter, request *http.Request) {
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
