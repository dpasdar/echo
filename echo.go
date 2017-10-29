package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"encoding/json"
)
type Answer struct {
	Status string
	Sentence string
}
func main() {

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", Index)
	router.HandleFunc("/answer", AnswerIndex).Methods("POST")
	http.Handle("/", router)
	log.Fatal(http.ListenAndServe(":8080", nil))

}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Nothing to see here!")
}

func AnswerIndex(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	answer := Answer{"ok", "hello guys"}
	fmt.Println(answer)
	answerJson,err := json.Marshal(answer)
	fmt.Println(answerJson)
	checkErr(err)
	w.Write(answerJson)
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}