package main

import (
	"github.com/eldario/smap/mapper"
	"github.com/eldario/smap/reader"
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"
	"runtime"
	"runtime/pprof"
	"tasks/task5/handlers/stat"
	"tasks/task5/handlers/stop"
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
	var (
		cpuProfile = "cpu.prof"
		memProfile = "mem.prof"
	)

	{
		f, err := os.Create(cpuProfile)
		if err != nil {
			log.Fatal("could not create CPU profile: ", err)
		}
		defer f.Close() // error handling omitted for example
		if err := pprof.StartCPUProfile(f); err != nil {
			log.Fatal("could not start CPU profile: ", err)
		}
		defer pprof.StopCPUProfile()
	}

	{
		f, err := os.Create(memProfile)
		if err != nil {
			log.Fatal("could not create memory profile: ", err)
		}
		defer f.Close()
		runtime.GC()
		if err := pprof.WriteHeapProfile(f); err != nil {
			log.Fatal("could not write memory profile: ", err)
		}
	}
}

func createServer() {
	sortedMap := mapper.New()
	lineReader := reader.New(sortedMap, 3)

	r := mux.NewRouter()
	r.Handle("/text", text.New(lineReader))
	r.Handle("/stat/{number}", stat.New(sortedMap))
	r.Handle("/stop", stop.New())

	log.Fatal(http.ListenAndServe(":4001", r))
}
