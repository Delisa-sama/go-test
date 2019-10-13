package service

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func StartWebServer(port string) {

	r := mux.NewRouter()

	for _, route := range routes {
		r.Methods(route.Method).Path(route.Pattern).Name(route.Name).Handler(route.HandlerFunc)
	}
	http.Handle("/", r)

	log.Println("Starting HTTP service at " + port)
	err := http.ListenAndServe(":"+port, nil)

	if err != nil {
		log.Fatalf("An error occured starting HTTP listener at port %d: %s", port, err)
	}
}
