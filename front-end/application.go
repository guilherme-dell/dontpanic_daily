package main

import (
	"net/http"
	"log"
)

func main(){
	http.Handle("/", http.FileServer(http.Dir("public")))
	log.Fatal(http.ListenAndServe(":8081", nil))
}