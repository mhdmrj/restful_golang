package main

import (
    "fmt"
    "log"
    "net/http"
		"encoding/json"
		"github.com/gorilla/mux"
)


type Kontak struct {
	Id      string    `json:"Id"`
	Name  string `json:"Name"`
	Gender    string `json:"gend"`
	Phone string `json:"phone"`
	Email string `json:"Email"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	
}

type Kontaks []Kontak

func lihatKontak(w http.ResponseWriter, r *http.Request) {
	kontaks := Kontaks{
		Kontak{Id:"1a5071bd-2960-4829-8adc-593e216b2de5", Name:"fulan", Gender:"male", Phone:"628123456789", Email:"fulan@email.com", CreatedAt: "2021-08-18T08:12:48+07:00", UpdatedAt:"2021-08-18T08:12:48+07:00"},
	}
	fmt.Println("Endpoint Hit: Nama Kontak")
	json.NewEncoder(w).Encode(kontaks)
}

func testPOSTKontak(w http.ResponseWriter, r *http.Request){
	fmt.Println("Test POST Endpoint Work")
}

func homePage(w http.ResponseWriter, r *http.Request){
    fmt.Fprintf(w, "Welcome to the Simple CRUD Golang!")
    fmt.Println("Endpoint Hit: homePage")
}

func handleRequests() {
		myRouter := mux.NewRouter().StrictSlash(true)
    myRouter.HandleFunc("/", homePage)
		myRouter.HandleFunc("/kontak", lihatKontak).Methods("GET")
		myRouter.HandleFunc("/kontak", testPOSTKontak).Methods("POST")		
    log.Fatal(http.ListenAndServe(":8081", myRouter))
}

func main(){	
    handleRequests()
}