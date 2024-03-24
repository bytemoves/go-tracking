package main

import (
	
	
	"log"
	"net/http"
	"github.com/bytemoves/tracking-service/handler"
)

func main() {
	//http server

	server := http.Server{
		Addr: ":8000",
		Handler:  handler.NewHandler(),
	}

	//runserver
	log.Printf("Starting http server listening at %q",server.Addr)
	if err := server.ListenAndServe(); err != nil {
		log.Printf("%v",err)

	} else {
		log.Printf("Server closed")
	}
}



