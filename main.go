package main

import (
	"log"
	"net/http"

	"TheLazyLemur/openfit/controllers"
	"TheLazyLemur/openfit/data"

	"github.com/gorilla/mux"
)

func handleRequests() {

	data.LoadById("1")

	myRouter := mux.NewRouter().StrictSlash(true)

	controllers.InitRoutes(myRouter)

	log.Println("Listening on port :8080")
	log.Fatal(http.ListenAndServe(":8080", myRouter))
}

func main() {
	handleRequests()
}
