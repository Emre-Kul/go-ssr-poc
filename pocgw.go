package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"fmt"
)

func sendIPC() {
	
}

func healtcheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func main() {
	fmt.Println("App started!");
    router := mux.NewRouter()
	router.HandleFunc("/", healtcheck).Methods("GET")
	http.ListenAndServe(":8080", router)
}
