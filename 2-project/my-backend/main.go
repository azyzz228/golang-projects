package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Class struct {
	Id      string   `json:"id"`
	Name    string   `json:"name"`
	Venue   string   `json:"location"`
	Teacher *Teacher `json:"teacher"`
}

type Teacher struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}

var classes []Class

func main() {

	var r *mux.Router = mux.NewRouter()

	classes = append(classes, Class{Id: "2", Name: "STATS101", Venue: "IB2104", Teacher: &Teacher{FirstName: "Lin", LastName: "Jiu"}})

	classes = append(classes, Class{Id: "1", Name: "MATH101", Venue: "AB3101", Teacher: &Teacher{FirstName: "Joe", LastName: "Doe"}})

	r.HandleFunc("/", handelHome).Methods("GET")
	r.HandleFunc("/classes", getClasses).Methods("GET")
	r.HandleFunc("/class/{id}", getClass).Methods("GET")
	r.HandleFunc("/class/{id}", deleteClass).Methods("DELETE")
	r.HandleFunc("/class/{id}", updateClass).Methods("PUT")
	r.HandleFunc("/class/new", addClass).Methods("POST")

	fmt.Println("Starting a new server")

	log.Fatal(http.ListenAndServe(":8080", r))
}

func handelHome(w http.ResponseWriter, r *http.Request) {
	fmt.Println("This is a simple server app with CRUD")
}

func getClasses(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(classes)
}

func getClass(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var params map[string]string = mux.Vars(r)

	for _, item := range classes {
		if item.Id == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}

	http.Error(w, "no class with this id was found", http.StatusNotFound)
}

func deleteClass(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var params map[string]string = mux.Vars(r)

	for index, item := range classes {
		if item.Id == params["id"] {
			classes = append(classes[:index], classes[index+1:]...)
			json.NewEncoder(w).Encode(classes)
			return
		}
	}
}

func updateClass(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var params map[string]string = mux.Vars(r)

	for index, item := range classes {
		if item.Id == params["id"] {
			classes = append(classes[:index], classes[index+1:]...)
		}
	}

	var updatedClass Class
	_ = json.NewDecoder(r.Body).Decode(&updatedClass)
	classes = append(classes, updatedClass)
	json.NewEncoder(w).Encode(updatedClass)

}

func addClass(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var newClass Class
	_ = json.NewDecoder(r.Body).Decode(&newClass)
	newClass.Id = strconv.Itoa(rand.Intn(10000))
	classes = append(classes, newClass)
	json.NewEncoder(w).Encode(newClass)
}
