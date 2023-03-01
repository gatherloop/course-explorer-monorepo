package main

import (
	"course-explorer-monorepo/libs/api/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	mux := mux.NewRouter()
	mux.HandleFunc("/hello", func(writer http.ResponseWriter, request *http.Request) {
		var response = struct {
			Message string `json:"message"`
			Data    string `json:"data"`
		}{
			Message: "success",
			Data:    "hello",
		}

		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusOK)
		err := json.NewEncoder(writer).Encode(response)
		if err != nil {
			panic(err)
		}

	}).Methods("GET")

	mux.HandleFunc("/hello", func(writer http.ResponseWriter, request *http.Request) {
		type requestBody struct {
			Name string `json:"name"`
		}

		type response struct {
			Message string      `json:"message"`
			Data    requestBody `json:"data"`
		}

		reqBody, err := ioutil.ReadAll(request.Body)
		if err != nil {
			panic(err)
		}

		var body requestBody
		err = json.Unmarshal(reqBody, &body)
		if err != nil {
			panic(err)
		}

		var res = response{
			Message: "success",
			Data:    body,
		}

		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusOK)
		err = json.NewEncoder(writer).Encode(res)
		if err != nil {
			panic(err)
		}

	}).Methods("POST")

	muxWithMiddlewares := utils.NewCorsMiddleware(mux.ServeHTTP)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("SERVER_PORT")), muxWithMiddlewares))
}
