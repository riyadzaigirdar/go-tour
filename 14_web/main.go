package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Person struct{
	Name string `json:"name"`
	Age string `json:"age"`
}

func index(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	person := Person{"riyad", "24"}
	// fmt.Fprintf(w, person)
	data, _ := json.Marshal(person)

	w.Write(data)
}

func main(){
	http.HandleFunc("/", index)
	fmt.Println("Server Starting...")
	http.ListenAndServe(":3000", nil)
}