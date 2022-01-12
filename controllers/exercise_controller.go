package controllers

import (
	"TheLazyLemur/openfit/data"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func InitRoutes(router *mux.Router) {
	router.HandleFunc("/api/exercise", returnAllExercises).Methods("GET")
	router.HandleFunc("/api/exercise", updateExercise).Methods("PUT")
	router.HandleFunc("/api/exercise", createNewExercise).Methods("POST")
	router.HandleFunc("/api/exercise/{id}", returnExercise).Methods("GET")
}

func returnAllExercises(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(data.ListExercises())
}

func returnExercise(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]
	intVar, err := strconv.Atoi(key)
	if err != nil {
		log.Fatal("ID need to be an integer")
	}

	json.NewEncoder(w).Encode(data.GetExercise(intVar))
}

func updateExercise(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var exercise data.Exercise
	json.Unmarshal(reqBody, &exercise)
	println("Updating")
}

func createNewExercise(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var exercise data.Exercise
	json.Unmarshal(reqBody, &exercise)

	data.AddExercise(exercise)
}
