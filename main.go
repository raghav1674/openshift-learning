package main

import (
	"fmt"
	"log"
	"net/http"
)



const (
	PORT = 8000
)

var (
	counter int 
)

func main(){

	r := http.NewServeMux()

	r.HandleFunc("/",func(w http.ResponseWriter, req *http.Request){
		fmt.Fprintf(w,"Hello World!!")
		log.Printf("%s %d %s",req.Method,200,req.URL.Path)
	})

	r.HandleFunc("/counter",func(w http.ResponseWriter, req *http.Request){
		fmt.Fprintf(w,"This page is visited %d times",counter)
		counter++
	})

	log.Printf("Server is listening on port %d",PORT)
	http.ListenAndServe(fmt.Sprintf(":%d",PORT),r)
}