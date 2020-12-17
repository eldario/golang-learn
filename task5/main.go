package main

import (
	"github.com/eldario/smap/mapper"
	"github.com/eldario/smap/reader"
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"
	"runtime/pprof"
	"tasks/task5/handlers/stat"
	"tasks/task5/handlers/text"

	"github.com/gorilla/mux"
)

func main() {
	go func() {
		createPprofServer()
	}()

	createServer()
}

func createPprofServer() {
	server := &http.Server{Addr: ":6060", Handler: nil}
	defer server.Close()

	var pProfFileName = "cpu.prof"
	pProfFile, err := os.Create(pProfFileName)
	if err != nil {
		log.Fatal("Cant create cpu prof file: ", err)
	}

	defer pProfFile.Close()

	if err := pprof.StartCPUProfile(pProfFile); err != nil {
		log.Fatal("Cant start profile: ", err)
	}

	defer pprof.StopCPUProfile()

	log.Fatal(server.ListenAndServe())
}

func createServer() {
	sortedMap := mapper.New()
	lineReader := reader.New(sortedMap, 3)

	r := mux.NewRouter()
	r.Handle("/text", text.New(lineReader))
	r.Handle("/stat/{number}", stat.New(sortedMap))
	//r.Handle("/test", textHandler)

	log.Fatal(http.ListenAndServe(":4001", r))
}
