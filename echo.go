package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"encoding/json"
)
type Answer struct {
	status string
	sentence string
}
func main() {

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", Index)
	router.HandleFunc("/answer", AnswerIndex)
	http.Handle("/", router)
	log.Fatal(http.ListenAndServe(":8080", nil))

}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Nothing to see here!")
}

func AnswerIndex(w http.ResponseWriter, r *http.Request) {
	answer := Answer{status : "ok", sentence: "hello guys"}
	json.NewEncoder(w).Encode(answer)
}