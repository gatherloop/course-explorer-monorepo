package main

import (
	"course-explorer-monorepo/apps/api/content/core/module"
	"course-explorer-monorepo/apps/api/content/core/repository"
	"course-explorer-monorepo/apps/api/content/handler"
	"course-explorer-monorepo/libs/api/middlewares"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("err load env", err)
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

	contentRepo := repository.NewContentRepository()
	contentUsecase := module.NewContentUsecase(contentRepo)
	contentHandler := handler.NewContentHandler(contentUsecase)

  mux.HandleFunc("/courses", contentHandler.GetContentList).Methods("GET")

	muxWithMiddlewares := middlewares.NewCorsMiddleware(mux.ServeHTTP)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("PORT")), muxWithMiddlewares))
}
