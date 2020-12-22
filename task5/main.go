package main

import (
	"github.com/eldario/smap/mapper"
	"github.com/eldario/smap/reader"
	"github.com/prometheus/client_golang/prometheus/promhttp"
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
	if err := createPprofServer(); err == nil {
		defer func() {
			pprof.StopCPUProfile()
		}()
	}
	// for go tool pprof http://localhost:6060/debug/pprof/profile?seconds=30
	//go func() {
	//	log.Println(http.ListenAndServe("localhost:6060", nil))
	//}()

	createServer()
}

func createPprofServer() error {
	var (
		cpuProfile = "cpu.prof"
		memProfile = "mem.prof"
	)

	{
		f, err := os.Create(cpuProfile)
		if err != nil {
			log.Fatal("could not create CPU profile: ", err)
		}

		//defer func() {
		//	_ = f.Close() // error handling omitted for example
		//}()

		if err := pprof.StartCPUProfile(f); err != nil {
			log.Println("could not start CPU profile: ", err)
			return err
		}
	}

	{
		f, err := os.Create(memProfile)
		if err != nil {
			log.Fatal("could not create memory profile: ", err)
		}

		//defer func() {
		//	_ = f.Close() // error handling omitted for example
		//}()

		runtime.GC()
		if err := pprof.WriteHeapProfile(f); err != nil {
			log.Println("could not write memory profile: ", err)
			return err
		}
	}

	return nil
}

func createServer() {
	sortedMap := mapper.New()
	lineReader := reader.New(sortedMap, 3)

	r := mux.NewRouter()
	r.Handle("/text", text.New(lineReader))
	r.Handle("/stat/{number}", stat.New(sortedMap))
	r.Handle("/metrics", promhttp.Handler())
	r.Handle("/stop", stop.New())

	log.Fatal(http.ListenAndServe(":4001", r))
}
