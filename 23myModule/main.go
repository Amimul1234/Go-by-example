package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func greeter() {
	fmt.Println("Hey there mod users")
}

func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to golang series on YT</h1>"))
}

func main() {
	fmt.Println("Hello mod in go lang")
	greeter()

	r := mux.NewRouter()
	r.HandleFunc("/", serveHome).Methods("GET")

	log.Fatalln(http.ListenAndServe(":4000", r))
}
