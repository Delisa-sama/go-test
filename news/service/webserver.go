package service

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

func StartWebServer(port int) {

	r := mux.NewRouter()

	for _, route := range routes {
		r.Methods(route.Method).Path(route.Pattern).Name(route.Name).Handler(route.HandlerFunc)
	}
	http.Handle("/", r)

	addr := ":" + strconv.Itoa(port)

	log.Println("Starting HTTP service at " + addr)
	err := http.ListenAndServe(addr, nil)

	if err != nil {
		log.Fatalf("An error occured starting HTTP listener [%s]: %s", addr, err)
	}
}
