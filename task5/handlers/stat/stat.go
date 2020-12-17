package stat

import (
	"encoding/json"
	"github.com/eldario/smap/mapper"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

type sortedMap interface {
	Insert([]string, int)
	Remove(string)
	SetTopCount(int)
	GetResults() []mapper.WordItem
}

type Handler struct {
	sortedMap sortedMap
}

func (h *Handler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	number, _ := strconv.Atoi(vars["number"])
	h.sortedMap.SetTopCount(number)

	var jsonData []byte
	result := h.sortedMap.GetResults()
	if len(result) == 0 {
		writer.Write([]byte(`Use endpoint /text firstly`))
		return
	}

	jsonData, err := json.Marshal(result)
	if err != nil {
		log.Println(err)
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	writer.Write(jsonData)
}

func New(mapper sortedMap) *Handler {
	return &Handler{sortedMap: mapper}
}
