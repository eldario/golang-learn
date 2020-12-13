package main

import (
	"encoding/json"
	"fmt"
	"github.com/eldario/smap/mapper"
	"github.com/eldario/smap/reader"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

type textOptions struct {
	Number int
	Text   string
}

var sortedMap = mapper.New()

func text(writer http.ResponseWriter, request *http.Request) {
	var t textOptions

	err := json.NewDecoder(request.Body).Decode(&t)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	lineReader := reader.New(sortedMap, 3)

	lineReader.Read(t.Text, t.Number)
}

func stat(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	number, _ := strconv.Atoi(vars["number"])

	sortedMap.SetTopCount(number)
	for _, word := range sortedMap.GetResults() {
		fmt.Fprintf(writer, "%s - count: %d \n", word.Word, word.Count)
	}
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/text", text)
	r.HandleFunc("/stat/{number}", stat)

	log.Fatal(http.ListenAndServe(":4000", r))
}
