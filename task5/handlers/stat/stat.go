package stat

import (
	"encoding/json"
	"github.com/eldario/smap/mapper"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
	"tasks/task5/counters"
)

type sortedMap interface {
	Insert([]string, int)
	Remove(string)
	SetTopCount(int)
	GetResults() []mapper.WordItem
}

type Handler struct {
	sortedMap sortedMap
	counter   *counters.Metric
}

func (h *Handler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	h.counter.Calls.Inc()
	vars := mux.Vars(request)
	number, _ := strconv.Atoi(vars["number"])
	h.sortedMap.SetTopCount(number)

	var jsonData []byte
	result := h.sortedMap.GetResults()
	if len(result) == 0 {
		if _, err := writer.Write([]byte(`Use endpoint /text firstly`)); err != nil {
			log.Println(err)
		}
		return
	}

	jsonData, err := json.Marshal(result)
	if err != nil {
		log.Println(err)
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	if _, err := writer.Write(jsonData); err != nil {
		log.Println(err)
	}
}

func New(sortedMap sortedMap) *Handler {
	var counter = counters.New("stat", "stat_func_calls")

	return &Handler{sortedMap: sortedMap, counter: counter}
}
