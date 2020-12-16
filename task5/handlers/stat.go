package handlers

import (
	"fmt"
	"github.com/eldario/smap/mapper"
	"github.com/eldario/smap/reader"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func (m *Mappa) Stat(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	number, _ := strconv.Atoi(vars["number"])

	sortedMap := mapper.New()
	sortedMap.SetTopCount(number)
	lineReader := reader.New(sortedMap, 3)

	m.result.Range(func(position, text interface{}) bool {
		lineReader.Read(text.(string), position.(int))
		return true
	})

	for _, word := range sortedMap.GetResults() {
		fmt.Fprintf(writer, "%s - count: %d \n", word.Word, word.Count)
	}
}
