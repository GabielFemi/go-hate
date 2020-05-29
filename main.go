package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", nil)

	server := http.Server{
		Addr: "127.0.0.1:5000",
		Handler: router,

		ReadTimeout: 15 * time.Second,
		WriteTimeout: 15 * time.Second,
	}
	log.Println("Hate Speech Detector is listening on", server.Addr)
	log.Fatalln(server.ListenAndServe())
}
