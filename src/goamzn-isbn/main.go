package main

import (
	"fmt"
	"net/http"

	"goamzn-isbn/handler"

	"github.com/gorilla/mux"
	"github.com/labstack/gommon/log"
)

const port = "8080"

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", handler.Top).Methods("GET")
	{
		router.HandleFunc("/search/isbn", handler.SearchISBN).Methods("GET")
	}
	router.Handle("/", router)
	if err := http.ListenAndServe(fmt.Sprintf(":%s", port), router); err != nil {
		log.Fatal("err: %v", err)
	}
}
