package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	_ "net/http/pprof"
	"tasks/task5/handlers"
)

func main() {
	//go func() {
	//	createPprofServer()
	//}()

	createServer()
}

func createPprofServer() {
	server := &http.Server{Addr: ":6060", Handler: nil}
	defer server.Close()

	log.Fatal(server.ListenAndServe())
}

func createServer() {
	newMap := new(handlers.Mappa)

	r := mux.NewRouter()
	r.HandleFunc("/text", newMap.Text)
	r.HandleFunc("/stat/{number}", newMap.Stat)
	r.HandleFunc("/test", newMap.Test)

	log.Fatal(http.ListenAndServe(":4001", r))
}
